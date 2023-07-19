package dcsinstancev1


type DcsInstanceV1Configuration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.3/docs/resources/dcs_instance_v1#parameter_id DcsInstanceV1#parameter_id}.
	ParameterId *string `field:"required" json:"parameterId" yaml:"parameterId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.3/docs/resources/dcs_instance_v1#parameter_name DcsInstanceV1#parameter_name}.
	ParameterName *string `field:"required" json:"parameterName" yaml:"parameterName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.3/docs/resources/dcs_instance_v1#parameter_value DcsInstanceV1#parameter_value}.
	ParameterValue *string `field:"required" json:"parameterValue" yaml:"parameterValue"`
}

