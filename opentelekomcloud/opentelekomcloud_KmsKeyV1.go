// Prebuilt opentelekomcloud Provider for Terraform CDK (cdktf)
package opentelekomcloud

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v2/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://www.terraform.io/docs/providers/opentelekomcloud/r/kms_key_v1 opentelekomcloud_kms_key_v1}.
type KmsKeyV1 interface {
	cdktf.TerraformResource
	AllowCancelDeletion() interface{}
	SetAllowCancelDeletion(val interface{})
	AllowCancelDeletionInput() interface{}
	// Experimental.
	CdktfStack() cdktf.TerraformStack
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
	CreationDate() *string
	DefaultKeyFlag() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	DomainId() *string
	ExpirationTime() *string
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
	IsEnabled() interface{}
	SetIsEnabled(val interface{})
	IsEnabledInput() interface{}
	KeyAlias() *string
	SetKeyAlias(val *string)
	KeyAliasInput() *string
	KeyDescription() *string
	SetKeyDescription(val *string)
	KeyDescriptionInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	// The tree node.
	Node() constructs.Node
	Origin() *string
	PendingDays() *string
	SetPendingDays(val *string)
	PendingDaysInput() *string
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
	Realm() *string
	SetRealm(val *string)
	RealmInput() *string
	RotationEnabled() interface{}
	SetRotationEnabled(val interface{})
	RotationEnabledInput() interface{}
	RotationInterval() *float64
	SetRotationInterval(val *float64)
	RotationIntervalInput() *float64
	RotationNumber() *float64
	ScheduledDeletionDate() *string
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsInput() *map[string]*string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
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
	ResetAllowCancelDeletion()
	ResetId()
	ResetIsEnabled()
	ResetKeyDescription()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPendingDays()
	ResetRealm()
	ResetRotationEnabled()
	ResetRotationInterval()
	ResetTags()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for KmsKeyV1
type jsiiProxy_KmsKeyV1 struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_KmsKeyV1) AllowCancelDeletion() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowCancelDeletion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KmsKeyV1) AllowCancelDeletionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowCancelDeletionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KmsKeyV1) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KmsKeyV1) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KmsKeyV1) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KmsKeyV1) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KmsKeyV1) CreationDate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"creationDate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KmsKeyV1) DefaultKeyFlag() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultKeyFlag",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KmsKeyV1) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KmsKeyV1) DomainId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KmsKeyV1) ExpirationTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"expirationTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KmsKeyV1) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KmsKeyV1) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KmsKeyV1) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KmsKeyV1) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KmsKeyV1) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KmsKeyV1) IsEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KmsKeyV1) IsEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KmsKeyV1) KeyAlias() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyAlias",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KmsKeyV1) KeyAliasInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyAliasInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KmsKeyV1) KeyDescription() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyDescription",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KmsKeyV1) KeyDescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyDescriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KmsKeyV1) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KmsKeyV1) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KmsKeyV1) Origin() *string {
	var returns *string
	_jsii_.Get(
		j,
		"origin",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KmsKeyV1) PendingDays() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pendingDays",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KmsKeyV1) PendingDaysInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pendingDaysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KmsKeyV1) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KmsKeyV1) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KmsKeyV1) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KmsKeyV1) Realm() *string {
	var returns *string
	_jsii_.Get(
		j,
		"realm",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KmsKeyV1) RealmInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"realmInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KmsKeyV1) RotationEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rotationEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KmsKeyV1) RotationEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rotationEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KmsKeyV1) RotationInterval() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"rotationInterval",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KmsKeyV1) RotationIntervalInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"rotationIntervalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KmsKeyV1) RotationNumber() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"rotationNumber",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KmsKeyV1) ScheduledDeletionDate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scheduledDeletionDate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KmsKeyV1) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KmsKeyV1) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KmsKeyV1) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KmsKeyV1) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KmsKeyV1) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}


