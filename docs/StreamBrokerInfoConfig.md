# StreamBrokerInfoConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Replicas** | Pointer to **int32** |  | [optional] 
**Storage** | Pointer to **string** |  | [optional] 

## Methods

### NewStreamBrokerInfoConfig

`func NewStreamBrokerInfoConfig() *StreamBrokerInfoConfig`

NewStreamBrokerInfoConfig instantiates a new StreamBrokerInfoConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStreamBrokerInfoConfigWithDefaults

`func NewStreamBrokerInfoConfigWithDefaults() *StreamBrokerInfoConfig`

NewStreamBrokerInfoConfigWithDefaults instantiates a new StreamBrokerInfoConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReplicas

`func (o *StreamBrokerInfoConfig) GetReplicas() int32`

GetReplicas returns the Replicas field if non-nil, zero value otherwise.

### GetReplicasOk

`func (o *StreamBrokerInfoConfig) GetReplicasOk() (*int32, bool)`

GetReplicasOk returns a tuple with the Replicas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicas

`func (o *StreamBrokerInfoConfig) SetReplicas(v int32)`

SetReplicas sets Replicas field to given value.

### HasReplicas

`func (o *StreamBrokerInfoConfig) HasReplicas() bool`

HasReplicas returns a boolean if a field has been set.

### GetStorage

`func (o *StreamBrokerInfoConfig) GetStorage() string`

GetStorage returns the Storage field if non-nil, zero value otherwise.

### GetStorageOk

`func (o *StreamBrokerInfoConfig) GetStorageOk() (*string, bool)`

GetStorageOk returns a tuple with the Storage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorage

`func (o *StreamBrokerInfoConfig) SetStorage(v string)`

SetStorage sets Storage field to given value.

### HasStorage

`func (o *StreamBrokerInfoConfig) HasStorage() bool`

HasStorage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


