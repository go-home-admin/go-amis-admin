package amis

import (
	"github.com/gin-gonic/gin"
	http2 "github.com/go-home-admin/home/app/http"
	"github.com/go-home-admin/home/bootstrap/services/database"
	"github.com/go-home-admin/home/bootstrap/services/filesystem"
	"github.com/sirupsen/logrus"
	"net/http"
)

func (f *Form) InputImage(name, label string) *FormItemImage {
	item := &FormItemImage{
		FormItem: NewItem(name, label, "input-image"),
		ctx:      f.ctx,
	}
	f.AddBody(item)
	return item
}

type FormItemImage struct {
	*FormItem
	ctx *gin.Context
}

// Update 图片保存地址, 同时设置闭包函数
func (i *FormItemImage) Update(path string) *FormItemImage {
	action := "_image_" + i.Name
	AddCallRoutes(i.ctx, action, func(c *gin.Context) {
		// 获取上传的文件
		url, err := filesystem.NewLocal().FormFile(i.ctx, "file", "/images/"+path+"/"+database.Now().Ymd())
		if err != nil {
			logrus.Error(err)
			c.String(http.StatusBadRequest, "上传失败")
			return
		}
		// 返回成功信息
		http2.NewContext(i.ctx).Success(map[string]interface{}{
			"url": url,
		})
	})
	return i.Receiver(GetUrl(i.ctx, "/"+action))
}

// Receiver 原始设置上传地址
// 选中文件后，就会自动调用 receiver 配置里的接口进行上传
//
//	response{
//	  "status": 0,
//	  "msg": "",
//	  "data": {
//	    "value": "https:/xxx.yy/zz.png"
//	  }
//	}
func (i *FormItemImage) Receiver(url string) *FormItemImage {
	i.SetOptions("receiver", url)
	return i
}

// Accept 想要限制多个类型，则用逗号分隔，例如：.jpg,.png
func (i *FormItemImage) Accept(ext string) *FormItemImage {
	i.SetOptions("accept", ext)
	return i
}

// Limit 限制文件大小
func (i *FormItemImage) Limit(minWidth int) *FormItemImage {
	i.SetOptions("limit", map[string]interface{}{
		"minWidth": minWidth,
	})
	return i
}

// Crop 支持裁剪
func (i *FormItemImage) Crop() *FormItemImage {
	i.SetOptions("crop", true)
	return i
}

// AutoFill 自动填充
//
//	response{
//	 "status": 0,
//	 "msg": "",
//	 "data": {
//	   "value": "xxxxxxx",
//	   "filename": "xxxx.jpg",
//	   "url": "http://xxxx.xxx.xxx"
//	 }
//	}
func (i *FormItemImage) AutoFill() *FormItemImage {
	i.SetOptions("autoFill", map[string]interface{}{
		"myUrl": "${url}",
	})
	return i
}
