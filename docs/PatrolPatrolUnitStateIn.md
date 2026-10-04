# PatrolPatrolUnitStateIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name is the unit&#39;s call sign, from the path. | [optional] 
**State** | Pointer to **string** | State is available, assigned, enroute, onsite, offline, break or emergency. | [optional] 

## Methods

### NewPatrolPatrolUnitStateIn

`func NewPatrolPatrolUnitStateIn() *PatrolPatrolUnitStateIn`

NewPatrolPatrolUnitStateIn instantiates a new PatrolPatrolUnitStateIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatrolPatrolUnitStateInWithDefaults

`func NewPatrolPatrolUnitStateInWithDefaults() *PatrolPatrolUnitStateIn`

NewPatrolPatrolUnitStateInWithDefaults instantiates a new PatrolPatrolUnitStateIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PatrolPatrolUnitStateIn) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatrolPatrolUnitStateIn) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatrolPatrolUnitStateIn) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatrolPatrolUnitStateIn) HasName() bool`

HasName returns a boolean if a field has been set.

### GetState

`func (o *PatrolPatrolUnitStateIn) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *PatrolPatrolUnitStateIn) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *PatrolPatrolUnitStateIn) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *PatrolPatrolUnitStateIn) HasState() bool`

HasState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


