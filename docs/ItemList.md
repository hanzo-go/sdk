# ItemList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]ItemView**](ItemView.md) | Data is the examples of the one dataset named in the path, archived ones included, so the caller sees the whole set rather than what a run would use. | [optional] 

## Methods

### NewItemList

`func NewItemList() *ItemList`

NewItemList instantiates a new ItemList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewItemListWithDefaults

`func NewItemListWithDefaults() *ItemList`

NewItemListWithDefaults instantiates a new ItemList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ItemList) GetData() []ItemView`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ItemList) GetDataOk() (*[]ItemView, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ItemList) SetData(v []ItemView)`

SetData sets Data field to given value.

### HasData

`func (o *ItemList) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


