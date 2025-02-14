// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cceclusterv3

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v11/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v11/cceclusterv3/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.31/docs/resources/cce_cluster_v3 opentelekomcloud_cce_cluster_v3}.
type CceClusterV3 interface {
	cdktf.TerraformResource
	Annotations() *map[string]*string
	SetAnnotations(val *map[string]*string)
	AnnotationsInput() *map[string]*string
	AuthenticatingProxy() CceClusterV3AuthenticatingProxyOutputReference
	AuthenticatingProxyCa() *string
	SetAuthenticatingProxyCa(val *string)
	AuthenticatingProxyCaInput() *string
	AuthenticatingProxyInput() *CceClusterV3AuthenticatingProxy
	AuthenticationMode() *string
	SetAuthenticationMode(val *string)
	AuthenticationModeInput() *string
	BillingMode() *float64
	SetBillingMode(val *float64)
	BillingModeInput() *float64
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	CertificateClusters() CceClusterV3CertificateClustersList
	CertificateUsers() CceClusterV3CertificateUsersList
	ClusterType() *string
	SetClusterType(val *string)
	ClusterTypeInput() *string
	ClusterVersion() *string
	SetClusterVersion(val *string)
	ClusterVersionInput() *string
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	ContainerNetworkCidr() *string
	SetContainerNetworkCidr(val *string)
	ContainerNetworkCidrInput() *string
	ContainerNetworkType() *string
	SetContainerNetworkType(val *string)
	ContainerNetworkTypeInput() *string
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	DeleteAllNetwork() *string
	SetDeleteAllNetwork(val *string)
	DeleteAllNetworkInput() *string
	DeleteAllStorage() *string
	SetDeleteAllStorage(val *string)
	DeleteAllStorageInput() *string
	DeleteEfs() *string
	SetDeleteEfs(val *string)
	DeleteEfsInput() *string
	DeleteEni() *string
	SetDeleteEni(val *string)
	DeleteEniInput() *string
	DeleteEvs() *string
	SetDeleteEvs(val *string)
	DeleteEvsInput() *string
	DeleteNet() *string
	SetDeleteNet(val *string)
	DeleteNetInput() *string
	DeleteObs() *string
	SetDeleteObs(val *string)
	DeleteObsInput() *string
	DeleteSfs() *string
	SetDeleteSfs(val *string)
	DeleteSfsInput() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	Eip() *string
	SetEip(val *string)
	EipInput() *string
	EnableVolumeEncryption() interface{}
	SetEnableVolumeEncryption(val interface{})
	EnableVolumeEncryptionInput() interface{}
	EniSubnetCidr() *string
	SetEniSubnetCidr(val *string)
	EniSubnetCidrInput() *string
	EniSubnetId() *string
	SetEniSubnetId(val *string)
	EniSubnetIdInput() *string
	ExtendParam() *map[string]*string
	SetExtendParam(val *map[string]*string)
	ExtendParamInput() *map[string]*string
	External() *string
	ExternalOtc() *string
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
	HighwaySubnetId() *string
	SetHighwaySubnetId(val *string)
	HighwaySubnetIdInput() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	IgnoreAddons() interface{}
	SetIgnoreAddons(val interface{})
	IgnoreAddonsInput() interface{}
	IgnoreCertificateClustersData() interface{}
	SetIgnoreCertificateClustersData(val interface{})
	IgnoreCertificateClustersDataInput() interface{}
	IgnoreCertificateUsersData() interface{}
	SetIgnoreCertificateUsersData(val interface{})
	IgnoreCertificateUsersDataInput() interface{}
	InstalledAddons() *[]*string
	Internal() *string
	KubeProxyMode() *string
	SetKubeProxyMode(val *string)
	KubeProxyModeInput() *string
	KubernetesSvcIpRange() *string
	SetKubernetesSvcIpRange(val *string)
	KubernetesSvcIpRangeInput() *string
	Labels() *map[string]*string
	SetLabels(val *map[string]*string)
	LabelsInput() *map[string]*string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Masters() CceClusterV3MastersList
	MastersInput() interface{}
	MultiAz() interface{}
	SetMultiAz(val interface{})
	MultiAzInput() interface{}
	Name() *string
	SetName(val *string)
	NameInput() *string
	NoAddons() interface{}
	SetNoAddons(val interface{})
	NoAddonsInput() interface{}
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
	SecurityGroupControl() *string
	SecurityGroupId() *string
	SetSecurityGroupId(val *string)
	SecurityGroupIdInput() *string
	SecurityGroupNode() *string
	Status() *string
	SubnetId() *string
	SetSubnetId(val *string)
	SubnetIdInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() CceClusterV3TimeoutsOutputReference
	TimeoutsInput() interface{}
	VpcId() *string
	SetVpcId(val *string)
	VpcIdInput() *string
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
	PutAuthenticatingProxy(value *CceClusterV3AuthenticatingProxy)
	PutMasters(value interface{})
	PutTimeouts(value *CceClusterV3Timeouts)
	ResetAnnotations()
	ResetAuthenticatingProxy()
	ResetAuthenticatingProxyCa()
	ResetAuthenticationMode()
	ResetBillingMode()
	ResetClusterVersion()
	ResetContainerNetworkCidr()
	ResetDeleteAllNetwork()
	ResetDeleteAllStorage()
	ResetDeleteEfs()
	ResetDeleteEni()
	ResetDeleteEvs()
	ResetDeleteNet()
	ResetDeleteObs()
	ResetDeleteSfs()
	ResetDescription()
	ResetEip()
	ResetEnableVolumeEncryption()
	ResetEniSubnetCidr()
	ResetEniSubnetId()
	ResetExtendParam()
	ResetHighwaySubnetId()
	ResetId()
	ResetIgnoreAddons()
	ResetIgnoreCertificateClustersData()
	ResetIgnoreCertificateUsersData()
	ResetKubeProxyMode()
	ResetKubernetesSvcIpRange()
	ResetLabels()
	ResetMasters()
	ResetMultiAz()
	ResetNoAddons()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetRegion()
	ResetSecurityGroupId()
	ResetTimeouts()
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

// The jsii proxy struct for CceClusterV3
type jsiiProxy_CceClusterV3 struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_CceClusterV3) Annotations() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"annotations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceClusterV3) AnnotationsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"annotationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceClusterV3) AuthenticatingProxy() CceClusterV3AuthenticatingProxyOutputReference {
	var returns CceClusterV3AuthenticatingProxyOutputReference
	_jsii_.Get(
		j,
		"authenticatingProxy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceClusterV3) AuthenticatingProxyCa() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authenticatingProxyCa",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceClusterV3) AuthenticatingProxyCaInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authenticatingProxyCaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceClusterV3) AuthenticatingProxyInput() *CceClusterV3AuthenticatingProxy {
	var returns *CceClusterV3AuthenticatingProxy
	_jsii_.Get(
		j,
		"authenticatingProxyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceClusterV3) AuthenticationMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authenticationMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceClusterV3) AuthenticationModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authenticationModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceClusterV3) BillingMode() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"billingMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceClusterV3) BillingModeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"billingModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceClusterV3) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceClusterV3) CertificateClusters() CceClusterV3CertificateClustersList {
	var returns CceClusterV3CertificateClustersList
	_jsii_.Get(
		j,
		"certificateClusters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceClusterV3) CertificateUsers() CceClusterV3CertificateUsersList {
	var returns CceClusterV3CertificateUsersList
	_jsii_.Get(
		j,
		"certificateUsers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceClusterV3) ClusterType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceClusterV3) ClusterTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceClusterV3) ClusterVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceClusterV3) ClusterVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceClusterV3) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceClusterV3) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceClusterV3) ContainerNetworkCidr() *string {
	var returns *string
	_jsii_.Get(
		j,
		"containerNetworkCidr",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceClusterV3) ContainerNetworkCidrInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"containerNetworkCidrInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceClusterV3) ContainerNetworkType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"containerNetworkType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceClusterV3) ContainerNetworkTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"containerNetworkTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceClusterV3) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceClusterV3) DeleteAllNetwork() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deleteAllNetwork",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceClusterV3) DeleteAllNetworkInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deleteAllNetworkInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceClusterV3) DeleteAllStorage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deleteAllStorage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceClusterV3) DeleteAllStorageInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deleteAllStorageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceClusterV3) DeleteEfs() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deleteEfs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceClusterV3) DeleteEfsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deleteEfsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceClusterV3) DeleteEni() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deleteEni",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceClusterV3) DeleteEniInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deleteEniInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceClusterV3) DeleteEvs() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deleteEvs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceClusterV3) DeleteEvsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deleteEvsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceClusterV3) DeleteNet() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deleteNet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceClusterV3) DeleteNetInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deleteNetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceClusterV3) DeleteObs() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deleteObs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceClusterV3) DeleteObsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deleteObsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceClusterV3) DeleteSfs() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deleteSfs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceClusterV3) DeleteSfsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deleteSfsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceClusterV3) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceClusterV3) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceClusterV3) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceClusterV3) Eip() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eip",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceClusterV3) EipInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eipInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceClusterV3) EnableVolumeEncryption() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableVolumeEncryption",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceClusterV3) EnableVolumeEncryptionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableVolumeEncryptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceClusterV3) EniSubnetCidr() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eniSubnetCidr",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceClusterV3) EniSubnetCidrInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eniSubnetCidrInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceClusterV3) EniSubnetId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eniSubnetId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceClusterV3) EniSubnetIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eniSubnetIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceClusterV3) ExtendParam() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"extendParam",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceClusterV3) ExtendParamInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"extendParamInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceClusterV3) External() *string {
	var returns *string
	_jsii_.Get(
		j,
		"external",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceClusterV3) ExternalOtc() *string {
	var returns *string
	_jsii_.Get(
		j,
		"externalOtc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceClusterV3) FlavorId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"flavorId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceClusterV3) FlavorIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"flavorIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceClusterV3) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceClusterV3) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceClusterV3) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceClusterV3) HighwaySubnetId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"highwaySubnetId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceClusterV3) HighwaySubnetIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"highwaySubnetIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceClusterV3) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceClusterV3) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceClusterV3) IgnoreAddons() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ignoreAddons",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceClusterV3) IgnoreAddonsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ignoreAddonsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceClusterV3) IgnoreCertificateClustersData() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ignoreCertificateClustersData",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceClusterV3) IgnoreCertificateClustersDataInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ignoreCertificateClustersDataInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceClusterV3) IgnoreCertificateUsersData() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ignoreCertificateUsersData",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceClusterV3) IgnoreCertificateUsersDataInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ignoreCertificateUsersDataInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceClusterV3) InstalledAddons() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"installedAddons",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceClusterV3) Internal() *string {
	var returns *string
	_jsii_.Get(
		j,
		"internal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceClusterV3) KubeProxyMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kubeProxyMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceClusterV3) KubeProxyModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kubeProxyModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceClusterV3) KubernetesSvcIpRange() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kubernetesSvcIpRange",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceClusterV3) KubernetesSvcIpRangeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kubernetesSvcIpRangeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceClusterV3) Labels() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceClusterV3) LabelsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labelsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceClusterV3) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceClusterV3) Masters() CceClusterV3MastersList {
	var returns CceClusterV3MastersList
	_jsii_.Get(
		j,
		"masters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceClusterV3) MastersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"mastersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceClusterV3) MultiAz() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"multiAz",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceClusterV3) MultiAzInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"multiAzInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceClusterV3) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceClusterV3) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceClusterV3) NoAddons() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noAddons",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceClusterV3) NoAddonsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noAddonsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceClusterV3) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceClusterV3) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceClusterV3) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceClusterV3) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceClusterV3) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceClusterV3) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceClusterV3) SecurityGroupControl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"securityGroupControl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceClusterV3) SecurityGroupId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"securityGroupId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceClusterV3) SecurityGroupIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"securityGroupIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceClusterV3) SecurityGroupNode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"securityGroupNode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceClusterV3) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceClusterV3) SubnetId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceClusterV3) SubnetIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceClusterV3) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceClusterV3) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceClusterV3) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceClusterV3) Timeouts() CceClusterV3TimeoutsOutputReference {
	var returns CceClusterV3TimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceClusterV3) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceClusterV3) VpcId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CceClusterV3) VpcIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcIdInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.31/docs/resources/cce_cluster_v3 opentelekomcloud_cce_cluster_v3} Resource.
