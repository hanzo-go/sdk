# ProviderGithubUserCompleteIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Grant** | Pointer to **string** | Grant is the id the callback handed the console in its redirect. It is single-use, lives ten minutes, and is readable only by the person who started the flow. | [optional] 

## Methods

### NewProviderGithubUserCompleteIn

`func NewProviderGithubUserCompleteIn() *ProviderGithubUserCompleteIn`

NewProviderGithubUserCompleteIn instantiates a new ProviderGithubUserCompleteIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProviderGithubUserCompleteInWithDefaults

`func NewProviderGithubUserCompleteInWithDefaults() *ProviderGithubUserCompleteIn`

NewProviderGithubUserCompleteInWithDefaults instantiates a new ProviderGithubUserCompleteIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGrant

`func (o *ProviderGithubUserCompleteIn) GetGrant() string`

GetGrant returns the Grant field if non-nil, zero value otherwise.

### GetGrantOk

`func (o *ProviderGithubUserCompleteIn) GetGrantOk() (*string, bool)`

GetGrantOk returns a tuple with the Grant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrant

`func (o *ProviderGithubUserCompleteIn) SetGrant(v string)`

SetGrant sets Grant field to given value.

### HasGrant

`func (o *ProviderGithubUserCompleteIn) HasGrant() bool`

HasGrant returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


