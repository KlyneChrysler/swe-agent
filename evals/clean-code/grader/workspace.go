package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

// workspace is every Go file the agent left in a module, parsed with comments.
type workspace struct {
	root  string
	fset  *token.FileSet
	files []sourceFile
}

type sourceFile struct {
	path string
	ast  *ast.File
}

func parseWorkspace(root string) (workspace, error) {
	sources := workspace{root: root, fset: token.NewFileSet()}
	err := filepath.WalkDir(root, func(path string, entry fs.DirEntry, err error) error {
		if err != nil || !isGoSource(entry) {
			return err
		}
		return sources.parse(path)
	})
	if err != nil {
		return workspace{}, fmt.Errorf("walk %s: %w", root, err)
	}
	return sources, nil
}

func isGoSource(entry fs.DirEntry) bool {
	return !entry.IsDir() && strings.HasSuffix(entry.Name(), ".go")
}

func (w *workspace) parse(path string) error {
	parsed, err := parser.ParseFile(w.fset, path, nil, parser.ParseComments)
	if err != nil {
		return fmt.Errorf("parse %s: %w", path, err)
	}
	w.files = append(w.files, sourceFile{path: path, ast: parsed})
	return nil
}

func (w workspace) production() []sourceFile {
	kept := []sourceFile{}
	for _, file := range w.files {
		if !strings.HasSuffix(file.path, "_test.go") {
			kept = append(kept, file)
		}
	}
	return kept
}

func (w workspace) tests() []sourceFile {
	kept := []sourceFile{}
	for _, file := range w.files {
		if strings.HasSuffix(file.path, "_test.go") {
			kept = append(kept, file)
		}
	}
	return kept
}

func (w workspace) functions() []declaredFunc {
	found := []declaredFunc{}
	for _, file := range w.production() {
		found = append(found, file.functions(w.fset)...)
	}
	return found
}

type declaredFunc struct {
	file string
	decl *ast.FuncDecl
	fset *token.FileSet
}

func (f sourceFile) functions(fset *token.FileSet) []declaredFunc {
	found := []declaredFunc{}
	for _, decl := range f.ast.Decls {
		if function, ok := decl.(*ast.FuncDecl); ok {
			found = append(found, declaredFunc{file: f.path, decl: function, fset: fset})
		}
	}
	return found
}

func (f declaredFunc) name() string {
	return f.decl.Name.Name
}

func (f declaredFunc) location() string {
	return fmt.Sprintf("%s:%d %s", filepath.Base(f.file), f.fset.Position(f.decl.Pos()).Line, f.name())
}

func (f declaredFunc) bodyLines() int {
	if f.decl.Body == nil {
		return 0
	}
	first := f.fset.Position(f.decl.Body.Lbrace).Line
	last := f.fset.Position(f.decl.Body.Rbrace).Line
	return max(last-first-1, 0)
}

func (f declaredFunc) params() []*ast.Field {
	return f.decl.Type.Params.List
}

func (f declaredFunc) paramCount() int {
	count := 0
	for _, field := range f.params() {
		count += max(len(field.Names), 1)
	}
	return count
}

func (f declaredFunc) results() []*ast.Field {
	if f.decl.Type.Results == nil {
		return nil
	}
	return f.decl.Type.Results.List
}

func (f declaredFunc) receiverType() string {
	if f.decl.Recv == nil {
		return ""
	}
	return typeName(f.decl.Recv.List[0].Type)
}

func typeName(expr ast.Expr) string {
	switch typed := expr.(type) {
	case *ast.StarExpr:
		return typeName(typed.X)
	case *ast.Ident:
		return typed.Name
	}
	return ""
}

func (w workspace) comments() []commentLine {
	found := []commentLine{}
	for _, file := range w.production() {
		found = append(found, file.commentLines(w.fset)...)
	}
	return found
}

type commentLine struct {
	file string
	line int
	text string
	pos  token.Pos
}

func (f sourceFile) commentLines(fset *token.FileSet) []commentLine {
	found := []commentLine{}
	for _, group := range f.ast.Comments {
		for _, comment := range group.List {
			found = append(found, splitComment(f.path, comment, fset)...)
		}
	}
	return found
}

func splitComment(file string, comment *ast.Comment, fset *token.FileSet) []commentLine {
	start := fset.Position(comment.Pos()).Line
	lines := []commentLine{}
	for offset, text := range strings.Split(stripMarkers(comment.Text), "\n") {
		lines = append(lines, commentLine{file: file, line: start + offset, text: strings.TrimSpace(text), pos: comment.Pos()})
	}
	return lines
}

func stripMarkers(text string) string {
	text = strings.TrimPrefix(text, "//")
	text = strings.TrimPrefix(text, "/*")
	return strings.TrimSuffix(text, "*/")
}

func (c commentLine) location() string {
	return fmt.Sprintf("%s:%d", filepath.Base(c.file), c.line)
}

func (w workspace) goFileCount() int {
	return len(w.production())
}

func (w workspace) exists() bool {
	_, err := os.Stat(w.root)
	return err == nil
}
