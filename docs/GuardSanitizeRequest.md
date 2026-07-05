# GuardSanitizeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Text** | **string** | Text to sanitize | 
**Direction** | Pointer to **string** | Sanitization direction (input runs all 5 stages, output runs PII + content filter) | [optional] [default to "input"]
**UserId** | Pointer to **string** | User ID for rate limiting and audit | [optional] 
**SessionId** | Pointer to **string** | Session ID for audit correlation | [optional] 
**Config** | Pointer to [**GuardSanitizeConfig**](GuardSanitizeConfig.md) |  | [optional] 

## Methods

### NewGuardSanitizeRequest

`func NewGuardSanitizeRequest(text string, ) *GuardSanitizeRequest`

NewGuardSanitizeRequest instantiates a new GuardSanitizeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGuardSanitizeRequestWithDefaults

`func NewGuardSanitizeRequestWithDefaults() *GuardSanitizeRequest`

NewGuardSanitizeRequestWithDefaults instantiates a new GuardSanitizeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetText

`func (o *GuardSanitizeRequest) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *GuardSanitizeRequest) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *GuardSanitizeRequest) SetText(v string)`

SetText sets Text field to given value.


### GetDirection

`func (o *GuardSanitizeRequest) GetDirection() string`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *GuardSanitizeRequest) GetDirectionOk() (*string, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *GuardSanitizeRequest) SetDirection(v string)`

SetDirection sets Direction field to given value.

### HasDirection

`func (o *GuardSanitizeRequest) HasDirection() bool`

HasDirection returns a boolean if a field has been set.

### GetUserId

`func (o *GuardSanitizeRequest) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *GuardSanitizeRequest) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *GuardSanitizeRequest) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *GuardSanitizeRequest) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetSessionId

`func (o *GuardSanitizeRequest) GetSessionId() string`

GetSessionId returns the SessionId field if non-nil, zero value otherwise.

### GetSessionIdOk

`func (o *GuardSanitizeRequest) GetSessionIdOk() (*string, bool)`

GetSessionIdOk returns a tuple with the SessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionId

`func (o *GuardSanitizeRequest) SetSessionId(v string)`

SetSessionId sets SessionId field to given value.

### HasSessionId

`func (o *GuardSanitizeRequest) HasSessionId() bool`

HasSessionId returns a boolean if a field has been set.

### GetConfig

`func (o *GuardSanitizeRequest) GetConfig() GuardSanitizeConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *GuardSanitizeRequest) GetConfigOk() (*GuardSanitizeConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *GuardSanitizeRequest) SetConfig(v GuardSanitizeConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *GuardSanitizeRequest) HasConfig() bool`

HasConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


