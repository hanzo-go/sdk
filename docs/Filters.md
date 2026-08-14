# Filters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Keywords** | Pointer to **[]string** | Keywords keeps only items whose TITLE contains one of these, case-insensitively. They are also the GDELT queries the feed fans out to, one per keyword of three characters or more — so a keyword both widens what is fetched and narrows what is kept. | [optional] 
**Regions** | Pointer to **[]string** | Regions keeps only items whose TITLE contains one of these, matched case-insensitively as a substring. Empty keeps every region. | [optional] 
**Sources** | Pointer to **[]string** | Sources keeps only items whose outlet name contains one of these, case-insensitively. Empty keeps every outlet. | [optional] 

## Methods

### NewFilters

`func NewFilters() *Filters`

NewFilters instantiates a new Filters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFiltersWithDefaults

`func NewFiltersWithDefaults() *Filters`

NewFiltersWithDefaults instantiates a new Filters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKeywords

`func (o *Filters) GetKeywords() []string`

GetKeywords returns the Keywords field if non-nil, zero value otherwise.

### GetKeywordsOk

`func (o *Filters) GetKeywordsOk() (*[]string, bool)`

GetKeywordsOk returns a tuple with the Keywords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeywords

`func (o *Filters) SetKeywords(v []string)`

SetKeywords sets Keywords field to given value.

### HasKeywords

`func (o *Filters) HasKeywords() bool`

HasKeywords returns a boolean if a field has been set.

### GetRegions

`func (o *Filters) GetRegions() []string`

GetRegions returns the Regions field if non-nil, zero value otherwise.

### GetRegionsOk

`func (o *Filters) GetRegionsOk() (*[]string, bool)`

GetRegionsOk returns a tuple with the Regions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegions

`func (o *Filters) SetRegions(v []string)`

SetRegions sets Regions field to given value.

### HasRegions

`func (o *Filters) HasRegions() bool`

HasRegions returns a boolean if a field has been set.

### GetSources

`func (o *Filters) GetSources() []string`

GetSources returns the Sources field if non-nil, zero value otherwise.

### GetSourcesOk

`func (o *Filters) GetSourcesOk() (*[]string, bool)`

GetSourcesOk returns a tuple with the Sources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSources

`func (o *Filters) SetSources(v []string)`

SetSources sets Sources field to given value.

### HasSources

`func (o *Filters) HasSources() bool`

HasSources returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


