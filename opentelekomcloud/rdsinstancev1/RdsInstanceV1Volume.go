package rdsinstancev1


type RdsInstanceV1Volume struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.34.2/docs/resources/rds_instance_v1#size RdsInstanceV1#size}.
	Size *float64 `field:"required" json:"size" yaml:"size"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.34.2/docs/resources/rds_instance_v1#type RdsInstanceV1#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
}

