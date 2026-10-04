# AiLimitsSet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreditsAfterAllowance** | Pointer to **bool** | CreditsAfterAllowance keeps a model working once the plan&#39;s included usage of it is spent, paid from prepaid credit (and granted credit where the model takes it), when true; when false the conversation moves to a Hanzo model and other calls are refused until the plan resets. Off until the payer turns it on. | [optional] 

## Methods

### NewAiLimitsSet

`func NewAiLimitsSet() *AiLimitsSet`

NewAiLimitsSet instantiates a new AiLimitsSet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiLimitsSetWithDefaults

`func NewAiLimitsSetWithDefaults() *AiLimitsSet`

NewAiLimitsSetWithDefaults instantiates a new AiLimitsSet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreditsAfterAllowance

`func (o *AiLimitsSet) GetCreditsAfterAllowance() bool`

GetCreditsAfterAllowance returns the CreditsAfterAllowance field if non-nil, zero value otherwise.

### GetCreditsAfterAllowanceOk

`func (o *AiLimitsSet) GetCreditsAfterAllowanceOk() (*bool, bool)`

GetCreditsAfterAllowanceOk returns a tuple with the CreditsAfterAllowance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditsAfterAllowance

`func (o *AiLimitsSet) SetCreditsAfterAllowance(v bool)`

SetCreditsAfterAllowance sets CreditsAfterAllowance field to given value.

### HasCreditsAfterAllowance

`func (o *AiLimitsSet) HasCreditsAfterAllowance() bool`

HasCreditsAfterAllowance returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


