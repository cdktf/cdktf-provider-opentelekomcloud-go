package lbpolicyv3


type LbPolicyV3RedirectPoolsConfig struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opentelekomcloud/r/lb_policy_v3#pool_id LbPolicyV3#pool_id}.
	PoolId *string `field:"required" json:"poolId" yaml:"poolId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opentelekomcloud/r/lb_policy_v3#weight LbPolicyV3#weight}.
	Weight *float64 `field:"required" json:"weight" yaml:"weight"`
}

