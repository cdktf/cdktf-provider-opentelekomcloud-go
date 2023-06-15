package s3bucket


type S3BucketLifecycleRuleNoncurrentVersionExpiration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.1/docs/resources/s3_bucket#days S3Bucket#days}.
	Days *float64 `field:"optional" json:"days" yaml:"days"`
}

