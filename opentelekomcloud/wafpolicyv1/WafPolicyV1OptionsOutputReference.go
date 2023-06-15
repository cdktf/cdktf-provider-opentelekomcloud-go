package wafpolicyv1

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v7/jsii"

	"github.com/cdktf/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v7/wafpolicyv1/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type WafPolicyV1OptionsOutputReference interface {
	cdktf.ComplexObject
	Antitamper() interface{}
	SetAntitamper(val interface{})
	AntitamperInput() interface{}
	Cc() interface{}
	SetCc(val interface{})
	CcInput() interface{}
	Common() interface{}
	SetCommon(val interface{})
	CommonInput() interface{}
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
	Crawler() interface{}
	SetCrawler(val interface{})
	CrawlerEngine() interface{}
	SetCrawlerEngine(val interface{})
	CrawlerEngineInput() interface{}
	CrawlerInput() interface{}
	CrawlerOther() interface{}
	SetCrawlerOther(val interface{})
	CrawlerOtherInput() interface{}
	CrawlerScanner() interface{}
	SetCrawlerScanner(val interface{})
	CrawlerScannerInput() interface{}
	CrawlerScript() interface{}
	SetCrawlerScript(val interface{})
	CrawlerScriptInput() interface{}
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	Custom() interface{}
	SetCustom(val interface{})
	CustomInput() interface{}
	// Experimental.
	Fqn() *string
	Ignore() interface{}
	SetIgnore(val interface{})
	IgnoreInput() interface{}
	InternalValue() *WafPolicyV1Options
	SetInternalValue(val *WafPolicyV1Options)
	Privacy() interface{}
	SetPrivacy(val interface{})
	PrivacyInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Webattack() interface{}
	SetWebattack(val interface{})
	WebattackInput() interface{}
	Webshell() interface{}
	SetWebshell(val interface{})
	WebshellInput() interface{}
	Whiteblackip() interface{}
	SetWhiteblackip(val interface{})
	WhiteblackipInput() interface{}
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
	ResetAntitamper()
	ResetCc()
	ResetCommon()
	ResetCrawler()
	ResetCrawlerEngine()
	ResetCrawlerOther()
	ResetCrawlerScanner()
	ResetCrawlerScript()
	ResetCustom()
	ResetIgnore()
	ResetPrivacy()
	ResetWebattack()
	ResetWebshell()
	ResetWhiteblackip()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for WafPolicyV1OptionsOutputReference
type jsiiProxy_WafPolicyV1OptionsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_WafPolicyV1OptionsOutputReference) Antitamper() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"antitamper",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafPolicyV1OptionsOutputReference) AntitamperInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"antitamperInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafPolicyV1OptionsOutputReference) Cc() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafPolicyV1OptionsOutputReference) CcInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ccInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafPolicyV1OptionsOutputReference) Common() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"common",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafPolicyV1OptionsOutputReference) CommonInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"commonInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafPolicyV1OptionsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafPolicyV1OptionsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafPolicyV1OptionsOutputReference) Crawler() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"crawler",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafPolicyV1OptionsOutputReference) CrawlerEngine() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"crawlerEngine",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafPolicyV1OptionsOutputReference) CrawlerEngineInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"crawlerEngineInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafPolicyV1OptionsOutputReference) CrawlerInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"crawlerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafPolicyV1OptionsOutputReference) CrawlerOther() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"crawlerOther",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafPolicyV1OptionsOutputReference) CrawlerOtherInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"crawlerOtherInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafPolicyV1OptionsOutputReference) CrawlerScanner() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"crawlerScanner",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafPolicyV1OptionsOutputReference) CrawlerScannerInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"crawlerScannerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafPolicyV1OptionsOutputReference) CrawlerScript() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"crawlerScript",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafPolicyV1OptionsOutputReference) CrawlerScriptInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"crawlerScriptInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafPolicyV1OptionsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafPolicyV1OptionsOutputReference) Custom() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"custom",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafPolicyV1OptionsOutputReference) CustomInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"customInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafPolicyV1OptionsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafPolicyV1OptionsOutputReference) Ignore() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ignore",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafPolicyV1OptionsOutputReference) IgnoreInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ignoreInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafPolicyV1OptionsOutputReference) InternalValue() *WafPolicyV1Options {
	var returns *WafPolicyV1Options
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafPolicyV1OptionsOutputReference) Privacy() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"privacy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafPolicyV1OptionsOutputReference) PrivacyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"privacyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafPolicyV1OptionsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafPolicyV1OptionsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafPolicyV1OptionsOutputReference) Webattack() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"webattack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafPolicyV1OptionsOutputReference) WebattackInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"webattackInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafPolicyV1OptionsOutputReference) Webshell() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"webshell",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafPolicyV1OptionsOutputReference) WebshellInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"webshellInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafPolicyV1OptionsOutputReference) Whiteblackip() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"whiteblackip",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafPolicyV1OptionsOutputReference) WhiteblackipInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"whiteblackipInput",
		&returns,
	)
	return returns
}


func NewWafPolicyV1OptionsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) WafPolicyV1OptionsOutputReference {
	_init_.Initialize()

	if err := validateNewWafPolicyV1OptionsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_WafPolicyV1OptionsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.wafPolicyV1.WafPolicyV1OptionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewWafPolicyV1OptionsOutputReference_Override(w WafPolicyV1OptionsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.wafPolicyV1.WafPolicyV1OptionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		w,
	)
}

func (j *jsiiProxy_WafPolicyV1OptionsOutputReference)SetAntitamper(val interface{}) {
	if err := j.validateSetAntitamperParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"antitamper",
		val,
	)
}

