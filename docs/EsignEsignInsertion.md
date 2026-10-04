# EsignEsignInsertion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FieldId** | Pointer to **string** | FieldID is the field that was filled. | [optional] 
**Inserted** | Pointer to **bool** | Inserted is true — the field now holds a value. Filling every field still leaves the document pending until the completion call. | [optional] 

## Methods

### NewEsignEsignInsertion

`func NewEsignEsignInsertion() *EsignEsignInsertion`

NewEsignEsignInsertion instantiates a new EsignEsignInsertion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEsignEsignInsertionWithDefaults

`func NewEsignEsignInsertionWithDefaults() *EsignEsignInsertion`

NewEsignEsignInsertionWithDefaults instantiates a new EsignEsignInsertion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFieldId

`func (o *EsignEsignInsertion) GetFieldId() string`

GetFieldId returns the FieldId field if non-nil, zero value otherwise.

### GetFieldIdOk

`func (o *EsignEsignInsertion) GetFieldIdOk() (*string, bool)`

GetFieldIdOk returns a tuple with the FieldId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldId

`func (o *EsignEsignInsertion) SetFieldId(v string)`

SetFieldId sets FieldId field to given value.

### HasFieldId

`func (o *EsignEsignInsertion) HasFieldId() bool`

HasFieldId returns a boolean if a field has been set.

### GetInserted

`func (o *EsignEsignInsertion) GetInserted() bool`

GetInserted returns the Inserted field if non-nil, zero value otherwise.

### GetInsertedOk

`func (o *EsignEsignInsertion) GetInsertedOk() (*bool, bool)`

GetInsertedOk returns a tuple with the Inserted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInserted

`func (o *EsignEsignInsertion) SetInserted(v bool)`

SetInserted sets Inserted field to given value.

### HasInserted

`func (o *EsignEsignInsertion) HasInserted() bool`

HasInserted returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


