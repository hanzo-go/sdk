# CloudChannelSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Account** | Pointer to **string** | provider account ref (ad-account/page/list id) | [optional] 
**Detail** | Pointer to **string** | honest last-outcome detail (never a secret) | [optional] 
**ExternalId** | Pointer to **string** |  | [optional] 
**Kind** | Pointer to **string** | paid | organic | email | [optional] 
**Platform** | Pointer to **string** | meta | google | x | instagram | (email provider) | [optional] 
**Status** | Pointer to **string** | pending | live | paused | failed | unavailable | [optional] 

## Methods

### NewCloudChannelSpec

`func NewCloudChannelSpec() *CloudChannelSpec`

NewCloudChannelSpec instantiates a new CloudChannelSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudChannelSpecWithDefaults

`func NewCloudChannelSpecWithDefaults() *CloudChannelSpec`

NewCloudChannelSpecWithDefaults instantiates a new CloudChannelSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccount

`func (o *CloudChannelSpec) GetAccount() string`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *CloudChannelSpec) GetAccountOk() (*string, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *CloudChannelSpec) SetAccount(v string)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *CloudChannelSpec) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetDetail

`func (o *CloudChannelSpec) GetDetail() string`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *CloudChannelSpec) GetDetailOk() (*string, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *CloudChannelSpec) SetDetail(v string)`

SetDetail sets Detail field to given value.

### HasDetail

`func (o *CloudChannelSpec) HasDetail() bool`

HasDetail returns a boolean if a field has been set.

### GetExternalId

`func (o *CloudChannelSpec) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *CloudChannelSpec) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *CloudChannelSpec) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *CloudChannelSpec) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetKind

`func (o *CloudChannelSpec) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *CloudChannelSpec) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *CloudChannelSpec) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *CloudChannelSpec) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetPlatform

`func (o *CloudChannelSpec) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *CloudChannelSpec) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *CloudChannelSpec) SetPlatform(v string)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *CloudChannelSpec) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### GetStatus

`func (o *CloudChannelSpec) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CloudChannelSpec) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CloudChannelSpec) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CloudChannelSpec) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


