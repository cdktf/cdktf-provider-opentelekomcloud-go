package vpcepservicev1


type VpcepServiceV1Port struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opentelekomcloud/r/vpcep_service_v1#client_port VpcepServiceV1#client_port}.
	ClientPort *float64 `field:"required" json:"clientPort" yaml:"clientPort"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opentelekomcloud/r/vpcep_service_v1#server_port VpcepServiceV1#server_port}.
	ServerPort *float64 `field:"required" json:"serverPort" yaml:"serverPort"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opentelekomcloud/r/vpcep_service_v1#protocol VpcepServiceV1#protocol}.
	Protocol *string `field:"optional" json:"protocol" yaml:"protocol"`
}

