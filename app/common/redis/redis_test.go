package redis

import (
	"github.com/go-home-admin/home/bootstrap/providers"
	"testing"
	"time"
)

func TestGetString(t *testing.T) {
	providers.NewRedisProvider()

	k, v := "test", "test1"
	Set(k, v, 3600*time.Second)

	nv := GetString(k)

	if nv != v {
		t.Fatal("无法设置redis缓存")
	}
}

func TestExists(t *testing.T) {
	providers.NewRedisProvider()

	k, v := "test", "test1"
	Set(k, v, 3600*time.Second)

	nv := Exists(k)

	if nv != true {
		t.Fatal("Exists 判断错误1")
	}

	Del(k)

	nv = Exists(k)

	if nv != false {
		t.Fatal("Exists 判断错误2")
	}
}

func TestSAdd(t *testing.T) {
	providers.NewRedisProvider()

	k, v := "test", "test1"
	SAdd(k, v, "test2")

	SRem(k, v)
}
