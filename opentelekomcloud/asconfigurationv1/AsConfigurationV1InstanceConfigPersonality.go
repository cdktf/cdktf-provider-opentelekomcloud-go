package asconfigurationv1


type AsConfigurationV1InstanceConfigPersonality struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opentelekomcloud/r/as_configuration_v1#content AsConfigurationV1#content}.
	Content *string `field:"required" json:"content" yaml:"content"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opentelekomcloud/r/as_configuration_v1#path AsConfigurationV1#path}.
	Path *string `field:"required" json:"path" yaml:"path"`
}

