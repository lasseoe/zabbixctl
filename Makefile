BUILD_ROOT	:= $(shell pwd)
BUILD_DATE	:= $(shell git log -1 --format="%cd" --date=short | sed s/-//g)
BUILD_NUM	:= $(shell git rev-list --count HEAD)
BUILD_HASH	:= $(shell git rev-parse --short HEAD)

LDFLAGS		:= "-X main.version=${BUILD_DATE}.${BUILD_NUM}_${BUILD_HASH}-1"
GCFLAGS		:= "-trimpath ${BUILD_ROOT}"

build:
	CGO_ENABLED=0 go build -x -ldflags=${LDFLAGS} -gcflags ${GCFLAGS} .

man:
	@ronn -r man.md

test:
	CGO_ENABLED=0 go test -v ./...

lint:
	CGO_ENABLED=0 go vet ./...
	staticcheck -checks=all ./...

vulncheck:
	govulncheck ./...

test-all:
	test lint vulncheck 

clean:
	@git clean -ffdx
