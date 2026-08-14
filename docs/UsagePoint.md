# UsagePoint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Date** | Pointer to **string** |  | [optional] 
**Requests** | Pointer to **int32** |  | [optional] 
**SpendCents** | Pointer to **int32** |  | [optional] 
**Tokens** | Pointer to **int32** |  | [optional] 

## Methods

### NewUsagePoint

`func NewUsagePoint() *UsagePoint`

NewUsagePoint instantiates a new UsagePoint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsagePointWithDefaults

`func NewUsagePointWithDefaults() *UsagePoint`

NewUsagePointWithDefaults instantiates a new UsagePoint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDate

`func (o *UsagePoint) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *UsagePoint) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *UsagePoint) SetDate(v string)`

SetDate sets Date field to given value.

### HasDate

`func (o *UsagePoint) HasDate() bool`

HasDate returns a boolean if a field has been set.

### GetRequests

`func (o *UsagePoint) GetRequests() int32`

GetRequests returns the Requests field if non-nil, zero value otherwise.

### GetRequestsOk

`func (o *UsagePoint) GetRequestsOk() (*int32, bool)`

GetRequestsOk returns a tuple with the Requests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequests

`func (o *UsagePoint) SetRequests(v int32)`

SetRequests sets Requests field to given value.

### HasRequests

`func (o *UsagePoint) HasRequests() bool`

HasRequests returns a boolean if a field has been set.

### GetSpendCents

`func (o *UsagePoint) GetSpendCents() int32`

GetSpendCents returns the SpendCents field if non-nil, zero value otherwise.

### GetSpendCentsOk

`func (o *UsagePoint) GetSpendCentsOk() (*int32, bool)`

GetSpendCentsOk returns a tuple with the SpendCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpendCents

`func (o *UsagePoint) SetSpendCents(v int32)`

SetSpendCents sets SpendCents field to given value.

### HasSpendCents

`func (o *UsagePoint) HasSpendCents() bool`

HasSpendCents returns a boolean if a field has been set.

### GetTokens

`func (o *UsagePoint) GetTokens() int32`

GetTokens returns the Tokens field if non-nil, zero value otherwise.

### GetTokensOk

`func (o *UsagePoint) GetTokensOk() (*int32, bool)`

GetTokensOk returns a tuple with the Tokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokens

`func (o *UsagePoint) SetTokens(v int32)`

SetTokens sets Tokens field to given value.

### HasTokens

`func (o *UsagePoint) HasTokens() bool`

HasTokens returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


