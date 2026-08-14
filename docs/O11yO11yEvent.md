# O11yO11yEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Culprit** | Pointer to **string** | Culprit is where it came from — the function or route blamed for it. | [optional] 
**Environment** | Pointer to **string** | Environment is the deployment it happened in. | [optional] 
**EventId** | Pointer to **string** | EventID is its own id. | [optional] 
**Fingerprint** | Pointer to **string** | Fingerprint is the grouping key that puts like errors in one issue. | [optional] 
**Frames** | Pointer to [**[]O11yO11yFrame**](O11yO11yFrame.md) | Frames are the stack, innermost first. | [optional] 
**Handled** | Pointer to **bool** | Handled says whether the application caught it. | [optional] 
**Level** | Pointer to **string** | Level is its severity, e.g. error, warning, info. | [optional] 
**Message** | Pointer to **string** | Message is the human-readable message. | [optional] 
**OrgId** | Pointer to **string** | OrgID is the org that owns it. | [optional] 
**Platform** | Pointer to **string** | Platform is the reporting runtime, e.g. go, python, javascript. | [optional] 
**ProjectId** | Pointer to **string** | ProjectID is the project it was captured for. | [optional] 
**ReceivedAt** | Pointer to **time.Time** | ReceivedAt is when it arrived here. | [optional] 
**Release** | Pointer to **string** | Release is the version that produced it. | [optional] 
**ServerName** | Pointer to **string** | ServerName is the host that reported it. | [optional] 
**ServiceName** | Pointer to **string** | ServiceName is the service that reported it. | [optional] 
**SpanId** | Pointer to **string** | SpanID is the span it belonged to. | [optional] 
**Tags** | Pointer to **map[string]string** | Tags are the reporter&#39;s own key/value labels. | [optional] 
**Timestamp** | Pointer to **time.Time** | Timestamp is when the error happened, as the reporter recorded it. | [optional] 
**TraceId** | Pointer to **string** | TraceID is the trace it belonged to. | [optional] 
**Transaction** | Pointer to **string** | Transaction is the operation it happened in. | [optional] 
**Type** | Pointer to **string** | Type is the exception type. | [optional] 
**UserEmail** | Pointer to **string** | UserEmail is that user&#39;s email, when attached. | [optional] 
**UserId** | Pointer to **string** | UserID identifies the affected end user, when the reporter attached one. | [optional] 
**UserIp** | Pointer to **string** | UserIP is that user&#39;s address, when attached. | [optional] 
**Value** | Pointer to **string** | Value is the exception value. | [optional] 

## Methods

### NewO11yO11yEvent

`func NewO11yO11yEvent() *O11yO11yEvent`

NewO11yO11yEvent instantiates a new O11yO11yEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yO11yEventWithDefaults

`func NewO11yO11yEventWithDefaults() *O11yO11yEvent`

NewO11yO11yEventWithDefaults instantiates a new O11yO11yEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCulprit

`func (o *O11yO11yEvent) GetCulprit() string`

GetCulprit returns the Culprit field if non-nil, zero value otherwise.

### GetCulpritOk

`func (o *O11yO11yEvent) GetCulpritOk() (*string, bool)`

GetCulpritOk returns a tuple with the Culprit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCulprit

`func (o *O11yO11yEvent) SetCulprit(v string)`

SetCulprit sets Culprit field to given value.

### HasCulprit

`func (o *O11yO11yEvent) HasCulprit() bool`

HasCulprit returns a boolean if a field has been set.

### GetEnvironment

