// Prebuilt opentelekomcloud Provider for Terraform CDK (cdktf)
package opentelekomcloud

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v2/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://www.terraform.io/docs/providers/opentelekomcloud/r/blockstorage_volume_v2 opentelekomcloud_blockstorage_volume_v2}.
type BlockstorageVolumeV2 interface {
	cdktf.TerraformResource
	Attachment() BlockstorageVolumeV2AttachmentList
	AvailabilityZone() *string
	SetAvailabilityZone(val *string)
	AvailabilityZoneInput() *string
	Cascade() interface{}
	SetCascade(val interface{})
	CascadeInput() interface{}
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	ConsistencyGroupId() *string
	SetConsistencyGroupId(val *string)
	ConsistencyGroupIdInput() *string
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
	DeviceType() *string
	SetDeviceType(val *string)
	DeviceTypeInput() *string
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
	Size() *float64
	SetSize(val *float64)
	SizeInput() *float64
	SnapshotId() *string
	SetSnapshotId(val *string)
	SnapshotIdInput() *string
	SourceReplica() *string
	SetSourceReplica(val *string)
	SourceReplicaInput() *string
	SourceVolId() *string
	SetSourceVolId(val *string)
	SourceVolIdInput() *string
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsInput() *map[string]*string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() BlockstorageVolumeV2TimeoutsOutputReference
	TimeoutsInput() interface{}
	VolumeType() *string
	SetVolumeType(val *string)
	VolumeTypeInput() *string
	Wwn() *string
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
	PutTimeouts(value *BlockstorageVolumeV2Timeouts)
	ResetAvailabilityZone()
	ResetCascade()
	ResetConsistencyGroupId()
	ResetDescription()
	ResetDeviceType()
	ResetId()
	ResetImageId()
	ResetMetadata()
	ResetName()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetRegion()
	ResetSnapshotId()
	ResetSourceReplica()
	ResetSourceVolId()
	ResetTags()
	ResetTimeouts()
	ResetVolumeType()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for BlockstorageVolumeV2
type jsiiProxy_BlockstorageVolumeV2 struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_BlockstorageVolumeV2) Attachment() BlockstorageVolumeV2AttachmentList {
	var returns BlockstorageVolumeV2AttachmentList
	_jsii_.Get(
		j,
		"attachment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BlockstorageVolumeV2) AvailabilityZone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"availabilityZone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BlockstorageVolumeV2) AvailabilityZoneInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"availabilityZoneInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BlockstorageVolumeV2) Cascade() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cascade",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BlockstorageVolumeV2) CascadeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cascadeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BlockstorageVolumeV2) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BlockstorageVolumeV2) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BlockstorageVolumeV2) ConsistencyGroupId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"consistencyGroupId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BlockstorageVolumeV2) ConsistencyGroupIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"consistencyGroupIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BlockstorageVolumeV2) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BlockstorageVolumeV2) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BlockstorageVolumeV2) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BlockstorageVolumeV2) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BlockstorageVolumeV2) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BlockstorageVolumeV2) DeviceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deviceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BlockstorageVolumeV2) DeviceTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deviceTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BlockstorageVolumeV2) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BlockstorageVolumeV2) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BlockstorageVolumeV2) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BlockstorageVolumeV2) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BlockstorageVolumeV2) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BlockstorageVolumeV2) ImageId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BlockstorageVolumeV2) ImageIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BlockstorageVolumeV2) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BlockstorageVolumeV2) Metadata() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"metadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BlockstorageVolumeV2) MetadataInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"metadataInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BlockstorageVolumeV2) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BlockstorageVolumeV2) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BlockstorageVolumeV2) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BlockstorageVolumeV2) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BlockstorageVolumeV2) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BlockstorageVolumeV2) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BlockstorageVolumeV2) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BlockstorageVolumeV2) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BlockstorageVolumeV2) Size() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"size",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BlockstorageVolumeV2) SizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BlockstorageVolumeV2) SnapshotId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"snapshotId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BlockstorageVolumeV2) SnapshotIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"snapshotIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BlockstorageVolumeV2) SourceReplica() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceReplica",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BlockstorageVolumeV2) SourceReplicaInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceReplicaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BlockstorageVolumeV2) SourceVolId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceVolId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BlockstorageVolumeV2) SourceVolIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceVolIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BlockstorageVolumeV2) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BlockstorageVolumeV2) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BlockstorageVolumeV2) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BlockstorageVolumeV2) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BlockstorageVolumeV2) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BlockstorageVolumeV2) Timeouts() BlockstorageVolumeV2TimeoutsOutputReference {
	var returns BlockstorageVolumeV2TimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BlockstorageVolumeV2) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BlockstorageVolumeV2) VolumeType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"volumeType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BlockstorageVolumeV2) VolumeTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"volumeTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BlockstorageVolumeV2) Wwn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"wwn",
		&returns,
	)
	return returns
}


