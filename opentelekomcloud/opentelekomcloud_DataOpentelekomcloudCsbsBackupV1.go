// Prebuilt opentelekomcloud Provider for Terraform CDK (cdktf)
package opentelekomcloud

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v2/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://www.terraform.io/docs/providers/opentelekomcloud/d/csbs_backup_v1 opentelekomcloud_csbs_backup_v1}.
type DataOpentelekomcloudCsbsBackupV1 interface {
	cdktf.TerraformDataSource
	AutoTrigger() cdktf.IResolvable
	AverageSpeed() *float64
	BackupName() *string
	SetBackupName(val *string)
	BackupNameInput() *string
	BackupRecordId() *string
	SetBackupRecordId(val *string)
	BackupRecordIdInput() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
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
	// The tree node.
	Node() constructs.Node
	PolicyId() *string
	SetPolicyId(val *string)
	PolicyIdInput() *string
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	RawOverrides() interface{}
	Region() *string
	SetRegion(val *string)
	RegionInput() *string
	ResourceId() *string
	SetResourceId(val *string)
	ResourceIdInput() *string
	ResourceName() *string
	SetResourceName(val *string)
	ResourceNameInput() *string
	ResourceType() *string
	SetResourceType(val *string)
	ResourceTypeInput() *string
	Size() *float64
	Status() *string
	SetStatus(val *string)
	StatusInput() *string
	Tags() DataOpentelekomcloudCsbsBackupV1TagsList
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	VmIp() *string
	SetVmIp(val *string)
	VmIpInput() *string
	VmMetadata() DataOpentelekomcloudCsbsBackupV1VmMetadataList
	VolumeBackups() DataOpentelekomcloudCsbsBackupV1VolumeBackupsList
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
	ResetBackupName()
	ResetBackupRecordId()
	ResetId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPolicyId()
	ResetRegion()
	ResetResourceId()
	ResetResourceName()
	ResetResourceType()
	ResetStatus()
	ResetVmIp()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for DataOpentelekomcloudCsbsBackupV1
type jsiiProxy_DataOpentelekomcloudCsbsBackupV1 struct {
	internal.Type__cdktfTerraformDataSource
}

func (j *jsiiProxy_DataOpentelekomcloudCsbsBackupV1) AutoTrigger() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"autoTrigger",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudCsbsBackupV1) AverageSpeed() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"averageSpeed",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudCsbsBackupV1) BackupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"backupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudCsbsBackupV1) BackupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"backupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudCsbsBackupV1) BackupRecordId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"backupRecordId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudCsbsBackupV1) BackupRecordIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"backupRecordIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudCsbsBackupV1) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudCsbsBackupV1) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudCsbsBackupV1) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudCsbsBackupV1) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudCsbsBackupV1) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudCsbsBackupV1) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudCsbsBackupV1) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudCsbsBackupV1) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudCsbsBackupV1) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudCsbsBackupV1) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudCsbsBackupV1) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudCsbsBackupV1) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudCsbsBackupV1) PolicyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudCsbsBackupV1) PolicyIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policyIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudCsbsBackupV1) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudCsbsBackupV1) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudCsbsBackupV1) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudCsbsBackupV1) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudCsbsBackupV1) ResourceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudCsbsBackupV1) ResourceIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudCsbsBackupV1) ResourceName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudCsbsBackupV1) ResourceNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudCsbsBackupV1) ResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudCsbsBackupV1) ResourceTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudCsbsBackupV1) Size() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"size",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudCsbsBackupV1) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudCsbsBackupV1) StatusInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"statusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudCsbsBackupV1) Tags() DataOpentelekomcloudCsbsBackupV1TagsList {
	var returns DataOpentelekomcloudCsbsBackupV1TagsList
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudCsbsBackupV1) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudCsbsBackupV1) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudCsbsBackupV1) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudCsbsBackupV1) VmIp() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vmIp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudCsbsBackupV1) VmIpInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vmIpInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudCsbsBackupV1) VmMetadata() DataOpentelekomcloudCsbsBackupV1VmMetadataList {
	var returns DataOpentelekomcloudCsbsBackupV1VmMetadataList
	_jsii_.Get(
		j,
		"vmMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudCsbsBackupV1) VolumeBackups() DataOpentelekomcloudCsbsBackupV1VolumeBackupsList {
	var returns DataOpentelekomcloudCsbsBackupV1VolumeBackupsList
	_jsii_.Get(
		j,
		"volumeBackups",
		&returns,
	)
	return returns
}


// Create a new {@link https://www.terraform.io/docs/providers/opentelekomcloud/d/csbs_backup_v1 opentelekomcloud_csbs_backup_v1} Data Source.
func NewDataOpentelekomcloudCsbsBackupV1(scope constructs.Construct, id *string, config *DataOpentelekomcloudCsbsBackupV1Config) DataOpentelekomcloudCsbsBackupV1 {
	_init_.Initialize()

	j := jsiiProxy_DataOpentelekomcloudCsbsBackupV1{}

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.DataOpentelekomcloudCsbsBackupV1",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/opentelekomcloud/d/csbs_backup_v1 opentelekomcloud_csbs_backup_v1} Data Source.
func NewDataOpentelekomcloudCsbsBackupV1_Override(d DataOpentelekomcloudCsbsBackupV1, scope constructs.Construct, id *string, config *DataOpentelekomcloudCsbsBackupV1Config) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.DataOpentelekomcloudCsbsBackupV1",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DataOpentelekomcloudCsbsBackupV1) SetBackupName(val *string) {
	_jsii_.Set(
		j,
		"backupName",
		val,
	)
}

func (j *jsiiProxy_DataOpentelekomcloudCsbsBackupV1) SetBackupRecordId(val *string) {
	_jsii_.Set(
		j,
		"backupRecordId",
		val,
	)
}

func (j *jsiiProxy_DataOpentelekomcloudCsbsBackupV1) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DataOpentelekomcloudCsbsBackupV1) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DataOpentelekomcloudCsbsBackupV1) SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DataOpentelekomcloudCsbsBackupV1) SetId(val *string) {
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_DataOpentelekomcloudCsbsBackupV1) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DataOpentelekomcloudCsbsBackupV1) SetPolicyId(val *string) {
	_jsii_.Set(
		j,
		"policyId",
		val,
	)
}

