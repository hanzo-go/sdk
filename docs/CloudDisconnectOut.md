# CloudDisconnectOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Disconnected** | Pointer to **bool** | Disconnected is always true — the org&#39;s secrets and connection row are gone. | [optional] 

## Methods

### NewCloudDisconnectOut

`func NewCloudDisconnectOut() *CloudDisconnectOut`

NewCloudDisconnectOut instantiates a new CloudDisconnectOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudDisconnectOutWithDefaults

`func NewCloudDisconnectOutWithDefaults() *CloudDisconnectOut`

NewCloudDisconnectOutWithDefaults instantiates a new CloudDisconnectOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisconnected

`func (o *CloudDisconnectOut) GetDisconnected() bool`

GetDisconnected returns the Disconnected field if non-nil, zero value otherwise.

### GetDisconnectedOk

`func (o *CloudDisconnectOut) GetDisconnectedOk() (*bool, bool)`

GetDisconnectedOk returns a tuple with the Disconnected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisconnected

`func (o *CloudDisconnectOut) SetDisconnected(v bool)`

SetDisconnected sets Disconnected field to given value.

### HasDisconnected

`func (o *CloudDisconnectOut) HasDisconnected() bool`

HasDisconnected returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


