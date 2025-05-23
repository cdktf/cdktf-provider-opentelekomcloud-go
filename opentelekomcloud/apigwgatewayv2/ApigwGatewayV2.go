// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package apigwgatewayv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v11/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v11/apigwgatewayv2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.39/docs/resources/apigw_gateway_v2 opentelekomcloud_apigw_gateway_v2}.
type ApigwGatewayV2 interface {
	cdktf.TerraformResource
	AvailabilityZones() *[]*string
	SetAvailabilityZones(val *[]*string)
	AvailabilityZonesInput() *[]*string
	BandwidthChargingMode() *string
	SetBandwidthChargingMode(val *string)
	BandwidthChargingModeInput() *string
	BandwidthSize() *float64
	SetBandwidthSize(val *float64)
	BandwidthSizeInput() *float64
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
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
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
	IngressBandwidthChargingMode() *string
	SetIngressBandwidthChargingMode(val *string)
	IngressBandwidthChargingModeInput() *string
	IngressBandwidthSize() *float64
	SetIngressBandwidthSize(val *float64)
	IngressBandwidthSizeInput() *float64
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	LoadbalancerProvider() *string
	SetLoadbalancerProvider(val *string)
	LoadbalancerProviderInput() *string
	MaintainBegin() *string
	SetMaintainBegin(val *string)
	MaintainBeginInput() *string
	MaintainEnd() *string
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	PrivateEgressAddresses() *[]*string
	ProjectId() *string
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	PublicEgressAddress() *string
	// Experimental.
	RawOverrides() interface{}
	Region() *string
	SecurityGroupId() *string
	SetSecurityGroupId(val *string)
	SecurityGroupIdInput() *string
	SpecId() *string
	SetSpecId(val *string)
	SpecIdInput() *string
	Status() *string
	SubnetId() *string
	SetSubnetId(val *string)
	SubnetIdInput() *string
	SupportedFeatures() *[]*string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() ApigwGatewayV2TimeoutsOutputReference
	TimeoutsInput() interface{}
	VpcepServiceName() *string
	VpcId() *string
	SetVpcId(val *string)
	VpcIdInput() *string
	VpcIngressAddress() *string
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
	HasResourceMove() interface{}
	// Experimental.
	ImportFrom(id *string, provider cdktf.TerraformProvider)
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	// Move the resource corresponding to "id" to this resource.
	//
	// Note that the resource being moved from must be marked as moved using it's instance function.
	// Experimental.
	MoveFromId(id *string)
	// Moves this resource to the target resource given by moveTarget.
	// Experimental.
	MoveTo(moveTarget *string, index interface{})
	// Moves this resource to the resource corresponding to "id".
	// Experimental.
	MoveToId(id *string)
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	PutTimeouts(value *ApigwGatewayV2Timeouts)
	ResetBandwidthChargingMode()
	ResetBandwidthSize()
	ResetDescription()
	ResetId()
	ResetIngressBandwidthChargingMode()
	ResetIngressBandwidthSize()
	ResetLoadbalancerProvider()
	ResetMaintainBegin()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetTimeouts()
	SynthesizeAttributes() *map[string]interface{}
	SynthesizeHclAttributes() *map[string]interface{}
	// Experimental.
	ToHclTerraform() interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for ApigwGatewayV2
type jsiiProxy_ApigwGatewayV2 struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_ApigwGatewayV2) AvailabilityZones() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"availabilityZones",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwGatewayV2) AvailabilityZonesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"availabilityZonesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwGatewayV2) BandwidthChargingMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bandwidthChargingMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwGatewayV2) BandwidthChargingModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bandwidthChargingModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwGatewayV2) BandwidthSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bandwidthSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwGatewayV2) BandwidthSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bandwidthSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwGatewayV2) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwGatewayV2) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwGatewayV2) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwGatewayV2) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwGatewayV2) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwGatewayV2) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwGatewayV2) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwGatewayV2) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwGatewayV2) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwGatewayV2) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwGatewayV2) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwGatewayV2) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwGatewayV2) IngressBandwidthChargingMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ingressBandwidthChargingMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwGatewayV2) IngressBandwidthChargingModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ingressBandwidthChargingModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwGatewayV2) IngressBandwidthSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ingressBandwidthSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwGatewayV2) IngressBandwidthSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ingressBandwidthSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwGatewayV2) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwGatewayV2) LoadbalancerProvider() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loadbalancerProvider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwGatewayV2) LoadbalancerProviderInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loadbalancerProviderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwGatewayV2) MaintainBegin() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maintainBegin",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwGatewayV2) MaintainBeginInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maintainBeginInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwGatewayV2) MaintainEnd() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maintainEnd",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwGatewayV2) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwGatewayV2) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwGatewayV2) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwGatewayV2) PrivateEgressAddresses() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"privateEgressAddresses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwGatewayV2) ProjectId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwGatewayV2) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwGatewayV2) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwGatewayV2) PublicEgressAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"publicEgressAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwGatewayV2) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwGatewayV2) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwGatewayV2) SecurityGroupId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"securityGroupId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwGatewayV2) SecurityGroupIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"securityGroupIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwGatewayV2) SpecId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"specId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwGatewayV2) SpecIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"specIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwGatewayV2) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwGatewayV2) SubnetId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwGatewayV2) SubnetIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwGatewayV2) SupportedFeatures() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"supportedFeatures",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwGatewayV2) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwGatewayV2) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwGatewayV2) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwGatewayV2) Timeouts() ApigwGatewayV2TimeoutsOutputReference {
	var returns ApigwGatewayV2TimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwGatewayV2) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwGatewayV2) VpcepServiceName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcepServiceName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwGatewayV2) VpcId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwGatewayV2) VpcIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigwGatewayV2) VpcIngressAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcIngressAddress",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.39/docs/resources/apigw_gateway_v2 opentelekomcloud_apigw_gateway_v2} Resource.
