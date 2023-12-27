#!/bin/sh
docker compose up &
curl --request POST \
    --url http://localhost:8181/onos/v1/applications/org.onosproject.fwd/active \
    --header 'Accept: application/json' \
    --header 'Authorization: Basic b25vczpyb2Nrcw=='
curl --request POST \
    --url http://localhost:8181/onos/v1/applications/org.onosproject.openflow/active \
    --header 'Accept: application/json' \
    --header 'Authorization: Basic b25vczpyb2Nrcw=='
docker exec -it mininet /bin/bash
mn --topo tree,2,2 --mac --switch ovs,protocols=OpenFlow14	  --controller remote,ip=onos
