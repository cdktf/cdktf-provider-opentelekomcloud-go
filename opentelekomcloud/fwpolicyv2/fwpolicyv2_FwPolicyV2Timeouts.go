package fwpolicyv2


type FwPolicyV2Timeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opentelekomcloud/r/fw_policy_v2#create FwPolicyV2#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opentelekomcloud/r/fw_policy_v2#delete FwPolicyV2#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

