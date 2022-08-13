// Prebuilt opentelekomcloud Provider for Terraform CDK (cdktf)
package opentelekomcloud


type RdsReadReplicaV3Volume struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opentelekomcloud/r/rds_read_replica_v3#type RdsReadReplicaV3#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opentelekomcloud/r/rds_read_replica_v3#disk_encryption_id RdsReadReplicaV3#disk_encryption_id}.
	DiskEncryptionId *string `field:"optional" json:"diskEncryptionId" yaml:"diskEncryptionId"`
}

