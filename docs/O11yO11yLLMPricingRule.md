# O11yO11yLLMPricingRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **time.Time** | CreatedAt is when the rule was stored. | [optional] 
**CreatedBy** | Pointer to **string** | CreatedBy is who created the rule. | [optional] 
**Enabled** | Pointer to **bool** | Enabled says whether the rule is on. | [optional] 
**Id** | Pointer to **string** | ID is the rule&#39;s id. | [optional] 
**IsOverride** | Pointer to **bool** | IsOverride marks the rule user-pinned; when true the sync job skips it. | [optional] 
**ModelName** | Pointer to **string** | Model is the model the rule prices. | [optional] 
**ModelPattern** | Pointer to **[]string** | ModelPattern are the model-name globs the rule matches. | [optional] 
**OrgId** | Pointer to **string** | OrgID is the org the rule belongs to. | [optional] 
**Pricing** | Pointer to [**O11yO11yLLMRulePricing**](O11yO11yLLMRulePricing.md) | Pricing is the per-unit cost. | [optional] 
**Provider** | Pointer to **string** | Provider is the model&#39;s provider. | [optional] 
**SourceId** | Pointer to **string** | SourceID is the upstream source the rule was synced from, when synced. | [optional] 
**SyncedAt** | Pointer to **time.Time** | SyncedAt is when the rule was last synced, when it is synced. | [optional] 
**Unit** | Pointer to **string** | Unit is the pricing unit, e.g. per_million_tokens. | [optional] 
**UpdatedAt** | Pointer to **time.Time** | UpdatedAt is when the rule last changed. | [optional] 
**UpdatedBy** | Pointer to **string** | UpdatedBy is who last changed it. | [optional] 

## Methods

### NewO11yO11yLLMPricingRule

`func NewO11yO11yLLMPricingRule() *O11yO11yLLMPricingRule`

NewO11yO11yLLMPricingRule instantiates a new O11yO11yLLMPricingRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yO11yLLMPricingRuleWithDefaults

`func NewO11yO11yLLMPricingRuleWithDefaults() *O11yO11yLLMPricingRule`

NewO11yO11yLLMPricingRuleWithDefaults instantiates a new O11yO11yLLMPricingRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *O11yO11yLLMPricingRule) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *O11yO11yLLMPricingRule) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *O11yO11yLLMPricingRule) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *O11yO11yLLMPricingRule) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCreatedBy

`func (o *O11yO11yLLMPricingRule) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *O11yO11yLLMPricingRule) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *O11yO11yLLMPricingRule) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *O11yO11yLLMPricingRule) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetEnabled

`func (o *O11yO11yLLMPricingRule) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *O11yO11yLLMPricingRule) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *O11yO11yLLMPricingRule) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *O11yO11yLLMPricingRule) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetId

