# DocType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Autoname** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **int64** |  | [optional] 
**Fields** | Pointer to [**[]DocField**](DocField.md) |  | [optional] 
**IsSingle** | Pointer to **bool** |  | [optional] 
**IsSubmittable** | Pointer to **bool** |  | [optional] 
**Module** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Permissions** | Pointer to [**[]DocPerm**](DocPerm.md) |  | [optional] 
**TitleField** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **int64** |  | [optional] 

## Methods

### NewDocType

`func NewDocType() *DocType`

NewDocType instantiates a new DocType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDocTypeWithDefaults

`func NewDocTypeWithDefaults() *DocType`

NewDocTypeWithDefaults instantiates a new DocType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutoname

`func (o *DocType) GetAutoname() string`

GetAutoname returns the Autoname field if non-nil, zero value otherwise.

### GetAutonameOk

`func (o *DocType) GetAutonameOk() (*string, bool)`

GetAutonameOk returns a tuple with the Autoname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoname

`func (o *DocType) SetAutoname(v string)`

SetAutoname sets Autoname field to given value.

### HasAutoname

`func (o *DocType) HasAutoname() bool`

HasAutoname returns a boolean if a field has been set.

### GetCreatedAt

`func (o *DocType) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DocType) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DocType) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *DocType) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetFields

`func (o *DocType) GetFields() []DocField`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *DocType) GetFieldsOk() (*[]DocField, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *DocType) SetFields(v []DocField)`

SetFields sets Fields field to given value.

### HasFields

`func (o *DocType) HasFields() bool`

HasFields returns a boolean if a field has been set.

### GetIsSingle

`func (o *DocType) GetIsSingle() bool`

GetIsSingle returns the IsSingle field if non-nil, zero value otherwise.

### GetIsSingleOk

`func (o *DocType) GetIsSingleOk() (*bool, bool)`

GetIsSingleOk returns a tuple with the IsSingle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSingle

`func (o *DocType) SetIsSingle(v bool)`

SetIsSingle sets IsSingle field to given value.

### HasIsSingle

`func (o *DocType) HasIsSingle() bool`

HasIsSingle returns a boolean if a field has been set.

### GetIsSubmittable

`func (o *DocType) GetIsSubmittable() bool`

GetIsSubmittable returns the IsSubmittable field if non-nil, zero value otherwise.

### GetIsSubmittableOk

`func (o *DocType) GetIsSubmittableOk() (*bool, bool)`

GetIsSubmittableOk returns a tuple with the IsSubmittable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSubmittable

`func (o *DocType) SetIsSubmittable(v bool)`

SetIsSubmittable sets IsSubmittable field to given value.

### HasIsSubmittable

`func (o *DocType) HasIsSubmittable() bool`

HasIsSubmittable returns a boolean if a field has been set.

### GetModule

`func (o *DocType) GetModule() string`

GetModule returns the Module field if non-nil, zero value otherwise.

### GetModuleOk

`func (o *DocType) GetModuleOk() (*string, bool)`

GetModuleOk returns a tuple with the Module field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModule

`func (o *DocType) SetModule(v string)`

SetModule sets Module field to given value.

### HasModule

`func (o *DocType) HasModule() bool`

HasModule returns a boolean if a field has been set.

### GetName

`func (o *DocType) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DocType) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DocType) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DocType) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPermissions

`func (o *DocType) GetPermissions() []DocPerm`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *DocType) GetPermissionsOk() (*[]DocPerm, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *DocType) SetPermissions(v []DocPerm)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *DocType) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### GetTitleField

`func (o *DocType) GetTitleField() string`

GetTitleField returns the TitleField field if non-nil, zero value otherwise.

### GetTitleFieldOk

`func (o *DocType) GetTitleFieldOk() (*string, bool)`

GetTitleFieldOk returns a tuple with the TitleField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitleField

`func (o *DocType) SetTitleField(v string)`

SetTitleField sets TitleField field to given value.

### HasTitleField

`func (o *DocType) HasTitleField() bool`

HasTitleField returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *DocType) GetUpdatedAt() int64`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *DocType) GetUpdatedAtOk() (*int64, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *DocType) SetUpdatedAt(v int64)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *DocType) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


