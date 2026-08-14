# TopModels

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Available** | Pointer to **bool** | Available is true whenever the ledger answered, including with no rows. | [optional] 
**Items** | Pointer to [**[]ModelRow**](ModelRow.md) | Items is the ranked models, highest spend first. | [optional] 
**Source** | Pointer to **string** | Source is the warehouse table the lens read. | [optional] 

## Methods

### NewTopModels

`func NewTopModels() *TopModels`

NewTopModels instantiates a new TopModels object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTopModelsWithDefaults

`func NewTopModelsWithDefaults() *TopModels`

NewTopModelsWithDefaults instantiates a new TopModels object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvailable

`func (o *TopModels) GetAvailable() bool`

GetAvailable returns the Available field if non-nil, zero value otherwise.

### GetAvailableOk

`func (o *TopModels) GetAvailableOk() (*bool, bool)`

GetAvailableOk returns a tuple with the Available field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailable

`func (o *TopModels) SetAvailable(v bool)`

SetAvailable sets Available field to given value.

### HasAvailable

`func (o *TopModels) HasAvailable() bool`

HasAvailable returns a boolean if a field has been set.

### GetItems

`func (o *TopModels) GetItems() []ModelRow`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *TopModels) GetItemsOk() (*[]ModelRow, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *TopModels) SetItems(v []ModelRow)`

SetItems sets Items field to given value.

### HasItems

`func (o *TopModels) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetSource

`func (o *TopModels) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *TopModels) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *TopModels) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *TopModels) HasSource() bool`

HasSource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


