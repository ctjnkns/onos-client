package main

import (
	"fmt"

	"github.com/ctjnkns/onosclient"
)

func main() {
	const HostURL string = "http://localhost:8181/onos/v1"
	username := "onos"
	password := "rocks"

	client, err := onosclient.NewClient(HostURL, username, password)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(client)

}
