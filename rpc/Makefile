gen-rpc:
	goctl rpc protoc user.proto --go_out=. --go-grpc_out=. --zrpc_out=.
gen-model:
	goctl model pg datasource --url="postgres://postgres:p4ssw0rd@127.0.0.1:5432/dev?sslmode=disable" -t="zero_user"  -dir="./model" -c
run:
	go run user.go
