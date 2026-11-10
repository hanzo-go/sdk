# SbomSbomView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ComponentCount** | Pointer to **int64** | ComponentCount is how many components this response carries — after the cap, so it matches components exactly rather than the image&#39;s true total. | [optional] 
**Components** | Pointer to [**[]SbomSbomComponent**](SbomSbomComponent.md) | Components is the flattened dependency set, ordered by type then name. | [optional] 
**GitSha** | Pointer to **string** | GitSha is the commit the image was built from. Empty on the same terms as sourceRepo. | [optional] 
**ImageDigest** | Pointer to **string** | ImageDigest is the content-addressed digest the components are keyed under, and is what identifies this bill of materials however it was asked for. | [optional] 
**ImageRef** | Pointer to **string** | ImageRef is the image reference recorded alongside the digest. Empty when the components were ingested against a digest with no ref. | [optional] 
**IngestedAt** | Pointer to **string** | IngestedAt is when the bill of materials was recorded, RFC 3339. | [optional] 
**SourceRepo** | Pointer to **string** | SourceRepo is the repository the image was built from. Empty unless the ingest declared it — a pull from the registry does not learn it. | [optional] 
**Truncated** | Pointer to **bool** | Truncated is true when the image has MORE components than the cap returns. | [optional] 

## Methods

### NewSbomSbomView

`func NewSbomSbomView() *SbomSbomView`

NewSbomSbomView instantiates a new SbomSbomView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSbomSbomViewWithDefaults

`func NewSbomSbomViewWithDefaults() *SbomSbomView`

NewSbomSbomViewWithDefaults instantiates a new SbomSbomView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComponentCount

`func (o *SbomSbomView) GetComponentCount() int64`

GetComponentCount returns the ComponentCount field if non-nil, zero value otherwise.

### GetComponentCountOk

`func (o *SbomSbomView) GetComponentCountOk() (*int64, bool)`

GetComponentCountOk returns a tuple with the ComponentCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentCount

`func (o *SbomSbomView) SetComponentCount(v int64)`

SetComponentCount sets ComponentCount field to given value.

### HasComponentCount

`func (o *SbomSbomView) HasComponentCount() bool`

HasComponentCount returns a boolean if a field has been set.

### GetComponents

`func (o *SbomSbomView) GetComponents() []SbomSbomComponent`

GetComponents returns the Components field if non-nil, zero value otherwise.

### GetComponentsOk

`func (o *SbomSbomView) GetComponentsOk() (*[]SbomSbomComponent, bool)`

GetComponentsOk returns a tuple with the Components field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponents

`func (o *SbomSbomView) SetComponents(v []SbomSbomComponent)`

SetComponents sets Components field to given value.

### HasComponents

`func (o *SbomSbomView) HasComponents() bool`

HasComponents returns a boolean if a field has been set.

### GetGitSha

`func (o *SbomSbomView) GetGitSha() string`

GetGitSha returns the GitSha field if non-nil, zero value otherwise.

### GetGitShaOk

`func (o *SbomSbomView) GetGitShaOk() (*string, bool)`

GetGitShaOk returns a tuple with the GitSha field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitSha

`func (o *SbomSbomView) SetGitSha(v string)`

SetGitSha sets GitSha field to given value.

### HasGitSha

`func (o *SbomSbomView) HasGitSha() bool`

HasGitSha returns a boolean if a field has been set.

### GetImageDigest

`func (o *SbomSbomView) GetImageDigest() string`

GetImageDigest returns the ImageDigest field if non-nil, zero value otherwise.

### GetImageDigestOk

`func (o *SbomSbomView) GetImageDigestOk() (*string, bool)`

GetImageDigestOk returns a tuple with the ImageDigest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageDigest

`func (o *SbomSbomView) SetImageDigest(v string)`

SetImageDigest sets ImageDigest field to given value.

### HasImageDigest

`func (o *SbomSbomView) HasImageDigest() bool`

HasImageDigest returns a boolean if a field has been set.

### GetImageRef

`func (o *SbomSbomView) GetImageRef() string`

GetImageRef returns the ImageRef field if non-nil, zero value otherwise.

### GetImageRefOk

`func (o *SbomSbomView) GetImageRefOk() (*string, bool)`

GetImageRefOk returns a tuple with the ImageRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageRef

`func (o *SbomSbomView) SetImageRef(v string)`

SetImageRef sets ImageRef field to given value.

### HasImageRef

`func (o *SbomSbomView) HasImageRef() bool`

HasImageRef returns a boolean if a field has been set.

### GetIngestedAt

`func (o *SbomSbomView) GetIngestedAt() string`

GetIngestedAt returns the IngestedAt field if non-nil, zero value otherwise.

### GetIngestedAtOk

`func (o *SbomSbomView) GetIngestedAtOk() (*string, bool)`

GetIngestedAtOk returns a tuple with the IngestedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIngestedAt

`func (o *SbomSbomView) SetIngestedAt(v string)`

SetIngestedAt sets IngestedAt field to given value.

### HasIngestedAt

`func (o *SbomSbomView) HasIngestedAt() bool`

HasIngestedAt returns a boolean if a field has been set.

### GetSourceRepo

`func (o *SbomSbomView) GetSourceRepo() string`

GetSourceRepo returns the SourceRepo field if non-nil, zero value otherwise.

### GetSourceRepoOk

`func (o *SbomSbomView) GetSourceRepoOk() (*string, bool)`

GetSourceRepoOk returns a tuple with the SourceRepo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceRepo

`func (o *SbomSbomView) SetSourceRepo(v string)`

SetSourceRepo sets SourceRepo field to given value.

### HasSourceRepo

`func (o *SbomSbomView) HasSourceRepo() bool`

HasSourceRepo returns a boolean if a field has been set.

### GetTruncated

`func (o *SbomSbomView) GetTruncated() bool`

GetTruncated returns the Truncated field if non-nil, zero value otherwise.

### GetTruncatedOk

`func (o *SbomSbomView) GetTruncatedOk() (*bool, bool)`

GetTruncatedOk returns a tuple with the Truncated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTruncated

`func (o *SbomSbomView) SetTruncated(v bool)`

SetTruncated sets Truncated field to given value.

### HasTruncated

`func (o *SbomSbomView) HasTruncated() bool`

HasTruncated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


