# SessionEnded

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LoggedIn** | Pointer to **bool** | LoggedIn is always false — this is the answer to having just signed out, so it states the resulting session state rather than reporting the request&#39;s outcome. It is not omitempty: false is the whole answer. | [optional] 
**LoginUrl** | Pointer to **string** | LoginURL is where to sign in again. Always present, because a caller that has just signed out is exactly the caller who needs it. | [optional] 

## Methods

### NewSessionEnded

`func NewSessionEnded() *SessionEnded`

NewSessionEnded instantiates a new SessionEnded object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSessionEndedWithDefaults

`func NewSessionEndedWithDefaults() *SessionEnded`

NewSessionEndedWithDefaults instantiates a new SessionEnded object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLoggedIn

`func (o *SessionEnded) GetLoggedIn() bool`

GetLoggedIn returns the LoggedIn field if non-nil, zero value otherwise.

### GetLoggedInOk

`func (o *SessionEnded) GetLoggedInOk() (*bool, bool)`

GetLoggedInOk returns a tuple with the LoggedIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoggedIn

`func (o *SessionEnded) SetLoggedIn(v bool)`

SetLoggedIn sets LoggedIn field to given value.

### HasLoggedIn

`func (o *SessionEnded) HasLoggedIn() bool`

HasLoggedIn returns a boolean if a field has been set.

### GetLoginUrl

`func (o *SessionEnded) GetLoginUrl() string`

GetLoginUrl returns the LoginUrl field if non-nil, zero value otherwise.

### GetLoginUrlOk

`func (o *SessionEnded) GetLoginUrlOk() (*string, bool)`

GetLoginUrlOk returns a tuple with the LoginUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginUrl

`func (o *SessionEnded) SetLoginUrl(v string)`

SetLoginUrl sets LoginUrl field to given value.

### HasLoginUrl

`func (o *SessionEnded) HasLoginUrl() bool`

HasLoginUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


