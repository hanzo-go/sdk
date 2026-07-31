# CloudOnboardResp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Additional** | Pointer to **bool** | Additional is true when the caller already had an organization and this one was created WITHOUT moving them into it — they reach it via the org switcher. | [optional] 
**DisplayName** | Pointer to **string** | DisplayName is the organization&#39;s human name. | [optional] 
**Org** | Pointer to **string** | Org is the created organization&#39;s slug, which is what X-Org-Id carries. | [optional] 

## Methods

### NewCloudOnboardResp

`func NewCloudOnboardResp() *CloudOnboardResp`

NewCloudOnboardResp instantiates a new CloudOnboardResp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudOnboardRespWithDefaults

`func NewCloudOnboardRespWithDefaults() *CloudOnboardResp`

NewCloudOnboardRespWithDefaults instantiates a new CloudOnboardResp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdditional

`func (o *CloudOnboardResp) GetAdditional() bool`

GetAdditional returns the Additional field if non-nil, zero value otherwise.

### GetAdditionalOk

`func (o *CloudOnboardResp) GetAdditionalOk() (*bool, bool)`

GetAdditionalOk returns a tuple with the Additional field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditional

`func (o *CloudOnboardResp) SetAdditional(v bool)`

SetAdditional sets Additional field to given value.

### HasAdditional

`func (o *CloudOnboardResp) HasAdditional() bool`

HasAdditional returns a boolean if a field has been set.

### GetDisplayName

`func (o *CloudOnboardResp) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *CloudOnboardResp) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *CloudOnboardResp) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *CloudOnboardResp) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetOrg

`func (o *CloudOnboardResp) GetOrg() string`

GetOrg returns the Org field if non-nil, zero value otherwise.

### GetOrgOk

`func (o *CloudOnboardResp) GetOrgOk() (*string, bool)`

GetOrgOk returns a tuple with the Org field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrg

`func (o *CloudOnboardResp) SetOrg(v string)`

SetOrg sets Org field to given value.

### HasOrg

`func (o *CloudOnboardResp) HasOrg() bool`

HasOrg returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


