#!/bin/bash
export GOOS=linux
export CGO_ENABLED=0

cd accountservice || exit;go build -o accountservice-linux-amd64;echo built $(pwd);cd ..

export GOOS=darwin

docker build -t ahereza/accountservice accountservice/

docker service rm accountservice
docker service create --name=accountservice --replicas=1 --network=my_network -p=6767:6767 someprefix/accountservice