// Prebuilt opentelekomcloud Provider for Terraform CDK (cdktf)
package opentelekomcloud


type RdsInstanceV3Volume struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opentelekomcloud/r/rds_instance_v3#size RdsInstanceV3#size}.
	Size *float64 `field:"required" json:"size" yaml:"size"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opentelekomcloud/r/rds_instance_v3#type RdsInstanceV3#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opentelekomcloud/r/rds_instance_v3#disk_encryption_id RdsInstanceV3#disk_encryption_id}.
	DiskEncryptionId *string `field:"optional" json:"diskEncryptionId" yaml:"diskEncryptionId"`
}

