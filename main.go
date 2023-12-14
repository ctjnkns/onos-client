package main

import (
	"fmt"
)

func main() {
	// HostURL - Default Onos URL
	const HostURL string = "http://localhost:8181/onos/v1"

	username := "onos"
	password := "rocks"

	client, err := NewClient(HostURL, username, password)
	if err != nil {
		fmt.Println(err)
	}

	intents, err := client.GetIntents()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Sprintf("%q", intents)

	intent := Intent{
		Type:  "HostToHostIntent",
		AppID: "org.onosproject.cli",
		One:   "00:00:00:00:00:01/None",
		Two:   "00:00:00:00:00:03/None",
	}
	fmt.Println("\nCreate intent")
	fmt.Println(intent)
	err = client.CreateIntent(intent)
	if err != nil {
		fmt.Println(err)
	}

	intents, err = client.GetIntents()
	fmt.Println("\nGetting intents")
	fmt.Println(intents)

	intent = Intent{
		AppID: intents.Intent[0].AppID,
		Key:   intents.Intent[0].Key,
	}

	fmt.Println("\nDelete Intent")
	fmt.Println(intent)
	err = client.DeleteIntent(intent)
	if err != nil {
		fmt.Println(err)
	}

	intents, err = client.GetIntents()
	fmt.Println("\nGetting intents")
	fmt.Println(intents)

	intent = Intent{
		AppID: intents.Intent[0].AppID,
		Key:   intents.Intent[0].Key,
	}

	intent, err = client.GetIntent(intent)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("\nGetting intent")
	fmt.Println(intent)
	fmt.Println(intent.Type)

}
