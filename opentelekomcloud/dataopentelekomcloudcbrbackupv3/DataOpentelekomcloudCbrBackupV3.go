// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataopentelekomcloudcbrbackupv3

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v11/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v11/dataopentelekomcloudcbrbackupv3/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.12/docs/data-sources/cbr_backup_v3 opentelekomcloud_cbr_backup_v3}.
type DataOpentelekomcloudCbrBackupV3 interface {
	cdktf.TerraformDataSource
	AutoTrigger() interface{}
	SetAutoTrigger(val interface{})
	AutoTriggerInput() interface{}
	Bootable() interface{}
	SetBootable(val interface{})
	BootableInput() interface{}
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	CheckpointId() *string
	SetCheckpointId(val *string)
	CheckpointIdInput() *string
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	ContainSystemDisk() interface{}
	SetContainSystemDisk(val interface{})
	ContainSystemDiskInput() interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	CreatedAt() *string
	SetCreatedAt(val *string)
	CreatedAtInput() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	Encrypted() interface{}
	SetEncrypted(val interface{})
	EncryptedInput() interface{}
	ExpiredAt() *string
	SetExpiredAt(val *string)
	ExpiredAtInput() *string
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
	ImageType() *string
	SetImageType(val *string)
	ImageTypeInput() *string
	Incremental() interface{}
	SetIncremental(val interface{})
	IncrementalInput() interface{}
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	ParentId() *string
	SetParentId(val *string)
	ParentIdInput() *string
	ProjectId() *string
	SetProjectId(val *string)
	ProjectIdInput() *string
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	ProviderId() *string
	SetProviderId(val *string)
	ProviderIdInput() *string
	// Experimental.
	RawOverrides() interface{}
	ResourceAz() *string
	SetResourceAz(val *string)
	ResourceAzInput() *string
	ResourceId() *string
	SetResourceId(val *string)
	ResourceIdInput() *string
	ResourceName() *string
	SetResourceName(val *string)
	ResourceNameInput() *string
	ResourceSize() *float64
	SetResourceSize(val *float64)
	ResourceSizeInput() *float64
	ResourceType() *string
	SetResourceType(val *string)
	ResourceTypeInput() *string
	SnapshotId() *string
	SetSnapshotId(val *string)
	SnapshotIdInput() *string
	Status() *string
	SetStatus(val *string)
	StatusInput() *string
	SupportedRestoreMode() *string
	SetSupportedRestoreMode(val *string)
	SupportedRestoreModeInput() *string
	SupportLld() interface{}
	SetSupportLld(val interface{})
	SupportLldInput() interface{}
	SystemDisk() interface{}
	SetSystemDisk(val interface{})
	SystemDiskInput() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	UpdatedAt() *string
	SetUpdatedAt(val *string)
	UpdatedAtInput() *string
	VaultId() *string
	SetVaultId(val *string)
	VaultIdInput() *string
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
	ResetAutoTrigger()
	ResetBootable()
	ResetCheckpointId()
	ResetContainSystemDisk()
	ResetCreatedAt()
	ResetDescription()
	ResetEncrypted()
	ResetExpiredAt()
	ResetId()
	ResetImageType()
	ResetIncremental()
	ResetName()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetParentId()
	ResetProjectId()
	ResetProviderId()
	ResetResourceAz()
	ResetResourceId()
	ResetResourceName()
	ResetResourceSize()
	ResetResourceType()
	ResetSnapshotId()
	ResetStatus()
	ResetSupportedRestoreMode()
	ResetSupportLld()
	ResetSystemDisk()
	ResetUpdatedAt()
	ResetVaultId()
	SynthesizeAttributes() *map[string]interface{}
	SynthesizeHclAttributes() *map[string]interface{}
	// Adds this resource to the terraform JSON output.
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

// The jsii proxy struct for DataOpentelekomcloudCbrBackupV3
type jsiiProxy_DataOpentelekomcloudCbrBackupV3 struct {
	internal.Type__cdktfTerraformDataSource
}

func (j *jsiiProxy_DataOpentelekomcloudCbrBackupV3) AutoTrigger() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoTrigger",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudCbrBackupV3) AutoTriggerInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoTriggerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudCbrBackupV3) Bootable() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"bootable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudCbrBackupV3) BootableInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"bootableInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudCbrBackupV3) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudCbrBackupV3) CheckpointId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"checkpointId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudCbrBackupV3) CheckpointIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"checkpointIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudCbrBackupV3) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudCbrBackupV3) ContainSystemDisk() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"containSystemDisk",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudCbrBackupV3) ContainSystemDiskInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"containSystemDiskInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudCbrBackupV3) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudCbrBackupV3) CreatedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudCbrBackupV3) CreatedAtInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdAtInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudCbrBackupV3) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudCbrBackupV3) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudCbrBackupV3) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudCbrBackupV3) Encrypted() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"encrypted",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudCbrBackupV3) EncryptedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"encryptedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudCbrBackupV3) ExpiredAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"expiredAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudCbrBackupV3) ExpiredAtInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"expiredAtInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudCbrBackupV3) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudCbrBackupV3) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudCbrBackupV3) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudCbrBackupV3) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudCbrBackupV3) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudCbrBackupV3) ImageType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudCbrBackupV3) ImageTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudCbrBackupV3) Incremental() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"incremental",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudCbrBackupV3) IncrementalInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"incrementalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudCbrBackupV3) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudCbrBackupV3) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudCbrBackupV3) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudCbrBackupV3) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudCbrBackupV3) ParentId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"parentId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudCbrBackupV3) ParentIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"parentIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudCbrBackupV3) ProjectId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudCbrBackupV3) ProjectIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudCbrBackupV3) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudCbrBackupV3) ProviderId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"providerId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudCbrBackupV3) ProviderIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"providerIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudCbrBackupV3) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudCbrBackupV3) ResourceAz() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceAz",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudCbrBackupV3) ResourceAzInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceAzInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudCbrBackupV3) ResourceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudCbrBackupV3) ResourceIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudCbrBackupV3) ResourceName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudCbrBackupV3) ResourceNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudCbrBackupV3) ResourceSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"resourceSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudCbrBackupV3) ResourceSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"resourceSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudCbrBackupV3) ResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudCbrBackupV3) ResourceTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudCbrBackupV3) SnapshotId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"snapshotId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudCbrBackupV3) SnapshotIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"snapshotIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudCbrBackupV3) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudCbrBackupV3) StatusInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"statusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudCbrBackupV3) SupportedRestoreMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"supportedRestoreMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudCbrBackupV3) SupportedRestoreModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"supportedRestoreModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudCbrBackupV3) SupportLld() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"supportLld",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudCbrBackupV3) SupportLldInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"supportLldInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudCbrBackupV3) SystemDisk() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"systemDisk",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudCbrBackupV3) SystemDiskInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"systemDiskInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudCbrBackupV3) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudCbrBackupV3) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudCbrBackupV3) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudCbrBackupV3) UpdatedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updatedAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudCbrBackupV3) UpdatedAtInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updatedAtInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudCbrBackupV3) VaultId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vaultId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudCbrBackupV3) VaultIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vaultIdInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.12/docs/data-sources/cbr_backup_v3 opentelekomcloud_cbr_backup_v3} Data Source.
func NewDataOpentelekomcloudCbrBackupV3(scope constructs.Construct, id *string, config *DataOpentelekomcloudCbrBackupV3Config) DataOpentelekomcloudCbrBackupV3 {
	_init_.Initialize()

	if err := validateNewDataOpentelekomcloudCbrBackupV3Parameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataOpentelekomcloudCbrBackupV3{}

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.dataOpentelekomcloudCbrBackupV3.DataOpentelekomcloudCbrBackupV3",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.12/docs/data-sources/cbr_backup_v3 opentelekomcloud_cbr_backup_v3} Data Source.
func NewDataOpentelekomcloudCbrBackupV3_Override(d DataOpentelekomcloudCbrBackupV3, scope constructs.Construct, id *string, config *DataOpentelekomcloudCbrBackupV3Config) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.dataOpentelekomcloudCbrBackupV3.DataOpentelekomcloudCbrBackupV3",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DataOpentelekomcloudCbrBackupV3)SetAutoTrigger(val interface{}) {
	if err := j.validateSetAutoTriggerParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoTrigger",
		val,
	)
}

