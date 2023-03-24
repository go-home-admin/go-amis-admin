package sys_queue

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/go-home-admin/amis"
	"github.com/go-home-admin/go-admin/app/common/redis"
	"github.com/go-home-admin/home/app"
	"github.com/go-home-admin/home/app/http"
	"github.com/go-home-admin/home/bootstrap/services/database"
	"strconv"
	"strings"
)

const pageSize = 15

func (c *CrudContext) Common() {
	// c.SetDb(admin.NewOrmAdminMenu())
}

func (c *CrudContext) Table(curd *amis.Crud) {
	curd.Column("排序", "id")
	curd.Column("JobID", "jid")
	curd.Column("JOB路由", "route")
	curd.Column("JOB参数", "data").Json()
	curd.Column("状态", "status").Mapping()
	curd.Column("时间", "date")

	curd.Api().Data(map[string]interface{}{
		"pre":     "${pre}",
		"next":    "${next}",
		"page":    "${page}",
		"perPage": "${perPage}",
	})
}

func (c *CrudContext) Form(form *amis.Form) {

}

func (c *CrudContext) List(ctx *gin.Context) {
	got := amis.NewCurdData()
	steam := app.Config("queue.queue.stream_name", "home_default_stream")

	list, nextId, total := c.GetStreamTask(ctx, steam)
	got.Items = list
	got.Total = total
	got.Page, _ = strconv.Atoi(ctx.Query("page"))
	got.SetOptions("next", nextId)
	got.SetOptions("pre", got.Page)
	http.NewContext(ctx).Success(got)
}

func (c *CrudContext) GetStreamTask(ctx *gin.Context, stream string) ([]interface{}, string, int64) {
	var start, stop = "+", "-"
	// 上一次执行的ID
	group := app.Config("queue.queue.group_name", "home_default_group")
	lastDeliveredID := getLastDeliveredID(stream, group)

	total := redis.Client().XLen(context.Background(), stream).Val()
	list := make([]interface{}, 0)
	var startId, nextId string
	res := redis.Client().XRevRangeN(context.Background(), stream, start, stop, int64(pageSize))
	if res.Err() != nil {
		return nil, "", 0
	}
	for i, item := range res.Val() {
		if startId == "" {
			startId = item.ID
		}
		nextId = item.ID
		unix, _ := strconv.ParseInt(string([]byte(item.ID)[:10]), 10, 64)

		if _, ok := item.Values["route"]; !ok {
			continue
		}
		status := 1
		route := item.Values["route"].(string)
		index := strings.LastIndex(route, "/message")
		if index > 0 {
			route = route[index+8:]
		}
		list = append(list, map[string]interface{}{
			"id":     i,
			"jid":    item.ID,
			"route":  route,
			"data":   item.Values["event"].(string),
			"status": int32(status),
			"date":   database.UnixToTime(int64(unix)).YmdHis(),
		})
	}

	pendingList := getPendingList(stream, group, startId, nextId)
	for i, v := range list {
		m := v.(map[string]interface{})
		taskID := m["jid"].(string)
		status := "schedule"
		if lastDeliveredID == "" || isDone(taskID, lastDeliveredID) {
			status = "success"
		}
		if _, ok := pendingList[taskID]; ok {
			status = "fail"
		}
		m["status"] = status
		list[i] = v
	}

	return list, nextId, total
}

func (c *Controller) GinHandleCurd(ctx *gin.Context) {
	var crud = &CrudContext{}
	crud.CurdController.Context = ctx
	crud.CurdController.Crud = crud
	amis.GinHandleCurd(ctx, crud)
}

type CrudContext struct {
	amis.CurdController
}
