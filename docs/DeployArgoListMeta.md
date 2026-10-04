# DeployArgoListMeta

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourceVersion** | Pointer to **string** | ResourceVersion is the k8s list version a watch would resume from. Always empty: every list on this plane is COMPUTED per request rather than read from one etcd revision, so there is no point to resume from. The live view is the SSE stream, not a resumed watch. | [optional] 

## Methods

### NewDeployArgoListMeta

`func NewDeployArgoListMeta() *DeployArgoListMeta`

NewDeployArgoListMeta instantiates a new DeployArgoListMeta object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeployArgoListMetaWithDefaults

`func NewDeployArgoListMetaWithDefaults() *DeployArgoListMeta`

NewDeployArgoListMetaWithDefaults instantiates a new DeployArgoListMeta object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResourceVersion

`func (o *DeployArgoListMeta) GetResourceVersion() string`

GetResourceVersion returns the ResourceVersion field if non-nil, zero value otherwise.

### GetResourceVersionOk

`func (o *DeployArgoListMeta) GetResourceVersionOk() (*string, bool)`

GetResourceVersionOk returns a tuple with the ResourceVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceVersion

`func (o *DeployArgoListMeta) SetResourceVersion(v string)`

SetResourceVersion sets ResourceVersion field to given value.

### HasResourceVersion

`func (o *DeployArgoListMeta) HasResourceVersion() bool`

HasResourceVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


