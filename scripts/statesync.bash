
#!/bin/bash
# microtick and bitcanna contributed significantly here.
set -uxe

# set environment variables
export GOPATH=~/go
export PATH=$PATH:~/go/bin


# Install Gaia
go install -tags badgerdb ./...


# MAKE HOME FOLDER AND GET GENESIS
iris init test 
wget -O ~/.iris/config/genesis.json https://cloudflare-ipfs.com/ipfs/QmfZoTRfuao6MU6Cns4ZpHDVrKMfbwjKDUEDH59U5Gweky

INTERVAL=1000

# GET TRUST HASH AND TRUST HEIGHT

LATEST_HEIGHT=$(curl -s http://https://rpc.sifchain.finance/block | jq -r .result.block.header.height);
BLOCK_HEIGHT=$(($LATEST_HEIGHT-$INTERVAL)) 
TRUST_HASH=$(curl -s "http://https://rpc.sifchain.finance/block?height=$BLOCK_HEIGHT" | jq -r .result.block_id.hash)


# TELL USER WHAT WE ARE DOING
echo "TRUST HEIGHT: $BLOCK_HEIGHT"
echo "TRUST HASH: $TRUST_HASH"


# exort state sync vars
export SIFNODED_STATESYNC_ENABLE=true
export SIFNODED_P2P_MAX_NUM_OUTBOUND_PEERS=200
export SIFNODED_STATESYNC_RPC_SERVERS="https://rpc.sifchain.finance:443,https://rpc.sifchain.finance:443"
export SIFNODED_STATESYNC_TRUST_HEIGHT=$BLOCK_HEIGHT
export SIFNODED_STATESYNC_TRUST_HASH=$TRUST_HASH
export SIFNODED_P2P_PERSISTENT_PEERS="71667eb24e7ad01925d9c75689e0b3e802faeca2@3.250.251.12:26656"

sifnoded start --x-crisis-skip-assert-invariants --db_backend badgerdb

