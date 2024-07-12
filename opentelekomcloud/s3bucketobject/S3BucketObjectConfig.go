// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package s3bucketobject

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type S3BucketObjectConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.14/docs/resources/s3_bucket_object#bucket S3BucketObject#bucket}.
	Bucket *string `field:"required" json:"bucket" yaml:"bucket"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.14/docs/resources/s3_bucket_object#key S3BucketObject#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.14/docs/resources/s3_bucket_object#acl S3BucketObject#acl}.
	Acl *string `field:"optional" json:"acl" yaml:"acl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.14/docs/resources/s3_bucket_object#cache_control S3BucketObject#cache_control}.
	CacheControl *string `field:"optional" json:"cacheControl" yaml:"cacheControl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.14/docs/resources/s3_bucket_object#content S3BucketObject#content}.
	Content *string `field:"optional" json:"content" yaml:"content"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.14/docs/resources/s3_bucket_object#content_disposition S3BucketObject#content_disposition}.
	ContentDisposition *string `field:"optional" json:"contentDisposition" yaml:"contentDisposition"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.14/docs/resources/s3_bucket_object#content_encoding S3BucketObject#content_encoding}.
	ContentEncoding *string `field:"optional" json:"contentEncoding" yaml:"contentEncoding"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.14/docs/resources/s3_bucket_object#content_language S3BucketObject#content_language}.
	ContentLanguage *string `field:"optional" json:"contentLanguage" yaml:"contentLanguage"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.14/docs/resources/s3_bucket_object#content_type S3BucketObject#content_type}.
	ContentType *string `field:"optional" json:"contentType" yaml:"contentType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.14/docs/resources/s3_bucket_object#etag S3BucketObject#etag}.
	Etag *string `field:"optional" json:"etag" yaml:"etag"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.14/docs/resources/s3_bucket_object#id S3BucketObject#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.14/docs/resources/s3_bucket_object#server_side_encryption S3BucketObject#server_side_encryption}.
	ServerSideEncryption *string `field:"optional" json:"serverSideEncryption" yaml:"serverSideEncryption"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.14/docs/resources/s3_bucket_object#source S3BucketObject#source}.
	Source *string `field:"optional" json:"source" yaml:"source"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.14/docs/resources/s3_bucket_object#sse_kms_key_id S3BucketObject#sse_kms_key_id}.
	SseKmsKeyId *string `field:"optional" json:"sseKmsKeyId" yaml:"sseKmsKeyId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.14/docs/resources/s3_bucket_object#website_redirect S3BucketObject#website_redirect}.
	WebsiteRedirect *string `field:"optional" json:"websiteRedirect" yaml:"websiteRedirect"`
}

