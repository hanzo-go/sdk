# ProviderField

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Default** | Pointer to **string** | Default is what an empty input stands for. | [optional] 
**Help** | Pointer to **string** | Help is the connector&#39;s own instruction for finding the value. | [optional] 
**Label** | Pointer to **string** | Label is what the input is called on the form. | [optional] 
**Name** | Pointer to **string** | Name is the key the value is submitted under. | [optional] 
**Options** | Pointer to [**[]ProviderFieldOption**](ProviderFieldOption.md) | Options are the values a choice field accepts. | [optional] 
**Required** | Pointer to **bool** | Required is whether the form refuses to submit without it. | [optional] 
**Secret** | Pointer to **bool** | Secret masks the input; its value is sealed and never read back. | [optional] 
**Type** | Pointer to **string** | Type is \&quot;number\&quot;, \&quot;boolean\&quot; or \&quot;text\&quot; (multi-line) when the value is not a single line of text. | [optional] 

## Methods

### NewProviderField

`func NewProviderField() *ProviderField`

NewProviderField instantiates a new ProviderField object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProviderFieldWithDefaults

`func NewProviderFieldWithDefaults() *ProviderField`

NewProviderFieldWithDefaults instantiates a new ProviderField object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefault

`func (o *ProviderField) GetDefault() string`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *ProviderField) GetDefaultOk() (*string, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *ProviderField) SetDefault(v string)`

SetDefault sets Default field to given value.

### HasDefault

`func (o *ProviderField) HasDefault() bool`

HasDefault returns a boolean if a field has been set.

### GetHelp

`func (o *ProviderField) GetHelp() string`

GetHelp returns the Help field if non-nil, zero value otherwise.

### GetHelpOk

`func (o *ProviderField) GetHelpOk() (*string, bool)`

GetHelpOk returns a tuple with the Help field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHelp

`func (o *ProviderField) SetHelp(v string)`

SetHelp sets Help field to given value.

### HasHelp

`func (o *ProviderField) HasHelp() bool`

HasHelp returns a boolean if a field has been set.

### GetLabel

`func (o *ProviderField) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *ProviderField) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *ProviderField) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *ProviderField) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetName

`func (o *ProviderField) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProviderField) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProviderField) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ProviderField) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOptions

`func (o *ProviderField) GetOptions() []ProviderFieldOption`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *ProviderField) GetOptionsOk() (*[]ProviderFieldOption, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *ProviderField) SetOptions(v []ProviderFieldOption)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *ProviderField) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### GetRequired

`func (o *ProviderField) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *ProviderField) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *ProviderField) SetRequired(v bool)`

SetRequired sets Required field to given value.

### HasRequired

`func (o *ProviderField) HasRequired() bool`

HasRequired returns a boolean if a field has been set.

### GetSecret

`func (o *ProviderField) GetSecret() bool`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *ProviderField) GetSecretOk() (*bool, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *ProviderField) SetSecret(v bool)`

SetSecret sets Secret field to given value.

### HasSecret

`func (o *ProviderField) HasSecret() bool`

HasSecret returns a boolean if a field has been set.

### GetType

`func (o *ProviderField) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ProviderField) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ProviderField) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ProviderField) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


