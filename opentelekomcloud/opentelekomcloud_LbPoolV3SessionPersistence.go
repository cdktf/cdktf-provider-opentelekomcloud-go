// Prebuilt opentelekomcloud Provider for Terraform CDK (cdktf)
package opentelekomcloud


type LbPoolV3SessionPersistence struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opentelekomcloud/r/lb_pool_v3#type LbPoolV3#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opentelekomcloud/r/lb_pool_v3#cookie_name LbPoolV3#cookie_name}.
	CookieName *string `field:"optional" json:"cookieName" yaml:"cookieName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opentelekomcloud/r/lb_pool_v3#persistence_timeout LbPoolV3#persistence_timeout}.
	PersistenceTimeout *float64 `field:"optional" json:"persistenceTimeout" yaml:"persistenceTimeout"`
}

