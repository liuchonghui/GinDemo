package main

import (
	"html/template"
	"github.com/gin-gonic/gin"
	"fmt"
	"net/http"
	"github.com/gorilla/websocket"
	"log"
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

	r.GET("/urlscheme", func(c *gin.Context) {
		var s []Plugin
		var plugin_urlscheme = Plugin{Id:"android.test.urlscheme",
			Md5:"1293f4412b5c63ef71b6d8d3ff3d1e5e",
			Cp:"us",
			Url:"https://gist.github.com/liuchonghui/b9757b65748eb42548213ec7b9572116/raw/b64e80fe78a2e14bf2cc4675a6def6f1ffd4a4d2/urlscheme.1293f4412b5c63ef71b6d8d3ff3d1e5e.zip"}
		s = append(s, plugin_urlscheme)

		var data Data
		data.Result = "success"
		data.Plugins = s
		var fp FetchPluginResult
		fp.Content = data
		c.JSON(http.StatusOK, fp)
	})

	r.GET("/home", func(c *gin.Context) {
		log.Printf("exec: %s", "ws://"+c.Request.Host+"/echo")
		if c.Request.Host == "localhost:8080" {
			log.Printf("localhost: %s", "ws://"+c.Request.Host+"/echo")
			err := homeTemplate.Execute(c.Writer, "ws://"+c.Request.Host+"/echo")
			if err != nil {
				log.Print("upgrade:", err)
			}
		} else {
			log.Printf("127.0.0.1: %s", "ws://45.32.40.65:8080/echo")
			err := homeTemplate.Execute(c.Writer, "ws://45.32.40.65:8080/echo")
			if err != nil {
				log.Print("upgrade:", err)
			}
		}

	})

	r.GET("/echo", func(a *gin.Context) {
		log.Printf(">>>>>>>>>>/echo")
		c, err := upgrader.Upgrade(a.Writer, a.Request, nil)
		if err != nil {
			log.Print("upgrade:", err)
			return
		}
		defer c.Close()
		for {
			mt, message, err := c.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				break
			}
			log.Printf("recv: %s", message)
			err = c.WriteMessage(mt, message)
			if err != nil {
				log.Println("write:", err)
				break
			}
		}
	})

	r.GET("/g/echo", func(a *gin.Context) {
		log.Printf(">>>>>>>>>>/g/echo")
		c, err := upgrader.Upgrade(a.Writer, a.Request, nil)
		if err != nil {
			log.Print("upgrade:", err)
			return
		}
		defer c.Close()
		for {
			mt, message, err := c.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				break
			}
			log.Printf("recv: %s", message)
			err = c.WriteMessage(mt, message)
			if err != nil {
				log.Println("write:", err)
				break
			}
		}
	})

	//fmt.Println(AAA + AAA)

	r.GET("/fetch_plugin", func(c *gin.Context) {
		ver := c.DefaultQuery("ver", "N/A")
		fmt.Println("ver is " + ver)
		var s []Plugin
		////////////////////1625version//////////////////////
		var plugin_chushou_1625 = Plugin{Id:"chushou.plugin",
			Md5:Chushou_1625_md5,
			Cp:ChuShou,
			Url:Chushou_1625_url}
		var plugin_fengxing_1625 = Plugin{Id:"fengxing.plugin",
			Md5:Fengxing_1625_md5,
			Cp:FengXing,
			Url:Fengxing_1625_url}
		var plugin_renren_1625 = Plugin{Id:"renren.plugin",
			Md5:Renren_1625_md5,
			Cp:RenRen,
			Url:Renren_1625_url}
		var plugin_yilan_1625 = Plugin{Id:"yilan.plugin",
			Md5:Yilan_1625_md5,
			Cp:YiLan,
			Url:Yilan_1625_url}
		var plugin_erlan_1625 = Plugin{Id:"erlan.plugin",
			Md5:Erlan_1625_md5,
			Cp:ErLan,
			Url:Erlan_1625_url}
		var plugin_fenghuang_1626 = Plugin{Id:"fenghuang.plugin",
			Md5:Fenghuang_1626_md5,
			Cp:FengHuang,
			Url:Fenghuang_1626_url}
		var plugin_youku_1625 = Plugin{Id:"youku.plugin",
			Md5:Youku_1625_md5,
			Cp:YouKu,
			Url:Youku_1625_url}
		var plugin_bobo_1626 = Plugin{Id:"bobo.plugin",
			Md5:Bobo_1626_md5,
			Cp:BoBo,
			Url:Bobo_1626_url}
		var plugin_weibo_1626 = Plugin{Id:"cp.weibo.plugin",
			Md5:Weibo_1626_md5,
			Cp:WeiBo,
			Url:Weibo_1626_url}
		var plugin_huashu_1625 = Plugin{Id:"cp.huashu.plugin",
			Md5:Huashu_1625_md5,
			Cp:HuaShu,
			Url:Huashu_1625_url}
		var plugin_pptv_1625 = Plugin{Id:"cp.pptv.plugin",
			Md5:Pptv_1625_md5,
			Cp:PpTv,
			Url:Pptv_1625_url}
		////////////////////1625version//////////////////////

		s = append(s, plugin_chushou_1625)
		s = append(s, plugin_fengxing_1625)
		s = append(s, plugin_renren_1625)
		s = append(s, plugin_yilan_1625)
		s = append(s, plugin_erlan_1625)
		s = append(s, plugin_fenghuang_1626)
		s = append(s, plugin_youku_1625)
		s = append(s, plugin_bobo_1626)
		s = append(s, plugin_weibo_1626)
		s = append(s, plugin_huashu_1625)
		s = append(s, plugin_pptv_1625)

		var data Data
		data.Result = "success"
		data.Plugins = s
		var fp FetchPluginResult
		fp.Content = data

		c.JSON(http.StatusOK, fp)
	})

	r.GET("/fetch_fm_plugin", func(c *gin.Context) {
		ver := c.DefaultQuery("ver", "N/A")
		fmt.Println("ver is " + ver)
		var s []Plugin
		////////////////////101version//////////////////////
		var plugin_qingting = Plugin{Id:"cp.qingtingfm.plugin",
			Md5:Qingting_178_md5,
			Cp:QingTing,
			Url:Qingting_178_url}
		////////////////////101version//////////////////////

		s = append(s, plugin_qingting)

		var data Data
		data.Result = "success"
		data.Plugins = s
		var fp FetchPluginResult
		fp.Content = data

		c.JSON(http.StatusOK, fp)
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

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

var homeTemplate = template.Must(template.New("").Parse(`
<!DOCTYPE html>
<html>
<head>
<meta charset="utf-8">
<script>
window.addEventListener("load", function(evt) {
    var output = document.getElementById("output");
    var input = document.getElementById("input");
    var ws;
    var print = function(message) {
        var d = document.createElement("div");
        d.innerHTML = message;
        output.appendChild(d);
    };
    document.getElementById("open").onclick = function(evt) {
        if (ws) {
            return false;
        }
        ws = new WebSocket("{{.}}");
        ws.onopen = function(evt) {
            print("OPEN");
        }
        ws.onclose = function(evt) {
            print("CLOSE");
            ws = null;
        }
        ws.onmessage = function(evt) {
            print("RESPONSE: " + evt.data);
        }
        ws.onerror = function(evt) {
            print("ERROR: " + evt.data);
        }
        return false;
    };
    document.getElementById("send").onclick = function(evt) {
        if (!ws) {
            return false;
        }
        print("SEND: " + input.value);
        ws.send(input.value);
        return false;
    };
    document.getElementById("close").onclick = function(evt) {
        if (!ws) {
            return false;
        }
        ws.close();
        return false;
    };
});
</script>
</head>
<body>
<table>
<tr><td valign="top" width="50%">
<p>Click "Open" to create a connection to the server,
"Send" to send a message to the server and "Close" to close the connection.
You can change the message and send multiple times.
<p>
<form>
<button id="open">Open</button>
<button id="close">Close</button>
<p><input id="input" type="text" value="Hello world!">
<button id="send">Send</button>
</form>
</td><td valign="top" width="50%">
<div id="output"></div>
</td></tr></table>
</body>
</html>
`))
