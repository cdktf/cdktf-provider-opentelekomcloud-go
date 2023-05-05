package mrsclusterv1

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v6/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v6/mrsclusterv1/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.34.3/docs/resources/mrs_cluster_v1 opentelekomcloud_mrs_cluster_v1}.
type MrsClusterV1 interface {
	cdktf.TerraformResource
	AddJobs() MrsClusterV1AddJobsList
	AddJobsInput() interface{}
	AvailableZoneId() *string
	SetAvailableZoneId(val *string)
	AvailableZoneIdInput() *string
	AvailableZoneName() *string
	BillingType() *float64
	SetBillingType(val *float64)
	BillingTypeInput() *float64
	BootstrapScripts() MrsClusterV1BootstrapScriptsList
	BootstrapScriptsInput() interface{}
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	ChargingStartTime() *string
	ClusterAdminSecret() *string
	SetClusterAdminSecret(val *string)
	ClusterAdminSecretInput() *string
	ClusterId() *string
	ClusterName() *string
	SetClusterName(val *string)
	ClusterNameInput() *string
	ClusterState() *string
	ClusterType() *float64
	SetClusterType(val *float64)
	ClusterTypeInput() *float64
	ClusterVersion() *string
	SetClusterVersion(val *string)
	ClusterVersionInput() *string
	ComponentList() MrsClusterV1ComponentListList
	ComponentListInput() interface{}
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	CoreDataVolumeCount() *float64
	SetCoreDataVolumeCount(val *float64)
	CoreDataVolumeCountInput() *float64
	CoreDataVolumeSize() *float64
	SetCoreDataVolumeSize(val *float64)
	CoreDataVolumeSizeInput() *float64
	CoreDataVolumeType() *string
	SetCoreDataVolumeType(val *string)
	CoreDataVolumeTypeInput() *string
	CoreNodeNum() *float64
	SetCoreNodeNum(val *float64)
	CoreNodeNumInput() *float64
	CoreNodeProductId() *string
	CoreNodeSize() *string
	SetCoreNodeSize(val *string)
	CoreNodeSizeInput() *string
	CoreNodeSpecId() *string
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	CreateAt() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	DeploymentId() *string
	ErrorInfo() *string
	ExternalAlternateIp() *string
	ExternalIp() *string
	Fee() *string
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	HadoopVersion() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	InstanceId() *string
	InternalIp() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	LogCollection() *float64
	SetLogCollection(val *float64)
	LogCollectionInput() *float64
	MasterDataVolumeCount() *float64
	SetMasterDataVolumeCount(val *float64)
	MasterDataVolumeCountInput() *float64
	MasterDataVolumeSize() *float64
	SetMasterDataVolumeSize(val *float64)
	MasterDataVolumeSizeInput() *float64
	MasterDataVolumeType() *string
	SetMasterDataVolumeType(val *string)
	MasterDataVolumeTypeInput() *string
	MasterNodeIp() *string
	MasterNodeNum() *float64
	SetMasterNodeNum(val *float64)
	MasterNodeNumInput() *float64
	MasterNodeProductId() *string
	MasterNodeSize() *string
	SetMasterNodeSize(val *string)
	MasterNodeSizeInput() *string
	MasterNodeSpecId() *string
	// The tree node.
	Node() constructs.Node
	NodePublicCertName() *string
	SetNodePublicCertName(val *string)
	NodePublicCertNameInput() *string
	OrderId() *string
	PrivateIpFirst() *string
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
	Remark() *string
	SafeMode() *float64
	SetSafeMode(val *float64)
	SafeModeInput() *float64
	SecurityGroupsId() *string
	SlaveSecurityGroupsId() *string
	SubnetId() *string
	SetSubnetId(val *string)
	SubnetIdInput() *string
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsInput() *map[string]*string
	TenantId() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() MrsClusterV1TimeoutsOutputReference
	TimeoutsInput() interface{}
	UpdateAt() *string
	Vnc() *string
	VolumeSize() *float64
	SetVolumeSize(val *float64)
	VolumeSizeInput() *float64
	VolumeType() *string
	SetVolumeType(val *string)
	VolumeTypeInput() *string
	VpcId() *string
	SetVpcId(val *string)
	VpcIdInput() *string
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
	PutAddJobs(value interface{})
	PutBootstrapScripts(value interface{})
	PutComponentList(value interface{})
	PutTimeouts(value *MrsClusterV1Timeouts)
	ResetAddJobs()
	ResetBootstrapScripts()
	ResetClusterAdminSecret()
	ResetClusterType()
	ResetCoreDataVolumeCount()
	ResetCoreDataVolumeSize()
	ResetCoreDataVolumeType()
	ResetId()
	ResetLogCollection()
	ResetMasterDataVolumeCount()
	ResetMasterDataVolumeSize()
	ResetMasterDataVolumeType()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetRegion()
	ResetTags()
	ResetTimeouts()
	ResetVolumeSize()
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

// The jsii proxy struct for MrsClusterV1
type jsiiProxy_MrsClusterV1 struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_MrsClusterV1) AddJobs() MrsClusterV1AddJobsList {
	var returns MrsClusterV1AddJobsList
	_jsii_.Get(
		j,
		"addJobs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrsClusterV1) AddJobsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"addJobsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrsClusterV1) AvailableZoneId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"availableZoneId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrsClusterV1) AvailableZoneIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"availableZoneIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrsClusterV1) AvailableZoneName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"availableZoneName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrsClusterV1) BillingType() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"billingType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrsClusterV1) BillingTypeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"billingTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrsClusterV1) BootstrapScripts() MrsClusterV1BootstrapScriptsList {
	var returns MrsClusterV1BootstrapScriptsList
	_jsii_.Get(
		j,
		"bootstrapScripts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrsClusterV1) BootstrapScriptsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"bootstrapScriptsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrsClusterV1) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrsClusterV1) ChargingStartTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"chargingStartTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrsClusterV1) ClusterAdminSecret() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterAdminSecret",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrsClusterV1) ClusterAdminSecretInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterAdminSecretInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrsClusterV1) ClusterId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrsClusterV1) ClusterName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrsClusterV1) ClusterNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrsClusterV1) ClusterState() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrsClusterV1) ClusterType() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"clusterType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrsClusterV1) ClusterTypeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"clusterTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrsClusterV1) ClusterVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrsClusterV1) ClusterVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrsClusterV1) ComponentList() MrsClusterV1ComponentListList {
	var returns MrsClusterV1ComponentListList
	_jsii_.Get(
		j,
		"componentList",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrsClusterV1) ComponentListInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"componentListInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrsClusterV1) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrsClusterV1) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrsClusterV1) CoreDataVolumeCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"coreDataVolumeCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrsClusterV1) CoreDataVolumeCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"coreDataVolumeCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrsClusterV1) CoreDataVolumeSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"coreDataVolumeSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrsClusterV1) CoreDataVolumeSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"coreDataVolumeSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrsClusterV1) CoreDataVolumeType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"coreDataVolumeType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrsClusterV1) CoreDataVolumeTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"coreDataVolumeTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrsClusterV1) CoreNodeNum() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"coreNodeNum",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrsClusterV1) CoreNodeNumInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"coreNodeNumInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrsClusterV1) CoreNodeProductId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"coreNodeProductId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrsClusterV1) CoreNodeSize() *string {
	var returns *string
	_jsii_.Get(
		j,
		"coreNodeSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrsClusterV1) CoreNodeSizeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"coreNodeSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrsClusterV1) CoreNodeSpecId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"coreNodeSpecId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrsClusterV1) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrsClusterV1) CreateAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrsClusterV1) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrsClusterV1) DeploymentId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deploymentId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrsClusterV1) ErrorInfo() *string {
	var returns *string
	_jsii_.Get(
		j,
		"errorInfo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrsClusterV1) ExternalAlternateIp() *string {
	var returns *string
	_jsii_.Get(
		j,
		"externalAlternateIp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrsClusterV1) ExternalIp() *string {
	var returns *string
	_jsii_.Get(
		j,
		"externalIp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrsClusterV1) Fee() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fee",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrsClusterV1) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrsClusterV1) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrsClusterV1) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrsClusterV1) HadoopVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hadoopVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrsClusterV1) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrsClusterV1) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrsClusterV1) InstanceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrsClusterV1) InternalIp() *string {
	var returns *string
	_jsii_.Get(
		j,
		"internalIp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrsClusterV1) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrsClusterV1) LogCollection() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"logCollection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrsClusterV1) LogCollectionInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"logCollectionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrsClusterV1) MasterDataVolumeCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"masterDataVolumeCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrsClusterV1) MasterDataVolumeCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"masterDataVolumeCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrsClusterV1) MasterDataVolumeSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"masterDataVolumeSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrsClusterV1) MasterDataVolumeSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"masterDataVolumeSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrsClusterV1) MasterDataVolumeType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"masterDataVolumeType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrsClusterV1) MasterDataVolumeTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"masterDataVolumeTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrsClusterV1) MasterNodeIp() *string {
	var returns *string
	_jsii_.Get(
		j,
		"masterNodeIp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrsClusterV1) MasterNodeNum() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"masterNodeNum",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrsClusterV1) MasterNodeNumInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"masterNodeNumInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrsClusterV1) MasterNodeProductId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"masterNodeProductId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrsClusterV1) MasterNodeSize() *string {
	var returns *string
	_jsii_.Get(
		j,
		"masterNodeSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrsClusterV1) MasterNodeSizeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"masterNodeSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrsClusterV1) MasterNodeSpecId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"masterNodeSpecId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrsClusterV1) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrsClusterV1) NodePublicCertName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nodePublicCertName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrsClusterV1) NodePublicCertNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nodePublicCertNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrsClusterV1) OrderId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"orderId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrsClusterV1) PrivateIpFirst() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateIpFirst",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrsClusterV1) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrsClusterV1) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrsClusterV1) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrsClusterV1) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrsClusterV1) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrsClusterV1) Remark() *string {
	var returns *string
	_jsii_.Get(
		j,
		"remark",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrsClusterV1) SafeMode() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"safeMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrsClusterV1) SafeModeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"safeModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrsClusterV1) SecurityGroupsId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"securityGroupsId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrsClusterV1) SlaveSecurityGroupsId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"slaveSecurityGroupsId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrsClusterV1) SubnetId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrsClusterV1) SubnetIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrsClusterV1) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrsClusterV1) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrsClusterV1) TenantId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tenantId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrsClusterV1) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrsClusterV1) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrsClusterV1) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrsClusterV1) Timeouts() MrsClusterV1TimeoutsOutputReference {
	var returns MrsClusterV1TimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrsClusterV1) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrsClusterV1) UpdateAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrsClusterV1) Vnc() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vnc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrsClusterV1) VolumeSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"volumeSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrsClusterV1) VolumeSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"volumeSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrsClusterV1) VolumeType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"volumeType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrsClusterV1) VolumeTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"volumeTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrsClusterV1) VpcId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MrsClusterV1) VpcIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcIdInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.34.3/docs/resources/mrs_cluster_v1 opentelekomcloud_mrs_cluster_v1} Resource.
