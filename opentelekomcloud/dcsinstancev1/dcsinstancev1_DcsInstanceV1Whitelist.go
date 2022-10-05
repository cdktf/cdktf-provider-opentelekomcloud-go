package dcsinstancev1


type DcsInstanceV1Whitelist struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opentelekomcloud/r/dcs_instance_v1#group_name DcsInstanceV1#group_name}.
	GroupName *string `field:"required" json:"groupName" yaml:"groupName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opentelekomcloud/r/dcs_instance_v1#ip_list DcsInstanceV1#ip_list}.
	IpList *[]*string `field:"required" json:"ipList" yaml:"ipList"`
}

