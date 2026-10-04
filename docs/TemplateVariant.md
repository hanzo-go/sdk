# TemplateVariant

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Framework** | Pointer to **string** | only when it differs from the template&#39;s | [optional] 
**Id** | Pointer to **string** | selector, unique within the template (\&quot;react\&quot;, \&quot;grid-3-fluid\&quot;) | [optional] 
**Kind** | Pointer to **string** | the axis it varies: format | page | theme | [optional] 
**Label** | Pointer to **string** | human label for the picker | [optional] 
**Source** | Pointer to **string** | the repository this shape is forked from; the synthesized default shape carries the template&#39;s own | [optional] 

## Methods

### NewTemplateVariant

`func NewTemplateVariant() *TemplateVariant`

NewTemplateVariant instantiates a new TemplateVariant object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTemplateVariantWithDefaults

`func NewTemplateVariantWithDefaults() *TemplateVariant`

NewTemplateVariantWithDefaults instantiates a new TemplateVariant object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFramework

`func (o *TemplateVariant) GetFramework() string`

GetFramework returns the Framework field if non-nil, zero value otherwise.

### GetFrameworkOk

`func (o *TemplateVariant) GetFrameworkOk() (*string, bool)`

GetFrameworkOk returns a tuple with the Framework field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFramework

`func (o *TemplateVariant) SetFramework(v string)`

SetFramework sets Framework field to given value.

### HasFramework

`func (o *TemplateVariant) HasFramework() bool`

HasFramework returns a boolean if a field has been set.

### GetId

`func (o *TemplateVariant) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TemplateVariant) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TemplateVariant) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TemplateVariant) HasId() bool`

HasId returns a boolean if a field has been set.

### GetKind

`func (o *TemplateVariant) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *TemplateVariant) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *TemplateVariant) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *TemplateVariant) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetLabel

`func (o *TemplateVariant) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *TemplateVariant) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *TemplateVariant) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *TemplateVariant) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetSource

`func (o *TemplateVariant) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *TemplateVariant) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *TemplateVariant) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *TemplateVariant) HasSource() bool`

HasSource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


