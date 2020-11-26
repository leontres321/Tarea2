FILES = ftp.proto

gen:
	protoc --proto_path=proto --go-grpc_out=. $(FILES)
	protoc --proto_path=proto --go_out=. $(FILES)

clean:
	rm ./pb/*

run:
	go run main.go
