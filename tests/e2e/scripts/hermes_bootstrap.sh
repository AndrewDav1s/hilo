#!/bin/bash

set -ex

# initialize Hermes relayer configuration
mkdir -p /root/.hermes/
touch /root/.hermes/config.toml

# setup Hermes relayer configuration
tee /root/.hermes/config.toml <<EOF

[global]
strategy = 'packets'
filter = false
log_level = 'info'
clear_packets_interval = 100
tx_confirmation = true

[rest]
enabled = true
host = '0.0.0.0'
port = 3031

[telemetry]
enabled = true
host = '127.0.0.1'
port = 3001

[[chains]]
id = '$HILO_E2E_HILO_CHAIN_ID'
rpc_addr = 'http://$HILO_E2E_HILO_VAL_HOST:26657'
grpc_addr = 'http://$HILO_E2E_HILO_VAL_HOST:9090'
websocket_addr = 'ws://$HILO_E2E_HILO_VAL_HOST:26657/websocket'
rpc_timeout = '10s'
account_prefix = 'hilo'
key_name = 'val01-hilo'
store_prefix = 'ibc'
max_gas = 3000000
gas_price = { price = 0.001, denom = 'photon' }
gas_adjustment = 0.1
clock_drift = '1m' # to accomdate docker containers
trusting_period = '14days'
trust_threshold = { numerator = '1', denominator = '3' }

[[chains]]
id = '$HILO_E2E_GAIA_CHAIN_ID'
rpc_addr = 'http://$HILO_E2E_GAIA_VAL_HOST:26657'
grpc_addr = 'http://$HILO_E2E_GAIA_VAL_HOST:9090'
websocket_addr = 'ws://$HILO_E2E_GAIA_VAL_HOST:26657/websocket'
rpc_timeout = '10s'
account_prefix = 'cosmos'
key_name = 'val01-gaia'
store_prefix = 'ibc'
max_gas = 3000000
gas_price = { price = 0.001, denom = 'stake' }
gas_adjustment = 0.1
clock_drift = '1m' # to accomdate docker containers
trusting_period = '14days'
trust_threshold = { numerator = '1', denominator = '3' }
EOF

# import gaia and hilo keys
hermes keys restore ${HILO_E2E_GAIA_CHAIN_ID} -n "val01-gaia" -m "${HILO_E2E_GAIA_VAL_MNEMONIC}"
hermes keys restore ${HILO_E2E_HILO_CHAIN_ID} -n "val01-hilo" -m "${HILO_E2E_HILO_VAL_MNEMONIC}"

# start Hermes relayer
hermes start