func (j *jsiiProxy_WafPolicyV1OptionsOutputReference)SetCc(val interface{}) {
	if err := j.validateSetCcParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cc",
		val,
	)
}

func (j *jsiiProxy_WafPolicyV1OptionsOutputReference)SetCommon(val interface{}) {
	if err := j.validateSetCommonParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"common",
		val,
	)
}

func (j *jsiiProxy_WafPolicyV1OptionsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_WafPolicyV1OptionsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_WafPolicyV1OptionsOutputReference)SetCrawler(val interface{}) {
	if err := j.validateSetCrawlerParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"crawler",
		val,
	)
}

func (j *jsiiProxy_WafPolicyV1OptionsOutputReference)SetCrawlerEngine(val interface{}) {
	if err := j.validateSetCrawlerEngineParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"crawlerEngine",
		val,
	)
}

func (j *jsiiProxy_WafPolicyV1OptionsOutputReference)SetCrawlerOther(val interface{}) {
	if err := j.validateSetCrawlerOtherParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"crawlerOther",
		val,
	)
}

func (j *jsiiProxy_WafPolicyV1OptionsOutputReference)SetCrawlerScanner(val interface{}) {
	if err := j.validateSetCrawlerScannerParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"crawlerScanner",
		val,
	)
}

func (j *jsiiProxy_WafPolicyV1OptionsOutputReference)SetCrawlerScript(val interface{}) {
	if err := j.validateSetCrawlerScriptParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"crawlerScript",
		val,
	)
}

func (j *jsiiProxy_WafPolicyV1OptionsOutputReference)SetCustom(val interface{}) {
	if err := j.validateSetCustomParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"custom",
		val,
	)
}

func (j *jsiiProxy_WafPolicyV1OptionsOutputReference)SetIgnore(val interface{}) {
	if err := j.validateSetIgnoreParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ignore",
		val,
	)
}

func (j *jsiiProxy_WafPolicyV1OptionsOutputReference)SetInternalValue(val *WafPolicyV1Options) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_WafPolicyV1OptionsOutputReference)SetPrivacy(val interface{}) {
	if err := j.validateSetPrivacyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"privacy",
		val,
	)
}

func (j *jsiiProxy_WafPolicyV1OptionsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_WafPolicyV1OptionsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_WafPolicyV1OptionsOutputReference)SetWebattack(val interface{}) {
	if err := j.validateSetWebattackParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"webattack",
		val,
	)
}

func (j *jsiiProxy_WafPolicyV1OptionsOutputReference)SetWebshell(val interface{}) {
	if err := j.validateSetWebshellParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"webshell",
		val,
	)
}

func (j *jsiiProxy_WafPolicyV1OptionsOutputReference)SetWhiteblackip(val interface{}) {
	if err := j.validateSetWhiteblackipParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"whiteblackip",
		val,
	)
}

func (w *jsiiProxy_WafPolicyV1OptionsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WafPolicyV1OptionsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (w *jsiiProxy_WafPolicyV1OptionsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (w *jsiiProxy_WafPolicyV1OptionsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (w *jsiiProxy_WafPolicyV1OptionsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (w *jsiiProxy_WafPolicyV1OptionsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (w *jsiiProxy_WafPolicyV1OptionsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (w *jsiiProxy_WafPolicyV1OptionsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (w *jsiiProxy_WafPolicyV1OptionsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (w *jsiiProxy_WafPolicyV1OptionsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (w *jsiiProxy_WafPolicyV1OptionsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WafPolicyV1OptionsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := w.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WafPolicyV1OptionsOutputReference) ResetAntitamper() {
	_jsii_.InvokeVoid(
		w,
		"resetAntitamper",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WafPolicyV1OptionsOutputReference) ResetCc() {
	_jsii_.InvokeVoid(
		w,
		"resetCc",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WafPolicyV1OptionsOutputReference) ResetCommon() {
	_jsii_.InvokeVoid(
		w,
		"resetCommon",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WafPolicyV1OptionsOutputReference) ResetCrawler() {
	_jsii_.InvokeVoid(
		w,
		"resetCrawler",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WafPolicyV1OptionsOutputReference) ResetCrawlerEngine() {
	_jsii_.InvokeVoid(
		w,
		"resetCrawlerEngine",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WafPolicyV1OptionsOutputReference) ResetCrawlerOther() {
	_jsii_.InvokeVoid(
		w,
		"resetCrawlerOther",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WafPolicyV1OptionsOutputReference) ResetCrawlerScanner() {
	_jsii_.InvokeVoid(
		w,
		"resetCrawlerScanner",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WafPolicyV1OptionsOutputReference) ResetCrawlerScript() {
	_jsii_.InvokeVoid(
		w,
		"resetCrawlerScript",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WafPolicyV1OptionsOutputReference) ResetCustom() {
	_jsii_.InvokeVoid(
		w,
		"resetCustom",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WafPolicyV1OptionsOutputReference) ResetIgnore() {
	_jsii_.InvokeVoid(
		w,
		"resetIgnore",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WafPolicyV1OptionsOutputReference) ResetPrivacy() {
	_jsii_.InvokeVoid(
		w,
		"resetPrivacy",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WafPolicyV1OptionsOutputReference) ResetWebattack() {
	_jsii_.InvokeVoid(
		w,
		"resetWebattack",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WafPolicyV1OptionsOutputReference) ResetWebshell() {
	_jsii_.InvokeVoid(
		w,
		"resetWebshell",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WafPolicyV1OptionsOutputReference) ResetWhiteblackip() {
	_jsii_.InvokeVoid(
		w,
		"resetWhiteblackip",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WafPolicyV1OptionsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := w.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		w,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WafPolicyV1OptionsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

