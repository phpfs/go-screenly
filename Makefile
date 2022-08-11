setup:
	go mod download
	go install github.com/go-swagger/go-swagger/cmd/swagger@latest

generate: setup clean
	swagger generate client --target=./screenly

clean:
	rm -rf screenly/client screenly/models
