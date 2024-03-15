package p

import (
	"fmt"
	// "html"
	// "io"
	"io/ioutil"
	"net/http"
	"strconv"
	"encoding/json"
	"time"
	// "strings"
)

// TaobaoResponse 定义了从淘宝API返回的JSON数据结构
type TaobaoResponse struct {
	Data struct {
		T string `json:"t"` // 时间戳字段
	} `json:"data"`
}

func TimeHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "text/html; charset=utf-8")

	resp, err := http.Get("https://api.m.taobao.com/rest/api3.do?api=mtop.common.getTimestamp")
	if err != nil {
		http.Error(w, "Failed to get timestamp from Taobao", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, "Failed to read response body", http.StatusInternalServerError)
		return
	}

	var taobaoResp TaobaoResponse
	if err := json.Unmarshal(body, &taobaoResp); err != nil {
		http.Error(w, "Failed to unmarshal JSON", http.StatusInternalServerError)
		return
	}

	serverTime, err := strconv.ParseInt(taobaoResp.Data.T, 10, 64)
	if err != nil {
		http.Error(w, "Failed to parse timestamp", http.StatusInternalServerError)
		return
	}

	// 淘宝API返回的时间戳是毫秒级的，转换为秒级
	beijingTime := time.Unix(serverTime/1000, 0).In(time.FixedZone("CST", 8*3600))



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
            displayLocalTime();
			document.write("<p>本地电脑时间: " + new Date().toLocaleString() + "</p>");
        </script>
		<p>淘宝服务器时间（北京时间）: %s</p>

    </body>
    </html>
    `
    fmt.Fprintln(w, js, beijingTime.Format(time.RFC1123))
}
