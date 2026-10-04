# PrincipalList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AttemptedAt** | Pointer to **time.Time** | AttemptedAt is when the last attempt happened, successful or not. | [optional] 
**Designations** | Pointer to **int64** | Designations is how many entries the last successful load carried. | [optional] 
**Digest** | Pointer to **string** | Digest is the digest of the payload that produced Designations. | [optional] 
**Error** | Pointer to **string** | Err is why the last attempt failed, empty if it succeeded. | [optional] 
**LoadedAt** | Pointer to **time.Time** | LoadedAt is when that load happened. Zero means this list has never loaded, which is not the same as a publisher who designates nobody. | [optional] 
**Source** | Pointer to **string** | Source is the publisher. | [optional] 

## Methods

### NewPrincipalList

`func NewPrincipalList() *PrincipalList`

NewPrincipalList instantiates a new PrincipalList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrincipalListWithDefaults

`func NewPrincipalListWithDefaults() *PrincipalList`

NewPrincipalListWithDefaults instantiates a new PrincipalList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttemptedAt

`func (o *PrincipalList) GetAttemptedAt() time.Time`

GetAttemptedAt returns the AttemptedAt field if non-nil, zero value otherwise.

### GetAttemptedAtOk

`func (o *PrincipalList) GetAttemptedAtOk() (*time.Time, bool)`

GetAttemptedAtOk returns a tuple with the AttemptedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttemptedAt

`func (o *PrincipalList) SetAttemptedAt(v time.Time)`

SetAttemptedAt sets AttemptedAt field to given value.

### HasAttemptedAt

`func (o *PrincipalList) HasAttemptedAt() bool`

HasAttemptedAt returns a boolean if a field has been set.

### GetDesignations

`func (o *PrincipalList) GetDesignations() int64`

GetDesignations returns the Designations field if non-nil, zero value otherwise.

### GetDesignationsOk

`func (o *PrincipalList) GetDesignationsOk() (*int64, bool)`

GetDesignationsOk returns a tuple with the Designations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesignations

`func (o *PrincipalList) SetDesignations(v int64)`

SetDesignations sets Designations field to given value.

### HasDesignations

`func (o *PrincipalList) HasDesignations() bool`

HasDesignations returns a boolean if a field has been set.

### GetDigest

`func (o *PrincipalList) GetDigest() string`

GetDigest returns the Digest field if non-nil, zero value otherwise.

### GetDigestOk

`func (o *PrincipalList) GetDigestOk() (*string, bool)`

GetDigestOk returns a tuple with the Digest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDigest

`func (o *PrincipalList) SetDigest(v string)`

SetDigest sets Digest field to given value.

### HasDigest

`func (o *PrincipalList) HasDigest() bool`

HasDigest returns a boolean if a field has been set.

### GetError

`func (o *PrincipalList) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *PrincipalList) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *PrincipalList) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *PrincipalList) HasError() bool`

HasError returns a boolean if a field has been set.

### GetLoadedAt

`func (o *PrincipalList) GetLoadedAt() time.Time`

GetLoadedAt returns the LoadedAt field if non-nil, zero value otherwise.

### GetLoadedAtOk

`func (o *PrincipalList) GetLoadedAtOk() (*time.Time, bool)`

GetLoadedAtOk returns a tuple with the LoadedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadedAt

`func (o *PrincipalList) SetLoadedAt(v time.Time)`

SetLoadedAt sets LoadedAt field to given value.

### HasLoadedAt

`func (o *PrincipalList) HasLoadedAt() bool`

HasLoadedAt returns a boolean if a field has been set.

### GetSource

`func (o *PrincipalList) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *PrincipalList) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *PrincipalList) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *PrincipalList) HasSource() bool`

HasSource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


