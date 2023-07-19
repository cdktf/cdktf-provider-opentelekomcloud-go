package provider

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v8/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v8/provider/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.3/docs opentelekomcloud}.
type OpentelekomcloudProvider interface {
	cdktf.TerraformProvider
	AccessKey() *string
	SetAccessKey(val *string)
	AccessKeyInput() *string
	AgencyDomainName() *string
	SetAgencyDomainName(val *string)
	AgencyDomainNameInput() *string
	AgencyName() *string
	SetAgencyName(val *string)
	AgencyNameInput() *string
	Alias() *string
	SetAlias(val *string)
	AliasInput() *string
	AuthUrl() *string
	SetAuthUrl(val *string)
	AuthUrlInput() *string
	CacertFile() *string
	SetCacertFile(val *string)
	CacertFileInput() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	Cert() *string
	SetCert(val *string)
	CertInput() *string
	Cloud() *string
	SetCloud(val *string)
	CloudInput() *string
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	DelegatedProject() *string
	SetDelegatedProject(val *string)
	DelegatedProjectInput() *string
	DomainId() *string
	SetDomainId(val *string)
	DomainIdInput() *string
	DomainName() *string
	SetDomainName(val *string)
	DomainNameInput() *string
	EndpointType() *string
	SetEndpointType(val *string)
	EndpointTypeInput() *string
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Insecure() interface{}
	SetInsecure(val interface{})
	InsecureInput() interface{}
	Key() *string
	SetKey(val *string)
	KeyInput() *string
	MaxRetries() *float64
	SetMaxRetries(val *float64)
	MaxRetriesInput() *float64
	// Experimental.
	MetaAttributes() *map[string]interface{}
	// The tree node.
	Node() constructs.Node
	Passcode() *string
	SetPasscode(val *string)
	PasscodeInput() *string
	Password() *string
	SetPassword(val *string)
	PasswordInput() *string
	// Experimental.
	RawOverrides() interface{}
	Region() *string
	SetRegion(val *string)
	RegionInput() *string
	SecretKey() *string
	SetSecretKey(val *string)
	SecretKeyInput() *string
	SecurityToken() *string
	SetSecurityToken(val *string)
	SecurityTokenInput() *string
	Swauth() interface{}
	SetSwauth(val interface{})
	SwauthInput() interface{}
	TenantId() *string
	SetTenantId(val *string)
	TenantIdInput() *string
	TenantName() *string
	SetTenantName(val *string)
	TenantNameInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformProviderSource() *string
	// Experimental.
	TerraformResourceType() *string
	Token() *string
	SetToken(val *string)
	TokenInput() *string
	UserId() *string
	SetUserId(val *string)
	UserIdInput() *string
	UserName() *string
	SetUserName(val *string)
	UserNameInput() *string
	// Experimental.
	AddOverride(path *string, value interface{})
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	ResetAccessKey()
	ResetAgencyDomainName()
	ResetAgencyName()
	ResetAlias()
	ResetAuthUrl()
	ResetCacertFile()
	ResetCert()
	ResetCloud()
	ResetDelegatedProject()
	ResetDomainId()
	ResetDomainName()
	ResetEndpointType()
	ResetInsecure()
	ResetKey()
	ResetMaxRetries()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPasscode()
	ResetPassword()
	ResetRegion()
	ResetSecretKey()
	ResetSecurityToken()
	ResetSwauth()
	ResetTenantId()
	ResetTenantName()
	ResetToken()
	ResetUserId()
	ResetUserName()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for OpentelekomcloudProvider
type jsiiProxy_OpentelekomcloudProvider struct {
	internal.Type__cdktfTerraformProvider
}

func (j *jsiiProxy_OpentelekomcloudProvider) AccessKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpentelekomcloudProvider) AccessKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpentelekomcloudProvider) AgencyDomainName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"agencyDomainName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpentelekomcloudProvider) AgencyDomainNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"agencyDomainNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpentelekomcloudProvider) AgencyName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"agencyName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpentelekomcloudProvider) AgencyNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"agencyNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpentelekomcloudProvider) Alias() *string {
	var returns *string
	_jsii_.Get(
		j,
		"alias",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpentelekomcloudProvider) AliasInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"aliasInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpentelekomcloudProvider) AuthUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpentelekomcloudProvider) AuthUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpentelekomcloudProvider) CacertFile() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cacertFile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpentelekomcloudProvider) CacertFileInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cacertFileInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpentelekomcloudProvider) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpentelekomcloudProvider) Cert() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cert",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpentelekomcloudProvider) CertInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpentelekomcloudProvider) Cloud() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloud",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpentelekomcloudProvider) CloudInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpentelekomcloudProvider) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpentelekomcloudProvider) DelegatedProject() *string {
	var returns *string
	_jsii_.Get(
		j,
		"delegatedProject",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpentelekomcloudProvider) DelegatedProjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"delegatedProjectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpentelekomcloudProvider) DomainId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpentelekomcloudProvider) DomainIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpentelekomcloudProvider) DomainName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpentelekomcloudProvider) DomainNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpentelekomcloudProvider) EndpointType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpointType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpentelekomcloudProvider) EndpointTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpointTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpentelekomcloudProvider) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpentelekomcloudProvider) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpentelekomcloudProvider) Insecure() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"insecure",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpentelekomcloudProvider) InsecureInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"insecureInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpentelekomcloudProvider) Key() *string {
	var returns *string
	_jsii_.Get(
		j,
		"key",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpentelekomcloudProvider) KeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpentelekomcloudProvider) MaxRetries() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxRetries",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpentelekomcloudProvider) MaxRetriesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxRetriesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpentelekomcloudProvider) MetaAttributes() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"metaAttributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpentelekomcloudProvider) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpentelekomcloudProvider) Passcode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"passcode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpentelekomcloudProvider) PasscodeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"passcodeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpentelekomcloudProvider) Password() *string {
	var returns *string
	_jsii_.Get(
		j,
		"password",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpentelekomcloudProvider) PasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"passwordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpentelekomcloudProvider) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpentelekomcloudProvider) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpentelekomcloudProvider) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpentelekomcloudProvider) SecretKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpentelekomcloudProvider) SecretKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpentelekomcloudProvider) SecurityToken() *string {
	var returns *string
	_jsii_.Get(
		j,
		"securityToken",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpentelekomcloudProvider) SecurityTokenInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"securityTokenInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpentelekomcloudProvider) Swauth() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"swauth",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpentelekomcloudProvider) SwauthInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"swauthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpentelekomcloudProvider) TenantId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tenantId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpentelekomcloudProvider) TenantIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tenantIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpentelekomcloudProvider) TenantName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tenantName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpentelekomcloudProvider) TenantNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tenantNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpentelekomcloudProvider) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpentelekomcloudProvider) TerraformProviderSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformProviderSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpentelekomcloudProvider) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpentelekomcloudProvider) Token() *string {
	var returns *string
	_jsii_.Get(
		j,
		"token",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpentelekomcloudProvider) TokenInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tokenInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpentelekomcloudProvider) UserId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpentelekomcloudProvider) UserIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpentelekomcloudProvider) UserName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpentelekomcloudProvider) UserNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userNameInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.3/docs opentelekomcloud} Resource.