func NewCceClusterV3(scope constructs.Construct, id *string, config *CceClusterV3Config) CceClusterV3 {
	_init_.Initialize()

	if err := validateNewCceClusterV3Parameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_CceClusterV3{}

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.cceClusterV3.CceClusterV3",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.31/docs/resources/cce_cluster_v3 opentelekomcloud_cce_cluster_v3} Resource.
func NewCceClusterV3_Override(c CceClusterV3, scope constructs.Construct, id *string, config *CceClusterV3Config) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.cceClusterV3.CceClusterV3",
		[]interface{}{scope, id, config},
		c,
	)
}

func (j *jsiiProxy_CceClusterV3)SetAnnotations(val *map[string]*string) {
	if err := j.validateSetAnnotationsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"annotations",
		val,
	)
}

func (j *jsiiProxy_CceClusterV3)SetAuthenticatingProxyCa(val *string) {
	if err := j.validateSetAuthenticatingProxyCaParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"authenticatingProxyCa",
		val,
	)
}

func (j *jsiiProxy_CceClusterV3)SetAuthenticationMode(val *string) {
	if err := j.validateSetAuthenticationModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"authenticationMode",
		val,
	)
}

func (j *jsiiProxy_CceClusterV3)SetBillingMode(val *float64) {
	if err := j.validateSetBillingModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"billingMode",
		val,
	)
}

