// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package imagesimagev2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v12/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v12/imagesimagev2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.44/docs/resources/images_image_v2 opentelekomcloud_images_image_v2}.
type ImagesImageV2 interface {
	cdktf.TerraformResource
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	Checksum() *string
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	ContainerFormat() *string
	SetContainerFormat(val *string)
	ContainerFormatInput() *string
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	CreatedAt() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	DiskFormat() *string
	SetDiskFormat(val *string)
	DiskFormatInput() *string
	File() *string
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	HwFirmwareType() *string
	SetHwFirmwareType(val *string)
	HwFirmwareTypeInput() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	ImageCachePath() *string
	SetImageCachePath(val *string)
	ImageCachePathInput() *string
	ImageSourceUrl() *string
	SetImageSourceUrl(val *string)
	ImageSourceUrlInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	LocalFilePath() *string
	SetLocalFilePath(val *string)
	LocalFilePathInput() *string
	MinDiskGb() *float64
	SetMinDiskGb(val *float64)
	MinDiskGbInput() *float64
	MinRamMb() *float64
	SetMinRamMb(val *float64)
	MinRamMbInput() *float64
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	Owner() *string
	Protected() interface{}
	SetProtected(val interface{})
	ProtectedInput() interface{}
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
	Schema() *string
	SizeBytes() *float64
	Status() *string
	Tags() *[]*string
	SetTags(val *[]*string)
	TagsInput() *[]*string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() ImagesImageV2TimeoutsOutputReference
	TimeoutsInput() interface{}
	UpdateAt() *string
	Visibility() *string
	SetVisibility(val *string)
	VisibilityInput() *string
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
	PutTimeouts(value *ImagesImageV2Timeouts)
	ResetHwFirmwareType()
	ResetId()
	ResetImageCachePath()
	ResetImageSourceUrl()
	ResetLocalFilePath()
	ResetMinDiskGb()
	ResetMinRamMb()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetProtected()
	ResetRegion()
	ResetTags()
	ResetTimeouts()
	ResetVisibility()
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

// The jsii proxy struct for ImagesImageV2
type jsiiProxy_ImagesImageV2 struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_ImagesImageV2) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagesImageV2) Checksum() *string {
	var returns *string
	_jsii_.Get(
		j,
		"checksum",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagesImageV2) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagesImageV2) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagesImageV2) ContainerFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"containerFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagesImageV2) ContainerFormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"containerFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagesImageV2) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagesImageV2) CreatedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagesImageV2) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagesImageV2) DiskFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"diskFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagesImageV2) DiskFormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"diskFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagesImageV2) File() *string {
	var returns *string
	_jsii_.Get(
		j,
		"file",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagesImageV2) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagesImageV2) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagesImageV2) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagesImageV2) HwFirmwareType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hwFirmwareType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagesImageV2) HwFirmwareTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hwFirmwareTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagesImageV2) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagesImageV2) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagesImageV2) ImageCachePath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageCachePath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagesImageV2) ImageCachePathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageCachePathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagesImageV2) ImageSourceUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageSourceUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagesImageV2) ImageSourceUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageSourceUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagesImageV2) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagesImageV2) LocalFilePath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"localFilePath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagesImageV2) LocalFilePathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"localFilePathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagesImageV2) MinDiskGb() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minDiskGb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagesImageV2) MinDiskGbInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minDiskGbInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagesImageV2) MinRamMb() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minRamMb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagesImageV2) MinRamMbInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minRamMbInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagesImageV2) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagesImageV2) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagesImageV2) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagesImageV2) Owner() *string {
	var returns *string
	_jsii_.Get(
		j,
		"owner",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagesImageV2) Protected() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"protected",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagesImageV2) ProtectedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"protectedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagesImageV2) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagesImageV2) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagesImageV2) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagesImageV2) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagesImageV2) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagesImageV2) Schema() *string {
	var returns *string
	_jsii_.Get(
		j,
		"schema",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagesImageV2) SizeBytes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sizeBytes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagesImageV2) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagesImageV2) Tags() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagesImageV2) TagsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagesImageV2) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagesImageV2) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagesImageV2) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagesImageV2) Timeouts() ImagesImageV2TimeoutsOutputReference {
	var returns ImagesImageV2TimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagesImageV2) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagesImageV2) UpdateAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagesImageV2) Visibility() *string {
	var returns *string
	_jsii_.Get(
		j,
		"visibility",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImagesImageV2) VisibilityInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"visibilityInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.44/docs/resources/images_image_v2 opentelekomcloud_images_image_v2} Resource.
func NewImagesImageV2(scope constructs.Construct, id *string, config *ImagesImageV2Config) ImagesImageV2 {
	_init_.Initialize()

	if err := validateNewImagesImageV2Parameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_ImagesImageV2{}

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.imagesImageV2.ImagesImageV2",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.44/docs/resources/images_image_v2 opentelekomcloud_images_image_v2} Resource.
func NewImagesImageV2_Override(i ImagesImageV2, scope constructs.Construct, id *string, config *ImagesImageV2Config) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.imagesImageV2.ImagesImageV2",
		[]interface{}{scope, id, config},
		i,
	)
}

