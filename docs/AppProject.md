# AppProject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Slug** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Org** | Pointer to **string** | Org slug (from the JWT owner claim). | [optional] 
**Domain** | Pointer to **string** | Custom or default hanzo.app domain. | [optional] 
**Url** | Pointer to **string** | Live URL of the latest deployment. | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewAppProject

`func NewAppProject() *AppProject`

NewAppProject instantiates a new AppProject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppProjectWithDefaults

`func NewAppProjectWithDefaults() *AppProject`

NewAppProjectWithDefaults instantiates a new AppProject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AppProject) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AppProject) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AppProject) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AppProject) HasId() bool`

HasId returns a boolean if a field has been set.

### GetSlug

`func (o *AppProject) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *AppProject) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *AppProject) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *AppProject) HasSlug() bool`

HasSlug returns a boolean if a field has been set.

### GetName

`func (o *AppProject) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AppProject) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AppProject) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AppProject) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOrg

`func (o *AppProject) GetOrg() string`

GetOrg returns the Org field if non-nil, zero value otherwise.

### GetOrgOk

`func (o *AppProject) GetOrgOk() (*string, bool)`

GetOrgOk returns a tuple with the Org field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrg

`func (o *AppProject) SetOrg(v string)`

SetOrg sets Org field to given value.

### HasOrg

`func (o *AppProject) HasOrg() bool`

HasOrg returns a boolean if a field has been set.

### GetDomain

`func (o *AppProject) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *AppProject) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *AppProject) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *AppProject) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetUrl

`func (o *AppProject) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *AppProject) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *AppProject) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *AppProject) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetCreatedAt

`func (o *AppProject) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AppProject) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AppProject) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *AppProject) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *AppProject) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *AppProject) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *AppProject) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *AppProject) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


