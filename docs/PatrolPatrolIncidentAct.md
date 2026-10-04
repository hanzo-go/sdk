# PatrolPatrolIncidentAct

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Act** | Pointer to **string** | Act is VERIFY, DISPATCH, ROLL, ARRIVE, CHECK, SECURE, FILE, CLOSE or ESCALATE. | [optional] 
**Detail** | Pointer to **string** | Detail is what to write on the trail with this act. | [optional] 
**False** | Pointer to **bool** | False says the activation turned out to be nothing, read by CLOSE. | [optional] 
**Name** | Pointer to **string** | Name is the incident&#39;s document name, from the path. | [optional] 
**Outcome** | Pointer to **string** | Outcome is what the incident came to, read by CLOSE. | [optional] 
**Unit** | Pointer to **string** | Unit is the call sign to send, read by DISPATCH. | [optional] 

## Methods

### NewPatrolPatrolIncidentAct

`func NewPatrolPatrolIncidentAct() *PatrolPatrolIncidentAct`

NewPatrolPatrolIncidentAct instantiates a new PatrolPatrolIncidentAct object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatrolPatrolIncidentActWithDefaults

`func NewPatrolPatrolIncidentActWithDefaults() *PatrolPatrolIncidentAct`

NewPatrolPatrolIncidentActWithDefaults instantiates a new PatrolPatrolIncidentAct object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAct

`func (o *PatrolPatrolIncidentAct) GetAct() string`

GetAct returns the Act field if non-nil, zero value otherwise.

### GetActOk

`func (o *PatrolPatrolIncidentAct) GetActOk() (*string, bool)`

GetActOk returns a tuple with the Act field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAct

`func (o *PatrolPatrolIncidentAct) SetAct(v string)`

SetAct sets Act field to given value.

### HasAct

`func (o *PatrolPatrolIncidentAct) HasAct() bool`

HasAct returns a boolean if a field has been set.

### GetDetail

`func (o *PatrolPatrolIncidentAct) GetDetail() string`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *PatrolPatrolIncidentAct) GetDetailOk() (*string, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *PatrolPatrolIncidentAct) SetDetail(v string)`

SetDetail sets Detail field to given value.

### HasDetail

`func (o *PatrolPatrolIncidentAct) HasDetail() bool`

HasDetail returns a boolean if a field has been set.

### GetFalse

`func (o *PatrolPatrolIncidentAct) GetFalse() bool`

GetFalse returns the False field if non-nil, zero value otherwise.

### GetFalseOk

`func (o *PatrolPatrolIncidentAct) GetFalseOk() (*bool, bool)`

GetFalseOk returns a tuple with the False field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFalse

`func (o *PatrolPatrolIncidentAct) SetFalse(v bool)`

SetFalse sets False field to given value.

### HasFalse

`func (o *PatrolPatrolIncidentAct) HasFalse() bool`

HasFalse returns a boolean if a field has been set.

### GetName

`func (o *PatrolPatrolIncidentAct) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatrolPatrolIncidentAct) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatrolPatrolIncidentAct) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatrolPatrolIncidentAct) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOutcome

`func (o *PatrolPatrolIncidentAct) GetOutcome() string`

GetOutcome returns the Outcome field if non-nil, zero value otherwise.

### GetOutcomeOk

`func (o *PatrolPatrolIncidentAct) GetOutcomeOk() (*string, bool)`

GetOutcomeOk returns a tuple with the Outcome field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutcome

`func (o *PatrolPatrolIncidentAct) SetOutcome(v string)`

SetOutcome sets Outcome field to given value.

### HasOutcome

`func (o *PatrolPatrolIncidentAct) HasOutcome() bool`

HasOutcome returns a boolean if a field has been set.

### GetUnit

`func (o *PatrolPatrolIncidentAct) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *PatrolPatrolIncidentAct) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *PatrolPatrolIncidentAct) SetUnit(v string)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *PatrolPatrolIncidentAct) HasUnit() bool`

HasUnit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


