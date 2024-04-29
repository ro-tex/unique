# https://github.com/casey/just

build:
	goreleaser build --clean

build version commitHash:
	go build -ldflags "-X main.version={{version}} -X main.commit={{commitHash}}" .

release version:
	@[[ "{{version}}" =~ ^v[0-9]+\.[0-9]+\.[0-9]+$ ]] # ensure semantic versioning
	@echo "Creating a release for version {{version}}:"
	git tag -a {{version}} -m "{{version}}"
	git push origin {{version}}
	goreleaser release --clean

sign-last-commit:
	git commit --amend --no-edit -S
