# CloudAllowlistView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessGroups** | Pointer to [**map[string]map[string][]string**](map.md) |  | [optional] 
**Dm** | Pointer to **[]string** |  | [optional] 
**DmPolicy** | Pointer to **string** |  | [optional] 
**Group** | Pointer to **[]string** |  | [optional] 
**GroupPolicy** | Pointer to **string** |  | [optional] 
**Paired** | Pointer to **[]string** |  | [optional] 

## Methods

### NewCloudAllowlistView

`func NewCloudAllowlistView() *CloudAllowlistView`

NewCloudAllowlistView instantiates a new CloudAllowlistView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudAllowlistViewWithDefaults

`func NewCloudAllowlistViewWithDefaults() *CloudAllowlistView`

NewCloudAllowlistViewWithDefaults instantiates a new CloudAllowlistView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessGroups

`func (o *CloudAllowlistView) GetAccessGroups() map[string]map[string][]string`

GetAccessGroups returns the AccessGroups field if non-nil, zero value otherwise.

### GetAccessGroupsOk

`func (o *CloudAllowlistView) GetAccessGroupsOk() (*map[string]map[string][]string, bool)`

GetAccessGroupsOk returns a tuple with the AccessGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessGroups

`func (o *CloudAllowlistView) SetAccessGroups(v map[string]map[string][]string)`

SetAccessGroups sets AccessGroups field to given value.

### HasAccessGroups

`func (o *CloudAllowlistView) HasAccessGroups() bool`

HasAccessGroups returns a boolean if a field has been set.

### GetDm

`func (o *CloudAllowlistView) GetDm() []string`

GetDm returns the Dm field if non-nil, zero value otherwise.

### GetDmOk

`func (o *CloudAllowlistView) GetDmOk() (*[]string, bool)`

GetDmOk returns a tuple with the Dm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDm

`func (o *CloudAllowlistView) SetDm(v []string)`

SetDm sets Dm field to given value.

### HasDm

`func (o *CloudAllowlistView) HasDm() bool`

HasDm returns a boolean if a field has been set.

### GetDmPolicy

`func (o *CloudAllowlistView) GetDmPolicy() string`

GetDmPolicy returns the DmPolicy field if non-nil, zero value otherwise.

### GetDmPolicyOk

`func (o *CloudAllowlistView) GetDmPolicyOk() (*string, bool)`

GetDmPolicyOk returns a tuple with the DmPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDmPolicy

`func (o *CloudAllowlistView) SetDmPolicy(v string)`

SetDmPolicy sets DmPolicy field to given value.

### HasDmPolicy

`func (o *CloudAllowlistView) HasDmPolicy() bool`

HasDmPolicy returns a boolean if a field has been set.

### GetGroup

`func (o *CloudAllowlistView) GetGroup() []string`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *CloudAllowlistView) GetGroupOk() (*[]string, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *CloudAllowlistView) SetGroup(v []string)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *CloudAllowlistView) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetGroupPolicy

`func (o *CloudAllowlistView) GetGroupPolicy() string`

GetGroupPolicy returns the GroupPolicy field if non-nil, zero value otherwise.

### GetGroupPolicyOk

`func (o *CloudAllowlistView) GetGroupPolicyOk() (*string, bool)`

GetGroupPolicyOk returns a tuple with the GroupPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupPolicy

`func (o *CloudAllowlistView) SetGroupPolicy(v string)`

SetGroupPolicy sets GroupPolicy field to given value.

### HasGroupPolicy

`func (o *CloudAllowlistView) HasGroupPolicy() bool`

HasGroupPolicy returns a boolean if a field has been set.

### GetPaired

`func (o *CloudAllowlistView) GetPaired() []string`

GetPaired returns the Paired field if non-nil, zero value otherwise.

### GetPairedOk

`func (o *CloudAllowlistView) GetPairedOk() (*[]string, bool)`

GetPairedOk returns a tuple with the Paired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaired

`func (o *CloudAllowlistView) SetPaired(v []string)`

SetPaired sets Paired field to given value.

### HasPaired

`func (o *CloudAllowlistView) HasPaired() bool`

HasPaired returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


