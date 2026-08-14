# SetEnablementBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BetaOrgs** | Pointer to **[]string** | BetaOrgs REPLACES the item&#39;s beta grant list when present. Omit it to leave the existing grants alone. | [optional] 
**Id** | Pointer to **string** | ID is the item within that namespace — a model id, a provider name, or a feature&#39;s key. | [optional] 
**Kind** | Pointer to **string** | Kind is the item&#39;s namespace: \&quot;model\&quot;, \&quot;provider\&quot; or \&quot;feature\&quot;. | [optional] 
**State** | Pointer to **string** | State is the item&#39;s global enablement: \&quot;off\&quot; (hidden from everyone, absolutely), \&quot;beta\&quot; (visible only to granted orgs) or \&quot;ga\&quot; (visible to everyone). Required. | [optional] 

## Methods

### NewSetEnablementBody

`func NewSetEnablementBody() *SetEnablementBody`

NewSetEnablementBody instantiates a new SetEnablementBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSetEnablementBodyWithDefaults

`func NewSetEnablementBodyWithDefaults() *SetEnablementBody`

NewSetEnablementBodyWithDefaults instantiates a new SetEnablementBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBetaOrgs

`func (o *SetEnablementBody) GetBetaOrgs() []string`

GetBetaOrgs returns the BetaOrgs field if non-nil, zero value otherwise.

### GetBetaOrgsOk

`func (o *SetEnablementBody) GetBetaOrgsOk() (*[]string, bool)`

GetBetaOrgsOk returns a tuple with the BetaOrgs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBetaOrgs

`func (o *SetEnablementBody) SetBetaOrgs(v []string)`

SetBetaOrgs sets BetaOrgs field to given value.

### HasBetaOrgs

`func (o *SetEnablementBody) HasBetaOrgs() bool`

HasBetaOrgs returns a boolean if a field has been set.

### GetId

`func (o *SetEnablementBody) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SetEnablementBody) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SetEnablementBody) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SetEnablementBody) HasId() bool`

HasId returns a boolean if a field has been set.

### GetKind

`func (o *SetEnablementBody) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *SetEnablementBody) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *SetEnablementBody) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *SetEnablementBody) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetState

`func (o *SetEnablementBody) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *SetEnablementBody) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *SetEnablementBody) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *SetEnablementBody) HasState() bool`

HasState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


