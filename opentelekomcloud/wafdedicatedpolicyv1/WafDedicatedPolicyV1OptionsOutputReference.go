// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package wafdedicatedpolicyv1

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v10/jsii"

	"github.com/cdktf/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v10/wafdedicatedpolicyv1/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type WafDedicatedPolicyV1OptionsOutputReference interface {
	cdktf.ComplexObject
	AntiCrawler() interface{}
	SetAntiCrawler(val interface{})
	AntiCrawlerInput() interface{}
	AntiLeakage() interface{}
	SetAntiLeakage(val interface{})
	AntiLeakageInput() interface{}
	AntiTamper() interface{}
	SetAntiTamper(val interface{})
	AntiTamperInput() interface{}
	Blacklist() interface{}
	SetBlacklist(val interface{})
	BlacklistInput() interface{}
	BotEnable() cdktf.IResolvable
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
	FollowedAction() interface{}
	SetFollowedAction(val interface{})
	FollowedActionInput() interface{}
	// Experimental.
	Fqn() *string
	GeolocationAccessControl() interface{}
	SetGeolocationAccessControl(val interface{})
	GeolocationAccessControlInput() interface{}
	Ignore() interface{}
	SetIgnore(val interface{})
	IgnoreInput() interface{}
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Precise() cdktf.IResolvable
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
	WebAttack() interface{}
	SetWebAttack(val interface{})
	WebAttackInput() interface{}
	WebShell() interface{}
	SetWebShell(val interface{})
	WebShellInput() interface{}
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
	ResetAntiCrawler()
	ResetAntiLeakage()
	ResetAntiTamper()
	ResetBlacklist()
	ResetCc()
	ResetCommon()
	ResetCrawler()
	ResetCrawlerEngine()
	ResetCrawlerOther()
	ResetCrawlerScanner()
	ResetCrawlerScript()
	ResetCustom()
	ResetFollowedAction()
	ResetGeolocationAccessControl()
	ResetIgnore()
	ResetPrivacy()
	ResetWebAttack()
	ResetWebShell()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for WafDedicatedPolicyV1OptionsOutputReference
type jsiiProxy_WafDedicatedPolicyV1OptionsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_WafDedicatedPolicyV1OptionsOutputReference) AntiCrawler() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"antiCrawler",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafDedicatedPolicyV1OptionsOutputReference) AntiCrawlerInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"antiCrawlerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafDedicatedPolicyV1OptionsOutputReference) AntiLeakage() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"antiLeakage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafDedicatedPolicyV1OptionsOutputReference) AntiLeakageInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"antiLeakageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafDedicatedPolicyV1OptionsOutputReference) AntiTamper() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"antiTamper",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafDedicatedPolicyV1OptionsOutputReference) AntiTamperInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"antiTamperInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafDedicatedPolicyV1OptionsOutputReference) Blacklist() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"blacklist",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafDedicatedPolicyV1OptionsOutputReference) BlacklistInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"blacklistInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafDedicatedPolicyV1OptionsOutputReference) BotEnable() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"botEnable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafDedicatedPolicyV1OptionsOutputReference) Cc() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafDedicatedPolicyV1OptionsOutputReference) CcInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ccInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafDedicatedPolicyV1OptionsOutputReference) Common() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"common",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafDedicatedPolicyV1OptionsOutputReference) CommonInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"commonInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafDedicatedPolicyV1OptionsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafDedicatedPolicyV1OptionsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafDedicatedPolicyV1OptionsOutputReference) Crawler() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"crawler",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafDedicatedPolicyV1OptionsOutputReference) CrawlerEngine() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"crawlerEngine",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafDedicatedPolicyV1OptionsOutputReference) CrawlerEngineInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"crawlerEngineInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafDedicatedPolicyV1OptionsOutputReference) CrawlerInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"crawlerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafDedicatedPolicyV1OptionsOutputReference) CrawlerOther() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"crawlerOther",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafDedicatedPolicyV1OptionsOutputReference) CrawlerOtherInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"crawlerOtherInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafDedicatedPolicyV1OptionsOutputReference) CrawlerScanner() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"crawlerScanner",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafDedicatedPolicyV1OptionsOutputReference) CrawlerScannerInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"crawlerScannerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafDedicatedPolicyV1OptionsOutputReference) CrawlerScript() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"crawlerScript",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafDedicatedPolicyV1OptionsOutputReference) CrawlerScriptInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"crawlerScriptInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafDedicatedPolicyV1OptionsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafDedicatedPolicyV1OptionsOutputReference) Custom() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"custom",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafDedicatedPolicyV1OptionsOutputReference) CustomInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"customInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafDedicatedPolicyV1OptionsOutputReference) FollowedAction() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"followedAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafDedicatedPolicyV1OptionsOutputReference) FollowedActionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"followedActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafDedicatedPolicyV1OptionsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafDedicatedPolicyV1OptionsOutputReference) GeolocationAccessControl() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"geolocationAccessControl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafDedicatedPolicyV1OptionsOutputReference) GeolocationAccessControlInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"geolocationAccessControlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafDedicatedPolicyV1OptionsOutputReference) Ignore() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ignore",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafDedicatedPolicyV1OptionsOutputReference) IgnoreInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ignoreInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafDedicatedPolicyV1OptionsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafDedicatedPolicyV1OptionsOutputReference) Precise() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"precise",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafDedicatedPolicyV1OptionsOutputReference) Privacy() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"privacy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafDedicatedPolicyV1OptionsOutputReference) PrivacyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"privacyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafDedicatedPolicyV1OptionsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafDedicatedPolicyV1OptionsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafDedicatedPolicyV1OptionsOutputReference) WebAttack() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"webAttack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafDedicatedPolicyV1OptionsOutputReference) WebAttackInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"webAttackInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafDedicatedPolicyV1OptionsOutputReference) WebShell() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"webShell",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafDedicatedPolicyV1OptionsOutputReference) WebShellInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"webShellInput",
		&returns,
	)
	return returns
}


