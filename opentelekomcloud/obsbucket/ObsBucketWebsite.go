// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package obsbucket


type ObsBucketWebsite struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.30/docs/resources/obs_bucket#error_document ObsBucket#error_document}.
	ErrorDocument *string `field:"optional" json:"errorDocument" yaml:"errorDocument"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.30/docs/resources/obs_bucket#index_document ObsBucket#index_document}.
	IndexDocument *string `field:"optional" json:"indexDocument" yaml:"indexDocument"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.30/docs/resources/obs_bucket#redirect_all_requests_to ObsBucket#redirect_all_requests_to}.
	RedirectAllRequestsTo *string `field:"optional" json:"redirectAllRequestsTo" yaml:"redirectAllRequestsTo"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.30/docs/resources/obs_bucket#routing_rules ObsBucket#routing_rules}.
	RoutingRules *string `field:"optional" json:"routingRules" yaml:"routingRules"`
}

