all: test

format:
	gofmt -w=true .

test: format
	golint .
	go vet .

readme:
	godoc2md github.com/mickep76/iodatafmt >README.md
