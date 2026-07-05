# KmsUpdateAppConnectionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Credentials** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewKmsUpdateAppConnectionRequest

`func NewKmsUpdateAppConnectionRequest() *KmsUpdateAppConnectionRequest`

NewKmsUpdateAppConnectionRequest instantiates a new KmsUpdateAppConnectionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKmsUpdateAppConnectionRequestWithDefaults

`func NewKmsUpdateAppConnectionRequestWithDefaults() *KmsUpdateAppConnectionRequest`

NewKmsUpdateAppConnectionRequestWithDefaults instantiates a new KmsUpdateAppConnectionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *KmsUpdateAppConnectionRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KmsUpdateAppConnectionRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *KmsUpdateAppConnectionRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *KmsUpdateAppConnectionRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCredentials

`func (o *KmsUpdateAppConnectionRequest) GetCredentials() map[string]interface{}`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *KmsUpdateAppConnectionRequest) GetCredentialsOk() (*map[string]interface{}, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *KmsUpdateAppConnectionRequest) SetCredentials(v map[string]interface{})`

SetCredentials sets Credentials field to given value.

### HasCredentials

`func (o *KmsUpdateAppConnectionRequest) HasCredentials() bool`

HasCredentials returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