`func (o *O11yO11yEvent) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *O11yO11yEvent) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *O11yO11yEvent) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *O11yO11yEvent) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetEventId

`func (o *O11yO11yEvent) GetEventId() string`

GetEventId returns the EventId field if non-nil, zero value otherwise.

### GetEventIdOk

`func (o *O11yO11yEvent) GetEventIdOk() (*string, bool)`

GetEventIdOk returns a tuple with the EventId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventId

`func (o *O11yO11yEvent) SetEventId(v string)`

SetEventId sets EventId field to given value.

### HasEventId

`func (o *O11yO11yEvent) HasEventId() bool`

HasEventId returns a boolean if a field has been set.

### GetFingerprint

`func (o *O11yO11yEvent) GetFingerprint() string`

GetFingerprint returns the Fingerprint field if non-nil, zero value otherwise.

### GetFingerprintOk

`func (o *O11yO11yEvent) GetFingerprintOk() (*string, bool)`

GetFingerprintOk returns a tuple with the Fingerprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFingerprint

`func (o *O11yO11yEvent) SetFingerprint(v string)`

SetFingerprint sets Fingerprint field to given value.

### HasFingerprint

`func (o *O11yO11yEvent) HasFingerprint() bool`

HasFingerprint returns a boolean if a field has been set.

### GetFrames

`func (o *O11yO11yEvent) GetFrames() []O11yO11yFrame`

GetFrames returns the Frames field if non-nil, zero value otherwise.

### GetFramesOk

`func (o *O11yO11yEvent) GetFramesOk() (*[]O11yO11yFrame, bool)`

GetFramesOk returns a tuple with the Frames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrames

`func (o *O11yO11yEvent) SetFrames(v []O11yO11yFrame)`

SetFrames sets Frames field to given value.

### HasFrames

`func (o *O11yO11yEvent) HasFrames() bool`

HasFrames returns a boolean if a field has been set.

### GetHandled

`func (o *O11yO11yEvent) GetHandled() bool`

GetHandled returns the Handled field if non-nil, zero value otherwise.

### GetHandledOk

`func (o *O11yO11yEvent) GetHandledOk() (*bool, bool)`

GetHandledOk returns a tuple with the Handled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHandled

`func (o *O11yO11yEvent) SetHandled(v bool)`

SetHandled sets Handled field to given value.

### HasHandled

`func (o *O11yO11yEvent) HasHandled() bool`

HasHandled returns a boolean if a field has been set.

### GetLevel

`func (o *O11yO11yEvent) GetLevel() string`

GetLevel returns the Level field if non-nil, zero value otherwise.

### GetLevelOk

`func (o *O11yO11yEvent) GetLevelOk() (*string, bool)`

GetLevelOk returns a tuple with the Level field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel

`func (o *O11yO11yEvent) SetLevel(v string)`

SetLevel sets Level field to given value.

### HasLevel

`func (o *O11yO11yEvent) HasLevel() bool`

HasLevel returns a boolean if a field has been set.

### GetMessage

`func (o *O11yO11yEvent) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *O11yO11yEvent) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *O11yO11yEvent) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *O11yO11yEvent) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetOrgId

`func (o *O11yO11yEvent) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *O11yO11yEvent) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *O11yO11yEvent) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *O11yO11yEvent) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### GetPlatform

`func (o *O11yO11yEvent) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *O11yO11yEvent) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *O11yO11yEvent) SetPlatform(v string)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *O11yO11yEvent) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### GetProjectId

`func (o *O11yO11yEvent) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *O11yO11yEvent) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *O11yO11yEvent) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *O11yO11yEvent) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetReceivedAt

`func (o *O11yO11yEvent) GetReceivedAt() time.Time`

GetReceivedAt returns the ReceivedAt field if non-nil, zero value otherwise.

### GetReceivedAtOk

`func (o *O11yO11yEvent) GetReceivedAtOk() (*time.Time, bool)`

GetReceivedAtOk returns a tuple with the ReceivedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceivedAt

`func (o *O11yO11yEvent) SetReceivedAt(v time.Time)`

SetReceivedAt sets ReceivedAt field to given value.

### HasReceivedAt

`func (o *O11yO11yEvent) HasReceivedAt() bool`

HasReceivedAt returns a boolean if a field has been set.

### GetRelease

`func (o *O11yO11yEvent) GetRelease() string`

GetRelease returns the Release field if non-nil, zero value otherwise.

### GetReleaseOk

`func (o *O11yO11yEvent) GetReleaseOk() (*string, bool)`

GetReleaseOk returns a tuple with the Release field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelease

`func (o *O11yO11yEvent) SetRelease(v string)`

SetRelease sets Release field to given value.

### HasRelease

`func (o *O11yO11yEvent) HasRelease() bool`

HasRelease returns a boolean if a field has been set.

### GetServerName

`func (o *O11yO11yEvent) GetServerName() string`

GetServerName returns the ServerName field if non-nil, zero value otherwise.

### GetServerNameOk

`func (o *O11yO11yEvent) GetServerNameOk() (*string, bool)`

GetServerNameOk returns a tuple with the ServerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerName

`func (o *O11yO11yEvent) SetServerName(v string)`

SetServerName sets ServerName field to given value.

### HasServerName

`func (o *O11yO11yEvent) HasServerName() bool`

HasServerName returns a boolean if a field has been set.

### GetServiceName

`func (o *O11yO11yEvent) GetServiceName() string`

GetServiceName returns the ServiceName field if non-nil, zero value otherwise.

### GetServiceNameOk

`func (o *O11yO11yEvent) GetServiceNameOk() (*string, bool)`

GetServiceNameOk returns a tuple with the ServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceName

`func (o *O11yO11yEvent) SetServiceName(v string)`

SetServiceName sets ServiceName field to given value.

### HasServiceName

`func (o *O11yO11yEvent) HasServiceName() bool`

HasServiceName returns a boolean if a field has been set.

### GetSpanId

`func (o *O11yO11yEvent) GetSpanId() string`

GetSpanId returns the SpanId field if non-nil, zero value otherwise.

### GetSpanIdOk

`func (o *O11yO11yEvent) GetSpanIdOk() (*string, bool)`

GetSpanIdOk returns a tuple with the SpanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpanId

`func (o *O11yO11yEvent) SetSpanId(v string)`

SetSpanId sets SpanId field to given value.

### HasSpanId

`func (o *O11yO11yEvent) HasSpanId() bool`

HasSpanId returns a boolean if a field has been set.

### GetTags

`func (o *O11yO11yEvent) GetTags() map[string]string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *O11yO11yEvent) GetTagsOk() (*map[string]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *O11yO11yEvent) SetTags(v map[string]string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *O11yO11yEvent) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTimestamp

