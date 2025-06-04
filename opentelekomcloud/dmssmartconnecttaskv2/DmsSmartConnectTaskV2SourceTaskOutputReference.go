// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dmssmartconnecttaskv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v12/jsii"

	"github.com/cdktf/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v12/dmssmartconnecttaskv2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DmsSmartConnectTaskV2SourceTaskOutputReference interface {
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
	CompressionType() *string
	SetCompressionType(val *string)
	CompressionTypeInput() *string
	ConsumerStrategy() *string
	SetConsumerStrategy(val *string)
	ConsumerStrategyInput() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	CurrentInstanceAlias() *string
	SetCurrentInstanceAlias(val *string)
	CurrentInstanceAliasInput() *string
	Direction() *string
	SetDirection(val *string)
	DirectionInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *DmsSmartConnectTaskV2SourceTask
	SetInternalValue(val *DmsSmartConnectTaskV2SourceTask)
	Password() *string
	SetPassword(val *string)
	PasswordInput() *string
	PeerInstanceAddress() *[]*string
	SetPeerInstanceAddress(val *[]*string)
	PeerInstanceAddressInput() *[]*string
	PeerInstanceAlias() *string
	SetPeerInstanceAlias(val *string)
	PeerInstanceAliasInput() *string
	PeerInstanceId() *string
	SetPeerInstanceId(val *string)
	PeerInstanceIdInput() *string
	ProvenanceHeaderEnabled() interface{}
	SetProvenanceHeaderEnabled(val interface{})
	ProvenanceHeaderEnabledInput() interface{}
	RenameTopicEnabled() interface{}
	SetRenameTopicEnabled(val interface{})
	RenameTopicEnabledInput() interface{}
	ReplicationFactor() *float64
	SetReplicationFactor(val *float64)
	ReplicationFactorInput() *float64
	SaslMechanism() *string
	SetSaslMechanism(val *string)
	SaslMechanismInput() *string
	SecurityProtocol() *string
	SetSecurityProtocol(val *string)
	SecurityProtocolInput() *string
	SyncConsumerOffsetsEnabled() interface{}
	SetSyncConsumerOffsetsEnabled(val interface{})
	SyncConsumerOffsetsEnabledInput() interface{}
	TaskNum() *float64
	SetTaskNum(val *float64)
	TaskNumInput() *float64
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TopicsMapping() *[]*string
	SetTopicsMapping(val *[]*string)
	TopicsMappingInput() *[]*string
	UserName() *string
	SetUserName(val *string)
	UserNameInput() *string
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
	ResetCompressionType()
	ResetConsumerStrategy()
	ResetCurrentInstanceAlias()
	ResetDirection()
	ResetPassword()
	ResetPeerInstanceAddress()
	ResetPeerInstanceAlias()
	ResetPeerInstanceId()
	ResetProvenanceHeaderEnabled()
	ResetRenameTopicEnabled()
	ResetReplicationFactor()
	ResetSaslMechanism()
	ResetSecurityProtocol()
	ResetSyncConsumerOffsetsEnabled()
	ResetTaskNum()
	ResetTopicsMapping()
	ResetUserName()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DmsSmartConnectTaskV2SourceTaskOutputReference
