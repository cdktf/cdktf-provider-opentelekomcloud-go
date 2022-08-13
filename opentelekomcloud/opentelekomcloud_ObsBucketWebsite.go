// Prebuilt opentelekomcloud Provider for Terraform CDK (cdktf)
package opentelekomcloud


type ObsBucketWebsite struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opentelekomcloud/r/obs_bucket#error_document ObsBucket#error_document}.
	ErrorDocument *string `field:"optional" json:"errorDocument" yaml:"errorDocument"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opentelekomcloud/r/obs_bucket#index_document ObsBucket#index_document}.
	IndexDocument *string `field:"optional" json:"indexDocument" yaml:"indexDocument"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opentelekomcloud/r/obs_bucket#redirect_all_requests_to ObsBucket#redirect_all_requests_to}.
	RedirectAllRequestsTo *string `field:"optional" json:"redirectAllRequestsTo" yaml:"redirectAllRequestsTo"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opentelekomcloud/r/obs_bucket#routing_rules ObsBucket#routing_rules}.
	RoutingRules *string `field:"optional" json:"routingRules" yaml:"routingRules"`
}

