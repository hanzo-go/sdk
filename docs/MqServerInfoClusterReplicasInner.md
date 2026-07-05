# MqServerInfoClusterReplicasInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Current** | Pointer to **bool** |  | [optional] 
**Active** | Pointer to **string** | Time since last activity. | [optional] 
**Lag** | Pointer to **int32** | Replication lag in operations. | [optional] 

## Methods

### NewMqServerInfoClusterReplicasInner

`func NewMqServerInfoClusterReplicasInner() *MqServerInfoClusterReplicasInner`

NewMqServerInfoClusterReplicasInner instantiates a new MqServerInfoClusterReplicasInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMqServerInfoClusterReplicasInnerWithDefaults

`func NewMqServerInfoClusterReplicasInnerWithDefaults() *MqServerInfoClusterReplicasInner`

NewMqServerInfoClusterReplicasInnerWithDefaults instantiates a new MqServerInfoClusterReplicasInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *MqServerInfoClusterReplicasInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MqServerInfoClusterReplicasInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MqServerInfoClusterReplicasInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MqServerInfoClusterReplicasInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCurrent

`func (o *MqServerInfoClusterReplicasInner) GetCurrent() bool`

GetCurrent returns the Current field if non-nil, zero value otherwise.

### GetCurrentOk

`func (o *MqServerInfoClusterReplicasInner) GetCurrentOk() (*bool, bool)`

GetCurrentOk returns a tuple with the Current field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrent

`func (o *MqServerInfoClusterReplicasInner) SetCurrent(v bool)`

SetCurrent sets Current field to given value.

### HasCurrent

`func (o *MqServerInfoClusterReplicasInner) HasCurrent() bool`

HasCurrent returns a boolean if a field has been set.

### GetActive

`func (o *MqServerInfoClusterReplicasInner) GetActive() string`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *MqServerInfoClusterReplicasInner) GetActiveOk() (*string, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *MqServerInfoClusterReplicasInner) SetActive(v string)`

SetActive sets Active field to given value.

### HasActive

`func (o *MqServerInfoClusterReplicasInner) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetLag

`func (o *MqServerInfoClusterReplicasInner) GetLag() int32`

GetLag returns the Lag field if non-nil, zero value otherwise.

### GetLagOk

`func (o *MqServerInfoClusterReplicasInner) GetLagOk() (*int32, bool)`

GetLagOk returns a tuple with the Lag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLag

`func (o *MqServerInfoClusterReplicasInner) SetLag(v int32)`

SetLag sets Lag field to given value.

### HasLag

`func (o *MqServerInfoClusterReplicasInner) HasLag() bool`

HasLag returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


