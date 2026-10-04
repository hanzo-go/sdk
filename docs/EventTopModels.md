# EventTopModels

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Available** | Pointer to **bool** | Available is true whenever the ledger answered, including with no rows. | [optional] 
**Items** | Pointer to [**[]EventModelRow**](EventModelRow.md) | Items is the ranked models, highest spend first. | [optional] 
**Source** | Pointer to **string** | Source is the warehouse table the lens read. | [optional] 

## Methods

### NewEventTopModels

`func NewEventTopModels() *EventTopModels`

NewEventTopModels instantiates a new EventTopModels object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventTopModelsWithDefaults

`func NewEventTopModelsWithDefaults() *EventTopModels`

NewEventTopModelsWithDefaults instantiates a new EventTopModels object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvailable

`func (o *EventTopModels) GetAvailable() bool`

GetAvailable returns the Available field if non-nil, zero value otherwise.

### GetAvailableOk

`func (o *EventTopModels) GetAvailableOk() (*bool, bool)`

GetAvailableOk returns a tuple with the Available field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailable

`func (o *EventTopModels) SetAvailable(v bool)`

SetAvailable sets Available field to given value.

### HasAvailable

`func (o *EventTopModels) HasAvailable() bool`

HasAvailable returns a boolean if a field has been set.

### GetItems

`func (o *EventTopModels) GetItems() []EventModelRow`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *EventTopModels) GetItemsOk() (*[]EventModelRow, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *EventTopModels) SetItems(v []EventModelRow)`

SetItems sets Items field to given value.

### HasItems

`func (o *EventTopModels) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetSource

`func (o *EventTopModels) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *EventTopModels) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *EventTopModels) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *EventTopModels) HasSource() bool`

HasSource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


