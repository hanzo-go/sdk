# SeoBacklinkOut

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

### NewSeoBacklinkOut

`func NewSeoBacklinkOut() *SeoBacklinkOut`

NewSeoBacklinkOut instantiates a new SeoBacklinkOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSeoBacklinkOutWithDefaults

`func NewSeoBacklinkOutWithDefaults() *SeoBacklinkOut`

NewSeoBacklinkOutWithDefaults instantiates a new SeoBacklinkOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBacklinks

`func (o *SeoBacklinkOut) GetBacklinks() int64`

GetBacklinks returns the Backlinks field if non-nil, zero value otherwise.

### GetBacklinksOk

`func (o *SeoBacklinkOut) GetBacklinksOk() (*int64, bool)`

GetBacklinksOk returns a tuple with the Backlinks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBacklinks

`func (o *SeoBacklinkOut) SetBacklinks(v int64)`

SetBacklinks sets Backlinks field to given value.

### HasBacklinks

`func (o *SeoBacklinkOut) HasBacklinks() bool`

HasBacklinks returns a boolean if a field has been set.

### GetBroken

`func (o *SeoBacklinkOut) GetBroken() int64`

GetBroken returns the Broken field if non-nil, zero value otherwise.

### GetBrokenOk

`func (o *SeoBacklinkOut) GetBrokenOk() (*int64, bool)`

GetBrokenOk returns a tuple with the Broken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBroken

`func (o *SeoBacklinkOut) SetBroken(v int64)`

SetBroken sets Broken field to given value.

### HasBroken

`func (o *SeoBacklinkOut) HasBroken() bool`

HasBroken returns a boolean if a field has been set.

### GetCost

`func (o *SeoBacklinkOut) GetCost() string`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *SeoBacklinkOut) GetCostOk() (*string, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *SeoBacklinkOut) SetCost(v string)`

SetCost sets Cost field to given value.

### HasCost

`func (o *SeoBacklinkOut) HasCost() bool`

HasCost returns a boolean if a field has been set.

### GetDomains

`func (o *SeoBacklinkOut) GetDomains() int64`

GetDomains returns the Domains field if non-nil, zero value otherwise.

### GetDomainsOk

`func (o *SeoBacklinkOut) GetDomainsOk() (*int64, bool)`

GetDomainsOk returns a tuple with the Domains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomains

`func (o *SeoBacklinkOut) SetDomains(v int64)`

SetDomains sets Domains field to given value.

### HasDomains

`func (o *SeoBacklinkOut) HasDomains() bool`

HasDomains returns a boolean if a field has been set.

### GetFirstSeen

`func (o *SeoBacklinkOut) GetFirstSeen() string`

GetFirstSeen returns the FirstSeen field if non-nil, zero value otherwise.

### GetFirstSeenOk

`func (o *SeoBacklinkOut) GetFirstSeenOk() (*string, bool)`

GetFirstSeenOk returns a tuple with the FirstSeen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstSeen

`func (o *SeoBacklinkOut) SetFirstSeen(v string)`

SetFirstSeen sets FirstSeen field to given value.

### HasFirstSeen

`func (o *SeoBacklinkOut) HasFirstSeen() bool`

HasFirstSeen returns a boolean if a field has been set.

### GetPages

`func (o *SeoBacklinkOut) GetPages() int64`

GetPages returns the Pages field if non-nil, zero value otherwise.

### GetPagesOk

`func (o *SeoBacklinkOut) GetPagesOk() (*int64, bool)`

GetPagesOk returns a tuple with the Pages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPages

`func (o *SeoBacklinkOut) SetPages(v int64)`

SetPages sets Pages field to given value.

### HasPages

`func (o *SeoBacklinkOut) HasPages() bool`

HasPages returns a boolean if a field has been set.

### GetRank

`func (o *SeoBacklinkOut) GetRank() int64`

GetRank returns the Rank field if non-nil, zero value otherwise.

### GetRankOk

`func (o *SeoBacklinkOut) GetRankOk() (*int64, bool)`

GetRankOk returns a tuple with the Rank field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRank

`func (o *SeoBacklinkOut) SetRank(v int64)`

SetRank sets Rank field to given value.

### HasRank

`func (o *SeoBacklinkOut) HasRank() bool`

HasRank returns a boolean if a field has been set.

### GetSpam

`func (o *SeoBacklinkOut) GetSpam() int64`

GetSpam returns the Spam field if non-nil, zero value otherwise.

### GetSpamOk

`func (o *SeoBacklinkOut) GetSpamOk() (*int64, bool)`

GetSpamOk returns a tuple with the Spam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpam

`func (o *SeoBacklinkOut) SetSpam(v int64)`

SetSpam sets Spam field to given value.

### HasSpam

`func (o *SeoBacklinkOut) HasSpam() bool`

HasSpam returns a boolean if a field has been set.

### GetTarget

`func (o *SeoBacklinkOut) GetTarget() string`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *SeoBacklinkOut) GetTargetOk() (*string, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *SeoBacklinkOut) SetTarget(v string)`

SetTarget sets Target field to given value.

### HasTarget

`func (o *SeoBacklinkOut) HasTarget() bool`

HasTarget returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


