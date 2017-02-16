package matchmaking

import (
	"flag"
	"fmt"

	"github.com/garyburd/redigo/redis"
)

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

func listenServerChannel() {
	con := pool.Get()
	psc := redis.PubSubConn{Conn: con}
	defer psc.Close()
	psc.Subscribe("ch:server:" + ServerName)
	for {
		switch v := psc.Receive().(type) {
		case redis.Message:
			fmt.Printf("%s: message: %s\n", v.Channel, v.Data)
		case redis.Subscription:
			fmt.Printf("%s: %s %d\n", v.Channel, v.Kind, v.Count)
		case error:
			return
		}
	}
}

func listenGlobalChannel() {
	con := pool.Get()
	psc := redis.PubSubConn{Conn: con}
	defer psc.Close()
	psc.Subscribe("ch:global")
	for {
		switch v := psc.Receive().(type) {
		case redis.Message:
			fmt.Printf("%s: message: %s\n", v.Channel, v.Data)
		case redis.Subscription:
			fmt.Printf("%s: %s %d\n", v.Channel, v.Kind, v.Count)
		case error:
			return
		}
	}
}
