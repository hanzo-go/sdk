# O11yO11yAuthDomain

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthNProviderInfo** | Pointer to [**O11yO11yAuthNProviderInfo**](O11yO11yAuthNProviderInfo.md) | AuthNProviderInfo is provider detail the console needs to finish setup. | [optional] 
**Config** | Pointer to [**O11yO11yAuthDomainConfig**](O11yO11yAuthDomainConfig.md) | Config is the domain&#39;s SSO configuration. | [optional] 
**CreatedAt** | Pointer to **time.Time** | CreatedAt is when it was claimed. | [optional] 
**Id** | Pointer to **string** | ID is the auth domain id. | [optional] 
**Name** | Pointer to **string** | Name is the email domain, e.g. example.com. | [optional] 
**OrgId** | Pointer to **string** | OrgID is the org that claimed it. | [optional] 
**UpdatedAt** | Pointer to **time.Time** | UpdatedAt is when its configuration last changed. | [optional] 

## Methods

### NewO11yO11yAuthDomain

`func NewO11yO11yAuthDomain() *O11yO11yAuthDomain`

NewO11yO11yAuthDomain instantiates a new O11yO11yAuthDomain object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yO11yAuthDomainWithDefaults

`func NewO11yO11yAuthDomainWithDefaults() *O11yO11yAuthDomain`

NewO11yO11yAuthDomainWithDefaults instantiates a new O11yO11yAuthDomain object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthNProviderInfo

`func (o *O11yO11yAuthDomain) GetAuthNProviderInfo() O11yO11yAuthNProviderInfo`

GetAuthNProviderInfo returns the AuthNProviderInfo field if non-nil, zero value otherwise.

### GetAuthNProviderInfoOk

`func (o *O11yO11yAuthDomain) GetAuthNProviderInfoOk() (*O11yO11yAuthNProviderInfo, bool)`

GetAuthNProviderInfoOk returns a tuple with the AuthNProviderInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthNProviderInfo

`func (o *O11yO11yAuthDomain) SetAuthNProviderInfo(v O11yO11yAuthNProviderInfo)`

SetAuthNProviderInfo sets AuthNProviderInfo field to given value.

### HasAuthNProviderInfo

`func (o *O11yO11yAuthDomain) HasAuthNProviderInfo() bool`

HasAuthNProviderInfo returns a boolean if a field has been set.

### GetConfig

`func (o *O11yO11yAuthDomain) GetConfig() O11yO11yAuthDomainConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *O11yO11yAuthDomain) GetConfigOk() (*O11yO11yAuthDomainConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *O11yO11yAuthDomain) SetConfig(v O11yO11yAuthDomainConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *O11yO11yAuthDomain) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetCreatedAt

`func (o *O11yO11yAuthDomain) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *O11yO11yAuthDomain) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *O11yO11yAuthDomain) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *O11yO11yAuthDomain) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetId

`func (o *O11yO11yAuthDomain) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *O11yO11yAuthDomain) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *O11yO11yAuthDomain) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *O11yO11yAuthDomain) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *O11yO11yAuthDomain) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *O11yO11yAuthDomain) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *O11yO11yAuthDomain) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *O11yO11yAuthDomain) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOrgId

`func (o *O11yO11yAuthDomain) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *O11yO11yAuthDomain) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *O11yO11yAuthDomain) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *O11yO11yAuthDomain) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *O11yO11yAuthDomain) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *O11yO11yAuthDomain) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *O11yO11yAuthDomain) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *O11yO11yAuthDomain) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


