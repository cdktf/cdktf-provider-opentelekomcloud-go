// Prebuilt opentelekomcloud Provider for Terraform CDK (cdktf)
package opentelekomcloud


type IdentityProtocolV3Metadata struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opentelekomcloud/r/identity_protocol_v3#domain_id IdentityProtocolV3#domain_id}.
	DomainId *string `field:"required" json:"domainId" yaml:"domainId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opentelekomcloud/r/identity_protocol_v3#metadata IdentityProtocolV3#metadata}.
	Metadata *string `field:"required" json:"metadata" yaml:"metadata"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opentelekomcloud/r/identity_protocol_v3#xaccount_type IdentityProtocolV3#xaccount_type}.
	XaccountType *string `field:"optional" json:"xaccountType" yaml:"xaccountType"`
}
