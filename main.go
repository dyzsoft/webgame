package main

import (
	"github.com/astaxie/beego"
	_ "webgame/routers"
	"net/http"
	"fmt"
	"strconv"
)

func main() {

	beego.ErrorHandler("404", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(writer,`
		<html>
		<head>
    		<meta charset="UTF-8">
    		<title>Web安全丛林之旅-挑战的不仅仅是智商</title>
		</head>
	   	<h1>系统禁地，禁止访问！</h1>	
       	<p style="font-size: 16px"> <span id="times"> 3 </span> 秒钟后将自动跳转  <span id="wait"> .</span> </p>
       	<script type="text/javascript">
         let times = document.getElementById("times");
         let wait = document.getElementById("wait");
        setInterval(function () {
            timevalue = times.innerText;
            times.innerText =(parseInt(timevalue)- 1 )+ "";
            wait .innerText = wait.innerText + "..."
        },1000);
        setTimeout(function () {
            window.location = "/"
        },2900)
       </script>
			</html>
		`)
	})
	fmt.Println("##########################################")
	fmt.Println()

	fmt.Println("系统启动正在启动中.... Version 0.2")
	fmt.Println("系统启动完成,请打开浏览器访问 http://127.0.0.1:"+strconv.Itoa(beego.BConfig.Listen.HTTPPort))

	fmt.Println()
	fmt.Println("##########################################")

	//go func() {
	//
	//	cmd :=exec.Command("explorer.exe","http://127.0.0.1:8080")
	//	cmd.Run();
	//
	//}()

	beego.Run()
}
