# KmsGetCurrentUser200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**User** | Pointer to [**KmsUser**](KmsUser.md) |  | [optional] 

## Methods

### NewKmsGetCurrentUser200Response

`func NewKmsGetCurrentUser200Response() *KmsGetCurrentUser200Response`

NewKmsGetCurrentUser200Response instantiates a new KmsGetCurrentUser200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKmsGetCurrentUser200ResponseWithDefaults

`func NewKmsGetCurrentUser200ResponseWithDefaults() *KmsGetCurrentUser200Response`

NewKmsGetCurrentUser200ResponseWithDefaults instantiates a new KmsGetCurrentUser200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUser

`func (o *KmsGetCurrentUser200Response) GetUser() KmsUser`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *KmsGetCurrentUser200Response) GetUserOk() (*KmsUser, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *KmsGetCurrentUser200Response) SetUser(v KmsUser)`

SetUser sets User field to given value.

### HasUser

`func (o *KmsGetCurrentUser200Response) HasUser() bool`

HasUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


