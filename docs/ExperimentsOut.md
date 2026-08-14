# ExperimentsOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]Experiment**](Experiment.md) | Data are the canonical experiment versions. | [optional] 
**Total** | Pointer to **int32** | Total is len(data) — the rows in this answer, not the store&#39;s history. | [optional] 

## Methods

### NewExperimentsOut

`func NewExperimentsOut() *ExperimentsOut`

NewExperimentsOut instantiates a new ExperimentsOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExperimentsOutWithDefaults

`func NewExperimentsOutWithDefaults() *ExperimentsOut`

NewExperimentsOutWithDefaults instantiates a new ExperimentsOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ExperimentsOut) GetData() []Experiment`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ExperimentsOut) GetDataOk() (*[]Experiment, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ExperimentsOut) SetData(v []Experiment)`

SetData sets Data field to given value.

### HasData

`func (o *ExperimentsOut) HasData() bool`

HasData returns a boolean if a field has been set.

### GetTotal

`func (o *ExperimentsOut) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *ExperimentsOut) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *ExperimentsOut) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *ExperimentsOut) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


