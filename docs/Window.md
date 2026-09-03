# Window

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Limit** | Pointer to **int64** |  | [optional] 
**Remaining** | Pointer to **int64** |  | [optional] 
**Resets** | Pointer to **string** |  | [optional] 
**Span** | Pointer to **string** |  | [optional] 
**Used** | Pointer to **int64** |  | [optional] 

## Methods

### NewWindow

`func NewWindow() *Window`

NewWindow instantiates a new Window object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWindowWithDefaults

`func NewWindowWithDefaults() *Window`

NewWindowWithDefaults instantiates a new Window object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLimit

`func (o *Window) GetLimit() int64`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *Window) GetLimitOk() (*int64, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *Window) SetLimit(v int64)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *Window) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetRemaining

`func (o *Window) GetRemaining() int64`

GetRemaining returns the Remaining field if non-nil, zero value otherwise.

### GetRemainingOk

`func (o *Window) GetRemainingOk() (*int64, bool)`

GetRemainingOk returns a tuple with the Remaining field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemaining

`func (o *Window) SetRemaining(v int64)`

SetRemaining sets Remaining field to given value.

### HasRemaining

`func (o *Window) HasRemaining() bool`

HasRemaining returns a boolean if a field has been set.

### GetResets

`func (o *Window) GetResets() string`

GetResets returns the Resets field if non-nil, zero value otherwise.

### GetResetsOk

`func (o *Window) GetResetsOk() (*string, bool)`

GetResetsOk returns a tuple with the Resets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResets

`func (o *Window) SetResets(v string)`

SetResets sets Resets field to given value.

### HasResets

`func (o *Window) HasResets() bool`

HasResets returns a boolean if a field has been set.

### GetSpan

`func (o *Window) GetSpan() string`

GetSpan returns the Span field if non-nil, zero value otherwise.

### GetSpanOk

`func (o *Window) GetSpanOk() (*string, bool)`

GetSpanOk returns a tuple with the Span field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpan

`func (o *Window) SetSpan(v string)`

SetSpan sets Span field to given value.

### HasSpan

`func (o *Window) HasSpan() bool`

HasSpan returns a boolean if a field has been set.

### GetUsed

`func (o *Window) GetUsed() int64`

GetUsed returns the Used field if non-nil, zero value otherwise.

### GetUsedOk

`func (o *Window) GetUsedOk() (*int64, bool)`

GetUsedOk returns a tuple with the Used field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsed

`func (o *Window) SetUsed(v int64)`

SetUsed sets Used field to given value.

### HasUsed

`func (o *Window) HasUsed() bool`

HasUsed returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


