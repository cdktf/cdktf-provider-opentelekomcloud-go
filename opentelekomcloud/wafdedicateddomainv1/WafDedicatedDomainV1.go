// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package wafdedicateddomainv1

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v10/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v10/wafdedicateddomainv1/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.13/docs/resources/waf_dedicated_domain_v1 opentelekomcloud_waf_dedicated_domain_v1}.
type WafDedicatedDomainV1 interface {
	cdktf.TerraformResource
	AccessStatus() *float64
	AlarmPage() cdktf.StringMap
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	CertificateId() *string
	SetCertificateId(val *string)
	CertificateIdInput() *string
	CertificateName() *string
	Cipher() *string
	SetCipher(val *string)
	CipherInput() *string
	ComplianceCertification() cdktf.BooleanMap
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	CreatedAt() *float64
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Domain() *string
	SetDomain(val *string)
	DomainInput() *string
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	KeepPolicy() interface{}
	SetKeepPolicy(val interface{})
	KeepPolicyInput() interface{}
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	// The tree node.
	Node() constructs.Node
	Pci3Ds() interface{}
	SetPci3Ds(val interface{})
	Pci3DsInput() interface{}
	PciDss() interface{}
	SetPciDss(val interface{})
	PciDssInput() interface{}
	PolicyId() *string
	SetPolicyId(val *string)
	PolicyIdInput() *string
	ProtectStatus() *float64
	SetProtectStatus(val *float64)
	ProtectStatusInput() *float64
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
	Region() *string
	SetRegion(val *string)
	RegionInput() *string
	Server() WafDedicatedDomainV1ServerList
	ServerInput() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Tls() *string
	SetTls(val *string)
	TlsInput() *string
	TrafficIdentifier() cdktf.StringMap
	// Adds a user defined moveTarget string to this resource to be later used in .moveTo(moveTarget) to resolve the location of the move.
	// Experimental.
	AddMoveTarget(moveTarget *string)
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
	ImportFrom(id *string, provider cdktf.TerraformProvider)
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	// Moves this resource to the target resource given by moveTarget.
	// Experimental.
	MoveTo(moveTarget *string, index interface{})
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	PutServer(value interface{})
	ResetCertificateId()
	ResetCipher()
	ResetId()
	ResetKeepPolicy()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPci3Ds()
	ResetPciDss()
	ResetPolicyId()
	ResetProtectStatus()
	ResetProxy()
	ResetRegion()
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

// The jsii proxy struct for WafDedicatedDomainV1
type jsiiProxy_WafDedicatedDomainV1 struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_WafDedicatedDomainV1) AccessStatus() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"accessStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafDedicatedDomainV1) AlarmPage() cdktf.StringMap {
	var returns cdktf.StringMap
	_jsii_.Get(
		j,
		"alarmPage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafDedicatedDomainV1) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafDedicatedDomainV1) CertificateId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificateId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafDedicatedDomainV1) CertificateIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificateIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafDedicatedDomainV1) CertificateName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificateName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafDedicatedDomainV1) Cipher() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cipher",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafDedicatedDomainV1) CipherInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cipherInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafDedicatedDomainV1) ComplianceCertification() cdktf.BooleanMap {
	var returns cdktf.BooleanMap
	_jsii_.Get(
		j,
		"complianceCertification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafDedicatedDomainV1) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafDedicatedDomainV1) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafDedicatedDomainV1) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafDedicatedDomainV1) CreatedAt() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"createdAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafDedicatedDomainV1) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafDedicatedDomainV1) Domain() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafDedicatedDomainV1) DomainInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafDedicatedDomainV1) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafDedicatedDomainV1) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafDedicatedDomainV1) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafDedicatedDomainV1) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafDedicatedDomainV1) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafDedicatedDomainV1) KeepPolicy() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"keepPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafDedicatedDomainV1) KeepPolicyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"keepPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafDedicatedDomainV1) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafDedicatedDomainV1) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafDedicatedDomainV1) Pci3Ds() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"pci3Ds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafDedicatedDomainV1) Pci3DsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"pci3DsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafDedicatedDomainV1) PciDss() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"pciDss",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafDedicatedDomainV1) PciDssInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"pciDssInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafDedicatedDomainV1) PolicyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafDedicatedDomainV1) PolicyIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policyIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafDedicatedDomainV1) ProtectStatus() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"protectStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafDedicatedDomainV1) ProtectStatusInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"protectStatusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafDedicatedDomainV1) Protocol() *string {
	var returns *string
	_jsii_.Get(
		j,
		"protocol",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafDedicatedDomainV1) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafDedicatedDomainV1) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafDedicatedDomainV1) Proxy() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"proxy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafDedicatedDomainV1) ProxyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"proxyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafDedicatedDomainV1) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafDedicatedDomainV1) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafDedicatedDomainV1) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafDedicatedDomainV1) Server() WafDedicatedDomainV1ServerList {
	var returns WafDedicatedDomainV1ServerList
	_jsii_.Get(
		j,
		"server",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafDedicatedDomainV1) ServerInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"serverInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafDedicatedDomainV1) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafDedicatedDomainV1) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafDedicatedDomainV1) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafDedicatedDomainV1) Tls() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tls",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafDedicatedDomainV1) TlsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tlsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafDedicatedDomainV1) TrafficIdentifier() cdktf.StringMap {
	var returns cdktf.StringMap
	_jsii_.Get(
		j,
		"trafficIdentifier",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.13/docs/resources/waf_dedicated_domain_v1 opentelekomcloud_waf_dedicated_domain_v1} Resource.
func NewWafDedicatedDomainV1(scope constructs.Construct, id *string, config *WafDedicatedDomainV1Config) WafDedicatedDomainV1 {
	_init_.Initialize()

	if err := validateNewWafDedicatedDomainV1Parameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_WafDedicatedDomainV1{}

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.wafDedicatedDomainV1.WafDedicatedDomainV1",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.13/docs/resources/waf_dedicated_domain_v1 opentelekomcloud_waf_dedicated_domain_v1} Resource.
func NewWafDedicatedDomainV1_Override(w WafDedicatedDomainV1, scope constructs.Construct, id *string, config *WafDedicatedDomainV1Config) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.wafDedicatedDomainV1.WafDedicatedDomainV1",
		[]interface{}{scope, id, config},
		w,
	)
}

