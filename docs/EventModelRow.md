# EventModelRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Model** | Pointer to **string** | Model is the model id, e.g. zen5-coder. | [optional] 
**Pct** | Pointer to **float64** | Pct is this model&#39;s share of the window&#39;s returned spend, 0..100, one decimal. | [optional] 
**Provider** | Pointer to **string** | Provider is who served it. | [optional] 
**Requests** | Pointer to **int64** | Requests is how many calls went to this model. | [optional] 
**SpendCents** | Pointer to **int64** | SpendCents is what they cost, in cents. | [optional] 
**Tokens** | Pointer to **int64** | Tokens is prompt plus completion tokens over those calls. | [optional] 

## Methods

### NewEventModelRow

`func NewEventModelRow() *EventModelRow`

NewEventModelRow instantiates a new EventModelRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventModelRowWithDefaults

`func NewEventModelRowWithDefaults() *EventModelRow`

NewEventModelRowWithDefaults instantiates a new EventModelRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetModel

`func (o *EventModelRow) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *EventModelRow) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *EventModelRow) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *EventModelRow) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetPct

`func (o *EventModelRow) GetPct() float64`

GetPct returns the Pct field if non-nil, zero value otherwise.

### GetPctOk

`func (o *EventModelRow) GetPctOk() (*float64, bool)`

GetPctOk returns a tuple with the Pct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPct

`func (o *EventModelRow) SetPct(v float64)`

SetPct sets Pct field to given value.

### HasPct

`func (o *EventModelRow) HasPct() bool`

HasPct returns a boolean if a field has been set.

### GetProvider

`func (o *EventModelRow) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *EventModelRow) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *EventModelRow) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *EventModelRow) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetRequests

`func (o *EventModelRow) GetRequests() int64`

GetRequests returns the Requests field if non-nil, zero value otherwise.

### GetRequestsOk

`func (o *EventModelRow) GetRequestsOk() (*int64, bool)`

GetRequestsOk returns a tuple with the Requests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequests

`func (o *EventModelRow) SetRequests(v int64)`

SetRequests sets Requests field to given value.

### HasRequests

`func (o *EventModelRow) HasRequests() bool`

HasRequests returns a boolean if a field has been set.

### GetSpendCents

`func (o *EventModelRow) GetSpendCents() int64`

GetSpendCents returns the SpendCents field if non-nil, zero value otherwise.

### GetSpendCentsOk

`func (o *EventModelRow) GetSpendCentsOk() (*int64, bool)`

GetSpendCentsOk returns a tuple with the SpendCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpendCents

`func (o *EventModelRow) SetSpendCents(v int64)`

SetSpendCents sets SpendCents field to given value.

### HasSpendCents

`func (o *EventModelRow) HasSpendCents() bool`

HasSpendCents returns a boolean if a field has been set.

### GetTokens

`func (o *EventModelRow) GetTokens() int64`

GetTokens returns the Tokens field if non-nil, zero value otherwise.

### GetTokensOk

`func (o *EventModelRow) GetTokensOk() (*int64, bool)`

GetTokensOk returns a tuple with the Tokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokens

`func (o *EventModelRow) SetTokens(v int64)`

SetTokens sets Tokens field to given value.

### HasTokens

`func (o *EventModelRow) HasTokens() bool`

HasTokens returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


