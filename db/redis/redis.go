package redis

import (
	"github.com/garyburd/redigo/redis"
	//"fmt"
)

func newDBRedis(host, dbname string) (redis.Conn, error) {
	rs, err := redis.Dial("tcp", host)
	if err != nil {
		return nil, err
	}
	//_, _ = rs.Do("SELECT", dbname)
	//fmt.Printf("err:%v rs:%v", err, rs)
	//if err != nil {
	//	return nil, err
	//}
	return rs, nil
}
