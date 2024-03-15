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

            // 使用fetch API获取服务器时间并显示
            const displayServerTime = () => {
                fetch("https://api.m.taobao.com/rest/api3.do?api=mtop.common.getTimestamp")
                .then(response => response.json())
                .then(data => {
                    const serverTime = parseInt(data.data.t);
                    const date = new Date(serverTime + (8 * 3600 * 1000)); // 转换到北京时间
                    document.body.innerHTML += "<p>服务器时间（北京时间）: " + date.toLocaleString() + "</p>";
                })
                .catch(error => console.error('Error:', error));
            }

            displayLocalTime();
            displayServerTime();
        </script>
    </body>
    </html>
    `
    fmt.Fprintln(w, js)
}