func (j *jsiiProxy_DataOpentelekomcloudCbrBackupV3)SetBootable(val interface{}) {
	if err := j.validateSetBootableParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bootable",
		val,
	)
}

func (j *jsiiProxy_DataOpentelekomcloudCbrBackupV3)SetCheckpointId(val *string) {
	if err := j.validateSetCheckpointIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"checkpointId",
		val,
	)
}

func (j *jsiiProxy_DataOpentelekomcloudCbrBackupV3)SetContainSystemDisk(val interface{}) {
	if err := j.validateSetContainSystemDiskParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"containSystemDisk",
		val,
	)
}

func (j *jsiiProxy_DataOpentelekomcloudCbrBackupV3)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DataOpentelekomcloudCbrBackupV3)SetCreatedAt(val *string) {
	if err := j.validateSetCreatedAtParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"createdAt",
		val,
	)
}

func (j *jsiiProxy_DataOpentelekomcloudCbrBackupV3)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DataOpentelekomcloudCbrBackupV3)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_DataOpentelekomcloudCbrBackupV3)SetEncrypted(val interface{}) {
	if err := j.validateSetEncryptedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"encrypted",
		val,
	)
}

func (j *jsiiProxy_DataOpentelekomcloudCbrBackupV3)SetExpiredAt(val *string) {
	if err := j.validateSetExpiredAtParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"expiredAt",
		val,
	)
}

