package dnszonev2


type DnsZoneV2Router struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.5/docs/resources/dns_zone_v2#router_id DnsZoneV2#router_id}.
	RouterId *string `field:"required" json:"routerId" yaml:"routerId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.5/docs/resources/dns_zone_v2#router_region DnsZoneV2#router_region}.
	RouterRegion *string `field:"required" json:"routerRegion" yaml:"routerRegion"`
}

