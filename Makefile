TEST_CMD=go test ./... -v

test:
	$(TEST_CMD)

test-uncached:
	$(TEST_CMD) -count=1