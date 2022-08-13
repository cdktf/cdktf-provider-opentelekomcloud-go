// Prebuilt opentelekomcloud Provider for Terraform CDK (cdktf)
package opentelekomcloud


type ComputeInstanceV2Network struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opentelekomcloud/r/compute_instance_v2#access_network ComputeInstanceV2#access_network}.
	AccessNetwork interface{} `field:"optional" json:"accessNetwork" yaml:"accessNetwork"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opentelekomcloud/r/compute_instance_v2#fixed_ip_v4 ComputeInstanceV2#fixed_ip_v4}.
	FixedIpV4 *string `field:"optional" json:"fixedIpV4" yaml:"fixedIpV4"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opentelekomcloud/r/compute_instance_v2#fixed_ip_v6 ComputeInstanceV2#fixed_ip_v6}.
	FixedIpV6 *string `field:"optional" json:"fixedIpV6" yaml:"fixedIpV6"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opentelekomcloud/r/compute_instance_v2#name ComputeInstanceV2#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opentelekomcloud/r/compute_instance_v2#port ComputeInstanceV2#port}.
	Port *string `field:"optional" json:"port" yaml:"port"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opentelekomcloud/r/compute_instance_v2#uuid ComputeInstanceV2#uuid}.
	Uuid *string `field:"optional" json:"uuid" yaml:"uuid"`
}

