# Upgrading to v0.9.5 on BetaNet.

## Overview

Version 0.9.5 of BetaNet includes an update to the Cosmos SDK that includes a recent fix for upgrades. In some instances, an upgrade will fail to be performed and will exit as follows:

```bash
panic: unexpected end of JSON input

goroutine 1 [running]:
github.com/Sifchain/sifnode/app.SetupHandlersForV095(0xc000eb1000)
        /github/workspace/app/setup_handlers.go:26 +0x455
github.com/Sifchain/sifnode/app.SetupHandlers(...)
        /github/workspace/app/setup_handlers.go:14
github.com/Sifchain/sifnode/app.NewSifApp(0x21b0da0, 0xc001233f20, 0x21c8fe0, 0xc0010e2050, 0x0, 0x0, 0x1, 0xc0012e1d10, 0xc000fa99a0, 0xf, ...)
        /github/workspace/app/app.go:319 +0x2bdd
github.com/Sifchain/sifnode/cmd/sifnoded/cmd.newApp(0x21b0da0, 0xc001233f20, 0x21c8fe0, 0xc0010e2050, 0x0, 0x0, 0x216ed00, 0xc0010307e0, 0xc001035400, 0x1d3194a)
        /github/workspace/cmd/sifnoded/cmd/root.go:180 +0x9cf
[...]
```

If this occurs, then the fix is relatively straight forward:

### Standalone/Docker

Create the file `~/.sifnoded/data/upgrade-info.json` with the following content:

```json
{"name":"0.9.5","height":2976500}
```

and then restart your node.

### k8s

If using Kubernetes, you can repair your node via helm. 

1. Simply pull down the latest updates from the `sifchain-deploy-public` repository (which should be sitting inside the `deploy` folder, if you followed our instructions [here](https://github.com/Sifchain/sifnode/blob/master/docs/chainOps/k8s/README.md#prerequisites--dependencies)), e.g.:

```
cd deploy
git checkout master && git pull
```

2. Fix your node by running:

```
helm upgrade sifnode ./helm/standalone/sifnode \
--install -n sifnode-api --create-namespace \
--set sifnode.args.upgrade.fix="true" \
--set sifnode.args.upgrade.version="0.9.5" \
--set sifnode.args.upgrade.height=2976500 \
--set sifnode.args.enableRpc="true" \
--set sifnode.args.enableExternalRpc="false" \
--set image.tag=mainnet-0.9.0 \
--set image.repository=sifchain/sifnoded
```