func (j *jsiiProxy_WafDedicatedDomainV1)SetCertificateId(val *string) {
	if err := j.validateSetCertificateIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"certificateId",
		val,
	)
}

func (j *jsiiProxy_WafDedicatedDomainV1)SetCipher(val *string) {
	if err := j.validateSetCipherParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cipher",
		val,
	)
}

func (j *jsiiProxy_WafDedicatedDomainV1)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_WafDedicatedDomainV1)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_WafDedicatedDomainV1)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_WafDedicatedDomainV1)SetDomain(val *string) {
	if err := j.validateSetDomainParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"domain",
		val,
	)
}

func (j *jsiiProxy_WafDedicatedDomainV1)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_WafDedicatedDomainV1)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_WafDedicatedDomainV1)SetKeepPolicy(val interface{}) {
	if err := j.validateSetKeepPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"keepPolicy",
		val,
	)
}

func (j *jsiiProxy_WafDedicatedDomainV1)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_WafDedicatedDomainV1)SetPci3Ds(val interface{}) {
	if err := j.validateSetPci3DsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pci3Ds",
		val,
	)
}

func (j *jsiiProxy_WafDedicatedDomainV1)SetPciDss(val interface{}) {
	if err := j.validateSetPciDssParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pciDss",
		val,
	)
}

func (j *jsiiProxy_WafDedicatedDomainV1)SetPolicyId(val *string) {
	if err := j.validateSetPolicyIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"policyId",
		val,
	)
}

func (j *jsiiProxy_WafDedicatedDomainV1)SetProtectStatus(val *float64) {
	if err := j.validateSetProtectStatusParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"protectStatus",
		val,
	)
}

func (j *jsiiProxy_WafDedicatedDomainV1)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_WafDedicatedDomainV1)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_WafDedicatedDomainV1)SetProxy(val interface{}) {
	if err := j.validateSetProxyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"proxy",
		val,
	)
}

func (j *jsiiProxy_WafDedicatedDomainV1)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_WafDedicatedDomainV1)SetTls(val *string) {
	if err := j.validateSetTlsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tls",
		val,
	)
}

