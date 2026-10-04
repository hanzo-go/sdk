# ProviderGithubUserDisconnectOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Disconnected** | Pointer to **bool** | Disconnected is always true: the sealed token and the connection are gone from this deployment. A custody failure is an HTTP error instead. | [optional] 
**Revoked** | Pointer to **bool** | Revoked is whether GitHub confirmed the token is dead. False means GitHub could not be told — the token lives until it expires or is revoked from the person&#39;s GitHub settings — or there was no token to revoke. | [optional] 

## Methods

### NewProviderGithubUserDisconnectOut

`func NewProviderGithubUserDisconnectOut() *ProviderGithubUserDisconnectOut`

NewProviderGithubUserDisconnectOut instantiates a new ProviderGithubUserDisconnectOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProviderGithubUserDisconnectOutWithDefaults

`func NewProviderGithubUserDisconnectOutWithDefaults() *ProviderGithubUserDisconnectOut`

NewProviderGithubUserDisconnectOutWithDefaults instantiates a new ProviderGithubUserDisconnectOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisconnected

`func (o *ProviderGithubUserDisconnectOut) GetDisconnected() bool`

GetDisconnected returns the Disconnected field if non-nil, zero value otherwise.

### GetDisconnectedOk

`func (o *ProviderGithubUserDisconnectOut) GetDisconnectedOk() (*bool, bool)`

GetDisconnectedOk returns a tuple with the Disconnected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisconnected

`func (o *ProviderGithubUserDisconnectOut) SetDisconnected(v bool)`

SetDisconnected sets Disconnected field to given value.

### HasDisconnected

`func (o *ProviderGithubUserDisconnectOut) HasDisconnected() bool`

HasDisconnected returns a boolean if a field has been set.

### GetRevoked

`func (o *ProviderGithubUserDisconnectOut) GetRevoked() bool`

GetRevoked returns the Revoked field if non-nil, zero value otherwise.

### GetRevokedOk

`func (o *ProviderGithubUserDisconnectOut) GetRevokedOk() (*bool, bool)`

GetRevokedOk returns a tuple with the Revoked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevoked

`func (o *ProviderGithubUserDisconnectOut) SetRevoked(v bool)`

SetRevoked sets Revoked field to given value.

### HasRevoked

`func (o *ProviderGithubUserDisconnectOut) HasRevoked() bool`

HasRevoked returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


