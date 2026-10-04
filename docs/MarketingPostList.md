# MarketingPostList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]MarketingCalendarPost**](MarketingCalendarPost.md) | Data is the page, ordered by scheduledAt descending — the furthest-out post first and unscheduled drafts (scheduledAt 0) last. An empty array when the org&#39;s calendar holds no matching post. | [optional] 

## Methods

### NewMarketingPostList

`func NewMarketingPostList() *MarketingPostList`

NewMarketingPostList instantiates a new MarketingPostList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarketingPostListWithDefaults

`func NewMarketingPostListWithDefaults() *MarketingPostList`

NewMarketingPostListWithDefaults instantiates a new MarketingPostList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *MarketingPostList) GetData() []MarketingCalendarPost`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *MarketingPostList) GetDataOk() (*[]MarketingCalendarPost, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *MarketingPostList) SetData(v []MarketingCalendarPost)`

SetData sets Data field to given value.

### HasData

`func (o *MarketingPostList) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


