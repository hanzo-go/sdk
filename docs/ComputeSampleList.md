# ComputeSampleList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Samples** | Pointer to [**[]ComputeSampleView**](ComputeSampleView.md) | Samples are the readings, OLDEST first — the order a chart plots. | [optional] 

## Methods

### NewComputeSampleList

`func NewComputeSampleList() *ComputeSampleList`

NewComputeSampleList instantiates a new ComputeSampleList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComputeSampleListWithDefaults

`func NewComputeSampleListWithDefaults() *ComputeSampleList`

NewComputeSampleListWithDefaults instantiates a new ComputeSampleList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSamples

`func (o *ComputeSampleList) GetSamples() []ComputeSampleView`

GetSamples returns the Samples field if non-nil, zero value otherwise.

### GetSamplesOk

`func (o *ComputeSampleList) GetSamplesOk() (*[]ComputeSampleView, bool)`

GetSamplesOk returns a tuple with the Samples field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamples

`func (o *ComputeSampleList) SetSamples(v []ComputeSampleView)`

SetSamples sets Samples field to given value.

### HasSamples

`func (o *ComputeSampleList) HasSamples() bool`

HasSamples returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