func (j *jsiiProxy_DataOpentelekomcloudCsbsBackupV1) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DataOpentelekomcloudCsbsBackupV1) SetRegion(val *string) {
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_DataOpentelekomcloudCsbsBackupV1) SetResourceId(val *string) {
	_jsii_.Set(
		j,
		"resourceId",
		val,
	)
}

func (j *jsiiProxy_DataOpentelekomcloudCsbsBackupV1) SetResourceName(val *string) {
	_jsii_.Set(
		j,
		"resourceName",
		val,
	)
}

func (j *jsiiProxy_DataOpentelekomcloudCsbsBackupV1) SetResourceType(val *string) {
	_jsii_.Set(
		j,
		"resourceType",
		val,
	)
}

func (j *jsiiProxy_DataOpentelekomcloudCsbsBackupV1) SetStatus(val *string) {
	_jsii_.Set(
		j,
		"status",
		val,
	)
}

func (j *jsiiProxy_DataOpentelekomcloudCsbsBackupV1) SetVmIp(val *string) {
	_jsii_.Set(
		j,
		"vmIp",
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
func DataOpentelekomcloudCsbsBackupV1_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-opentelekomcloud.DataOpentelekomcloudCsbsBackupV1",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataOpentelekomcloudCsbsBackupV1_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-opentelekomcloud.DataOpentelekomcloudCsbsBackupV1",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DataOpentelekomcloudCsbsBackupV1) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DataOpentelekomcloudCsbsBackupV1) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataOpentelekomcloudCsbsBackupV1) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataOpentelekomcloudCsbsBackupV1) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataOpentelekomcloudCsbsBackupV1) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataOpentelekomcloudCsbsBackupV1) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataOpentelekomcloudCsbsBackupV1) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataOpentelekomcloudCsbsBackupV1) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataOpentelekomcloudCsbsBackupV1) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataOpentelekomcloudCsbsBackupV1) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataOpentelekomcloudCsbsBackupV1) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataOpentelekomcloudCsbsBackupV1) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DataOpentelekomcloudCsbsBackupV1) ResetBackupName() {
	_jsii_.InvokeVoid(
		d,
		"resetBackupName",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataOpentelekomcloudCsbsBackupV1) ResetBackupRecordId() {
	_jsii_.InvokeVoid(
		d,
		"resetBackupRecordId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataOpentelekomcloudCsbsBackupV1) ResetId() {
	_jsii_.InvokeVoid(
		d,
		"resetId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataOpentelekomcloudCsbsBackupV1) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataOpentelekomcloudCsbsBackupV1) ResetPolicyId() {
	_jsii_.InvokeVoid(
		d,
		"resetPolicyId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataOpentelekomcloudCsbsBackupV1) ResetRegion() {
	_jsii_.InvokeVoid(
		d,
		"resetRegion",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataOpentelekomcloudCsbsBackupV1) ResetResourceId() {
	_jsii_.InvokeVoid(
		d,
		"resetResourceId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataOpentelekomcloudCsbsBackupV1) ResetResourceName() {
	_jsii_.InvokeVoid(
		d,
		"resetResourceName",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataOpentelekomcloudCsbsBackupV1) ResetResourceType() {
	_jsii_.InvokeVoid(
		d,
		"resetResourceType",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataOpentelekomcloudCsbsBackupV1) ResetStatus() {
	_jsii_.InvokeVoid(
		d,
		"resetStatus",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataOpentelekomcloudCsbsBackupV1) ResetVmIp() {
	_jsii_.InvokeVoid(
		d,
		"resetVmIp",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataOpentelekomcloudCsbsBackupV1) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataOpentelekomcloudCsbsBackupV1) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataOpentelekomcloudCsbsBackupV1) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataOpentelekomcloudCsbsBackupV1) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}
