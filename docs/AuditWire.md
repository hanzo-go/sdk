# AuditWire

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to **string** | Action is the verb that was performed. It is the event&#39;s name, not the HTTP method — a request-sourced record carries both, and the pair is what makes a row readable (\&quot;grant.create\&quot; at POST /v1/admin/grants). | [optional] 
**After** | Pointer to **interface{}** |  | [optional] 
**AuthMethod** | Pointer to **string** | Auth is the credential the actor presented: \&quot;jwt\&quot;, \&quot;api-key\&quot;, or \&quot;none\&quot;. | [optional] 
**Before** | Pointer to **interface{}** |  | [optional] 
**Email** | Pointer to **string** | Email is the actor&#39;s validated address, absent when the credential carried none. It comes from the verified token, never from a client header. | [optional] 
**Hash** | Pointer to **string** | Hash is this record&#39;s SHA-256 over its own canonical JSON with both hash fields cleared, folded with prevHash. Recomputing it from the row&#39;s other fields is what proves the row has not been edited. | [optional] 
**Home** | Pointer to **string** | Home is present ONLY on a cross-org action: the org the actor came FROM, while Org is the org they acted IN. A console row carrying &#x60;home&#x60; is a platform-admin impersonation and should be rendered as one. | [optional] 
**IsAdmin** | Pointer to **bool** | IsAdmin is the VALIDATED platform-SuperAdmin bit at decision time (owner &#x3D;&#x3D; \&quot;admin\&quot;, principal.Super), never the client&#39;s own claim to be one. | [optional] 
**Method** | Pointer to **string** | Method is the HTTP verb, on a record a request produced. Absent on an event emitted from inside the binary with no request behind it. | [optional] 
**Org** | Pointer to **string** | Org is the tenant the action was taken IN — the effective org, which for everyone but an impersonating SuperAdmin is also the actor&#39;s own. Empty on an unauthenticated request. | [optional] 
**Path** | Pointer to **string** | Path is the request&#39;s route. Any segment shaped like a credential is replaced before the record is written, so a key that rides in a path is not preserved here by the very control meant to watch it. | [optional] 
**PrevHash** | Pointer to **string** | PrevHash is the hash of record seq-1, which is what links the rows into a chain: a deleted or reordered record breaks the recomputation at that point. The first record of a chain carries 64 zeros rather than an empty string, so \&quot;start of chain\&quot; and \&quot;field missing\&quot; cannot look alike. | [optional] 
**Reason** | Pointer to **string** | Reason is a short explanation for a deny or an error (\&quot;SuperAdmin required\&quot;, \&quot;insufficient_balance\&quot;). It is never a secret and never a raw upstream error body; absent on a success. | [optional] 
**RequestId** | Pointer to **string** | RequestID ties this row to the request-line log and any downstream trace — the X-Request-Id the pipeline minted for that request. | [optional] 
**Resource** | Pointer to **string** | Resource is the KIND of thing acted upon (\&quot;org\&quot;, \&quot;role\&quot;, \&quot;secret\&quot;, \&quot;provider-config\&quot;, \&quot;credit\&quot;). Where a mutation has no finer semantics than its route, this is the route family and resourceId is empty — the action and the path already pin the object. | [optional] 
**ResourceId** | Pointer to **string** | ResourceID is the specific instance, absent when the kind alone identifies it. | [optional] 
**Result** | Pointer to **string** | Result is how the action ended: \&quot;success\&quot;, \&quot;deny\&quot; or \&quot;error\&quot;. A deny is a decision this binary made and is as much evidence as a success. | [optional] 
**Seq** | Pointer to **int32** | Seq is the record&#39;s position in the chain, 0-based and gapless. The Recorder assigns it under its own lock, so it is a true total order: seq n+1 was written after seq n, and a missing number is a missing record. | [optional] 
**SourceIp** | Pointer to **string** | SourceIP is the client address the edge resolved for the request, after the proxy chain — the address a responder would act on, not the socket peer. | [optional] 
**Status** | Pointer to **int64** | Status is the HTTP status the caller received. It is the outcome as the client saw it, so a 200 carrying a domain refusal still reads 200 here. | [optional] 
**Sub** | Pointer to **string** | Sub is the acting user (the IAM subject). Empty for a machine principal or an anonymous request, which is how a service action is told from a person&#39;s. | [optional] 
**Time** | Pointer to **string** | Time is when the action happened, RFC3339Nano in UTC. The stored column has the same precision and sorts the same way, so a client can range and order on this string verbatim. | [optional] 
**UserAgent** | Pointer to **string** | UserAgent is the client the request announced itself as. Client-supplied, so it is evidence about what claimed to act, not proof of it. | [optional] 

## Methods

### NewAuditWire

`func NewAuditWire() *AuditWire`

NewAuditWire instantiates a new AuditWire object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuditWireWithDefaults

`func NewAuditWireWithDefaults() *AuditWire`

NewAuditWireWithDefaults instantiates a new AuditWire object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *AuditWire) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *AuditWire) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *AuditWire) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *AuditWire) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetAfter

