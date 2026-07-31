# CloudWire

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to **string** |  | [optional] 
**AuthMethod** | Pointer to **string** |  | [optional] 
**Email** | Pointer to **string** |  | [optional] 
**Hash** | Pointer to **string** |  | [optional] 
**Home** | Pointer to **string** | Home is present ONLY on a cross-org action: the org the actor came FROM, while Org is the org they acted IN. A console row carrying &#x60;home&#x60; is a platform-admin impersonation and should be rendered as one. | [optional] 
**IsAdmin** | Pointer to **bool** |  | [optional] 
**Method** | Pointer to **string** |  | [optional] 
**Org** | Pointer to **string** |  | [optional] 
**Path** | Pointer to **string** |  | [optional] 
**PrevHash** | Pointer to **string** |  | [optional] 
**Reason** | Pointer to **string** |  | [optional] 
**RequestId** | Pointer to **string** |  | [optional] 
**Resource** | Pointer to **string** |  | [optional] 
**ResourceId** | Pointer to **string** |  | [optional] 
**Result** | Pointer to **string** |  | [optional] 
**Seq** | Pointer to **int32** |  | [optional] 
**SourceIp** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **int32** |  | [optional] 
**Sub** | Pointer to **string** |  | [optional] 
**Time** | Pointer to **string** |  | [optional] 
**UserAgent** | Pointer to **string** |  | [optional] 

## Methods

### NewCloudWire

`func NewCloudWire() *CloudWire`

NewCloudWire instantiates a new CloudWire object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudWireWithDefaults

`func NewCloudWireWithDefaults() *CloudWire`

NewCloudWireWithDefaults instantiates a new CloudWire object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *CloudWire) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *CloudWire) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *CloudWire) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *CloudWire) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetAuthMethod

`func (o *CloudWire) GetAuthMethod() string`

GetAuthMethod returns the AuthMethod field if non-nil, zero value otherwise.

### GetAuthMethodOk

`func (o *CloudWire) GetAuthMethodOk() (*string, bool)`

GetAuthMethodOk returns a tuple with the AuthMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthMethod

`func (o *CloudWire) SetAuthMethod(v string)`

SetAuthMethod sets AuthMethod field to given value.

### HasAuthMethod

`func (o *CloudWire) HasAuthMethod() bool`

HasAuthMethod returns a boolean if a field has been set.

### GetEmail

`func (o *CloudWire) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *CloudWire) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *CloudWire) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *CloudWire) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetHash

`func (o *CloudWire) GetHash() string`

GetHash returns the Hash field if non-nil, zero value otherwise.

### GetHashOk

`func (o *CloudWire) GetHashOk() (*string, bool)`

GetHashOk returns a tuple with the Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHash

`func (o *CloudWire) SetHash(v string)`

SetHash sets Hash field to given value.

### HasHash

`func (o *CloudWire) HasHash() bool`

HasHash returns a boolean if a field has been set.

### GetHome

`func (o *CloudWire) GetHome() string`

GetHome returns the Home field if non-nil, zero value otherwise.

### GetHomeOk

`func (o *CloudWire) GetHomeOk() (*string, bool)`

GetHomeOk returns a tuple with the Home field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHome

`func (o *CloudWire) SetHome(v string)`

SetHome sets Home field to given value.

### HasHome

`func (o *CloudWire) HasHome() bool`

HasHome returns a boolean if a field has been set.

### GetIsAdmin

`func (o *CloudWire) GetIsAdmin() bool`

GetIsAdmin returns the IsAdmin field if non-nil, zero value otherwise.

### GetIsAdminOk

`func (o *CloudWire) GetIsAdminOk() (*bool, bool)`

GetIsAdminOk returns a tuple with the IsAdmin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAdmin

`func (o *CloudWire) SetIsAdmin(v bool)`

SetIsAdmin sets IsAdmin field to given value.

### HasIsAdmin

`func (o *CloudWire) HasIsAdmin() bool`

HasIsAdmin returns a boolean if a field has been set.

### GetMethod

`func (o *CloudWire) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *CloudWire) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *CloudWire) SetMethod(v string)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *CloudWire) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### GetOrg

`func (o *CloudWire) GetOrg() string`

GetOrg returns the Org field if non-nil, zero value otherwise.

### GetOrgOk

`func (o *CloudWire) GetOrgOk() (*string, bool)`

GetOrgOk returns a tuple with the Org field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrg

`func (o *CloudWire) SetOrg(v string)`

SetOrg sets Org field to given value.

### HasOrg

`func (o *CloudWire) HasOrg() bool`

