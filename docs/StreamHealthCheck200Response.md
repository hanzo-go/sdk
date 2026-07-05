# StreamHealthCheck200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** |  | [optional] 
**Version** | Pointer to **string** |  | [optional] 
**PubsubConnected** | Pointer to **bool** |  | [optional] 
**Topics** | Pointer to **int32** |  | [optional] 

## Methods

### NewStreamHealthCheck200Response

`func NewStreamHealthCheck200Response() *StreamHealthCheck200Response`

NewStreamHealthCheck200Response instantiates a new StreamHealthCheck200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStreamHealthCheck200ResponseWithDefaults

`func NewStreamHealthCheck200ResponseWithDefaults() *StreamHealthCheck200Response`

NewStreamHealthCheck200ResponseWithDefaults instantiates a new StreamHealthCheck200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *StreamHealthCheck200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *StreamHealthCheck200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *StreamHealthCheck200Response) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *StreamHealthCheck200Response) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetVersion

`func (o *StreamHealthCheck200Response) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *StreamHealthCheck200Response) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *StreamHealthCheck200Response) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *StreamHealthCheck200Response) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetPubsubConnected

`func (o *StreamHealthCheck200Response) GetPubsubConnected() bool`

GetPubsubConnected returns the PubsubConnected field if non-nil, zero value otherwise.

### GetPubsubConnectedOk

`func (o *StreamHealthCheck200Response) GetPubsubConnectedOk() (*bool, bool)`

GetPubsubConnectedOk returns a tuple with the PubsubConnected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPubsubConnected

`func (o *StreamHealthCheck200Response) SetPubsubConnected(v bool)`

SetPubsubConnected sets PubsubConnected field to given value.

### HasPubsubConnected

`func (o *StreamHealthCheck200Response) HasPubsubConnected() bool`

HasPubsubConnected returns a boolean if a field has been set.

### GetTopics

`func (o *StreamHealthCheck200Response) GetTopics() int32`

GetTopics returns the Topics field if non-nil, zero value otherwise.

### GetTopicsOk

`func (o *StreamHealthCheck200Response) GetTopicsOk() (*int32, bool)`

GetTopicsOk returns a tuple with the Topics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopics

`func (o *StreamHealthCheck200Response) SetTopics(v int32)`

SetTopics sets Topics field to given value.

### HasTopics

`func (o *StreamHealthCheck200Response) HasTopics() bool`

HasTopics returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


