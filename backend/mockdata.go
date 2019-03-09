package backend

import (
	"encoding/json"
	"strconv"
)

// mock of refreshTable
func mockRefreshTable() (err error) {
	// introduce delays and possible error
	err = mockDelayWithPossibleError()
	if err != nil {
		log.Errorf("mock error refreshing table")
		return
	}

	// we only need to load initial test data when cachedTable is empty
	if len(cachedTable) == 0 {
		err = json.Unmarshal([]byte(mockEnvDetailsJSON), &cachedTable)
		if err != nil {
			log.Fatalf("mock API is enabled, but can't unmarshal json file: %s", err)
		}
	}

	if experimentalEnabled {
		calculateEnvBills()
		cachedTableTemp := cachedTable
		cachedTable = cachedTable[:0]
		for _, env := range cachedTableTemp {
			for _, instanceObj := range env.Instances {
				// determine instance cpu and memory
				if details, found := getInstanceTypeDetails(instanceObj.InstanceType); found {
					instanceObj.MemoryGB = details.MemoryGB
					instanceObj.VCPU = details.VCPU
					if pricingstr, ok := details.PricingHourlyByRegion[instanceObj.Region]; ok {
						pricing, err := strconv.ParseFloat(pricingstr, 64)
						if err != nil {
							log.Errorf("failed to parse pricing info to float: %s", pricingstr)
						}
						instanceObj.PricingHourly = pricing
					}
				}
				if validateEnvName(instanceObj.Environment) {
					addInstance(&instanceObj)
				}
			}
		}
	}

	updateEnvDetails()
	log.Debugf("MOCK: valid environment(s) in cache: %d", len(cachedTable))
	return
}

