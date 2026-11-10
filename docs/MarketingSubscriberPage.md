# MarketingSubscriberPage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]MarketingSubscriber**](MarketingSubscriber.md) | Data is the page. | [optional] 
**Next** | Pointer to **string** | Next continues the walk; empty when this is the last page. | [optional] 
**Total** | Pointer to **int64** | Total is how many subscribers match, when the list knows without reading them: the whole list or a set of statuses. Absent otherwise. | [optional] 

## Methods

### NewMarketingSubscriberPage

`func NewMarketingSubscriberPage() *MarketingSubscriberPage`

NewMarketingSubscriberPage instantiates a new MarketingSubscriberPage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarketingSubscriberPageWithDefaults

`func NewMarketingSubscriberPageWithDefaults() *MarketingSubscriberPage`

NewMarketingSubscriberPageWithDefaults instantiates a new MarketingSubscriberPage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *MarketingSubscriberPage) GetData() []MarketingSubscriber`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *MarketingSubscriberPage) GetDataOk() (*[]MarketingSubscriber, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *MarketingSubscriberPage) SetData(v []MarketingSubscriber)`

SetData sets Data field to given value.

### HasData

`func (o *MarketingSubscriberPage) HasData() bool`

HasData returns a boolean if a field has been set.

### GetNext

`func (o *MarketingSubscriberPage) GetNext() string`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *MarketingSubscriberPage) GetNextOk() (*string, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *MarketingSubscriberPage) SetNext(v string)`

SetNext sets Next field to given value.

### HasNext

`func (o *MarketingSubscriberPage) HasNext() bool`

HasNext returns a boolean if a field has been set.

### GetTotal

`func (o *MarketingSubscriberPage) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *MarketingSubscriberPage) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *MarketingSubscriberPage) SetTotal(v int64)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *MarketingSubscriberPage) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


