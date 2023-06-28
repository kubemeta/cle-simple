INTERNAL_PROTO_FILES=$(shell find . -name *.proto)

.PHONY: config
# generate internal proto
config:
	protoc --proto_path=./config \
 	       --go_out=paths=source_relative:./config \
	       $(INTERNAL_PROTO_FILES)