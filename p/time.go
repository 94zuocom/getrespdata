package p

import (
	"fmt"
	// "html"
	// "io"
	// "io/ioutil"
	"net/http"
	// "strconv"
	// "strings"
)


func TimeHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "text/html; charset=utf-8")
    js := `
    <!DOCTYPE html>
    <html>
    <head>
        <title>Time Page</title>
    </head>
    <body>
        <script>
            // 显示电脑当前时间戳转换成北京时间
            const displayLocalTime = () => {
                const now = new Date();
                const localTime = now.getTime() + (now.getTimezoneOffset() * 60000) + (8 * 3600 * 1000); // 转换到北京时间
                const localDate = new Date(localTime);
                document.body.innerHTML += "<p>本地电脑时间（北京时间）: " + localDate.toLocaleString() + "</p>";
            }

            // 使用JSONP获取服务器时间并显示
            const displayServerTime = () => {
                const script = document.createElement('script');
                script.src = "https://api.m.taobao.com/rest/api3.do?api=mtop.common.getTimestamp&callback=handleResponse";
                document.body.appendChild(script);
            }

            // 定义处理响应的函数
            window.handleResponse = (response) => {
                const serverTime = parseInt(response.data.t);
                const date = new Date(serverTime + (8 * 3600 * 1000)); // 转换到北京时间
                document.body.innerHTML += "<p>服务器时间（北京时间）: " + date.toLocaleString() + "</p>";
            }

            displayLocalTime();
            displayServerTime();
        </script>
    </body>
    </html>
    `
    fmt.Fprintln(w, js)
}