func (j *jsiiProxy_DataOpentelekomcloudCbrBackupV3)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DataOpentelekomcloudCbrBackupV3)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_DataOpentelekomcloudCbrBackupV3)SetImageType(val *string) {
	if err := j.validateSetImageTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"imageType",
		val,
	)
}

func (j *jsiiProxy_DataOpentelekomcloudCbrBackupV3)SetIncremental(val interface{}) {
	if err := j.validateSetIncrementalParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"incremental",
		val,
	)
}

func (j *jsiiProxy_DataOpentelekomcloudCbrBackupV3)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DataOpentelekomcloudCbrBackupV3)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_DataOpentelekomcloudCbrBackupV3)SetParentId(val *string) {
	if err := j.validateSetParentIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"parentId",
		val,
	)
}

func (j *jsiiProxy_DataOpentelekomcloudCbrBackupV3)SetProjectId(val *string) {
	if err := j.validateSetProjectIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"projectId",
		val,
	)
}

func (j *jsiiProxy_DataOpentelekomcloudCbrBackupV3)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DataOpentelekomcloudCbrBackupV3)SetProviderId(val *string) {
	if err := j.validateSetProviderIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"providerId",
		val,
	)
}

func (j *jsiiProxy_DataOpentelekomcloudCbrBackupV3)SetResourceAz(val *string) {
	if err := j.validateSetResourceAzParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceAz",
		val,
	)
}

func (j *jsiiProxy_DataOpentelekomcloudCbrBackupV3)SetResourceId(val *string) {
	if err := j.validateSetResourceIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceId",
		val,
	)
}

func (j *jsiiProxy_DataOpentelekomcloudCbrBackupV3)SetResourceName(val *string) {
	if err := j.validateSetResourceNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceName",
		val,
	)
}

func (j *jsiiProxy_DataOpentelekomcloudCbrBackupV3)SetResourceSize(val *float64) {
	if err := j.validateSetResourceSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceSize",
		val,
	)
}

func (j *jsiiProxy_DataOpentelekomcloudCbrBackupV3)SetResourceType(val *string) {
	if err := j.validateSetResourceTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceType",
		val,
	)
}

