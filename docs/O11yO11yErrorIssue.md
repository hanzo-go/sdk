# O11yO11yErrorIssue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Assignee** | Pointer to **string** | Assignee is who the issue is assigned to. | [optional] 
**Count** | Pointer to **int32** | Count is how many occurrences have landed on the issue. | [optional] 
**CreatedAt** | Pointer to **time.Time** | CreatedAt is when the issue was first recorded. | [optional] 
**Culprit** | Pointer to **string** | Culprit is where it came from — the function or route blamed for it. | [optional] 
**Environment** | Pointer to **string** | Environment is the deployment the issue was seen in. | [optional] 
**Fingerprint** | Pointer to **string** | Fingerprint is the grouping key that puts like errors in one issue. | [optional] 
**FirstSeen** | Pointer to **time.Time** | FirstSeen is when the earliest occurrence was recorded. | [optional] 
**Id** | Pointer to **string** | ID is the issue id. | [optional] 
**LastSeen** | Pointer to **time.Time** | LastSeen is when the latest was. | [optional] 
**Level** | Pointer to **string** | Level is the issue&#39;s severity, e.g. error, warning, info. | [optional] 
**Platform** | Pointer to **string** | Platform is the reporting runtime, e.g. go, python, javascript. | [optional] 
**Regressed** | Pointer to **bool** | Regressed marks an issue that reopened after being resolved. | [optional] 
**Release** | Pointer to **string** | Release is the version that produced it. | [optional] 
**ResolvedAt** | Pointer to **time.Time** | ResolvedAt is when the issue was resolved, if it is. | [optional] 
**ServiceName** | Pointer to **string** | ServiceName is the service that reported it. | [optional] 
**Status** | Pointer to **string** | Status is the lifecycle state: unresolved, resolved or ignored. | [optional] 
**Type** | Pointer to **string** | Type is the exception type. | [optional] 
**UpdatedAt** | Pointer to **time.Time** | UpdatedAt is when the issue last changed. | [optional] 
**Value** | Pointer to **string** | Value is the exception value. | [optional] 

## Methods

### NewO11yO11yErrorIssue

`func NewO11yO11yErrorIssue() *O11yO11yErrorIssue`

NewO11yO11yErrorIssue instantiates a new O11yO11yErrorIssue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yO11yErrorIssueWithDefaults

`func NewO11yO11yErrorIssueWithDefaults() *O11yO11yErrorIssue`

NewO11yO11yErrorIssueWithDefaults instantiates a new O11yO11yErrorIssue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssignee

`func (o *O11yO11yErrorIssue) GetAssignee() string`

GetAssignee returns the Assignee field if non-nil, zero value otherwise.

### GetAssigneeOk

`func (o *O11yO11yErrorIssue) GetAssigneeOk() (*string, bool)`

GetAssigneeOk returns a tuple with the Assignee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignee

`func (o *O11yO11yErrorIssue) SetAssignee(v string)`

SetAssignee sets Assignee field to given value.

### HasAssignee

`func (o *O11yO11yErrorIssue) HasAssignee() bool`

HasAssignee returns a boolean if a field has been set.

### GetCount

`func (o *O11yO11yErrorIssue) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *O11yO11yErrorIssue) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *O11yO11yErrorIssue) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *O11yO11yErrorIssue) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetCreatedAt

`func (o *O11yO11yErrorIssue) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *O11yO11yErrorIssue) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *O11yO11yErrorIssue) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *O11yO11yErrorIssue) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCulprit

`func (o *O11yO11yErrorIssue) GetCulprit() string`

GetCulprit returns the Culprit field if non-nil, zero value otherwise.

### GetCulpritOk

`func (o *O11yO11yErrorIssue) GetCulpritOk() (*string, bool)`

GetCulpritOk returns a tuple with the Culprit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCulprit

`func (o *O11yO11yErrorIssue) SetCulprit(v string)`

SetCulprit sets Culprit field to given value.

### HasCulprit

`func (o *O11yO11yErrorIssue) HasCulprit() bool`

HasCulprit returns a boolean if a field has been set.

### GetEnvironment

`func (o *O11yO11yErrorIssue) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *O11yO11yErrorIssue) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *O11yO11yErrorIssue) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *O11yO11yErrorIssue) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetFingerprint

`func (o *O11yO11yErrorIssue) GetFingerprint() string`

GetFingerprint returns the Fingerprint field if non-nil, zero value otherwise.

### GetFingerprintOk

`func (o *O11yO11yErrorIssue) GetFingerprintOk() (*string, bool)`

GetFingerprintOk returns a tuple with the Fingerprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFingerprint

`func (o *O11yO11yErrorIssue) SetFingerprint(v string)`

SetFingerprint sets Fingerprint field to given value.

### HasFingerprint

`func (o *O11yO11yErrorIssue) HasFingerprint() bool`

HasFingerprint returns a boolean if a field has been set.

### GetFirstSeen

`func (o *O11yO11yErrorIssue) GetFirstSeen() time.Time`

GetFirstSeen returns the FirstSeen field if non-nil, zero value otherwise.

### GetFirstSeenOk

`func (o *O11yO11yErrorIssue) GetFirstSeenOk() (*time.Time, bool)`

GetFirstSeenOk returns a tuple with the FirstSeen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstSeen

`func (o *O11yO11yErrorIssue) SetFirstSeen(v time.Time)`

SetFirstSeen sets FirstSeen field to given value.

### HasFirstSeen

`func (o *O11yO11yErrorIssue) HasFirstSeen() bool`

HasFirstSeen returns a boolean if a field has been set.

### GetId

`func (o *O11yO11yErrorIssue) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *O11yO11yErrorIssue) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *O11yO11yErrorIssue) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *O11yO11yErrorIssue) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastSeen

`func (o *O11yO11yErrorIssue) GetLastSeen() time.Time`

GetLastSeen returns the LastSeen field if non-nil, zero value otherwise.

### GetLastSeenOk

`func (o *O11yO11yErrorIssue) GetLastSeenOk() (*time.Time, bool)`

GetLastSeenOk returns a tuple with the LastSeen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSeen

`func (o *O11yO11yErrorIssue) SetLastSeen(v time.Time)`

SetLastSeen sets LastSeen field to given value.

### HasLastSeen

`func (o *O11yO11yErrorIssue) HasLastSeen() bool`

HasLastSeen returns a boolean if a field has been set.

### GetLevel

`func (o *O11yO11yErrorIssue) GetLevel() string`

GetLevel returns the Level field if non-nil, zero value otherwise.

### GetLevelOk

`func (o *O11yO11yErrorIssue) GetLevelOk() (*string, bool)`

GetLevelOk returns a tuple with the Level field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel

`func (o *O11yO11yErrorIssue) SetLevel(v string)`

SetLevel sets Level field to given value.

### HasLevel

`func (o *O11yO11yErrorIssue) HasLevel() bool`

HasLevel returns a boolean if a field has been set.

### GetPlatform

`func (o *O11yO11yErrorIssue) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *O11yO11yErrorIssue) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *O11yO11yErrorIssue) SetPlatform(v string)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *O11yO11yErrorIssue) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### GetRegressed

`func (o *O11yO11yErrorIssue) GetRegressed() bool`

GetRegressed returns the Regressed field if non-nil, zero value otherwise.

### GetRegressedOk

`func (o *O11yO11yErrorIssue) GetRegressedOk() (*bool, bool)`

GetRegressedOk returns a tuple with the Regressed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegressed

`func (o *O11yO11yErrorIssue) SetRegressed(v bool)`

SetRegressed sets Regressed field to given value.

### HasRegressed

`func (o *O11yO11yErrorIssue) HasRegressed() bool`

HasRegressed returns a boolean if a field has been set.

### GetRelease

`func (o *O11yO11yErrorIssue) GetRelease() string`

GetRelease returns the Release field if non-nil, zero value otherwise.

### GetReleaseOk

`func (o *O11yO11yErrorIssue) GetReleaseOk() (*string, bool)`

GetReleaseOk returns a tuple with the Release field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelease

`func (o *O11yO11yErrorIssue) SetRelease(v string)`

SetRelease sets Release field to given value.

### HasRelease

`func (o *O11yO11yErrorIssue) HasRelease() bool`

HasRelease returns a boolean if a field has been set.

### GetResolvedAt

`func (o *O11yO11yErrorIssue) GetResolvedAt() time.Time`

GetResolvedAt returns the ResolvedAt field if non-nil, zero value otherwise.

### GetResolvedAtOk

`func (o *O11yO11yErrorIssue) GetResolvedAtOk() (*time.Time, bool)`

GetResolvedAtOk returns a tuple with the ResolvedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolvedAt

`func (o *O11yO11yErrorIssue) SetResolvedAt(v time.Time)`

SetResolvedAt sets ResolvedAt field to given value.

### HasResolvedAt

`func (o *O11yO11yErrorIssue) HasResolvedAt() bool`

HasResolvedAt returns a boolean if a field has been set.

### GetServiceName

`func (o *O11yO11yErrorIssue) GetServiceName() string`

GetServiceName returns the ServiceName field if non-nil, zero value otherwise.

### GetServiceNameOk

`func (o *O11yO11yErrorIssue) GetServiceNameOk() (*string, bool)`

GetServiceNameOk returns a tuple with the ServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceName

`func (o *O11yO11yErrorIssue) SetServiceName(v string)`

SetServiceName sets ServiceName field to given value.

### HasServiceName

`func (o *O11yO11yErrorIssue) HasServiceName() bool`

HasServiceName returns a boolean if a field has been set.

### GetStatus

`func (o *O11yO11yErrorIssue) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *O11yO11yErrorIssue) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *O11yO11yErrorIssue) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *O11yO11yErrorIssue) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetType

`func (o *O11yO11yErrorIssue) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *O11yO11yErrorIssue) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *O11yO11yErrorIssue) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *O11yO11yErrorIssue) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *O11yO11yErrorIssue) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *O11yO11yErrorIssue) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *O11yO11yErrorIssue) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *O11yO11yErrorIssue) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetValue

`func (o *O11yO11yErrorIssue) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *O11yO11yErrorIssue) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *O11yO11yErrorIssue) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *O11yO11yErrorIssue) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