func (j *jsiiProxy_CceClusterV3)SetClusterType(val *string) {
	if err := j.validateSetClusterTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clusterType",
		val,
	)
}

func (j *jsiiProxy_CceClusterV3)SetClusterVersion(val *string) {
	if err := j.validateSetClusterVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clusterVersion",
		val,
	)
}

func (j *jsiiProxy_CceClusterV3)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_CceClusterV3)SetContainerNetworkCidr(val *string) {
	if err := j.validateSetContainerNetworkCidrParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"containerNetworkCidr",
		val,
	)
}

func (j *jsiiProxy_CceClusterV3)SetContainerNetworkType(val *string) {
	if err := j.validateSetContainerNetworkTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"containerNetworkType",
		val,
	)
}

func (j *jsiiProxy_CceClusterV3)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_CceClusterV3)SetDeleteAllNetwork(val *string) {
	if err := j.validateSetDeleteAllNetworkParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deleteAllNetwork",
		val,
	)
}

func (j *jsiiProxy_CceClusterV3)SetDeleteAllStorage(val *string) {
	if err := j.validateSetDeleteAllStorageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deleteAllStorage",
		val,
	)
}

func (j *jsiiProxy_CceClusterV3)SetDeleteEfs(val *string) {
	if err := j.validateSetDeleteEfsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deleteEfs",
		val,
	)
}

