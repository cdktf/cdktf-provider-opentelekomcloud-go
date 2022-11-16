package ccenodev3

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v4/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v4/ccenodev3/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://www.terraform.io/docs/providers/opentelekomcloud/r/cce_node_v3 opentelekomcloud_cce_node_v3}.
type CceNodeV3 interface {
	cdktf.TerraformResource
	Annotations() *map[string]*string
	SetAnnotations(val *map[string]*string)
	AnnotationsInput() *map[string]*string
	AvailabilityZone() *string
	SetAvailabilityZone(val *string)
	AvailabilityZoneInput() *string
	BandwidthChargeMode() *string
	SetBandwidthChargeMode(val *string)
	BandwidthChargeModeInput() *string
	BandwidthSize() *float64
	SetBandwidthSize(val *float64)
	BandwidthSizeInput() *float64
	BillingMode() *float64
	SetBillingMode(val *float64)
	BillingModeInput() *float64
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
	Count() *float64
	// Experimental.
	SetCount(val *float64)
	DataVolumes() CceNodeV3DataVolumesList
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
	EcsPerformanceType() *string
	SetEcsPerformanceType(val *string)
	EcsPerformanceTypeInput() *string
	EipCount() *float64
	SetEipCount(val *float64)
	EipCountInput() *float64
	EipIds() *[]*string
	SetEipIds(val *[]*string)
	EipIdsInput() *[]*string
	ExtendParamChargingMode() *float64
	SetExtendParamChargingMode(val *float64)
	ExtendParamChargingModeInput() *float64
	FlavorId() *string
	SetFlavorId(val *string)
	FlavorIdInput() *string
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
	Iptype() *string
	SetIptype(val *string)
	IptypeInput() *string
	K8STags() *map[string]*string
	SetK8STags(val *map[string]*string)
	K8STagsInput() *map[string]*string
	KeyPair() *string
	SetKeyPair(val *string)
	KeyPairInput() *string
	Labels() *map[string]*string
	SetLabels(val *map[string]*string)
	LabelsInput() *map[string]*string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	MaxPods() *float64
	SetMaxPods(val *float64)
	MaxPodsInput() *float64
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	OrderId() *string
	SetOrderId(val *string)
	OrderIdInput() *string
	Os() *string
	SetOs(val *string)
	OsInput() *string
	Postinstall() *string
	SetPostinstall(val *string)
	PostinstallInput() *string
	Preinstall() *string
	SetPreinstall(val *string)
	PreinstallInput() *string
	PrivateIp() *string
	SetPrivateIp(val *string)
	PrivateIpInput() *string
	ProductId() *string
	SetProductId(val *string)
	ProductIdInput() *string
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	PublicIp() *string
	PublicKey() *string
	SetPublicKey(val *string)
	PublicKeyInput() *string
	// Experimental.
	RawOverrides() interface{}
	Region() *string
	SetRegion(val *string)
	RegionInput() *string
	RootVolume() CceNodeV3RootVolumeOutputReference
	RootVolumeInput() *CceNodeV3RootVolume
	ServerId() *string
	Sharetype() *string
	SetSharetype(val *string)
	SharetypeInput() *string
	Status() *string
	SubnetId() *string
	SetSubnetId(val *string)
	SubnetIdInput() *string
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsInput() *map[string]*string
	Taints() CceNodeV3TaintsList
	TaintsInput() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() CceNodeV3TimeoutsOutputReference
	TimeoutsInput() interface{}
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
	PutDataVolumes(value interface{})
	PutRootVolume(value *CceNodeV3RootVolume)
	PutTaints(value interface{})
	PutTimeouts(value *CceNodeV3Timeouts)
	ResetAnnotations()
	ResetBandwidthChargeMode()
	ResetBandwidthSize()
	ResetBillingMode()
	ResetDockerBaseSize()
	ResetDockerLvmConfigOverride()
	ResetEcsPerformanceType()
	ResetEipCount()
	ResetEipIds()
	ResetExtendParamChargingMode()
	ResetId()
	ResetIptype()
	ResetK8STags()
	ResetLabels()
	ResetMaxPods()
	ResetName()
	ResetOrderId()
	ResetOs()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPostinstall()
	ResetPreinstall()
	ResetPrivateIp()
	ResetProductId()
	ResetPublicKey()
	ResetRegion()
	ResetSharetype()
	ResetSubnetId()
	ResetTags()
	ResetTaints()
	ResetTimeouts()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for CceNodeV3
type jsiiProxy_CceNodeV3 struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_CceNodeV3) Annotations() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"annotations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceNodeV3) AnnotationsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"annotationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceNodeV3) AvailabilityZone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"availabilityZone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceNodeV3) AvailabilityZoneInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"availabilityZoneInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceNodeV3) BandwidthChargeMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bandwidthChargeMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceNodeV3) BandwidthChargeModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bandwidthChargeModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceNodeV3) BandwidthSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bandwidthSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceNodeV3) BandwidthSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bandwidthSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceNodeV3) BillingMode() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"billingMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceNodeV3) BillingModeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"billingModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceNodeV3) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceNodeV3) ClusterId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceNodeV3) ClusterIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceNodeV3) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceNodeV3) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceNodeV3) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceNodeV3) DataVolumes() CceNodeV3DataVolumesList {
	var returns CceNodeV3DataVolumesList
	_jsii_.Get(
		j,
		"dataVolumes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceNodeV3) DataVolumesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dataVolumesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceNodeV3) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceNodeV3) DockerBaseSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"dockerBaseSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceNodeV3) DockerBaseSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"dockerBaseSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceNodeV3) DockerLvmConfigOverride() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dockerLvmConfigOverride",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceNodeV3) DockerLvmConfigOverrideInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dockerLvmConfigOverrideInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceNodeV3) EcsPerformanceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ecsPerformanceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceNodeV3) EcsPerformanceTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ecsPerformanceTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceNodeV3) EipCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"eipCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceNodeV3) EipCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"eipCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceNodeV3) EipIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"eipIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceNodeV3) EipIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"eipIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceNodeV3) ExtendParamChargingMode() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"extendParamChargingMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceNodeV3) ExtendParamChargingModeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"extendParamChargingModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceNodeV3) FlavorId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"flavorId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceNodeV3) FlavorIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"flavorIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceNodeV3) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceNodeV3) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceNodeV3) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceNodeV3) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceNodeV3) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceNodeV3) Iptype() *string {
	var returns *string
	_jsii_.Get(
		j,
		"iptype",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceNodeV3) IptypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"iptypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceNodeV3) K8STags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"k8STags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceNodeV3) K8STagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"k8STagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceNodeV3) KeyPair() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyPair",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceNodeV3) KeyPairInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyPairInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceNodeV3) Labels() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceNodeV3) LabelsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labelsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceNodeV3) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceNodeV3) MaxPods() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxPods",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceNodeV3) MaxPodsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxPodsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceNodeV3) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceNodeV3) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceNodeV3) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceNodeV3) OrderId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"orderId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceNodeV3) OrderIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"orderIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceNodeV3) Os() *string {
	var returns *string
	_jsii_.Get(
		j,
		"os",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceNodeV3) OsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"osInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceNodeV3) Postinstall() *string {
	var returns *string
	_jsii_.Get(
		j,
		"postinstall",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceNodeV3) PostinstallInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"postinstallInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceNodeV3) Preinstall() *string {
	var returns *string
	_jsii_.Get(
		j,
		"preinstall",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceNodeV3) PreinstallInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"preinstallInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceNodeV3) PrivateIp() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateIp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceNodeV3) PrivateIpInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateIpInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceNodeV3) ProductId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"productId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceNodeV3) ProductIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"productIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceNodeV3) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceNodeV3) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceNodeV3) PublicIp() *string {
	var returns *string
	_jsii_.Get(
		j,
		"publicIp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceNodeV3) PublicKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"publicKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceNodeV3) PublicKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"publicKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceNodeV3) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceNodeV3) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceNodeV3) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceNodeV3) RootVolume() CceNodeV3RootVolumeOutputReference {
	var returns CceNodeV3RootVolumeOutputReference
	_jsii_.Get(
		j,
		"rootVolume",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceNodeV3) RootVolumeInput() *CceNodeV3RootVolume {
	var returns *CceNodeV3RootVolume
	_jsii_.Get(
		j,
		"rootVolumeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceNodeV3) ServerId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serverId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceNodeV3) Sharetype() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sharetype",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceNodeV3) SharetypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sharetypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceNodeV3) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceNodeV3) SubnetId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceNodeV3) SubnetIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceNodeV3) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceNodeV3) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceNodeV3) Taints() CceNodeV3TaintsList {
	var returns CceNodeV3TaintsList
	_jsii_.Get(
		j,
		"taints",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceNodeV3) TaintsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"taintsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceNodeV3) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceNodeV3) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceNodeV3) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceNodeV3) Timeouts() CceNodeV3TimeoutsOutputReference {
	var returns CceNodeV3TimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceNodeV3) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://www.terraform.io/docs/providers/opentelekomcloud/r/cce_node_v3 opentelekomcloud_cce_node_v3} Resource.
