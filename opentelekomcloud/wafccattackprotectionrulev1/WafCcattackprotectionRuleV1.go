package wafccattackprotectionrulev1

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v6/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v6/wafccattackprotectionrulev1/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.34.1/docs/resources/waf_ccattackprotection_rule_v1 opentelekomcloud_waf_ccattackprotection_rule_v1}.
type WafCcattackprotectionRuleV1 interface {
	cdktf.TerraformResource
	ActionCategory() *string
	SetActionCategory(val *string)
	ActionCategoryInput() *string
	BlockContent() *string
	SetBlockContent(val *string)
	BlockContentInput() *string
	BlockContentType() *string
	SetBlockContentType(val *string)
	BlockContentTypeInput() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	Default() cdktf.IResolvable
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
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
	LimitNum() *float64
	SetLimitNum(val *float64)
	LimitNumInput() *float64
	LimitPeriod() *float64
	SetLimitPeriod(val *float64)
	LimitPeriodInput() *float64
	LockTime() *float64
	SetLockTime(val *float64)
	LockTimeInput() *float64
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
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	// Experimental.
	RawOverrides() interface{}
	TagCategory() *string
	SetTagCategory(val *string)
	TagCategoryInput() *string
	TagContents() *[]*string
	SetTagContents(val *[]*string)
	TagContentsInput() *[]*string
	TagIndex() *string
	SetTagIndex(val *string)
	TagIndexInput() *string
	TagType() *string
	SetTagType(val *string)
	TagTypeInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() WafCcattackprotectionRuleV1TimeoutsOutputReference
	TimeoutsInput() interface{}
	Url() *string
	SetUrl(val *string)
	UrlInput() *string
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
	PutTimeouts(value *WafCcattackprotectionRuleV1Timeouts)
	ResetBlockContent()
	ResetBlockContentType()
	ResetId()
	ResetLockTime()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetTagCategory()
	ResetTagContents()
	ResetTagIndex()
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

// The jsii proxy struct for WafCcattackprotectionRuleV1
type jsiiProxy_WafCcattackprotectionRuleV1 struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_WafCcattackprotectionRuleV1) ActionCategory() *string {
	var returns *string
	_jsii_.Get(
		j,
		"actionCategory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafCcattackprotectionRuleV1) ActionCategoryInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"actionCategoryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafCcattackprotectionRuleV1) BlockContent() *string {
	var returns *string
	_jsii_.Get(
		j,
		"blockContent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafCcattackprotectionRuleV1) BlockContentInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"blockContentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafCcattackprotectionRuleV1) BlockContentType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"blockContentType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafCcattackprotectionRuleV1) BlockContentTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"blockContentTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafCcattackprotectionRuleV1) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafCcattackprotectionRuleV1) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafCcattackprotectionRuleV1) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafCcattackprotectionRuleV1) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafCcattackprotectionRuleV1) Default() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"default",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafCcattackprotectionRuleV1) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafCcattackprotectionRuleV1) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafCcattackprotectionRuleV1) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafCcattackprotectionRuleV1) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafCcattackprotectionRuleV1) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafCcattackprotectionRuleV1) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafCcattackprotectionRuleV1) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafCcattackprotectionRuleV1) LimitNum() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"limitNum",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafCcattackprotectionRuleV1) LimitNumInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"limitNumInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafCcattackprotectionRuleV1) LimitPeriod() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"limitPeriod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafCcattackprotectionRuleV1) LimitPeriodInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"limitPeriodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafCcattackprotectionRuleV1) LockTime() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"lockTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafCcattackprotectionRuleV1) LockTimeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"lockTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafCcattackprotectionRuleV1) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafCcattackprotectionRuleV1) PolicyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafCcattackprotectionRuleV1) PolicyIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policyIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafCcattackprotectionRuleV1) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafCcattackprotectionRuleV1) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafCcattackprotectionRuleV1) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafCcattackprotectionRuleV1) TagCategory() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tagCategory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafCcattackprotectionRuleV1) TagCategoryInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tagCategoryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafCcattackprotectionRuleV1) TagContents() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tagContents",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafCcattackprotectionRuleV1) TagContentsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tagContentsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafCcattackprotectionRuleV1) TagIndex() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tagIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafCcattackprotectionRuleV1) TagIndexInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tagIndexInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafCcattackprotectionRuleV1) TagType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tagType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafCcattackprotectionRuleV1) TagTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tagTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafCcattackprotectionRuleV1) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafCcattackprotectionRuleV1) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafCcattackprotectionRuleV1) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafCcattackprotectionRuleV1) Timeouts() WafCcattackprotectionRuleV1TimeoutsOutputReference {
	var returns WafCcattackprotectionRuleV1TimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafCcattackprotectionRuleV1) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafCcattackprotectionRuleV1) Url() *string {
	var returns *string
	_jsii_.Get(
		j,
		"url",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafCcattackprotectionRuleV1) UrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"urlInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.34.1/docs/resources/waf_ccattackprotection_rule_v1 opentelekomcloud_waf_ccattackprotection_rule_v1} Resource.
func NewWafCcattackprotectionRuleV1(scope constructs.Construct, id *string, config *WafCcattackprotectionRuleV1Config) WafCcattackprotectionRuleV1 {
	_init_.Initialize()

	if err := validateNewWafCcattackprotectionRuleV1Parameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_WafCcattackprotectionRuleV1{}

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.wafCcattackprotectionRuleV1.WafCcattackprotectionRuleV1",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.34.1/docs/resources/waf_ccattackprotection_rule_v1 opentelekomcloud_waf_ccattackprotection_rule_v1} Resource.
func NewWafCcattackprotectionRuleV1_Override(w WafCcattackprotectionRuleV1, scope constructs.Construct, id *string, config *WafCcattackprotectionRuleV1Config) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.wafCcattackprotectionRuleV1.WafCcattackprotectionRuleV1",
		[]interface{}{scope, id, config},
		w,
	)
}

