// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dwsclusterv1

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v12/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v12/dwsclusterv1/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.52/docs/resources/dws_cluster_v1 opentelekomcloud_dws_cluster_v1}.
type DwsClusterV1 interface {
	cdktf.TerraformResource
	AvailabilityZone() *string
	SetAvailabilityZone(val *string)
	AvailabilityZoneInput() *string
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
	Created() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Endpoints() DwsClusterV1EndpointsList
	EndpointsInput() interface{}
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
	KeepLastManualSnapshot() *float64
	SetKeepLastManualSnapshot(val *float64)
	KeepLastManualSnapshotInput() *float64
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Name() *string
	SetName(val *string)
	NameInput() *string
	NetworkId() *string
	SetNetworkId(val *string)
	NetworkIdInput() *string
	// The tree node.
	Node() constructs.Node
	NodeType() *string
	SetNodeType(val *string)
	NodeTypeInput() *string
	NumberOfCn() *float64
	SetNumberOfCn(val *float64)
	NumberOfCnInput() *float64
	NumberOfNode() *float64
	SetNumberOfNode(val *float64)
	NumberOfNodeInput() *float64
	Port() *float64
	SetPort(val *float64)
	PortInput() *float64
	PrivateIp() *[]*string
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	PublicEndpoints() DwsClusterV1PublicEndpointsList
	PublicEndpointsInput() interface{}
	PublicIp() DwsClusterV1PublicIpOutputReference
	PublicIpInput() *DwsClusterV1PublicIp
	// Experimental.
	RawOverrides() interface{}
	RecentEvent() *float64
	SecurityGroupId() *string
	SetSecurityGroupId(val *string)
	SecurityGroupIdInput() *string
	Status() *string
	SubStatus() *string
	TaskStatus() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() DwsClusterV1TimeoutsOutputReference
	TimeoutsInput() interface{}
	Updated() *string
	UserName() *string
	SetUserName(val *string)
	UserNameInput() *string
	UserPwd() *string
	SetUserPwd(val *string)
	UserPwdInput() *string
	Version() *string
	VpcId() *string
	SetVpcId(val *string)
	VpcIdInput() *string
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
	PutEndpoints(value interface{})
	PutPublicEndpoints(value interface{})
	PutPublicIp(value *DwsClusterV1PublicIp)
	PutTimeouts(value *DwsClusterV1Timeouts)
	ResetAvailabilityZone()
	ResetEndpoints()
	ResetId()
	ResetKeepLastManualSnapshot()
	ResetNumberOfCn()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPort()
	ResetPublicEndpoints()
	ResetPublicIp()
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

// The jsii proxy struct for DwsClusterV1
type jsiiProxy_DwsClusterV1 struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_DwsClusterV1) AvailabilityZone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"availabilityZone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DwsClusterV1) AvailabilityZoneInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"availabilityZoneInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DwsClusterV1) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DwsClusterV1) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DwsClusterV1) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DwsClusterV1) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DwsClusterV1) Created() *string {
	var returns *string
	_jsii_.Get(
		j,
		"created",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DwsClusterV1) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DwsClusterV1) Endpoints() DwsClusterV1EndpointsList {
	var returns DwsClusterV1EndpointsList
	_jsii_.Get(
		j,
		"endpoints",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DwsClusterV1) EndpointsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"endpointsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DwsClusterV1) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DwsClusterV1) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DwsClusterV1) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DwsClusterV1) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DwsClusterV1) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DwsClusterV1) KeepLastManualSnapshot() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"keepLastManualSnapshot",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DwsClusterV1) KeepLastManualSnapshotInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"keepLastManualSnapshotInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DwsClusterV1) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DwsClusterV1) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DwsClusterV1) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DwsClusterV1) NetworkId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DwsClusterV1) NetworkIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DwsClusterV1) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DwsClusterV1) NodeType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nodeType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DwsClusterV1) NodeTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nodeTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DwsClusterV1) NumberOfCn() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"numberOfCn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DwsClusterV1) NumberOfCnInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"numberOfCnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DwsClusterV1) NumberOfNode() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"numberOfNode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DwsClusterV1) NumberOfNodeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"numberOfNodeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DwsClusterV1) Port() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"port",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DwsClusterV1) PortInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"portInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DwsClusterV1) PrivateIp() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"privateIp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DwsClusterV1) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DwsClusterV1) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DwsClusterV1) PublicEndpoints() DwsClusterV1PublicEndpointsList {
	var returns DwsClusterV1PublicEndpointsList
	_jsii_.Get(
		j,
		"publicEndpoints",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DwsClusterV1) PublicEndpointsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"publicEndpointsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DwsClusterV1) PublicIp() DwsClusterV1PublicIpOutputReference {
	var returns DwsClusterV1PublicIpOutputReference
	_jsii_.Get(
		j,
		"publicIp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DwsClusterV1) PublicIpInput() *DwsClusterV1PublicIp {
	var returns *DwsClusterV1PublicIp
	_jsii_.Get(
		j,
		"publicIpInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DwsClusterV1) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DwsClusterV1) RecentEvent() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"recentEvent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DwsClusterV1) SecurityGroupId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"securityGroupId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DwsClusterV1) SecurityGroupIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"securityGroupIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DwsClusterV1) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DwsClusterV1) SubStatus() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DwsClusterV1) TaskStatus() *string {
	var returns *string
	_jsii_.Get(
		j,
		"taskStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DwsClusterV1) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DwsClusterV1) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DwsClusterV1) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DwsClusterV1) Timeouts() DwsClusterV1TimeoutsOutputReference {
	var returns DwsClusterV1TimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DwsClusterV1) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DwsClusterV1) Updated() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updated",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DwsClusterV1) UserName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DwsClusterV1) UserNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DwsClusterV1) UserPwd() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userPwd",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DwsClusterV1) UserPwdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userPwdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DwsClusterV1) Version() *string {
	var returns *string
	_jsii_.Get(
		j,
		"version",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DwsClusterV1) VpcId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DwsClusterV1) VpcIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcIdInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.52/docs/resources/dws_cluster_v1 opentelekomcloud_dws_cluster_v1} Resource.
