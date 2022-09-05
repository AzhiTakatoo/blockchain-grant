#!/bin/bash

priv_sk_path=$(ls ../crypto-config/peerOrganizations/org1.blockchaingrant.com/users/Admin@org1.blockchaingrant.com/msp/keystore/)

cp -rf ./connection-profile/network_temp.json ./connection-profile/network.json

sed -i "" "s/priv_sk/$priv_sk_path/" ./connection-profile/network.json

docker-compose down -v
docker-compose up -d