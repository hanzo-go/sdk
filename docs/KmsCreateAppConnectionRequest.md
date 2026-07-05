# KmsCreateAppConnectionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**App** | **string** |  | 
**Credentials** | **map[string]interface{}** |  | 

## Methods

### NewKmsCreateAppConnectionRequest

`func NewKmsCreateAppConnectionRequest(name string, app string, credentials map[string]interface{}, ) *KmsCreateAppConnectionRequest`

NewKmsCreateAppConnectionRequest instantiates a new KmsCreateAppConnectionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKmsCreateAppConnectionRequestWithDefaults

`func NewKmsCreateAppConnectionRequestWithDefaults() *KmsCreateAppConnectionRequest`

NewKmsCreateAppConnectionRequestWithDefaults instantiates a new KmsCreateAppConnectionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *KmsCreateAppConnectionRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KmsCreateAppConnectionRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *KmsCreateAppConnectionRequest) SetName(v string)`

SetName sets Name field to given value.


### GetApp

`func (o *KmsCreateAppConnectionRequest) GetApp() string`

GetApp returns the App field if non-nil, zero value otherwise.

### GetAppOk

`func (o *KmsCreateAppConnectionRequest) GetAppOk() (*string, bool)`

GetAppOk returns a tuple with the App field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApp

`func (o *KmsCreateAppConnectionRequest) SetApp(v string)`

SetApp sets App field to given value.


### GetCredentials

`func (o *KmsCreateAppConnectionRequest) GetCredentials() map[string]interface{}`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *KmsCreateAppConnectionRequest) GetCredentialsOk() (*map[string]interface{}, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *KmsCreateAppConnectionRequest) SetCredentials(v map[string]interface{})`

SetCredentials sets Credentials field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


