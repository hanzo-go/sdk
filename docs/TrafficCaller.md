# TrafficCaller

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to **string** | Action is the verdict currently held against it, if any. | [optional] 
**Cred** | Pointer to **string** | Cred is the caller&#39;s key: a credential fingerprint (a per-process one-way digest, not a key) for a validated caller, and \&quot;ip:&lt;addr&gt;\&quot; for one that presented no credential we could validate. | [optional] 
**Failures** | Pointer to **int32** | Failures is how many ended 401 or 403. | [optional] 
**HeldUntil** | Pointer to **int32** | HeldUntil is when the held verdict lapses, unix seconds. | [optional] 
**Paths** | Pointer to **int32** | Paths is the approximate number of distinct paths it touched (max 64). | [optional] 
**Reason** | Pointer to **string** | Reason is why that verdict was reached. | [optional] 
**Requests** | Pointer to **int32** | Requests is its request count in the window. | [optional] 

## Methods

### NewTrafficCaller

`func NewTrafficCaller() *TrafficCaller`

NewTrafficCaller instantiates a new TrafficCaller object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrafficCallerWithDefaults

`func NewTrafficCallerWithDefaults() *TrafficCaller`

NewTrafficCallerWithDefaults instantiates a new TrafficCaller object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *TrafficCaller) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *TrafficCaller) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *TrafficCaller) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *TrafficCaller) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetCred

`func (o *TrafficCaller) GetCred() string`

GetCred returns the Cred field if non-nil, zero value otherwise.

### GetCredOk

`func (o *TrafficCaller) GetCredOk() (*string, bool)`

GetCredOk returns a tuple with the Cred field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCred

`func (o *TrafficCaller) SetCred(v string)`

SetCred sets Cred field to given value.

### HasCred

`func (o *TrafficCaller) HasCred() bool`

HasCred returns a boolean if a field has been set.

### GetFailures

`func (o *TrafficCaller) GetFailures() int32`

GetFailures returns the Failures field if non-nil, zero value otherwise.

### GetFailuresOk

`func (o *TrafficCaller) GetFailuresOk() (*int32, bool)`

GetFailuresOk returns a tuple with the Failures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailures

`func (o *TrafficCaller) SetFailures(v int32)`

SetFailures sets Failures field to given value.

### HasFailures

`func (o *TrafficCaller) HasFailures() bool`

HasFailures returns a boolean if a field has been set.

### GetHeldUntil

`func (o *TrafficCaller) GetHeldUntil() int32`

GetHeldUntil returns the HeldUntil field if non-nil, zero value otherwise.

### GetHeldUntilOk

`func (o *TrafficCaller) GetHeldUntilOk() (*int32, bool)`

GetHeldUntilOk returns a tuple with the HeldUntil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeldUntil

`func (o *TrafficCaller) SetHeldUntil(v int32)`

SetHeldUntil sets HeldUntil field to given value.

### HasHeldUntil

`func (o *TrafficCaller) HasHeldUntil() bool`

HasHeldUntil returns a boolean if a field has been set.

### GetPaths

`func (o *TrafficCaller) GetPaths() int32`

GetPaths returns the Paths field if non-nil, zero value otherwise.

### GetPathsOk

`func (o *TrafficCaller) GetPathsOk() (*int32, bool)`

GetPathsOk returns a tuple with the Paths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaths

`func (o *TrafficCaller) SetPaths(v int32)`

SetPaths sets Paths field to given value.

### HasPaths

`func (o *TrafficCaller) HasPaths() bool`

HasPaths returns a boolean if a field has been set.

### GetReason

`func (o *TrafficCaller) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *TrafficCaller) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *TrafficCaller) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *TrafficCaller) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetRequests

`func (o *TrafficCaller) GetRequests() int32`

GetRequests returns the Requests field if non-nil, zero value otherwise.

### GetRequestsOk

`func (o *TrafficCaller) GetRequestsOk() (*int32, bool)`

GetRequestsOk returns a tuple with the Requests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequests

`func (o *TrafficCaller) SetRequests(v int32)`

SetRequests sets Requests field to given value.

### HasRequests

`func (o *TrafficCaller) HasRequests() bool`

HasRequests returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


