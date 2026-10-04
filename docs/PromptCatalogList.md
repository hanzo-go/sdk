# PromptCatalogList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]PromptCatalogEntry**](PromptCatalogEntry.md) | Data is every starter prompt, each importable as-is with POST /v1/prompt. | [optional] 

## Methods

### NewPromptCatalogList

`func NewPromptCatalogList() *PromptCatalogList`

NewPromptCatalogList instantiates a new PromptCatalogList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPromptCatalogListWithDefaults

`func NewPromptCatalogListWithDefaults() *PromptCatalogList`

NewPromptCatalogListWithDefaults instantiates a new PromptCatalogList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *PromptCatalogList) GetData() []PromptCatalogEntry`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *PromptCatalogList) GetDataOk() (*[]PromptCatalogEntry, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *PromptCatalogList) SetData(v []PromptCatalogEntry)`

SetData sets Data field to given value.

### HasData

`func (o *PromptCatalogList) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


