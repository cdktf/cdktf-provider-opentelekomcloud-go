// Prebuilt opentelekomcloud Provider for Terraform CDK (cdktf)
package opentelekomcloud

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v2/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://www.terraform.io/docs/providers/opentelekomcloud/r/compute_instance_v2 opentelekomcloud_compute_instance_v2}.
type ComputeInstanceV2 interface {
	cdktf.TerraformResource
	AccessIpV4() *string
	SetAccessIpV4(val *string)
	AccessIpV4Input() *string
	AccessIpV6() *string
	SetAccessIpV6(val *string)
	AccessIpV6Input() *string
	AdminPass() *string
	SetAdminPass(val *string)
	AdminPassInput() *string
	AllMetadata() cdktf.StringMap
	AutoRecovery() interface{}
	SetAutoRecovery(val interface{})
	AutoRecoveryInput() interface{}
	AvailabilityZone() *string
	SetAvailabilityZone(val *string)
	AvailabilityZoneInput() *string
	BlockDevice() ComputeInstanceV2BlockDeviceList
	BlockDeviceInput() interface{}
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	ConfigDrive() interface{}
	SetConfigDrive(val interface{})
	ConfigDriveInput() interface{}
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
	FlavorId() *string
	SetFlavorId(val *string)
	FlavorIdInput() *string
	FlavorName() *string
	SetFlavorName(val *string)
	FlavorNameInput() *string
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
	ImageId() *string
	SetImageId(val *string)
	ImageIdInput() *string
	ImageName() *string
	SetImageName(val *string)
	ImageNameInput() *string
	KeyPair() *string
	SetKeyPair(val *string)
	KeyPairInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Metadata() *map[string]*string
	SetMetadata(val *map[string]*string)
	MetadataInput() *map[string]*string
	Name() *string
	SetName(val *string)
	NameInput() *string
	Network() ComputeInstanceV2NetworkList
	NetworkInput() interface{}
	// The tree node.
	Node() constructs.Node
	PowerState() *string
	SetPowerState(val *string)
	PowerStateInput() *string
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
	SchedulerHints() ComputeInstanceV2SchedulerHintsList
	SchedulerHintsInput() interface{}
	SecurityGroups() *[]*string
	SetSecurityGroups(val *[]*string)
	SecurityGroupsInput() *[]*string
	StopBeforeDestroy() interface{}
	SetStopBeforeDestroy(val interface{})
	StopBeforeDestroyInput() interface{}
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsInput() *map[string]*string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() ComputeInstanceV2TimeoutsOutputReference
	TimeoutsInput() interface{}
	UserData() *string
	SetUserData(val *string)
	UserDataInput() *string
	VolumeAttached() ComputeInstanceV2VolumeAttachedList
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
	PutBlockDevice(value interface{})
	PutNetwork(value interface{})
	PutSchedulerHints(value interface{})
	PutTimeouts(value *ComputeInstanceV2Timeouts)
	ResetAccessIpV4()
	ResetAccessIpV6()
	ResetAdminPass()
	ResetAutoRecovery()
	ResetAvailabilityZone()
	ResetBlockDevice()
	ResetConfigDrive()
	ResetFlavorId()
	ResetFlavorName()
	ResetId()
	ResetImageId()
	ResetImageName()
	ResetKeyPair()
	ResetMetadata()
	ResetNetwork()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPowerState()
	ResetRegion()
	ResetSchedulerHints()
	ResetSecurityGroups()
	ResetStopBeforeDestroy()
	ResetTags()
	ResetTimeouts()
	ResetUserData()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for ComputeInstanceV2
type jsiiProxy_ComputeInstanceV2 struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_ComputeInstanceV2) AccessIpV4() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessIpV4",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceV2) AccessIpV4Input() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessIpV4Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceV2) AccessIpV6() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessIpV6",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceV2) AccessIpV6Input() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessIpV6Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceV2) AdminPass() *string {
	var returns *string
	_jsii_.Get(
		j,
		"adminPass",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceV2) AdminPassInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"adminPassInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceV2) AllMetadata() cdktf.StringMap {
	var returns cdktf.StringMap
	_jsii_.Get(
		j,
		"allMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceV2) AutoRecovery() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoRecovery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceV2) AutoRecoveryInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoRecoveryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceV2) AvailabilityZone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"availabilityZone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceV2) AvailabilityZoneInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"availabilityZoneInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceV2) BlockDevice() ComputeInstanceV2BlockDeviceList {
	var returns ComputeInstanceV2BlockDeviceList
	_jsii_.Get(
		j,
		"blockDevice",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceV2) BlockDeviceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"blockDeviceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceV2) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceV2) ConfigDrive() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"configDrive",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceV2) ConfigDriveInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"configDriveInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceV2) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceV2) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceV2) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceV2) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceV2) FlavorId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"flavorId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceV2) FlavorIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"flavorIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceV2) FlavorName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"flavorName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceV2) FlavorNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"flavorNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceV2) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceV2) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceV2) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceV2) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceV2) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceV2) ImageId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceV2) ImageIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceV2) ImageName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceV2) ImageNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceV2) KeyPair() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyPair",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceV2) KeyPairInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyPairInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceV2) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceV2) Metadata() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"metadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceV2) MetadataInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"metadataInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceV2) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceV2) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceV2) Network() ComputeInstanceV2NetworkList {
	var returns ComputeInstanceV2NetworkList
	_jsii_.Get(
		j,
		"network",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceV2) NetworkInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"networkInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceV2) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceV2) PowerState() *string {
	var returns *string
	_jsii_.Get(
		j,
		"powerState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceV2) PowerStateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"powerStateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceV2) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceV2) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceV2) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceV2) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceV2) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceV2) SchedulerHints() ComputeInstanceV2SchedulerHintsList {
	var returns ComputeInstanceV2SchedulerHintsList
	_jsii_.Get(
		j,
		"schedulerHints",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceV2) SchedulerHintsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"schedulerHintsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceV2) SecurityGroups() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceV2) SecurityGroupsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityGroupsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceV2) StopBeforeDestroy() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"stopBeforeDestroy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceV2) StopBeforeDestroyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"stopBeforeDestroyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceV2) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceV2) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceV2) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceV2) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceV2) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceV2) Timeouts() ComputeInstanceV2TimeoutsOutputReference {
	var returns ComputeInstanceV2TimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceV2) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceV2) UserData() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userData",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceV2) UserDataInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userDataInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceV2) VolumeAttached() ComputeInstanceV2VolumeAttachedList {
	var returns ComputeInstanceV2VolumeAttachedList
	_jsii_.Get(
		j,
		"volumeAttached",
		&returns,
	)
	return returns
}


