package wafdomainv1

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v3/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v3/wafdomainv1/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://www.terraform.io/docs/providers/opentelekomcloud/r/waf_domain_v1 opentelekomcloud_waf_domain_v1}.
type WafDomainV1 interface {
	cdktf.TerraformResource
	AccessCode() *string
	AccessStatus() *float64
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	CertificateId() *string
	SetCertificateId(val *string)
	CertificateIdInput() *string
	Cipher() *string
	SetCipher(val *string)
	CipherInput() *string
	Cname() *string
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() *float64
	// Experimental.
	SetCount(val *float64)
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Hostname() *string
	SetHostname(val *string)
	HostnameInput() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	// The tree node.
	Node() constructs.Node
	PolicyId() *string
	SetPolicyId(val *string)
	PolicyIdInput() *string
	ProtectStatus() *float64
	Protocol() *string
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	Proxy() interface{}
	SetProxy(val interface{})
	ProxyInput() interface{}
	// Experimental.
	RawOverrides() interface{}
	Server() WafDomainV1ServerList
	ServerInput() interface{}
	SipHeaderList() *[]*string
	SetSipHeaderList(val *[]*string)
	SipHeaderListInput() *[]*string
	SipHeaderName() *string
	SetSipHeaderName(val *string)
	SipHeaderNameInput() *string
	SubDomain() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() WafDomainV1TimeoutsOutputReference
	TimeoutsInput() interface{}
	Tls() *string
	SetTls(val *string)
	TlsInput() *string
	TxtCode() *string
	// Experimental.
	AddOverride(path *string, value interface{})
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	// Experimental.
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	// Experimental.
	GetListAttribute(terraformAttribute *string) *[]*string
	// Experimental.
	GetNumberAttribute(terraformAttribute *string) *float64
	// Experimental.
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	// Experimental.
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	// Experimental.
	GetStringAttribute(terraformAttribute *string) *string
	// Experimental.
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	PutServer(value interface{})
	PutTimeouts(value *WafDomainV1Timeouts)
	ResetCertificateId()
	ResetCipher()
	ResetId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPolicyId()
	ResetSipHeaderList()
	ResetSipHeaderName()
	ResetTimeouts()
	ResetTls()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for WafDomainV1
type jsiiProxy_WafDomainV1 struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_WafDomainV1) AccessCode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessCode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafDomainV1) AccessStatus() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"accessStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafDomainV1) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafDomainV1) CertificateId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificateId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafDomainV1) CertificateIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificateIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafDomainV1) Cipher() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cipher",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafDomainV1) CipherInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cipherInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafDomainV1) Cname() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cname",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafDomainV1) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafDomainV1) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafDomainV1) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafDomainV1) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafDomainV1) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafDomainV1) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafDomainV1) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafDomainV1) Hostname() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostname",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafDomainV1) HostnameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostnameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafDomainV1) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafDomainV1) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafDomainV1) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafDomainV1) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafDomainV1) PolicyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafDomainV1) PolicyIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policyIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafDomainV1) ProtectStatus() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"protectStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafDomainV1) Protocol() *string {
	var returns *string
	_jsii_.Get(
		j,
		"protocol",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafDomainV1) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafDomainV1) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafDomainV1) Proxy() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"proxy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafDomainV1) ProxyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"proxyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafDomainV1) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafDomainV1) Server() WafDomainV1ServerList {
	var returns WafDomainV1ServerList
	_jsii_.Get(
		j,
		"server",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafDomainV1) ServerInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"serverInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafDomainV1) SipHeaderList() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"sipHeaderList",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafDomainV1) SipHeaderListInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"sipHeaderListInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafDomainV1) SipHeaderName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sipHeaderName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafDomainV1) SipHeaderNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sipHeaderNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafDomainV1) SubDomain() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subDomain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafDomainV1) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafDomainV1) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafDomainV1) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafDomainV1) Timeouts() WafDomainV1TimeoutsOutputReference {
	var returns WafDomainV1TimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafDomainV1) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafDomainV1) Tls() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tls",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafDomainV1) TlsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tlsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafDomainV1) TxtCode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"txtCode",
		&returns,
	)
	return returns
}