func NewOpentelekomcloudProvider(scope constructs.Construct, id *string, config *OpentelekomcloudProviderConfig) OpentelekomcloudProvider {
	_init_.Initialize()

	if err := validateNewOpentelekomcloudProviderParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_OpentelekomcloudProvider{}

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.provider.OpentelekomcloudProvider",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.3/docs opentelekomcloud} Resource.
func NewOpentelekomcloudProvider_Override(o OpentelekomcloudProvider, scope constructs.Construct, id *string, config *OpentelekomcloudProviderConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.provider.OpentelekomcloudProvider",
		[]interface{}{scope, id, config},
		o,
	)
}

func (j *jsiiProxy_OpentelekomcloudProvider)SetAccessKey(val *string) {
	_jsii_.Set(
		j,
		"accessKey",
		val,
	)
}

func (j *jsiiProxy_OpentelekomcloudProvider)SetAgencyDomainName(val *string) {
	_jsii_.Set(
		j,
		"agencyDomainName",
		val,
	)
}

func (j *jsiiProxy_OpentelekomcloudProvider)SetAgencyName(val *string) {
	_jsii_.Set(
		j,
		"agencyName",
		val,
	)
}

func (j *jsiiProxy_OpentelekomcloudProvider)SetAlias(val *string) {
	_jsii_.Set(
		j,
		"alias",
		val,
	)
}

