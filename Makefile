# Settings
# Disable makefile default values and rules
# Unconditionally make all targets.
MAKEFLAGS=--no-builtin-rules --no-builtin-variables --always-make

# Rules
.DEFAULT_GOAL := gen

################################################################################

gen:
	./scripts/gen.sh
