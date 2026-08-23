# Detachment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Deleted** | Pointer to **bool** | Deleted is whether the method was actually removed. False with no error means it was already gone, which is a successful detach rather than a failure — a retry must not be an error. | [optional] 
**Id** | Pointer to **string** | ID is the method that was detached, echoed so a caller batching several can tell the answers apart. | [optional] 

## Methods

### NewDetachment

`func NewDetachment() *Detachment`

NewDetachment instantiates a new Detachment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDetachmentWithDefaults

`func NewDetachmentWithDefaults() *Detachment`

NewDetachmentWithDefaults instantiates a new Detachment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeleted

`func (o *Detachment) GetDeleted() bool`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *Detachment) GetDeletedOk() (*bool, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *Detachment) SetDeleted(v bool)`

SetDeleted sets Deleted field to given value.

### HasDeleted

`func (o *Detachment) HasDeleted() bool`

HasDeleted returns a boolean if a field has been set.

### GetId

`func (o *Detachment) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Detachment) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Detachment) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Detachment) HasId() bool`

HasId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


