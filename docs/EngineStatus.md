# EngineStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Reachable** | Pointer to **bool** | Reachable is true when the engine answered its health probe. | [optional] 
**Revision** | Pointer to **string** | Revision is the engine build&#39;s git revision, present only when reachable (the server&#39;s own build identity — it publishes no semver). | [optional] 

## Methods

### NewEngineStatus

`func NewEngineStatus() *EngineStatus`

NewEngineStatus instantiates a new EngineStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEngineStatusWithDefaults

`func NewEngineStatusWithDefaults() *EngineStatus`

NewEngineStatusWithDefaults instantiates a new EngineStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReachable

`func (o *EngineStatus) GetReachable() bool`

GetReachable returns the Reachable field if non-nil, zero value otherwise.

### GetReachableOk

`func (o *EngineStatus) GetReachableOk() (*bool, bool)`

GetReachableOk returns a tuple with the Reachable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReachable

`func (o *EngineStatus) SetReachable(v bool)`

SetReachable sets Reachable field to given value.

### HasReachable

`func (o *EngineStatus) HasReachable() bool`

HasReachable returns a boolean if a field has been set.

### GetRevision

`func (o *EngineStatus) GetRevision() string`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *EngineStatus) GetRevisionOk() (*string, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *EngineStatus) SetRevision(v string)`

SetRevision sets Revision field to given value.

### HasRevision

`func (o *EngineStatus) HasRevision() bool`

HasRevision returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


