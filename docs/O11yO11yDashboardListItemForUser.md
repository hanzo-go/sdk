# O11yO11yDashboardListItemForUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**CreatedBy** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Image** | Pointer to **string** |  | [optional] 
**Locked** | Pointer to **bool** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**OrgId** | Pointer to **string** |  | [optional] 
**Pinned** | Pointer to **bool** | Pinned reports whether the calling user has pinned this dashboard. | [optional] 
**SchemaVersion** | Pointer to **string** |  | [optional] 
**Source** | Pointer to **string** |  | [optional] 
**Spec** | Pointer to [**O11yO11yDashboardListSpec**](O11yO11yDashboardListSpec.md) |  | [optional] 
**Tags** | Pointer to [**[]O11yO11yDashboardTag**](O11yO11yDashboardTag.md) |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedBy** | Pointer to **string** |  | [optional] 

## Methods

### NewO11yO11yDashboardListItemForUser

`func NewO11yO11yDashboardListItemForUser() *O11yO11yDashboardListItemForUser`

NewO11yO11yDashboardListItemForUser instantiates a new O11yO11yDashboardListItemForUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yO11yDashboardListItemForUserWithDefaults

`func NewO11yO11yDashboardListItemForUserWithDefaults() *O11yO11yDashboardListItemForUser`

NewO11yO11yDashboardListItemForUserWithDefaults instantiates a new O11yO11yDashboardListItemForUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *O11yO11yDashboardListItemForUser) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *O11yO11yDashboardListItemForUser) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *O11yO11yDashboardListItemForUser) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *O11yO11yDashboardListItemForUser) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCreatedBy

`func (o *O11yO11yDashboardListItemForUser) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *O11yO11yDashboardListItemForUser) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *O11yO11yDashboardListItemForUser) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *O11yO11yDashboardListItemForUser) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetId

`func (o *O11yO11yDashboardListItemForUser) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *O11yO11yDashboardListItemForUser) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *O11yO11yDashboardListItemForUser) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *O11yO11yDashboardListItemForUser) HasId() bool`

HasId returns a boolean if a field has been set.

### GetImage

`func (o *O11yO11yDashboardListItemForUser) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *O11yO11yDashboardListItemForUser) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *O11yO11yDashboardListItemForUser) SetImage(v string)`

SetImage sets Image field to given value.

### HasImage

`func (o *O11yO11yDashboardListItemForUser) HasImage() bool`

HasImage returns a boolean if a field has been set.

### GetLocked

`func (o *O11yO11yDashboardListItemForUser) GetLocked() bool`

GetLocked returns the Locked field if non-nil, zero value otherwise.

### GetLockedOk

`func (o *O11yO11yDashboardListItemForUser) GetLockedOk() (*bool, bool)`

GetLockedOk returns a tuple with the Locked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocked

`func (o *O11yO11yDashboardListItemForUser) SetLocked(v bool)`

SetLocked sets Locked field to given value.

### HasLocked

`func (o *O11yO11yDashboardListItemForUser) HasLocked() bool`

HasLocked returns a boolean if a field has been set.

### GetName

`func (o *O11yO11yDashboardListItemForUser) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *O11yO11yDashboardListItemForUser) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *O11yO11yDashboardListItemForUser) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *O11yO11yDashboardListItemForUser) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOrgId

`func (o *O11yO11yDashboardListItemForUser) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *O11yO11yDashboardListItemForUser) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *O11yO11yDashboardListItemForUser) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *O11yO11yDashboardListItemForUser) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### GetPinned

`func (o *O11yO11yDashboardListItemForUser) GetPinned() bool`

GetPinned returns the Pinned field if non-nil, zero value otherwise.

### GetPinnedOk

`func (o *O11yO11yDashboardListItemForUser) GetPinnedOk() (*bool, bool)`

GetPinnedOk returns a tuple with the Pinned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPinned

`func (o *O11yO11yDashboardListItemForUser) SetPinned(v bool)`

SetPinned sets Pinned field to given value.

### HasPinned

`func (o *O11yO11yDashboardListItemForUser) HasPinned() bool`

HasPinned returns a boolean if a field has been set.

### GetSchemaVersion

`func (o *O11yO11yDashboardListItemForUser) GetSchemaVersion() string`

GetSchemaVersion returns the SchemaVersion field if non-nil, zero value otherwise.

### GetSchemaVersionOk

`func (o *O11yO11yDashboardListItemForUser) GetSchemaVersionOk() (*string, bool)`

GetSchemaVersionOk returns a tuple with the SchemaVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaVersion

`func (o *O11yO11yDashboardListItemForUser) SetSchemaVersion(v string)`

SetSchemaVersion sets SchemaVersion field to given value.

### HasSchemaVersion

`func (o *O11yO11yDashboardListItemForUser) HasSchemaVersion() bool`

HasSchemaVersion returns a boolean if a field has been set.

### GetSource

`func (o *O11yO11yDashboardListItemForUser) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *O11yO11yDashboardListItemForUser) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *O11yO11yDashboardListItemForUser) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *O11yO11yDashboardListItemForUser) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetSpec

`func (o *O11yO11yDashboardListItemForUser) GetSpec() O11yO11yDashboardListSpec`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *O11yO11yDashboardListItemForUser) GetSpecOk() (*O11yO11yDashboardListSpec, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *O11yO11yDashboardListItemForUser) SetSpec(v O11yO11yDashboardListSpec)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *O11yO11yDashboardListItemForUser) HasSpec() bool`

HasSpec returns a boolean if a field has been set.

### GetTags

`func (o *O11yO11yDashboardListItemForUser) GetTags() []O11yO11yDashboardTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *O11yO11yDashboardListItemForUser) GetTagsOk() (*[]O11yO11yDashboardTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *O11yO11yDashboardListItemForUser) SetTags(v []O11yO11yDashboardTag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *O11yO11yDashboardListItemForUser) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *O11yO11yDashboardListItemForUser) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *O11yO11yDashboardListItemForUser) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *O11yO11yDashboardListItemForUser) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *O11yO11yDashboardListItemForUser) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUpdatedBy

`func (o *O11yO11yDashboardListItemForUser) GetUpdatedBy() string`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *O11yO11yDashboardListItemForUser) GetUpdatedByOk() (*string, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *O11yO11yDashboardListItemForUser) SetUpdatedBy(v string)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *O11yO11yDashboardListItemForUser) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