// Create a new {@link https://www.terraform.io/docs/providers/opentelekomcloud/r/waf_domain_v1 opentelekomcloud_waf_domain_v1} Resource.
func NewWafDomainV1(scope constructs.Construct, id *string, config *WafDomainV1Config) WafDomainV1 {
	_init_.Initialize()

	if err := validateNewWafDomainV1Parameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_WafDomainV1{}

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.wafDomainV1.WafDomainV1",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/opentelekomcloud/r/waf_domain_v1 opentelekomcloud_waf_domain_v1} Resource.
func NewWafDomainV1_Override(w WafDomainV1, scope constructs.Construct, id *string, config *WafDomainV1Config) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.wafDomainV1.WafDomainV1",
		[]interface{}{scope, id, config},
		w,
	)
}

func (j *jsiiProxy_WafDomainV1)SetCertificateId(val *string) {
	if err := j.validateSetCertificateIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"certificateId",
		val,
	)
}

func (j *jsiiProxy_WafDomainV1)SetCipher(val *string) {
	if err := j.validateSetCipherParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cipher",
		val,
	)
}

func (j *jsiiProxy_WafDomainV1)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_WafDomainV1)SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_WafDomainV1)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_WafDomainV1)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_WafDomainV1)SetHostname(val *string) {
	if err := j.validateSetHostnameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hostname",
		val,
	)
}

func (j *jsiiProxy_WafDomainV1)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_WafDomainV1)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_WafDomainV1)SetPolicyId(val *string) {
	if err := j.validateSetPolicyIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"policyId",
		val,
	)
}

func (j *jsiiProxy_WafDomainV1)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_WafDomainV1)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_WafDomainV1)SetProxy(val interface{}) {
	if err := j.validateSetProxyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"proxy",
		val,
	)
}

func (j *jsiiProxy_WafDomainV1)SetSipHeaderList(val *[]*string) {
	if err := j.validateSetSipHeaderListParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sipHeaderList",
		val,
	)
}

func (j *jsiiProxy_WafDomainV1)SetSipHeaderName(val *string) {
	if err := j.validateSetSipHeaderNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sipHeaderName",
		val,
	)
}

func (j *jsiiProxy_WafDomainV1)SetTls(val *string) {
	if err := j.validateSetTlsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tls",
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
func WafDomainV1_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateWafDomainV1_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-opentelekomcloud.wafDomainV1.WafDomainV1",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func WafDomainV1_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-opentelekomcloud.wafDomainV1.WafDomainV1",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (w *jsiiProxy_WafDomainV1) AddOverride(path *string, value interface{}) {
	if err := w.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (w *jsiiProxy_WafDomainV1) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := w.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		w,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WafDomainV1) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := w.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WafDomainV1) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := w.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		w,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WafDomainV1) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := w.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		w,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WafDomainV1) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := w.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		w,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WafDomainV1) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := w.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		w,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WafDomainV1) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := w.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		w,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WafDomainV1) GetStringAttribute(terraformAttribute *string) *string {
	if err := w.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		w,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WafDomainV1) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := w.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		w,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WafDomainV1) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := w.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WafDomainV1) OverrideLogicalId(newLogicalId *string) {
	if err := w.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (w *jsiiProxy_WafDomainV1) PutServer(value interface{}) {
	if err := w.validatePutServerParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putServer",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WafDomainV1) PutTimeouts(value *WafDomainV1Timeouts) {
	if err := w.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WafDomainV1) ResetCertificateId() {
	_jsii_.InvokeVoid(
		w,
		"resetCertificateId",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WafDomainV1) ResetCipher() {
	_jsii_.InvokeVoid(
		w,
		"resetCipher",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WafDomainV1) ResetId() {
	_jsii_.InvokeVoid(
		w,
		"resetId",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WafDomainV1) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		w,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WafDomainV1) ResetPolicyId() {
	_jsii_.InvokeVoid(
		w,
		"resetPolicyId",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WafDomainV1) ResetSipHeaderList() {
	_jsii_.InvokeVoid(
		w,
		"resetSipHeaderList",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WafDomainV1) ResetSipHeaderName() {
	_jsii_.InvokeVoid(
		w,
		"resetSipHeaderName",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WafDomainV1) ResetTimeouts() {
	_jsii_.InvokeVoid(
		w,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WafDomainV1) ResetTls() {
	_jsii_.InvokeVoid(
		w,
		"resetTls",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WafDomainV1) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		w,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WafDomainV1) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		w,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WafDomainV1) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WafDomainV1) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		w,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

