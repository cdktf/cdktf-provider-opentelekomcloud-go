package natgatewayv2


type NatGatewayV2Timeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.3/docs/resources/nat_gateway_v2#create NatGatewayV2#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.3/docs/resources/nat_gateway_v2#delete NatGatewayV2#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

