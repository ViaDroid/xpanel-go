package init

import (
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/viadroid/xpanel-go/global"
)

func initRedis(conf map[string]any) {
	redis_host := conf["redis_host"]
	redis_port := conf["redis_port"]
	redis_connect_timeout, _ := conf["redis_connect_timeout"].(time.Duration)
	redis_read_timeout, _ := conf["redis_read_timeout"].(time.Duration)
	redis_username := conf["redis_username"]
	redis_password := conf["redis_password"]
	// redis_ssl := conf["redis_ssl"]
	// redis_ssl_context := conf["redis_ssl_context"]

	// cert, err := tls.LoadX509KeyPair("testdata/example-cert.pem", "testdata/example-key.pem")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// cfg := &tls.Config{Certificates: []tls.Certificate{cert}}
	global.Redis = redis.NewClient(&redis.Options{
		Addr:        fmt.Sprintf("%s:%d", redis_host, redis_port),
		DB:          0, // 默认DB 0
		DialTimeout: redis_connect_timeout,
		ReadTimeout: redis_read_timeout,
		Username:    redis_username.(string),
		Password:    redis_password.(string),
		// TLSConfig: cfg,
	})

}
