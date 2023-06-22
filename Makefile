REPL_PATH ?= cmd/golisp

.PHONY: repl
repl:
	@go run $(REPL_PATH)/main.go
