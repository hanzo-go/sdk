# ComputeJobCanceled

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Canceled** | Pointer to **string** | Canceled is the job id that was canceled. | [optional] 
**Run** | Pointer to **string** | Run is the run id the cancel was applied to. | [optional] 

## Methods

### NewComputeJobCanceled

`func NewComputeJobCanceled() *ComputeJobCanceled`

NewComputeJobCanceled instantiates a new ComputeJobCanceled object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComputeJobCanceledWithDefaults

`func NewComputeJobCanceledWithDefaults() *ComputeJobCanceled`

NewComputeJobCanceledWithDefaults instantiates a new ComputeJobCanceled object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCanceled

`func (o *ComputeJobCanceled) GetCanceled() string`

GetCanceled returns the Canceled field if non-nil, zero value otherwise.

### GetCanceledOk

`func (o *ComputeJobCanceled) GetCanceledOk() (*string, bool)`

GetCanceledOk returns a tuple with the Canceled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanceled

`func (o *ComputeJobCanceled) SetCanceled(v string)`

SetCanceled sets Canceled field to given value.

### HasCanceled

`func (o *ComputeJobCanceled) HasCanceled() bool`

HasCanceled returns a boolean if a field has been set.

### GetRun

`func (o *ComputeJobCanceled) GetRun() string`

GetRun returns the Run field if non-nil, zero value otherwise.

### GetRunOk

`func (o *ComputeJobCanceled) GetRunOk() (*string, bool)`

GetRunOk returns a tuple with the Run field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRun

`func (o *ComputeJobCanceled) SetRun(v string)`

SetRun sets Run field to given value.

### HasRun

`func (o *ComputeJobCanceled) HasRun() bool`

HasRun returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