func NewMrsClusterV1(scope constructs.Construct, id *string, config *MrsClusterV1Config) MrsClusterV1 {
	_init_.Initialize()

	if err := validateNewMrsClusterV1Parameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_MrsClusterV1{}

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.mrsClusterV1.MrsClusterV1",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.34.3/docs/resources/mrs_cluster_v1 opentelekomcloud_mrs_cluster_v1} Resource.
func NewMrsClusterV1_Override(m MrsClusterV1, scope constructs.Construct, id *string, config *MrsClusterV1Config) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.mrsClusterV1.MrsClusterV1",
		[]interface{}{scope, id, config},
		m,
	)
}

func (j *jsiiProxy_MrsClusterV1)SetAvailableZoneId(val *string) {
	if err := j.validateSetAvailableZoneIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"availableZoneId",
		val,
	)
}

func (j *jsiiProxy_MrsClusterV1)SetBillingType(val *float64) {
	if err := j.validateSetBillingTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"billingType",
		val,
	)
}

func (j *jsiiProxy_MrsClusterV1)SetClusterAdminSecret(val *string) {
	if err := j.validateSetClusterAdminSecretParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clusterAdminSecret",
		val,
	)
}

func (j *jsiiProxy_MrsClusterV1)SetClusterName(val *string) {
	if err := j.validateSetClusterNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clusterName",
		val,
	)
}

