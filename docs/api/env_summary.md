# Get Summary of a SINGLE Environment

Retrieves a summary of the specified environment

**URL** : `/api/v1/env/{env-id}/summary`

**Method** : `GET`

## Success Response

**Code** : `200 OK`

**Example Response Body**

response of request: `/api/v1/env/931decfe6fd5/summary`

```json
{
  "id": "931decfe6fd5",
  "name": "kube",
  "provider": "aws",
  "region": "ca-central-1",
  "running_instances": 0,
  "state": "stopped",
  "stopped_instances": 10,
  "total_instances": 10,
  "total_memory_gb": 94,
  "total_vcpu": 30,
  "billsAccrued":"1.00",
  "billsSaved":"1.00"
}
```
