package p

import (

	"net/http"

)

// 测试数据

// ServeStatic 定义静态文件服务
func ServeStatic(routePrefix, staticDir string) {
	fs := http.FileServer(http.Dir(staticDir))
	http.Handle(routePrefix, http.StripPrefix(routePrefix, fs))
}

// IndexHandler 处理根目录请求，重定向到index.html
func IndexHandler(staticDir string) http.HandlerFunc {
	// return func(w http.ResponseWriter, r *http.Request) {
	// 	http.ServeFile(w, r, staticDir+"/index.html")
	// }
	return func(w http.ResponseWriter, r *http.Request) {
		// Only serve the index file if the path is exactly "/"
		if r.URL.Path == "/" {
			http.ServeFile(w, r, staticDir+"/index.html")
		} else {
			// Otherwise, let the FileServer handler deal with the request
			http.NotFound(w, r)
		}
	}
}