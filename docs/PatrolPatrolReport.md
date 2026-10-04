# PatrolPatrolReport

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**At** | Pointer to **string** | At is when it was filed. | [optional] 
**Body** | Pointer to **string** | Body is the report as text. | [optional] 
**Client** | Pointer to **string** | Client is the customer it is about, or empty. | [optional] 
**From** | Pointer to **string** | From is the start of the window the report covers. | [optional] 
**Incident** | Pointer to **string** | Incident is the incident it is about, or empty. | [optional] 
**Kind** | Pointer to **string** | Kind is shift, incident or client. | [optional] 
**Name** | Pointer to **string** | Name is the report&#39;s document name and the segment /v1/patrol/report/{name} addresses it by. | [optional] 
**To** | Pointer to **string** | To is the end of that window. | [optional] 

## Methods

### NewPatrolPatrolReport

`func NewPatrolPatrolReport() *PatrolPatrolReport`

NewPatrolPatrolReport instantiates a new PatrolPatrolReport object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatrolPatrolReportWithDefaults

`func NewPatrolPatrolReportWithDefaults() *PatrolPatrolReport`

NewPatrolPatrolReportWithDefaults instantiates a new PatrolPatrolReport object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAt

`func (o *PatrolPatrolReport) GetAt() string`

GetAt returns the At field if non-nil, zero value otherwise.

### GetAtOk

`func (o *PatrolPatrolReport) GetAtOk() (*string, bool)`

GetAtOk returns a tuple with the At field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAt

`func (o *PatrolPatrolReport) SetAt(v string)`

SetAt sets At field to given value.

### HasAt

`func (o *PatrolPatrolReport) HasAt() bool`

HasAt returns a boolean if a field has been set.

### GetBody

`func (o *PatrolPatrolReport) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *PatrolPatrolReport) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *PatrolPatrolReport) SetBody(v string)`

SetBody sets Body field to given value.

### HasBody

`func (o *PatrolPatrolReport) HasBody() bool`

HasBody returns a boolean if a field has been set.

### GetClient

`func (o *PatrolPatrolReport) GetClient() string`

GetClient returns the Client field if non-nil, zero value otherwise.

### GetClientOk

`func (o *PatrolPatrolReport) GetClientOk() (*string, bool)`

GetClientOk returns a tuple with the Client field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClient

`func (o *PatrolPatrolReport) SetClient(v string)`

SetClient sets Client field to given value.

### HasClient

`func (o *PatrolPatrolReport) HasClient() bool`

HasClient returns a boolean if a field has been set.

### GetFrom

`func (o *PatrolPatrolReport) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *PatrolPatrolReport) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *PatrolPatrolReport) SetFrom(v string)`

SetFrom sets From field to given value.

### HasFrom

`func (o *PatrolPatrolReport) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetIncident

`func (o *PatrolPatrolReport) GetIncident() string`

GetIncident returns the Incident field if non-nil, zero value otherwise.

### GetIncidentOk

`func (o *PatrolPatrolReport) GetIncidentOk() (*string, bool)`

GetIncidentOk returns a tuple with the Incident field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncident

`func (o *PatrolPatrolReport) SetIncident(v string)`

SetIncident sets Incident field to given value.

### HasIncident

`func (o *PatrolPatrolReport) HasIncident() bool`

HasIncident returns a boolean if a field has been set.

### GetKind

`func (o *PatrolPatrolReport) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *PatrolPatrolReport) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *PatrolPatrolReport) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *PatrolPatrolReport) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetName

`func (o *PatrolPatrolReport) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatrolPatrolReport) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatrolPatrolReport) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatrolPatrolReport) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTo

`func (o *PatrolPatrolReport) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *PatrolPatrolReport) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *PatrolPatrolReport) SetTo(v string)`

SetTo sets To field to given value.

### HasTo

`func (o *PatrolPatrolReport) HasTo() bool`

HasTo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


