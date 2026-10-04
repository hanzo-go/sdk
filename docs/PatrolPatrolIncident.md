# PatrolPatrolIncident

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Checked** | Pointer to **string** | Checked is when the premises had been walked. | [optional] 
**Clocks** | Pointer to [**[]PatrolPatrolClock**](PatrolPatrolClock.md) | Clocks is the SLA clocks, present on a single-incident read. | [optional] 
**Closed** | Pointer to **string** | Closed is when the incident was closed. | [optional] 
**Dispatched** | Pointer to **string** | Dispatched is when a unit was assigned. | [optional] 
**Enroute** | Pointer to **string** | Enroute is when that unit started moving. | [optional] 
**Escalated** | Pointer to **string** | Escalated is when it was escalated, or empty. | [optional] 
**False** | Pointer to **bool** | False says the activation turned out to be nothing. | [optional] 
**Filed** | Pointer to **string** | Filed is when the report was written. | [optional] 
**Name** | Pointer to **string** | Name is the incident&#39;s document name. | [optional] 
**Onsite** | Pointer to **string** | Onsite is when it arrived. | [optional] 
**Opened** | Pointer to **string** | Opened is when the activation landed. | [optional] 
**Outcome** | Pointer to **string** | Outcome is what it came to. | [optional] 
**Ref** | Pointer to **string** | Ref is the receiving centre&#39;s own reference. | [optional] 
**Secured** | Pointer to **string** | Secured is when the site was made safe. | [optional] 
**Site** | Pointer to **string** | Site is where it is happening. | [optional] 
**Sla** | Pointer to **int64** | SLA is the contracted response time in minutes. | [optional] 
**State** | Pointer to **string** | State is one of received, verified, dispatched, enroute, onsite, checked, secured, filed and closed. | [optional] 
**Steps** | Pointer to [**[]PatrolPatrolAct**](PatrolPatrolAct.md) | Steps is the act trail, present on a single-incident read. | [optional] 
**Tier** | Pointer to **string** | Tier is the site&#39;s contracted service tier. | [optional] 
**Trigger** | Pointer to **string** | Trigger is what the panel reported. | [optional] 
**Unit** | Pointer to **string** | Unit is the call sign attending, or empty. | [optional] 
**Verified** | Pointer to **string** | Verified is when the activation was confirmed real. | [optional] 
**Zone** | Pointer to **string** | Zone is the panel zone that reported. | [optional] 

## Methods

### NewPatrolPatrolIncident

`func NewPatrolPatrolIncident() *PatrolPatrolIncident`

NewPatrolPatrolIncident instantiates a new PatrolPatrolIncident object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatrolPatrolIncidentWithDefaults

`func NewPatrolPatrolIncidentWithDefaults() *PatrolPatrolIncident`

NewPatrolPatrolIncidentWithDefaults instantiates a new PatrolPatrolIncident object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChecked

`func (o *PatrolPatrolIncident) GetChecked() string`

GetChecked returns the Checked field if non-nil, zero value otherwise.

### GetCheckedOk

`func (o *PatrolPatrolIncident) GetCheckedOk() (*string, bool)`

GetCheckedOk returns a tuple with the Checked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChecked

`func (o *PatrolPatrolIncident) SetChecked(v string)`

SetChecked sets Checked field to given value.

### HasChecked

`func (o *PatrolPatrolIncident) HasChecked() bool`

HasChecked returns a boolean if a field has been set.

### GetClocks

`func (o *PatrolPatrolIncident) GetClocks() []PatrolPatrolClock`

GetClocks returns the Clocks field if non-nil, zero value otherwise.

### GetClocksOk

`func (o *PatrolPatrolIncident) GetClocksOk() (*[]PatrolPatrolClock, bool)`

GetClocksOk returns a tuple with the Clocks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClocks

`func (o *PatrolPatrolIncident) SetClocks(v []PatrolPatrolClock)`

SetClocks sets Clocks field to given value.

### HasClocks

`func (o *PatrolPatrolIncident) HasClocks() bool`

HasClocks returns a boolean if a field has been set.

### GetClosed

`func (o *PatrolPatrolIncident) GetClosed() string`

GetClosed returns the Closed field if non-nil, zero value otherwise.

### GetClosedOk

`func (o *PatrolPatrolIncident) GetClosedOk() (*string, bool)`

GetClosedOk returns a tuple with the Closed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClosed

`func (o *PatrolPatrolIncident) SetClosed(v string)`

SetClosed sets Closed field to given value.

### HasClosed

`func (o *PatrolPatrolIncident) HasClosed() bool`

HasClosed returns a boolean if a field has been set.

### GetDispatched

`func (o *PatrolPatrolIncident) GetDispatched() string`

GetDispatched returns the Dispatched field if non-nil, zero value otherwise.

### GetDispatchedOk

