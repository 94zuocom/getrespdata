package p

import (

	"net/http"

)


func Status502(w http.ResponseWriter, r *http.Request) {
        // 设置响应头的内容类型为 text/html; charset=utf-8
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		// 写入502状态码
		w.WriteHeader(http.StatusBadGateway)
		// 响应体中填充错误信息
		msg := `<html>
		<head>
			<title>502 Bad Gateway</title>
		</head>
		<body>
			<h1>502 Bad Gateway</h1>
			<p>抱歉，服务器遇到错误。</p>
		</body>
	</html>`
		w.Write([]byte(msg))
}
