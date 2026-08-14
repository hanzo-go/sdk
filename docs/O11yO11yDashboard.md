# O11yO11yDashboard

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **time.Time** | CreatedAt is when the dashboard was created. | [optional] 
**CreatedBy** | Pointer to **string** | CreatedBy is who created it. | [optional] 
**Id** | Pointer to **string** | ID is the dashboard&#39;s id. | [optional] 
**Image** | Pointer to **string** | Image is an optional cover image reference. | [optional] 
**Locked** | Pointer to **bool** | Locked reports whether the dashboard is locked against edits. | [optional] 
**Name** | Pointer to **string** | Name is the dashboard&#39;s unique internal name. | [optional] 
**OrgId** | Pointer to **string** | OrgID is the org the dashboard belongs to. | [optional] 
**SchemaVersion** | Pointer to **string** | SchemaVersion is the dashboard schema version. | [optional] 
**Source** | Pointer to **string** | Source is where the dashboard came from: user, system or integration. | [optional] 
**Spec** | Pointer to **interface{}** |  | [optional] 
**Tags** | Pointer to [**[]O11yO11yDashboardTag**](O11yO11yDashboardTag.md) | Tags are the dashboard&#39;s tags. | [optional] 
**UpdatedAt** | Pointer to **time.Time** | UpdatedAt is when it last changed. | [optional] 
**UpdatedBy** | Pointer to **string** | UpdatedBy is who last changed it. | [optional] 

## Methods

### NewO11yO11yDashboard

`func NewO11yO11yDashboard() *O11yO11yDashboard`

NewO11yO11yDashboard instantiates a new O11yO11yDashboard object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yO11yDashboardWithDefaults

`func NewO11yO11yDashboardWithDefaults() *O11yO11yDashboard`

NewO11yO11yDashboardWithDefaults instantiates a new O11yO11yDashboard object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *O11yO11yDashboard) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *O11yO11yDashboard) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *O11yO11yDashboard) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *O11yO11yDashboard) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCreatedBy

`func (o *O11yO11yDashboard) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *O11yO11yDashboard) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *O11yO11yDashboard) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *O11yO11yDashboard) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetId

`func (o *O11yO11yDashboard) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *O11yO11yDashboard) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *O11yO11yDashboard) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *O11yO11yDashboard) HasId() bool`

HasId returns a boolean if a field has been set.

### GetImage

`func (o *O11yO11yDashboard) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *O11yO11yDashboard) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *O11yO11yDashboard) SetImage(v string)`

SetImage sets Image field to given value.

### HasImage

`func (o *O11yO11yDashboard) HasImage() bool`

HasImage returns a boolean if a field has been set.

### GetLocked

`func (o *O11yO11yDashboard) GetLocked() bool`

GetLocked returns the Locked field if non-nil, zero value otherwise.

### GetLockedOk

`func (o *O11yO11yDashboard) GetLockedOk() (*bool, bool)`

GetLockedOk returns a tuple with the Locked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocked

`func (o *O11yO11yDashboard) SetLocked(v bool)`

SetLocked sets Locked field to given value.

### HasLocked

`func (o *O11yO11yDashboard) HasLocked() bool`

HasLocked returns a boolean if a field has been set.

### GetName

`func (o *O11yO11yDashboard) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *O11yO11yDashboard) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *O11yO11yDashboard) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *O11yO11yDashboard) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOrgId

`func (o *O11yO11yDashboard) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *O11yO11yDashboard) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *O11yO11yDashboard) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *O11yO11yDashboard) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### GetSchemaVersion

`func (o *O11yO11yDashboard) GetSchemaVersion() string`

GetSchemaVersion returns the SchemaVersion field if non-nil, zero value otherwise.

### GetSchemaVersionOk

`func (o *O11yO11yDashboard) GetSchemaVersionOk() (*string, bool)`

GetSchemaVersionOk returns a tuple with the SchemaVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaVersion

`func (o *O11yO11yDashboard) SetSchemaVersion(v string)`

SetSchemaVersion sets SchemaVersion field to given value.

### HasSchemaVersion

`func (o *O11yO11yDashboard) HasSchemaVersion() bool`

HasSchemaVersion returns a boolean if a field has been set.

### GetSource

`func (o *O11yO11yDashboard) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *O11yO11yDashboard) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *O11yO11yDashboard) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *O11yO11yDashboard) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetSpec

`func (o *O11yO11yDashboard) GetSpec() interface{}`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *O11yO11yDashboard) GetSpecOk() (*interface{}, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *O11yO11yDashboard) SetSpec(v interface{})`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *O11yO11yDashboard) HasSpec() bool`

HasSpec returns a boolean if a field has been set.

### SetSpecNil

`func (o *O11yO11yDashboard) SetSpecNil(b bool)`

 SetSpecNil sets the value for Spec to be an explicit nil

### UnsetSpec
`func (o *O11yO11yDashboard) UnsetSpec()`

UnsetSpec ensures that no value is present for Spec, not even an explicit nil
### GetTags

`func (o *O11yO11yDashboard) GetTags() []O11yO11yDashboardTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *O11yO11yDashboard) GetTagsOk() (*[]O11yO11yDashboardTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *O11yO11yDashboard) SetTags(v []O11yO11yDashboardTag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *O11yO11yDashboard) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *O11yO11yDashboard) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *O11yO11yDashboard) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *O11yO11yDashboard) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *O11yO11yDashboard) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUpdatedBy

`func (o *O11yO11yDashboard) GetUpdatedBy() string`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *O11yO11yDashboard) GetUpdatedByOk() (*string, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *O11yO11yDashboard) SetUpdatedBy(v string)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *O11yO11yDashboard) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


