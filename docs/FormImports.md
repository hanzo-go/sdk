# FormImports

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Created** | Pointer to **int64** | Created counts the responses made new. | [optional] 
**Existing** | Pointer to **int64** | Existing counts the ones already held: imported before, or the same subscriber. | [optional] 
**Failed** | Pointer to [**[]FormImportFailure**](FormImportFailure.md) | Failed are the ones refused, with why. | [optional] 

## Methods

### NewFormImports

`func NewFormImports() *FormImports`

NewFormImports instantiates a new FormImports object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFormImportsWithDefaults

`func NewFormImportsWithDefaults() *FormImports`

NewFormImportsWithDefaults instantiates a new FormImports object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreated

`func (o *FormImports) GetCreated() int64`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *FormImports) GetCreatedOk() (*int64, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *FormImports) SetCreated(v int64)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *FormImports) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetExisting

`func (o *FormImports) GetExisting() int64`

GetExisting returns the Existing field if non-nil, zero value otherwise.

### GetExistingOk

`func (o *FormImports) GetExistingOk() (*int64, bool)`

GetExistingOk returns a tuple with the Existing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExisting

`func (o *FormImports) SetExisting(v int64)`

SetExisting sets Existing field to given value.

### HasExisting

`func (o *FormImports) HasExisting() bool`

HasExisting returns a boolean if a field has been set.

### GetFailed

`func (o *FormImports) GetFailed() []FormImportFailure`

GetFailed returns the Failed field if non-nil, zero value otherwise.

### GetFailedOk

`func (o *FormImports) GetFailedOk() (*[]FormImportFailure, bool)`

GetFailedOk returns a tuple with the Failed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailed

`func (o *FormImports) SetFailed(v []FormImportFailure)`

SetFailed sets Failed field to given value.

### HasFailed

`func (o *FormImports) HasFailed() bool`

HasFailed returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


