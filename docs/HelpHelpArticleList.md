# HelpHelpArticleList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]HelpHelpArticleCard**](HelpHelpArticleCard.md) | Data is the matching Published, public articles, newest write order last — the store&#39;s order, not a ranking. Empty when the center has none. | [optional] 

## Methods

### NewHelpHelpArticleList

`func NewHelpHelpArticleList() *HelpHelpArticleList`

NewHelpHelpArticleList instantiates a new HelpHelpArticleList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHelpHelpArticleListWithDefaults

`func NewHelpHelpArticleListWithDefaults() *HelpHelpArticleList`

NewHelpHelpArticleListWithDefaults instantiates a new HelpHelpArticleList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *HelpHelpArticleList) GetData() []HelpHelpArticleCard`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *HelpHelpArticleList) GetDataOk() (*[]HelpHelpArticleCard, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *HelpHelpArticleList) SetData(v []HelpHelpArticleCard)`

SetData sets Data field to given value.

### HasData

`func (o *HelpHelpArticleList) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


