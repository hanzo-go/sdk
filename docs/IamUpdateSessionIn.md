# IamUpdateSessionIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Application** | **string** |  | 
**Name** | **string** |  | 
**Owner** | **string** |  | 
**SessionId** | Pointer to **[]string** |  | [optional] 

## Methods

### NewIamUpdateSessionIn

`func NewIamUpdateSessionIn(application string, name string, owner string, ) *IamUpdateSessionIn`

NewIamUpdateSessionIn instantiates a new IamUpdateSessionIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamUpdateSessionInWithDefaults

`func NewIamUpdateSessionInWithDefaults() *IamUpdateSessionIn`

NewIamUpdateSessionInWithDefaults instantiates a new IamUpdateSessionIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplication

`func (o *IamUpdateSessionIn) GetApplication() string`

GetApplication returns the Application field if non-nil, zero value otherwise.

### GetApplicationOk

`func (o *IamUpdateSessionIn) GetApplicationOk() (*string, bool)`

GetApplicationOk returns a tuple with the Application field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplication

`func (o *IamUpdateSessionIn) SetApplication(v string)`

SetApplication sets Application field to given value.


### GetName

`func (o *IamUpdateSessionIn) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IamUpdateSessionIn) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IamUpdateSessionIn) SetName(v string)`

SetName sets Name field to given value.


### GetOwner

`func (o *IamUpdateSessionIn) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *IamUpdateSessionIn) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *IamUpdateSessionIn) SetOwner(v string)`

SetOwner sets Owner field to given value.


### GetSessionId

`func (o *IamUpdateSessionIn) GetSessionId() []string`

GetSessionId returns the SessionId field if non-nil, zero value otherwise.

### GetSessionIdOk

`func (o *IamUpdateSessionIn) GetSessionIdOk() (*[]string, bool)`

GetSessionIdOk returns a tuple with the SessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionId

`func (o *IamUpdateSessionIn) SetSessionId(v []string)`

SetSessionId sets SessionId field to given value.

### HasSessionId

`func (o *IamUpdateSessionIn) HasSessionId() bool`

HasSessionId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


