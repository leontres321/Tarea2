FILES = ftp.proto

gen:
	protoc --proto_path=proto --go-grpc_out=. $(FILES)
	protoc --proto_path=proto --go_out=. $(FILES)

clean:
	rm ./pb/ftp.pb.go ./pb/ftp_grpc.pb.go

run:
	go run main.go
