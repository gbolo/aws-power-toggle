# aws-power-toggle
web UI and API for quickly starting and stopping AWS environments

<img src="https://raw.githubusercontent.com/gbolo/aws-power-toggle/master/testdata/screenshots/v1-0-poc_mobile.png" width="432" />

## Getting Started
aws-power-toggle groups your instances by environments (via instance tag [described below](#Required-Tags)), then allows you power toggle
them with a single action. This can be very useful when you have to deal with a ton of developers who have multiple environments
and need to use them sporadically. While you could leave them on all the time, it may result in your boss(es) freaking out about the bill :)
Just hand this web UI and/or API to the devs and let them decide when to start/stop the environment(s).
Don't forget to check up on them to make sure they are actually turning them down when not in use :)

### Required Tags
The backend polls the aws API periodically (or on demand through web ui). In order for your instances to show up the following
`instance tags` are **required** (all other instances are ignored):

* `power-toggle-enabled` set to `true`
* `Environment` set to **non-empty** value

Both the tags listed above are configurable via the config file (see [power-toggle-config.yaml](testdata/sampleconfig/power-toggle-config.yaml)).
Instances are grouped by the value of `Environment` tag. Please note that tag values are **case-sensitive*.

### AWS API Key
the backend requires an API key to successfully poll AWS (*shock*). Once you have obtained it, set the following environment variables:`AWS_ACCESS_KEY_ID` and `AWS_SECRET_ACCESS_KEY`

### Running the docker image
Once you have tagged your AWS instances appropriately (hopefully with [terraform](https://www.terraform.io) or the aws cli) then your ready to deploy.
Ofcourse, this is done quickest via docker:
```
docker run -d --name "aws-power-toggle" \
 -p 8080:8080 \
 -e "AWS_ACCESS_KEY_ID=<your_key_id>" \
 -e "AWS_SECRET_ACCESS_KEY=<your_secret_key>" \
 gbolo/aws-power-toggle:1.1
```

Then open your browser to: [http://127.0.0.1:8080](http://127.0.0.1:8080)

## Developer Guide
ensure you have go and make installed (with `GOPATH` set) then follow these steps:
```
# clone source
mkdir -p ${GOPATH}/src/github.com/gbolo/aws-power-toggle
git clone https://github.com/gbolo/aws-power-toggle.git ${GOPATH}/src/github.com/gbolo/aws-power-toggle
cd ${GOPATH}/src/github.com/gbolo/aws-power-toggle

# make your changes to source code or config then export your aws api key
export AWS_ACCESS_KEY_ID=<your_key_id>
export AWS_SECRET_ACCESS_KEY=<your_secret_key>

# build it
make all

# run it
./bin/aws-power-toggle -config testdata/sampleconfig/power-toggle-config.yaml

# do a test API call
curl -v 127.0.0.1:8080/api/env/summary
```

### API Documentation
For further details on an API endpoint (including example responses), click on the endpoint's name.

* [EnvAllSummary](docs/api/env_all_summary.md): `GET /api/v1/env/summary` retrieves a summary of all known environments

* [EnvSummary](docs/api/env_summary.md): `GET /api/v1/env/{env-id}/summary` retrieves a summary of a single environment

* [EnvAllDetails](docs/api/env_all_details.md): `GET /api/v1/env/details` retrieves full details of all known environments (including list of instances)

* [EnvDetails](docs/api/env_details.md): `GET /api/v1/env/{env-id}/details` retrieves full details of a single environment (including list of instances)

* [StopEnv](docs/api/env_stop.md): `POST /api/v1/env/{env-id}/stop` triggers a shutdown of an environment

* [StartEnv](docs/api/env_start.md): `POST /api/v1/env/{env-id}/start` triggers a startup of an environment

* [StopInstance](docs/api/instance_stop.md): `POST /api/v1/instance/{instance-id}/stop` triggers a shutdown of a single instance

* [StartInstance](docs/api/instance_start.md): `POST /api/v1/instance/{instance-id}/start` triggers a startup of a single instance

* [Refresh](docs/api/refresh.md): `POST /api/v1/refresh` forces backend to refresh it's cache

* [Version](docs/api/version.md): `GET /api/v1/version` returns backend version information


#### TODO:
Need to finish these...

* [Config](docs/api/config.md): `GET /api/v1/config` returns relevant backend configuration

### Enabling AWS API mocking (web dev mode)
It may be useful to mock the aws API when doing development work against the API (like for web ui development).
This means you don't need an aws api key. To enable this feature:
```
# modify aws.go and set this constant to true:
MOCK_ENABLED = true

# set fake AWS API keys
export AWS_ACCESS_KEY_ID=DOESNT_MATTER
export AWS_SECRET_ACCESS_KEY=DOESNT_MATTER

# build then start
make all
./bin/aws-power-toggle -config testdata/sampleconfig/power-toggle-config.yaml
```

if you would like to add/remove/change any of the fake inventory, then modify this file:
`testdata/mock_env_cachedTable.json`