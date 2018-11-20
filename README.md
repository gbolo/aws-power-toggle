# aws-power-toggle
web UI and API for quickly starting and stopping AWS environments

## Getting Started
- clone to `GOPATH` then `go get` dependencies.
- modify the config file: `testdata/sampleconfig/power-toggle-config.yaml`
- set aws API keys with following environment variables: `AWS_ACCESS_KEY_ID` and `AWS_SECRET_ACCESS_KEY`
- run the script: `./testdata/scripts/dev-start.sh`

## API examples
install `jq` to make your life easier

```
### List environments and their state
curl -s -X GET 127.0.0.1:8080/api/env | jq -r '.[] | "\(.Name): \(.State)"'

### Shutdown an environment (example kube)
curl -s -X POST 127.0.0.1:8080/api/env/powerdown/kube | jq .

### Refresh environment cache
curl -s -X POST 127.0.0.1:8080/api/env/refresh | jq .

### Start up and environment (example kube)
curl -s -X POST 127.0.0.1:8080/api/env/startup/kube | jq .
```