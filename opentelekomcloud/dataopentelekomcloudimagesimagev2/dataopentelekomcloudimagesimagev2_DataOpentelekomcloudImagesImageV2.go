package dataopentelekomcloudimagesimagev2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v5/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v5/dataopentelekomcloudimagesimagev2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://www.terraform.io/docs/providers/opentelekomcloud/d/images_image_v2 opentelekomcloud_images_image_v2}.
type DataOpentelekomcloudImagesImageV2 interface {
	cdktf.TerraformDataSource
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	Checksum() *string
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	ContainerFormat() *string
	// Experimental.
	Count() *float64
	// Experimental.
	SetCount(val *float64)
	CreatedAt() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	DiskFormat() *string
	File() *string
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
	Metadata() cdktf.StringMap
	MinDiskGb() *float64
	MinRamMb() *float64
	MostRecent() interface{}
	SetMostRecent(val interface{})
	MostRecentInput() interface{}
	Name() *string
	SetName(val *string)
	NameInput() *string
	NameRegex() *string
	SetNameRegex(val *string)
	NameRegexInput() *string
	// The tree node.
	Node() constructs.Node
	Owner() *string
	SetOwner(val *string)
	OwnerInput() *string
	Properties() *map[string]*string
	SetProperties(val *map[string]*string)
	PropertiesInput() *map[string]*string
	Protected() cdktf.IResolvable
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	RawOverrides() interface{}
	Region() *string
	SetRegion(val *string)
	RegionInput() *string
	Schema() *string
	SizeBytes() *float64
	SizeMax() *float64
	SetSizeMax(val *float64)
	SizeMaxInput() *float64
	SizeMin() *float64
	SetSizeMin(val *float64)
	SizeMinInput() *float64
	SortDirection() *string
	SetSortDirection(val *string)
	SortDirectionInput() *string
	SortKey() *string
	SetSortKey(val *string)
	SortKeyInput() *string
	Tag() *string
	SetTag(val *string)
	TagInput() *string
	Tags() *[]*string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	UpdatedAt() *string
	Visibility() *string
	SetVisibility(val *string)
	VisibilityInput() *string
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
	ResetId()
	ResetMostRecent()
	ResetName()
	ResetNameRegex()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetOwner()
	ResetProperties()
	ResetRegion()
	ResetSizeMax()
	ResetSizeMin()
	ResetSortDirection()
	ResetSortKey()
	ResetTag()
	ResetVisibility()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for DataOpentelekomcloudImagesImageV2
type jsiiProxy_DataOpentelekomcloudImagesImageV2 struct {
	internal.Type__cdktfTerraformDataSource
}

func (j *jsiiProxy_DataOpentelekomcloudImagesImageV2) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudImagesImageV2) Checksum() *string {
	var returns *string
	_jsii_.Get(
		j,
		"checksum",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudImagesImageV2) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudImagesImageV2) ContainerFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"containerFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudImagesImageV2) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudImagesImageV2) CreatedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudImagesImageV2) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudImagesImageV2) DiskFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"diskFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudImagesImageV2) File() *string {
	var returns *string
	_jsii_.Get(
		j,
		"file",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudImagesImageV2) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudImagesImageV2) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudImagesImageV2) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudImagesImageV2) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudImagesImageV2) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudImagesImageV2) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudImagesImageV2) Metadata() cdktf.StringMap {
	var returns cdktf.StringMap
	_jsii_.Get(
		j,
		"metadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudImagesImageV2) MinDiskGb() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minDiskGb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudImagesImageV2) MinRamMb() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minRamMb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudImagesImageV2) MostRecent() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"mostRecent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudImagesImageV2) MostRecentInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"mostRecentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudImagesImageV2) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudImagesImageV2) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudImagesImageV2) NameRegex() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameRegex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudImagesImageV2) NameRegexInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameRegexInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudImagesImageV2) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudImagesImageV2) Owner() *string {
	var returns *string
	_jsii_.Get(
		j,
		"owner",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudImagesImageV2) OwnerInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ownerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudImagesImageV2) Properties() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"properties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudImagesImageV2) PropertiesInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"propertiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudImagesImageV2) Protected() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"protected",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudImagesImageV2) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudImagesImageV2) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudImagesImageV2) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudImagesImageV2) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudImagesImageV2) Schema() *string {
	var returns *string
	_jsii_.Get(
		j,
		"schema",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudImagesImageV2) SizeBytes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sizeBytes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudImagesImageV2) SizeMax() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sizeMax",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudImagesImageV2) SizeMaxInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sizeMaxInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudImagesImageV2) SizeMin() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sizeMin",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudImagesImageV2) SizeMinInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sizeMinInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudImagesImageV2) SortDirection() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sortDirection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudImagesImageV2) SortDirectionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sortDirectionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudImagesImageV2) SortKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sortKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudImagesImageV2) SortKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sortKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudImagesImageV2) Tag() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tag",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudImagesImageV2) TagInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tagInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudImagesImageV2) Tags() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudImagesImageV2) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudImagesImageV2) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudImagesImageV2) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudImagesImageV2) UpdatedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updatedAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudImagesImageV2) Visibility() *string {
	var returns *string
	_jsii_.Get(
		j,
		"visibility",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudImagesImageV2) VisibilityInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"visibilityInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://www.terraform.io/docs/providers/opentelekomcloud/d/images_image_v2 opentelekomcloud_images_image_v2} Data Source.
