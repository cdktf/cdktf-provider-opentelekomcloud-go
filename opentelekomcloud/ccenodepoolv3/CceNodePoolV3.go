// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ccenodepoolv3

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v11/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v11/ccenodepoolv3/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.12/docs/resources/cce_node_pool_v3 opentelekomcloud_cce_node_pool_v3}.
type CceNodePoolV3 interface {
	cdktf.TerraformResource
	AgencyName() *string
	SetAgencyName(val *string)
	AgencyNameInput() *string
	AvailabilityZone() *string
	SetAvailabilityZone(val *string)
	AvailabilityZoneInput() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	ClusterId() *string
	SetClusterId(val *string)
	ClusterIdInput() *string
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
	DataVolumes() CceNodePoolV3DataVolumesList
	DataVolumesInput() interface{}
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	DockerBaseSize() *float64
	SetDockerBaseSize(val *float64)
	DockerBaseSizeInput() *float64
	DockerLvmConfigOverride() *string
	SetDockerLvmConfigOverride(val *string)
	DockerLvmConfigOverrideInput() *string
	Flavor() *string
	SetFlavor(val *string)
	FlavorInput() *string
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
	InitialNodeCount() *float64
	SetInitialNodeCount(val *float64)
	InitialNodeCountInput() *float64
	K8STags() *map[string]*string
	SetK8STags(val *map[string]*string)
	K8STagsInput() *map[string]*string
	KeyPair() *string
	SetKeyPair(val *string)
	KeyPairInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	MaxNodeCount() *float64
	SetMaxNodeCount(val *float64)
	MaxNodeCountInput() *float64
	MaxPods() *float64
	SetMaxPods(val *float64)
	MaxPodsInput() *float64
	MinNodeCount() *float64
	SetMinNodeCount(val *float64)
	MinNodeCountInput() *float64
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	Os() *string
	SetOs(val *string)
	OsInput() *string
	Password() *string
	SetPassword(val *string)
	PasswordInput() *string
	Postinstall() *string
	SetPostinstall(val *string)
	PostinstallInput() *string
	Preinstall() *string
	SetPreinstall(val *string)
	PreinstallInput() *string
	Priority() *float64
	SetPriority(val *float64)
	PriorityInput() *float64
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
	RootVolume() CceNodePoolV3RootVolumeOutputReference
	RootVolumeInput() *CceNodePoolV3RootVolume
	Runtime() *string
	SetRuntime(val *string)
	RuntimeInput() *string
	ScaleDownCooldownTime() *float64
	SetScaleDownCooldownTime(val *float64)
	ScaleDownCooldownTimeInput() *float64
	ScaleEnable() interface{}
	SetScaleEnable(val interface{})
	ScaleEnableInput() interface{}
	ServerGroupReference() *string
	SetServerGroupReference(val *string)
	ServerGroupReferenceInput() *string
	Status() *string
	SubnetId() *string
	SetSubnetId(val *string)
	SubnetIdInput() *string
	Taints() CceNodePoolV3TaintsList
	TaintsInput() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() CceNodePoolV3TimeoutsOutputReference
	TimeoutsInput() interface{}
	UserTags() *map[string]*string
	SetUserTags(val *map[string]*string)
	UserTagsInput() *map[string]*string
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
	PutDataVolumes(value interface{})
	PutRootVolume(value *CceNodePoolV3RootVolume)
	PutTaints(value interface{})
	PutTimeouts(value *CceNodePoolV3Timeouts)
	ResetAgencyName()
	ResetAvailabilityZone()
	ResetDockerBaseSize()
	ResetDockerLvmConfigOverride()
	ResetId()
	ResetK8STags()
	ResetKeyPair()
	ResetMaxNodeCount()
	ResetMaxPods()
	ResetMinNodeCount()
	ResetOs()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPassword()
	ResetPostinstall()
	ResetPreinstall()
	ResetPriority()
	ResetRuntime()
	ResetScaleDownCooldownTime()
	ResetScaleEnable()
	ResetServerGroupReference()
	ResetSubnetId()
	ResetTaints()
	ResetTimeouts()
	ResetUserTags()
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

// The jsii proxy struct for CceNodePoolV3
type jsiiProxy_CceNodePoolV3 struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_CceNodePoolV3) AgencyName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"agencyName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceNodePoolV3) AgencyNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"agencyNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceNodePoolV3) AvailabilityZone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"availabilityZone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceNodePoolV3) AvailabilityZoneInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"availabilityZoneInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceNodePoolV3) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceNodePoolV3) ClusterId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceNodePoolV3) ClusterIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceNodePoolV3) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceNodePoolV3) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceNodePoolV3) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceNodePoolV3) DataVolumes() CceNodePoolV3DataVolumesList {
	var returns CceNodePoolV3DataVolumesList
	_jsii_.Get(
		j,
		"dataVolumes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceNodePoolV3) DataVolumesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dataVolumesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceNodePoolV3) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceNodePoolV3) DockerBaseSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"dockerBaseSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceNodePoolV3) DockerBaseSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"dockerBaseSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceNodePoolV3) DockerLvmConfigOverride() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dockerLvmConfigOverride",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceNodePoolV3) DockerLvmConfigOverrideInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dockerLvmConfigOverrideInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceNodePoolV3) Flavor() *string {
	var returns *string
	_jsii_.Get(
		j,
		"flavor",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceNodePoolV3) FlavorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"flavorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceNodePoolV3) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceNodePoolV3) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceNodePoolV3) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceNodePoolV3) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceNodePoolV3) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceNodePoolV3) InitialNodeCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"initialNodeCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceNodePoolV3) InitialNodeCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"initialNodeCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceNodePoolV3) K8STags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"k8STags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceNodePoolV3) K8STagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"k8STagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceNodePoolV3) KeyPair() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyPair",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceNodePoolV3) KeyPairInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyPairInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceNodePoolV3) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceNodePoolV3) MaxNodeCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxNodeCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceNodePoolV3) MaxNodeCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxNodeCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceNodePoolV3) MaxPods() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxPods",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceNodePoolV3) MaxPodsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxPodsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceNodePoolV3) MinNodeCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minNodeCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceNodePoolV3) MinNodeCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minNodeCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceNodePoolV3) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceNodePoolV3) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceNodePoolV3) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceNodePoolV3) Os() *string {
	var returns *string
	_jsii_.Get(
		j,
		"os",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceNodePoolV3) OsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"osInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceNodePoolV3) Password() *string {
	var returns *string
	_jsii_.Get(
		j,
		"password",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceNodePoolV3) PasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"passwordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceNodePoolV3) Postinstall() *string {
	var returns *string
	_jsii_.Get(
		j,
		"postinstall",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceNodePoolV3) PostinstallInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"postinstallInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceNodePoolV3) Preinstall() *string {
	var returns *string
	_jsii_.Get(
		j,
		"preinstall",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceNodePoolV3) PreinstallInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"preinstallInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceNodePoolV3) Priority() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"priority",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceNodePoolV3) PriorityInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"priorityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceNodePoolV3) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceNodePoolV3) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceNodePoolV3) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceNodePoolV3) RootVolume() CceNodePoolV3RootVolumeOutputReference {
	var returns CceNodePoolV3RootVolumeOutputReference
	_jsii_.Get(
		j,
		"rootVolume",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceNodePoolV3) RootVolumeInput() *CceNodePoolV3RootVolume {
	var returns *CceNodePoolV3RootVolume
	_jsii_.Get(
		j,
		"rootVolumeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceNodePoolV3) Runtime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runtime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceNodePoolV3) RuntimeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runtimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceNodePoolV3) ScaleDownCooldownTime() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"scaleDownCooldownTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceNodePoolV3) ScaleDownCooldownTimeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"scaleDownCooldownTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceNodePoolV3) ScaleEnable() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"scaleEnable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceNodePoolV3) ScaleEnableInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"scaleEnableInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceNodePoolV3) ServerGroupReference() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serverGroupReference",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceNodePoolV3) ServerGroupReferenceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serverGroupReferenceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceNodePoolV3) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceNodePoolV3) SubnetId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceNodePoolV3) SubnetIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceNodePoolV3) Taints() CceNodePoolV3TaintsList {
	var returns CceNodePoolV3TaintsList
	_jsii_.Get(
		j,
		"taints",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceNodePoolV3) TaintsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"taintsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceNodePoolV3) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceNodePoolV3) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceNodePoolV3) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceNodePoolV3) Timeouts() CceNodePoolV3TimeoutsOutputReference {
	var returns CceNodePoolV3TimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceNodePoolV3) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceNodePoolV3) UserTags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"userTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceNodePoolV3) UserTagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"userTagsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.12/docs/resources/cce_node_pool_v3 opentelekomcloud_cce_node_pool_v3} Resource.
