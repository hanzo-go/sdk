# CloudSelfRank

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CostCents** | Pointer to **int32** |  | [optional] 
**Handle** | Pointer to **string** |  | [optional] 
**Listed** | Pointer to **bool** | is the caller publicly listed (opted in) | [optional] 
**Metric** | Pointer to **int32** |  | [optional] 
**OfTotal** | Pointer to **int32** |  | [optional] 
**Rank** | Pointer to **int32** |  | [optional] 
**Ranked** | Pointer to **bool** |  | [optional] 
**Requests** | Pointer to **int32** |  | [optional] 
**Tokens** | Pointer to **int32** |  | [optional] 

## Methods

### NewCloudSelfRank

`func NewCloudSelfRank() *CloudSelfRank`

NewCloudSelfRank instantiates a new CloudSelfRank object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudSelfRankWithDefaults

`func NewCloudSelfRankWithDefaults() *CloudSelfRank`

NewCloudSelfRankWithDefaults instantiates a new CloudSelfRank object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCostCents

`func (o *CloudSelfRank) GetCostCents() int32`

GetCostCents returns the CostCents field if non-nil, zero value otherwise.

### GetCostCentsOk

`func (o *CloudSelfRank) GetCostCentsOk() (*int32, bool)`

GetCostCentsOk returns a tuple with the CostCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostCents

`func (o *CloudSelfRank) SetCostCents(v int32)`

SetCostCents sets CostCents field to given value.

### HasCostCents

`func (o *CloudSelfRank) HasCostCents() bool`

HasCostCents returns a boolean if a field has been set.

### GetHandle

`func (o *CloudSelfRank) GetHandle() string`

GetHandle returns the Handle field if non-nil, zero value otherwise.

### GetHandleOk

`func (o *CloudSelfRank) GetHandleOk() (*string, bool)`

GetHandleOk returns a tuple with the Handle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHandle

`func (o *CloudSelfRank) SetHandle(v string)`

SetHandle sets Handle field to given value.

### HasHandle

`func (o *CloudSelfRank) HasHandle() bool`

HasHandle returns a boolean if a field has been set.

### GetListed

`func (o *CloudSelfRank) GetListed() bool`

GetListed returns the Listed field if non-nil, zero value otherwise.

### GetListedOk

`func (o *CloudSelfRank) GetListedOk() (*bool, bool)`

GetListedOk returns a tuple with the Listed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListed

`func (o *CloudSelfRank) SetListed(v bool)`

SetListed sets Listed field to given value.

### HasListed

`func (o *CloudSelfRank) HasListed() bool`

HasListed returns a boolean if a field has been set.

### GetMetric

`func (o *CloudSelfRank) GetMetric() int32`

GetMetric returns the Metric field if non-nil, zero value otherwise.

### GetMetricOk

`func (o *CloudSelfRank) GetMetricOk() (*int32, bool)`

GetMetricOk returns a tuple with the Metric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetric

`func (o *CloudSelfRank) SetMetric(v int32)`

SetMetric sets Metric field to given value.

### HasMetric

`func (o *CloudSelfRank) HasMetric() bool`

HasMetric returns a boolean if a field has been set.

### GetOfTotal

`func (o *CloudSelfRank) GetOfTotal() int32`

GetOfTotal returns the OfTotal field if non-nil, zero value otherwise.

### GetOfTotalOk

`func (o *CloudSelfRank) GetOfTotalOk() (*int32, bool)`

GetOfTotalOk returns a tuple with the OfTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfTotal

`func (o *CloudSelfRank) SetOfTotal(v int32)`

SetOfTotal sets OfTotal field to given value.

### HasOfTotal

`func (o *CloudSelfRank) HasOfTotal() bool`

HasOfTotal returns a boolean if a field has been set.

### GetRank

`func (o *CloudSelfRank) GetRank() int32`

GetRank returns the Rank field if non-nil, zero value otherwise.

### GetRankOk

`func (o *CloudSelfRank) GetRankOk() (*int32, bool)`

GetRankOk returns a tuple with the Rank field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRank

`func (o *CloudSelfRank) SetRank(v int32)`

SetRank sets Rank field to given value.

### HasRank

`func (o *CloudSelfRank) HasRank() bool`

HasRank returns a boolean if a field has been set.

### GetRanked

`func (o *CloudSelfRank) GetRanked() bool`

GetRanked returns the Ranked field if non-nil, zero value otherwise.

### GetRankedOk

`func (o *CloudSelfRank) GetRankedOk() (*bool, bool)`

GetRankedOk returns a tuple with the Ranked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRanked

`func (o *CloudSelfRank) SetRanked(v bool)`

SetRanked sets Ranked field to given value.

### HasRanked

`func (o *CloudSelfRank) HasRanked() bool`

HasRanked returns a boolean if a field has been set.

### GetRequests

`func (o *CloudSelfRank) GetRequests() int32`

GetRequests returns the Requests field if non-nil, zero value otherwise.

### GetRequestsOk

`func (o *CloudSelfRank) GetRequestsOk() (*int32, bool)`

GetRequestsOk returns a tuple with the Requests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequests

`func (o *CloudSelfRank) SetRequests(v int32)`

SetRequests sets Requests field to given value.

### HasRequests

`func (o *CloudSelfRank) HasRequests() bool`

HasRequests returns a boolean if a field has been set.

### GetTokens

`func (o *CloudSelfRank) GetTokens() int32`

GetTokens returns the Tokens field if non-nil, zero value otherwise.

### GetTokensOk

`func (o *CloudSelfRank) GetTokensOk() (*int32, bool)`

GetTokensOk returns a tuple with the Tokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokens

`func (o *CloudSelfRank) SetTokens(v int32)`

SetTokens sets Tokens field to given value.

### HasTokens

`func (o *CloudSelfRank) HasTokens() bool`

HasTokens returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


