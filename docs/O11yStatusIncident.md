# O11yStatusIncident

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AffectedComponents** | Pointer to [**[]O11yStatusComponent**](O11yStatusComponent.md) | AffectedComponents is what this incident covers. It is COUNTED rather than classified: some services down is a partial outage and every probed service down is a full one, because deciding that one service is critical and another is not would need a judgement nobody has measured. | [optional] 
**CurrentWorstImpact** | Pointer to **string** | CurrentWorstImpact is the incident&#39;s impact on the PLATFORM, which is not the same question as the component&#39;s own condition above. | [optional] 
**Id** | Pointer to **string** | ID is derived from the service, so the same outage keeps one id across reads rather than being reported as a new incident every 15 seconds. | [optional] 
**LastUpdateAt** | Pointer to **string** | LastUpdateAt is when the failing measurement this incident reports was read, RFC3339 UTC. | [optional] 
**LastUpdateMessage** | Pointer to **string** | LastUpdateMessage says what was observed, not what is being done about it — there is no operator writing updates here, only the probe that failed. | [optional] 
**Name** | Pointer to **string** | Name is the one-line headline, built from the service that stopped answering. | [optional] 
**Status** | Pointer to **string** | Status is always \&quot;investigating\&quot; — the member of the client&#39;s closed set that means detected, cause not yet established, which is exactly what an automated prober knows. Nothing here ever claims \&quot;identified\&quot;: that would assert a diagnosis no measurement made. | [optional] 
**Url** | Pointer to **string** | URL points at the HUMAN status page, not back at this JSON. Every link in this document goes to the same place. | [optional] 

## Methods

### NewO11yStatusIncident

`func NewO11yStatusIncident() *O11yStatusIncident`

NewO11yStatusIncident instantiates a new O11yStatusIncident object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yStatusIncidentWithDefaults

`func NewO11yStatusIncidentWithDefaults() *O11yStatusIncident`

NewO11yStatusIncidentWithDefaults instantiates a new O11yStatusIncident object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAffectedComponents

`func (o *O11yStatusIncident) GetAffectedComponents() []O11yStatusComponent`

GetAffectedComponents returns the AffectedComponents field if non-nil, zero value otherwise.

### GetAffectedComponentsOk

`func (o *O11yStatusIncident) GetAffectedComponentsOk() (*[]O11yStatusComponent, bool)`

GetAffectedComponentsOk returns a tuple with the AffectedComponents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffectedComponents

`func (o *O11yStatusIncident) SetAffectedComponents(v []O11yStatusComponent)`

SetAffectedComponents sets AffectedComponents field to given value.

### HasAffectedComponents

`func (o *O11yStatusIncident) HasAffectedComponents() bool`

HasAffectedComponents returns a boolean if a field has been set.

### GetCurrentWorstImpact

`func (o *O11yStatusIncident) GetCurrentWorstImpact() string`

GetCurrentWorstImpact returns the CurrentWorstImpact field if non-nil, zero value otherwise.

### GetCurrentWorstImpactOk

`func (o *O11yStatusIncident) GetCurrentWorstImpactOk() (*string, bool)`

GetCurrentWorstImpactOk returns a tuple with the CurrentWorstImpact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentWorstImpact

`func (o *O11yStatusIncident) SetCurrentWorstImpact(v string)`

SetCurrentWorstImpact sets CurrentWorstImpact field to given value.

### HasCurrentWorstImpact

`func (o *O11yStatusIncident) HasCurrentWorstImpact() bool`

HasCurrentWorstImpact returns a boolean if a field has been set.

### GetId

`func (o *O11yStatusIncident) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *O11yStatusIncident) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *O11yStatusIncident) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *O11yStatusIncident) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastUpdateAt

`func (o *O11yStatusIncident) GetLastUpdateAt() string`

GetLastUpdateAt returns the LastUpdateAt field if non-nil, zero value otherwise.

### GetLastUpdateAtOk

`func (o *O11yStatusIncident) GetLastUpdateAtOk() (*string, bool)`

GetLastUpdateAtOk returns a tuple with the LastUpdateAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdateAt

`func (o *O11yStatusIncident) SetLastUpdateAt(v string)`

SetLastUpdateAt sets LastUpdateAt field to given value.

### HasLastUpdateAt

`func (o *O11yStatusIncident) HasLastUpdateAt() bool`

HasLastUpdateAt returns a boolean if a field has been set.

### GetLastUpdateMessage

`func (o *O11yStatusIncident) GetLastUpdateMessage() string`

GetLastUpdateMessage returns the LastUpdateMessage field if non-nil, zero value otherwise.

### GetLastUpdateMessageOk

`func (o *O11yStatusIncident) GetLastUpdateMessageOk() (*string, bool)`

GetLastUpdateMessageOk returns a tuple with the LastUpdateMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdateMessage

`func (o *O11yStatusIncident) SetLastUpdateMessage(v string)`

SetLastUpdateMessage sets LastUpdateMessage field to given value.

### HasLastUpdateMessage

`func (o *O11yStatusIncident) HasLastUpdateMessage() bool`

HasLastUpdateMessage returns a boolean if a field has been set.

### GetName

`func (o *O11yStatusIncident) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *O11yStatusIncident) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *O11yStatusIncident) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *O11yStatusIncident) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStatus

`func (o *O11yStatusIncident) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *O11yStatusIncident) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *O11yStatusIncident) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *O11yStatusIncident) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUrl

`func (o *O11yStatusIncident) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *O11yStatusIncident) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *O11yStatusIncident) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *O11yStatusIncident) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


