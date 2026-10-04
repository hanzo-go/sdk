# PatrolPatrolActivation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Alarm** | Pointer to **string** | Alarm is the recorded activation&#39;s document name. | [optional] 
**Incident** | Pointer to [**PatrolPatrolIncident**](PatrolPatrolIncident.md) | Incident is the incident it opened, or null when it matched no site. | [optional] 
**Matched** | Pointer to **bool** | Matched says whether a site was found for it. | [optional] 
**Site** | Pointer to **string** | Site is the site it matched, or empty. | [optional] 

## Methods

### NewPatrolPatrolActivation

`func NewPatrolPatrolActivation() *PatrolPatrolActivation`

NewPatrolPatrolActivation instantiates a new PatrolPatrolActivation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatrolPatrolActivationWithDefaults

`func NewPatrolPatrolActivationWithDefaults() *PatrolPatrolActivation`

NewPatrolPatrolActivationWithDefaults instantiates a new PatrolPatrolActivation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlarm

`func (o *PatrolPatrolActivation) GetAlarm() string`

GetAlarm returns the Alarm field if non-nil, zero value otherwise.

### GetAlarmOk

`func (o *PatrolPatrolActivation) GetAlarmOk() (*string, bool)`

GetAlarmOk returns a tuple with the Alarm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlarm

`func (o *PatrolPatrolActivation) SetAlarm(v string)`

SetAlarm sets Alarm field to given value.

### HasAlarm

`func (o *PatrolPatrolActivation) HasAlarm() bool`

HasAlarm returns a boolean if a field has been set.

### GetIncident

`func (o *PatrolPatrolActivation) GetIncident() PatrolPatrolIncident`

GetIncident returns the Incident field if non-nil, zero value otherwise.

### GetIncidentOk

`func (o *PatrolPatrolActivation) GetIncidentOk() (*PatrolPatrolIncident, bool)`

GetIncidentOk returns a tuple with the Incident field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncident

`func (o *PatrolPatrolActivation) SetIncident(v PatrolPatrolIncident)`

SetIncident sets Incident field to given value.

### HasIncident

`func (o *PatrolPatrolActivation) HasIncident() bool`

HasIncident returns a boolean if a field has been set.

### GetMatched

`func (o *PatrolPatrolActivation) GetMatched() bool`

GetMatched returns the Matched field if non-nil, zero value otherwise.

### GetMatchedOk

`func (o *PatrolPatrolActivation) GetMatchedOk() (*bool, bool)`

GetMatchedOk returns a tuple with the Matched field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatched

`func (o *PatrolPatrolActivation) SetMatched(v bool)`

SetMatched sets Matched field to given value.

### HasMatched

`func (o *PatrolPatrolActivation) HasMatched() bool`

HasMatched returns a boolean if a field has been set.

### GetSite

`func (o *PatrolPatrolActivation) GetSite() string`

GetSite returns the Site field if non-nil, zero value otherwise.

### GetSiteOk

`func (o *PatrolPatrolActivation) GetSiteOk() (*string, bool)`

GetSiteOk returns a tuple with the Site field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSite

`func (o *PatrolPatrolActivation) SetSite(v string)`

SetSite sets Site field to given value.

### HasSite

`func (o *PatrolPatrolActivation) HasSite() bool`

HasSite returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


