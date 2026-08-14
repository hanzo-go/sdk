# O11yJiraConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NotifierConfig** | Pointer to [**O11yNotifierConfig**](O11yNotifierConfig.md) |  | [optional] 
**ApiType** | Pointer to **string** |  | [optional] 
**ApiUrl** | Pointer to **interface{}** |  | [optional] 
**CustomFields** | Pointer to **map[string]map[string]interface{}** |  | [optional] 
**Description** | Pointer to [**O11yJiraFieldConfig**](O11yJiraFieldConfig.md) |  | [optional] 
**HttpConfig** | Pointer to [**O11yHTTPClientConfig**](O11yHTTPClientConfig.md) |  | [optional] 
**IssueType** | Pointer to **string** |  | [optional] 
**Labels** | Pointer to **[]string** |  | [optional] 
**Priority** | Pointer to **string** |  | [optional] 
**Project** | Pointer to **string** |  | [optional] 
**ReopenDuration** | Pointer to **interface{}** |  | [optional] 
**ReopenTransition** | Pointer to **string** |  | [optional] 
**ResolveTransition** | Pointer to **string** |  | [optional] 
**Summary** | Pointer to [**O11yJiraFieldConfig**](O11yJiraFieldConfig.md) |  | [optional] 
**WontFixResolution** | Pointer to **string** |  | [optional] 

## Methods

### NewO11yJiraConfig

`func NewO11yJiraConfig() *O11yJiraConfig`

NewO11yJiraConfig instantiates a new O11yJiraConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yJiraConfigWithDefaults

`func NewO11yJiraConfigWithDefaults() *O11yJiraConfig`

NewO11yJiraConfigWithDefaults instantiates a new O11yJiraConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNotifierConfig

`func (o *O11yJiraConfig) GetNotifierConfig() O11yNotifierConfig`

GetNotifierConfig returns the NotifierConfig field if non-nil, zero value otherwise.

### GetNotifierConfigOk

`func (o *O11yJiraConfig) GetNotifierConfigOk() (*O11yNotifierConfig, bool)`

GetNotifierConfigOk returns a tuple with the NotifierConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifierConfig

`func (o *O11yJiraConfig) SetNotifierConfig(v O11yNotifierConfig)`

SetNotifierConfig sets NotifierConfig field to given value.

### HasNotifierConfig

`func (o *O11yJiraConfig) HasNotifierConfig() bool`

HasNotifierConfig returns a boolean if a field has been set.

### GetApiType

`func (o *O11yJiraConfig) GetApiType() string`

GetApiType returns the ApiType field if non-nil, zero value otherwise.

### GetApiTypeOk

`func (o *O11yJiraConfig) GetApiTypeOk() (*string, bool)`

GetApiTypeOk returns a tuple with the ApiType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiType

`func (o *O11yJiraConfig) SetApiType(v string)`

SetApiType sets ApiType field to given value.

### HasApiType

`func (o *O11yJiraConfig) HasApiType() bool`

HasApiType returns a boolean if a field has been set.

### GetApiUrl

`func (o *O11yJiraConfig) GetApiUrl() interface{}`

GetApiUrl returns the ApiUrl field if non-nil, zero value otherwise.

### GetApiUrlOk

`func (o *O11yJiraConfig) GetApiUrlOk() (*interface{}, bool)`

GetApiUrlOk returns a tuple with the ApiUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiUrl

`func (o *O11yJiraConfig) SetApiUrl(v interface{})`

SetApiUrl sets ApiUrl field to given value.

### HasApiUrl

`func (o *O11yJiraConfig) HasApiUrl() bool`

HasApiUrl returns a boolean if a field has been set.

### SetApiUrlNil

`func (o *O11yJiraConfig) SetApiUrlNil(b bool)`

 SetApiUrlNil sets the value for ApiUrl to be an explicit nil

### UnsetApiUrl
`func (o *O11yJiraConfig) UnsetApiUrl()`

UnsetApiUrl ensures that no value is present for ApiUrl, not even an explicit nil
### GetCustomFields

