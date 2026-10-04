# MarketplaceStep

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**At** | Pointer to **int64** | At is when, unix seconds. | [optional] 
**By** | Pointer to **string** | By is the org whose act it was, or \&quot;marketplace\&quot; for a clock that ran out. | [optional] 
**Status** | Pointer to **string** | Status is the state the job entered. | [optional] 

## Methods

### NewMarketplaceStep

`func NewMarketplaceStep() *MarketplaceStep`

NewMarketplaceStep instantiates a new MarketplaceStep object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarketplaceStepWithDefaults

`func NewMarketplaceStepWithDefaults() *MarketplaceStep`

NewMarketplaceStepWithDefaults instantiates a new MarketplaceStep object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAt

`func (o *MarketplaceStep) GetAt() int64`

GetAt returns the At field if non-nil, zero value otherwise.

### GetAtOk

`func (o *MarketplaceStep) GetAtOk() (*int64, bool)`

GetAtOk returns a tuple with the At field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAt

`func (o *MarketplaceStep) SetAt(v int64)`

SetAt sets At field to given value.

### HasAt

`func (o *MarketplaceStep) HasAt() bool`

HasAt returns a boolean if a field has been set.

### GetBy

`func (o *MarketplaceStep) GetBy() string`

GetBy returns the By field if non-nil, zero value otherwise.

### GetByOk

`func (o *MarketplaceStep) GetByOk() (*string, bool)`

GetByOk returns a tuple with the By field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBy

`func (o *MarketplaceStep) SetBy(v string)`

SetBy sets By field to given value.

### HasBy

`func (o *MarketplaceStep) HasBy() bool`

HasBy returns a boolean if a field has been set.

### GetStatus

`func (o *MarketplaceStep) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *MarketplaceStep) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *MarketplaceStep) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *MarketplaceStep) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


