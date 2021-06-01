module social/app/grpc

go 1.16

replace social/repo => ../db/repo

require (
	github.com/stretchr/testify v1.5.1
	google.golang.org/grpc v1.38.0
	google.golang.org/protobuf v1.25.0
	social/repo v0.0.0-00010101000000-000000000000
)
