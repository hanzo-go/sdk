# CloudSampleList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Samples** | Pointer to [**[]CloudSampleView**](CloudSampleView.md) | Samples are the readings, OLDEST first — the order a chart plots. | [optional] 

## Methods

### NewCloudSampleList

`func NewCloudSampleList() *CloudSampleList`

NewCloudSampleList instantiates a new CloudSampleList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudSampleListWithDefaults

`func NewCloudSampleListWithDefaults() *CloudSampleList`

NewCloudSampleListWithDefaults instantiates a new CloudSampleList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSamples

`func (o *CloudSampleList) GetSamples() []CloudSampleView`

GetSamples returns the Samples field if non-nil, zero value otherwise.

### GetSamplesOk

`func (o *CloudSampleList) GetSamplesOk() (*[]CloudSampleView, bool)`

GetSamplesOk returns a tuple with the Samples field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamples

`func (o *CloudSampleList) SetSamples(v []CloudSampleView)`

SetSamples sets Samples field to given value.

### HasSamples

`func (o *CloudSampleList) HasSamples() bool`

HasSamples returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


