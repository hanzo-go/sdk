# MarketingSample

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Reason** | Pointer to **string** | Reason is why it was rejected. | [optional] 
**Row** | Pointer to **int64** | Row is the row&#39;s number in the file, the header not counted. | [optional] 

## Methods

### NewMarketingSample

`func NewMarketingSample() *MarketingSample`

NewMarketingSample instantiates a new MarketingSample object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarketingSampleWithDefaults

`func NewMarketingSampleWithDefaults() *MarketingSample`

NewMarketingSampleWithDefaults instantiates a new MarketingSample object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReason

`func (o *MarketingSample) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *MarketingSample) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *MarketingSample) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *MarketingSample) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetRow

`func (o *MarketingSample) GetRow() int64`

GetRow returns the Row field if non-nil, zero value otherwise.

### GetRowOk

`func (o *MarketingSample) GetRowOk() (*int64, bool)`

GetRowOk returns a tuple with the Row field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRow

`func (o *MarketingSample) SetRow(v int64)`

SetRow sets Row field to given value.

### HasRow

`func (o *MarketingSample) HasRow() bool`

HasRow returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


