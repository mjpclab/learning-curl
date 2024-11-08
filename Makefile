.PHONY: build dev clean init

build:
	mdbook build
dev:
	mdbook serve -n [::]
clean:
	rm -rf dist/
init:
	cargo install mdbook mdbook-chapter-zero
