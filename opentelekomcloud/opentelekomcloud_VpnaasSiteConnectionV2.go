// Prebuilt opentelekomcloud Provider for Terraform CDK (cdktf)
package opentelekomcloud

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v2/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/cdktf-provider-opentelekomcloud-go/opentelekomcloud/v2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://www.terraform.io/docs/providers/opentelekomcloud/r/vpnaas_site_connection_v2 opentelekomcloud_vpnaas_site_connection_v2}.
type VpnaasSiteConnectionV2 interface {
	cdktf.TerraformResource
	AdminStateUp() interface{}
	SetAdminStateUp(val interface{})
	AdminStateUpInput() interface{}
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
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	Dpd() VpnaasSiteConnectionV2DpdList
	DpdInput() interface{}
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
	IkepolicyId() *string
	SetIkepolicyId(val *string)
	IkepolicyIdInput() *string
	Initiator() *string
	SetInitiator(val *string)
	InitiatorInput() *string
	IpsecpolicyId() *string
	SetIpsecpolicyId(val *string)
	IpsecpolicyIdInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	LocalEpGroupId() *string
	SetLocalEpGroupId(val *string)
	LocalEpGroupIdInput() *string
	LocalId() *string
	SetLocalId(val *string)
	LocalIdInput() *string
	Mtu() *float64
	SetMtu(val *float64)
	MtuInput() *float64
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	PeerAddress() *string
	SetPeerAddress(val *string)
	PeerAddressInput() *string
	PeerCidrs() *[]*string
	SetPeerCidrs(val *[]*string)
	PeerCidrsInput() *[]*string
	PeerEpGroupId() *string
	SetPeerEpGroupId(val *string)
	PeerEpGroupIdInput() *string
	PeerId() *string
	SetPeerId(val *string)
	PeerIdInput() *string
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	Psk() *string
	SetPsk(val *string)
	PskInput() *string
	// Experimental.
	RawOverrides() interface{}
	Region() *string
	SetRegion(val *string)
	RegionInput() *string
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsInput() *map[string]*string
	TenantId() *string
	SetTenantId(val *string)
	TenantIdInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() VpnaasSiteConnectionV2TimeoutsOutputReference
	TimeoutsInput() interface{}
	ValueSpecs() *map[string]*string
	SetValueSpecs(val *map[string]*string)
	ValueSpecsInput() *map[string]*string
	VpnserviceId() *string
	SetVpnserviceId(val *string)
	VpnserviceIdInput() *string
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
	PutDpd(value interface{})
	PutTimeouts(value *VpnaasSiteConnectionV2Timeouts)
	ResetAdminStateUp()
	ResetDescription()
	ResetDpd()
	ResetId()
	ResetInitiator()
	ResetLocalEpGroupId()
	ResetLocalId()
	ResetMtu()
	ResetName()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPeerCidrs()
	ResetPeerEpGroupId()
	ResetRegion()
	ResetTags()
	ResetTenantId()
	ResetTimeouts()
	ResetValueSpecs()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for VpnaasSiteConnectionV2
type jsiiProxy_VpnaasSiteConnectionV2 struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_VpnaasSiteConnectionV2) AdminStateUp() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"adminStateUp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnaasSiteConnectionV2) AdminStateUpInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"adminStateUpInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnaasSiteConnectionV2) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnaasSiteConnectionV2) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnaasSiteConnectionV2) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnaasSiteConnectionV2) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnaasSiteConnectionV2) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnaasSiteConnectionV2) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnaasSiteConnectionV2) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnaasSiteConnectionV2) Dpd() VpnaasSiteConnectionV2DpdList {
	var returns VpnaasSiteConnectionV2DpdList
	_jsii_.Get(
		j,
		"dpd",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnaasSiteConnectionV2) DpdInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dpdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnaasSiteConnectionV2) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnaasSiteConnectionV2) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnaasSiteConnectionV2) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnaasSiteConnectionV2) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnaasSiteConnectionV2) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnaasSiteConnectionV2) IkepolicyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ikepolicyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnaasSiteConnectionV2) IkepolicyIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ikepolicyIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnaasSiteConnectionV2) Initiator() *string {
	var returns *string
	_jsii_.Get(
		j,
		"initiator",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnaasSiteConnectionV2) InitiatorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"initiatorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnaasSiteConnectionV2) IpsecpolicyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipsecpolicyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnaasSiteConnectionV2) IpsecpolicyIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipsecpolicyIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnaasSiteConnectionV2) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnaasSiteConnectionV2) LocalEpGroupId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"localEpGroupId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnaasSiteConnectionV2) LocalEpGroupIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"localEpGroupIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnaasSiteConnectionV2) LocalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"localId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnaasSiteConnectionV2) LocalIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"localIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnaasSiteConnectionV2) Mtu() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"mtu",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnaasSiteConnectionV2) MtuInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"mtuInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnaasSiteConnectionV2) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnaasSiteConnectionV2) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnaasSiteConnectionV2) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnaasSiteConnectionV2) PeerAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"peerAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnaasSiteConnectionV2) PeerAddressInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"peerAddressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnaasSiteConnectionV2) PeerCidrs() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"peerCidrs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnaasSiteConnectionV2) PeerCidrsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"peerCidrsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnaasSiteConnectionV2) PeerEpGroupId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"peerEpGroupId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnaasSiteConnectionV2) PeerEpGroupIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"peerEpGroupIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnaasSiteConnectionV2) PeerId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"peerId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnaasSiteConnectionV2) PeerIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"peerIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnaasSiteConnectionV2) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnaasSiteConnectionV2) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnaasSiteConnectionV2) Psk() *string {
	var returns *string
	_jsii_.Get(
		j,
		"psk",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnaasSiteConnectionV2) PskInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pskInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnaasSiteConnectionV2) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnaasSiteConnectionV2) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnaasSiteConnectionV2) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnaasSiteConnectionV2) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnaasSiteConnectionV2) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnaasSiteConnectionV2) TenantId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tenantId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnaasSiteConnectionV2) TenantIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tenantIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnaasSiteConnectionV2) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnaasSiteConnectionV2) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnaasSiteConnectionV2) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnaasSiteConnectionV2) Timeouts() VpnaasSiteConnectionV2TimeoutsOutputReference {
	var returns VpnaasSiteConnectionV2TimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnaasSiteConnectionV2) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnaasSiteConnectionV2) ValueSpecs() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"valueSpecs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnaasSiteConnectionV2) ValueSpecsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"valueSpecsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnaasSiteConnectionV2) VpnserviceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpnserviceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnaasSiteConnectionV2) VpnserviceIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpnserviceIdInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://www.terraform.io/docs/providers/opentelekomcloud/r/vpnaas_site_connection_v2 opentelekomcloud_vpnaas_site_connection_v2} Resource.
