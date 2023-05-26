API_FILE := user.api
API_JSON := user.json

gen:
	@goctl api go -api ${API_FILE} -dir . -style gozero

swagger:
	goctl api plugin -plugin goctl-swagger="swagger -filename ${API_JSON}" -api ${API_FILE} -dir .     
	docker run --rm -p 8083:8080 -e SWAGGER_JSON=/foo/${API_JSON} -v ${PWD}:/foo swaggerapi/swagger-ui
run:
	go run main.go
