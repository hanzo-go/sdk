# KmsProject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Slug** | Pointer to **string** |  | [optional] 
**OrgId** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Environments** | Pointer to [**[]KmsEnvironment**](KmsEnvironment.md) |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewKmsProject

`func NewKmsProject() *KmsProject`

NewKmsProject instantiates a new KmsProject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKmsProjectWithDefaults

`func NewKmsProjectWithDefaults() *KmsProject`

NewKmsProjectWithDefaults instantiates a new KmsProject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *KmsProject) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *KmsProject) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *KmsProject) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *KmsProject) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *KmsProject) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KmsProject) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *KmsProject) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *KmsProject) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSlug

`func (o *KmsProject) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *KmsProject) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *KmsProject) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *KmsProject) HasSlug() bool`

HasSlug returns a boolean if a field has been set.

### GetOrgId

`func (o *KmsProject) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *KmsProject) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *KmsProject) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *KmsProject) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### GetType

`func (o *KmsProject) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *KmsProject) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *KmsProject) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *KmsProject) HasType() bool`

HasType returns a boolean if a field has been set.

### GetEnvironments

`func (o *KmsProject) GetEnvironments() []KmsEnvironment`

GetEnvironments returns the Environments field if non-nil, zero value otherwise.

### GetEnvironmentsOk

`func (o *KmsProject) GetEnvironmentsOk() (*[]KmsEnvironment, bool)`

GetEnvironmentsOk returns a tuple with the Environments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironments

`func (o *KmsProject) SetEnvironments(v []KmsEnvironment)`

SetEnvironments sets Environments field to given value.

### HasEnvironments

`func (o *KmsProject) HasEnvironments() bool`

HasEnvironments returns a boolean if a field has been set.

### GetCreatedAt

`func (o *KmsProject) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *KmsProject) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *KmsProject) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *KmsProject) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *KmsProject) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *KmsProject) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *KmsProject) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *KmsProject) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


