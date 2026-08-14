# ModelRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Model** | Pointer to **string** | Model is the model id, e.g. zen5-coder. | [optional] 
**Pct** | Pointer to **float32** | Pct is this model&#39;s share of the window&#39;s returned spend, 0..100, one decimal. | [optional] 
**Provider** | Pointer to **string** | Provider is who served it. | [optional] 
**Requests** | Pointer to **int32** | Requests is how many calls went to this model. | [optional] 
**SpendCents** | Pointer to **int32** | SpendCents is what they cost, in cents. | [optional] 
**Tokens** | Pointer to **int32** | Tokens is prompt plus completion tokens over those calls. | [optional] 

## Methods

### NewModelRow

`func NewModelRow() *ModelRow`

NewModelRow instantiates a new ModelRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelRowWithDefaults

`func NewModelRowWithDefaults() *ModelRow`

NewModelRowWithDefaults instantiates a new ModelRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetModel

`func (o *ModelRow) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *ModelRow) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *ModelRow) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *ModelRow) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetPct

`func (o *ModelRow) GetPct() float32`

GetPct returns the Pct field if non-nil, zero value otherwise.

### GetPctOk

`func (o *ModelRow) GetPctOk() (*float32, bool)`

GetPctOk returns a tuple with the Pct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPct

`func (o *ModelRow) SetPct(v float32)`

SetPct sets Pct field to given value.

### HasPct

`func (o *ModelRow) HasPct() bool`

HasPct returns a boolean if a field has been set.

### GetProvider

`func (o *ModelRow) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *ModelRow) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *ModelRow) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *ModelRow) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetRequests

`func (o *ModelRow) GetRequests() int32`

GetRequests returns the Requests field if non-nil, zero value otherwise.

### GetRequestsOk

`func (o *ModelRow) GetRequestsOk() (*int32, bool)`

GetRequestsOk returns a tuple with the Requests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequests

`func (o *ModelRow) SetRequests(v int32)`

SetRequests sets Requests field to given value.

### HasRequests

`func (o *ModelRow) HasRequests() bool`

HasRequests returns a boolean if a field has been set.

### GetSpendCents

`func (o *ModelRow) GetSpendCents() int32`

GetSpendCents returns the SpendCents field if non-nil, zero value otherwise.

### GetSpendCentsOk

`func (o *ModelRow) GetSpendCentsOk() (*int32, bool)`

GetSpendCentsOk returns a tuple with the SpendCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpendCents

`func (o *ModelRow) SetSpendCents(v int32)`

SetSpendCents sets SpendCents field to given value.

### HasSpendCents

`func (o *ModelRow) HasSpendCents() bool`

HasSpendCents returns a boolean if a field has been set.

### GetTokens

`func (o *ModelRow) GetTokens() int32`

GetTokens returns the Tokens field if non-nil, zero value otherwise.

### GetTokensOk

`func (o *ModelRow) GetTokensOk() (*int32, bool)`

GetTokensOk returns a tuple with the Tokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokens

`func (o *ModelRow) SetTokens(v int32)`

SetTokens sets Tokens field to given value.

### HasTokens

`func (o *ModelRow) HasTokens() bool`

HasTokens returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


