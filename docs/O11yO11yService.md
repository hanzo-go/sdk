# O11yO11yService

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AvgDuration** | Pointer to **float64** | AvgDuration is their average latency, nanoseconds. | [optional] 
**CallRate** | Pointer to **float64** | CallRate is calls per second over the window. | [optional] 
**DataWarning** | Pointer to [**O11yO11yServiceWarning**](O11yO11yServiceWarning.md) | DataWarning carries the entry-point operations the numbers were computed over. | [optional] 
**ErrorRate** | Pointer to **float64** | ErrorRate is the percentage of calls that errored. | [optional] 
**FourXXRate** | Pointer to **float64** | FourXXRate is the percentage of calls that answered 4xx. | [optional] 
**Num4XX** | Pointer to **int32** | Num4XX is how many of the calls answered 4xx. | [optional] 
**NumCalls** | Pointer to **int32** | NumCalls is how many entry-point spans landed in the window. | [optional] 
**NumErrors** | Pointer to **int32** | NumErrors is how many of the calls errored. | [optional] 
**P99** | Pointer to **float64** | Percentile99 is the p99 latency of its entry-point spans, nanoseconds. | [optional] 
**ServiceName** | Pointer to **string** | ServiceName is the service. | [optional] 

## Methods

### NewO11yO11yService

`func NewO11yO11yService() *O11yO11yService`

NewO11yO11yService instantiates a new O11yO11yService object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yO11yServiceWithDefaults

`func NewO11yO11yServiceWithDefaults() *O11yO11yService`

NewO11yO11yServiceWithDefaults instantiates a new O11yO11yService object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvgDuration

`func (o *O11yO11yService) GetAvgDuration() float64`

GetAvgDuration returns the AvgDuration field if non-nil, zero value otherwise.

### GetAvgDurationOk

`func (o *O11yO11yService) GetAvgDurationOk() (*float64, bool)`

GetAvgDurationOk returns a tuple with the AvgDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvgDuration

`func (o *O11yO11yService) SetAvgDuration(v float64)`

SetAvgDuration sets AvgDuration field to given value.

### HasAvgDuration

`func (o *O11yO11yService) HasAvgDuration() bool`

HasAvgDuration returns a boolean if a field has been set.

### GetCallRate

`func (o *O11yO11yService) GetCallRate() float64`

GetCallRate returns the CallRate field if non-nil, zero value otherwise.

### GetCallRateOk

`func (o *O11yO11yService) GetCallRateOk() (*float64, bool)`

GetCallRateOk returns a tuple with the CallRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallRate

`func (o *O11yO11yService) SetCallRate(v float64)`

SetCallRate sets CallRate field to given value.

### HasCallRate

`func (o *O11yO11yService) HasCallRate() bool`

HasCallRate returns a boolean if a field has been set.

### GetDataWarning

`func (o *O11yO11yService) GetDataWarning() O11yO11yServiceWarning`

GetDataWarning returns the DataWarning field if non-nil, zero value otherwise.

### GetDataWarningOk

`func (o *O11yO11yService) GetDataWarningOk() (*O11yO11yServiceWarning, bool)`

GetDataWarningOk returns a tuple with the DataWarning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataWarning

`func (o *O11yO11yService) SetDataWarning(v O11yO11yServiceWarning)`

SetDataWarning sets DataWarning field to given value.

### HasDataWarning

`func (o *O11yO11yService) HasDataWarning() bool`

HasDataWarning returns a boolean if a field has been set.

### GetErrorRate

`func (o *O11yO11yService) GetErrorRate() float64`

GetErrorRate returns the ErrorRate field if non-nil, zero value otherwise.

### GetErrorRateOk

`func (o *O11yO11yService) GetErrorRateOk() (*float64, bool)`

GetErrorRateOk returns a tuple with the ErrorRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorRate

`func (o *O11yO11yService) SetErrorRate(v float64)`

SetErrorRate sets ErrorRate field to given value.

### HasErrorRate

`func (o *O11yO11yService) HasErrorRate() bool`

HasErrorRate returns a boolean if a field has been set.

### GetFourXXRate

`func (o *O11yO11yService) GetFourXXRate() float64`

GetFourXXRate returns the FourXXRate field if non-nil, zero value otherwise.

### GetFourXXRateOk

`func (o *O11yO11yService) GetFourXXRateOk() (*float64, bool)`

GetFourXXRateOk returns a tuple with the FourXXRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFourXXRate

`func (o *O11yO11yService) SetFourXXRate(v float64)`

SetFourXXRate sets FourXXRate field to given value.

### HasFourXXRate

`func (o *O11yO11yService) HasFourXXRate() bool`

HasFourXXRate returns a boolean if a field has been set.

### GetNum4XX

`func (o *O11yO11yService) GetNum4XX() int32`

GetNum4XX returns the Num4XX field if non-nil, zero value otherwise.

### GetNum4XXOk

`func (o *O11yO11yService) GetNum4XXOk() (*int32, bool)`

GetNum4XXOk returns a tuple with the Num4XX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNum4XX

`func (o *O11yO11yService) SetNum4XX(v int32)`

SetNum4XX sets Num4XX field to given value.

### HasNum4XX

`func (o *O11yO11yService) HasNum4XX() bool`

HasNum4XX returns a boolean if a field has been set.

### GetNumCalls

`func (o *O11yO11yService) GetNumCalls() int32`

GetNumCalls returns the NumCalls field if non-nil, zero value otherwise.

### GetNumCallsOk

`func (o *O11yO11yService) GetNumCallsOk() (*int32, bool)`

GetNumCallsOk returns a tuple with the NumCalls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumCalls

`func (o *O11yO11yService) SetNumCalls(v int32)`

SetNumCalls sets NumCalls field to given value.

### HasNumCalls

`func (o *O11yO11yService) HasNumCalls() bool`

HasNumCalls returns a boolean if a field has been set.

### GetNumErrors

`func (o *O11yO11yService) GetNumErrors() int32`

GetNumErrors returns the NumErrors field if non-nil, zero value otherwise.

### GetNumErrorsOk

`func (o *O11yO11yService) GetNumErrorsOk() (*int32, bool)`

GetNumErrorsOk returns a tuple with the NumErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumErrors

`func (o *O11yO11yService) SetNumErrors(v int32)`

SetNumErrors sets NumErrors field to given value.

### HasNumErrors

`func (o *O11yO11yService) HasNumErrors() bool`

HasNumErrors returns a boolean if a field has been set.

### GetP99

`func (o *O11yO11yService) GetP99() float64`

GetP99 returns the P99 field if non-nil, zero value otherwise.

### GetP99Ok

`func (o *O11yO11yService) GetP99Ok() (*float64, bool)`

GetP99Ok returns a tuple with the P99 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetP99

`func (o *O11yO11yService) SetP99(v float64)`

SetP99 sets P99 field to given value.

### HasP99

`func (o *O11yO11yService) HasP99() bool`

HasP99 returns a boolean if a field has been set.

### GetServiceName

`func (o *O11yO11yService) GetServiceName() string`

GetServiceName returns the ServiceName field if non-nil, zero value otherwise.

### GetServiceNameOk

`func (o *O11yO11yService) GetServiceNameOk() (*string, bool)`

GetServiceNameOk returns a tuple with the ServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceName

`func (o *O11yO11yService) SetServiceName(v string)`

SetServiceName sets ServiceName field to given value.

### HasServiceName

`func (o *O11yO11yService) HasServiceName() bool`

HasServiceName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


