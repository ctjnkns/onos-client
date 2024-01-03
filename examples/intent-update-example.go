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
		Type:  "HostToHostIntent",
		AppID: "org.onosproject.cli",
		Key:   "0x300009",
		One:   "00:00:00:00:00:01/None",
		Two:   "00:00:00:00:00:03/None",
	}

	intent, err = client.UpdateIntent(intent)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(intent)
}
