package wafcertificatev1


type WafCertificateV1Timeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opentelekomcloud/r/waf_certificate_v1#create WafCertificateV1#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opentelekomcloud/r/waf_certificate_v1#delete WafCertificateV1#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}
