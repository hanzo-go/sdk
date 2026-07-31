# CloudModelRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Model** | Pointer to **string** | Model is the model id, e.g. claude-opus-4. | [optional] 
**Pct** | Pointer to **float32** | Pct is this model&#39;s share of the window&#39;s returned spend, 0..100, one decimal. | [optional] 
**Provider** | Pointer to **string** | Provider is who served it. | [optional] 
**Requests** | Pointer to **int32** | Requests is how many calls went to this model. | [optional] 
**SpendCents** | Pointer to **int32** | SpendCents is what they cost, in cents. | [optional] 
**Tokens** | Pointer to **int32** | Tokens is prompt plus completion tokens over those calls. | [optional] 

## Methods

### NewCloudModelRow

`func NewCloudModelRow() *CloudModelRow`

NewCloudModelRow instantiates a new CloudModelRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudModelRowWithDefaults

`func NewCloudModelRowWithDefaults() *CloudModelRow`

NewCloudModelRowWithDefaults instantiates a new CloudModelRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetModel

`func (o *CloudModelRow) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *CloudModelRow) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *CloudModelRow) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *CloudModelRow) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetPct

`func (o *CloudModelRow) GetPct() float32`

GetPct returns the Pct field if non-nil, zero value otherwise.

### GetPctOk

`func (o *CloudModelRow) GetPctOk() (*float32, bool)`

GetPctOk returns a tuple with the Pct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPct

`func (o *CloudModelRow) SetPct(v float32)`

SetPct sets Pct field to given value.

### HasPct

`func (o *CloudModelRow) HasPct() bool`

HasPct returns a boolean if a field has been set.

### GetProvider

`func (o *CloudModelRow) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *CloudModelRow) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *CloudModelRow) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *CloudModelRow) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetRequests

`func (o *CloudModelRow) GetRequests() int32`

GetRequests returns the Requests field if non-nil, zero value otherwise.

### GetRequestsOk

`func (o *CloudModelRow) GetRequestsOk() (*int32, bool)`

GetRequestsOk returns a tuple with the Requests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequests

`func (o *CloudModelRow) SetRequests(v int32)`

SetRequests sets Requests field to given value.

### HasRequests

`func (o *CloudModelRow) HasRequests() bool`

HasRequests returns a boolean if a field has been set.

### GetSpendCents

`func (o *CloudModelRow) GetSpendCents() int32`

GetSpendCents returns the SpendCents field if non-nil, zero value otherwise.

### GetSpendCentsOk

`func (o *CloudModelRow) GetSpendCentsOk() (*int32, bool)`

GetSpendCentsOk returns a tuple with the SpendCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpendCents

`func (o *CloudModelRow) SetSpendCents(v int32)`

SetSpendCents sets SpendCents field to given value.

### HasSpendCents

`func (o *CloudModelRow) HasSpendCents() bool`

HasSpendCents returns a boolean if a field has been set.

### GetTokens

`func (o *CloudModelRow) GetTokens() int32`

GetTokens returns the Tokens field if non-nil, zero value otherwise.

### GetTokensOk

`func (o *CloudModelRow) GetTokensOk() (*int32, bool)`

GetTokensOk returns a tuple with the Tokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokens

`func (o *CloudModelRow) SetTokens(v int32)`

SetTokens sets Tokens field to given value.

### HasTokens

`func (o *CloudModelRow) HasTokens() bool`

HasTokens returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


