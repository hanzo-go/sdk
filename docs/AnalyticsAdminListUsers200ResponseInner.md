# AnalyticsAdminListUsers200ResponseInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Username** | Pointer to **string** |  | [optional] 
**Role** | Pointer to **string** |  | [optional] 
**LogoUrl** | Pointer to **string** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 
**DeletedAt** | Pointer to **time.Time** |  | [optional] 
**Count** | Pointer to [**AnalyticsAdminListUsers200ResponseInnerAllOfCount**](AnalyticsAdminListUsers200ResponseInnerAllOfCount.md) |  | [optional] 

## Methods

### NewAnalyticsAdminListUsers200ResponseInner

`func NewAnalyticsAdminListUsers200ResponseInner() *AnalyticsAdminListUsers200ResponseInner`

NewAnalyticsAdminListUsers200ResponseInner instantiates a new AnalyticsAdminListUsers200ResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnalyticsAdminListUsers200ResponseInnerWithDefaults

`func NewAnalyticsAdminListUsers200ResponseInnerWithDefaults() *AnalyticsAdminListUsers200ResponseInner`

NewAnalyticsAdminListUsers200ResponseInnerWithDefaults instantiates a new AnalyticsAdminListUsers200ResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AnalyticsAdminListUsers200ResponseInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AnalyticsAdminListUsers200ResponseInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AnalyticsAdminListUsers200ResponseInner) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AnalyticsAdminListUsers200ResponseInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUsername

`func (o *AnalyticsAdminListUsers200ResponseInner) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *AnalyticsAdminListUsers200ResponseInner) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *AnalyticsAdminListUsers200ResponseInner) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *AnalyticsAdminListUsers200ResponseInner) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetRole

`func (o *AnalyticsAdminListUsers200ResponseInner) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *AnalyticsAdminListUsers200ResponseInner) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *AnalyticsAdminListUsers200ResponseInner) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *AnalyticsAdminListUsers200ResponseInner) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetLogoUrl

`func (o *AnalyticsAdminListUsers200ResponseInner) GetLogoUrl() string`

GetLogoUrl returns the LogoUrl field if non-nil, zero value otherwise.

### GetLogoUrlOk

`func (o *AnalyticsAdminListUsers200ResponseInner) GetLogoUrlOk() (*string, bool)`

GetLogoUrlOk returns a tuple with the LogoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoUrl

`func (o *AnalyticsAdminListUsers200ResponseInner) SetLogoUrl(v string)`

SetLogoUrl sets LogoUrl field to given value.

### HasLogoUrl

`func (o *AnalyticsAdminListUsers200ResponseInner) HasLogoUrl() bool`

HasLogoUrl returns a boolean if a field has been set.

### GetDisplayName

`func (o *AnalyticsAdminListUsers200ResponseInner) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *AnalyticsAdminListUsers200ResponseInner) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *AnalyticsAdminListUsers200ResponseInner) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *AnalyticsAdminListUsers200ResponseInner) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetCreatedAt

`func (o *AnalyticsAdminListUsers200ResponseInner) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AnalyticsAdminListUsers200ResponseInner) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AnalyticsAdminListUsers200ResponseInner) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *AnalyticsAdminListUsers200ResponseInner) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *AnalyticsAdminListUsers200ResponseInner) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *AnalyticsAdminListUsers200ResponseInner) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *AnalyticsAdminListUsers200ResponseInner) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *AnalyticsAdminListUsers200ResponseInner) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetDeletedAt

`func (o *AnalyticsAdminListUsers200ResponseInner) GetDeletedAt() time.Time`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *AnalyticsAdminListUsers200ResponseInner) GetDeletedAtOk() (*time.Time, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *AnalyticsAdminListUsers200ResponseInner) SetDeletedAt(v time.Time)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *AnalyticsAdminListUsers200ResponseInner) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetCount

`func (o *AnalyticsAdminListUsers200ResponseInner) GetCount() AnalyticsAdminListUsers200ResponseInnerAllOfCount`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *AnalyticsAdminListUsers200ResponseInner) GetCountOk() (*AnalyticsAdminListUsers200ResponseInnerAllOfCount, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *AnalyticsAdminListUsers200ResponseInner) SetCount(v AnalyticsAdminListUsers200ResponseInnerAllOfCount)`

SetCount sets Count field to given value.

### HasCount

`func (o *AnalyticsAdminListUsers200ResponseInner) HasCount() bool`

HasCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


