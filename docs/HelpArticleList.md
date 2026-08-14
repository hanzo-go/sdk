# HelpArticleList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]HelpArticleCard**](HelpArticleCard.md) | Data is the matching Published, public articles, newest write order last — the store&#39;s order, not a ranking. Empty when the center has none. | [optional] 

## Methods

### NewHelpArticleList

`func NewHelpArticleList() *HelpArticleList`

NewHelpArticleList instantiates a new HelpArticleList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHelpArticleListWithDefaults

`func NewHelpArticleListWithDefaults() *HelpArticleList`

NewHelpArticleListWithDefaults instantiates a new HelpArticleList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *HelpArticleList) GetData() []HelpArticleCard`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *HelpArticleList) GetDataOk() (*[]HelpArticleCard, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *HelpArticleList) SetData(v []HelpArticleCard)`

SetData sets Data field to given value.

### HasData

`func (o *HelpArticleList) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


