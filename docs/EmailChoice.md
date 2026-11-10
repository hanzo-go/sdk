# EmailChoice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Provider** | Pointer to **string** | Provider is ses or sendgrid. | [optional] 

## Methods

### NewEmailChoice

`func NewEmailChoice() *EmailChoice`

NewEmailChoice instantiates a new EmailChoice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmailChoiceWithDefaults

`func NewEmailChoiceWithDefaults() *EmailChoice`

NewEmailChoiceWithDefaults instantiates a new EmailChoice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProvider

`func (o *EmailChoice) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *EmailChoice) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *EmailChoice) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *EmailChoice) HasProvider() bool`

HasProvider returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


