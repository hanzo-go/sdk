# MarketplaceJobCounts

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Completed** | Pointer to **int64** | Completed is released jobs. | [optional] 
**Disputed** | Pointer to **int64** | Disputed is jobs ever disputed. | [optional] 

## Methods

### NewMarketplaceJobCounts

`func NewMarketplaceJobCounts() *MarketplaceJobCounts`

NewMarketplaceJobCounts instantiates a new MarketplaceJobCounts object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarketplaceJobCountsWithDefaults

`func NewMarketplaceJobCountsWithDefaults() *MarketplaceJobCounts`

NewMarketplaceJobCountsWithDefaults instantiates a new MarketplaceJobCounts object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompleted

`func (o *MarketplaceJobCounts) GetCompleted() int64`

GetCompleted returns the Completed field if non-nil, zero value otherwise.

### GetCompletedOk

`func (o *MarketplaceJobCounts) GetCompletedOk() (*int64, bool)`

GetCompletedOk returns a tuple with the Completed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompleted

`func (o *MarketplaceJobCounts) SetCompleted(v int64)`

SetCompleted sets Completed field to given value.

### HasCompleted

`func (o *MarketplaceJobCounts) HasCompleted() bool`

HasCompleted returns a boolean if a field has been set.

### GetDisputed

`func (o *MarketplaceJobCounts) GetDisputed() int64`

GetDisputed returns the Disputed field if non-nil, zero value otherwise.

### GetDisputedOk

`func (o *MarketplaceJobCounts) GetDisputedOk() (*int64, bool)`

GetDisputedOk returns a tuple with the Disputed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisputed

`func (o *MarketplaceJobCounts) SetDisputed(v int64)`

SetDisputed sets Disputed field to given value.

### HasDisputed

`func (o *MarketplaceJobCounts) HasDisputed() bool`

HasDisputed returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


