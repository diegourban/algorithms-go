# Same commands as CI. `make check` is the counterpart of `./mvnw verify` in algorithms-java.

COVERAGE_MIN := 90
COVERAGE_FILE := coverage.out

.PHONY: check fmt lint test cover

check: lint cover

fmt:
	golangci-lint fmt

lint:
	golangci-lint run

test:
	go test -race ./...

# Runs the tests with coverage and fails if total statement coverage is below COVERAGE_MIN.
# A module with no statements yet (only the mode line in the profile) passes.
cover:
	go test -race -covermode=atomic -coverprofile=$(COVERAGE_FILE) ./...
	@if [ $$(wc -l < $(COVERAGE_FILE)) -le 1 ]; then \
		echo "coverage: no statements yet"; \
	else \
		total=$$(go tool cover -func=$(COVERAGE_FILE) | awk '/^total:/ { sub("%", "", $$3); print $$3 }'); \
		echo "coverage: $$total% (minimum $(COVERAGE_MIN)%)"; \
		awk -v total="$$total" -v min="$(COVERAGE_MIN)" 'BEGIN { exit (total + 0 < min + 0) }' \
			|| { echo "coverage below $(COVERAGE_MIN)%"; exit 1; }; \
	fi
