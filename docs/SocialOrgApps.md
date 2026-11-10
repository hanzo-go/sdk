# SocialOrgApps

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Apps** | Pointer to [**[]SocialOrgApp**](SocialOrgApp.md) | Apps is the org&#39;s sealed developer apps. | [optional] 
**Data** | Pointer to [**[]SocialOrgApp**](SocialOrgApp.md) | Data is Apps again, so a client that reads either envelope sees the list. | [optional] 

## Methods

### NewSocialOrgApps

`func NewSocialOrgApps() *SocialOrgApps`

NewSocialOrgApps instantiates a new SocialOrgApps object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSocialOrgAppsWithDefaults

`func NewSocialOrgAppsWithDefaults() *SocialOrgApps`

NewSocialOrgAppsWithDefaults instantiates a new SocialOrgApps object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApps

`func (o *SocialOrgApps) GetApps() []SocialOrgApp`

GetApps returns the Apps field if non-nil, zero value otherwise.

### GetAppsOk

`func (o *SocialOrgApps) GetAppsOk() (*[]SocialOrgApp, bool)`

GetAppsOk returns a tuple with the Apps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApps

`func (o *SocialOrgApps) SetApps(v []SocialOrgApp)`

SetApps sets Apps field to given value.

### HasApps

`func (o *SocialOrgApps) HasApps() bool`

HasApps returns a boolean if a field has been set.

### GetData

`func (o *SocialOrgApps) GetData() []SocialOrgApp`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *SocialOrgApps) GetDataOk() (*[]SocialOrgApp, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *SocialOrgApps) SetData(v []SocialOrgApp)`

SetData sets Data field to given value.

### HasData

`func (o *SocialOrgApps) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


