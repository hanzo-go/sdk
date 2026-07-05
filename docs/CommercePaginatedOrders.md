# CommercePaginatedOrders

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Page** | Pointer to **string** |  | [optional] 
**Display** | Pointer to **string** |  | [optional] 
**Count** | Pointer to **int32** |  | [optional] 
**Models** | Pointer to [**[]CommerceOrder**](CommerceOrder.md) |  | [optional] 

## Methods

### NewCommercePaginatedOrders

`func NewCommercePaginatedOrders() *CommercePaginatedOrders`

NewCommercePaginatedOrders instantiates a new CommercePaginatedOrders object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommercePaginatedOrdersWithDefaults

`func NewCommercePaginatedOrdersWithDefaults() *CommercePaginatedOrders`

NewCommercePaginatedOrdersWithDefaults instantiates a new CommercePaginatedOrders object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPage

`func (o *CommercePaginatedOrders) GetPage() string`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *CommercePaginatedOrders) GetPageOk() (*string, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *CommercePaginatedOrders) SetPage(v string)`

SetPage sets Page field to given value.

### HasPage

`func (o *CommercePaginatedOrders) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetDisplay

`func (o *CommercePaginatedOrders) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *CommercePaginatedOrders) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *CommercePaginatedOrders) SetDisplay(v string)`

SetDisplay sets Display field to given value.

### HasDisplay

`func (o *CommercePaginatedOrders) HasDisplay() bool`

HasDisplay returns a boolean if a field has been set.

### GetCount

`func (o *CommercePaginatedOrders) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *CommercePaginatedOrders) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *CommercePaginatedOrders) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *CommercePaginatedOrders) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetModels

`func (o *CommercePaginatedOrders) GetModels() []CommerceOrder`

GetModels returns the Models field if non-nil, zero value otherwise.

### GetModelsOk

`func (o *CommercePaginatedOrders) GetModelsOk() (*[]CommerceOrder, bool)`

GetModelsOk returns a tuple with the Models field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModels

`func (o *CommercePaginatedOrders) SetModels(v []CommerceOrder)`

SetModels sets Models field to given value.

### HasModels

`func (o *CommercePaginatedOrders) HasModels() bool`

HasModels returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


