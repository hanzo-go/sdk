# MlMlResourceList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | Pointer to [**[]MlMlResource**](MlMlResource.md) | Items is one entry per object, newest LAST (the Kubernetes list order). | [optional] 

## Methods

### NewMlMlResourceList

`func NewMlMlResourceList() *MlMlResourceList`

NewMlMlResourceList instantiates a new MlMlResourceList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMlMlResourceListWithDefaults

`func NewMlMlResourceListWithDefaults() *MlMlResourceList`

NewMlMlResourceListWithDefaults instantiates a new MlMlResourceList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *MlMlResourceList) GetItems() []MlMlResource`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *MlMlResourceList) GetItemsOk() (*[]MlMlResource, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *MlMlResourceList) SetItems(v []MlMlResource)`

SetItems sets Items field to given value.

### HasItems

`func (o *MlMlResourceList) HasItems() bool`

HasItems returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


