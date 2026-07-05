# AdminO11ySvcStat

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Service** | Pointer to **string** |  | [optional] 
**Requests** | Pointer to **int64** |  | [optional] 
**ErrorRate** | Pointer to **float64** |  | [optional] 
**LatencyP95Ms** | Pointer to **float64** |  | [optional] 

## Methods

### NewAdminO11ySvcStat

`func NewAdminO11ySvcStat() *AdminO11ySvcStat`

NewAdminO11ySvcStat instantiates a new AdminO11ySvcStat object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdminO11ySvcStatWithDefaults

`func NewAdminO11ySvcStatWithDefaults() *AdminO11ySvcStat`

NewAdminO11ySvcStatWithDefaults instantiates a new AdminO11ySvcStat object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetService

`func (o *AdminO11ySvcStat) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *AdminO11ySvcStat) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *AdminO11ySvcStat) SetService(v string)`

SetService sets Service field to given value.

### HasService

`func (o *AdminO11ySvcStat) HasService() bool`

HasService returns a boolean if a field has been set.

### GetRequests

`func (o *AdminO11ySvcStat) GetRequests() int64`

GetRequests returns the Requests field if non-nil, zero value otherwise.

### GetRequestsOk

`func (o *AdminO11ySvcStat) GetRequestsOk() (*int64, bool)`

GetRequestsOk returns a tuple with the Requests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequests

`func (o *AdminO11ySvcStat) SetRequests(v int64)`

SetRequests sets Requests field to given value.

### HasRequests

`func (o *AdminO11ySvcStat) HasRequests() bool`

HasRequests returns a boolean if a field has been set.

### GetErrorRate

`func (o *AdminO11ySvcStat) GetErrorRate() float64`

GetErrorRate returns the ErrorRate field if non-nil, zero value otherwise.

### GetErrorRateOk

`func (o *AdminO11ySvcStat) GetErrorRateOk() (*float64, bool)`

GetErrorRateOk returns a tuple with the ErrorRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorRate

`func (o *AdminO11ySvcStat) SetErrorRate(v float64)`

SetErrorRate sets ErrorRate field to given value.

### HasErrorRate

`func (o *AdminO11ySvcStat) HasErrorRate() bool`

HasErrorRate returns a boolean if a field has been set.

### GetLatencyP95Ms

`func (o *AdminO11ySvcStat) GetLatencyP95Ms() float64`

GetLatencyP95Ms returns the LatencyP95Ms field if non-nil, zero value otherwise.

### GetLatencyP95MsOk

`func (o *AdminO11ySvcStat) GetLatencyP95MsOk() (*float64, bool)`

GetLatencyP95MsOk returns a tuple with the LatencyP95Ms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatencyP95Ms

`func (o *AdminO11ySvcStat) SetLatencyP95Ms(v float64)`

SetLatencyP95Ms sets LatencyP95Ms field to given value.

### HasLatencyP95Ms

`func (o *AdminO11ySvcStat) HasLatencyP95Ms() bool`

HasLatencyP95Ms returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


