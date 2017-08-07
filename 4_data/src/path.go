package main

import (
	"fmt"
	"net/url"
)

func main() {
	u,err:=url.Parse("http://news.zone.mn")
	if err==nil {
		fmt.Println(u.Host)
	} else {
		fmt.Println(err)
	}
}