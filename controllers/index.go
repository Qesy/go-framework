package controllers

import (
	"fmt"
	"net"
)

// Index_index 首页
func (e *Entry) Index_index() {
	ip, _, _ := net.SplitHostPort(e.R.RemoteAddr)
	fmt.Fprintf(e.W, "<br><br><br><h1><center> API SERVER INTERFACE 1.0.0 </center></h1><h2><center>GOLANG Version 1.0</center></h2><center><p>Your IP:"+ip+"</p></center>")
}
