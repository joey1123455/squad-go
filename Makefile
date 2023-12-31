# Run tests
test:
	go test ./...

verbose_test:
	go test -v ./...

# Generate test coverage
test_coverage:
	go test -coverprofile=coverage.out ./...

# Generate coverage report
cover: test_coverage
	go tool cover -func=coverage.out

# Generate HTML coverage report
cover-html: test_coverage
	go tool cover -html=coverage.out -o coverage.html
	open coverage.html

# Remove coverage file
clean-coverage:
	rm coverage.out

