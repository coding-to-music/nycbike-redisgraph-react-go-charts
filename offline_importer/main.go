package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/gomodule/redigo/redis"
	"github.com/mitchsw/citibike-journeys/offline_importer/importer"
)

func main() {
	redisAddress := flag.String("redis", "172.18.0.2:6379", "host:port address of Redis")
	resetGraph := flag.Bool("reset_graph", false, "Reset graph before importing. Should be true for first import")
	log.SetOutput(os.Stdout)
	pool := &redis.Pool{
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", *redisAddress)
		},
	}
	defer pool.Close()

	imp, err := importer.NewImporter(pool, 1, 10000)
	if err != nil {
		panic(err)
	}
	err = imp.Run(*resetGraph)
	if err != nil {
		panic(err)
	}

	fmt.Println("Done!")
}