// test data for cachedTable
const mockEnvDetailsJSON = `
[
  {
    "id": "356f6265efcc",
    "instances": [
      {
        "environment": "mockenv1",
        "region": "ca-central-1",
        "id": "2e4f6625ed4f",
        "instance_id": "i-0a5c68f08c53f5a81",
        "instance_type": "c5.xlarge",
        "memory_gb": 8,
        "name": "mockenv1-org3peer4",
        "state": "stopped",
        "vcpu": 4
      },
      {
        "environment": "mockenv1",
        "region": "ca-central-1",
        "id": "d428482da3c6",
        "instance_id": "i-0ab609a3c8fb70025",
        "instance_type": "c5.xlarge",
        "memory_gb": 8,
        "name": "mockenv1-org3peer3",
        "state": "stopped",
        "vcpu": 4
      },
      {
        "environment": "mockenv1",
        "region": "ca-central-1",
        "id": "62559ac6a49f",
        "instance_id": "i-00061d4c8875b263d",
        "instance_type": "c5.2xlarge",
        "memory_gb": 16,
        "name": "mockenv1-org2peer2",
        "state": "stopped",
        "vcpu": 8
      },
      {
        "environment": "mockenv1",
        "region": "ca-central-1",
        "id": "7a629ac90e94",
        "instance_id": "i-0d1fee52c9633fac6",
        "instance_type": "c5.xlarge",
        "memory_gb": 8,
        "name": "mockenv1-org3peer1",
        "state": "stopped",
        "vcpu": 4
      },
      {
        "environment": "mockenv1",
        "region": "ca-central-1",
        "id": "69370f81dc94",
        "instance_id": "i-04a4c3c42912783bf",
        "instance_type": "c5.2xlarge",
        "memory_gb": 16,
        "name": "mockenv1-org3peer2",
        "state": "stopped",
        "vcpu": 8
      },
      {
        "environment": "mockenv1",
        "region": "ca-central-1",
        "id": "e37539798466",
        "instance_id": "i-00b7f2acde2b257f4",
        "instance_type": "t2.medium",
        "memory_gb": 4,
        "name": "mockenv1-kafka1",
        "state": "stopped",
        "vcpu": 2
      },
      {
        "environment": "mockenv1",
        "region": "ca-central-1",
        "id": "a6259e9e5b8c",
        "instance_id": "i-0f5dd015ec55df23d",
        "instance_type": "c5.xlarge",
        "memory_gb": 8,
        "name": "mockenv1-org2peer3",
        "state": "stopped",
        "vcpu": 4
      },
      {
        "environment": "mockenv1",
        "region": "ca-central-1",
        "id": "34dbab61afac",
        "instance_id": "i-01e281e66008412c5",
        "instance_type": "c5.2xlarge",
        "memory_gb": 16,
        "name": "mockenv1-org4peer2",
        "state": "stopped",
        "vcpu": 8
      },
      {
        "environment": "mockenv1",
        "region": "ca-central-1",
        "id": "29a2cdd81788",
        "instance_id": "i-023d466a8c9802aed",
        "instance_type": "t2.2xlarge",
        "memory_gb": 32,
        "name": "mockenv1-admin",
        "state": "stopped",
        "vcpu": 8
      },
      {
        "environment": "mockenv1",
        "region": "ca-central-1",
        "id": "78453ab3186e",
        "instance_id": "i-0a7e0901d8a388830",
        "instance_type": "c5.xlarge",
        "memory_gb": 8,
        "name": "mockenv1-org2peer1",
        "state": "stopped",
        "vcpu": 4
      },
      {
        "environment": "mockenv1",
        "region": "ca-central-1",
        "id": "d374e4df5632",
        "instance_id": "i-00e9d9726ae79b57c",
        "instance_type": "c5.xlarge",
        "memory_gb": 8,
        "name": "mockenv1-org4peer5",
        "state": "stopped",
        "vcpu": 4
      },
      {
        "environment": "mockenv1",
        "region": "ca-central-1",
        "id": "d9d04aae31d1",
        "instance_id": "i-067d9e1b712c97808",
        "instance_type": "t2.micro",
        "memory_gb": 1,
        "name": "mockenv1-orderer1",
        "state": "stopped",
        "vcpu": 1
      },
      {
        "environment": "mockenv1",
        "region": "ca-central-1",
        "id": "c796ea6fe680",
        "instance_id": "i-0607964069781218d",
        "instance_type": "c5.xlarge",
        "memory_gb": 8,
        "name": "mockenv1-org3peer5",
        "state": "stopped",
        "vcpu": 4
      },
      {
        "environment": "mockenv1",
        "region": "ca-central-1",
        "id": "a07d171267e6",
        "instance_id": "i-0b9dff025ebe98954",
        "instance_type": "c5.xlarge",
        "memory_gb": 8,
        "name": "mockenv1-org2peer4",
        "state": "stopped",
        "vcpu": 4
      },
      {
        "environment": "mockenv1",
        "region": "ca-central-1",
        "id": "03aa1724fcfa",
        "instance_id": "i-0db1f3649cc30252a",
        "instance_type": "c5.xlarge",
        "memory_gb": 8,
        "name": "mockenv1-org1peer1",
        "state": "stopped",
        "vcpu": 4
      },
      {
        "environment": "mockenv1",
        "region": "ca-central-1",
        "id": "e9e5b6224663",
        "instance_id": "i-006cbdad6c27666b5",
        "instance_type": "c5.xlarge",
        "memory_gb": 8,
        "name": "mockenv1-org1peer5",
        "state": "stopped",
        "vcpu": 4
      },
      {
        "environment": "mockenv1",
        "region": "ca-central-1",
        "id": "bf2193deedd0",
        "instance_id": "i-0b635062614f41902",
        "instance_type": "c5.xlarge",
        "memory_gb": 8,
        "name": "mockenv1-org1peer4",
        "state": "stopped",
        "vcpu": 4
      },
      {
        "environment": "mockenv1",
        "region": "ca-central-1",
        "id": "02a5e4d01641",
        "instance_id": "i-0ce165e54c97354af",
        "instance_type": "c5.2xlarge",
        "memory_gb": 16,
        "name": "mockenv1-org1peer2",
        "state": "stopped",
        "vcpu": 8
      },
      {
        "environment": "mockenv1",
        "region": "ca-central-1",
        "id": "740d75980033",
        "instance_id": "i-092bf46daa1ba864d",
        "instance_type": "t2.micro",
        "memory_gb": 1,
        "name": "mockenv1-orderer2",
        "state": "stopped",
        "vcpu": 1
      },
      {
        "environment": "mockenv1",
        "region": "ca-central-1",
        "id": "b6386966b577",
        "instance_id": "i-007fe1dee4659994d",
        "instance_type": "t2.small",
        "memory_gb": 2,
        "name": "mockenv1-zoo2",
        "state": "stopped",
        "vcpu": 1
      },
      {
        "environment": "mockenv1",
        "region": "ca-central-1",
        "id": "f931c27c9f88",
        "instance_id": "i-02a865c984a00457b",
        "instance_type": "c5.xlarge",
        "memory_gb": 8,
        "name": "mockenv1-org4peer3",
        "state": "stopped",
        "vcpu": 4
      },
      {
        "environment": "mockenv1",
        "region": "ca-central-1",
        "id": "9dd331c3de9c",
        "instance_id": "i-0fcd4b9773bd76d98",
        "instance_type": "t2.medium",
        "memory_gb": 4,
        "name": "mockenv1-kafka2",
        "state": "stopped",
        "vcpu": 2
      },
      {
        "environment": "mockenv1",
        "region": "ca-central-1",
        "id": "776872084bcf",
        "instance_id": "i-0537759a63512bbb7",
        "instance_type": "c5.xlarge",
        "memory_gb": 8,
        "name": "mockenv1-org4peer1",
        "state": "stopped",
        "vcpu": 4
      },
      {
        "environment": "mockenv1",
        "region": "ca-central-1",
        "id": "53faeeb6a0eb",
        "instance_id": "i-081c2a979f70ef34e",
        "instance_type": "c5.xlarge",
        "memory_gb": 8,
        "name": "mockenv1-org4peer4",
        "state": "stopped",
        "vcpu": 4
      },
      {
        "environment": "mockenv1",
        "region": "ca-central-1",
        "id": "cac7286dbc01",
        "instance_id": "i-03e949f0b8df6b8fa",
        "instance_type": "c5.xlarge",
        "memory_gb": 8,
        "name": "mockenv1-org1peer3",
        "state": "stopped",
        "vcpu": 4
      },
      {
        "environment": "mockenv1",
        "region": "ca-central-1",
        "id": "66d2149e9634",
        "instance_id": "i-06ec7f5c26d0a21ac",
        "instance_type": "t2.small",
        "memory_gb": 2,
        "name": "mockenv1-zoo1",
        "state": "stopped",
        "vcpu": 1
      },
      {
        "environment": "mockenv1",
        "region": "ca-central-1",
        "id": "c188b66b7c3a",
        "instance_id": "i-04b0e0df81ddbf819",
        "instance_type": "c5.xlarge",
        "memory_gb": 8,
        "name": "mockenv1-org2peer5",
        "state": "stopped",
        "vcpu": 4
      }
    ],
    "name": "mockenv1",
    "provider": "aws",
    "region": "ca-central-1",
    "running_instances": 0,
    "state": "stopped",
    "stopped_instances": 27,
    "total_instances": 27,
    "total_memory_gb": 238,
    "total_vcpu": 112
  },
  {
    "id": "944d72e873f2",
    "instances": [
      {
        "environment": "mockenv2",
        "region": "ca-central-1",
        "id": "67cfa28a1663",
        "instance_id": "i-0655e4242278bf53c",
        "instance_type": "t2.small",
        "memory_gb": 2,
        "name": "mockenv2-org3admin",
        "state": "stopped",
        "vcpu": 1
      },
      {
        "environment": "mockenv2",
        "region": "ca-central-1",
        "id": "02c0a17846d2",
        "instance_id": "i-0beb81b9b12b095bd",
        "instance_type": "t2.small",
        "memory_gb": 2,
        "name": "mockenv2-org5app2",
        "state": "stopped",
        "vcpu": 1
      },
      {
        "environment": "mockenv2",
        "region": "ca-central-1",
        "id": "1af2b2d6210b",
        "instance_id": "i-047abbc75e0df27c0",
        "instance_type": "t2.micro",
        "memory_gb": 1,
        "name": "mockenv2-org6ca1",
        "state": "stopped",
        "vcpu": 1
      },
      {
        "environment": "mockenv2",
        "region": "ca-central-1",
        "id": "8dccf21bc575",
        "instance_id": "i-011b515645ad4dc00",
        "instance_type": "t2.small",
        "memory_gb": 2,
        "name": "mockenv2-kafka2",
        "state": "stopped",
        "vcpu": 1
      },
      {
        "environment": "mockenv2",
        "region": "ca-central-1",
        "id": "72c3926c63e4",
        "instance_id": "i-090b1a428ec4bf9eb",
        "instance_type": "t2.micro",
        "memory_gb": 1,
        "name": "mockenv2-org4ca1",
        "state": "stopped",
        "vcpu": 1
      },
      {
        "environment": "mockenv2",
        "region": "ca-central-1",
        "id": "4918922522aa",
        "instance_id": "i-0e7abdd4b4081803d",
        "instance_type": "t2.small",
        "memory_gb": 2,
        "name": "mockenv2-org6admin",
        "state": "stopped",
        "vcpu": 1
      },
      {
        "environment": "mockenv2",
        "region": "ca-central-1",
        "id": "6c1c373e8c1b",
        "instance_id": "i-0a89bd1a56bd3fd9f",
        "instance_type": "c5.large",
        "memory_gb": 4,
        "name": "mockenv2-org3peer2",
        "state": "stopped",
        "vcpu": 2
      },
      {
        "environment": "mockenv2",
        "region": "ca-central-1",
        "id": "8aad5b358018",
        "instance_id": "i-097f9c01c92880f4a",
        "instance_type": "t2.small",
        "memory_gb": 2,
        "name": "mockenv2-org4admin",
        "state": "stopped",
        "vcpu": 1
      },
      {
        "environment": "mockenv2",
        "region": "ca-central-1",
        "id": "20dbf22acd3b",
        "instance_id": "i-058935adf9d3e022d",
        "instance_type": "t2.small",
        "memory_gb": 2,
        "name": "mockenv2-org5app1",
        "state": "stopped",
        "vcpu": 1
      },
      {
        "environment": "mockenv2",
        "region": "ca-central-1",
        "id": "c6fd96b989ce",
        "instance_id": "i-0e3e419c9c5c5b033",
        "instance_type": "t2.small",
        "memory_gb": 2,
        "name": "mockenv2-orderer2",
        "state": "stopped",
        "vcpu": 1
      },
      {
        "environment": "mockenv2",
        "region": "ca-central-1",
        "id": "a4b5b0a4fa37",
        "instance_id": "i-0a345f92722456c2d",
        "instance_type": "t2.micro",
        "memory_gb": 1,
        "name": "mockenv2-org4ca2",
        "state": "stopped",
        "vcpu": 1
      },
      {
        "environment": "mockenv2",
        "region": "ca-central-1",
        "id": "1fdfe474d073",
        "instance_id": "i-0963676849262ad9b",
        "instance_type": "t2.small",
        "memory_gb": 2,
        "name": "mockenv2-org2app1",
        "state": "stopped",
        "vcpu": 1
      },
      {
        "environment": "mockenv2",
        "region": "ca-central-1",
        "id": "541a3812283d",
        "instance_id": "i-0a36afd331e1c8707",
        "instance_type": "t2.small",
        "memory_gb": 2,
        "name": "mockenv2-org4app1",
        "state": "stopped",
        "vcpu": 1
      },
      {
        "environment": "mockenv2",
        "region": "ca-central-1",
        "id": "d372dd3eb695",
        "instance_id": "i-0473f26821165011b",
        "instance_type": "t2.micro",
        "memory_gb": 1,
        "name": "mockenv2-admin",
        "state": "stopped",
        "vcpu": 1
      },
      {
        "environment": "mockenv2",
        "region": "ca-central-1",
        "id": "819583a44a9c",
        "instance_id": "i-0275344d2e6448169",
        "instance_type": "t2.small",
        "memory_gb": 2,
        "name": "mockenv2-zoo2",
        "state": "stopped",
        "vcpu": 1
      },
      {
        "environment": "mockenv2",
        "region": "ca-central-1",
        "id": "3705490defbf",
        "instance_id": "i-0fedcaac13e21f830",
        "instance_type": "t2.small",
        "memory_gb": 2,
        "name": "mockenv2-org2app2",
        "state": "stopped",
        "vcpu": 1
      },
      {
        "environment": "mockenv2",
        "region": "ca-central-1",
        "id": "7816f5bba557",
        "instance_id": "i-045641ef1ced05a75",
        "instance_type": "t2.micro",
        "memory_gb": 1,
        "name": "mockenv2-org5ca1",
        "state": "stopped",
        "vcpu": 1
      },
      {
        "environment": "mockenv2",
        "region": "ca-central-1",
        "id": "4cd25585a0be",
        "instance_id": "i-03cf891551793bb19",
        "instance_type": "c5.large",
        "memory_gb": 4,
        "name": "mockenv2-org4peer2",
        "state": "stopped",
        "vcpu": 2
      },
      {
        "environment": "mockenv2",
        "region": "ca-central-1",
        "id": "3fec87e8a9b3",
        "instance_id": "i-0daaed23ff0e0863c",
        "instance_type": "t2.small",
        "memory_gb": 2,
        "name": "mockenv2-zoo1",
        "state": "stopped",
        "vcpu": 1
      },
      {
        "environment": "mockenv2",
        "region": "ca-central-1",
        "id": "0e14edd9340c",
        "instance_id": "i-020bd33c68e581b66",
        "instance_type": "t2.small",
        "memory_gb": 2,
        "name": "mockenv2-orderer1",
        "state": "stopped",
        "vcpu": 1
      },
      {
        "environment": "mockenv2",
        "region": "ca-central-1",
        "id": "6e4cc5b2c5ed",
        "instance_id": "i-03300f1580f195ae2",
        "instance_type": "t2.micro",
        "memory_gb": 1,
        "name": "mockenv2-org1ca1",
        "state": "stopped",
        "vcpu": 1
      },
      {
        "environment": "mockenv2",
        "region": "ca-central-1",
        "id": "1b4c23d33ece",
        "instance_id": "i-0203967889caf3e6e",
        "instance_type": "c5.large",
        "memory_gb": 4,
        "name": "mockenv2-org6peer2",
        "state": "stopped",
        "vcpu": 2
      },
      {
        "environment": "mockenv2",
        "region": "ca-central-1",
        "id": "6b3fb980e4e8",
        "instance_id": "i-0ec603a71d27d4f0f",
        "instance_type": "c5.large",
        "memory_gb": 4,
        "name": "mockenv2-org5peer1",
        "state": "stopped",
        "vcpu": 2
      },
      {
        "environment": "mockenv2",
        "region": "ca-central-1",
        "id": "501ac75066f8",
        "instance_id": "i-0416dd9eceedf6c15",
        "instance_type": "c5.large",
        "memory_gb": 4,
        "name": "mockenv2-org6peer1",
        "state": "stopped",
        "vcpu": 2
      },
      {
        "environment": "mockenv2",
        "region": "ca-central-1",
        "id": "85607652104d",
        "instance_id": "i-0e0a36c172e5b77d4",
        "instance_type": "t2.small",
        "memory_gb": 2,
        "name": "mockenv2-org4app2",
        "state": "stopped",
        "vcpu": 1
      },
      {
        "environment": "mockenv2",
        "region": "ca-central-1",
        "id": "d61daf373255",
        "instance_id": "i-035852a154db41563",
        "instance_type": "t2.small",
        "memory_gb": 2,
        "name": "mockenv2-zoo3",
        "state": "stopped",
        "vcpu": 1
      },
      {
        "environment": "mockenv2",
        "region": "ca-central-1",
        "id": "3f88e9d7b8c8",
        "instance_id": "i-01b1e49f8934e18ea",
        "instance_type": "t2.small",
        "memory_gb": 2,
        "name": "mockenv2-kafka1",
        "state": "stopped",
        "vcpu": 1
      },
      {
        "environment": "mockenv2",
        "region": "ca-central-1",
        "id": "f3e3801b9c67",
        "instance_id": "i-020627cfcce463843",
        "instance_type": "t2.small",
        "memory_gb": 2,
        "name": "mockenv2-org2admin",
        "state": "stopped",
        "vcpu": 1
      },
      {
        "environment": "mockenv2",
        "region": "ca-central-1",
        "id": "3ca2efc1b6bb",
        "instance_id": "i-061d3922415a19815",
        "instance_type": "t2.small",
        "memory_gb": 2,
        "name": "mockenv2-org5admin",
        "state": "stopped",
        "vcpu": 1
      },
      {
        "environment": "mockenv2",
        "region": "ca-central-1",
        "id": "753d23277247",
        "instance_id": "i-034562753e3a38e68",
        "instance_type": "t2.micro",
        "memory_gb": 1,
        "name": "mockenv2-org6ca2",
        "state": "stopped",
        "vcpu": 1
      },
      {
        "environment": "mockenv2",
        "region": "ca-central-1",
        "id": "b9d7e388cdc1",
        "instance_id": "i-0f5b5445292b07afc",
        "instance_type": "c5.large",
        "memory_gb": 4,
        "name": "mockenv2-org2peer2",
        "state": "stopped",
        "vcpu": 2
      },
      {
        "environment": "mockenv2",
        "region": "ca-central-1",
        "id": "c34afe8f6301",
        "instance_id": "i-0cf3a247c4fabf863",
        "instance_type": "t2.small",
        "memory_gb": 2,
        "name": "mockenv2-kafka3",
        "state": "stopped",
        "vcpu": 1
      },
      {
        "environment": "mockenv2",
        "region": "ca-central-1",
        "id": "840d9f42e10c",
        "instance_id": "i-065a19d641e34e950",
        "instance_type": "t2.micro",
        "memory_gb": 1,
        "name": "mockenv2-org2ca1",
        "state": "stopped",
        "vcpu": 1
      },
      {
        "environment": "mockenv2",
        "region": "ca-central-1",
        "id": "16e6c2fb0bf6",
        "instance_id": "i-0015edfc2a8cccc6d",
        "instance_type": "t2.micro",
        "memory_gb": 1,
        "name": "mockenv2-org5ca2",
        "state": "stopped",
        "vcpu": 1
      },
      {
        "environment": "mockenv2",
        "region": "ca-central-1",
        "id": "72a284a8954b",
        "instance_id": "i-0a94c398e3526a635",
        "instance_type": "t2.small",
        "memory_gb": 2,
        "name": "mockenv2-kafka4",
        "state": "stopped",
        "vcpu": 1
      },
      {
        "environment": "mockenv2",
        "region": "ca-central-1",
        "id": "a86e7df62e18",
        "instance_id": "i-0d253caf816c362fb",
        "instance_type": "t2.small",
        "memory_gb": 2,
        "name": "mockenv2-org1admin",
        "state": "stopped",
        "vcpu": 1
      },
      {
        "environment": "mockenv2",
        "region": "ca-central-1",
        "id": "c6fc18b5e2a9",
        "instance_id": "i-09529203675e50ba4",
        "instance_type": "c5.large",
        "memory_gb": 4,
        "name": "mockenv2-org1peer1",
        "state": "stopped",
        "vcpu": 2
      },
      {
        "environment": "mockenv2",
        "region": "ca-central-1",
        "id": "09e8a4914d94",
        "instance_id": "i-0b24846f4a1d19177",
        "instance_type": "c5.large",
        "memory_gb": 4,
        "name": "mockenv2-org3peer1",
        "state": "stopped",
        "vcpu": 2
      },
      {
        "environment": "mockenv2",
        "region": "ca-central-1",
        "id": "23b47c794ae6",
        "instance_id": "i-0adf79a2e32e0bca8",
        "instance_type": "t2.micro",
        "memory_gb": 1,
        "name": "mockenv2-org1ca2",
        "state": "stopped",
        "vcpu": 1
      },
      {
        "environment": "mockenv2",
        "region": "ca-central-1",
        "id": "01d3da0d8663",
        "instance_id": "i-096a5f9780c78dcfb",
        "instance_type": "t2.small",
        "memory_gb": 2,
        "name": "mockenv2-org3app2",
        "state": "stopped",
        "vcpu": 1
      },
      {
        "environment": "mockenv2",
        "region": "ca-central-1",
        "id": "1e974c348d57",
        "instance_id": "i-0b8b0192c8e989471",
        "instance_type": "c5.large",
        "memory_gb": 4,
        "name": "mockenv2-org4peer1",
        "state": "stopped",
        "vcpu": 2
      },
      {
        "environment": "mockenv2",
        "region": "ca-central-1",
        "id": "0c0dfb6b206d",
        "instance_id": "i-07d051737f3ac0937",
        "instance_type": "c5.large",
        "memory_gb": 4,
        "name": "mockenv2-org1peer2",
        "state": "stopped",
        "vcpu": 2
      },
      {
        "environment": "mockenv2",
        "region": "ca-central-1",
        "id": "520212ae9e38",
        "instance_id": "i-0adee2dbd9228b2cb",
        "instance_type": "t2.micro",
        "memory_gb": 1,
        "name": "mockenv2-org3ca1",
        "state": "stopped",
        "vcpu": 1
      },
      {
        "environment": "mockenv2",
        "region": "ca-central-1",
        "id": "624d4be9668b",
        "instance_id": "i-0175ff2b9410c5d4e",
        "instance_type": "t2.small",
        "memory_gb": 2,
        "name": "mockenv2-org6app1",
        "state": "stopped",
        "vcpu": 1
      },
      {
        "environment": "mockenv2",
        "region": "ca-central-1",
        "id": "2c158eab73e0",
        "instance_id": "i-06113a14aa50fc45b",
        "instance_type": "t2.micro",
        "memory_gb": 1,
        "name": "mockenv2-org3ca2",
        "state": "stopped",
        "vcpu": 1
      },
      {
        "environment": "mockenv2",
        "region": "ca-central-1",
        "id": "5716a1e722dc",
        "instance_id": "i-0e14f1405355c4da9",
        "instance_type": "t2.micro",
        "memory_gb": 1,
        "name": "mockenv2-org2ca2",
        "state": "stopped",
        "vcpu": 1
      },
      {
        "environment": "mockenv2",
        "region": "ca-central-1",
        "id": "fdb12eb75302",
        "instance_id": "i-042d6e25f9758773b",
        "instance_type": "t2.small",
        "memory_gb": 2,
        "name": "mockenv2-org6app2",
        "state": "stopped",
        "vcpu": 1
      },
      {
        "environment": "mockenv2",
        "region": "ca-central-1",
        "id": "39692c417594",
        "instance_id": "i-0a1a722c320915c3d",
        "instance_type": "t2.small",
        "memory_gb": 2,
        "name": "mockenv2-org1app2",
        "state": "stopped",
        "vcpu": 1
      },
      {
        "environment": "mockenv2",
        "region": "ca-central-1",
        "id": "7c103bd1d392",
        "instance_id": "i-0f7faa94638ccef20",
        "instance_type": "t2.small",
        "memory_gb": 2,
        "name": "mockenv2-orderer3",
        "state": "stopped",
        "vcpu": 1
      },
      {
        "environment": "mockenv2",
        "region": "ca-central-1",
        "id": "584442b420e1",
        "instance_id": "i-0b23ddf592991b8ee",
        "instance_type": "t2.small",
        "memory_gb": 2,
        "name": "mockenv2-org3app1",
        "state": "stopped",
        "vcpu": 1
      },
      {
        "environment": "mockenv2",
        "region": "ca-central-1",
        "id": "eb05fe01c0ec",
        "instance_id": "i-0ccc10ef212856dd3",
        "instance_type": "c5.large",
        "memory_gb": 4,
        "name": "mockenv2-org2peer1",
        "state": "stopped",
        "vcpu": 2
      },
      {
        "environment": "mockenv2",
        "region": "ca-central-1",
        "id": "150ac25c4eac",
        "instance_id": "i-0e528e724ceef797b",
        "instance_type": "t2.small",
        "memory_gb": 2,
        "name": "mockenv2-org1app1",
        "state": "stopped",
        "vcpu": 1
      },
      {
        "environment": "mockenv2",
        "region": "ca-central-1",
        "id": "24de7ffae987",
        "instance_id": "i-009f5086307f4e57c",
        "instance_type": "c5.large",
        "memory_gb": 4,
        "name": "mockenv2-org5peer2",
        "state": "stopped",
        "vcpu": 2
      }
    ],
    "name": "mockenv2",
    "provider": "aws",
    "region": "ca-central-1",
    "running_instances": 0,
    "state": "stopped",
    "stopped_instances": 53,
    "total_instances": 53,
    "total_memory_gb": 117,
    "total_vcpu": 65
  },
  {
    "id": "77713d1aa299",
    "instances": [
      {
        "environment": "mockenv3",
        "region": "ca-central-1",
        "id": "368d0c1a098a",
        "instance_id": "i-04a0c55913889362d",
        "instance_type": "t2.medium",
        "memory_gb": 4,
        "name": "mockenv3-org2app1",
        "state": "stopped",
        "vcpu": 2
      },
      {
        "environment": "mockenv3",
        "region": "ca-central-1",
        "id": "b99e88285b48",
        "instance_id": "i-0b5eb71f782dc761a",
        "instance_type": "c5.2xlarge",
        "memory_gb": 16,
        "name": "mockenv3-org3couchdb1",
        "state": "stopped",
        "vcpu": 8
      },
      {
        "environment": "mockenv3",
        "region": "ca-central-1",
        "id": "ddeb5d37d839",
        "instance_id": "i-04d62321c72da87d8",
        "instance_type": "t2.medium",
        "memory_gb": 4,
        "name": "mockenv3-org4app1",
        "state": "stopped",
        "vcpu": 2
      },
      {
        "environment": "mockenv3",
        "region": "ca-central-1",
        "id": "a0924362bec8",
        "instance_id": "i-0e992c1a81a293ff6",
        "instance_type": "c5.2xlarge",
        "memory_gb": 16,
        "name": "mockenv3-org3peer3",
        "state": "stopped",
        "vcpu": 8
      },
      {
        "environment": "mockenv3",
        "region": "ca-central-1",
        "id": "28e77c87f2e9",
        "instance_id": "i-041ec7c305cbb61db",
        "instance_type": "t2.micro",
        "memory_gb": 1,
        "name": "mockenv3-org3ca1",
        "state": "stopped",
        "vcpu": 1
      },
      {
        "environment": "mockenv3",
        "region": "ca-central-1",
        "id": "86efa8cf0a70",
        "instance_id": "i-0cbca67848c5c8db8",
        "instance_type": "t2.medium",
        "memory_gb": 4,
        "name": "mockenv3-org1app2",
        "state": "stopped",
        "vcpu": 2
      },
      {
        "environment": "mockenv3",
        "region": "ca-central-1",
        "id": "9a8125db390f",
        "instance_id": "i-077bcb143f38cbcfe",
        "instance_type": "c5.xlarge",
        "memory_gb": 8,
        "name": "mockenv3-org2admin",
        "state": "stopped",
        "vcpu": 4
      },
      {
        "environment": "mockenv3",
        "region": "ca-central-1",
        "id": "981694c9be97",
        "instance_id": "i-04d04ff500e6521af",
        "instance_type": "c5.2xlarge",
        "memory_gb": 16,
        "name": "mockenv3-org3peer4",
        "state": "stopped",
        "vcpu": 8
      },
      {
        "environment": "mockenv3",
        "region": "ca-central-1",
        "id": "7c32299b2106",
        "instance_id": "i-0ed6ffc835322e749",
        "instance_type": "c5.2xlarge",
        "memory_gb": 16,
        "name": "mockenv3-org3peer1",
        "state": "stopped",
        "vcpu": 8
      },
      {
        "environment": "mockenv3",
        "region": "ca-central-1",
        "id": "b068b2b8bf9e",
        "instance_id": "i-0f8ab29b72c5a05a5",
        "instance_type": "c5.2xlarge",
        "memory_gb": 16,
        "name": "mockenv3-org6couchdb1",
        "state": "stopped",
        "vcpu": 8
      },
      {
        "environment": "mockenv3",
        "region": "ca-central-1",
        "id": "a28226832a30",
        "instance_id": "i-0b7077e7a39e2eec7",
        "instance_type": "t2.medium",
        "memory_gb": 4,
        "name": "mockenv3-org3app2",
        "state": "stopped",
        "vcpu": 2
      },
      {
        "environment": "mockenv3",
        "region": "ca-central-1",
        "id": "bb197abe31c7",
        "instance_id": "i-03bfff149eab04a22",
        "instance_type": "c5.2xlarge",
        "memory_gb": 16,
        "name": "mockenv3-org5peer2",
        "state": "stopped",
        "vcpu": 8
      },
      {
        "environment": "mockenv3",
        "region": "ca-central-1",
        "id": "0328768a5860",
        "instance_id": "i-086100c7d2aa948c3",
        "instance_type": "c5.2xlarge",
        "memory_gb": 16,
        "name": "mockenv3-org1couchdb1",
        "state": "stopped",
        "vcpu": 8
      },
      {
        "environment": "mockenv3",
        "region": "ca-central-1",
        "id": "f633018d3487",
        "instance_id": "i-0b3b9ba0054811960",
        "instance_type": "c5.2xlarge",
        "memory_gb": 16,
        "name": "mockenv3-org4peer3",
        "state": "stopped",
        "vcpu": 8
      },
      {
        "environment": "mockenv3",
        "region": "ca-central-1",
        "id": "815b6897229a",
        "instance_id": "i-0205dd65fafa04b52",
        "instance_type": "c5.2xlarge",
        "memory_gb": 16,
        "name": "mockenv3-org2peer1",
        "state": "stopped",
        "vcpu": 8
      },
      {
        "environment": "mockenv3",
        "region": "ca-central-1",
        "id": "3c4e53e14b32",
        "instance_id": "i-00cb37df46a3a1b9c",
        "instance_type": "c5.2xlarge",
        "memory_gb": 16,
        "name": "mockenv3-org6couchdb3",
        "state": "stopped",
        "vcpu": 8
      },
      {
        "environment": "mockenv3",
        "region": "ca-central-1",
        "id": "17df03e1de47",
        "instance_id": "i-04c26fe5fa5a461f6",
        "instance_type": "t2.micro",
        "memory_gb": 1,
        "name": "mockenv3-org1ca1",
        "state": "stopped",
        "vcpu": 1
      },
      {
        "environment": "mockenv3",
        "region": "ca-central-1",
        "id": "28cd0c4f2b2e",
        "instance_id": "i-0f457f643e5600eed",
        "instance_type": "t2.medium",
        "memory_gb": 4,
        "name": "mockenv3-org2app2",
        "state": "stopped",
        "vcpu": 2
      },
      {
        "environment": "mockenv3",
        "region": "ca-central-1",
        "id": "560c6974ba1e",
        "instance_id": "i-0dce8b6aef982df76",
        "instance_type": "c5.xlarge",
        "memory_gb": 8,
        "name": "mockenv3-org4admin",
        "state": "stopped",
        "vcpu": 4
      },
      {
        "environment": "mockenv3",
        "region": "ca-central-1",
        "id": "1f1cb8738028",
        "instance_id": "i-0d7e29155113217d0",
        "instance_type": "c5.2xlarge",
        "memory_gb": 16,
        "name": "mockenv3-org3peer2",
        "state": "stopped",
        "vcpu": 8
      },
      {
        "environment": "mockenv3",
        "region": "ca-central-1",
        "id": "f5e1952e958d",
        "instance_id": "i-0ee9e80cfc74a0898",
        "instance_type": "t2.micro",
        "memory_gb": 1,
        "name": "mockenv3-org2ca2",
        "state": "stopped",
        "vcpu": 1
      },
      {
        "environment": "mockenv3",
        "region": "ca-central-1",
        "id": "bd567538f943",
        "instance_id": "i-0f390537f4bb2c526",
        "instance_type": "t2.medium",
        "memory_gb": 4,
        "name": "mockenv3-org5app2",
        "state": "stopped",
        "vcpu": 2
      },
      {
        "environment": "mockenv3",
        "region": "ca-central-1",
        "id": "d4854833fb1c",
        "instance_id": "i-0f235911b547bfe34",
        "instance_type": "c5.2xlarge",
        "memory_gb": 16,
        "name": "mockenv3-org1couchdb3",
        "state": "stopped",
        "vcpu": 8
      },
      {
        "environment": "mockenv3",
        "region": "ca-central-1",
        "id": "c7f5a8ae775d",
        "instance_id": "i-0401cae45aeca3f7a",
        "instance_type": "t2.micro",
        "memory_gb": 1,
        "name": "mockenv3-org1ca2",
        "state": "stopped",
        "vcpu": 1
      },
      {
        "environment": "mockenv3",
        "region": "ca-central-1",
        "id": "aafdff6dee02",
        "instance_id": "i-0545fcc68d9aec05d",
        "instance_type": "c5.2xlarge",
        "memory_gb": 16,
        "name": "mockenv3-org5peer4",
        "state": "stopped",
        "vcpu": 8
      },
      {
        "environment": "mockenv3",
        "region": "ca-central-1",
        "id": "723036bc917b",
        "instance_id": "i-0f854d7467340384a",
        "instance_type": "c5.2xlarge",
        "memory_gb": 16,
        "name": "mockenv3-org5couchdb2",
        "state": "stopped",
        "vcpu": 8
      },
      {
        "environment": "mockenv3",
        "region": "ca-central-1",
        "id": "dcdd1e922406",
        "instance_id": "i-09947283d689501ed",
        "instance_type": "c5.xlarge",
        "memory_gb": 8,
        "name": "mockenv3-org5admin",
        "state": "stopped",
        "vcpu": 4
      },
      {
        "environment": "mockenv3",
        "region": "ca-central-1",
        "id": "895655593b82",
        "instance_id": "i-0a91741855fd8d21b",
        "instance_type": "t2.small",
        "memory_gb": 2,
        "name": "mockenv3-kafka2",
        "state": "stopped",
        "vcpu": 1
      },
      {
        "environment": "mockenv3",
        "region": "ca-central-1",
        "id": "f5c7edda05f6",
        "instance_id": "i-0ee85ded758bbb133",
        "instance_type": "t2.medium",
        "memory_gb": 4,
        "name": "mockenv3-org5app1",
        "state": "stopped",
        "vcpu": 2
      },
      {
        "environment": "mockenv3",
        "region": "ca-central-1",
        "id": "f49a75e7f5b5",
        "instance_id": "i-0d1203b346e6f3a8b",
        "instance_type": "t2.micro",
        "memory_gb": 1,
        "name": "mockenv3-org6ca1",
        "state": "stopped",
        "vcpu": 1
      },
      {
        "environment": "mockenv3",
        "region": "ca-central-1",
        "id": "8adc48958551",
        "instance_id": "i-0899f1bdba6104271",
        "instance_type": "c5.xlarge",
        "memory_gb": 8,
        "name": "mockenv3-org1admin",
        "state": "stopped",
        "vcpu": 4
      },
      {
        "environment": "mockenv3",
        "region": "ca-central-1",
        "id": "0ae28ec366b2",
        "instance_id": "i-04e6cd6c653b9703f",
        "instance_type": "c5.2xlarge",
        "memory_gb": 16,
        "name": "mockenv3-org4couchdb2",
        "state": "stopped",
        "vcpu": 8
      },
      {
        "environment": "mockenv3",
        "region": "ca-central-1",
        "id": "e4771369a825",
        "instance_id": "i-00717355a5b09d92d",
        "instance_type": "c5.2xlarge",
        "memory_gb": 16,
        "name": "mockenv3-org1peer4",
        "state": "stopped",
        "vcpu": 8
      },
      {
        "environment": "mockenv3",
        "region": "ca-central-1",
        "id": "f49c01b1d530",
        "instance_id": "i-08bab54792edb3108",
        "instance_type": "c5.2xlarge",
        "memory_gb": 16,
        "name": "mockenv3-org1peer1",
        "state": "stopped",
        "vcpu": 8
      },
      {
        "environment": "mockenv3",
        "region": "ca-central-1",
        "id": "ac11df53473c",
        "instance_id": "i-0e5b44c3c8d46d40a",
        "instance_type": "t2.micro",
        "memory_gb": 1,
        "name": "mockenv3-org4ca1",
        "state": "stopped",
        "vcpu": 1
      },
      {
        "environment": "mockenv3",
        "region": "ca-central-1",
        "id": "9eb0f154b0b4",
        "instance_id": "i-0528bfb7a1373cb75",
        "instance_type": "t2.micro",
        "memory_gb": 1,
        "name": "mockenv3-admin",
        "state": "stopped",
        "vcpu": 1
      },
      {
        "environment": "mockenv3",
        "region": "ca-central-1",
        "id": "11064d01c5c4",
        "instance_id": "i-0ee88bbe88e410e73",
        "instance_type": "t2.micro",
        "memory_gb": 1,
        "name": "mockenv3-org3ca2",
        "state": "stopped",
        "vcpu": 1
      },
      {
        "environment": "mockenv3",
        "region": "ca-central-1",
        "id": "59c097fa90f6",
        "instance_id": "i-00cf5cc42d686b39c",
        "instance_type": "c5.2xlarge",
        "memory_gb": 16,
        "name": "mockenv3-org1peer3",
        "state": "stopped",
        "vcpu": 8
      },
      {
        "environment": "mockenv3",
        "region": "ca-central-1",
        "id": "92c8e8e722f5",
        "instance_id": "i-018c532ea8ab6b5ef",
        "instance_type": "c5.2xlarge",
        "memory_gb": 16,
        "name": "mockenv3-org1couchdb2",
        "state": "stopped",
        "vcpu": 8
      },
      {
        "environment": "mockenv3",
        "region": "ca-central-1",
        "id": "ddf59daad7cc",
        "instance_id": "i-0cb53cad9d5c130a4",
        "instance_type": "t2.small",
        "memory_gb": 2,
        "name": "mockenv3-kafka3",
        "state": "stopped",
        "vcpu": 1
      },
      {
        "environment": "mockenv3",
        "region": "ca-central-1",
        "id": "a4212906b7a1",
        "instance_id": "i-0dc93c407b06e3829",
        "instance_type": "t2.small",
        "memory_gb": 2,
        "name": "mockenv3-zoo1",
        "state": "stopped",
        "vcpu": 1
      },
      {
        "environment": "mockenv3",
        "region": "ca-central-1",
        "id": "0abed0f705eb",
        "instance_id": "i-078797fa4054f26f8",
        "instance_type": "c5.2xlarge",
        "memory_gb": 16,
        "name": "mockenv3-org4peer4",
        "state": "stopped",
        "vcpu": 8
      },
      {
        "environment": "mockenv3",
        "region": "ca-central-1",
        "id": "b1b253674502",
        "instance_id": "i-02db01c3e4f8af91c",
        "instance_type": "t2.micro",
        "memory_gb": 1,
        "name": "mockenv3-org4ca2",
        "state": "stopped",
        "vcpu": 1
      },
      {
        "environment": "mockenv3",
        "region": "ca-central-1",
        "id": "f4aa815e906e",
        "instance_id": "i-03701407b8e15e994",
        "instance_type": "t2.small",
        "memory_gb": 2,
        "name": "mockenv3-kafka1",
        "state": "stopped",
        "vcpu": 1
      },
      {
        "environment": "mockenv3",
        "region": "ca-central-1",
        "id": "c03975e7cc7f",
        "instance_id": "i-0879eea460a169142",
        "instance_type": "c5.2xlarge",
        "memory_gb": 16,
        "name": "mockenv3-org5peer1",
        "state": "stopped",
        "vcpu": 8
      },
      {
        "environment": "mockenv3",
        "region": "ca-central-1",
        "id": "6cfbfe38e78b",
        "instance_id": "i-07a9303107eced81a",
        "instance_type": "c5.2xlarge",
        "memory_gb": 16,
        "name": "mockenv3-org6peer2",
        "state": "stopped",
        "vcpu": 8
      },
      {
        "environment": "mockenv3",
        "region": "ca-central-1",
        "id": "12e2d04a9d29",
        "instance_id": "i-09cc94659bb0d55d1",
        "instance_type": "t2.micro",
        "memory_gb": 1,
        "name": "mockenv3-org5ca2",
        "state": "stopped",
        "vcpu": 1
      },
      {
        "environment": "mockenv3",
        "region": "ca-central-1",
        "id": "82823f2e0e67",
        "instance_id": "i-0b200d1501d25b189",
        "instance_type": "t2.micro",
        "memory_gb": 1,
        "name": "mockenv3-orderer2",
        "state": "stopped",
        "vcpu": 1
      },
      {
        "environment": "mockenv3",
        "region": "ca-central-1",
        "id": "aaa9c42c58c8",
        "instance_id": "i-006e41966ec273d5c",
        "instance_type": "c5.xlarge",
        "memory_gb": 8,
        "name": "mockenv3-org6admin",
        "state": "stopped",
        "vcpu": 4
      },
      {
        "environment": "mockenv3",
        "region": "ca-central-1",
        "id": "2c30a2ab9cc7",
        "instance_id": "i-07fa42e247e7e4db1",
        "instance_type": "t2.micro",
        "memory_gb": 1,
        "name": "mockenv3-orderer1",
        "state": "stopped",
        "vcpu": 1
      },
      {
        "environment": "mockenv3",
        "region": "ca-central-1",
        "id": "8dc6df454d78",
        "instance_id": "i-02b321c4b35f6bc7b",
        "instance_type": "t2.micro",
        "memory_gb": 1,
        "name": "mockenv3-org2ca1",
        "state": "stopped",
        "vcpu": 1
      },
      {
        "environment": "mockenv3",
        "region": "ca-central-1",
        "id": "7d0c93d31827",
        "instance_id": "i-07249f14c67c1f620",
        "instance_type": "c5.2xlarge",
        "memory_gb": 16,
        "name": "mockenv3-org4peer1",
        "state": "stopped",
        "vcpu": 8
      },
      {
        "environment": "mockenv3",
        "region": "ca-central-1",
        "id": "bdb0987a9a94",
        "instance_id": "i-0356839e916342f5f",
        "instance_type": "c5.2xlarge",
        "memory_gb": 16,
        "name": "mockenv3-org2couchdb3",
        "state": "stopped",
        "vcpu": 8
      },
      {
        "environment": "mockenv3",
        "region": "ca-central-1",
        "id": "865111360bec",
        "instance_id": "i-0b582a96eea58a0a8",
        "instance_type": "c5.2xlarge",
        "memory_gb": 16,
        "name": "mockenv3-org4peer2",
        "state": "stopped",
        "vcpu": 8
      },
      {
        "environment": "mockenv3",
        "region": "ca-central-1",
        "id": "9382a5e06532",
        "instance_id": "i-0161984571e97ed45",
        "instance_type": "c5.2xlarge",
        "memory_gb": 16,
        "name": "mockenv3-org5peer3",
        "state": "stopped",
        "vcpu": 8
      },
      {
        "environment": "mockenv3",
        "region": "ca-central-1",
        "id": "7834625a89ab",
        "instance_id": "i-08fbe3e5549c08451",
        "instance_type": "c5.2xlarge",
        "memory_gb": 16,
        "name": "mockenv3-org2peer4",
        "state": "stopped",
        "vcpu": 8
      },
      {
        "environment": "mockenv3",
        "region": "ca-central-1",
        "id": "a13f8480a799",
        "instance_id": "i-089c5c96127b0edc4",
        "instance_type": "c5.xlarge",
        "memory_gb": 8,
        "name": "mockenv3-org3admin",
        "state": "stopped",
        "vcpu": 4
      },
      {
        "environment": "mockenv3",
        "region": "ca-central-1",
        "id": "2697ddc6dbd8",
        "instance_id": "i-0a0b3a0b84f8594d1",
        "instance_type": "c5.2xlarge",
        "memory_gb": 16,
        "name": "mockenv3-org4couchdb1",
        "state": "stopped",
        "vcpu": 8
      },
      {
        "environment": "mockenv3",
        "region": "ca-central-1",
        "id": "85cf08e8fe6d",
        "instance_id": "i-0c21480388f04fffa",
        "instance_type": "c5.2xlarge",
        "memory_gb": 16,
        "name": "mockenv3-org6couchdb2",
        "state": "stopped",
        "vcpu": 8
      },
      {
        "environment": "mockenv3",
        "region": "ca-central-1",
        "id": "6c196e92f40d",
        "instance_id": "i-020313c05c70c7661",
        "instance_type": "c5.2xlarge",
        "memory_gb": 16,
        "name": "mockenv3-org5couchdb3",
        "state": "stopped",
        "vcpu": 8
      },
      {
        "environment": "mockenv3",
        "region": "ca-central-1",
        "id": "5a29a8eebf5e",
        "instance_id": "i-0d7ea2bd3066f9d24",
        "instance_type": "t2.small",
        "memory_gb": 2,
        "name": "mockenv3-zoo2",
        "state": "stopped",
        "vcpu": 1
      },
      {
        "environment": "mockenv3",
        "region": "ca-central-1",
        "id": "6edb40e7b2c4",
        "instance_id": "i-0434f16f019623890",
        "instance_type": "c5.2xlarge",
        "memory_gb": 16,
        "name": "mockenv3-org2peer2",
        "state": "stopped",
        "vcpu": 8
      },
      {
        "environment": "mockenv3",
        "region": "ca-central-1",
        "id": "326bba19346b",
        "instance_id": "i-09e3660555eebe3d7",
        "instance_type": "t2.medium",
        "memory_gb": 4,
        "name": "mockenv3-org6app1",
        "state": "stopped",
        "vcpu": 2
      },
      {
        "environment": "mockenv3",
        "region": "ca-central-1",
        "id": "d1569ecb0dc3",
        "instance_id": "i-089f46ea175744c99",
        "instance_type": "c5.2xlarge",
        "memory_gb": 16,
        "name": "mockenv3-org2couchdb1",
        "state": "stopped",
        "vcpu": 8
      },
      {
        "environment": "mockenv3",
        "region": "ca-central-1",
        "id": "6c6e1862ad74",
        "instance_id": "i-05bc3fba82b4c798b",
        "instance_type": "t2.micro",
        "memory_gb": 1,
        "name": "mockenv3-org6ca2",
        "state": "stopped",
        "vcpu": 1
      },
      {
        "environment": "mockenv3",
        "region": "ca-central-1",
        "id": "720b1e1cc400",
        "instance_id": "i-0525d9ec317111755",
        "instance_type": "t2.medium",
        "memory_gb": 4,
        "name": "mockenv3-org4app2",
        "state": "stopped",
        "vcpu": 2
      },
      {
        "environment": "mockenv3",
        "region": "ca-central-1",
        "id": "056356a56f3e",
        "instance_id": "i-09dac5f53e4d100da",
        "instance_type": "t2.small",
        "memory_gb": 2,
        "name": "mockenv3-kafka4",
        "state": "stopped",
        "vcpu": 1
      },
      {
        "environment": "mockenv3",
        "region": "ca-central-1",
        "id": "a3564ab183bd",
        "instance_id": "i-0338c2fa2043749b3",
        "instance_type": "c5.2xlarge",
        "memory_gb": 16,
        "name": "mockenv3-org1peer2",
        "state": "stopped",
        "vcpu": 8
      },
      {
        "environment": "mockenv3",
        "region": "ca-central-1",
        "id": "5b8023ad596a",
        "instance_id": "i-0ca6aff28e4f09ea2",
        "instance_type": "t2.micro",
        "memory_gb": 1,
        "name": "mockenv3-org5ca1",
        "state": "stopped",
        "vcpu": 1
      },
      {
        "environment": "mockenv3",
        "region": "ca-central-1",
        "id": "22ae040cedc1",
        "instance_id": "i-030aff848e70dd91d",
        "instance_type": "c5.2xlarge",
        "memory_gb": 16,
        "name": "mockenv3-org2couchdb2",
        "state": "stopped",
        "vcpu": 8
      },
      {
        "environment": "mockenv3",
        "region": "ca-central-1",
        "id": "905bfe5ac7f2",
        "instance_id": "i-0ff59189bf3e18e97",
        "instance_type": "c5.2xlarge",
        "memory_gb": 16,
        "name": "mockenv3-org6peer4",
        "state": "stopped",
        "vcpu": 8
      },
      {
        "environment": "mockenv3",
        "region": "ca-central-1",
        "id": "c55c61c4b20f",
        "instance_id": "i-058b482b1646467da",
        "instance_type": "c5.2xlarge",
        "memory_gb": 16,
        "name": "mockenv3-org2peer3",
        "state": "stopped",
        "vcpu": 8
      },
      {
        "environment": "mockenv3",
        "region": "ca-central-1",
        "id": "91e3b6a8c345",
        "instance_id": "i-05c79685437b8a258",
        "instance_type": "c5.2xlarge",
        "memory_gb": 16,
        "name": "mockenv3-org5couchdb1",
        "state": "stopped",
        "vcpu": 8
      },
      {
        "environment": "mockenv3",
        "region": "ca-central-1",
        "id": "a16e816844bd",
        "instance_id": "i-053fda95ac228f8c5",
        "instance_type": "t2.micro",
        "memory_gb": 1,
        "name": "mockenv3-orderer3",
        "state": "stopped",
        "vcpu": 1
      },
      {
        "environment": "mockenv3",
        "region": "ca-central-1",
        "id": "86370dcefdae",
        "instance_id": "i-0c42520bdd6a7a534",
        "instance_type": "c5.2xlarge",
        "memory_gb": 16,
        "name": "mockenv3-org3couchdb3",
        "state": "stopped",
        "vcpu": 8
      },
      {
        "environment": "mockenv3",
        "region": "ca-central-1",
        "id": "b711cd8650ce",
        "instance_id": "i-00a1c56505c9b2ee4",
        "instance_type": "c5.2xlarge",
        "memory_gb": 16,
        "name": "mockenv3-org6peer1",
        "state": "stopped",
        "vcpu": 8
      },
      {
        "environment": "mockenv3",
        "region": "ca-central-1",
        "id": "281d8324d260",
        "instance_id": "i-07d1c73ff4c942a3c",
        "instance_type": "t2.medium",
        "memory_gb": 4,
        "name": "mockenv3-org6app2",
        "state": "stopped",
        "vcpu": 2
      },
      {
        "environment": "mockenv3",
        "region": "ca-central-1",
        "id": "a706d8c90d86",
        "instance_id": "i-00782b79caaa38ab5",
        "instance_type": "c5.2xlarge",
        "memory_gb": 16,
        "name": "mockenv3-org6peer3",
        "state": "stopped",
        "vcpu": 8
      },
      {
        "environment": "mockenv3",
        "region": "ca-central-1",
        "id": "dfc3263ab3b2",
        "instance_id": "i-086c974df6e4a532f",
        "instance_type": "t2.small",
        "memory_gb": 2,
        "name": "mockenv3-zoo3",
        "state": "stopped",
        "vcpu": 1
      },
      {
        "environment": "mockenv3",
        "region": "ca-central-1",
        "id": "a67ae6a3d3ab",
        "instance_id": "i-017549e8a39609389",
        "instance_type": "c5.2xlarge",
        "memory_gb": 16,
        "name": "mockenv3-org3couchdb2",
        "state": "stopped",
        "vcpu": 8
      },
      {
        "environment": "mockenv3",
        "region": "ca-central-1",
        "id": "2d32a214bbd3",
        "instance_id": "i-08b898af6356e76ac",
        "instance_type": "t2.medium",
        "memory_gb": 4,
        "name": "mockenv3-org3app1",
        "state": "stopped",
        "vcpu": 2
      },
      {
        "environment": "mockenv3",
        "region": "ca-central-1",
        "id": "10f573bb2806",
        "instance_id": "i-00f3211b37f283517",
        "instance_type": "t2.medium",
        "memory_gb": 4,
        "name": "mockenv3-org1app1",
        "state": "stopped",
        "vcpu": 2
      },
      {
        "environment": "mockenv3",
        "region": "ca-central-1",
        "id": "05cf92915f0b",
        "instance_id": "i-0ea0560a364ec6e1b",
        "instance_type": "c5.2xlarge",
        "memory_gb": 16,
        "name": "mockenv3-org4couchdb3",
        "state": "stopped",
        "vcpu": 8
      }
    ],
    "name": "mockenv3",
    "provider": "aws",
    "region": "ca-central-1",
    "running_instances": 0,
    "state": "stopped",
    "stopped_instances": 83,
    "total_instances": 83,
    "total_memory_gb": 798,
    "total_vcpu": 407
  },
  {
    "id": "ad5c448d114c",
    "instances": [
      {
        "environment": "mockenv4",
        "region": "ca-central-1",
        "id": "719d499ff57c",
        "instance_id": "i-071c75646d226dc73",
        "instance_type": "t3.small",
        "memory_gb": 2,
        "name": "mockenv4-org4app1",
        "state": "stopped",
        "vcpu": 2
      },
      {
        "environment": "mockenv4",
        "region": "ca-central-1",
        "id": "44058603bcbb",
        "instance_id": "i-01d86244293b4d12c",
        "instance_type": "t2.small",
        "memory_gb": 2,
        "name": "mockenv4-kafka3",
        "state": "stopped",
        "vcpu": 1
      },
      {
        "environment": "mockenv4",
        "region": "ca-central-1",
        "id": "b80300a8d5bd",
        "instance_id": "i-0c6ed3dc150f3e568",
        "instance_type": "t2.micro",
        "memory_gb": 1,
        "name": "mockenv4-org4ca1",
        "state": "stopped",
        "vcpu": 1
      },
      {
        "environment": "mockenv4",
        "region": "ca-central-1",
        "id": "a2149768473d",
        "instance_id": "i-02943039d40e0fbfd",
        "instance_type": "t3.small",
        "memory_gb": 2,
        "name": "mockenv4-org5app1",
        "state": "stopped",
        "vcpu": 2
      },
      {
        "environment": "mockenv4",
        "region": "ca-central-1",
        "id": "bec2ce3a8050",
        "instance_id": "i-0d368037e5b6709db",
        "instance_type": "t3.small",
        "memory_gb": 2,
        "name": "mockenv4-org6app2",
        "state": "stopped",
        "vcpu": 2
      },
      {
        "environment": "mockenv4",
        "region": "ca-central-1",
        "id": "d465a5f6a736",
        "instance_id": "i-07bf50a4c01ed1313",
        "instance_type": "c5.2xlarge",
        "memory_gb": 16,
        "name": "mockenv4-org3peer1",
        "state": "stopped",
        "vcpu": 8
      },
      {
        "environment": "mockenv4",
        "region": "ca-central-1",
        "id": "3dca77fd504b",
        "instance_id": "i-01c55ddb1637e2f07",
        "instance_type": "t2.micro",
        "memory_gb": 1,
        "name": "mockenv4-org2ca2",
        "state": "stopped",
        "vcpu": 1
      },
      {
        "environment": "mockenv4",
        "region": "ca-central-1",
        "id": "26fdaf50904e",
        "instance_id": "i-09e2c3809819b2570",
        "instance_type": "t3.small",
        "memory_gb": 2,
        "name": "mockenv4-org2app2",
        "state": "stopped",
        "vcpu": 2
      },
      {
        "environment": "mockenv4",
        "region": "ca-central-1",
        "id": "3fa46dda2457",
        "instance_id": "i-00d36b6dc2f22faad",
        "instance_type": "t2.micro",
        "memory_gb": 1,
        "name": "mockenv4-org6ca2",
        "state": "stopped",
        "vcpu": 1
      },
      {
        "environment": "mockenv4",
        "region": "ca-central-1",
        "id": "d4d11910b653",
        "instance_id": "i-05284abd159ef0390",
        "instance_type": "t3.medium",
        "memory_gb": 4,
        "name": "mockenv4-org4admin",
        "state": "stopped",
        "vcpu": 2
      },
      {
        "environment": "mockenv4",
        "region": "ca-central-1",
        "id": "14143792ac53",
        "instance_id": "i-01839ccd34342d68c",
        "instance_type": "t3.small",
        "memory_gb": 2,
        "name": "mockenv4-org5app2",
        "state": "stopped",
        "vcpu": 2
      },
      {
        "environment": "mockenv4",
        "region": "ca-central-1",
        "id": "4e7cf8a8dac8",
        "instance_id": "i-02eb42e1e948cfe55",
        "instance_type": "t3.medium",
        "memory_gb": 4,
        "name": "mockenv4-org5admin",
        "state": "stopped",
        "vcpu": 2
      },
      {
        "environment": "mockenv4",
        "region": "ca-central-1",
        "id": "1e141f4505fa",
        "instance_id": "i-076937378cb7f9d39",
        "instance_type": "c5.2xlarge",
        "memory_gb": 16,
        "name": "mockenv4-org1peer1",
        "state": "stopped",
        "vcpu": 8
      },
      {
        "environment": "mockenv4",
        "region": "ca-central-1",
        "id": "485765191c69",
        "instance_id": "i-00752d1a65895a363",
        "instance_type": "t2.small",
        "memory_gb": 2,
        "name": "mockenv4-zoo2",
        "state": "stopped",
        "vcpu": 1
      },
      {
        "environment": "mockenv4",
        "region": "ca-central-1",
        "id": "42b98ef6cc51",
        "instance_id": "i-002098f454b94a739",
        "instance_type": "c5.2xlarge",
        "memory_gb": 16,
        "name": "mockenv4-org6peer2",
        "state": "stopped",
        "vcpu": 8
      },
      {
        "environment": "mockenv4",
        "region": "ca-central-1",
        "id": "53ddae7614a5",
        "instance_id": "i-0ad04d9e1427db1f9",
        "instance_type": "c5.2xlarge",
        "memory_gb": 16,
        "name": "mockenv4-org5peer2",
        "state": "stopped",
        "vcpu": 8
      },
      {
        "environment": "mockenv4",
        "region": "ca-central-1",
        "id": "50f195e139f2",
        "instance_id": "i-025a223d40dea49fa",
        "instance_type": "t3.small",
        "memory_gb": 2,
        "name": "mockenv4-org3app1",
        "state": "stopped",
        "vcpu": 2
      },
      {
        "environment": "mockenv4",
        "region": "ca-central-1",
        "id": "24060f508242",
        "instance_id": "i-0bb5e20947cd0e861",
        "instance_type": "t2.micro",
        "memory_gb": 1,
        "name": "mockenv4-org1ca1",
        "state": "stopped",
        "vcpu": 1
      },
      {
        "environment": "mockenv4",
        "region": "ca-central-1",
        "id": "1d9d8ccea28d",
        "instance_id": "i-030ad523576890bdb",
        "instance_type": "t2.micro",
        "memory_gb": 1,
        "name": "mockenv4-admin",
        "state": "stopped",
        "vcpu": 1
      },
      {
        "environment": "mockenv4",
        "region": "ca-central-1",
        "id": "02fe4798f9c1",
        "instance_id": "i-01d50deae235d4715",
        "instance_type": "t2.micro",
        "memory_gb": 1,
        "name": "mockenv4-org3ca2",
        "state": "stopped",
        "vcpu": 1
      },
      {
        "environment": "mockenv4",
        "region": "ca-central-1",
        "id": "3390c63d7662",
        "instance_id": "i-043ddbc8cfbb8041f",
        "instance_type": "t2.small",
        "memory_gb": 2,
        "name": "mockenv4-kafka1",
        "state": "stopped",
        "vcpu": 1
      },
      {
        "environment": "mockenv4",
        "region": "ca-central-1",
        "id": "64931b087c7e",
        "instance_id": "i-05f8d0d8bb63b1a6b",
        "instance_type": "t2.micro",
        "memory_gb": 1,
        "name": "mockenv4-org1ca2",
        "state": "stopped",
        "vcpu": 1
      },
      {
        "environment": "mockenv4",
        "region": "ca-central-1",
        "id": "329040e5ec7b",
        "instance_id": "i-0a30d4808fccbdcb8",
        "instance_type": "t3.small",
        "memory_gb": 2,
        "name": "mockenv4-org1app2",
        "state": "stopped",
        "vcpu": 2
      },
      {
        "environment": "mockenv4",
        "region": "ca-central-1",
        "id": "93c942ca8169",
        "instance_id": "i-0e71ee77d15693b4b",
        "instance_type": "t2.small",
        "memory_gb": 2,
        "name": "mockenv4-zoo1",
        "state": "stopped",
        "vcpu": 1
      },
      {
        "environment": "mockenv4",
        "region": "ca-central-1",
        "id": "b1b473dd3d5d",
        "instance_id": "i-03a14738e9f555da9",
        "instance_type": "t2.small",
        "memory_gb": 2,
        "name": "mockenv4-zoo3",
        "state": "stopped",
        "vcpu": 1
      },
      {
        "environment": "mockenv4",
        "region": "ca-central-1",
        "id": "fc00168e1eed",
        "instance_id": "i-0351e24fd1298b777",
        "instance_type": "t2.micro",
        "memory_gb": 1,
        "name": "mockenv4-org4ca2",
        "state": "stopped",
        "vcpu": 1
      },
      {
        "environment": "mockenv4",
        "region": "ca-central-1",
        "id": "17f1f300ef91",
        "instance_id": "i-0d87b7d24b4b966c2",
        "instance_type": "t3.medium",
        "memory_gb": 4,
        "name": "mockenv4-org1admin",
        "state": "stopped",
        "vcpu": 2
      },
      {
        "environment": "mockenv4",
        "region": "ca-central-1",
        "id": "fd1557ad91ca",
        "instance_id": "i-0f12d013aeffeb095",
        "instance_type": "t3.small",
        "memory_gb": 2,
        "name": "mockenv4-org6app1",
        "state": "stopped",
        "vcpu": 2
      },
      {
        "environment": "mockenv4",
        "region": "ca-central-1",
        "id": "35c9a04ae78c",
        "instance_id": "i-0343cc018585ecddf",
        "instance_type": "t3.small",
        "memory_gb": 2,
        "name": "mockenv4-org1app1",
        "state": "stopped",
        "vcpu": 2
      },
      {
        "environment": "mockenv4",
        "region": "ca-central-1",
        "id": "b19a56a8015e",
        "instance_id": "i-058f9365e36ff1dab",
        "instance_type": "t3.medium",
        "memory_gb": 4,
        "name": "mockenv4-org2admin",
        "state": "stopped",
        "vcpu": 2
      },
      {
        "environment": "mockenv4",
        "region": "ca-central-1",
        "id": "a898a9734e99",
        "instance_id": "i-04ab76023830d8087",
        "instance_type": "t2.small",
        "memory_gb": 2,
        "name": "mockenv4-kafka2",
        "state": "stopped",
        "vcpu": 1
      },
      {
        "environment": "mockenv4",
        "region": "ca-central-1",
        "id": "4839843afd04",
        "instance_id": "i-0f95a04fab9d9c2e4",
        "instance_type": "t3.small",
        "memory_gb": 2,
        "name": "mockenv4-org3app2",
        "state": "stopped",
        "vcpu": 2
      },
      {
        "environment": "mockenv4",
        "region": "ca-central-1",
        "id": "9d67afd75936",
        "instance_id": "i-0ba3d969566a38577",
        "instance_type": "c5.2xlarge",
        "memory_gb": 16,
        "name": "mockenv4-org3peer2",
        "state": "stopped",
        "vcpu": 8
      },
      {
        "environment": "mockenv4",
        "region": "ca-central-1",
        "id": "5df3c122f86a",
        "instance_id": "i-039d594905b6da0ff",
        "instance_type": "t2.micro",
        "memory_gb": 1,
        "name": "mockenv4-org3ca1",
        "state": "stopped",
        "vcpu": 1
      },
      {
        "environment": "mockenv4",
        "region": "ca-central-1",
        "id": "a290b2852595",
        "instance_id": "i-0a81b04963eea5416",
        "instance_type": "c5.2xlarge",
        "memory_gb": 16,
        "name": "mockenv4-org6peer1",
        "state": "stopped",
        "vcpu": 8
      },
      {
        "environment": "mockenv4",
        "region": "ca-central-1",
        "id": "e1fb91d5d19b",
        "instance_id": "i-05dba6198a6f8af29",
        "instance_type": "t2.micro",
        "memory_gb": 1,
        "name": "mockenv4-orderer2",
        "state": "stopped",
        "vcpu": 1
      },
      {
        "environment": "mockenv4",
        "region": "ca-central-1",
        "id": "36153956507b",
        "instance_id": "i-0a3f075db62dc422c",
        "instance_type": "t2.micro",
        "memory_gb": 1,
        "name": "mockenv4-org6ca1",
        "state": "stopped",
        "vcpu": 1
      },
      {
        "environment": "mockenv4",
        "region": "ca-central-1",
        "id": "89095c491388",
        "instance_id": "i-08b622bd87725f7fe",
        "instance_type": "c5.2xlarge",
        "memory_gb": 16,
        "name": "mockenv4-org1peer2",
        "state": "stopped",
        "vcpu": 8
      },
      {
        "environment": "mockenv4",
        "region": "ca-central-1",
        "id": "ce2d171929f5",
        "instance_id": "i-0eee987e3c043c618",
        "instance_type": "t3.medium",
        "memory_gb": 4,
        "name": "mockenv4-org3admin",
        "state": "stopped",
        "vcpu": 2
      },
      {
        "environment": "mockenv4",
        "region": "ca-central-1",
        "id": "0591a0a4267e",
        "instance_id": "i-0535d0aab0edb9990",
        "instance_type": "t2.micro",
        "memory_gb": 1,
        "name": "mockenv4-org5ca1",
        "state": "stopped",
        "vcpu": 1
      },
      {
        "environment": "mockenv4",
        "region": "ca-central-1",
        "id": "8be970c14ac2",
        "instance_id": "i-07bf6fe55d4a3f552",
        "instance_type": "c5.2xlarge",
        "memory_gb": 16,
        "name": "mockenv4-org4peer2",
        "state": "stopped",
        "vcpu": 8
      },
      {
        "environment": "mockenv4",
        "region": "ca-central-1",
        "id": "2c1ca946b212",
        "instance_id": "i-016d42a175975ffa4",
        "instance_type": "t3.small",
        "memory_gb": 2,
        "name": "mockenv4-org4app2",
        "state": "stopped",
        "vcpu": 2
      },
      {
        "environment": "mockenv4",
        "region": "ca-central-1",
        "id": "e10f7f3af226",
        "instance_id": "i-0d24f4a1432038057",
        "instance_type": "c5.2xlarge",
        "memory_gb": 16,
        "name": "mockenv4-org4peer1",
        "state": "stopped",
        "vcpu": 8
      },
      {
        "environment": "mockenv4",
        "region": "ca-central-1",
        "id": "2b3dc834ca2f",
        "instance_id": "i-0192e936f569da9e2",
        "instance_type": "c5.2xlarge",
        "memory_gb": 16,
        "name": "mockenv4-org2peer2",
        "state": "stopped",
        "vcpu": 8
      },
      {
        "environment": "mockenv4",
        "region": "ca-central-1",
        "id": "08e6b40a9477",
        "instance_id": "i-007e3984a32f6d19b",
        "instance_type": "c5.2xlarge",
        "memory_gb": 16,
        "name": "mockenv4-org2peer1",
        "state": "stopped",
        "vcpu": 8
      },
      {
        "environment": "mockenv4",
        "region": "ca-central-1",
        "id": "aced52b96272",
        "instance_id": "i-0f1785ae32fd8f301",
        "instance_type": "t2.small",
        "memory_gb": 2,
        "name": "mockenv4-kafka4",
        "state": "stopped",
        "vcpu": 1
      },
      {
        "environment": "mockenv4",
        "region": "ca-central-1",
        "id": "b25ec0a16445",
        "instance_id": "i-00743cb11c40f7121",
        "instance_type": "t3.small",
        "memory_gb": 2,
        "name": "mockenv4-org2app1",
        "state": "stopped",
        "vcpu": 2
      },
      {
        "environment": "mockenv4",
        "region": "ca-central-1",
        "id": "0de339c06537",
        "instance_id": "i-08acff738b6eefbb3",
        "instance_type": "c5.2xlarge",
        "memory_gb": 16,
        "name": "mockenv4-org5peer1",
        "state": "stopped",
        "vcpu": 8
      },
      {
        "environment": "mockenv4",
        "region": "ca-central-1",
        "id": "3a5f01cfed89",
        "instance_id": "i-0bc7b098b29c79754",
        "instance_type": "t2.micro",
        "memory_gb": 1,
        "name": "mockenv4-org5ca2",
        "state": "stopped",
        "vcpu": 1
      },
      {
        "environment": "mockenv4",
        "region": "ca-central-1",
        "id": "75fbc908f595",
        "instance_id": "i-0f3d2ad78d2a08c37",
        "instance_type": "t2.micro",
        "memory_gb": 1,
        "name": "mockenv4-orderer1",
        "state": "stopped",
        "vcpu": 1
      },
      {
        "environment": "mockenv4",
        "region": "ca-central-1",
        "id": "f2457833f058",
        "instance_id": "i-0ddd509fd7f0d7998",
        "instance_type": "t3.medium",
        "memory_gb": 4,
        "name": "mockenv4-org6admin",
        "state": "stopped",
        "vcpu": 2
      },
      {
        "environment": "mockenv4",
        "region": "ca-central-1",
        "id": "5e59ce4cbb11",
        "instance_id": "i-026a0cdf8433d74f3",
        "instance_type": "t2.micro",
        "memory_gb": 1,
        "name": "mockenv4-orderer3",
        "state": "stopped",
        "vcpu": 1
      },
      {
        "environment": "mockenv4",
        "region": "ca-central-1",
        "id": "6f62a3dd3cce",
        "instance_id": "i-09e1c63fdeb70a0ea",
        "instance_type": "t2.micro",
        "memory_gb": 1,
        "name": "mockenv4-org2ca1",
        "state": "stopped",
        "vcpu": 1
      }
    ],
    "name": "mockenv4",
    "provider": "aws",
    "region": "ca-central-1",
    "running_instances": 0,
    "state": "stopped",
    "stopped_instances": 53,
    "total_instances": 53,
    "total_memory_gb": 270,
    "total_vcpu": 155
  },
  {
    "id": "3c8c9bb51bc8",
    "instances": [
      {
        "environment": "mockenv5",
        "region": "ca-central-1",
        "id": "ea92070884db",
        "instance_id": "i-044a7742565cd9040",
        "instance_type": "t2.medium",
        "memory_gb": 4,
        "name": "mockenv5-org1peer2",
        "state": "stopped",
        "vcpu": 2
      },
      {
        "environment": "mockenv5",
        "region": "ca-central-1",
        "id": "55aee1db6ff3",
        "instance_id": "i-08b4337605830c376",
        "instance_type": "t2.medium",
        "memory_gb": 4,
        "name": "mockenv5-org4peer1",
        "state": "stopped",
        "vcpu": 2
      },
      {
        "environment": "mockenv5",
        "region": "ca-central-1",
        "id": "7ec4ac55ee54",
        "instance_id": "i-074ffb151220acb37",
        "instance_type": "t2.medium",
        "memory_gb": 4,
        "name": "mockenv5-org4peer2",
        "state": "stopped",
        "vcpu": 2
      },
      {
        "environment": "mockenv5",
        "region": "ca-central-1",
        "id": "149c8b99c551",
        "instance_id": "i-09c3ae7f709d026f9",
        "instance_type": "t2.medium",
        "memory_gb": 4,
        "name": "mockenv5-kafka2",
        "state": "stopped",
        "vcpu": 2
      },
      {
        "environment": "mockenv5",
        "region": "ca-central-1",
        "id": "13bfa7760af5",
        "instance_id": "i-05e41007a825cf74f",
        "instance_type": "t2.medium",
        "memory_gb": 4,
        "name": "mockenv5-org1peer1",
        "state": "stopped",
        "vcpu": 2
      },
      {
        "environment": "mockenv5",
        "region": "ca-central-1",
        "id": "3708c0881c91",
        "instance_id": "i-0fdc8b231f1335e88",
        "instance_type": "t2.medium",
        "memory_gb": 4,
        "name": "mockenv5-org2peer2",
        "state": "stopped",
        "vcpu": 2
      },
      {
        "environment": "mockenv5",
        "region": "ca-central-1",
        "id": "4f280dca90ba",
        "instance_id": "i-03fd31a538a26cd4c",
        "instance_type": "t2.medium",
        "memory_gb": 4,
        "name": "mockenv5-org3peer1",
        "state": "stopped",
        "vcpu": 2
      },
      {
        "environment": "mockenv5",
        "region": "ca-central-1",
        "id": "3130c6f1d59d",
        "instance_id": "i-0335492d5b88f5ba0",
        "instance_type": "t2.micro",
        "memory_gb": 1,
        "name": "mockenv5-orderer1",
        "state": "stopped",
        "vcpu": 1
      },
      {
        "environment": "mockenv5",
        "region": "ca-central-1",
        "id": "d28218f8939d",
        "instance_id": "i-09c51ed034d69d6f0",
        "instance_type": "t2.small",
        "memory_gb": 2,
        "name": "mockenv5-zoo1",
        "state": "stopped",
        "vcpu": 1
      },
      {
        "environment": "mockenv5",
        "region": "ca-central-1",
        "id": "0088b60add67",
        "instance_id": "i-08041d386acfd0b2f",
        "instance_type": "t2.medium",
        "memory_gb": 4,
        "name": "mockenv5-org3peer2",
        "state": "stopped",
        "vcpu": 2
      },
      {
        "environment": "mockenv5",
        "region": "ca-central-1",
        "id": "cac29dcec57a",
        "instance_id": "i-055df7f5477314394",
        "instance_type": "t2.small",
        "memory_gb": 2,
        "name": "mockenv5-zoo2",
        "state": "stopped",
        "vcpu": 1
      },
      {
        "environment": "mockenv5",
        "region": "ca-central-1",
        "id": "66f8c8fb5f8e",
        "instance_id": "i-07e831d45c42fd193",
        "instance_type": "t2.medium",
        "memory_gb": 4,
        "name": "mockenv5-kafka1",
        "state": "stopped",
        "vcpu": 2
      },
      {
        "environment": "mockenv5",
        "region": "ca-central-1",
        "id": "b1f6eb042314",
        "instance_id": "i-0b075aeb37bc972a1",
        "instance_type": "t2.micro",
        "memory_gb": 1,
        "name": "mockenv5-admin",
        "state": "stopped",
        "vcpu": 1
      },
      {
        "environment": "mockenv5",
        "region": "ca-central-1",
        "id": "038cfbf9f69c",
        "instance_id": "i-05d597db2a6ba54e6",
        "instance_type": "t2.medium",
        "memory_gb": 4,
        "name": "mockenv5-org2peer1",
        "state": "stopped",
        "vcpu": 2
      },
      {
        "environment": "mockenv5",
        "region": "ca-central-1",
        "id": "7263b8184f63",
        "instance_id": "i-07a108898968f0007",
        "instance_type": "t2.micro",
        "memory_gb": 1,
        "name": "mockenv5-orderer2",
        "state": "stopped",
        "vcpu": 1
      }
    ],
    "name": "mockenv5",
    "provider": "aws",
    "region": "ca-central-1",
    "running_instances": 0,
    "state": "stopped",
    "stopped_instances": 15,
    "total_instances": 15,
    "total_memory_gb": 47,
    "total_vcpu": 25
  },
  {
    "id": "7aacc4c405b9",
    "instances": [
      {
        "environment": "mockenv6",
        "region": "ca-central-1",
        "id": "1aef6299109b",
        "instance_id": "i-0008ad1bfd83a52eb",
        "instance_type": "t3.xlarge",
        "memory_gb": 16,
        "name": "mockenv6-k8node2",
        "state": "stopped",
        "vcpu": 4
      },
      {
        "environment": "mockenv6",
        "region": "ca-central-1",
        "id": "9b97d53ab004",
        "instance_id": "i-04b69b4ec92d548df",
        "instance_type": "t3.medium",
        "memory_gb": 4,
        "name": "mockenv6-k8master1",
        "state": "stopped",
        "vcpu": 2
      },
      {
        "environment": "mockenv6",
        "region": "ca-central-1",
        "id": "be8225018847",
        "instance_id": "i-0dedc66553eaecfc3",
        "instance_type": "t3.medium",
        "memory_gb": 4,
        "name": "mockenv6-k8master2",
        "state": "stopped",
        "vcpu": 2
      },
      {
        "environment": "mockenv6",
        "region": "ca-central-1",
        "id": "dd4c000552b2",
        "instance_id": "i-0778ddf78ec72bffd",
        "instance_type": "t3.small",
        "memory_gb": 2,
        "name": "mockenv6-etcd3",
        "state": "stopped",
        "vcpu": 2
      },
      {
        "environment": "mockenv6",
        "region": "ca-central-1",
        "id": "63296c73e207",
        "instance_id": "i-03b3782294e848647",
        "instance_type": "t3.small",
        "memory_gb": 2,
        "name": "mockenv6-etcd1",
        "state": "stopped",
        "vcpu": 2
      },
      {
        "environment": "mockenv6",
        "region": "ca-central-1",
        "id": "a2f4635a834b",
        "instance_id": "i-0872ea25c0766d383",
        "instance_type": "t3.xlarge",
        "memory_gb": 16,
        "name": "mockenv6-k8node5",
        "state": "stopped",
        "vcpu": 4
      },
      {
        "environment": "mockenv6",
        "region": "ca-central-1",
        "id": "98a9867235f3",
        "instance_id": "i-0e9fcaf646ce779ab",
        "instance_type": "t3.xlarge",
        "memory_gb": 16,
        "name": "mockenv6-k8node3",
        "state": "stopped",
        "vcpu": 4
      },
      {
        "environment": "mockenv6",
        "region": "ca-central-1",
        "id": "906d663b6ecd",
        "instance_id": "i-0eba74077ac760573",
        "instance_type": "t3.small",
        "memory_gb": 2,
        "name": "mockenv6-etcd2",
        "state": "stopped",
        "vcpu": 2
      },
      {
        "environment": "mockenv6",
        "region": "ca-central-1",
        "id": "8a0a2a0d8725",
        "instance_id": "i-0d67730effbee8b3b",
        "instance_type": "t3.xlarge",
        "memory_gb": 16,
        "name": "mockenv6-k8node1",
        "state": "stopped",
        "vcpu": 4
      },
      {
        "environment": "mockenv6",
        "region": "ca-central-1",
        "id": "8c65980f5b5f",
        "instance_id": "i-009517923633b51f4",
        "instance_type": "t3.xlarge",
        "memory_gb": 16,
        "name": "mockenv6-k8node4",
        "state": "stopped",
        "vcpu": 4
      }
    ],
    "name": "mockenv6",
    "provider": "aws",
    "region": "ca-central-1",
    "running_instances": 0,
    "state": "stopped",
    "stopped_instances": 10,
    "total_instances": 10,
    "total_memory_gb": 94,
    "total_vcpu": 30
  },
  {
    "id": "4f9f1afb29f1",
    "instances": [
      {
        "environment": "mockenv7",
        "region": "ca-central-1",
        "id": "26f653b0d42a",
        "instance_id": "i-01901f7ff1afc2365",
        "instance_type": "t2.medium",
        "memory_gb": 4,
        "name": "mockenv7-org1-app",
        "state": "stopped",
        "vcpu": 2
      },
      {
        "environment": "mockenv7",
        "region": "ca-central-1",
        "id": "6397591647fc",
        "instance_id": "i-03f161f3b27709891",
        "instance_type": "t2.medium",
        "memory_gb": 4,
        "name": "mockenv7-org5-app",
        "state": "stopped",
        "vcpu": 2
      },
      {
        "environment": "mockenv7",
        "region": "ca-central-1",
        "id": "b7f7bee3c531",
        "instance_id": "i-0b1d0fcde41d75806",
        "instance_type": "t2.medium",
        "memory_gb": 4,
        "name": "mockenv7-org2-app",
        "state": "stopped",
        "vcpu": 2
      },
      {
        "environment": "mockenv7",
        "region": "ca-central-1",
        "id": "6c7afa9ae90e",
        "instance_id": "i-0517e563d41f75d18",
        "instance_type": "t2.medium",
        "memory_gb": 4,
        "name": "mockenv7-org4-app",
        "state": "stopped",
        "vcpu": 2
      },
      {
        "environment": "mockenv7",
        "region": "ca-central-1",
        "id": "30fc2b2d1f89",
        "instance_id": "i-0c4f22923d9f58d0a",
        "instance_type": "t2.medium",
        "memory_gb": 4,
        "name": "mockenv7-org3-app",
        "state": "stopped",
        "vcpu": 2
      },
      {
        "environment": "mockenv7",
        "region": "ca-central-1",
        "id": "86670aa353c9",
        "instance_id": "i-048b24df4aa79e266",
        "instance_type": "t2.medium",
        "memory_gb": 4,
        "name": "mockenv7-org6-app",
        "state": "stopped",
        "vcpu": 2
      }
    ],
    "name": "mockenv7",
    "provider": "aws",
    "region": "ca-central-1",
    "running_instances": 0,
    "state": "stopped",
    "stopped_instances": 6,
    "total_instances": 6,
    "total_memory_gb": 24,
    "total_vcpu": 12
  }
]
`