func NewDataOpentelekomcloudImagesImageV2(scope constructs.Construct, id *string, config *DataOpentelekomcloudImagesImageV2Config) DataOpentelekomcloudImagesImageV2 {
	_init_.Initialize()

	if err := validateNewDataOpentelekomcloudImagesImageV2Parameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataOpentelekomcloudImagesImageV2{}

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.dataOpentelekomcloudImagesImageV2.DataOpentelekomcloudImagesImageV2",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/opentelekomcloud/d/images_image_v2 opentelekomcloud_images_image_v2} Data Source.
func NewDataOpentelekomcloudImagesImageV2_Override(d DataOpentelekomcloudImagesImageV2, scope constructs.Construct, id *string, config *DataOpentelekomcloudImagesImageV2Config) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.dataOpentelekomcloudImagesImageV2.DataOpentelekomcloudImagesImageV2",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DataOpentelekomcloudImagesImageV2)SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DataOpentelekomcloudImagesImageV2)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DataOpentelekomcloudImagesImageV2)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DataOpentelekomcloudImagesImageV2)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_DataOpentelekomcloudImagesImageV2)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DataOpentelekomcloudImagesImageV2)SetMostRecent(val interface{}) {
	if err := j.validateSetMostRecentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mostRecent",
		val,
	)
}

func (j *jsiiProxy_DataOpentelekomcloudImagesImageV2)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_DataOpentelekomcloudImagesImageV2)SetNameRegex(val *string) {
	if err := j.validateSetNameRegexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nameRegex",
		val,
	)
}

func (j *jsiiProxy_DataOpentelekomcloudImagesImageV2)SetOwner(val *string) {
	if err := j.validateSetOwnerParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"owner",
		val,
	)
}

func (j *jsiiProxy_DataOpentelekomcloudImagesImageV2)SetProperties(val *map[string]*string) {
	if err := j.validateSetPropertiesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"properties",
		val,
	)
}

func (j *jsiiProxy_DataOpentelekomcloudImagesImageV2)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DataOpentelekomcloudImagesImageV2)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_DataOpentelekomcloudImagesImageV2)SetSizeMax(val *float64) {
	if err := j.validateSetSizeMaxParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sizeMax",
		val,
	)
}

func (j *jsiiProxy_DataOpentelekomcloudImagesImageV2)SetSizeMin(val *float64) {
	if err := j.validateSetSizeMinParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sizeMin",
		val,
	)
}

func (j *jsiiProxy_DataOpentelekomcloudImagesImageV2)SetSortDirection(val *string) {
	if err := j.validateSetSortDirectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sortDirection",
		val,
	)
}

func (j *jsiiProxy_DataOpentelekomcloudImagesImageV2)SetSortKey(val *string) {
	if err := j.validateSetSortKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sortKey",
		val,
	)
}

