package main

import (
	"fmt"
	"hik_gateway/hik"
)

func main() {
	hik_gateway := hik.NewHikGateway()

	hik_gateway.RegisterHikGateway(hik.HikConfig{
		Ip:       "",
		Port:     80,
		Username: "",
		Password: "",
	})

	broadClient := make(chan hik.Msg)
	hik_gateway.RegisterBroadClient(broadClient)
	select {
	case msg := <-broadClient:
		{
			fmt.Println(msg)
		}
	}
}
