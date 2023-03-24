package sys_queue

import (
	"context"
	redis2 "github.com/go-home-admin/go-admin/app/common/redis"
	"github.com/go-home-admin/home/app"
	"github.com/go-redis/redis/v8"
	"github.com/sirupsen/logrus"
	"strconv"
	"strings"
)

func getPendingList2(stream, start string) []redis.XPendingExt {
	group := app.Config("queue.queue.group_name", "home_default_group")
	XPending := redis2.Client().XPending(context.Background(), stream, group).Val()
	var pendingList []redis.XPendingExt
	if XPending != nil && XPending.Count > 0 {
		count := int64(pageSize + 1)
		if start == "" {
			start = XPending.Lower
		} else if start == "all" {
			start = XPending.Lower
			count = XPending.Count
		}
		pendingList = redis2.Client().XPendingExt(context.Background(), &redis.XPendingExtArgs{
			Stream: stream,
			Group:  group,
			Start:  start,
			End:    XPending.Higher,
			Count:  count,
		}).Val()
	}
	return pendingList
}

func getPendingList(stream, group, start, end string) map[string]bool {
	// 失败任务ID
	pendingIdMap := map[string]bool{}
	if start == "" || "" == end {
		return pendingIdMap
	}
	XPending := redis2.Client().XPending(context.Background(), stream, group).Val()
	var pendingList []redis.XPendingExt
	if XPending != nil && XPending.Count > 0 {
		// 检查是否在区间
		if !isDone(end, XPending.Lower) || (isDone(start, XPending.Higher) && start != XPending.Higher) {
			return pendingIdMap
		}
		if !isDone(start, XPending.Lower) {
			// 在区间内部, 减少查询
			start = XPending.Lower
		}
		if isDone(end, XPending.Higher) {
			// 超出区间, 减少查询
			end = XPending.Higher
		}

		pendingList = redis2.Client().XPendingExt(context.Background(), &redis.XPendingExtArgs{
			Stream: stream,
			Group:  group,
			Start:  start,
			End:    end,
			Count:  pageSize,
		}).Val()
	}

	for _, task := range pendingList {
		pendingIdMap[task.ID] = true
	}
	return pendingIdMap
}

// isDone 检查任务是否已执行，false为未执行
// (id <= lastId) == true 就是执行过
func isDone(id, lastId string) bool {
	if lastId == "" {
		return false
	}
	arr1 := strings.Split(id, "-")
	arr2 := strings.Split(lastId, "-")
	if strToInt(arr1[0]) > strToInt(arr2[0]) {
		return false
	} else if strToInt(arr1[0]) == strToInt(arr2[0]) {
		if strToInt(arr1[1]) > strToInt(arr2[1]) {
			return false
		}
	}
	return true
}

func strToInt(str string) int {
	num, _ := strconv.Atoi(str)
	return num
}

func getLastDeliveredID(stream, group string) string {
	XInfo, err := redis2.Client().XInfoGroups(context.Background(), stream).Result()
	if err != nil {
		logrus.Error(err)
	}
	lastDeliveredID := ""
	for _, item := range XInfo {
		if item.Name == group {
			lastDeliveredID = item.LastDeliveredID
		}
	}
	return lastDeliveredID
}
