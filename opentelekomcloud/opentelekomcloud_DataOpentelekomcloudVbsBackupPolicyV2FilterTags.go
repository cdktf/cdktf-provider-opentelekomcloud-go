// Prebuilt opentelekomcloud Provider for Terraform CDK (cdktf)
package opentelekomcloud


type DataOpentelekomcloudVbsBackupPolicyV2FilterTags struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opentelekomcloud/d/vbs_backup_policy_v2#key DataOpentelekomcloudVbsBackupPolicyV2#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opentelekomcloud/d/vbs_backup_policy_v2#values DataOpentelekomcloudVbsBackupPolicyV2#values}.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}
