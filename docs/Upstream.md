# Upstream

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Backends** | Pointer to [**[]Backend**](Backend.md) | Backends are the upstream servers to balance across: 1..32 of them. | [optional] 
**Id** | Pointer to **string** | ID identifies the pool within the org: [A-Za-z0-9-_.], at most 128 chars. A create that omits it gets a generated one. Routes reference it by this id. | [optional] 
**PassHostHeader** | Pointer to **bool** | PassHostHeader forwards the client&#39;s original Host header upstream instead of rewriting it to the backend&#39;s. | [optional] 

## Methods

### NewUpstream

`func NewUpstream() *Upstream`

NewUpstream instantiates a new Upstream object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpstreamWithDefaults

`func NewUpstreamWithDefaults() *Upstream`

NewUpstreamWithDefaults instantiates a new Upstream object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBackends

`func (o *Upstream) GetBackends() []Backend`

GetBackends returns the Backends field if non-nil, zero value otherwise.

### GetBackendsOk

`func (o *Upstream) GetBackendsOk() (*[]Backend, bool)`

GetBackendsOk returns a tuple with the Backends field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackends

`func (o *Upstream) SetBackends(v []Backend)`

SetBackends sets Backends field to given value.

### HasBackends

`func (o *Upstream) HasBackends() bool`

HasBackends returns a boolean if a field has been set.

### GetId

`func (o *Upstream) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Upstream) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Upstream) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Upstream) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPassHostHeader

`func (o *Upstream) GetPassHostHeader() bool`

GetPassHostHeader returns the PassHostHeader field if non-nil, zero value otherwise.

### GetPassHostHeaderOk

`func (o *Upstream) GetPassHostHeaderOk() (*bool, bool)`

GetPassHostHeaderOk returns a tuple with the PassHostHeader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassHostHeader

`func (o *Upstream) SetPassHostHeader(v bool)`

SetPassHostHeader sets PassHostHeader field to given value.

### HasPassHostHeader

`func (o *Upstream) HasPassHostHeader() bool`

HasPassHostHeader returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