func (j *jsiiProxy_WafCcattackprotectionRuleV1)SetActionCategory(val *string) {
	if err := j.validateSetActionCategoryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"actionCategory",
		val,
	)
}

func (j *jsiiProxy_WafCcattackprotectionRuleV1)SetBlockContent(val *string) {
	if err := j.validateSetBlockContentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"blockContent",
		val,
	)
}

func (j *jsiiProxy_WafCcattackprotectionRuleV1)SetBlockContentType(val *string) {
	if err := j.validateSetBlockContentTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"blockContentType",
		val,
	)
}

func (j *jsiiProxy_WafCcattackprotectionRuleV1)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_WafCcattackprotectionRuleV1)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_WafCcattackprotectionRuleV1)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_WafCcattackprotectionRuleV1)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_WafCcattackprotectionRuleV1)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_WafCcattackprotectionRuleV1)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_WafCcattackprotectionRuleV1)SetLimitNum(val *float64) {
	if err := j.validateSetLimitNumParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"limitNum",
		val,
	)
}

func (j *jsiiProxy_WafCcattackprotectionRuleV1)SetLimitPeriod(val *float64) {
	if err := j.validateSetLimitPeriodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"limitPeriod",
		val,
	)
}

func (j *jsiiProxy_WafCcattackprotectionRuleV1)SetLockTime(val *float64) {
	if err := j.validateSetLockTimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lockTime",
		val,
	)
}

func (j *jsiiProxy_WafCcattackprotectionRuleV1)SetPolicyId(val *string) {
	if err := j.validateSetPolicyIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"policyId",
		val,
	)
}

func (j *jsiiProxy_WafCcattackprotectionRuleV1)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_WafCcattackprotectionRuleV1)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_WafCcattackprotectionRuleV1)SetTagCategory(val *string) {
	if err := j.validateSetTagCategoryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagCategory",
		val,
	)
}