func NewDwsClusterV1(scope constructs.Construct, id *string, config *DwsClusterV1Config) DwsClusterV1 {
	_init_.Initialize()

	if err := validateNewDwsClusterV1Parameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_DwsClusterV1{}

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.dwsClusterV1.DwsClusterV1",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.52/docs/resources/dws_cluster_v1 opentelekomcloud_dws_cluster_v1} Resource.
func NewDwsClusterV1_Override(d DwsClusterV1, scope constructs.Construct, id *string, config *DwsClusterV1Config) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.dwsClusterV1.DwsClusterV1",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DwsClusterV1)SetAvailabilityZone(val *string) {
	if err := j.validateSetAvailabilityZoneParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"availabilityZone",
		val,
	)
}

func (j *jsiiProxy_DwsClusterV1)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_DwsClusterV1)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DwsClusterV1)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DwsClusterV1)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DwsClusterV1)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_DwsClusterV1)SetKeepLastManualSnapshot(val *float64) {
	if err := j.validateSetKeepLastManualSnapshotParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"keepLastManualSnapshot",
		val,
	)
}

func (j *jsiiProxy_DwsClusterV1)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DwsClusterV1)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_DwsClusterV1)SetNetworkId(val *string) {
	if err := j.validateSetNetworkIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"networkId",
		val,
	)
}

func (j *jsiiProxy_DwsClusterV1)SetNodeType(val *string) {
	if err := j.validateSetNodeTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nodeType",
		val,
	)
}

func (j *jsiiProxy_DwsClusterV1)SetNumberOfCn(val *float64) {
	if err := j.validateSetNumberOfCnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"numberOfCn",
		val,
	)
}

func (j *jsiiProxy_DwsClusterV1)SetNumberOfNode(val *float64) {
	if err := j.validateSetNumberOfNodeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"numberOfNode",
		val,
	)
}

func (j *jsiiProxy_DwsClusterV1)SetPort(val *float64) {
	if err := j.validateSetPortParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"port",
		val,
	)
}

func (j *jsiiProxy_DwsClusterV1)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DwsClusterV1)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_DwsClusterV1)SetSecurityGroupId(val *string) {
	if err := j.validateSetSecurityGroupIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"securityGroupId",
		val,
	)
}