// Create a new {@link https://www.terraform.io/docs/providers/opentelekomcloud/r/kms_key_v1 opentelekomcloud_kms_key_v1} Resource.
func NewKmsKeyV1(scope constructs.Construct, id *string, config *KmsKeyV1Config) KmsKeyV1 {
	_init_.Initialize()

	j := jsiiProxy_KmsKeyV1{}

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.KmsKeyV1",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/opentelekomcloud/r/kms_key_v1 opentelekomcloud_kms_key_v1} Resource.
func NewKmsKeyV1_Override(k KmsKeyV1, scope constructs.Construct, id *string, config *KmsKeyV1Config) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.KmsKeyV1",
		[]interface{}{scope, id, config},
		k,
	)
}

func (j *jsiiProxy_KmsKeyV1) SetAllowCancelDeletion(val interface{}) {
	_jsii_.Set(
		j,
		"allowCancelDeletion",
		val,
	)
}

func (j *jsiiProxy_KmsKeyV1) SetConnection(val interface{}) {
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_KmsKeyV1) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_KmsKeyV1) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_KmsKeyV1) SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_KmsKeyV1) SetId(val *string) {
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_KmsKeyV1) SetIsEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"isEnabled",
		val,
	)
}

func (j *jsiiProxy_KmsKeyV1) SetKeyAlias(val *string) {
	_jsii_.Set(
		j,
		"keyAlias",
		val,
	)
}

func (j *jsiiProxy_KmsKeyV1) SetKeyDescription(val *string) {
	_jsii_.Set(
		j,
		"keyDescription",
		val,
	)
}

func (j *jsiiProxy_KmsKeyV1) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_KmsKeyV1) SetPendingDays(val *string) {
	_jsii_.Set(
		j,
		"pendingDays",
		val,
	)
}

func (j *jsiiProxy_KmsKeyV1) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_KmsKeyV1) SetProvisioners(val *[]interface{}) {
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_KmsKeyV1) SetRealm(val *string) {
	_jsii_.Set(
		j,
		"realm",
		val,
	)
}

func (j *jsiiProxy_KmsKeyV1) SetRotationEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"rotationEnabled",
		val,
	)
}

func (j *jsiiProxy_KmsKeyV1) SetRotationInterval(val *float64) {
	_jsii_.Set(
		j,
		"rotationInterval",
		val,
	)
}

func (j *jsiiProxy_KmsKeyV1) SetTags(val *map[string]*string) {
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
func KmsKeyV1_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-opentelekomcloud.KmsKeyV1",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func KmsKeyV1_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-opentelekomcloud.KmsKeyV1",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (k *jsiiProxy_KmsKeyV1) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		k,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (k *jsiiProxy_KmsKeyV1) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		k,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KmsKeyV1) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KmsKeyV1) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		k,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KmsKeyV1) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		k,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KmsKeyV1) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		k,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KmsKeyV1) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		k,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KmsKeyV1) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		k,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KmsKeyV1) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KmsKeyV1) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		k,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KmsKeyV1) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KmsKeyV1) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		k,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (k *jsiiProxy_KmsKeyV1) ResetAllowCancelDeletion() {
	_jsii_.InvokeVoid(
		k,
		"resetAllowCancelDeletion",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KmsKeyV1) ResetId() {
	_jsii_.InvokeVoid(
		k,
		"resetId",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KmsKeyV1) ResetIsEnabled() {
	_jsii_.InvokeVoid(
		k,
		"resetIsEnabled",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KmsKeyV1) ResetKeyDescription() {
	_jsii_.InvokeVoid(
		k,
		"resetKeyDescription",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KmsKeyV1) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		k,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KmsKeyV1) ResetPendingDays() {
	_jsii_.InvokeVoid(
		k,
		"resetPendingDays",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KmsKeyV1) ResetRealm() {
	_jsii_.InvokeVoid(
		k,
		"resetRealm",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KmsKeyV1) ResetRotationEnabled() {
	_jsii_.InvokeVoid(
		k,
		"resetRotationEnabled",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KmsKeyV1) ResetRotationInterval() {
	_jsii_.InvokeVoid(
		k,
		"resetRotationInterval",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KmsKeyV1) ResetTags() {
	_jsii_.InvokeVoid(
		k,
		"resetTags",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KmsKeyV1) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		k,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KmsKeyV1) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		k,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KmsKeyV1) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KmsKeyV1) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		k,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}
