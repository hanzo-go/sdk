# IndexVersion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CommitDate** | Pointer to **string** | CommitDate is empty here: this surface is a dialect implementation, not a build of Meilisearch, so there is no upstream commit to date. | [optional] 
**CommitSha** | Pointer to **string** | CommitSha names the implementation (&#x60;hanzo-cloud&#x60;) rather than a build hash, so a client logging it records which server answered instead of implying a Meilisearch release. | [optional] 
**PkgVersion** | Pointer to **string** | PkgVersion is this dialect implementation&#39;s own version. | [optional] 

## Methods

### NewIndexVersion

`func NewIndexVersion() *IndexVersion`

NewIndexVersion instantiates a new IndexVersion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIndexVersionWithDefaults

`func NewIndexVersionWithDefaults() *IndexVersion`

NewIndexVersionWithDefaults instantiates a new IndexVersion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommitDate

`func (o *IndexVersion) GetCommitDate() string`

GetCommitDate returns the CommitDate field if non-nil, zero value otherwise.

### GetCommitDateOk

`func (o *IndexVersion) GetCommitDateOk() (*string, bool)`

GetCommitDateOk returns a tuple with the CommitDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitDate

`func (o *IndexVersion) SetCommitDate(v string)`

SetCommitDate sets CommitDate field to given value.

### HasCommitDate

`func (o *IndexVersion) HasCommitDate() bool`

HasCommitDate returns a boolean if a field has been set.

### GetCommitSha

`func (o *IndexVersion) GetCommitSha() string`

GetCommitSha returns the CommitSha field if non-nil, zero value otherwise.

### GetCommitShaOk

`func (o *IndexVersion) GetCommitShaOk() (*string, bool)`

GetCommitShaOk returns a tuple with the CommitSha field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitSha

`func (o *IndexVersion) SetCommitSha(v string)`

SetCommitSha sets CommitSha field to given value.

### HasCommitSha

`func (o *IndexVersion) HasCommitSha() bool`

HasCommitSha returns a boolean if a field has been set.

### GetPkgVersion

`func (o *IndexVersion) GetPkgVersion() string`

GetPkgVersion returns the PkgVersion field if non-nil, zero value otherwise.

### GetPkgVersionOk

`func (o *IndexVersion) GetPkgVersionOk() (*string, bool)`

GetPkgVersionOk returns a tuple with the PkgVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkgVersion

`func (o *IndexVersion) SetPkgVersion(v string)`

SetPkgVersion sets PkgVersion field to given value.

### HasPkgVersion

`func (o *IndexVersion) HasPkgVersion() bool`

HasPkgVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