`func (o *O11yO11yEvent) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *O11yO11yEvent) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *O11yO11yEvent) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *O11yO11yEvent) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetTraceId

`func (o *O11yO11yEvent) GetTraceId() string`

GetTraceId returns the TraceId field if non-nil, zero value otherwise.

### GetTraceIdOk

`func (o *O11yO11yEvent) GetTraceIdOk() (*string, bool)`

GetTraceIdOk returns a tuple with the TraceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraceId

`func (o *O11yO11yEvent) SetTraceId(v string)`

SetTraceId sets TraceId field to given value.

### HasTraceId

`func (o *O11yO11yEvent) HasTraceId() bool`

HasTraceId returns a boolean if a field has been set.

### GetTransaction

`func (o *O11yO11yEvent) GetTransaction() string`

GetTransaction returns the Transaction field if non-nil, zero value otherwise.

### GetTransactionOk

`func (o *O11yO11yEvent) GetTransactionOk() (*string, bool)`

GetTransactionOk returns a tuple with the Transaction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransaction

`func (o *O11yO11yEvent) SetTransaction(v string)`

SetTransaction sets Transaction field to given value.

### HasTransaction

`func (o *O11yO11yEvent) HasTransaction() bool`

HasTransaction returns a boolean if a field has been set.

### GetType

`func (o *O11yO11yEvent) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *O11yO11yEvent) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *O11yO11yEvent) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *O11yO11yEvent) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUserEmail

`func (o *O11yO11yEvent) GetUserEmail() string`

GetUserEmail returns the UserEmail field if non-nil, zero value otherwise.

### GetUserEmailOk

`func (o *O11yO11yEvent) GetUserEmailOk() (*string, bool)`

GetUserEmailOk returns a tuple with the UserEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserEmail

`func (o *O11yO11yEvent) SetUserEmail(v string)`

SetUserEmail sets UserEmail field to given value.

### HasUserEmail

`func (o *O11yO11yEvent) HasUserEmail() bool`

HasUserEmail returns a boolean if a field has been set.

### GetUserId

`func (o *O11yO11yEvent) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *O11yO11yEvent) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *O11yO11yEvent) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *O11yO11yEvent) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetUserIp

`func (o *O11yO11yEvent) GetUserIp() string`

GetUserIp returns the UserIp field if non-nil, zero value otherwise.

### GetUserIpOk

`func (o *O11yO11yEvent) GetUserIpOk() (*string, bool)`

GetUserIpOk returns a tuple with the UserIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserIp

`func (o *O11yO11yEvent) SetUserIp(v string)`

SetUserIp sets UserIp field to given value.

### HasUserIp

`func (o *O11yO11yEvent) HasUserIp() bool`

HasUserIp returns a boolean if a field has been set.

### GetValue

`func (o *O11yO11yEvent) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *O11yO11yEvent) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *O11yO11yEvent) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *O11yO11yEvent) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


