# O11ySvcStat

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ErrorRate** | Pointer to **float32** | percent (0..100) | [optional] 
**LatencyP95Ms** | Pointer to **float32** |  | [optional] 
**Requests** | Pointer to **int32** |  | [optional] 
**Service** | Pointer to **string** |  | [optional] 

## Methods

### NewO11ySvcStat

`func NewO11ySvcStat() *O11ySvcStat`

NewO11ySvcStat instantiates a new O11ySvcStat object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11ySvcStatWithDefaults

`func NewO11ySvcStatWithDefaults() *O11ySvcStat`

NewO11ySvcStatWithDefaults instantiates a new O11ySvcStat object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrorRate

`func (o *O11ySvcStat) GetErrorRate() float32`

GetErrorRate returns the ErrorRate field if non-nil, zero value otherwise.

### GetErrorRateOk

`func (o *O11ySvcStat) GetErrorRateOk() (*float32, bool)`

GetErrorRateOk returns a tuple with the ErrorRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorRate

`func (o *O11ySvcStat) SetErrorRate(v float32)`

SetErrorRate sets ErrorRate field to given value.

### HasErrorRate

`func (o *O11ySvcStat) HasErrorRate() bool`

HasErrorRate returns a boolean if a field has been set.

### GetLatencyP95Ms

`func (o *O11ySvcStat) GetLatencyP95Ms() float32`

GetLatencyP95Ms returns the LatencyP95Ms field if non-nil, zero value otherwise.

### GetLatencyP95MsOk

`func (o *O11ySvcStat) GetLatencyP95MsOk() (*float32, bool)`

GetLatencyP95MsOk returns a tuple with the LatencyP95Ms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatencyP95Ms

`func (o *O11ySvcStat) SetLatencyP95Ms(v float32)`

SetLatencyP95Ms sets LatencyP95Ms field to given value.

### HasLatencyP95Ms

`func (o *O11ySvcStat) HasLatencyP95Ms() bool`

HasLatencyP95Ms returns a boolean if a field has been set.

### GetRequests

`func (o *O11ySvcStat) GetRequests() int32`

GetRequests returns the Requests field if non-nil, zero value otherwise.

### GetRequestsOk

`func (o *O11ySvcStat) GetRequestsOk() (*int32, bool)`

GetRequestsOk returns a tuple with the Requests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequests

`func (o *O11ySvcStat) SetRequests(v int32)`

SetRequests sets Requests field to given value.

### HasRequests

`func (o *O11ySvcStat) HasRequests() bool`

HasRequests returns a boolean if a field has been set.

### GetService

`func (o *O11ySvcStat) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *O11ySvcStat) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *O11ySvcStat) SetService(v string)`

SetService sets Service field to given value.

### HasService

`func (o *O11ySvcStat) HasService() bool`

HasService returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


