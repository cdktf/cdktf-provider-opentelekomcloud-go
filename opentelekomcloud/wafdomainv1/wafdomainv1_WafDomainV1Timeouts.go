package wafdomainv1


type WafDomainV1Timeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opentelekomcloud/r/waf_domain_v1#create WafDomainV1#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opentelekomcloud/r/waf_domain_v1#delete WafDomainV1#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}
