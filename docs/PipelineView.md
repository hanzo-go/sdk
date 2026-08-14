# PipelineView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **string** | CreatedAt is when the pipeline was first stored, RFC3339 UTC. Absent on the default. | [optional] 
**Default** | Pointer to **bool** | Default is true when no pipeline is stored for this project and these are the built-in world feeds. Writing one turns it false. | [optional] 
**Feeds** | Pointer to **[]string** | Feeds is the RSS/Atom feed URLs the pipeline reads. Every host is on the server&#39;s allowlist — a URL that is not cannot be stored. | [optional] 
**Filters** | Pointer to [**Filters**](Filters.md) | Filters narrows the merged feed. | [optional] 
**Org** | Pointer to **string** | Org is the tenant the pipeline belongs to, resolved server-side from the validated principal. | [optional] 
**Project** | Pointer to **string** | Project is the org sub-scope the pipeline belongs to. | [optional] 
**UpdatedAt** | Pointer to **string** | UpdatedAt is when it was last written, RFC3339 UTC. Absent on the default. | [optional] 

## Methods

### NewPipelineView

`func NewPipelineView() *PipelineView`

NewPipelineView instantiates a new PipelineView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPipelineViewWithDefaults

`func NewPipelineViewWithDefaults() *PipelineView`

NewPipelineViewWithDefaults instantiates a new PipelineView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *PipelineView) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PipelineView) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PipelineView) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *PipelineView) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDefault

`func (o *PipelineView) GetDefault() bool`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *PipelineView) GetDefaultOk() (*bool, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *PipelineView) SetDefault(v bool)`

SetDefault sets Default field to given value.

### HasDefault

`func (o *PipelineView) HasDefault() bool`

HasDefault returns a boolean if a field has been set.

### GetFeeds

`func (o *PipelineView) GetFeeds() []string`

GetFeeds returns the Feeds field if non-nil, zero value otherwise.

### GetFeedsOk

`func (o *PipelineView) GetFeedsOk() (*[]string, bool)`

GetFeedsOk returns a tuple with the Feeds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeds

`func (o *PipelineView) SetFeeds(v []string)`

SetFeeds sets Feeds field to given value.

### HasFeeds

`func (o *PipelineView) HasFeeds() bool`

HasFeeds returns a boolean if a field has been set.

### GetFilters

`func (o *PipelineView) GetFilters() Filters`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *PipelineView) GetFiltersOk() (*Filters, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *PipelineView) SetFilters(v Filters)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *PipelineView) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### GetOrg

`func (o *PipelineView) GetOrg() string`

GetOrg returns the Org field if non-nil, zero value otherwise.

### GetOrgOk

`func (o *PipelineView) GetOrgOk() (*string, bool)`

GetOrgOk returns a tuple with the Org field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrg

`func (o *PipelineView) SetOrg(v string)`

SetOrg sets Org field to given value.

### HasOrg

`func (o *PipelineView) HasOrg() bool`

HasOrg returns a boolean if a field has been set.

### GetProject

`func (o *PipelineView) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *PipelineView) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *PipelineView) SetProject(v string)`

SetProject sets Project field to given value.

### HasProject

`func (o *PipelineView) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *PipelineView) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *PipelineView) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *PipelineView) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *PipelineView) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


