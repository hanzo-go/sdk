# RefreshOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Connector** | Pointer to [**ConnView**](ConnView.md) | Connection is the connector with its new expiry. | [optional] 
**Refreshed** | Pointer to **bool** | Refreshed is always true — a failed rotation is an HTTP error. | [optional] 

## Methods

### NewRefreshOut

`func NewRefreshOut() *RefreshOut`

NewRefreshOut instantiates a new RefreshOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRefreshOutWithDefaults

`func NewRefreshOutWithDefaults() *RefreshOut`

NewRefreshOutWithDefaults instantiates a new RefreshOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnector

`func (o *RefreshOut) GetConnector() ConnView`

GetConnector returns the Connector field if non-nil, zero value otherwise.

### GetConnectorOk

`func (o *RefreshOut) GetConnectorOk() (*ConnView, bool)`

GetConnectorOk returns a tuple with the Connector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnector

`func (o *RefreshOut) SetConnector(v ConnView)`

SetConnector sets Connector field to given value.

### HasConnector

`func (o *RefreshOut) HasConnector() bool`

HasConnector returns a boolean if a field has been set.

### GetRefreshed

`func (o *RefreshOut) GetRefreshed() bool`

GetRefreshed returns the Refreshed field if non-nil, zero value otherwise.

### GetRefreshedOk

`func (o *RefreshOut) GetRefreshedOk() (*bool, bool)`

GetRefreshedOk returns a tuple with the Refreshed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshed

`func (o *RefreshOut) SetRefreshed(v bool)`

SetRefreshed sets Refreshed field to given value.

### HasRefreshed

`func (o *RefreshOut) HasRefreshed() bool`

HasRefreshed returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


