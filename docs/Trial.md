# Trial

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **string** | when it started assigning | [optional] 
**CreatedBy** | Pointer to **string** | the credential that registered it | [optional] 
**DecidedAt** | Pointer to **string** | when the promotion took effect | [optional] 
**DecidedBy** | Pointer to **string** | the credential that promoted the winner | [optional] 
**ExposureEvent** | Pointer to **string** | the event that enrols a subject — the analysis denominator | [optional] 
**FlagKey** | Pointer to **string** | the assignment flag this experiment drives | [optional] 
**Id** | Pointer to **string** | the experiment&#39;s slug, unique within the project | [optional] 
**MetricEvent** | Pointer to **string** | the event that counts as a conversion — the numerator | [optional] 
**Name** | Pointer to **string** | free text for a reader | [optional] 
**Project** | Pointer to **string** | the sub-scope within the org, stamped from the principal | [optional] 
**Status** | Pointer to **string** | running while it assigns and measures, decided once a winner is promoted | [optional] 
**SubjectKind** | Pointer to **string** | the unit assigned and measured: user, org, session or audience | [optional] 
**Variants** | Pointer to [**[]Arm**](Arm.md) | the arms, weighted, one of them the control | [optional] 
**Winner** | Pointer to **string** | the arm promoted to the whole rollout | [optional] 

## Methods

### NewTrial

`func NewTrial() *Trial`

NewTrial instantiates a new Trial object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrialWithDefaults

`func NewTrialWithDefaults() *Trial`

NewTrialWithDefaults instantiates a new Trial object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *Trial) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Trial) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Trial) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Trial) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCreatedBy

`func (o *Trial) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *Trial) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *Trial) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *Trial) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetDecidedAt

`func (o *Trial) GetDecidedAt() string`

GetDecidedAt returns the DecidedAt field if non-nil, zero value otherwise.

### GetDecidedAtOk

`func (o *Trial) GetDecidedAtOk() (*string, bool)`

GetDecidedAtOk returns a tuple with the DecidedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecidedAt

`func (o *Trial) SetDecidedAt(v string)`

SetDecidedAt sets DecidedAt field to given value.

### HasDecidedAt

`func (o *Trial) HasDecidedAt() bool`

HasDecidedAt returns a boolean if a field has been set.

### GetDecidedBy

`func (o *Trial) GetDecidedBy() string`

GetDecidedBy returns the DecidedBy field if non-nil, zero value otherwise.

### GetDecidedByOk

`func (o *Trial) GetDecidedByOk() (*string, bool)`

GetDecidedByOk returns a tuple with the DecidedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecidedBy

`func (o *Trial) SetDecidedBy(v string)`

SetDecidedBy sets DecidedBy field to given value.

### HasDecidedBy

`func (o *Trial) HasDecidedBy() bool`

HasDecidedBy returns a boolean if a field has been set.

### GetExposureEvent

`func (o *Trial) GetExposureEvent() string`

GetExposureEvent returns the ExposureEvent field if non-nil, zero value otherwise.

### GetExposureEventOk

`func (o *Trial) GetExposureEventOk() (*string, bool)`

GetExposureEventOk returns a tuple with the ExposureEvent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExposureEvent

`func (o *Trial) SetExposureEvent(v string)`

SetExposureEvent sets ExposureEvent field to given value.

### HasExposureEvent

`func (o *Trial) HasExposureEvent() bool`

HasExposureEvent returns a boolean if a field has been set.

### GetFlagKey

`func (o *Trial) GetFlagKey() string`

GetFlagKey returns the FlagKey field if non-nil, zero value otherwise.

### GetFlagKeyOk

`func (o *Trial) GetFlagKeyOk() (*string, bool)`

GetFlagKeyOk returns a tuple with the FlagKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlagKey

`func (o *Trial) SetFlagKey(v string)`

SetFlagKey sets FlagKey field to given value.

### HasFlagKey

`func (o *Trial) HasFlagKey() bool`

HasFlagKey returns a boolean if a field has been set.

### GetId

`func (o *Trial) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Trial) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Trial) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Trial) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMetricEvent

`func (o *Trial) GetMetricEvent() string`

GetMetricEvent returns the MetricEvent field if non-nil, zero value otherwise.

### GetMetricEventOk

`func (o *Trial) GetMetricEventOk() (*string, bool)`

GetMetricEventOk returns a tuple with the MetricEvent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricEvent

`func (o *Trial) SetMetricEvent(v string)`

SetMetricEvent sets MetricEvent field to given value.

### HasMetricEvent

`func (o *Trial) HasMetricEvent() bool`

HasMetricEvent returns a boolean if a field has been set.

### GetName

`func (o *Trial) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Trial) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Trial) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Trial) HasName() bool`

HasName returns a boolean if a field has been set.

### GetProject

`func (o *Trial) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *Trial) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *Trial) SetProject(v string)`

SetProject sets Project field to given value.

### HasProject

`func (o *Trial) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetStatus

`func (o *Trial) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Trial) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Trial) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Trial) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSubjectKind

`func (o *Trial) GetSubjectKind() string`

GetSubjectKind returns the SubjectKind field if non-nil, zero value otherwise.

### GetSubjectKindOk

`func (o *Trial) GetSubjectKindOk() (*string, bool)`

GetSubjectKindOk returns a tuple with the SubjectKind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectKind

`func (o *Trial) SetSubjectKind(v string)`

SetSubjectKind sets SubjectKind field to given value.

### HasSubjectKind

`func (o *Trial) HasSubjectKind() bool`

HasSubjectKind returns a boolean if a field has been set.

### GetVariants

`func (o *Trial) GetVariants() []Arm`

GetVariants returns the Variants field if non-nil, zero value otherwise.

### GetVariantsOk

`func (o *Trial) GetVariantsOk() (*[]Arm, bool)`

GetVariantsOk returns a tuple with the Variants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariants

`func (o *Trial) SetVariants(v []Arm)`

SetVariants sets Variants field to given value.

### HasVariants

`func (o *Trial) HasVariants() bool`

HasVariants returns a boolean if a field has been set.

### GetWinner

`func (o *Trial) GetWinner() string`

GetWinner returns the Winner field if non-nil, zero value otherwise.

### GetWinnerOk

`func (o *Trial) GetWinnerOk() (*string, bool)`

GetWinnerOk returns a tuple with the Winner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWinner

`func (o *Trial) SetWinner(v string)`

SetWinner sets Winner field to given value.

### HasWinner

`func (o *Trial) HasWinner() bool`

HasWinner returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


