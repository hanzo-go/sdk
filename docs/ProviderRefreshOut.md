# ProviderRefreshOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Connector** | Pointer to [**ProviderConnectionView**](ProviderConnectionView.md) | Connection is the connector with its new expiry. | [optional] 
**Refreshed** | Pointer to **bool** | Refreshed is always true — a failed rotation is an HTTP error. | [optional] 

## Methods

### NewProviderRefreshOut

`func NewProviderRefreshOut() *ProviderRefreshOut`

NewProviderRefreshOut instantiates a new ProviderRefreshOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProviderRefreshOutWithDefaults

`func NewProviderRefreshOutWithDefaults() *ProviderRefreshOut`

NewProviderRefreshOutWithDefaults instantiates a new ProviderRefreshOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnector

`func (o *ProviderRefreshOut) GetConnector() ProviderConnectionView`

GetConnector returns the Connector field if non-nil, zero value otherwise.

### GetConnectorOk

`func (o *ProviderRefreshOut) GetConnectorOk() (*ProviderConnectionView, bool)`

GetConnectorOk returns a tuple with the Connector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnector

`func (o *ProviderRefreshOut) SetConnector(v ProviderConnectionView)`

SetConnector sets Connector field to given value.

### HasConnector

`func (o *ProviderRefreshOut) HasConnector() bool`

HasConnector returns a boolean if a field has been set.

### GetRefreshed

`func (o *ProviderRefreshOut) GetRefreshed() bool`

GetRefreshed returns the Refreshed field if non-nil, zero value otherwise.

### GetRefreshedOk

`func (o *ProviderRefreshOut) GetRefreshedOk() (*bool, bool)`

GetRefreshedOk returns a tuple with the Refreshed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshed

`func (o *ProviderRefreshOut) SetRefreshed(v bool)`

SetRefreshed sets Refreshed field to given value.

### HasRefreshed

`func (o *ProviderRefreshOut) HasRefreshed() bool`

HasRefreshed returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