func NewApigwGatewayV2(scope constructs.Construct, id *string, config *ApigwGatewayV2Config) ApigwGatewayV2 {
	_init_.Initialize()

	if err := validateNewApigwGatewayV2Parameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_ApigwGatewayV2{}

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.apigwGatewayV2.ApigwGatewayV2",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.39/docs/resources/apigw_gateway_v2 opentelekomcloud_apigw_gateway_v2} Resource.
func NewApigwGatewayV2_Override(a ApigwGatewayV2, scope constructs.Construct, id *string, config *ApigwGatewayV2Config) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.apigwGatewayV2.ApigwGatewayV2",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_ApigwGatewayV2)SetAvailabilityZones(val *[]*string) {
	if err := j.validateSetAvailabilityZonesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"availabilityZones",
		val,
	)
}

func (j *jsiiProxy_ApigwGatewayV2)SetBandwidthChargingMode(val *string) {
	if err := j.validateSetBandwidthChargingModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bandwidthChargingMode",
		val,
	)
}

func (j *jsiiProxy_ApigwGatewayV2)SetBandwidthSize(val *float64) {
	if err := j.validateSetBandwidthSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bandwidthSize",
		val,
	)
}

func (j *jsiiProxy_ApigwGatewayV2)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_ApigwGatewayV2)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_ApigwGatewayV2)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_ApigwGatewayV2)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_ApigwGatewayV2)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_ApigwGatewayV2)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_ApigwGatewayV2)SetIngressBandwidthChargingMode(val *string) {
	if err := j.validateSetIngressBandwidthChargingModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ingressBandwidthChargingMode",
		val,
	)
}

func (j *jsiiProxy_ApigwGatewayV2)SetIngressBandwidthSize(val *float64) {
	if err := j.validateSetIngressBandwidthSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ingressBandwidthSize",
		val,
	)
}

func (j *jsiiProxy_ApigwGatewayV2)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_ApigwGatewayV2)SetLoadbalancerProvider(val *string) {
	if err := j.validateSetLoadbalancerProviderParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"loadbalancerProvider",
		val,
	)
}

func (j *jsiiProxy_ApigwGatewayV2)SetMaintainBegin(val *string) {
	if err := j.validateSetMaintainBeginParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maintainBegin",
		val,
	)
}

func (j *jsiiProxy_ApigwGatewayV2)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_ApigwGatewayV2)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_ApigwGatewayV2)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_ApigwGatewayV2)SetSecurityGroupId(val *string) {
	if err := j.validateSetSecurityGroupIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"securityGroupId",
		val,
	)
}