func (j *jsiiProxy_ImagesImageV2)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_ImagesImageV2)SetContainerFormat(val *string) {
	if err := j.validateSetContainerFormatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"containerFormat",
		val,
	)
}

func (j *jsiiProxy_ImagesImageV2)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_ImagesImageV2)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_ImagesImageV2)SetDiskFormat(val *string) {
	if err := j.validateSetDiskFormatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"diskFormat",
		val,
	)
}

func (j *jsiiProxy_ImagesImageV2)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_ImagesImageV2)SetHwFirmwareType(val *string) {
	if err := j.validateSetHwFirmwareTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hwFirmwareType",
		val,
	)
}

func (j *jsiiProxy_ImagesImageV2)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_ImagesImageV2)SetImageCachePath(val *string) {
	if err := j.validateSetImageCachePathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"imageCachePath",
		val,
	)
}

func (j *jsiiProxy_ImagesImageV2)SetImageSourceUrl(val *string) {
	if err := j.validateSetImageSourceUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"imageSourceUrl",
		val,
	)
}

func (j *jsiiProxy_ImagesImageV2)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_ImagesImageV2)SetLocalFilePath(val *string) {
	if err := j.validateSetLocalFilePathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"localFilePath",
		val,
	)
}

func (j *jsiiProxy_ImagesImageV2)SetMinDiskGb(val *float64) {
	if err := j.validateSetMinDiskGbParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minDiskGb",
		val,
	)
}

func (j *jsiiProxy_ImagesImageV2)SetMinRamMb(val *float64) {
	if err := j.validateSetMinRamMbParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minRamMb",
		val,
	)
}

func (j *jsiiProxy_ImagesImageV2)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_ImagesImageV2)SetProtected(val interface{}) {
	if err := j.validateSetProtectedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"protected",
		val,
	)
}

func (j *jsiiProxy_ImagesImageV2)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_ImagesImageV2)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_ImagesImageV2)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_ImagesImageV2)SetTags(val *[]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_ImagesImageV2)SetVisibility(val *string) {
	if err := j.validateSetVisibilityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"visibility",
		val,
	)
}

