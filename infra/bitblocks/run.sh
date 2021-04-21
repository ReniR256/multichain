#!/bin/bash
ADDRESS=$1

# Start
/app/bin/bitblocksd -conf=/root/.bitblocks/bitblocks.conf # -server -rpcbind=0.0.0.0 -rpcallowip=0.0.0.0/0 -rpcuser=user -rpcpassword=password
sleep 10

# Print setup
echo "BITBLOCKS_ADDRESS=$ADDRESS"

# Import the address
/app/bin/bitblocks-cli importaddress $ADDRESS

# Generate enough block to pass the maturation time
/app/bin/bitblocks-cli generatetoaddress 101 $ADDRESS

# Simulate mining
while :
do
    /app/bin/bitblocks-cli generatetoaddress 1 $ADDRESS
    sleep 10
done
