# Get Summary of ALL Environments

Retrieves a list of environments that are currently available to control

**URL** : `/api/v1/env/summary`

**Method** : `GET`

## Success Response

**Code** : `200 OK`

**Example Response Body**

```json
{
  "envList:": [
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
    },
    {
      "billsAccrued":"1.00",
      "billsSaved":"1.00",
      "id": "36436c027202",
      "name": "bench6",
      "provider": "aws",
      "region": "ca-central-1",
      "running_instances": 0,
      "state": "stopped",
      "stopped_instances": 53,
      "total_instances": 53,
      "total_memory_gb": 270,
      "total_vcpu": 155,
      "billsAccrued":"1.00",
      "billsSaved":"1.00"
    }
  ],
  "totalBillsAccrued":"2.00",
  "totalBillsSaved":"2.00",
}
```
