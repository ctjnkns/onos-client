package main

import (
	"fmt"
	"os"

	"github.com/ctjnkns/onosclient"
)

func main() {
	host := os.Getenv("ONOS_HOST")
	username := os.Getenv("ONOS_USERNAME")
	password := os.Getenv("ONOS_PASSWORD")

	client, err := onosclient.NewClient(host, username, password)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(client)

}
