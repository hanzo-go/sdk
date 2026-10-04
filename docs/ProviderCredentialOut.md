# ProviderCredentialOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Connected** | Pointer to **bool** | Connected is always true — a failed verification is a 400 and stores nothing. | [optional] 
**Connector** | Pointer to [**ProviderConnectionView**](ProviderConnectionView.md) | Connection is the connector as it now stands. | [optional] 

## Methods

### NewProviderCredentialOut

`func NewProviderCredentialOut() *ProviderCredentialOut`

NewProviderCredentialOut instantiates a new ProviderCredentialOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProviderCredentialOutWithDefaults

`func NewProviderCredentialOutWithDefaults() *ProviderCredentialOut`

NewProviderCredentialOutWithDefaults instantiates a new ProviderCredentialOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnected

`func (o *ProviderCredentialOut) GetConnected() bool`

GetConnected returns the Connected field if non-nil, zero value otherwise.

### GetConnectedOk

`func (o *ProviderCredentialOut) GetConnectedOk() (*bool, bool)`

GetConnectedOk returns a tuple with the Connected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnected

`func (o *ProviderCredentialOut) SetConnected(v bool)`

SetConnected sets Connected field to given value.

### HasConnected

`func (o *ProviderCredentialOut) HasConnected() bool`

HasConnected returns a boolean if a field has been set.

### GetConnector

`func (o *ProviderCredentialOut) GetConnector() ProviderConnectionView`

GetConnector returns the Connector field if non-nil, zero value otherwise.

### GetConnectorOk

`func (o *ProviderCredentialOut) GetConnectorOk() (*ProviderConnectionView, bool)`

GetConnectorOk returns a tuple with the Connector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnector

`func (o *ProviderCredentialOut) SetConnector(v ProviderConnectionView)`

SetConnector sets Connector field to given value.

### HasConnector

`func (o *ProviderCredentialOut) HasConnector() bool`

HasConnector returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