`func (o *PatrolPatrolIncident) GetDispatchedOk() (*string, bool)`

GetDispatchedOk returns a tuple with the Dispatched field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDispatched

`func (o *PatrolPatrolIncident) SetDispatched(v string)`

SetDispatched sets Dispatched field to given value.

### HasDispatched

`func (o *PatrolPatrolIncident) HasDispatched() bool`

HasDispatched returns a boolean if a field has been set.

### GetEnroute

`func (o *PatrolPatrolIncident) GetEnroute() string`

GetEnroute returns the Enroute field if non-nil, zero value otherwise.

### GetEnrouteOk

`func (o *PatrolPatrolIncident) GetEnrouteOk() (*string, bool)`

GetEnrouteOk returns a tuple with the Enroute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnroute

`func (o *PatrolPatrolIncident) SetEnroute(v string)`

SetEnroute sets Enroute field to given value.

### HasEnroute

`func (o *PatrolPatrolIncident) HasEnroute() bool`

HasEnroute returns a boolean if a field has been set.

### GetEscalated

`func (o *PatrolPatrolIncident) GetEscalated() string`

GetEscalated returns the Escalated field if non-nil, zero value otherwise.

### GetEscalatedOk

`func (o *PatrolPatrolIncident) GetEscalatedOk() (*string, bool)`

GetEscalatedOk returns a tuple with the Escalated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEscalated

`func (o *PatrolPatrolIncident) SetEscalated(v string)`

SetEscalated sets Escalated field to given value.

### HasEscalated

`func (o *PatrolPatrolIncident) HasEscalated() bool`

HasEscalated returns a boolean if a field has been set.

### GetFalse

`func (o *PatrolPatrolIncident) GetFalse() bool`

GetFalse returns the False field if non-nil, zero value otherwise.

### GetFalseOk

`func (o *PatrolPatrolIncident) GetFalseOk() (*bool, bool)`

GetFalseOk returns a tuple with the False field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFalse

`func (o *PatrolPatrolIncident) SetFalse(v bool)`

SetFalse sets False field to given value.

### HasFalse

`func (o *PatrolPatrolIncident) HasFalse() bool`

HasFalse returns a boolean if a field has been set.

### GetFiled

`func (o *PatrolPatrolIncident) GetFiled() string`

GetFiled returns the Filed field if non-nil, zero value otherwise.

### GetFiledOk

`func (o *PatrolPatrolIncident) GetFiledOk() (*string, bool)`

GetFiledOk returns a tuple with the Filed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiled

`func (o *PatrolPatrolIncident) SetFiled(v string)`

SetFiled sets Filed field to given value.

### HasFiled

`func (o *PatrolPatrolIncident) HasFiled() bool`

HasFiled returns a boolean if a field has been set.

### GetName

`func (o *PatrolPatrolIncident) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatrolPatrolIncident) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatrolPatrolIncident) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatrolPatrolIncident) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOnsite

`func (o *PatrolPatrolIncident) GetOnsite() string`

GetOnsite returns the Onsite field if non-nil, zero value otherwise.

### GetOnsiteOk

`func (o *PatrolPatrolIncident) GetOnsiteOk() (*string, bool)`

GetOnsiteOk returns a tuple with the Onsite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnsite

`func (o *PatrolPatrolIncident) SetOnsite(v string)`

SetOnsite sets Onsite field to given value.

### HasOnsite

`func (o *PatrolPatrolIncident) HasOnsite() bool`

HasOnsite returns a boolean if a field has been set.

### GetOpened

`func (o *PatrolPatrolIncident) GetOpened() string`

GetOpened returns the Opened field if non-nil, zero value otherwise.

### GetOpenedOk

`func (o *PatrolPatrolIncident) GetOpenedOk() (*string, bool)`

GetOpenedOk returns a tuple with the Opened field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpened

`func (o *PatrolPatrolIncident) SetOpened(v string)`

SetOpened sets Opened field to given value.

### HasOpened

`func (o *PatrolPatrolIncident) HasOpened() bool`

HasOpened returns a boolean if a field has been set.

### GetOutcome

`func (o *PatrolPatrolIncident) GetOutcome() string`

GetOutcome returns the Outcome field if non-nil, zero value otherwise.

### GetOutcomeOk

`func (o *PatrolPatrolIncident) GetOutcomeOk() (*string, bool)`

GetOutcomeOk returns a tuple with the Outcome field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutcome

`func (o *PatrolPatrolIncident) SetOutcome(v string)`

SetOutcome sets Outcome field to given value.

### HasOutcome

`func (o *PatrolPatrolIncident) HasOutcome() bool`

HasOutcome returns a boolean if a field has been set.

### GetRef

`func (o *PatrolPatrolIncident) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *PatrolPatrolIncident) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *PatrolPatrolIncident) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *PatrolPatrolIncident) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetSecured

