# O11yAddItemsIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | ID is the annotation queue to add to, from the path. | [optional] 
**Items** | Pointer to [**[]O11yItemInput**](O11yItemInput.md) | Items are the objects to enqueue for review, 1–200 per request. Each names exactly one object. | [optional] 

## Methods

### NewO11yAddItemsIn

`func NewO11yAddItemsIn() *O11yAddItemsIn`

NewO11yAddItemsIn instantiates a new O11yAddItemsIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yAddItemsInWithDefaults

`func NewO11yAddItemsInWithDefaults() *O11yAddItemsIn`

NewO11yAddItemsInWithDefaults instantiates a new O11yAddItemsIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *O11yAddItemsIn) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *O11yAddItemsIn) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *O11yAddItemsIn) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *O11yAddItemsIn) HasId() bool`

HasId returns a boolean if a field has been set.

### GetItems

`func (o *O11yAddItemsIn) GetItems() []O11yItemInput`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *O11yAddItemsIn) GetItemsOk() (*[]O11yItemInput, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *O11yAddItemsIn) SetItems(v []O11yItemInput)`

SetItems sets Items field to given value.

### HasItems

`func (o *O11yAddItemsIn) HasItems() bool`

HasItems returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


