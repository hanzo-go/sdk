# UserEnablementItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CanOptIn** | Pointer to **bool** | beta &amp;&amp; not yet opted in | [optional] 
**Effective** | Pointer to **bool** | visible to the caller&#39;s org | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Kind** | Pointer to **string** |  | [optional] 
**OptedIn** | Pointer to **bool** | caller&#39;s org on the beta list | [optional] 
**State** | Pointer to **string** | off|beta|ga | [optional] 

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


