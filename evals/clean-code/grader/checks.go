package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os/exec"
	"regexp"
	"strings"
)

type check func(sources workspace, argument string) []string

var checks = map[string]check{
	"tests_exist":            testsExist,
	"tests_pass":             command("go", "test", "./..."),
	"vet_clean":              command("go", "vet", "./..."),
	"gofmt_clean":            gofmtClean,
	"max_func_lines":         maxFuncLines,
	"max_params":             maxParams,
	"max_nesting":            maxNesting,
	"no_bool_params":         noBoolParams,
	"no_nil_return":          noNilReturn,
	"no_comments_in_funcs":   noCommentsInFuncs,
	"no_banner_comments":     noBannerComments,
	"no_commented_code":      noCommentedCode,
	"no_todo":                noTodo,
	"no_smell_names":         noSmellNames,
	"symbol_exists":          symbolExists,
	"symbol_absent":          symbolAbsent,
	"max_files":              maxFiles,
	"no_panic":               noPanic,
	"no_global_mutable":      noGlobalMutable,
	"single_switch":          singleSwitch,
	"no_getter_setter_pairs": noGetterSetterPairs,
}

func testsExist(sources workspace, _ string) []string {
	for _, file := range sources.tests() {
		for _, function := range file.functions(sources.fset) {
			if strings.HasPrefix(function.name(), "Test") {
				return nil
			}
		}
	}
	return []string{"no _test.go file declares a Test function"}
}

func command(name string, args ...string) check {
	return func(sources workspace, _ string) []string {
		run := exec.Command(name, args...)
		run.Dir = sources.root
		output, err := run.CombinedOutput()
		if err != nil {
			return []string{strings.TrimSpace(string(output))}
		}
		return nil
	}
}

func gofmtClean(sources workspace, _ string) []string {
	run := exec.Command("gofmt", "-l", ".")
	run.Dir = sources.root
	output, err := run.CombinedOutput()
	unformatted := strings.TrimSpace(string(output))
	if err != nil || unformatted != "" {
		return []string{"gofmt -l: " + unformatted}
	}
	return nil
}

func maxFuncLines(sources workspace, argument string) []string {
	problems := []string{}
	for _, function := range sources.functions() {
		if lines := function.bodyLines(); lines > limit(argument) {
			problems = append(problems, fmt.Sprintf("%s has %d body lines", function.location(), lines))
		}
	}
	return problems
}

func maxParams(sources workspace, argument string) []string {
	problems := []string{}
	for _, function := range sources.functions() {
		if count := function.paramCount(); count > limit(argument) {
			problems = append(problems, fmt.Sprintf("%s has %d parameters", function.location(), count))
		}
	}
	return problems
}

func maxNesting(sources workspace, argument string) []string {
	problems := []string{}
	for _, function := range sources.functions() {
		if depth := nestingDepth(function.decl.Body); depth > limit(argument) {
			problems = append(problems, fmt.Sprintf("%s nests %d deep", function.location(), depth))
		}
	}
	return problems
}

func nestingDepth(body *ast.BlockStmt) int {
	if body == nil {
		return 0
	}
	deepest := 0
	for _, statement := range body.List {
		deepest = max(deepest, statementDepth(statement))
	}
	return deepest
}

func statementDepth(statement ast.Stmt) int {
	switch typed := statement.(type) {
	case *ast.IfStmt:
		return 1 + max(nestingDepth(typed.Body), elseDepth(typed.Else))
	case *ast.ForStmt:
		return 1 + nestingDepth(typed.Body)
	case *ast.RangeStmt:
		return 1 + nestingDepth(typed.Body)
	case *ast.SwitchStmt:
		return 1 + clausesDepth(typed.Body)
	case *ast.TypeSwitchStmt:
		return 1 + clausesDepth(typed.Body)
	case *ast.SelectStmt:
		return 1 + clausesDepth(typed.Body)
	case *ast.BlockStmt:
		return 1 + nestingDepth(typed)
	}
	return 0
}

func elseDepth(branch ast.Stmt) int {
	if branch == nil {
		return 0
	}
	if block, ok := branch.(*ast.BlockStmt); ok {
		return nestingDepth(block)
	}
	return statementDepth(branch) - 1
}

