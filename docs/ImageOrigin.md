# ImageOrigin

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Repository** | Pointer to **string** | Repository is the image repository. Required for source &#x60;image&#x60;. | [optional] 
**Tag** | Pointer to **string** | Tag is the image tag to deploy; &#x60;latest&#x60; when omitted. | [optional] 

## Methods

### NewImageOrigin

`func NewImageOrigin() *ImageOrigin`

NewImageOrigin instantiates a new ImageOrigin object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImageOriginWithDefaults

`func NewImageOriginWithDefaults() *ImageOrigin`

NewImageOriginWithDefaults instantiates a new ImageOrigin object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRepository

`func (o *ImageOrigin) GetRepository() string`

GetRepository returns the Repository field if non-nil, zero value otherwise.

### GetRepositoryOk

`func (o *ImageOrigin) GetRepositoryOk() (*string, bool)`

GetRepositoryOk returns a tuple with the Repository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepository

`func (o *ImageOrigin) SetRepository(v string)`

SetRepository sets Repository field to given value.

### HasRepository

`func (o *ImageOrigin) HasRepository() bool`

HasRepository returns a boolean if a field has been set.

### GetTag

`func (o *ImageOrigin) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *ImageOrigin) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *ImageOrigin) SetTag(v string)`

SetTag sets Tag field to given value.

### HasTag

`func (o *ImageOrigin) HasTag() bool`

HasTag returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


