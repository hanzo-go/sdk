# PatrolPatrolEstate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Events** | Pointer to [**[]PatrolPatrolEvent**](PatrolPatrolEvent.md) | Events is the recent feed, newest first. | [optional] 
**Incident** | Pointer to [**PatrolPatrolIncident**](PatrolPatrolIncident.md) | Incident is the incident currently running, or null when none is. | [optional] 
**Keys** | Pointer to [**[]PatrolPatrolKey**](PatrolPatrolKey.md) | Keys is the key register, empty for a caller whose roles do not reach it. | [optional] 
**Sites** | Pointer to [**[]PatrolPatrolSite**](PatrolPatrolSite.md) | Sites is the estate under contract. | [optional] 
**Tours** | Pointer to [**[]PatrolPatrolTour**](PatrolPatrolTour.md) | Tours is today&#39;s rounds with their checkpoints. | [optional] 
**Units** | Pointer to [**[]PatrolPatrolUnit**](PatrolPatrolUnit.md) | Units is the fleet, each with the tail of its position trail. | [optional] 

## Methods

### NewPatrolPatrolEstate

`func NewPatrolPatrolEstate() *PatrolPatrolEstate`

NewPatrolPatrolEstate instantiates a new PatrolPatrolEstate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatrolPatrolEstateWithDefaults

`func NewPatrolPatrolEstateWithDefaults() *PatrolPatrolEstate`

NewPatrolPatrolEstateWithDefaults instantiates a new PatrolPatrolEstate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEvents

`func (o *PatrolPatrolEstate) GetEvents() []PatrolPatrolEvent`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *PatrolPatrolEstate) GetEventsOk() (*[]PatrolPatrolEvent, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *PatrolPatrolEstate) SetEvents(v []PatrolPatrolEvent)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *PatrolPatrolEstate) HasEvents() bool`

HasEvents returns a boolean if a field has been set.

### GetIncident

`func (o *PatrolPatrolEstate) GetIncident() PatrolPatrolIncident`

GetIncident returns the Incident field if non-nil, zero value otherwise.

### GetIncidentOk

`func (o *PatrolPatrolEstate) GetIncidentOk() (*PatrolPatrolIncident, bool)`

GetIncidentOk returns a tuple with the Incident field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncident

`func (o *PatrolPatrolEstate) SetIncident(v PatrolPatrolIncident)`

SetIncident sets Incident field to given value.

### HasIncident

`func (o *PatrolPatrolEstate) HasIncident() bool`

HasIncident returns a boolean if a field has been set.

### GetKeys

`func (o *PatrolPatrolEstate) GetKeys() []PatrolPatrolKey`

GetKeys returns the Keys field if non-nil, zero value otherwise.

### GetKeysOk

`func (o *PatrolPatrolEstate) GetKeysOk() (*[]PatrolPatrolKey, bool)`

GetKeysOk returns a tuple with the Keys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeys

`func (o *PatrolPatrolEstate) SetKeys(v []PatrolPatrolKey)`

SetKeys sets Keys field to given value.

### HasKeys

`func (o *PatrolPatrolEstate) HasKeys() bool`

HasKeys returns a boolean if a field has been set.

### GetSites

`func (o *PatrolPatrolEstate) GetSites() []PatrolPatrolSite`

GetSites returns the Sites field if non-nil, zero value otherwise.

### GetSitesOk

`func (o *PatrolPatrolEstate) GetSitesOk() (*[]PatrolPatrolSite, bool)`

GetSitesOk returns a tuple with the Sites field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSites

`func (o *PatrolPatrolEstate) SetSites(v []PatrolPatrolSite)`

SetSites sets Sites field to given value.

### HasSites

`func (o *PatrolPatrolEstate) HasSites() bool`

HasSites returns a boolean if a field has been set.

### GetTours

`func (o *PatrolPatrolEstate) GetTours() []PatrolPatrolTour`

GetTours returns the Tours field if non-nil, zero value otherwise.

### GetToursOk

`func (o *PatrolPatrolEstate) GetToursOk() (*[]PatrolPatrolTour, bool)`

GetToursOk returns a tuple with the Tours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTours

`func (o *PatrolPatrolEstate) SetTours(v []PatrolPatrolTour)`

SetTours sets Tours field to given value.

### HasTours

`func (o *PatrolPatrolEstate) HasTours() bool`

HasTours returns a boolean if a field has been set.

### GetUnits

`func (o *PatrolPatrolEstate) GetUnits() []PatrolPatrolUnit`

GetUnits returns the Units field if non-nil, zero value otherwise.

### GetUnitsOk

`func (o *PatrolPatrolEstate) GetUnitsOk() (*[]PatrolPatrolUnit, bool)`

GetUnitsOk returns a tuple with the Units field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnits

`func (o *PatrolPatrolEstate) SetUnits(v []PatrolPatrolUnit)`

SetUnits sets Units field to given value.

### HasUnits

`func (o *PatrolPatrolEstate) HasUnits() bool`

HasUnits returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


