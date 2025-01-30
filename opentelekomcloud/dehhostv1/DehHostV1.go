// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dehhostv1

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v11/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v11/dehhostv1/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.30/docs/resources/deh_host_v1 opentelekomcloud_deh_host_v1}.
type DehHostV1 interface {
	cdktf.TerraformResource
	AutoPlacement() *string
	SetAutoPlacement(val *string)
	AutoPlacementInput() *string
	AvailabilityZone() *string
	SetAvailabilityZone(val *string)
	AvailabilityZoneInput() *string
	AvailableInstanceCapacities() DehHostV1AvailableInstanceCapacitiesList
	AvailableInstanceCapacitiesInput() interface{}
	AvailableMemory() *float64
	SetAvailableMemory(val *float64)
	AvailableMemoryInput() *float64
	AvailableVcpus() *float64
	SetAvailableVcpus(val *float64)
	AvailableVcpusInput() *float64
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	Cores() *float64
	SetCores(val *float64)
	CoresInput() *float64
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
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
	HostType() *string
	SetHostType(val *string)
	HostTypeInput() *string
	HostTypeName() *string
	SetHostTypeName(val *string)
	HostTypeNameInput() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	InstanceTotal() *float64
	SetInstanceTotal(val *float64)
	InstanceTotalInput() *float64
	InstanceUuids() *[]*string
	SetInstanceUuids(val *[]*string)
	InstanceUuidsInput() *[]*string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Memory() *float64
	SetMemory(val *float64)
	MemoryInput() *float64
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
	Sockets() *float64
	SetSockets(val *float64)
	SocketsInput() *float64
	Status() *string
	SetStatus(val *string)
	StatusInput() *string
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsInput() *map[string]*string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() DehHostV1TimeoutsOutputReference
	TimeoutsInput() interface{}
	Vcpus() *float64
	SetVcpus(val *float64)
	VcpusInput() *float64
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
	PutAvailableInstanceCapacities(value interface{})
	PutTimeouts(value *DehHostV1Timeouts)
	ResetAutoPlacement()
	ResetAvailableInstanceCapacities()
	ResetAvailableMemory()
	ResetAvailableVcpus()
	ResetCores()
	ResetHostTypeName()
	ResetId()
	ResetInstanceTotal()
	ResetInstanceUuids()
	ResetMemory()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetRegion()
	ResetSockets()
	ResetStatus()
	ResetTags()
	ResetTimeouts()
	ResetVcpus()
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

// The jsii proxy struct for DehHostV1
type jsiiProxy_DehHostV1 struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_DehHostV1) AutoPlacement() *string {
	var returns *string
	_jsii_.Get(
		j,
		"autoPlacement",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DehHostV1) AutoPlacementInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"autoPlacementInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DehHostV1) AvailabilityZone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"availabilityZone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DehHostV1) AvailabilityZoneInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"availabilityZoneInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DehHostV1) AvailableInstanceCapacities() DehHostV1AvailableInstanceCapacitiesList {
	var returns DehHostV1AvailableInstanceCapacitiesList
	_jsii_.Get(
		j,
		"availableInstanceCapacities",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DehHostV1) AvailableInstanceCapacitiesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"availableInstanceCapacitiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DehHostV1) AvailableMemory() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"availableMemory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DehHostV1) AvailableMemoryInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"availableMemoryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DehHostV1) AvailableVcpus() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"availableVcpus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DehHostV1) AvailableVcpusInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"availableVcpusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DehHostV1) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DehHostV1) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DehHostV1) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DehHostV1) Cores() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"cores",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DehHostV1) CoresInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"coresInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DehHostV1) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DehHostV1) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DehHostV1) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DehHostV1) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DehHostV1) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DehHostV1) HostType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DehHostV1) HostTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DehHostV1) HostTypeName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostTypeName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DehHostV1) HostTypeNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostTypeNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DehHostV1) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DehHostV1) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DehHostV1) InstanceTotal() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"instanceTotal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DehHostV1) InstanceTotalInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"instanceTotalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DehHostV1) InstanceUuids() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"instanceUuids",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DehHostV1) InstanceUuidsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"instanceUuidsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DehHostV1) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DehHostV1) Memory() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"memory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DehHostV1) MemoryInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"memoryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DehHostV1) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DehHostV1) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DehHostV1) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DehHostV1) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DehHostV1) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DehHostV1) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DehHostV1) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DehHostV1) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DehHostV1) Sockets() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sockets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DehHostV1) SocketsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"socketsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DehHostV1) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DehHostV1) StatusInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"statusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DehHostV1) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DehHostV1) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DehHostV1) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DehHostV1) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DehHostV1) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DehHostV1) Timeouts() DehHostV1TimeoutsOutputReference {
	var returns DehHostV1TimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DehHostV1) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DehHostV1) Vcpus() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"vcpus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DehHostV1) VcpusInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"vcpusInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.30/docs/resources/deh_host_v1 opentelekomcloud_deh_host_v1} Resource.
