# FrameworkDocType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Module** | Pointer to **string** |  | [optional] 
**IsSingle** | Pointer to **bool** | Exactly one document exists | [optional] 
**IsSubmittable** | Pointer to **bool** | Has the submit/cancel lifecycle | [optional] 
**Autoname** | Pointer to **string** | \&quot;\&quot; or hash → random id; \&quot;field:x\&quot;; \&quot;prompt\&quot;; or a series pattern e.g. \&quot;INV-.YYYY.-.#####\&quot; | [optional] 
**TitleField** | Pointer to **string** |  | [optional] 
**Fields** | [**[]FrameworkDocField**](FrameworkDocField.md) |  | 
**Permissions** | Pointer to [**[]FrameworkDocPerm**](FrameworkDocPerm.md) |  | [optional] 
**CreatedAt** | Pointer to **int64** |  | [optional] 
**UpdatedAt** | Pointer to **int64** |  | [optional] 

## Methods

### NewFrameworkDocType

`func NewFrameworkDocType(name string, fields []FrameworkDocField, ) *FrameworkDocType`

NewFrameworkDocType instantiates a new FrameworkDocType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFrameworkDocTypeWithDefaults

`func NewFrameworkDocTypeWithDefaults() *FrameworkDocType`

NewFrameworkDocTypeWithDefaults instantiates a new FrameworkDocType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *FrameworkDocType) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FrameworkDocType) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FrameworkDocType) SetName(v string)`

SetName sets Name field to given value.


### GetModule

`func (o *FrameworkDocType) GetModule() string`

GetModule returns the Module field if non-nil, zero value otherwise.

### GetModuleOk

`func (o *FrameworkDocType) GetModuleOk() (*string, bool)`

GetModuleOk returns a tuple with the Module field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModule

`func (o *FrameworkDocType) SetModule(v string)`

SetModule sets Module field to given value.

### HasModule

`func (o *FrameworkDocType) HasModule() bool`

HasModule returns a boolean if a field has been set.

### GetIsSingle

`func (o *FrameworkDocType) GetIsSingle() bool`

GetIsSingle returns the IsSingle field if non-nil, zero value otherwise.

### GetIsSingleOk

`func (o *FrameworkDocType) GetIsSingleOk() (*bool, bool)`

GetIsSingleOk returns a tuple with the IsSingle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSingle

`func (o *FrameworkDocType) SetIsSingle(v bool)`

SetIsSingle sets IsSingle field to given value.

### HasIsSingle

`func (o *FrameworkDocType) HasIsSingle() bool`

HasIsSingle returns a boolean if a field has been set.

### GetIsSubmittable

`func (o *FrameworkDocType) GetIsSubmittable() bool`

GetIsSubmittable returns the IsSubmittable field if non-nil, zero value otherwise.

### GetIsSubmittableOk

`func (o *FrameworkDocType) GetIsSubmittableOk() (*bool, bool)`

GetIsSubmittableOk returns a tuple with the IsSubmittable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSubmittable

`func (o *FrameworkDocType) SetIsSubmittable(v bool)`

SetIsSubmittable sets IsSubmittable field to given value.

### HasIsSubmittable

`func (o *FrameworkDocType) HasIsSubmittable() bool`

HasIsSubmittable returns a boolean if a field has been set.

### GetAutoname

`func (o *FrameworkDocType) GetAutoname() string`

GetAutoname returns the Autoname field if non-nil, zero value otherwise.

### GetAutonameOk

`func (o *FrameworkDocType) GetAutonameOk() (*string, bool)`

GetAutonameOk returns a tuple with the Autoname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoname

`func (o *FrameworkDocType) SetAutoname(v string)`

SetAutoname sets Autoname field to given value.

### HasAutoname

`func (o *FrameworkDocType) HasAutoname() bool`

HasAutoname returns a boolean if a field has been set.

### GetTitleField

`func (o *FrameworkDocType) GetTitleField() string`

GetTitleField returns the TitleField field if non-nil, zero value otherwise.

### GetTitleFieldOk

`func (o *FrameworkDocType) GetTitleFieldOk() (*string, bool)`

GetTitleFieldOk returns a tuple with the TitleField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitleField

`func (o *FrameworkDocType) SetTitleField(v string)`

SetTitleField sets TitleField field to given value.

### HasTitleField

`func (o *FrameworkDocType) HasTitleField() bool`

HasTitleField returns a boolean if a field has been set.

### GetFields

`func (o *FrameworkDocType) GetFields() []FrameworkDocField`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *FrameworkDocType) GetFieldsOk() (*[]FrameworkDocField, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *FrameworkDocType) SetFields(v []FrameworkDocField)`

SetFields sets Fields field to given value.


### GetPermissions

`func (o *FrameworkDocType) GetPermissions() []FrameworkDocPerm`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *FrameworkDocType) GetPermissionsOk() (*[]FrameworkDocPerm, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *FrameworkDocType) SetPermissions(v []FrameworkDocPerm)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *FrameworkDocType) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### GetCreatedAt

`func (o *FrameworkDocType) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *FrameworkDocType) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *FrameworkDocType) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *FrameworkDocType) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *FrameworkDocType) GetUpdatedAt() int64`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *FrameworkDocType) GetUpdatedAtOk() (*int64, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *FrameworkDocType) SetUpdatedAt(v int64)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *FrameworkDocType) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


