# Settings
# Disable makefile default values and rules
# Unconditionally make all targets.
MAKEFLAGS=--no-builtin-rules --no-builtin-variables --always-make

# Rules
.DEFAULT_GOAL := gen

################################################################################

# プロバイダのスキーマが大きいためメモリ上限を引き上げないと out of memory になる
gen:
	NODE_OPTIONS="--max-old-space-size=4096" npx cdktf-cli get
	go mod tidy
