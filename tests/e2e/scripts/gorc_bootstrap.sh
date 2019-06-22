#!/busybox/sh

set -ex

# import orchestrator Hilo key
gorc --config=/root/gorc/config.toml keys cosmos recover orch-hilo-key "$HILO_E2E_ORCH_MNEMONIC"

# import orchestrator Ethereum key
gorc --config=/root/gorc/config.toml keys eth import orch-eth-key $HILO_E2E_ETH_PRIV_KEY

# start gorc orchestrator
gorc --config=/root/gorc/config.toml orchestrator start --cosmos-key=orch-hilo-key --ethereum-key=orch-eth-key