func NewCceNodePoolV3(scope constructs.Construct, id *string, config *CceNodePoolV3Config) CceNodePoolV3 {
	_init_.Initialize()

	if err := validateNewCceNodePoolV3Parameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_CceNodePoolV3{}

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.cceNodePoolV3.CceNodePoolV3",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.12/docs/resources/cce_node_pool_v3 opentelekomcloud_cce_node_pool_v3} Resource.
func NewCceNodePoolV3_Override(c CceNodePoolV3, scope constructs.Construct, id *string, config *CceNodePoolV3Config) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.cceNodePoolV3.CceNodePoolV3",
		[]interface{}{scope, id, config},
		c,
	)
}

func (j *jsiiProxy_CceNodePoolV3)SetAgencyName(val *string) {
	if err := j.validateSetAgencyNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"agencyName",
		val,
	)
}

func (j *jsiiProxy_CceNodePoolV3)SetAvailabilityZone(val *string) {
	if err := j.validateSetAvailabilityZoneParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"availabilityZone",
		val,
	)
}

func (j *jsiiProxy_CceNodePoolV3)SetClusterId(val *string) {
	if err := j.validateSetClusterIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clusterId",
		val,
	)
}

func (j *jsiiProxy_CceNodePoolV3)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_CceNodePoolV3)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_CceNodePoolV3)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_CceNodePoolV3)SetDockerBaseSize(val *float64) {
	if err := j.validateSetDockerBaseSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dockerBaseSize",
		val,
	)
}

