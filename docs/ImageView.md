# ImageView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Repository** | Pointer to **string** |  | [optional] 
**Tag** | Pointer to **string** |  | [optional] 

## Methods

### NewImageView

`func NewImageView() *ImageView`

NewImageView instantiates a new ImageView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImageViewWithDefaults

`func NewImageViewWithDefaults() *ImageView`

NewImageViewWithDefaults instantiates a new ImageView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRepository

`func (o *ImageView) GetRepository() string`

GetRepository returns the Repository field if non-nil, zero value otherwise.

### GetRepositoryOk

`func (o *ImageView) GetRepositoryOk() (*string, bool)`

GetRepositoryOk returns a tuple with the Repository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepository

`func (o *ImageView) SetRepository(v string)`

SetRepository sets Repository field to given value.

### HasRepository

`func (o *ImageView) HasRepository() bool`

HasRepository returns a boolean if a field has been set.

### GetTag

`func (o *ImageView) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *ImageView) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *ImageView) SetTag(v string)`

SetTag sets Tag field to given value.

### HasTag

`func (o *ImageView) HasTag() bool`

HasTag returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


