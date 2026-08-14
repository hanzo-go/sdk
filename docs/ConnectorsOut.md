# ConnectorsOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Connectors** | Pointer to [**[]ConnView**](ConnView.md) | Connectors is the caller&#39;s own set. Never null; [] when they have none. | [optional] 

## Methods

### NewConnectorsOut

`func NewConnectorsOut() *ConnectorsOut`

NewConnectorsOut instantiates a new ConnectorsOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectorsOutWithDefaults

`func NewConnectorsOutWithDefaults() *ConnectorsOut`

NewConnectorsOutWithDefaults instantiates a new ConnectorsOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnectors

`func (o *ConnectorsOut) GetConnectors() []ConnView`

GetConnectors returns the Connectors field if non-nil, zero value otherwise.

### GetConnectorsOk

`func (o *ConnectorsOut) GetConnectorsOk() (*[]ConnView, bool)`

GetConnectorsOk returns a tuple with the Connectors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectors

`func (o *ConnectorsOut) SetConnectors(v []ConnView)`

SetConnectors sets Connectors field to given value.

### HasConnectors

`func (o *ConnectorsOut) HasConnectors() bool`

HasConnectors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


