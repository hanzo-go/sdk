# PromptPromptList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]PromptPromptMeta**](PromptPromptMeta.md) | Data is one row per prompt the org owns, each with its version numbers and taxonomy — never the template bodies. | [optional] 

## Methods

### NewPromptPromptList

`func NewPromptPromptList() *PromptPromptList`

NewPromptPromptList instantiates a new PromptPromptList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPromptPromptListWithDefaults

`func NewPromptPromptListWithDefaults() *PromptPromptList`

NewPromptPromptListWithDefaults instantiates a new PromptPromptList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *PromptPromptList) GetData() []PromptPromptMeta`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *PromptPromptList) GetDataOk() (*[]PromptPromptMeta, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *PromptPromptList) SetData(v []PromptPromptMeta)`

SetData sets Data field to given value.

### HasData

`func (o *PromptPromptList) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


