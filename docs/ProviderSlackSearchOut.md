# ProviderSlackSearchOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Matches** | Pointer to [**[]ProviderSlackSearchHit**](ProviderSlackSearchHit.md) | Matches are this page&#39;s hits, most relevant first. | [optional] 
**Page** | Pointer to **int64** | Page is the page these matches came from. | [optional] 
**Pages** | Pointer to **int64** | Pages is how many pages the query has in total. | [optional] 
**Total** | Pointer to **int64** | Total is how many messages matched, across every page. | [optional] 

## Methods

### NewProviderSlackSearchOut

`func NewProviderSlackSearchOut() *ProviderSlackSearchOut`

NewProviderSlackSearchOut instantiates a new ProviderSlackSearchOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProviderSlackSearchOutWithDefaults

`func NewProviderSlackSearchOutWithDefaults() *ProviderSlackSearchOut`

NewProviderSlackSearchOutWithDefaults instantiates a new ProviderSlackSearchOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMatches

`func (o *ProviderSlackSearchOut) GetMatches() []ProviderSlackSearchHit`

GetMatches returns the Matches field if non-nil, zero value otherwise.

### GetMatchesOk

`func (o *ProviderSlackSearchOut) GetMatchesOk() (*[]ProviderSlackSearchHit, bool)`

GetMatchesOk returns a tuple with the Matches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatches

`func (o *ProviderSlackSearchOut) SetMatches(v []ProviderSlackSearchHit)`

SetMatches sets Matches field to given value.

### HasMatches

`func (o *ProviderSlackSearchOut) HasMatches() bool`

HasMatches returns a boolean if a field has been set.

### GetPage

`func (o *ProviderSlackSearchOut) GetPage() int64`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *ProviderSlackSearchOut) GetPageOk() (*int64, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *ProviderSlackSearchOut) SetPage(v int64)`

SetPage sets Page field to given value.

### HasPage

`func (o *ProviderSlackSearchOut) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetPages

`func (o *ProviderSlackSearchOut) GetPages() int64`

GetPages returns the Pages field if non-nil, zero value otherwise.

### GetPagesOk

`func (o *ProviderSlackSearchOut) GetPagesOk() (*int64, bool)`

GetPagesOk returns a tuple with the Pages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPages

`func (o *ProviderSlackSearchOut) SetPages(v int64)`

SetPages sets Pages field to given value.

### HasPages

`func (o *ProviderSlackSearchOut) HasPages() bool`

HasPages returns a boolean if a field has been set.

### GetTotal

`func (o *ProviderSlackSearchOut) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *ProviderSlackSearchOut) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *ProviderSlackSearchOut) SetTotal(v int64)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *ProviderSlackSearchOut) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


