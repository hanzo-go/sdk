# SbomIngest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Document** | Pointer to **interface{}** |  | [optional] 
**Format** | Pointer to **string** | Format names the document format; \&quot;cyclonedx\&quot; is the only one parsed. | [optional] 
**GitSha** | Pointer to **string** | GitSha is the commit the image was built from. | [optional] 
**ImageDigest** | Pointer to **string** | ImageDigest is the content-addressed digest (sha256:…) the components are keyed under. Required — it, not a tenant, is what an SBOM belongs to. | [optional] 
**ImageRef** | Pointer to **string** | ImageRef is the human-readable image reference the digest was published as. A resolve matches on either this or the digest. | [optional] 
**SourceRepo** | Pointer to **string** | SourceRepo is the repository the image was built from. | [optional] 

## Methods

### NewSbomIngest

`func NewSbomIngest() *SbomIngest`

NewSbomIngest instantiates a new SbomIngest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSbomIngestWithDefaults

`func NewSbomIngestWithDefaults() *SbomIngest`

NewSbomIngestWithDefaults instantiates a new SbomIngest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDocument

`func (o *SbomIngest) GetDocument() interface{}`

GetDocument returns the Document field if non-nil, zero value otherwise.

### GetDocumentOk

`func (o *SbomIngest) GetDocumentOk() (*interface{}, bool)`

GetDocumentOk returns a tuple with the Document field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocument

`func (o *SbomIngest) SetDocument(v interface{})`

SetDocument sets Document field to given value.

### HasDocument

`func (o *SbomIngest) HasDocument() bool`

HasDocument returns a boolean if a field has been set.

### SetDocumentNil

`func (o *SbomIngest) SetDocumentNil(b bool)`

 SetDocumentNil sets the value for Document to be an explicit nil

### UnsetDocument
`func (o *SbomIngest) UnsetDocument()`

UnsetDocument ensures that no value is present for Document, not even an explicit nil
### GetFormat

`func (o *SbomIngest) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *SbomIngest) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *SbomIngest) SetFormat(v string)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *SbomIngest) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### GetGitSha

`func (o *SbomIngest) GetGitSha() string`

GetGitSha returns the GitSha field if non-nil, zero value otherwise.

### GetGitShaOk

`func (o *SbomIngest) GetGitShaOk() (*string, bool)`

GetGitShaOk returns a tuple with the GitSha field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitSha

`func (o *SbomIngest) SetGitSha(v string)`

SetGitSha sets GitSha field to given value.

### HasGitSha

`func (o *SbomIngest) HasGitSha() bool`

HasGitSha returns a boolean if a field has been set.

### GetImageDigest

`func (o *SbomIngest) GetImageDigest() string`

GetImageDigest returns the ImageDigest field if non-nil, zero value otherwise.

### GetImageDigestOk

`func (o *SbomIngest) GetImageDigestOk() (*string, bool)`

GetImageDigestOk returns a tuple with the ImageDigest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageDigest

`func (o *SbomIngest) SetImageDigest(v string)`

SetImageDigest sets ImageDigest field to given value.

### HasImageDigest

`func (o *SbomIngest) HasImageDigest() bool`

HasImageDigest returns a boolean if a field has been set.

### GetImageRef

`func (o *SbomIngest) GetImageRef() string`

GetImageRef returns the ImageRef field if non-nil, zero value otherwise.

### GetImageRefOk

`func (o *SbomIngest) GetImageRefOk() (*string, bool)`

GetImageRefOk returns a tuple with the ImageRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageRef

`func (o *SbomIngest) SetImageRef(v string)`

SetImageRef sets ImageRef field to given value.

### HasImageRef

`func (o *SbomIngest) HasImageRef() bool`

HasImageRef returns a boolean if a field has been set.

### GetSourceRepo

`func (o *SbomIngest) GetSourceRepo() string`

GetSourceRepo returns the SourceRepo field if non-nil, zero value otherwise.

### GetSourceRepoOk

`func (o *SbomIngest) GetSourceRepoOk() (*string, bool)`

GetSourceRepoOk returns a tuple with the SourceRepo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceRepo

`func (o *SbomIngest) SetSourceRepo(v string)`

SetSourceRepo sets SourceRepo field to given value.

### HasSourceRepo

`func (o *SbomIngest) HasSourceRepo() bool`

HasSourceRepo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