`func (o *PatrolPatrolIncident) GetSecured() string`

GetSecured returns the Secured field if non-nil, zero value otherwise.

### GetSecuredOk

`func (o *PatrolPatrolIncident) GetSecuredOk() (*string, bool)`

GetSecuredOk returns a tuple with the Secured field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecured

`func (o *PatrolPatrolIncident) SetSecured(v string)`

SetSecured sets Secured field to given value.

### HasSecured

`func (o *PatrolPatrolIncident) HasSecured() bool`

HasSecured returns a boolean if a field has been set.

### GetSite

`func (o *PatrolPatrolIncident) GetSite() string`

GetSite returns the Site field if non-nil, zero value otherwise.

### GetSiteOk

`func (o *PatrolPatrolIncident) GetSiteOk() (*string, bool)`

GetSiteOk returns a tuple with the Site field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSite

`func (o *PatrolPatrolIncident) SetSite(v string)`

SetSite sets Site field to given value.

### HasSite

`func (o *PatrolPatrolIncident) HasSite() bool`

HasSite returns a boolean if a field has been set.

### GetSla

`func (o *PatrolPatrolIncident) GetSla() int64`

GetSla returns the Sla field if non-nil, zero value otherwise.

### GetSlaOk

`func (o *PatrolPatrolIncident) GetSlaOk() (*int64, bool)`

GetSlaOk returns a tuple with the Sla field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSla

`func (o *PatrolPatrolIncident) SetSla(v int64)`

SetSla sets Sla field to given value.

### HasSla

`func (o *PatrolPatrolIncident) HasSla() bool`

HasSla returns a boolean if a field has been set.

### GetState

`func (o *PatrolPatrolIncident) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *PatrolPatrolIncident) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *PatrolPatrolIncident) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *PatrolPatrolIncident) HasState() bool`

HasState returns a boolean if a field has been set.

### GetSteps

`func (o *PatrolPatrolIncident) GetSteps() []PatrolPatrolAct`

GetSteps returns the Steps field if non-nil, zero value otherwise.

### GetStepsOk

`func (o *PatrolPatrolIncident) GetStepsOk() (*[]PatrolPatrolAct, bool)`

GetStepsOk returns a tuple with the Steps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSteps

`func (o *PatrolPatrolIncident) SetSteps(v []PatrolPatrolAct)`

SetSteps sets Steps field to given value.

### HasSteps

`func (o *PatrolPatrolIncident) HasSteps() bool`

HasSteps returns a boolean if a field has been set.

### GetTier

`func (o *PatrolPatrolIncident) GetTier() string`

GetTier returns the Tier field if non-nil, zero value otherwise.

### GetTierOk

`func (o *PatrolPatrolIncident) GetTierOk() (*string, bool)`

GetTierOk returns a tuple with the Tier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTier

`func (o *PatrolPatrolIncident) SetTier(v string)`

SetTier sets Tier field to given value.

### HasTier

`func (o *PatrolPatrolIncident) HasTier() bool`

HasTier returns a boolean if a field has been set.

### GetTrigger

`func (o *PatrolPatrolIncident) GetTrigger() string`

GetTrigger returns the Trigger field if non-nil, zero value otherwise.

### GetTriggerOk

`func (o *PatrolPatrolIncident) GetTriggerOk() (*string, bool)`

GetTriggerOk returns a tuple with the Trigger field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrigger

`func (o *PatrolPatrolIncident) SetTrigger(v string)`

SetTrigger sets Trigger field to given value.

### HasTrigger

`func (o *PatrolPatrolIncident) HasTrigger() bool`

HasTrigger returns a boolean if a field has been set.

### GetUnit

`func (o *PatrolPatrolIncident) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *PatrolPatrolIncident) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *PatrolPatrolIncident) SetUnit(v string)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *PatrolPatrolIncident) HasUnit() bool`

HasUnit returns a boolean if a field has been set.

### GetVerified

`func (o *PatrolPatrolIncident) GetVerified() string`

GetVerified returns the Verified field if non-nil, zero value otherwise.

### GetVerifiedOk

`func (o *PatrolPatrolIncident) GetVerifiedOk() (*string, bool)`

GetVerifiedOk returns a tuple with the Verified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerified

`func (o *PatrolPatrolIncident) SetVerified(v string)`

SetVerified sets Verified field to given value.

### HasVerified

`func (o *PatrolPatrolIncident) HasVerified() bool`

HasVerified returns a boolean if a field has been set.

### GetZone

`func (o *PatrolPatrolIncident) GetZone() string`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *PatrolPatrolIncident) GetZoneOk() (*string, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *PatrolPatrolIncident) SetZone(v string)`

SetZone sets Zone field to given value.

### HasZone

`func (o *PatrolPatrolIncident) HasZone() bool`

HasZone returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


