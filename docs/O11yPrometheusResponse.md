# O11yPrometheusResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** |  | [optional] 
**Data** | Pointer to [**O11yPrometheusResponseData**](O11yPrometheusResponseData.md) |  | [optional] 
**ErrorType** | Pointer to **string** |  | [optional] 
**Error** | Pointer to **string** |  | [optional] 

## Methods

### NewO11yPrometheusResponse

`func NewO11yPrometheusResponse() *O11yPrometheusResponse`

NewO11yPrometheusResponse instantiates a new O11yPrometheusResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yPrometheusResponseWithDefaults

`func NewO11yPrometheusResponseWithDefaults() *O11yPrometheusResponse`

NewO11yPrometheusResponseWithDefaults instantiates a new O11yPrometheusResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *O11yPrometheusResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *O11yPrometheusResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *O11yPrometheusResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *O11yPrometheusResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetData

`func (o *O11yPrometheusResponse) GetData() O11yPrometheusResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *O11yPrometheusResponse) GetDataOk() (*O11yPrometheusResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *O11yPrometheusResponse) SetData(v O11yPrometheusResponseData)`

SetData sets Data field to given value.

### HasData

`func (o *O11yPrometheusResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetErrorType

`func (o *O11yPrometheusResponse) GetErrorType() string`

GetErrorType returns the ErrorType field if non-nil, zero value otherwise.

### GetErrorTypeOk

`func (o *O11yPrometheusResponse) GetErrorTypeOk() (*string, bool)`

GetErrorTypeOk returns a tuple with the ErrorType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorType

`func (o *O11yPrometheusResponse) SetErrorType(v string)`

SetErrorType sets ErrorType field to given value.

### HasErrorType

`func (o *O11yPrometheusResponse) HasErrorType() bool`

HasErrorType returns a boolean if a field has been set.

### GetError

`func (o *O11yPrometheusResponse) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *O11yPrometheusResponse) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *O11yPrometheusResponse) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *O11yPrometheusResponse) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


