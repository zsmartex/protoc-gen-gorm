gen: generate
generate: gormpb install

.PHONY: example
example:
	buf generate --template example/user/buf.gen.yaml --path example/user

.PHONY: gormpb
gormpb:
	# Generate the standalone module and update/lock dependencies.
	cd proto && buf generate
	find proto -type f -name "*.pb.go" -exec mv {} gormpb \;

i: install
install:
	go install

# Remove all generated files.
clean:
	go clean
	find -name '*.pb.go' -delete
	$(MAKE) gormpb
