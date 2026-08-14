# ModelRoute

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CostInPerMillion** | Pointer to **float32** |  | [optional] 
**CostOutPerMillion** | Pointer to **float32** |  | [optional] 
**CreatedTime** | Pointer to **string** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**Fallback1Provider** | Pointer to **string** |  | [optional] 
**Fallback1Upstream** | Pointer to **string** |  | [optional] 
**Fallback2Provider** | Pointer to **string** |  | [optional] 
**Fallback2Upstream** | Pointer to **string** |  | [optional] 
**Hidden** | Pointer to **bool** |  | [optional] 
**InputPricePerMillion** | Pointer to **float32** |  | [optional] 
**ModelName** | Pointer to **string** |  | [optional] 
**OutputPricePerMillion** | Pointer to **float32** |  | [optional] 
**OwnedBy** | Pointer to **string** |  | [optional] 
**Owner** | Pointer to **string** |  | [optional] 
**Premium** | Pointer to **bool** |  | [optional] 
**Provider** | Pointer to **string** |  | [optional] 
**UpdatedTime** | Pointer to **string** |  | [optional] 
**Upstream** | Pointer to **string** |  | [optional] 

## Methods

### NewModelRoute

`func NewModelRoute() *ModelRoute`

NewModelRoute instantiates a new ModelRoute object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelRouteWithDefaults

`func NewModelRouteWithDefaults() *ModelRoute`

NewModelRouteWithDefaults instantiates a new ModelRoute object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCostInPerMillion

`func (o *ModelRoute) GetCostInPerMillion() float32`

GetCostInPerMillion returns the CostInPerMillion field if non-nil, zero value otherwise.

### GetCostInPerMillionOk

`func (o *ModelRoute) GetCostInPerMillionOk() (*float32, bool)`

GetCostInPerMillionOk returns a tuple with the CostInPerMillion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostInPerMillion

`func (o *ModelRoute) SetCostInPerMillion(v float32)`

SetCostInPerMillion sets CostInPerMillion field to given value.

### HasCostInPerMillion

`func (o *ModelRoute) HasCostInPerMillion() bool`

HasCostInPerMillion returns a boolean if a field has been set.

### GetCostOutPerMillion

`func (o *ModelRoute) GetCostOutPerMillion() float32`

GetCostOutPerMillion returns the CostOutPerMillion field if non-nil, zero value otherwise.

### GetCostOutPerMillionOk

`func (o *ModelRoute) GetCostOutPerMillionOk() (*float32, bool)`

GetCostOutPerMillionOk returns a tuple with the CostOutPerMillion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostOutPerMillion

`func (o *ModelRoute) SetCostOutPerMillion(v float32)`

SetCostOutPerMillion sets CostOutPerMillion field to given value.

### HasCostOutPerMillion

`func (o *ModelRoute) HasCostOutPerMillion() bool`

HasCostOutPerMillion returns a boolean if a field has been set.

### GetCreatedTime

`func (o *ModelRoute) GetCreatedTime() string`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *ModelRoute) GetCreatedTimeOk() (*string, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *ModelRoute) SetCreatedTime(v string)`

SetCreatedTime sets CreatedTime field to given value.

### HasCreatedTime

`func (o *ModelRoute) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.

### GetEnabled

