package cbrvaultv3

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v8/jsii"

	"github.com/cdktf/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v8/cbrvaultv3/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CbrVaultV3BillingOutputReference interface {
	cdktf.ComplexObject
	Allocated() *float64
	ChargingMode() *string
	SetChargingMode(val *string)
	ChargingModeInput() *string
	CloudType() *string
	SetCloudType(val *string)
	CloudTypeInput() *string
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
	ConsistentLevel() *string
	SetConsistentLevel(val *string)
	ConsistentLevelInput() *string
	ConsoleUrl() *string
	SetConsoleUrl(val *string)
	ConsoleUrlInput() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	ExtraInfo() *map[string]*string
	SetExtraInfo(val *map[string]*string)
	ExtraInfoInput() *map[string]*string
	// Experimental.
	Fqn() *string
	FrozenScene() *string
	InternalValue() *CbrVaultV3Billing
	SetInternalValue(val *CbrVaultV3Billing)
	IsAutoPay() interface{}
	SetIsAutoPay(val interface{})
	IsAutoPayInput() interface{}
	IsAutoRenew() interface{}
	SetIsAutoRenew(val interface{})
	IsAutoRenewInput() interface{}
	ObjectType() *string
	SetObjectType(val *string)
	ObjectTypeInput() *string
	OrderId() *string
	PeriodNum() *float64
	SetPeriodNum(val *float64)
	PeriodNumInput() *float64
	PeriodType() *string
	SetPeriodType(val *string)
	PeriodTypeInput() *string
	ProductId() *string
	ProtectType() *string
	SetProtectType(val *string)
	ProtectTypeInput() *string
	Size() *float64
	SetSize(val *float64)
	SizeInput() *float64
	SpecCode() *string
	Status() *string
	StorageUnit() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Used() *float64
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
	ResetChargingMode()
	ResetCloudType()
	ResetConsistentLevel()
	ResetConsoleUrl()
	ResetExtraInfo()
	ResetIsAutoPay()
	ResetIsAutoRenew()
	ResetPeriodNum()
	ResetPeriodType()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for CbrVaultV3BillingOutputReference
type jsiiProxy_CbrVaultV3BillingOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CbrVaultV3BillingOutputReference) Allocated() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"allocated",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CbrVaultV3BillingOutputReference) ChargingMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"chargingMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CbrVaultV3BillingOutputReference) ChargingModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"chargingModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CbrVaultV3BillingOutputReference) CloudType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CbrVaultV3BillingOutputReference) CloudTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CbrVaultV3BillingOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CbrVaultV3BillingOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CbrVaultV3BillingOutputReference) ConsistentLevel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"consistentLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CbrVaultV3BillingOutputReference) ConsistentLevelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"consistentLevelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CbrVaultV3BillingOutputReference) ConsoleUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"consoleUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CbrVaultV3BillingOutputReference) ConsoleUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"consoleUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CbrVaultV3BillingOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CbrVaultV3BillingOutputReference) ExtraInfo() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"extraInfo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CbrVaultV3BillingOutputReference) ExtraInfoInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"extraInfoInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CbrVaultV3BillingOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CbrVaultV3BillingOutputReference) FrozenScene() *string {
	var returns *string
	_jsii_.Get(
		j,
		"frozenScene",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CbrVaultV3BillingOutputReference) InternalValue() *CbrVaultV3Billing {
	var returns *CbrVaultV3Billing
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CbrVaultV3BillingOutputReference) IsAutoPay() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isAutoPay",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CbrVaultV3BillingOutputReference) IsAutoPayInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isAutoPayInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CbrVaultV3BillingOutputReference) IsAutoRenew() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isAutoRenew",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CbrVaultV3BillingOutputReference) IsAutoRenewInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isAutoRenewInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CbrVaultV3BillingOutputReference) ObjectType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"objectType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CbrVaultV3BillingOutputReference) ObjectTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"objectTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CbrVaultV3BillingOutputReference) OrderId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"orderId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CbrVaultV3BillingOutputReference) PeriodNum() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"periodNum",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CbrVaultV3BillingOutputReference) PeriodNumInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"periodNumInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CbrVaultV3BillingOutputReference) PeriodType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"periodType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CbrVaultV3BillingOutputReference) PeriodTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"periodTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CbrVaultV3BillingOutputReference) ProductId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"productId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CbrVaultV3BillingOutputReference) ProtectType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"protectType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CbrVaultV3BillingOutputReference) ProtectTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"protectTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CbrVaultV3BillingOutputReference) Size() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"size",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CbrVaultV3BillingOutputReference) SizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CbrVaultV3BillingOutputReference) SpecCode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"specCode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CbrVaultV3BillingOutputReference) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CbrVaultV3BillingOutputReference) StorageUnit() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageUnit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CbrVaultV3BillingOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CbrVaultV3BillingOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CbrVaultV3BillingOutputReference) Used() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"used",
		&returns,
	)
	return returns
}


