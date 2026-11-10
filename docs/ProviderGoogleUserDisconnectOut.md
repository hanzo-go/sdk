# ProviderGoogleUserDisconnectOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Disconnected** | Pointer to **bool** | Disconnected is always true: the sealed tokens and the connection are gone from this deployment. A custody failure is an HTTP error instead. | [optional] 
**Revoked** | Pointer to **bool** | Revoked is whether Google confirmed the grant is dead. False means Google could not be told — the grant lives until the person removes it from their Google account — or there was nothing to revoke. | [optional] 

## Methods

### NewProviderGoogleUserDisconnectOut

`func NewProviderGoogleUserDisconnectOut() *ProviderGoogleUserDisconnectOut`

NewProviderGoogleUserDisconnectOut instantiates a new ProviderGoogleUserDisconnectOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProviderGoogleUserDisconnectOutWithDefaults

`func NewProviderGoogleUserDisconnectOutWithDefaults() *ProviderGoogleUserDisconnectOut`

NewProviderGoogleUserDisconnectOutWithDefaults instantiates a new ProviderGoogleUserDisconnectOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisconnected

`func (o *ProviderGoogleUserDisconnectOut) GetDisconnected() bool`

GetDisconnected returns the Disconnected field if non-nil, zero value otherwise.

### GetDisconnectedOk

`func (o *ProviderGoogleUserDisconnectOut) GetDisconnectedOk() (*bool, bool)`

GetDisconnectedOk returns a tuple with the Disconnected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisconnected

`func (o *ProviderGoogleUserDisconnectOut) SetDisconnected(v bool)`

SetDisconnected sets Disconnected field to given value.

### HasDisconnected

`func (o *ProviderGoogleUserDisconnectOut) HasDisconnected() bool`

HasDisconnected returns a boolean if a field has been set.

### GetRevoked

`func (o *ProviderGoogleUserDisconnectOut) GetRevoked() bool`

GetRevoked returns the Revoked field if non-nil, zero value otherwise.

### GetRevokedOk

`func (o *ProviderGoogleUserDisconnectOut) GetRevokedOk() (*bool, bool)`

GetRevokedOk returns a tuple with the Revoked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevoked

`func (o *ProviderGoogleUserDisconnectOut) SetRevoked(v bool)`

SetRevoked sets Revoked field to given value.

### HasRevoked

`func (o *ProviderGoogleUserDisconnectOut) HasRevoked() bool`

HasRevoked returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


