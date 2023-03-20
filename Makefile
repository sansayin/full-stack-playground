all: bin/main
test: lint unit-test

PLATFORM=local

.PHONY: ./bin/main
bin/main:
	@docker build -f ./dev.dockerfile . --target bin \
	--output ./bin/ \
	--platform ${PLATFORM}
	#CGO_ENABLED=1 go build -o bin/main -a -ldflags '-linkmode external -extldflags "-static"' ./cmd 
.PHONY: unit-test
unit-test:
	@docker build -f ./dev.dockerfile . --target unit-test

.PHONY: unit-test-coverage
unit-test-coverage:
	@docker build -f ./dev.dockerfile . --target unit-test-coverage \
	--output coverage/
	cat coverage/cover.out

.PHONY: lint
lint:
	@docker build -f ./dev.dockerfile .   --target lint

.PHONY: kind-cluster
kind-cluster:
	kind create cluster --config ./kind-cluster.yaml --name dev-cluster
	kubectl apply -f ./db-endpoint.yaml
	kubectl apply -f https://raw.githubusercontent.com/metallb/metallb/v0.13.7/config/manifests/metallb-native.yaml
	kubectl apply -f ./metallb-config.yaml

.PHONY: delete-cluster
delete-cluster:
	kkind delete cluster --name dev-cluster 

.PHONY: kind-deploy
kind-deploy:
	echo "Deploy sasayin/goapp:latest as metallb loadbalancer demo\n"
	kubectl delete -f ./kind-deploy.yaml
	kubectl apply -f ./kind-deploy.yaml
	kubectl get svc/rest-service

.PHONY: image
image: 
	@docker build -t sansayin/goapp .

.PHONY: docker-run
docker-run:
	docker run -p8081:8081 --network=postgre_default sansayin/goapp

.PHONY: push
push: image
	docker push sansayin/goapp:latest


LB_IP=$(shell kubectl get svc/rest-service -o=jsonpath='{.status.loadBalancer.ingress[0].ip}')

.PHONY: create
create:
	curl -X POST http://$(LB_IP):8081/users \
   -H 'Content-Type: application/json' \
   -d '{"name":"BY"}'

.PHONYN: delete
delete:
	curl -X DELETE http://$(LB_IP):8081/users/2 \
   -H 'Content-Type: application/json'

.PHONYN: update
update:
	curl -X PUT http://$(LB_IP):8081/users/11 \
   -H 'Content-Type: application/json'    \
   -d '{"name":"bing"}'

.PHONY: getall
getall:
	curl -X GET http://$(LB_IP):8081/users


.ONESHELL:
.PHONY: robot-env 
robor-env:
	cd test
	python3 -m venv venv
	. venv/bin/activate 
	pip3 install -r requirements.txt

.PHONY: robot-test
robot-test:
	cd test && robot --outputdir results atest
