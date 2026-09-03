# ActivityPoint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CostCents** | Pointer to **int64** | CostCents is the day&#39;s spend in whole US cents. A series is only ever returned for a subject the caller is authorized to see, so this is never withheld: 0 means no spend that day. | [optional] 
**Day** | Pointer to **string** | Day is the UTC calendar day this point covers, \&quot;2006-01-02\&quot;. | [optional] 
**Requests** | Pointer to **int64** | Requests is the subject&#39;s request count on this day. 0 is a real, quiet day: the series is gap-filled, so every day in the range is present whether or not anything happened. | [optional] 
**Tokens** | Pointer to **int64** | Tokens is prompt+completion tokens on this day — normally the heatmap&#39;s intensity, scaled against ActivityTotals.MaxTokens. | [optional] 

## Methods

### NewActivityPoint

`func NewActivityPoint() *ActivityPoint`

NewActivityPoint instantiates a new ActivityPoint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActivityPointWithDefaults

`func NewActivityPointWithDefaults() *ActivityPoint`

NewActivityPointWithDefaults instantiates a new ActivityPoint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCostCents

`func (o *ActivityPoint) GetCostCents() int64`

GetCostCents returns the CostCents field if non-nil, zero value otherwise.

### GetCostCentsOk

`func (o *ActivityPoint) GetCostCentsOk() (*int64, bool)`

GetCostCentsOk returns a tuple with the CostCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostCents

`func (o *ActivityPoint) SetCostCents(v int64)`

SetCostCents sets CostCents field to given value.

### HasCostCents

`func (o *ActivityPoint) HasCostCents() bool`

HasCostCents returns a boolean if a field has been set.

### GetDay

`func (o *ActivityPoint) GetDay() string`

GetDay returns the Day field if non-nil, zero value otherwise.

### GetDayOk

`func (o *ActivityPoint) GetDayOk() (*string, bool)`

GetDayOk returns a tuple with the Day field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDay

`func (o *ActivityPoint) SetDay(v string)`

SetDay sets Day field to given value.

### HasDay

`func (o *ActivityPoint) HasDay() bool`

HasDay returns a boolean if a field has been set.

### GetRequests

`func (o *ActivityPoint) GetRequests() int64`

GetRequests returns the Requests field if non-nil, zero value otherwise.

### GetRequestsOk

`func (o *ActivityPoint) GetRequestsOk() (*int64, bool)`

GetRequestsOk returns a tuple with the Requests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequests

`func (o *ActivityPoint) SetRequests(v int64)`

SetRequests sets Requests field to given value.

### HasRequests

`func (o *ActivityPoint) HasRequests() bool`

HasRequests returns a boolean if a field has been set.

### GetTokens

`func (o *ActivityPoint) GetTokens() int64`

GetTokens returns the Tokens field if non-nil, zero value otherwise.

### GetTokensOk

`func (o *ActivityPoint) GetTokensOk() (*int64, bool)`

GetTokensOk returns a tuple with the Tokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokens

`func (o *ActivityPoint) SetTokens(v int64)`

SetTokens sets Tokens field to given value.

### HasTokens

`func (o *ActivityPoint) HasTokens() bool`

HasTokens returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


