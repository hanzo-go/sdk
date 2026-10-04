# ProviderOpView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | Description says what it does. Only on GET /v1/provider/{provider}. | [optional] 
**DisplayName** | Pointer to **string** | DisplayName is what the card calls it. | [optional] 
**Name** | Pointer to **string** | Name is the action&#39;s id — what a run or a flow step names. | [optional] 
**Props** | Pointer to [**[]ProviderPropSpec**](ProviderPropSpec.md) | Props are its inputs — what POST /v1/provider/{provider}/run takes as input. Only on GET /v1/provider/{provider}; the list names actions only. | [optional] 

## Methods

### NewProviderOpView

`func NewProviderOpView() *ProviderOpView`

NewProviderOpView instantiates a new ProviderOpView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProviderOpViewWithDefaults

`func NewProviderOpViewWithDefaults() *ProviderOpView`

NewProviderOpViewWithDefaults instantiates a new ProviderOpView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *ProviderOpView) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ProviderOpView) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ProviderOpView) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ProviderOpView) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDisplayName

`func (o *ProviderOpView) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *ProviderOpView) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *ProviderOpView) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *ProviderOpView) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetName

`func (o *ProviderOpView) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProviderOpView) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProviderOpView) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ProviderOpView) HasName() bool`

HasName returns a boolean if a field has been set.

### GetProps

`func (o *ProviderOpView) GetProps() []ProviderPropSpec`

GetProps returns the Props field if non-nil, zero value otherwise.

### GetPropsOk

`func (o *ProviderOpView) GetPropsOk() (*[]ProviderPropSpec, bool)`

GetPropsOk returns a tuple with the Props field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProps

`func (o *ProviderOpView) SetProps(v []ProviderPropSpec)`

SetProps sets Props field to given value.

### HasProps

`func (o *ProviderOpView) HasProps() bool`

HasProps returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


