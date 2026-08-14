# O11yO11yDashboardPostable

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GenerateName** | Pointer to **bool** | GenerateName derives a fresh unique name from spec.display.name instead of taking Name. | [optional] 
**Image** | Pointer to **string** | Image is an optional cover image reference. | [optional] 
**Name** | Pointer to **string** | Name is the dashboard&#39;s unique internal name (a DNS-1123 label). Omit it with generateName to derive one from the display name. | [optional] 
**SchemaVersion** | Pointer to **string** | SchemaVersion is the dashboard schema version; must be the current v6. | [optional] 
**Spec** | Pointer to **interface{}** |  | [optional] 
**Tags** | Pointer to [**[]O11yO11yDashboardPostableTag**](O11yO11yDashboardPostableTag.md) | Tags are the dashboard&#39;s tags; at most ten, and none may use a reserved DSL key. | [optional] 

## Methods

### NewO11yO11yDashboardPostable

`func NewO11yO11yDashboardPostable() *O11yO11yDashboardPostable`

NewO11yO11yDashboardPostable instantiates a new O11yO11yDashboardPostable object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yO11yDashboardPostableWithDefaults

`func NewO11yO11yDashboardPostableWithDefaults() *O11yO11yDashboardPostable`

NewO11yO11yDashboardPostableWithDefaults instantiates a new O11yO11yDashboardPostable object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGenerateName

`func (o *O11yO11yDashboardPostable) GetGenerateName() bool`

GetGenerateName returns the GenerateName field if non-nil, zero value otherwise.

### GetGenerateNameOk

`func (o *O11yO11yDashboardPostable) GetGenerateNameOk() (*bool, bool)`

GetGenerateNameOk returns a tuple with the GenerateName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenerateName

`func (o *O11yO11yDashboardPostable) SetGenerateName(v bool)`

SetGenerateName sets GenerateName field to given value.

### HasGenerateName

`func (o *O11yO11yDashboardPostable) HasGenerateName() bool`

HasGenerateName returns a boolean if a field has been set.

### GetImage

`func (o *O11yO11yDashboardPostable) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *O11yO11yDashboardPostable) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *O11yO11yDashboardPostable) SetImage(v string)`

SetImage sets Image field to given value.

### HasImage

`func (o *O11yO11yDashboardPostable) HasImage() bool`

HasImage returns a boolean if a field has been set.

### GetName

`func (o *O11yO11yDashboardPostable) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *O11yO11yDashboardPostable) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *O11yO11yDashboardPostable) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *O11yO11yDashboardPostable) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSchemaVersion

`func (o *O11yO11yDashboardPostable) GetSchemaVersion() string`

GetSchemaVersion returns the SchemaVersion field if non-nil, zero value otherwise.

### GetSchemaVersionOk

`func (o *O11yO11yDashboardPostable) GetSchemaVersionOk() (*string, bool)`

GetSchemaVersionOk returns a tuple with the SchemaVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaVersion

`func (o *O11yO11yDashboardPostable) SetSchemaVersion(v string)`

SetSchemaVersion sets SchemaVersion field to given value.

### HasSchemaVersion

`func (o *O11yO11yDashboardPostable) HasSchemaVersion() bool`

HasSchemaVersion returns a boolean if a field has been set.

### GetSpec

`func (o *O11yO11yDashboardPostable) GetSpec() interface{}`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *O11yO11yDashboardPostable) GetSpecOk() (*interface{}, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *O11yO11yDashboardPostable) SetSpec(v interface{})`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *O11yO11yDashboardPostable) HasSpec() bool`

HasSpec returns a boolean if a field has been set.

### SetSpecNil

`func (o *O11yO11yDashboardPostable) SetSpecNil(b bool)`

 SetSpecNil sets the value for Spec to be an explicit nil

### UnsetSpec
`func (o *O11yO11yDashboardPostable) UnsetSpec()`

UnsetSpec ensures that no value is present for Spec, not even an explicit nil
### GetTags

`func (o *O11yO11yDashboardPostable) GetTags() []O11yO11yDashboardPostableTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *O11yO11yDashboardPostable) GetTagsOk() (*[]O11yO11yDashboardPostableTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *O11yO11yDashboardPostable) SetTags(v []O11yO11yDashboardPostableTag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *O11yO11yDashboardPostable) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