func NewDehHostV1(scope constructs.Construct, id *string, config *DehHostV1Config) DehHostV1 {
	_init_.Initialize()

	if err := validateNewDehHostV1Parameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_DehHostV1{}

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.dehHostV1.DehHostV1",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.30/docs/resources/deh_host_v1 opentelekomcloud_deh_host_v1} Resource.
func NewDehHostV1_Override(d DehHostV1, scope constructs.Construct, id *string, config *DehHostV1Config) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.dehHostV1.DehHostV1",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DehHostV1)SetAutoPlacement(val *string) {
	if err := j.validateSetAutoPlacementParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoPlacement",
		val,
	)
}

func (j *jsiiProxy_DehHostV1)SetAvailabilityZone(val *string) {
	if err := j.validateSetAvailabilityZoneParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"availabilityZone",
		val,
	)
}

func (j *jsiiProxy_DehHostV1)SetAvailableMemory(val *float64) {
	if err := j.validateSetAvailableMemoryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"availableMemory",
		val,
	)
}

func (j *jsiiProxy_DehHostV1)SetAvailableVcpus(val *float64) {
	if err := j.validateSetAvailableVcpusParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"availableVcpus",
		val,
	)
}

func (j *jsiiProxy_DehHostV1)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_DehHostV1)SetCores(val *float64) {
	if err := j.validateSetCoresParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cores",
		val,
	)
}

func (j *jsiiProxy_DehHostV1)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DehHostV1)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DehHostV1)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DehHostV1)SetHostType(val *string) {
	if err := j.validateSetHostTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hostType",
		val,
	)
}

func (j *jsiiProxy_DehHostV1)SetHostTypeName(val *string) {
	if err := j.validateSetHostTypeNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hostTypeName",
		val,
	)
}

func (j *jsiiProxy_DehHostV1)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_DehHostV1)SetInstanceTotal(val *float64) {
	if err := j.validateSetInstanceTotalParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instanceTotal",
		val,
	)
}

func (j *jsiiProxy_DehHostV1)SetInstanceUuids(val *[]*string) {
	if err := j.validateSetInstanceUuidsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instanceUuids",
		val,
	)
}

func (j *jsiiProxy_DehHostV1)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DehHostV1)SetMemory(val *float64) {
	if err := j.validateSetMemoryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"memory",
		val,
	)
}

func (j *jsiiProxy_DehHostV1)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_DehHostV1)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DehHostV1)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_DehHostV1)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_DehHostV1)SetSockets(val *float64) {
	if err := j.validateSetSocketsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sockets",
		val,
	)
}

func (j *jsiiProxy_DehHostV1)SetStatus(val *string) {
	if err := j.validateSetStatusParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"status",
		val,
	)
}

func (j *jsiiProxy_DehHostV1)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_DehHostV1)SetVcpus(val *float64) {
	if err := j.validateSetVcpusParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vcpus",
		val,
	)
}

