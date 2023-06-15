package cssclusterv1

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CssClusterV1Config struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.1/docs/resources/css_cluster_v1#name CssClusterV1#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// node_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.1/docs/resources/css_cluster_v1#node_config CssClusterV1#node_config}
	NodeConfig *CssClusterV1NodeConfig `field:"required" json:"nodeConfig" yaml:"nodeConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.1/docs/resources/css_cluster_v1#admin_pass CssClusterV1#admin_pass}.
	AdminPass *string `field:"optional" json:"adminPass" yaml:"adminPass"`
	// datastore block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.1/docs/resources/css_cluster_v1#datastore CssClusterV1#datastore}
	Datastore *CssClusterV1Datastore `field:"optional" json:"datastore" yaml:"datastore"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.1/docs/resources/css_cluster_v1#enable_authority CssClusterV1#enable_authority}.
	EnableAuthority interface{} `field:"optional" json:"enableAuthority" yaml:"enableAuthority"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.1/docs/resources/css_cluster_v1#enable_https CssClusterV1#enable_https}.
	EnableHttps interface{} `field:"optional" json:"enableHttps" yaml:"enableHttps"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.1/docs/resources/css_cluster_v1#expect_node_num CssClusterV1#expect_node_num}.
	ExpectNodeNum *float64 `field:"optional" json:"expectNodeNum" yaml:"expectNodeNum"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.1/docs/resources/css_cluster_v1#id CssClusterV1#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.1/docs/resources/css_cluster_v1#tags CssClusterV1#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.1/docs/resources/css_cluster_v1#timeouts CssClusterV1#timeouts}
	Timeouts *CssClusterV1Timeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

