package rdsinstancev1


type RdsInstanceV1Ha struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opentelekomcloud/r/rds_instance_v1#enable RdsInstanceV1#enable}.
	Enable interface{} `field:"optional" json:"enable" yaml:"enable"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opentelekomcloud/r/rds_instance_v1#replicationmode RdsInstanceV1#replicationmode}.
	Replicationmode *string `field:"optional" json:"replicationmode" yaml:"replicationmode"`
}

