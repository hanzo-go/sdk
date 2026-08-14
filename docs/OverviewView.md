# OverviewView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Custom** | Pointer to **bool** |  | [optional] 
**Funnel** | Pointer to [**Funnel**](Funnel.md) |  | [optional] 
**Progress** | Pointer to [**ProgressView**](ProgressView.md) |  | [optional] 
**Steps** | Pointer to [**[]StepView**](StepView.md) |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**Version** | Pointer to **string** |  | [optional] 

## Methods

### NewOverviewView

`func NewOverviewView() *OverviewView`

NewOverviewView instantiates a new OverviewView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOverviewViewWithDefaults

`func NewOverviewViewWithDefaults() *OverviewView`

NewOverviewViewWithDefaults instantiates a new OverviewView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustom

`func (o *OverviewView) GetCustom() bool`

GetCustom returns the Custom field if non-nil, zero value otherwise.

### GetCustomOk

`func (o *OverviewView) GetCustomOk() (*bool, bool)`

GetCustomOk returns a tuple with the Custom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustom

`func (o *OverviewView) SetCustom(v bool)`

SetCustom sets Custom field to given value.

### HasCustom

`func (o *OverviewView) HasCustom() bool`

HasCustom returns a boolean if a field has been set.

### GetFunnel

`func (o *OverviewView) GetFunnel() Funnel`

GetFunnel returns the Funnel field if non-nil, zero value otherwise.

### GetFunnelOk

`func (o *OverviewView) GetFunnelOk() (*Funnel, bool)`

GetFunnelOk returns a tuple with the Funnel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunnel

`func (o *OverviewView) SetFunnel(v Funnel)`

SetFunnel sets Funnel field to given value.

### HasFunnel

`func (o *OverviewView) HasFunnel() bool`

HasFunnel returns a boolean if a field has been set.

### GetProgress

`func (o *OverviewView) GetProgress() ProgressView`

GetProgress returns the Progress field if non-nil, zero value otherwise.

### GetProgressOk

`func (o *OverviewView) GetProgressOk() (*ProgressView, bool)`

GetProgressOk returns a tuple with the Progress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgress

`func (o *OverviewView) SetProgress(v ProgressView)`

SetProgress sets Progress field to given value.

### HasProgress

`func (o *OverviewView) HasProgress() bool`

HasProgress returns a boolean if a field has been set.

### GetSteps

`func (o *OverviewView) GetSteps() []StepView`

GetSteps returns the Steps field if non-nil, zero value otherwise.

### GetStepsOk

`func (o *OverviewView) GetStepsOk() (*[]StepView, bool)`

GetStepsOk returns a tuple with the Steps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSteps

`func (o *OverviewView) SetSteps(v []StepView)`

SetSteps sets Steps field to given value.

### HasSteps

`func (o *OverviewView) HasSteps() bool`

HasSteps returns a boolean if a field has been set.

### GetTitle

`func (o *OverviewView) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *OverviewView) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *OverviewView) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *OverviewView) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetVersion

`func (o *OverviewView) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *OverviewView) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *OverviewView) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *OverviewView) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


