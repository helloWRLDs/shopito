# Shopito

## Code Gens
- protoc

    -   ```
        export GOPATH=$$HOME/go
        export PATH=$$PATH:$$GOPATH/bin
        sudo apt install -y protobuf-compiler
        go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
        go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
        ```
    -   ```
        find ./services -name "*.proto" -print0 | xargs -0 protoc --go_out=./pkg/protobuf --go-grpc_out=./pkg/protobuf
        ```
- sqlc
    -   ```
        docker run --rm -v $(pwd):/src -w /src sqlc/sqlc generate
        ```