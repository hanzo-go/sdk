# PatrolPatrolSite

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Account** | Pointer to **string** | Account is the alarm-panel account number the receiving centre reports. | [optional] 
**Cameras** | Pointer to **int64** | Cameras is how many cameras are registered here. | [optional] 
**Category** | Pointer to **string** | Category is what kind of premises this is. | [optional] 
**Client** | Pointer to **string** | Client is the customer this site belongs to. | [optional] 
**Contact** | Pointer to **string** | Contact is the keyholder to reach. | [optional] 
**Eircode** | Pointer to **string** | Eircode is the postal code. | [optional] 
**Hazards** | Pointer to **[]string** | Hazards is the standing warnings, present on a single-site read. | [optional] 
**Keys** | Pointer to **bool** | Keys says whether the operation holds a key set for this site. | [optional] 
**Lat** | Pointer to **float64** | Lat is the site&#39;s latitude in degrees. | [optional] 
**Lon** | Pointer to **float64** | Lon is the site&#39;s longitude in degrees. | [optional] 
**Name** | Pointer to **string** | Name is the site&#39;s document name and the segment /v1/patrol/site/{name} addresses it by. | [optional] 
**Notes** | Pointer to **string** | Notes is the standing brief for the site. | [optional] 
**Ref** | Pointer to **int64** | Ref is the site&#39;s number in the register. | [optional] 
**Risk** | Pointer to **int64** | Risk is the assessed risk score. | [optional] 
**Sector** | Pointer to **string** | Sector is the patrol sector the site sits in. | [optional] 
**Sla** | Pointer to **int64** | SLA is the contracted response time in minutes. | [optional] 
**Status** | Pointer to **string** | Status is ok, fault, alarm or offline. | [optional] 
**Tier** | Pointer to **string** | Tier is the contracted service tier. | [optional] 
**Town** | Pointer to **string** | Town is the town the site is in. | [optional] 
**Unit** | Pointer to **string** | Unit is the call sign of the unit that covers the site, or empty. | [optional] 
**Visits** | Pointer to **int64** | Visits is how many patrol visits the contract calls for. | [optional] 
**Zones** | Pointer to [**[]PatrolPatrolZone**](PatrolPatrolZone.md) | Zones is the panel&#39;s detection zones, present on a single-site read. | [optional] 

## Methods

### NewPatrolPatrolSite

`func NewPatrolPatrolSite() *PatrolPatrolSite`

NewPatrolPatrolSite instantiates a new PatrolPatrolSite object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatrolPatrolSiteWithDefaults

`func NewPatrolPatrolSiteWithDefaults() *PatrolPatrolSite`

NewPatrolPatrolSiteWithDefaults instantiates a new PatrolPatrolSite object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccount

`func (o *PatrolPatrolSite) GetAccount() string`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *PatrolPatrolSite) GetAccountOk() (*string, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *PatrolPatrolSite) SetAccount(v string)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *PatrolPatrolSite) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetCameras

`func (o *PatrolPatrolSite) GetCameras() int64`

GetCameras returns the Cameras field if non-nil, zero value otherwise.

### GetCamerasOk

`func (o *PatrolPatrolSite) GetCamerasOk() (*int64, bool)`

GetCamerasOk returns a tuple with the Cameras field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCameras

`func (o *PatrolPatrolSite) SetCameras(v int64)`

SetCameras sets Cameras field to given value.

### HasCameras

`func (o *PatrolPatrolSite) HasCameras() bool`

HasCameras returns a boolean if a field has been set.

### GetCategory

