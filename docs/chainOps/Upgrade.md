# Upgrading to Cosmos 0.42 on BetaNet (External Validator Operators)

## Standalone

_These instructions assume that you're running our docker standalone solution_.

### Upgrade

1. Halt your node.

2. Open up a terminal to your node using `docker run`. You must switch to the parent directory where the node's `.sifnoded` and `.sifnodecli` directories live, and run:

```bash
docker run -v `pwd`:/root -it sifchain/sifnoded:mainnet-genesis /bin/sh
```

3. Take a backup of the `config` and `data` directories:

```bash
BACKUP_DIR="${HOME}"/.sifnoded/backups/$(date +%s%N)/
mkdir -p "${BACKUP_DIR}"
cp -avr "${HOME}"/.sifnoded/data/ "${BACKUP_DIR}"
cp -avr "${HOME}"/.sifnoded/config/ "${BACKUP_DIR}"
```

4. Download the new binary:

```bash
mkdir -p "${HOME}"/.sifnoded/cosmovisor/upgrades/0.9.0/bin
cd "${HOME}"/.sifnoded/cosmovisor/upgrades/0.9.0/bin
wget -O sifnoded.zip https://github.com/Sifchain/sifnode/releases/download/mainnet-0.9.0/sifnoded-mainnet-0.9.0-linux-amd64.zip
```

5. Very the hash:

```bash
sha256sum sifnoded.zip
```

and ensure that it matches:

```
2c82a37c62e185b9211a01fe831d35beb4b01fb7438fee555dd61e3c47c0cf8c
```

6. Unzip the binary:

```bash
unzip sifnoded.zip
```

7. Reset the state of your node:

```bash
"${HOME}"/.sifnoded/cosmovisor/upgrades/0.9.0/bin/sifnoded unsafe-reset-all
```

8. Download the updated `app.toml`:

```bash
wget -O "${HOME}"/.sifnoded/config/app.toml https://raw.githubusercontent.com/Sifchain/networks/master/config/sifchain-1/app.toml
```

9. Download the new genesis:

```bash
wget -O "${HOME}"/.sifnoded/config/genesis.json.gz https://raw.githubusercontent.com/Sifchain/networks/master/mainnet/sifchain-1/genesis.json.gz
gunzip "${HOME}"/.sifnoded/config/genesis.json.gz
```

10. Verify the genesis hash:

```bash
jq -S -c -M '.' genesis.json | sha256sum
```

and ensure that it matches:

```
42f07301312f6211bb95d8eb87d34770ad0ed712932e4c99dea7b01dd493dc1a
```

11. Update the `config.toml`:

```bash
sed -ri 's/log_level.*/log_level = \"info\"/g' "${HOME}"/.sifnoded/config/config.toml
sed -ri 's/persistent_peers.*/persistent_peers = \"87688830f890e5374fd4638942397a65d05f703b@13.213.156.252:26656\"/g' "${HOME}"/.sifnoded/config/config.toml
```

12. Tell cosmovisor to use the new binary:

```bash
rm "${HOME}"/.sifnoded/cosmovisor/current
ln -s "${HOME}"/.sifnoded/cosmovisor/upgrades/0.9.0 "${HOME}"/.sifnoded/cosmovisor/current
```

13. Exit the shell and restart your node. It'll now connect and start synchronising.

## k8s

### Pre-upgrade

Like the standalone solution above, Sifchain will publish the time at which your node should halt, via Discord. 

1. Checkout the latest code:

```bash
git checkout master && git pull
```

2. Checkout the new `sifnode-deploy-public` repository (from the root of where you checked out the sifnode repository to), as these rake tasks and helm charts have been moved to a new repository:

```bash
git clone ssh://git@github.com/Sifchain/sifchain-deploy-public ./deploy
```

3. To update your node so that it halts at the correct time, run (from the root of the sifnode repository):

```bash
export CLUSTER_NAME=<cluster_name>
export KUBECONFIG=./.live/${CLUSTER_NAME}/kubeconfig_${CLUSTER_NAME}

helm upgrade sifnode ./deploy/helm/standalone/sifnode \
--set sifnode.env.chainnet=sifchain \
--install -n sifnode --create-namespace \
--set sifnode.args.enableRpc="true" \
--set sifnode.args.enableExternalRpc="false" \
--set sifnode.upgrade.timestamp="<timestamp>" \
--set image.tag=mainnet-genesis \
--set image.repository=sifchain/sifnoded
```

Replace `<cluster_name>` with the full name of your cluster and `<timestamp>` with the value provided by Sifchain in Discord.

### Upgrade

1. Once the timestamp has been reached, your node will halt and sit idle (this will also be announced on Discord). You will then need to shut down the node completely:

```bash
export CLUSTER_NAME=<cluster_name>
export KUBECONFIG=./.live/${CLUSTER_NAME}/kubeconfig_${CLUSTER_NAME}

kubectl scale deployment sifnode --replicas=0 -n sifnode
```

As above, replace `<cluster_name>` with the name of your cluster.

2. To check your node has terminated, run:

```bash
kubectl get pods -n sifnode
```

3. Once terminated, you can perform the upgrade by running:

```bash
helm upgrade sifnode ./deploy/helm/standalone/sifnode \
--install -n sifnode --create-namespace \
--set sifnode.args.enableRpc="true" \
--set sifnode.args.enableExternalRpc="false" \
--set sifnode.args.upgradeNode="true" \
--set sifnode.upgrade.initialHeight="<initial_height>" \
--set sifnode.upgrade.chainnet="sifchain-1" \
--set sifnode.upgrade.cosmosSDKVersion="0.40" \
--set sifnode.upgrade.genesisTime="<genesis_time>" \
--set sifnode.upgrade.version="0.9.0" \
--set sifnode.upgrade.currentVersion="0.8.8" \
--set sifnode.upgrade.dataMigrateVersion="0.9" \
--set image.tag=mainnet-0.9.0 \
--set image.repository=sifchain/sifnoded
```

Replace `<initial_height>` and `<genesis_time>` with the values provided by Sifchain in Discord.

4. The node should then perform the upgrade (data export, migration and restart) and reconnect back into the network.