func NewWafDedicatedPolicyV1OptionsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) WafDedicatedPolicyV1OptionsOutputReference {
	_init_.Initialize()

	if err := validateNewWafDedicatedPolicyV1OptionsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_WafDedicatedPolicyV1OptionsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.wafDedicatedPolicyV1.WafDedicatedPolicyV1OptionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewWafDedicatedPolicyV1OptionsOutputReference_Override(w WafDedicatedPolicyV1OptionsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.wafDedicatedPolicyV1.WafDedicatedPolicyV1OptionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		w,
	)
}

func (j *jsiiProxy_WafDedicatedPolicyV1OptionsOutputReference)SetAntiCrawler(val interface{}) {
	if err := j.validateSetAntiCrawlerParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"antiCrawler",
		val,
	)
}

func (j *jsiiProxy_WafDedicatedPolicyV1OptionsOutputReference)SetAntiLeakage(val interface{}) {
	if err := j.validateSetAntiLeakageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"antiLeakage",
		val,
	)
}

func (j *jsiiProxy_WafDedicatedPolicyV1OptionsOutputReference)SetAntiTamper(val interface{}) {
	if err := j.validateSetAntiTamperParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"antiTamper",
		val,
	)
}

func (j *jsiiProxy_WafDedicatedPolicyV1OptionsOutputReference)SetBlacklist(val interface{}) {
	if err := j.validateSetBlacklistParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"blacklist",
		val,
	)
}

func (j *jsiiProxy_WafDedicatedPolicyV1OptionsOutputReference)SetCc(val interface{}) {
	if err := j.validateSetCcParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cc",
		val,
	)
}

func (j *jsiiProxy_WafDedicatedPolicyV1OptionsOutputReference)SetCommon(val interface{}) {
	if err := j.validateSetCommonParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"common",
		val,
	)
}

func (j *jsiiProxy_WafDedicatedPolicyV1OptionsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_WafDedicatedPolicyV1OptionsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_WafDedicatedPolicyV1OptionsOutputReference)SetCrawler(val interface{}) {
	if err := j.validateSetCrawlerParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"crawler",
		val,
	)
}

func (j *jsiiProxy_WafDedicatedPolicyV1OptionsOutputReference)SetCrawlerEngine(val interface{}) {
	if err := j.validateSetCrawlerEngineParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"crawlerEngine",
		val,
	)
}

func (j *jsiiProxy_WafDedicatedPolicyV1OptionsOutputReference)SetCrawlerOther(val interface{}) {
	if err := j.validateSetCrawlerOtherParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"crawlerOther",
		val,
	)
}

func (j *jsiiProxy_WafDedicatedPolicyV1OptionsOutputReference)SetCrawlerScanner(val interface{}) {
	if err := j.validateSetCrawlerScannerParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"crawlerScanner",
		val,
	)
}

func (j *jsiiProxy_WafDedicatedPolicyV1OptionsOutputReference)SetCrawlerScript(val interface{}) {
	if err := j.validateSetCrawlerScriptParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"crawlerScript",
		val,
	)
}

func (j *jsiiProxy_WafDedicatedPolicyV1OptionsOutputReference)SetCustom(val interface{}) {
	if err := j.validateSetCustomParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"custom",
		val,
	)
}

func (j *jsiiProxy_WafDedicatedPolicyV1OptionsOutputReference)SetFollowedAction(val interface{}) {
	if err := j.validateSetFollowedActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"followedAction",
		val,
	)
}

func (j *jsiiProxy_WafDedicatedPolicyV1OptionsOutputReference)SetGeolocationAccessControl(val interface{}) {
	if err := j.validateSetGeolocationAccessControlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"geolocationAccessControl",
		val,
	)
}