`func (o *AuditWire) GetAfter() interface{}`

GetAfter returns the After field if non-nil, zero value otherwise.

### GetAfterOk

`func (o *AuditWire) GetAfterOk() (*interface{}, bool)`

GetAfterOk returns a tuple with the After field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAfter

`func (o *AuditWire) SetAfter(v interface{})`

SetAfter sets After field to given value.

### HasAfter

`func (o *AuditWire) HasAfter() bool`

HasAfter returns a boolean if a field has been set.

### SetAfterNil

`func (o *AuditWire) SetAfterNil(b bool)`

 SetAfterNil sets the value for After to be an explicit nil

### UnsetAfter
`func (o *AuditWire) UnsetAfter()`

UnsetAfter ensures that no value is present for After, not even an explicit nil
### GetAuthMethod

`func (o *AuditWire) GetAuthMethod() string`

GetAuthMethod returns the AuthMethod field if non-nil, zero value otherwise.

### GetAuthMethodOk

`func (o *AuditWire) GetAuthMethodOk() (*string, bool)`

GetAuthMethodOk returns a tuple with the AuthMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthMethod

`func (o *AuditWire) SetAuthMethod(v string)`

SetAuthMethod sets AuthMethod field to given value.

### HasAuthMethod

`func (o *AuditWire) HasAuthMethod() bool`

HasAuthMethod returns a boolean if a field has been set.

### GetBefore

`func (o *AuditWire) GetBefore() interface{}`

GetBefore returns the Before field if non-nil, zero value otherwise.

### GetBeforeOk

`func (o *AuditWire) GetBeforeOk() (*interface{}, bool)`

GetBeforeOk returns a tuple with the Before field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBefore

`func (o *AuditWire) SetBefore(v interface{})`

SetBefore sets Before field to given value.

### HasBefore

`func (o *AuditWire) HasBefore() bool`

HasBefore returns a boolean if a field has been set.

### SetBeforeNil

`func (o *AuditWire) SetBeforeNil(b bool)`

 SetBeforeNil sets the value for Before to be an explicit nil

### UnsetBefore
`func (o *AuditWire) UnsetBefore()`

UnsetBefore ensures that no value is present for Before, not even an explicit nil
### GetEmail

`func (o *AuditWire) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *AuditWire) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *AuditWire) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *AuditWire) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetHash

`func (o *AuditWire) GetHash() string`

GetHash returns the Hash field if non-nil, zero value otherwise.

### GetHashOk

`func (o *AuditWire) GetHashOk() (*string, bool)`

GetHashOk returns a tuple with the Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHash

`func (o *AuditWire) SetHash(v string)`

SetHash sets Hash field to given value.

### HasHash

`func (o *AuditWire) HasHash() bool`

HasHash returns a boolean if a field has been set.

### GetHome

`func (o *AuditWire) GetHome() string`

GetHome returns the Home field if non-nil, zero value otherwise.

### GetHomeOk

`func (o *AuditWire) GetHomeOk() (*string, bool)`

GetHomeOk returns a tuple with the Home field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHome

`func (o *AuditWire) SetHome(v string)`

SetHome sets Home field to given value.

### HasHome

`func (o *AuditWire) HasHome() bool`

HasHome returns a boolean if a field has been set.

### GetIsAdmin

`func (o *AuditWire) GetIsAdmin() bool`

GetIsAdmin returns the IsAdmin field if non-nil, zero value otherwise.

### GetIsAdminOk

`func (o *AuditWire) GetIsAdminOk() (*bool, bool)`

GetIsAdminOk returns a tuple with the IsAdmin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAdmin

`func (o *AuditWire) SetIsAdmin(v bool)`

SetIsAdmin sets IsAdmin field to given value.

### HasIsAdmin

`func (o *AuditWire) HasIsAdmin() bool`

HasIsAdmin returns a boolean if a field has been set.

### GetMethod

`func (o *AuditWire) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *AuditWire) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *AuditWire) SetMethod(v string)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *AuditWire) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### GetOrg

`func (o *AuditWire) GetOrg() string`

GetOrg returns the Org field if non-nil, zero value otherwise.

### GetOrgOk

`func (o *AuditWire) GetOrgOk() (*string, bool)`

GetOrgOk returns a tuple with the Org field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrg

`func (o *AuditWire) SetOrg(v string)`

SetOrg sets Org field to given value.

