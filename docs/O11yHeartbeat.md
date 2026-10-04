# O11yHeartbeat

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AgeSeconds** | Pointer to **int64** | AgeSeconds is how long ago the last heartbeat arrived — or this process started, before its first. | [optional] 
**LimitSeconds** | Pointer to **int64** | LimitSeconds is how old the heartbeat may get before this answers 503. | [optional] 

## Methods

### NewO11yHeartbeat

`func NewO11yHeartbeat() *O11yHeartbeat`

NewO11yHeartbeat instantiates a new O11yHeartbeat object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yHeartbeatWithDefaults

`func NewO11yHeartbeatWithDefaults() *O11yHeartbeat`

NewO11yHeartbeatWithDefaults instantiates a new O11yHeartbeat object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAgeSeconds

`func (o *O11yHeartbeat) GetAgeSeconds() int64`

GetAgeSeconds returns the AgeSeconds field if non-nil, zero value otherwise.

### GetAgeSecondsOk

`func (o *O11yHeartbeat) GetAgeSecondsOk() (*int64, bool)`

GetAgeSecondsOk returns a tuple with the AgeSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgeSeconds

`func (o *O11yHeartbeat) SetAgeSeconds(v int64)`

SetAgeSeconds sets AgeSeconds field to given value.

### HasAgeSeconds

`func (o *O11yHeartbeat) HasAgeSeconds() bool`

HasAgeSeconds returns a boolean if a field has been set.

### GetLimitSeconds

`func (o *O11yHeartbeat) GetLimitSeconds() int64`

GetLimitSeconds returns the LimitSeconds field if non-nil, zero value otherwise.

### GetLimitSecondsOk

`func (o *O11yHeartbeat) GetLimitSecondsOk() (*int64, bool)`

GetLimitSecondsOk returns a tuple with the LimitSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimitSeconds

`func (o *O11yHeartbeat) SetLimitSeconds(v int64)`

SetLimitSeconds sets LimitSeconds field to given value.

### HasLimitSeconds

`func (o *O11yHeartbeat) HasLimitSeconds() bool`

HasLimitSeconds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


