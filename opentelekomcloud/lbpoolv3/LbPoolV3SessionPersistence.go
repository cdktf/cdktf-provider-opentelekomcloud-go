// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package lbpoolv3


type LbPoolV3SessionPersistence struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.35/docs/resources/lb_pool_v3#type LbPoolV3#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.35/docs/resources/lb_pool_v3#cookie_name LbPoolV3#cookie_name}.
	CookieName *string `field:"optional" json:"cookieName" yaml:"cookieName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.35/docs/resources/lb_pool_v3#persistence_timeout LbPoolV3#persistence_timeout}.
	PersistenceTimeout *float64 `field:"optional" json:"persistenceTimeout" yaml:"persistenceTimeout"`
}

