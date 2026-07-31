# CloudSetEnablementBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BetaOrgs** | Pointer to **[]string** | BetaOrgs REPLACES the item&#39;s beta grant list when present. Omit it to leave the existing grants alone. | [optional] 
**Id** | Pointer to **string** | ID is the item within that namespace — a model id, a provider name, or a feature&#39;s key. | [optional] 
**Kind** | Pointer to **string** | Kind is the item&#39;s namespace: \&quot;model\&quot;, \&quot;provider\&quot; or \&quot;feature\&quot;. | [optional] 
**State** | Pointer to **string** | State is the item&#39;s global enablement: \&quot;off\&quot; (hidden from everyone, absolutely), \&quot;beta\&quot; (visible only to granted orgs) or \&quot;ga\&quot; (visible to everyone). Required. | [optional] 

## Methods

### NewCloudSetEnablementBody

`func NewCloudSetEnablementBody() *CloudSetEnablementBody`

NewCloudSetEnablementBody instantiates a new CloudSetEnablementBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudSetEnablementBodyWithDefaults

`func NewCloudSetEnablementBodyWithDefaults() *CloudSetEnablementBody`

NewCloudSetEnablementBodyWithDefaults instantiates a new CloudSetEnablementBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBetaOrgs

`func (o *CloudSetEnablementBody) GetBetaOrgs() []string`

GetBetaOrgs returns the BetaOrgs field if non-nil, zero value otherwise.

### GetBetaOrgsOk

`func (o *CloudSetEnablementBody) GetBetaOrgsOk() (*[]string, bool)`

GetBetaOrgsOk returns a tuple with the BetaOrgs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBetaOrgs

`func (o *CloudSetEnablementBody) SetBetaOrgs(v []string)`

SetBetaOrgs sets BetaOrgs field to given value.

### HasBetaOrgs

`func (o *CloudSetEnablementBody) HasBetaOrgs() bool`

HasBetaOrgs returns a boolean if a field has been set.

### GetId

`func (o *CloudSetEnablementBody) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CloudSetEnablementBody) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CloudSetEnablementBody) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CloudSetEnablementBody) HasId() bool`

HasId returns a boolean if a field has been set.

### GetKind

`func (o *CloudSetEnablementBody) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *CloudSetEnablementBody) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *CloudSetEnablementBody) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *CloudSetEnablementBody) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetState

`func (o *CloudSetEnablementBody) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *CloudSetEnablementBody) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *CloudSetEnablementBody) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *CloudSetEnablementBody) HasState() bool`

HasState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


