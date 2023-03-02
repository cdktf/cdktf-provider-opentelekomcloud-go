package dnsptrrecordv2


type DnsPtrrecordV2Timeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opentelekomcloud/r/dns_ptrrecord_v2#create DnsPtrrecordV2#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opentelekomcloud/r/dns_ptrrecord_v2#delete DnsPtrrecordV2#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opentelekomcloud/r/dns_ptrrecord_v2#update DnsPtrrecordV2#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