// Create a new {@link https://www.terraform.io/docs/providers/opentelekomcloud/r/blockstorage_volume_v2 opentelekomcloud_blockstorage_volume_v2} Resource.
func NewBlockstorageVolumeV2(scope constructs.Construct, id *string, config *BlockstorageVolumeV2Config) BlockstorageVolumeV2 {
	_init_.Initialize()

	j := jsiiProxy_BlockstorageVolumeV2{}

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.BlockstorageVolumeV2",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/opentelekomcloud/r/blockstorage_volume_v2 opentelekomcloud_blockstorage_volume_v2} Resource.
func NewBlockstorageVolumeV2_Override(b BlockstorageVolumeV2, scope constructs.Construct, id *string, config *BlockstorageVolumeV2Config) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.BlockstorageVolumeV2",
		[]interface{}{scope, id, config},
		b,
	)
}

func (j *jsiiProxy_BlockstorageVolumeV2) SetAvailabilityZone(val *string) {
	_jsii_.Set(
		j,
		"availabilityZone",
		val,
	)
}

func (j *jsiiProxy_BlockstorageVolumeV2) SetCascade(val interface{}) {
	_jsii_.Set(
		j,
		"cascade",
		val,
	)
}

func (j *jsiiProxy_BlockstorageVolumeV2) SetConnection(val interface{}) {
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_BlockstorageVolumeV2) SetConsistencyGroupId(val *string) {
	_jsii_.Set(
		j,
		"consistencyGroupId",
		val,
	)
}

func (j *jsiiProxy_BlockstorageVolumeV2) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_BlockstorageVolumeV2) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_BlockstorageVolumeV2) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_BlockstorageVolumeV2) SetDeviceType(val *string) {
	_jsii_.Set(
		j,
		"deviceType",
		val,
	)
}

func (j *jsiiProxy_BlockstorageVolumeV2) SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_BlockstorageVolumeV2) SetId(val *string) {
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_BlockstorageVolumeV2) SetImageId(val *string) {
	_jsii_.Set(
		j,
		"imageId",
		val,
	)
}

func (j *jsiiProxy_BlockstorageVolumeV2) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_BlockstorageVolumeV2) SetMetadata(val *map[string]*string) {
	_jsii_.Set(
		j,
		"metadata",
		val,
	)
}

func (j *jsiiProxy_BlockstorageVolumeV2) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_BlockstorageVolumeV2) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_BlockstorageVolumeV2) SetProvisioners(val *[]interface{}) {
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_BlockstorageVolumeV2) SetRegion(val *string) {
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_BlockstorageVolumeV2) SetSize(val *float64) {
	_jsii_.Set(
		j,
		"size",
		val,
	)
}

func (j *jsiiProxy_BlockstorageVolumeV2) SetSnapshotId(val *string) {
	_jsii_.Set(
		j,
		"snapshotId",
		val,
	)
}

