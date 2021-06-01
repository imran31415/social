## Social Network

Build Protos
```bash
protoc -I protos/ protos/social.proto --go_out=plugins=grpc:protos
protoc --proto_path=protos --js_out=import_style=commonjs,binary:../http/frontend/src/static/js --grpc-web_out=import_style=commonjs,mode=grpcwebtext:../http/frontend/src/static/js protos/social.proto 
protoc --doc_out=protos  protos/social.proto 
```

Run with docker compose:
`docker-compose up --build`

## Documentation:
See [blog/BLOG.md](blog/BLOG.md) for dev process