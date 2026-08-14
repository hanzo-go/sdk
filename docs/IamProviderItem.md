# IamProviderItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BindingRule** | Pointer to **[]string** |  | [optional] 
**CanSignIn** | Pointer to **bool** |  | [optional] 
**CanSignUp** | Pointer to **bool** |  | [optional] 
**CanUnlink** | Pointer to **bool** |  | [optional] 
**CountryCodes** | Pointer to **[]string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Owner** | Pointer to **string** |  | [optional] 
**Prompted** | Pointer to **bool** |  | [optional] 
**Provider** | Pointer to [**IamProvider**](IamProvider.md) |  | [optional] 
**Rule** | Pointer to **string** |  | [optional] 
**SignupGroup** | Pointer to **string** |  | [optional] 

## Methods

### NewIamProviderItem

`func NewIamProviderItem() *IamProviderItem`

NewIamProviderItem instantiates a new IamProviderItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamProviderItemWithDefaults

`func NewIamProviderItemWithDefaults() *IamProviderItem`

NewIamProviderItemWithDefaults instantiates a new IamProviderItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBindingRule

`func (o *IamProviderItem) GetBindingRule() []string`

GetBindingRule returns the BindingRule field if non-nil, zero value otherwise.

### GetBindingRuleOk

`func (o *IamProviderItem) GetBindingRuleOk() (*[]string, bool)`

GetBindingRuleOk returns a tuple with the BindingRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBindingRule

`func (o *IamProviderItem) SetBindingRule(v []string)`

SetBindingRule sets BindingRule field to given value.

### HasBindingRule

`func (o *IamProviderItem) HasBindingRule() bool`

HasBindingRule returns a boolean if a field has been set.

### GetCanSignIn

`func (o *IamProviderItem) GetCanSignIn() bool`

GetCanSignIn returns the CanSignIn field if non-nil, zero value otherwise.

### GetCanSignInOk

`func (o *IamProviderItem) GetCanSignInOk() (*bool, bool)`

GetCanSignInOk returns a tuple with the CanSignIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanSignIn

`func (o *IamProviderItem) SetCanSignIn(v bool)`

SetCanSignIn sets CanSignIn field to given value.

### HasCanSignIn

`func (o *IamProviderItem) HasCanSignIn() bool`

HasCanSignIn returns a boolean if a field has been set.

### GetCanSignUp

`func (o *IamProviderItem) GetCanSignUp() bool`

GetCanSignUp returns the CanSignUp field if non-nil, zero value otherwise.

### GetCanSignUpOk

`func (o *IamProviderItem) GetCanSignUpOk() (*bool, bool)`

GetCanSignUpOk returns a tuple with the CanSignUp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanSignUp

`func (o *IamProviderItem) SetCanSignUp(v bool)`

SetCanSignUp sets CanSignUp field to given value.

### HasCanSignUp

`func (o *IamProviderItem) HasCanSignUp() bool`

HasCanSignUp returns a boolean if a field has been set.

### GetCanUnlink

`func (o *IamProviderItem) GetCanUnlink() bool`

GetCanUnlink returns the CanUnlink field if non-nil, zero value otherwise.

### GetCanUnlinkOk

`func (o *IamProviderItem) GetCanUnlinkOk() (*bool, bool)`

GetCanUnlinkOk returns a tuple with the CanUnlink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanUnlink

`func (o *IamProviderItem) SetCanUnlink(v bool)`

SetCanUnlink sets CanUnlink field to given value.

### HasCanUnlink

`func (o *IamProviderItem) HasCanUnlink() bool`

HasCanUnlink returns a boolean if a field has been set.

### GetCountryCodes

`func (o *IamProviderItem) GetCountryCodes() []string`

GetCountryCodes returns the CountryCodes field if non-nil, zero value otherwise.

### GetCountryCodesOk

`func (o *IamProviderItem) GetCountryCodesOk() (*[]string, bool)`

GetCountryCodesOk returns a tuple with the CountryCodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCodes

`func (o *IamProviderItem) SetCountryCodes(v []string)`

SetCountryCodes sets CountryCodes field to given value.

### HasCountryCodes

`func (o *IamProviderItem) HasCountryCodes() bool`

HasCountryCodes returns a boolean if a field has been set.

### GetName

`func (o *IamProviderItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IamProviderItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IamProviderItem) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IamProviderItem) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOwner

`func (o *IamProviderItem) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *IamProviderItem) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *IamProviderItem) SetOwner(v string)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *IamProviderItem) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetPrompted

`func (o *IamProviderItem) GetPrompted() bool`

GetPrompted returns the Prompted field if non-nil, zero value otherwise.

### GetPromptedOk

`func (o *IamProviderItem) GetPromptedOk() (*bool, bool)`

GetPromptedOk returns a tuple with the Prompted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrompted

`func (o *IamProviderItem) SetPrompted(v bool)`

SetPrompted sets Prompted field to given value.

### HasPrompted

`func (o *IamProviderItem) HasPrompted() bool`

HasPrompted returns a boolean if a field has been set.

### GetProvider

`func (o *IamProviderItem) GetProvider() IamProvider`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *IamProviderItem) GetProviderOk() (*IamProvider, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *IamProviderItem) SetProvider(v IamProvider)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *IamProviderItem) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetRule

`func (o *IamProviderItem) GetRule() string`

GetRule returns the Rule field if non-nil, zero value otherwise.

### GetRuleOk

`func (o *IamProviderItem) GetRuleOk() (*string, bool)`

GetRuleOk returns a tuple with the Rule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRule

`func (o *IamProviderItem) SetRule(v string)`

SetRule sets Rule field to given value.

### HasRule

`func (o *IamProviderItem) HasRule() bool`

HasRule returns a boolean if a field has been set.

### GetSignupGroup

`func (o *IamProviderItem) GetSignupGroup() string`

GetSignupGroup returns the SignupGroup field if non-nil, zero value otherwise.

### GetSignupGroupOk

`func (o *IamProviderItem) GetSignupGroupOk() (*string, bool)`

GetSignupGroupOk returns a tuple with the SignupGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignupGroup

`func (o *IamProviderItem) SetSignupGroup(v string)`

SetSignupGroup sets SignupGroup field to given value.

### HasSignupGroup

`func (o *IamProviderItem) HasSignupGroup() bool`

HasSignupGroup returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


