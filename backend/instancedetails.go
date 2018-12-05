package backend

import "encoding/json"

type awsInstanceTypeDetails struct {
	InstanceType string  `json:"instance_type"`
	VCPU         int     `json:"vCPU"`
	MemoryGB     float32 `json:"memory"`
}

var instanceTypeDetailsCache []awsInstanceTypeDetails

func loadAwsInstanceDetailsJson() error {
	return json.Unmarshal([]byte(AWS_INSTANCE_TYPE_DETAILS_JSON), &instanceTypeDetailsCache)
}

func getInstanceTypeDetails(instanceType string) (typeDetails awsInstanceTypeDetails, found bool) {
	for _, details := range instanceTypeDetailsCache {
		if details.InstanceType == instanceType {
			typeDetails = details
			found = true
			break
		}
	}
	return
}

// generated with command:
// curl -s https://raw.githubusercontent.com/powdahound/ec2instances.info/master/www/instances.json | jq '.[] | {instance_type, vCPU, memory}' | jq -s .
// !! remove i3.metal since it uses a string for vCPU !!
const AWS_INSTANCE_TYPE_DETAILS_JSON = `
[
  {
    "instance_type": "m1.small",
    "vCPU": 1,
    "memory": 1.7
  },
  {
    "instance_type": "m1.medium",
    "vCPU": 1,
    "memory": 3.75
  },
  {
    "instance_type": "m1.large",
    "vCPU": 2,
    "memory": 7.5
  },
  {
    "instance_type": "m1.xlarge",
    "vCPU": 4,
    "memory": 15
  },
  {
    "instance_type": "m3.medium",
    "vCPU": 1,
    "memory": 3.75
  },
  {
    "instance_type": "m3.large",
    "vCPU": 2,
    "memory": 7.5
  },
  {
    "instance_type": "m3.xlarge",
    "vCPU": 4,
    "memory": 15
  },
  {
    "instance_type": "m3.2xlarge",
    "vCPU": 8,
    "memory": 30
  },
  {
    "instance_type": "c1.medium",
    "vCPU": 2,
    "memory": 1.7
  },
  {
    "instance_type": "c1.xlarge",
    "vCPU": 8,
    "memory": 7
  },
  {
    "instance_type": "cc2.8xlarge",
    "vCPU": 32,
    "memory": 60.5
  },
  {
    "instance_type": "c3.large",
    "vCPU": 2,
    "memory": 3.75
  },
  {
    "instance_type": "c3.xlarge",
    "vCPU": 4,
    "memory": 7.5
  },
  {
    "instance_type": "c3.2xlarge",
    "vCPU": 8,
    "memory": 15
  },
  {
    "instance_type": "c3.4xlarge",
    "vCPU": 16,
    "memory": 30
  },
  {
    "instance_type": "c3.8xlarge",
    "vCPU": 32,
    "memory": 60
  },
  {
    "instance_type": "g2.2xlarge",
    "vCPU": 8,
    "memory": 15
  },
  {
    "instance_type": "g2.8xlarge",
    "vCPU": 32,
    "memory": 60
  },
  {
    "instance_type": "m2.xlarge",
    "vCPU": 2,
    "memory": 17.1
  },
  {
    "instance_type": "m2.2xlarge",
    "vCPU": 4,
    "memory": 34.2
  },
  {
    "instance_type": "m2.4xlarge",
    "vCPU": 8,
    "memory": 68.4
  },
  {
    "instance_type": "cr1.8xlarge",
    "vCPU": 32,
    "memory": 244
  },
  {
    "instance_type": "r3.large",
    "vCPU": 2,
    "memory": 15.25
  },
  {
    "instance_type": "r3.xlarge",
    "vCPU": 4,
    "memory": 30.5
  },
  {
    "instance_type": "r3.2xlarge",
    "vCPU": 8,
    "memory": 61
  },
  {
    "instance_type": "r3.4xlarge",
    "vCPU": 16,
    "memory": 122
  },
  {
    "instance_type": "r3.8xlarge",
    "vCPU": 32,
    "memory": 244
  },
  {
    "instance_type": "i2.xlarge",
    "vCPU": 4,
    "memory": 30.5
  },
  {
    "instance_type": "i2.2xlarge",
    "vCPU": 8,
    "memory": 61
  },
  {
    "instance_type": "i2.4xlarge",
    "vCPU": 16,
    "memory": 122
  },
  {
    "instance_type": "i2.8xlarge",
    "vCPU": 32,
    "memory": 244
  },
  {
    "instance_type": "hs1.8xlarge",
    "vCPU": 16,
    "memory": 117
  },
  {
    "instance_type": "t1.micro",
    "vCPU": 1,
    "memory": 0.613
  },
  {
    "instance_type": "t3.nano",
    "vCPU": 2,
    "memory": 0.5
  },
  {
    "instance_type": "t3.micro",
    "vCPU": 2,
    "memory": 1
  },
  {
    "instance_type": "t3.small",
    "vCPU": 2,
    "memory": 2
  },
  {
    "instance_type": "t3.medium",
    "vCPU": 2,
    "memory": 4
  },
  {
    "instance_type": "t3.large",
    "vCPU": 2,
    "memory": 8
  },
  {
    "instance_type": "t3.xlarge",
    "vCPU": 4,
    "memory": 16
  },
  {
    "instance_type": "t3.2xlarge",
    "vCPU": 8,
    "memory": 32
  },
  {
    "instance_type": "t2.nano",
    "vCPU": 1,
    "memory": 0.5
  },
  {
    "instance_type": "t2.micro",
    "vCPU": 1,
    "memory": 1
  },
  {
    "instance_type": "t2.small",
    "vCPU": 1,
    "memory": 2
  },
  {
    "instance_type": "t2.medium",
    "vCPU": 2,
    "memory": 4
  },
  {
    "instance_type": "t2.large",
    "vCPU": 2,
    "memory": 8
  },
  {
    "instance_type": "t2.xlarge",
    "vCPU": 4,
    "memory": 16
  },
  {
    "instance_type": "t2.2xlarge",
    "vCPU": 8,
    "memory": 32
  },
  {
    "instance_type": "m5.large",
    "vCPU": 2,
    "memory": 8
  },
  {
    "instance_type": "m5.xlarge",
    "vCPU": 4,
    "memory": 16
  },
  {
    "instance_type": "m5.2xlarge",
    "vCPU": 8,
    "memory": 32
  },
  {
    "instance_type": "m5.4xlarge",
    "vCPU": 16,
    "memory": 64
  },
  {
    "instance_type": "m5.12xlarge",
    "vCPU": 48,
    "memory": 192
  },
  {
    "instance_type": "m5.24xlarge",
    "vCPU": 96,
    "memory": 384
  },
  {
    "instance_type": "m5d.large",
    "vCPU": 2,
    "memory": 8
  },
  {
    "instance_type": "m5d.xlarge",
    "vCPU": 4,
    "memory": 16
  },
  {
    "instance_type": "m5d.2xlarge",
    "vCPU": 8,
    "memory": 32
  },
  {
    "instance_type": "m5d.4xlarge",
    "vCPU": 16,
    "memory": 64
  },
  {
    "instance_type": "m5d.12xlarge",
    "vCPU": 48,
    "memory": 192
  },
  {
    "instance_type": "m5d.24xlarge",
    "vCPU": 96,
    "memory": 384
  },
  {
    "instance_type": "m4.large",
    "vCPU": 2,
    "memory": 8
  },
  {
    "instance_type": "m4.xlarge",
    "vCPU": 4,
    "memory": 16
  },
  {
    "instance_type": "m4.2xlarge",
    "vCPU": 8,
    "memory": 32
  },
  {
    "instance_type": "m4.4xlarge",
    "vCPU": 16,
    "memory": 64
  },
  {
    "instance_type": "m4.10xlarge",
    "vCPU": 40,
    "memory": 160
  },
  {
    "instance_type": "m4.16xlarge",
    "vCPU": 64,
    "memory": 256
  },
  {
    "instance_type": "c5.large",
    "vCPU": 2,
    "memory": 4
  },
  {
    "instance_type": "c5.xlarge",
    "vCPU": 4,
    "memory": 8
  },
  {
    "instance_type": "c5.2xlarge",
    "vCPU": 8,
    "memory": 16
  },
  {
    "instance_type": "c5.4xlarge",
    "vCPU": 16,
    "memory": 32
  },
  {
    "instance_type": "c5.9xlarge",
    "vCPU": 36,
    "memory": 72
  },
  {
    "instance_type": "c5.18xlarge",
    "vCPU": 72,
    "memory": 144
  },
  {
    "instance_type": "c5d.large",
    "vCPU": 2,
    "memory": 4
  },
  {
    "instance_type": "c5d.xlarge",
    "vCPU": 4,
    "memory": 8
  },
  {
    "instance_type": "c5d.2xlarge",
    "vCPU": 8,
    "memory": 16
  },
  {
    "instance_type": "c5d.4xlarge",
    "vCPU": 16,
    "memory": 32
  },
  {
    "instance_type": "c5d.9xlarge",
    "vCPU": 36,
    "memory": 72
  },
  {
    "instance_type": "c5d.18xlarge",
    "vCPU": 72,
    "memory": 144
  },
  {
    "instance_type": "c4.large",
    "vCPU": 2,
    "memory": 3.75
  },
  {
    "instance_type": "c4.xlarge",
    "vCPU": 4,
    "memory": 7.5
  },
  {
    "instance_type": "c4.2xlarge",
    "vCPU": 8,
    "memory": 15
  },
  {
    "instance_type": "c4.4xlarge",
    "vCPU": 16,
    "memory": 30
  },
  {
    "instance_type": "c4.8xlarge",
    "vCPU": 36,
    "memory": 60
  },
  {
    "instance_type": "r5.large",
    "vCPU": 2,
    "memory": 16
  },
  {
    "instance_type": "r5.xlarge",
    "vCPU": 4,
    "memory": 32
  },
  {
    "instance_type": "r5.2xlarge",
    "vCPU": 8,
    "memory": 64
  },
  {
    "instance_type": "r5.4xlarge",
    "vCPU": 16,
    "memory": 128
  },
  {
    "instance_type": "r5.12xlarge",
    "vCPU": 48,
    "memory": 384
  },
  {
    "instance_type": "r5.24xlarge",
    "vCPU": 96,
    "memory": 768
  },
  {
    "instance_type": "r5d.large",
    "vCPU": 2,
    "memory": 16
  },
  {
    "instance_type": "r5d.xlarge",
    "vCPU": 4,
    "memory": 32
  },
  {
    "instance_type": "r5d.2xlarge",
    "vCPU": 8,
    "memory": 64
  },
  {
    "instance_type": "r5d.4xlarge",
    "vCPU": 16,
    "memory": 128
  },
  {
    "instance_type": "r5d.12xlarge",
    "vCPU": 48,
    "memory": 384
  },
  {
    "instance_type": "r5d.24xlarge",
    "vCPU": 96,
    "memory": 768
  },
  {
    "instance_type": "r4.large",
    "vCPU": 2,
    "memory": 15.25
  },
  {
    "instance_type": "r4.xlarge",
    "vCPU": 4,
    "memory": 30.5
  },
  {
    "instance_type": "r4.2xlarge",
    "vCPU": 8,
    "memory": 61
  },
  {
    "instance_type": "r4.4xlarge",
    "vCPU": 16,
    "memory": 122
  },
  {
    "instance_type": "r4.8xlarge",
    "vCPU": 32,
    "memory": 244
  },
  {
    "instance_type": "r4.16xlarge",
    "vCPU": 64,
    "memory": 488
  },
  {
    "instance_type": "x1e.xlarge",
    "vCPU": 4,
    "memory": 122
  },
  {
    "instance_type": "x1e.2xlarge",
    "vCPU": 8,
    "memory": 244
  },
  {
    "instance_type": "x1e.4xlarge",
    "vCPU": 16,
    "memory": 488
  },
  {
    "instance_type": "x1e.8xlarge",
    "vCPU": 32,
    "memory": 976
  },
  {
    "instance_type": "x1e.16xlarge",
    "vCPU": 64,
    "memory": 1952
  },
  {
    "instance_type": "x1e.32xlarge",
    "vCPU": 128,
    "memory": 3904
  },
  {
    "instance_type": "x1.16xlarge",
    "vCPU": 64,
    "memory": 976
  },
  {
    "instance_type": "x1.32xlarge",
    "vCPU": 128,
    "memory": 1952
  },
  {
    "instance_type": "u-6tb1.metal",
    "vCPU": 448,
    "memory": 6144
  },
  {
    "instance_type": "u-9tb1.metal",
    "vCPU": 448,
    "memory": 9216
  },
  {
    "instance_type": "u-12tb1.metal",
    "vCPU": 448,
    "memory": 12288
  },
  {
    "instance_type": "z1d.large",
    "vCPU": 2,
    "memory": 16
  },
  {
    "instance_type": "z1d.xlarge",
    "vCPU": 4,
    "memory": 32
  },
  {
    "instance_type": "z1d.2xlarge",
    "vCPU": 8,
    "memory": 64
  },
  {
    "instance_type": "z1d.3xlarge",
    "vCPU": 12,
    "memory": 96
  },
  {
    "instance_type": "z1d.6xlarge",
    "vCPU": 24,
    "memory": 192
  },
  {
    "instance_type": "z1d.12xlarge",
    "vCPU": 48,
    "memory": 384
  },
  {
    "instance_type": "p3.2xlarge",
    "vCPU": 8,
    "memory": 61
  },
  {
    "instance_type": "p3.8xlarge",
    "vCPU": 32,
    "memory": 244
  },
  {
    "instance_type": "p3.16xlarge",
    "vCPU": 64,
    "memory": 488
  },
  {
    "instance_type": "p2.xlarge",
    "vCPU": 4,
    "memory": 61
  },
  {
    "instance_type": "p2.8xlarge",
    "vCPU": 32,
    "memory": 488
  },
  {
    "instance_type": "p2.16xlarge",
    "vCPU": 64,
    "memory": 732
  },
  {
    "instance_type": "g3s.xlarge",
    "vCPU": 4,
    "memory": 30.5
  },
  {
    "instance_type": "g3.4xlarge",
    "vCPU": 16,
    "memory": 122
  },
  {
    "instance_type": "g3.8xlarge",
    "vCPU": 32,
    "memory": 244
  },
  {
    "instance_type": "g3.16xlarge",
    "vCPU": 64,
    "memory": 488
  },
  {
    "instance_type": "f1.2xlarge",
    "vCPU": 8,
    "memory": 122
  },
  {
    "instance_type": "f1.4xlarge",
    "vCPU": 16,
    "memory": 244
  },
  {
    "instance_type": "f1.16xlarge",
    "vCPU": 64,
    "memory": 976
  },
  {
    "instance_type": "h1.2xlarge",
    "vCPU": 8,
    "memory": 32
  },
  {
    "instance_type": "h1.4xlarge",
    "vCPU": 16,
    "memory": 64
  },
  {
    "instance_type": "h1.8xlarge",
    "vCPU": 32,
    "memory": 128
  },
  {
    "instance_type": "h1.16xlarge",
    "vCPU": 64,
    "memory": 256
  },
  {
    "instance_type": "i3.large",
    "vCPU": 2,
    "memory": 15.25
  },
  {
    "instance_type": "i3.xlarge",
    "vCPU": 4,
    "memory": 30.5
  },
  {
    "instance_type": "i3.2xlarge",
    "vCPU": 8,
    "memory": 61
  },
  {
    "instance_type": "i3.4xlarge",
    "vCPU": 16,
    "memory": 122
  },
  {
    "instance_type": "i3.8xlarge",
    "vCPU": 32,
    "memory": 244
  },
  {
    "instance_type": "i3.16xlarge",
    "vCPU": 64,
    "memory": 488
  },
  {
    "instance_type": "d2.xlarge",
    "vCPU": 4,
    "memory": 30.5
  },
  {
    "instance_type": "d2.2xlarge",
    "vCPU": 8,
    "memory": 61
  },
  {
    "instance_type": "d2.4xlarge",
    "vCPU": 16,
    "memory": 122
  },
  {
    "instance_type": "d2.8xlarge",
    "vCPU": 36,
    "memory": 244
  }
]
`
