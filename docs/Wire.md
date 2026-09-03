# Wire

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to **string** | Action is the verb that was performed. It is the event&#39;s name, not the HTTP method — a request-sourced record carries both, and the pair is what makes a row readable (\&quot;grant.create\&quot; at POST /v1/admin/grants). | [optional] 
**AuthMethod** | Pointer to **string** | Auth is the credential the actor presented: \&quot;jwt\&quot;, \&quot;api-key\&quot;, or \&quot;none\&quot;. | [optional] 
**Email** | Pointer to **string** | Email is the actor&#39;s validated address, absent when the credential carried none. It comes from the verified token, never from a client header. | [optional] 
**Hash** | Pointer to **string** | Hash is this record&#39;s SHA-256 over its own canonical JSON with both hash fields cleared, folded with prevHash. Recomputing it from the row&#39;s other fields is what proves the row has not been edited. | [optional] 
**Home** | Pointer to **string** | Home is present ONLY on a cross-org action: the org the actor came FROM, while Org is the org they acted IN. A console row carrying &#x60;home&#x60; is a platform-admin impersonation and should be rendered as one. | [optional] 
**IsAdmin** | Pointer to **bool** | IsAdmin is the VALIDATED platform-SuperAdmin bit at decision time (membership of the reserved admin org), never the client&#39;s own claim to be one. | [optional] 
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

### NewWire

`func NewWire() *Wire`

NewWire instantiates a new Wire object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWireWithDefaults

`func NewWireWithDefaults() *Wire`

NewWireWithDefaults instantiates a new Wire object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *Wire) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *Wire) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *Wire) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *Wire) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetAuthMethod

`func (o *Wire) GetAuthMethod() string`

GetAuthMethod returns the AuthMethod field if non-nil, zero value otherwise.

### GetAuthMethodOk

`func (o *Wire) GetAuthMethodOk() (*string, bool)`

GetAuthMethodOk returns a tuple with the AuthMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthMethod

`func (o *Wire) SetAuthMethod(v string)`

SetAuthMethod sets AuthMethod field to given value.

### HasAuthMethod

`func (o *Wire) HasAuthMethod() bool`

HasAuthMethod returns a boolean if a field has been set.

### GetEmail

`func (o *Wire) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *Wire) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *Wire) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *Wire) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetHash

`func (o *Wire) GetHash() string`

GetHash returns the Hash field if non-nil, zero value otherwise.

### GetHashOk

`func (o *Wire) GetHashOk() (*string, bool)`

GetHashOk returns a tuple with the Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHash

`func (o *Wire) SetHash(v string)`

SetHash sets Hash field to given value.

### HasHash

`func (o *Wire) HasHash() bool`

HasHash returns a boolean if a field has been set.

### GetHome

`func (o *Wire) GetHome() string`

GetHome returns the Home field if non-nil, zero value otherwise.

### GetHomeOk

`func (o *Wire) GetHomeOk() (*string, bool)`

GetHomeOk returns a tuple with the Home field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHome

`func (o *Wire) SetHome(v string)`

SetHome sets Home field to given value.

### HasHome

`func (o *Wire) HasHome() bool`

HasHome returns a boolean if a field has been set.

### GetIsAdmin

`func (o *Wire) GetIsAdmin() bool`

GetIsAdmin returns the IsAdmin field if non-nil, zero value otherwise.

### GetIsAdminOk

`func (o *Wire) GetIsAdminOk() (*bool, bool)`

GetIsAdminOk returns a tuple with the IsAdmin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAdmin

`func (o *Wire) SetIsAdmin(v bool)`

SetIsAdmin sets IsAdmin field to given value.

### HasIsAdmin

`func (o *Wire) HasIsAdmin() bool`

HasIsAdmin returns a boolean if a field has been set.

### GetMethod

`func (o *Wire) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *Wire) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *Wire) SetMethod(v string)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *Wire) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### GetOrg

`func (o *Wire) GetOrg() string`

GetOrg returns the Org field if non-nil, zero value otherwise.

### GetOrgOk

`func (o *Wire) GetOrgOk() (*string, bool)`

GetOrgOk returns a tuple with the Org field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrg

`func (o *Wire) SetOrg(v string)`

SetOrg sets Org field to given value.

### HasOrg

`func (o *Wire) HasOrg() bool`

HasOrg returns a boolean if a field has been set.