HasOrg returns a boolean if a field has been set.

### GetPath

`func (o *CloudWire) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *CloudWire) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *CloudWire) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *CloudWire) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetPrevHash

`func (o *CloudWire) GetPrevHash() string`

GetPrevHash returns the PrevHash field if non-nil, zero value otherwise.

### GetPrevHashOk

`func (o *CloudWire) GetPrevHashOk() (*string, bool)`

GetPrevHashOk returns a tuple with the PrevHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrevHash

`func (o *CloudWire) SetPrevHash(v string)`

SetPrevHash sets PrevHash field to given value.

### HasPrevHash

`func (o *CloudWire) HasPrevHash() bool`

HasPrevHash returns a boolean if a field has been set.

### GetReason

`func (o *CloudWire) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *CloudWire) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *CloudWire) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *CloudWire) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetRequestId

`func (o *CloudWire) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *CloudWire) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *CloudWire) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *CloudWire) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.

### GetResource

`func (o *CloudWire) GetResource() string`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *CloudWire) GetResourceOk() (*string, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *CloudWire) SetResource(v string)`

SetResource sets Resource field to given value.

### HasResource

`func (o *CloudWire) HasResource() bool`

HasResource returns a boolean if a field has been set.

### GetResourceId

`func (o *CloudWire) GetResourceId() string`

GetResourceId returns the ResourceId field if non-nil, zero value otherwise.

### GetResourceIdOk

`func (o *CloudWire) GetResourceIdOk() (*string, bool)`

GetResourceIdOk returns a tuple with the ResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceId

`func (o *CloudWire) SetResourceId(v string)`

SetResourceId sets ResourceId field to given value.

### HasResourceId

`func (o *CloudWire) HasResourceId() bool`

HasResourceId returns a boolean if a field has been set.

### GetResult

`func (o *CloudWire) GetResult() string`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *CloudWire) GetResultOk() (*string, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *CloudWire) SetResult(v string)`

SetResult sets Result field to given value.

### HasResult

`func (o *CloudWire) HasResult() bool`

HasResult returns a boolean if a field has been set.

### GetSeq

`func (o *CloudWire) GetSeq() int32`

GetSeq returns the Seq field if non-nil, zero value otherwise.

### GetSeqOk

`func (o *CloudWire) GetSeqOk() (*int32, bool)`

GetSeqOk returns a tuple with the Seq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeq

`func (o *CloudWire) SetSeq(v int32)`

SetSeq sets Seq field to given value.

### HasSeq

`func (o *CloudWire) HasSeq() bool`

HasSeq returns a boolean if a field has been set.

### GetSourceIp

`func (o *CloudWire) GetSourceIp() string`

GetSourceIp returns the SourceIp field if non-nil, zero value otherwise.

### GetSourceIpOk

`func (o *CloudWire) GetSourceIpOk() (*string, bool)`

GetSourceIpOk returns a tuple with the SourceIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceIp

`func (o *CloudWire) SetSourceIp(v string)`

SetSourceIp sets SourceIp field to given value.

### HasSourceIp

`func (o *CloudWire) HasSourceIp() bool`

HasSourceIp returns a boolean if a field has been set.

### GetStatus

`func (o *CloudWire) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CloudWire) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CloudWire) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CloudWire) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSub

`func (o *CloudWire) GetSub() string`

GetSub returns the Sub field if non-nil, zero value otherwise.

### GetSubOk

`func (o *CloudWire) GetSubOk() (*string, bool)`

GetSubOk returns a tuple with the Sub field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSub

`func (o *CloudWire) SetSub(v string)`

SetSub sets Sub field to given value.

### HasSub

`func (o *CloudWire) HasSub() bool`

HasSub returns a boolean if a field has been set.

### GetTime

`func (o *CloudWire) GetTime() string`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *CloudWire) GetTimeOk() (*string, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *CloudWire) SetTime(v string)`

SetTime sets Time field to given value.

### HasTime

`func (o *CloudWire) HasTime() bool`

HasTime returns a boolean if a field has been set.

### GetUserAgent

`func (o *CloudWire) GetUserAgent() string`

GetUserAgent returns the UserAgent field if non-nil, zero value otherwise.

### GetUserAgentOk

`func (o *CloudWire) GetUserAgentOk() (*string, bool)`

GetUserAgentOk returns a tuple with the UserAgent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAgent

`func (o *CloudWire) SetUserAgent(v string)`

SetUserAgent sets UserAgent field to given value.

### HasUserAgent

`func (o *CloudWire) HasUserAgent() bool`

HasUserAgent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


