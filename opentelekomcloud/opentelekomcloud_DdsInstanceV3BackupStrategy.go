// Prebuilt opentelekomcloud Provider for Terraform CDK (cdktf)
package opentelekomcloud


type DdsInstanceV3BackupStrategy struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opentelekomcloud/r/dds_instance_v3#keep_days DdsInstanceV3#keep_days}.
	KeepDays *float64 `field:"required" json:"keepDays" yaml:"keepDays"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opentelekomcloud/r/dds_instance_v3#start_time DdsInstanceV3#start_time}.
	StartTime *string `field:"required" json:"startTime" yaml:"startTime"`
}
