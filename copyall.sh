#!/bin/bash
export GOOS=linux
export CGO_ENABLED=0

cd accountservice || exit;go build -o accountservice-linux-amd64;echo built $(pwd);cd ..
cd vipservice || exit;go build -o vipservice-linux-amd64;echo built $(pwd);cd ..
cd healthchecker || exit;go build -o healthchecker-linux-amd64;echo built $(pwd);cd ..

export GOOS=darwin

cp healthchecker/healthchecker-linux-amd64 accountservice/
cp healthchecker/healthchecker-linux-amd64 vipservice/


docker build -t ahereza/accountservice accountservice/
docker push ahereza/accountservice:latest
docker service rm accountservice
docker service create --log-driver=gelf --log-opt gelf-address=udp://192.168.99.112:12202 --log-opt gelf-compression-type=none --name=accountservice --replicas=1 --network=my_network -p=6767:6767 ahereza/accountservice

docker build -t ahereza/vipservice vipservice/
docker push ahereza/vipservice:latest
docker service rm vipservice
docker service create --log-driver=gelf --log-opt gelf-address=udp://192.168.99.112:12202 --log-opt gelf-compression-type=none --name=vipservice --replicas=1 --network=my_network -p=6868:6868 ahereza/vipservice
