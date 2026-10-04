# LinkRevokeResp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to [**[]LinkLinkView**](LinkLinkView.md) | Links is each revoked row with its new status — retained, not deleted, so usage history and the audit trail survive the log-out. | [optional] 
**Revoked** | Pointer to **int64** | Revoked is how many links this call revoked. | [optional] 
**SessionsStopped** | Pointer to **int64** | SessionsStopped is how many of the caller&#39;s own agent sessions stopped. A stop that fails does not fail the revoke, so this may honestly report fewer. | [optional] 

## Methods

### NewLinkRevokeResp

`func NewLinkRevokeResp() *LinkRevokeResp`

NewLinkRevokeResp instantiates a new LinkRevokeResp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLinkRevokeRespWithDefaults

`func NewLinkRevokeRespWithDefaults() *LinkRevokeResp`

NewLinkRevokeRespWithDefaults instantiates a new LinkRevokeResp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *LinkRevokeResp) GetLinks() []LinkLinkView`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *LinkRevokeResp) GetLinksOk() (*[]LinkLinkView, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *LinkRevokeResp) SetLinks(v []LinkLinkView)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *LinkRevokeResp) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetRevoked

`func (o *LinkRevokeResp) GetRevoked() int64`

GetRevoked returns the Revoked field if non-nil, zero value otherwise.

### GetRevokedOk

`func (o *LinkRevokeResp) GetRevokedOk() (*int64, bool)`

GetRevokedOk returns a tuple with the Revoked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevoked

`func (o *LinkRevokeResp) SetRevoked(v int64)`

SetRevoked sets Revoked field to given value.

### HasRevoked

`func (o *LinkRevokeResp) HasRevoked() bool`

HasRevoked returns a boolean if a field has been set.

### GetSessionsStopped

`func (o *LinkRevokeResp) GetSessionsStopped() int64`

GetSessionsStopped returns the SessionsStopped field if non-nil, zero value otherwise.

### GetSessionsStoppedOk

`func (o *LinkRevokeResp) GetSessionsStoppedOk() (*int64, bool)`

GetSessionsStoppedOk returns a tuple with the SessionsStopped field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionsStopped

`func (o *LinkRevokeResp) SetSessionsStopped(v int64)`

SetSessionsStopped sets SessionsStopped field to given value.

### HasSessionsStopped

`func (o *LinkRevokeResp) HasSessionsStopped() bool`

HasSessionsStopped returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


