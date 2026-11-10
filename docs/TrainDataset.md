# TrainDataset

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Splits** | Pointer to **map[string]string** | Splits maps a role (train, validation) to the dataset&#39;s own split name. | [optional] 
**Uri** | Pointer to **string** | URI names the data. | [optional] 

## Methods

### NewTrainDataset

`func NewTrainDataset() *TrainDataset`

NewTrainDataset instantiates a new TrainDataset object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrainDatasetWithDefaults

`func NewTrainDatasetWithDefaults() *TrainDataset`

NewTrainDatasetWithDefaults instantiates a new TrainDataset object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSplits

`func (o *TrainDataset) GetSplits() map[string]string`

GetSplits returns the Splits field if non-nil, zero value otherwise.

### GetSplitsOk

`func (o *TrainDataset) GetSplitsOk() (*map[string]string, bool)`

GetSplitsOk returns a tuple with the Splits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSplits

`func (o *TrainDataset) SetSplits(v map[string]string)`

SetSplits sets Splits field to given value.

### HasSplits

`func (o *TrainDataset) HasSplits() bool`

HasSplits returns a boolean if a field has been set.

### GetUri

`func (o *TrainDataset) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *TrainDataset) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *TrainDataset) SetUri(v string)`

SetUri sets Uri field to given value.

### HasUri

`func (o *TrainDataset) HasUri() bool`

HasUri returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