type jsiiProxy_DmsSmartConnectTaskV2SourceTaskOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DmsSmartConnectTaskV2SourceTaskOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsSmartConnectTaskV2SourceTaskOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsSmartConnectTaskV2SourceTaskOutputReference) CompressionType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"compressionType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsSmartConnectTaskV2SourceTaskOutputReference) CompressionTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"compressionTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsSmartConnectTaskV2SourceTaskOutputReference) ConsumerStrategy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"consumerStrategy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsSmartConnectTaskV2SourceTaskOutputReference) ConsumerStrategyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"consumerStrategyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsSmartConnectTaskV2SourceTaskOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsSmartConnectTaskV2SourceTaskOutputReference) CurrentInstanceAlias() *string {
	var returns *string
	_jsii_.Get(
		j,
		"currentInstanceAlias",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsSmartConnectTaskV2SourceTaskOutputReference) CurrentInstanceAliasInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"currentInstanceAliasInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsSmartConnectTaskV2SourceTaskOutputReference) Direction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"direction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsSmartConnectTaskV2SourceTaskOutputReference) DirectionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"directionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsSmartConnectTaskV2SourceTaskOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsSmartConnectTaskV2SourceTaskOutputReference) InternalValue() *DmsSmartConnectTaskV2SourceTask {
	var returns *DmsSmartConnectTaskV2SourceTask
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsSmartConnectTaskV2SourceTaskOutputReference) Password() *string {
	var returns *string
	_jsii_.Get(
		j,
		"password",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsSmartConnectTaskV2SourceTaskOutputReference) PasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"passwordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsSmartConnectTaskV2SourceTaskOutputReference) PeerInstanceAddress() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"peerInstanceAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsSmartConnectTaskV2SourceTaskOutputReference) PeerInstanceAddressInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"peerInstanceAddressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsSmartConnectTaskV2SourceTaskOutputReference) PeerInstanceAlias() *string {
	var returns *string
	_jsii_.Get(
		j,
		"peerInstanceAlias",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsSmartConnectTaskV2SourceTaskOutputReference) PeerInstanceAliasInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"peerInstanceAliasInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsSmartConnectTaskV2SourceTaskOutputReference) PeerInstanceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"peerInstanceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsSmartConnectTaskV2SourceTaskOutputReference) PeerInstanceIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"peerInstanceIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsSmartConnectTaskV2SourceTaskOutputReference) ProvenanceHeaderEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"provenanceHeaderEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsSmartConnectTaskV2SourceTaskOutputReference) ProvenanceHeaderEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"provenanceHeaderEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsSmartConnectTaskV2SourceTaskOutputReference) RenameTopicEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"renameTopicEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsSmartConnectTaskV2SourceTaskOutputReference) RenameTopicEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"renameTopicEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsSmartConnectTaskV2SourceTaskOutputReference) ReplicationFactor() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"replicationFactor",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsSmartConnectTaskV2SourceTaskOutputReference) ReplicationFactorInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"replicationFactorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsSmartConnectTaskV2SourceTaskOutputReference) SaslMechanism() *string {
	var returns *string
	_jsii_.Get(
		j,
		"saslMechanism",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsSmartConnectTaskV2SourceTaskOutputReference) SaslMechanismInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"saslMechanismInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsSmartConnectTaskV2SourceTaskOutputReference) SecurityProtocol() *string {
	var returns *string
	_jsii_.Get(
		j,
		"securityProtocol",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsSmartConnectTaskV2SourceTaskOutputReference) SecurityProtocolInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"securityProtocolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsSmartConnectTaskV2SourceTaskOutputReference) SyncConsumerOffsetsEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"syncConsumerOffsetsEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsSmartConnectTaskV2SourceTaskOutputReference) SyncConsumerOffsetsEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"syncConsumerOffsetsEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsSmartConnectTaskV2SourceTaskOutputReference) TaskNum() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"taskNum",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsSmartConnectTaskV2SourceTaskOutputReference) TaskNumInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"taskNumInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsSmartConnectTaskV2SourceTaskOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsSmartConnectTaskV2SourceTaskOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsSmartConnectTaskV2SourceTaskOutputReference) TopicsMapping() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"topicsMapping",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsSmartConnectTaskV2SourceTaskOutputReference) TopicsMappingInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"topicsMappingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsSmartConnectTaskV2SourceTaskOutputReference) UserName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsSmartConnectTaskV2SourceTaskOutputReference) UserNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userNameInput",
		&returns,
	)
	return returns
}


func NewDmsSmartConnectTaskV2SourceTaskOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DmsSmartConnectTaskV2SourceTaskOutputReference {
	_init_.Initialize()

	if err := validateNewDmsSmartConnectTaskV2SourceTaskOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DmsSmartConnectTaskV2SourceTaskOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.dmsSmartConnectTaskV2.DmsSmartConnectTaskV2SourceTaskOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDmsSmartConnectTaskV2SourceTaskOutputReference_Override(d DmsSmartConnectTaskV2SourceTaskOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.dmsSmartConnectTaskV2.DmsSmartConnectTaskV2SourceTaskOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DmsSmartConnectTaskV2SourceTaskOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DmsSmartConnectTaskV2SourceTaskOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DmsSmartConnectTaskV2SourceTaskOutputReference)SetCompressionType(val *string) {
	if err := j.validateSetCompressionTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"compressionType",
		val,
	)
}