func NewVpnaasSiteConnectionV2(scope constructs.Construct, id *string, config *VpnaasSiteConnectionV2Config) VpnaasSiteConnectionV2 {
	_init_.Initialize()

	j := jsiiProxy_VpnaasSiteConnectionV2{}

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.VpnaasSiteConnectionV2",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/opentelekomcloud/r/vpnaas_site_connection_v2 opentelekomcloud_vpnaas_site_connection_v2} Resource.
func NewVpnaasSiteConnectionV2_Override(v VpnaasSiteConnectionV2, scope constructs.Construct, id *string, config *VpnaasSiteConnectionV2Config) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-opentelekomcloud.VpnaasSiteConnectionV2",
		[]interface{}{scope, id, config},
		v,
	)
}

func (j *jsiiProxy_VpnaasSiteConnectionV2) SetAdminStateUp(val interface{}) {
	_jsii_.Set(
		j,
		"adminStateUp",
		val,
	)
}

func (j *jsiiProxy_VpnaasSiteConnectionV2) SetConnection(val interface{}) {
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_VpnaasSiteConnectionV2) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_VpnaasSiteConnectionV2) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_VpnaasSiteConnectionV2) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_VpnaasSiteConnectionV2) SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_VpnaasSiteConnectionV2) SetId(val *string) {
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_VpnaasSiteConnectionV2) SetIkepolicyId(val *string) {
	_jsii_.Set(
		j,
		"ikepolicyId",
		val,
	)
}

func (j *jsiiProxy_VpnaasSiteConnectionV2) SetInitiator(val *string) {
	_jsii_.Set(
		j,
		"initiator",
		val,
	)
}

func (j *jsiiProxy_VpnaasSiteConnectionV2) SetIpsecpolicyId(val *string) {
	_jsii_.Set(
		j,
		"ipsecpolicyId",
		val,
	)
}

func (j *jsiiProxy_VpnaasSiteConnectionV2) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_VpnaasSiteConnectionV2) SetLocalEpGroupId(val *string) {
	_jsii_.Set(
		j,
		"localEpGroupId",
		val,
	)
}

func (j *jsiiProxy_VpnaasSiteConnectionV2) SetLocalId(val *string) {
	_jsii_.Set(
		j,
		"localId",
		val,
	)
}

func (j *jsiiProxy_VpnaasSiteConnectionV2) SetMtu(val *float64) {
	_jsii_.Set(
		j,
		"mtu",
		val,
	)
}

func (j *jsiiProxy_VpnaasSiteConnectionV2) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_VpnaasSiteConnectionV2) SetPeerAddress(val *string) {
	_jsii_.Set(
		j,
		"peerAddress",
		val,
	)
}

func (j *jsiiProxy_VpnaasSiteConnectionV2) SetPeerCidrs(val *[]*string) {
	_jsii_.Set(
		j,
		"peerCidrs",
		val,
	)
}

func (j *jsiiProxy_VpnaasSiteConnectionV2) SetPeerEpGroupId(val *string) {
	_jsii_.Set(
		j,
		"peerEpGroupId",
		val,
	)
}

