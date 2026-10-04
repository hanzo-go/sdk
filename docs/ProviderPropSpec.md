# ProviderPropSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | Description says what the value is. | [optional] 
**DisplayName** | Pointer to **string** | DisplayName is what a form calls it. | [optional] 
**Name** | Pointer to **string** | Name is the input&#39;s key in the action&#39;s input object. | [optional] 
**Required** | Pointer to **bool** | Required is whether the action refuses to run without it. | [optional] 
**Type** | Pointer to **string** | Type is the JSON type the value takes: string|number|boolean|object|array. | [optional] 

## Methods

### NewProviderPropSpec

`func NewProviderPropSpec() *ProviderPropSpec`

NewProviderPropSpec instantiates a new ProviderPropSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProviderPropSpecWithDefaults

`func NewProviderPropSpecWithDefaults() *ProviderPropSpec`

NewProviderPropSpecWithDefaults instantiates a new ProviderPropSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *ProviderPropSpec) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ProviderPropSpec) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ProviderPropSpec) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ProviderPropSpec) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDisplayName

`func (o *ProviderPropSpec) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *ProviderPropSpec) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *ProviderPropSpec) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *ProviderPropSpec) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetName

`func (o *ProviderPropSpec) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProviderPropSpec) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProviderPropSpec) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ProviderPropSpec) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRequired

`func (o *ProviderPropSpec) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *ProviderPropSpec) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *ProviderPropSpec) SetRequired(v bool)`

SetRequired sets Required field to given value.

### HasRequired

`func (o *ProviderPropSpec) HasRequired() bool`

HasRequired returns a boolean if a field has been set.

### GetType

`func (o *ProviderPropSpec) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ProviderPropSpec) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ProviderPropSpec) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ProviderPropSpec) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


