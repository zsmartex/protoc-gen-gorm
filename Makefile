gen: generate
generate: gormpb install

.PHONY: gormpb
gormpb:
	# Generate the standalone module and update/lock dependencies.
	cd proto && buf generate
	find proto/gorm/options -type f -name "*.pb.go" -exec mv {} options \;
	find proto/gorm/types -type f -name "*.pb.go" -exec mv {} types \;

i: install
install:
	go install

# Remove all generated files.
clean:
	go clean
	find -name '*.pb.go' -delete
	$(MAKE) gormpb
