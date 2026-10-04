# WorldPipelineView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **string** | CreatedAt is when the pipeline was first stored, RFC3339 UTC. Absent on the default. | [optional] 
**Default** | Pointer to **bool** | Default is true when no pipeline is stored for this project and these are the built-in world feeds. Writing one turns it false. | [optional] 
**Feeds** | Pointer to **[]string** | Feeds is the RSS/Atom feed URLs the pipeline reads. Every host is on the server&#39;s allowlist — a URL that is not cannot be stored. | [optional] 
**Filters** | Pointer to [**WorldFilters**](WorldFilters.md) | Filters narrows the merged feed. | [optional] 
**Org** | Pointer to **string** | Org is the tenant the pipeline belongs to, resolved server-side from the validated principal. | [optional] 
**Project** | Pointer to **string** | Project is the org sub-scope the pipeline belongs to. | [optional] 
**UpdatedAt** | Pointer to **string** | UpdatedAt is when it was last written, RFC3339 UTC. Absent on the default. | [optional] 

## Methods

### NewWorldPipelineView

`func NewWorldPipelineView() *WorldPipelineView`

NewWorldPipelineView instantiates a new WorldPipelineView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorldPipelineViewWithDefaults

`func NewWorldPipelineViewWithDefaults() *WorldPipelineView`

NewWorldPipelineViewWithDefaults instantiates a new WorldPipelineView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *WorldPipelineView) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *WorldPipelineView) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *WorldPipelineView) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *WorldPipelineView) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDefault

`func (o *WorldPipelineView) GetDefault() bool`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *WorldPipelineView) GetDefaultOk() (*bool, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *WorldPipelineView) SetDefault(v bool)`

SetDefault sets Default field to given value.

### HasDefault

`func (o *WorldPipelineView) HasDefault() bool`

HasDefault returns a boolean if a field has been set.

### GetFeeds

`func (o *WorldPipelineView) GetFeeds() []string`

GetFeeds returns the Feeds field if non-nil, zero value otherwise.

### GetFeedsOk

`func (o *WorldPipelineView) GetFeedsOk() (*[]string, bool)`

GetFeedsOk returns a tuple with the Feeds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeds

`func (o *WorldPipelineView) SetFeeds(v []string)`

SetFeeds sets Feeds field to given value.

### HasFeeds

`func (o *WorldPipelineView) HasFeeds() bool`

HasFeeds returns a boolean if a field has been set.

### GetFilters

`func (o *WorldPipelineView) GetFilters() WorldFilters`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *WorldPipelineView) GetFiltersOk() (*WorldFilters, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *WorldPipelineView) SetFilters(v WorldFilters)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *WorldPipelineView) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### GetOrg

`func (o *WorldPipelineView) GetOrg() string`

GetOrg returns the Org field if non-nil, zero value otherwise.

### GetOrgOk

`func (o *WorldPipelineView) GetOrgOk() (*string, bool)`

GetOrgOk returns a tuple with the Org field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrg

`func (o *WorldPipelineView) SetOrg(v string)`

SetOrg sets Org field to given value.

### HasOrg

`func (o *WorldPipelineView) HasOrg() bool`

HasOrg returns a boolean if a field has been set.

### GetProject

`func (o *WorldPipelineView) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *WorldPipelineView) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *WorldPipelineView) SetProject(v string)`

SetProject sets Project field to given value.

### HasProject

`func (o *WorldPipelineView) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *WorldPipelineView) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *WorldPipelineView) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *WorldPipelineView) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *WorldPipelineView) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


