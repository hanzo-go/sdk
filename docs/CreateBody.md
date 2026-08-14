# CreateBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExposureEvent** | Pointer to **string** | ExposureEvent is the event that marks a subject as enrolled — the analysis denominator — defaulting to the SDK&#39;s $feature_flag_called marker. | [optional] 
**FlagKey** | Pointer to **string** | FlagKey names the assignment flag this experiment writes, defaulting to exp_&lt;id&gt;. It must be a slug. | [optional] 
**Id** | **string** | ID is the experiment&#39;s slug, claimed once per project. It must match [A-Za-z0-9][A-Za-z0-9._-]{0,127}. | 
**MetricEvent** | **string** | MetricEvent is the event that counts as a conversion — the analysis numerator. | 
**Name** | Pointer to **string** | Name is free text for a reader; the id is what addresses the experiment. | [optional] 
**SubjectKind** | Pointer to **string** | SubjectKind is the unit assigned and measured: user (the default), org, session or audience. | [optional] 
**Variants** | [**[]Arm**](Arm.md) | Arms are the arms, at least two. Weights that are all zero become an even split; otherwise they must sum to 100, and at most one arm may be flagged control. | 

## Methods

### NewCreateBody

`func NewCreateBody(id string, metricEvent string, variants []Arm, ) *CreateBody`

NewCreateBody instantiates a new CreateBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateBodyWithDefaults

`func NewCreateBodyWithDefaults() *CreateBody`

NewCreateBodyWithDefaults instantiates a new CreateBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExposureEvent

`func (o *CreateBody) GetExposureEvent() string`

GetExposureEvent returns the ExposureEvent field if non-nil, zero value otherwise.

### GetExposureEventOk

`func (o *CreateBody) GetExposureEventOk() (*string, bool)`

GetExposureEventOk returns a tuple with the ExposureEvent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExposureEvent

`func (o *CreateBody) SetExposureEvent(v string)`

SetExposureEvent sets ExposureEvent field to given value.

### HasExposureEvent

`func (o *CreateBody) HasExposureEvent() bool`

HasExposureEvent returns a boolean if a field has been set.

### GetFlagKey

`func (o *CreateBody) GetFlagKey() string`

GetFlagKey returns the FlagKey field if non-nil, zero value otherwise.

### GetFlagKeyOk

`func (o *CreateBody) GetFlagKeyOk() (*string, bool)`

GetFlagKeyOk returns a tuple with the FlagKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlagKey

`func (o *CreateBody) SetFlagKey(v string)`

SetFlagKey sets FlagKey field to given value.

### HasFlagKey

`func (o *CreateBody) HasFlagKey() bool`

HasFlagKey returns a boolean if a field has been set.

### GetId

`func (o *CreateBody) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CreateBody) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CreateBody) SetId(v string)`

SetId sets Id field to given value.


### GetMetricEvent

`func (o *CreateBody) GetMetricEvent() string`

GetMetricEvent returns the MetricEvent field if non-nil, zero value otherwise.

### GetMetricEventOk

`func (o *CreateBody) GetMetricEventOk() (*string, bool)`

GetMetricEventOk returns a tuple with the MetricEvent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricEvent

`func (o *CreateBody) SetMetricEvent(v string)`

SetMetricEvent sets MetricEvent field to given value.


### GetName

`func (o *CreateBody) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateBody) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateBody) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CreateBody) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSubjectKind

`func (o *CreateBody) GetSubjectKind() string`

GetSubjectKind returns the SubjectKind field if non-nil, zero value otherwise.

### GetSubjectKindOk

`func (o *CreateBody) GetSubjectKindOk() (*string, bool)`

GetSubjectKindOk returns a tuple with the SubjectKind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectKind

`func (o *CreateBody) SetSubjectKind(v string)`

SetSubjectKind sets SubjectKind field to given value.

### HasSubjectKind

`func (o *CreateBody) HasSubjectKind() bool`

HasSubjectKind returns a boolean if a field has been set.

### GetVariants

`func (o *CreateBody) GetVariants() []Arm`

GetVariants returns the Variants field if non-nil, zero value otherwise.

### GetVariantsOk

`func (o *CreateBody) GetVariantsOk() (*[]Arm, bool)`

GetVariantsOk returns a tuple with the Variants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariants

`func (o *CreateBody) SetVariants(v []Arm)`

SetVariants sets Variants field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


