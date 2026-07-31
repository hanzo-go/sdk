# CloudDocType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Autoname** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **int32** |  | [optional] 
**Fields** | Pointer to [**[]CloudDocField**](CloudDocField.md) |  | [optional] 
**IsSingle** | Pointer to **bool** |  | [optional] 
**IsSubmittable** | Pointer to **bool** |  | [optional] 
**Module** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Permissions** | Pointer to [**[]CloudDocPerm**](CloudDocPerm.md) |  | [optional] 
**TitleField** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **int32** |  | [optional] 

## Methods

### NewCloudDocType

`func NewCloudDocType() *CloudDocType`

NewCloudDocType instantiates a new CloudDocType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudDocTypeWithDefaults

`func NewCloudDocTypeWithDefaults() *CloudDocType`

NewCloudDocTypeWithDefaults instantiates a new CloudDocType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutoname

`func (o *CloudDocType) GetAutoname() string`

GetAutoname returns the Autoname field if non-nil, zero value otherwise.

### GetAutonameOk

`func (o *CloudDocType) GetAutonameOk() (*string, bool)`

GetAutonameOk returns a tuple with the Autoname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoname

`func (o *CloudDocType) SetAutoname(v string)`

SetAutoname sets Autoname field to given value.

### HasAutoname

`func (o *CloudDocType) HasAutoname() bool`

HasAutoname returns a boolean if a field has been set.

### GetCreatedAt

`func (o *CloudDocType) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CloudDocType) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CloudDocType) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *CloudDocType) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetFields

`func (o *CloudDocType) GetFields() []CloudDocField`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *CloudDocType) GetFieldsOk() (*[]CloudDocField, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *CloudDocType) SetFields(v []CloudDocField)`

SetFields sets Fields field to given value.

### HasFields

`func (o *CloudDocType) HasFields() bool`

HasFields returns a boolean if a field has been set.

### GetIsSingle

`func (o *CloudDocType) GetIsSingle() bool`

GetIsSingle returns the IsSingle field if non-nil, zero value otherwise.

### GetIsSingleOk

`func (o *CloudDocType) GetIsSingleOk() (*bool, bool)`

GetIsSingleOk returns a tuple with the IsSingle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSingle

`func (o *CloudDocType) SetIsSingle(v bool)`

SetIsSingle sets IsSingle field to given value.

### HasIsSingle

`func (o *CloudDocType) HasIsSingle() bool`

HasIsSingle returns a boolean if a field has been set.

### GetIsSubmittable

`func (o *CloudDocType) GetIsSubmittable() bool`

GetIsSubmittable returns the IsSubmittable field if non-nil, zero value otherwise.

### GetIsSubmittableOk

`func (o *CloudDocType) GetIsSubmittableOk() (*bool, bool)`

GetIsSubmittableOk returns a tuple with the IsSubmittable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSubmittable

`func (o *CloudDocType) SetIsSubmittable(v bool)`

SetIsSubmittable sets IsSubmittable field to given value.

### HasIsSubmittable

`func (o *CloudDocType) HasIsSubmittable() bool`

HasIsSubmittable returns a boolean if a field has been set.

### GetModule

`func (o *CloudDocType) GetModule() string`

GetModule returns the Module field if non-nil, zero value otherwise.

### GetModuleOk

`func (o *CloudDocType) GetModuleOk() (*string, bool)`

GetModuleOk returns a tuple with the Module field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModule

`func (o *CloudDocType) SetModule(v string)`

SetModule sets Module field to given value.

### HasModule

`func (o *CloudDocType) HasModule() bool`

HasModule returns a boolean if a field has been set.

### GetName

`func (o *CloudDocType) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CloudDocType) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CloudDocType) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CloudDocType) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPermissions

`func (o *CloudDocType) GetPermissions() []CloudDocPerm`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *CloudDocType) GetPermissionsOk() (*[]CloudDocPerm, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *CloudDocType) SetPermissions(v []CloudDocPerm)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *CloudDocType) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### GetTitleField

`func (o *CloudDocType) GetTitleField() string`

GetTitleField returns the TitleField field if non-nil, zero value otherwise.

### GetTitleFieldOk

`func (o *CloudDocType) GetTitleFieldOk() (*string, bool)`

GetTitleFieldOk returns a tuple with the TitleField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitleField

`func (o *CloudDocType) SetTitleField(v string)`

SetTitleField sets TitleField field to given value.

### HasTitleField

`func (o *CloudDocType) HasTitleField() bool`

HasTitleField returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *CloudDocType) GetUpdatedAt() int32`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *CloudDocType) GetUpdatedAtOk() (*int32, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *CloudDocType) SetUpdatedAt(v int32)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *CloudDocType) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


