# AdminAuditRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Seq** | Pointer to **int64** |  | [optional] 
**Time** | Pointer to **string** |  | [optional] 
**Org** | Pointer to **string** |  | [optional] 
**Sub** | Pointer to **string** |  | [optional] 
**Email** | Pointer to **string** |  | [optional] 
**Action** | Pointer to **string** |  | [optional] 
**Resource** | Pointer to **string** |  | [optional] 
**ResourceId** | Pointer to **string** |  | [optional] 
**Method** | Pointer to **string** |  | [optional] 
**Path** | Pointer to **string** |  | [optional] 
**Result** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **int32** |  | [optional] 
**Reason** | Pointer to **string** |  | [optional] 
**SourceIp** | Pointer to **string** |  | [optional] 
**UserAgent** | Pointer to **string** |  | [optional] 
**RequestId** | Pointer to **string** |  | [optional] 
**IsAdmin** | Pointer to **bool** |  | [optional] 
**AuthMethod** | Pointer to **string** |  | [optional] 
**Hash** | Pointer to **string** |  | [optional] 
**PrevHash** | Pointer to **string** |  | [optional] 

## Methods

### NewAdminAuditRow

`func NewAdminAuditRow() *AdminAuditRow`

NewAdminAuditRow instantiates a new AdminAuditRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdminAuditRowWithDefaults

`func NewAdminAuditRowWithDefaults() *AdminAuditRow`

NewAdminAuditRowWithDefaults instantiates a new AdminAuditRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSeq

`func (o *AdminAuditRow) GetSeq() int64`

GetSeq returns the Seq field if non-nil, zero value otherwise.

### GetSeqOk

`func (o *AdminAuditRow) GetSeqOk() (*int64, bool)`

GetSeqOk returns a tuple with the Seq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeq

`func (o *AdminAuditRow) SetSeq(v int64)`

SetSeq sets Seq field to given value.

### HasSeq

`func (o *AdminAuditRow) HasSeq() bool`

HasSeq returns a boolean if a field has been set.

### GetTime

`func (o *AdminAuditRow) GetTime() string`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *AdminAuditRow) GetTimeOk() (*string, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *AdminAuditRow) SetTime(v string)`

SetTime sets Time field to given value.

### HasTime

`func (o *AdminAuditRow) HasTime() bool`

HasTime returns a boolean if a field has been set.

### GetOrg

`func (o *AdminAuditRow) GetOrg() string`

GetOrg returns the Org field if non-nil, zero value otherwise.

### GetOrgOk

`func (o *AdminAuditRow) GetOrgOk() (*string, bool)`

GetOrgOk returns a tuple with the Org field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrg

`func (o *AdminAuditRow) SetOrg(v string)`

SetOrg sets Org field to given value.

### HasOrg

`func (o *AdminAuditRow) HasOrg() bool`

HasOrg returns a boolean if a field has been set.

### GetSub

`func (o *AdminAuditRow) GetSub() string`

GetSub returns the Sub field if non-nil, zero value otherwise.

### GetSubOk

`func (o *AdminAuditRow) GetSubOk() (*string, bool)`

GetSubOk returns a tuple with the Sub field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSub

`func (o *AdminAuditRow) SetSub(v string)`

SetSub sets Sub field to given value.

### HasSub

`func (o *AdminAuditRow) HasSub() bool`

HasSub returns a boolean if a field has been set.

### GetEmail

`func (o *AdminAuditRow) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *AdminAuditRow) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *AdminAuditRow) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *AdminAuditRow) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetAction

`func (o *AdminAuditRow) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *AdminAuditRow) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *AdminAuditRow) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *AdminAuditRow) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetResource

`func (o *AdminAuditRow) GetResource() string`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *AdminAuditRow) GetResourceOk() (*string, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *AdminAuditRow) SetResource(v string)`

SetResource sets Resource field to given value.

### HasResource

`func (o *AdminAuditRow) HasResource() bool`

HasResource returns a boolean if a field has been set.

### GetResourceId

`func (o *AdminAuditRow) GetResourceId() string`

GetResourceId returns the ResourceId field if non-nil, zero value otherwise.

### GetResourceIdOk

`func (o *AdminAuditRow) GetResourceIdOk() (*string, bool)`

GetResourceIdOk returns a tuple with the ResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceId

`func (o *AdminAuditRow) SetResourceId(v string)`

SetResourceId sets ResourceId field to given value.

### HasResourceId

`func (o *AdminAuditRow) HasResourceId() bool`

HasResourceId returns a boolean if a field has been set.

### GetMethod

