# PatrolPatrolSiteEdit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Category** | Pointer to **string** | Category is what kind of premises this is. | [optional] 
**Contact** | Pointer to **string** | Contact is the keyholder to reach. | [optional] 
**Eircode** | Pointer to **string** | Eircode is the postal code. | [optional] 
**Name** | Pointer to **string** | Name is the site&#39;s document name, from the path. | [optional] 
**Notes** | Pointer to **string** | Notes is the standing brief for the site. | [optional] 
**Sector** | Pointer to **string** | Sector is the patrol sector the site sits in. | [optional] 
**Sla** | Pointer to **int64** | SLA is the contracted response time in minutes. | [optional] 
**Status** | Pointer to **string** | Status is ok, fault, alarm or offline. | [optional] 
**Tier** | Pointer to **string** | Tier is the contracted service tier. | [optional] 
**Town** | Pointer to **string** | Town is the town the site is in. | [optional] 

## Methods

### NewPatrolPatrolSiteEdit

`func NewPatrolPatrolSiteEdit() *PatrolPatrolSiteEdit`

NewPatrolPatrolSiteEdit instantiates a new PatrolPatrolSiteEdit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatrolPatrolSiteEditWithDefaults

`func NewPatrolPatrolSiteEditWithDefaults() *PatrolPatrolSiteEdit`

NewPatrolPatrolSiteEditWithDefaults instantiates a new PatrolPatrolSiteEdit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategory

`func (o *PatrolPatrolSiteEdit) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *PatrolPatrolSiteEdit) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *PatrolPatrolSiteEdit) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *PatrolPatrolSiteEdit) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetContact

`func (o *PatrolPatrolSiteEdit) GetContact() string`

GetContact returns the Contact field if non-nil, zero value otherwise.

### GetContactOk

`func (o *PatrolPatrolSiteEdit) GetContactOk() (*string, bool)`

GetContactOk returns a tuple with the Contact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContact

`func (o *PatrolPatrolSiteEdit) SetContact(v string)`

SetContact sets Contact field to given value.

### HasContact

`func (o *PatrolPatrolSiteEdit) HasContact() bool`

HasContact returns a boolean if a field has been set.

### GetEircode

`func (o *PatrolPatrolSiteEdit) GetEircode() string`

GetEircode returns the Eircode field if non-nil, zero value otherwise.

### GetEircodeOk

`func (o *PatrolPatrolSiteEdit) GetEircodeOk() (*string, bool)`

GetEircodeOk returns a tuple with the Eircode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEircode

`func (o *PatrolPatrolSiteEdit) SetEircode(v string)`

SetEircode sets Eircode field to given value.

### HasEircode

`func (o *PatrolPatrolSiteEdit) HasEircode() bool`

HasEircode returns a boolean if a field has been set.

### GetName

`func (o *PatrolPatrolSiteEdit) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatrolPatrolSiteEdit) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatrolPatrolSiteEdit) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatrolPatrolSiteEdit) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNotes

`func (o *PatrolPatrolSiteEdit) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *PatrolPatrolSiteEdit) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *PatrolPatrolSiteEdit) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *PatrolPatrolSiteEdit) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetSector

`func (o *PatrolPatrolSiteEdit) GetSector() string`

GetSector returns the Sector field if non-nil, zero value otherwise.

### GetSectorOk

`func (o *PatrolPatrolSiteEdit) GetSectorOk() (*string, bool)`

GetSectorOk returns a tuple with the Sector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSector

`func (o *PatrolPatrolSiteEdit) SetSector(v string)`

SetSector sets Sector field to given value.

### HasSector

`func (o *PatrolPatrolSiteEdit) HasSector() bool`

HasSector returns a boolean if a field has been set.

### GetSla

`func (o *PatrolPatrolSiteEdit) GetSla() int64`

GetSla returns the Sla field if non-nil, zero value otherwise.

### GetSlaOk

`func (o *PatrolPatrolSiteEdit) GetSlaOk() (*int64, bool)`

GetSlaOk returns a tuple with the Sla field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSla

`func (o *PatrolPatrolSiteEdit) SetSla(v int64)`

SetSla sets Sla field to given value.

### HasSla

`func (o *PatrolPatrolSiteEdit) HasSla() bool`

HasSla returns a boolean if a field has been set.

### GetStatus

`func (o *PatrolPatrolSiteEdit) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PatrolPatrolSiteEdit) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PatrolPatrolSiteEdit) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PatrolPatrolSiteEdit) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTier

`func (o *PatrolPatrolSiteEdit) GetTier() string`

GetTier returns the Tier field if non-nil, zero value otherwise.

### GetTierOk

`func (o *PatrolPatrolSiteEdit) GetTierOk() (*string, bool)`

GetTierOk returns a tuple with the Tier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTier

`func (o *PatrolPatrolSiteEdit) SetTier(v string)`

SetTier sets Tier field to given value.

### HasTier

`func (o *PatrolPatrolSiteEdit) HasTier() bool`

HasTier returns a boolean if a field has been set.

### GetTown

`func (o *PatrolPatrolSiteEdit) GetTown() string`

GetTown returns the Town field if non-nil, zero value otherwise.

### GetTownOk

`func (o *PatrolPatrolSiteEdit) GetTownOk() (*string, bool)`

GetTownOk returns a tuple with the Town field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTown

`func (o *PatrolPatrolSiteEdit) SetTown(v string)`

SetTown sets Town field to given value.

### HasTown

`func (o *PatrolPatrolSiteEdit) HasTown() bool`

HasTown returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


