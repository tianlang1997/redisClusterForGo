# redisClusterForGo
use the go to connect to redis cluster

use docker to create redis cluster 
#!/bin/bash
for ((i=0; i<6; i++))
do
docker run --name redis700$i -p 700$i:6379 -d redis redis-server --appendonly yes --protected-mode no --cluster-enabled yes
done

find ip of redis 
echo "">clusterIP
for ((i=0; i<6; i++))
do
docker inspect redis700$i | grep IPAddress | egrep -o "([0-9]{1,3}.){3}[0-9]{1,3}" >> clusterIP
done
"redis-cli --cluster create 172.17.0.3:6379 172.17.0.4:6379 172.17.0.5:6379 172.17.0.6:6379 172.17.0.7:6379 172.17.0.8:6379 --cluster-replicas 1" >> clusterIP
echo "redis-cli --cluster create 172.17.0.3:6379 172.17.0.4:6379 172.17.0.5:6379 172.17.0.6:6379 172.17.0.7:6379 172.17.0.8:6379 --cluster-replicas 1"

buid go environment in docker
FROM golang:alpine AS development
WORKDIR $GOPATH/src

run 
docker run -it -v your go workspace path:/go/src go-development:v0.1
go run main.go