func (j *jsiiProxy_WafDedicatedPolicyV1OptionsOutputReference)SetIgnore(val interface{}) {
	if err := j.validateSetIgnoreParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ignore",
		val,
	)
}

func (j *jsiiProxy_WafDedicatedPolicyV1OptionsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_WafDedicatedPolicyV1OptionsOutputReference)SetPrivacy(val interface{}) {
	if err := j.validateSetPrivacyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"privacy",
		val,
	)
}

func (j *jsiiProxy_WafDedicatedPolicyV1OptionsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_WafDedicatedPolicyV1OptionsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_WafDedicatedPolicyV1OptionsOutputReference)SetWebAttack(val interface{}) {
	if err := j.validateSetWebAttackParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"webAttack",
		val,
	)
}

func (j *jsiiProxy_WafDedicatedPolicyV1OptionsOutputReference)SetWebShell(val interface{}) {
	if err := j.validateSetWebShellParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"webShell",
		val,
	)
}

func (w *jsiiProxy_WafDedicatedPolicyV1OptionsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WafDedicatedPolicyV1OptionsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (w *jsiiProxy_WafDedicatedPolicyV1OptionsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (w *jsiiProxy_WafDedicatedPolicyV1OptionsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (w *jsiiProxy_WafDedicatedPolicyV1OptionsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (w *jsiiProxy_WafDedicatedPolicyV1OptionsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (w *jsiiProxy_WafDedicatedPolicyV1OptionsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (w *jsiiProxy_WafDedicatedPolicyV1OptionsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (w *jsiiProxy_WafDedicatedPolicyV1OptionsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (w *jsiiProxy_WafDedicatedPolicyV1OptionsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (w *jsiiProxy_WafDedicatedPolicyV1OptionsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WafDedicatedPolicyV1OptionsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (w *jsiiProxy_WafDedicatedPolicyV1OptionsOutputReference) ResetAntiCrawler() {
	_jsii_.InvokeVoid(
		w,
		"resetAntiCrawler",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WafDedicatedPolicyV1OptionsOutputReference) ResetAntiLeakage() {
	_jsii_.InvokeVoid(
		w,
		"resetAntiLeakage",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WafDedicatedPolicyV1OptionsOutputReference) ResetAntiTamper() {
	_jsii_.InvokeVoid(
		w,
		"resetAntiTamper",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WafDedicatedPolicyV1OptionsOutputReference) ResetBlacklist() {
	_jsii_.InvokeVoid(
		w,
		"resetBlacklist",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WafDedicatedPolicyV1OptionsOutputReference) ResetCc() {
	_jsii_.InvokeVoid(
		w,
		"resetCc",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WafDedicatedPolicyV1OptionsOutputReference) ResetCommon() {
	_jsii_.InvokeVoid(
		w,
		"resetCommon",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WafDedicatedPolicyV1OptionsOutputReference) ResetCrawler() {
	_jsii_.InvokeVoid(
		w,
		"resetCrawler",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WafDedicatedPolicyV1OptionsOutputReference) ResetCrawlerEngine() {
	_jsii_.InvokeVoid(
		w,
		"resetCrawlerEngine",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WafDedicatedPolicyV1OptionsOutputReference) ResetCrawlerOther() {
	_jsii_.InvokeVoid(
		w,
		"resetCrawlerOther",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WafDedicatedPolicyV1OptionsOutputReference) ResetCrawlerScanner() {
	_jsii_.InvokeVoid(
		w,
		"resetCrawlerScanner",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WafDedicatedPolicyV1OptionsOutputReference) ResetCrawlerScript() {
	_jsii_.InvokeVoid(
		w,
		"resetCrawlerScript",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WafDedicatedPolicyV1OptionsOutputReference) ResetCustom() {
	_jsii_.InvokeVoid(
		w,
		"resetCustom",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WafDedicatedPolicyV1OptionsOutputReference) ResetFollowedAction() {
	_jsii_.InvokeVoid(
		w,
		"resetFollowedAction",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WafDedicatedPolicyV1OptionsOutputReference) ResetGeolocationAccessControl() {
	_jsii_.InvokeVoid(
		w,
		"resetGeolocationAccessControl",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WafDedicatedPolicyV1OptionsOutputReference) ResetIgnore() {
	_jsii_.InvokeVoid(
		w,
		"resetIgnore",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WafDedicatedPolicyV1OptionsOutputReference) ResetPrivacy() {
	_jsii_.InvokeVoid(
		w,
		"resetPrivacy",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WafDedicatedPolicyV1OptionsOutputReference) ResetWebAttack() {
	_jsii_.InvokeVoid(
		w,
		"resetWebAttack",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WafDedicatedPolicyV1OptionsOutputReference) ResetWebShell() {
	_jsii_.InvokeVoid(
		w,
		"resetWebShell",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WafDedicatedPolicyV1OptionsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (w *jsiiProxy_WafDedicatedPolicyV1OptionsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

