# O11yMetricQueryResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** |  | [optional] 
**Data** | Pointer to [**O11yMetricQueryResultData**](O11yMetricQueryResultData.md) |  | [optional] 
**ErrorType** | Pointer to **string** |  | [optional] 
**Error** | Pointer to **string** |  | [optional] 

## Methods

### NewO11yMetricQueryResult

`func NewO11yMetricQueryResult() *O11yMetricQueryResult`

NewO11yMetricQueryResult instantiates a new O11yMetricQueryResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yMetricQueryResultWithDefaults

`func NewO11yMetricQueryResultWithDefaults() *O11yMetricQueryResult`

NewO11yMetricQueryResultWithDefaults instantiates a new O11yMetricQueryResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *O11yMetricQueryResult) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *O11yMetricQueryResult) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *O11yMetricQueryResult) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *O11yMetricQueryResult) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetData

`func (o *O11yMetricQueryResult) GetData() O11yMetricQueryResultData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *O11yMetricQueryResult) GetDataOk() (*O11yMetricQueryResultData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *O11yMetricQueryResult) SetData(v O11yMetricQueryResultData)`

SetData sets Data field to given value.

### HasData

`func (o *O11yMetricQueryResult) HasData() bool`

HasData returns a boolean if a field has been set.

### GetErrorType

`func (o *O11yMetricQueryResult) GetErrorType() string`

GetErrorType returns the ErrorType field if non-nil, zero value otherwise.

### GetErrorTypeOk

`func (o *O11yMetricQueryResult) GetErrorTypeOk() (*string, bool)`

GetErrorTypeOk returns a tuple with the ErrorType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorType

`func (o *O11yMetricQueryResult) SetErrorType(v string)`

SetErrorType sets ErrorType field to given value.

### HasErrorType

`func (o *O11yMetricQueryResult) HasErrorType() bool`

HasErrorType returns a boolean if a field has been set.

### GetError

`func (o *O11yMetricQueryResult) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *O11yMetricQueryResult) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *O11yMetricQueryResult) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *O11yMetricQueryResult) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