func (j *jsiiProxy_DwsClusterV1)SetUserName(val *string) {
	if err := j.validateSetUserNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"userName",
		val,
	)
}

func (j *jsiiProxy_DwsClusterV1)SetUserPwd(val *string) {
	if err := j.validateSetUserPwdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"userPwd",
		val,
	)
}

func (j *jsiiProxy_DwsClusterV1)SetVpcId(val *string) {
	if err := j.validateSetVpcIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vpcId",
		val,
	)
}

// Generates CDKTF code for importing a DwsClusterV1 resource upon running "cdktf plan <stack-name>".
func DwsClusterV1_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateDwsClusterV1_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-opentelekomcloud.dwsClusterV1.DwsClusterV1",
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
func DwsClusterV1_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDwsClusterV1_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-opentelekomcloud.dwsClusterV1.DwsClusterV1",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DwsClusterV1_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDwsClusterV1_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-opentelekomcloud.dwsClusterV1.DwsClusterV1",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DwsClusterV1_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDwsClusterV1_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-opentelekomcloud.dwsClusterV1.DwsClusterV1",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DwsClusterV1_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-opentelekomcloud.dwsClusterV1.DwsClusterV1",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DwsClusterV1) AddMoveTarget(moveTarget *string) {
	if err := d.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (d *jsiiProxy_DwsClusterV1) AddOverride(path *string, value interface{}) {
	if err := d.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DwsClusterV1) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := d.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DwsClusterV1) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DwsClusterV1) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := d.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DwsClusterV1) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := d.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DwsClusterV1) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := d.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DwsClusterV1) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := d.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DwsClusterV1) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := d.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DwsClusterV1) GetStringAttribute(terraformAttribute *string) *string {
	if err := d.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DwsClusterV1) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := d.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DwsClusterV1) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DwsClusterV1) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := d.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (d *jsiiProxy_DwsClusterV1) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DwsClusterV1) MoveFromId(id *string) {
	if err := d.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"moveFromId",
		[]interface{}{id},
	)
}

func (d *jsiiProxy_DwsClusterV1) MoveTo(moveTarget *string, index interface{}) {
	if err := d.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (d *jsiiProxy_DwsClusterV1) MoveToId(id *string) {
	if err := d.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"moveToId",
		[]interface{}{id},
	)
}

func (d *jsiiProxy_DwsClusterV1) OverrideLogicalId(newLogicalId *string) {
	if err := d.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DwsClusterV1) PutEndpoints(value interface{}) {
	if err := d.validatePutEndpointsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putEndpoints",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DwsClusterV1) PutPublicEndpoints(value interface{}) {
	if err := d.validatePutPublicEndpointsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putPublicEndpoints",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DwsClusterV1) PutPublicIp(value *DwsClusterV1PublicIp) {
	if err := d.validatePutPublicIpParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putPublicIp",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DwsClusterV1) PutTimeouts(value *DwsClusterV1Timeouts) {
	if err := d.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DwsClusterV1) ResetAvailabilityZone() {
	_jsii_.InvokeVoid(
		d,
		"resetAvailabilityZone",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DwsClusterV1) ResetEndpoints() {
	_jsii_.InvokeVoid(
		d,
		"resetEndpoints",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DwsClusterV1) ResetId() {
	_jsii_.InvokeVoid(
		d,
		"resetId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DwsClusterV1) ResetKeepLastManualSnapshot() {
	_jsii_.InvokeVoid(
		d,
		"resetKeepLastManualSnapshot",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DwsClusterV1) ResetNumberOfCn() {
	_jsii_.InvokeVoid(
		d,
		"resetNumberOfCn",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DwsClusterV1) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DwsClusterV1) ResetPort() {
	_jsii_.InvokeVoid(
		d,
		"resetPort",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DwsClusterV1) ResetPublicEndpoints() {
	_jsii_.InvokeVoid(
		d,
		"resetPublicEndpoints",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DwsClusterV1) ResetPublicIp() {
	_jsii_.InvokeVoid(
		d,
		"resetPublicIp",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DwsClusterV1) ResetTimeouts() {
	_jsii_.InvokeVoid(
		d,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DwsClusterV1) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DwsClusterV1) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DwsClusterV1) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DwsClusterV1) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DwsClusterV1) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DwsClusterV1) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

