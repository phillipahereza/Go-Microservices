# shellcheck disable=SC2046
eval $(docker-machine env swarm-manager-1)

export ManagerIP=$(docker-machine ip swarm-manager-1)