func (j *jsiiProxy_OpentelekomcloudProvider)SetAuthUrl(val *string) {
	_jsii_.Set(
		j,
		"authUrl",
		val,
	)
}

func (j *jsiiProxy_OpentelekomcloudProvider)SetCacertFile(val *string) {
	_jsii_.Set(
		j,
		"cacertFile",
		val,
	)
}

func (j *jsiiProxy_OpentelekomcloudProvider)SetCert(val *string) {
	_jsii_.Set(
		j,
		"cert",
		val,
	)
}

func (j *jsiiProxy_OpentelekomcloudProvider)SetCloud(val *string) {
	_jsii_.Set(
		j,
		"cloud",
		val,
	)
}

func (j *jsiiProxy_OpentelekomcloudProvider)SetDelegatedProject(val *string) {
	_jsii_.Set(
		j,
		"delegatedProject",
		val,
	)
}

func (j *jsiiProxy_OpentelekomcloudProvider)SetDomainId(val *string) {
	_jsii_.Set(
		j,
		"domainId",
		val,
	)
}

func (j *jsiiProxy_OpentelekomcloudProvider)SetDomainName(val *string) {
	_jsii_.Set(
		j,
		"domainName",
		val,
	)
}

func (j *jsiiProxy_OpentelekomcloudProvider)SetEndpointType(val *string) {
	_jsii_.Set(
		j,
		"endpointType",
		val,
	)
}

func (j *jsiiProxy_OpentelekomcloudProvider)SetInsecure(val interface{}) {
	if err := j.validateSetInsecureParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"insecure",
		val,
	)
}

func (j *jsiiProxy_OpentelekomcloudProvider)SetKey(val *string) {
	_jsii_.Set(
		j,
		"key",
		val,
	)
}

func (j *jsiiProxy_OpentelekomcloudProvider)SetMaxRetries(val *float64) {
	_jsii_.Set(
		j,
		"maxRetries",
		val,
	)
}

func (j *jsiiProxy_OpentelekomcloudProvider)SetPasscode(val *string) {
	_jsii_.Set(
		j,
		"passcode",
		val,
	)
}

func (j *jsiiProxy_OpentelekomcloudProvider)SetPassword(val *string) {
	_jsii_.Set(
		j,
		"password",
		val,
	)
}

func (j *jsiiProxy_OpentelekomcloudProvider)SetRegion(val *string) {
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_OpentelekomcloudProvider)SetSecretKey(val *string) {
	_jsii_.Set(
		j,
		"secretKey",
		val,
	)
}

func (j *jsiiProxy_OpentelekomcloudProvider)SetSecurityToken(val *string) {
	_jsii_.Set(
		j,
		"securityToken",
		val,
	)
}

func (j *jsiiProxy_OpentelekomcloudProvider)SetSwauth(val interface{}) {
	if err := j.validateSetSwauthParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"swauth",
		val,
	)
}

func (j *jsiiProxy_OpentelekomcloudProvider)SetTenantId(val *string) {
	_jsii_.Set(
		j,
		"tenantId",
		val,
	)
}

func (j *jsiiProxy_OpentelekomcloudProvider)SetTenantName(val *string) {
	_jsii_.Set(
		j,
		"tenantName",
		val,
	)
}

func (j *jsiiProxy_OpentelekomcloudProvider)SetToken(val *string) {
	_jsii_.Set(
		j,
		"token",
		val,
	)
}

func (j *jsiiProxy_OpentelekomcloudProvider)SetUserId(val *string) {
	_jsii_.Set(
		j,
		"userId",
		val,
	)
}