func (j *jsiiProxy_MrsClusterV1)SetClusterType(val *float64) {
	if err := j.validateSetClusterTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clusterType",
		val,
	)
}

func (j *jsiiProxy_MrsClusterV1)SetClusterVersion(val *string) {
	if err := j.validateSetClusterVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clusterVersion",
		val,
	)
}

func (j *jsiiProxy_MrsClusterV1)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_MrsClusterV1)SetCoreDataVolumeCount(val *float64) {
	if err := j.validateSetCoreDataVolumeCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"coreDataVolumeCount",
		val,
	)
}

func (j *jsiiProxy_MrsClusterV1)SetCoreDataVolumeSize(val *float64) {
	if err := j.validateSetCoreDataVolumeSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"coreDataVolumeSize",
		val,
	)
}

func (j *jsiiProxy_MrsClusterV1)SetCoreDataVolumeType(val *string) {
	if err := j.validateSetCoreDataVolumeTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"coreDataVolumeType",
		val,
	)
}

func (j *jsiiProxy_MrsClusterV1)SetCoreNodeNum(val *float64) {
	if err := j.validateSetCoreNodeNumParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"coreNodeNum",
		val,
	)
}

func (j *jsiiProxy_MrsClusterV1)SetCoreNodeSize(val *string) {
	if err := j.validateSetCoreNodeSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"coreNodeSize",
		val,
	)
}

