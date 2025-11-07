// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataopentelekomcloudasmservicemeshv1

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v12/jsii"

	"github.com/cdktf/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v12/dataopentelekomcloudasmservicemeshv1/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataOpentelekomcloudAsmServiceMeshV1ServiceMeshesTelemetryConfigTracingExtensionProvidersList interface {
	cdktf.ComplexList
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	// The attribute on the parent resource this class is referencing.
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	// The parent resource.
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// whether the list is wrapping a set (will add tolist() to be able to access an item via an index).
	WrapsSet() *bool
	SetWrapsSet(val *bool)
	// Creating an iterator for this complex list.
	//
	// The list will be converted into a map with the mapKeyAttributeName as the key.
	// Experimental.
	AllWithMapKey(mapKeyAttributeName *string) cdktf.DynamicListTerraformIterator
	// Experimental.
	ComputeFqn() *string
	Get(index *float64) DataOpentelekomcloudAsmServiceMeshV1ServiceMeshesTelemetryConfigTracingExtensionProvidersOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataOpentelekomcloudAsmServiceMeshV1ServiceMeshesTelemetryConfigTracingExtensionProvidersList
type jsiiProxy_DataOpentelekomcloudAsmServiceMeshV1ServiceMeshesTelemetryConfigTracingExtensionProvidersList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_DataOpentelekomcloudAsmServiceMeshV1ServiceMeshesTelemetryConfigTracingExtensionProvidersList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudAsmServiceMeshV1ServiceMeshesTelemetryConfigTracingExtensionProvidersList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudAsmServiceMeshV1ServiceMeshesTelemetryConfigTracingExtensionProvidersList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudAsmServiceMeshV1ServiceMeshesTelemetryConfigTracingExtensionProvidersList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudAsmServiceMeshV1ServiceMeshesTelemetryConfigTracingExtensionProvidersList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewDataOpentelekomcloudAsmServiceMeshV1ServiceMeshesTelemetryConfigTracingExtensionProvidersList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) DataOpentelekomcloudAsmServiceMeshV1ServiceMeshesTelemetryConfigTracingExtensionProvidersList {
	_init_.Initialize()

	if err := validateNewDataOpentelekomcloudAsmServiceMeshV1ServiceMeshesTelemetryConfigTracingExtensionProvidersListParameters(terraformResource, terraformAttribute, wrapsSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataOpentelekomcloudAsmServiceMeshV1ServiceMeshesTelemetryConfigTracingExtensionProvidersList{}

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.dataOpentelekomcloudAsmServiceMeshV1.DataOpentelekomcloudAsmServiceMeshV1ServiceMeshesTelemetryConfigTracingExtensionProvidersList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewDataOpentelekomcloudAsmServiceMeshV1ServiceMeshesTelemetryConfigTracingExtensionProvidersList_Override(d DataOpentelekomcloudAsmServiceMeshV1ServiceMeshesTelemetryConfigTracingExtensionProvidersList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.dataOpentelekomcloudAsmServiceMeshV1.DataOpentelekomcloudAsmServiceMeshV1ServiceMeshesTelemetryConfigTracingExtensionProvidersList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		d,
	)
}

func (j *jsiiProxy_DataOpentelekomcloudAsmServiceMeshV1ServiceMeshesTelemetryConfigTracingExtensionProvidersList)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataOpentelekomcloudAsmServiceMeshV1ServiceMeshesTelemetryConfigTracingExtensionProvidersList)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DataOpentelekomcloudAsmServiceMeshV1ServiceMeshesTelemetryConfigTracingExtensionProvidersList)SetWrapsSet(val *bool) {
	if err := j.validateSetWrapsSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (d *jsiiProxy_DataOpentelekomcloudAsmServiceMeshV1ServiceMeshesTelemetryConfigTracingExtensionProvidersList) AllWithMapKey(mapKeyAttributeName *string) cdktf.DynamicListTerraformIterator {
	if err := d.validateAllWithMapKeyParameters(mapKeyAttributeName); err != nil {
		panic(err)
	}
	var returns cdktf.DynamicListTerraformIterator

	_jsii_.Invoke(
		d,
		"allWithMapKey",
		[]interface{}{mapKeyAttributeName},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataOpentelekomcloudAsmServiceMeshV1ServiceMeshesTelemetryConfigTracingExtensionProvidersList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataOpentelekomcloudAsmServiceMeshV1ServiceMeshesTelemetryConfigTracingExtensionProvidersList) Get(index *float64) DataOpentelekomcloudAsmServiceMeshV1ServiceMeshesTelemetryConfigTracingExtensionProvidersOutputReference {
	if err := d.validateGetParameters(index); err != nil {
		panic(err)
	}
	var returns DataOpentelekomcloudAsmServiceMeshV1ServiceMeshesTelemetryConfigTracingExtensionProvidersOutputReference

	_jsii_.Invoke(
		d,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataOpentelekomcloudAsmServiceMeshV1ServiceMeshesTelemetryConfigTracingExtensionProvidersList) Resolve(context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataOpentelekomcloudAsmServiceMeshV1ServiceMeshesTelemetryConfigTracingExtensionProvidersList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

