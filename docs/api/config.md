# Get Configuration Information

Retrieves the current running configuration of the backend server.
This **excludes** sensitive information such as AWS API keys, exc.


**URL** : `/api/v1/config`

**Method** : `GET`

## Success Response

**Code** : `200 OK`

**Example Response Body**

```json
{
  "aws_environment_tag_key": "Environment",
  "aws_ignore_environments": [
    "prod"
  ],
  "aws_ignore_instance_types": [
    "c5d.large",
    "c5d.xlarge",
    "c5d.2xlarge",
    "c5d.4xlarge",
    "c5d.9xlarge",
    "c5d.18xlarge"
  ],
  "aws_max_instances_to_shutdown": 100,
  "aws_polling_interval": 5,
  "aws_regions": [
    "ca-central-1",
    "us-east-1"
  ],
  "aws_required_tag_key": "power-toggle-enabled",
  "aws_required_tag_value": "true",
  "mock_delay": true,
  "mock_enabled": false,
  "mock_errors": true,
  "slack_enabled": false
}
```
