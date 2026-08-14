# O11yStatusSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CheckedAt** | Pointer to **string** | CheckedAt is when the underlying availability read was taken, RFC3339 UTC. Not part of the status-page schema the panel parses (which ignores unknown fields); it is here because a status document with no timestamp cannot be told apart from a stale one. | [optional] 
**InProgressMaintenances** | Pointer to [**[]O11yStatusMaintenance**](O11yStatusMaintenance.md) |  | [optional] 
**OngoingIncidents** | Pointer to [**[]O11yStatusIncident**](O11yStatusIncident.md) |  | [optional] 
**PageTitle** | Pointer to **string** |  | [optional] 
**PageUrl** | Pointer to **string** | PageURL is the HUMAN status page — an HTML page for people, distinct from this JSON endpoint. Every link in this document points there. | [optional] 
**ScheduledMaintenances** | Pointer to [**[]O11yStatusMaintenance**](O11yStatusMaintenance.md) |  | [optional] 

## Methods

### NewO11yStatusSummary

`func NewO11yStatusSummary() *O11yStatusSummary`

NewO11yStatusSummary instantiates a new O11yStatusSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yStatusSummaryWithDefaults

`func NewO11yStatusSummaryWithDefaults() *O11yStatusSummary`

NewO11yStatusSummaryWithDefaults instantiates a new O11yStatusSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCheckedAt

`func (o *O11yStatusSummary) GetCheckedAt() string`

GetCheckedAt returns the CheckedAt field if non-nil, zero value otherwise.

### GetCheckedAtOk

`func (o *O11yStatusSummary) GetCheckedAtOk() (*string, bool)`

GetCheckedAtOk returns a tuple with the CheckedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckedAt

`func (o *O11yStatusSummary) SetCheckedAt(v string)`

SetCheckedAt sets CheckedAt field to given value.

### HasCheckedAt

`func (o *O11yStatusSummary) HasCheckedAt() bool`

HasCheckedAt returns a boolean if a field has been set.

### GetInProgressMaintenances

`func (o *O11yStatusSummary) GetInProgressMaintenances() []O11yStatusMaintenance`

GetInProgressMaintenances returns the InProgressMaintenances field if non-nil, zero value otherwise.

### GetInProgressMaintenancesOk

`func (o *O11yStatusSummary) GetInProgressMaintenancesOk() (*[]O11yStatusMaintenance, bool)`

GetInProgressMaintenancesOk returns a tuple with the InProgressMaintenances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInProgressMaintenances

`func (o *O11yStatusSummary) SetInProgressMaintenances(v []O11yStatusMaintenance)`

SetInProgressMaintenances sets InProgressMaintenances field to given value.

### HasInProgressMaintenances

`func (o *O11yStatusSummary) HasInProgressMaintenances() bool`

HasInProgressMaintenances returns a boolean if a field has been set.

### GetOngoingIncidents

`func (o *O11yStatusSummary) GetOngoingIncidents() []O11yStatusIncident`

GetOngoingIncidents returns the OngoingIncidents field if non-nil, zero value otherwise.

### GetOngoingIncidentsOk

`func (o *O11yStatusSummary) GetOngoingIncidentsOk() (*[]O11yStatusIncident, bool)`

GetOngoingIncidentsOk returns a tuple with the OngoingIncidents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOngoingIncidents

`func (o *O11yStatusSummary) SetOngoingIncidents(v []O11yStatusIncident)`

SetOngoingIncidents sets OngoingIncidents field to given value.

### HasOngoingIncidents

`func (o *O11yStatusSummary) HasOngoingIncidents() bool`

HasOngoingIncidents returns a boolean if a field has been set.

### GetPageTitle

`func (o *O11yStatusSummary) GetPageTitle() string`

GetPageTitle returns the PageTitle field if non-nil, zero value otherwise.

### GetPageTitleOk

`func (o *O11yStatusSummary) GetPageTitleOk() (*string, bool)`

GetPageTitleOk returns a tuple with the PageTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageTitle

`func (o *O11yStatusSummary) SetPageTitle(v string)`

SetPageTitle sets PageTitle field to given value.

### HasPageTitle

`func (o *O11yStatusSummary) HasPageTitle() bool`

HasPageTitle returns a boolean if a field has been set.

### GetPageUrl

`func (o *O11yStatusSummary) GetPageUrl() string`

GetPageUrl returns the PageUrl field if non-nil, zero value otherwise.

### GetPageUrlOk

`func (o *O11yStatusSummary) GetPageUrlOk() (*string, bool)`

GetPageUrlOk returns a tuple with the PageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageUrl

`func (o *O11yStatusSummary) SetPageUrl(v string)`

SetPageUrl sets PageUrl field to given value.

### HasPageUrl

`func (o *O11yStatusSummary) HasPageUrl() bool`

HasPageUrl returns a boolean if a field has been set.

### GetScheduledMaintenances

`func (o *O11yStatusSummary) GetScheduledMaintenances() []O11yStatusMaintenance`

GetScheduledMaintenances returns the ScheduledMaintenances field if non-nil, zero value otherwise.

### GetScheduledMaintenancesOk

`func (o *O11yStatusSummary) GetScheduledMaintenancesOk() (*[]O11yStatusMaintenance, bool)`

GetScheduledMaintenancesOk returns a tuple with the ScheduledMaintenances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledMaintenances

`func (o *O11yStatusSummary) SetScheduledMaintenances(v []O11yStatusMaintenance)`

SetScheduledMaintenances sets ScheduledMaintenances field to given value.

### HasScheduledMaintenances

`func (o *O11yStatusSummary) HasScheduledMaintenances() bool`

HasScheduledMaintenances returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


