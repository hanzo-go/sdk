# MqPurgeStreamRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Filter** | Pointer to **string** | Subject filter to purge only matching messages. | [optional] 
**Keep** | Pointer to **int32** | Number of messages to keep (from the end). | [optional] 

## Methods

### NewMqPurgeStreamRequest

`func NewMqPurgeStreamRequest() *MqPurgeStreamRequest`

NewMqPurgeStreamRequest instantiates a new MqPurgeStreamRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMqPurgeStreamRequestWithDefaults

`func NewMqPurgeStreamRequestWithDefaults() *MqPurgeStreamRequest`

NewMqPurgeStreamRequestWithDefaults instantiates a new MqPurgeStreamRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilter

`func (o *MqPurgeStreamRequest) GetFilter() string`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *MqPurgeStreamRequest) GetFilterOk() (*string, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *MqPurgeStreamRequest) SetFilter(v string)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *MqPurgeStreamRequest) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### GetKeep

`func (o *MqPurgeStreamRequest) GetKeep() int32`

GetKeep returns the Keep field if non-nil, zero value otherwise.

### GetKeepOk

`func (o *MqPurgeStreamRequest) GetKeepOk() (*int32, bool)`

GetKeepOk returns a tuple with the Keep field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeep

`func (o *MqPurgeStreamRequest) SetKeep(v int32)`

SetKeep sets Keep field to given value.

### HasKeep

`func (o *MqPurgeStreamRequest) HasKeep() bool`

HasKeep returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


