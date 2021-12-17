all: install-go-tools build

ROOTDIR := $(shell pwd)
OUTPUTDIR = $(ROOTDIR)/_output
DEPSGOBIN = $(OUTPUTDIR)/.bin

.PHONY: mod-download
mod-download:
	go mod download

.PHONY: install-go-tools
install-go-tools: mod-download
	mkdir -p $(DEPSGOBIN)
	GOBIN=$(DEPSGOBIN) go install github.com/golang/protobuf/protoc-gen-go

build:
	mkdir -p $(OUTPUTDIR)
	PATH=$(DEPSGOBIN):$$PATH protoc -I=. -I=./imports/ --go_out=$(OUTPUTDIR) cue.proto
	cp $(OUTPUTDIR)/github.com/solo-io/cue/encoding/protobuf/cue/cue.pb.go .
	rm -rf $(OUTPUTDIR)