compile:
	protoc api/vi/*.proto --go_out=. --go_opt=paths=source_relative --proto_path=.

test:
	go test -race ./...
