# PlatformImageView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Repository** | Pointer to **string** | Repository is the image path without a tag (ghcr.io/acme/api). Required for source &#x60;image&#x60;, which runs it as-is. A git app&#39;s built image is NOT this: the build pushes to a path derived from the org and slug, and the deployment records that full ref. | [optional] 
**Tag** | Pointer to **string** | Tag is the tag to run: what the create declared, then RE-STAMPED on every transition to live with the tag that actually went live. So after a deploy it names what is running, not what was asked for. | [optional] 

## Methods

### NewPlatformImageView

`func NewPlatformImageView() *PlatformImageView`

NewPlatformImageView instantiates a new PlatformImageView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlatformImageViewWithDefaults

`func NewPlatformImageViewWithDefaults() *PlatformImageView`

NewPlatformImageViewWithDefaults instantiates a new PlatformImageView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRepository

`func (o *PlatformImageView) GetRepository() string`

GetRepository returns the Repository field if non-nil, zero value otherwise.

### GetRepositoryOk

`func (o *PlatformImageView) GetRepositoryOk() (*string, bool)`

GetRepositoryOk returns a tuple with the Repository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepository

`func (o *PlatformImageView) SetRepository(v string)`

SetRepository sets Repository field to given value.

### HasRepository

`func (o *PlatformImageView) HasRepository() bool`

HasRepository returns a boolean if a field has been set.

### GetTag

`func (o *PlatformImageView) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *PlatformImageView) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *PlatformImageView) SetTag(v string)`

SetTag sets Tag field to given value.

### HasTag

`func (o *PlatformImageView) HasTag() bool`

HasTag returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


