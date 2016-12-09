
test:
	govendor test +local -race -cover

cover: coverage.out
	go tool cover -html=$<

coverage.out: $(wildcard *.go)
	govendor test +local -race -cover -coverprofile=$@


.PHONY: test cover
.SILENT:
