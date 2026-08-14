# AuthorizeOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthorizeUrl** | Pointer to **string** | AuthorizeURL is the provider consent (or bot deep-link) URL. | [optional] 

## Methods

### NewAuthorizeOut

`func NewAuthorizeOut() *AuthorizeOut`

NewAuthorizeOut instantiates a new AuthorizeOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthorizeOutWithDefaults

`func NewAuthorizeOutWithDefaults() *AuthorizeOut`

NewAuthorizeOutWithDefaults instantiates a new AuthorizeOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthorizeUrl

`func (o *AuthorizeOut) GetAuthorizeUrl() string`

GetAuthorizeUrl returns the AuthorizeUrl field if non-nil, zero value otherwise.

### GetAuthorizeUrlOk

`func (o *AuthorizeOut) GetAuthorizeUrlOk() (*string, bool)`

GetAuthorizeUrlOk returns a tuple with the AuthorizeUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizeUrl

`func (o *AuthorizeOut) SetAuthorizeUrl(v string)`

SetAuthorizeUrl sets AuthorizeUrl field to given value.

### HasAuthorizeUrl

`func (o *AuthorizeOut) HasAuthorizeUrl() bool`

HasAuthorizeUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