func (j *jsiiProxy_CceNodePoolV3)SetDockerLvmConfigOverride(val *string) {
	if err := j.validateSetDockerLvmConfigOverrideParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dockerLvmConfigOverride",
		val,
	)
}

func (j *jsiiProxy_CceNodePoolV3)SetFlavor(val *string) {
	if err := j.validateSetFlavorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"flavor",
		val,
	)
}

func (j *jsiiProxy_CceNodePoolV3)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_CceNodePoolV3)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_CceNodePoolV3)SetInitialNodeCount(val *float64) {
	if err := j.validateSetInitialNodeCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"initialNodeCount",
		val,
	)
}

func (j *jsiiProxy_CceNodePoolV3)SetK8STags(val *map[string]*string) {
	if err := j.validateSetK8STagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"k8STags",
		val,
	)
}

func (j *jsiiProxy_CceNodePoolV3)SetKeyPair(val *string) {
	if err := j.validateSetKeyPairParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"keyPair",
		val,
	)
}

func (j *jsiiProxy_CceNodePoolV3)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_CceNodePoolV3)SetMaxNodeCount(val *float64) {
	if err := j.validateSetMaxNodeCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxNodeCount",
		val,
	)
}

func (j *jsiiProxy_CceNodePoolV3)SetMaxPods(val *float64) {
	if err := j.validateSetMaxPodsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxPods",
		val,
	)
}

