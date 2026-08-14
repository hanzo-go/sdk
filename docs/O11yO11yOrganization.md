# O11yO11yOrganization

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Alias** | Pointer to **string** | Alias is an alternate name the org also answers to. | [optional] 
**CreatedAt** | Pointer to **time.Time** | CreatedAt is when the org was created. | [optional] 
**DisplayName** | Pointer to **string** | DisplayName is what the console shows for the org. | [optional] 
**Id** | Pointer to **string** | ID is the org id. On update it is ignored: the call always addresses the caller&#39;s own org. | [optional] 
**Key** | Pointer to **int32** | Key is the org&#39;s stable numeric key, derived from its id. | [optional] 
**Name** | Pointer to **string** | Name is the org&#39;s short name. | [optional] 
**UpdatedAt** | Pointer to **time.Time** | UpdatedAt is when it last changed. | [optional] 

## Methods

### NewO11yO11yOrganization

`func NewO11yO11yOrganization() *O11yO11yOrganization`

NewO11yO11yOrganization instantiates a new O11yO11yOrganization object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yO11yOrganizationWithDefaults

`func NewO11yO11yOrganizationWithDefaults() *O11yO11yOrganization`

NewO11yO11yOrganizationWithDefaults instantiates a new O11yO11yOrganization object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlias

`func (o *O11yO11yOrganization) GetAlias() string`

GetAlias returns the Alias field if non-nil, zero value otherwise.

### GetAliasOk

`func (o *O11yO11yOrganization) GetAliasOk() (*string, bool)`

GetAliasOk returns a tuple with the Alias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlias

`func (o *O11yO11yOrganization) SetAlias(v string)`

SetAlias sets Alias field to given value.

### HasAlias

`func (o *O11yO11yOrganization) HasAlias() bool`

HasAlias returns a boolean if a field has been set.

### GetCreatedAt

`func (o *O11yO11yOrganization) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *O11yO11yOrganization) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *O11yO11yOrganization) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *O11yO11yOrganization) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDisplayName

`func (o *O11yO11yOrganization) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *O11yO11yOrganization) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *O11yO11yOrganization) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *O11yO11yOrganization) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetId

`func (o *O11yO11yOrganization) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *O11yO11yOrganization) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *O11yO11yOrganization) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *O11yO11yOrganization) HasId() bool`

HasId returns a boolean if a field has been set.

### GetKey

`func (o *O11yO11yOrganization) GetKey() int32`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *O11yO11yOrganization) GetKeyOk() (*int32, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *O11yO11yOrganization) SetKey(v int32)`

SetKey sets Key field to given value.

### HasKey

`func (o *O11yO11yOrganization) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetName

`func (o *O11yO11yOrganization) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *O11yO11yOrganization) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *O11yO11yOrganization) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *O11yO11yOrganization) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *O11yO11yOrganization) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *O11yO11yOrganization) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *O11yO11yOrganization) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *O11yO11yOrganization) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