func (j *jsiiProxy_DataOpentelekomcloudCbrBackupV3)SetSnapshotId(val *string) {
	if err := j.validateSetSnapshotIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"snapshotId",
		val,
	)
}

func (j *jsiiProxy_DataOpentelekomcloudCbrBackupV3)SetStatus(val *string) {
	if err := j.validateSetStatusParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"status",
		val,
	)
}

func (j *jsiiProxy_DataOpentelekomcloudCbrBackupV3)SetSupportedRestoreMode(val *string) {
	if err := j.validateSetSupportedRestoreModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"supportedRestoreMode",
		val,
	)
}

func (j *jsiiProxy_DataOpentelekomcloudCbrBackupV3)SetSupportLld(val interface{}) {
	if err := j.validateSetSupportLldParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"supportLld",
		val,
	)
}

func (j *jsiiProxy_DataOpentelekomcloudCbrBackupV3)SetSystemDisk(val interface{}) {
	if err := j.validateSetSystemDiskParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"systemDisk",
		val,
	)
}

func (j *jsiiProxy_DataOpentelekomcloudCbrBackupV3)SetUpdatedAt(val *string) {
	if err := j.validateSetUpdatedAtParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"updatedAt",
		val,
	)
}

func (j *jsiiProxy_DataOpentelekomcloudCbrBackupV3)SetVaultId(val *string) {
	if err := j.validateSetVaultIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vaultId",
		val,
	)
}

