TEST?=$$(go list ./... | grep -v 'vendor')
HOSTNAME=example.com
NAMESPACE=joerx
NAME=cheesecake
BINARY=terraform-provider-${NAME}
VERSION=0.2
OS_ARCH=darwin_amd64

default: install

build:
	go build -o ${BINARY}

install: build
	mkdir -p ~/.terraform.d/plugins/${HOSTNAME}/${NAMESPACE}/${NAME}/${VERSION}/${OS_ARCH}
	mv ${BINARY} ~/.terraform.d/plugins/${HOSTNAME}/${NAMESPACE}/${NAME}/${VERSION}/${OS_ARCH}

test: 
	go test -i $(TEST) || exit 1                                                   
	echo $(TEST) | xargs -t -n4 go test $(TESTARGS) -timeout=30s -parallel=4

release:
ifndef GPG_FINGERPRINT
	@ echo "GPG_FINGERPRINT must be set"
	@ exit 1
endif
	goreleaser release --rm-dist -p2

testacc: 
	TF_ACC=1 go test $(TEST) -v $(TESTARGS) -timeout 120m
