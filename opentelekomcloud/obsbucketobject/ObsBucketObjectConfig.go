package obsbucketobject

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ObsBucketObjectConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.34.3/docs/resources/obs_bucket_object#bucket ObsBucketObject#bucket}.
	Bucket *string `field:"required" json:"bucket" yaml:"bucket"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.34.3/docs/resources/obs_bucket_object#key ObsBucketObject#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.34.3/docs/resources/obs_bucket_object#acl ObsBucketObject#acl}.
	Acl *string `field:"optional" json:"acl" yaml:"acl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.34.3/docs/resources/obs_bucket_object#content ObsBucketObject#content}.
	Content *string `field:"optional" json:"content" yaml:"content"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.34.3/docs/resources/obs_bucket_object#content_type ObsBucketObject#content_type}.
	ContentType *string `field:"optional" json:"contentType" yaml:"contentType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.34.3/docs/resources/obs_bucket_object#encryption ObsBucketObject#encryption}.
	Encryption interface{} `field:"optional" json:"encryption" yaml:"encryption"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.34.3/docs/resources/obs_bucket_object#etag ObsBucketObject#etag}.
	Etag *string `field:"optional" json:"etag" yaml:"etag"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.34.3/docs/resources/obs_bucket_object#id ObsBucketObject#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.34.3/docs/resources/obs_bucket_object#kms_key_id ObsBucketObject#kms_key_id}.
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.34.3/docs/resources/obs_bucket_object#source ObsBucketObject#source}.
	Source *string `field:"optional" json:"source" yaml:"source"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.34.3/docs/resources/obs_bucket_object#storage_class ObsBucketObject#storage_class}.
	StorageClass *string `field:"optional" json:"storageClass" yaml:"storageClass"`
}

