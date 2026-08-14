# IamCreateSessionIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Application** | **string** |  | 
**ExclusiveSignin** | Pointer to **bool** |  | [optional] 
**Name** | **string** |  | 
**Owner** | **string** |  | 
**SessionId** | Pointer to **[]string** |  | [optional] 

## Methods

### NewIamCreateSessionIn

`func NewIamCreateSessionIn(application string, name string, owner string, ) *IamCreateSessionIn`

NewIamCreateSessionIn instantiates a new IamCreateSessionIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamCreateSessionInWithDefaults

`func NewIamCreateSessionInWithDefaults() *IamCreateSessionIn`

NewIamCreateSessionInWithDefaults instantiates a new IamCreateSessionIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplication

`func (o *IamCreateSessionIn) GetApplication() string`

GetApplication returns the Application field if non-nil, zero value otherwise.

### GetApplicationOk

`func (o *IamCreateSessionIn) GetApplicationOk() (*string, bool)`

GetApplicationOk returns a tuple with the Application field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplication

`func (o *IamCreateSessionIn) SetApplication(v string)`

SetApplication sets Application field to given value.


### GetExclusiveSignin

`func (o *IamCreateSessionIn) GetExclusiveSignin() bool`

GetExclusiveSignin returns the ExclusiveSignin field if non-nil, zero value otherwise.

### GetExclusiveSigninOk

`func (o *IamCreateSessionIn) GetExclusiveSigninOk() (*bool, bool)`

GetExclusiveSigninOk returns a tuple with the ExclusiveSignin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExclusiveSignin

`func (o *IamCreateSessionIn) SetExclusiveSignin(v bool)`

SetExclusiveSignin sets ExclusiveSignin field to given value.

### HasExclusiveSignin

`func (o *IamCreateSessionIn) HasExclusiveSignin() bool`

HasExclusiveSignin returns a boolean if a field has been set.

### GetName

`func (o *IamCreateSessionIn) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IamCreateSessionIn) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IamCreateSessionIn) SetName(v string)`

SetName sets Name field to given value.


### GetOwner

`func (o *IamCreateSessionIn) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *IamCreateSessionIn) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *IamCreateSessionIn) SetOwner(v string)`

SetOwner sets Owner field to given value.


### GetSessionId

`func (o *IamCreateSessionIn) GetSessionId() []string`

GetSessionId returns the SessionId field if non-nil, zero value otherwise.

### GetSessionIdOk

`func (o *IamCreateSessionIn) GetSessionIdOk() (*[]string, bool)`

GetSessionIdOk returns a tuple with the SessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionId

`func (o *IamCreateSessionIn) SetSessionId(v []string)`

SetSessionId sets SessionId field to given value.

### HasSessionId

`func (o *IamCreateSessionIn) HasSessionId() bool`

HasSessionId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


