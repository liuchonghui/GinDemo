package main

import (
	"html/template"
	"github.com/gin-gonic/gin"
	"fmt"
	"net/http"
	"github.com/gorilla/websocket"
	"log"
	"io/ioutil"
	"encoding/json"
	"strconv"
	"net"
)

var DB = make(map[string]string)

var conn *websocket.Conn

func logcat(format string, v ...interface{}) {
	if conn != nil {
		var err = conn.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf(format, v...)))
		if err != nil {
			log.Printf("[C]conn.writemessage:", err)
		}
	} else {
		log.Printf("[C]conn == nil")
	}
}

func XUserAgent(req *http.Request) string {
	return req.Header.Get("User-Agent")
}

func RemoteIp(req *http.Request) string {
	remoteAddr := req.RemoteAddr
	if ip := req.Header.Get("X-Real-IP"); ip != "" {
		remoteAddr = ip
	} else if ip = req.Header.Get("X-Forwarded-For"); ip != "" {
		remoteAddr = ip
	} else {
		remoteAddr, _, _ = net.SplitHostPort(remoteAddr)
	}

	if remoteAddr == "::1" {
		remoteAddr = "127.0.0.1"
	}

	return remoteAddr
}

func XRealIp(req *http.Request) string {
	remoteAddr := req.RemoteAddr
	if ip := req.Header.Get("X-Real-IP"); ip != "" {
		remoteAddr = ip
	} else {
		remoteAddr, _, _ = net.SplitHostPort(remoteAddr)
	}

	if remoteAddr == "::1" {
		remoteAddr = "127.0.0.1"
	}
	return remoteAddr
}

func XForwardedFor(req *http.Request) string {
	remoteAddr := req.RemoteAddr
	if ip := req.Header.Get("X-Forwarded-For"); ip != "" {
		remoteAddr = ip
	} else {
		remoteAddr, _, _ = net.SplitHostPort(remoteAddr)
	}

	if remoteAddr == "::1" {
		remoteAddr = "127.0.0.1"
	}
	return remoteAddr
}

