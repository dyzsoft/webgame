/**
	author  dyz
 */

package controllers

import (
	"github.com/astaxie/beego"
	"math/rand"
	"time"
	"strconv"
	"fmt"
	"encoding/hex"
	"crypto/md5"
	"strings"
)

const (
	// for session
	LEVEL    = "level"
	PASSWORD = "password"
	// End  for session

	// for c.Data[""]
	MESSAGE  = "message"
	// end for c.Data

	// final level  the tpl is gameover.html
	FinalLevel = 8
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Prepare() {

	//  init level in  session ,default 0
	if c.GetSession(LEVEL)  == nil {
		c.SetSession(LEVEL,0)
	}

	c.Data[LEVEL] = c.GetSession(LEVEL).(int)
}

func (c *MainController) Get() {

	// restart，clear session data,
	// such as http://127.0.0.1/?method=restart

	method := c.GetString("method","")

	if method == "restart" {
		c.DestroySession()
		c.Redirect("/", 302)
	}
	// end  restart

	// for debug ,use http://127.0.0.1/?level=n n is a number[0-finalLevel)
	//

	beginLevel, err := c.GetInt("level",1)
	if err == nil {
		if beginLevel != 1 {
			if beginLevel > FinalLevel {
				beginLevel  = FinalLevel
			}
			c.SetSession(LEVEL,beginLevel)
			c.Redirect("/", 302)
		}
	}
	// end for debug

	//  get the level
	level := c.GetSession(LEVEL).(int)

	//  setup  the level data

	switch level {
	case 1:
		// random password  such as  pass1234
		rand.Seed(time.Now().Unix())
		password := "pass" + strconv.Itoa(rand.Intn(10)) +
							strconv.Itoa(rand.Intn(10)) +
							strconv.Itoa(rand.Intn(10)) +
							strconv.Itoa(rand.Intn(10))
		c.SetSession(PASSWORD, password)
		c.Data[PASSWORD] = password
	case 5:
		//  set  cookie  data
		c.Ctx.SetCookie("login","no")
	case 7:
		//  get the  beego  listen port
		c.Data["port"] = beego.BConfig.Listen.HTTPPort
	default:

	}

	// layout  and  tpl files
	c.Layout = "layout.html"

	if level == FinalLevel { // final level use gameover.html as  tpl
		c.TplName = "gameover.html"
	} else {
		c.TplName = fmt.Sprintf("level%d.html", level)
	}

}

func (c *MainController) Post() {

	// get level
	level := c.GetSession(LEVEL).(int)

	var  success =  "恭喜您，通过第"+ strconv.Itoa(level) +"关"
	var failed = "闯关失败，请再尝试一下! Come on baby!"

	var  msg  =failed //  the default msg is failed

	switch level {
	case 0:
		postPassword,err := c.GetInt("password",0)
		if err ==  nil {
			if postPassword == 2 {
				msg = success
				c.SetSession(LEVEL,level+1)  // level up
			} else {
				msg = failed
			}
		}else {
			msg = "闯关失败, 请输入正确的数字 "
		}

	case 1:
		postPassword := c.GetString("password","")
		sessPassword := c.GetSession(PASSWORD)
		if postPassword == sessPassword {
			msg = success
			c.SetSession(LEVEL,level+1)
			c.DelSession(PASSWORD) //通关后删掉level 1 session 中的 password
		} else {
			msg = failed
		}
	case 2:
		postPassword := c.GetString("password","")
		if postPassword != "" {
			msg = success
			c.SetSession(LEVEL,level+1)
		} else {
			msg = failed
		}
	case 3:
		postPassword := c.GetString("password","")

		result := md5.Sum([]byte(postPassword))

		if  hex.EncodeToString(result[0:]) == "0192023a7bbd73250516f069df18b500" {
			msg = success
			c.SetSession(LEVEL,level+1)
		} else {
			msg = failed
		}
	case 4:
		postPassword := c.GetString("password","")

		result := md5.Sum([]byte(postPassword+"dyz"))

		if  hex.EncodeToString(result[0:]) == "6ddd502c96cdca7b63037bdd3441e783" {
			msg = success
			c.SetSession(LEVEL,level+1)
		} else {
			msg = failed
		}
	case 5:
		login := c.Ctx.GetCookie("login")
		if login == "yes" {
			msg = success
			c.SetSession(LEVEL,level+1)
			c.Ctx.SetCookie("login","",0)
		} else {
			msg = failed
		}
	case 6:
		useragent := c.Ctx.Request.UserAgent()

		postPassword := c.GetString("password","")

		if postPassword == "dyz123456" {
			if !strings.Contains(strings.ToLower(useragent),"windows") {
				msg = success
				c.SetSession(LEVEL,level+1)
			}else {
				msg = "系统禁止 windows 用户登录...O(∩_∩)O "
			}
		} else {
			msg = "密码错误，请继续努力 "
		}
	case 7:
		postPassword := c.GetString("password","")

		var port string

		fmt.Println(c.Ctx.Request.RemoteAddr)

		if  index := strings.LastIndex(c.Ctx.Request.RemoteAddr,":") ; index != -1 {
			port = c.Ctx.Request.RemoteAddr[index+1:] // 获取来源 ip 和 端口 ， 判断端口
		}
		if postPassword == port {
			msg = success
			c.SetSession(LEVEL,level+1)
		} else {
			msg = failed
		}
	default:
	}

	c.Data[MESSAGE]  = msg
	c.Layout  = "layout.html"
	c.TplName = "result.html"
}