func (j *jsiiProxy_DmsSmartConnectTaskV2SourceTaskOutputReference)SetConsumerStrategy(val *string) {
	if err := j.validateSetConsumerStrategyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"consumerStrategy",
		val,
	)
}

func (j *jsiiProxy_DmsSmartConnectTaskV2SourceTaskOutputReference)SetCurrentInstanceAlias(val *string) {
	if err := j.validateSetCurrentInstanceAliasParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"currentInstanceAlias",
		val,
	)
}

func (j *jsiiProxy_DmsSmartConnectTaskV2SourceTaskOutputReference)SetDirection(val *string) {
	if err := j.validateSetDirectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"direction",
		val,
	)
}

func (j *jsiiProxy_DmsSmartConnectTaskV2SourceTaskOutputReference)SetInternalValue(val *DmsSmartConnectTaskV2SourceTask) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DmsSmartConnectTaskV2SourceTaskOutputReference)SetPassword(val *string) {
	if err := j.validateSetPasswordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"password",
		val,
	)
}

func (j *jsiiProxy_DmsSmartConnectTaskV2SourceTaskOutputReference)SetPeerInstanceAddress(val *[]*string) {
	if err := j.validateSetPeerInstanceAddressParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"peerInstanceAddress",
		val,
	)
}

func (j *jsiiProxy_DmsSmartConnectTaskV2SourceTaskOutputReference)SetPeerInstanceAlias(val *string) {
	if err := j.validateSetPeerInstanceAliasParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"peerInstanceAlias",
		val,
	)
}

func (j *jsiiProxy_DmsSmartConnectTaskV2SourceTaskOutputReference)SetPeerInstanceId(val *string) {
	if err := j.validateSetPeerInstanceIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"peerInstanceId",
		val,
	)
}

func (j *jsiiProxy_DmsSmartConnectTaskV2SourceTaskOutputReference)SetProvenanceHeaderEnabled(val interface{}) {
	if err := j.validateSetProvenanceHeaderEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provenanceHeaderEnabled",
		val,
	)
}

func (j *jsiiProxy_DmsSmartConnectTaskV2SourceTaskOutputReference)SetRenameTopicEnabled(val interface{}) {
	if err := j.validateSetRenameTopicEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"renameTopicEnabled",
		val,
	)
}

func (j *jsiiProxy_DmsSmartConnectTaskV2SourceTaskOutputReference)SetReplicationFactor(val *float64) {
	if err := j.validateSetReplicationFactorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"replicationFactor",
		val,
	)
}

func (j *jsiiProxy_DmsSmartConnectTaskV2SourceTaskOutputReference)SetSaslMechanism(val *string) {
	if err := j.validateSetSaslMechanismParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"saslMechanism",
		val,
	)
}

func (j *jsiiProxy_DmsSmartConnectTaskV2SourceTaskOutputReference)SetSecurityProtocol(val *string) {
	if err := j.validateSetSecurityProtocolParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"securityProtocol",
		val,
	)
}

func (j *jsiiProxy_DmsSmartConnectTaskV2SourceTaskOutputReference)SetSyncConsumerOffsetsEnabled(val interface{}) {
	if err := j.validateSetSyncConsumerOffsetsEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"syncConsumerOffsetsEnabled",
		val,
	)
}

func (j *jsiiProxy_DmsSmartConnectTaskV2SourceTaskOutputReference)SetTaskNum(val *float64) {
	if err := j.validateSetTaskNumParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"taskNum",
		val,
	)
}

func (j *jsiiProxy_DmsSmartConnectTaskV2SourceTaskOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DmsSmartConnectTaskV2SourceTaskOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DmsSmartConnectTaskV2SourceTaskOutputReference)SetTopicsMapping(val *[]*string) {
	if err := j.validateSetTopicsMappingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"topicsMapping",
		val,
	)
}

