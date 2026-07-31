# CloudCredentialOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Connected** | Pointer to **bool** | Connected is always true — a failed verification is a 400 and stores nothing. | [optional] 
**Connector** | Pointer to [**CloudConnView**](CloudConnView.md) | Connector is the connector as it now stands. | [optional] 

## Methods

### NewCloudCredentialOut

`func NewCloudCredentialOut() *CloudCredentialOut`

NewCloudCredentialOut instantiates a new CloudCredentialOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudCredentialOutWithDefaults

`func NewCloudCredentialOutWithDefaults() *CloudCredentialOut`

NewCloudCredentialOutWithDefaults instantiates a new CloudCredentialOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnected

`func (o *CloudCredentialOut) GetConnected() bool`

GetConnected returns the Connected field if non-nil, zero value otherwise.

### GetConnectedOk

`func (o *CloudCredentialOut) GetConnectedOk() (*bool, bool)`

GetConnectedOk returns a tuple with the Connected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnected

`func (o *CloudCredentialOut) SetConnected(v bool)`

SetConnected sets Connected field to given value.

### HasConnected

`func (o *CloudCredentialOut) HasConnected() bool`

HasConnected returns a boolean if a field has been set.

### GetConnector

`func (o *CloudCredentialOut) GetConnector() CloudConnView`

GetConnector returns the Connector field if non-nil, zero value otherwise.

### GetConnectorOk

`func (o *CloudCredentialOut) GetConnectorOk() (*CloudConnView, bool)`

GetConnectorOk returns a tuple with the Connector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnector

`func (o *CloudCredentialOut) SetConnector(v CloudConnView)`

SetConnector sets Connector field to given value.

### HasConnector

`func (o *CloudCredentialOut) HasConnector() bool`

HasConnector returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


