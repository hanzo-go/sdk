# AuthorsVerifyRepoResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Repo** | Pointer to [**AuthorsRepoView**](AuthorsRepoView.md) |  | [optional] 
**Created** | Pointer to **bool** |  | [optional] 

## Methods

### NewAuthorsVerifyRepoResponse

`func NewAuthorsVerifyRepoResponse() *AuthorsVerifyRepoResponse`

NewAuthorsVerifyRepoResponse instantiates a new AuthorsVerifyRepoResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthorsVerifyRepoResponseWithDefaults

`func NewAuthorsVerifyRepoResponseWithDefaults() *AuthorsVerifyRepoResponse`

NewAuthorsVerifyRepoResponseWithDefaults instantiates a new AuthorsVerifyRepoResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRepo

`func (o *AuthorsVerifyRepoResponse) GetRepo() AuthorsRepoView`

GetRepo returns the Repo field if non-nil, zero value otherwise.

### GetRepoOk

`func (o *AuthorsVerifyRepoResponse) GetRepoOk() (*AuthorsRepoView, bool)`

GetRepoOk returns a tuple with the Repo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepo

`func (o *AuthorsVerifyRepoResponse) SetRepo(v AuthorsRepoView)`

SetRepo sets Repo field to given value.

### HasRepo

`func (o *AuthorsVerifyRepoResponse) HasRepo() bool`

HasRepo returns a boolean if a field has been set.

### GetCreated

`func (o *AuthorsVerifyRepoResponse) GetCreated() bool`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *AuthorsVerifyRepoResponse) GetCreatedOk() (*bool, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *AuthorsVerifyRepoResponse) SetCreated(v bool)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *AuthorsVerifyRepoResponse) HasCreated() bool`

HasCreated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


