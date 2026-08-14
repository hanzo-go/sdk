# O11yO11yOccurrence

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Culprit** | Pointer to **string** | Culprit is where it came from. | [optional] 
**Environment** | Pointer to **string** | Environment is the deployment it happened in. | [optional] 
**EventId** | Pointer to **string** | EventID is the occurrence&#39;s id. | [optional] 
**Fingerprint** | Pointer to **string** | Fingerprint is the grouping key it was bucketed by. | [optional] 
**Frames** | Pointer to [**[]O11yO11yOccurrenceFrame**](O11yO11yOccurrenceFrame.md) | Frames are the stack, innermost first. | [optional] 
**Level** | Pointer to **string** | Level is its severity, e.g. error, warning, info. | [optional] 
**Platform** | Pointer to **string** | Platform is the reporting runtime. | [optional] 
**Release** | Pointer to **string** | Release is the version that produced it. | [optional] 
**ServerName** | Pointer to **string** | ServerName is the host that reported it. | [optional] 
**ServiceName** | Pointer to **string** | ServiceName is the service that reported it. | [optional] 
**SpanId** | Pointer to **string** | SpanID is the span it belonged to. | [optional] 
**Tags** | Pointer to **map[string]string** | Tags are the reporter&#39;s own key/value labels. | [optional] 
**Timestamp** | Pointer to **time.Time** | Timestamp is when the error happened. | [optional] 
**TraceId** | Pointer to **string** | TraceID is the trace it belonged to. | [optional] 
**Transaction** | Pointer to **string** | Transaction is the operation it happened in. | [optional] 
**Type** | Pointer to **string** | Type is the exception type. | [optional] 
**User** | Pointer to [**O11yO11yEventUser**](O11yO11yEventUser.md) | User is the affected end-user context, when the reporter attached one. | [optional] 
**Value** | Pointer to **string** | Value is the exception value. | [optional] 

## Methods

### NewO11yO11yOccurrence

`func NewO11yO11yOccurrence() *O11yO11yOccurrence`

NewO11yO11yOccurrence instantiates a new O11yO11yOccurrence object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yO11yOccurrenceWithDefaults

`func NewO11yO11yOccurrenceWithDefaults() *O11yO11yOccurrence`

NewO11yO11yOccurrenceWithDefaults instantiates a new O11yO11yOccurrence object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCulprit

`func (o *O11yO11yOccurrence) GetCulprit() string`

GetCulprit returns the Culprit field if non-nil, zero value otherwise.

### GetCulpritOk

`func (o *O11yO11yOccurrence) GetCulpritOk() (*string, bool)`

GetCulpritOk returns a tuple with the Culprit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCulprit

`func (o *O11yO11yOccurrence) SetCulprit(v string)`

SetCulprit sets Culprit field to given value.

### HasCulprit

`func (o *O11yO11yOccurrence) HasCulprit() bool`

HasCulprit returns a boolean if a field has been set.

### GetEnvironment

`func (o *O11yO11yOccurrence) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *O11yO11yOccurrence) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *O11yO11yOccurrence) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *O11yO11yOccurrence) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetEventId

`func (o *O11yO11yOccurrence) GetEventId() string`

GetEventId returns the EventId field if non-nil, zero value otherwise.

### GetEventIdOk

`func (o *O11yO11yOccurrence) GetEventIdOk() (*string, bool)`

GetEventIdOk returns a tuple with the EventId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventId

`func (o *O11yO11yOccurrence) SetEventId(v string)`

SetEventId sets EventId field to given value.

### HasEventId

`func (o *O11yO11yOccurrence) HasEventId() bool`

HasEventId returns a boolean if a field has been set.

### GetFingerprint

`func (o *O11yO11yOccurrence) GetFingerprint() string`

GetFingerprint returns the Fingerprint field if non-nil, zero value otherwise.

### GetFingerprintOk

`func (o *O11yO11yOccurrence) GetFingerprintOk() (*string, bool)`

GetFingerprintOk returns a tuple with the Fingerprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFingerprint

`func (o *O11yO11yOccurrence) SetFingerprint(v string)`

SetFingerprint sets Fingerprint field to given value.

### HasFingerprint

`func (o *O11yO11yOccurrence) HasFingerprint() bool`

HasFingerprint returns a boolean if a field has been set.

### GetFrames

`func (o *O11yO11yOccurrence) GetFrames() []O11yO11yOccurrenceFrame`

GetFrames returns the Frames field if non-nil, zero value otherwise.

### GetFramesOk

`func (o *O11yO11yOccurrence) GetFramesOk() (*[]O11yO11yOccurrenceFrame, bool)`

GetFramesOk returns a tuple with the Frames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrames

`func (o *O11yO11yOccurrence) SetFrames(v []O11yO11yOccurrenceFrame)`

SetFrames sets Frames field to given value.

### HasFrames

`func (o *O11yO11yOccurrence) HasFrames() bool`

HasFrames returns a boolean if a field has been set.

### GetLevel

`func (o *O11yO11yOccurrence) GetLevel() string`

GetLevel returns the Level field if non-nil, zero value otherwise.

### GetLevelOk

`func (o *O11yO11yOccurrence) GetLevelOk() (*string, bool)`

GetLevelOk returns a tuple with the Level field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel

`func (o *O11yO11yOccurrence) SetLevel(v string)`

SetLevel sets Level field to given value.

### HasLevel

`func (o *O11yO11yOccurrence) HasLevel() bool`

HasLevel returns a boolean if a field has been set.

### GetPlatform

`func (o *O11yO11yOccurrence) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *O11yO11yOccurrence) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *O11yO11yOccurrence) SetPlatform(v string)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *O11yO11yOccurrence) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### GetRelease

