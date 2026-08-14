# SubsystemRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** |  | [optional] 
**ErrorRate** | Pointer to **float32** | percent (0..100) | [optional] 
**Errors** | Pointer to **int32** |  | [optional] 
**LastErrorAt** | Pointer to **string** |  | [optional] 
**LastErrorMessage** | Pointer to **string** |  | [optional] 
**LastErrorRoute** | Pointer to **string** |  | [optional] 
**LastErrorStatus** | Pointer to **string** |  | [optional] 
**LatencyP50Ms** | Pointer to **float32** |  | [optional] 
**LatencyP95Ms** | Pointer to **float32** |  | [optional] 
**LatencyP99Ms** | Pointer to **float32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Prefixes** | Pointer to **[]string** |  | [optional] 
**Requests** | Pointer to **int32** |  | [optional] 
**RequestsPerMin** | Pointer to **float32** |  | [optional] 

## Methods

### NewSubsystemRow

`func NewSubsystemRow() *SubsystemRow`

NewSubsystemRow instantiates a new SubsystemRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubsystemRowWithDefaults

`func NewSubsystemRowWithDefaults() *SubsystemRow`

NewSubsystemRowWithDefaults instantiates a new SubsystemRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *SubsystemRow) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *SubsystemRow) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *SubsystemRow) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *SubsystemRow) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetErrorRate

`func (o *SubsystemRow) GetErrorRate() float32`

GetErrorRate returns the ErrorRate field if non-nil, zero value otherwise.

### GetErrorRateOk

`func (o *SubsystemRow) GetErrorRateOk() (*float32, bool)`

GetErrorRateOk returns a tuple with the ErrorRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorRate

`func (o *SubsystemRow) SetErrorRate(v float32)`

SetErrorRate sets ErrorRate field to given value.

### HasErrorRate

`func (o *SubsystemRow) HasErrorRate() bool`

HasErrorRate returns a boolean if a field has been set.

### GetErrors

`func (o *SubsystemRow) GetErrors() int32`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *SubsystemRow) GetErrorsOk() (*int32, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *SubsystemRow) SetErrors(v int32)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *SubsystemRow) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetLastErrorAt

`func (o *SubsystemRow) GetLastErrorAt() string`

GetLastErrorAt returns the LastErrorAt field if non-nil, zero value otherwise.

### GetLastErrorAtOk

`func (o *SubsystemRow) GetLastErrorAtOk() (*string, bool)`

GetLastErrorAtOk returns a tuple with the LastErrorAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastErrorAt

`func (o *SubsystemRow) SetLastErrorAt(v string)`

SetLastErrorAt sets LastErrorAt field to given value.

### HasLastErrorAt

`func (o *SubsystemRow) HasLastErrorAt() bool`

HasLastErrorAt returns a boolean if a field has been set.

### GetLastErrorMessage

`func (o *SubsystemRow) GetLastErrorMessage() string`

GetLastErrorMessage returns the LastErrorMessage field if non-nil, zero value otherwise.

### GetLastErrorMessageOk

`func (o *SubsystemRow) GetLastErrorMessageOk() (*string, bool)`

GetLastErrorMessageOk returns a tuple with the LastErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastErrorMessage

`func (o *SubsystemRow) SetLastErrorMessage(v string)`

SetLastErrorMessage sets LastErrorMessage field to given value.

### HasLastErrorMessage

`func (o *SubsystemRow) HasLastErrorMessage() bool`

HasLastErrorMessage returns a boolean if a field has been set.

### GetLastErrorRoute

`func (o *SubsystemRow) GetLastErrorRoute() string`

GetLastErrorRoute returns the LastErrorRoute field if non-nil, zero value otherwise.

### GetLastErrorRouteOk

`func (o *SubsystemRow) GetLastErrorRouteOk() (*string, bool)`

GetLastErrorRouteOk returns a tuple with the LastErrorRoute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastErrorRoute

`func (o *SubsystemRow) SetLastErrorRoute(v string)`

SetLastErrorRoute sets LastErrorRoute field to given value.

### HasLastErrorRoute

`func (o *SubsystemRow) HasLastErrorRoute() bool`

HasLastErrorRoute returns a boolean if a field has been set.

### GetLastErrorStatus

`func (o *SubsystemRow) GetLastErrorStatus() string`

GetLastErrorStatus returns the LastErrorStatus field if non-nil, zero value otherwise.

### GetLastErrorStatusOk

`func (o *SubsystemRow) GetLastErrorStatusOk() (*string, bool)`

