all: build

build:
	mkdir -p ./_output
	protoc -I=. -I=./imports/ --go_out=./_output cue.proto
	cp ./_output/cuelang.org/go/encoding/protobuf/cue/cue.pb.go .
	rm -rf ./_output