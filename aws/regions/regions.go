package regions

import ()

const (
	GovCloud       = "us-gov-west-1"
	US_EAST_1      = "us-east-1"
	US_WEST_1      = "us-west-1"
	US_WEST_2      = "us-west-2"
	EU_WEST_1      = "eu-west-1"
	AP_SOUTHEAST_1 = "ap-southeast-1"
	AP_SOUTHEAST_2 = "ap-southeast-2"
	AP_NORTHEAST_1 = "ap-northeast-1"
	SA_EAST_1      = "sa-east-1"
	CN_NORTH_1     = "cn-north-1"
	DEFAULT_REGION = US_EAST_1
)

type Region struct {
	name string
}

func (r Region) Name() string {
	return r.name
}

func Config(regionName string) *Region {
	return &Region{regionName}
}
