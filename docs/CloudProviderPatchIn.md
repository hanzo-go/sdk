# CloudProviderPatchIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Beta** | Pointer to **bool** |  | [optional] 
**BetaOrgs** | Pointer to **[]string** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**Name** | Pointer to **string** | Name is the provider the overlay belongs to, from the URL. | [optional] 
**Overrides** | Pointer to **interface{}** |  | [optional] 
**State** | Pointer to **string** |  | [optional] 

## Methods

### NewCloudProviderPatchIn

`func NewCloudProviderPatchIn() *CloudProviderPatchIn`

NewCloudProviderPatchIn instantiates a new CloudProviderPatchIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudProviderPatchInWithDefaults

`func NewCloudProviderPatchInWithDefaults() *CloudProviderPatchIn`

NewCloudProviderPatchInWithDefaults instantiates a new CloudProviderPatchIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBeta

`func (o *CloudProviderPatchIn) GetBeta() bool`

GetBeta returns the Beta field if non-nil, zero value otherwise.

### GetBetaOk

`func (o *CloudProviderPatchIn) GetBetaOk() (*bool, bool)`

GetBetaOk returns a tuple with the Beta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBeta

`func (o *CloudProviderPatchIn) SetBeta(v bool)`

SetBeta sets Beta field to given value.

### HasBeta

`func (o *CloudProviderPatchIn) HasBeta() bool`

HasBeta returns a boolean if a field has been set.

### GetBetaOrgs

`func (o *CloudProviderPatchIn) GetBetaOrgs() []string`

GetBetaOrgs returns the BetaOrgs field if non-nil, zero value otherwise.

### GetBetaOrgsOk

`func (o *CloudProviderPatchIn) GetBetaOrgsOk() (*[]string, bool)`

GetBetaOrgsOk returns a tuple with the BetaOrgs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBetaOrgs

`func (o *CloudProviderPatchIn) SetBetaOrgs(v []string)`

SetBetaOrgs sets BetaOrgs field to given value.

### HasBetaOrgs

`func (o *CloudProviderPatchIn) HasBetaOrgs() bool`

HasBetaOrgs returns a boolean if a field has been set.

### GetEnabled

`func (o *CloudProviderPatchIn) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *CloudProviderPatchIn) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *CloudProviderPatchIn) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *CloudProviderPatchIn) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetName

`func (o *CloudProviderPatchIn) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CloudProviderPatchIn) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CloudProviderPatchIn) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CloudProviderPatchIn) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOverrides

`func (o *CloudProviderPatchIn) GetOverrides() interface{}`

GetOverrides returns the Overrides field if non-nil, zero value otherwise.

### GetOverridesOk

`func (o *CloudProviderPatchIn) GetOverridesOk() (*interface{}, bool)`

GetOverridesOk returns a tuple with the Overrides field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrides

`func (o *CloudProviderPatchIn) SetOverrides(v interface{})`

SetOverrides sets Overrides field to given value.

### HasOverrides

`func (o *CloudProviderPatchIn) HasOverrides() bool`

HasOverrides returns a boolean if a field has been set.

### SetOverridesNil

`func (o *CloudProviderPatchIn) SetOverridesNil(b bool)`

 SetOverridesNil sets the value for Overrides to be an explicit nil

### UnsetOverrides
`func (o *CloudProviderPatchIn) UnsetOverrides()`

UnsetOverrides ensures that no value is present for Overrides, not even an explicit nil
### GetState

`func (o *CloudProviderPatchIn) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *CloudProviderPatchIn) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *CloudProviderPatchIn) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *CloudProviderPatchIn) HasState() bool`

HasState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


