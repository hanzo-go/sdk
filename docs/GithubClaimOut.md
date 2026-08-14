# GithubClaimOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Already** | Pointer to **[]string** | Already were bound before the call and are unchanged by it. | [optional] 
**Claimed** | Pointer to **[]string** | Claimed are the accounts this call bound. Never null; [] when none. | [optional] 

## Methods

### NewGithubClaimOut

`func NewGithubClaimOut() *GithubClaimOut`

NewGithubClaimOut instantiates a new GithubClaimOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGithubClaimOutWithDefaults

`func NewGithubClaimOutWithDefaults() *GithubClaimOut`

NewGithubClaimOutWithDefaults instantiates a new GithubClaimOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlready

`func (o *GithubClaimOut) GetAlready() []string`

GetAlready returns the Already field if non-nil, zero value otherwise.

### GetAlreadyOk

`func (o *GithubClaimOut) GetAlreadyOk() (*[]string, bool)`

GetAlreadyOk returns a tuple with the Already field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlready

`func (o *GithubClaimOut) SetAlready(v []string)`

SetAlready sets Already field to given value.

### HasAlready

`func (o *GithubClaimOut) HasAlready() bool`

HasAlready returns a boolean if a field has been set.

### GetClaimed

`func (o *GithubClaimOut) GetClaimed() []string`

GetClaimed returns the Claimed field if non-nil, zero value otherwise.

### GetClaimedOk

`func (o *GithubClaimOut) GetClaimedOk() (*[]string, bool)`

GetClaimedOk returns a tuple with the Claimed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClaimed

`func (o *GithubClaimOut) SetClaimed(v []string)`

SetClaimed sets Claimed field to given value.

### HasClaimed

`func (o *GithubClaimOut) HasClaimed() bool`

HasClaimed returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


