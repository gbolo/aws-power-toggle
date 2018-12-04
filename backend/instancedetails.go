package backend

import "encoding/json"

type awsInstanceTypeDetails struct {
	InstanceType          string            `json:"instance_type"`
	VCPU                  int               `json:"vCPU"`
	MemoryGB              float32           `json:"memory"`
	PricingHourlyByRegion map[string]string `json:"pricing"`
}

var instanceTypeDetailsCache []awsInstanceTypeDetails

func loadAwsInstanceDetailsJSON() error {
	return json.Unmarshal([]byte(awsInstanceTypeDetailsJSON), &instanceTypeDetailsCache)
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
// curl -s https://raw.githubusercontent.com/powdahound/ec2instances.info/master/www/instances.json | jq '.[] | {instance_type, vCPU, memory, pricing:.pricing | map_values(.linux|.ondemand) }' | jq -s .
// the JSON has been simplified from the original version. pricing is for linux.ondemand ONLY
const awsInstanceTypeDetailsJSON = `
[
  {
    "instance_type": "c5d.xlarge",
    "vCPU": 4,
    "memory": 8,
    "pricing": {
      "ap-northeast-1": "0.244",
      "ap-northeast-2": "0.22",
      "ap-southeast-1": "0.224",
      "ap-southeast-2": "0.252",
      "ca-central-1": "0.212",
      "eu-central-1": "0.222",
      "eu-west-1": "0.218",
      "eu-west-2": "0.23",
      "us-east-1": "0.192",
      "us-east-2": "0.192",
      "us-gov-west-1": "0.232",
      "us-west-1": "0.24",
      "us-west-2": "0.192"
    }
  },
  {
    "instance_type": "m5a.2xlarge",
    "vCPU": 8,
    "memory": 32,
    "pricing": {
      "ap-southeast-1": "0.432",
      "eu-west-1": "0.384",
      "us-east-1": "0.344",
      "us-east-2": "0.344",
      "us-west-2": "0.344"
    }
  },
  {
    "instance_type": "c5.9xlarge",
    "vCPU": 36,
    "memory": 72,
    "pricing": {
      "ap-northeast-1": "1.926",
      "ap-northeast-2": "1.728",
      "ap-south-1": "1.53",
      "ap-southeast-1": "1.764",
      "ap-southeast-2": "1.998",
      "ca-central-1": "1.674",
      "eu-central-1": "1.746",
      "eu-west-1": "1.728",
      "eu-west-2": "1.818",
      "eu-west-3": "1.818",
      "sa-east-1": "2.358",
      "us-east-1": "1.53",
      "us-east-2": "1.53",
      "us-gov-east-1": "1.836",
      "us-gov-west-1": "1.836",
      "us-west-1": "1.908",
      "us-west-2": "1.53"
    }
  },
  {
    "instance_type": "m5.24xlarge",
    "vCPU": 96,
    "memory": 384,
    "pricing": {
      "ap-northeast-1": "5.952",
      "ap-northeast-2": "5.664",
      "ap-south-1": "4.848",
      "ap-southeast-1": "5.76",
      "ap-southeast-2": "5.76",
      "ca-central-1": "5.136",
      "eu-central-1": "5.52",
      "eu-west-1": "5.136",
      "eu-west-2": "5.328",
      "eu-west-3": "5.376",
      "sa-east-1": "7.344",
      "us-east-1": "4.608",
      "us-east-2": "4.608",
      "us-gov-east-1": "5.808",
      "us-gov-west-1": "5.808",
      "us-west-1": "5.376",
      "us-west-2": "4.608"
    }
  },
  {
    "instance_type": "m5d.12xlarge",
    "vCPU": 48,
    "memory": 192,
    "pricing": {
      "ap-northeast-1": "3.504",
      "ap-northeast-2": "3.336",
      "ap-southeast-1": "3.384",
      "ap-southeast-2": "3.408",
      "ca-central-1": "3.024",
      "eu-central-1": "3.264",
      "eu-west-1": "3.024",
      "eu-west-2": "3.144",
      "us-east-1": "2.712",
      "us-east-2": "2.712",
      "us-gov-west-1": "3.432",
      "us-west-1": "3.192",
      "us-west-2": "2.712"
    }
  },
  {
    "instance_type": "c5.large",
    "vCPU": 2,
    "memory": 4,
    "pricing": {
      "ap-northeast-1": "0.107",
      "ap-northeast-2": "0.096",
      "ap-south-1": "0.085",
      "ap-southeast-1": "0.098",
      "ap-southeast-2": "0.111",
      "ca-central-1": "0.093",
      "eu-central-1": "0.097",
      "eu-west-1": "0.096",
      "eu-west-2": "0.101",
      "eu-west-3": "0.101",
      "sa-east-1": "0.131",
      "us-east-1": "0.085",
      "us-east-2": "0.085",
      "us-gov-east-1": "0.102",
      "us-gov-west-1": "0.102",
      "us-west-1": "0.106",
      "us-west-2": "0.085"
    }
  },
  {
    "instance_type": "c5n.large",
    "vCPU": 2,
    "memory": 5.25,
    "pricing": {
      "eu-west-1": "0.122",
      "us-east-1": "0.108",
      "us-east-2": "0.108",
      "us-gov-west-1": "0.13",
      "us-west-2": "0.108"
    }
  },
  {
    "instance_type": "i2.xlarge",
    "vCPU": 4,
    "memory": 30.5,
    "pricing": {
      "ap-northeast-1": "1.001",
      "ap-northeast-2": "1.001",
      "ap-south-1": "0.967",
      "ap-southeast-1": "1.018",
      "ap-southeast-2": "1.018",
      "eu-central-1": "1.013",
      "eu-west-1": "0.938",
      "us-east-1": "0.853",
      "us-east-2": "0.853",
      "us-gov-west-1": "1.023",
      "us-west-1": "0.938",
      "us-west-2": "0.853"
    }
  },
  {
    "instance_type": "t2.micro",
    "vCPU": 1,
    "memory": 1,
    "pricing": {
      "ap-northeast-1": "0.0152",
      "ap-northeast-2": "0.0144",
      "ap-northeast-3": "0.0152",
      "ap-south-1": "0.0124",
      "ap-southeast-1": "0.0146",
      "ap-southeast-2": "0.0146",
      "ca-central-1": "0.0128",
      "eu-central-1": "0.0134",
      "eu-west-1": "0.0126",
      "eu-west-2": "0.0132",
      "eu-west-3": "0.0132",
      "sa-east-1": "0.0186",
      "us-east-1": "0.0116",
      "us-east-2": "0.0116",
      "us-gov-west-1": "0.0136",
      "us-west-1": "0.0138",
      "us-west-2": "0.0116"
    }
  },
  {
    "instance_type": "d2.8xlarge",
    "vCPU": 36,
    "memory": 244,
    "pricing": {
      "ap-northeast-1": "6.752",
      "ap-northeast-2": "6.752",
      "ap-northeast-3": "6.752",
      "ap-south-1": "6.612",
      "ap-southeast-1": "6.96",
      "ap-southeast-2": "6.96",
      "ca-central-1": "6.072",
      "eu-central-1": "6.352",
      "eu-west-1": "5.88",
      "eu-west-2": "6.174",
      "eu-west-3": "6.176",
      "us-east-1": "5.52",
      "us-east-2": "5.52",
      "us-gov-west-1": "6.624",
      "us-west-1": "6.25",
      "us-west-2": "5.52"
    }
  },
  {
    "instance_type": "z1d.3xlarge",
    "vCPU": 12,
    "memory": 96,
    "pricing": {
      "ap-northeast-1": "1.362",
      "ap-southeast-1": "1.356",
      "eu-west-1": "1.248",
      "us-east-1": "1.116",
      "us-west-1": "1.266",
      "us-west-2": "1.116"
    }
  },
  {
    "instance_type": "m5.2xlarge",
    "vCPU": 8,
    "memory": 32,
    "pricing": {
      "ap-northeast-1": "0.496",
      "ap-northeast-2": "0.472",
      "ap-south-1": "0.404",
      "ap-southeast-1": "0.48",
      "ap-southeast-2": "0.48",
      "ca-central-1": "0.428",
      "eu-central-1": "0.46",
      "eu-west-1": "0.428",
      "eu-west-2": "0.444",
      "eu-west-3": "0.448",
      "sa-east-1": "0.612",
      "us-east-1": "0.384",
      "us-east-2": "0.384",
      "us-gov-east-1": "0.484",
      "us-gov-west-1": "0.484",
      "us-west-1": "0.448",
      "us-west-2": "0.384"
    }
  },
  {
    "instance_type": "c5.18xlarge",
    "vCPU": 72,
    "memory": 144,
    "pricing": {
      "ap-northeast-1": "3.852",
      "ap-northeast-2": "3.456",
      "ap-south-1": "3.06",
      "ap-southeast-1": "3.528",
      "ap-southeast-2": "3.996",
      "ca-central-1": "3.348",
      "eu-central-1": "3.492",
      "eu-west-1": "3.456",
      "eu-west-2": "3.636",
      "eu-west-3": "3.636",
      "sa-east-1": "4.716",
      "us-east-1": "3.06",
      "us-east-2": "3.06",
      "us-gov-east-1": "3.672",
      "us-gov-west-1": "3.672",
      "us-west-1": "3.816",
      "us-west-2": "3.06"
    }
  },
  {
    "instance_type": "x1e.16xlarge",
    "vCPU": 64,
    "memory": 1952,
    "pricing": {
      "ap-northeast-1": "19.344",
      "ap-southeast-2": "19.344",
      "eu-central-1": "18.672",
      "eu-west-1": "16",
      "us-east-1": "13.344",
      "us-gov-west-1": "16",
      "us-west-2": "13.344"
    }
  },
  {
    "instance_type": "i2.8xlarge",
    "vCPU": 32,
    "memory": 244,
    "pricing": {
      "ap-northeast-1": "8.004",
      "ap-northeast-2": "8.004",
      "ap-south-1": "7.733",
      "ap-southeast-1": "8.14",
      "ap-southeast-2": "8.14",
      "eu-central-1": "8.102",
      "eu-west-1": "7.502",
      "us-east-1": "6.82",
      "us-east-2": "6.82",
      "us-gov-west-1": "8.184",
      "us-west-1": "7.502",
      "us-west-2": "6.82"
    }
  },
  {
    "instance_type": "i2.2xlarge",
    "vCPU": 8,
    "memory": 61,
    "pricing": {
      "ap-northeast-1": "2.001",
      "ap-northeast-2": "2.001",
      "ap-south-1": "1.933",
      "ap-southeast-1": "2.035",
      "ap-southeast-2": "2.035",
      "eu-central-1": "2.026",
      "eu-west-1": "1.876",
      "us-east-1": "1.705",
      "us-east-2": "1.705",
      "us-gov-west-1": "2.046",
      "us-west-1": "1.876",
      "us-west-2": "1.705"
    }
  },
  {
    "instance_type": "m5a.xlarge",
    "vCPU": 4,
    "memory": 16,
    "pricing": {
      "ap-southeast-1": "0.216",
      "eu-west-1": "0.192",
      "us-east-1": "0.172",
      "us-east-2": "0.172",
      "us-west-2": "0.172"
    }
  },
  {
    "instance_type": "p3.2xlarge",
    "vCPU": 8,
    "memory": 61,
    "pricing": {
      "ap-northeast-1": "5.243",
      "ap-northeast-2": "4.981",
      "ap-southeast-1": "4.981",
      "ap-southeast-2": "4.981",
      "ca-central-1": "3.366",
      "eu-central-1": "3.823",
      "eu-west-1": "3.305",
      "eu-west-2": "3.589",
      "us-east-1": "3.06",
      "us-east-2": "3.06",
      "us-gov-west-1": "3.672",
      "us-west-2": "3.06"
    }
  },
  {
    "instance_type": "t2.2xlarge",
    "vCPU": 8,
    "memory": 32,
    "pricing": {
      "ap-northeast-1": "0.4864",
      "ap-northeast-2": "0.4608",
      "ap-northeast-3": "0.4864",
      "ap-south-1": "0.3968",
      "ap-southeast-1": "0.4672",
      "ap-southeast-2": "0.4672",
      "ca-central-1": "0.4096",
      "eu-central-1": "0.4288",
      "eu-west-1": "0.4032",
      "eu-west-2": "0.4224",
      "eu-west-3": "0.4224",
      "sa-east-1": "0.5952",
      "us-east-1": "0.3712",
      "us-east-2": "0.3712",
      "us-gov-west-1": "0.4352",
      "us-west-1": "0.4416",
      "us-west-2": "0.3712"
    }
  },
  {
    "instance_type": "h1.8xlarge",
    "vCPU": 32,
    "memory": 128,
    "pricing": {
      "eu-west-1": "2.076",
      "us-east-1": "1.872",
      "us-east-2": "1.872",
      "us-west-2": "1.872"
    }
  },
  {
    "instance_type": "r5d.24xlarge",
    "vCPU": 96,
    "memory": 768,
    "pricing": {
      "ap-northeast-1": "8.352",
      "ap-northeast-2": "8.304",
      "ap-southeast-1": "8.352",
      "ap-southeast-2": "8.352",
      "ca-central-1": "7.584",
      "eu-central-1": "8.304",
      "eu-west-1": "7.68",
      "eu-west-2": "8.112",
      "us-east-1": "6.912",
      "us-east-2": "6.912",
      "us-west-1": "7.776",
      "us-west-2": "6.912"
    }
  },
  {
    "instance_type": "r4.8xlarge",
    "vCPU": 32,
    "memory": 244,
    "pricing": {
      "ap-northeast-1": "2.56",
      "ap-northeast-2": "2.56",
      "ap-northeast-3": "2.56",
      "ap-south-1": "2.192",
      "ap-southeast-1": "2.56",
      "ap-southeast-2": "2.5536",
      "ca-central-1": "2.336",
      "eu-central-1": "2.5608",
      "eu-west-1": "2.3712",
      "eu-west-2": "2.496",
      "eu-west-3": "2.496",
      "sa-east-1": "4.48",
      "us-east-1": "2.128",
      "us-east-2": "2.128",
      "us-gov-west-1": "2.5536",
      "us-west-1": "2.3712",
      "us-west-2": "2.128"
    }
  },
  {
    "instance_type": "t2.large",
    "vCPU": 2,
    "memory": 8,
    "pricing": {
      "ap-northeast-1": "0.1216",
      "ap-northeast-2": "0.1152",
      "ap-northeast-3": "0.1216",
      "ap-south-1": "0.0992",
      "ap-southeast-1": "0.1168",
      "ap-southeast-2": "0.1168",
      "ca-central-1": "0.1024",
      "eu-central-1": "0.1072",
      "eu-west-1": "0.1008",
      "eu-west-2": "0.1056",
      "eu-west-3": "0.1056",
      "sa-east-1": "0.1488",
      "us-east-1": "0.0928",
      "us-east-2": "0.0928",
      "us-gov-west-1": "0.1088",
      "us-west-1": "0.1104",
      "us-west-2": "0.0928"
    }
  },
  {
    "instance_type": "x1.16xlarge",
    "vCPU": 64,
    "memory": 976,
    "pricing": {
      "ap-northeast-1": "9.671",
      "ap-northeast-2": "9.671",
      "ap-south-1": "9.187",
      "ap-southeast-1": "9.671",
      "ap-southeast-2": "9.671",
      "ca-central-1": "7.336",
      "eu-central-1": "9.337",
      "eu-west-1": "8.003",
      "eu-west-2": "8.403",
      "eu-west-3": "8.403",
      "sa-east-1": "13.005",
      "us-east-1": "6.669",
      "us-east-2": "6.669",
      "us-gov-west-1": "8.003",
      "us-west-2": "6.669"
    }
  },
  {
    "instance_type": "c5d.18xlarge",
    "vCPU": 72,
    "memory": 144,
    "pricing": {
      "ap-northeast-1": "4.392",
      "ap-northeast-2": "3.96",
      "ap-southeast-1": "4.032",
      "ap-southeast-2": "4.536",
      "ca-central-1": "3.816",
      "eu-central-1": "3.996",
      "eu-west-1": "3.924",
      "eu-west-2": "4.14",
      "us-east-1": "3.456",
      "us-east-2": "3.456",
      "us-gov-west-1": "4.176",
      "us-west-1": "4.32",
      "us-west-2": "3.456"
    }
  },
  {
    "instance_type": "r5a.large",
    "vCPU": 2,
    "memory": 16,
    "pricing": {
      "ap-southeast-1": "0.136",
      "eu-west-1": "0.127",
      "us-east-1": "0.113",
      "us-east-2": "0.113",
      "us-west-2": "0.113"
    }
  },
  {
    "instance_type": "c3.large",
    "vCPU": 2,
    "memory": 3.75,
    "pricing": {
      "ap-northeast-1": "0.128",
      "ap-northeast-2": "0.115",
      "ap-northeast-3": null,
      "ap-southeast-1": "0.132",
      "ap-southeast-2": "0.132",
      "eu-central-1": "0.129",
      "eu-west-1": "0.12",
      "sa-east-1": "0.163",
      "us-east-1": "0.105",
      "us-gov-west-1": "0.126",
      "us-west-1": "0.12",
      "us-west-2": "0.105"
    }
  },
  {
    "instance_type": "r5a.24xlarge",
    "vCPU": 96,
    "memory": 768,
    "pricing": {
      "ap-southeast-1": "6.528",
      "eu-west-1": "6.096",
      "us-east-1": "5.424",
      "us-east-2": "5.424",
      "us-west-2": "5.424"
    }
  },
  {
    "instance_type": "g3.16xlarge",
    "vCPU": 64,
    "memory": 488,
    "pricing": {
      "ap-northeast-1": "6.32",
      "ap-southeast-1": "6.68",
      "ap-southeast-2": "7.016",
      "ca-central-1": "5.664",
      "eu-central-1": "5.7",
      "eu-west-1": "4.84",
      "us-east-1": "4.56",
      "us-east-2": "4.56",
      "us-gov-west-1": "5.28",
      "us-west-1": "6.136",
      "us-west-2": "4.56"
    }
  },
  {
    "instance_type": "c4.xlarge",
    "vCPU": 4,
    "memory": 7.5,
    "pricing": {
      "ap-northeast-1": "0.252",
      "ap-northeast-2": "0.227",
      "ap-northeast-3": "0.252",
      "ap-south-1": "0.2",
      "ap-southeast-1": "0.231",
      "ap-southeast-2": "0.261",
      "ca-central-1": "0.218",
      "eu-central-1": "0.227",
      "eu-west-1": "0.226",
      "eu-west-2": "0.237",
      "sa-east-1": "0.309",
      "us-east-1": "0.199",
      "us-east-2": "0.199",
      "us-gov-west-1": "0.239",
      "us-west-1": "0.249",
      "us-west-2": "0.199"
    }
  },
  {
    "instance_type": "x1e.4xlarge",
    "vCPU": 16,
    "memory": 488,
    "pricing": {
      "ap-northeast-1": "4.836",
      "ap-southeast-2": "4.836",
      "eu-central-1": "4.668",
      "eu-west-1": "4",
      "us-east-1": "3.336",
      "us-gov-west-1": "4",
      "us-west-2": "3.336"
    }
  },
  {
    "instance_type": "c5n.18xlarge",
    "vCPU": 72,
    "memory": 192,
    "pricing": {
      "eu-west-1": "4.392",
      "us-east-1": "3.888",
      "us-east-2": "3.888",
      "us-gov-west-1": "4.68",
      "us-west-2": "3.888"
    }
  },
  {
    "instance_type": "m4.large",
    "vCPU": 2,
    "memory": 8,
    "pricing": {
      "ap-northeast-1": "0.129",
      "ap-northeast-2": "0.123",
      "ap-northeast-3": "0.129",
      "ap-south-1": "0.105",
      "ap-southeast-1": "0.125",
      "ap-southeast-2": "0.125",
      "ca-central-1": "0.111",
      "eu-central-1": "0.12",
      "eu-west-1": "0.111",
      "eu-west-2": "0.116",
      "sa-east-1": "0.159",
      "us-east-1": "0.1",
      "us-east-2": "0.1",
      "us-gov-west-1": "0.126",
      "us-west-1": "0.117",
      "us-west-2": "0.1"
    }
  },
  {
    "instance_type": "h1.4xlarge",
    "vCPU": 16,
    "memory": 64,
    "pricing": {
      "eu-west-1": "1.038",
      "us-east-1": "0.936",
      "us-east-2": "0.936",
      "us-west-2": "0.936"
    }
  },
  {
    "instance_type": "x1e.xlarge",
    "vCPU": 4,
    "memory": 122,
    "pricing": {
      "ap-northeast-1": "1.209",
      "ap-southeast-2": "1.209",
      "eu-central-1": "1.167",
      "eu-west-1": "1",
      "us-east-1": "0.834",
      "us-gov-west-1": "1",
      "us-west-2": "0.834"
    }
  },
  {
    "instance_type": "m5.large",
    "vCPU": 2,
    "memory": 8,
    "pricing": {
      "ap-northeast-1": "0.124",
      "ap-northeast-2": "0.118",
      "ap-south-1": "0.101",
      "ap-southeast-1": "0.12",
      "ap-southeast-2": "0.12",
      "ca-central-1": "0.107",
      "eu-central-1": "0.115",
      "eu-west-1": "0.107",
      "eu-west-2": "0.111",
      "eu-west-3": "0.112",
      "sa-east-1": "0.153",
      "us-east-1": "0.096",
      "us-east-2": "0.096",
      "us-gov-east-1": "0.121",
      "us-gov-west-1": "0.121",
      "us-west-1": "0.112",
      "us-west-2": "0.096"
    }
  },
  {
    "instance_type": "c5.4xlarge",
    "vCPU": 16,
    "memory": 32,
    "pricing": {
      "ap-northeast-1": "0.856",
      "ap-northeast-2": "0.768",
      "ap-south-1": "0.68",
      "ap-southeast-1": "0.784",
      "ap-southeast-2": "0.888",
      "ca-central-1": "0.744",
      "eu-central-1": "0.776",
      "eu-west-1": "0.768",
      "eu-west-2": "0.808",
      "eu-west-3": "0.808",
      "sa-east-1": "1.048",
      "us-east-1": "0.68",
      "us-east-2": "0.68",
      "us-gov-east-1": "0.816",
      "us-gov-west-1": "0.816",
      "us-west-1": "0.848",
      "us-west-2": "0.68"
    }
  },
  {
    "instance_type": "m5d.24xlarge",
    "vCPU": 96,
    "memory": 384,
    "pricing": {
      "ap-northeast-1": "7.008",
      "ap-northeast-2": "6.672",
      "ap-southeast-1": "6.768",
      "ap-southeast-2": "6.816",
      "ca-central-1": "6.048",
      "eu-central-1": "6.528",
      "eu-west-1": "6.048",
      "eu-west-2": "6.288",
      "us-east-1": "5.424",
      "us-east-2": "5.424",
      "us-gov-west-1": "6.864",
      "us-west-1": "6.384",
      "us-west-2": "5.424"
    }
  },
  {
    "instance_type": "r3.large",
    "vCPU": 2,
    "memory": 15.25,
    "pricing": {
      "ap-northeast-1": "0.2",
      "ap-northeast-2": "0.2",
      "ap-northeast-3": null,
      "ap-south-1": "0.19",
      "ap-southeast-1": "0.2",
      "ap-southeast-2": "0.2",
      "eu-central-1": "0.2",
      "eu-west-1": "0.185",
      "sa-east-1": "0.35",
      "us-east-1": "0.166",
      "us-east-2": "0.166",
      "us-gov-west-1": "0.2",
      "us-west-1": "0.185",
      "us-west-2": "0.166"
    }
  },
  {
    "instance_type": "c4.large",
    "vCPU": 2,
    "memory": 3.75,
    "pricing": {
      "ap-northeast-1": "0.126",
      "ap-northeast-2": "0.114",
      "ap-northeast-3": "0.126",
      "ap-south-1": "0.1",
      "ap-southeast-1": "0.115",
      "ap-southeast-2": "0.13",
      "ca-central-1": "0.11",
      "eu-central-1": "0.114",
      "eu-west-1": "0.113",
      "eu-west-2": "0.119",
      "sa-east-1": "0.155",
      "us-east-1": "0.1",
      "us-east-2": "0.1",
      "us-gov-west-1": "0.12",
      "us-west-1": "0.124",
      "us-west-2": "0.1"
    }
  },
  {
    "instance_type": "r5d.xlarge",
    "vCPU": 4,
    "memory": 32,
    "pricing": {
      "ap-northeast-1": "0.348",
      "ap-northeast-2": "0.346",
      "ap-southeast-1": "0.348",
      "ap-southeast-2": "0.348",
      "ca-central-1": "0.316",
      "eu-central-1": "0.346",
      "eu-west-1": "0.32",
      "eu-west-2": "0.338",
      "us-east-1": "0.288",
      "us-east-2": "0.288",
      "us-west-1": "0.324",
      "us-west-2": "0.288"
    }
  },
  {
    "instance_type": "m5d.large",
    "vCPU": 2,
    "memory": 8,
    "pricing": {
      "ap-northeast-1": "0.146",
      "ap-northeast-2": "0.139",
      "ap-southeast-1": "0.141",
      "ap-southeast-2": "0.142",
      "ca-central-1": "0.126",
      "eu-central-1": "0.136",
      "eu-west-1": "0.126",
      "eu-west-2": "0.131",
      "us-east-1": "0.113",
      "us-east-2": "0.113",
      "us-gov-west-1": "0.143",
      "us-west-1": "0.133",
      "us-west-2": "0.113"
    }
  },
  {
    "instance_type": "r5.2xlarge",
    "vCPU": 8,
    "memory": 64,
    "pricing": {
      "ap-northeast-1": "0.608",
      "ap-northeast-2": "0.608",
      "ap-southeast-1": "0.608",
      "ap-southeast-2": "0.604",
      "ca-central-1": "0.552",
      "eu-central-1": "0.608",
      "eu-west-1": "0.564",
      "eu-west-2": "0.592",
      "eu-west-3": "0.592",
      "us-east-1": "0.504",
      "us-east-2": "0.504",
      "us-gov-east-1": "0.604",
      "us-gov-west-1": "0.604",
      "us-west-1": "0.56",
      "us-west-2": "0.504"
    }
  },
  {
    "instance_type": "m3.2xlarge",
    "vCPU": 8,
    "memory": 30,
    "pricing": {
      "ap-northeast-1": "0.77",
      "ap-northeast-2": "0.732",
      "ap-northeast-3": null,
      "ap-southeast-1": "0.784",
      "ap-southeast-2": "0.745",
      "eu-central-1": "0.632",
      "eu-west-1": "0.585",
      "sa-east-1": "0.761",
      "us-east-1": "0.532",
      "us-gov-west-1": "0.672",
      "us-west-1": "0.616",
      "us-west-2": "0.532"
    }
  },
  {
    "instance_type": "m5d.2xlarge",
    "vCPU": 8,
    "memory": 32,
    "pricing": {
      "ap-northeast-1": "0.584",
      "ap-northeast-2": "0.556",
      "ap-southeast-1": "0.564",
      "ap-southeast-2": "0.568",
      "ca-central-1": "0.504",
      "eu-central-1": "0.544",
      "eu-west-1": "0.504",
      "eu-west-2": "0.524",
      "us-east-1": "0.452",
      "us-east-2": "0.452",
      "us-gov-west-1": "0.572",
      "us-west-1": "0.532",
      "us-west-2": "0.452"
    }
  },
  {
    "instance_type": "m4.10xlarge",
    "vCPU": 40,
    "memory": 160,
    "pricing": {
      "ap-northeast-1": "2.58",
      "ap-northeast-2": "2.46",
      "ap-northeast-3": "2.58",
      "ap-south-1": "2.1",
      "ap-southeast-1": "2.5",
      "ap-southeast-2": "2.5",
      "ca-central-1": "2.22",
      "eu-central-1": "2.4",
      "eu-west-1": "2.22",
      "eu-west-2": "2.32",
      "sa-east-1": "3.18",
      "us-east-1": "2",
      "us-east-2": "2",
      "us-gov-west-1": "2.52",
      "us-west-1": "2.34",
      "us-west-2": "2"
    }
  },
  {
    "instance_type": "c5d.large",
    "vCPU": 2,
    "memory": 4,
    "pricing": {
      "ap-northeast-1": "0.122",
      "ap-northeast-2": "0.11",
      "ap-southeast-1": "0.112",
      "ap-southeast-2": "0.126",
      "ca-central-1": "0.106",
      "eu-central-1": "0.111",
      "eu-west-1": "0.109",
      "eu-west-2": "0.115",
      "us-east-1": "0.096",
      "us-east-2": "0.096",
      "us-gov-west-1": "0.116",
      "us-west-1": "0.12",
      "us-west-2": "0.096"
    }
  },
  {
    "instance_type": "x1.32xlarge",
    "vCPU": 128,
    "memory": 1952,
    "pricing": {
      "ap-northeast-1": "19.341",
      "ap-northeast-2": "19.341",
      "ap-south-1": "18.374",
      "ap-southeast-1": "19.341",
      "ap-southeast-2": "19.341",
      "ca-central-1": "14.672",
      "eu-central-1": "18.674",
      "eu-west-1": "16.006",
      "eu-west-2": "16.806",
      "eu-west-3": "16.806",
      "sa-east-1": "26.01",
      "us-east-1": "13.338",
      "us-east-2": "13.338",
      "us-gov-west-1": "16.006",
      "us-west-2": "13.338"
    }
  },
  {
    "instance_type": "c5n.2xlarge",
    "vCPU": 8,
    "memory": 21,
    "pricing": {
      "eu-west-1": "0.488",
      "us-east-1": "0.432",
      "us-east-2": "0.432",
      "us-gov-west-1": "0.52",
      "us-west-2": "0.432"
    }
  },
  {
    "instance_type": "m5d.xlarge",
    "vCPU": 4,
    "memory": 16,
    "pricing": {
      "ap-northeast-1": "0.292",
      "ap-northeast-2": "0.278",
      "ap-southeast-1": "0.282",
      "ap-southeast-2": "0.284",
      "ca-central-1": "0.252",
      "eu-central-1": "0.272",
      "eu-west-1": "0.252",
      "eu-west-2": "0.262",
      "us-east-1": "0.226",
      "us-east-2": "0.226",
      "us-gov-west-1": "0.286",
      "us-west-1": "0.266",
      "us-west-2": "0.226"
    }
  },
  {
    "instance_type": "m5.12xlarge",
    "vCPU": 48,
    "memory": 192,
    "pricing": {
      "ap-northeast-1": "2.976",
      "ap-northeast-2": "2.832",
      "ap-south-1": "2.424",
      "ap-southeast-1": "2.88",
      "ap-southeast-2": "2.88",
      "ca-central-1": "2.568",
      "eu-central-1": "2.76",
      "eu-west-1": "2.568",
      "eu-west-2": "2.664",
      "eu-west-3": "2.688",
      "sa-east-1": "3.672",
      "us-east-1": "2.304",
      "us-east-2": "2.304",
      "us-gov-east-1": "2.904",
      "us-gov-west-1": "2.904",
      "us-west-1": "2.688",
      "us-west-2": "2.304"
    }
  },
  {
    "instance_type": "p3.16xlarge",
    "vCPU": 64,
    "memory": 488,
    "pricing": {
      "ap-northeast-1": "41.944",
      "ap-northeast-2": "39.848",
      "ap-southeast-1": "39.848",
      "ap-southeast-2": "39.848",
      "ca-central-1": "26.928",
      "eu-central-1": "30.584",
      "eu-west-1": "26.44",
      "eu-west-2": "28.712",
      "us-east-1": "24.48",
      "us-east-2": "24.48",
      "us-gov-west-1": "29.376",
      "us-west-2": "24.48"
    }
  },
  {
    "instance_type": "m3.large",
    "vCPU": 2,
    "memory": 7.5,
    "pricing": {
      "ap-northeast-1": "0.193",
      "ap-northeast-2": "0.183",
      "ap-northeast-3": null,
      "ap-southeast-1": "0.196",
      "ap-southeast-2": "0.186",
      "eu-central-1": "0.158",
      "eu-west-1": "0.146",
      "sa-east-1": "0.19",
      "us-east-1": "0.133",
      "us-gov-west-1": "0.168",
      "us-west-1": "0.154",
      "us-west-2": "0.133"
    }
  },
  {
    "instance_type": "c5d.4xlarge",
    "vCPU": 16,
    "memory": 32,
    "pricing": {
      "ap-northeast-1": "0.976",
      "ap-northeast-2": "0.88",
      "ap-southeast-1": "0.896",
      "ap-southeast-2": "1.008",
      "ca-central-1": "0.848",
      "eu-central-1": "0.888",
      "eu-west-1": "0.872",
      "eu-west-2": "0.92",
      "us-east-1": "0.768",
      "us-east-2": "0.768",
      "us-gov-west-1": "0.928",
      "us-west-1": "0.96",
      "us-west-2": "0.768"
    }
  },
  {
    "instance_type": "r5.24xlarge",
    "vCPU": 96,
    "memory": 768,
    "pricing": {
      "ap-northeast-1": "7.296",
      "ap-northeast-2": "7.296",
      "ap-southeast-1": "7.296",
      "ap-southeast-2": "7.248",
      "ca-central-1": "6.624",
      "eu-central-1": "7.296",
      "eu-west-1": "6.768",
      "eu-west-2": "7.104",
      "eu-west-3": "7.104",
      "us-east-1": "6.048",
      "us-east-2": "6.048",
      "us-gov-east-1": "7.248",
      "us-gov-west-1": "7.248",
      "us-west-1": "6.72",
      "us-west-2": "6.048"
    }
  },
  {
    "instance_type": "r4.large",
    "vCPU": 2,
    "memory": 15.25,
    "pricing": {
      "ap-northeast-1": "0.16",
      "ap-northeast-2": "0.16",
      "ap-northeast-3": "0.16",
      "ap-south-1": "0.137",
      "ap-southeast-1": "0.16",
      "ap-southeast-2": "0.1596",
      "ca-central-1": "0.146",
      "eu-central-1": "0.16005",
      "eu-west-1": "0.1482",
      "eu-west-2": "0.156",
      "eu-west-3": "0.156",
      "sa-east-1": "0.28",
      "us-east-1": "0.133",
      "us-east-2": "0.133",
      "us-gov-west-1": "0.1596",
      "us-west-1": "0.1482",
      "us-west-2": "0.133"
    }
  },
  {
    "instance_type": "r3.xlarge",
    "vCPU": 4,
    "memory": 30.5,
    "pricing": {
      "ap-northeast-1": "0.399",
      "ap-northeast-2": "0.399",
      "ap-northeast-3": null,
      "ap-south-1": "0.379",
      "ap-southeast-1": "0.399",
      "ap-southeast-2": "0.399",
      "eu-central-1": "0.4",
      "eu-west-1": "0.371",
      "sa-east-1": "0.7",
      "us-east-1": "0.333",
      "us-east-2": "0.333",
      "us-gov-west-1": "0.399",
      "us-west-1": "0.371",
      "us-west-2": "0.333"
    }
  },
  {
    "instance_type": "x1e.32xlarge",
    "vCPU": 128,
    "memory": 3904,
    "pricing": {
      "ap-northeast-1": "38.688",
      "ap-southeast-2": "38.688",
      "eu-central-1": "37.344",
      "eu-west-1": "32",
      "us-east-1": "26.688",
      "us-gov-west-1": "32",
      "us-west-2": "26.688"
    }
  },
  {
    "instance_type": "c5n.xlarge",
    "vCPU": 4,
    "memory": 10.5,
    "pricing": {
      "eu-west-1": "0.244",
      "us-east-1": "0.216",
      "us-east-2": "0.216",
      "us-gov-west-1": "0.26",
      "us-west-2": "0.216"
    }
  },
  {
    "instance_type": "m5a.4xlarge",
    "vCPU": 16,
    "memory": 64,
    "pricing": {
      "ap-southeast-1": "0.864",
      "eu-west-1": "0.768",
      "us-east-1": "0.688",
      "us-east-2": "0.688",
      "us-west-2": "0.688"
    }
  },
  {
    "instance_type": "t2.xlarge",
    "vCPU": 4,
    "memory": 16,
    "pricing": {
      "ap-northeast-1": "0.2432",
      "ap-northeast-2": "0.2304",
      "ap-northeast-3": "0.2432",
      "ap-south-1": "0.1984",
      "ap-southeast-1": "0.2336",
      "ap-southeast-2": "0.2336",
      "ca-central-1": "0.2048",
      "eu-central-1": "0.2144",
      "eu-west-1": "0.2016",
      "eu-west-2": "0.2112",
      "eu-west-3": "0.2112",
      "sa-east-1": "0.2976",
      "us-east-1": "0.1856",
      "us-east-2": "0.1856",
      "us-gov-west-1": "0.2176",
      "us-west-1": "0.2208",
      "us-west-2": "0.1856"
    }
  },
  {
    "instance_type": "p2.16xlarge",
    "vCPU": 64,
    "memory": 768,
    "pricing": {
      "ap-northeast-1": "24.672",
      "ap-northeast-2": "23.44",
      "ap-south-1": "27.488",
      "ap-southeast-1": "27.488",
      "ap-southeast-2": "24.672",
      "eu-central-1": "21.216",
      "eu-west-1": "15.552",
      "us-east-1": "14.4",
      "us-east-2": "14.4",
      "us-gov-west-1": "17.28",
      "us-west-2": "14.4"
    }
  },
  {
    "instance_type": "c3.8xlarge",
    "vCPU": 32,
    "memory": 60,
    "pricing": {
      "ap-northeast-1": "2.043",
      "ap-northeast-2": "1.839",
      "ap-northeast-3": null,
      "ap-southeast-1": "2.117",
      "ap-southeast-2": "2.117",
      "eu-central-1": "2.064",
      "eu-west-1": "1.912",
      "sa-east-1": "2.6",
      "us-east-1": "1.68",
      "us-gov-west-1": "2.016",
      "us-west-1": "1.912",
      "us-west-2": "1.68"
    }
  },
  {
    "instance_type": "m3.medium",
    "vCPU": 1,
    "memory": 3.75,
    "pricing": {
      "ap-northeast-1": "0.096",
      "ap-northeast-2": "0.091",
      "ap-northeast-3": null,
      "ap-southeast-1": "0.098",
      "ap-southeast-2": "0.093",
      "eu-central-1": "0.079",
      "eu-west-1": "0.073",
      "sa-east-1": "0.095",
      "us-east-1": "0.067",
      "us-gov-west-1": "0.084",
      "us-west-1": "0.077",
      "us-west-2": "0.067"
    }
  },
  {
    "instance_type": "x1e.2xlarge",
    "vCPU": 8,
    "memory": 244,
    "pricing": {
      "ap-northeast-1": "2.418",
      "ap-southeast-2": "2.418",
      "eu-central-1": "2.334",
      "eu-west-1": "2",
      "us-east-1": "1.668",
      "us-gov-west-1": "2",
      "us-west-2": "1.668"
    }
  },
  {
    "instance_type": "r5a.2xlarge",
    "vCPU": 8,
    "memory": 64,
    "pricing": {
      "ap-southeast-1": "0.544",
      "eu-west-1": "0.508",
      "us-east-1": "0.452",
      "us-east-2": "0.452",
      "us-west-2": "0.452"
    }
  },
  {
    "instance_type": "m5a.12xlarge",
    "vCPU": 48,
    "memory": 192,
    "pricing": {
      "ap-southeast-1": "2.592",
      "eu-west-1": "2.304",
      "us-east-1": "2.064",
      "us-east-2": "2.064",
      "us-west-2": "2.064"
    }
  },
  {
    "instance_type": "t2.small",
    "vCPU": 1,
    "memory": 2,
    "pricing": {
      "ap-northeast-1": "0.0304",
      "ap-northeast-2": "0.0288",
      "ap-northeast-3": "0.0304",
      "ap-south-1": "0.0248",
      "ap-southeast-1": "0.0292",
      "ap-southeast-2": "0.0292",
      "ca-central-1": "0.0256",
      "eu-central-1": "0.0268",
      "eu-west-1": "0.025",
      "eu-west-2": "0.026",
      "eu-west-3": "0.0264",
      "sa-east-1": "0.0372",
      "us-east-1": "0.023",
      "us-east-2": "0.023",
      "us-gov-west-1": "0.0272",
      "us-west-1": "0.0276",
      "us-west-2": "0.023"
    }
  },
  {
    "instance_type": "d2.xlarge",
    "vCPU": 4,
    "memory": 30.5,
    "pricing": {
      "ap-northeast-1": "0.844",
      "ap-northeast-2": "0.844",
      "ap-northeast-3": "0.844",
      "ap-south-1": "0.827",
      "ap-southeast-1": "0.87",
      "ap-southeast-2": "0.87",
      "ca-central-1": "0.759",
      "eu-central-1": "0.794",
      "eu-west-1": "0.735",
      "eu-west-2": "0.772",
      "eu-west-3": "0.772",
      "us-east-1": "0.69",
      "us-east-2": "0.69",
      "us-gov-west-1": "0.828",
      "us-west-1": "0.781",
      "us-west-2": "0.69"
    }
  },
  {
    "instance_type": "c5.2xlarge",
    "vCPU": 8,
    "memory": 16,
    "pricing": {
      "ap-northeast-1": "0.428",
      "ap-northeast-2": "0.384",
      "ap-south-1": "0.34",
      "ap-southeast-1": "0.392",
      "ap-southeast-2": "0.444",
      "ca-central-1": "0.372",
      "eu-central-1": "0.388",
      "eu-west-1": "0.384",
      "eu-west-2": "0.404",
      "eu-west-3": "0.404",
      "sa-east-1": "0.524",
      "us-east-1": "0.34",
      "us-east-2": "0.34",
      "us-gov-east-1": "0.408",
      "us-gov-west-1": "0.408",
      "us-west-1": "0.424",
      "us-west-2": "0.34"
    }
  },
  {
    "instance_type": "t3.medium",
    "vCPU": 2,
    "memory": 4,
    "pricing": {
      "ap-northeast-1": "0.0544",
      "ap-northeast-2": "0.052",
      "ap-southeast-1": "0.0528",
      "ap-southeast-2": "0.0528",
      "ca-central-1": "0.0464",
      "eu-central-1": "0.048",
      "eu-west-1": "0.0456",
      "eu-west-2": "0.0472",
      "eu-west-3": "0.0472",
      "sa-east-1": "0.0672",
      "us-east-1": "0.0416",
      "us-east-2": "0.0416",
      "us-gov-east-1": "0.0488",
      "us-gov-west-1": "0.0488",
      "us-west-1": "0.0496",
      "us-west-2": "0.0416"
    }
  },
  {
    "instance_type": "m1.xlarge",
    "vCPU": 4,
    "memory": 15,
    "pricing": {
      "ap-northeast-1": "0.486",
      "ap-southeast-1": "0.467",
      "ap-southeast-2": "0.467",
      "eu-west-1": "0.379",
      "sa-east-1": "0.467",
      "us-east-1": "0.35",
      "us-gov-west-1": "0.423",
      "us-west-1": "0.379",
      "us-west-2": "0.35"
    }
  },
  {
    "instance_type": "c5d.2xlarge",
    "vCPU": 8,
    "memory": 16,
    "pricing": {
      "ap-northeast-1": "0.488",
      "ap-northeast-2": "0.44",
      "ap-southeast-1": "0.448",
      "ap-southeast-2": "0.504",
      "ca-central-1": "0.424",
      "eu-central-1": "0.444",
      "eu-west-1": "0.436",
      "eu-west-2": "0.46",
      "us-east-1": "0.384",
      "us-east-2": "0.384",
      "us-gov-west-1": "0.464",
      "us-west-1": "0.48",
      "us-west-2": "0.384"
    }
  },
  {
    "instance_type": "m3.xlarge",
    "vCPU": 4,
    "memory": 15,
    "pricing": {
      "ap-northeast-1": "0.385",
      "ap-northeast-2": "0.366",
      "ap-northeast-3": null,
      "ap-southeast-1": "0.392",
      "ap-southeast-2": "0.372",
      "eu-central-1": "0.315",
      "eu-west-1": "0.293",
      "sa-east-1": "0.381",
      "us-east-1": "0.266",
      "us-gov-west-1": "0.336",
      "us-west-1": "0.308",
      "us-west-2": "0.266"
    }
  },
  {
    "instance_type": "m1.medium",
    "vCPU": 1,
    "memory": 3.75,
    "pricing": {
      "ap-northeast-1": "0.122",
      "ap-southeast-1": "0.117",
      "ap-southeast-2": "0.117",
      "eu-west-1": "0.095",
      "sa-east-1": "0.117",
      "us-east-1": "0.087",
      "us-gov-west-1": "0.106",
      "us-west-1": "0.095",
      "us-west-2": "0.087"
    }
  },
  {
    "instance_type": "r4.16xlarge",
    "vCPU": 64,
    "memory": 488,
    "pricing": {
      "ap-northeast-1": "5.12",
      "ap-northeast-2": "5.12",
      "ap-northeast-3": "5.12",
      "ap-south-1": "4.384",
      "ap-southeast-1": "5.12",
      "ap-southeast-2": "5.1072",
      "ca-central-1": "4.672",
      "eu-central-1": "5.1216",
      "eu-west-1": "4.7424",
      "eu-west-2": "4.992",
      "eu-west-3": "4.992",
      "sa-east-1": "8.96",
      "us-east-1": "4.256",
      "us-east-2": "4.256",
      "us-gov-west-1": "5.1072",
      "us-west-1": "4.7424",
      "us-west-2": "4.256"
    }
  },
  {
    "instance_type": "i3.xlarge",
    "vCPU": 4,
    "memory": 30.5,
    "pricing": {
      "ap-northeast-1": "0.366",
      "ap-northeast-2": "0.366",
      "ap-northeast-3": "0.366",
      "ap-south-1": "0.354",
      "ap-southeast-1": "0.374",
      "ap-southeast-2": "0.374",
      "ca-central-1": "0.344",
      "eu-central-1": "0.372",
      "eu-west-1": "0.344",
      "eu-west-2": "0.362",
      "eu-west-3": "0.362",
      "sa-east-1": "0.572",
      "us-east-1": "0.312",
      "us-east-2": "0.312",
      "us-gov-east-1": "0.376",
      "us-gov-west-1": "0.376",
      "us-west-1": "0.344",
      "us-west-2": "0.312"
    }
  },
  {
    "instance_type": "z1d.2xlarge",
    "vCPU": 8,
    "memory": 64,
    "pricing": {
      "ap-northeast-1": "0.908",
      "ap-southeast-1": "0.904",
      "eu-west-1": "0.832",
      "us-east-1": "0.744",
      "us-west-1": "0.844",
      "us-west-2": "0.744"
    }
  },
  {
    "instance_type": "c3.xlarge",
    "vCPU": 4,
    "memory": 7.5,
    "pricing": {
      "ap-northeast-1": "0.255",
      "ap-northeast-2": "0.23",
      "ap-northeast-3": null,
      "ap-southeast-1": "0.265",
      "ap-southeast-2": "0.265",
      "eu-central-1": "0.258",
      "eu-west-1": "0.239",
      "sa-east-1": "0.325",
      "us-east-1": "0.21",
      "us-gov-west-1": "0.252",
      "us-west-1": "0.239",
      "us-west-2": "0.21"
    }
  },
  {
    "instance_type": "c3.2xlarge",
    "vCPU": 8,
    "memory": 15,
    "pricing": {
      "ap-northeast-1": "0.511",
      "ap-northeast-2": "0.46",
      "ap-northeast-3": null,
      "ap-southeast-1": "0.529",
      "ap-southeast-2": "0.529",
      "eu-central-1": "0.516",
      "eu-west-1": "0.478",
      "sa-east-1": "0.65",
      "us-east-1": "0.42",
      "us-gov-west-1": "0.504",
      "us-west-1": "0.478",
      "us-west-2": "0.42"
    }
  },
  {
    "instance_type": "r3.2xlarge",
    "vCPU": 8,
    "memory": 61,
    "pricing": {
      "ap-northeast-1": "0.798",
      "ap-northeast-2": "0.798",
      "ap-northeast-3": null,
      "ap-south-1": "0.758",
      "ap-southeast-1": "0.798",
      "ap-southeast-2": "0.798",
      "eu-central-1": "0.8",
      "eu-west-1": "0.741",
      "sa-east-1": "1.399",
      "us-east-1": "0.665",
      "us-east-2": "0.665",
      "us-gov-west-1": "0.798",
      "us-west-1": "0.741",
      "us-west-2": "0.665"
    }
  },
  {
    "instance_type": "r4.2xlarge",
    "vCPU": 8,
    "memory": 61,
    "pricing": {
      "ap-northeast-1": "0.64",
      "ap-northeast-2": "0.64",
      "ap-northeast-3": "0.64",
      "ap-south-1": "0.548",
      "ap-southeast-1": "0.64",
      "ap-southeast-2": "0.6384",
      "ca-central-1": "0.584",
      "eu-central-1": "0.6402",
      "eu-west-1": "0.5928",
      "eu-west-2": "0.624",
      "eu-west-3": "0.624",
      "sa-east-1": "1.12",
      "us-east-1": "0.532",
      "us-east-2": "0.532",
      "us-gov-west-1": "0.6384",
      "us-west-1": "0.5928",
      "us-west-2": "0.532"
    }
  },
  {
    "instance_type": "p2.8xlarge",
    "vCPU": 32,
    "memory": 488,
    "pricing": {
      "ap-northeast-1": "12.336",
      "ap-northeast-2": "11.72",
      "ap-south-1": "13.744",
      "ap-southeast-1": "13.744",
      "ap-southeast-2": "12.336",
      "eu-central-1": "10.608",
      "eu-west-1": "7.776",
      "us-east-1": "7.2",
      "us-east-2": "7.2",
      "us-gov-west-1": "8.64",
      "us-west-2": "7.2"
    }
  },
  {
    "instance_type": "m5.xlarge",
    "vCPU": 4,
    "memory": 16,
    "pricing": {
      "ap-northeast-1": "0.248",
      "ap-northeast-2": "0.236",
      "ap-south-1": "0.202",
      "ap-southeast-1": "0.24",
      "ap-southeast-2": "0.24",
      "ca-central-1": "0.214",
      "eu-central-1": "0.23",
      "eu-west-1": "0.214",
      "eu-west-2": "0.222",
      "eu-west-3": "0.224",
      "sa-east-1": "0.306",
      "us-east-1": "0.192",
      "us-east-2": "0.192",
      "us-gov-east-1": "0.242",
      "us-gov-west-1": "0.242",
      "us-west-1": "0.224",
      "us-west-2": "0.192"
    }
  },
  {
    "instance_type": "c4.8xlarge",
    "vCPU": 36,
    "memory": 60,
    "pricing": {
      "ap-northeast-1": "2.016",
      "ap-northeast-2": "1.815",
      "ap-northeast-3": "2.016",
      "ap-south-1": "1.6",
      "ap-southeast-1": "1.848",
      "ap-southeast-2": "2.085",
      "ca-central-1": "1.75",
      "eu-central-1": "1.817",
      "eu-west-1": "1.811",
      "eu-west-2": "1.902",
      "sa-east-1": "2.47",
      "us-east-1": "1.591",
      "us-east-2": "1.591",
      "us-gov-west-1": "1.915",
      "us-west-1": "1.993",
      "us-west-2": "1.591"
    }
  },
  {
    "instance_type": "c5n.9xlarge",
    "vCPU": 36,
    "memory": 96,
    "pricing": {
      "eu-west-1": "2.196",
      "us-east-1": "1.944",
      "us-east-2": "1.944",
      "us-gov-west-1": "2.34",
      "us-west-2": "1.944"
    }
  },
  {
    "instance_type": "d2.2xlarge",
    "vCPU": 8,
    "memory": 61,
    "pricing": {
      "ap-northeast-1": "1.688",
      "ap-northeast-2": "1.688",
      "ap-northeast-3": "1.688",
      "ap-south-1": "1.653",
      "ap-southeast-1": "1.74",
      "ap-southeast-2": "1.74",
      "ca-central-1": "1.518",
      "eu-central-1": "1.588",
      "eu-west-1": "1.47",
      "eu-west-2": "1.544",
      "eu-west-3": "1.544",
      "us-east-1": "1.38",
      "us-east-2": "1.38",
      "us-gov-west-1": "1.656",
      "us-west-1": "1.563",
      "us-west-2": "1.38"
    }
  },
  {
    "instance_type": "d2.4xlarge",
    "vCPU": 16,
    "memory": 122,
    "pricing": {
      "ap-northeast-1": "3.376",
      "ap-northeast-2": "3.376",
      "ap-northeast-3": "3.376",
      "ap-south-1": "3.306",
      "ap-southeast-1": "3.48",
      "ap-southeast-2": "3.48",
      "ca-central-1": "3.036",
      "eu-central-1": "3.176",
      "eu-west-1": "2.94",
      "eu-west-2": "3.087",
      "eu-west-3": "3.088",
      "us-east-1": "2.76",
      "us-east-2": "2.76",
      "us-gov-west-1": "3.312",
      "us-west-1": "3.125",
      "us-west-2": "2.76"
    }
  },
  {
    "instance_type": "c5.xlarge",
    "vCPU": 4,
    "memory": 8,
    "pricing": {
      "ap-northeast-1": "0.214",
      "ap-northeast-2": "0.192",
      "ap-south-1": "0.17",
      "ap-southeast-1": "0.196",
      "ap-southeast-2": "0.222",
      "ca-central-1": "0.186",
      "eu-central-1": "0.194",
      "eu-west-1": "0.192",
      "eu-west-2": "0.202",
      "eu-west-3": "0.202",
      "sa-east-1": "0.262",
      "us-east-1": "0.17",
      "us-east-2": "0.17",
      "us-gov-east-1": "0.204",
      "us-gov-west-1": "0.204",
      "us-west-1": "0.212",
      "us-west-2": "0.17"
    }
  },
  {
    "instance_type": "m5d.4xlarge",
    "vCPU": 16,
    "memory": 64,
    "pricing": {
      "ap-northeast-1": "1.168",
      "ap-northeast-2": "1.112",
      "ap-southeast-1": "1.128",
      "ap-southeast-2": "1.136",
      "ca-central-1": "1.008",
      "eu-central-1": "1.088",
      "eu-west-1": "1.008",
      "eu-west-2": "1.048",
      "us-east-1": "0.904",
      "us-east-2": "0.904",
      "us-gov-west-1": "1.144",
      "us-west-1": "1.064",
      "us-west-2": "0.904"
    }
  },
  {
    "instance_type": "m4.xlarge",
    "vCPU": 4,
    "memory": 16,
    "pricing": {
      "ap-northeast-1": "0.258",
      "ap-northeast-2": "0.246",
      "ap-northeast-3": "0.258",
      "ap-south-1": "0.21",
      "ap-southeast-1": "0.25",
      "ap-southeast-2": "0.25",
      "ca-central-1": "0.222",
      "eu-central-1": "0.24",
      "eu-west-1": "0.222",
      "eu-west-2": "0.232",
      "sa-east-1": "0.318",
      "us-east-1": "0.2",
      "us-east-2": "0.2",
      "us-gov-west-1": "0.252",
      "us-west-1": "0.234",
      "us-west-2": "0.2"
    }
  },
  {
    "instance_type": "r5a.4xlarge",
    "vCPU": 16,
    "memory": 128,
    "pricing": {
      "ap-southeast-1": "1.088",
      "eu-west-1": "1.016",
      "us-east-1": "0.904",
      "us-east-2": "0.904",
      "us-west-2": "0.904"
    }
  },
  {
    "instance_type": "t3.xlarge",
    "vCPU": 4,
    "memory": 16,
    "pricing": {
      "ap-northeast-1": "0.2176",
      "ap-northeast-2": "0.208",
      "ap-southeast-1": "0.2112",
      "ap-southeast-2": "0.2112",
      "ca-central-1": "0.1856",
      "eu-central-1": "0.192",
      "eu-west-1": "0.1824",
      "eu-west-2": "0.1888",
      "eu-west-3": "0.1888",
      "sa-east-1": "0.2688",
      "us-east-1": "0.1664",
      "us-east-2": "0.1664",
      "us-gov-east-1": "0.1952",
      "us-gov-west-1": "0.1952",
      "us-west-1": "0.1984",
      "us-west-2": "0.1664"
    }
  },
  {
    "instance_type": "z1d.12xlarge",
    "vCPU": 48,
    "memory": 384,
    "pricing": {
      "ap-northeast-1": "5.448",
      "ap-southeast-1": "5.424",
      "eu-west-1": "4.992",
      "us-east-1": "4.464",
      "us-west-1": "5.064",
      "us-west-2": "4.464"
    }
  },
  {
    "instance_type": "m4.16xlarge",
    "vCPU": 64,
    "memory": 256,
    "pricing": {
      "ap-northeast-1": "4.128",
      "ap-northeast-2": "3.936",
      "ap-northeast-3": "4.128",
      "ap-south-1": "3.36",
      "ap-southeast-1": "4",
      "ap-southeast-2": "4",
      "ca-central-1": "3.552",
      "eu-central-1": "3.84",
      "eu-west-1": "3.552",
      "eu-west-2": "3.712",
      "sa-east-1": "5.088",
      "us-east-1": "3.2",
      "us-east-2": "3.2",
      "us-gov-west-1": "4.032",
      "us-west-1": "3.744",
      "us-west-2": "3.2"
    }
  },
  {
    "instance_type": "r5.12xlarge",
    "vCPU": 48,
    "memory": 384,
    "pricing": {
      "ap-northeast-1": "3.648",
      "ap-northeast-2": "3.648",
      "ap-southeast-1": "3.648",
      "ap-southeast-2": "3.624",
      "ca-central-1": "3.312",
      "eu-central-1": "3.648",
      "eu-west-1": "3.384",
      "eu-west-2": "3.552",
      "eu-west-3": "3.552",
      "us-east-1": "3.024",
      "us-east-2": "3.024",
      "us-gov-east-1": "3.624",
      "us-gov-west-1": "3.624",
      "us-west-1": "3.36",
      "us-west-2": "3.024"
    }
  },
  {
    "instance_type": "m4.4xlarge",
    "vCPU": 16,
    "memory": 64,
    "pricing": {
      "ap-northeast-1": "1.032",
      "ap-northeast-2": "0.984",
      "ap-northeast-3": "1.032",
      "ap-south-1": "0.84",
      "ap-southeast-1": "1",
      "ap-southeast-2": "1",
      "ca-central-1": "0.888",
      "eu-central-1": "0.96",
      "eu-west-1": "0.888",
      "eu-west-2": "0.928",
      "sa-east-1": "1.272",
      "us-east-1": "0.8",
      "us-east-2": "0.8",
      "us-gov-west-1": "1.008",
      "us-west-1": "0.936",
      "us-west-2": "0.8"
    }
  },
  {
    "instance_type": "r4.xlarge",
    "vCPU": 4,
    "memory": 30.5,
    "pricing": {
      "ap-northeast-1": "0.32",
      "ap-northeast-2": "0.32",
      "ap-northeast-3": "0.32",
      "ap-south-1": "0.274",
      "ap-southeast-1": "0.32",
      "ap-southeast-2": "0.3192",
      "ca-central-1": "0.292",
      "eu-central-1": "0.3201",
      "eu-west-1": "0.2964",
      "eu-west-2": "0.312",
      "eu-west-3": "0.312",
      "sa-east-1": "0.56",
      "us-east-1": "0.266",
      "us-east-2": "0.266",
      "us-gov-west-1": "0.3192",
      "us-west-1": "0.2964",
      "us-west-2": "0.266"
    }
  },
  {
    "instance_type": "m1.small",
    "vCPU": 1,
    "memory": 1.7,
    "pricing": {
      "ap-northeast-1": "0.061",
      "ap-southeast-1": "0.058",
      "ap-southeast-2": "0.058",
      "eu-west-1": "0.047",
      "sa-east-1": "0.058",
      "us-east-1": "0.044",
      "us-gov-west-1": "0.053",
      "us-west-1": "0.047",
      "us-west-2": "0.044"
    }
  },
  {
    "instance_type": "a1.xlarge",
    "vCPU": 4,
    "memory": 8,
    "pricing": {
      "eu-west-1": "0.1152",
      "us-east-1": "0.102",
      "us-east-2": "0.102",
      "us-west-2": "0.102"
    }
  },
  {
    "instance_type": "z1d.6xlarge",
    "vCPU": 24,
    "memory": 192,
    "pricing": {
      "ap-northeast-1": "2.724",
      "ap-southeast-1": "2.712",
      "eu-west-1": "2.496",
      "us-east-1": "2.232",
      "us-west-1": "2.532",
      "us-west-2": "2.232"
    }
  },
  {
    "instance_type": "r5d.4xlarge",
    "vCPU": 16,
    "memory": 128,
    "pricing": {
      "ap-northeast-1": "1.392",
      "ap-northeast-2": "1.384",
      "ap-southeast-1": "1.392",
      "ap-southeast-2": "1.392",
      "ca-central-1": "1.264",
      "eu-central-1": "1.384",
      "eu-west-1": "1.28",
      "eu-west-2": "1.352",
      "us-east-1": "1.152",
      "us-east-2": "1.152",
      "us-west-1": "1.296",
      "us-west-2": "1.152"
    }
  },
  {
    "instance_type": "p2.xlarge",
    "vCPU": 4,
    "memory": 61,
    "pricing": {
      "ap-northeast-1": "1.542",
      "ap-northeast-2": "1.465",
      "ap-south-1": "1.718",
      "ap-southeast-1": "1.718",
      "ap-southeast-2": "1.542",
      "eu-central-1": "1.326",
      "eu-west-1": "0.972",
      "us-east-1": "0.9",
      "us-east-2": "0.9",
      "us-gov-west-1": "1.08",
      "us-west-2": "0.9"
    }
  },
  {
    "instance_type": "c3.4xlarge",
    "vCPU": 16,
    "memory": 30,
    "pricing": {
      "ap-northeast-1": "1.021",
      "ap-northeast-2": "0.919",
      "ap-northeast-3": null,
      "ap-southeast-1": "1.058",
      "ap-southeast-2": "1.058",
      "eu-central-1": "1.032",
      "eu-west-1": "0.956",
      "sa-east-1": "1.3",
      "us-east-1": "0.84",
      "us-gov-west-1": "1.008",
      "us-west-1": "0.956",
      "us-west-2": "0.84"
    }
  },
  {
    "instance_type": "r4.4xlarge",
    "vCPU": 16,
    "memory": 122,
    "pricing": {
      "ap-northeast-1": "1.28",
      "ap-northeast-2": "1.28",
      "ap-northeast-3": "1.28",
      "ap-south-1": "1.096",
      "ap-southeast-1": "1.28",
      "ap-southeast-2": "1.2768",
      "ca-central-1": "1.168",
      "eu-central-1": "1.2804",
      "eu-west-1": "1.1856",
      "eu-west-2": "1.248",
      "eu-west-3": "1.248",
      "sa-east-1": "2.24",
      "us-east-1": "1.064",
      "us-east-2": "1.064",
      "us-gov-west-1": "1.2768",
      "us-west-1": "1.1856",
      "us-west-2": "1.064"
    }
  },
  {
    "instance_type": "r5a.xlarge",
    "vCPU": 4,
    "memory": 32,
    "pricing": {
      "ap-southeast-1": "0.272",
      "eu-west-1": "0.254",
      "us-east-1": "0.226",
      "us-east-2": "0.226",
      "us-west-2": "0.226"
    }
  },
  {
    "instance_type": "h1.2xlarge",
    "vCPU": 8,
    "memory": 32,
    "pricing": {
      "eu-west-1": "0.519",
      "us-east-1": "0.468",
      "us-east-2": "0.468",
      "us-west-2": "0.468"
    }
  },
  {
    "instance_type": "f1.4xlarge",
    "vCPU": 16,
    "memory": 244,
    "pricing": {
      "eu-west-1": "3.63",
      "us-east-1": "3.3",
      "us-gov-west-1": "3.96",
      "us-west-1": "3.826",
      "us-west-2": "3.3"
    }
  },
  {
    "instance_type": "r5a.12xlarge",
    "vCPU": 48,
    "memory": 384,
    "pricing": {
      "ap-southeast-1": "3.264",
      "eu-west-1": "3.048",
      "us-east-1": "2.712",
      "us-east-2": "2.712",
      "us-west-2": "2.712"
    }
  },
  {
    "instance_type": "g2.2xlarge",
    "vCPU": 8,
    "memory": 15,
    "pricing": {
      "ap-northeast-1": "0.898",
      "ap-northeast-2": "0.898",
      "ap-southeast-1": "1",
      "ap-southeast-2": "0.898",
      "eu-central-1": "0.772",
      "eu-west-1": "0.702",
      "us-east-1": "0.65",
      "us-west-1": "0.702",
      "us-west-2": "0.65"
    }
  },
  {
    "instance_type": "c4.2xlarge",
    "vCPU": 8,
    "memory": 15,
    "pricing": {
      "ap-northeast-1": "0.504",
      "ap-northeast-2": "0.454",
      "ap-northeast-3": "0.504",
      "ap-south-1": "0.4",
      "ap-southeast-1": "0.462",
      "ap-southeast-2": "0.522",
      "ca-central-1": "0.438",
      "eu-central-1": "0.454",
      "eu-west-1": "0.453",
      "eu-west-2": "0.476",
      "sa-east-1": "0.618",
      "us-east-1": "0.398",
      "us-east-2": "0.398",
      "us-gov-west-1": "0.479",
      "us-west-1": "0.498",
      "us-west-2": "0.398"
    }
  },
  {
    "instance_type": "x1e.8xlarge",
    "vCPU": 32,
    "memory": 976,
    "pricing": {
      "ap-northeast-1": "9.672",
      "ap-southeast-2": "9.672",
      "eu-central-1": "9.336",
      "eu-west-1": "8",
      "us-east-1": "6.672",
      "us-gov-west-1": "8",
      "us-west-2": "6.672"
    }
  },
  {
    "instance_type": "m5.4xlarge",
    "vCPU": 16,
    "memory": 64,
    "pricing": {
      "ap-northeast-1": "0.992",
      "ap-northeast-2": "0.944",
      "ap-south-1": "0.808",
      "ap-southeast-1": "0.96",
      "ap-southeast-2": "0.96",
      "ca-central-1": "0.856",
      "eu-central-1": "0.92",
      "eu-west-1": "0.856",
      "eu-west-2": "0.888",
      "eu-west-3": "0.896",
      "sa-east-1": "1.224",
      "us-east-1": "0.768",
      "us-east-2": "0.768",
      "us-gov-east-1": "0.968",
      "us-gov-west-1": "0.968",
      "us-west-1": "0.896",
      "us-west-2": "0.768"
    }
  },
  {
    "instance_type": "h1.16xlarge",
    "vCPU": 64,
    "memory": 256,
    "pricing": {
      "eu-west-1": "4.152",
      "us-east-1": "3.744",
      "us-east-2": "3.744",
      "us-west-2": "3.744"
    }
  },
  {
    "instance_type": "z1d.xlarge",
    "vCPU": 4,
    "memory": 32,
    "pricing": {
      "ap-northeast-1": "0.454",
      "ap-southeast-1": "0.452",
      "eu-west-1": "0.416",
      "us-east-1": "0.372",
      "us-west-1": "0.422",
      "us-west-2": "0.372"
    }
  },
  {
    "instance_type": "a1.2xlarge",
    "vCPU": 8,
    "memory": 16,
    "pricing": {
      "eu-west-1": "0.2304",
      "us-east-1": "0.204",
      "us-east-2": "0.204",
      "us-west-2": "0.204"
    }
  },
  {
    "instance_type": "t3.2xlarge",
    "vCPU": 8,
    "memory": 32,
    "pricing": {
      "ap-northeast-1": "0.4352",
      "ap-northeast-2": "0.416",
      "ap-southeast-1": "0.4224",
      "ap-southeast-2": "0.4224",
      "ca-central-1": "0.3712",
      "eu-central-1": "0.384",
      "eu-west-1": "0.3648",
      "eu-west-2": "0.3776",
      "eu-west-3": "0.3776",
      "sa-east-1": "0.5376",
      "us-east-1": "0.3328",
      "us-east-2": "0.3328",
      "us-gov-east-1": "0.3904",
      "us-gov-west-1": "0.3904",
      "us-west-1": "0.3968",
      "us-west-2": "0.3328"
    }
  },
  {
    "instance_type": "g3.8xlarge",
    "vCPU": 32,
    "memory": 244,
    "pricing": {
      "ap-northeast-1": "3.16",
      "ap-southeast-1": "3.34",
      "ap-southeast-2": "3.508",
      "ca-central-1": "2.832",
      "eu-central-1": "2.85",
      "eu-west-1": "2.42",
      "us-east-1": "2.28",
      "us-east-2": "2.28",
      "us-gov-west-1": "2.64",
      "us-west-1": "3.068",
      "us-west-2": "2.28"
    }
  },
  {
    "instance_type": "i3.metal",
    "vCPU": 72,
    "memory": 512,
    "pricing": {
      "eu-central-1": "5.952",
      "eu-west-1": "5.504",
      "eu-west-2": "5.792",
      "us-east-1": "4.992",
      "us-east-2": "4.992",
      "us-west-2": "4.992"
    }
  },
  {
    "instance_type": "m4.2xlarge",
    "vCPU": 8,
    "memory": 32,
    "pricing": {
      "ap-northeast-1": "0.516",
      "ap-northeast-2": "0.492",
      "ap-northeast-3": "0.516",
      "ap-south-1": "0.42",
      "ap-southeast-1": "0.5",
      "ap-southeast-2": "0.5",
      "ca-central-1": "0.444",
      "eu-central-1": "0.48",
      "eu-west-1": "0.444",
      "eu-west-2": "0.464",
      "sa-east-1": "0.636",
      "us-east-1": "0.4",
      "us-east-2": "0.4",
      "us-gov-west-1": "0.504",
      "us-west-1": "0.468",
      "us-west-2": "0.4"
    }
  },
  {
    "instance_type": "r5.4xlarge",
    "vCPU": 16,
    "memory": 128,
    "pricing": {
      "ap-northeast-1": "1.216",
      "ap-northeast-2": "1.216",
      "ap-southeast-1": "1.216",
      "ap-southeast-2": "1.208",
      "ca-central-1": "1.104",
      "eu-central-1": "1.216",
      "eu-west-1": "1.128",
      "eu-west-2": "1.184",
      "eu-west-3": "1.184",
      "us-east-1": "1.008",
      "us-east-2": "1.008",
      "us-gov-east-1": "1.208",
      "us-gov-west-1": "1.208",
      "us-west-1": "1.12",
      "us-west-2": "1.008"
    }
  },
  {
    "instance_type": "r5d.2xlarge",
    "vCPU": 8,
    "memory": 64,
    "pricing": {
      "ap-northeast-1": "0.696",
      "ap-northeast-2": "0.692",
      "ap-southeast-1": "0.696",
      "ap-southeast-2": "0.696",
      "ca-central-1": "0.632",
      "eu-central-1": "0.692",
      "eu-west-1": "0.64",
      "eu-west-2": "0.676",
      "us-east-1": "0.576",
      "us-east-2": "0.576",
      "us-west-1": "0.648",
      "us-west-2": "0.576"
    }
  },
  {
    "instance_type": "r5.large",
    "vCPU": 2,
    "memory": 16,
    "pricing": {
      "ap-northeast-1": "0.152",
      "ap-northeast-2": "0.152",
      "ap-southeast-1": "0.152",
      "ap-southeast-2": "0.151",
      "ca-central-1": "0.138",
      "eu-central-1": "0.152",
      "eu-west-1": "0.141",
      "eu-west-2": "0.148",
      "eu-west-3": "0.148",
      "us-east-1": "0.126",
      "us-east-2": "0.126",
      "us-gov-east-1": "0.151",
      "us-gov-west-1": "0.151",
      "us-west-1": "0.14",
      "us-west-2": "0.126"
    }
  },
  {
    "instance_type": "i3.large",
    "vCPU": 2,
    "memory": 15.25,
    "pricing": {
      "ap-northeast-1": "0.183",
      "ap-northeast-2": "0.183",
      "ap-northeast-3": "0.183",
      "ap-south-1": "0.177",
      "ap-southeast-1": "0.187",
      "ap-southeast-2": "0.187",
      "ca-central-1": "0.172",
      "eu-central-1": "0.186",
      "eu-west-1": "0.172",
      "eu-west-2": "0.181",
      "eu-west-3": "0.181",
      "sa-east-1": "0.286",
      "us-east-1": "0.156",
      "us-east-2": "0.156",
      "us-gov-east-1": "0.188",
      "us-gov-west-1": "0.188",
      "us-west-1": "0.172",
      "us-west-2": "0.156"
    }
  },
  {
    "instance_type": "z1d.large",
    "vCPU": 2,
    "memory": 16,
    "pricing": {
      "ap-northeast-1": "0.227",
      "ap-southeast-1": "0.226",
      "eu-west-1": "0.208",
      "us-east-1": "0.186",
      "us-west-1": "0.211",
      "us-west-2": "0.186"
    }
  },
  {
    "instance_type": "i3.16xlarge",
    "vCPU": 64,
    "memory": 488,
    "pricing": {
      "ap-northeast-1": "5.856",
      "ap-northeast-2": "5.856",
      "ap-northeast-3": "5.856",
      "ap-south-1": "5.664",
      "ap-southeast-1": "5.984",
      "ap-southeast-2": "5.984",
      "ca-central-1": "5.504",
      "eu-central-1": "5.952",
      "eu-west-1": "5.504",
      "eu-west-2": "5.792",
      "eu-west-3": "5.792",
      "sa-east-1": "9.152",
      "us-east-1": "4.992",
      "us-east-2": "4.992",
      "us-gov-east-1": "6.016",
      "us-gov-west-1": "6.016",
      "us-west-1": "5.504",
      "us-west-2": "4.992"
    }
  },
  {
    "instance_type": "i2.4xlarge",
    "vCPU": 16,
    "memory": 122,
    "pricing": {
      "ap-northeast-1": "4.002",
      "ap-northeast-2": "4.002",
      "ap-south-1": "3.867",
      "ap-southeast-1": "4.07",
      "ap-southeast-2": "4.07",
      "eu-central-1": "4.051",
      "eu-west-1": "3.751",
      "us-east-1": "3.41",
      "us-east-2": "3.41",
      "us-gov-west-1": "4.092",
      "us-west-1": "3.751",
      "us-west-2": "3.41"
    }
  },
  {
    "instance_type": "c5d.9xlarge",
    "vCPU": 36,
    "memory": 72,
    "pricing": {
      "ap-northeast-1": "2.196",
      "ap-northeast-2": "1.98",
      "ap-southeast-1": "2.016",
      "ap-southeast-2": "2.268",
      "ca-central-1": "1.908",
      "eu-central-1": "1.998",
      "eu-west-1": "1.962",
      "eu-west-2": "2.07",
      "us-east-1": "1.728",
      "us-east-2": "1.728",
      "us-gov-west-1": "2.088",
      "us-west-1": "2.16",
      "us-west-2": "1.728"
    }
  },
  {
    "instance_type": "r5.xlarge",
    "vCPU": 4,
    "memory": 32,
    "pricing": {
      "ap-northeast-1": "0.304",
      "ap-northeast-2": "0.304",
      "ap-southeast-1": "0.304",
      "ap-southeast-2": "0.302",
      "ca-central-1": "0.276",
      "eu-central-1": "0.304",
      "eu-west-1": "0.282",
      "eu-west-2": "0.296",
      "eu-west-3": "0.296",
      "us-east-1": "0.252",
      "us-east-2": "0.252",
      "us-gov-east-1": "0.302",
      "us-gov-west-1": "0.302",
      "us-west-1": "0.28",
      "us-west-2": "0.252"
    }
  },
  {
    "instance_type": "i3.2xlarge",
    "vCPU": 8,
    "memory": 61,
    "pricing": {
      "ap-northeast-1": "0.732",
      "ap-northeast-2": "0.732",
      "ap-northeast-3": "0.732",
      "ap-south-1": "0.708",
      "ap-southeast-1": "0.748",
      "ap-southeast-2": "0.748",
      "ca-central-1": "0.688",
      "eu-central-1": "0.744",
      "eu-west-1": "0.688",
      "eu-west-2": "0.724",
      "eu-west-3": "0.724",
      "sa-east-1": "1.144",
      "us-east-1": "0.624",
      "us-east-2": "0.624",
      "us-gov-east-1": "0.752",
      "us-gov-west-1": "0.752",
      "us-west-1": "0.688",
      "us-west-2": "0.624"
    }
  },
  {
    "instance_type": "r5d.large",
    "vCPU": 2,
    "memory": 16,
    "pricing": {
      "ap-northeast-1": "0.174",
      "ap-northeast-2": "0.173",
      "ap-southeast-1": "0.174",
      "ap-southeast-2": "0.174",
      "ca-central-1": "0.158",
      "eu-central-1": "0.173",
      "eu-west-1": "0.16",
      "eu-west-2": "0.169",
      "us-east-1": "0.144",
      "us-east-2": "0.144",
      "us-west-1": "0.162",
      "us-west-2": "0.144"
    }
  },
  {
    "instance_type": "g3.4xlarge",
    "vCPU": 16,
    "memory": 122,
    "pricing": {
      "ap-northeast-1": "1.58",
      "ap-southeast-1": "1.67",
      "ap-southeast-2": "1.754",
      "ca-central-1": "1.416",
      "eu-central-1": "1.425",
      "eu-west-1": "1.21",
      "us-east-1": "1.14",
      "us-east-2": "1.14",
      "us-gov-west-1": "1.32",
      "us-west-1": "1.534",
      "us-west-2": "1.14"
    }
  },
  {
    "instance_type": "r5d.12xlarge",
    "vCPU": 48,
    "memory": 384,
    "pricing": {
      "ap-northeast-1": "4.176",
      "ap-northeast-2": "4.152",
      "ap-southeast-1": "4.176",
      "ap-southeast-2": "4.176",
      "ca-central-1": "3.792",
      "eu-central-1": "4.152",
      "eu-west-1": "3.84",
      "eu-west-2": "4.056",
      "us-east-1": "3.456",
      "us-east-2": "3.456",
      "us-west-1": "3.888",
      "us-west-2": "3.456"
    }
  },
  {
    "instance_type": "g2.8xlarge",
    "vCPU": 32,
    "memory": 60,
    "pricing": {
      "ap-northeast-1": "3.592",
      "ap-northeast-2": "3.592",
      "ap-southeast-1": "4",
      "ap-southeast-2": "3.592",
      "eu-central-1": "3.088",
      "eu-west-1": "2.808",
      "us-east-1": "2.6",
      "us-west-1": "2.808",
      "us-west-2": "2.6"
    }
  },
  {
    "instance_type": "i3.4xlarge",
    "vCPU": 16,
    "memory": 122,
    "pricing": {
      "ap-northeast-1": "1.464",
      "ap-northeast-2": "1.464",
      "ap-northeast-3": "1.464",
      "ap-south-1": "1.416",
      "ap-southeast-1": "1.496",
      "ap-southeast-2": "1.496",
      "ca-central-1": "1.376",
      "eu-central-1": "1.488",
      "eu-west-1": "1.376",
      "eu-west-2": "1.448",
      "eu-west-3": "1.448",
      "sa-east-1": "2.288",
      "us-east-1": "1.248",
      "us-east-2": "1.248",
      "us-gov-east-1": "1.504",
      "us-gov-west-1": "1.504",
      "us-west-1": "1.376",
      "us-west-2": "1.248"
    }
  },
  {
    "instance_type": "c5n.4xlarge",
    "vCPU": 16,
    "memory": 42,
    "pricing": {
      "eu-west-1": "0.976",
      "us-east-1": "0.864",
      "us-east-2": "0.864",
      "us-gov-west-1": "1.04",
      "us-west-2": "0.864"
    }
  },
  {
    "instance_type": "r3.4xlarge",
    "vCPU": 16,
    "memory": 122,
    "pricing": {
      "ap-northeast-1": "1.596",
      "ap-northeast-2": "1.596",
      "ap-northeast-3": null,
      "ap-south-1": "1.516",
      "ap-southeast-1": "1.596",
      "ap-southeast-2": "1.596",
      "eu-central-1": "1.6",
      "eu-west-1": "1.482",
      "sa-east-1": "2.799",
      "us-east-1": "1.33",
      "us-east-2": "1.33",
      "us-gov-west-1": "1.596",
      "us-west-1": "1.482",
      "us-west-2": "1.33"
    }
  },
  {
    "instance_type": "m1.large",
    "vCPU": 2,
    "memory": 7.5,
    "pricing": {
      "ap-northeast-1": "0.243",
      "ap-southeast-1": "0.233",
      "ap-southeast-2": "0.233",
      "eu-west-1": "0.19",
      "sa-east-1": "0.233",
      "us-east-1": "0.175",
      "us-gov-west-1": "0.211",
      "us-west-1": "0.19",
      "us-west-2": "0.175"
    }
  },
  {
    "instance_type": "m2.4xlarge",
    "vCPU": 8,
    "memory": 68.4,
    "pricing": {
      "ap-northeast-1": "1.15",
      "ap-southeast-1": "1.183",
      "ap-southeast-2": "1.183",
      "eu-west-1": "1.1",
      "sa-east-1": "1.291",
      "us-east-1": "0.98",
      "us-gov-west-1": "1.171",
      "us-west-1": "1.1",
      "us-west-2": "0.98"
    }
  },
  {
    "instance_type": "r3.8xlarge",
    "vCPU": 32,
    "memory": 244,
    "pricing": {
      "ap-northeast-1": "3.192",
      "ap-northeast-2": "3.192",
      "ap-northeast-3": null,
      "ap-south-1": "3.032",
      "ap-southeast-1": "3.192",
      "ap-southeast-2": "3.192",
      "eu-central-1": "3.201",
      "eu-west-1": "2.964",
      "sa-east-1": "5.597",
      "us-east-1": "2.66",
      "us-east-2": "2.66",
      "us-gov-west-1": "3.192",
      "us-west-1": "2.964",
      "us-west-2": "2.66"
    }
  },
  {
    "instance_type": "p3.8xlarge",
    "vCPU": 32,
    "memory": 244,
    "pricing": {
      "ap-northeast-1": "20.972",
      "ap-northeast-2": "19.924",
      "ap-southeast-1": "19.924",
      "ap-southeast-2": "19.924",
      "ca-central-1": "13.464",
      "eu-central-1": "15.292",
      "eu-west-1": "13.22",
      "eu-west-2": "14.356",
      "us-east-1": "12.24",
      "us-east-2": "12.24",
      "us-gov-west-1": "14.688",
      "us-west-2": "12.24"
    }
  },
  {
    "instance_type": "m2.xlarge",
    "vCPU": 2,
    "memory": 17.1,
    "pricing": {
      "ap-northeast-1": "0.287",
      "ap-southeast-1": "0.296",
      "ap-southeast-2": "0.296",
      "eu-west-1": "0.275",
      "sa-east-1": "0.323",
      "us-east-1": "0.245",
      "us-gov-west-1": "0.293",
      "us-west-1": "0.275",
      "us-west-2": "0.245"
    }
  },
  {
    "instance_type": "c1.medium",
    "vCPU": 2,
    "memory": 1.7,
    "pricing": {
      "ap-northeast-1": "0.158",
      "ap-southeast-1": "0.164",
      "ap-southeast-2": "0.164",
      "eu-west-1": "0.148",
      "sa-east-1": "0.179",
      "us-east-1": "0.13",
      "us-gov-west-1": "0.157",
      "us-west-1": "0.148",
      "us-west-2": "0.13"
    }
  },
  {
    "instance_type": "cc2.8xlarge",
    "vCPU": 32,
    "memory": 60.5,
    "pricing": {
      "ap-northeast-1": "2.349",
      "eu-west-1": "2.25",
      "us-east-1": "2",
      "us-gov-west-1": "2.25",
      "us-west-2": "2"
    }
  },
  {
    "instance_type": "c1.xlarge",
    "vCPU": 8,
    "memory": 7,
    "pricing": {
      "ap-northeast-1": "0.632",
      "ap-southeast-1": "0.655",
      "ap-southeast-2": "0.655",
      "eu-west-1": "0.592",
      "sa-east-1": "0.718",
      "us-east-1": "0.52",
      "us-gov-west-1": "0.628",
      "us-west-1": "0.592",
      "us-west-2": "0.52"
    }
  },
  {
    "instance_type": "m2.2xlarge",
    "vCPU": 4,
    "memory": 34.2,
    "pricing": {
      "ap-northeast-1": "0.575",
      "ap-southeast-1": "0.592",
      "ap-southeast-2": "0.592",
      "eu-west-1": "0.55",
      "sa-east-1": "0.645",
      "us-east-1": "0.49",
      "us-gov-west-1": "0.586",
      "us-west-1": "0.55",
      "us-west-2": "0.49"
    }
  },
  {
    "instance_type": "m5a.large",
    "vCPU": 2,
    "memory": 8,
    "pricing": {
      "ap-southeast-1": "0.108",
      "eu-west-1": "0.096",
      "us-east-1": "0.086",
      "us-east-2": "0.086",
      "us-west-2": "0.086"
    }
  },
  {
    "instance_type": "i3.8xlarge",
    "vCPU": 32,
    "memory": 244,
    "pricing": {
      "ap-northeast-1": "2.928",
      "ap-northeast-2": "2.928",
      "ap-northeast-3": "2.928",
      "ap-south-1": "2.832",
      "ap-southeast-1": "2.992",
      "ap-southeast-2": "2.992",
      "ca-central-1": "2.752",
      "eu-central-1": "2.976",
      "eu-west-1": "2.752",
      "eu-west-2": "2.896",
      "eu-west-3": "2.896",
      "sa-east-1": "4.576",
      "us-east-1": "2.496",
      "us-east-2": "2.496",
      "us-gov-east-1": "3.008",
      "us-gov-west-1": "3.008",
      "us-west-1": "2.752",
      "us-west-2": "2.496"
    }
  },
  {
    "instance_type": "cr1.8xlarge",
    "vCPU": 32,
    "memory": 244,
    "pricing": {
      "ap-northeast-1": "4.105",
      "eu-west-1": "3.75",
      "us-east-1": "3.5",
      "us-west-2": "3.5"
    }
  },
  {
    "instance_type": "m5a.24xlarge",
    "vCPU": 96,
    "memory": 384,
    "pricing": {
      "ap-southeast-1": "5.184",
      "eu-west-1": "4.608",
      "us-east-1": "4.128",
      "us-east-2": "4.128",
      "us-west-2": "4.128"
    }
  },
  {
    "instance_type": "c4.4xlarge",
    "vCPU": 16,
    "memory": 30,
    "pricing": {
      "ap-northeast-1": "1.008",
      "ap-northeast-2": "0.907",
      "ap-northeast-3": "1.008",
      "ap-south-1": "0.8",
      "ap-southeast-1": "0.924",
      "ap-southeast-2": "1.042",
      "ca-central-1": "0.876",
      "eu-central-1": "0.909",
      "eu-west-1": "0.905",
      "eu-west-2": "0.95",
      "sa-east-1": "1.235",
      "us-east-1": "0.796",
      "us-east-2": "0.796",
      "us-gov-west-1": "0.958",
      "us-west-1": "0.997",
      "us-west-2": "0.796"
    }
  },
  {
    "instance_type": "t3.small",
    "vCPU": 2,
    "memory": 2,
    "pricing": {
      "ap-northeast-1": "0.0272",
      "ap-northeast-2": "0.026",
      "ap-southeast-1": "0.0264",
      "ap-southeast-2": "0.0264",
      "ca-central-1": "0.0232",
      "eu-central-1": "0.024",
      "eu-west-1": "0.0228",
      "eu-west-2": "0.0236",
      "eu-west-3": "0.0236",
      "sa-east-1": "0.0336",
      "us-east-1": "0.0208",
      "us-east-2": "0.0208",
      "us-gov-east-1": "0.0244",
      "us-gov-west-1": "0.0244",
      "us-west-1": "0.0248",
      "us-west-2": "0.0208"
    }
  },
  {
    "instance_type": "f1.16xlarge",
    "vCPU": 64,
    "memory": 976,
    "pricing": {
      "eu-west-1": "14.52",
      "us-east-1": "13.2",
      "us-gov-west-1": "15.84",
      "us-west-1": "15.304",
      "us-west-2": "13.2"
    }
  },
  {
    "instance_type": "t3.micro",
    "vCPU": 2,
    "memory": 1,
    "pricing": {
      "ap-northeast-1": "0.0136",
      "ap-northeast-2": "0.013",
      "ap-southeast-1": "0.0132",
      "ap-southeast-2": "0.0132",
      "ca-central-1": "0.0116",
      "eu-central-1": "0.012",
      "eu-west-1": "0.0114",
      "eu-west-2": "0.0118",
      "eu-west-3": "0.0118",
      "sa-east-1": "0.0168",
      "us-east-1": "0.0104",
      "us-east-2": "0.0104",
      "us-gov-east-1": "0.0122",
      "us-gov-west-1": "0.0122",
      "us-west-1": "0.0124",
      "us-west-2": "0.0104"
    }
  },
  {
    "instance_type": "t3.large",
    "vCPU": 2,
    "memory": 8,
    "pricing": {
      "ap-northeast-1": "0.1088",
      "ap-northeast-2": "0.104",
      "ap-southeast-1": "0.1056",
      "ap-southeast-2": "0.1056",
      "ca-central-1": "0.0928",
      "eu-central-1": "0.096",
      "eu-west-1": "0.0912",
      "eu-west-2": "0.0944",
      "eu-west-3": "0.0944",
      "sa-east-1": "0.1344",
      "us-east-1": "0.0832",
      "us-east-2": "0.0832",
      "us-gov-east-1": "0.0976",
      "us-gov-west-1": "0.0976",
      "us-west-1": "0.0992",
      "us-west-2": "0.0832"
    }
  },
  {
    "instance_type": "g3s.xlarge",
    "vCPU": 4,
    "memory": 30.5,
    "pricing": {
      "ap-northeast-1": "1.04",
      "ap-southeast-2": "1.154",
      "eu-central-1": "0.938",
      "eu-west-1": "0.796",
      "us-east-1": "0.75",
      "us-east-2": "0.75",
      "us-west-1": "1.009",
      "us-west-2": "0.75"
    }
  },
  {
    "instance_type": "f1.2xlarge",
    "vCPU": 8,
    "memory": 122,
    "pricing": {
      "eu-west-1": "1.815",
      "us-east-1": "1.65",
      "us-gov-west-1": "1.98",
      "us-west-1": "1.913",
      "us-west-2": "1.65"
    }
  },
  {
    "instance_type": "a1.medium",
    "vCPU": 1,
    "memory": 2,
    "pricing": {
      "eu-west-1": "0.0288",
      "us-east-1": "0.0255",
      "us-east-2": "0.0255",
      "us-west-2": "0.0255"
    }
  },
  {
    "instance_type": "t2.medium",
    "vCPU": 2,
    "memory": 4,
    "pricing": {
      "ap-northeast-1": "0.0608",
      "ap-northeast-2": "0.0576",
      "ap-northeast-3": "0.0608",
      "ap-south-1": "0.0496",
      "ap-southeast-1": "0.0584",
      "ap-southeast-2": "0.0584",
      "ca-central-1": "0.0512",
      "eu-central-1": "0.0536",
      "eu-west-1": "0.05",
      "eu-west-2": "0.052",
      "eu-west-3": "0.0528",
      "sa-east-1": "0.0744",
      "us-east-1": "0.0464",
      "us-east-2": "0.0464",
      "us-gov-west-1": "0.0544",
      "us-west-1": "0.0552",
      "us-west-2": "0.0464"
    }
  },
  {
    "instance_type": "a1.large",
    "vCPU": 2,
    "memory": 4,
    "pricing": {
      "eu-west-1": "0.0576",
      "us-east-1": "0.051",
      "us-east-2": "0.051",
      "us-west-2": "0.051"
    }
  },
  {
    "instance_type": "hs1.8xlarge",
    "vCPU": 17,
    "memory": 117,
    "pricing": {
      "ap-northeast-1": "5.4",
      "ap-southeast-1": "5.57",
      "ap-southeast-2": "5.57",
      "eu-west-1": "4.9",
      "us-east-1": "4.6",
      "us-gov-west-1": "5.52",
      "us-west-2": "4.6"
    }
  },
  {
    "instance_type": "t1.micro",
    "vCPU": 1,
    "memory": 0.613,
    "pricing": {
      "ap-northeast-1": "0.026",
      "ap-southeast-1": "0.02",
      "ap-southeast-2": "0.02",
      "eu-west-1": "0.02",
      "sa-east-1": "0.027",
      "us-east-1": "0.02",
      "us-gov-west-1": "0.024",
      "us-west-1": "0.025",
      "us-west-2": "0.02"
    }
  },
  {
    "instance_type": "u-6tb1.metal",
    "vCPU": 448,
    "memory": 6144,
    "pricing": {}
  },
  {
    "instance_type": "a1.4xlarge",
    "vCPU": 16,
    "memory": 32,
    "pricing": {
      "eu-west-1": "0.4608",
      "us-east-1": "0.408",
      "us-east-2": "0.408",
      "us-west-2": "0.408"
    }
  },
  {
    "instance_type": "t3.nano",
    "vCPU": 2,
    "memory": 0.5,
    "pricing": {
      "ap-northeast-1": "0.0068",
      "ap-northeast-2": "0.0065",
      "ap-southeast-1": "0.0066",
      "ap-southeast-2": "0.0066",
      "ca-central-1": "0.0058",
      "eu-central-1": "0.006",
      "eu-west-1": "0.0057",
      "eu-west-2": "0.0059",
      "eu-west-3": "0.0059",
      "sa-east-1": "0.0084",
      "us-east-1": "0.0052",
      "us-east-2": "0.0052",
      "us-gov-east-1": "0.0061",
      "us-gov-west-1": "0.0061",
      "us-west-1": "0.0062",
      "us-west-2": "0.0052"
    }
  },
  {
    "instance_type": "t2.nano",
    "vCPU": 1,
    "memory": 0.5,
    "pricing": {
      "ap-northeast-1": "0.0076",
      "ap-northeast-2": "0.0072",
      "ap-northeast-3": "0.0076",
      "ap-south-1": "0.0062",
      "ap-southeast-1": "0.0073",
      "ap-southeast-2": "0.0073",
      "ca-central-1": "0.0064",
      "eu-central-1": "0.0067",
      "eu-west-1": "0.0063",
      "eu-west-2": "0.0066",
      "eu-west-3": "0.0066",
      "sa-east-1": "0.0093",
      "us-east-1": "0.0058",
      "us-east-2": "0.0058",
      "us-gov-west-1": "0.0068",
      "us-west-1": "0.0069",
      "us-west-2": "0.0058"
    }
  },
  {
    "instance_type": "u-9tb1.metal",
    "vCPU": 448,
    "memory": 9216,
    "pricing": {}
  },
  {
    "instance_type": "u-12tb1.metal",
    "vCPU": 448,
    "memory": 12288,
    "pricing": {}
  }
]
`
