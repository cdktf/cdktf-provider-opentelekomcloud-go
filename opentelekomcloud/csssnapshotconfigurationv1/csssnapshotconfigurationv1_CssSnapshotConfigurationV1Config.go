package csssnapshotconfigurationv1

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CssSnapshotConfigurationV1Config struct {
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
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opentelekomcloud/r/css_snapshot_configuration_v1#cluster_id CssSnapshotConfigurationV1#cluster_id}.
	ClusterId *string `field:"required" json:"clusterId" yaml:"clusterId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opentelekomcloud/r/css_snapshot_configuration_v1#automatic CssSnapshotConfigurationV1#automatic}.
	Automatic interface{} `field:"optional" json:"automatic" yaml:"automatic"`
	// configuration block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opentelekomcloud/r/css_snapshot_configuration_v1#configuration CssSnapshotConfigurationV1#configuration}
	Configuration *CssSnapshotConfigurationV1Configuration `field:"optional" json:"configuration" yaml:"configuration"`
	// creation_policy block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opentelekomcloud/r/css_snapshot_configuration_v1#creation_policy CssSnapshotConfigurationV1#creation_policy}
	CreationPolicy *CssSnapshotConfigurationV1CreationPolicy `field:"optional" json:"creationPolicy" yaml:"creationPolicy"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opentelekomcloud/r/css_snapshot_configuration_v1#id CssSnapshotConfigurationV1#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opentelekomcloud/r/css_snapshot_configuration_v1#timeouts CssSnapshotConfigurationV1#timeouts}
	Timeouts *CssSnapshotConfigurationV1Timeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}
