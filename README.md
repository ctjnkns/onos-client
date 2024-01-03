# ONOS Client
This is a Go client for interacting with the [ONOS](https://opennetworking.org/onos/) API. It was developed alongside the [ONOS Terraform Provider](https://github.com/ctjnkns/terraform-provider-onos).

## Example Using Docker Containers running on Linux (Ubuntu 22.04.3 LTS)
These examples require a current version of [go](https://go.dev/doc/install) and [docker](https://docs.docker.com/engine/install/ubuntu/).

### Docker Environment Setup
Navigate to the examples directory and run:

```bash
sudo docker compose up
```

This runs onos and mininet in docker containers and links them.  

From a new terminal, paste in the following curl commands to activate openflow and fwd in ONOS:

```bash
	curl --request POST \
	--url http://localhost:8181/onos/v1/applications/org.onosproject.fwd/active \
	--header 'Accept: application/json' \
	--header 'Authorization: Basic b25vczpyb2Nrcw=='

	curl --request POST \
	--url http://localhost:8181/onos/v1/applications/org.onosproject.openflow/active \
	--header 'Accept: application/json' \
	--header 'Authorization: Basic b25vczpyb2Nrcw=='
```

The calls should return output similar to this, indicating the activation was successful:

```json
{"name":"org.onosproject.fwd","id":78,"version":"2.7.1.SNAPSHOT","category":"Traffic Engineering","description":"Provisions traffic between end-stations using hop-by-hop flow programming by intercepting packets for which there are currently no matching flow objectives on the data plane.","readme":"Provisions traffic between end-stations using hop-by-hop flow programming by intercepting packets for which there are currently no matching flow objectives on the data plane. The paths paved in this manner are short-lived, i.e. they expire a few seconds after the flow on whose behalf they were programmed stops.\\n\\nThe application relies on the ONOS path service to compute the shortest paths. In the event of negative topology events (link loss, device disconnect, etc.), the application will proactively invalidate any paths that it had programmed to lead through the resources that are no longer available.","origin":"ONOS Community","url":"http://onosproject.org","featuresRepo":"mvn:org.onosproject/onos-apps-fwd/2.7.1-SNAPSHOT/xml/features","state":"ACTIVE","features":["onos-apps-fwd"],"permissions":[],"requiredApps":[]}

{"name":"org.onosproject.openflow","id":46,"version":"2.7.1.SNAPSHOT","category":"Provider","description":"Suite of the OpenFlow base providers bundled together with ARP\\/NDP host location provider and LLDP link provider.","readme":"Suite of the OpenFlow base providers bundled together with ARP\\/NDP host location provider and LLDP link provider.","origin":"ONOS Community","url":"http://onosproject.org","featuresRepo":"mvn:org.onosproject/onos-providers-openflow-app/2.7.1-SNAPSHOT/xml/features","state":"ACTIVE","features":["onos-providers-openflow-app"],"permissions":[],"requiredApps":["org.onosproject.hostprovider","org.onosproject.lldpprovider","org.onosproject.openflow-base"]}
```

Launch a terminal in the Mininet container:

```bash
sudo docker exec -it mininet-compose /bin/bash
```

Start a basic mininet topology with onos as the SDN controller:

```bash
mn --topo tree,2,2 --mac --switch ovs,protocols=OpenFlow14	  --controller remote,ip=onos-compose
```

Mininet should successfully connect to the ONOS controller, start the switches and hosts, and display the mininet CLI prompt:

```bash
root@c0b74ca1c8a7:~# mn --topo tree,2,2 --mac --switch ovs,protocols=OpenFlow14   --controller remote,ip=onos-compose
*** Error setting resource limits. Mininet's performance may be affected.
*** Creating network
*** Adding controller
Connecting to remote controller at onos-compose:6653
*** Adding hosts:
h1 h2 h3 h4 
*** Adding switches:
s1 s2 s3 
*** Adding links:
(s1, s2) (s1, s3) (s2, h1) (s2, h2) (s3, h3) (s3, h4) 
*** Configuring hosts
h1 h2 h3 h4 
*** Starting controller
c0 
*** Starting 3 switches
s1 s2 s3 ...
*** Starting CLI:
mininet> 
```

From the mininet CLI, run pingall and confirm that all hosts can communicate:
```bash
mininet> pingall
*** Ping: testing ping reachability
h1 -> h2 h3 h4 
h2 -> h1 h3 h4 
h3 -> h1 h2 h4 
h4 -> h1 h2 h3 
*** Results: 0% dropped (12/12 received)
```

Switch to an ubuntu terminal (not the mininet container) and paste in the following curl command to deactivate onos fwd:
```bash
curl -X DELETE --header 'Accept: application/json' 'http://localhost:8181/onos/v1/applications/org.onosproject.fwd/active'
```

This disabled reactive forwarding in the onos controller, causing to to behave more like a firewall than a router.

From the mininet CLI, run pingall again and confirm that the traffic is now being blocked:
```bash
mininet> pingall
*** Ping: testing ping reachability
h1 -> X X X 
h2 -> X X X 
h3 -> X X X 
h4 -> X X X 
*** Results: 100% dropped (0/12 received)
```

### Examples Demonstration 

#### Hosts
```bash
go run hosts-get-example.go
```

There should be several hosts listed:
```json
{[{00:00:00:00:00:03/None 00:00:00:00:00:03 None None 0x0000 false false [10.0.0.3] [{of:0000000000000003 1}]} {00:00:00:00:00:04/None 00:00:00:00:00:04 None None 0x0000 false false [10.0.0.4] [{of:0000000000000003 2}]} {00:00:00:00:00:01/None 00:00:00:00:00:01 None None 0x0000 false false [10.0.0.1] [{of:0000000000000002 1}]} {00:00:00:00:00:02/None 00:00:00:00:00:02 None None 0x0000 false false [10.0.0.2] [{of:0000000000000002 2}]}]}
```

#### Flows
```bash
go run flows-get-example.go
```

There should be flows representing connections between the hosts and switches:
```json
{[{org.onosproject.core 1078 of:0000000000000003 0 281478170942982 true 1704298704107 3100 UNKNOWN 11 5 ADDED 0 0 0 {[{0x800  0 ETH_TYPE}]} {true [] [{CONTROLLER OUTPUT}]}} {org.onosproject.core 139000 of:0000000000000003 0 281476661728682 true 1704298704107 3100 UNKNOWN 1000 40000 ADDED 0 0 0 {[{0x88cc  0 ETH_TYPE}]} {true [] [{CONTROLLER OUTPUT}]}} {org.onosproject.core 924 of:0000000000000003 0 281477764386537 true 1704298704107 3100 UNKNOWN 22 40000 ADDED 0 0 0 {[{0x806  0 ETH_TYPE}]} {true [] [{CONTROLLER OUTPUT}]}} {org.onosproject.core 139000 of:0000000000000003 0 281476156249461 true 1704298704107 3100 UNKNOWN 1000 40000 ADDED 0 0 0 {[{0x8942  0 ETH_TYPE}]} {true [] [{CONTROLLER OUTPUT}]}} {org.onosproject.core 798 of:0000000000000001 0 281478909873038 true 1704298704186 3100 UNKNOWN 19 40000 ADDED 0 0 0 {[{0x806  0 ETH_TYPE}]} {true [] [{CONTROLLER OUTPUT}]}} {org.onosproject.core 278000 of:0000000000000001 0 281477466379610 true 1704298704186 3100 UNKNOWN 2000 40000 ADDED 0 0 0 {[{0x88cc  0 ETH_TYPE}]} {true [] [{CONTROLLER OUTPUT}]}} {org.onosproject.core 980 of:0000000000000001 0 281475012051420 true 1704298704186 3100 UNKNOWN 10 5 ADDED 0 0 0 {[{0x800  0 ETH_TYPE}]} {true [] [{CONTROLLER OUTPUT}]}} {org.onosproject.core 278000 of:0000000000000001 0 281477029321583 true 1704298704186 3100 UNKNOWN 2000 40000 ADDED 0 0 0 {[{0x8942  0 ETH_TYPE}]} {true [] [{CONTROLLER OUTPUT}]}} {org.onosproject.core 882 of:0000000000000002 0 281478316350853 true 1704298704106 3100 UNKNOWN 21 40000 ADDED 0 0 0 {[{0x806  0 ETH_TYPE}]} {true [] [{CONTROLLER OUTPUT}]}} {org.onosproject.core 139000 of:0000000000000002 0 281478673389323 true 1704298704107 3100 UNKNOWN 1000 40000 ADDED 0 0 0 {[{0x8942  0 ETH_TYPE}]} {true [] [{CONTROLLER OUTPUT}]}} {org.onosproject.core 139000 of:0000000000000002 0 281475022575828 true 1704298704107 3100 UNKNOWN 1000 40000 ADDED 0 0 0 {[{0x88cc  0 ETH_TYPE}]} {true [] [{CONTROLLER OUTPUT}]}} {org.onosproject.core 1078 of:0000000000000002 0 281475568191580 true 1704298704107 3100 UNKNOWN 11 5 ADDED 0 0 0 {[{0x800  0 ETH_TYPE}]} {true [] [{CONTROLLER OUTPUT}]}}]}
```

If you do not see any flows, run pingall from the mininet terminal to refresh the flows table and try again.

#### Intents
```bash
go run intents-get-example.go
```

There should not be any intents listed since none have been created:
```json
{[]}
```

##### Create an Intent
```bash
go run intent-create-example.go
```

This creates an intent that allows h1 and h2 to communicate.
```json
{
    Type:  "HostToHostIntent",
    AppID: "org.onosproject.cli",
    Key:   "0x300009",
    One:   "00:00:00:00:00:01/None",
    Two:   "00:00:00:00:00:02/None",
}
```

The new intent should be returned from ONOS:
```json
{org.onosproject.cli 0xa 0x300009 INSTALLED HostToHostIntent [00:00:00:00:00:01/None 00:00:00:00:00:02/None] 0xc0000124b0 0xc000068580 100 [{false [OPTICAL] LinkTypeConstraint}] 00:00:00:00:00:01/None 00:00:00:00:00:02/None}
```

**Note: if the "One" and "Two" strings do not match any hosts from your hosts output above, the intent may fail to install. Edit the struct in intent-create-example.go to match two of the hosts in your environment.**

From the mininet CLI, test connectivity again and confirm that the traffic between h1 and h2 is now allowed:
```bash
mininet> h1 ping -c 4 h2
PING 10.0.0.2 (10.0.0.2) 56(84) bytes of data.
64 bytes from 10.0.0.2: icmp_seq=1 ttl=64 time=0.147 ms
64 bytes from 10.0.0.2: icmp_seq=2 ttl=64 time=0.038 ms
64 bytes from 10.0.0.2: icmp_seq=3 ttl=64 time=0.039 ms
64 bytes from 10.0.0.2: icmp_seq=4 ttl=64 time=0.038 ms

--- 10.0.0.2 ping statistics ---
4 packets transmitted, 4 received, 0% packet loss, time 3058ms
rtt min/avg/max/mdev = 0.038/0.065/0.147/0.047 ms
```

Confirm that traffic not allowed by the intent is still blocked:
```bash
mininet> h1 ping -c 4 h3
PING 10.0.0.3 (10.0.0.3) 56(84) bytes of data.

--- 10.0.0.3 ping statistics ---
4 packets transmitted, 0 received, 100% packet loss, time 3079ms
```

#### Get a specific Inent
```bash
go run intent-get-example.go
```

This looks up an intent with the following AppID and Key, and returns the details:
```json
{
    AppID: "org.onosproject.cli",
    Key:   "0x300009",
}
```

The intent details should be returned from ONOS:
```json
{org.onosproject.cli 0xa 0x300009 INSTALLED HostToHostIntent [00:00:00:00:00:01/None 00:00:00:00:00:02/None] 0xc0000124b0 0xc000068580 100 [{false [OPTICAL] LinkTypeConstraint}] 00:00:00:00:00:01/None 00:00:00:00:00:02/None}
```

#### Update the intent
```bash
go run intent-update-example.go
```

This updates the intent to allow h1 and h3 to communicate.
```json
{
    Type:  "HostToHostIntent",
    AppID: "org.onosproject.cli",
    Key:   "0x300009",
    One:   "00:00:00:00:00:01/None",
    Two:   "00:00:00:00:00:03/None",
}
```
Note: if the "One" and "Two" strings do not match any hosts from your hosts output above, edit the struct in intent-create-example.go.

From the mininet CLI, test connectivity again and confirm that the traffic between h1 and h2 is now blocked:
```bash
mininet> h1 ping -c 4 h2
PING 10.0.0.2 (10.0.0.2) 56(84) bytes of data.
From 10.0.0.1 icmp_seq=1 Destination Host Unreachable
From 10.0.0.1 icmp_seq=2 Destination Host Unreachable
From 10.0.0.1 icmp_seq=3 Destination Host Unreachable
From 10.0.0.1 icmp_seq=4 Destination Host Unreachable

--- 10.0.0.2 ping statistics ---
4 packets transmitted, 0 received, +4 errors, 100% packet loss, time 3069ms
pipe 4
```

Confirm that traffic from h1 to h3 is now allowed:
```bash
mininet> h1 ping -c 4 h3
PING 10.0.0.3 (10.0.0.3) 56(84) bytes of data.
64 bytes from 10.0.0.3: icmp_seq=1 ttl=64 time=0.512 ms
64 bytes from 10.0.0.3: icmp_seq=2 ttl=64 time=0.075 ms
64 bytes from 10.0.0.3: icmp_seq=3 ttl=64 time=0.042 ms
64 bytes from 10.0.0.3: icmp_seq=4 ttl=64 time=0.044 ms

--- 10.0.0.3 ping statistics ---
4 packets transmitted, 4 received, 0% packet loss, time 3059ms
rtt min/avg/max/mdev = 0.042/0.168/0.512/0.198 ms
```

#### Delete the intent
```bash
go run intent-delete-example.go
```

This looks up an intent with the following AppID and Key, and removes it:
```json
{
    AppID: "org.onosproject.cli",
    Key:   "0x300009",
}
```
Re-run intents-get-example.go or intent-get-example and confirm that the intent is no longer there.


### Basic Usage
Working examples are located in the examples direcotry. The folloiwng are snippets to demonstrate basic usage.

#### Creating a client
```go
const HostURL string = "http://localhost:8181/onos/v1"
username := "onos"
password := "rocks"

client, err := onosclient.NewClient(HostURL, username, password)
if err != nil {
    fmt.Println(err)
}
```

#### Creating a client using Environment Variables
```bash
#bash
export ONOS_HOST=http://localhost:8181/onos/
export ONOS_USERNAME=onos
export ONOS_PASSWORD=rocks
```

```go
host := os.Getenv("ONOS_HOST")
username := os.Getenv("ONOS_USERNAME")
password := os.Getenv("ONOS_PASSWORD")

client, err := onosclient.NewClient(host, username, password)
if err != nil {
    fmt.Println(err)
}
```

#### Get Hosts
```go
hosts, err := client.GetHosts()
if err != nil {
    fmt.Println(err)
}

fmt.Println(hosts)
```

#### Get Flows
```go
flows, err := client.GetFlows()
if err != nil {
    fmt.Println(err)
}

fmt.Println(flows)
```

#### Get Intents
```go
intents, err := client.GetIntents()
if err != nil {
    fmt.Println(err)
}

fmt.Println(intents)
```

#### Get a single Intent
The AppID and Key are required to lookup the intent in ONOS.

```go
intent := onosclient.Intent{
    AppID: "org.onosproject.cli",
    Key:   "0x300009",
}

intent, err = client.GetIntent(intent)
if err != nil {
    fmt.Println(err)
}
fmt.Println(intent)
```

#### Create an Intent

```go
intent := onosclient.Intent{
    Type:  "HostToHostIntent",
    AppID: "org.onosproject.cli",
    Key:   "0x300009",
    One:   "00:00:00:00:00:01/None",
    Two:   "00:00:00:00:00:02/None",
}

intent, err = client.CreateIntent(intent)
if err != nil {
    fmt.Println(err)
}
fmt.Println(intent)
```

#### Update an Intent
The AppID and Key must match an exisitng intent in ONOS.

```go
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
```

#### Delete an Intent
The AppID and Key are required to lookup the intent in ONOS.

```go
intent := onosclient.Intent{
    AppID: "org.onosproject.cli",
    Key:   "0x300009",
}

err = client.DeleteIntent(intent)
if err != nil {
    fmt.Println(err)
}
```