func (j *jsiiProxy_MrsClusterV1)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_MrsClusterV1)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_MrsClusterV1)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_MrsClusterV1)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_MrsClusterV1)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_MrsClusterV1)SetLogCollection(val *float64) {
	if err := j.validateSetLogCollectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"logCollection",
		val,
	)
}

func (j *jsiiProxy_MrsClusterV1)SetMasterDataVolumeCount(val *float64) {
	if err := j.validateSetMasterDataVolumeCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"masterDataVolumeCount",
		val,
	)
}

func (j *jsiiProxy_MrsClusterV1)SetMasterDataVolumeSize(val *float64) {
	if err := j.validateSetMasterDataVolumeSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"masterDataVolumeSize",
		val,
	)
}

func (j *jsiiProxy_MrsClusterV1)SetMasterDataVolumeType(val *string) {
	if err := j.validateSetMasterDataVolumeTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"masterDataVolumeType",
		val,
	)
}

func (j *jsiiProxy_MrsClusterV1)SetMasterNodeNum(val *float64) {
	if err := j.validateSetMasterNodeNumParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"masterNodeNum",
		val,
	)
}

func (j *jsiiProxy_MrsClusterV1)SetMasterNodeSize(val *string) {
	if err := j.validateSetMasterNodeSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"masterNodeSize",
		val,
	)
}

func (j *jsiiProxy_MrsClusterV1)SetNodePublicCertName(val *string) {
	if err := j.validateSetNodePublicCertNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nodePublicCertName",
		val,
	)
}

func (j *jsiiProxy_MrsClusterV1)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_MrsClusterV1)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_MrsClusterV1)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_MrsClusterV1)SetSafeMode(val *float64) {
	if err := j.validateSetSafeModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"safeMode",
		val,
	)
}

func (j *jsiiProxy_MrsClusterV1)SetSubnetId(val *string) {
	if err := j.validateSetSubnetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"subnetId",
		val,
	)
}

func (j *jsiiProxy_MrsClusterV1)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_MrsClusterV1)SetVolumeSize(val *float64) {
	if err := j.validateSetVolumeSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"volumeSize",
		val,
	)
}

func (j *jsiiProxy_MrsClusterV1)SetVolumeType(val *string) {
	if err := j.validateSetVolumeTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"volumeType",
		val,
	)
}

