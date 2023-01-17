package vpcependpointv1

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v5/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v5/vpcependpointv1/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://www.terraform.io/docs/providers/opentelekomcloud/r/vpcep_endpoint_v1 opentelekomcloud_vpcep_endpoint_v1}.
type VpcepEndpointV1 interface {
	cdktf.TerraformResource
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
	DnsNames() *[]*string
	EnableDns() interface{}
	SetEnableDns(val interface{})
	EnableDnsInput() interface{}
	EnableWhitelist() interface{}
	SetEnableWhitelist(val interface{})
	EnableWhitelistInput() interface{}
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
	MarkerId() *float64
	// The tree node.
	Node() constructs.Node
	PortIp() *string
	SetPortIp(val *string)
	PortIpInput() *string
	ProjectId() *string
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
	RouteTables() *[]*string
	SetRouteTables(val *[]*string)
	RouteTablesInput() *[]*string
	ServiceId() *string
	SetServiceId(val *string)
	ServiceIdInput() *string
	ServiceName() *string
	ServiceType() *string
	SubnetId() *string
	SetSubnetId(val *string)
	SubnetIdInput() *string
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsInput() *map[string]*string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() VpcepEndpointV1TimeoutsOutputReference
	TimeoutsInput() interface{}
	VpcId() *string
	SetVpcId(val *string)
	VpcIdInput() *string
	Whitelist() *[]*string
	SetWhitelist(val *[]*string)
	WhitelistInput() *[]*string
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
	PutTimeouts(value *VpcepEndpointV1Timeouts)
	ResetEnableDns()
	ResetEnableWhitelist()
	ResetId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPortIp()
	ResetRouteTables()
	ResetSubnetId()
	ResetTags()
	ResetTimeouts()
	ResetWhitelist()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for VpcepEndpointV1
type jsiiProxy_VpcepEndpointV1 struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_VpcepEndpointV1) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcepEndpointV1) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcepEndpointV1) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcepEndpointV1) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcepEndpointV1) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcepEndpointV1) DnsNames() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dnsNames",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcepEndpointV1) EnableDns() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableDns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcepEndpointV1) EnableDnsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableDnsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcepEndpointV1) EnableWhitelist() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableWhitelist",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcepEndpointV1) EnableWhitelistInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableWhitelistInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcepEndpointV1) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcepEndpointV1) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcepEndpointV1) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcepEndpointV1) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcepEndpointV1) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcepEndpointV1) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcepEndpointV1) MarkerId() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"markerId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcepEndpointV1) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcepEndpointV1) PortIp() *string {
	var returns *string
	_jsii_.Get(
		j,
		"portIp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcepEndpointV1) PortIpInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"portIpInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcepEndpointV1) ProjectId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcepEndpointV1) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcepEndpointV1) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcepEndpointV1) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcepEndpointV1) RouteTables() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"routeTables",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcepEndpointV1) RouteTablesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"routeTablesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcepEndpointV1) ServiceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcepEndpointV1) ServiceIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcepEndpointV1) ServiceName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcepEndpointV1) ServiceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcepEndpointV1) SubnetId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcepEndpointV1) SubnetIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcepEndpointV1) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcepEndpointV1) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcepEndpointV1) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcepEndpointV1) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcepEndpointV1) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcepEndpointV1) Timeouts() VpcepEndpointV1TimeoutsOutputReference {
	var returns VpcepEndpointV1TimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcepEndpointV1) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcepEndpointV1) VpcId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcepEndpointV1) VpcIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcepEndpointV1) Whitelist() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"whitelist",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcepEndpointV1) WhitelistInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"whitelistInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://www.terraform.io/docs/providers/opentelekomcloud/r/vpcep_endpoint_v1 opentelekomcloud_vpcep_endpoint_v1} Resource.
func NewVpcepEndpointV1(scope constructs.Construct, id *string, config *VpcepEndpointV1Config) VpcepEndpointV1 {
	_init_.Initialize()

	if err := validateNewVpcepEndpointV1Parameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_VpcepEndpointV1{}

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.vpcepEndpointV1.VpcepEndpointV1",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/opentelekomcloud/r/vpcep_endpoint_v1 opentelekomcloud_vpcep_endpoint_v1} Resource.
func NewVpcepEndpointV1_Override(v VpcepEndpointV1, scope constructs.Construct, id *string, config *VpcepEndpointV1Config) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.vpcepEndpointV1.VpcepEndpointV1",
		[]interface{}{scope, id, config},
		v,
	)
}

