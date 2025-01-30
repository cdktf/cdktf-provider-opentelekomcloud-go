// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package identityprovider


type IdentityProviderAccessConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.30/docs/resources/identity_provider#access_type IdentityProvider#access_type}.
	AccessType *string `field:"required" json:"accessType" yaml:"accessType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.30/docs/resources/identity_provider#client_id IdentityProvider#client_id}.
	ClientId *string `field:"required" json:"clientId" yaml:"clientId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.30/docs/resources/identity_provider#provider_url IdentityProvider#provider_url}.
	ProviderUrl *string `field:"required" json:"providerUrl" yaml:"providerUrl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.30/docs/resources/identity_provider#signing_key IdentityProvider#signing_key}.
	SigningKey *string `field:"required" json:"signingKey" yaml:"signingKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.30/docs/resources/identity_provider#authorization_endpoint IdentityProvider#authorization_endpoint}.
	AuthorizationEndpoint *string `field:"optional" json:"authorizationEndpoint" yaml:"authorizationEndpoint"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.30/docs/resources/identity_provider#response_mode IdentityProvider#response_mode}.
	ResponseMode *string `field:"optional" json:"responseMode" yaml:"responseMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.30/docs/resources/identity_provider#response_type IdentityProvider#response_type}.
	ResponseType *string `field:"optional" json:"responseType" yaml:"responseType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.30/docs/resources/identity_provider#scopes IdentityProvider#scopes}.
	Scopes *[]*string `field:"optional" json:"scopes" yaml:"scopes"`
}

