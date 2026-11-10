# FormFieldSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Answered** | Pointer to **int64** | Answered is how many responses answer it. | [optional] 
**Distribution** | Pointer to [**[]FormTally**](FormTally.md) | Distribution counts each value: every choice of a multiple_choice or dropdown (with its label), every point of a rating or opinion_scale, and yes and no. | [optional] 
**DropOff** | Pointer to **int64** | DropOff is how many open responses stopped in front of it: this is the first field they left unanswered after the last one they answered. | [optional] 
**Field** | Pointer to **string** | Field is the field&#39;s ref; Type its type; Title its question. | [optional] 
**Max** | Pointer to **float64** |  | [optional] 
**Mean** | Pointer to **float64** | Mean, Min and Max are a number, rating or opinion_scale&#39;s. | [optional] 
**Min** | Pointer to **float64** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 

## Methods

### NewFormFieldSummary

`func NewFormFieldSummary() *FormFieldSummary`

NewFormFieldSummary instantiates a new FormFieldSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFormFieldSummaryWithDefaults

`func NewFormFieldSummaryWithDefaults() *FormFieldSummary`

NewFormFieldSummaryWithDefaults instantiates a new FormFieldSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnswered

`func (o *FormFieldSummary) GetAnswered() int64`

GetAnswered returns the Answered field if non-nil, zero value otherwise.

### GetAnsweredOk

`func (o *FormFieldSummary) GetAnsweredOk() (*int64, bool)`

GetAnsweredOk returns a tuple with the Answered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnswered

`func (o *FormFieldSummary) SetAnswered(v int64)`

SetAnswered sets Answered field to given value.

### HasAnswered

`func (o *FormFieldSummary) HasAnswered() bool`

HasAnswered returns a boolean if a field has been set.

### GetDistribution

`func (o *FormFieldSummary) GetDistribution() []FormTally`

GetDistribution returns the Distribution field if non-nil, zero value otherwise.

### GetDistributionOk

`func (o *FormFieldSummary) GetDistributionOk() (*[]FormTally, bool)`

GetDistributionOk returns a tuple with the Distribution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistribution

`func (o *FormFieldSummary) SetDistribution(v []FormTally)`

SetDistribution sets Distribution field to given value.

### HasDistribution

`func (o *FormFieldSummary) HasDistribution() bool`

HasDistribution returns a boolean if a field has been set.

### GetDropOff

`func (o *FormFieldSummary) GetDropOff() int64`

GetDropOff returns the DropOff field if non-nil, zero value otherwise.

### GetDropOffOk

`func (o *FormFieldSummary) GetDropOffOk() (*int64, bool)`

GetDropOffOk returns a tuple with the DropOff field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDropOff

`func (o *FormFieldSummary) SetDropOff(v int64)`

SetDropOff sets DropOff field to given value.

### HasDropOff

`func (o *FormFieldSummary) HasDropOff() bool`

HasDropOff returns a boolean if a field has been set.

### GetField

`func (o *FormFieldSummary) GetField() string`

GetField returns the Field field if non-nil, zero value otherwise.

### GetFieldOk

`func (o *FormFieldSummary) GetFieldOk() (*string, bool)`

GetFieldOk returns a tuple with the Field field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField

`func (o *FormFieldSummary) SetField(v string)`

SetField sets Field field to given value.

### HasField

`func (o *FormFieldSummary) HasField() bool`

HasField returns a boolean if a field has been set.

### GetMax

`func (o *FormFieldSummary) GetMax() float64`

GetMax returns the Max field if non-nil, zero value otherwise.

### GetMaxOk

`func (o *FormFieldSummary) GetMaxOk() (*float64, bool)`

GetMaxOk returns a tuple with the Max field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMax

`func (o *FormFieldSummary) SetMax(v float64)`

SetMax sets Max field to given value.

### HasMax

`func (o *FormFieldSummary) HasMax() bool`

HasMax returns a boolean if a field has been set.

### GetMean

`func (o *FormFieldSummary) GetMean() float64`

GetMean returns the Mean field if non-nil, zero value otherwise.

### GetMeanOk

`func (o *FormFieldSummary) GetMeanOk() (*float64, bool)`

GetMeanOk returns a tuple with the Mean field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMean

`func (o *FormFieldSummary) SetMean(v float64)`

SetMean sets Mean field to given value.

### HasMean

`func (o *FormFieldSummary) HasMean() bool`

HasMean returns a boolean if a field has been set.

### GetMin

`func (o *FormFieldSummary) GetMin() float64`

GetMin returns the Min field if non-nil, zero value otherwise.

### GetMinOk

`func (o *FormFieldSummary) GetMinOk() (*float64, bool)`

GetMinOk returns a tuple with the Min field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMin

`func (o *FormFieldSummary) SetMin(v float64)`

SetMin sets Min field to given value.

### HasMin

`func (o *FormFieldSummary) HasMin() bool`

HasMin returns a boolean if a field has been set.

### GetTitle

`func (o *FormFieldSummary) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *FormFieldSummary) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *FormFieldSummary) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *FormFieldSummary) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetType

`func (o *FormFieldSummary) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FormFieldSummary) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FormFieldSummary) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *FormFieldSummary) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


