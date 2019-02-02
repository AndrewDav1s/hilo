#!/busybox/sh

set -ex

# import orchestrator key
gorc --config=/root/gorc/config.toml keys cosmos recover orchestrator "$HILO_E2E_ORCH_MNEMONIC"

# import Ethereum key
gorc --config=/root/gorc/config.toml keys eth import signer $HILO_E2E_ETH_PRIV_KEY

# start gorc orchestrator
gorc --config=/root/gorc/config.toml orchestrator start --cosmos-key=orchestrator --ethereum-key=signer
