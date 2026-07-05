# IamObjectProviderItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CanSignIn** | Pointer to **bool** |  | [optional] 
**CanSignUp** | Pointer to **bool** |  | [optional] 
**CanUnlink** | Pointer to **bool** |  | [optional] 
**CountryCodes** | Pointer to **[]string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Owner** | Pointer to **string** |  | [optional] 
**Prompted** | Pointer to **bool** |  | [optional] 
**Provider** | Pointer to [**IamObjectProvider**](IamObjectProvider.md) |  | [optional] 
**Rule** | Pointer to **string** |  | [optional] 
**SignupGroup** | Pointer to **string** |  | [optional] 

## Methods

### NewIamObjectProviderItem

`func NewIamObjectProviderItem() *IamObjectProviderItem`

NewIamObjectProviderItem instantiates a new IamObjectProviderItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamObjectProviderItemWithDefaults

`func NewIamObjectProviderItemWithDefaults() *IamObjectProviderItem`

NewIamObjectProviderItemWithDefaults instantiates a new IamObjectProviderItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCanSignIn

`func (o *IamObjectProviderItem) GetCanSignIn() bool`

GetCanSignIn returns the CanSignIn field if non-nil, zero value otherwise.

### GetCanSignInOk

`func (o *IamObjectProviderItem) GetCanSignInOk() (*bool, bool)`

GetCanSignInOk returns a tuple with the CanSignIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanSignIn

`func (o *IamObjectProviderItem) SetCanSignIn(v bool)`

SetCanSignIn sets CanSignIn field to given value.

### HasCanSignIn

`func (o *IamObjectProviderItem) HasCanSignIn() bool`

HasCanSignIn returns a boolean if a field has been set.

### GetCanSignUp

`func (o *IamObjectProviderItem) GetCanSignUp() bool`

GetCanSignUp returns the CanSignUp field if non-nil, zero value otherwise.

### GetCanSignUpOk

`func (o *IamObjectProviderItem) GetCanSignUpOk() (*bool, bool)`

GetCanSignUpOk returns a tuple with the CanSignUp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanSignUp

`func (o *IamObjectProviderItem) SetCanSignUp(v bool)`

SetCanSignUp sets CanSignUp field to given value.

### HasCanSignUp

`func (o *IamObjectProviderItem) HasCanSignUp() bool`

HasCanSignUp returns a boolean if a field has been set.

### GetCanUnlink

`func (o *IamObjectProviderItem) GetCanUnlink() bool`

GetCanUnlink returns the CanUnlink field if non-nil, zero value otherwise.

### GetCanUnlinkOk

`func (o *IamObjectProviderItem) GetCanUnlinkOk() (*bool, bool)`

GetCanUnlinkOk returns a tuple with the CanUnlink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanUnlink

`func (o *IamObjectProviderItem) SetCanUnlink(v bool)`

SetCanUnlink sets CanUnlink field to given value.

### HasCanUnlink

`func (o *IamObjectProviderItem) HasCanUnlink() bool`

HasCanUnlink returns a boolean if a field has been set.

### GetCountryCodes

`func (o *IamObjectProviderItem) GetCountryCodes() []string`

GetCountryCodes returns the CountryCodes field if non-nil, zero value otherwise.

### GetCountryCodesOk

`func (o *IamObjectProviderItem) GetCountryCodesOk() (*[]string, bool)`

GetCountryCodesOk returns a tuple with the CountryCodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCodes

`func (o *IamObjectProviderItem) SetCountryCodes(v []string)`

SetCountryCodes sets CountryCodes field to given value.

### HasCountryCodes

`func (o *IamObjectProviderItem) HasCountryCodes() bool`

HasCountryCodes returns a boolean if a field has been set.

### GetName

`func (o *IamObjectProviderItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IamObjectProviderItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IamObjectProviderItem) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IamObjectProviderItem) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOwner

`func (o *IamObjectProviderItem) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *IamObjectProviderItem) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *IamObjectProviderItem) SetOwner(v string)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *IamObjectProviderItem) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetPrompted

`func (o *IamObjectProviderItem) GetPrompted() bool`

GetPrompted returns the Prompted field if non-nil, zero value otherwise.

### GetPromptedOk

`func (o *IamObjectProviderItem) GetPromptedOk() (*bool, bool)`

GetPromptedOk returns a tuple with the Prompted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrompted

`func (o *IamObjectProviderItem) SetPrompted(v bool)`

SetPrompted sets Prompted field to given value.

### HasPrompted

`func (o *IamObjectProviderItem) HasPrompted() bool`

HasPrompted returns a boolean if a field has been set.

### GetProvider

`func (o *IamObjectProviderItem) GetProvider() IamObjectProvider`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *IamObjectProviderItem) GetProviderOk() (*IamObjectProvider, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *IamObjectProviderItem) SetProvider(v IamObjectProvider)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *IamObjectProviderItem) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetRule

`func (o *IamObjectProviderItem) GetRule() string`

GetRule returns the Rule field if non-nil, zero value otherwise.

### GetRuleOk

`func (o *IamObjectProviderItem) GetRuleOk() (*string, bool)`

GetRuleOk returns a tuple with the Rule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRule

`func (o *IamObjectProviderItem) SetRule(v string)`

SetRule sets Rule field to given value.

### HasRule

`func (o *IamObjectProviderItem) HasRule() bool`

HasRule returns a boolean if a field has been set.

### GetSignupGroup

`func (o *IamObjectProviderItem) GetSignupGroup() string`

GetSignupGroup returns the SignupGroup field if non-nil, zero value otherwise.

### GetSignupGroupOk

`func (o *IamObjectProviderItem) GetSignupGroupOk() (*string, bool)`

GetSignupGroupOk returns a tuple with the SignupGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignupGroup

`func (o *IamObjectProviderItem) SetSignupGroup(v string)`

SetSignupGroup sets SignupGroup field to given value.

### HasSignupGroup

`func (o *IamObjectProviderItem) HasSignupGroup() bool`

HasSignupGroup returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


