# CloudVariant

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Framework** | Pointer to **string** | only when it differs from the template&#39;s | [optional] 
**Id** | Pointer to **string** | selector, unique within the template (\&quot;react\&quot;, \&quot;grid-3-fluid\&quot;) | [optional] 
**Kind** | Pointer to **string** | the axis it varies: format | page | theme | [optional] 
**Label** | Pointer to **string** | human label for the picker | [optional] 
**Source** | Pointer to **string** | the repository this shape is forked from; the synthesized default shape carries the template&#39;s own | [optional] 

## Methods

### NewCloudVariant

`func NewCloudVariant() *CloudVariant`

NewCloudVariant instantiates a new CloudVariant object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudVariantWithDefaults

`func NewCloudVariantWithDefaults() *CloudVariant`

NewCloudVariantWithDefaults instantiates a new CloudVariant object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFramework

`func (o *CloudVariant) GetFramework() string`

GetFramework returns the Framework field if non-nil, zero value otherwise.

### GetFrameworkOk

`func (o *CloudVariant) GetFrameworkOk() (*string, bool)`

GetFrameworkOk returns a tuple with the Framework field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFramework

`func (o *CloudVariant) SetFramework(v string)`

SetFramework sets Framework field to given value.

### HasFramework

`func (o *CloudVariant) HasFramework() bool`

HasFramework returns a boolean if a field has been set.

### GetId

`func (o *CloudVariant) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CloudVariant) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CloudVariant) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CloudVariant) HasId() bool`

HasId returns a boolean if a field has been set.

### GetKind

`func (o *CloudVariant) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *CloudVariant) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *CloudVariant) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *CloudVariant) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetLabel

`func (o *CloudVariant) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *CloudVariant) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *CloudVariant) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *CloudVariant) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetSource

`func (o *CloudVariant) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *CloudVariant) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *CloudVariant) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *CloudVariant) HasSource() bool`

HasSource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


