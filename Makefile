all: bin/main
test: lint unit-test

PLATFORM=local

init-docker:
	@docker-compose -f ./env.compose.yaml up -d  

bin/main:
	@docker build -f ./dev.dockerfile . --target bin \
	--output ./bin/ \
	--platform ${PLATFORM}
	#CGO_ENABLED=1 go build -o bin/main -a -ldflags '-linkmode external -extldflags "-static"' ./cmd 
unit-test:
	@docker build -f ./dev.dockerfile . --target unit-test

unit-test-coverage:
	@docker build -f ./dev.dockerfile . --target unit-test-coverage \
	--output coverage/
	cat coverage/cover.out

lint:
	@docker build -f ./dev.dockerfile .   --target lint

kind-cluster:
	kind create cluster --config ./kind-cluster.yaml --name dev-cluster
	#kubectl apply -f ./db-endpoint.yaml
	kubectl apply -f https://raw.githubusercontent.com/metallb/metallb/v0.13.9/config/manifests/metallb-native.yaml
	kubectl apply -f ./metallb-config.yaml

delete-cluster:
	kind delete cluster --name dev-cluster 
	kubectl delete -f https://raw.githubusercontent.com/metallb/metallb/v0.13.9/config/manifests/metallb-native.yaml
	kubectl delete  -f ./metallb-config.yaml

kind-deploy-rpc:
	echo "Deploy sasayin/zero-rpc:latest as metallb loadbalancer demo\n"
	kubectl delete -f ./kind-deploy-rpc.yaml
	kubectl apply -f ./kind-deploy-rpc.yaml

kind-deploy-api:
	echo "Deploy sasayin/zero-app:latest as metallb loadbalancer demo\n"
	kubectl delete -f ./kind-deploy-api.yaml
	kubectl apply -f ./kind-deploy-api.yaml
	kubectl get svc/zero-rest-service

image: 
	@docker build -t sansayin/zero-app  -f ./rest-docker . 
	@docker build -t sansayin/zero-rpc  -f ./rpc-docker .


docker-run:
	docker run -p8080:8080 --network=microservices sansayin/zero-app

push: image
	docker push sansayin/zero-rpc:latest
	docker push sansayin/zero-app:latest


LB_IP=$(shell kubectl get svc/zero-rest-service -o=jsonpath='{.status.loadBalancer.ingress[0].ip}')

create:
	curl -X POST http://$(LB_IP):8888/api/user \
   -H 'Content-Type: application/json' \
   -d '{"name":"BY"}'

delete:
	curl -X DELETE http://$(LB_IP):8888/users/2 \
   -H 'Content-Type: application/json'

update:
	curl -X PUT http://$(LB_IP):8888/user \
   -H 'Content-Type: application/json'    \
   -d '{"name":"bing"}'

getall:
	curl -X GET http://$(LB_IP):8888/users

stress:
	#ab -p user.json -T application/json -c 100 -n 2000 \
	http://localhost:8888/api/users
	ab -c 100 -n 20000 -m GET http://$(LB_IP):8888/api/users

.ONESHELL:
robor-env:
	cd test
	python3 -m venv venv
	. venv/bin/activate 
	pip3 install -r requirements.txt

robot-test:
	cd test && robot --outputdir results atest