func (j *jsiiProxy_MrsClusterV1)SetVpcId(val *string) {
	if err := j.validateSetVpcIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vpcId",
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
func MrsClusterV1_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateMrsClusterV1_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-opentelekomcloud.mrsClusterV1.MrsClusterV1",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func MrsClusterV1_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateMrsClusterV1_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-opentelekomcloud.mrsClusterV1.MrsClusterV1",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func MrsClusterV1_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateMrsClusterV1_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-opentelekomcloud.mrsClusterV1.MrsClusterV1",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func MrsClusterV1_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-opentelekomcloud.mrsClusterV1.MrsClusterV1",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (m *jsiiProxy_MrsClusterV1) AddOverride(path *string, value interface{}) {
	if err := m.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (m *jsiiProxy_MrsClusterV1) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := m.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MrsClusterV1) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := m.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MrsClusterV1) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := m.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		m,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MrsClusterV1) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := m.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		m,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MrsClusterV1) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := m.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		m,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MrsClusterV1) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := m.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		m,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MrsClusterV1) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := m.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		m,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MrsClusterV1) GetStringAttribute(terraformAttribute *string) *string {
	if err := m.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		m,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MrsClusterV1) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := m.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		m,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MrsClusterV1) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := m.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MrsClusterV1) OverrideLogicalId(newLogicalId *string) {
	if err := m.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (m *jsiiProxy_MrsClusterV1) PutAddJobs(value interface{}) {
	if err := m.validatePutAddJobsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putAddJobs",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MrsClusterV1) PutBootstrapScripts(value interface{}) {
	if err := m.validatePutBootstrapScriptsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putBootstrapScripts",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MrsClusterV1) PutComponentList(value interface{}) {
	if err := m.validatePutComponentListParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putComponentList",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MrsClusterV1) PutTimeouts(value *MrsClusterV1Timeouts) {
	if err := m.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MrsClusterV1) ResetAddJobs() {
	_jsii_.InvokeVoid(
		m,
		"resetAddJobs",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MrsClusterV1) ResetBootstrapScripts() {
	_jsii_.InvokeVoid(
		m,
		"resetBootstrapScripts",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MrsClusterV1) ResetClusterAdminSecret() {
	_jsii_.InvokeVoid(
		m,
		"resetClusterAdminSecret",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MrsClusterV1) ResetClusterType() {
	_jsii_.InvokeVoid(
		m,
		"resetClusterType",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MrsClusterV1) ResetCoreDataVolumeCount() {
	_jsii_.InvokeVoid(
		m,
		"resetCoreDataVolumeCount",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MrsClusterV1) ResetCoreDataVolumeSize() {
	_jsii_.InvokeVoid(
		m,
		"resetCoreDataVolumeSize",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MrsClusterV1) ResetCoreDataVolumeType() {
	_jsii_.InvokeVoid(
		m,
		"resetCoreDataVolumeType",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MrsClusterV1) ResetId() {
	_jsii_.InvokeVoid(
		m,
		"resetId",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MrsClusterV1) ResetLogCollection() {
	_jsii_.InvokeVoid(
		m,
		"resetLogCollection",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MrsClusterV1) ResetMasterDataVolumeCount() {
	_jsii_.InvokeVoid(
		m,
		"resetMasterDataVolumeCount",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MrsClusterV1) ResetMasterDataVolumeSize() {
	_jsii_.InvokeVoid(
		m,
		"resetMasterDataVolumeSize",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MrsClusterV1) ResetMasterDataVolumeType() {
	_jsii_.InvokeVoid(
		m,
		"resetMasterDataVolumeType",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MrsClusterV1) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		m,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MrsClusterV1) ResetRegion() {
	_jsii_.InvokeVoid(
		m,
		"resetRegion",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MrsClusterV1) ResetTags() {
	_jsii_.InvokeVoid(
		m,
		"resetTags",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MrsClusterV1) ResetTimeouts() {
	_jsii_.InvokeVoid(
		m,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MrsClusterV1) ResetVolumeSize() {
	_jsii_.InvokeVoid(
		m,
		"resetVolumeSize",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MrsClusterV1) ResetVolumeType() {
	_jsii_.InvokeVoid(
		m,
		"resetVolumeType",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MrsClusterV1) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MrsClusterV1) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MrsClusterV1) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MrsClusterV1) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

