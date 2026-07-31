# CloudAddItemsIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | ID is the annotation queue to add to, from the path. | [optional] 
**Items** | Pointer to [**[]CloudItemInput**](CloudItemInput.md) | Items are the objects to enqueue for review, 1–200 per request. Each names exactly one object. | [optional] 

## Methods

### NewCloudAddItemsIn

`func NewCloudAddItemsIn() *CloudAddItemsIn`

NewCloudAddItemsIn instantiates a new CloudAddItemsIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudAddItemsInWithDefaults

`func NewCloudAddItemsInWithDefaults() *CloudAddItemsIn`

NewCloudAddItemsInWithDefaults instantiates a new CloudAddItemsIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CloudAddItemsIn) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CloudAddItemsIn) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CloudAddItemsIn) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CloudAddItemsIn) HasId() bool`

HasId returns a boolean if a field has been set.

### GetItems

`func (o *CloudAddItemsIn) GetItems() []CloudItemInput`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *CloudAddItemsIn) GetItemsOk() (*[]CloudItemInput, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *CloudAddItemsIn) SetItems(v []CloudItemInput)`

SetItems sets Items field to given value.

### HasItems

`func (o *CloudAddItemsIn) HasItems() bool`

HasItems returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