func (j *jsiiProxy_VpnaasSiteConnectionV2) SetPeerId(val *string) {
	_jsii_.Set(
		j,
		"peerId",
		val,
	)
}

func (j *jsiiProxy_VpnaasSiteConnectionV2) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_VpnaasSiteConnectionV2) SetProvisioners(val *[]interface{}) {
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_VpnaasSiteConnectionV2) SetPsk(val *string) {
	_jsii_.Set(
		j,
		"psk",
		val,
	)
}

func (j *jsiiProxy_VpnaasSiteConnectionV2) SetRegion(val *string) {
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_VpnaasSiteConnectionV2) SetTags(val *map[string]*string) {
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_VpnaasSiteConnectionV2) SetTenantId(val *string) {
	_jsii_.Set(
		j,
		"tenantId",
		val,
	)
}

func (j *jsiiProxy_VpnaasSiteConnectionV2) SetValueSpecs(val *map[string]*string) {
	_jsii_.Set(
		j,
		"valueSpecs",
		val,
	)
}

func (j *jsiiProxy_VpnaasSiteConnectionV2) SetVpnserviceId(val *string) {
	_jsii_.Set(
		j,
		"vpnserviceId",
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
func VpnaasSiteConnectionV2_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-opentelekomcloud.VpnaasSiteConnectionV2",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func VpnaasSiteConnectionV2_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-opentelekomcloud.VpnaasSiteConnectionV2",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (v *jsiiProxy_VpnaasSiteConnectionV2) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		v,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (v *jsiiProxy_VpnaasSiteConnectionV2) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		v,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpnaasSiteConnectionV2) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpnaasSiteConnectionV2) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		v,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpnaasSiteConnectionV2) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		v,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpnaasSiteConnectionV2) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		v,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpnaasSiteConnectionV2) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		v,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpnaasSiteConnectionV2) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		v,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpnaasSiteConnectionV2) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpnaasSiteConnectionV2) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		v,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpnaasSiteConnectionV2) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpnaasSiteConnectionV2) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		v,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (v *jsiiProxy_VpnaasSiteConnectionV2) PutDpd(value interface{}) {
	_jsii_.InvokeVoid(
		v,
		"putDpd",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VpnaasSiteConnectionV2) PutTimeouts(value *VpnaasSiteConnectionV2Timeouts) {
	_jsii_.InvokeVoid(
		v,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VpnaasSiteConnectionV2) ResetAdminStateUp() {
	_jsii_.InvokeVoid(
		v,
		"resetAdminStateUp",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpnaasSiteConnectionV2) ResetDescription() {
	_jsii_.InvokeVoid(
		v,
		"resetDescription",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpnaasSiteConnectionV2) ResetDpd() {
	_jsii_.InvokeVoid(
		v,
		"resetDpd",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpnaasSiteConnectionV2) ResetId() {
	_jsii_.InvokeVoid(
		v,
		"resetId",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpnaasSiteConnectionV2) ResetInitiator() {
	_jsii_.InvokeVoid(
		v,
		"resetInitiator",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpnaasSiteConnectionV2) ResetLocalEpGroupId() {
	_jsii_.InvokeVoid(
		v,
		"resetLocalEpGroupId",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpnaasSiteConnectionV2) ResetLocalId() {
	_jsii_.InvokeVoid(
		v,
		"resetLocalId",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpnaasSiteConnectionV2) ResetMtu() {
	_jsii_.InvokeVoid(
		v,
		"resetMtu",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpnaasSiteConnectionV2) ResetName() {
	_jsii_.InvokeVoid(
		v,
		"resetName",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpnaasSiteConnectionV2) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		v,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpnaasSiteConnectionV2) ResetPeerCidrs() {
	_jsii_.InvokeVoid(
		v,
		"resetPeerCidrs",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpnaasSiteConnectionV2) ResetPeerEpGroupId() {
	_jsii_.InvokeVoid(
		v,
		"resetPeerEpGroupId",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpnaasSiteConnectionV2) ResetRegion() {
	_jsii_.InvokeVoid(
		v,
		"resetRegion",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpnaasSiteConnectionV2) ResetTags() {
	_jsii_.InvokeVoid(
		v,
		"resetTags",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpnaasSiteConnectionV2) ResetTenantId() {
	_jsii_.InvokeVoid(
		v,
		"resetTenantId",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpnaasSiteConnectionV2) ResetTimeouts() {
	_jsii_.InvokeVoid(
		v,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpnaasSiteConnectionV2) ResetValueSpecs() {
	_jsii_.InvokeVoid(
		v,
		"resetValueSpecs",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpnaasSiteConnectionV2) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		v,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpnaasSiteConnectionV2) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		v,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpnaasSiteConnectionV2) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpnaasSiteConnectionV2) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		v,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}
