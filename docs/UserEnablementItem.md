# UserEnablementItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CanOptIn** | Pointer to **bool** | CanOptIn is whether POST /v1/pricing/enablement/optin would do anything here: the item is in beta and this org is not on its list yet. False for a caller with no validated org, who has no org to enrol. | [optional] 
**Effective** | Pointer to **bool** | Effective is whether the caller&#39;s org may use the item right now, which is the field to branch on: true for any ga item, for a beta this org holds, and never for an off one. | [optional] 
**Id** | Pointer to **string** | ID is the item within that namespace — a model id, a provider name, or a feature key. | [optional] 
**Kind** | Pointer to **string** | Kind is the namespace the id lives in: \&quot;model\&quot;, \&quot;provider\&quot; or \&quot;feature\&quot;. | [optional] 
**OptedIn** | Pointer to **bool** | OptedIn is whether the caller&#39;s org is on this item&#39;s beta grant list. It can be true on an \&quot;off\&quot; item — the list survives the kill switch and is simply ignored while it is thrown — so it does not imply Effective. | [optional] 
**State** | Pointer to **string** | State is the item&#39;s GLOBAL availability — \&quot;off\&quot;, \&quot;beta\&quot; or \&quot;ga\&quot; — which is the operator&#39;s setting and not this caller&#39;s answer. Effective is that. | [optional] 

## Methods

### NewUserEnablementItem

`func NewUserEnablementItem() *UserEnablementItem`

NewUserEnablementItem instantiates a new UserEnablementItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserEnablementItemWithDefaults

`func NewUserEnablementItemWithDefaults() *UserEnablementItem`

NewUserEnablementItemWithDefaults instantiates a new UserEnablementItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCanOptIn

`func (o *UserEnablementItem) GetCanOptIn() bool`

GetCanOptIn returns the CanOptIn field if non-nil, zero value otherwise.

### GetCanOptInOk

`func (o *UserEnablementItem) GetCanOptInOk() (*bool, bool)`

GetCanOptInOk returns a tuple with the CanOptIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanOptIn

`func (o *UserEnablementItem) SetCanOptIn(v bool)`

SetCanOptIn sets CanOptIn field to given value.

### HasCanOptIn

`func (o *UserEnablementItem) HasCanOptIn() bool`

HasCanOptIn returns a boolean if a field has been set.

### GetEffective

`func (o *UserEnablementItem) GetEffective() bool`

GetEffective returns the Effective field if non-nil, zero value otherwise.

### GetEffectiveOk

`func (o *UserEnablementItem) GetEffectiveOk() (*bool, bool)`

GetEffectiveOk returns a tuple with the Effective field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffective

`func (o *UserEnablementItem) SetEffective(v bool)`

SetEffective sets Effective field to given value.

### HasEffective

`func (o *UserEnablementItem) HasEffective() bool`

HasEffective returns a boolean if a field has been set.

### GetId

`func (o *UserEnablementItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UserEnablementItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UserEnablementItem) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *UserEnablementItem) HasId() bool`

HasId returns a boolean if a field has been set.

### GetKind

`func (o *UserEnablementItem) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *UserEnablementItem) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *UserEnablementItem) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *UserEnablementItem) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetOptedIn

`func (o *UserEnablementItem) GetOptedIn() bool`

GetOptedIn returns the OptedIn field if non-nil, zero value otherwise.

### GetOptedInOk

`func (o *UserEnablementItem) GetOptedInOk() (*bool, bool)`

GetOptedInOk returns a tuple with the OptedIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptedIn

`func (o *UserEnablementItem) SetOptedIn(v bool)`

SetOptedIn sets OptedIn field to given value.

### HasOptedIn

`func (o *UserEnablementItem) HasOptedIn() bool`

HasOptedIn returns a boolean if a field has been set.

### GetState

`func (o *UserEnablementItem) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *UserEnablementItem) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *UserEnablementItem) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *UserEnablementItem) HasState() bool`

HasState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


