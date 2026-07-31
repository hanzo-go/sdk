# AutomationsFlowVersion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**FlowId** | Pointer to **string** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**Trigger** | Pointer to [**AutomationsFlowTrigger**](AutomationsFlowTrigger.md) |  | [optional] 
**Valid** | Pointer to **bool** |  | [optional] 
**State** | Pointer to **string** |  | [optional] 
**SchemaVersion** | Pointer to **string** |  | [optional] 
**Created** | Pointer to **int64** |  | [optional] 
**Updated** | Pointer to **int64** |  | [optional] 

## Methods

### NewAutomationsFlowVersion

`func NewAutomationsFlowVersion() *AutomationsFlowVersion`

NewAutomationsFlowVersion instantiates a new AutomationsFlowVersion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutomationsFlowVersionWithDefaults

`func NewAutomationsFlowVersionWithDefaults() *AutomationsFlowVersion`

NewAutomationsFlowVersionWithDefaults instantiates a new AutomationsFlowVersion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AutomationsFlowVersion) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AutomationsFlowVersion) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AutomationsFlowVersion) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AutomationsFlowVersion) HasId() bool`

HasId returns a boolean if a field has been set.

### GetFlowId

`func (o *AutomationsFlowVersion) GetFlowId() string`

GetFlowId returns the FlowId field if non-nil, zero value otherwise.

### GetFlowIdOk

`func (o *AutomationsFlowVersion) GetFlowIdOk() (*string, bool)`

GetFlowIdOk returns a tuple with the FlowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowId

`func (o *AutomationsFlowVersion) SetFlowId(v string)`

SetFlowId sets FlowId field to given value.

### HasFlowId

`func (o *AutomationsFlowVersion) HasFlowId() bool`

HasFlowId returns a boolean if a field has been set.

### GetDisplayName

`func (o *AutomationsFlowVersion) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *AutomationsFlowVersion) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *AutomationsFlowVersion) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *AutomationsFlowVersion) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetTrigger

`func (o *AutomationsFlowVersion) GetTrigger() AutomationsFlowTrigger`

GetTrigger returns the Trigger field if non-nil, zero value otherwise.

### GetTriggerOk

`func (o *AutomationsFlowVersion) GetTriggerOk() (*AutomationsFlowTrigger, bool)`

GetTriggerOk returns a tuple with the Trigger field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrigger

`func (o *AutomationsFlowVersion) SetTrigger(v AutomationsFlowTrigger)`

SetTrigger sets Trigger field to given value.

### HasTrigger

`func (o *AutomationsFlowVersion) HasTrigger() bool`

HasTrigger returns a boolean if a field has been set.

### GetValid

`func (o *AutomationsFlowVersion) GetValid() bool`

GetValid returns the Valid field if non-nil, zero value otherwise.

### GetValidOk

`func (o *AutomationsFlowVersion) GetValidOk() (*bool, bool)`

GetValidOk returns a tuple with the Valid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValid

`func (o *AutomationsFlowVersion) SetValid(v bool)`

SetValid sets Valid field to given value.

### HasValid

`func (o *AutomationsFlowVersion) HasValid() bool`

HasValid returns a boolean if a field has been set.

### GetState

`func (o *AutomationsFlowVersion) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *AutomationsFlowVersion) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *AutomationsFlowVersion) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *AutomationsFlowVersion) HasState() bool`

HasState returns a boolean if a field has been set.

### GetSchemaVersion

`func (o *AutomationsFlowVersion) GetSchemaVersion() string`

GetSchemaVersion returns the SchemaVersion field if non-nil, zero value otherwise.

### GetSchemaVersionOk

`func (o *AutomationsFlowVersion) GetSchemaVersionOk() (*string, bool)`

GetSchemaVersionOk returns a tuple with the SchemaVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaVersion

`func (o *AutomationsFlowVersion) SetSchemaVersion(v string)`

SetSchemaVersion sets SchemaVersion field to given value.

### HasSchemaVersion

`func (o *AutomationsFlowVersion) HasSchemaVersion() bool`

HasSchemaVersion returns a boolean if a field has been set.

### GetCreated

`func (o *AutomationsFlowVersion) GetCreated() int64`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *AutomationsFlowVersion) GetCreatedOk() (*int64, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *AutomationsFlowVersion) SetCreated(v int64)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *AutomationsFlowVersion) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetUpdated

`func (o *AutomationsFlowVersion) GetUpdated() int64`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *AutomationsFlowVersion) GetUpdatedOk() (*int64, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *AutomationsFlowVersion) SetUpdated(v int64)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *AutomationsFlowVersion) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


