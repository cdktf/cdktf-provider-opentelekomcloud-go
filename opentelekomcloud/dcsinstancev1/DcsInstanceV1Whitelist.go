package dcsinstancev1


type DcsInstanceV1Whitelist struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.34.2/docs/resources/dcs_instance_v1#group_name DcsInstanceV1#group_name}.
	GroupName *string `field:"required" json:"groupName" yaml:"groupName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.34.2/docs/resources/dcs_instance_v1#ip_list DcsInstanceV1#ip_list}.
	IpList *[]*string `field:"required" json:"ipList" yaml:"ipList"`
}

