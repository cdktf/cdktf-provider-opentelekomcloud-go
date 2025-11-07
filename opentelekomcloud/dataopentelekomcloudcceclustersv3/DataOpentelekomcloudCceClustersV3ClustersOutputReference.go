// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataopentelekomcloudcceclustersv3

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v12/jsii"

	"github.com/cdktf/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v12/dataopentelekomcloudcceclustersv3/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataOpentelekomcloudCceClustersV3ClustersOutputReference interface {
	cdktf.ComplexObject
	AuthenticationMode() *string
	BillingMode() *float64
	CertificateClusters() DataOpentelekomcloudCceClustersV3ClustersCertificateClustersList
	CertificateUsers() DataOpentelekomcloudCceClustersV3ClustersCertificateUsersList
	ClusterType() *string
	ClusterVersion() *string
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	ContainerNetworkCidr() *string
	ContainerNetworkType() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	Description() *string
	Endpoints() DataOpentelekomcloudCceClustersV3ClustersEndpointsList
	EniSubnetCidr() *string
	EniSubnetId() *string
	EnterpriseProjectId() *string
	FlavorId() *string
	// Experimental.
	Fqn() *string
	Id() *string
	InternalValue() *DataOpentelekomcloudCceClustersV3Clusters
	SetInternalValue(val *DataOpentelekomcloudCceClustersV3Clusters)
	KubeConfigRaw() *string
	Masters() DataOpentelekomcloudCceClustersV3ClustersMastersList
	Name() *string
	SecurityGroupId() *string
	ServiceNetworkCidr() *string
	Status() *string
	SubnetId() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	VpcId() *string
	// Experimental.
	ComputeFqn() *string
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
	InterpolationAsList() cdktf.IResolvable
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataOpentelekomcloudCceClustersV3ClustersOutputReference
type jsiiProxy_DataOpentelekomcloudCceClustersV3ClustersOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataOpentelekomcloudCceClustersV3ClustersOutputReference) AuthenticationMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authenticationMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudCceClustersV3ClustersOutputReference) BillingMode() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"billingMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudCceClustersV3ClustersOutputReference) CertificateClusters() DataOpentelekomcloudCceClustersV3ClustersCertificateClustersList {
	var returns DataOpentelekomcloudCceClustersV3ClustersCertificateClustersList
	_jsii_.Get(
		j,
		"certificateClusters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudCceClustersV3ClustersOutputReference) CertificateUsers() DataOpentelekomcloudCceClustersV3ClustersCertificateUsersList {
	var returns DataOpentelekomcloudCceClustersV3ClustersCertificateUsersList
	_jsii_.Get(
		j,
		"certificateUsers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudCceClustersV3ClustersOutputReference) ClusterType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudCceClustersV3ClustersOutputReference) ClusterVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudCceClustersV3ClustersOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudCceClustersV3ClustersOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudCceClustersV3ClustersOutputReference) ContainerNetworkCidr() *string {
	var returns *string
	_jsii_.Get(
		j,
		"containerNetworkCidr",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudCceClustersV3ClustersOutputReference) ContainerNetworkType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"containerNetworkType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudCceClustersV3ClustersOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudCceClustersV3ClustersOutputReference) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudCceClustersV3ClustersOutputReference) Endpoints() DataOpentelekomcloudCceClustersV3ClustersEndpointsList {
	var returns DataOpentelekomcloudCceClustersV3ClustersEndpointsList
	_jsii_.Get(
		j,
		"endpoints",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudCceClustersV3ClustersOutputReference) EniSubnetCidr() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eniSubnetCidr",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudCceClustersV3ClustersOutputReference) EniSubnetId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eniSubnetId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudCceClustersV3ClustersOutputReference) EnterpriseProjectId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"enterpriseProjectId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudCceClustersV3ClustersOutputReference) FlavorId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"flavorId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudCceClustersV3ClustersOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudCceClustersV3ClustersOutputReference) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudCceClustersV3ClustersOutputReference) InternalValue() *DataOpentelekomcloudCceClustersV3Clusters {
	var returns *DataOpentelekomcloudCceClustersV3Clusters
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudCceClustersV3ClustersOutputReference) KubeConfigRaw() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kubeConfigRaw",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudCceClustersV3ClustersOutputReference) Masters() DataOpentelekomcloudCceClustersV3ClustersMastersList {
	var returns DataOpentelekomcloudCceClustersV3ClustersMastersList
	_jsii_.Get(
		j,
		"masters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudCceClustersV3ClustersOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudCceClustersV3ClustersOutputReference) SecurityGroupId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"securityGroupId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudCceClustersV3ClustersOutputReference) ServiceNetworkCidr() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceNetworkCidr",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudCceClustersV3ClustersOutputReference) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudCceClustersV3ClustersOutputReference) SubnetId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudCceClustersV3ClustersOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudCceClustersV3ClustersOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudCceClustersV3ClustersOutputReference) VpcId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcId",
		&returns,
	)
	return returns
}


func NewDataOpentelekomcloudCceClustersV3ClustersOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataOpentelekomcloudCceClustersV3ClustersOutputReference {
	_init_.Initialize()

	if err := validateNewDataOpentelekomcloudCceClustersV3ClustersOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataOpentelekomcloudCceClustersV3ClustersOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.dataOpentelekomcloudCceClustersV3.DataOpentelekomcloudCceClustersV3ClustersOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataOpentelekomcloudCceClustersV3ClustersOutputReference_Override(d DataOpentelekomcloudCceClustersV3ClustersOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.dataOpentelekomcloudCceClustersV3.DataOpentelekomcloudCceClustersV3ClustersOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataOpentelekomcloudCceClustersV3ClustersOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataOpentelekomcloudCceClustersV3ClustersOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataOpentelekomcloudCceClustersV3ClustersOutputReference)SetInternalValue(val *DataOpentelekomcloudCceClustersV3Clusters) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataOpentelekomcloudCceClustersV3ClustersOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataOpentelekomcloudCceClustersV3ClustersOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataOpentelekomcloudCceClustersV3ClustersOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataOpentelekomcloudCceClustersV3ClustersOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataOpentelekomcloudCceClustersV3ClustersOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataOpentelekomcloudCceClustersV3ClustersOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataOpentelekomcloudCceClustersV3ClustersOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataOpentelekomcloudCceClustersV3ClustersOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataOpentelekomcloudCceClustersV3ClustersOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataOpentelekomcloudCceClustersV3ClustersOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataOpentelekomcloudCceClustersV3ClustersOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataOpentelekomcloudCceClustersV3ClustersOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataOpentelekomcloudCceClustersV3ClustersOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataOpentelekomcloudCceClustersV3ClustersOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataOpentelekomcloudCceClustersV3ClustersOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
	if err := d.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataOpentelekomcloudCceClustersV3ClustersOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

