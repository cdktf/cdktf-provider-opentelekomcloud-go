package obsbucket


type ObsBucketLifecycleRuleNoncurrentVersionExpiration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.34.4/docs/resources/obs_bucket#days ObsBucket#days}.
	Days *float64 `field:"required" json:"days" yaml:"days"`
}

