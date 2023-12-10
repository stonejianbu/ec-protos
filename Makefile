gen-proto:
	protoc -I./src --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative --go_out=./goout --go-grpc_out=./goout src/*/*.proto