func NewCbrVaultV3BillingOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) CbrVaultV3BillingOutputReference {
	_init_.Initialize()

	if err := validateNewCbrVaultV3BillingOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_CbrVaultV3BillingOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.cbrVaultV3.CbrVaultV3BillingOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewCbrVaultV3BillingOutputReference_Override(c CbrVaultV3BillingOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.cbrVaultV3.CbrVaultV3BillingOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_CbrVaultV3BillingOutputReference)SetChargingMode(val *string) {
	if err := j.validateSetChargingModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"chargingMode",
		val,
	)
}

func (j *jsiiProxy_CbrVaultV3BillingOutputReference)SetCloudType(val *string) {
	if err := j.validateSetCloudTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cloudType",
		val,
	)
}

func (j *jsiiProxy_CbrVaultV3BillingOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CbrVaultV3BillingOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CbrVaultV3BillingOutputReference)SetConsistentLevel(val *string) {
	if err := j.validateSetConsistentLevelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"consistentLevel",
		val,
	)
}

func (j *jsiiProxy_CbrVaultV3BillingOutputReference)SetConsoleUrl(val *string) {
	if err := j.validateSetConsoleUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"consoleUrl",
		val,
	)
}

func (j *jsiiProxy_CbrVaultV3BillingOutputReference)SetExtraInfo(val *map[string]*string) {
	if err := j.validateSetExtraInfoParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"extraInfo",
		val,
	)
}

func (j *jsiiProxy_CbrVaultV3BillingOutputReference)SetInternalValue(val *CbrVaultV3Billing) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CbrVaultV3BillingOutputReference)SetIsAutoPay(val interface{}) {
	if err := j.validateSetIsAutoPayParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"isAutoPay",
		val,
	)
}

func (j *jsiiProxy_CbrVaultV3BillingOutputReference)SetIsAutoRenew(val interface{}) {
	if err := j.validateSetIsAutoRenewParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"isAutoRenew",
		val,
	)
}

func (j *jsiiProxy_CbrVaultV3BillingOutputReference)SetObjectType(val *string) {
	if err := j.validateSetObjectTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"objectType",
		val,
	)
}

func (j *jsiiProxy_CbrVaultV3BillingOutputReference)SetPeriodNum(val *float64) {
	if err := j.validateSetPeriodNumParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"periodNum",
		val,
	)
}

func (j *jsiiProxy_CbrVaultV3BillingOutputReference)SetPeriodType(val *string) {
	if err := j.validateSetPeriodTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"periodType",
		val,
	)
}

func (j *jsiiProxy_CbrVaultV3BillingOutputReference)SetProtectType(val *string) {
	if err := j.validateSetProtectTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"protectType",
		val,
	)
}

func (j *jsiiProxy_CbrVaultV3BillingOutputReference)SetSize(val *float64) {
	if err := j.validateSetSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"size",
		val,
	)
}

func (j *jsiiProxy_CbrVaultV3BillingOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CbrVaultV3BillingOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_CbrVaultV3BillingOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CbrVaultV3BillingOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_CbrVaultV3BillingOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CbrVaultV3BillingOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_CbrVaultV3BillingOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_CbrVaultV3BillingOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_CbrVaultV3BillingOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_CbrVaultV3BillingOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_CbrVaultV3BillingOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_CbrVaultV3BillingOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_CbrVaultV3BillingOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CbrVaultV3BillingOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := c.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CbrVaultV3BillingOutputReference) ResetChargingMode() {
	_jsii_.InvokeVoid(
		c,
		"resetChargingMode",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CbrVaultV3BillingOutputReference) ResetCloudType() {
	_jsii_.InvokeVoid(
		c,
		"resetCloudType",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CbrVaultV3BillingOutputReference) ResetConsistentLevel() {
	_jsii_.InvokeVoid(
		c,
		"resetConsistentLevel",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CbrVaultV3BillingOutputReference) ResetConsoleUrl() {
	_jsii_.InvokeVoid(
		c,
		"resetConsoleUrl",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CbrVaultV3BillingOutputReference) ResetExtraInfo() {
	_jsii_.InvokeVoid(
		c,
		"resetExtraInfo",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CbrVaultV3BillingOutputReference) ResetIsAutoPay() {
	_jsii_.InvokeVoid(
		c,
		"resetIsAutoPay",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CbrVaultV3BillingOutputReference) ResetIsAutoRenew() {
	_jsii_.InvokeVoid(
		c,
		"resetIsAutoRenew",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CbrVaultV3BillingOutputReference) ResetPeriodNum() {
	_jsii_.InvokeVoid(
		c,
		"resetPeriodNum",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CbrVaultV3BillingOutputReference) ResetPeriodType() {
	_jsii_.InvokeVoid(
		c,
		"resetPeriodType",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CbrVaultV3BillingOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := c.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		c,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CbrVaultV3BillingOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

