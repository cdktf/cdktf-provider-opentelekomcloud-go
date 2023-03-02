package rdsinstancev3


type RdsInstanceV3RestorePoint struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opentelekomcloud/r/rds_instance_v3#instance_id RdsInstanceV3#instance_id}.
	InstanceId *string `field:"required" json:"instanceId" yaml:"instanceId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opentelekomcloud/r/rds_instance_v3#backup_id RdsInstanceV3#backup_id}.
	BackupId *string `field:"optional" json:"backupId" yaml:"backupId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opentelekomcloud/r/rds_instance_v3#restore_time RdsInstanceV3#restore_time}.
	RestoreTime *float64 `field:"optional" json:"restoreTime" yaml:"restoreTime"`
}

