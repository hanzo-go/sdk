# O11yAlertRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**Datasource** | Pointer to **string** |  | [optional] 
**Query** | Pointer to **string** | LogQL or PromQL expression | [optional] 
**Condition** | Pointer to [**O11yAlertRuleCondition**](O11yAlertRuleCondition.md) |  | [optional] 
**ForDuration** | Pointer to **string** | Duration condition must be true before firing (e.g. 5m) | [optional] 
**Severity** | Pointer to **string** |  | [optional] 
**Labels** | Pointer to **map[string]string** |  | [optional] 
**Annotations** | Pointer to **map[string]string** | Template-based annotations (summary, description, runbook) | [optional] 
**ChannelIds** | Pointer to **[]string** | Notification channels to alert | [optional] 
**State** | Pointer to **string** |  | [optional] 
**LastEvaluated** | Pointer to **time.Time** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewO11yAlertRule

`func NewO11yAlertRule() *O11yAlertRule`

NewO11yAlertRule instantiates a new O11yAlertRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yAlertRuleWithDefaults

`func NewO11yAlertRuleWithDefaults() *O11yAlertRule`

NewO11yAlertRuleWithDefaults instantiates a new O11yAlertRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *O11yAlertRule) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *O11yAlertRule) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *O11yAlertRule) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *O11yAlertRule) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *O11yAlertRule) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *O11yAlertRule) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *O11yAlertRule) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *O11yAlertRule) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *O11yAlertRule) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *O11yAlertRule) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *O11yAlertRule) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *O11yAlertRule) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *O11yAlertRule) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *O11yAlertRule) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *O11yAlertRule) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *O11yAlertRule) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetDatasource

`func (o *O11yAlertRule) GetDatasource() string`

GetDatasource returns the Datasource field if non-nil, zero value otherwise.

### GetDatasourceOk

`func (o *O11yAlertRule) GetDatasourceOk() (*string, bool)`

GetDatasourceOk returns a tuple with the Datasource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatasource

`func (o *O11yAlertRule) SetDatasource(v string)`

SetDatasource sets Datasource field to given value.

### HasDatasource

`func (o *O11yAlertRule) HasDatasource() bool`

HasDatasource returns a boolean if a field has been set.

### GetQuery

`func (o *O11yAlertRule) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *O11yAlertRule) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *O11yAlertRule) SetQuery(v string)`

SetQuery sets Query field to given value.

### HasQuery

`func (o *O11yAlertRule) HasQuery() bool`

HasQuery returns a boolean if a field has been set.

### GetCondition

`func (o *O11yAlertRule) GetCondition() O11yAlertRuleCondition`

GetCondition returns the Condition field if non-nil, zero value otherwise.

### GetConditionOk

`func (o *O11yAlertRule) GetConditionOk() (*O11yAlertRuleCondition, bool)`

GetConditionOk returns a tuple with the Condition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCondition

`func (o *O11yAlertRule) SetCondition(v O11yAlertRuleCondition)`

SetCondition sets Condition field to given value.

### HasCondition

`func (o *O11yAlertRule) HasCondition() bool`

HasCondition returns a boolean if a field has been set.

### GetForDuration

`func (o *O11yAlertRule) GetForDuration() string`

GetForDuration returns the ForDuration field if non-nil, zero value otherwise.

### GetForDurationOk

`func (o *O11yAlertRule) GetForDurationOk() (*string, bool)`

GetForDurationOk returns a tuple with the ForDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForDuration

`func (o *O11yAlertRule) SetForDuration(v string)`

SetForDuration sets ForDuration field to given value.

### HasForDuration

`func (o *O11yAlertRule) HasForDuration() bool`

HasForDuration returns a boolean if a field has been set.

### GetSeverity

`func (o *O11yAlertRule) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *O11yAlertRule) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *O11yAlertRule) SetSeverity(v string)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *O11yAlertRule) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### GetLabels

`func (o *O11yAlertRule) GetLabels() map[string]string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *O11yAlertRule) GetLabelsOk() (*map[string]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *O11yAlertRule) SetLabels(v map[string]string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *O11yAlertRule) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetAnnotations

`func (o *O11yAlertRule) GetAnnotations() map[string]string`

GetAnnotations returns the Annotations field if non-nil, zero value otherwise.

### GetAnnotationsOk

`func (o *O11yAlertRule) GetAnnotationsOk() (*map[string]string, bool)`

GetAnnotationsOk returns a tuple with the Annotations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotations

`func (o *O11yAlertRule) SetAnnotations(v map[string]string)`

SetAnnotations sets Annotations field to given value.

### HasAnnotations

`func (o *O11yAlertRule) HasAnnotations() bool`

HasAnnotations returns a boolean if a field has been set.

### GetChannelIds

`func (o *O11yAlertRule) GetChannelIds() []string`

GetChannelIds returns the ChannelIds field if non-nil, zero value otherwise.

### GetChannelIdsOk

`func (o *O11yAlertRule) GetChannelIdsOk() (*[]string, bool)`

GetChannelIdsOk returns a tuple with the ChannelIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelIds

`func (o *O11yAlertRule) SetChannelIds(v []string)`

SetChannelIds sets ChannelIds field to given value.

### HasChannelIds

`func (o *O11yAlertRule) HasChannelIds() bool`

HasChannelIds returns a boolean if a field has been set.

### GetState

`func (o *O11yAlertRule) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *O11yAlertRule) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *O11yAlertRule) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *O11yAlertRule) HasState() bool`

HasState returns a boolean if a field has been set.

### GetLastEvaluated

`func (o *O11yAlertRule) GetLastEvaluated() time.Time`

GetLastEvaluated returns the LastEvaluated field if non-nil, zero value otherwise.

### GetLastEvaluatedOk

`func (o *O11yAlertRule) GetLastEvaluatedOk() (*time.Time, bool)`

GetLastEvaluatedOk returns a tuple with the LastEvaluated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastEvaluated

`func (o *O11yAlertRule) SetLastEvaluated(v time.Time)`

SetLastEvaluated sets LastEvaluated field to given value.

### HasLastEvaluated

`func (o *O11yAlertRule) HasLastEvaluated() bool`

HasLastEvaluated returns a boolean if a field has been set.

### GetCreatedAt

`func (o *O11yAlertRule) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *O11yAlertRule) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *O11yAlertRule) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *O11yAlertRule) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *O11yAlertRule) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *O11yAlertRule) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *O11yAlertRule) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *O11yAlertRule) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