`func (o *O11yO11yOccurrence) GetRelease() string`

GetRelease returns the Release field if non-nil, zero value otherwise.

### GetReleaseOk

`func (o *O11yO11yOccurrence) GetReleaseOk() (*string, bool)`

GetReleaseOk returns a tuple with the Release field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelease

`func (o *O11yO11yOccurrence) SetRelease(v string)`

SetRelease sets Release field to given value.

### HasRelease

`func (o *O11yO11yOccurrence) HasRelease() bool`

HasRelease returns a boolean if a field has been set.

### GetServerName

`func (o *O11yO11yOccurrence) GetServerName() string`

GetServerName returns the ServerName field if non-nil, zero value otherwise.

### GetServerNameOk

`func (o *O11yO11yOccurrence) GetServerNameOk() (*string, bool)`

GetServerNameOk returns a tuple with the ServerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerName

`func (o *O11yO11yOccurrence) SetServerName(v string)`

SetServerName sets ServerName field to given value.

### HasServerName

`func (o *O11yO11yOccurrence) HasServerName() bool`

HasServerName returns a boolean if a field has been set.

### GetServiceName

`func (o *O11yO11yOccurrence) GetServiceName() string`

GetServiceName returns the ServiceName field if non-nil, zero value otherwise.

### GetServiceNameOk

`func (o *O11yO11yOccurrence) GetServiceNameOk() (*string, bool)`

GetServiceNameOk returns a tuple with the ServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceName

`func (o *O11yO11yOccurrence) SetServiceName(v string)`

SetServiceName sets ServiceName field to given value.

### HasServiceName

`func (o *O11yO11yOccurrence) HasServiceName() bool`

HasServiceName returns a boolean if a field has been set.

### GetSpanId

`func (o *O11yO11yOccurrence) GetSpanId() string`

GetSpanId returns the SpanId field if non-nil, zero value otherwise.

### GetSpanIdOk

`func (o *O11yO11yOccurrence) GetSpanIdOk() (*string, bool)`

GetSpanIdOk returns a tuple with the SpanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpanId

`func (o *O11yO11yOccurrence) SetSpanId(v string)`

SetSpanId sets SpanId field to given value.

### HasSpanId

`func (o *O11yO11yOccurrence) HasSpanId() bool`

HasSpanId returns a boolean if a field has been set.

### GetTags

`func (o *O11yO11yOccurrence) GetTags() map[string]string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *O11yO11yOccurrence) GetTagsOk() (*map[string]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *O11yO11yOccurrence) SetTags(v map[string]string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *O11yO11yOccurrence) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTimestamp

`func (o *O11yO11yOccurrence) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *O11yO11yOccurrence) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *O11yO11yOccurrence) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *O11yO11yOccurrence) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetTraceId

`func (o *O11yO11yOccurrence) GetTraceId() string`

GetTraceId returns the TraceId field if non-nil, zero value otherwise.

### GetTraceIdOk

`func (o *O11yO11yOccurrence) GetTraceIdOk() (*string, bool)`

GetTraceIdOk returns a tuple with the TraceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraceId

`func (o *O11yO11yOccurrence) SetTraceId(v string)`

SetTraceId sets TraceId field to given value.

### HasTraceId

`func (o *O11yO11yOccurrence) HasTraceId() bool`

HasTraceId returns a boolean if a field has been set.

### GetTransaction

`func (o *O11yO11yOccurrence) GetTransaction() string`

GetTransaction returns the Transaction field if non-nil, zero value otherwise.

### GetTransactionOk

`func (o *O11yO11yOccurrence) GetTransactionOk() (*string, bool)`

GetTransactionOk returns a tuple with the Transaction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransaction

`func (o *O11yO11yOccurrence) SetTransaction(v string)`

SetTransaction sets Transaction field to given value.

### HasTransaction

`func (o *O11yO11yOccurrence) HasTransaction() bool`

HasTransaction returns a boolean if a field has been set.

### GetType

`func (o *O11yO11yOccurrence) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *O11yO11yOccurrence) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *O11yO11yOccurrence) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *O11yO11yOccurrence) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUser

`func (o *O11yO11yOccurrence) GetUser() O11yO11yEventUser`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *O11yO11yOccurrence) GetUserOk() (*O11yO11yEventUser, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *O11yO11yOccurrence) SetUser(v O11yO11yEventUser)`

SetUser sets User field to given value.

### HasUser

`func (o *O11yO11yOccurrence) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetValue

`func (o *O11yO11yOccurrence) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *O11yO11yOccurrence) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *O11yO11yOccurrence) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *O11yO11yOccurrence) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


