// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package smntopicattributev2

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SmnTopicAttributeV2Config struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.19/docs/resources/smn_topic_attribute_v2#attribute_name SmnTopicAttributeV2#attribute_name}.
	AttributeName *string `field:"required" json:"attributeName" yaml:"attributeName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.19/docs/resources/smn_topic_attribute_v2#topic_attribute SmnTopicAttributeV2#topic_attribute}.
	TopicAttribute *string `field:"required" json:"topicAttribute" yaml:"topicAttribute"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.19/docs/resources/smn_topic_attribute_v2#topic_urn SmnTopicAttributeV2#topic_urn}.
	TopicUrn *string `field:"required" json:"topicUrn" yaml:"topicUrn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.19/docs/resources/smn_topic_attribute_v2#id SmnTopicAttributeV2#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
}

