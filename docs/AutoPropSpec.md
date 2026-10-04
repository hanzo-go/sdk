# AutoPropSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | Description says what the value is. | [optional] 
**DisplayName** | Pointer to **string** | DisplayName is what a form calls it. | [optional] 
**Name** | Pointer to **string** | Name is the input&#39;s key in the action&#39;s input object. | [optional] 
**Required** | Pointer to **bool** | Required is whether the action refuses to run without it. | [optional] 
**Type** | Pointer to **string** | Type is the JSON type the value takes: string|number|boolean|object|array. | [optional] 

## Methods

### NewAutoPropSpec

`func NewAutoPropSpec() *AutoPropSpec`

NewAutoPropSpec instantiates a new AutoPropSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutoPropSpecWithDefaults

`func NewAutoPropSpecWithDefaults() *AutoPropSpec`

NewAutoPropSpecWithDefaults instantiates a new AutoPropSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *AutoPropSpec) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AutoPropSpec) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AutoPropSpec) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AutoPropSpec) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDisplayName

`func (o *AutoPropSpec) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *AutoPropSpec) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *AutoPropSpec) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *AutoPropSpec) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetName

`func (o *AutoPropSpec) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AutoPropSpec) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AutoPropSpec) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AutoPropSpec) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRequired

`func (o *AutoPropSpec) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *AutoPropSpec) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *AutoPropSpec) SetRequired(v bool)`

SetRequired sets Required field to given value.

### HasRequired

`func (o *AutoPropSpec) HasRequired() bool`

HasRequired returns a boolean if a field has been set.

### GetType

`func (o *AutoPropSpec) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AutoPropSpec) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AutoPropSpec) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *AutoPropSpec) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


