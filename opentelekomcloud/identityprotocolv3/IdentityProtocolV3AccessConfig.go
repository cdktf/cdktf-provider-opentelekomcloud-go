// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package identityprotocolv3


type IdentityProtocolV3AccessConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.52/docs/resources/identity_protocol_v3#access_type IdentityProtocolV3#access_type}.
	AccessType *string `field:"required" json:"accessType" yaml:"accessType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.52/docs/resources/identity_protocol_v3#client_id IdentityProtocolV3#client_id}.
	ClientId *string `field:"required" json:"clientId" yaml:"clientId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.52/docs/resources/identity_protocol_v3#provider_url IdentityProtocolV3#provider_url}.
	ProviderUrl *string `field:"required" json:"providerUrl" yaml:"providerUrl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.52/docs/resources/identity_protocol_v3#signing_key IdentityProtocolV3#signing_key}.
	SigningKey *string `field:"required" json:"signingKey" yaml:"signingKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.52/docs/resources/identity_protocol_v3#authorization_endpoint IdentityProtocolV3#authorization_endpoint}.
	AuthorizationEndpoint *string `field:"optional" json:"authorizationEndpoint" yaml:"authorizationEndpoint"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.52/docs/resources/identity_protocol_v3#response_mode IdentityProtocolV3#response_mode}.
	ResponseMode *string `field:"optional" json:"responseMode" yaml:"responseMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.52/docs/resources/identity_protocol_v3#response_type IdentityProtocolV3#response_type}.
	ResponseType *string `field:"optional" json:"responseType" yaml:"responseType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.52/docs/resources/identity_protocol_v3#scopes IdentityProtocolV3#scopes}.
	Scopes *[]*string `field:"optional" json:"scopes" yaml:"scopes"`
}

