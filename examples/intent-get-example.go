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

	intent := onosclient.Intent{
		AppID: "org.onosproject.cli",
		Key:   "0x300009",
	}

	intent, err = client.GetIntent(intent)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(intent)
}
