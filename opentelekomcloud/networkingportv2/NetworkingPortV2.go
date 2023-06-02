package networkingportv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v6/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v6/networkingportv2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.0/docs/resources/networking_port_v2 opentelekomcloud_networking_port_v2}.
type NetworkingPortV2 interface {
	cdktf.TerraformResource
	AdminStateUp() interface{}
	SetAdminStateUp(val interface{})
	AdminStateUpInput() interface{}
	AllFixedIps() *[]*string
	AllowedAddressPairs() NetworkingPortV2AllowedAddressPairsList
	AllowedAddressPairsInput() interface{}
	// Experimental.
	CdktfStack() cdktf.TerraformStack
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
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	DeviceId() *string
	SetDeviceId(val *string)
	DeviceIdInput() *string
	DeviceOwner() *string
	SetDeviceOwner(val *string)
	DeviceOwnerInput() *string
	FixedIp() NetworkingPortV2FixedIpList
	FixedIpInput() interface{}
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
	MacAddress() *string
	SetMacAddress(val *string)
	MacAddressInput() *string
	Name() *string
	SetName(val *string)
	NameInput() *string
	NetworkId() *string
	SetNetworkId(val *string)
	NetworkIdInput() *string
	// The tree node.
	Node() constructs.Node
	NoSecurityGroups() interface{}
	SetNoSecurityGroups(val interface{})
	NoSecurityGroupsInput() interface{}
	PortSecurityEnabled() interface{}
	SetPortSecurityEnabled(val interface{})
	PortSecurityEnabledInput() interface{}
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
	SecurityGroupIds() *[]*string
	SetSecurityGroupIds(val *[]*string)
	SecurityGroupIdsInput() *[]*string
	TenantId() *string
	SetTenantId(val *string)
	TenantIdInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() NetworkingPortV2TimeoutsOutputReference
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
	PutAllowedAddressPairs(value interface{})
	PutFixedIp(value interface{})
	PutTimeouts(value *NetworkingPortV2Timeouts)
	ResetAdminStateUp()
	ResetAllowedAddressPairs()
	ResetDeviceId()
	ResetDeviceOwner()
	ResetFixedIp()
	ResetId()
	ResetMacAddress()
	ResetName()
	ResetNoSecurityGroups()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPortSecurityEnabled()
	ResetRegion()
	ResetSecurityGroupIds()
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

// The jsii proxy struct for NetworkingPortV2
type jsiiProxy_NetworkingPortV2 struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_NetworkingPortV2) AdminStateUp() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"adminStateUp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkingPortV2) AdminStateUpInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"adminStateUpInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkingPortV2) AllFixedIps() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allFixedIps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkingPortV2) AllowedAddressPairs() NetworkingPortV2AllowedAddressPairsList {
	var returns NetworkingPortV2AllowedAddressPairsList
	_jsii_.Get(
		j,
		"allowedAddressPairs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkingPortV2) AllowedAddressPairsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowedAddressPairsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkingPortV2) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkingPortV2) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkingPortV2) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkingPortV2) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkingPortV2) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkingPortV2) DeviceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deviceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkingPortV2) DeviceIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deviceIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkingPortV2) DeviceOwner() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deviceOwner",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkingPortV2) DeviceOwnerInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deviceOwnerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkingPortV2) FixedIp() NetworkingPortV2FixedIpList {
	var returns NetworkingPortV2FixedIpList
	_jsii_.Get(
		j,
		"fixedIp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkingPortV2) FixedIpInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"fixedIpInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkingPortV2) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkingPortV2) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkingPortV2) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkingPortV2) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkingPortV2) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkingPortV2) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkingPortV2) MacAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"macAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkingPortV2) MacAddressInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"macAddressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkingPortV2) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkingPortV2) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkingPortV2) NetworkId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkingPortV2) NetworkIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkingPortV2) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkingPortV2) NoSecurityGroups() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noSecurityGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkingPortV2) NoSecurityGroupsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noSecurityGroupsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkingPortV2) PortSecurityEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"portSecurityEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkingPortV2) PortSecurityEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"portSecurityEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkingPortV2) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkingPortV2) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkingPortV2) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkingPortV2) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkingPortV2) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkingPortV2) SecurityGroupIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityGroupIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkingPortV2) SecurityGroupIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityGroupIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkingPortV2) TenantId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tenantId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkingPortV2) TenantIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tenantIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkingPortV2) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkingPortV2) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkingPortV2) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkingPortV2) Timeouts() NetworkingPortV2TimeoutsOutputReference {
	var returns NetworkingPortV2TimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkingPortV2) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkingPortV2) ValueSpecs() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"valueSpecs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkingPortV2) ValueSpecsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"valueSpecsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.0/docs/resources/networking_port_v2 opentelekomcloud_networking_port_v2} Resource.
func NewNetworkingPortV2(scope constructs.Construct, id *string, config *NetworkingPortV2Config) NetworkingPortV2 {
	_init_.Initialize()

	if err := validateNewNetworkingPortV2Parameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_NetworkingPortV2{}

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.networkingPortV2.NetworkingPortV2",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.0/docs/resources/networking_port_v2 opentelekomcloud_networking_port_v2} Resource.
func NewNetworkingPortV2_Override(n NetworkingPortV2, scope constructs.Construct, id *string, config *NetworkingPortV2Config) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.networkingPortV2.NetworkingPortV2",
		[]interface{}{scope, id, config},
		n,
	)
}

func (j *jsiiProxy_NetworkingPortV2)SetAdminStateUp(val interface{}) {
	if err := j.validateSetAdminStateUpParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"adminStateUp",
		val,
	)
}

