# CloudSubsystemRow

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

### NewCloudSubsystemRow

`func NewCloudSubsystemRow() *CloudSubsystemRow`

NewCloudSubsystemRow instantiates a new CloudSubsystemRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudSubsystemRowWithDefaults

`func NewCloudSubsystemRowWithDefaults() *CloudSubsystemRow`

NewCloudSubsystemRowWithDefaults instantiates a new CloudSubsystemRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *CloudSubsystemRow) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *CloudSubsystemRow) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *CloudSubsystemRow) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *CloudSubsystemRow) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetErrorRate

`func (o *CloudSubsystemRow) GetErrorRate() float32`

GetErrorRate returns the ErrorRate field if non-nil, zero value otherwise.

### GetErrorRateOk

`func (o *CloudSubsystemRow) GetErrorRateOk() (*float32, bool)`

GetErrorRateOk returns a tuple with the ErrorRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorRate

`func (o *CloudSubsystemRow) SetErrorRate(v float32)`

SetErrorRate sets ErrorRate field to given value.

### HasErrorRate

`func (o *CloudSubsystemRow) HasErrorRate() bool`

HasErrorRate returns a boolean if a field has been set.

### GetErrors

`func (o *CloudSubsystemRow) GetErrors() int32`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *CloudSubsystemRow) GetErrorsOk() (*int32, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *CloudSubsystemRow) SetErrors(v int32)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *CloudSubsystemRow) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetLastErrorAt

`func (o *CloudSubsystemRow) GetLastErrorAt() string`

GetLastErrorAt returns the LastErrorAt field if non-nil, zero value otherwise.

### GetLastErrorAtOk

`func (o *CloudSubsystemRow) GetLastErrorAtOk() (*string, bool)`

GetLastErrorAtOk returns a tuple with the LastErrorAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastErrorAt

`func (o *CloudSubsystemRow) SetLastErrorAt(v string)`

SetLastErrorAt sets LastErrorAt field to given value.

### HasLastErrorAt

`func (o *CloudSubsystemRow) HasLastErrorAt() bool`

HasLastErrorAt returns a boolean if a field has been set.

### GetLastErrorMessage

`func (o *CloudSubsystemRow) GetLastErrorMessage() string`

GetLastErrorMessage returns the LastErrorMessage field if non-nil, zero value otherwise.

### GetLastErrorMessageOk

`func (o *CloudSubsystemRow) GetLastErrorMessageOk() (*string, bool)`

GetLastErrorMessageOk returns a tuple with the LastErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastErrorMessage

`func (o *CloudSubsystemRow) SetLastErrorMessage(v string)`

SetLastErrorMessage sets LastErrorMessage field to given value.

### HasLastErrorMessage

`func (o *CloudSubsystemRow) HasLastErrorMessage() bool`

HasLastErrorMessage returns a boolean if a field has been set.

### GetLastErrorRoute

`func (o *CloudSubsystemRow) GetLastErrorRoute() string`

GetLastErrorRoute returns the LastErrorRoute field if non-nil, zero value otherwise.

### GetLastErrorRouteOk

`func (o *CloudSubsystemRow) GetLastErrorRouteOk() (*string, bool)`

GetLastErrorRouteOk returns a tuple with the LastErrorRoute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastErrorRoute

`func (o *CloudSubsystemRow) SetLastErrorRoute(v string)`

SetLastErrorRoute sets LastErrorRoute field to given value.

### HasLastErrorRoute

`func (o *CloudSubsystemRow) HasLastErrorRoute() bool`

HasLastErrorRoute returns a boolean if a field has been set.

### GetLastErrorStatus

`func (o *CloudSubsystemRow) GetLastErrorStatus() string`

GetLastErrorStatus returns the LastErrorStatus field if non-nil, zero value otherwise.

### GetLastErrorStatusOk

`func (o *CloudSubsystemRow) GetLastErrorStatusOk() (*string, bool)`

GetLastErrorStatusOk returns a tuple with the LastErrorStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastErrorStatus

`func (o *CloudSubsystemRow) SetLastErrorStatus(v string)`

SetLastErrorStatus sets LastErrorStatus field to given value.

### HasLastErrorStatus

`func (o *CloudSubsystemRow) HasLastErrorStatus() bool`

HasLastErrorStatus returns a boolean if a field has been set.

### GetLatencyP50Ms

