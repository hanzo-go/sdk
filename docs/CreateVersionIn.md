# CreateVersionIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | Pointer to **string** | DisplayName names the new version. | [optional] 
**Id** | Pointer to **string** | ID is the flow to add a version to, from the path. | [optional] 
**Trigger** | Pointer to [**FlowTrigger**](FlowTrigger.md) | Trigger is the root of the version&#39;s step tree. Optional: a version with no trigger is created invalid, and cannot run until one is set. | [optional] 

## Methods

### NewCreateVersionIn

`func NewCreateVersionIn() *CreateVersionIn`

NewCreateVersionIn instantiates a new CreateVersionIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateVersionInWithDefaults

`func NewCreateVersionInWithDefaults() *CreateVersionIn`

NewCreateVersionInWithDefaults instantiates a new CreateVersionIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayName

`func (o *CreateVersionIn) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *CreateVersionIn) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *CreateVersionIn) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *CreateVersionIn) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetId

`func (o *CreateVersionIn) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CreateVersionIn) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CreateVersionIn) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CreateVersionIn) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTrigger

`func (o *CreateVersionIn) GetTrigger() FlowTrigger`

GetTrigger returns the Trigger field if non-nil, zero value otherwise.

### GetTriggerOk

`func (o *CreateVersionIn) GetTriggerOk() (*FlowTrigger, bool)`

GetTriggerOk returns a tuple with the Trigger field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrigger

`func (o *CreateVersionIn) SetTrigger(v FlowTrigger)`

SetTrigger sets Trigger field to given value.

### HasTrigger

`func (o *CreateVersionIn) HasTrigger() bool`

HasTrigger returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


