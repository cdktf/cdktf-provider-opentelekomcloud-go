// Prebuilt opentelekomcloud Provider for Terraform CDK (cdktf)
package opentelekomcloud

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v2/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://www.terraform.io/docs/providers/opentelekomcloud/r/vpnaas_service_v2 opentelekomcloud_vpnaas_service_v2}.
type VpnaasServiceV2 interface {
	cdktf.TerraformResource
	AdminStateUp() interface{}
	SetAdminStateUp(val interface{})
	AdminStateUpInput() interface{}
	// Experimental.
	CdktfStack() cdktf.TerraformStack
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
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	ExternalV4Ip() *string
	ExternalV6Ip() *string
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
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	// Experimental.
	RawOverrides() interface{}
	Region() *string
	SetRegion(val *string)
	RegionInput() *string
	RouterId() *string
	SetRouterId(val *string)
	RouterIdInput() *string
	Status() *string
	SubnetId() *string
	SetSubnetId(val *string)
	SubnetIdInput() *string
	TenantId() *string
	SetTenantId(val *string)
	TenantIdInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() VpnaasServiceV2TimeoutsOutputReference
	TimeoutsInput() interface{}
	ValueSpecs() *map[string]*string
	SetValueSpecs(val *map[string]*string)
	ValueSpecsInput() *map[string]*string
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
	PutTimeouts(value *VpnaasServiceV2Timeouts)
	ResetAdminStateUp()
	ResetDescription()
	ResetId()
	ResetName()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetRegion()
	ResetSubnetId()
	ResetTenantId()
	ResetTimeouts()
	ResetValueSpecs()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for VpnaasServiceV2
type jsiiProxy_VpnaasServiceV2 struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_VpnaasServiceV2) AdminStateUp() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"adminStateUp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnaasServiceV2) AdminStateUpInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"adminStateUpInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnaasServiceV2) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnaasServiceV2) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnaasServiceV2) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnaasServiceV2) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnaasServiceV2) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnaasServiceV2) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnaasServiceV2) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnaasServiceV2) ExternalV4Ip() *string {
	var returns *string
	_jsii_.Get(
		j,
		"externalV4Ip",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnaasServiceV2) ExternalV6Ip() *string {
	var returns *string
	_jsii_.Get(
		j,
		"externalV6Ip",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnaasServiceV2) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnaasServiceV2) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnaasServiceV2) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnaasServiceV2) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnaasServiceV2) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnaasServiceV2) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnaasServiceV2) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnaasServiceV2) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnaasServiceV2) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnaasServiceV2) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnaasServiceV2) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnaasServiceV2) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnaasServiceV2) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnaasServiceV2) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnaasServiceV2) RouterId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routerId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnaasServiceV2) RouterIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routerIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnaasServiceV2) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnaasServiceV2) SubnetId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnaasServiceV2) SubnetIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnaasServiceV2) TenantId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tenantId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnaasServiceV2) TenantIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tenantIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnaasServiceV2) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnaasServiceV2) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnaasServiceV2) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnaasServiceV2) Timeouts() VpnaasServiceV2TimeoutsOutputReference {
	var returns VpnaasServiceV2TimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnaasServiceV2) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnaasServiceV2) ValueSpecs() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"valueSpecs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnaasServiceV2) ValueSpecsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"valueSpecsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://www.terraform.io/docs/providers/opentelekomcloud/r/vpnaas_service_v2 opentelekomcloud_vpnaas_service_v2} Resource.
func NewVpnaasServiceV2(scope constructs.Construct, id *string, config *VpnaasServiceV2Config) VpnaasServiceV2 {
	_init_.Initialize()

	j := jsiiProxy_VpnaasServiceV2{}

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.VpnaasServiceV2",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/opentelekomcloud/r/vpnaas_service_v2 opentelekomcloud_vpnaas_service_v2} Resource.
func NewVpnaasServiceV2_Override(v VpnaasServiceV2, scope constructs.Construct, id *string, config *VpnaasServiceV2Config) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.VpnaasServiceV2",
		[]interface{}{scope, id, config},
		v,
	)
}

func (j *jsiiProxy_VpnaasServiceV2) SetAdminStateUp(val interface{}) {
	_jsii_.Set(
		j,
		"adminStateUp",
		val,
	)
}

func (j *jsiiProxy_VpnaasServiceV2) SetConnection(val interface{}) {
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_VpnaasServiceV2) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_VpnaasServiceV2) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_VpnaasServiceV2) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_VpnaasServiceV2) SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_VpnaasServiceV2) SetId(val *string) {
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_VpnaasServiceV2) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_VpnaasServiceV2) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_VpnaasServiceV2) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_VpnaasServiceV2) SetProvisioners(val *[]interface{}) {
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_VpnaasServiceV2) SetRegion(val *string) {
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_VpnaasServiceV2) SetRouterId(val *string) {
	_jsii_.Set(
		j,
		"routerId",
		val,
	)
}

func (j *jsiiProxy_VpnaasServiceV2) SetSubnetId(val *string) {
	_jsii_.Set(
		j,
		"subnetId",
		val,
	)
}

func (j *jsiiProxy_VpnaasServiceV2) SetTenantId(val *string) {
	_jsii_.Set(
		j,
		"tenantId",
		val,
	)
}

func (j *jsiiProxy_VpnaasServiceV2) SetValueSpecs(val *map[string]*string) {
	_jsii_.Set(
		j,
		"valueSpecs",
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
func VpnaasServiceV2_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-opentelekomcloud.VpnaasServiceV2",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func VpnaasServiceV2_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-opentelekomcloud.VpnaasServiceV2",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (v *jsiiProxy_VpnaasServiceV2) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		v,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (v *jsiiProxy_VpnaasServiceV2) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		v,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpnaasServiceV2) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpnaasServiceV2) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		v,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpnaasServiceV2) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		v,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpnaasServiceV2) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		v,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpnaasServiceV2) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		v,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpnaasServiceV2) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		v,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpnaasServiceV2) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpnaasServiceV2) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		v,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpnaasServiceV2) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpnaasServiceV2) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		v,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (v *jsiiProxy_VpnaasServiceV2) PutTimeouts(value *VpnaasServiceV2Timeouts) {
	_jsii_.InvokeVoid(
		v,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VpnaasServiceV2) ResetAdminStateUp() {
	_jsii_.InvokeVoid(
		v,
		"resetAdminStateUp",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpnaasServiceV2) ResetDescription() {
	_jsii_.InvokeVoid(
		v,
		"resetDescription",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpnaasServiceV2) ResetId() {
	_jsii_.InvokeVoid(
		v,
		"resetId",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpnaasServiceV2) ResetName() {
	_jsii_.InvokeVoid(
		v,
		"resetName",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpnaasServiceV2) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		v,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpnaasServiceV2) ResetRegion() {
	_jsii_.InvokeVoid(
		v,
		"resetRegion",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpnaasServiceV2) ResetSubnetId() {
	_jsii_.InvokeVoid(
		v,
		"resetSubnetId",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpnaasServiceV2) ResetTenantId() {
	_jsii_.InvokeVoid(
		v,
		"resetTenantId",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpnaasServiceV2) ResetTimeouts() {
	_jsii_.InvokeVoid(
		v,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpnaasServiceV2) ResetValueSpecs() {
	_jsii_.InvokeVoid(
		v,
		"resetValueSpecs",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpnaasServiceV2) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		v,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpnaasServiceV2) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		v,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpnaasServiceV2) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpnaasServiceV2) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		v,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}
