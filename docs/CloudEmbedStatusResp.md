# CloudEmbedStatusResp

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

### NewCloudEmbedStatusResp

`func NewCloudEmbedStatusResp() *CloudEmbedStatusResp`

NewCloudEmbedStatusResp instantiates a new CloudEmbedStatusResp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudEmbedStatusRespWithDefaults

`func NewCloudEmbedStatusRespWithDefaults() *CloudEmbedStatusResp`

NewCloudEmbedStatusRespWithDefaults instantiates a new CloudEmbedStatusResp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApp

`func (o *CloudEmbedStatusResp) GetApp() string`

GetApp returns the App field if non-nil, zero value otherwise.

### GetAppOk

`func (o *CloudEmbedStatusResp) GetAppOk() (*string, bool)`

GetAppOk returns a tuple with the App field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApp

`func (o *CloudEmbedStatusResp) SetApp(v string)`

SetApp sets App field to given value.

### HasApp

`func (o *CloudEmbedStatusResp) HasApp() bool`

HasApp returns a boolean if a field has been set.

### GetEmbedUrl

`func (o *CloudEmbedStatusResp) GetEmbedUrl() string`

GetEmbedUrl returns the EmbedUrl field if non-nil, zero value otherwise.

### GetEmbedUrlOk

`func (o *CloudEmbedStatusResp) GetEmbedUrlOk() (*string, bool)`

GetEmbedUrlOk returns a tuple with the EmbedUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbedUrl

`func (o *CloudEmbedStatusResp) SetEmbedUrl(v string)`

SetEmbedUrl sets EmbedUrl field to given value.

### HasEmbedUrl

`func (o *CloudEmbedStatusResp) HasEmbedUrl() bool`

HasEmbedUrl returns a boolean if a field has been set.

### GetEntitled

`func (o *CloudEmbedStatusResp) GetEntitled() bool`

GetEntitled returns the Entitled field if non-nil, zero value otherwise.

### GetEntitledOk

`func (o *CloudEmbedStatusResp) GetEntitledOk() (*bool, bool)`

GetEntitledOk returns a tuple with the Entitled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitled

`func (o *CloudEmbedStatusResp) SetEntitled(v bool)`

SetEntitled sets Entitled field to given value.

### HasEntitled

`func (o *CloudEmbedStatusResp) HasEntitled() bool`

HasEntitled returns a boolean if a field has been set.

### GetOrigin

`func (o *CloudEmbedStatusResp) GetOrigin() string`

GetOrigin returns the Origin field if non-nil, zero value otherwise.

### GetOriginOk

`func (o *CloudEmbedStatusResp) GetOriginOk() (*string, bool)`

GetOriginOk returns a tuple with the Origin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigin

`func (o *CloudEmbedStatusResp) SetOrigin(v string)`

SetOrigin sets Origin field to given value.

### HasOrigin

`func (o *CloudEmbedStatusResp) HasOrigin() bool`

HasOrigin returns a boolean if a field has been set.

### GetPhase

`func (o *CloudEmbedStatusResp) GetPhase() string`

GetPhase returns the Phase field if non-nil, zero value otherwise.

### GetPhaseOk

`func (o *CloudEmbedStatusResp) GetPhaseOk() (*string, bool)`

GetPhaseOk returns a tuple with the Phase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase

`func (o *CloudEmbedStatusResp) SetPhase(v string)`

SetPhase sets Phase field to given value.

### HasPhase

`func (o *CloudEmbedStatusResp) HasPhase() bool`

HasPhase returns a boolean if a field has been set.

### GetReachable

`func (o *CloudEmbedStatusResp) GetReachable() bool`

GetReachable returns the Reachable field if non-nil, zero value otherwise.

### GetReachableOk

`func (o *CloudEmbedStatusResp) GetReachableOk() (*bool, bool)`

GetReachableOk returns a tuple with the Reachable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReachable

`func (o *CloudEmbedStatusResp) SetReachable(v bool)`

SetReachable sets Reachable field to given value.

### HasReachable

`func (o *CloudEmbedStatusResp) HasReachable() bool`

HasReachable returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