`func (o *PatrolPatrolSite) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *PatrolPatrolSite) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *PatrolPatrolSite) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *PatrolPatrolSite) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetClient

`func (o *PatrolPatrolSite) GetClient() string`

GetClient returns the Client field if non-nil, zero value otherwise.

### GetClientOk

`func (o *PatrolPatrolSite) GetClientOk() (*string, bool)`

GetClientOk returns a tuple with the Client field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClient

`func (o *PatrolPatrolSite) SetClient(v string)`

SetClient sets Client field to given value.

### HasClient

`func (o *PatrolPatrolSite) HasClient() bool`

HasClient returns a boolean if a field has been set.

### GetContact

`func (o *PatrolPatrolSite) GetContact() string`

GetContact returns the Contact field if non-nil, zero value otherwise.

### GetContactOk

`func (o *PatrolPatrolSite) GetContactOk() (*string, bool)`

GetContactOk returns a tuple with the Contact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContact

`func (o *PatrolPatrolSite) SetContact(v string)`

SetContact sets Contact field to given value.

### HasContact

`func (o *PatrolPatrolSite) HasContact() bool`

HasContact returns a boolean if a field has been set.

### GetEircode

`func (o *PatrolPatrolSite) GetEircode() string`

GetEircode returns the Eircode field if non-nil, zero value otherwise.

### GetEircodeOk

`func (o *PatrolPatrolSite) GetEircodeOk() (*string, bool)`

GetEircodeOk returns a tuple with the Eircode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEircode

`func (o *PatrolPatrolSite) SetEircode(v string)`

SetEircode sets Eircode field to given value.

### HasEircode

`func (o *PatrolPatrolSite) HasEircode() bool`

HasEircode returns a boolean if a field has been set.

### GetHazards

`func (o *PatrolPatrolSite) GetHazards() []string`

GetHazards returns the Hazards field if non-nil, zero value otherwise.

### GetHazardsOk

`func (o *PatrolPatrolSite) GetHazardsOk() (*[]string, bool)`

GetHazardsOk returns a tuple with the Hazards field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHazards

`func (o *PatrolPatrolSite) SetHazards(v []string)`

SetHazards sets Hazards field to given value.

### HasHazards

`func (o *PatrolPatrolSite) HasHazards() bool`

HasHazards returns a boolean if a field has been set.

### GetKeys

`func (o *PatrolPatrolSite) GetKeys() bool`

GetKeys returns the Keys field if non-nil, zero value otherwise.

### GetKeysOk

`func (o *PatrolPatrolSite) GetKeysOk() (*bool, bool)`

GetKeysOk returns a tuple with the Keys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeys

`func (o *PatrolPatrolSite) SetKeys(v bool)`

SetKeys sets Keys field to given value.

### HasKeys

`func (o *PatrolPatrolSite) HasKeys() bool`

HasKeys returns a boolean if a field has been set.

### GetLat

`func (o *PatrolPatrolSite) GetLat() float64`

GetLat returns the Lat field if non-nil, zero value otherwise.

### GetLatOk

`func (o *PatrolPatrolSite) GetLatOk() (*float64, bool)`

GetLatOk returns a tuple with the Lat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLat

`func (o *PatrolPatrolSite) SetLat(v float64)`

SetLat sets Lat field to given value.

### HasLat

`func (o *PatrolPatrolSite) HasLat() bool`

HasLat returns a boolean if a field has been set.

### GetLon

`func (o *PatrolPatrolSite) GetLon() float64`

GetLon returns the Lon field if non-nil, zero value otherwise.

### GetLonOk

`func (o *PatrolPatrolSite) GetLonOk() (*float64, bool)`

GetLonOk returns a tuple with the Lon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLon

`func (o *PatrolPatrolSite) SetLon(v float64)`

SetLon sets Lon field to given value.

### HasLon

`func (o *PatrolPatrolSite) HasLon() bool`

HasLon returns a boolean if a field has been set.

### GetName

`func (o *PatrolPatrolSite) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatrolPatrolSite) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatrolPatrolSite) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatrolPatrolSite) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNotes

`func (o *PatrolPatrolSite) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *PatrolPatrolSite) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *PatrolPatrolSite) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *PatrolPatrolSite) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetRef

`func (o *PatrolPatrolSite) GetRef() int64`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *PatrolPatrolSite) GetRefOk() (*int64, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *PatrolPatrolSite) SetRef(v int64)`

SetRef sets Ref field to given value.

### HasRef

