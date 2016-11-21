package main

import (
	"fmt"
	"github.com/juliancheal/smart_plug/tplink"
)

func main() {
	plug := tplink.HS110{IPAddress: "192.168.1.12"}
	results, err := plug.MeterInfo()
	if err != nil {
		fmt.Println("err:", err)
	}
	fmt.Println(results)

	powerStatus, err := plug.PowerStatus()
	if err != nil {
		fmt.Println("err:", err)
	}
	fmt.Println(powerStatus)
}
