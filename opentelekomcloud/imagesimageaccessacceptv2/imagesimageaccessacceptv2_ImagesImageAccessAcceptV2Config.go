package imagesimageaccessacceptv2

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ImagesImageAccessAcceptV2Config struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count *float64 `field:"optional" json:"count" yaml:"count"`
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
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opentelekomcloud/r/images_image_access_accept_v2#image_id ImagesImageAccessAcceptV2#image_id}.
	ImageId *string `field:"required" json:"imageId" yaml:"imageId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opentelekomcloud/r/images_image_access_accept_v2#member_id ImagesImageAccessAcceptV2#member_id}.
	MemberId *string `field:"required" json:"memberId" yaml:"memberId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opentelekomcloud/r/images_image_access_accept_v2#status ImagesImageAccessAcceptV2#status}.
	Status *string `field:"required" json:"status" yaml:"status"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opentelekomcloud/r/images_image_access_accept_v2#id ImagesImageAccessAcceptV2#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
}
