package lbcertificatev2


type LbCertificateV2Timeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opentelekomcloud/r/lb_certificate_v2#create LbCertificateV2#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opentelekomcloud/r/lb_certificate_v2#delete LbCertificateV2#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opentelekomcloud/r/lb_certificate_v2#update LbCertificateV2#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

