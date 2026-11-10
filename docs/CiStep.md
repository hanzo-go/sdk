# CiStep

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Conclusion** | Pointer to **string** | Conclusion is how the step ended; empty until it has. | [optional] 
**Name** | Pointer to **string** | Name is the step&#39;s name as the workflow declares it. | [optional] 
**Number** | Pointer to **int64** | Number is the step&#39;s position in the job, from 1. | [optional] 
**Status** | Pointer to **string** | Status is whether the step is over. | [optional] 

## Methods

### NewCiStep

`func NewCiStep() *CiStep`

NewCiStep instantiates a new CiStep object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCiStepWithDefaults

`func NewCiStepWithDefaults() *CiStep`

NewCiStepWithDefaults instantiates a new CiStep object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConclusion

`func (o *CiStep) GetConclusion() string`

GetConclusion returns the Conclusion field if non-nil, zero value otherwise.

### GetConclusionOk

`func (o *CiStep) GetConclusionOk() (*string, bool)`

GetConclusionOk returns a tuple with the Conclusion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConclusion

`func (o *CiStep) SetConclusion(v string)`

SetConclusion sets Conclusion field to given value.

### HasConclusion

`func (o *CiStep) HasConclusion() bool`

HasConclusion returns a boolean if a field has been set.

### GetName

`func (o *CiStep) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CiStep) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CiStep) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CiStep) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNumber

`func (o *CiStep) GetNumber() int64`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *CiStep) GetNumberOk() (*int64, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *CiStep) SetNumber(v int64)`

SetNumber sets Number field to given value.

### HasNumber

`func (o *CiStep) HasNumber() bool`

HasNumber returns a boolean if a field has been set.

### GetStatus

`func (o *CiStep) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CiStep) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CiStep) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CiStep) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


