# MqListAccountConnections200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Connections** | Pointer to [**[]MqConnection**](MqConnection.md) |  | [optional] 
**Total** | Pointer to **int32** |  | [optional] 

## Methods

### NewMqListAccountConnections200Response

`func NewMqListAccountConnections200Response() *MqListAccountConnections200Response`

NewMqListAccountConnections200Response instantiates a new MqListAccountConnections200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMqListAccountConnections200ResponseWithDefaults

`func NewMqListAccountConnections200ResponseWithDefaults() *MqListAccountConnections200Response`

NewMqListAccountConnections200ResponseWithDefaults instantiates a new MqListAccountConnections200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnections

`func (o *MqListAccountConnections200Response) GetConnections() []MqConnection`

GetConnections returns the Connections field if non-nil, zero value otherwise.

### GetConnectionsOk

`func (o *MqListAccountConnections200Response) GetConnectionsOk() (*[]MqConnection, bool)`

GetConnectionsOk returns a tuple with the Connections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnections

`func (o *MqListAccountConnections200Response) SetConnections(v []MqConnection)`

SetConnections sets Connections field to given value.

### HasConnections

`func (o *MqListAccountConnections200Response) HasConnections() bool`

HasConnections returns a boolean if a field has been set.

### GetTotal

`func (o *MqListAccountConnections200Response) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *MqListAccountConnections200Response) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *MqListAccountConnections200Response) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *MqListAccountConnections200Response) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


