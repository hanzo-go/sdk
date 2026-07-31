# CloudExperimentsOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]CloudExperiment**](CloudExperiment.md) | Data are the canonical experiment versions. | [optional] 
**Total** | Pointer to **int32** | Total is len(data) — the rows in this answer, not the store&#39;s history. | [optional] 

## Methods

### NewCloudExperimentsOut

`func NewCloudExperimentsOut() *CloudExperimentsOut`

NewCloudExperimentsOut instantiates a new CloudExperimentsOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudExperimentsOutWithDefaults

`func NewCloudExperimentsOutWithDefaults() *CloudExperimentsOut`

NewCloudExperimentsOutWithDefaults instantiates a new CloudExperimentsOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *CloudExperimentsOut) GetData() []CloudExperiment`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CloudExperimentsOut) GetDataOk() (*[]CloudExperiment, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CloudExperimentsOut) SetData(v []CloudExperiment)`

SetData sets Data field to given value.

### HasData

`func (o *CloudExperimentsOut) HasData() bool`

HasData returns a boolean if a field has been set.

### GetTotal

`func (o *CloudExperimentsOut) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *CloudExperimentsOut) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *CloudExperimentsOut) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *CloudExperimentsOut) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