`func (o *PatrolPatrolSite) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetRisk

`func (o *PatrolPatrolSite) GetRisk() int64`

GetRisk returns the Risk field if non-nil, zero value otherwise.

### GetRiskOk

`func (o *PatrolPatrolSite) GetRiskOk() (*int64, bool)`

GetRiskOk returns a tuple with the Risk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRisk

`func (o *PatrolPatrolSite) SetRisk(v int64)`

SetRisk sets Risk field to given value.

### HasRisk

`func (o *PatrolPatrolSite) HasRisk() bool`

HasRisk returns a boolean if a field has been set.

### GetSector

`func (o *PatrolPatrolSite) GetSector() string`

GetSector returns the Sector field if non-nil, zero value otherwise.

### GetSectorOk

`func (o *PatrolPatrolSite) GetSectorOk() (*string, bool)`

GetSectorOk returns a tuple with the Sector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSector

`func (o *PatrolPatrolSite) SetSector(v string)`

SetSector sets Sector field to given value.

### HasSector

`func (o *PatrolPatrolSite) HasSector() bool`

HasSector returns a boolean if a field has been set.

### GetSla

`func (o *PatrolPatrolSite) GetSla() int64`

GetSla returns the Sla field if non-nil, zero value otherwise.

### GetSlaOk

`func (o *PatrolPatrolSite) GetSlaOk() (*int64, bool)`

GetSlaOk returns a tuple with the Sla field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSla

`func (o *PatrolPatrolSite) SetSla(v int64)`

SetSla sets Sla field to given value.

### HasSla

`func (o *PatrolPatrolSite) HasSla() bool`

HasSla returns a boolean if a field has been set.

### GetStatus

`func (o *PatrolPatrolSite) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PatrolPatrolSite) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PatrolPatrolSite) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PatrolPatrolSite) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTier

`func (o *PatrolPatrolSite) GetTier() string`

GetTier returns the Tier field if non-nil, zero value otherwise.

### GetTierOk

`func (o *PatrolPatrolSite) GetTierOk() (*string, bool)`

GetTierOk returns a tuple with the Tier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTier

`func (o *PatrolPatrolSite) SetTier(v string)`

SetTier sets Tier field to given value.

### HasTier

`func (o *PatrolPatrolSite) HasTier() bool`

HasTier returns a boolean if a field has been set.

### GetTown

`func (o *PatrolPatrolSite) GetTown() string`

GetTown returns the Town field if non-nil, zero value otherwise.

### GetTownOk

`func (o *PatrolPatrolSite) GetTownOk() (*string, bool)`

GetTownOk returns a tuple with the Town field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTown

`func (o *PatrolPatrolSite) SetTown(v string)`

SetTown sets Town field to given value.

### HasTown

`func (o *PatrolPatrolSite) HasTown() bool`

HasTown returns a boolean if a field has been set.

### GetUnit

`func (o *PatrolPatrolSite) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *PatrolPatrolSite) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *PatrolPatrolSite) SetUnit(v string)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *PatrolPatrolSite) HasUnit() bool`

HasUnit returns a boolean if a field has been set.

### GetVisits

`func (o *PatrolPatrolSite) GetVisits() int64`

GetVisits returns the Visits field if non-nil, zero value otherwise.

### GetVisitsOk

`func (o *PatrolPatrolSite) GetVisitsOk() (*int64, bool)`

GetVisitsOk returns a tuple with the Visits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisits

`func (o *PatrolPatrolSite) SetVisits(v int64)`

SetVisits sets Visits field to given value.

### HasVisits

`func (o *PatrolPatrolSite) HasVisits() bool`

HasVisits returns a boolean if a field has been set.

### GetZones

`func (o *PatrolPatrolSite) GetZones() []PatrolPatrolZone`

GetZones returns the Zones field if non-nil, zero value otherwise.

### GetZonesOk

`func (o *PatrolPatrolSite) GetZonesOk() (*[]PatrolPatrolZone, bool)`

GetZonesOk returns a tuple with the Zones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZones

`func (o *PatrolPatrolSite) SetZones(v []PatrolPatrolZone)`

SetZones sets Zones field to given value.

### HasZones

`func (o *PatrolPatrolSite) HasZones() bool`

HasZones returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


