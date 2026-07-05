# O11yLogQueryResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** |  | [optional] 
**Data** | Pointer to [**O11yLogQueryResultData**](O11yLogQueryResultData.md) |  | [optional] 

## Methods

### NewO11yLogQueryResult

`func NewO11yLogQueryResult() *O11yLogQueryResult`

NewO11yLogQueryResult instantiates a new O11yLogQueryResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yLogQueryResultWithDefaults

`func NewO11yLogQueryResultWithDefaults() *O11yLogQueryResult`

NewO11yLogQueryResultWithDefaults instantiates a new O11yLogQueryResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *O11yLogQueryResult) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *O11yLogQueryResult) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *O11yLogQueryResult) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *O11yLogQueryResult) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetData

`func (o *O11yLogQueryResult) GetData() O11yLogQueryResultData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *O11yLogQueryResult) GetDataOk() (*O11yLogQueryResultData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *O11yLogQueryResult) SetData(v O11yLogQueryResultData)`

SetData sets Data field to given value.

### HasData

`func (o *O11yLogQueryResult) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


