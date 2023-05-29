#!/usr/bin/env bash
# Scrip works with gqldenring version >= 1.1.0
# Usage ./scripts/dl_gqldenring
# Optionally specify a specific version ./scripts/dl_gqldenring v1.1.0

set -exo pipefail

OS=$(uname -o)
ARCH=$(uname -m)
GQLDENRING_DIR=gqldenring

ROOT_GITHUB_URL="https://api.github.com/repos/kamsandhu93/gqldenring/releases/"

if [ -z "$1" ] ; then
  VERSION=latest
  GITHUB_URL=${ROOT_GITHUB_URL}${VERSION}
else
  VERSION=$1
  GITHUB_URL=${ROOT_GITHUB_URL}"tags/"${VERSION}
fi

if [[ "$OS" == "GNU/Linux" ]] ; then
  OS="Linux"
fi

printf "[INFO] Expanding Vars OS=%s, ARCH=%s, VERSION=%s GITHUB_URL=%s\n" "$OS" "$ARCH" "$VERSION" "$GITHUB_URL"

ARTEFACT_URI=$(curl -s "$GITHUB_URL" \
  | grep "${OS}_${ARCH}" \
  | grep "https" \
  | cut -d : -f 2,3 \
  | tr -d \" \
  | tr -d " ")

printf "[INFO] ARTEFACT_URI=%s\n" "${ARTEFACT_URI}"

rm -f ${GQLDENRING_DIR}/gqldenring
rm -f ${GQLDENRING_DIR}/*.tar.gz || true

wget -cO - "${ARTEFACT_URI}" > ${GQLDENRING_DIR}/gqldenring_"${VERSION}".tar.gz

tar -C ${GQLDENRING_DIR} -zxvf ${GQLDENRING_DIR}/gqldenring_"${VERSION}".tar.gz gqldenring
