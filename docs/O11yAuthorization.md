# O11yAuthorization

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Credentials** | Pointer to **interface{}** |  | [optional] 
**CredentialsFile** | Pointer to **string** |  | [optional] 
**CredentialsRef** | Pointer to **string** | CredentialsRef is the name of the secret within the secret manager to use as credentials. | [optional] 
**Type** | Pointer to **string** |  | [optional] 

## Methods

### NewO11yAuthorization

`func NewO11yAuthorization() *O11yAuthorization`

NewO11yAuthorization instantiates a new O11yAuthorization object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yAuthorizationWithDefaults

`func NewO11yAuthorizationWithDefaults() *O11yAuthorization`

NewO11yAuthorizationWithDefaults instantiates a new O11yAuthorization object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCredentials

`func (o *O11yAuthorization) GetCredentials() interface{}`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *O11yAuthorization) GetCredentialsOk() (*interface{}, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *O11yAuthorization) SetCredentials(v interface{})`

SetCredentials sets Credentials field to given value.

### HasCredentials

`func (o *O11yAuthorization) HasCredentials() bool`

HasCredentials returns a boolean if a field has been set.

### SetCredentialsNil

`func (o *O11yAuthorization) SetCredentialsNil(b bool)`

 SetCredentialsNil sets the value for Credentials to be an explicit nil

### UnsetCredentials
`func (o *O11yAuthorization) UnsetCredentials()`

UnsetCredentials ensures that no value is present for Credentials, not even an explicit nil
### GetCredentialsFile

`func (o *O11yAuthorization) GetCredentialsFile() string`

GetCredentialsFile returns the CredentialsFile field if non-nil, zero value otherwise.

### GetCredentialsFileOk

`func (o *O11yAuthorization) GetCredentialsFileOk() (*string, bool)`

GetCredentialsFileOk returns a tuple with the CredentialsFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentialsFile

`func (o *O11yAuthorization) SetCredentialsFile(v string)`

SetCredentialsFile sets CredentialsFile field to given value.

### HasCredentialsFile

`func (o *O11yAuthorization) HasCredentialsFile() bool`

HasCredentialsFile returns a boolean if a field has been set.

### GetCredentialsRef

`func (o *O11yAuthorization) GetCredentialsRef() string`

GetCredentialsRef returns the CredentialsRef field if non-nil, zero value otherwise.

### GetCredentialsRefOk

`func (o *O11yAuthorization) GetCredentialsRefOk() (*string, bool)`

GetCredentialsRefOk returns a tuple with the CredentialsRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentialsRef

`func (o *O11yAuthorization) SetCredentialsRef(v string)`

SetCredentialsRef sets CredentialsRef field to given value.

### HasCredentialsRef

`func (o *O11yAuthorization) HasCredentialsRef() bool`

HasCredentialsRef returns a boolean if a field has been set.

### GetType

`func (o *O11yAuthorization) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *O11yAuthorization) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *O11yAuthorization) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *O11yAuthorization) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