func NewCceNodeV3(scope constructs.Construct, id *string, config *CceNodeV3Config) CceNodeV3 {
	_init_.Initialize()

	if err := validateNewCceNodeV3Parameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_CceNodeV3{}

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.cceNodeV3.CceNodeV3",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/opentelekomcloud/r/cce_node_v3 opentelekomcloud_cce_node_v3} Resource.
func NewCceNodeV3_Override(c CceNodeV3, scope constructs.Construct, id *string, config *CceNodeV3Config) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.cceNodeV3.CceNodeV3",
		[]interface{}{scope, id, config},
		c,
	)
}

func (j *jsiiProxy_CceNodeV3)SetAnnotations(val *map[string]*string) {
	if err := j.validateSetAnnotationsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"annotations",
		val,
	)
}

func (j *jsiiProxy_CceNodeV3)SetAvailabilityZone(val *string) {
	if err := j.validateSetAvailabilityZoneParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"availabilityZone",
		val,
	)
}

func (j *jsiiProxy_CceNodeV3)SetBandwidthChargeMode(val *string) {
	if err := j.validateSetBandwidthChargeModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bandwidthChargeMode",
		val,
	)
}

func (j *jsiiProxy_CceNodeV3)SetBandwidthSize(val *float64) {
	if err := j.validateSetBandwidthSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bandwidthSize",
		val,
	)
}