`func (o *CloudSubsystemRow) GetLatencyP50Ms() float32`

GetLatencyP50Ms returns the LatencyP50Ms field if non-nil, zero value otherwise.

### GetLatencyP50MsOk

`func (o *CloudSubsystemRow) GetLatencyP50MsOk() (*float32, bool)`

GetLatencyP50MsOk returns a tuple with the LatencyP50Ms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatencyP50Ms

`func (o *CloudSubsystemRow) SetLatencyP50Ms(v float32)`

SetLatencyP50Ms sets LatencyP50Ms field to given value.

### HasLatencyP50Ms

`func (o *CloudSubsystemRow) HasLatencyP50Ms() bool`

HasLatencyP50Ms returns a boolean if a field has been set.

### GetLatencyP95Ms

`func (o *CloudSubsystemRow) GetLatencyP95Ms() float32`

GetLatencyP95Ms returns the LatencyP95Ms field if non-nil, zero value otherwise.

### GetLatencyP95MsOk

`func (o *CloudSubsystemRow) GetLatencyP95MsOk() (*float32, bool)`

GetLatencyP95MsOk returns a tuple with the LatencyP95Ms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatencyP95Ms

`func (o *CloudSubsystemRow) SetLatencyP95Ms(v float32)`

SetLatencyP95Ms sets LatencyP95Ms field to given value.

### HasLatencyP95Ms

`func (o *CloudSubsystemRow) HasLatencyP95Ms() bool`

HasLatencyP95Ms returns a boolean if a field has been set.

### GetLatencyP99Ms

`func (o *CloudSubsystemRow) GetLatencyP99Ms() float32`

GetLatencyP99Ms returns the LatencyP99Ms field if non-nil, zero value otherwise.

### GetLatencyP99MsOk

`func (o *CloudSubsystemRow) GetLatencyP99MsOk() (*float32, bool)`

GetLatencyP99MsOk returns a tuple with the LatencyP99Ms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatencyP99Ms

`func (o *CloudSubsystemRow) SetLatencyP99Ms(v float32)`

SetLatencyP99Ms sets LatencyP99Ms field to given value.

### HasLatencyP99Ms

`func (o *CloudSubsystemRow) HasLatencyP99Ms() bool`

HasLatencyP99Ms returns a boolean if a field has been set.

### GetName

`func (o *CloudSubsystemRow) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CloudSubsystemRow) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CloudSubsystemRow) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CloudSubsystemRow) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPrefixes

`func (o *CloudSubsystemRow) GetPrefixes() []string`

GetPrefixes returns the Prefixes field if non-nil, zero value otherwise.

### GetPrefixesOk

`func (o *CloudSubsystemRow) GetPrefixesOk() (*[]string, bool)`

GetPrefixesOk returns a tuple with the Prefixes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefixes

`func (o *CloudSubsystemRow) SetPrefixes(v []string)`

SetPrefixes sets Prefixes field to given value.

### HasPrefixes

`func (o *CloudSubsystemRow) HasPrefixes() bool`

HasPrefixes returns a boolean if a field has been set.

### GetRequests

`func (o *CloudSubsystemRow) GetRequests() int32`

GetRequests returns the Requests field if non-nil, zero value otherwise.

### GetRequestsOk

`func (o *CloudSubsystemRow) GetRequestsOk() (*int32, bool)`

GetRequestsOk returns a tuple with the Requests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequests

`func (o *CloudSubsystemRow) SetRequests(v int32)`

SetRequests sets Requests field to given value.

### HasRequests

`func (o *CloudSubsystemRow) HasRequests() bool`

HasRequests returns a boolean if a field has been set.

### GetRequestsPerMin

`func (o *CloudSubsystemRow) GetRequestsPerMin() float32`

GetRequestsPerMin returns the RequestsPerMin field if non-nil, zero value otherwise.

### GetRequestsPerMinOk

`func (o *CloudSubsystemRow) GetRequestsPerMinOk() (*float32, bool)`

GetRequestsPerMinOk returns a tuple with the RequestsPerMin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestsPerMin

`func (o *CloudSubsystemRow) SetRequestsPerMin(v float32)`

SetRequestsPerMin sets RequestsPerMin field to given value.

### HasRequestsPerMin

`func (o *CloudSubsystemRow) HasRequestsPerMin() bool`

HasRequestsPerMin returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