// Generates CDKTF code for importing a DehHostV1 resource upon running "cdktf plan <stack-name>".
func DehHostV1_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateDehHostV1_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-opentelekomcloud.dehHostV1.DehHostV1",
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
func DehHostV1_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDehHostV1_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-opentelekomcloud.dehHostV1.DehHostV1",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DehHostV1_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDehHostV1_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-opentelekomcloud.dehHostV1.DehHostV1",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DehHostV1_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDehHostV1_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-opentelekomcloud.dehHostV1.DehHostV1",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DehHostV1_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-opentelekomcloud.dehHostV1.DehHostV1",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DehHostV1) AddMoveTarget(moveTarget *string) {
	if err := d.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (d *jsiiProxy_DehHostV1) AddOverride(path *string, value interface{}) {
	if err := d.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DehHostV1) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DehHostV1) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DehHostV1) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DehHostV1) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DehHostV1) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DehHostV1) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DehHostV1) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DehHostV1) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DehHostV1) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DehHostV1) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DehHostV1) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := d.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (d *jsiiProxy_DehHostV1) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DehHostV1) MoveFromId(id *string) {
	if err := d.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"moveFromId",
		[]interface{}{id},
	)
}

func (d *jsiiProxy_DehHostV1) MoveTo(moveTarget *string, index interface{}) {
	if err := d.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (d *jsiiProxy_DehHostV1) MoveToId(id *string) {
	if err := d.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"moveToId",
		[]interface{}{id},
	)
}

func (d *jsiiProxy_DehHostV1) OverrideLogicalId(newLogicalId *string) {
	if err := d.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DehHostV1) PutAvailableInstanceCapacities(value interface{}) {
	if err := d.validatePutAvailableInstanceCapacitiesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putAvailableInstanceCapacities",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DehHostV1) PutTimeouts(value *DehHostV1Timeouts) {
	if err := d.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DehHostV1) ResetAutoPlacement() {
	_jsii_.InvokeVoid(
		d,
		"resetAutoPlacement",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DehHostV1) ResetAvailableInstanceCapacities() {
	_jsii_.InvokeVoid(
		d,
		"resetAvailableInstanceCapacities",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DehHostV1) ResetAvailableMemory() {
	_jsii_.InvokeVoid(
		d,
		"resetAvailableMemory",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DehHostV1) ResetAvailableVcpus() {
	_jsii_.InvokeVoid(
		d,
		"resetAvailableVcpus",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DehHostV1) ResetCores() {
	_jsii_.InvokeVoid(
		d,
		"resetCores",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DehHostV1) ResetHostTypeName() {
	_jsii_.InvokeVoid(
		d,
		"resetHostTypeName",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DehHostV1) ResetId() {
	_jsii_.InvokeVoid(
		d,
		"resetId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DehHostV1) ResetInstanceTotal() {
	_jsii_.InvokeVoid(
		d,
		"resetInstanceTotal",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DehHostV1) ResetInstanceUuids() {
	_jsii_.InvokeVoid(
		d,
		"resetInstanceUuids",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DehHostV1) ResetMemory() {
	_jsii_.InvokeVoid(
		d,
		"resetMemory",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DehHostV1) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DehHostV1) ResetRegion() {
	_jsii_.InvokeVoid(
		d,
		"resetRegion",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DehHostV1) ResetSockets() {
	_jsii_.InvokeVoid(
		d,
		"resetSockets",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DehHostV1) ResetStatus() {
	_jsii_.InvokeVoid(
		d,
		"resetStatus",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DehHostV1) ResetTags() {
	_jsii_.InvokeVoid(
		d,
		"resetTags",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DehHostV1) ResetTimeouts() {
	_jsii_.InvokeVoid(
		d,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DehHostV1) ResetVcpus() {
	_jsii_.InvokeVoid(
		d,
		"resetVcpus",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DehHostV1) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DehHostV1) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DehHostV1) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DehHostV1) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DehHostV1) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DehHostV1) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