func (j *jsiiProxy_CceNodePoolV3)SetMinNodeCount(val *float64) {
	if err := j.validateSetMinNodeCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minNodeCount",
		val,
	)
}

func (j *jsiiProxy_CceNodePoolV3)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CceNodePoolV3)SetOs(val *string) {
	if err := j.validateSetOsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"os",
		val,
	)
}

func (j *jsiiProxy_CceNodePoolV3)SetPassword(val *string) {
	if err := j.validateSetPasswordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"password",
		val,
	)
}

func (j *jsiiProxy_CceNodePoolV3)SetPostinstall(val *string) {
	if err := j.validateSetPostinstallParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"postinstall",
		val,
	)
}

func (j *jsiiProxy_CceNodePoolV3)SetPreinstall(val *string) {
	if err := j.validateSetPreinstallParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"preinstall",
		val,
	)
}

func (j *jsiiProxy_CceNodePoolV3)SetPriority(val *float64) {
	if err := j.validateSetPriorityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"priority",
		val,
	)
}

func (j *jsiiProxy_CceNodePoolV3)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_CceNodePoolV3)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_CceNodePoolV3)SetRuntime(val *string) {
	if err := j.validateSetRuntimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"runtime",
		val,
	)
}

func (j *jsiiProxy_CceNodePoolV3)SetScaleDownCooldownTime(val *float64) {
	if err := j.validateSetScaleDownCooldownTimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scaleDownCooldownTime",
		val,
	)
}

func (j *jsiiProxy_CceNodePoolV3)SetScaleEnable(val interface{}) {
	if err := j.validateSetScaleEnableParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scaleEnable",
		val,
	)
}

func (j *jsiiProxy_CceNodePoolV3)SetServerGroupReference(val *string) {
	if err := j.validateSetServerGroupReferenceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serverGroupReference",
		val,
	)
}

func (j *jsiiProxy_CceNodePoolV3)SetSubnetId(val *string) {
	if err := j.validateSetSubnetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"subnetId",
		val,
	)
}

func (j *jsiiProxy_CceNodePoolV3)SetUserTags(val *map[string]*string) {
	if err := j.validateSetUserTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"userTags",
		val,
	)
}

// Generates CDKTF code for importing a CceNodePoolV3 resource upon running "cdktf plan <stack-name>".
func CceNodePoolV3_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateCceNodePoolV3_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-opentelekomcloud.cceNodePoolV3.CceNodePoolV3",
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
func CceNodePoolV3_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCceNodePoolV3_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-opentelekomcloud.cceNodePoolV3.CceNodePoolV3",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func CceNodePoolV3_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCceNodePoolV3_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-opentelekomcloud.cceNodePoolV3.CceNodePoolV3",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func CceNodePoolV3_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCceNodePoolV3_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-opentelekomcloud.cceNodePoolV3.CceNodePoolV3",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CceNodePoolV3_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-opentelekomcloud.cceNodePoolV3.CceNodePoolV3",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CceNodePoolV3) AddMoveTarget(moveTarget *string) {
	if err := c.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (c *jsiiProxy_CceNodePoolV3) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CceNodePoolV3) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := c.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CceNodePoolV3) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := c.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CceNodePoolV3) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := c.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CceNodePoolV3) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := c.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CceNodePoolV3) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := c.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CceNodePoolV3) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := c.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CceNodePoolV3) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := c.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CceNodePoolV3) GetStringAttribute(terraformAttribute *string) *string {
	if err := c.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CceNodePoolV3) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := c.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CceNodePoolV3) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CceNodePoolV3) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := c.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (c *jsiiProxy_CceNodePoolV3) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := c.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CceNodePoolV3) MoveFromId(id *string) {
	if err := c.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"moveFromId",
		[]interface{}{id},
	)
}

