# SeoSeoBacklinkOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Backlinks** | Pointer to **int64** | Backlinks is how many links point at it. | [optional] 
**Broken** | Pointer to **int64** | Broken is how many of those links point at something that no longer answers. | [optional] 
**Cost** | Pointer to **string** | Cost is what this call cost, in USD, as an exact decimal string. | [optional] 
**Domains** | Pointer to **int64** | Domains is how many distinct sites those links come from — the number that matters, since a thousand links from one site is one site. | [optional] 
**FirstSeen** | Pointer to **string** | FirstSeen is when the upstream first saw a link to this target, RFC 3339. | [optional] 
**Pages** | Pointer to **int64** | Pages is how many distinct pages link in. | [optional] 
**Rank** | Pointer to **int64** | Rank is the upstream&#39;s authority score for the target, 0 to 1000. | [optional] 
**Spam** | Pointer to **int64** | Spam is the share of the profile judged spam, 0 to 100. | [optional] 
**Target** | Pointer to **string** | Target is the target as the upstream resolved it. | [optional] 

## Methods

### NewSeoSeoBacklinkOut

`func NewSeoSeoBacklinkOut() *SeoSeoBacklinkOut`

NewSeoSeoBacklinkOut instantiates a new SeoSeoBacklinkOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSeoSeoBacklinkOutWithDefaults

`func NewSeoSeoBacklinkOutWithDefaults() *SeoSeoBacklinkOut`

NewSeoSeoBacklinkOutWithDefaults instantiates a new SeoSeoBacklinkOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBacklinks

`func (o *SeoSeoBacklinkOut) GetBacklinks() int64`

GetBacklinks returns the Backlinks field if non-nil, zero value otherwise.

### GetBacklinksOk

`func (o *SeoSeoBacklinkOut) GetBacklinksOk() (*int64, bool)`

GetBacklinksOk returns a tuple with the Backlinks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBacklinks

`func (o *SeoSeoBacklinkOut) SetBacklinks(v int64)`

SetBacklinks sets Backlinks field to given value.

### HasBacklinks

`func (o *SeoSeoBacklinkOut) HasBacklinks() bool`

HasBacklinks returns a boolean if a field has been set.

### GetBroken

`func (o *SeoSeoBacklinkOut) GetBroken() int64`

GetBroken returns the Broken field if non-nil, zero value otherwise.

### GetBrokenOk

`func (o *SeoSeoBacklinkOut) GetBrokenOk() (*int64, bool)`

GetBrokenOk returns a tuple with the Broken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBroken

`func (o *SeoSeoBacklinkOut) SetBroken(v int64)`

SetBroken sets Broken field to given value.

### HasBroken

`func (o *SeoSeoBacklinkOut) HasBroken() bool`

HasBroken returns a boolean if a field has been set.

### GetCost

`func (o *SeoSeoBacklinkOut) GetCost() string`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *SeoSeoBacklinkOut) GetCostOk() (*string, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *SeoSeoBacklinkOut) SetCost(v string)`

SetCost sets Cost field to given value.

### HasCost

`func (o *SeoSeoBacklinkOut) HasCost() bool`

HasCost returns a boolean if a field has been set.

### GetDomains

`func (o *SeoSeoBacklinkOut) GetDomains() int64`

GetDomains returns the Domains field if non-nil, zero value otherwise.

### GetDomainsOk

`func (o *SeoSeoBacklinkOut) GetDomainsOk() (*int64, bool)`

GetDomainsOk returns a tuple with the Domains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomains

`func (o *SeoSeoBacklinkOut) SetDomains(v int64)`

SetDomains sets Domains field to given value.

### HasDomains

`func (o *SeoSeoBacklinkOut) HasDomains() bool`

HasDomains returns a boolean if a field has been set.

### GetFirstSeen

`func (o *SeoSeoBacklinkOut) GetFirstSeen() string`

GetFirstSeen returns the FirstSeen field if non-nil, zero value otherwise.

### GetFirstSeenOk

`func (o *SeoSeoBacklinkOut) GetFirstSeenOk() (*string, bool)`

GetFirstSeenOk returns a tuple with the FirstSeen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstSeen

`func (o *SeoSeoBacklinkOut) SetFirstSeen(v string)`

SetFirstSeen sets FirstSeen field to given value.

### HasFirstSeen

`func (o *SeoSeoBacklinkOut) HasFirstSeen() bool`

HasFirstSeen returns a boolean if a field has been set.

### GetPages

`func (o *SeoSeoBacklinkOut) GetPages() int64`

GetPages returns the Pages field if non-nil, zero value otherwise.

### GetPagesOk

`func (o *SeoSeoBacklinkOut) GetPagesOk() (*int64, bool)`

GetPagesOk returns a tuple with the Pages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPages

`func (o *SeoSeoBacklinkOut) SetPages(v int64)`

SetPages sets Pages field to given value.

### HasPages

`func (o *SeoSeoBacklinkOut) HasPages() bool`

HasPages returns a boolean if a field has been set.

### GetRank

`func (o *SeoSeoBacklinkOut) GetRank() int64`

GetRank returns the Rank field if non-nil, zero value otherwise.

### GetRankOk

`func (o *SeoSeoBacklinkOut) GetRankOk() (*int64, bool)`

GetRankOk returns a tuple with the Rank field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRank

`func (o *SeoSeoBacklinkOut) SetRank(v int64)`

SetRank sets Rank field to given value.

### HasRank

`func (o *SeoSeoBacklinkOut) HasRank() bool`

HasRank returns a boolean if a field has been set.

### GetSpam

`func (o *SeoSeoBacklinkOut) GetSpam() int64`

GetSpam returns the Spam field if non-nil, zero value otherwise.

### GetSpamOk

`func (o *SeoSeoBacklinkOut) GetSpamOk() (*int64, bool)`

GetSpamOk returns a tuple with the Spam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpam

`func (o *SeoSeoBacklinkOut) SetSpam(v int64)`

SetSpam sets Spam field to given value.

### HasSpam

`func (o *SeoSeoBacklinkOut) HasSpam() bool`

HasSpam returns a boolean if a field has been set.

### GetTarget

`func (o *SeoSeoBacklinkOut) GetTarget() string`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *SeoSeoBacklinkOut) GetTargetOk() (*string, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *SeoSeoBacklinkOut) SetTarget(v string)`

SetTarget sets Target field to given value.

### HasTarget

`func (o *SeoSeoBacklinkOut) HasTarget() bool`

HasTarget returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


