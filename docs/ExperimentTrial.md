# ExperimentTrial

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
**Variants** | Pointer to [**[]ExperimentArm**](ExperimentArm.md) | the arms, weighted, one of them the control | [optional] 
**Winner** | Pointer to **string** | the arm promoted to the whole rollout | [optional] 

## Methods

### NewExperimentTrial

`func NewExperimentTrial() *ExperimentTrial`

NewExperimentTrial instantiates a new ExperimentTrial object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExperimentTrialWithDefaults

`func NewExperimentTrialWithDefaults() *ExperimentTrial`

NewExperimentTrialWithDefaults instantiates a new ExperimentTrial object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *ExperimentTrial) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ExperimentTrial) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ExperimentTrial) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ExperimentTrial) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCreatedBy

`func (o *ExperimentTrial) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *ExperimentTrial) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *ExperimentTrial) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *ExperimentTrial) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetDecidedAt

`func (o *ExperimentTrial) GetDecidedAt() string`

GetDecidedAt returns the DecidedAt field if non-nil, zero value otherwise.

### GetDecidedAtOk

`func (o *ExperimentTrial) GetDecidedAtOk() (*string, bool)`

GetDecidedAtOk returns a tuple with the DecidedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecidedAt

`func (o *ExperimentTrial) SetDecidedAt(v string)`

SetDecidedAt sets DecidedAt field to given value.

### HasDecidedAt

`func (o *ExperimentTrial) HasDecidedAt() bool`

HasDecidedAt returns a boolean if a field has been set.

### GetDecidedBy

`func (o *ExperimentTrial) GetDecidedBy() string`

GetDecidedBy returns the DecidedBy field if non-nil, zero value otherwise.

### GetDecidedByOk

`func (o *ExperimentTrial) GetDecidedByOk() (*string, bool)`

GetDecidedByOk returns a tuple with the DecidedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecidedBy

`func (o *ExperimentTrial) SetDecidedBy(v string)`

SetDecidedBy sets DecidedBy field to given value.

### HasDecidedBy

`func (o *ExperimentTrial) HasDecidedBy() bool`

HasDecidedBy returns a boolean if a field has been set.

### GetExposureEvent

`func (o *ExperimentTrial) GetExposureEvent() string`

GetExposureEvent returns the ExposureEvent field if non-nil, zero value otherwise.

### GetExposureEventOk

`func (o *ExperimentTrial) GetExposureEventOk() (*string, bool)`

GetExposureEventOk returns a tuple with the ExposureEvent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExposureEvent

`func (o *ExperimentTrial) SetExposureEvent(v string)`

SetExposureEvent sets ExposureEvent field to given value.

### HasExposureEvent

`func (o *ExperimentTrial) HasExposureEvent() bool`

HasExposureEvent returns a boolean if a field has been set.

### GetFlagKey

`func (o *ExperimentTrial) GetFlagKey() string`

GetFlagKey returns the FlagKey field if non-nil, zero value otherwise.

### GetFlagKeyOk

`func (o *ExperimentTrial) GetFlagKeyOk() (*string, bool)`

GetFlagKeyOk returns a tuple with the FlagKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlagKey

`func (o *ExperimentTrial) SetFlagKey(v string)`

SetFlagKey sets FlagKey field to given value.

### HasFlagKey

`func (o *ExperimentTrial) HasFlagKey() bool`

HasFlagKey returns a boolean if a field has been set.

### GetId

`func (o *ExperimentTrial) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ExperimentTrial) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ExperimentTrial) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ExperimentTrial) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMetricEvent

`func (o *ExperimentTrial) GetMetricEvent() string`

GetMetricEvent returns the MetricEvent field if non-nil, zero value otherwise.

### GetMetricEventOk

`func (o *ExperimentTrial) GetMetricEventOk() (*string, bool)`

GetMetricEventOk returns a tuple with the MetricEvent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricEvent

`func (o *ExperimentTrial) SetMetricEvent(v string)`

SetMetricEvent sets MetricEvent field to given value.

### HasMetricEvent

`func (o *ExperimentTrial) HasMetricEvent() bool`

HasMetricEvent returns a boolean if a field has been set.

### GetName

`func (o *ExperimentTrial) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ExperimentTrial) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ExperimentTrial) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ExperimentTrial) HasName() bool`

HasName returns a boolean if a field has been set.

### GetProject

`func (o *ExperimentTrial) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *ExperimentTrial) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *ExperimentTrial) SetProject(v string)`

SetProject sets Project field to given value.

### HasProject

`func (o *ExperimentTrial) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetStatus

`func (o *ExperimentTrial) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ExperimentTrial) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ExperimentTrial) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ExperimentTrial) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSubjectKind

`func (o *ExperimentTrial) GetSubjectKind() string`

GetSubjectKind returns the SubjectKind field if non-nil, zero value otherwise.

### GetSubjectKindOk

`func (o *ExperimentTrial) GetSubjectKindOk() (*string, bool)`

GetSubjectKindOk returns a tuple with the SubjectKind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectKind

`func (o *ExperimentTrial) SetSubjectKind(v string)`

SetSubjectKind sets SubjectKind field to given value.

### HasSubjectKind

`func (o *ExperimentTrial) HasSubjectKind() bool`

HasSubjectKind returns a boolean if a field has been set.

### GetVariants

`func (o *ExperimentTrial) GetVariants() []ExperimentArm`

GetVariants returns the Variants field if non-nil, zero value otherwise.

### GetVariantsOk

`func (o *ExperimentTrial) GetVariantsOk() (*[]ExperimentArm, bool)`

GetVariantsOk returns a tuple with the Variants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariants

`func (o *ExperimentTrial) SetVariants(v []ExperimentArm)`

SetVariants sets Variants field to given value.

### HasVariants

`func (o *ExperimentTrial) HasVariants() bool`

HasVariants returns a boolean if a field has been set.

### GetWinner

`func (o *ExperimentTrial) GetWinner() string`

GetWinner returns the Winner field if non-nil, zero value otherwise.

### GetWinnerOk

`func (o *ExperimentTrial) GetWinnerOk() (*string, bool)`

GetWinnerOk returns a tuple with the Winner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWinner

`func (o *ExperimentTrial) SetWinner(v string)`

SetWinner sets Winner field to given value.

### HasWinner

`func (o *ExperimentTrial) HasWinner() bool`

HasWinner returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


