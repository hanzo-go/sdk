# CloudTopModels

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Available** | Pointer to **bool** | Available is true whenever the ledger answered, including with no rows. | [optional] 
**Items** | Pointer to [**[]CloudModelRow**](CloudModelRow.md) | Items is the ranked models, highest spend first. | [optional] 
**Source** | Pointer to **string** | Source is the warehouse table the lens read. | [optional] 

## Methods

### NewCloudTopModels

`func NewCloudTopModels() *CloudTopModels`

NewCloudTopModels instantiates a new CloudTopModels object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudTopModelsWithDefaults

`func NewCloudTopModelsWithDefaults() *CloudTopModels`

NewCloudTopModelsWithDefaults instantiates a new CloudTopModels object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvailable

`func (o *CloudTopModels) GetAvailable() bool`

GetAvailable returns the Available field if non-nil, zero value otherwise.

### GetAvailableOk

`func (o *CloudTopModels) GetAvailableOk() (*bool, bool)`

GetAvailableOk returns a tuple with the Available field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailable

`func (o *CloudTopModels) SetAvailable(v bool)`

SetAvailable sets Available field to given value.

### HasAvailable

`func (o *CloudTopModels) HasAvailable() bool`

HasAvailable returns a boolean if a field has been set.

### GetItems

`func (o *CloudTopModels) GetItems() []CloudModelRow`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *CloudTopModels) GetItemsOk() (*[]CloudModelRow, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *CloudTopModels) SetItems(v []CloudModelRow)`

SetItems sets Items field to given value.

### HasItems

`func (o *CloudTopModels) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetSource

`func (o *CloudTopModels) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *CloudTopModels) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *CloudTopModels) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *CloudTopModels) HasSource() bool`

HasSource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