func clausesDepth(body *ast.BlockStmt) int {
	deepest := 0
	for _, clause := range body.List {
		deepest = max(deepest, nestingDepth(&ast.BlockStmt{List: clauseBody(clause)}))
	}
	return deepest
}

func clauseBody(clause ast.Stmt) []ast.Stmt {
	switch typed := clause.(type) {
	case *ast.CaseClause:
		return typed.Body
	case *ast.CommClause:
		return typed.Body
	}
	return nil
}

func noBoolParams(sources workspace, _ string) []string {
	problems := []string{}
	for _, function := range sources.functions() {
		for _, param := range function.params() {
			if isIdent(param.Type, "bool") {
				problems = append(problems, function.location()+" takes a bool parameter")
			}
		}
	}
	return problems
}

func isIdent(expr ast.Expr, name string) bool {
	ident, ok := expr.(*ast.Ident)
	return ok && ident.Name == name
}

func noNilReturn(sources workspace, _ string) []string {
	problems := []string{}
	for _, function := range sources.functions() {
		if returnsBareNil(function) {
			problems = append(problems, function.location()+" returns nil for a reference result")
		}
	}
	return problems
}

func returnsBareNil(function declaredFunc) bool {
	if function.decl.Body == nil {
		return false
	}
	found := false
	ast.Inspect(function.decl.Body, func(node ast.Node) bool {
		if ret, ok := node.(*ast.ReturnStmt); ok && isBareNilReturn(ret, function.results()) {
			found = true
		}
		return !found
	})
	return found
}

// isBareNilReturn is true when a reference-typed result is the literal nil
// and no other result carries a value, so the caller receives nil and nothing else.
func isBareNilReturn(ret *ast.ReturnStmt, results []*ast.Field) bool {
	types := resultTypes(results)
	if len(ret.Results) != len(types) {
		return false
	}
	nilReference, other := false, false
	for index, expr := range ret.Results {
		if isIdent(expr, "nil") && isReference(types[index]) {
			nilReference = true
		} else if !isIdent(expr, "nil") {
			other = true
		}
	}
	return nilReference && !other
}

func resultTypes(results []*ast.Field) []ast.Expr {
	types := []ast.Expr{}
	for _, field := range results {
		for range max(len(field.Names), 1) {
			types = append(types, field.Type)
		}
	}
	return types
}

func isReference(expr ast.Expr) bool {
	switch typed := expr.(type) {
	case *ast.StarExpr, *ast.MapType, *ast.InterfaceType:
		return true
	case *ast.ArrayType:
		return typed.Len == nil
	}
	return false
}

func noCommentsInFuncs(sources workspace, _ string) []string {
	problems := []string{}
	for _, comment := range sources.comments() {
		if function, inside := enclosingFunction(sources, comment.pos); inside {
			problems = append(problems, fmt.Sprintf("%s comment inside %s", comment.location(), function.name()))
		}
	}
	return problems
}

func enclosingFunction(sources workspace, pos token.Pos) (declaredFunc, bool) {
	for _, function := range sources.functions() {
		body := function.decl.Body
		if body != nil && body.Lbrace < pos && pos < body.Rbrace {
			return function, true
		}
	}
	return declaredFunc{}, false
}

func noBannerComments(sources workspace, _ string) []string {
	problems := []string{}
	for _, comment := range sources.comments() {
		if isBanner(comment.text) {
			problems = append(problems, comment.location()+" banner comment")
		}
	}
	return problems
}

func isBanner(text string) bool {
	if len(text) < 5 {
		return false
	}
	decorative := strings.Count(text, "-") + strings.Count(text, "=") + strings.Count(text, "*") + strings.Count(text, "#")
	return decorative*10 >= len(text)*6
}

func noCommentedCode(sources workspace, _ string) []string {
	problems := []string{}
	for _, comment := range sources.comments() {
		if looksLikeCode(comment.text) {
			problems = append(problems, comment.location()+" commented-out code")
		}
	}
	return problems
}

