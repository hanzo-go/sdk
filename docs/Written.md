# Written

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | ID is the record&#39;s id — the one supplied, or the one minted when the caller supplied none. | [optional] 
**Kind** | Pointer to **string** | Kind is the section written. | [optional] 
**Updated** | Pointer to **int64** | Updated is when it was written, unix milliseconds. | [optional] 

## Methods

### NewWritten

`func NewWritten() *Written`

NewWritten instantiates a new Written object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWrittenWithDefaults

`func NewWrittenWithDefaults() *Written`

NewWrittenWithDefaults instantiates a new Written object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Written) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Written) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Written) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Written) HasId() bool`

HasId returns a boolean if a field has been set.

### GetKind

`func (o *Written) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *Written) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *Written) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *Written) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetUpdated

`func (o *Written) GetUpdated() int64`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *Written) GetUpdatedOk() (*int64, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *Written) SetUpdated(v int64)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *Written) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


