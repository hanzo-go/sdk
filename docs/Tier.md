# Tier

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Balance** | Pointer to [**TierBalance**](TierBalance.md) |  | [optional] 
**Tier** | Pointer to [**TierLimits**](TierLimits.md) |  | [optional] 
**User** | Pointer to **string** |  | [optional] 
**Windows** | Pointer to [**[]Window**](Window.md) |  | [optional] 

## Methods

### NewTier

`func NewTier() *Tier`

NewTier instantiates a new Tier object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTierWithDefaults

`func NewTierWithDefaults() *Tier`

NewTierWithDefaults instantiates a new Tier object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBalance

`func (o *Tier) GetBalance() TierBalance`

GetBalance returns the Balance field if non-nil, zero value otherwise.

### GetBalanceOk

`func (o *Tier) GetBalanceOk() (*TierBalance, bool)`

GetBalanceOk returns a tuple with the Balance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalance

`func (o *Tier) SetBalance(v TierBalance)`

SetBalance sets Balance field to given value.

### HasBalance

`func (o *Tier) HasBalance() bool`

HasBalance returns a boolean if a field has been set.

### GetTier

`func (o *Tier) GetTier() TierLimits`

GetTier returns the Tier field if non-nil, zero value otherwise.

### GetTierOk

`func (o *Tier) GetTierOk() (*TierLimits, bool)`

GetTierOk returns a tuple with the Tier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTier

`func (o *Tier) SetTier(v TierLimits)`

SetTier sets Tier field to given value.

### HasTier

`func (o *Tier) HasTier() bool`

HasTier returns a boolean if a field has been set.

### GetUser

`func (o *Tier) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *Tier) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *Tier) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *Tier) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetWindows

`func (o *Tier) GetWindows() []Window`

GetWindows returns the Windows field if non-nil, zero value otherwise.

### GetWindowsOk

`func (o *Tier) GetWindowsOk() (*[]Window, bool)`

GetWindowsOk returns a tuple with the Windows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindows

`func (o *Tier) SetWindows(v []Window)`

SetWindows sets Windows field to given value.

### HasWindows

`func (o *Tier) HasWindows() bool`

HasWindows returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


