package redisCluster

import (
	"fmt"
	"github.com/go-redis/redis"
	"math/rand"
	"strings"
	"time"
)
type RedisNode struct {
	Addr string
	DB int
	Password string
}

type RedisCluster struct {
	nodes map[string]*redis.Client
	nodesAddr []redis.Options
}

func (rc *RedisCluster) initNodes()  {
	for i:= 0;i<len(rc.nodesAddr);i++{
		rc.nodes[rc.nodesAddr[i].Addr] =  redis.NewClient(&rc.nodesAddr[i])
		pong, err :=rc.nodes[rc.nodesAddr[i].Addr].Ping().Result()
		fmt.Println(pong, err)
	}
}

func (rc *RedisCluster)Set(key string,val string,expiration int) error {
	index := rand.Int63n(int64(len(rc.nodesAddr)))
	return rc.nodes[rc.nodesAddr[index].Addr].Set(key, val, time.Duration(expiration)).Err()
}

func (rc *RedisCluster)Get(key string) (interface{},error)  {
	index := rand.Int63n(int64(len(rc.nodesAddr)))
	val , err := rc.nodes[rc.nodesAddr[index].Addr].Get(key).Result()
	if err != nil {
		fmt.Println(err)
		a := strings.Split(err.Error(), " ")
		if a[0]=="MOVED" {
			return rc.nodes[a[2]].Get(key).Result()
		}else{
			return nil,err
		}
	}
	return val ,err
}
func NewRedisCluster(nodesAddr []redis.Options)* RedisCluster {
	rc:=&RedisCluster{
		nodes:make(map[string]*redis.Client),
		nodesAddr:nodesAddr,
	}
	rc.initNodes()
	return  rc
}