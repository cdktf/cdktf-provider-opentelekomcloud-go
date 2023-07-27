package dcsinstancev1


type DcsInstanceV1BackupPolicy struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.4/docs/resources/dcs_instance_v1#backup_at DcsInstanceV1#backup_at}.
	BackupAt *[]*float64 `field:"required" json:"backupAt" yaml:"backupAt"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.4/docs/resources/dcs_instance_v1#begin_at DcsInstanceV1#begin_at}.
	BeginAt *string `field:"required" json:"beginAt" yaml:"beginAt"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.4/docs/resources/dcs_instance_v1#period_type DcsInstanceV1#period_type}.
	PeriodType *string `field:"required" json:"periodType" yaml:"periodType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.4/docs/resources/dcs_instance_v1#backup_type DcsInstanceV1#backup_type}.
	BackupType *string `field:"optional" json:"backupType" yaml:"backupType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.4/docs/resources/dcs_instance_v1#save_days DcsInstanceV1#save_days}.
	SaveDays *float64 `field:"optional" json:"saveDays" yaml:"saveDays"`
}

