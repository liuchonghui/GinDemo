package main

import (
	"github.com/gin-gonic/gin"
	"fmt"
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

	//fmt.Println(AAA + AAA)

	r.GET("/fetch_plugin", func(c *gin.Context) {
		ver := c.DefaultQuery("ver", "N/A")
		fmt.Println("ver is " + ver)
		//c.JSON(200, gin.H{"id": ch, "state": "success", "md5": "819f24d8d44fb678e0b4c5cbfe3aca68", "url": "https://gist.github.com/liuchonghui/d671bb312dceb6540e8987578f09e3b1/raw/7ff6ea477d81e044f0eb100c2da14ddbbaf94457/chushou.plugin.pl.apk"})
		var s []Plugin
		////////////////////134version//////////////////////
		//var plugin_chushou_134 = Plugin{Id:"chushou.plugin",
		//	Md5:Chushou_134_md5,
		//	Cp:ChuShou,
		//	Url:Chushou_134_url}
		//var plugin_fengxing_134 = Plugin{Id:"fengxing.plugin",
		//	Md5:Fengxing_134_md5,
		//	Cp:FengXing,
		//	Url:Fengxing_134_url}
		//var plugin_renren_134 = Plugin{Id:"renren.plugin",
		//	Md5:Renren_134_md5,
		//	Cp:RenRen,
		//	Url:Renren_134_url}
		//var plugin_yilan_135 = Plugin{Id:"yilan.plugin",
		//	Md5:Yilan_135_md5,
		//	Cp:YiLan,
		//	Url:Yilan_135_url}
		////////////////////134version//////////////////////
		////////////////////146version//////////////////////
		//var plugin_chushou_146 = Plugin{Id:"chushou.plugin",
		//	Md5:Chushou_146_md5,
		//	Cp:ChuShou,
		//	Url:Chushou_146_url}
		//var plugin_fengxing_146 = Plugin{Id:"fengxing.plugin",
		//	Md5:Fengxing_146_md5,
		//	Cp:FengXing,
		//	Url:Fengxing_146_url}
		//var plugin_renren_146 = Plugin{Id:"renren.plugin",
		//	Md5:Renren_146_md5,
		//	Cp:RenRen,
		//	Url:Renren_146_url}
		//var plugin_yilan_146 = Plugin{Id:"yilan.plugin",
		//	Md5:Yilan_146_md5,
		//	Cp:YiLan,
		//	Url:Yilan_146_url}
		////////////////////146version//////////////////////
		////////////////////158version//////////////////////
		var plugin_chushou_158 = Plugin{Id:"chushou.plugin",
			Md5:Chushou_158_md5,
			Cp:ChuShou,
			Url:Chushou_158_url}
		var plugin_fengxing_158 = Plugin{Id:"fengxing.plugin",
			Md5:Fengxing_158_md5,
			Cp:FengXing,
			Url:Fengxing_158_url}
		var plugin_renren_158 = Plugin{Id:"renren.plugin",
			Md5:Renren_158_md5,
			Cp:RenRen,
			Url:Renren_158_url}
		var plugin_yilan_158 = Plugin{Id:"yilan.plugin",
			Md5:Yilan_158_md5,
			Cp:YiLan,
			Url:Yilan_158_url}
		var plugin_fenghuang_158 = Plugin{Id:"fenghuang.plugin",
			Md5:Fenghuang_158_md5,
			Cp:FengHuang,
			Url:Fenghuang_158_url}
		////////////////////158version//////////////////////


		s = append(s, plugin_chushou_158)
		s = append(s, plugin_fengxing_158)
		s = append(s, plugin_renren_158)
		s = append(s, plugin_yilan_158)
		s = append(s, plugin_fenghuang_158)

		//var ret PluginResult
		//ret.Plugins = s
		//ret.State = "success"

		var data Data
		data.Result = "success"
		data.Plugins = s
		var fp FetchPluginResult
		fp.Content = data

		c.JSON(http.StatusOK, fp)
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

func main() {
	r := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}

type Plugin struct {
	Id string `json:"_id"`
	Md5 string `json:"md5"`
	Cp string `json:"cp"`
	Url string `json:"url"`
}

//type PluginResult struct {
//	State string `json:"state"`
//	Plugins []Plugin `json:"plugins"`
//}

type FetchPluginResult struct {
	Content Data `json:"data"`
}

type Data struct {
	Result string `json:"result"`
	Plugins []Plugin `json:"cp_plugin"`
}
