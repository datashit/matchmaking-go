package matchmaking

import "github.com/garyburd/redigo/redis"
import "flag"

var redisServer = flag.String("redisServer", "192.168.99.100:6379", "Redis server adress")

var pool = redis.Pool{MaxIdle: 20, MaxActive: 100, Dial: func() (redis.Conn, error) {
	return redis.Dial("tcp", *redisServer)
}}

var c = pool.Get()

func getKey(key string) string {
	v, err := redis.String(c.Do("GET", key))
	if err != nil {

	}
	return v

}
