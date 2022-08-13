// Prebuilt opentelekomcloud Provider for Terraform CDK (cdktf)
package opentelekomcloud


type RdsInstanceV1Backupstrategy struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opentelekomcloud/r/rds_instance_v1#keepdays RdsInstanceV1#keepdays}.
	Keepdays *float64 `field:"optional" json:"keepdays" yaml:"keepdays"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opentelekomcloud/r/rds_instance_v1#starttime RdsInstanceV1#starttime}.
	Starttime *string `field:"optional" json:"starttime" yaml:"starttime"`
}