`func (o *O11yO11yLLMPricingRule) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *O11yO11yLLMPricingRule) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *O11yO11yLLMPricingRule) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *O11yO11yLLMPricingRule) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsOverride

`func (o *O11yO11yLLMPricingRule) GetIsOverride() bool`

GetIsOverride returns the IsOverride field if non-nil, zero value otherwise.

### GetIsOverrideOk

`func (o *O11yO11yLLMPricingRule) GetIsOverrideOk() (*bool, bool)`

GetIsOverrideOk returns a tuple with the IsOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsOverride

`func (o *O11yO11yLLMPricingRule) SetIsOverride(v bool)`

SetIsOverride sets IsOverride field to given value.

### HasIsOverride

`func (o *O11yO11yLLMPricingRule) HasIsOverride() bool`

HasIsOverride returns a boolean if a field has been set.

### GetModelName

`func (o *O11yO11yLLMPricingRule) GetModelName() string`

GetModelName returns the ModelName field if non-nil, zero value otherwise.

### GetModelNameOk

`func (o *O11yO11yLLMPricingRule) GetModelNameOk() (*string, bool)`

GetModelNameOk returns a tuple with the ModelName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelName

`func (o *O11yO11yLLMPricingRule) SetModelName(v string)`

SetModelName sets ModelName field to given value.

### HasModelName

`func (o *O11yO11yLLMPricingRule) HasModelName() bool`

HasModelName returns a boolean if a field has been set.

### GetModelPattern

`func (o *O11yO11yLLMPricingRule) GetModelPattern() []string`

GetModelPattern returns the ModelPattern field if non-nil, zero value otherwise.

### GetModelPatternOk

`func (o *O11yO11yLLMPricingRule) GetModelPatternOk() (*[]string, bool)`

GetModelPatternOk returns a tuple with the ModelPattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelPattern

`func (o *O11yO11yLLMPricingRule) SetModelPattern(v []string)`

SetModelPattern sets ModelPattern field to given value.

### HasModelPattern

`func (o *O11yO11yLLMPricingRule) HasModelPattern() bool`

HasModelPattern returns a boolean if a field has been set.

### GetOrgId

`func (o *O11yO11yLLMPricingRule) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *O11yO11yLLMPricingRule) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *O11yO11yLLMPricingRule) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *O11yO11yLLMPricingRule) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### GetPricing

`func (o *O11yO11yLLMPricingRule) GetPricing() O11yO11yLLMRulePricing`

GetPricing returns the Pricing field if non-nil, zero value otherwise.

### GetPricingOk

`func (o *O11yO11yLLMPricingRule) GetPricingOk() (*O11yO11yLLMRulePricing, bool)`

GetPricingOk returns a tuple with the Pricing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPricing

`func (o *O11yO11yLLMPricingRule) SetPricing(v O11yO11yLLMRulePricing)`

SetPricing sets Pricing field to given value.

### HasPricing

`func (o *O11yO11yLLMPricingRule) HasPricing() bool`

HasPricing returns a boolean if a field has been set.

### GetProvider

`func (o *O11yO11yLLMPricingRule) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *O11yO11yLLMPricingRule) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *O11yO11yLLMPricingRule) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *O11yO11yLLMPricingRule) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetSourceId

`func (o *O11yO11yLLMPricingRule) GetSourceId() string`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *O11yO11yLLMPricingRule) GetSourceIdOk() (*string, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *O11yO11yLLMPricingRule) SetSourceId(v string)`

SetSourceId sets SourceId field to given value.

### HasSourceId

`func (o *O11yO11yLLMPricingRule) HasSourceId() bool`

HasSourceId returns a boolean if a field has been set.

### GetSyncedAt

`func (o *O11yO11yLLMPricingRule) GetSyncedAt() time.Time`

GetSyncedAt returns the SyncedAt field if non-nil, zero value otherwise.

### GetSyncedAtOk

`func (o *O11yO11yLLMPricingRule) GetSyncedAtOk() (*time.Time, bool)`

GetSyncedAtOk returns a tuple with the SyncedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncedAt

`func (o *O11yO11yLLMPricingRule) SetSyncedAt(v time.Time)`

SetSyncedAt sets SyncedAt field to given value.

### HasSyncedAt

`func (o *O11yO11yLLMPricingRule) HasSyncedAt() bool`

HasSyncedAt returns a boolean if a field has been set.

### GetUnit

`func (o *O11yO11yLLMPricingRule) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *O11yO11yLLMPricingRule) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *O11yO11yLLMPricingRule) SetUnit(v string)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *O11yO11yLLMPricingRule) HasUnit() bool`

HasUnit returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *O11yO11yLLMPricingRule) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *O11yO11yLLMPricingRule) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *O11yO11yLLMPricingRule) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *O11yO11yLLMPricingRule) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUpdatedBy

`func (o *O11yO11yLLMPricingRule) GetUpdatedBy() string`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *O11yO11yLLMPricingRule) GetUpdatedByOk() (*string, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *O11yO11yLLMPricingRule) SetUpdatedBy(v string)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *O11yO11yLLMPricingRule) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


