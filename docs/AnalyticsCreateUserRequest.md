# AnalyticsCreateUserRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Username** | **string** |  | 
**Password** | **string** |  | 
**Role** | **string** |  | 

## Methods

### NewAnalyticsCreateUserRequest

`func NewAnalyticsCreateUserRequest(username string, password string, role string, ) *AnalyticsCreateUserRequest`

NewAnalyticsCreateUserRequest instantiates a new AnalyticsCreateUserRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnalyticsCreateUserRequestWithDefaults

`func NewAnalyticsCreateUserRequestWithDefaults() *AnalyticsCreateUserRequest`

NewAnalyticsCreateUserRequestWithDefaults instantiates a new AnalyticsCreateUserRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AnalyticsCreateUserRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AnalyticsCreateUserRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AnalyticsCreateUserRequest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AnalyticsCreateUserRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUsername

`func (o *AnalyticsCreateUserRequest) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *AnalyticsCreateUserRequest) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *AnalyticsCreateUserRequest) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetPassword

`func (o *AnalyticsCreateUserRequest) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *AnalyticsCreateUserRequest) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *AnalyticsCreateUserRequest) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetRole

`func (o *AnalyticsCreateUserRequest) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *AnalyticsCreateUserRequest) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *AnalyticsCreateUserRequest) SetRole(v string)`

SetRole sets Role field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


