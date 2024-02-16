#!/usr/bin/env bash

# This script should be run from the root of the repository (./scripts/generateClient.sh)

set -x
set -e

# First remove any existing files that were generated. If we want to edit files post-generation, this section will need to be removed
# rm ./api_*.go
# rm ./model_*.go
# rm -r docs/

docker run --rm \
    -v "$PWD:/local" openapitools/openapi-generator-cli:v7.3.0 generate \
    -i /local/api/pf-swagger.yaml \
    -g go \
    -o /local/configurationapi \
    --git-host github.com \
    --git-repo-id pingfederate-go-client \
    --git-user-id pingidentity \
    --type-mappings=integer=int64,number=float64 \
    --additional-properties=enumClassPrefix=true,packageName=configurationapi,useOneOfDiscriminatorLookup=true \
    --skip-validate-spec

rm -r configurationapi/test/

# Run any code generators
go mod tidy
go generate ./...

# Clean things up
go fmt ./...
