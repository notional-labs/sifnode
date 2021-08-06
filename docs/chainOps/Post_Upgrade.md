# Upgrading to Cosmos 0.42 on BetaNet (External Validator Operators)

This is for those validators running Kubernetes, who recently upgraded from Cosmos 0.39 to Cosmos 0.42.

The upgrade scripts are no longer required and as such, can be removed from your node.

1. Ensure that you're already cloned the `Sifchain/sifchain-deploy-public` repository into `./deploy`, in the root of where you checked out the main sifnode repository to.

2. Go into the `./deploy` folder and pull down the latest updates:

```bash
git checkout master && git pull
```

Please check that the latest commit hash is `e37729f5fa737c2b8b3411bf0c64f1373a046a2d`.

3. Deploy the latest helm updates to update your node:

```bash
export CLUSTER_NAME=<cluster_name>
export KUBECONFIG=./.live/${CLUSTER_NAME}/kubeconfig_${CLUSTER_NAME}

helm upgrade sifnode ../deploy/helm/standalone/sifnode \
--install -n sifnode --create-namespace \
--set sifnode.args.enableRpc="true" \
--set sifnode.args.enableExternalRpc="false" \
--set image.tag=mainnet-0.9.0 \
--set image.repository=sifchain/sifnoded
```

Replace `<cluster_name>` with the full name of your cluster.

4. Your node will restart and rejoin the network within 30-60 seconds.