// Create a new {@link https://www.terraform.io/docs/providers/opentelekomcloud/r/compute_instance_v2 opentelekomcloud_compute_instance_v2} Resource.
func NewComputeInstanceV2(scope constructs.Construct, id *string, config *ComputeInstanceV2Config) ComputeInstanceV2 {
	_init_.Initialize()

	j := jsiiProxy_ComputeInstanceV2{}

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.ComputeInstanceV2",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/opentelekomcloud/r/compute_instance_v2 opentelekomcloud_compute_instance_v2} Resource.
func NewComputeInstanceV2_Override(c ComputeInstanceV2, scope constructs.Construct, id *string, config *ComputeInstanceV2Config) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.ComputeInstanceV2",
		[]interface{}{scope, id, config},
		c,
	)
}

func (j *jsiiProxy_ComputeInstanceV2) SetAccessIpV4(val *string) {
	_jsii_.Set(
		j,
		"accessIpV4",
		val,
	)
}

func (j *jsiiProxy_ComputeInstanceV2) SetAccessIpV6(val *string) {
	_jsii_.Set(
		j,
		"accessIpV6",
		val,
	)
}

func (j *jsiiProxy_ComputeInstanceV2) SetAdminPass(val *string) {
	_jsii_.Set(
		j,
		"adminPass",
		val,
	)
}

func (j *jsiiProxy_ComputeInstanceV2) SetAutoRecovery(val interface{}) {
	_jsii_.Set(
		j,
		"autoRecovery",
		val,
	)
}

func (j *jsiiProxy_ComputeInstanceV2) SetAvailabilityZone(val *string) {
	_jsii_.Set(
		j,
		"availabilityZone",
		val,
	)
}

func (j *jsiiProxy_ComputeInstanceV2) SetConfigDrive(val interface{}) {
	_jsii_.Set(
		j,
		"configDrive",
		val,
	)
}

func (j *jsiiProxy_ComputeInstanceV2) SetConnection(val interface{}) {
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_ComputeInstanceV2) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_ComputeInstanceV2) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_ComputeInstanceV2) SetFlavorId(val *string) {
	_jsii_.Set(
		j,
		"flavorId",
		val,
	)
}

func (j *jsiiProxy_ComputeInstanceV2) SetFlavorName(val *string) {
	_jsii_.Set(
		j,
		"flavorName",
		val,
	)
}

func (j *jsiiProxy_ComputeInstanceV2) SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_ComputeInstanceV2) SetId(val *string) {
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_ComputeInstanceV2) SetImageId(val *string) {
	_jsii_.Set(
		j,
		"imageId",
		val,
	)
}

func (j *jsiiProxy_ComputeInstanceV2) SetImageName(val *string) {
	_jsii_.Set(
		j,
		"imageName",
		val,
	)
}

func (j *jsiiProxy_ComputeInstanceV2) SetKeyPair(val *string) {
	_jsii_.Set(
		j,
		"keyPair",
		val,
	)
}

