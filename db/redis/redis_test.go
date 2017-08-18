package redis

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func newData(msg, key, version string) (string, error) {
	data := map[string]interface{}{}
	data["message"] = msg
	data["@verison"] = version
	data["key"] = key
	data["@timestamp"] = time.Now().Format("2006-01-02T15:04:05Z")
	v, err := json.Marshal(data)
	if err != nil {
		return "", err
	}
	return string(v), nil
}

func Test(t *testing.T) {
	con, err := newDBRedis("localhost:6379", "db0")
	if err != nil {
		t.Errorf("err:%v", err)
		return
	}
	t.Logf("con:%v", con)
	ch := "logstash-2"
	n := rand.Intn(100)
	for i := 1; i < n; i++ {
		mesg := fmt.Sprintf("go_%03v", i)
		key := fmt.Sprintf("key_%03v", i)
		version := fmt.Sprintf("version_%02v", i)
		v, err := newData(mesg, key, version)
		if err != nil {
			t.Fatal(err)
		}
		r, err := con.Do("PUBLISH", ch, v)
		if err != nil {
			t.Fatal(err)
		}
		if r.(int64) == 0 {
			t.Errorf("no cli sub ch:%v", ch)
			return
		}
		t.Logf("r:%v", r.(int64))
	}
}