func (j *jsiiProxy_NetworkingPortV2)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_NetworkingPortV2)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_NetworkingPortV2)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_NetworkingPortV2)SetDeviceId(val *string) {
	if err := j.validateSetDeviceIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deviceId",
		val,
	)
}

func (j *jsiiProxy_NetworkingPortV2)SetDeviceOwner(val *string) {
	if err := j.validateSetDeviceOwnerParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deviceOwner",
		val,
	)
}

func (j *jsiiProxy_NetworkingPortV2)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_NetworkingPortV2)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_NetworkingPortV2)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_NetworkingPortV2)SetMacAddress(val *string) {
	if err := j.validateSetMacAddressParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"macAddress",
		val,
	)
}

func (j *jsiiProxy_NetworkingPortV2)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_NetworkingPortV2)SetNetworkId(val *string) {
	if err := j.validateSetNetworkIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"networkId",
		val,
	)
}

func (j *jsiiProxy_NetworkingPortV2)SetNoSecurityGroups(val interface{}) {
	if err := j.validateSetNoSecurityGroupsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noSecurityGroups",
		val,
	)
}

func (j *jsiiProxy_NetworkingPortV2)SetPortSecurityEnabled(val interface{}) {
	if err := j.validateSetPortSecurityEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"portSecurityEnabled",
		val,
	)
}

func (j *jsiiProxy_NetworkingPortV2)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_NetworkingPortV2)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_NetworkingPortV2)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_NetworkingPortV2)SetSecurityGroupIds(val *[]*string) {
	if err := j.validateSetSecurityGroupIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"securityGroupIds",
		val,
	)
}

func (j *jsiiProxy_NetworkingPortV2)SetTenantId(val *string) {
	if err := j.validateSetTenantIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tenantId",
		val,
	)
}

func (j *jsiiProxy_NetworkingPortV2)SetValueSpecs(val *map[string]*string) {
	if err := j.validateSetValueSpecsParameters(val); err != nil {
		panic(err)
	}
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
func NetworkingPortV2_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateNetworkingPortV2_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-opentelekomcloud.networkingPortV2.NetworkingPortV2",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func NetworkingPortV2_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateNetworkingPortV2_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-opentelekomcloud.networkingPortV2.NetworkingPortV2",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func NetworkingPortV2_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateNetworkingPortV2_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-opentelekomcloud.networkingPortV2.NetworkingPortV2",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func NetworkingPortV2_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-opentelekomcloud.networkingPortV2.NetworkingPortV2",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (n *jsiiProxy_NetworkingPortV2) AddOverride(path *string, value interface{}) {
	if err := n.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (n *jsiiProxy_NetworkingPortV2) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := n.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		n,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkingPortV2) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := n.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		n,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkingPortV2) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := n.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		n,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkingPortV2) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := n.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		n,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkingPortV2) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := n.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		n,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkingPortV2) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := n.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		n,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkingPortV2) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := n.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		n,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkingPortV2) GetStringAttribute(terraformAttribute *string) *string {
	if err := n.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		n,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkingPortV2) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := n.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		n,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkingPortV2) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := n.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		n,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkingPortV2) OverrideLogicalId(newLogicalId *string) {
	if err := n.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (n *jsiiProxy_NetworkingPortV2) PutAllowedAddressPairs(value interface{}) {
	if err := n.validatePutAllowedAddressPairsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"putAllowedAddressPairs",
		[]interface{}{value},
	)
}

func (n *jsiiProxy_NetworkingPortV2) PutFixedIp(value interface{}) {
	if err := n.validatePutFixedIpParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"putFixedIp",
		[]interface{}{value},
	)
}

func (n *jsiiProxy_NetworkingPortV2) PutTimeouts(value *NetworkingPortV2Timeouts) {
	if err := n.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (n *jsiiProxy_NetworkingPortV2) ResetAdminStateUp() {
	_jsii_.InvokeVoid(
		n,
		"resetAdminStateUp",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkingPortV2) ResetAllowedAddressPairs() {
	_jsii_.InvokeVoid(
		n,
		"resetAllowedAddressPairs",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkingPortV2) ResetDeviceId() {
	_jsii_.InvokeVoid(
		n,
		"resetDeviceId",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkingPortV2) ResetDeviceOwner() {
	_jsii_.InvokeVoid(
		n,
		"resetDeviceOwner",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkingPortV2) ResetFixedIp() {
	_jsii_.InvokeVoid(
		n,
		"resetFixedIp",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkingPortV2) ResetId() {
	_jsii_.InvokeVoid(
		n,
		"resetId",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkingPortV2) ResetMacAddress() {
	_jsii_.InvokeVoid(
		n,
		"resetMacAddress",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkingPortV2) ResetName() {
	_jsii_.InvokeVoid(
		n,
		"resetName",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkingPortV2) ResetNoSecurityGroups() {
	_jsii_.InvokeVoid(
		n,
		"resetNoSecurityGroups",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkingPortV2) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		n,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkingPortV2) ResetPortSecurityEnabled() {
	_jsii_.InvokeVoid(
		n,
		"resetPortSecurityEnabled",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkingPortV2) ResetRegion() {
	_jsii_.InvokeVoid(
		n,
		"resetRegion",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkingPortV2) ResetSecurityGroupIds() {
	_jsii_.InvokeVoid(
		n,
		"resetSecurityGroupIds",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkingPortV2) ResetTenantId() {
	_jsii_.InvokeVoid(
		n,
		"resetTenantId",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkingPortV2) ResetTimeouts() {
	_jsii_.InvokeVoid(
		n,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkingPortV2) ResetValueSpecs() {
	_jsii_.InvokeVoid(
		n,
		"resetValueSpecs",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkingPortV2) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		n,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkingPortV2) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		n,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkingPortV2) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkingPortV2) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		n,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

