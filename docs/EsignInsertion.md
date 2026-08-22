# EsignInsertion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FieldId** | Pointer to **string** | FieldID is the field that was filled. | [optional] 
**Inserted** | Pointer to **bool** | Inserted is true — the field now holds a value. Filling every field still leaves the document pending until the completion call. | [optional] 

## Methods

### NewEsignInsertion

`func NewEsignInsertion() *EsignInsertion`

NewEsignInsertion instantiates a new EsignInsertion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEsignInsertionWithDefaults

`func NewEsignInsertionWithDefaults() *EsignInsertion`

NewEsignInsertionWithDefaults instantiates a new EsignInsertion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFieldId

`func (o *EsignInsertion) GetFieldId() string`

GetFieldId returns the FieldId field if non-nil, zero value otherwise.

### GetFieldIdOk

`func (o *EsignInsertion) GetFieldIdOk() (*string, bool)`

GetFieldIdOk returns a tuple with the FieldId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldId

`func (o *EsignInsertion) SetFieldId(v string)`

SetFieldId sets FieldId field to given value.

### HasFieldId

`func (o *EsignInsertion) HasFieldId() bool`

HasFieldId returns a boolean if a field has been set.

### GetInserted

`func (o *EsignInsertion) GetInserted() bool`

GetInserted returns the Inserted field if non-nil, zero value otherwise.

### GetInsertedOk

`func (o *EsignInsertion) GetInsertedOk() (*bool, bool)`

GetInsertedOk returns a tuple with the Inserted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInserted

`func (o *EsignInsertion) SetInserted(v bool)`

SetInserted sets Inserted field to given value.

### HasInserted

`func (o *EsignInsertion) HasInserted() bool`

HasInserted returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


