# Get Details of ALL Environments

retrieves full details of all known environments (including list of instances)

**URL** : `/api/v1/env/details`

**Method** : `GET`

## Success Response

**Code** : `200 OK`

**Example Response Body**

```json
{
	"envList":[
		{
			"id": "931decfe6fd5",
			"provider": "aws",
			"region": "ca-central-1",
			"name": "kube",
			"instances": [
				{
					"id": "1aef6299109b",
					"instance_id": "i-0008ad1bfd83a52eb",
					"instance_type": "t3.xlarge",
					"name": "kube-k8node2",
					"state": "stopped",
					"environment": "kube",
					"vcpu": 4,
					"memory_gb": 16
				},
				{
					"id": "9b97d53ab004",
					"instance_id": "i-04b69b4ec92d548df",
					"instance_type": "t3.medium",
					"name": "kube-k8master1",
					"state": "stopped",
					"environment": "kube",
					"vcpu": 2,
					"memory_gb": 4
				},
				{
					"id": "be8225018847",
					"instance_id": "i-0dedc66553eaecfc3",
					"instance_type": "t3.medium",
					"name": "kube-k8master2",
					"state": "stopped",
					"environment": "kube",
					"vcpu": 2,
					"memory_gb": 4
				},
				{
					"id": "dd4c000552b2",
					"instance_id": "i-0778ddf78ec72bffd",
					"instance_type": "t3.small",
					"name": "kube-etcd3",
					"state": "stopped",
					"environment": "kube",
					"vcpu": 2,
					"memory_gb": 2
				},
				{
					"id": "63296c73e207",
					"instance_id": "i-03b3782294e848647",
					"instance_type": "t3.small",
					"name": "kube-etcd1",
					"state": "stopped",
					"environment": "kube",
					"vcpu": 2,
					"memory_gb": 2
				},
				{
					"id": "a2f4635a834b",
					"instance_id": "i-0872ea25c0766d383",
					"instance_type": "t3.xlarge",
					"name": "kube-k8node5",
					"state": "stopped",
					"environment": "kube",
					"vcpu": 4,
					"memory_gb": 16
				},
				{
					"id": "98a9867235f3",
					"instance_id": "i-0e9fcaf646ce779ab",
					"instance_type": "t3.xlarge",
					"name": "kube-k8node3",
					"state": "stopped",
					"environment": "kube",
					"vcpu": 4,
					"memory_gb": 16
				},
				{
					"id": "906d663b6ecd",
					"instance_id": "i-0eba74077ac760573",
					"instance_type": "t3.small",
					"name": "kube-etcd2",
					"state": "stopped",
					"environment": "kube",
					"vcpu": 2,
					"memory_gb": 2
				},
				{
					"id": "8a0a2a0d8725",
					"instance_id": "i-0d67730effbee8b3b",
					"instance_type": "t3.xlarge",
					"name": "kube-k8node1",
					"state": "stopped",
					"environment": "kube",
					"vcpu": 4,
					"memory_gb": 16
				},
				{
					"id": "8c65980f5b5f",
					"instance_id": "i-009517923633b51f4",
					"instance_type": "t3.xlarge",
					"name": "kube-k8node4",
					"state": "stopped",
					"environment": "kube",
					"vcpu": 4,
					"memory_gb": 16
				}
			],
			"running_instances": 0,
			"stopped_instances": 10,
			"total_instances": 10,
			"total_vcpu": 30,
			"total_memory_gb": 94,
			"state": "stopped",
			"billsAccrued":"1.00",
			"billsSaved":"1.00"
		}
	],
	"totalBillsAccrued":"1.00",
	"totalBillsSaved":"1.00"
}
```
