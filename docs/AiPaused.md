# AiPaused

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Fallback** | Pointer to **string** | Fallback is the Hanzo model that answers a conversation in its place. | [optional] 
**Model** | Pointer to **string** | Model is the model id, or a pattern ending in &#x60;*&#x60; naming a family of them. | [optional] 
**ResetsAt** | Pointer to **string** | ResetsAt is when its share resets (RFC3339). | [optional] 

## Methods

### NewAiPaused

`func NewAiPaused() *AiPaused`

NewAiPaused instantiates a new AiPaused object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiPausedWithDefaults

`func NewAiPausedWithDefaults() *AiPaused`

NewAiPausedWithDefaults instantiates a new AiPaused object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFallback

`func (o *AiPaused) GetFallback() string`

GetFallback returns the Fallback field if non-nil, zero value otherwise.

### GetFallbackOk

`func (o *AiPaused) GetFallbackOk() (*string, bool)`

GetFallbackOk returns a tuple with the Fallback field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFallback

`func (o *AiPaused) SetFallback(v string)`

SetFallback sets Fallback field to given value.

### HasFallback

`func (o *AiPaused) HasFallback() bool`

HasFallback returns a boolean if a field has been set.

### GetModel

`func (o *AiPaused) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *AiPaused) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *AiPaused) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *AiPaused) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetResetsAt

`func (o *AiPaused) GetResetsAt() string`

GetResetsAt returns the ResetsAt field if non-nil, zero value otherwise.

### GetResetsAtOk

`func (o *AiPaused) GetResetsAtOk() (*string, bool)`

GetResetsAtOk returns a tuple with the ResetsAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResetsAt

`func (o *AiPaused) SetResetsAt(v string)`

SetResetsAt sets ResetsAt field to given value.

### HasResetsAt

`func (o *AiPaused) HasResetsAt() bool`

HasResetsAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