func (j *jsiiProxy_CceClusterV3)SetDeleteEni(val *string) {
	if err := j.validateSetDeleteEniParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deleteEni",
		val,
	)
}

func (j *jsiiProxy_CceClusterV3)SetDeleteEvs(val *string) {
	if err := j.validateSetDeleteEvsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deleteEvs",
		val,
	)
}

func (j *jsiiProxy_CceClusterV3)SetDeleteNet(val *string) {
	if err := j.validateSetDeleteNetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deleteNet",
		val,
	)
}

func (j *jsiiProxy_CceClusterV3)SetDeleteObs(val *string) {
	if err := j.validateSetDeleteObsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deleteObs",
		val,
	)
}

func (j *jsiiProxy_CceClusterV3)SetDeleteSfs(val *string) {
	if err := j.validateSetDeleteSfsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deleteSfs",
		val,
	)
}

func (j *jsiiProxy_CceClusterV3)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_CceClusterV3)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_CceClusterV3)SetEip(val *string) {
	if err := j.validateSetEipParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"eip",
		val,
	)
}

func (j *jsiiProxy_CceClusterV3)SetEnableVolumeEncryption(val interface{}) {
	if err := j.validateSetEnableVolumeEncryptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableVolumeEncryption",
		val,
	)
}

func (j *jsiiProxy_CceClusterV3)SetEniSubnetCidr(val *string) {
	if err := j.validateSetEniSubnetCidrParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"eniSubnetCidr",
		val,
	)
}

func (j *jsiiProxy_CceClusterV3)SetEniSubnetId(val *string) {
	if err := j.validateSetEniSubnetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"eniSubnetId",
		val,
	)
}

func (j *jsiiProxy_CceClusterV3)SetExtendParam(val *map[string]*string) {
	if err := j.validateSetExtendParamParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"extendParam",
		val,
	)
}

func (j *jsiiProxy_CceClusterV3)SetFlavorId(val *string) {
	if err := j.validateSetFlavorIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"flavorId",
		val,
	)
}

func (j *jsiiProxy_CceClusterV3)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_CceClusterV3)SetHighwaySubnetId(val *string) {
	if err := j.validateSetHighwaySubnetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"highwaySubnetId",
		val,
	)
}

func (j *jsiiProxy_CceClusterV3)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_CceClusterV3)SetIgnoreAddons(val interface{}) {
	if err := j.validateSetIgnoreAddonsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ignoreAddons",
		val,
	)
}

