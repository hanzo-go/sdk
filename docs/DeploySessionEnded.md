# DeploySessionEnded

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LoggedIn** | Pointer to **bool** | LoggedIn is always false — this is the answer to having just signed out, so it states the resulting session state rather than reporting the request&#39;s outcome. It is not omitempty: false is the whole answer. | [optional] 
**LoginUrl** | Pointer to **string** | LoginURL is where to sign in again. Always present, because a caller that has just signed out is exactly the caller who needs it. | [optional] 

## Methods

### NewDeploySessionEnded

`func NewDeploySessionEnded() *DeploySessionEnded`

NewDeploySessionEnded instantiates a new DeploySessionEnded object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeploySessionEndedWithDefaults

`func NewDeploySessionEndedWithDefaults() *DeploySessionEnded`

NewDeploySessionEndedWithDefaults instantiates a new DeploySessionEnded object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLoggedIn

`func (o *DeploySessionEnded) GetLoggedIn() bool`

GetLoggedIn returns the LoggedIn field if non-nil, zero value otherwise.

### GetLoggedInOk

`func (o *DeploySessionEnded) GetLoggedInOk() (*bool, bool)`

GetLoggedInOk returns a tuple with the LoggedIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoggedIn

`func (o *DeploySessionEnded) SetLoggedIn(v bool)`

SetLoggedIn sets LoggedIn field to given value.

### HasLoggedIn

`func (o *DeploySessionEnded) HasLoggedIn() bool`

HasLoggedIn returns a boolean if a field has been set.

### GetLoginUrl

`func (o *DeploySessionEnded) GetLoginUrl() string`

GetLoginUrl returns the LoginUrl field if non-nil, zero value otherwise.

### GetLoginUrlOk

`func (o *DeploySessionEnded) GetLoginUrlOk() (*string, bool)`

GetLoginUrlOk returns a tuple with the LoginUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginUrl

`func (o *DeploySessionEnded) SetLoginUrl(v string)`

SetLoginUrl sets LoginUrl field to given value.

### HasLoginUrl

`func (o *DeploySessionEnded) HasLoginUrl() bool`

HasLoginUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