func (j *jsiiProxy_WafCcattackprotectionRuleV1)SetTagContents(val *[]*string) {
	if err := j.validateSetTagContentsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagContents",
		val,
	)
}

func (j *jsiiProxy_WafCcattackprotectionRuleV1)SetTagIndex(val *string) {
	if err := j.validateSetTagIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagIndex",
		val,
	)
}

func (j *jsiiProxy_WafCcattackprotectionRuleV1)SetTagType(val *string) {
	if err := j.validateSetTagTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagType",
		val,
	)
}

func (j *jsiiProxy_WafCcattackprotectionRuleV1)SetUrl(val *string) {
	if err := j.validateSetUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"url",
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
func WafCcattackprotectionRuleV1_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateWafCcattackprotectionRuleV1_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-opentelekomcloud.wafCcattackprotectionRuleV1.WafCcattackprotectionRuleV1",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func WafCcattackprotectionRuleV1_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateWafCcattackprotectionRuleV1_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-opentelekomcloud.wafCcattackprotectionRuleV1.WafCcattackprotectionRuleV1",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func WafCcattackprotectionRuleV1_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateWafCcattackprotectionRuleV1_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-opentelekomcloud.wafCcattackprotectionRuleV1.WafCcattackprotectionRuleV1",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func WafCcattackprotectionRuleV1_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-opentelekomcloud.wafCcattackprotectionRuleV1.WafCcattackprotectionRuleV1",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (w *jsiiProxy_WafCcattackprotectionRuleV1) AddOverride(path *string, value interface{}) {
	if err := w.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (w *jsiiProxy_WafCcattackprotectionRuleV1) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := w.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		w,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WafCcattackprotectionRuleV1) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := w.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WafCcattackprotectionRuleV1) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := w.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		w,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WafCcattackprotectionRuleV1) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := w.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		w,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WafCcattackprotectionRuleV1) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := w.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		w,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WafCcattackprotectionRuleV1) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := w.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		w,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WafCcattackprotectionRuleV1) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := w.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		w,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WafCcattackprotectionRuleV1) GetStringAttribute(terraformAttribute *string) *string {
	if err := w.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		w,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WafCcattackprotectionRuleV1) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := w.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		w,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WafCcattackprotectionRuleV1) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := w.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WafCcattackprotectionRuleV1) OverrideLogicalId(newLogicalId *string) {
	if err := w.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (w *jsiiProxy_WafCcattackprotectionRuleV1) PutTimeouts(value *WafCcattackprotectionRuleV1Timeouts) {
	if err := w.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WafCcattackprotectionRuleV1) ResetBlockContent() {
	_jsii_.InvokeVoid(
		w,
		"resetBlockContent",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WafCcattackprotectionRuleV1) ResetBlockContentType() {
	_jsii_.InvokeVoid(
		w,
		"resetBlockContentType",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WafCcattackprotectionRuleV1) ResetId() {
	_jsii_.InvokeVoid(
		w,
		"resetId",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WafCcattackprotectionRuleV1) ResetLockTime() {
	_jsii_.InvokeVoid(
		w,
		"resetLockTime",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WafCcattackprotectionRuleV1) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		w,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WafCcattackprotectionRuleV1) ResetTagCategory() {
	_jsii_.InvokeVoid(
		w,
		"resetTagCategory",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WafCcattackprotectionRuleV1) ResetTagContents() {
	_jsii_.InvokeVoid(
		w,
		"resetTagContents",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WafCcattackprotectionRuleV1) ResetTagIndex() {
	_jsii_.InvokeVoid(
		w,
		"resetTagIndex",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WafCcattackprotectionRuleV1) ResetTimeouts() {
	_jsii_.InvokeVoid(
		w,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WafCcattackprotectionRuleV1) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		w,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WafCcattackprotectionRuleV1) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		w,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WafCcattackprotectionRuleV1) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WafCcattackprotectionRuleV1) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		w,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

