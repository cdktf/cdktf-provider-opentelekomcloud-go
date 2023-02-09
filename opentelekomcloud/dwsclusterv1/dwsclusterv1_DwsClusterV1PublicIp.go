package dwsclusterv1


type DwsClusterV1PublicIp struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opentelekomcloud/r/dws_cluster_v1#eip_id DwsClusterV1#eip_id}.
	EipId *string `field:"optional" json:"eipId" yaml:"eipId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opentelekomcloud/r/dws_cluster_v1#public_bind_type DwsClusterV1#public_bind_type}.
	PublicBindType *string `field:"optional" json:"publicBindType" yaml:"publicBindType"`
}

