#!/usr/bin/env sh

# This script defines a wrapper around starting an hilod process that is used
# in local network integration testing. Typically, a Docker image is built all
# necessary depencies to run the hilod process and which includes this script
# with a default command, where the entrypoint is the execution of this script.
# 
# Via docker-compose, you can start a local network by building the appropriate
# hilod binary and linking it via an environment variable.

BINARY=/hilod/${BINARY:-hilod}
ID=${ID:-0}

if ! [ -f "${BINARY}" ]; then
	echo "The binary $(basename "${BINARY}") cannot be found.
Please add the binary to the shared folder. Use the BINARY environment variable
if the name of the binary is not 'hilod'."
	exit 1
fi

BINARY_CHECK="$(file "$BINARY" | grep 'ELF 64-bit LSB executable, x86-64')"

if [ -z "${BINARY_CHECK}" ]; then
	echo "Binary needs to be OS linux, ARCH amd64"
	exit 1
fi

export HILODHOME="/hilod/node${ID}/hilod"

"${BINARY}" --home "${HILODHOME}" "$@"