func (c *jsiiProxy_CceNodePoolV3) MoveTo(moveTarget *string, index interface{}) {
	if err := c.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (c *jsiiProxy_CceNodePoolV3) MoveToId(id *string) {
	if err := c.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"moveToId",
		[]interface{}{id},
	)
}

func (c *jsiiProxy_CceNodePoolV3) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CceNodePoolV3) PutDataVolumes(value interface{}) {
	if err := c.validatePutDataVolumesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putDataVolumes",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CceNodePoolV3) PutRootVolume(value *CceNodePoolV3RootVolume) {
	if err := c.validatePutRootVolumeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putRootVolume",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CceNodePoolV3) PutTaints(value interface{}) {
	if err := c.validatePutTaintsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putTaints",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CceNodePoolV3) PutTimeouts(value *CceNodePoolV3Timeouts) {
	if err := c.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CceNodePoolV3) ResetAgencyName() {
	_jsii_.InvokeVoid(
		c,
		"resetAgencyName",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CceNodePoolV3) ResetAvailabilityZone() {
	_jsii_.InvokeVoid(
		c,
		"resetAvailabilityZone",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CceNodePoolV3) ResetDockerBaseSize() {
	_jsii_.InvokeVoid(
		c,
		"resetDockerBaseSize",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CceNodePoolV3) ResetDockerLvmConfigOverride() {
	_jsii_.InvokeVoid(
		c,
		"resetDockerLvmConfigOverride",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CceNodePoolV3) ResetId() {
	_jsii_.InvokeVoid(
		c,
		"resetId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CceNodePoolV3) ResetK8STags() {
	_jsii_.InvokeVoid(
		c,
		"resetK8STags",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CceNodePoolV3) ResetKeyPair() {
	_jsii_.InvokeVoid(
		c,
		"resetKeyPair",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CceNodePoolV3) ResetMaxNodeCount() {
	_jsii_.InvokeVoid(
		c,
		"resetMaxNodeCount",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CceNodePoolV3) ResetMaxPods() {
	_jsii_.InvokeVoid(
		c,
		"resetMaxPods",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CceNodePoolV3) ResetMinNodeCount() {
	_jsii_.InvokeVoid(
		c,
		"resetMinNodeCount",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CceNodePoolV3) ResetOs() {
	_jsii_.InvokeVoid(
		c,
		"resetOs",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CceNodePoolV3) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		c,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CceNodePoolV3) ResetPassword() {
	_jsii_.InvokeVoid(
		c,
		"resetPassword",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CceNodePoolV3) ResetPostinstall() {
	_jsii_.InvokeVoid(
		c,
		"resetPostinstall",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CceNodePoolV3) ResetPreinstall() {
	_jsii_.InvokeVoid(
		c,
		"resetPreinstall",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CceNodePoolV3) ResetPriority() {
	_jsii_.InvokeVoid(
		c,
		"resetPriority",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CceNodePoolV3) ResetRuntime() {
	_jsii_.InvokeVoid(
		c,
		"resetRuntime",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CceNodePoolV3) ResetScaleDownCooldownTime() {
	_jsii_.InvokeVoid(
		c,
		"resetScaleDownCooldownTime",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CceNodePoolV3) ResetScaleEnable() {
	_jsii_.InvokeVoid(
		c,
		"resetScaleEnable",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CceNodePoolV3) ResetServerGroupReference() {
	_jsii_.InvokeVoid(
		c,
		"resetServerGroupReference",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CceNodePoolV3) ResetSubnetId() {
	_jsii_.InvokeVoid(
		c,
		"resetSubnetId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CceNodePoolV3) ResetTaints() {
	_jsii_.InvokeVoid(
		c,
		"resetTaints",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CceNodePoolV3) ResetTimeouts() {
	_jsii_.InvokeVoid(
		c,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CceNodePoolV3) ResetUserTags() {
	_jsii_.InvokeVoid(
		c,
		"resetUserTags",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CceNodePoolV3) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CceNodePoolV3) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CceNodePoolV3) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CceNodePoolV3) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CceNodePoolV3) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CceNodePoolV3) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

