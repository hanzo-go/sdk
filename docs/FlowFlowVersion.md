# FlowFlowVersion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**FlowId** | Pointer to **string** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**Trigger** | Pointer to **map[string]interface{}** |  | [optional] 
**Valid** | Pointer to **bool** |  | [optional] 
**State** | Pointer to **string** |  | [optional] 
**UpdatedBy** | Pointer to **string** |  | [optional] 
**Created** | Pointer to **time.Time** |  | [optional] 
**Updated** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewFlowFlowVersion

`func NewFlowFlowVersion() *FlowFlowVersion`

NewFlowFlowVersion instantiates a new FlowFlowVersion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFlowFlowVersionWithDefaults

`func NewFlowFlowVersionWithDefaults() *FlowFlowVersion`

NewFlowFlowVersionWithDefaults instantiates a new FlowFlowVersion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *FlowFlowVersion) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FlowFlowVersion) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FlowFlowVersion) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *FlowFlowVersion) HasId() bool`

HasId returns a boolean if a field has been set.

### GetFlowId

`func (o *FlowFlowVersion) GetFlowId() string`

GetFlowId returns the FlowId field if non-nil, zero value otherwise.

### GetFlowIdOk

`func (o *FlowFlowVersion) GetFlowIdOk() (*string, bool)`

GetFlowIdOk returns a tuple with the FlowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowId

`func (o *FlowFlowVersion) SetFlowId(v string)`

SetFlowId sets FlowId field to given value.

### HasFlowId

`func (o *FlowFlowVersion) HasFlowId() bool`

HasFlowId returns a boolean if a field has been set.

### GetDisplayName

`func (o *FlowFlowVersion) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *FlowFlowVersion) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *FlowFlowVersion) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *FlowFlowVersion) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetTrigger

`func (o *FlowFlowVersion) GetTrigger() map[string]interface{}`

GetTrigger returns the Trigger field if non-nil, zero value otherwise.

### GetTriggerOk

`func (o *FlowFlowVersion) GetTriggerOk() (*map[string]interface{}, bool)`

GetTriggerOk returns a tuple with the Trigger field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrigger

`func (o *FlowFlowVersion) SetTrigger(v map[string]interface{})`

SetTrigger sets Trigger field to given value.

### HasTrigger

`func (o *FlowFlowVersion) HasTrigger() bool`

HasTrigger returns a boolean if a field has been set.

### GetValid

`func (o *FlowFlowVersion) GetValid() bool`

GetValid returns the Valid field if non-nil, zero value otherwise.

### GetValidOk

`func (o *FlowFlowVersion) GetValidOk() (*bool, bool)`

GetValidOk returns a tuple with the Valid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValid

`func (o *FlowFlowVersion) SetValid(v bool)`

SetValid sets Valid field to given value.

### HasValid

`func (o *FlowFlowVersion) HasValid() bool`

HasValid returns a boolean if a field has been set.

### GetState

`func (o *FlowFlowVersion) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *FlowFlowVersion) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *FlowFlowVersion) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *FlowFlowVersion) HasState() bool`

HasState returns a boolean if a field has been set.

### GetUpdatedBy

`func (o *FlowFlowVersion) GetUpdatedBy() string`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *FlowFlowVersion) GetUpdatedByOk() (*string, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *FlowFlowVersion) SetUpdatedBy(v string)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *FlowFlowVersion) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.

### GetCreated

`func (o *FlowFlowVersion) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *FlowFlowVersion) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *FlowFlowVersion) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *FlowFlowVersion) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetUpdated

`func (o *FlowFlowVersion) GetUpdated() time.Time`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *FlowFlowVersion) GetUpdatedOk() (*time.Time, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *FlowFlowVersion) SetUpdated(v time.Time)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *FlowFlowVersion) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


