FILES = ftp.proto

gen:
	protoc --proto_path=proto --go_out=pb --go_opt=paths=source_relative $(FILES)

clean:
	rm ./pb/*

run:
	go run main.go
