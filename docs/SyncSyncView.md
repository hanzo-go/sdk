# SyncSyncView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Actor** | Pointer to **string** | Actor is the identity a reconcile writes AS. It is the loop guard: a change this identity made is one we already have, so it is not synced back. | [optional] 
**CreatedAt** | Pointer to **string** | CreatedAt is when the link was first declared, RFC3339 in UTC. | [optional] 
**Direction** | Pointer to **string** | Direction is which way work flows: \&quot;both\&quot;, \&quot;pull\&quot; (target ← source), \&quot;push\&quot; (source → target), or \&quot;off\&quot; — which keeps the link declared and moves nothing. | [optional] 
**Id** | Pointer to **string** | ID is the link&#39;s handle, derived from its source and target — which is what makes re-declaring the same pair an update rather than a duplicate. | [optional] 
**Kind** | Pointer to **string** | Kind is what is being synced. \&quot;git\&quot; today; the field exists so a storage or database link is a value here rather than a second route family. | [optional] 
**Native** | Pointer to [**SyncNativeView**](SyncNativeView.md) | Native is the repo link&#39;s copy on the forge: where it lives and how it stands. Absent for an account link, and absent when the forge could not be read — an unread forge is not reported as a healthy one. | [optional] 
**Scope** | Pointer to **string** | Scope is \&quot;repo\&quot; for a link to one repository, or \&quot;account\&quot; for a link to a whole GitHub account, which declares a repo link for every repository the account holds — the ones created later included. | [optional] 
**Source** | Pointer to [**SyncEndpointView**](SyncEndpointView.md) | Source is the side read FROM on a pull. | [optional] 
**Target** | Pointer to [**SyncEndpointView**](SyncEndpointView.md) | Target is the side written TO on a push. | [optional] 
**Trigger** | Pointer to **string** | Trigger is what starts a reconcile: \&quot;webhook\&quot; (the provider tells us), \&quot;poll\&quot; (we ask on a schedule), or \&quot;manual\&quot; (only an explicit call). | [optional] 
**UpdatedAt** | Pointer to **string** | UpdatedAt is bumped by every reconcile, so it reads as the LAST-SYNCED time rather than the last edit. Absent until the first one runs. | [optional] 

## Methods

### NewSyncSyncView

`func NewSyncSyncView() *SyncSyncView`

NewSyncSyncView instantiates a new SyncSyncView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSyncSyncViewWithDefaults

`func NewSyncSyncViewWithDefaults() *SyncSyncView`

NewSyncSyncViewWithDefaults instantiates a new SyncSyncView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActor

`func (o *SyncSyncView) GetActor() string`

GetActor returns the Actor field if non-nil, zero value otherwise.

### GetActorOk

`func (o *SyncSyncView) GetActorOk() (*string, bool)`

GetActorOk returns a tuple with the Actor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActor

`func (o *SyncSyncView) SetActor(v string)`

SetActor sets Actor field to given value.

### HasActor

`func (o *SyncSyncView) HasActor() bool`

HasActor returns a boolean if a field has been set.

### GetCreatedAt

`func (o *SyncSyncView) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *SyncSyncView) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *SyncSyncView) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *SyncSyncView) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDirection

`func (o *SyncSyncView) GetDirection() string`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *SyncSyncView) GetDirectionOk() (*string, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *SyncSyncView) SetDirection(v string)`

SetDirection sets Direction field to given value.

### HasDirection

`func (o *SyncSyncView) HasDirection() bool`

HasDirection returns a boolean if a field has been set.

### GetId

`func (o *SyncSyncView) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SyncSyncView) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SyncSyncView) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SyncSyncView) HasId() bool`

HasId returns a boolean if a field has been set.

### GetKind

`func (o *SyncSyncView) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *SyncSyncView) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *SyncSyncView) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *SyncSyncView) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetNative

`func (o *SyncSyncView) GetNative() SyncNativeView`

GetNative returns the Native field if non-nil, zero value otherwise.

### GetNativeOk

`func (o *SyncSyncView) GetNativeOk() (*SyncNativeView, bool)`

GetNativeOk returns a tuple with the Native field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNative

`func (o *SyncSyncView) SetNative(v SyncNativeView)`

SetNative sets Native field to given value.

### HasNative

`func (o *SyncSyncView) HasNative() bool`

HasNative returns a boolean if a field has been set.

### GetScope

`func (o *SyncSyncView) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *SyncSyncView) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *SyncSyncView) SetScope(v string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *SyncSyncView) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetSource

`func (o *SyncSyncView) GetSource() SyncEndpointView`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *SyncSyncView) GetSourceOk() (*SyncEndpointView, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *SyncSyncView) SetSource(v SyncEndpointView)`

SetSource sets Source field to given value.

### HasSource

`func (o *SyncSyncView) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetTarget

`func (o *SyncSyncView) GetTarget() SyncEndpointView`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *SyncSyncView) GetTargetOk() (*SyncEndpointView, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *SyncSyncView) SetTarget(v SyncEndpointView)`

SetTarget sets Target field to given value.

### HasTarget

`func (o *SyncSyncView) HasTarget() bool`

HasTarget returns a boolean if a field has been set.

### GetTrigger

`func (o *SyncSyncView) GetTrigger() string`

GetTrigger returns the Trigger field if non-nil, zero value otherwise.

### GetTriggerOk

`func (o *SyncSyncView) GetTriggerOk() (*string, bool)`

GetTriggerOk returns a tuple with the Trigger field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrigger

`func (o *SyncSyncView) SetTrigger(v string)`

SetTrigger sets Trigger field to given value.

### HasTrigger

`func (o *SyncSyncView) HasTrigger() bool`

HasTrigger returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *SyncSyncView) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *SyncSyncView) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *SyncSyncView) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *SyncSyncView) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