func (j *jsiiProxy_OpentelekomcloudProvider)SetUserName(val *string) {
	_jsii_.Set(
		j,
		"userName",
		val,
	)
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func OpentelekomcloudProvider_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateOpentelekomcloudProvider_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-opentelekomcloud.provider.OpentelekomcloudProvider",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func OpentelekomcloudProvider_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateOpentelekomcloudProvider_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-opentelekomcloud.provider.OpentelekomcloudProvider",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func OpentelekomcloudProvider_IsTerraformProvider(x interface{}) *bool {
	_init_.Initialize()

	if err := validateOpentelekomcloudProvider_IsTerraformProviderParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-opentelekomcloud.provider.OpentelekomcloudProvider",
		"isTerraformProvider",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func OpentelekomcloudProvider_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-opentelekomcloud.provider.OpentelekomcloudProvider",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (o *jsiiProxy_OpentelekomcloudProvider) AddOverride(path *string, value interface{}) {
	if err := o.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (o *jsiiProxy_OpentelekomcloudProvider) OverrideLogicalId(newLogicalId *string) {
	if err := o.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (o *jsiiProxy_OpentelekomcloudProvider) ResetAccessKey() {
	_jsii_.InvokeVoid(
		o,
		"resetAccessKey",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpentelekomcloudProvider) ResetAgencyDomainName() {
	_jsii_.InvokeVoid(
		o,
		"resetAgencyDomainName",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpentelekomcloudProvider) ResetAgencyName() {
	_jsii_.InvokeVoid(
		o,
		"resetAgencyName",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpentelekomcloudProvider) ResetAlias() {
	_jsii_.InvokeVoid(
		o,
		"resetAlias",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpentelekomcloudProvider) ResetAuthUrl() {
	_jsii_.InvokeVoid(
		o,
		"resetAuthUrl",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpentelekomcloudProvider) ResetCacertFile() {
	_jsii_.InvokeVoid(
		o,
		"resetCacertFile",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpentelekomcloudProvider) ResetCert() {
	_jsii_.InvokeVoid(
		o,
		"resetCert",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpentelekomcloudProvider) ResetCloud() {
	_jsii_.InvokeVoid(
		o,
		"resetCloud",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpentelekomcloudProvider) ResetDelegatedProject() {
	_jsii_.InvokeVoid(
		o,
		"resetDelegatedProject",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpentelekomcloudProvider) ResetDomainId() {
	_jsii_.InvokeVoid(
		o,
		"resetDomainId",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpentelekomcloudProvider) ResetDomainName() {
	_jsii_.InvokeVoid(
		o,
		"resetDomainName",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpentelekomcloudProvider) ResetEndpointType() {
	_jsii_.InvokeVoid(
		o,
		"resetEndpointType",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpentelekomcloudProvider) ResetInsecure() {
	_jsii_.InvokeVoid(
		o,
		"resetInsecure",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpentelekomcloudProvider) ResetKey() {
	_jsii_.InvokeVoid(
		o,
		"resetKey",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpentelekomcloudProvider) ResetMaxRetries() {
	_jsii_.InvokeVoid(
		o,
		"resetMaxRetries",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpentelekomcloudProvider) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		o,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpentelekomcloudProvider) ResetPasscode() {
	_jsii_.InvokeVoid(
		o,
		"resetPasscode",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpentelekomcloudProvider) ResetPassword() {
	_jsii_.InvokeVoid(
		o,
		"resetPassword",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpentelekomcloudProvider) ResetRegion() {
	_jsii_.InvokeVoid(
		o,
		"resetRegion",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpentelekomcloudProvider) ResetSecretKey() {
	_jsii_.InvokeVoid(
		o,
		"resetSecretKey",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpentelekomcloudProvider) ResetSecurityToken() {
	_jsii_.InvokeVoid(
		o,
		"resetSecurityToken",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpentelekomcloudProvider) ResetSwauth() {
	_jsii_.InvokeVoid(
		o,
		"resetSwauth",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpentelekomcloudProvider) ResetTenantId() {
	_jsii_.InvokeVoid(
		o,
		"resetTenantId",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpentelekomcloudProvider) ResetTenantName() {
	_jsii_.InvokeVoid(
		o,
		"resetTenantName",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpentelekomcloudProvider) ResetToken() {
	_jsii_.InvokeVoid(
		o,
		"resetToken",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpentelekomcloudProvider) ResetUserId() {
	_jsii_.InvokeVoid(
		o,
		"resetUserId",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpentelekomcloudProvider) ResetUserName() {
	_jsii_.InvokeVoid(
		o,
		"resetUserName",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpentelekomcloudProvider) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		o,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpentelekomcloudProvider) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		o,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpentelekomcloudProvider) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpentelekomcloudProvider) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		o,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

