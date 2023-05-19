package dnszonev2


type DnsZoneV2Timeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.34.4/docs/resources/dns_zone_v2#create DnsZoneV2#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.34.4/docs/resources/dns_zone_v2#delete DnsZoneV2#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.34.4/docs/resources/dns_zone_v2#update DnsZoneV2#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

