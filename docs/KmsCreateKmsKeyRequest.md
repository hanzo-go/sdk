# KmsCreateKmsKeyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**ProjectId** | **string** |  | 

## Methods

### NewKmsCreateKmsKeyRequest

`func NewKmsCreateKmsKeyRequest(name string, projectId string, ) *KmsCreateKmsKeyRequest`

NewKmsCreateKmsKeyRequest instantiates a new KmsCreateKmsKeyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKmsCreateKmsKeyRequestWithDefaults

`func NewKmsCreateKmsKeyRequestWithDefaults() *KmsCreateKmsKeyRequest`

NewKmsCreateKmsKeyRequestWithDefaults instantiates a new KmsCreateKmsKeyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *KmsCreateKmsKeyRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KmsCreateKmsKeyRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *KmsCreateKmsKeyRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *KmsCreateKmsKeyRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *KmsCreateKmsKeyRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *KmsCreateKmsKeyRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *KmsCreateKmsKeyRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetProjectId

`func (o *KmsCreateKmsKeyRequest) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *KmsCreateKmsKeyRequest) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *KmsCreateKmsKeyRequest) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


