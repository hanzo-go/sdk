# O11ySeriesResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** |  | [optional] 
**Data** | Pointer to **[]map[string]string** |  | [optional] 

## Methods

### NewO11ySeriesResult

`func NewO11ySeriesResult() *O11ySeriesResult`

NewO11ySeriesResult instantiates a new O11ySeriesResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11ySeriesResultWithDefaults

`func NewO11ySeriesResultWithDefaults() *O11ySeriesResult`

NewO11ySeriesResultWithDefaults instantiates a new O11ySeriesResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *O11ySeriesResult) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *O11ySeriesResult) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *O11ySeriesResult) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *O11ySeriesResult) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetData

`func (o *O11ySeriesResult) GetData() []map[string]string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *O11ySeriesResult) GetDataOk() (*[]map[string]string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *O11ySeriesResult) SetData(v []map[string]string)`

SetData sets Data field to given value.

### HasData

`func (o *O11ySeriesResult) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