`func (o *O11yJiraConfig) GetCustomFields() map[string]map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *O11yJiraConfig) GetCustomFieldsOk() (*map[string]map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *O11yJiraConfig) SetCustomFields(v map[string]map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *O11yJiraConfig) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetDescription

`func (o *O11yJiraConfig) GetDescription() O11yJiraFieldConfig`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *O11yJiraConfig) GetDescriptionOk() (*O11yJiraFieldConfig, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *O11yJiraConfig) SetDescription(v O11yJiraFieldConfig)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *O11yJiraConfig) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetHttpConfig

`func (o *O11yJiraConfig) GetHttpConfig() O11yHTTPClientConfig`

GetHttpConfig returns the HttpConfig field if non-nil, zero value otherwise.

### GetHttpConfigOk

`func (o *O11yJiraConfig) GetHttpConfigOk() (*O11yHTTPClientConfig, bool)`

GetHttpConfigOk returns a tuple with the HttpConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpConfig

`func (o *O11yJiraConfig) SetHttpConfig(v O11yHTTPClientConfig)`

SetHttpConfig sets HttpConfig field to given value.

### HasHttpConfig

`func (o *O11yJiraConfig) HasHttpConfig() bool`

HasHttpConfig returns a boolean if a field has been set.

### GetIssueType

`func (o *O11yJiraConfig) GetIssueType() string`

GetIssueType returns the IssueType field if non-nil, zero value otherwise.

### GetIssueTypeOk

`func (o *O11yJiraConfig) GetIssueTypeOk() (*string, bool)`

GetIssueTypeOk returns a tuple with the IssueType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueType

`func (o *O11yJiraConfig) SetIssueType(v string)`

SetIssueType sets IssueType field to given value.

### HasIssueType

`func (o *O11yJiraConfig) HasIssueType() bool`

HasIssueType returns a boolean if a field has been set.

### GetLabels

`func (o *O11yJiraConfig) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *O11yJiraConfig) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *O11yJiraConfig) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *O11yJiraConfig) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetPriority

`func (o *O11yJiraConfig) GetPriority() string`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *O11yJiraConfig) GetPriorityOk() (*string, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *O11yJiraConfig) SetPriority(v string)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *O11yJiraConfig) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetProject

`func (o *O11yJiraConfig) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *O11yJiraConfig) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *O11yJiraConfig) SetProject(v string)`

SetProject sets Project field to given value.

### HasProject

`func (o *O11yJiraConfig) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetReopenDuration

`func (o *O11yJiraConfig) GetReopenDuration() interface{}`

GetReopenDuration returns the ReopenDuration field if non-nil, zero value otherwise.

### GetReopenDurationOk

`func (o *O11yJiraConfig) GetReopenDurationOk() (*interface{}, bool)`

GetReopenDurationOk returns a tuple with the ReopenDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReopenDuration

`func (o *O11yJiraConfig) SetReopenDuration(v interface{})`

SetReopenDuration sets ReopenDuration field to given value.

### HasReopenDuration

`func (o *O11yJiraConfig) HasReopenDuration() bool`

HasReopenDuration returns a boolean if a field has been set.

### SetReopenDurationNil

`func (o *O11yJiraConfig) SetReopenDurationNil(b bool)`

 SetReopenDurationNil sets the value for ReopenDuration to be an explicit nil

### UnsetReopenDuration
`func (o *O11yJiraConfig) UnsetReopenDuration()`

UnsetReopenDuration ensures that no value is present for ReopenDuration, not even an explicit nil
### GetReopenTransition

`func (o *O11yJiraConfig) GetReopenTransition() string`

GetReopenTransition returns the ReopenTransition field if non-nil, zero value otherwise.

### GetReopenTransitionOk

`func (o *O11yJiraConfig) GetReopenTransitionOk() (*string, bool)`

GetReopenTransitionOk returns a tuple with the ReopenTransition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReopenTransition

`func (o *O11yJiraConfig) SetReopenTransition(v string)`

SetReopenTransition sets ReopenTransition field to given value.

### HasReopenTransition

`func (o *O11yJiraConfig) HasReopenTransition() bool`

HasReopenTransition returns a boolean if a field has been set.

### GetResolveTransition

`func (o *O11yJiraConfig) GetResolveTransition() string`

GetResolveTransition returns the ResolveTransition field if non-nil, zero value otherwise.

### GetResolveTransitionOk

`func (o *O11yJiraConfig) GetResolveTransitionOk() (*string, bool)`

GetResolveTransitionOk returns a tuple with the ResolveTransition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolveTransition

`func (o *O11yJiraConfig) SetResolveTransition(v string)`

SetResolveTransition sets ResolveTransition field to given value.

### HasResolveTransition

`func (o *O11yJiraConfig) HasResolveTransition() bool`

HasResolveTransition returns a boolean if a field has been set.

### GetSummary

`func (o *O11yJiraConfig) GetSummary() O11yJiraFieldConfig`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *O11yJiraConfig) GetSummaryOk() (*O11yJiraFieldConfig, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *O11yJiraConfig) SetSummary(v O11yJiraFieldConfig)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *O11yJiraConfig) HasSummary() bool`

HasSummary returns a boolean if a field has been set.

### GetWontFixResolution

`func (o *O11yJiraConfig) GetWontFixResolution() string`

GetWontFixResolution returns the WontFixResolution field if non-nil, zero value otherwise.

### GetWontFixResolutionOk

`func (o *O11yJiraConfig) GetWontFixResolutionOk() (*string, bool)`

GetWontFixResolutionOk returns a tuple with the WontFixResolution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWontFixResolution

`func (o *O11yJiraConfig) SetWontFixResolution(v string)`

SetWontFixResolution sets WontFixResolution field to given value.

### HasWontFixResolution

`func (o *O11yJiraConfig) HasWontFixResolution() bool`

HasWontFixResolution returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


