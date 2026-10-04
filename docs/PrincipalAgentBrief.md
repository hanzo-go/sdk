# PrincipalAgentBrief

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | ID is the agent id — the ref a run names and the value a member id derives from, so it must be stable across a reconcile. | [optional] 
**Name** | Pointer to **string** | Name is the display name a person reads in a member list. | [optional] 
**Parent** | Pointer to **string** | Parent is the agent that spawned this one; empty for an agent a person defined. | [optional] 
**Status** | Pointer to **string** | Status is the registry status VERBATIM, never a boolean. The caller decides what counts as live: apps/team reads empty/active/ready as a live member and anything else as one that keeps its authorship and drops out of the roster, and folding that judgement into a bool here would move the policy to the wrong side of the wire. | [optional] 

## Methods

### NewPrincipalAgentBrief

`func NewPrincipalAgentBrief() *PrincipalAgentBrief`

NewPrincipalAgentBrief instantiates a new PrincipalAgentBrief object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrincipalAgentBriefWithDefaults

`func NewPrincipalAgentBriefWithDefaults() *PrincipalAgentBrief`

NewPrincipalAgentBriefWithDefaults instantiates a new PrincipalAgentBrief object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PrincipalAgentBrief) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PrincipalAgentBrief) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PrincipalAgentBrief) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PrincipalAgentBrief) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *PrincipalAgentBrief) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PrincipalAgentBrief) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PrincipalAgentBrief) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PrincipalAgentBrief) HasName() bool`

HasName returns a boolean if a field has been set.

### GetParent

`func (o *PrincipalAgentBrief) GetParent() string`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *PrincipalAgentBrief) GetParentOk() (*string, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *PrincipalAgentBrief) SetParent(v string)`

SetParent sets Parent field to given value.

### HasParent

`func (o *PrincipalAgentBrief) HasParent() bool`

HasParent returns a boolean if a field has been set.

### GetStatus

`func (o *PrincipalAgentBrief) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PrincipalAgentBrief) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PrincipalAgentBrief) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PrincipalAgentBrief) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


