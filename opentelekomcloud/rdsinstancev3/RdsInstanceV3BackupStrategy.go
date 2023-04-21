package rdsinstancev3


type RdsInstanceV3BackupStrategy struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.34.2/docs/resources/rds_instance_v3#start_time RdsInstanceV3#start_time}.
	StartTime *string `field:"required" json:"startTime" yaml:"startTime"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.34.2/docs/resources/rds_instance_v3#keep_days RdsInstanceV3#keep_days}.
	KeepDays *float64 `field:"optional" json:"keepDays" yaml:"keepDays"`
}

