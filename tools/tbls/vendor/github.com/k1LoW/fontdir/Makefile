export GO111MODULE=on

default: test

ci: depsdev test sec

test:
	go test ./... -coverprofile=coverage.txt -covermode=count

sec:
	gosec ./...

depsdev:
	go get github.com/Songmu/ghch/cmd/ghch
	go get github.com/Songmu/gocredits/cmd/gocredits
	go get github.com/securego/gosec/cmd/gosec

prerelease:
	ghch -w -N ${VER}
	gocredits . > CREDITS
	cat _EXTRA_CREDITS >> CREDITS
	git add CHANGELOG.md CREDITS
	git commit -m'Bump up version number'
	git tag ${VER}

.PHONY: default test
