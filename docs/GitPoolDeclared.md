# GitPoolDeclared

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pool** | Pointer to [**GitPoolView**](GitPoolView.md) | Pool is the capacity that now exists. | [optional] 
**Secret** | Pointer to **string** | Secret is what a runner daemon presents to enter the pool. It is answered HERE AND NOWHERE ELSE — only a digest is stored — so a lost one is replaced by declaring the pool again. | [optional] 

## Methods

### NewGitPoolDeclared

`func NewGitPoolDeclared() *GitPoolDeclared`

NewGitPoolDeclared instantiates a new GitPoolDeclared object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGitPoolDeclaredWithDefaults

`func NewGitPoolDeclaredWithDefaults() *GitPoolDeclared`

NewGitPoolDeclaredWithDefaults instantiates a new GitPoolDeclared object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPool

`func (o *GitPoolDeclared) GetPool() GitPoolView`

GetPool returns the Pool field if non-nil, zero value otherwise.

### GetPoolOk

`func (o *GitPoolDeclared) GetPoolOk() (*GitPoolView, bool)`

GetPoolOk returns a tuple with the Pool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPool

`func (o *GitPoolDeclared) SetPool(v GitPoolView)`

SetPool sets Pool field to given value.

### HasPool

`func (o *GitPoolDeclared) HasPool() bool`

HasPool returns a boolean if a field has been set.

### GetSecret

`func (o *GitPoolDeclared) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *GitPoolDeclared) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *GitPoolDeclared) SetSecret(v string)`

SetSecret sets Secret field to given value.

### HasSecret

`func (o *GitPoolDeclared) HasSecret() bool`

HasSecret returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


