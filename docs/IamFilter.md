# IamFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MaxResults** | Pointer to **int32** |  | [optional] 
**Supported** | Pointer to **bool** |  | [optional] 

## Methods

### NewIamFilter

`func NewIamFilter() *IamFilter`

NewIamFilter instantiates a new IamFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamFilterWithDefaults

`func NewIamFilterWithDefaults() *IamFilter`

NewIamFilterWithDefaults instantiates a new IamFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMaxResults

`func (o *IamFilter) GetMaxResults() int32`

GetMaxResults returns the MaxResults field if non-nil, zero value otherwise.

### GetMaxResultsOk

`func (o *IamFilter) GetMaxResultsOk() (*int32, bool)`

GetMaxResultsOk returns a tuple with the MaxResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxResults

`func (o *IamFilter) SetMaxResults(v int32)`

SetMaxResults sets MaxResults field to given value.

### HasMaxResults

`func (o *IamFilter) HasMaxResults() bool`

HasMaxResults returns a boolean if a field has been set.

### GetSupported

`func (o *IamFilter) GetSupported() bool`

GetSupported returns the Supported field if non-nil, zero value otherwise.

### GetSupportedOk

`func (o *IamFilter) GetSupportedOk() (*bool, bool)`

GetSupportedOk returns a tuple with the Supported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupported

`func (o *IamFilter) SetSupported(v bool)`

SetSupported sets Supported field to given value.

### HasSupported

`func (o *IamFilter) HasSupported() bool`

HasSupported returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


