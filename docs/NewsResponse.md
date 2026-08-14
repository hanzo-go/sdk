# NewsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | Pointer to [**[]NewsItem**](NewsItem.md) | Items is the merged, filtered, deduped feed, freshest first and capped at 50. A source that failed is skipped rather than failing the read, so this can be shorter than the pipeline&#39;s reach — it is never an error. | [optional] 

## Methods

### NewNewsResponse

`func NewNewsResponse() *NewsResponse`

NewNewsResponse instantiates a new NewsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNewsResponseWithDefaults

`func NewNewsResponseWithDefaults() *NewsResponse`

NewNewsResponseWithDefaults instantiates a new NewsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *NewsResponse) GetItems() []NewsItem`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *NewsResponse) GetItemsOk() (*[]NewsItem, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *NewsResponse) SetItems(v []NewsItem)`

SetItems sets Items field to given value.

### HasItems

`func (o *NewsResponse) HasItems() bool`

HasItems returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


