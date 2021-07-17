#!/usr/bin/bash

set -eux
set -o pipefail

SERVERPORT=8080
SERVERADDR=localhost:${SERVERPORT}
GQLPATH=http://${SERVERADDR}/query

# Clean up the task store initially
curl ${GQLPATH} \
    -w "\n" -H 'Content-Type: application/json' \
    --data-binary '{"query":"mutation {\n  DeleteAllProvinces\n}\n"}'

curl ${GQLPATH} \
    -w "\n" -H 'Content-Type: application/json' \
    --data-binary '{"query":"mutation {\n  CreateProvince(input:\n    {Id:\"1\",\n     Name:\"Stockholm\",\n      admin_dev: 5,\n   dip_dev: 5,\n     mil_dev: 3,\n      trade_good:\"grain\",\n  trade_node:\"Baltic Sea\",\n  modifiers:[\"Entrepot\"]\n })\n  {\n    Id\n  }\n}"}'

curl ${GQLPATH} \
    -w "\n" -H 'Content-Type: application/json' \
    --data-binary '{"query":"mutation {\n  CreateProvince(input:\n    {Id:\"2\",\n     Name:\"Osterotland\",\n      admin_dev: 2,\n   dip_dev: 2,\n     mil_dev: 2,\n      trade_good:\"grain\",\n  trade_node:\"Baltic Sea\",\n  modifiers:[]\n })\n  {\n    Id\n  }\n}"}'

curl ${GQLPATH} \
    -w "\n" -H 'Content-Type: application/json' \
    --data-binary '{"query":"mutation {\n  CreateProvince(input:\n    {Id:\"3\",\n     Name:\"Kalmar\",\n      admin_dev: 2,\n   dip_dev: 2,\n     mil_dev: 1,\n      trade_good:\"naval supplies\",\n  trade_node:\"Baltic Sea\",\n  modifiers:[]\n })\n  {\n    Id\n  }\n}"}'