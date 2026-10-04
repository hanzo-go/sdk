# DeployArgoGroupKind

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Group** | Pointer to **string** | Group is the API group a project admits, \&quot;*\&quot; for any. Empty names the core group. | [optional] 
**Kind** | Pointer to **string** | Kind is the kind it admits, \&quot;*\&quot; for any. | [optional] 

## Methods

### NewDeployArgoGroupKind

`func NewDeployArgoGroupKind() *DeployArgoGroupKind`

NewDeployArgoGroupKind instantiates a new DeployArgoGroupKind object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeployArgoGroupKindWithDefaults

`func NewDeployArgoGroupKindWithDefaults() *DeployArgoGroupKind`

NewDeployArgoGroupKindWithDefaults instantiates a new DeployArgoGroupKind object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroup

`func (o *DeployArgoGroupKind) GetGroup() string`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *DeployArgoGroupKind) GetGroupOk() (*string, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *DeployArgoGroupKind) SetGroup(v string)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *DeployArgoGroupKind) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetKind

`func (o *DeployArgoGroupKind) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *DeployArgoGroupKind) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *DeployArgoGroupKind) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *DeployArgoGroupKind) HasKind() bool`

HasKind returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