func (j *jsiiProxy_CceNodeV3)SetBillingMode(val *float64) {
	if err := j.validateSetBillingModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"billingMode",
		val,
	)
}

func (j *jsiiProxy_CceNodeV3)SetClusterId(val *string) {
	if err := j.validateSetClusterIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clusterId",
		val,
	)
}

func (j *jsiiProxy_CceNodeV3)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_CceNodeV3)SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_CceNodeV3)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_CceNodeV3)SetDockerBaseSize(val *float64) {
	if err := j.validateSetDockerBaseSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dockerBaseSize",
		val,
	)
}

func (j *jsiiProxy_CceNodeV3)SetDockerLvmConfigOverride(val *string) {
	if err := j.validateSetDockerLvmConfigOverrideParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dockerLvmConfigOverride",
		val,
	)
}

func (j *jsiiProxy_CceNodeV3)SetEcsPerformanceType(val *string) {
	if err := j.validateSetEcsPerformanceTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ecsPerformanceType",
		val,
	)
}

func (j *jsiiProxy_CceNodeV3)SetEipCount(val *float64) {
	if err := j.validateSetEipCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"eipCount",
		val,
	)
}

func (j *jsiiProxy_CceNodeV3)SetEipIds(val *[]*string) {
	if err := j.validateSetEipIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"eipIds",
		val,
	)
}

