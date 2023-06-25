REPL_PATH  ?= cmd/golisp
TEST_FLAGS ?=

.PHONY: repl
repl:
	@go run $(REPL_PATH)/main.go

.PHONY: test
test:
	@go test $(TEST_FLAGS) golisp/evaluator
