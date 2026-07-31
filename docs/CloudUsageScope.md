# CloudUsageScope

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Org** | Pointer to **string** | Org is the IAM org slug the rows were read under — the validated principal&#39;s, never a client header. | [optional] 
**User** | Pointer to **string** | User is the caller&#39;s own subject, whose linked-account rows the accounts block carries. Absent on a read that is org-scoped only. | [optional] 

## Methods

### NewCloudUsageScope

`func NewCloudUsageScope() *CloudUsageScope`

NewCloudUsageScope instantiates a new CloudUsageScope object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudUsageScopeWithDefaults

`func NewCloudUsageScopeWithDefaults() *CloudUsageScope`

NewCloudUsageScopeWithDefaults instantiates a new CloudUsageScope object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrg

`func (o *CloudUsageScope) GetOrg() string`

GetOrg returns the Org field if non-nil, zero value otherwise.

### GetOrgOk

`func (o *CloudUsageScope) GetOrgOk() (*string, bool)`

GetOrgOk returns a tuple with the Org field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrg

`func (o *CloudUsageScope) SetOrg(v string)`

SetOrg sets Org field to given value.

### HasOrg

`func (o *CloudUsageScope) HasOrg() bool`

HasOrg returns a boolean if a field has been set.

### GetUser

`func (o *CloudUsageScope) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *CloudUsageScope) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *CloudUsageScope) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *CloudUsageScope) HasUser() bool`

HasUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


