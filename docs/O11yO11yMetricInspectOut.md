# O11yO11yMetricInspectOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**O11yO11yMetricSeriesSet**](O11yO11yMetricSeriesSet.md) | Data holds the series. | [optional] 
**Status** | Pointer to **string** | Status is \&quot;success\&quot;. | [optional] 

## Methods

### NewO11yO11yMetricInspectOut

`func NewO11yO11yMetricInspectOut() *O11yO11yMetricInspectOut`

NewO11yO11yMetricInspectOut instantiates a new O11yO11yMetricInspectOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yO11yMetricInspectOutWithDefaults

`func NewO11yO11yMetricInspectOutWithDefaults() *O11yO11yMetricInspectOut`

NewO11yO11yMetricInspectOutWithDefaults instantiates a new O11yO11yMetricInspectOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *O11yO11yMetricInspectOut) GetData() O11yO11yMetricSeriesSet`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *O11yO11yMetricInspectOut) GetDataOk() (*O11yO11yMetricSeriesSet, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *O11yO11yMetricInspectOut) SetData(v O11yO11yMetricSeriesSet)`

SetData sets Data field to given value.

### HasData

`func (o *O11yO11yMetricInspectOut) HasData() bool`

HasData returns a boolean if a field has been set.

### GetStatus

`func (o *O11yO11yMetricInspectOut) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *O11yO11yMetricInspectOut) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *O11yO11yMetricInspectOut) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *O11yO11yMetricInspectOut) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


