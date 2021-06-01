# run tests for repo module
cd db/repo
go mod tidy
go fmt ./...
go test ./...
cd ..
cd ..

# run tests for grpc module
cd grpc
go mod tidy
go fmt ./...
go test ./...
cd ..

