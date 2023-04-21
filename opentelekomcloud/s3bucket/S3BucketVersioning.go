package s3bucket


type S3BucketVersioning struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.34.2/docs/resources/s3_bucket#enabled S3Bucket#enabled}.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.34.2/docs/resources/s3_bucket#mfa_delete S3Bucket#mfa_delete}.
	MfaDelete interface{} `field:"optional" json:"mfaDelete" yaml:"mfaDelete"`
}

