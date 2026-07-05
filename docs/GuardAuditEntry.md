# GuardAuditEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Timestamp** | Pointer to **time.Time** |  | [optional] 
**UserId** | Pointer to **string** |  | [optional] 
**SessionId** | Pointer to **string** |  | [optional] 
**SourceIp** | Pointer to **string** |  | [optional] 
**Direction** | Pointer to **string** |  | [optional] 
**Result** | Pointer to **string** |  | [optional] 
**RedactionCount** | Pointer to **int32** |  | [optional] 
**InjectionDetected** | Pointer to **bool** |  | [optional] 
**InjectionConfidence** | Pointer to **float32** |  | [optional] 
**ContentCategory** | Pointer to **string** |  | [optional] 
**ContentHash** | Pointer to **string** | SHA-256 hash of original content | [optional] 
**ProcessingTimeUs** | Pointer to **int32** |  | [optional] 

## Methods

### NewGuardAuditEntry

`func NewGuardAuditEntry() *GuardAuditEntry`

NewGuardAuditEntry instantiates a new GuardAuditEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGuardAuditEntryWithDefaults

`func NewGuardAuditEntryWithDefaults() *GuardAuditEntry`

NewGuardAuditEntryWithDefaults instantiates a new GuardAuditEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GuardAuditEntry) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GuardAuditEntry) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GuardAuditEntry) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GuardAuditEntry) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTimestamp

`func (o *GuardAuditEntry) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *GuardAuditEntry) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *GuardAuditEntry) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *GuardAuditEntry) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetUserId

`func (o *GuardAuditEntry) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *GuardAuditEntry) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *GuardAuditEntry) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *GuardAuditEntry) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetSessionId

`func (o *GuardAuditEntry) GetSessionId() string`

GetSessionId returns the SessionId field if non-nil, zero value otherwise.

### GetSessionIdOk

`func (o *GuardAuditEntry) GetSessionIdOk() (*string, bool)`

GetSessionIdOk returns a tuple with the SessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionId

`func (o *GuardAuditEntry) SetSessionId(v string)`

SetSessionId sets SessionId field to given value.

### HasSessionId

`func (o *GuardAuditEntry) HasSessionId() bool`

HasSessionId returns a boolean if a field has been set.

### GetSourceIp

`func (o *GuardAuditEntry) GetSourceIp() string`

GetSourceIp returns the SourceIp field if non-nil, zero value otherwise.

### GetSourceIpOk

`func (o *GuardAuditEntry) GetSourceIpOk() (*string, bool)`

GetSourceIpOk returns a tuple with the SourceIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceIp

`func (o *GuardAuditEntry) SetSourceIp(v string)`

SetSourceIp sets SourceIp field to given value.

### HasSourceIp

`func (o *GuardAuditEntry) HasSourceIp() bool`

HasSourceIp returns a boolean if a field has been set.

### GetDirection

`func (o *GuardAuditEntry) GetDirection() string`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *GuardAuditEntry) GetDirectionOk() (*string, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *GuardAuditEntry) SetDirection(v string)`

SetDirection sets Direction field to given value.

### HasDirection

`func (o *GuardAuditEntry) HasDirection() bool`

HasDirection returns a boolean if a field has been set.

### GetResult

`func (o *GuardAuditEntry) GetResult() string`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *GuardAuditEntry) GetResultOk() (*string, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *GuardAuditEntry) SetResult(v string)`

SetResult sets Result field to given value.

### HasResult

`func (o *GuardAuditEntry) HasResult() bool`

HasResult returns a boolean if a field has been set.

### GetRedactionCount

`func (o *GuardAuditEntry) GetRedactionCount() int32`

GetRedactionCount returns the RedactionCount field if non-nil, zero value otherwise.

### GetRedactionCountOk

`func (o *GuardAuditEntry) GetRedactionCountOk() (*int32, bool)`

GetRedactionCountOk returns a tuple with the RedactionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedactionCount

`func (o *GuardAuditEntry) SetRedactionCount(v int32)`

SetRedactionCount sets RedactionCount field to given value.

### HasRedactionCount

`func (o *GuardAuditEntry) HasRedactionCount() bool`

HasRedactionCount returns a boolean if a field has been set.

### GetInjectionDetected

`func (o *GuardAuditEntry) GetInjectionDetected() bool`

GetInjectionDetected returns the InjectionDetected field if non-nil, zero value otherwise.

### GetInjectionDetectedOk

`func (o *GuardAuditEntry) GetInjectionDetectedOk() (*bool, bool)`

GetInjectionDetectedOk returns a tuple with the InjectionDetected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInjectionDetected

`func (o *GuardAuditEntry) SetInjectionDetected(v bool)`

SetInjectionDetected sets InjectionDetected field to given value.

### HasInjectionDetected

`func (o *GuardAuditEntry) HasInjectionDetected() bool`

HasInjectionDetected returns a boolean if a field has been set.

### GetInjectionConfidence

`func (o *GuardAuditEntry) GetInjectionConfidence() float32`

GetInjectionConfidence returns the InjectionConfidence field if non-nil, zero value otherwise.

### GetInjectionConfidenceOk

`func (o *GuardAuditEntry) GetInjectionConfidenceOk() (*float32, bool)`

GetInjectionConfidenceOk returns a tuple with the InjectionConfidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInjectionConfidence

`func (o *GuardAuditEntry) SetInjectionConfidence(v float32)`

SetInjectionConfidence sets InjectionConfidence field to given value.

### HasInjectionConfidence

`func (o *GuardAuditEntry) HasInjectionConfidence() bool`

HasInjectionConfidence returns a boolean if a field has been set.

### GetContentCategory

`func (o *GuardAuditEntry) GetContentCategory() string`

GetContentCategory returns the ContentCategory field if non-nil, zero value otherwise.

### GetContentCategoryOk

`func (o *GuardAuditEntry) GetContentCategoryOk() (*string, bool)`

GetContentCategoryOk returns a tuple with the ContentCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentCategory

`func (o *GuardAuditEntry) SetContentCategory(v string)`

SetContentCategory sets ContentCategory field to given value.

### HasContentCategory

`func (o *GuardAuditEntry) HasContentCategory() bool`

HasContentCategory returns a boolean if a field has been set.

### GetContentHash

`func (o *GuardAuditEntry) GetContentHash() string`

GetContentHash returns the ContentHash field if non-nil, zero value otherwise.

### GetContentHashOk

`func (o *GuardAuditEntry) GetContentHashOk() (*string, bool)`

GetContentHashOk returns a tuple with the ContentHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentHash

`func (o *GuardAuditEntry) SetContentHash(v string)`

SetContentHash sets ContentHash field to given value.

### HasContentHash

`func (o *GuardAuditEntry) HasContentHash() bool`

HasContentHash returns a boolean if a field has been set.

### GetProcessingTimeUs

`func (o *GuardAuditEntry) GetProcessingTimeUs() int32`

GetProcessingTimeUs returns the ProcessingTimeUs field if non-nil, zero value otherwise.

### GetProcessingTimeUsOk

`func (o *GuardAuditEntry) GetProcessingTimeUsOk() (*int32, bool)`

GetProcessingTimeUsOk returns a tuple with the ProcessingTimeUs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessingTimeUs

`func (o *GuardAuditEntry) SetProcessingTimeUs(v int32)`

SetProcessingTimeUs sets ProcessingTimeUs field to given value.

### HasProcessingTimeUs

`func (o *GuardAuditEntry) HasProcessingTimeUs() bool`

HasProcessingTimeUs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


