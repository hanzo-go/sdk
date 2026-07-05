# O11yAlertRuleCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] [default to true]
**Datasource** | **string** |  | 
**Query** | **string** |  | 
**Condition** | [**O11yAlertRuleCreateCondition**](O11yAlertRuleCreateCondition.md) |  | 
**ForDuration** | Pointer to **string** |  | [optional] [default to "5m"]
**Severity** | Pointer to **string** |  | [optional] [default to "warning"]
**Labels** | Pointer to **map[string]string** |  | [optional] 
**Annotations** | Pointer to **map[string]string** |  | [optional] 
**ChannelIds** | Pointer to **[]string** |  | [optional] 

## Methods

### NewO11yAlertRuleCreate

`func NewO11yAlertRuleCreate(name string, datasource string, query string, condition O11yAlertRuleCreateCondition, ) *O11yAlertRuleCreate`

NewO11yAlertRuleCreate instantiates a new O11yAlertRuleCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yAlertRuleCreateWithDefaults

`func NewO11yAlertRuleCreateWithDefaults() *O11yAlertRuleCreate`

NewO11yAlertRuleCreateWithDefaults instantiates a new O11yAlertRuleCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *O11yAlertRuleCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *O11yAlertRuleCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *O11yAlertRuleCreate) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *O11yAlertRuleCreate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *O11yAlertRuleCreate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *O11yAlertRuleCreate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *O11yAlertRuleCreate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *O11yAlertRuleCreate) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *O11yAlertRuleCreate) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *O11yAlertRuleCreate) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *O11yAlertRuleCreate) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetDatasource

`func (o *O11yAlertRuleCreate) GetDatasource() string`

GetDatasource returns the Datasource field if non-nil, zero value otherwise.

### GetDatasourceOk

`func (o *O11yAlertRuleCreate) GetDatasourceOk() (*string, bool)`

GetDatasourceOk returns a tuple with the Datasource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatasource

`func (o *O11yAlertRuleCreate) SetDatasource(v string)`

SetDatasource sets Datasource field to given value.


### GetQuery

`func (o *O11yAlertRuleCreate) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *O11yAlertRuleCreate) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *O11yAlertRuleCreate) SetQuery(v string)`

SetQuery sets Query field to given value.


### GetCondition

`func (o *O11yAlertRuleCreate) GetCondition() O11yAlertRuleCreateCondition`

GetCondition returns the Condition field if non-nil, zero value otherwise.

### GetConditionOk

`func (o *O11yAlertRuleCreate) GetConditionOk() (*O11yAlertRuleCreateCondition, bool)`

GetConditionOk returns a tuple with the Condition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCondition

`func (o *O11yAlertRuleCreate) SetCondition(v O11yAlertRuleCreateCondition)`

SetCondition sets Condition field to given value.


### GetForDuration

`func (o *O11yAlertRuleCreate) GetForDuration() string`

GetForDuration returns the ForDuration field if non-nil, zero value otherwise.

### GetForDurationOk

`func (o *O11yAlertRuleCreate) GetForDurationOk() (*string, bool)`

GetForDurationOk returns a tuple with the ForDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForDuration

`func (o *O11yAlertRuleCreate) SetForDuration(v string)`

SetForDuration sets ForDuration field to given value.

### HasForDuration

`func (o *O11yAlertRuleCreate) HasForDuration() bool`

HasForDuration returns a boolean if a field has been set.

### GetSeverity

`func (o *O11yAlertRuleCreate) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *O11yAlertRuleCreate) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *O11yAlertRuleCreate) SetSeverity(v string)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *O11yAlertRuleCreate) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### GetLabels

`func (o *O11yAlertRuleCreate) GetLabels() map[string]string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *O11yAlertRuleCreate) GetLabelsOk() (*map[string]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *O11yAlertRuleCreate) SetLabels(v map[string]string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *O11yAlertRuleCreate) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetAnnotations

`func (o *O11yAlertRuleCreate) GetAnnotations() map[string]string`

GetAnnotations returns the Annotations field if non-nil, zero value otherwise.

### GetAnnotationsOk

`func (o *O11yAlertRuleCreate) GetAnnotationsOk() (*map[string]string, bool)`

GetAnnotationsOk returns a tuple with the Annotations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotations

`func (o *O11yAlertRuleCreate) SetAnnotations(v map[string]string)`

SetAnnotations sets Annotations field to given value.

### HasAnnotations

`func (o *O11yAlertRuleCreate) HasAnnotations() bool`

HasAnnotations returns a boolean if a field has been set.

### GetChannelIds

`func (o *O11yAlertRuleCreate) GetChannelIds() []string`

GetChannelIds returns the ChannelIds field if non-nil, zero value otherwise.

### GetChannelIdsOk

`func (o *O11yAlertRuleCreate) GetChannelIdsOk() (*[]string, bool)`

GetChannelIdsOk returns a tuple with the ChannelIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelIds

`func (o *O11yAlertRuleCreate) SetChannelIds(v []string)`

SetChannelIds sets ChannelIds field to given value.

### HasChannelIds

`func (o *O11yAlertRuleCreate) HasChannelIds() bool`

HasChannelIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