func (j *jsiiProxy_BlockstorageVolumeV2) SetSourceReplica(val *string) {
	_jsii_.Set(
		j,
		"sourceReplica",
		val,
	)
}

func (j *jsiiProxy_BlockstorageVolumeV2) SetSourceVolId(val *string) {
	_jsii_.Set(
		j,
		"sourceVolId",
		val,
	)
}

func (j *jsiiProxy_BlockstorageVolumeV2) SetTags(val *map[string]*string) {
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_BlockstorageVolumeV2) SetVolumeType(val *string) {
	_jsii_.Set(
		j,
		"volumeType",
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
func BlockstorageVolumeV2_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-opentelekomcloud.BlockstorageVolumeV2",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func BlockstorageVolumeV2_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-opentelekomcloud.BlockstorageVolumeV2",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (b *jsiiProxy_BlockstorageVolumeV2) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		b,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (b *jsiiProxy_BlockstorageVolumeV2) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		b,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BlockstorageVolumeV2) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BlockstorageVolumeV2) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		b,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BlockstorageVolumeV2) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		b,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BlockstorageVolumeV2) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		b,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BlockstorageVolumeV2) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		b,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BlockstorageVolumeV2) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		b,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BlockstorageVolumeV2) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BlockstorageVolumeV2) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		b,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BlockstorageVolumeV2) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BlockstorageVolumeV2) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		b,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (b *jsiiProxy_BlockstorageVolumeV2) PutTimeouts(value *BlockstorageVolumeV2Timeouts) {
	_jsii_.InvokeVoid(
		b,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BlockstorageVolumeV2) ResetAvailabilityZone() {
	_jsii_.InvokeVoid(
		b,
		"resetAvailabilityZone",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BlockstorageVolumeV2) ResetCascade() {
	_jsii_.InvokeVoid(
		b,
		"resetCascade",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BlockstorageVolumeV2) ResetConsistencyGroupId() {
	_jsii_.InvokeVoid(
		b,
		"resetConsistencyGroupId",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BlockstorageVolumeV2) ResetDescription() {
	_jsii_.InvokeVoid(
		b,
		"resetDescription",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BlockstorageVolumeV2) ResetDeviceType() {
	_jsii_.InvokeVoid(
		b,
		"resetDeviceType",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BlockstorageVolumeV2) ResetId() {
	_jsii_.InvokeVoid(
		b,
		"resetId",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BlockstorageVolumeV2) ResetImageId() {
	_jsii_.InvokeVoid(
		b,
		"resetImageId",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BlockstorageVolumeV2) ResetMetadata() {
	_jsii_.InvokeVoid(
		b,
		"resetMetadata",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BlockstorageVolumeV2) ResetName() {
	_jsii_.InvokeVoid(
		b,
		"resetName",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BlockstorageVolumeV2) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		b,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BlockstorageVolumeV2) ResetRegion() {
	_jsii_.InvokeVoid(
		b,
		"resetRegion",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BlockstorageVolumeV2) ResetSnapshotId() {
	_jsii_.InvokeVoid(
		b,
		"resetSnapshotId",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BlockstorageVolumeV2) ResetSourceReplica() {
	_jsii_.InvokeVoid(
		b,
		"resetSourceReplica",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BlockstorageVolumeV2) ResetSourceVolId() {
	_jsii_.InvokeVoid(
		b,
		"resetSourceVolId",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BlockstorageVolumeV2) ResetTags() {
	_jsii_.InvokeVoid(
		b,
		"resetTags",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BlockstorageVolumeV2) ResetTimeouts() {
	_jsii_.InvokeVoid(
		b,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BlockstorageVolumeV2) ResetVolumeType() {
	_jsii_.InvokeVoid(
		b,
		"resetVolumeType",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BlockstorageVolumeV2) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		b,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BlockstorageVolumeV2) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		b,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BlockstorageVolumeV2) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BlockstorageVolumeV2) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		b,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

