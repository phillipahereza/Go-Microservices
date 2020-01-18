#!/bin/bash

# shellcheck disable=SC2164
#cd support/config-server
#././gradlew build
#
#cd ../..
#docker build -t ahereza/configserver support/config-server/
#docker push ahereza/configserver:latest
docker service rm configserver
docker service create --replicas 1 --name configserver -p 8888:8888 --network my_network --update-delay 10s --with-registry-auth --update-parallelism 1 ahereza/configserver
