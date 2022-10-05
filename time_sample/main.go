package main

import (
	"fmt"
	"net/http"
	"strconv"
	"time"
)

func main() {
	t := time.Now()
	fmt.Println(t)                         // 2022-09-26 11:01:32.1806072 +0800 CST m=+0.001637001
	fmt.Println(t.Format(time.RFC3339))    // 2022-09-26T11:03:08+08:00
	fmt.Println(t.Format(http.TimeFormat)) // Mon, 26 Sep 2022 11:04:00 GMT

	fmt.Println(t.Unix())     // 1664161530 // second
	fmt.Println(t.UnixNano()) // 1664161551799919000 // micro second

	// timestatmp to time.Time
	timestamp, _ := strconv.ParseInt("1597215563", 10, 64)
	t1 := time.Unix(timestamp, 0)
	fmt.Println(t1)

	// ISO8601 to time.Time
	t2, _ := time.Parse(time.RFC3339, "2020-08-12T07:09:44.975Z")
	fmt.Println(t2)
}
