package identityagencyv3


type IdentityAgencyV3ProjectRole struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.34.3/docs/resources/identity_agency_v3#project IdentityAgencyV3#project}.
	Project *string `field:"required" json:"project" yaml:"project"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.34.3/docs/resources/identity_agency_v3#roles IdentityAgencyV3#roles}.
	Roles *[]*string `field:"required" json:"roles" yaml:"roles"`
}