`func (o *AdminAuditRow) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *AdminAuditRow) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *AdminAuditRow) SetMethod(v string)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *AdminAuditRow) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### GetPath

`func (o *AdminAuditRow) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *AdminAuditRow) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *AdminAuditRow) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *AdminAuditRow) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetResult

`func (o *AdminAuditRow) GetResult() string`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *AdminAuditRow) GetResultOk() (*string, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *AdminAuditRow) SetResult(v string)`

SetResult sets Result field to given value.

### HasResult

`func (o *AdminAuditRow) HasResult() bool`

HasResult returns a boolean if a field has been set.

### GetStatus

`func (o *AdminAuditRow) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AdminAuditRow) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AdminAuditRow) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AdminAuditRow) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetReason

`func (o *AdminAuditRow) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *AdminAuditRow) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *AdminAuditRow) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *AdminAuditRow) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetSourceIp

`func (o *AdminAuditRow) GetSourceIp() string`

GetSourceIp returns the SourceIp field if non-nil, zero value otherwise.

### GetSourceIpOk

`func (o *AdminAuditRow) GetSourceIpOk() (*string, bool)`

GetSourceIpOk returns a tuple with the SourceIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceIp

`func (o *AdminAuditRow) SetSourceIp(v string)`

SetSourceIp sets SourceIp field to given value.

### HasSourceIp

`func (o *AdminAuditRow) HasSourceIp() bool`

HasSourceIp returns a boolean if a field has been set.

### GetUserAgent

`func (o *AdminAuditRow) GetUserAgent() string`

GetUserAgent returns the UserAgent field if non-nil, zero value otherwise.

### GetUserAgentOk

`func (o *AdminAuditRow) GetUserAgentOk() (*string, bool)`

GetUserAgentOk returns a tuple with the UserAgent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAgent

`func (o *AdminAuditRow) SetUserAgent(v string)`

SetUserAgent sets UserAgent field to given value.

### HasUserAgent

`func (o *AdminAuditRow) HasUserAgent() bool`

HasUserAgent returns a boolean if a field has been set.

### GetRequestId

`func (o *AdminAuditRow) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *AdminAuditRow) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *AdminAuditRow) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *AdminAuditRow) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.

### GetIsAdmin

`func (o *AdminAuditRow) GetIsAdmin() bool`

GetIsAdmin returns the IsAdmin field if non-nil, zero value otherwise.

### GetIsAdminOk

`func (o *AdminAuditRow) GetIsAdminOk() (*bool, bool)`

GetIsAdminOk returns a tuple with the IsAdmin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAdmin

`func (o *AdminAuditRow) SetIsAdmin(v bool)`

SetIsAdmin sets IsAdmin field to given value.

### HasIsAdmin

`func (o *AdminAuditRow) HasIsAdmin() bool`

HasIsAdmin returns a boolean if a field has been set.

### GetAuthMethod

`func (o *AdminAuditRow) GetAuthMethod() string`

GetAuthMethod returns the AuthMethod field if non-nil, zero value otherwise.

### GetAuthMethodOk

`func (o *AdminAuditRow) GetAuthMethodOk() (*string, bool)`

GetAuthMethodOk returns a tuple with the AuthMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthMethod

`func (o *AdminAuditRow) SetAuthMethod(v string)`

SetAuthMethod sets AuthMethod field to given value.

### HasAuthMethod

`func (o *AdminAuditRow) HasAuthMethod() bool`

HasAuthMethod returns a boolean if a field has been set.

### GetHash

`func (o *AdminAuditRow) GetHash() string`

GetHash returns the Hash field if non-nil, zero value otherwise.

### GetHashOk

`func (o *AdminAuditRow) GetHashOk() (*string, bool)`

GetHashOk returns a tuple with the Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHash

`func (o *AdminAuditRow) SetHash(v string)`

SetHash sets Hash field to given value.

### HasHash

`func (o *AdminAuditRow) HasHash() bool`

HasHash returns a boolean if a field has been set.

### GetPrevHash

`func (o *AdminAuditRow) GetPrevHash() string`

GetPrevHash returns the PrevHash field if non-nil, zero value otherwise.

### GetPrevHashOk

`func (o *AdminAuditRow) GetPrevHashOk() (*string, bool)`

GetPrevHashOk returns a tuple with the PrevHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrevHash

`func (o *AdminAuditRow) SetPrevHash(v string)`

SetPrevHash sets PrevHash field to given value.

### HasPrevHash

`func (o *AdminAuditRow) HasPrevHash() bool`

HasPrevHash returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


