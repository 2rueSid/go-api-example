
#!/usr/bin/env bash

set -o errexit
set -o nounset

if [ ! $(command -v golangci-lint) ]
then
   sudo curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sudo sh -s -- -b $(go env GOPATH)/bin v1.41.1
    golangci-lint --version
fi

golangci-lint run