// Generates CDKTF code for importing a DataOpentelekomcloudCbrBackupV3 resource upon running "cdktf plan <stack-name>".
func DataOpentelekomcloudCbrBackupV3_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateDataOpentelekomcloudCbrBackupV3_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-opentelekomcloud.dataOpentelekomcloudCbrBackupV3.DataOpentelekomcloudCbrBackupV3",
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
func DataOpentelekomcloudCbrBackupV3_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataOpentelekomcloudCbrBackupV3_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-opentelekomcloud.dataOpentelekomcloudCbrBackupV3.DataOpentelekomcloudCbrBackupV3",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataOpentelekomcloudCbrBackupV3_IsTerraformDataSource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataOpentelekomcloudCbrBackupV3_IsTerraformDataSourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-opentelekomcloud.dataOpentelekomcloudCbrBackupV3.DataOpentelekomcloudCbrBackupV3",
		"isTerraformDataSource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataOpentelekomcloudCbrBackupV3_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataOpentelekomcloudCbrBackupV3_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-opentelekomcloud.dataOpentelekomcloudCbrBackupV3.DataOpentelekomcloudCbrBackupV3",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataOpentelekomcloudCbrBackupV3_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-opentelekomcloud.dataOpentelekomcloudCbrBackupV3.DataOpentelekomcloudCbrBackupV3",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DataOpentelekomcloudCbrBackupV3) AddOverride(path *string, value interface{}) {
	if err := d.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DataOpentelekomcloudCbrBackupV3) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataOpentelekomcloudCbrBackupV3) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataOpentelekomcloudCbrBackupV3) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataOpentelekomcloudCbrBackupV3) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataOpentelekomcloudCbrBackupV3) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataOpentelekomcloudCbrBackupV3) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataOpentelekomcloudCbrBackupV3) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataOpentelekomcloudCbrBackupV3) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataOpentelekomcloudCbrBackupV3) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataOpentelekomcloudCbrBackupV3) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataOpentelekomcloudCbrBackupV3) OverrideLogicalId(newLogicalId *string) {
	if err := d.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DataOpentelekomcloudCbrBackupV3) ResetAutoTrigger() {
	_jsii_.InvokeVoid(
		d,
		"resetAutoTrigger",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataOpentelekomcloudCbrBackupV3) ResetBootable() {
	_jsii_.InvokeVoid(
		d,
		"resetBootable",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataOpentelekomcloudCbrBackupV3) ResetCheckpointId() {
	_jsii_.InvokeVoid(
		d,
		"resetCheckpointId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataOpentelekomcloudCbrBackupV3) ResetContainSystemDisk() {
	_jsii_.InvokeVoid(
		d,
		"resetContainSystemDisk",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataOpentelekomcloudCbrBackupV3) ResetCreatedAt() {
	_jsii_.InvokeVoid(
		d,
		"resetCreatedAt",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataOpentelekomcloudCbrBackupV3) ResetDescription() {
	_jsii_.InvokeVoid(
		d,
		"resetDescription",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataOpentelekomcloudCbrBackupV3) ResetEncrypted() {
	_jsii_.InvokeVoid(
		d,
		"resetEncrypted",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataOpentelekomcloudCbrBackupV3) ResetExpiredAt() {
	_jsii_.InvokeVoid(
		d,
		"resetExpiredAt",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataOpentelekomcloudCbrBackupV3) ResetId() {
	_jsii_.InvokeVoid(
		d,
		"resetId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataOpentelekomcloudCbrBackupV3) ResetImageType() {
	_jsii_.InvokeVoid(
		d,
		"resetImageType",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataOpentelekomcloudCbrBackupV3) ResetIncremental() {
	_jsii_.InvokeVoid(
		d,
		"resetIncremental",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataOpentelekomcloudCbrBackupV3) ResetName() {
	_jsii_.InvokeVoid(
		d,
		"resetName",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataOpentelekomcloudCbrBackupV3) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataOpentelekomcloudCbrBackupV3) ResetParentId() {
	_jsii_.InvokeVoid(
		d,
		"resetParentId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataOpentelekomcloudCbrBackupV3) ResetProjectId() {
	_jsii_.InvokeVoid(
		d,
		"resetProjectId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataOpentelekomcloudCbrBackupV3) ResetProviderId() {
	_jsii_.InvokeVoid(
		d,
		"resetProviderId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataOpentelekomcloudCbrBackupV3) ResetResourceAz() {
	_jsii_.InvokeVoid(
		d,
		"resetResourceAz",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataOpentelekomcloudCbrBackupV3) ResetResourceId() {
	_jsii_.InvokeVoid(
		d,
		"resetResourceId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataOpentelekomcloudCbrBackupV3) ResetResourceName() {
	_jsii_.InvokeVoid(
		d,
		"resetResourceName",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataOpentelekomcloudCbrBackupV3) ResetResourceSize() {
	_jsii_.InvokeVoid(
		d,
		"resetResourceSize",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataOpentelekomcloudCbrBackupV3) ResetResourceType() {
	_jsii_.InvokeVoid(
		d,
		"resetResourceType",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataOpentelekomcloudCbrBackupV3) ResetSnapshotId() {
	_jsii_.InvokeVoid(
		d,
		"resetSnapshotId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataOpentelekomcloudCbrBackupV3) ResetStatus() {
	_jsii_.InvokeVoid(
		d,
		"resetStatus",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataOpentelekomcloudCbrBackupV3) ResetSupportedRestoreMode() {
	_jsii_.InvokeVoid(
		d,
		"resetSupportedRestoreMode",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataOpentelekomcloudCbrBackupV3) ResetSupportLld() {
	_jsii_.InvokeVoid(
		d,
		"resetSupportLld",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataOpentelekomcloudCbrBackupV3) ResetSystemDisk() {
	_jsii_.InvokeVoid(
		d,
		"resetSystemDisk",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataOpentelekomcloudCbrBackupV3) ResetUpdatedAt() {
	_jsii_.InvokeVoid(
		d,
		"resetUpdatedAt",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataOpentelekomcloudCbrBackupV3) ResetVaultId() {
	_jsii_.InvokeVoid(
		d,
		"resetVaultId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataOpentelekomcloudCbrBackupV3) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataOpentelekomcloudCbrBackupV3) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataOpentelekomcloudCbrBackupV3) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataOpentelekomcloudCbrBackupV3) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataOpentelekomcloudCbrBackupV3) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataOpentelekomcloudCbrBackupV3) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