GetLastErrorStatusOk returns a tuple with the LastErrorStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastErrorStatus

`func (o *SubsystemRow) SetLastErrorStatus(v string)`

SetLastErrorStatus sets LastErrorStatus field to given value.

### HasLastErrorStatus

`func (o *SubsystemRow) HasLastErrorStatus() bool`

HasLastErrorStatus returns a boolean if a field has been set.

### GetLatencyP50Ms

`func (o *SubsystemRow) GetLatencyP50Ms() float32`

GetLatencyP50Ms returns the LatencyP50Ms field if non-nil, zero value otherwise.

### GetLatencyP50MsOk

`func (o *SubsystemRow) GetLatencyP50MsOk() (*float32, bool)`

GetLatencyP50MsOk returns a tuple with the LatencyP50Ms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatencyP50Ms

`func (o *SubsystemRow) SetLatencyP50Ms(v float32)`

SetLatencyP50Ms sets LatencyP50Ms field to given value.

### HasLatencyP50Ms

`func (o *SubsystemRow) HasLatencyP50Ms() bool`

HasLatencyP50Ms returns a boolean if a field has been set.

### GetLatencyP95Ms

`func (o *SubsystemRow) GetLatencyP95Ms() float32`

GetLatencyP95Ms returns the LatencyP95Ms field if non-nil, zero value otherwise.

### GetLatencyP95MsOk

`func (o *SubsystemRow) GetLatencyP95MsOk() (*float32, bool)`

GetLatencyP95MsOk returns a tuple with the LatencyP95Ms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatencyP95Ms

`func (o *SubsystemRow) SetLatencyP95Ms(v float32)`

SetLatencyP95Ms sets LatencyP95Ms field to given value.

### HasLatencyP95Ms

`func (o *SubsystemRow) HasLatencyP95Ms() bool`

HasLatencyP95Ms returns a boolean if a field has been set.

### GetLatencyP99Ms

`func (o *SubsystemRow) GetLatencyP99Ms() float32`

GetLatencyP99Ms returns the LatencyP99Ms field if non-nil, zero value otherwise.

### GetLatencyP99MsOk

`func (o *SubsystemRow) GetLatencyP99MsOk() (*float32, bool)`

GetLatencyP99MsOk returns a tuple with the LatencyP99Ms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatencyP99Ms

`func (o *SubsystemRow) SetLatencyP99Ms(v float32)`

SetLatencyP99Ms sets LatencyP99Ms field to given value.

### HasLatencyP99Ms

`func (o *SubsystemRow) HasLatencyP99Ms() bool`

HasLatencyP99Ms returns a boolean if a field has been set.

### GetName

`func (o *SubsystemRow) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SubsystemRow) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SubsystemRow) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SubsystemRow) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPrefixes

`func (o *SubsystemRow) GetPrefixes() []string`

GetPrefixes returns the Prefixes field if non-nil, zero value otherwise.

### GetPrefixesOk

`func (o *SubsystemRow) GetPrefixesOk() (*[]string, bool)`

GetPrefixesOk returns a tuple with the Prefixes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefixes

`func (o *SubsystemRow) SetPrefixes(v []string)`

SetPrefixes sets Prefixes field to given value.

### HasPrefixes

`func (o *SubsystemRow) HasPrefixes() bool`

HasPrefixes returns a boolean if a field has been set.

### GetRequests

`func (o *SubsystemRow) GetRequests() int32`

GetRequests returns the Requests field if non-nil, zero value otherwise.

### GetRequestsOk

`func (o *SubsystemRow) GetRequestsOk() (*int32, bool)`

GetRequestsOk returns a tuple with the Requests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequests

`func (o *SubsystemRow) SetRequests(v int32)`

SetRequests sets Requests field to given value.

### HasRequests

`func (o *SubsystemRow) HasRequests() bool`

HasRequests returns a boolean if a field has been set.

### GetRequestsPerMin

`func (o *SubsystemRow) GetRequestsPerMin() float32`

GetRequestsPerMin returns the RequestsPerMin field if non-nil, zero value otherwise.

### GetRequestsPerMinOk

`func (o *SubsystemRow) GetRequestsPerMinOk() (*float32, bool)`

GetRequestsPerMinOk returns a tuple with the RequestsPerMin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestsPerMin

`func (o *SubsystemRow) SetRequestsPerMin(v float32)`

SetRequestsPerMin sets RequestsPerMin field to given value.

### HasRequestsPerMin

`func (o *SubsystemRow) HasRequestsPerMin() bool`

HasRequestsPerMin returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


