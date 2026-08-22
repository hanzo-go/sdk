# SeoAuditOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Checks** | Pointer to **map[string]bool** | Checks is every named finding, each a yes or no — \&quot;is_https\&quot;, \&quot;no_h1_tag\&quot;, \&quot;high_loading_time\&quot;, around fifty of them. It is an OPEN set the upstream adds to, so it is published as an object of booleans rather than as fifty declared fields that would be wrong the next time they name a fifty-first. | [optional] 
**Cost** | Pointer to **string** | Cost is what this call cost, in USD, as an exact decimal string. | [optional] 
**Description** | Pointer to **string** | Description is its meta description. | [optional] 
**Score** | Pointer to **float32** | Score is the upstream&#39;s on-page score, 0 to 100. | [optional] 
**Status** | Pointer to **int32** | Status is the HTTP status the page answered with. | [optional] 
**Title** | Pointer to **string** | Title is the page&#39;s title. | [optional] 
**Url** | Pointer to **string** | URL is the address actually read, after redirects. | [optional] 
**Words** | Pointer to **int32** | Words is how many words of readable text the page carries. | [optional] 

## Methods

### NewSeoAuditOut

`func NewSeoAuditOut() *SeoAuditOut`

NewSeoAuditOut instantiates a new SeoAuditOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSeoAuditOutWithDefaults

`func NewSeoAuditOutWithDefaults() *SeoAuditOut`

NewSeoAuditOutWithDefaults instantiates a new SeoAuditOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChecks

`func (o *SeoAuditOut) GetChecks() map[string]bool`

GetChecks returns the Checks field if non-nil, zero value otherwise.

### GetChecksOk

`func (o *SeoAuditOut) GetChecksOk() (*map[string]bool, bool)`

GetChecksOk returns a tuple with the Checks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChecks

`func (o *SeoAuditOut) SetChecks(v map[string]bool)`

SetChecks sets Checks field to given value.

### HasChecks

`func (o *SeoAuditOut) HasChecks() bool`

HasChecks returns a boolean if a field has been set.

### GetCost

`func (o *SeoAuditOut) GetCost() string`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *SeoAuditOut) GetCostOk() (*string, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *SeoAuditOut) SetCost(v string)`

SetCost sets Cost field to given value.

### HasCost

`func (o *SeoAuditOut) HasCost() bool`

HasCost returns a boolean if a field has been set.

### GetDescription

`func (o *SeoAuditOut) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SeoAuditOut) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SeoAuditOut) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SeoAuditOut) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetScore

`func (o *SeoAuditOut) GetScore() float32`

GetScore returns the Score field if non-nil, zero value otherwise.

### GetScoreOk

`func (o *SeoAuditOut) GetScoreOk() (*float32, bool)`

GetScoreOk returns a tuple with the Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScore

`func (o *SeoAuditOut) SetScore(v float32)`

SetScore sets Score field to given value.

### HasScore

`func (o *SeoAuditOut) HasScore() bool`

HasScore returns a boolean if a field has been set.

### GetStatus

`func (o *SeoAuditOut) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SeoAuditOut) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SeoAuditOut) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SeoAuditOut) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTitle

`func (o *SeoAuditOut) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *SeoAuditOut) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *SeoAuditOut) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *SeoAuditOut) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetUrl

`func (o *SeoAuditOut) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *SeoAuditOut) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *SeoAuditOut) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *SeoAuditOut) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetWords

`func (o *SeoAuditOut) GetWords() int32`

GetWords returns the Words field if non-nil, zero value otherwise.

### GetWordsOk

`func (o *SeoAuditOut) GetWordsOk() (*int32, bool)`

GetWordsOk returns a tuple with the Words field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWords

`func (o *SeoAuditOut) SetWords(v int32)`

SetWords sets Words field to given value.

### HasWords

`func (o *SeoAuditOut) HasWords() bool`

HasWords returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