func (j *jsiiProxy_VpcepEndpointV1)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_VpcepEndpointV1)SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_VpcepEndpointV1)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_VpcepEndpointV1)SetEnableDns(val interface{}) {
	if err := j.validateSetEnableDnsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableDns",
		val,
	)
}

func (j *jsiiProxy_VpcepEndpointV1)SetEnableWhitelist(val interface{}) {
	if err := j.validateSetEnableWhitelistParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableWhitelist",
		val,
	)
}

func (j *jsiiProxy_VpcepEndpointV1)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_VpcepEndpointV1)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_VpcepEndpointV1)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_VpcepEndpointV1)SetPortIp(val *string) {
	if err := j.validateSetPortIpParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"portIp",
		val,
	)
}

func (j *jsiiProxy_VpcepEndpointV1)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_VpcepEndpointV1)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_VpcepEndpointV1)SetRouteTables(val *[]*string) {
	if err := j.validateSetRouteTablesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"routeTables",
		val,
	)
}

func (j *jsiiProxy_VpcepEndpointV1)SetServiceId(val *string) {
	if err := j.validateSetServiceIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serviceId",
		val,
	)
}

func (j *jsiiProxy_VpcepEndpointV1)SetSubnetId(val *string) {
	if err := j.validateSetSubnetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"subnetId",
		val,
	)
}

func (j *jsiiProxy_VpcepEndpointV1)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_VpcepEndpointV1)SetVpcId(val *string) {
	if err := j.validateSetVpcIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vpcId",
		val,
	)
}

func (j *jsiiProxy_VpcepEndpointV1)SetWhitelist(val *[]*string) {
	if err := j.validateSetWhitelistParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"whitelist",
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
func VpcepEndpointV1_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateVpcepEndpointV1_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-opentelekomcloud.vpcepEndpointV1.VpcepEndpointV1",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func VpcepEndpointV1_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateVpcepEndpointV1_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-opentelekomcloud.vpcepEndpointV1.VpcepEndpointV1",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func VpcepEndpointV1_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateVpcepEndpointV1_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-opentelekomcloud.vpcepEndpointV1.VpcepEndpointV1",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func VpcepEndpointV1_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-opentelekomcloud.vpcepEndpointV1.VpcepEndpointV1",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (v *jsiiProxy_VpcepEndpointV1) AddOverride(path *string, value interface{}) {
	if err := v.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (v *jsiiProxy_VpcepEndpointV1) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := v.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		v,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpcepEndpointV1) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := v.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpcepEndpointV1) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := v.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		v,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpcepEndpointV1) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := v.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		v,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpcepEndpointV1) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := v.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		v,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpcepEndpointV1) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := v.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		v,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpcepEndpointV1) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := v.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		v,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpcepEndpointV1) GetStringAttribute(terraformAttribute *string) *string {
	if err := v.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		v,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpcepEndpointV1) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := v.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		v,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpcepEndpointV1) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := v.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpcepEndpointV1) OverrideLogicalId(newLogicalId *string) {
	if err := v.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (v *jsiiProxy_VpcepEndpointV1) PutTimeouts(value *VpcepEndpointV1Timeouts) {
	if err := v.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VpcepEndpointV1) ResetEnableDns() {
	_jsii_.InvokeVoid(
		v,
		"resetEnableDns",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpcepEndpointV1) ResetEnableWhitelist() {
	_jsii_.InvokeVoid(
		v,
		"resetEnableWhitelist",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpcepEndpointV1) ResetId() {
	_jsii_.InvokeVoid(
		v,
		"resetId",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpcepEndpointV1) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		v,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpcepEndpointV1) ResetPortIp() {
	_jsii_.InvokeVoid(
		v,
		"resetPortIp",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpcepEndpointV1) ResetRouteTables() {
	_jsii_.InvokeVoid(
		v,
		"resetRouteTables",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpcepEndpointV1) ResetSubnetId() {
	_jsii_.InvokeVoid(
		v,
		"resetSubnetId",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpcepEndpointV1) ResetTags() {
	_jsii_.InvokeVoid(
		v,
		"resetTags",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpcepEndpointV1) ResetTimeouts() {
	_jsii_.InvokeVoid(
		v,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpcepEndpointV1) ResetWhitelist() {
	_jsii_.InvokeVoid(
		v,
		"resetWhitelist",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpcepEndpointV1) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		v,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpcepEndpointV1) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		v,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpcepEndpointV1) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpcepEndpointV1) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		v,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

