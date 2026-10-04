# AiWindow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Percent** | Pointer to **int64** | Percent is the share used, 0 to 100, rounded up to the next five so any use shows. | [optional] 
**ResetsAt** | Pointer to **string** | ResetsAt is when the window starts again (RFC3339), null for a session that is not running: it starts at the next request. | [optional] 
**State** | Pointer to **string** | State is ok, near (four fifths used) or limited (used up). | [optional] 

## Methods

### NewAiWindow

`func NewAiWindow() *AiWindow`

NewAiWindow instantiates a new AiWindow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiWindowWithDefaults

`func NewAiWindowWithDefaults() *AiWindow`

NewAiWindowWithDefaults instantiates a new AiWindow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPercent

`func (o *AiWindow) GetPercent() int64`

GetPercent returns the Percent field if non-nil, zero value otherwise.

### GetPercentOk

`func (o *AiWindow) GetPercentOk() (*int64, bool)`

GetPercentOk returns a tuple with the Percent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercent

`func (o *AiWindow) SetPercent(v int64)`

SetPercent sets Percent field to given value.

### HasPercent

`func (o *AiWindow) HasPercent() bool`

HasPercent returns a boolean if a field has been set.

### GetResetsAt

`func (o *AiWindow) GetResetsAt() string`

GetResetsAt returns the ResetsAt field if non-nil, zero value otherwise.

### GetResetsAtOk

`func (o *AiWindow) GetResetsAtOk() (*string, bool)`

GetResetsAtOk returns a tuple with the ResetsAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResetsAt

`func (o *AiWindow) SetResetsAt(v string)`

SetResetsAt sets ResetsAt field to given value.

### HasResetsAt

`func (o *AiWindow) HasResetsAt() bool`

HasResetsAt returns a boolean if a field has been set.

### GetState

`func (o *AiWindow) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *AiWindow) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *AiWindow) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *AiWindow) HasState() bool`

HasState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


