// Prebuilt opentelekomcloud Provider for Terraform CDK (cdktf)
package opentelekomcloud


type DnsZoneV2Router struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opentelekomcloud/r/dns_zone_v2#router_id DnsZoneV2#router_id}.
	RouterId *string `field:"required" json:"routerId" yaml:"routerId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opentelekomcloud/r/dns_zone_v2#router_region DnsZoneV2#router_region}.
	RouterRegion *string `field:"required" json:"routerRegion" yaml:"routerRegion"`
}