func (j *jsiiProxy_CceNodeV3)SetExtendParamChargingMode(val *float64) {
	if err := j.validateSetExtendParamChargingModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"extendParamChargingMode",
		val,
	)
}

func (j *jsiiProxy_CceNodeV3)SetFlavorId(val *string) {
	if err := j.validateSetFlavorIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"flavorId",
		val,
	)
}

func (j *jsiiProxy_CceNodeV3)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_CceNodeV3)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_CceNodeV3)SetIptype(val *string) {
	if err := j.validateSetIptypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"iptype",
		val,
	)
}

func (j *jsiiProxy_CceNodeV3)SetK8STags(val *map[string]*string) {
	if err := j.validateSetK8STagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"k8STags",
		val,
	)
}

func (j *jsiiProxy_CceNodeV3)SetKeyPair(val *string) {
	if err := j.validateSetKeyPairParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"keyPair",
		val,
	)
}

func (j *jsiiProxy_CceNodeV3)SetLabels(val *map[string]*string) {
	if err := j.validateSetLabelsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"labels",
		val,
	)
}

func (j *jsiiProxy_CceNodeV3)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_CceNodeV3)SetMaxPods(val *float64) {
	if err := j.validateSetMaxPodsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxPods",
		val,
	)
}

func (j *jsiiProxy_CceNodeV3)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CceNodeV3)SetOrderId(val *string) {
	if err := j.validateSetOrderIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"orderId",
		val,
	)
}

func (j *jsiiProxy_CceNodeV3)SetOs(val *string) {
	if err := j.validateSetOsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"os",
		val,
	)
}

func (j *jsiiProxy_CceNodeV3)SetPostinstall(val *string) {
	if err := j.validateSetPostinstallParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"postinstall",
		val,
	)
}

func (j *jsiiProxy_CceNodeV3)SetPreinstall(val *string) {
	if err := j.validateSetPreinstallParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"preinstall",
		val,
	)
}

func (j *jsiiProxy_CceNodeV3)SetPrivateIp(val *string) {
	if err := j.validateSetPrivateIpParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"privateIp",
		val,
	)
}

func (j *jsiiProxy_CceNodeV3)SetProductId(val *string) {
	if err := j.validateSetProductIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"productId",
		val,
	)
}

func (j *jsiiProxy_CceNodeV3)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_CceNodeV3)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_CceNodeV3)SetPublicKey(val *string) {
	if err := j.validateSetPublicKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"publicKey",
		val,
	)
}

func (j *jsiiProxy_CceNodeV3)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_CceNodeV3)SetSharetype(val *string) {
	if err := j.validateSetSharetypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sharetype",
		val,
	)
}

func (j *jsiiProxy_CceNodeV3)SetSubnetId(val *string) {
	if err := j.validateSetSubnetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"subnetId",
		val,
	)
}

