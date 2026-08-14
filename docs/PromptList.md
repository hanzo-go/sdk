# PromptList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]PromptMeta**](PromptMeta.md) | Data is one row per prompt the org owns, each with its version numbers and taxonomy — never the template bodies. | [optional] 

## Methods

### NewPromptList

`func NewPromptList() *PromptList`

NewPromptList instantiates a new PromptList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPromptListWithDefaults

`func NewPromptListWithDefaults() *PromptList`

NewPromptListWithDefaults instantiates a new PromptList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *PromptList) GetData() []PromptMeta`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *PromptList) GetDataOk() (*[]PromptMeta, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *PromptList) SetData(v []PromptMeta)`

SetData sets Data field to given value.

### HasData

`func (o *PromptList) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


