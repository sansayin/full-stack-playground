![](wordcloud.png#center)
- Generated by [Word Cloud](https://www.wordclouds.com/).

# Build and Run with Docker
## Docker Dev ENV
simply run make
```
make
```

this will build within docker env, and out put executable in bin folder.
Taking longer time to build after using Sqlite DB, pretty fast when docker image was cached.

## Make Docker Image
```
make image
```

## Run with Docker
```
make docker-run
```

# Deploy in KinD(K8s in Docker)
## Install KinD
```
make kind-cluster
```

## Push To Docker Hub 
```
make push
```

## Deploy
```
make kind-deploy 
```
Use k9s to check nodes
![](k9s.png#center)
## Get IP
```
kubectl get svc/rest-service -o=jsonpath='{.status.loadBalancer.ingress[0].ip}'
```
## Simple Test
Replace url with IP in prev step
```
curl -X POST http://$(LB_IP):8081/users \
   -H 'Content-Type: application/json' \
   -d '{"name":"BY"}'

```

# Robotframework Testing

Simple script under test folder
after robot env was setup
run 
```
cd test
python -m venv
source venv/bin/activate
cd ..
make robot-test
```
![](test.png#center)


