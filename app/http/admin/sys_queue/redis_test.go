package sys_queue

import (
	"github.com/go-home-admin/go-admin/app/message"
	"github.com/go-home-admin/home/bootstrap/providers"
	"github.com/go-home-admin/home/bootstrap/servers"
	"github.com/go-home-admin/home/bootstrap/utils"
	"testing"
)

func Test_getPendingList10000(t *testing.T) {
	providers.NewRedisProvider()

	servers.NewQueue().Push(&message.DemoMessage{})
	servers.NewQueue().Push(&message.DemoMessage{})
	servers.NewQueue().Push(&message.DemoMessage{})
	servers.NewQueue().Push(&message.DemoMessage{})
	servers.NewQueue().Push(&message.DemoMessage{})
	servers.NewQueue().Push(&message.DemoMessage{})
	servers.NewQueue().Push(&message.DemoMessage{})
	servers.NewQueue().Push(&message.DemoMessage{})
	servers.NewQueue().Push(&message.DemoMessage{})
	servers.NewQueue().Push(&message.DemoMessage{})
	servers.NewQueue().Push(&message.DemoMessage{})
	servers.NewQueue().Push(&message.DemoMessage{})
	servers.NewQueue().Push(&message.DemoMessage{})
	servers.NewQueue().Push(&message.DemoMessage{})
	servers.NewQueue().Push(&message.DemoMessage{})
	servers.NewQueue().Push(&message.DemoMessage{})
	servers.NewQueue().Push(&message.DemoMessage{})
	servers.NewQueue().Push(&message.DemoMessage{})
	servers.NewQueue().Push(&message.DemoMessage{})
	servers.NewQueue().Push(&message.DemoMessage{})
	servers.NewQueue().Push(&message.DemoMessage{})
	servers.NewQueue().Push(&message.DemoMessage{})
	servers.NewQueue().Push(&message.DemoMessage{})
	servers.NewQueue().Push(&message.DemoMessage{})
	servers.NewQueue().Push(&message.DemoMessage{})
	servers.NewQueue().Push(&message.DemoMessage{})
	servers.NewQueue().Push(&message.DemoMessage{})
	servers.NewQueue().Push(&message.DemoMessage{})

	getLastDeliveredID("home_default_stream", "home_default_group")

	list := getPendingList2("home_default_stream", "all")

	utils.Dump(list)
}
