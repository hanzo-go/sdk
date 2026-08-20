# EdgeState

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Configured** | Pointer to **bool** | Configured is whether the edge holds credentials to act at all. False means every purge is a no-op. | [optional] 
**Error** | Pointer to **string** | Error is the blocker, so an operator reads it instead of guessing at it. | [optional] 
**Freshness** | Pointer to **string** | Freshness says, in one phrase, how long after a publish a reader sees it. It is the sentence an operator actually wants; the booleans above are how a machine reads the same fact. | [optional] 
**Policy** | Pointer to **map[string]string** | Policy is the Cache-Control this edge serves each class of object with. It is DERIVED from the one canonical function, never a second copy: half the confusion when a publish looks stale is not knowing what the TTLs are, and reading them out of the source is not something an operator should have to do to answer \&quot;how long until this is live\&quot;. | [optional] 
**Provider** | Pointer to **string** | Provider is the CDN behind this edge, or \&quot;none\&quot;. It is the first thing an operator wants and the only vendor name this API returns. | [optional] 
**Reach** | Pointer to **[]string** | Reach is the apexes a publish is invalidated on. A site is served on more than one — the site plane&#39;s own and the first-party apex — and a purge that covers one of them looks identical from here to a purge that covers both. | [optional] 
**Status** | Pointer to **string** | Status is \&quot;ok\&quot; when a publish reaches readers immediately, else \&quot;degraded\&quot;. | [optional] 

## Methods

### NewEdgeState

`func NewEdgeState() *EdgeState`

NewEdgeState instantiates a new EdgeState object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEdgeStateWithDefaults

`func NewEdgeStateWithDefaults() *EdgeState`

NewEdgeStateWithDefaults instantiates a new EdgeState object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfigured

`func (o *EdgeState) GetConfigured() bool`

GetConfigured returns the Configured field if non-nil, zero value otherwise.

### GetConfiguredOk

`func (o *EdgeState) GetConfiguredOk() (*bool, bool)`

GetConfiguredOk returns a tuple with the Configured field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigured

`func (o *EdgeState) SetConfigured(v bool)`

SetConfigured sets Configured field to given value.

### HasConfigured

`func (o *EdgeState) HasConfigured() bool`

HasConfigured returns a boolean if a field has been set.

### GetError

`func (o *EdgeState) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *EdgeState) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *EdgeState) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *EdgeState) HasError() bool`

HasError returns a boolean if a field has been set.

### GetFreshness

`func (o *EdgeState) GetFreshness() string`

GetFreshness returns the Freshness field if non-nil, zero value otherwise.

### GetFreshnessOk

`func (o *EdgeState) GetFreshnessOk() (*string, bool)`

GetFreshnessOk returns a tuple with the Freshness field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreshness

`func (o *EdgeState) SetFreshness(v string)`

SetFreshness sets Freshness field to given value.

### HasFreshness

`func (o *EdgeState) HasFreshness() bool`

HasFreshness returns a boolean if a field has been set.

### GetPolicy

`func (o *EdgeState) GetPolicy() map[string]string`

GetPolicy returns the Policy field if non-nil, zero value otherwise.

### GetPolicyOk

`func (o *EdgeState) GetPolicyOk() (*map[string]string, bool)`

GetPolicyOk returns a tuple with the Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicy

`func (o *EdgeState) SetPolicy(v map[string]string)`

SetPolicy sets Policy field to given value.

### HasPolicy

`func (o *EdgeState) HasPolicy() bool`

HasPolicy returns a boolean if a field has been set.

### GetProvider

`func (o *EdgeState) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *EdgeState) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *EdgeState) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *EdgeState) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetReach

`func (o *EdgeState) GetReach() []string`

GetReach returns the Reach field if non-nil, zero value otherwise.

### GetReachOk

`func (o *EdgeState) GetReachOk() (*[]string, bool)`

GetReachOk returns a tuple with the Reach field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReach

`func (o *EdgeState) SetReach(v []string)`

SetReach sets Reach field to given value.

### HasReach

`func (o *EdgeState) HasReach() bool`

HasReach returns a boolean if a field has been set.

### GetStatus

`func (o *EdgeState) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *EdgeState) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *EdgeState) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *EdgeState) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