func looksLikeCode(text string) bool {
	if !strings.ContainsAny(lastChar(text), "{};)") {
		return false
	}
	_, err := parser.ParseFile(token.NewFileSet(), "", "package p\nfunc _() {\n"+text+"\n}", 0)
	return err == nil
}

func lastChar(text string) string {
	if text == "" {
		return ""
	}
	return text[len(text)-1:]
}

var taskNote = regexp.MustCompile(`\b(TODO|FIXME|XXX)\b`)

func noTodo(sources workspace, _ string) []string {
	problems := []string{}
	for _, comment := range sources.comments() {
		if taskNote.MatchString(comment.text) {
			problems = append(problems, comment.location()+" task note")
		}
	}
	return problems
}

var smellSuffixes = []string{"Manager", "Helper", "Util", "Utils", "Processor", "Data", "Info"}

func noSmellNames(sources workspace, _ string) []string {
	problems := []string{}
	for _, file := range sources.production() {
		for _, name := range declaredNames(file.ast) {
			if hasSmellSuffix(name) {
				problems = append(problems, fmt.Sprintf("%s declares %s", file.path, name))
			}
		}
	}
	return problems
}

func hasSmellSuffix(name string) bool {
	for _, suffix := range smellSuffixes {
		if strings.HasSuffix(name, suffix) && name != suffix {
			return true
		}
	}
	return false
}

func declaredNames(file *ast.File) []string {
	names := []string{}
	ast.Inspect(file, func(node ast.Node) bool {
		names = append(names, declaredName(node)...)
		return true
	})
	return names
}

func declaredName(node ast.Node) []string {
	switch typed := node.(type) {
	case *ast.FuncDecl:
		return []string{typed.Name.Name}
	case *ast.TypeSpec:
		return []string{typed.Name.Name}
	case *ast.ValueSpec:
		return identNames(typed.Names)
	case *ast.Field:
		return identNames(typed.Names)
	}
	return nil
}

func identNames(idents []*ast.Ident) []string {
	names := []string{}
	for _, ident := range idents {
		names = append(names, ident.Name)
	}
	return names
}

func symbolExists(sources workspace, name string) []string {
	if hasSymbol(sources, name) {
		return nil
	}
	return []string{"no func, method, or type named " + name}
}

func symbolAbsent(sources workspace, name string) []string {
	if hasSymbol(sources, name) {
		return []string{"a func, method, or type is named " + name}
	}
	return nil
}

func hasSymbol(sources workspace, name string) bool {
	for _, function := range sources.functions() {
		if function.name() == name {
			return true
		}
	}
	return hasType(sources, name)
}

func hasType(sources workspace, name string) bool {
	for _, file := range sources.production() {
		for _, declared := range declaredTypes(file.ast) {
			if declared == name {
				return true
			}
		}
	}
	return false
}

func declaredTypes(file *ast.File) []string {
	names := []string{}
	for _, decl := range file.Decls {
		if generic, ok := decl.(*ast.GenDecl); ok && generic.Tok == token.TYPE {
			names = append(names, specNames(generic)...)
		}
	}
	return names
}

func specNames(generic *ast.GenDecl) []string {
	names := []string{}
	for _, spec := range generic.Specs {
		if typeSpec, ok := spec.(*ast.TypeSpec); ok {
			names = append(names, typeSpec.Name.Name)
		}
	}
	return names
}

func maxFiles(sources workspace, argument string) []string {
	if count := sources.goFileCount(); count > limit(argument) {
		return []string{fmt.Sprintf("%d non-test Go files", count)}
	}
	return nil
}

func noPanic(sources workspace, _ string) []string {
	problems := []string{}
	for _, function := range sources.functions() {
		if callsPanic(function) {
			problems = append(problems, function.location()+" calls panic")
		}
	}
	return problems
}

func callsPanic(function declaredFunc) bool {
	found := false
	ast.Inspect(function.decl, func(node ast.Node) bool {
		if call, ok := node.(*ast.CallExpr); ok && isIdent(call.Fun, "panic") {
			found = true
		}
		return !found
	})
	return found
}

