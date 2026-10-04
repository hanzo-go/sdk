# MarketingSuppressionList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]MarketingSuppression**](MarketingSuppression.md) | Data is the page: every (channel, address) this org&#39;s send gate refuses, newest opt-out first. Absence from it is not permission to mail someone — it only means no opt-out was recorded on that channel. | [optional] 

## Methods

### NewMarketingSuppressionList

`func NewMarketingSuppressionList() *MarketingSuppressionList`

NewMarketingSuppressionList instantiates a new MarketingSuppressionList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarketingSuppressionListWithDefaults

`func NewMarketingSuppressionListWithDefaults() *MarketingSuppressionList`

NewMarketingSuppressionListWithDefaults instantiates a new MarketingSuppressionList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *MarketingSuppressionList) GetData() []MarketingSuppression`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *MarketingSuppressionList) GetDataOk() (*[]MarketingSuppression, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *MarketingSuppressionList) SetData(v []MarketingSuppression)`

SetData sets Data field to given value.

### HasData

`func (o *MarketingSuppressionList) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


