# PrincipalView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Agents** | Pointer to [**[]PrincipalAgentBrief**](PrincipalAgentBrief.md) | Agents are the agents that answer to it, spawned ones with their parent. | [optional] 
**Compliance** | Pointer to [**PrincipalSummary**](PrincipalSummary.md) | Compliance is what it lacks to pay and be paid on the platform&#39;s rails. | [optional] 
**Entity** | Pointer to [**PrincipalEntity**](PrincipalEntity.md) | Entity is its legal entity, when it formed or imported one here. | [optional] 
**Identity** | Pointer to [**PrincipalIdentity**](PrincipalIdentity.md) | Identity is whether its founders are verified. | [optional] 
**Org** | Pointer to **string** | Org is the org — the principal itself. | [optional] 
**Sanctions** | Pointer to [**PrincipalScreening**](PrincipalScreening.md) | Sanctions is its screening against the OFAC, UN, EU and UK lists and the embargoed jurisdictions, run on this read and recorded by none: a match is recorded for a reviewer when a payment to or from the org is cleared. Matches are answered to the org&#39;s admins. | [optional] 
**Sources** | Pointer to [**[]PrincipalSource**](PrincipalSource.md) | Sources says which owning apps answered. | [optional] 
**Tax** | Pointer to [**PrincipalForm**](PrincipalForm.md) | Tax is the form it certified about itself. | [optional] 
**Wallets** | Pointer to [**[]PrincipalWalletBrief**](PrincipalWalletBrief.md) | Wallets are its addresses. | [optional] 

## Methods

### NewPrincipalView

`func NewPrincipalView() *PrincipalView`

NewPrincipalView instantiates a new PrincipalView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrincipalViewWithDefaults

`func NewPrincipalViewWithDefaults() *PrincipalView`

NewPrincipalViewWithDefaults instantiates a new PrincipalView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAgents

`func (o *PrincipalView) GetAgents() []PrincipalAgentBrief`

GetAgents returns the Agents field if non-nil, zero value otherwise.

### GetAgentsOk

`func (o *PrincipalView) GetAgentsOk() (*[]PrincipalAgentBrief, bool)`

GetAgentsOk returns a tuple with the Agents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgents

`func (o *PrincipalView) SetAgents(v []PrincipalAgentBrief)`

SetAgents sets Agents field to given value.

### HasAgents

`func (o *PrincipalView) HasAgents() bool`

HasAgents returns a boolean if a field has been set.

### GetCompliance

`func (o *PrincipalView) GetCompliance() PrincipalSummary`

GetCompliance returns the Compliance field if non-nil, zero value otherwise.

### GetComplianceOk

`func (o *PrincipalView) GetComplianceOk() (*PrincipalSummary, bool)`

GetComplianceOk returns a tuple with the Compliance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompliance

`func (o *PrincipalView) SetCompliance(v PrincipalSummary)`

SetCompliance sets Compliance field to given value.

### HasCompliance

`func (o *PrincipalView) HasCompliance() bool`

HasCompliance returns a boolean if a field has been set.

### GetEntity

`func (o *PrincipalView) GetEntity() PrincipalEntity`

GetEntity returns the Entity field if non-nil, zero value otherwise.

### GetEntityOk

`func (o *PrincipalView) GetEntityOk() (*PrincipalEntity, bool)`

GetEntityOk returns a tuple with the Entity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntity

`func (o *PrincipalView) SetEntity(v PrincipalEntity)`

SetEntity sets Entity field to given value.

### HasEntity

`func (o *PrincipalView) HasEntity() bool`

HasEntity returns a boolean if a field has been set.

### GetIdentity

`func (o *PrincipalView) GetIdentity() PrincipalIdentity`

GetIdentity returns the Identity field if non-nil, zero value otherwise.

### GetIdentityOk

`func (o *PrincipalView) GetIdentityOk() (*PrincipalIdentity, bool)`

GetIdentityOk returns a tuple with the Identity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentity

`func (o *PrincipalView) SetIdentity(v PrincipalIdentity)`

SetIdentity sets Identity field to given value.

### HasIdentity

`func (o *PrincipalView) HasIdentity() bool`

HasIdentity returns a boolean if a field has been set.

### GetOrg

`func (o *PrincipalView) GetOrg() string`

GetOrg returns the Org field if non-nil, zero value otherwise.

### GetOrgOk

`func (o *PrincipalView) GetOrgOk() (*string, bool)`

GetOrgOk returns a tuple with the Org field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrg

`func (o *PrincipalView) SetOrg(v string)`

SetOrg sets Org field to given value.

### HasOrg

`func (o *PrincipalView) HasOrg() bool`

HasOrg returns a boolean if a field has been set.

### GetSanctions

`func (o *PrincipalView) GetSanctions() PrincipalScreening`

GetSanctions returns the Sanctions field if non-nil, zero value otherwise.

### GetSanctionsOk

`func (o *PrincipalView) GetSanctionsOk() (*PrincipalScreening, bool)`

GetSanctionsOk returns a tuple with the Sanctions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSanctions

`func (o *PrincipalView) SetSanctions(v PrincipalScreening)`

SetSanctions sets Sanctions field to given value.

### HasSanctions

`func (o *PrincipalView) HasSanctions() bool`

HasSanctions returns a boolean if a field has been set.

### GetSources

`func (o *PrincipalView) GetSources() []PrincipalSource`

GetSources returns the Sources field if non-nil, zero value otherwise.

### GetSourcesOk

`func (o *PrincipalView) GetSourcesOk() (*[]PrincipalSource, bool)`

GetSourcesOk returns a tuple with the Sources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSources

`func (o *PrincipalView) SetSources(v []PrincipalSource)`

SetSources sets Sources field to given value.

### HasSources

`func (o *PrincipalView) HasSources() bool`

HasSources returns a boolean if a field has been set.

### GetTax

`func (o *PrincipalView) GetTax() PrincipalForm`

GetTax returns the Tax field if non-nil, zero value otherwise.

### GetTaxOk

`func (o *PrincipalView) GetTaxOk() (*PrincipalForm, bool)`

GetTaxOk returns a tuple with the Tax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTax

`func (o *PrincipalView) SetTax(v PrincipalForm)`

SetTax sets Tax field to given value.

### HasTax

`func (o *PrincipalView) HasTax() bool`

HasTax returns a boolean if a field has been set.

### GetWallets

`func (o *PrincipalView) GetWallets() []PrincipalWalletBrief`

GetWallets returns the Wallets field if non-nil, zero value otherwise.

### GetWalletsOk

`func (o *PrincipalView) GetWalletsOk() (*[]PrincipalWalletBrief, bool)`

GetWalletsOk returns a tuple with the Wallets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWallets

`func (o *PrincipalView) SetWallets(v []PrincipalWalletBrief)`

SetWallets sets Wallets field to given value.

### HasWallets

`func (o *PrincipalView) HasWallets() bool`

HasWallets returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