// Generates CDKTF code for importing a WafDedicatedDomainV1 resource upon running "cdktf plan <stack-name>".
func WafDedicatedDomainV1_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateWafDedicatedDomainV1_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-opentelekomcloud.wafDedicatedDomainV1.WafDedicatedDomainV1",
		"generateConfigForImport",
		[]interface{}{scope, importToId, importFromId, provider},
		&returns,
	)

	return returns
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
func WafDedicatedDomainV1_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateWafDedicatedDomainV1_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-opentelekomcloud.wafDedicatedDomainV1.WafDedicatedDomainV1",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func WafDedicatedDomainV1_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateWafDedicatedDomainV1_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-opentelekomcloud.wafDedicatedDomainV1.WafDedicatedDomainV1",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func WafDedicatedDomainV1_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateWafDedicatedDomainV1_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-opentelekomcloud.wafDedicatedDomainV1.WafDedicatedDomainV1",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func WafDedicatedDomainV1_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-opentelekomcloud.wafDedicatedDomainV1.WafDedicatedDomainV1",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (w *jsiiProxy_WafDedicatedDomainV1) AddMoveTarget(moveTarget *string) {
	if err := w.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (w *jsiiProxy_WafDedicatedDomainV1) AddOverride(path *string, value interface{}) {
	if err := w.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (w *jsiiProxy_WafDedicatedDomainV1) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (w *jsiiProxy_WafDedicatedDomainV1) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (w *jsiiProxy_WafDedicatedDomainV1) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (w *jsiiProxy_WafDedicatedDomainV1) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (w *jsiiProxy_WafDedicatedDomainV1) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (w *jsiiProxy_WafDedicatedDomainV1) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (w *jsiiProxy_WafDedicatedDomainV1) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (w *jsiiProxy_WafDedicatedDomainV1) GetStringAttribute(terraformAttribute *string) *string {
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

func (w *jsiiProxy_WafDedicatedDomainV1) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (w *jsiiProxy_WafDedicatedDomainV1) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := w.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (w *jsiiProxy_WafDedicatedDomainV1) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (w *jsiiProxy_WafDedicatedDomainV1) MoveTo(moveTarget *string, index interface{}) {
	if err := w.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (w *jsiiProxy_WafDedicatedDomainV1) OverrideLogicalId(newLogicalId *string) {
	if err := w.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (w *jsiiProxy_WafDedicatedDomainV1) PutServer(value interface{}) {
	if err := w.validatePutServerParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putServer",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WafDedicatedDomainV1) ResetCertificateId() {
	_jsii_.InvokeVoid(
		w,
		"resetCertificateId",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WafDedicatedDomainV1) ResetCipher() {
	_jsii_.InvokeVoid(
		w,
		"resetCipher",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WafDedicatedDomainV1) ResetId() {
	_jsii_.InvokeVoid(
		w,
		"resetId",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WafDedicatedDomainV1) ResetKeepPolicy() {
	_jsii_.InvokeVoid(
		w,
		"resetKeepPolicy",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WafDedicatedDomainV1) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		w,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WafDedicatedDomainV1) ResetPci3Ds() {
	_jsii_.InvokeVoid(
		w,
		"resetPci3Ds",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WafDedicatedDomainV1) ResetPciDss() {
	_jsii_.InvokeVoid(
		w,
		"resetPciDss",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WafDedicatedDomainV1) ResetPolicyId() {
	_jsii_.InvokeVoid(
		w,
		"resetPolicyId",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WafDedicatedDomainV1) ResetProtectStatus() {
	_jsii_.InvokeVoid(
		w,
		"resetProtectStatus",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WafDedicatedDomainV1) ResetProxy() {
	_jsii_.InvokeVoid(
		w,
		"resetProxy",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WafDedicatedDomainV1) ResetRegion() {
	_jsii_.InvokeVoid(
		w,
		"resetRegion",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WafDedicatedDomainV1) ResetTls() {
	_jsii_.InvokeVoid(
		w,
		"resetTls",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WafDedicatedDomainV1) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		w,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WafDedicatedDomainV1) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		w,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WafDedicatedDomainV1) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WafDedicatedDomainV1) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		w,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

