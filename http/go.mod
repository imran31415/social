module social/http

go 1.16

replace social/repo => ../db/repo

replace social/app/grpc => ../grpc

require (
	social/app/grpc v0.0.0-00010101000000-000000000000
	social/repo v0.0.0-00010101000000-000000000000
)
