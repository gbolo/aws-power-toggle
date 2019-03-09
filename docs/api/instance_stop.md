# Initialize a Shutdown of a SINGLE Instance

Sends a shutdown request to the AWS API for the specified instance id.

`{instance-id}` is the **aws-power-toggle internal `id`, NOT the AWS `instance_id`**

**URL** : `/api/v1/instance/{instance-id}/stop`

**Method** : `POST`

## Success Response

**Code** : `200 OK`

## Example Request

Given the following instance:
```json
{
  "id": "1aef6299109b",
  "instance_id": "i-0008ad1bfd83a52eb",
  "instance_type": "t3.xlarge",
  "name": "kube-k8node2",
  "state": "running",
  "environment": "kube",
  "vcpu": 4,
  "memory_gb": 16,
  "pricing": 0.1856
}
```

Use the `id` field (**NOT the `instance_id` field**) to place the following request:
`/api/v1/instance/1aef6299109b/stop`

## Notes

Responses vary depending on upstream AWS API
since we pass the exact aws response.
**This will be addressed in an upcomming release for consistency.**
