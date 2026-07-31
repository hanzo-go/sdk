# CloudCsrfResp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CsrfToken** | Pointer to **string** | Token is the value to send back in the X-CSRF-Token header. It is bound to the caller&#39;s identity, so it authorizes writes as them and as nobody else. | [optional] 
**ExpiresIn** | Pointer to **int32** | ExpiresIn is the token&#39;s lifetime in seconds. Fetch a new one when it lapses; a write with an expired token is refused. | [optional] 

## Methods

### NewCloudCsrfResp

`func NewCloudCsrfResp() *CloudCsrfResp`

NewCloudCsrfResp instantiates a new CloudCsrfResp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudCsrfRespWithDefaults

`func NewCloudCsrfRespWithDefaults() *CloudCsrfResp`

NewCloudCsrfRespWithDefaults instantiates a new CloudCsrfResp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCsrfToken

`func (o *CloudCsrfResp) GetCsrfToken() string`

GetCsrfToken returns the CsrfToken field if non-nil, zero value otherwise.

### GetCsrfTokenOk

`func (o *CloudCsrfResp) GetCsrfTokenOk() (*string, bool)`

GetCsrfTokenOk returns a tuple with the CsrfToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCsrfToken

`func (o *CloudCsrfResp) SetCsrfToken(v string)`

SetCsrfToken sets CsrfToken field to given value.

### HasCsrfToken

`func (o *CloudCsrfResp) HasCsrfToken() bool`

HasCsrfToken returns a boolean if a field has been set.

### GetExpiresIn

`func (o *CloudCsrfResp) GetExpiresIn() int32`

GetExpiresIn returns the ExpiresIn field if non-nil, zero value otherwise.

### GetExpiresInOk

`func (o *CloudCsrfResp) GetExpiresInOk() (*int32, bool)`

GetExpiresInOk returns a tuple with the ExpiresIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresIn

`func (o *CloudCsrfResp) SetExpiresIn(v int32)`

SetExpiresIn sets ExpiresIn field to given value.

### HasExpiresIn

`func (o *CloudCsrfResp) HasExpiresIn() bool`

HasExpiresIn returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