func noGlobalMutable(sources workspace, _ string) []string {
	problems := []string{}
	for _, file := range sources.production() {
		for _, spec := range packageVars(file.ast) {
			if isMutableGlobal(spec) {
				problems = append(problems, fmt.Sprintf("%s declares package var %s", file.path, identNames(spec.Names)))
			}
		}
	}
	return problems
}

func packageVars(file *ast.File) []*ast.ValueSpec {
	specs := []*ast.ValueSpec{}
	for _, decl := range file.Decls {
		if generic, ok := decl.(*ast.GenDecl); ok && generic.Tok == token.VAR {
			specs = append(specs, valueSpecs(generic)...)
		}
	}
	return specs
}

func valueSpecs(generic *ast.GenDecl) []*ast.ValueSpec {
	specs := []*ast.ValueSpec{}
	for _, spec := range generic.Specs {
		if valueSpec, ok := spec.(*ast.ValueSpec); ok {
			specs = append(specs, valueSpec)
		}
	}
	return specs
}

func isMutableGlobal(spec *ast.ValueSpec) bool {
	if spec.Type != nil {
		return !isIdent(spec.Type, "error") && !isFuncType(spec.Type)
	}
	for _, value := range spec.Values {
		if !isErrorConstructor(value) && !isFuncLiteral(value) {
			return true
		}
	}
	return false
}

func isFuncType(expr ast.Expr) bool {
	_, ok := expr.(*ast.FuncType)
	return ok
}

func isFuncLiteral(expr ast.Expr) bool {
	_, ok := expr.(*ast.FuncLit)
	return ok
}

func isErrorConstructor(expr ast.Expr) bool {
	call, ok := expr.(*ast.CallExpr)
	if !ok {
		return false
	}
	selector, ok := call.Fun.(*ast.SelectorExpr)
	return ok && (isIdent(selector.X, "errors") && selector.Sel.Name == "New" ||
		isIdent(selector.X, "fmt") && selector.Sel.Name == "Errorf")
}

func singleSwitch(sources workspace, field string) []string {
	switching := []string{}
	for _, function := range sources.functions() {
		if switchesOn(function, field) {
			switching = append(switching, function.location())
		}
	}
	if len(switching) <= 1 {
		return nil
	}
	return []string{fmt.Sprintf("%s is switched on in %d functions: %s", field, len(switching), strings.Join(switching, "; "))}
}

func switchesOn(function declaredFunc, field string) bool {
	found := false
	ast.Inspect(function.decl, func(node ast.Node) bool {
		if discriminates(node, field) {
			found = true
		}
		return !found
	})
	return found
}

func discriminates(node ast.Node, field string) bool {
	switch typed := node.(type) {
	case *ast.SwitchStmt:
		return typed.Tag != nil && namesField(typed.Tag, field)
	case *ast.BinaryExpr:
		return (typed.Op == token.EQL || typed.Op == token.NEQ) && (namesField(typed.X, field) || namesField(typed.Y, field))
	}
	return false
}

func namesField(expr ast.Expr, field string) bool {
	if selector, ok := expr.(*ast.SelectorExpr); ok {
		return selector.Sel.Name == field
	}
	return isIdent(expr, field)
}

func noGetterSetterPairs(sources workspace, _ string) []string {
	problems := []string{}
	for receiver, methods := range methodsByReceiver(sources) {
		for _, field := range setterFields(methods) {
			if methods["Get"+field] || methods[field] {
				problems = append(problems, fmt.Sprintf("%s has a getter and setter for %s", receiver, field))
			}
		}
	}
	return problems
}

func methodsByReceiver(sources workspace) map[string]map[string]bool {
	grouped := map[string]map[string]bool{}
	for _, function := range sources.functions() {
		receiver := function.receiverType()
		if receiver == "" {
			continue
		}
		if grouped[receiver] == nil {
			grouped[receiver] = map[string]bool{}
		}
		grouped[receiver][function.name()] = true
	}
	return grouped
}

func setterFields(methods map[string]bool) []string {
	fields := []string{}
	for name := range methods {
		if strings.HasPrefix(name, "Set") && len(name) > 3 {
			fields = append(fields, name[3:])
		}
	}
	return fields
}
