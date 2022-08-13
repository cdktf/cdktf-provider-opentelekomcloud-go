// Prebuilt opentelekomcloud Provider for Terraform CDK (cdktf)
package opentelekomcloud


type RdsInstanceV3BackupStrategy struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opentelekomcloud/r/rds_instance_v3#start_time RdsInstanceV3#start_time}.
	StartTime *string `field:"required" json:"startTime" yaml:"startTime"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opentelekomcloud/r/rds_instance_v3#keep_days RdsInstanceV3#keep_days}.
	KeepDays *float64 `field:"optional" json:"keepDays" yaml:"keepDays"`
}