// Generates CDKTF code for importing a ImagesImageV2 resource upon running "cdktf plan <stack-name>".
func ImagesImageV2_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateImagesImageV2_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-opentelekomcloud.imagesImageV2.ImagesImageV2",
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
func ImagesImageV2_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateImagesImageV2_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-opentelekomcloud.imagesImageV2.ImagesImageV2",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ImagesImageV2_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateImagesImageV2_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-opentelekomcloud.imagesImageV2.ImagesImageV2",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ImagesImageV2_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateImagesImageV2_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-opentelekomcloud.imagesImageV2.ImagesImageV2",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func ImagesImageV2_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-opentelekomcloud.imagesImageV2.ImagesImageV2",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (i *jsiiProxy_ImagesImageV2) AddMoveTarget(moveTarget *string) {
	if err := i.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (i *jsiiProxy_ImagesImageV2) AddOverride(path *string, value interface{}) {
	if err := i.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (i *jsiiProxy_ImagesImageV2) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := i.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		i,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ImagesImageV2) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := i.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		i,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ImagesImageV2) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := i.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		i,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ImagesImageV2) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := i.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		i,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ImagesImageV2) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := i.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		i,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ImagesImageV2) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := i.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		i,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ImagesImageV2) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := i.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		i,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ImagesImageV2) GetStringAttribute(terraformAttribute *string) *string {
	if err := i.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		i,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ImagesImageV2) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := i.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		i,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ImagesImageV2) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ImagesImageV2) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := i.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (i *jsiiProxy_ImagesImageV2) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := i.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ImagesImageV2) MoveFromId(id *string) {
	if err := i.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"moveFromId",
		[]interface{}{id},
	)
}

func (i *jsiiProxy_ImagesImageV2) MoveTo(moveTarget *string, index interface{}) {
	if err := i.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (i *jsiiProxy_ImagesImageV2) MoveToId(id *string) {
	if err := i.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"moveToId",
		[]interface{}{id},
	)
}

func (i *jsiiProxy_ImagesImageV2) OverrideLogicalId(newLogicalId *string) {
	if err := i.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (i *jsiiProxy_ImagesImageV2) PutTimeouts(value *ImagesImageV2Timeouts) {
	if err := i.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_ImagesImageV2) ResetHwFirmwareType() {
	_jsii_.InvokeVoid(
		i,
		"resetHwFirmwareType",
		nil, // no parameters
	)
}

func (i *jsiiProxy_ImagesImageV2) ResetId() {
	_jsii_.InvokeVoid(
		i,
		"resetId",
		nil, // no parameters
	)
}

func (i *jsiiProxy_ImagesImageV2) ResetImageCachePath() {
	_jsii_.InvokeVoid(
		i,
		"resetImageCachePath",
		nil, // no parameters
	)
}

func (i *jsiiProxy_ImagesImageV2) ResetImageSourceUrl() {
	_jsii_.InvokeVoid(
		i,
		"resetImageSourceUrl",
		nil, // no parameters
	)
}

func (i *jsiiProxy_ImagesImageV2) ResetLocalFilePath() {
	_jsii_.InvokeVoid(
		i,
		"resetLocalFilePath",
		nil, // no parameters
	)
}

func (i *jsiiProxy_ImagesImageV2) ResetMinDiskGb() {
	_jsii_.InvokeVoid(
		i,
		"resetMinDiskGb",
		nil, // no parameters
	)
}

func (i *jsiiProxy_ImagesImageV2) ResetMinRamMb() {
	_jsii_.InvokeVoid(
		i,
		"resetMinRamMb",
		nil, // no parameters
	)
}

func (i *jsiiProxy_ImagesImageV2) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		i,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (i *jsiiProxy_ImagesImageV2) ResetProtected() {
	_jsii_.InvokeVoid(
		i,
		"resetProtected",
		nil, // no parameters
	)
}

func (i *jsiiProxy_ImagesImageV2) ResetRegion() {
	_jsii_.InvokeVoid(
		i,
		"resetRegion",
		nil, // no parameters
	)
}

func (i *jsiiProxy_ImagesImageV2) ResetTags() {
	_jsii_.InvokeVoid(
		i,
		"resetTags",
		nil, // no parameters
	)
}

func (i *jsiiProxy_ImagesImageV2) ResetTimeouts() {
	_jsii_.InvokeVoid(
		i,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (i *jsiiProxy_ImagesImageV2) ResetVisibility() {
	_jsii_.InvokeVoid(
		i,
		"resetVisibility",
		nil, // no parameters
	)
}

func (i *jsiiProxy_ImagesImageV2) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		i,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ImagesImageV2) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		i,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ImagesImageV2) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ImagesImageV2) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ImagesImageV2) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ImagesImageV2) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