func (j *jsiiProxy_CceNodeV3)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
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
func CceNodeV3_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCceNodeV3_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-opentelekomcloud.cceNodeV3.CceNodeV3",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func CceNodeV3_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCceNodeV3_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-opentelekomcloud.cceNodeV3.CceNodeV3",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func CceNodeV3_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCceNodeV3_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-opentelekomcloud.cceNodeV3.CceNodeV3",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CceNodeV3_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-opentelekomcloud.cceNodeV3.CceNodeV3",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CceNodeV3) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CceNodeV3) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_CceNodeV3) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CceNodeV3) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_CceNodeV3) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_CceNodeV3) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_CceNodeV3) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_CceNodeV3) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_CceNodeV3) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_CceNodeV3) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_CceNodeV3) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CceNodeV3) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CceNodeV3) PutDataVolumes(value interface{}) {
	if err := c.validatePutDataVolumesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putDataVolumes",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CceNodeV3) PutRootVolume(value *CceNodeV3RootVolume) {
	if err := c.validatePutRootVolumeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putRootVolume",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CceNodeV3) PutTaints(value interface{}) {
	if err := c.validatePutTaintsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putTaints",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CceNodeV3) PutTimeouts(value *CceNodeV3Timeouts) {
	if err := c.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CceNodeV3) ResetAnnotations() {
	_jsii_.InvokeVoid(
		c,
		"resetAnnotations",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CceNodeV3) ResetBandwidthChargeMode() {
	_jsii_.InvokeVoid(
		c,
		"resetBandwidthChargeMode",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CceNodeV3) ResetBandwidthSize() {
	_jsii_.InvokeVoid(
		c,
		"resetBandwidthSize",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CceNodeV3) ResetBillingMode() {
	_jsii_.InvokeVoid(
		c,
		"resetBillingMode",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CceNodeV3) ResetDockerBaseSize() {
	_jsii_.InvokeVoid(
		c,
		"resetDockerBaseSize",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CceNodeV3) ResetDockerLvmConfigOverride() {
	_jsii_.InvokeVoid(
		c,
		"resetDockerLvmConfigOverride",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CceNodeV3) ResetEcsPerformanceType() {
	_jsii_.InvokeVoid(
		c,
		"resetEcsPerformanceType",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CceNodeV3) ResetEipCount() {
	_jsii_.InvokeVoid(
		c,
		"resetEipCount",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CceNodeV3) ResetEipIds() {
	_jsii_.InvokeVoid(
		c,
		"resetEipIds",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CceNodeV3) ResetExtendParamChargingMode() {
	_jsii_.InvokeVoid(
		c,
		"resetExtendParamChargingMode",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CceNodeV3) ResetId() {
	_jsii_.InvokeVoid(
		c,
		"resetId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CceNodeV3) ResetIptype() {
	_jsii_.InvokeVoid(
		c,
		"resetIptype",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CceNodeV3) ResetK8STags() {
	_jsii_.InvokeVoid(
		c,
		"resetK8STags",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CceNodeV3) ResetLabels() {
	_jsii_.InvokeVoid(
		c,
		"resetLabels",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CceNodeV3) ResetMaxPods() {
	_jsii_.InvokeVoid(
		c,
		"resetMaxPods",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CceNodeV3) ResetName() {
	_jsii_.InvokeVoid(
		c,
		"resetName",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CceNodeV3) ResetOrderId() {
	_jsii_.InvokeVoid(
		c,
		"resetOrderId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CceNodeV3) ResetOs() {
	_jsii_.InvokeVoid(
		c,
		"resetOs",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CceNodeV3) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		c,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CceNodeV3) ResetPostinstall() {
	_jsii_.InvokeVoid(
		c,
		"resetPostinstall",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CceNodeV3) ResetPreinstall() {
	_jsii_.InvokeVoid(
		c,
		"resetPreinstall",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CceNodeV3) ResetPrivateIp() {
	_jsii_.InvokeVoid(
		c,
		"resetPrivateIp",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CceNodeV3) ResetProductId() {
	_jsii_.InvokeVoid(
		c,
		"resetProductId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CceNodeV3) ResetPublicKey() {
	_jsii_.InvokeVoid(
		c,
		"resetPublicKey",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CceNodeV3) ResetRegion() {
	_jsii_.InvokeVoid(
		c,
		"resetRegion",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CceNodeV3) ResetSharetype() {
	_jsii_.InvokeVoid(
		c,
		"resetSharetype",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CceNodeV3) ResetSubnetId() {
	_jsii_.InvokeVoid(
		c,
		"resetSubnetId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CceNodeV3) ResetTags() {
	_jsii_.InvokeVoid(
		c,
		"resetTags",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CceNodeV3) ResetTaints() {
	_jsii_.InvokeVoid(
		c,
		"resetTaints",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CceNodeV3) ResetTimeouts() {
	_jsii_.InvokeVoid(
		c,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CceNodeV3) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CceNodeV3) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CceNodeV3) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CceNodeV3) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