### GetPath

`func (o *Wire) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *Wire) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *Wire) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *Wire) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetPrevHash

`func (o *Wire) GetPrevHash() string`

GetPrevHash returns the PrevHash field if non-nil, zero value otherwise.

### GetPrevHashOk

`func (o *Wire) GetPrevHashOk() (*string, bool)`

GetPrevHashOk returns a tuple with the PrevHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrevHash

`func (o *Wire) SetPrevHash(v string)`

SetPrevHash sets PrevHash field to given value.

### HasPrevHash

`func (o *Wire) HasPrevHash() bool`

HasPrevHash returns a boolean if a field has been set.

### GetReason

`func (o *Wire) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *Wire) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *Wire) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *Wire) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetRequestId

`func (o *Wire) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *Wire) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *Wire) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *Wire) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.

### GetResource

`func (o *Wire) GetResource() string`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *Wire) GetResourceOk() (*string, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *Wire) SetResource(v string)`

SetResource sets Resource field to given value.

### HasResource

`func (o *Wire) HasResource() bool`

HasResource returns a boolean if a field has been set.

### GetResourceId

`func (o *Wire) GetResourceId() string`

GetResourceId returns the ResourceId field if non-nil, zero value otherwise.

### GetResourceIdOk

`func (o *Wire) GetResourceIdOk() (*string, bool)`

GetResourceIdOk returns a tuple with the ResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceId

`func (o *Wire) SetResourceId(v string)`

SetResourceId sets ResourceId field to given value.

### HasResourceId

`func (o *Wire) HasResourceId() bool`

HasResourceId returns a boolean if a field has been set.

### GetResult

`func (o *Wire) GetResult() string`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *Wire) GetResultOk() (*string, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *Wire) SetResult(v string)`

SetResult sets Result field to given value.

### HasResult

`func (o *Wire) HasResult() bool`

HasResult returns a boolean if a field has been set.

### GetSeq

`func (o *Wire) GetSeq() int32`

GetSeq returns the Seq field if non-nil, zero value otherwise.

### GetSeqOk

`func (o *Wire) GetSeqOk() (*int32, bool)`

GetSeqOk returns a tuple with the Seq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeq

`func (o *Wire) SetSeq(v int32)`

SetSeq sets Seq field to given value.

### HasSeq

`func (o *Wire) HasSeq() bool`

HasSeq returns a boolean if a field has been set.

### GetSourceIp

`func (o *Wire) GetSourceIp() string`

GetSourceIp returns the SourceIp field if non-nil, zero value otherwise.

### GetSourceIpOk

`func (o *Wire) GetSourceIpOk() (*string, bool)`

GetSourceIpOk returns a tuple with the SourceIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceIp

`func (o *Wire) SetSourceIp(v string)`

SetSourceIp sets SourceIp field to given value.

### HasSourceIp

`func (o *Wire) HasSourceIp() bool`

HasSourceIp returns a boolean if a field has been set.

### GetStatus

`func (o *Wire) GetStatus() int64`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Wire) GetStatusOk() (*int64, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Wire) SetStatus(v int64)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Wire) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSub

`func (o *Wire) GetSub() string`

GetSub returns the Sub field if non-nil, zero value otherwise.

### GetSubOk

`func (o *Wire) GetSubOk() (*string, bool)`

GetSubOk returns a tuple with the Sub field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSub

`func (o *Wire) SetSub(v string)`

SetSub sets Sub field to given value.

### HasSub

`func (o *Wire) HasSub() bool`

HasSub returns a boolean if a field has been set.

### GetTime

`func (o *Wire) GetTime() string`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *Wire) GetTimeOk() (*string, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *Wire) SetTime(v string)`

SetTime sets Time field to given value.

### HasTime

`func (o *Wire) HasTime() bool`

HasTime returns a boolean if a field has been set.

### GetUserAgent

`func (o *Wire) GetUserAgent() string`

GetUserAgent returns the UserAgent field if non-nil, zero value otherwise.

### GetUserAgentOk

`func (o *Wire) GetUserAgentOk() (*string, bool)`

GetUserAgentOk returns a tuple with the UserAgent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAgent

`func (o *Wire) SetUserAgent(v string)`

SetUserAgent sets UserAgent field to given value.

### HasUserAgent

`func (o *Wire) HasUserAgent() bool`

HasUserAgent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


