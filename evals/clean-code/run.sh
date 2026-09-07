#!/usr/bin/env bash
# Run clean-code agent scenarios headlessly and grade them.
#
#   evals/clean-code/run.sh review [slug...]      default: every review scenario
#   evals/clean-code/run.sh implement [slug...]   default: every implement scenario
#
# Environment:
#   PARALLEL   concurrent agent runs (default 4)
#   MODEL      model alias passed to claude (default: the CLI default)
#   PERMISSIONS  extra permission flags (default: acceptEdits with go, gofmt, git allowed)
#   SCENARIOS  directory holding review/ and implement/ (default: this directory)
#
# Results land in $SCENARIOS/results/<mode>/<slug>/{output.md,grade.json,workspace/}
# and a summary in $SCENARIOS/results/<mode>/summary.md.
set -euo pipefail

here="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
plugin="$(cd "$here/../.." && pwd)"
mode="${1:?usage: run.sh review|implement [slug...]}"
shift
: "${SCENARIOS:=$here}"
results="$SCENARIOS/results/$mode"
grader="$here/grader/bin/grader"
: "${PARALLEL:=4}"
: "${MODEL:=}"
: "${PERMISSIONS:=--permission-mode acceptEdits --allowedTools Bash(go:*),Bash(gofmt:*),Bash(git:*),Bash(ls:*),Bash(cat:*),Bash(head:*),Bash(tail:*),Bash(grep:*),Bash(rg:*),Bash(sed:*),Bash(awk:*),Bash(sort:*),Bash(uniq:*),Bash(wc:*),Bash(cut:*),Bash(tr:*),Bash(echo:*),Bash(printf:*),Bash(tee:*),Bash(xargs:*),Bash(find:*),Bash(mkdir:*),Bash(rm:*),Bash(mv:*),Bash(cp:*),Bash(touch:*),Bash(diff:*),Bash(pwd:*),Bash(test:*),Bash(true:*),Bash(which:*),Bash(env:*),Bash(date:*),Bash(cd:*)}"

review_prompt='Review the uncommitted changes in this repository against the Clean Code standard. Report findings only; do not edit any file.'

build_grader() {
	(cd "$here/grader" && go build -o bin/grader .)
}

scenarios() {
	if [ "$#" -gt 0 ]; then
		printf '%s\n' "$@"
		return
	fi
	case "$mode" in
	review) ls "$SCENARIOS/review" ;;
	implement) ls "$SCENARIOS/implement" | sed 's/\.json$//' ;;
	*) echo "unknown mode $mode" >&2; exit 1 ;;
	esac
}

invoke_agent() {
	local workspace="$1" prompt="$2" output="$3"
	local model_flag=()
	[ -n "$MODEL" ] && model_flag=(--model "$MODEL")
	(cd "$workspace" && claude -p "$prompt" \
		--plugin-dir "$plugin" --agent clean-code --strict-mcp-config --setting-sources project \
		--append-system-prompt-file "$plugin/skills/clean-code/SKILL.md" \
		$PERMISSIONS "${model_flag[@]}" \
		--output-format text <"/dev/null" >"$output" 2>"$output.stderr") || true
}

run_review() {
	local slug="$1" out="$results/$1" workspace
	workspace="$(mktemp -d)"
	mkdir -p "$out"
	(cd "$workspace" && git init -q && git commit -q --allow-empty -m init)
	cp -R "$SCENARIOS/review/$slug/fixture/." "$workspace/"
	printf 'module example.com/fixture\n\ngo 1.26\n' >"$workspace/go.mod"
	(cd "$workspace" && git add -A)
	invoke_agent "$workspace" "$review_prompt" "$out/output.md"
	"$grader" review "$SCENARIOS/review/$slug" "$out/output.md" >"$out/grade.json"
	rm -rf "$workspace"
}

run_implement() {
	local slug="$1" out="$results/$1" workspace task
	workspace="$(mktemp -d)"
	mkdir -p "$out"
	printf 'module example.com/eval\n\ngo 1.26\n' >"$workspace/go.mod"
	"$grader" seed "$SCENARIOS/implement/$slug.json" "$workspace"
	(cd "$workspace" && git init -q && git add -A && git commit -q --allow-empty -m seed)
	task="$(python3 -c 'import json,sys; print(json.load(open(sys.argv[1]))["task"])' "$SCENARIOS/implement/$slug.json")"
	invoke_agent "$workspace" "$task" "$out/output.md"
	rm -rf "$out/workspace"
	cp -R "$workspace" "$out/workspace"
	"$grader" implement "$SCENARIOS/implement/$slug.json" "$out/workspace" >"$out/grade.json"
	rm -rf "$workspace"
}

run_one() {
	local slug="$1"
	echo "[$mode] $slug"
	case "$mode" in
	review) run_review "$slug" ;;
	implement) run_implement "$slug" ;;
	esac
	python3 -c 'import json,sys; g=json.load(open(sys.argv[1])); print("  ", "PASS" if g["pass"] else "FAIL", sys.argv[2])' "$results/$slug/grade.json" "$slug"
}

export -f run_one run_review run_implement invoke_agent
export here plugin mode results grader SCENARIOS MODEL PERMISSIONS review_prompt

build_grader
mkdir -p "$results"
scenarios "$@" | xargs -P "$PARALLEL" -I {} bash -c 'run_one "$@"' _ {}
"$grader" summary "$results" | tee "$results/summary.md"
