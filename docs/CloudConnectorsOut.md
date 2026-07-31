# CloudConnectorsOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Connectors** | Pointer to [**[]CloudConnView**](CloudConnView.md) | Connectors is the caller&#39;s own set. Never null; [] when they have none. | [optional] 

## Methods

### NewCloudConnectorsOut

`func NewCloudConnectorsOut() *CloudConnectorsOut`

NewCloudConnectorsOut instantiates a new CloudConnectorsOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudConnectorsOutWithDefaults

`func NewCloudConnectorsOutWithDefaults() *CloudConnectorsOut`

NewCloudConnectorsOutWithDefaults instantiates a new CloudConnectorsOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnectors

`func (o *CloudConnectorsOut) GetConnectors() []CloudConnView`

GetConnectors returns the Connectors field if non-nil, zero value otherwise.

### GetConnectorsOk

`func (o *CloudConnectorsOut) GetConnectorsOk() (*[]CloudConnView, bool)`

GetConnectorsOk returns a tuple with the Connectors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectors

`func (o *CloudConnectorsOut) SetConnectors(v []CloudConnView)`

SetConnectors sets Connectors field to given value.

### HasConnectors

`func (o *CloudConnectorsOut) HasConnectors() bool`

HasConnectors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


