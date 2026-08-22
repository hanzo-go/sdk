# Dropped

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Deleted** | Pointer to **bool** | Deleted is always true; a record that was not there is a 404 instead. | [optional] 
**Id** | Pointer to **string** | ID is the record removed. | [optional] 
**Kind** | Pointer to **string** | Kind is the section. | [optional] 

## Methods

### NewDropped

`func NewDropped() *Dropped`

NewDropped instantiates a new Dropped object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDroppedWithDefaults

`func NewDroppedWithDefaults() *Dropped`

NewDroppedWithDefaults instantiates a new Dropped object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeleted

`func (o *Dropped) GetDeleted() bool`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *Dropped) GetDeletedOk() (*bool, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *Dropped) SetDeleted(v bool)`

SetDeleted sets Deleted field to given value.

### HasDeleted

`func (o *Dropped) HasDeleted() bool`

HasDeleted returns a boolean if a field has been set.

### GetId

`func (o *Dropped) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Dropped) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Dropped) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Dropped) HasId() bool`

HasId returns a boolean if a field has been set.

### GetKind

`func (o *Dropped) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *Dropped) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *Dropped) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *Dropped) HasKind() bool`

HasKind returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