### HasOrg

`func (o *AuditWire) HasOrg() bool`

HasOrg returns a boolean if a field has been set.

### GetPath

`func (o *AuditWire) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *AuditWire) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *AuditWire) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *AuditWire) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetPrevHash

`func (o *AuditWire) GetPrevHash() string`

GetPrevHash returns the PrevHash field if non-nil, zero value otherwise.

### GetPrevHashOk

`func (o *AuditWire) GetPrevHashOk() (*string, bool)`

GetPrevHashOk returns a tuple with the PrevHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrevHash

`func (o *AuditWire) SetPrevHash(v string)`

SetPrevHash sets PrevHash field to given value.

### HasPrevHash

`func (o *AuditWire) HasPrevHash() bool`

HasPrevHash returns a boolean if a field has been set.

### GetReason

`func (o *AuditWire) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *AuditWire) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *AuditWire) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *AuditWire) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetRequestId

`func (o *AuditWire) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *AuditWire) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *AuditWire) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *AuditWire) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.

### GetResource

`func (o *AuditWire) GetResource() string`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *AuditWire) GetResourceOk() (*string, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *AuditWire) SetResource(v string)`

SetResource sets Resource field to given value.

### HasResource

`func (o *AuditWire) HasResource() bool`

HasResource returns a boolean if a field has been set.

### GetResourceId

`func (o *AuditWire) GetResourceId() string`

GetResourceId returns the ResourceId field if non-nil, zero value otherwise.

### GetResourceIdOk

`func (o *AuditWire) GetResourceIdOk() (*string, bool)`

GetResourceIdOk returns a tuple with the ResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceId

`func (o *AuditWire) SetResourceId(v string)`

SetResourceId sets ResourceId field to given value.

### HasResourceId

`func (o *AuditWire) HasResourceId() bool`

HasResourceId returns a boolean if a field has been set.

### GetResult

`func (o *AuditWire) GetResult() string`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *AuditWire) GetResultOk() (*string, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *AuditWire) SetResult(v string)`

SetResult sets Result field to given value.

### HasResult

`func (o *AuditWire) HasResult() bool`

HasResult returns a boolean if a field has been set.

### GetSeq

`func (o *AuditWire) GetSeq() int32`

GetSeq returns the Seq field if non-nil, zero value otherwise.

### GetSeqOk

`func (o *AuditWire) GetSeqOk() (*int32, bool)`

GetSeqOk returns a tuple with the Seq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeq

`func (o *AuditWire) SetSeq(v int32)`

SetSeq sets Seq field to given value.

### HasSeq

`func (o *AuditWire) HasSeq() bool`

HasSeq returns a boolean if a field has been set.

### GetSourceIp

`func (o *AuditWire) GetSourceIp() string`

GetSourceIp returns the SourceIp field if non-nil, zero value otherwise.

### GetSourceIpOk

`func (o *AuditWire) GetSourceIpOk() (*string, bool)`

GetSourceIpOk returns a tuple with the SourceIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceIp

`func (o *AuditWire) SetSourceIp(v string)`

SetSourceIp sets SourceIp field to given value.

### HasSourceIp

`func (o *AuditWire) HasSourceIp() bool`

HasSourceIp returns a boolean if a field has been set.

### GetStatus

`func (o *AuditWire) GetStatus() int64`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AuditWire) GetStatusOk() (*int64, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AuditWire) SetStatus(v int64)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AuditWire) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSub

`func (o *AuditWire) GetSub() string`

GetSub returns the Sub field if non-nil, zero value otherwise.

### GetSubOk

`func (o *AuditWire) GetSubOk() (*string, bool)`

GetSubOk returns a tuple with the Sub field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSub

`func (o *AuditWire) SetSub(v string)`

SetSub sets Sub field to given value.

### HasSub

`func (o *AuditWire) HasSub() bool`

HasSub returns a boolean if a field has been set.

### GetTime

`func (o *AuditWire) GetTime() string`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *AuditWire) GetTimeOk() (*string, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *AuditWire) SetTime(v string)`

SetTime sets Time field to given value.

### HasTime

`func (o *AuditWire) HasTime() bool`

HasTime returns a boolean if a field has been set.

### GetUserAgent

`func (o *AuditWire) GetUserAgent() string`

GetUserAgent returns the UserAgent field if non-nil, zero value otherwise.

### GetUserAgentOk

`func (o *AuditWire) GetUserAgentOk() (*string, bool)`

GetUserAgentOk returns a tuple with the UserAgent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAgent

`func (o *AuditWire) SetUserAgent(v string)`

SetUserAgent sets UserAgent field to given value.

### HasUserAgent

`func (o *AuditWire) HasUserAgent() bool`

HasUserAgent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


