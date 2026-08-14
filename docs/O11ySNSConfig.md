# O11ySNSConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NotifierConfig** | Pointer to [**O11yNotifierConfig**](O11yNotifierConfig.md) |  | [optional] 
**ApiUrl** | Pointer to **string** |  | [optional] 
**Attributes** | Pointer to **map[string]string** |  | [optional] 
**HttpConfig** | Pointer to [**O11yHTTPClientConfig**](O11yHTTPClientConfig.md) |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**PhoneNumber** | Pointer to **string** |  | [optional] 
**Sigv4** | Pointer to [**O11ySigV4Config**](O11ySigV4Config.md) |  | [optional] 
**Subject** | Pointer to **string** |  | [optional] 
**TargetArn** | Pointer to **string** |  | [optional] 
**TopicArn** | Pointer to **string** |  | [optional] 

## Methods

### NewO11ySNSConfig

`func NewO11ySNSConfig() *O11ySNSConfig`

NewO11ySNSConfig instantiates a new O11ySNSConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11ySNSConfigWithDefaults

`func NewO11ySNSConfigWithDefaults() *O11ySNSConfig`

NewO11ySNSConfigWithDefaults instantiates a new O11ySNSConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNotifierConfig

`func (o *O11ySNSConfig) GetNotifierConfig() O11yNotifierConfig`

GetNotifierConfig returns the NotifierConfig field if non-nil, zero value otherwise.

### GetNotifierConfigOk

`func (o *O11ySNSConfig) GetNotifierConfigOk() (*O11yNotifierConfig, bool)`

GetNotifierConfigOk returns a tuple with the NotifierConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifierConfig

`func (o *O11ySNSConfig) SetNotifierConfig(v O11yNotifierConfig)`

SetNotifierConfig sets NotifierConfig field to given value.

### HasNotifierConfig

`func (o *O11ySNSConfig) HasNotifierConfig() bool`

HasNotifierConfig returns a boolean if a field has been set.

### GetApiUrl

`func (o *O11ySNSConfig) GetApiUrl() string`

GetApiUrl returns the ApiUrl field if non-nil, zero value otherwise.

### GetApiUrlOk

`func (o *O11ySNSConfig) GetApiUrlOk() (*string, bool)`

GetApiUrlOk returns a tuple with the ApiUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiUrl

`func (o *O11ySNSConfig) SetApiUrl(v string)`

SetApiUrl sets ApiUrl field to given value.

### HasApiUrl

`func (o *O11ySNSConfig) HasApiUrl() bool`

HasApiUrl returns a boolean if a field has been set.

### GetAttributes

`func (o *O11ySNSConfig) GetAttributes() map[string]string`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *O11ySNSConfig) GetAttributesOk() (*map[string]string, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *O11ySNSConfig) SetAttributes(v map[string]string)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *O11ySNSConfig) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetHttpConfig

`func (o *O11ySNSConfig) GetHttpConfig() O11yHTTPClientConfig`

GetHttpConfig returns the HttpConfig field if non-nil, zero value otherwise.

### GetHttpConfigOk

`func (o *O11ySNSConfig) GetHttpConfigOk() (*O11yHTTPClientConfig, bool)`

GetHttpConfigOk returns a tuple with the HttpConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpConfig

`func (o *O11ySNSConfig) SetHttpConfig(v O11yHTTPClientConfig)`

SetHttpConfig sets HttpConfig field to given value.

### HasHttpConfig

`func (o *O11ySNSConfig) HasHttpConfig() bool`

HasHttpConfig returns a boolean if a field has been set.

### GetMessage

`func (o *O11ySNSConfig) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *O11ySNSConfig) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *O11ySNSConfig) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *O11ySNSConfig) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetPhoneNumber

`func (o *O11ySNSConfig) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *O11ySNSConfig) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *O11ySNSConfig) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.

### HasPhoneNumber

`func (o *O11ySNSConfig) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.

### GetSigv4

`func (o *O11ySNSConfig) GetSigv4() O11ySigV4Config`

GetSigv4 returns the Sigv4 field if non-nil, zero value otherwise.

### GetSigv4Ok

`func (o *O11ySNSConfig) GetSigv4Ok() (*O11ySigV4Config, bool)`

GetSigv4Ok returns a tuple with the Sigv4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigv4

`func (o *O11ySNSConfig) SetSigv4(v O11ySigV4Config)`

SetSigv4 sets Sigv4 field to given value.

### HasSigv4

`func (o *O11ySNSConfig) HasSigv4() bool`

HasSigv4 returns a boolean if a field has been set.

### GetSubject

`func (o *O11ySNSConfig) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *O11ySNSConfig) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *O11ySNSConfig) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *O11ySNSConfig) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### GetTargetArn

`func (o *O11ySNSConfig) GetTargetArn() string`

GetTargetArn returns the TargetArn field if non-nil, zero value otherwise.

### GetTargetArnOk

`func (o *O11ySNSConfig) GetTargetArnOk() (*string, bool)`

GetTargetArnOk returns a tuple with the TargetArn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetArn

`func (o *O11ySNSConfig) SetTargetArn(v string)`

SetTargetArn sets TargetArn field to given value.

### HasTargetArn

`func (o *O11ySNSConfig) HasTargetArn() bool`

HasTargetArn returns a boolean if a field has been set.

### GetTopicArn

`func (o *O11ySNSConfig) GetTopicArn() string`

GetTopicArn returns the TopicArn field if non-nil, zero value otherwise.

### GetTopicArnOk

`func (o *O11ySNSConfig) GetTopicArnOk() (*string, bool)`

GetTopicArnOk returns a tuple with the TopicArn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopicArn

`func (o *O11ySNSConfig) SetTopicArn(v string)`

SetTopicArn sets TopicArn field to given value.

### HasTopicArn

`func (o *O11ySNSConfig) HasTopicArn() bool`

HasTopicArn returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


