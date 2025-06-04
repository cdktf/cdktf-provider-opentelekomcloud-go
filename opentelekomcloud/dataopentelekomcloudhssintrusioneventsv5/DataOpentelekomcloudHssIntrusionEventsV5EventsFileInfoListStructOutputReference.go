// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataopentelekomcloudhssintrusioneventsv5

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v12/jsii"

	"github.com/cdktf/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v12/dataopentelekomcloudhssintrusioneventsv5/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataOpentelekomcloudHssIntrusionEventsV5EventsFileInfoListStructOutputReference interface {
	cdktf.ComplexObject
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
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	FdCount() *string
	FdInfo() *string
	FileAction() *string
	FileAlias() *string
	FileAtime() *string
	FileAttr() *string
	FileChangeAttr() *string
	FileContent() *string
	FileCtime() *string
	FileDesc() *string
	FileHash() *string
	FileKeyWord() *string
	FileMd5() *string
	FileMtime() *string
	FileNewPath() *string
	FileOperation() *string
	FilePath() *string
	FileSha256() *string
	FileSize() *float64
	FileType() *string
	// Experimental.
	Fqn() *string
	InternalValue() *DataOpentelekomcloudHssIntrusionEventsV5EventsFileInfoListStruct
	SetInternalValue(val *DataOpentelekomcloudHssIntrusionEventsV5EventsFileInfoListStruct)
	IsDir() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
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
	InterpolationForAttribute(property *string) cdktf.IResolvable
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataOpentelekomcloudHssIntrusionEventsV5EventsFileInfoListStructOutputReference
type jsiiProxy_DataOpentelekomcloudHssIntrusionEventsV5EventsFileInfoListStructOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataOpentelekomcloudHssIntrusionEventsV5EventsFileInfoListStructOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudHssIntrusionEventsV5EventsFileInfoListStructOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudHssIntrusionEventsV5EventsFileInfoListStructOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudHssIntrusionEventsV5EventsFileInfoListStructOutputReference) FdCount() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fdCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudHssIntrusionEventsV5EventsFileInfoListStructOutputReference) FdInfo() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fdInfo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudHssIntrusionEventsV5EventsFileInfoListStructOutputReference) FileAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fileAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudHssIntrusionEventsV5EventsFileInfoListStructOutputReference) FileAlias() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fileAlias",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudHssIntrusionEventsV5EventsFileInfoListStructOutputReference) FileAtime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fileAtime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudHssIntrusionEventsV5EventsFileInfoListStructOutputReference) FileAttr() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fileAttr",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudHssIntrusionEventsV5EventsFileInfoListStructOutputReference) FileChangeAttr() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fileChangeAttr",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudHssIntrusionEventsV5EventsFileInfoListStructOutputReference) FileContent() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fileContent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudHssIntrusionEventsV5EventsFileInfoListStructOutputReference) FileCtime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fileCtime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudHssIntrusionEventsV5EventsFileInfoListStructOutputReference) FileDesc() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fileDesc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudHssIntrusionEventsV5EventsFileInfoListStructOutputReference) FileHash() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fileHash",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudHssIntrusionEventsV5EventsFileInfoListStructOutputReference) FileKeyWord() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fileKeyWord",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudHssIntrusionEventsV5EventsFileInfoListStructOutputReference) FileMd5() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fileMd5",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudHssIntrusionEventsV5EventsFileInfoListStructOutputReference) FileMtime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fileMtime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudHssIntrusionEventsV5EventsFileInfoListStructOutputReference) FileNewPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fileNewPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudHssIntrusionEventsV5EventsFileInfoListStructOutputReference) FileOperation() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fileOperation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudHssIntrusionEventsV5EventsFileInfoListStructOutputReference) FilePath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"filePath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudHssIntrusionEventsV5EventsFileInfoListStructOutputReference) FileSha256() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fileSha256",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudHssIntrusionEventsV5EventsFileInfoListStructOutputReference) FileSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"fileSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudHssIntrusionEventsV5EventsFileInfoListStructOutputReference) FileType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fileType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudHssIntrusionEventsV5EventsFileInfoListStructOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudHssIntrusionEventsV5EventsFileInfoListStructOutputReference) InternalValue() *DataOpentelekomcloudHssIntrusionEventsV5EventsFileInfoListStruct {
	var returns *DataOpentelekomcloudHssIntrusionEventsV5EventsFileInfoListStruct
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudHssIntrusionEventsV5EventsFileInfoListStructOutputReference) IsDir() *string {
	var returns *string
	_jsii_.Get(
		j,
		"isDir",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudHssIntrusionEventsV5EventsFileInfoListStructOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpentelekomcloudHssIntrusionEventsV5EventsFileInfoListStructOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataOpentelekomcloudHssIntrusionEventsV5EventsFileInfoListStructOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataOpentelekomcloudHssIntrusionEventsV5EventsFileInfoListStructOutputReference {
	_init_.Initialize()

	if err := validateNewDataOpentelekomcloudHssIntrusionEventsV5EventsFileInfoListStructOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataOpentelekomcloudHssIntrusionEventsV5EventsFileInfoListStructOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.dataOpentelekomcloudHssIntrusionEventsV5.DataOpentelekomcloudHssIntrusionEventsV5EventsFileInfoListStructOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataOpentelekomcloudHssIntrusionEventsV5EventsFileInfoListStructOutputReference_Override(d DataOpentelekomcloudHssIntrusionEventsV5EventsFileInfoListStructOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.dataOpentelekomcloudHssIntrusionEventsV5.DataOpentelekomcloudHssIntrusionEventsV5EventsFileInfoListStructOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataOpentelekomcloudHssIntrusionEventsV5EventsFileInfoListStructOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataOpentelekomcloudHssIntrusionEventsV5EventsFileInfoListStructOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataOpentelekomcloudHssIntrusionEventsV5EventsFileInfoListStructOutputReference)SetInternalValue(val *DataOpentelekomcloudHssIntrusionEventsV5EventsFileInfoListStruct) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataOpentelekomcloudHssIntrusionEventsV5EventsFileInfoListStructOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataOpentelekomcloudHssIntrusionEventsV5EventsFileInfoListStructOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataOpentelekomcloudHssIntrusionEventsV5EventsFileInfoListStructOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataOpentelekomcloudHssIntrusionEventsV5EventsFileInfoListStructOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataOpentelekomcloudHssIntrusionEventsV5EventsFileInfoListStructOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataOpentelekomcloudHssIntrusionEventsV5EventsFileInfoListStructOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataOpentelekomcloudHssIntrusionEventsV5EventsFileInfoListStructOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataOpentelekomcloudHssIntrusionEventsV5EventsFileInfoListStructOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataOpentelekomcloudHssIntrusionEventsV5EventsFileInfoListStructOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataOpentelekomcloudHssIntrusionEventsV5EventsFileInfoListStructOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataOpentelekomcloudHssIntrusionEventsV5EventsFileInfoListStructOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataOpentelekomcloudHssIntrusionEventsV5EventsFileInfoListStructOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataOpentelekomcloudHssIntrusionEventsV5EventsFileInfoListStructOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataOpentelekomcloudHssIntrusionEventsV5EventsFileInfoListStructOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataOpentelekomcloudHssIntrusionEventsV5EventsFileInfoListStructOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := d.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataOpentelekomcloudHssIntrusionEventsV5EventsFileInfoListStructOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

