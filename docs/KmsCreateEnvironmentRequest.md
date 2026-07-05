# KmsCreateEnvironmentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Slug** | **string** |  | 
**Position** | Pointer to **int32** |  | [optional] 

## Methods

### NewKmsCreateEnvironmentRequest

`func NewKmsCreateEnvironmentRequest(name string, slug string, ) *KmsCreateEnvironmentRequest`

NewKmsCreateEnvironmentRequest instantiates a new KmsCreateEnvironmentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKmsCreateEnvironmentRequestWithDefaults

`func NewKmsCreateEnvironmentRequestWithDefaults() *KmsCreateEnvironmentRequest`

NewKmsCreateEnvironmentRequestWithDefaults instantiates a new KmsCreateEnvironmentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *KmsCreateEnvironmentRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KmsCreateEnvironmentRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *KmsCreateEnvironmentRequest) SetName(v string)`

SetName sets Name field to given value.


### GetSlug

`func (o *KmsCreateEnvironmentRequest) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *KmsCreateEnvironmentRequest) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *KmsCreateEnvironmentRequest) SetSlug(v string)`

SetSlug sets Slug field to given value.


### GetPosition

`func (o *KmsCreateEnvironmentRequest) GetPosition() int32`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *KmsCreateEnvironmentRequest) GetPositionOk() (*int32, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *KmsCreateEnvironmentRequest) SetPosition(v int32)`

SetPosition sets Position field to given value.

### HasPosition

`func (o *KmsCreateEnvironmentRequest) HasPosition() bool`

HasPosition returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