func (j *jsiiProxy_ComputeInstanceV2) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_ComputeInstanceV2) SetMetadata(val *map[string]*string) {
	_jsii_.Set(
		j,
		"metadata",
		val,
	)
}

func (j *jsiiProxy_ComputeInstanceV2) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_ComputeInstanceV2) SetPowerState(val *string) {
	_jsii_.Set(
		j,
		"powerState",
		val,
	)
}

func (j *jsiiProxy_ComputeInstanceV2) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_ComputeInstanceV2) SetProvisioners(val *[]interface{}) {
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_ComputeInstanceV2) SetRegion(val *string) {
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_ComputeInstanceV2) SetSecurityGroups(val *[]*string) {
	_jsii_.Set(
		j,
		"securityGroups",
		val,
	)
}

func (j *jsiiProxy_ComputeInstanceV2) SetStopBeforeDestroy(val interface{}) {
	_jsii_.Set(
		j,
		"stopBeforeDestroy",
		val,
	)
}

func (j *jsiiProxy_ComputeInstanceV2) SetTags(val *map[string]*string) {
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_ComputeInstanceV2) SetUserData(val *string) {
	_jsii_.Set(
		j,
		"userData",
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
func ComputeInstanceV2_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-opentelekomcloud.ComputeInstanceV2",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func ComputeInstanceV2_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-opentelekomcloud.ComputeInstanceV2",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_ComputeInstanceV2) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_ComputeInstanceV2) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeInstanceV2) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeInstanceV2) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeInstanceV2) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeInstanceV2) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeInstanceV2) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeInstanceV2) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeInstanceV2) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeInstanceV2) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeInstanceV2) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeInstanceV2) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_ComputeInstanceV2) PutBlockDevice(value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"putBlockDevice",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComputeInstanceV2) PutNetwork(value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"putNetwork",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComputeInstanceV2) PutSchedulerHints(value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"putSchedulerHints",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComputeInstanceV2) PutTimeouts(value *ComputeInstanceV2Timeouts) {
	_jsii_.InvokeVoid(
		c,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComputeInstanceV2) ResetAccessIpV4() {
	_jsii_.InvokeVoid(
		c,
		"resetAccessIpV4",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeInstanceV2) ResetAccessIpV6() {
	_jsii_.InvokeVoid(
		c,
		"resetAccessIpV6",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeInstanceV2) ResetAdminPass() {
	_jsii_.InvokeVoid(
		c,
		"resetAdminPass",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeInstanceV2) ResetAutoRecovery() {
	_jsii_.InvokeVoid(
		c,
		"resetAutoRecovery",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeInstanceV2) ResetAvailabilityZone() {
	_jsii_.InvokeVoid(
		c,
		"resetAvailabilityZone",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeInstanceV2) ResetBlockDevice() {
	_jsii_.InvokeVoid(
		c,
		"resetBlockDevice",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeInstanceV2) ResetConfigDrive() {
	_jsii_.InvokeVoid(
		c,
		"resetConfigDrive",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeInstanceV2) ResetFlavorId() {
	_jsii_.InvokeVoid(
		c,
		"resetFlavorId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeInstanceV2) ResetFlavorName() {
	_jsii_.InvokeVoid(
		c,
		"resetFlavorName",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeInstanceV2) ResetId() {
	_jsii_.InvokeVoid(
		c,
		"resetId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeInstanceV2) ResetImageId() {
	_jsii_.InvokeVoid(
		c,
		"resetImageId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeInstanceV2) ResetImageName() {
	_jsii_.InvokeVoid(
		c,
		"resetImageName",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeInstanceV2) ResetKeyPair() {
	_jsii_.InvokeVoid(
		c,
		"resetKeyPair",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeInstanceV2) ResetMetadata() {
	_jsii_.InvokeVoid(
		c,
		"resetMetadata",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeInstanceV2) ResetNetwork() {
	_jsii_.InvokeVoid(
		c,
		"resetNetwork",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeInstanceV2) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		c,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeInstanceV2) ResetPowerState() {
	_jsii_.InvokeVoid(
		c,
		"resetPowerState",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeInstanceV2) ResetRegion() {
	_jsii_.InvokeVoid(
		c,
		"resetRegion",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeInstanceV2) ResetSchedulerHints() {
	_jsii_.InvokeVoid(
		c,
		"resetSchedulerHints",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeInstanceV2) ResetSecurityGroups() {
	_jsii_.InvokeVoid(
		c,
		"resetSecurityGroups",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeInstanceV2) ResetStopBeforeDestroy() {
	_jsii_.InvokeVoid(
		c,
		"resetStopBeforeDestroy",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeInstanceV2) ResetTags() {
	_jsii_.InvokeVoid(
		c,
		"resetTags",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeInstanceV2) ResetTimeouts() {
	_jsii_.InvokeVoid(
		c,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeInstanceV2) ResetUserData() {
	_jsii_.InvokeVoid(
		c,
		"resetUserData",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeInstanceV2) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeInstanceV2) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeInstanceV2) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeInstanceV2) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}
