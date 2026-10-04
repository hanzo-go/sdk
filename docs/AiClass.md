# AiClass

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Paying** | Pointer to **string** | Paying is plan, prepaid, credits, or none. | [optional] 
**Percent** | Pointer to **int64** |  | [optional] 
**ResetsAt** | Pointer to **string** |  | [optional] 
**State** | Pointer to **string** | State is ok, near, or limited when the class is used up and nothing else pays. | [optional] 
**Window** | Pointer to [**AiWindow**](AiWindow.md) |  | [optional] 

## Methods

### NewAiClass

`func NewAiClass() *AiClass`

NewAiClass instantiates a new AiClass object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiClassWithDefaults

`func NewAiClassWithDefaults() *AiClass`

NewAiClassWithDefaults instantiates a new AiClass object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPaying

`func (o *AiClass) GetPaying() string`

GetPaying returns the Paying field if non-nil, zero value otherwise.

### GetPayingOk

`func (o *AiClass) GetPayingOk() (*string, bool)`

GetPayingOk returns a tuple with the Paying field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaying

`func (o *AiClass) SetPaying(v string)`

SetPaying sets Paying field to given value.

### HasPaying

`func (o *AiClass) HasPaying() bool`

HasPaying returns a boolean if a field has been set.

### GetPercent

`func (o *AiClass) GetPercent() int64`

GetPercent returns the Percent field if non-nil, zero value otherwise.

### GetPercentOk

`func (o *AiClass) GetPercentOk() (*int64, bool)`

GetPercentOk returns a tuple with the Percent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercent

`func (o *AiClass) SetPercent(v int64)`

SetPercent sets Percent field to given value.

### HasPercent

`func (o *AiClass) HasPercent() bool`

HasPercent returns a boolean if a field has been set.

### GetResetsAt

`func (o *AiClass) GetResetsAt() string`

GetResetsAt returns the ResetsAt field if non-nil, zero value otherwise.

### GetResetsAtOk

`func (o *AiClass) GetResetsAtOk() (*string, bool)`

GetResetsAtOk returns a tuple with the ResetsAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResetsAt

`func (o *AiClass) SetResetsAt(v string)`

SetResetsAt sets ResetsAt field to given value.

### HasResetsAt

`func (o *AiClass) HasResetsAt() bool`

HasResetsAt returns a boolean if a field has been set.

### GetState

`func (o *AiClass) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *AiClass) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *AiClass) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *AiClass) HasState() bool`

HasState returns a boolean if a field has been set.

### GetWindow

`func (o *AiClass) GetWindow() AiWindow`

GetWindow returns the Window field if non-nil, zero value otherwise.

### GetWindowOk

`func (o *AiClass) GetWindowOk() (*AiWindow, bool)`

GetWindowOk returns a tuple with the Window field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindow

`func (o *AiClass) SetWindow(v AiWindow)`

SetWindow sets Window field to given value.

### HasWindow

`func (o *AiClass) HasWindow() bool`

HasWindow returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


