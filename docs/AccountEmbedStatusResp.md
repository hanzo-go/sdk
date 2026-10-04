# AccountEmbedStatusResp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**App** | Pointer to **string** | App is the app this verdict is about. | [optional] 
**EmbedUrl** | Pointer to **string** | EmbedURL is the in-app landing URL to frame. Empty when the caller is not entitled — a non-entitled caller never receives it. | [optional] 
**Entitled** | Pointer to **bool** | Entitled is whether the caller&#39;s org may frame this brand-owned app. | [optional] 
**Origin** | Pointer to **string** | Origin is the app&#39;s origin on this deployment&#39;s own brand domain. | [optional] 
**Phase** | Pointer to **string** | Phase is the verdict in one word: not-entitled, not-provisioned or ready. | [optional] 
**Reachable** | Pointer to **bool** | Reachable is whether the app answered the liveness probe. | [optional] 

## Methods

### NewAccountEmbedStatusResp

`func NewAccountEmbedStatusResp() *AccountEmbedStatusResp`

NewAccountEmbedStatusResp instantiates a new AccountEmbedStatusResp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountEmbedStatusRespWithDefaults

`func NewAccountEmbedStatusRespWithDefaults() *AccountEmbedStatusResp`

NewAccountEmbedStatusRespWithDefaults instantiates a new AccountEmbedStatusResp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApp

`func (o *AccountEmbedStatusResp) GetApp() string`

GetApp returns the App field if non-nil, zero value otherwise.

### GetAppOk

`func (o *AccountEmbedStatusResp) GetAppOk() (*string, bool)`

GetAppOk returns a tuple with the App field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApp

`func (o *AccountEmbedStatusResp) SetApp(v string)`

SetApp sets App field to given value.

### HasApp

`func (o *AccountEmbedStatusResp) HasApp() bool`

HasApp returns a boolean if a field has been set.

### GetEmbedUrl

`func (o *AccountEmbedStatusResp) GetEmbedUrl() string`

GetEmbedUrl returns the EmbedUrl field if non-nil, zero value otherwise.

### GetEmbedUrlOk

`func (o *AccountEmbedStatusResp) GetEmbedUrlOk() (*string, bool)`

GetEmbedUrlOk returns a tuple with the EmbedUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbedUrl

`func (o *AccountEmbedStatusResp) SetEmbedUrl(v string)`

SetEmbedUrl sets EmbedUrl field to given value.

### HasEmbedUrl

`func (o *AccountEmbedStatusResp) HasEmbedUrl() bool`

HasEmbedUrl returns a boolean if a field has been set.

### GetEntitled

`func (o *AccountEmbedStatusResp) GetEntitled() bool`

GetEntitled returns the Entitled field if non-nil, zero value otherwise.

### GetEntitledOk

`func (o *AccountEmbedStatusResp) GetEntitledOk() (*bool, bool)`

GetEntitledOk returns a tuple with the Entitled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitled

`func (o *AccountEmbedStatusResp) SetEntitled(v bool)`

SetEntitled sets Entitled field to given value.

### HasEntitled

`func (o *AccountEmbedStatusResp) HasEntitled() bool`

HasEntitled returns a boolean if a field has been set.

### GetOrigin

`func (o *AccountEmbedStatusResp) GetOrigin() string`

GetOrigin returns the Origin field if non-nil, zero value otherwise.

### GetOriginOk

`func (o *AccountEmbedStatusResp) GetOriginOk() (*string, bool)`

GetOriginOk returns a tuple with the Origin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigin

`func (o *AccountEmbedStatusResp) SetOrigin(v string)`

SetOrigin sets Origin field to given value.

### HasOrigin

`func (o *AccountEmbedStatusResp) HasOrigin() bool`

HasOrigin returns a boolean if a field has been set.

### GetPhase

`func (o *AccountEmbedStatusResp) GetPhase() string`

GetPhase returns the Phase field if non-nil, zero value otherwise.

### GetPhaseOk

`func (o *AccountEmbedStatusResp) GetPhaseOk() (*string, bool)`

GetPhaseOk returns a tuple with the Phase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase

`func (o *AccountEmbedStatusResp) SetPhase(v string)`

SetPhase sets Phase field to given value.

### HasPhase

`func (o *AccountEmbedStatusResp) HasPhase() bool`

HasPhase returns a boolean if a field has been set.

### GetReachable

`func (o *AccountEmbedStatusResp) GetReachable() bool`

GetReachable returns the Reachable field if non-nil, zero value otherwise.

### GetReachableOk

`func (o *AccountEmbedStatusResp) GetReachableOk() (*bool, bool)`

GetReachableOk returns a tuple with the Reachable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReachable

`func (o *AccountEmbedStatusResp) SetReachable(v bool)`

SetReachable sets Reachable field to given value.

### HasReachable

`func (o *AccountEmbedStatusResp) HasReachable() bool`

HasReachable returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