func (j *jsiiProxy_ApigwGatewayV2)SetSpecId(val *string) {
	if err := j.validateSetSpecIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"specId",
		val,
	)
}

func (j *jsiiProxy_ApigwGatewayV2)SetSubnetId(val *string) {
	if err := j.validateSetSubnetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"subnetId",
		val,
	)
}

func (j *jsiiProxy_ApigwGatewayV2)SetVpcId(val *string) {
	if err := j.validateSetVpcIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vpcId",
		val,
	)
}

// Generates CDKTF code for importing a ApigwGatewayV2 resource upon running "cdktf plan <stack-name>".
func ApigwGatewayV2_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateApigwGatewayV2_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-opentelekomcloud.apigwGatewayV2.ApigwGatewayV2",
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
func ApigwGatewayV2_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateApigwGatewayV2_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-opentelekomcloud.apigwGatewayV2.ApigwGatewayV2",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ApigwGatewayV2_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateApigwGatewayV2_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-opentelekomcloud.apigwGatewayV2.ApigwGatewayV2",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ApigwGatewayV2_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateApigwGatewayV2_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-opentelekomcloud.apigwGatewayV2.ApigwGatewayV2",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func ApigwGatewayV2_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-opentelekomcloud.apigwGatewayV2.ApigwGatewayV2",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (a *jsiiProxy_ApigwGatewayV2) AddMoveTarget(moveTarget *string) {
	if err := a.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (a *jsiiProxy_ApigwGatewayV2) AddOverride(path *string, value interface{}) {
	if err := a.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (a *jsiiProxy_ApigwGatewayV2) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := a.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApigwGatewayV2) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := a.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApigwGatewayV2) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := a.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApigwGatewayV2) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := a.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApigwGatewayV2) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := a.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApigwGatewayV2) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := a.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApigwGatewayV2) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := a.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApigwGatewayV2) GetStringAttribute(terraformAttribute *string) *string {
	if err := a.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApigwGatewayV2) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := a.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApigwGatewayV2) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApigwGatewayV2) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := a.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (a *jsiiProxy_ApigwGatewayV2) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := a.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApigwGatewayV2) MoveFromId(id *string) {
	if err := a.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveFromId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_ApigwGatewayV2) MoveTo(moveTarget *string, index interface{}) {
	if err := a.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (a *jsiiProxy_ApigwGatewayV2) MoveToId(id *string) {
	if err := a.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveToId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_ApigwGatewayV2) OverrideLogicalId(newLogicalId *string) {
	if err := a.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_ApigwGatewayV2) PutTimeouts(value *ApigwGatewayV2Timeouts) {
	if err := a.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ApigwGatewayV2) ResetBandwidthChargingMode() {
	_jsii_.InvokeVoid(
		a,
		"resetBandwidthChargingMode",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApigwGatewayV2) ResetBandwidthSize() {
	_jsii_.InvokeVoid(
		a,
		"resetBandwidthSize",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApigwGatewayV2) ResetDescription() {
	_jsii_.InvokeVoid(
		a,
		"resetDescription",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApigwGatewayV2) ResetId() {
	_jsii_.InvokeVoid(
		a,
		"resetId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApigwGatewayV2) ResetIngressBandwidthChargingMode() {
	_jsii_.InvokeVoid(
		a,
		"resetIngressBandwidthChargingMode",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApigwGatewayV2) ResetIngressBandwidthSize() {
	_jsii_.InvokeVoid(
		a,
		"resetIngressBandwidthSize",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApigwGatewayV2) ResetLoadbalancerProvider() {
	_jsii_.InvokeVoid(
		a,
		"resetLoadbalancerProvider",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApigwGatewayV2) ResetMaintainBegin() {
	_jsii_.InvokeVoid(
		a,
		"resetMaintainBegin",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApigwGatewayV2) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApigwGatewayV2) ResetTimeouts() {
	_jsii_.InvokeVoid(
		a,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApigwGatewayV2) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApigwGatewayV2) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApigwGatewayV2) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApigwGatewayV2) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApigwGatewayV2) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApigwGatewayV2) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

