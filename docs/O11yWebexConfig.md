# O11yWebexConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NotifierConfig** | Pointer to [**O11yNotifierConfig**](O11yNotifierConfig.md) |  | [optional] 
**ApiUrl** | Pointer to **interface{}** |  | [optional] 
**HttpConfig** | Pointer to [**O11yHTTPClientConfig**](O11yHTTPClientConfig.md) |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**RoomId** | Pointer to **string** |  | [optional] 

## Methods

### NewO11yWebexConfig

`func NewO11yWebexConfig() *O11yWebexConfig`

NewO11yWebexConfig instantiates a new O11yWebexConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yWebexConfigWithDefaults

`func NewO11yWebexConfigWithDefaults() *O11yWebexConfig`

NewO11yWebexConfigWithDefaults instantiates a new O11yWebexConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNotifierConfig

`func (o *O11yWebexConfig) GetNotifierConfig() O11yNotifierConfig`

GetNotifierConfig returns the NotifierConfig field if non-nil, zero value otherwise.

### GetNotifierConfigOk

`func (o *O11yWebexConfig) GetNotifierConfigOk() (*O11yNotifierConfig, bool)`

GetNotifierConfigOk returns a tuple with the NotifierConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifierConfig

`func (o *O11yWebexConfig) SetNotifierConfig(v O11yNotifierConfig)`

SetNotifierConfig sets NotifierConfig field to given value.

### HasNotifierConfig

`func (o *O11yWebexConfig) HasNotifierConfig() bool`

HasNotifierConfig returns a boolean if a field has been set.

### GetApiUrl

`func (o *O11yWebexConfig) GetApiUrl() interface{}`

GetApiUrl returns the ApiUrl field if non-nil, zero value otherwise.

### GetApiUrlOk

`func (o *O11yWebexConfig) GetApiUrlOk() (*interface{}, bool)`

GetApiUrlOk returns a tuple with the ApiUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiUrl

`func (o *O11yWebexConfig) SetApiUrl(v interface{})`

SetApiUrl sets ApiUrl field to given value.

### HasApiUrl

`func (o *O11yWebexConfig) HasApiUrl() bool`

HasApiUrl returns a boolean if a field has been set.

### SetApiUrlNil

`func (o *O11yWebexConfig) SetApiUrlNil(b bool)`

 SetApiUrlNil sets the value for ApiUrl to be an explicit nil

### UnsetApiUrl
`func (o *O11yWebexConfig) UnsetApiUrl()`

UnsetApiUrl ensures that no value is present for ApiUrl, not even an explicit nil
### GetHttpConfig

`func (o *O11yWebexConfig) GetHttpConfig() O11yHTTPClientConfig`

GetHttpConfig returns the HttpConfig field if non-nil, zero value otherwise.

### GetHttpConfigOk

`func (o *O11yWebexConfig) GetHttpConfigOk() (*O11yHTTPClientConfig, bool)`

GetHttpConfigOk returns a tuple with the HttpConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpConfig

`func (o *O11yWebexConfig) SetHttpConfig(v O11yHTTPClientConfig)`

SetHttpConfig sets HttpConfig field to given value.

### HasHttpConfig

`func (o *O11yWebexConfig) HasHttpConfig() bool`

HasHttpConfig returns a boolean if a field has been set.

### GetMessage

`func (o *O11yWebexConfig) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *O11yWebexConfig) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *O11yWebexConfig) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *O11yWebexConfig) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetRoomId

`func (o *O11yWebexConfig) GetRoomId() string`

GetRoomId returns the RoomId field if non-nil, zero value otherwise.

### GetRoomIdOk

`func (o *O11yWebexConfig) GetRoomIdOk() (*string, bool)`

GetRoomIdOk returns a tuple with the RoomId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoomId

`func (o *O11yWebexConfig) SetRoomId(v string)`

SetRoomId sets RoomId field to given value.

### HasRoomId

`func (o *O11yWebexConfig) HasRoomId() bool`

HasRoomId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