`func (o *ModelRoute) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ModelRoute) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ModelRoute) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *ModelRoute) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetFallback1Provider

`func (o *ModelRoute) GetFallback1Provider() string`

GetFallback1Provider returns the Fallback1Provider field if non-nil, zero value otherwise.

### GetFallback1ProviderOk

`func (o *ModelRoute) GetFallback1ProviderOk() (*string, bool)`

GetFallback1ProviderOk returns a tuple with the Fallback1Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFallback1Provider

`func (o *ModelRoute) SetFallback1Provider(v string)`

SetFallback1Provider sets Fallback1Provider field to given value.

### HasFallback1Provider

`func (o *ModelRoute) HasFallback1Provider() bool`

HasFallback1Provider returns a boolean if a field has been set.

### GetFallback1Upstream

`func (o *ModelRoute) GetFallback1Upstream() string`

GetFallback1Upstream returns the Fallback1Upstream field if non-nil, zero value otherwise.

### GetFallback1UpstreamOk

`func (o *ModelRoute) GetFallback1UpstreamOk() (*string, bool)`

GetFallback1UpstreamOk returns a tuple with the Fallback1Upstream field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFallback1Upstream

`func (o *ModelRoute) SetFallback1Upstream(v string)`

SetFallback1Upstream sets Fallback1Upstream field to given value.

### HasFallback1Upstream

`func (o *ModelRoute) HasFallback1Upstream() bool`

HasFallback1Upstream returns a boolean if a field has been set.

### GetFallback2Provider

`func (o *ModelRoute) GetFallback2Provider() string`

GetFallback2Provider returns the Fallback2Provider field if non-nil, zero value otherwise.

### GetFallback2ProviderOk

`func (o *ModelRoute) GetFallback2ProviderOk() (*string, bool)`

GetFallback2ProviderOk returns a tuple with the Fallback2Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFallback2Provider

`func (o *ModelRoute) SetFallback2Provider(v string)`

SetFallback2Provider sets Fallback2Provider field to given value.

### HasFallback2Provider

`func (o *ModelRoute) HasFallback2Provider() bool`

HasFallback2Provider returns a boolean if a field has been set.

### GetFallback2Upstream

`func (o *ModelRoute) GetFallback2Upstream() string`

GetFallback2Upstream returns the Fallback2Upstream field if non-nil, zero value otherwise.

### GetFallback2UpstreamOk

`func (o *ModelRoute) GetFallback2UpstreamOk() (*string, bool)`

GetFallback2UpstreamOk returns a tuple with the Fallback2Upstream field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFallback2Upstream

`func (o *ModelRoute) SetFallback2Upstream(v string)`

SetFallback2Upstream sets Fallback2Upstream field to given value.

### HasFallback2Upstream

`func (o *ModelRoute) HasFallback2Upstream() bool`

HasFallback2Upstream returns a boolean if a field has been set.

### GetHidden

`func (o *ModelRoute) GetHidden() bool`

GetHidden returns the Hidden field if non-nil, zero value otherwise.

### GetHiddenOk

`func (o *ModelRoute) GetHiddenOk() (*bool, bool)`

GetHiddenOk returns a tuple with the Hidden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHidden

`func (o *ModelRoute) SetHidden(v bool)`

SetHidden sets Hidden field to given value.

### HasHidden

`func (o *ModelRoute) HasHidden() bool`

HasHidden returns a boolean if a field has been set.

### GetInputPricePerMillion

`func (o *ModelRoute) GetInputPricePerMillion() float32`

GetInputPricePerMillion returns the InputPricePerMillion field if non-nil, zero value otherwise.

### GetInputPricePerMillionOk

`func (o *ModelRoute) GetInputPricePerMillionOk() (*float32, bool)`

GetInputPricePerMillionOk returns a tuple with the InputPricePerMillion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputPricePerMillion

`func (o *ModelRoute) SetInputPricePerMillion(v float32)`

SetInputPricePerMillion sets InputPricePerMillion field to given value.

### HasInputPricePerMillion

`func (o *ModelRoute) HasInputPricePerMillion() bool`

HasInputPricePerMillion returns a boolean if a field has been set.

### GetModelName

`func (o *ModelRoute) GetModelName() string`

GetModelName returns the ModelName field if non-nil, zero value otherwise.

### GetModelNameOk

`func (o *ModelRoute) GetModelNameOk() (*string, bool)`

GetModelNameOk returns a tuple with the ModelName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelName

`func (o *ModelRoute) SetModelName(v string)`

SetModelName sets ModelName field to given value.

### HasModelName

`func (o *ModelRoute) HasModelName() bool`

HasModelName returns a boolean if a field has been set.

### GetOutputPricePerMillion

`func (o *ModelRoute) GetOutputPricePerMillion() float32`

GetOutputPricePerMillion returns the OutputPricePerMillion field if non-nil, zero value otherwise.

### GetOutputPricePerMillionOk

`func (o *ModelRoute) GetOutputPricePerMillionOk() (*float32, bool)`

GetOutputPricePerMillionOk returns a tuple with the OutputPricePerMillion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputPricePerMillion

`func (o *ModelRoute) SetOutputPricePerMillion(v float32)`

SetOutputPricePerMillion sets OutputPricePerMillion field to given value.

### HasOutputPricePerMillion

`func (o *ModelRoute) HasOutputPricePerMillion() bool`

HasOutputPricePerMillion returns a boolean if a field has been set.

### GetOwnedBy

`func (o *ModelRoute) GetOwnedBy() string`

GetOwnedBy returns the OwnedBy field if non-nil, zero value otherwise.

### GetOwnedByOk

`func (o *ModelRoute) GetOwnedByOk() (*string, bool)`

GetOwnedByOk returns a tuple with the OwnedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnedBy

`func (o *ModelRoute) SetOwnedBy(v string)`

SetOwnedBy sets OwnedBy field to given value.

### HasOwnedBy

`func (o *ModelRoute) HasOwnedBy() bool`

HasOwnedBy returns a boolean if a field has been set.

### GetOwner

`func (o *ModelRoute) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *ModelRoute) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *ModelRoute) SetOwner(v string)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *ModelRoute) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetPremium

`func (o *ModelRoute) GetPremium() bool`

GetPremium returns the Premium field if non-nil, zero value otherwise.

### GetPremiumOk

`func (o *ModelRoute) GetPremiumOk() (*bool, bool)`

GetPremiumOk returns a tuple with the Premium field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPremium

`func (o *ModelRoute) SetPremium(v bool)`

SetPremium sets Premium field to given value.

### HasPremium

`func (o *ModelRoute) HasPremium() bool`

HasPremium returns a boolean if a field has been set.

### GetProvider

`func (o *ModelRoute) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *ModelRoute) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *ModelRoute) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *ModelRoute) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetUpdatedTime

`func (o *ModelRoute) GetUpdatedTime() string`

GetUpdatedTime returns the UpdatedTime field if non-nil, zero value otherwise.

### GetUpdatedTimeOk

`func (o *ModelRoute) GetUpdatedTimeOk() (*string, bool)`

GetUpdatedTimeOk returns a tuple with the UpdatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedTime

`func (o *ModelRoute) SetUpdatedTime(v string)`

SetUpdatedTime sets UpdatedTime field to given value.

### HasUpdatedTime

`func (o *ModelRoute) HasUpdatedTime() bool`

HasUpdatedTime returns a boolean if a field has been set.

### GetUpstream

`func (o *ModelRoute) GetUpstream() string`

GetUpstream returns the Upstream field if non-nil, zero value otherwise.

### GetUpstreamOk

`func (o *ModelRoute) GetUpstreamOk() (*string, bool)`

GetUpstreamOk returns a tuple with the Upstream field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpstream

`func (o *ModelRoute) SetUpstream(v string)`

SetUpstream sets Upstream field to given value.

### HasUpstream

`func (o *ModelRoute) HasUpstream() bool`

HasUpstream returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


