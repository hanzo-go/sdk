# AdminEnablementBoard

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | Pointer to [**[]AdminEnablementItem**](AdminEnablementItem.md) | Items is every item an operator has set a state on. An item nobody has touched is absent: it is generally available by default. | [optional] 

## Methods

### NewAdminEnablementBoard

`func NewAdminEnablementBoard() *AdminEnablementBoard`

NewAdminEnablementBoard instantiates a new AdminEnablementBoard object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdminEnablementBoardWithDefaults

`func NewAdminEnablementBoardWithDefaults() *AdminEnablementBoard`

NewAdminEnablementBoardWithDefaults instantiates a new AdminEnablementBoard object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *AdminEnablementBoard) GetItems() []AdminEnablementItem`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *AdminEnablementBoard) GetItemsOk() (*[]AdminEnablementItem, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *AdminEnablementBoard) SetItems(v []AdminEnablementItem)`

SetItems sets Items field to given value.

### HasItems

`func (o *AdminEnablementBoard) HasItems() bool`

HasItems returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


