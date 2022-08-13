// Prebuilt opentelekomcloud Provider for Terraform CDK (cdktf)
package opentelekomcloud

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CtsTrackerV1Config struct {
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
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opentelekomcloud/r/cts_tracker_v1#bucket_name CtsTrackerV1#bucket_name}.
	BucketName *string `field:"required" json:"bucketName" yaml:"bucketName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opentelekomcloud/r/cts_tracker_v1#is_support_smn CtsTrackerV1#is_support_smn}.
	IsSupportSmn interface{} `field:"required" json:"isSupportSmn" yaml:"isSupportSmn"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opentelekomcloud/r/cts_tracker_v1#file_prefix_name CtsTrackerV1#file_prefix_name}.
	FilePrefixName *string `field:"optional" json:"filePrefixName" yaml:"filePrefixName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opentelekomcloud/r/cts_tracker_v1#id CtsTrackerV1#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opentelekomcloud/r/cts_tracker_v1#is_send_all_key_operation CtsTrackerV1#is_send_all_key_operation}.
	IsSendAllKeyOperation interface{} `field:"optional" json:"isSendAllKeyOperation" yaml:"isSendAllKeyOperation"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opentelekomcloud/r/cts_tracker_v1#need_notify_user_list CtsTrackerV1#need_notify_user_list}.
	NeedNotifyUserList *[]*string `field:"optional" json:"needNotifyUserList" yaml:"needNotifyUserList"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opentelekomcloud/r/cts_tracker_v1#operations CtsTrackerV1#operations}.
	Operations *[]*string `field:"optional" json:"operations" yaml:"operations"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opentelekomcloud/r/cts_tracker_v1#project_name CtsTrackerV1#project_name}.
	ProjectName *string `field:"optional" json:"projectName" yaml:"projectName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opentelekomcloud/r/cts_tracker_v1#region CtsTrackerV1#region}.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opentelekomcloud/r/cts_tracker_v1#status CtsTrackerV1#status}.
	Status *string `field:"optional" json:"status" yaml:"status"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opentelekomcloud/r/cts_tracker_v1#timeouts CtsTrackerV1#timeouts}
	Timeouts *CtsTrackerV1Timeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opentelekomcloud/r/cts_tracker_v1#topic_id CtsTrackerV1#topic_id}.
	TopicId *string `field:"optional" json:"topicId" yaml:"topicId"`
}

