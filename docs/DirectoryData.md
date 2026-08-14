# DirectoryData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Affiliates** | Pointer to [**[]AdminAffiliateView**](AdminAffiliateView.md) | Affiliates is one row per affiliate across the whole fleet, ORG EXPOSED, oldest first and bounded by the request&#39;s limit. | [optional] 
**Summary** | Pointer to [**Totals**](Totals.md) | Summary tallies exactly the rows above — not the whole table — so a limit that truncates the page truncates the tally with it. | [optional] 

## Methods

### NewDirectoryData

`func NewDirectoryData() *DirectoryData`

NewDirectoryData instantiates a new DirectoryData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDirectoryDataWithDefaults

`func NewDirectoryDataWithDefaults() *DirectoryData`

NewDirectoryDataWithDefaults instantiates a new DirectoryData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAffiliates

`func (o *DirectoryData) GetAffiliates() []AdminAffiliateView`

GetAffiliates returns the Affiliates field if non-nil, zero value otherwise.

### GetAffiliatesOk

`func (o *DirectoryData) GetAffiliatesOk() (*[]AdminAffiliateView, bool)`

GetAffiliatesOk returns a tuple with the Affiliates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffiliates

`func (o *DirectoryData) SetAffiliates(v []AdminAffiliateView)`

SetAffiliates sets Affiliates field to given value.

### HasAffiliates

`func (o *DirectoryData) HasAffiliates() bool`

HasAffiliates returns a boolean if a field has been set.

### GetSummary

`func (o *DirectoryData) GetSummary() Totals`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *DirectoryData) GetSummaryOk() (*Totals, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *DirectoryData) SetSummary(v Totals)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *DirectoryData) HasSummary() bool`

HasSummary returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


