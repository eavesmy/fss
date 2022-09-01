/*
# File Name: ./login.go
# Author : eavesmy
# Email:eavesmy@gmail.com
# Created Time: Thu Sep  1 10:08:04 2022
*/

package service

import (
	"net/http"
)

var Path string

type HttpHandler struct {
}

func (h HttpHandler) ServeHTTP(resp http.ResponseWriter, req *http.Request) {

	if req.URL.Query().Get("pwd") == "123qwe" {
		handler := http.FileServer(http.Dir(Path))
		handler.ServeHTTP(resp, req)
		return
	}

	html := `
<!doctype html>
<html>
<head>
    <meta charset="utf-8">
    <title>Verify Page</title>
		</head>
<body>
<script>
        window.location += "?pwd=" + prompt("Input password");
</script>
</body>


</html>
	`

	// Jump to login page.
	resp.Write([]byte(html))
}
