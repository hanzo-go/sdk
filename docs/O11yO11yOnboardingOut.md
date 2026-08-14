# O11yO11yOnboardingOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**O11yO11yK8sOnboarding**](O11yO11yK8sOnboarding.md) | Data is the progress report. | [optional] 
**Status** | Pointer to **string** | Status is \&quot;success\&quot;. | [optional] 

## Methods

### NewO11yO11yOnboardingOut

`func NewO11yO11yOnboardingOut() *O11yO11yOnboardingOut`

NewO11yO11yOnboardingOut instantiates a new O11yO11yOnboardingOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yO11yOnboardingOutWithDefaults

`func NewO11yO11yOnboardingOutWithDefaults() *O11yO11yOnboardingOut`

NewO11yO11yOnboardingOutWithDefaults instantiates a new O11yO11yOnboardingOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *O11yO11yOnboardingOut) GetData() O11yO11yK8sOnboarding`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *O11yO11yOnboardingOut) GetDataOk() (*O11yO11yK8sOnboarding, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *O11yO11yOnboardingOut) SetData(v O11yO11yK8sOnboarding)`

SetData sets Data field to given value.

### HasData

`func (o *O11yO11yOnboardingOut) HasData() bool`

HasData returns a boolean if a field has been set.

### GetStatus

`func (o *O11yO11yOnboardingOut) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *O11yO11yOnboardingOut) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *O11yO11yOnboardingOut) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *O11yO11yOnboardingOut) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