func setupRouter() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()

	// Ping test
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	r.GET("/test", func(c *gin.Context) {
		c.String(200, RemoteIp(c.Request))
	})

	r.GET("/header", func(c *gin.Context) {
		var header = Header{
			RealIp: XRealIp(c.Request),
			ForwardedFor: XForwardedFor(c.Request),
			RemoteAddr: RemoteIp(c.Request),
			UserAgent: XUserAgent(c.Request),
		}
		c.JSON(http.StatusOK, header)
	})

	r.GET("/shadow", func(c *gin.Context) {
		var shdata = ShData{Id:"shadow_model",
			Md5:"725d58a7adc6a5fe5265dba87435bbc8",
			Ver:"model_s12a_117_type1",
			Url:"https://gist.github.com/liuchonghui/277cd9fac31b8cff8c9ccbc3600b55fd/raw/dce6b640b6f29e79d9093f9cca1702ec0c87ea4a/model_s12a_117_type1.bin",
		}
		var shadow = Shadow{
			ShResult: 0,
			ShMsg:"",
			ShData: shdata,
		}
		// also:
		//var shadow Shadow
		//shadow.ShResult = "success"
		//shadow.ShData = shdata
		c.JSON(http.StatusOK, shadow)
	})

	r.GET("/shadows", func(c *gin.Context) {
		var s []ShData
		var shadow_data = ShData{Id:"shadow_model",
			Md5:"725d58a7adc6a5fe5265dba87435bbc8",
			Ver:"model_s12a_117_type1",
			Url:"https://gist.github.com/liuchonghui/277cd9fac31b8cff8c9ccbc3600b55fd/raw/dce6b640b6f29e79d9093f9cca1702ec0c87ea4a/model_s12a_117_type1.bin",
		}
		s = append(s, shadow_data)
		var shadows Shadows
		shadows.ShResult = 0
		shadows.ShMsg = ""
		shadows.ShDatas = s
		c.JSON(http.StatusOK, shadows)
	})

	r.GET("/shadows_err", func(c *gin.Context) {
		var s []ShData = []ShData{}
		//var shadow_data = ShData{Id:"shadow_model",
		//	Md5:"725d58a7adc6a5fe5265dba87435bbc8",
		//	Ver:"model_s12a_117_type1",
		//	Url:"https://gist.github.com/liuchonghui/277cd9fac31b8cff8c9ccbc3600b55fd/raw/dce6b640b6f29e79d9093f9cca1702ec0c87ea4a/model_s12a_117_type1.bin",
		//}
		//s = append(s, shadow_data)
		var shadows Shadows
		shadows.ShResult = 0
		shadows.ShMsg = ""
		shadows.ShDatas = s
		c.JSON(http.StatusOK, shadows)
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

	r.POST("/bot617804206", func(c *gin.Context) {
		log.Printf("post /bot617804206 here:")
		data, _ := ioutil.ReadAll(c.Request.Body)
		log.Printf("request-body: %v", string(data))

		logcat(string(data))

		var tele Telegram
		if json.Unmarshal(data, &tele) == nil {
			log.Printf("update_id: %d", tele.UpdateId)
			log.Printf("message-date: %d", tele.Content.Date)
			log.Printf("message-message_id: %d", tele.Content.MessageId)
			log.Printf("message-text: %s", tele.Content.Text)
			logcat("message-text: %s", tele.Content.Text)
			log.Printf("message-chat-lastname: %s", tele.Content.ChatContent.LastName)
			log.Printf("message-chat-id: %d", tele.Content.ChatContent.Id)
			log.Printf("message-chat-type: %s", tele.Content.ChatContent.Type)
			log.Printf("message-chat-firstname: %s", tele.Content.ChatContent.FirstName)
			log.Printf("message-chat-username: %s", tele.Content.ChatContent.UserName)
			log.Printf("message-from-lastname: %s", tele.Content.FromContent.LastName)
			log.Printf("message-from-id: %d", tele.Content.FromContent.Id)
			log.Printf("message-from-type: %s", tele.Content.FromContent.Type)
			log.Printf("message-from-firstname: %s", tele.Content.FromContent.FirstName)
			log.Printf("message-from-username: %s", tele.Content.FromContent.UserName)
		}

		c.String(200, "success")
	})


	r.GET("/console", func(c *gin.Context) {
		log.Printf("[C]exec: %s", "ws://"+c.Request.Host+"/cat")
		if c.Request.Host == "localhost:8080" {
			log.Printf("[C]localhost: %s", "ws://"+c.Request.Host+"/cat")
			err := consoleTemplate.Execute(c.Writer, "ws://"+c.Request.Host+"/cat")
			if err != nil {
				log.Print("[C]upgrade:", err)
			}
		} else {
			log.Printf("[C]127.0.0.1: %s", "ws://45.32.40.65:8080/cat")
			err := consoleTemplate.Execute(c.Writer, "ws://45.32.40.65:8080/cat")
			if err != nil {
				log.Print("[C]upgrade:", err)
			}
		}

	})

	r.GET("/cat", func(a *gin.Context) {
		log.Printf("[C]>>>>>>>>>>/cat")
		tempConn, err := upgrader.Upgrade(a.Writer, a.Request, nil)
		if err != nil {
			log.Print("[C]upgrade:", err)
			return
		} else {
			log.Print("[C]tempConn is availiable")
		}
		defer tempConn.Close()
		conn = tempConn
		for {
			mt, message, err := conn.ReadMessage()
			if err != nil {
				log.Println("[C]read:", err)
				break
			}
			log.Printf("[C]mt: %d", mt)
			log.Printf("[C]recv: %s", message)
			err = conn.WriteMessage(mt, message)
			if err != nil {
				log.Println("[C]write:", err)
				break
			}
		}
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
//////////////////////////////

	r.GET("/vc_review", func(c *gin.Context) {
		ver := c.DefaultQuery("ver", "N/A")
		fmt.Println("ver is " + ver)
		var s []Review
		////////////////////1625version//////////////////////
		var review_interaction_100 = Review{Id:"vc.lab.interaction",
			Md5:Interaction_100_md5,
			Cp:Interaction,
			Url:Interaction_100_url}
		s = append(s, review_interaction_100)

		var data Dat
		data.Result = "success"
		data.Reviews = s
		var lr LabReviewResult
		lr.Content = data

		c.JSON(http.StatusOK, lr)
	})
/////////////////////////////

	r.GET("/fetch_plugin", func(c *gin.Context) {
		ver := c.DefaultQuery("ver", "N/A")
		fmt.Println("ver is " + ver)
		var s []Plugin
		////////////////////1625version//////////////////////
		var plugin_chushou_1626 = Plugin{Id:"chushou.plugin",
			Md5:Chushou_1626_md5,
			Cp:ChuShou,
			Url:Chushou_1626_url}
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
		var plugin_youku_1626 = Plugin{Id:"youku.plugin",
			Md5:Youku_1626_md5,
			Cp:YouKu,
			Url:Youku_1626_url}
		var plugin_bobo_1626 = Plugin{Id:"bobo.plugin",
			Md5:Bobo_1626_md5,
			Cp:BoBo,
			Url:Bobo_1626_url}
		var plugin_weibo_1628 = Plugin{Id:"cp.weibo.plugin",
			Md5:Weibo_1628_md5,
			Cp:WeiBo,
			Url:Weibo_1628_url}
		var plugin_huashu_1626 = Plugin{Id:"cp.huashu.plugin",
			Md5:Huashu_1626_md5,
			Cp:HuaShu,
			Url:Huashu_1626_url}
		var plugin_pptv_1625 = Plugin{Id:"cp.pptv.plugin",
			Md5:Pptv_1625_md5,
			Cp:PpTv,
			Url:Pptv_1625_url}
		var plugin_fxpgc_1629 = Plugin{Id:"fxpgc.plugin",
			Md5:Fxpgc_1629_md5,
			Cp:FxPgc,
			Url:Fxpgc_1629_url}
        var plugin_mangguo_1626 = Plugin{Id:"mangguo.plugin",
            Md5:Mangguo_1626_md5,
            Cp:MangGuo,
            Url:Mangguo_1626_url}
		////////////////////1625version//////////////////////

		s = append(s, plugin_chushou_1626)
		s = append(s, plugin_fengxing_1625)
		s = append(s, plugin_renren_1625)
		s = append(s, plugin_yilan_1625)
		s = append(s, plugin_erlan_1625)
		s = append(s, plugin_fenghuang_1626)
		s = append(s, plugin_youku_1626)
		s = append(s, plugin_bobo_1626)
		s = append(s, plugin_weibo_1628)
		s = append(s, plugin_huashu_1626)
		s = append(s, plugin_pptv_1625)
		s = append(s, plugin_fxpgc_1629)
        s = append(s, plugin_mangguo_1626)

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
			Md5:Qingting_189_md5,
			Cp:QingTing,
			Url:Qingting_189_url}
		////////////////////101version//////////////////////

		s = append(s, plugin_qingting)

		var data Data
		data.Result = "success"
		data.Plugins = s
		var fp FetchPluginResult
		fp.Content = data

		c.JSON(http.StatusOK, fp)
	})

	r.GET("/signature", func(c *gin.Context) {
		dat0 := c.DefaultQuery("dat0", "0")
		fmt.Println("dat0 is " + dat0)
		var s []Signature
		var sig_chengdu = Signature{City:"成都市",
			Code:510500,
			District:"成都市"}
		var sig_yibin = Signature{City:"宜宾市",
			Code:511501,
			District:"宜宾市"}
		var sig_ziyang = Signature{City:"资阳市",
			Code:511502,
			District:"资阳市"}
		var sig_dazhou = Signature{City:"达州市",
			Code:511503,
			District:"达州市"}
		var sig_chengdu2 = Signature{City:"成都市2成都市2成都市2成都市2成都市2成都市2成都市2成都市2",
			Code:510110,
			District:"成都市2成都市2成都市2成都市2成都市2成都市2成都市2成都市2成都市2"}
		var sig_yibin2 = Signature{City:"宜宾市2宜宾市2宜宾市2宜宾市2宜宾市2",
			Code:511520,
			District:"宜宾市2宜宾市2"}
		var sig_ziyang2 = Signature{City:"资阳市2",
			Code:511531,
			District:"资阳市2资阳市2资阳市2资阳市2资阳市2资阳市2资阳市2资阳市2"}
		var sig_dazhou2 = Signature{City:"达州市2达州市2达州市2",
			Code:511542,
			District:"达州市2达州市2达州市2达州市2达州市2达州市2达州市2达州市2达州市2达州市2达州市2达州市2"}
		var sig_chengdu3 = Signature{City:"成都市3",
			Code:510100,
			District:"成都市3"}
		var sig_yibin3 = Signature{City:"宜宾市3",
			Code:511200,
			District:"宜宾市3"}
		var sig_ziyang3 = Signature{City:"资阳市3",
			Code:511301,
			District:"资阳市3"}
		var sig_dazhou3 = Signature{City:"达州市3",
			Code:511402,
			District:"达州市3"}
		var sig_chengdu4 = Signature{City:"成都市4",
			Code:511500,
			District:"成都市4"}
		var sig_yibin4 = Signature{City:"宜宾市4",
			Code:512501,
			District:"宜宾市4"}
		var sig_ziyang4 = Signature{City:"资阳市4",
			Code:513502,
			District:"资阳市4"}
		var sig_dazhou4 = Signature{City:"达州市4",
			Code:514503,
			District:"达州市4"}
		var sig_chengdu5 = Signature{City:"成都市5",
			Code:500500,
			District:"成都市5"}
		var sig_yibin5 = Signature{City:"宜宾市5",
			Code:521501,
			District:"宜宾市5"}
		var sig_ziyang5 = Signature{City:"资阳市5",
			Code:531502,
			District:"资阳市5"}
		var sig_dazhou5 = Signature{City:"达州市5",
			Code:541503,
			District:"达州市"}


		s = append(s, sig_chengdu)
		s = append(s, sig_yibin)
		s = append(s, sig_ziyang)
		s = append(s, sig_dazhou)
		s = append(s, sig_chengdu2)
		s = append(s, sig_yibin2)
		s = append(s, sig_ziyang2)
		s = append(s, sig_dazhou2)
		s = append(s, sig_chengdu3)
		s = append(s, sig_yibin3)
		s = append(s, sig_ziyang3)
		s = append(s, sig_dazhou3)
		s = append(s, sig_chengdu4)
		s = append(s, sig_yibin4)
		s = append(s, sig_ziyang4)
		s = append(s, sig_dazhou4)
		s = append(s, sig_chengdu5)
		s = append(s, sig_yibin5)
		s = append(s, sig_ziyang5)
		s = append(s, sig_dazhou5)

		var  sr SignatureResult
		sr.Code, _ = strconv.Atoi(dat0)
		sr.CurrentSign = []string {}
		sr.Signature = s

		c.JSON(http.StatusOK, sr)
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

type Header struct {
	RealIp string `json:"real_ip"`
	ForwardedFor string `json:"forwarded_for"`
	RemoteAddr string `json:"remote_addr"`
	UserAgent string `json:"user_agent"`
}

type Shadow struct {
	ShResult int `json:"result"`
	ShMsg string `json:"message"`
	ShData ShData `json:"data"`
}

type ShData struct {
	Id string `json:"_id"`
	Md5 string `json:"md5"`
	Ver string `json:"ver"`
	Url string `json:"url"`
}

type Shadows struct {
	ShResult int `json:"result"`
	ShMsg string `json:"message"`
	ShDatas []ShData `json:"data"`
}

//type PluginResult struct {
//	State string `json:"state"`
//	Plugins []Plugin `json:"plugins"`
//}
type LabReviewResult struct {
	Content Dat `json:"data"`
}


type Dat struct {
	Result string `json:"result"`
	Reviews []Review `json:"vc_review"`
}

type Review struct {
	Id string `json:"_id"`
	Md5 string `json:"md5"`
	Cp string `json:"vc"`
	Url string `json:"url"`
}
////////////////////////////////////
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

//////////////telegram/////////////
type Person struct {
	Name string `form:"name"`
	Address string `form:"address"`
}

type Telegram struct {
	UpdateId int `json:"update_id"`
	Content Message `json:"message"`
}

type Message struct {
	Date int `json:"date"`
	ChatContent Chat `json:"chat"`
	MessageId int `json:"message_id"`
	FromContent From `json:"from"`
	Text string `json:"text"`
}

type Chat struct {
	LastName string `json:"last_name"`
	Id int `json:"id"`
	Type string `json:"type"`
	FirstName string `json:"first_name"`
	UserName string `json:"username"`
}

type From struct {
	LastName string `json:"last_name"`
	Id int `json:"id"`
	Type string `json:"type"`
	FirstName string `json:"first_name"`
	UserName string `json:"username"`
}

/////////signature/////////

type SignatureResult struct {
	Code int `json:"code"`
	CurrentSign []string `json:"current_sign"`
	Signature []Signature `json:"signature"`
}

type Signature struct {
	City string `json:"city"`
	Code int `json:"code"`
	District string `json:"district"`
}

////////////////////////////

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

var consoleTemplate = template.Must(template.New("").Parse(`
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
        output.appendChild(d).scrollIntoView();
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
            print("RESP| " + evt.data);
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
