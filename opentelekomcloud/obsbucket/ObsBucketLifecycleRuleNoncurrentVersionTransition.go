package obsbucket


type ObsBucketLifecycleRuleNoncurrentVersionTransition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.34.3/docs/resources/obs_bucket#days ObsBucket#days}.
	Days *float64 `field:"required" json:"days" yaml:"days"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.34.3/docs/resources/obs_bucket#storage_class ObsBucket#storage_class}.
	StorageClass *string `field:"required" json:"storageClass" yaml:"storageClass"`
}

