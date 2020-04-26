# To compile locally:
1) `protoc --proto_path=api/proto/v1 --proto_path=third_party --go_out=plugins=grpc:pkg/api/v1 cernlearn.proto`
2) `cd pkg/cmd/server`
3) `CGO_ENABLED=0 GOOS=linux go build`

To run, build it and exec something like this:\
`sudo docker run -e "SERVER_PORT=1111" -e "DB_HOST=137.138.32.116:31762" -e "OAUTH_SECRET_KEY=bla" -e "OAUTH_CLIENT_ID=cernlearn" -e "OAUTH_REDIRECT_URL=https://cern.ch/test-cernlearn" gitlab-registry.cern.ch/ddolot/cernlearn/cernlearn-server:1.0`