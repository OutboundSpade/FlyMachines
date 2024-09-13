#!/bin/bash

docker run --rm -v "${PWD}:/local" openapitools/openapi-generator-cli generate \
    -i /local/spec.json \
    -g go \
    -o /local/ \
    --package-name=machines \
    -p packageName=machines \
    --skip-validate-spec \
    --git-repo-id FlyMachines --git-user-id OutboundSpade

go mod init github.com/OutboundSpade/FlyMachines
go mod tidy
