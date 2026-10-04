# PromptMetricRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **string** | CreatedAt is when version 1 was written, RFC 3339 UTC. | [optional] 
**CurrentVersion** | Pointer to **int64** | CurrentVer is the version number served as current. It always equals &#x60;versions&#x60;: numbering is dense from 1, and deleting a prompt takes its whole history with it rather than leaving a gap. | [optional] 
**LastUpdatedAt** | Pointer to **string** | LastUpdatedAt is when the newest version was appended, RFC 3339 UTC — the age of the template you would get today. | [optional] 
**Name** | Pointer to **string** | Name is the prompt this row is about — its org-unique handle. | [optional] 
**Type** | Pointer to **string** | Type is the current version&#39;s kind. | [optional] 
**Versions** | Pointer to **int64** | Versions is how many revisions the prompt has, COUNTED in the store and uncapped — so it can exceed the 100 entries a list row or a detail response carries. Note the type: here &#x60;versions&#x60; is a number, while on a list row it is the list of version numbers. | [optional] 

## Methods

### NewPromptMetricRow

`func NewPromptMetricRow() *PromptMetricRow`

NewPromptMetricRow instantiates a new PromptMetricRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPromptMetricRowWithDefaults

`func NewPromptMetricRowWithDefaults() *PromptMetricRow`

NewPromptMetricRowWithDefaults instantiates a new PromptMetricRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *PromptMetricRow) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PromptMetricRow) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PromptMetricRow) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *PromptMetricRow) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCurrentVersion

`func (o *PromptMetricRow) GetCurrentVersion() int64`

GetCurrentVersion returns the CurrentVersion field if non-nil, zero value otherwise.

### GetCurrentVersionOk

`func (o *PromptMetricRow) GetCurrentVersionOk() (*int64, bool)`

GetCurrentVersionOk returns a tuple with the CurrentVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentVersion

`func (o *PromptMetricRow) SetCurrentVersion(v int64)`

SetCurrentVersion sets CurrentVersion field to given value.

### HasCurrentVersion

`func (o *PromptMetricRow) HasCurrentVersion() bool`

HasCurrentVersion returns a boolean if a field has been set.

### GetLastUpdatedAt

`func (o *PromptMetricRow) GetLastUpdatedAt() string`

GetLastUpdatedAt returns the LastUpdatedAt field if non-nil, zero value otherwise.

### GetLastUpdatedAtOk

`func (o *PromptMetricRow) GetLastUpdatedAtOk() (*string, bool)`

GetLastUpdatedAtOk returns a tuple with the LastUpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedAt

`func (o *PromptMetricRow) SetLastUpdatedAt(v string)`

SetLastUpdatedAt sets LastUpdatedAt field to given value.

### HasLastUpdatedAt

`func (o *PromptMetricRow) HasLastUpdatedAt() bool`

HasLastUpdatedAt returns a boolean if a field has been set.

### GetName

`func (o *PromptMetricRow) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PromptMetricRow) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PromptMetricRow) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PromptMetricRow) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *PromptMetricRow) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PromptMetricRow) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PromptMetricRow) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *PromptMetricRow) HasType() bool`

HasType returns a boolean if a field has been set.

### GetVersions

`func (o *PromptMetricRow) GetVersions() int64`

GetVersions returns the Versions field if non-nil, zero value otherwise.

### GetVersionsOk

`func (o *PromptMetricRow) GetVersionsOk() (*int64, bool)`

GetVersionsOk returns a tuple with the Versions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersions

`func (o *PromptMetricRow) SetVersions(v int64)`

SetVersions sets Versions field to given value.

### HasVersions

`func (o *PromptMetricRow) HasVersions() bool`

HasVersions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


