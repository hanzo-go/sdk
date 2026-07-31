# CloudScope

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Org** | Pointer to **string** | Org is the IAM org slug the rows were read under: the validated principal&#39;s, resolved server-side. | [optional] 

## Methods

### NewCloudScope

`func NewCloudScope() *CloudScope`

NewCloudScope instantiates a new CloudScope object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudScopeWithDefaults

`func NewCloudScopeWithDefaults() *CloudScope`

NewCloudScopeWithDefaults instantiates a new CloudScope object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrg

`func (o *CloudScope) GetOrg() string`

GetOrg returns the Org field if non-nil, zero value otherwise.

### GetOrgOk

`func (o *CloudScope) GetOrgOk() (*string, bool)`

GetOrgOk returns a tuple with the Org field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrg

`func (o *CloudScope) SetOrg(v string)`

SetOrg sets Org field to given value.

### HasOrg

`func (o *CloudScope) HasOrg() bool`

HasOrg returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


