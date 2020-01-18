#!/bin/bash

# RabbitMQ
docker service rm rabbitmq
#docker build -t ahereza/rabbitmq support/rabbitmq/
#docker push ahereza/rabbitmq:latest
docker service create --name=rabbitmq --replicas=1 --network=my_network -p 1883:1883 -p 5672:5672 -p 15672:15672 ahereza/rabbitmq