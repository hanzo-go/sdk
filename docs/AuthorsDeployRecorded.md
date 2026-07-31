# AuthorsDeployRecorded

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Recorded** | Pointer to **bool** |  | [optional] 
**Created** | Pointer to **bool** |  | [optional] 
**Self** | Pointer to **bool** | Whether the deploying org is the author&#39;s own org (excluded from accrual). | [optional] 
**DeployId** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **int64** |  | [optional] 

## Methods

### NewAuthorsDeployRecorded

`func NewAuthorsDeployRecorded() *AuthorsDeployRecorded`

NewAuthorsDeployRecorded instantiates a new AuthorsDeployRecorded object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthorsDeployRecordedWithDefaults

`func NewAuthorsDeployRecordedWithDefaults() *AuthorsDeployRecorded`

NewAuthorsDeployRecordedWithDefaults instantiates a new AuthorsDeployRecorded object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRecorded

`func (o *AuthorsDeployRecorded) GetRecorded() bool`

GetRecorded returns the Recorded field if non-nil, zero value otherwise.

### GetRecordedOk

`func (o *AuthorsDeployRecorded) GetRecordedOk() (*bool, bool)`

GetRecordedOk returns a tuple with the Recorded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecorded

`func (o *AuthorsDeployRecorded) SetRecorded(v bool)`

SetRecorded sets Recorded field to given value.

### HasRecorded

`func (o *AuthorsDeployRecorded) HasRecorded() bool`

HasRecorded returns a boolean if a field has been set.

### GetCreated

`func (o *AuthorsDeployRecorded) GetCreated() bool`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *AuthorsDeployRecorded) GetCreatedOk() (*bool, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *AuthorsDeployRecorded) SetCreated(v bool)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *AuthorsDeployRecorded) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetSelf

`func (o *AuthorsDeployRecorded) GetSelf() bool`

GetSelf returns the Self field if non-nil, zero value otherwise.

### GetSelfOk

`func (o *AuthorsDeployRecorded) GetSelfOk() (*bool, bool)`

GetSelfOk returns a tuple with the Self field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelf

`func (o *AuthorsDeployRecorded) SetSelf(v bool)`

SetSelf sets Self field to given value.

### HasSelf

`func (o *AuthorsDeployRecorded) HasSelf() bool`

HasSelf returns a boolean if a field has been set.

### GetDeployId

`func (o *AuthorsDeployRecorded) GetDeployId() string`

GetDeployId returns the DeployId field if non-nil, zero value otherwise.

### GetDeployIdOk

`func (o *AuthorsDeployRecorded) GetDeployIdOk() (*string, bool)`

GetDeployIdOk returns a tuple with the DeployId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeployId

`func (o *AuthorsDeployRecorded) SetDeployId(v string)`

SetDeployId sets DeployId field to given value.

### HasDeployId

`func (o *AuthorsDeployRecorded) HasDeployId() bool`

HasDeployId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *AuthorsDeployRecorded) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AuthorsDeployRecorded) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AuthorsDeployRecorded) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *AuthorsDeployRecorded) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