func (j *jsiiProxy_CceClusterV3)SetIgnoreCertificateClustersData(val interface{}) {
	if err := j.validateSetIgnoreCertificateClustersDataParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ignoreCertificateClustersData",
		val,
	)
}

func (j *jsiiProxy_CceClusterV3)SetIgnoreCertificateUsersData(val interface{}) {
	if err := j.validateSetIgnoreCertificateUsersDataParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ignoreCertificateUsersData",
		val,
	)
}

func (j *jsiiProxy_CceClusterV3)SetKubeProxyMode(val *string) {
	if err := j.validateSetKubeProxyModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kubeProxyMode",
		val,
	)
}

func (j *jsiiProxy_CceClusterV3)SetKubernetesSvcIpRange(val *string) {
	if err := j.validateSetKubernetesSvcIpRangeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kubernetesSvcIpRange",
		val,
	)
}

func (j *jsiiProxy_CceClusterV3)SetLabels(val *map[string]*string) {
	if err := j.validateSetLabelsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"labels",
		val,
	)
}

func (j *jsiiProxy_CceClusterV3)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_CceClusterV3)SetMultiAz(val interface{}) {
	if err := j.validateSetMultiAzParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"multiAz",
		val,
	)
}

func (j *jsiiProxy_CceClusterV3)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CceClusterV3)SetNoAddons(val interface{}) {
	if err := j.validateSetNoAddonsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noAddons",
		val,
	)
}

func (j *jsiiProxy_CceClusterV3)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_CceClusterV3)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_CceClusterV3)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_CceClusterV3)SetSecurityGroupId(val *string) {
	if err := j.validateSetSecurityGroupIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"securityGroupId",
		val,
	)
}

func (j *jsiiProxy_CceClusterV3)SetSubnetId(val *string) {
	if err := j.validateSetSubnetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"subnetId",
		val,
	)
}

func (j *jsiiProxy_CceClusterV3)SetVpcId(val *string) {
	if err := j.validateSetVpcIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vpcId",
		val,
	)
}

// Generates CDKTF code for importing a CceClusterV3 resource upon running "cdktf plan <stack-name>".
func CceClusterV3_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateCceClusterV3_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-opentelekomcloud.cceClusterV3.CceClusterV3",
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
func CceClusterV3_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCceClusterV3_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-opentelekomcloud.cceClusterV3.CceClusterV3",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func CceClusterV3_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCceClusterV3_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-opentelekomcloud.cceClusterV3.CceClusterV3",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func CceClusterV3_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCceClusterV3_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-opentelekomcloud.cceClusterV3.CceClusterV3",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CceClusterV3_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-opentelekomcloud.cceClusterV3.CceClusterV3",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CceClusterV3) AddMoveTarget(moveTarget *string) {
	if err := c.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (c *jsiiProxy_CceClusterV3) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CceClusterV3) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_CceClusterV3) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CceClusterV3) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_CceClusterV3) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_CceClusterV3) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_CceClusterV3) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_CceClusterV3) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_CceClusterV3) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_CceClusterV3) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_CceClusterV3) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CceClusterV3) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := c.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (c *jsiiProxy_CceClusterV3) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CceClusterV3) MoveFromId(id *string) {
	if err := c.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"moveFromId",
		[]interface{}{id},
	)
}

func (c *jsiiProxy_CceClusterV3) MoveTo(moveTarget *string, index interface{}) {
	if err := c.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (c *jsiiProxy_CceClusterV3) MoveToId(id *string) {
	if err := c.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"moveToId",
		[]interface{}{id},
	)
}

