// Prebuilt opentelekomcloud Provider for Terraform CDK (cdktf)
package opentelekomcloud


type CssSnapshotConfigurationV1Configuration struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opentelekomcloud/r/css_snapshot_configuration_v1#agency CssSnapshotConfigurationV1#agency}.
	Agency *string `field:"required" json:"agency" yaml:"agency"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opentelekomcloud/r/css_snapshot_configuration_v1#bucket CssSnapshotConfigurationV1#bucket}.
	Bucket *string `field:"required" json:"bucket" yaml:"bucket"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opentelekomcloud/r/css_snapshot_configuration_v1#kms_id CssSnapshotConfigurationV1#kms_id}.
	KmsId *string `field:"optional" json:"kmsId" yaml:"kmsId"`
}
