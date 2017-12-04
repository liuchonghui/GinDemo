package main

import (
	"github.com/gin-gonic/gin"
	"fmt"
	"encoding/json"
	"net/http"
)

var DB = make(map[string]string)

func setupRouter() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()

	// Ping test
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	r.GET("/fetch_plugin", func(c *gin.Context) {
		ver := c.DefaultQuery("ver", "N/A")
		fmt.Println("ver is " + ver)
		//c.JSON(200, gin.H{"id": ch, "state": "success", "md5": "819f24d8d44fb678e0b4c5cbfe3aca68", "url": "https://gist.github.com/liuchonghui/d671bb312dceb6540e8987578f09e3b1/raw/7ff6ea477d81e044f0eb100c2da14ddbbaf94457/chushou.plugin.pl.apk"})
		var s []Plugin
		s = append(s, Plugin{Id:"chushou.plugin",Md5:"819f24d8d44fb678e0b4c5cbfe3aca68",Url:"https://gist.github.com/liuchonghui/d671bb312dceb6540e8987578f09e3b1/raw/7ff6ea477d81e044f0eb100c2da14ddbbaf94457/chushou.plugin.pl.apk"})
		s = append(s, Plugin{Id:"yilan.plugin",Md5:"2ca0e18b7c624fae402a6988a2494d39",Url:"https://gist.github.com/liuchonghui/d671bb312dceb6540e8987578f09e3b1/raw/6cfbd5852f00644fb8d1ab8222bda0abafdab706/yilan.plugin.pl.apk"})
		s = append(s, Plugin{Id:"renren.plugin",Md5:"8ac11e966aef40c77722b57e3aebc8ac",Url:"https://gist.github.com/liuchonghui/d671bb312dceb6540e8987578f09e3b1/raw/6cfbd5852f00644fb8d1ab8222bda0abafdab706/renren.plugin.pl.apk"})
		s = append(s, Plugin{Id:"fengxing.plugin",Md5:"e4f897ab3b08c365c4c01d46b631d81c",Url:"https://gist.github.com/liuchonghui/d671bb312dceb6540e8987578f09e3b1/raw/6cfbd5852f00644fb8d1ab8222bda0abafdab706/fengxing.plugin.pl.apk"})
		var ret PluginResult
		ret.Plugins = s
		ret.State = "success"
		c.JSON(http.StatusOK, ret)
		//b, err := json.Marshal(ret)
		//if err != nil {
		//	c.JSON(100, "{}")
		//} else {
		//	fmt.Println(string(b))
		//	c.JSON(200, string(b))
		//}
	})

	// Get user value
	r.GET("/user/:name", func(c *gin.Context) {
		user := c.Params.ByName("name")
		value, ok := DB[user]
		if ok {
			c.JSON(200, gin.H{"user": user, "value": value})
		} else {
			c.JSON(200, gin.H{"user": user, "status": "no value"})
		}
	})

	// Authorized group (uses gin.BasicAuth() middleware)
	// Same than:
	// authorized := r.Group("/")
	// authorized.Use(gin.BasicAuth(gin.Credentials{
	//	  "foo":  "bar",
	//	  "manu": "123",
	//}))
	authorized := r.Group("/", gin.BasicAuth(gin.Accounts{
		"foo":  "bar", // user:foo password:bar
		"manu": "123", // user:manu password:123
	}))

	authorized.POST("admin", func(c *gin.Context) {
		user := c.MustGet(gin.AuthUserKey).(string)

		// Parse JSON
		var json struct {
			Value string `json:"value" binding:"required"`
		}

		if c.Bind(&json) == nil {
			DB[user] = json.Value
			c.JSON(200, gin.H{"status": "ok"})
		}
	})

	return r
}

func FetchPluginHandler(c *gin.Context) {
	//type JsonHolder struct {
	//	Id   int    `json:"id"`
	//	Name string `json:"name"`
	//}
	//holder := JsonHolder{Id: 1, Name: "my name"}
	////若返回json数据，可以直接使用gin封装好的JSON方法
	//c.JSON(http.StatusOK, holder)

	var s []Plugin
	s = append(s, Plugin{Id:"chushou.plugin",Md5:"819f24d8d44fb678e0b4c5cbfe3aca68",Url:"https://gist.github.com/liuchonghui/d671bb312dceb6540e8987578f09e3b1/raw/7ff6ea477d81e044f0eb100c2da14ddbbaf94457/chushou.plugin.pl.apk"})
	s = append(s, Plugin{Id:"yilan.plugin",Md5:"2ca0e18b7c624fae402a6988a2494d39",Url:"https://gist.github.com/liuchonghui/d671bb312dceb6540e8987578f09e3b1/raw/9855826c39d4e5a7b3c38338f0d907ed400d3081/YiLanPlugin-debug-pl.apk"})
	s = append(s, Plugin{Id:"renren.plugin",Md5:"8ac11e966aef40c77722b57e3aebc8ac",Url:"https://gist.github.com/liuchonghui/d671bb312dceb6540e8987578f09e3b1/raw/9855826c39d4e5a7b3c38338f0d907ed400d3081/RenRenPlugin-debug-pl.apk"})
	s = append(s, Plugin{Id:"fengxing.plugin",Md5:"e4f897ab3b08c365c4c01d46b631d81c",Url:"https://gist.github.com/liuchonghui/d671bb312dceb6540e8987578f09e3b1/raw/9855826c39d4e5a7b3c38338f0d907ed400d3081/FengXingPlugin-debug-pl.apk"})
	var ret PluginResult
	ret.Plugins = s
	ret.State = "success"
	b, err := json.Marshal(ret)
	if err != nil {
		c.JSON(http.StatusOK, "{}")
	} else {
		fmt.Println(string(b))
		c.JSON(http.StatusOK, string(b))
	}
	return
}

func main() {
	r := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}

type Plugin struct {
	Id string `json:"id"`
	Md5 string `json:"md5"`
	Url string `json:"url"`
}

type PluginList struct {
	Plugins []Plugin `json:"plugins"`
}

type PluginResult struct {
	State string `json:"state"`
	Plugins []Plugin `json:"plugins"`
}