func (j *jsiiProxy_DataOpentelekomcloudImagesImageV2)SetTag(val *string) {
	if err := j.validateSetTagParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tag",
		val,
	)
}

func (j *jsiiProxy_DataOpentelekomcloudImagesImageV2)SetVisibility(val *string) {
	if err := j.validateSetVisibilityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"visibility",
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
func DataOpentelekomcloudImagesImageV2_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataOpentelekomcloudImagesImageV2_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-opentelekomcloud.dataOpentelekomcloudImagesImageV2.DataOpentelekomcloudImagesImageV2",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataOpentelekomcloudImagesImageV2_IsTerraformDataSource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataOpentelekomcloudImagesImageV2_IsTerraformDataSourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-opentelekomcloud.dataOpentelekomcloudImagesImageV2.DataOpentelekomcloudImagesImageV2",
		"isTerraformDataSource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataOpentelekomcloudImagesImageV2_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataOpentelekomcloudImagesImageV2_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-opentelekomcloud.dataOpentelekomcloudImagesImageV2.DataOpentelekomcloudImagesImageV2",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataOpentelekomcloudImagesImageV2_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-opentelekomcloud.dataOpentelekomcloudImagesImageV2.DataOpentelekomcloudImagesImageV2",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DataOpentelekomcloudImagesImageV2) AddOverride(path *string, value interface{}) {
	if err := d.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DataOpentelekomcloudImagesImageV2) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataOpentelekomcloudImagesImageV2) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataOpentelekomcloudImagesImageV2) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataOpentelekomcloudImagesImageV2) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataOpentelekomcloudImagesImageV2) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataOpentelekomcloudImagesImageV2) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataOpentelekomcloudImagesImageV2) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataOpentelekomcloudImagesImageV2) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataOpentelekomcloudImagesImageV2) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataOpentelekomcloudImagesImageV2) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataOpentelekomcloudImagesImageV2) OverrideLogicalId(newLogicalId *string) {
	if err := d.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DataOpentelekomcloudImagesImageV2) ResetId() {
	_jsii_.InvokeVoid(
		d,
		"resetId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataOpentelekomcloudImagesImageV2) ResetMostRecent() {
	_jsii_.InvokeVoid(
		d,
		"resetMostRecent",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataOpentelekomcloudImagesImageV2) ResetName() {
	_jsii_.InvokeVoid(
		d,
		"resetName",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataOpentelekomcloudImagesImageV2) ResetNameRegex() {
	_jsii_.InvokeVoid(
		d,
		"resetNameRegex",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataOpentelekomcloudImagesImageV2) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataOpentelekomcloudImagesImageV2) ResetOwner() {
	_jsii_.InvokeVoid(
		d,
		"resetOwner",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataOpentelekomcloudImagesImageV2) ResetProperties() {
	_jsii_.InvokeVoid(
		d,
		"resetProperties",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataOpentelekomcloudImagesImageV2) ResetRegion() {
	_jsii_.InvokeVoid(
		d,
		"resetRegion",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataOpentelekomcloudImagesImageV2) ResetSizeMax() {
	_jsii_.InvokeVoid(
		d,
		"resetSizeMax",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataOpentelekomcloudImagesImageV2) ResetSizeMin() {
	_jsii_.InvokeVoid(
		d,
		"resetSizeMin",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataOpentelekomcloudImagesImageV2) ResetSortDirection() {
	_jsii_.InvokeVoid(
		d,
		"resetSortDirection",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataOpentelekomcloudImagesImageV2) ResetSortKey() {
	_jsii_.InvokeVoid(
		d,
		"resetSortKey",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataOpentelekomcloudImagesImageV2) ResetTag() {
	_jsii_.InvokeVoid(
		d,
		"resetTag",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataOpentelekomcloudImagesImageV2) ResetVisibility() {
	_jsii_.InvokeVoid(
		d,
		"resetVisibility",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataOpentelekomcloudImagesImageV2) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataOpentelekomcloudImagesImageV2) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataOpentelekomcloudImagesImageV2) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataOpentelekomcloudImagesImageV2) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