func (c *jsiiProxy_CceClusterV3) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CceClusterV3) PutAuthenticatingProxy(value *CceClusterV3AuthenticatingProxy) {
	if err := c.validatePutAuthenticatingProxyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putAuthenticatingProxy",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CceClusterV3) PutMasters(value interface{}) {
	if err := c.validatePutMastersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putMasters",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CceClusterV3) PutTimeouts(value *CceClusterV3Timeouts) {
	if err := c.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CceClusterV3) ResetAnnotations() {
	_jsii_.InvokeVoid(
		c,
		"resetAnnotations",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CceClusterV3) ResetAuthenticatingProxy() {
	_jsii_.InvokeVoid(
		c,
		"resetAuthenticatingProxy",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CceClusterV3) ResetAuthenticatingProxyCa() {
	_jsii_.InvokeVoid(
		c,
		"resetAuthenticatingProxyCa",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CceClusterV3) ResetAuthenticationMode() {
	_jsii_.InvokeVoid(
		c,
		"resetAuthenticationMode",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CceClusterV3) ResetBillingMode() {
	_jsii_.InvokeVoid(
		c,
		"resetBillingMode",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CceClusterV3) ResetClusterVersion() {
	_jsii_.InvokeVoid(
		c,
		"resetClusterVersion",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CceClusterV3) ResetContainerNetworkCidr() {
	_jsii_.InvokeVoid(
		c,
		"resetContainerNetworkCidr",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CceClusterV3) ResetDeleteAllNetwork() {
	_jsii_.InvokeVoid(
		c,
		"resetDeleteAllNetwork",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CceClusterV3) ResetDeleteAllStorage() {
	_jsii_.InvokeVoid(
		c,
		"resetDeleteAllStorage",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CceClusterV3) ResetDeleteEfs() {
	_jsii_.InvokeVoid(
		c,
		"resetDeleteEfs",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CceClusterV3) ResetDeleteEni() {
	_jsii_.InvokeVoid(
		c,
		"resetDeleteEni",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CceClusterV3) ResetDeleteEvs() {
	_jsii_.InvokeVoid(
		c,
		"resetDeleteEvs",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CceClusterV3) ResetDeleteNet() {
	_jsii_.InvokeVoid(
		c,
		"resetDeleteNet",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CceClusterV3) ResetDeleteObs() {
	_jsii_.InvokeVoid(
		c,
		"resetDeleteObs",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CceClusterV3) ResetDeleteSfs() {
	_jsii_.InvokeVoid(
		c,
		"resetDeleteSfs",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CceClusterV3) ResetDescription() {
	_jsii_.InvokeVoid(
		c,
		"resetDescription",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CceClusterV3) ResetEip() {
	_jsii_.InvokeVoid(
		c,
		"resetEip",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CceClusterV3) ResetEnableVolumeEncryption() {
	_jsii_.InvokeVoid(
		c,
		"resetEnableVolumeEncryption",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CceClusterV3) ResetEniSubnetCidr() {
	_jsii_.InvokeVoid(
		c,
		"resetEniSubnetCidr",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CceClusterV3) ResetEniSubnetId() {
	_jsii_.InvokeVoid(
		c,
		"resetEniSubnetId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CceClusterV3) ResetExtendParam() {
	_jsii_.InvokeVoid(
		c,
		"resetExtendParam",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CceClusterV3) ResetHighwaySubnetId() {
	_jsii_.InvokeVoid(
		c,
		"resetHighwaySubnetId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CceClusterV3) ResetId() {
	_jsii_.InvokeVoid(
		c,
		"resetId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CceClusterV3) ResetIgnoreAddons() {
	_jsii_.InvokeVoid(
		c,
		"resetIgnoreAddons",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CceClusterV3) ResetIgnoreCertificateClustersData() {
	_jsii_.InvokeVoid(
		c,
		"resetIgnoreCertificateClustersData",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CceClusterV3) ResetIgnoreCertificateUsersData() {
	_jsii_.InvokeVoid(
		c,
		"resetIgnoreCertificateUsersData",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CceClusterV3) ResetKubeProxyMode() {
	_jsii_.InvokeVoid(
		c,
		"resetKubeProxyMode",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CceClusterV3) ResetKubernetesSvcIpRange() {
	_jsii_.InvokeVoid(
		c,
		"resetKubernetesSvcIpRange",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CceClusterV3) ResetLabels() {
	_jsii_.InvokeVoid(
		c,
		"resetLabels",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CceClusterV3) ResetMasters() {
	_jsii_.InvokeVoid(
		c,
		"resetMasters",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CceClusterV3) ResetMultiAz() {
	_jsii_.InvokeVoid(
		c,
		"resetMultiAz",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CceClusterV3) ResetNoAddons() {
	_jsii_.InvokeVoid(
		c,
		"resetNoAddons",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CceClusterV3) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		c,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CceClusterV3) ResetRegion() {
	_jsii_.InvokeVoid(
		c,
		"resetRegion",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CceClusterV3) ResetSecurityGroupId() {
	_jsii_.InvokeVoid(
		c,
		"resetSecurityGroupId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CceClusterV3) ResetTimeouts() {
	_jsii_.InvokeVoid(
		c,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CceClusterV3) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CceClusterV3) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CceClusterV3) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CceClusterV3) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CceClusterV3) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CceClusterV3) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

