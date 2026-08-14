# O11yO11yDashboardUpdateIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | ID is the dashboard id from the path. | [optional] 
**Image** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**SchemaVersion** | Pointer to **string** |  | [optional] 
**Spec** | Pointer to **interface{}** |  | [optional] 
**Tags** | Pointer to [**[]O11yO11yDashboardPostableTag**](O11yO11yDashboardPostableTag.md) |  | [optional] 

## Methods

### NewO11yO11yDashboardUpdateIn

`func NewO11yO11yDashboardUpdateIn() *O11yO11yDashboardUpdateIn`

NewO11yO11yDashboardUpdateIn instantiates a new O11yO11yDashboardUpdateIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yO11yDashboardUpdateInWithDefaults

`func NewO11yO11yDashboardUpdateInWithDefaults() *O11yO11yDashboardUpdateIn`

NewO11yO11yDashboardUpdateInWithDefaults instantiates a new O11yO11yDashboardUpdateIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *O11yO11yDashboardUpdateIn) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *O11yO11yDashboardUpdateIn) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *O11yO11yDashboardUpdateIn) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *O11yO11yDashboardUpdateIn) HasId() bool`

HasId returns a boolean if a field has been set.

### GetImage

`func (o *O11yO11yDashboardUpdateIn) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *O11yO11yDashboardUpdateIn) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *O11yO11yDashboardUpdateIn) SetImage(v string)`

SetImage sets Image field to given value.

### HasImage

`func (o *O11yO11yDashboardUpdateIn) HasImage() bool`

HasImage returns a boolean if a field has been set.

### GetName

`func (o *O11yO11yDashboardUpdateIn) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *O11yO11yDashboardUpdateIn) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *O11yO11yDashboardUpdateIn) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *O11yO11yDashboardUpdateIn) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSchemaVersion

`func (o *O11yO11yDashboardUpdateIn) GetSchemaVersion() string`

GetSchemaVersion returns the SchemaVersion field if non-nil, zero value otherwise.

### GetSchemaVersionOk

`func (o *O11yO11yDashboardUpdateIn) GetSchemaVersionOk() (*string, bool)`

GetSchemaVersionOk returns a tuple with the SchemaVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaVersion

`func (o *O11yO11yDashboardUpdateIn) SetSchemaVersion(v string)`

SetSchemaVersion sets SchemaVersion field to given value.

### HasSchemaVersion

`func (o *O11yO11yDashboardUpdateIn) HasSchemaVersion() bool`

HasSchemaVersion returns a boolean if a field has been set.

### GetSpec

`func (o *O11yO11yDashboardUpdateIn) GetSpec() interface{}`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *O11yO11yDashboardUpdateIn) GetSpecOk() (*interface{}, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *O11yO11yDashboardUpdateIn) SetSpec(v interface{})`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *O11yO11yDashboardUpdateIn) HasSpec() bool`

HasSpec returns a boolean if a field has been set.

### SetSpecNil

`func (o *O11yO11yDashboardUpdateIn) SetSpecNil(b bool)`

 SetSpecNil sets the value for Spec to be an explicit nil

### UnsetSpec
`func (o *O11yO11yDashboardUpdateIn) UnsetSpec()`

UnsetSpec ensures that no value is present for Spec, not even an explicit nil
### GetTags

`func (o *O11yO11yDashboardUpdateIn) GetTags() []O11yO11yDashboardPostableTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *O11yO11yDashboardUpdateIn) GetTagsOk() (*[]O11yO11yDashboardPostableTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *O11yO11yDashboardUpdateIn) SetTags(v []O11yO11yDashboardPostableTag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *O11yO11yDashboardUpdateIn) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


