# CloudInboxOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | Pointer to [**[]CloudInboxItem**](CloudInboxItem.md) | Items is every document still unsorted or in draft, newest first. | [optional] 

## Methods

### NewCloudInboxOut

`func NewCloudInboxOut() *CloudInboxOut`

NewCloudInboxOut instantiates a new CloudInboxOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudInboxOutWithDefaults

`func NewCloudInboxOutWithDefaults() *CloudInboxOut`

NewCloudInboxOutWithDefaults instantiates a new CloudInboxOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *CloudInboxOut) GetItems() []CloudInboxItem`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *CloudInboxOut) GetItemsOk() (*[]CloudInboxItem, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *CloudInboxOut) SetItems(v []CloudInboxItem)`

SetItems sets Items field to given value.

### HasItems

`func (o *CloudInboxOut) HasItems() bool`

HasItems returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


