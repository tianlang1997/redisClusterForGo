package main

import (
	"./redisCluster"
	"github.com/go-redis/redis"
	"fmt"
)
var nodesAddr = []redis.Options{
	{Addr:"172.17.0.3:6379",Password:"",DB:0},
	{Addr:"172.17.0.4:6379",Password:"",DB:0},
	{Addr:"172.17.0.5:6379",Password:"",DB:0},
	{Addr:"172.17.0.6:6379",Password:"",DB:0},
	{Addr:"172.17.0.7:6379",Password:"",DB:0},
	{Addr:"172.17.0.8:6379",Password:"",DB:0},

}

func main() {

	clusterClient:=redisCluster.NewRedisCluster(nodesAddr)
	err:=clusterClient.Set("aaa","aaa1",0)
	fmt.Println(err)
	ret,err:=clusterClient.Get("aaa")
	fmt.Println(ret,err)
}
