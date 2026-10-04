# PatrolPatrolReportIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Client** | Pointer to **string** | Client is the customer to report on, read by kind client. Naming one files an account that customer may download; leaving it empty covers every customer, which is an account of the operation and stays with it. | [optional] 
**From** | Pointer to **string** | From is the start of the window to cover. | [optional] 
**Incident** | Pointer to **string** | Incident is the incident to report on, required for kind incident. | [optional] 
**Kind** | Pointer to **string** | Kind is shift, incident or client. | [optional] 
**To** | Pointer to **string** | To is the end of that window. | [optional] 

## Methods

### NewPatrolPatrolReportIn

`func NewPatrolPatrolReportIn() *PatrolPatrolReportIn`

NewPatrolPatrolReportIn instantiates a new PatrolPatrolReportIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatrolPatrolReportInWithDefaults

`func NewPatrolPatrolReportInWithDefaults() *PatrolPatrolReportIn`

NewPatrolPatrolReportInWithDefaults instantiates a new PatrolPatrolReportIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClient

`func (o *PatrolPatrolReportIn) GetClient() string`

GetClient returns the Client field if non-nil, zero value otherwise.

### GetClientOk

`func (o *PatrolPatrolReportIn) GetClientOk() (*string, bool)`

GetClientOk returns a tuple with the Client field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClient

`func (o *PatrolPatrolReportIn) SetClient(v string)`

SetClient sets Client field to given value.

### HasClient

`func (o *PatrolPatrolReportIn) HasClient() bool`

HasClient returns a boolean if a field has been set.

### GetFrom

`func (o *PatrolPatrolReportIn) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *PatrolPatrolReportIn) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *PatrolPatrolReportIn) SetFrom(v string)`

SetFrom sets From field to given value.

### HasFrom

`func (o *PatrolPatrolReportIn) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetIncident

`func (o *PatrolPatrolReportIn) GetIncident() string`

GetIncident returns the Incident field if non-nil, zero value otherwise.

### GetIncidentOk

`func (o *PatrolPatrolReportIn) GetIncidentOk() (*string, bool)`

GetIncidentOk returns a tuple with the Incident field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncident

`func (o *PatrolPatrolReportIn) SetIncident(v string)`

SetIncident sets Incident field to given value.

### HasIncident

`func (o *PatrolPatrolReportIn) HasIncident() bool`

HasIncident returns a boolean if a field has been set.

### GetKind

`func (o *PatrolPatrolReportIn) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *PatrolPatrolReportIn) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *PatrolPatrolReportIn) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *PatrolPatrolReportIn) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetTo

`func (o *PatrolPatrolReportIn) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *PatrolPatrolReportIn) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *PatrolPatrolReportIn) SetTo(v string)`

SetTo sets To field to given value.

### HasTo

`func (o *PatrolPatrolReportIn) HasTo() bool`

HasTo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


