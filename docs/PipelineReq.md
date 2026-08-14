# PipelineReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Feeds** | Pointer to **[]string** | Feeds is the RSS/Atom feed URLs to read, at most 64. Each must be an http(s) URL whose host is on the server&#39;s allowlist — the SSRF guard is applied here, at the write, so a stored pipeline can never name a host the fetcher would refuse. Blank entries are dropped and duplicates collapse. | [optional] 
**Filters** | Pointer to [**Filters**](Filters.md) | Filters narrows the merged feed. Terms are trimmed, de-duplicated case-insensitively, and capped at 64 per axis. | [optional] 

## Methods

### NewPipelineReq

`func NewPipelineReq() *PipelineReq`

NewPipelineReq instantiates a new PipelineReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPipelineReqWithDefaults

`func NewPipelineReqWithDefaults() *PipelineReq`

NewPipelineReqWithDefaults instantiates a new PipelineReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFeeds

`func (o *PipelineReq) GetFeeds() []string`

GetFeeds returns the Feeds field if non-nil, zero value otherwise.

### GetFeedsOk

`func (o *PipelineReq) GetFeedsOk() (*[]string, bool)`

GetFeedsOk returns a tuple with the Feeds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeds

`func (o *PipelineReq) SetFeeds(v []string)`

SetFeeds sets Feeds field to given value.

### HasFeeds

`func (o *PipelineReq) HasFeeds() bool`

HasFeeds returns a boolean if a field has been set.

### GetFilters

`func (o *PipelineReq) GetFilters() Filters`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *PipelineReq) GetFiltersOk() (*Filters, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *PipelineReq) SetFilters(v Filters)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *PipelineReq) HasFilters() bool`

HasFilters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


