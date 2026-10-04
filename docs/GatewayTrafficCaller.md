# GatewayTrafficCaller

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to **string** | Action is the verdict currently held against it, if any. | [optional] 
**Cred** | Pointer to **string** | Cred is the caller&#39;s key: a credential fingerprint (a per-process one-way digest, not a key) for a validated caller, and \&quot;ip:&lt;addr&gt;\&quot; for one that presented no credential we could validate. | [optional] 
**Failures** | Pointer to **int64** | Failures is how many ended 401 or 403. | [optional] 
**HeldUntil** | Pointer to **int64** | HeldUntil is when the held verdict lapses, unix seconds. | [optional] 
**Paths** | Pointer to **int64** | Paths is the approximate number of distinct paths it touched (max 64). | [optional] 
**Reason** | Pointer to **string** | Reason is why that verdict was reached. | [optional] 
**Requests** | Pointer to **int64** | Requests is its request count in the window. | [optional] 

## Methods

### NewGatewayTrafficCaller

`func NewGatewayTrafficCaller() *GatewayTrafficCaller`

NewGatewayTrafficCaller instantiates a new GatewayTrafficCaller object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGatewayTrafficCallerWithDefaults

`func NewGatewayTrafficCallerWithDefaults() *GatewayTrafficCaller`

NewGatewayTrafficCallerWithDefaults instantiates a new GatewayTrafficCaller object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *GatewayTrafficCaller) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *GatewayTrafficCaller) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *GatewayTrafficCaller) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *GatewayTrafficCaller) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetCred

`func (o *GatewayTrafficCaller) GetCred() string`

GetCred returns the Cred field if non-nil, zero value otherwise.

### GetCredOk

`func (o *GatewayTrafficCaller) GetCredOk() (*string, bool)`

GetCredOk returns a tuple with the Cred field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCred

`func (o *GatewayTrafficCaller) SetCred(v string)`

SetCred sets Cred field to given value.

### HasCred

`func (o *GatewayTrafficCaller) HasCred() bool`

HasCred returns a boolean if a field has been set.

### GetFailures

`func (o *GatewayTrafficCaller) GetFailures() int64`

GetFailures returns the Failures field if non-nil, zero value otherwise.

### GetFailuresOk

`func (o *GatewayTrafficCaller) GetFailuresOk() (*int64, bool)`

GetFailuresOk returns a tuple with the Failures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailures

`func (o *GatewayTrafficCaller) SetFailures(v int64)`

SetFailures sets Failures field to given value.

### HasFailures

`func (o *GatewayTrafficCaller) HasFailures() bool`

HasFailures returns a boolean if a field has been set.

### GetHeldUntil

`func (o *GatewayTrafficCaller) GetHeldUntil() int64`

GetHeldUntil returns the HeldUntil field if non-nil, zero value otherwise.

### GetHeldUntilOk

`func (o *GatewayTrafficCaller) GetHeldUntilOk() (*int64, bool)`

GetHeldUntilOk returns a tuple with the HeldUntil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeldUntil

`func (o *GatewayTrafficCaller) SetHeldUntil(v int64)`

SetHeldUntil sets HeldUntil field to given value.

### HasHeldUntil

`func (o *GatewayTrafficCaller) HasHeldUntil() bool`

HasHeldUntil returns a boolean if a field has been set.

### GetPaths

`func (o *GatewayTrafficCaller) GetPaths() int64`

GetPaths returns the Paths field if non-nil, zero value otherwise.

### GetPathsOk

`func (o *GatewayTrafficCaller) GetPathsOk() (*int64, bool)`

GetPathsOk returns a tuple with the Paths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaths

`func (o *GatewayTrafficCaller) SetPaths(v int64)`

SetPaths sets Paths field to given value.

### HasPaths

`func (o *GatewayTrafficCaller) HasPaths() bool`

HasPaths returns a boolean if a field has been set.

### GetReason

`func (o *GatewayTrafficCaller) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *GatewayTrafficCaller) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *GatewayTrafficCaller) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *GatewayTrafficCaller) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetRequests

`func (o *GatewayTrafficCaller) GetRequests() int64`

GetRequests returns the Requests field if non-nil, zero value otherwise.

### GetRequestsOk

`func (o *GatewayTrafficCaller) GetRequestsOk() (*int64, bool)`

GetRequestsOk returns a tuple with the Requests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequests

`func (o *GatewayTrafficCaller) SetRequests(v int64)`

SetRequests sets Requests field to given value.

### HasRequests

`func (o *GatewayTrafficCaller) HasRequests() bool`

HasRequests returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


