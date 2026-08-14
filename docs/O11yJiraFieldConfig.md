# O11yJiraFieldConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnableUpdate** | Pointer to **bool** | EnableUpdate indicates whether this field should be omitted when updating an existing issue. | [optional] 
**Template** | Pointer to **string** | Template is the template string used to render the field. | [optional] 

## Methods

### NewO11yJiraFieldConfig

`func NewO11yJiraFieldConfig() *O11yJiraFieldConfig`

NewO11yJiraFieldConfig instantiates a new O11yJiraFieldConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yJiraFieldConfigWithDefaults

`func NewO11yJiraFieldConfigWithDefaults() *O11yJiraFieldConfig`

NewO11yJiraFieldConfigWithDefaults instantiates a new O11yJiraFieldConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnableUpdate

`func (o *O11yJiraFieldConfig) GetEnableUpdate() bool`

GetEnableUpdate returns the EnableUpdate field if non-nil, zero value otherwise.

### GetEnableUpdateOk

`func (o *O11yJiraFieldConfig) GetEnableUpdateOk() (*bool, bool)`

GetEnableUpdateOk returns a tuple with the EnableUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableUpdate

`func (o *O11yJiraFieldConfig) SetEnableUpdate(v bool)`

SetEnableUpdate sets EnableUpdate field to given value.

### HasEnableUpdate

`func (o *O11yJiraFieldConfig) HasEnableUpdate() bool`

HasEnableUpdate returns a boolean if a field has been set.

### GetTemplate

`func (o *O11yJiraFieldConfig) GetTemplate() string`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *O11yJiraFieldConfig) GetTemplateOk() (*string, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *O11yJiraFieldConfig) SetTemplate(v string)`

SetTemplate sets Template field to given value.

### HasTemplate

`func (o *O11yJiraFieldConfig) HasTemplate() bool`

HasTemplate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


