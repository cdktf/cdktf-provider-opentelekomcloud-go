package provider


type OpentelekomcloudProviderConfig struct {
	// The access key for API operations. You can retrieve this from the 'My Credential' section of the console.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opentelekomcloud#access_key OpentelekomcloudProvider#access_key}
	AccessKey *string `field:"optional" json:"accessKey" yaml:"accessKey"`
	// The name of domain who created the agency (Identity v3).
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opentelekomcloud#agency_domain_name OpentelekomcloudProvider#agency_domain_name}
	AgencyDomainName *string `field:"optional" json:"agencyDomainName" yaml:"agencyDomainName"`
	// The name of agency.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opentelekomcloud#agency_name OpentelekomcloudProvider#agency_name}
	AgencyName *string `field:"optional" json:"agencyName" yaml:"agencyName"`
	// Alias name.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opentelekomcloud#alias OpentelekomcloudProvider#alias}
	Alias *string `field:"optional" json:"alias" yaml:"alias"`
	// The Identity authentication URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opentelekomcloud#auth_url OpentelekomcloudProvider#auth_url}
	AuthUrl *string `field:"optional" json:"authUrl" yaml:"authUrl"`
	// A Custom CA certificate.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opentelekomcloud#cacert_file OpentelekomcloudProvider#cacert_file}
	CacertFile *string `field:"optional" json:"cacertFile" yaml:"cacertFile"`
	// A client certificate to authenticate with.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opentelekomcloud#cert OpentelekomcloudProvider#cert}
	Cert *string `field:"optional" json:"cert" yaml:"cert"`
	// An entry in a `clouds.yaml` file to use.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opentelekomcloud#cloud OpentelekomcloudProvider#cloud}
	Cloud *string `field:"optional" json:"cloud" yaml:"cloud"`
	// The name of delegated project (Identity v3).
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opentelekomcloud#delegated_project OpentelekomcloudProvider#delegated_project}
	DelegatedProject *string `field:"optional" json:"delegatedProject" yaml:"delegatedProject"`
	// The ID of the Domain to scope to (Identity v3).
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opentelekomcloud#domain_id OpentelekomcloudProvider#domain_id}
	DomainId *string `field:"optional" json:"domainId" yaml:"domainId"`
	// The name of the Domain to scope to (Identity v3).
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opentelekomcloud#domain_name OpentelekomcloudProvider#domain_name}
	DomainName *string `field:"optional" json:"domainName" yaml:"domainName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opentelekomcloud#endpoint_type OpentelekomcloudProvider#endpoint_type}.
	EndpointType *string `field:"optional" json:"endpointType" yaml:"endpointType"`
	// Trust self-signed certificates.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opentelekomcloud#insecure OpentelekomcloudProvider#insecure}
	Insecure interface{} `field:"optional" json:"insecure" yaml:"insecure"`
	// A client private key to authenticate with.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opentelekomcloud#key OpentelekomcloudProvider#key}
	Key *string `field:"optional" json:"key" yaml:"key"`
	// How many times HTTP connection should be retried until giving up.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opentelekomcloud#max_retries OpentelekomcloudProvider#max_retries}
	MaxRetries *float64 `field:"optional" json:"maxRetries" yaml:"maxRetries"`
	// One-time MFA passcode.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opentelekomcloud#passcode OpentelekomcloudProvider#passcode}
	Passcode *string `field:"optional" json:"passcode" yaml:"passcode"`
	// Password to login with.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opentelekomcloud#password OpentelekomcloudProvider#password}
	Password *string `field:"optional" json:"password" yaml:"password"`
	// The OpenTelekomCloud region to connect to.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opentelekomcloud#region OpentelekomcloudProvider#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// The secret key for API operations. You can retrieve this from the 'My Credential' section of the console.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opentelekomcloud#secret_key OpentelekomcloudProvider#secret_key}
	SecretKey *string `field:"optional" json:"secretKey" yaml:"secretKey"`
	// Security token to use for OBS federated authentication.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opentelekomcloud#security_token OpentelekomcloudProvider#security_token}
	SecurityToken *string `field:"optional" json:"securityToken" yaml:"securityToken"`
	// Use Swift's authentication system instead of Keystone. Only used for interaction with Swift.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opentelekomcloud#swauth OpentelekomcloudProvider#swauth}
	Swauth interface{} `field:"optional" json:"swauth" yaml:"swauth"`
	// The ID of the Tenant (Identity v2) or Project (Identity v3) to login with.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opentelekomcloud#tenant_id OpentelekomcloudProvider#tenant_id}
	TenantId *string `field:"optional" json:"tenantId" yaml:"tenantId"`
	// The name of the Tenant (Identity v2) or Project (Identity v3) to login with.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opentelekomcloud#tenant_name OpentelekomcloudProvider#tenant_name}
	TenantName *string `field:"optional" json:"tenantName" yaml:"tenantName"`
	// Authentication token to use as an alternative to username/password.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opentelekomcloud#token OpentelekomcloudProvider#token}
	Token *string `field:"optional" json:"token" yaml:"token"`
	// User ID to login with.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opentelekomcloud#user_id OpentelekomcloudProvider#user_id}
	UserId *string `field:"optional" json:"userId" yaml:"userId"`
	// Username to login with.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opentelekomcloud#user_name OpentelekomcloudProvider#user_name}
	UserName *string `field:"optional" json:"userName" yaml:"userName"`
}