func (j *jsiiProxy_DmsSmartConnectTaskV2SourceTaskOutputReference)SetUserName(val *string) {
	if err := j.validateSetUserNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"userName",
		val,
	)
}

func (d *jsiiProxy_DmsSmartConnectTaskV2SourceTaskOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DmsSmartConnectTaskV2SourceTaskOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DmsSmartConnectTaskV2SourceTaskOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DmsSmartConnectTaskV2SourceTaskOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DmsSmartConnectTaskV2SourceTaskOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DmsSmartConnectTaskV2SourceTaskOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DmsSmartConnectTaskV2SourceTaskOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DmsSmartConnectTaskV2SourceTaskOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DmsSmartConnectTaskV2SourceTaskOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DmsSmartConnectTaskV2SourceTaskOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DmsSmartConnectTaskV2SourceTaskOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DmsSmartConnectTaskV2SourceTaskOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DmsSmartConnectTaskV2SourceTaskOutputReference) ResetCompressionType() {
	_jsii_.InvokeVoid(
		d,
		"resetCompressionType",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsSmartConnectTaskV2SourceTaskOutputReference) ResetConsumerStrategy() {
	_jsii_.InvokeVoid(
		d,
		"resetConsumerStrategy",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsSmartConnectTaskV2SourceTaskOutputReference) ResetCurrentInstanceAlias() {
	_jsii_.InvokeVoid(
		d,
		"resetCurrentInstanceAlias",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsSmartConnectTaskV2SourceTaskOutputReference) ResetDirection() {
	_jsii_.InvokeVoid(
		d,
		"resetDirection",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsSmartConnectTaskV2SourceTaskOutputReference) ResetPassword() {
	_jsii_.InvokeVoid(
		d,
		"resetPassword",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsSmartConnectTaskV2SourceTaskOutputReference) ResetPeerInstanceAddress() {
	_jsii_.InvokeVoid(
		d,
		"resetPeerInstanceAddress",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsSmartConnectTaskV2SourceTaskOutputReference) ResetPeerInstanceAlias() {
	_jsii_.InvokeVoid(
		d,
		"resetPeerInstanceAlias",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsSmartConnectTaskV2SourceTaskOutputReference) ResetPeerInstanceId() {
	_jsii_.InvokeVoid(
		d,
		"resetPeerInstanceId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsSmartConnectTaskV2SourceTaskOutputReference) ResetProvenanceHeaderEnabled() {
	_jsii_.InvokeVoid(
		d,
		"resetProvenanceHeaderEnabled",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsSmartConnectTaskV2SourceTaskOutputReference) ResetRenameTopicEnabled() {
	_jsii_.InvokeVoid(
		d,
		"resetRenameTopicEnabled",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsSmartConnectTaskV2SourceTaskOutputReference) ResetReplicationFactor() {
	_jsii_.InvokeVoid(
		d,
		"resetReplicationFactor",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsSmartConnectTaskV2SourceTaskOutputReference) ResetSaslMechanism() {
	_jsii_.InvokeVoid(
		d,
		"resetSaslMechanism",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsSmartConnectTaskV2SourceTaskOutputReference) ResetSecurityProtocol() {
	_jsii_.InvokeVoid(
		d,
		"resetSecurityProtocol",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsSmartConnectTaskV2SourceTaskOutputReference) ResetSyncConsumerOffsetsEnabled() {
	_jsii_.InvokeVoid(
		d,
		"resetSyncConsumerOffsetsEnabled",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsSmartConnectTaskV2SourceTaskOutputReference) ResetTaskNum() {
	_jsii_.InvokeVoid(
		d,
		"resetTaskNum",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsSmartConnectTaskV2SourceTaskOutputReference) ResetTopicsMapping() {
	_jsii_.InvokeVoid(
		d,
		"resetTopicsMapping",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsSmartConnectTaskV2SourceTaskOutputReference) ResetUserName() {
	_jsii_.InvokeVoid(
		d,
		"resetUserName",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsSmartConnectTaskV2SourceTaskOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DmsSmartConnectTaskV2SourceTaskOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

