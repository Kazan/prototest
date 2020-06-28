.PHONY: protos

protos:
	protoc -I ../prototest_spec/ ../prototest_spec/finance/currency.proto --go-grpc_out=gen/