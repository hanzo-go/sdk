# ArgoListMeta

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourceVersion** | Pointer to **string** | ResourceVersion is the k8s list version a watch would resume from. Always empty: every list on this plane is COMPUTED per request rather than read from one etcd revision, so there is no point to resume from. The live view is the SSE stream, not a resumed watch. | [optional] 

## Methods

### NewArgoListMeta

`func NewArgoListMeta() *ArgoListMeta`

NewArgoListMeta instantiates a new ArgoListMeta object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArgoListMetaWithDefaults

`func NewArgoListMetaWithDefaults() *ArgoListMeta`

NewArgoListMetaWithDefaults instantiates a new ArgoListMeta object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResourceVersion

`func (o *ArgoListMeta) GetResourceVersion() string`

GetResourceVersion returns the ResourceVersion field if non-nil, zero value otherwise.

### GetResourceVersionOk

`func (o *ArgoListMeta) GetResourceVersionOk() (*string, bool)`

GetResourceVersionOk returns a tuple with the ResourceVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceVersion

`func (o *ArgoListMeta) SetResourceVersion(v string)`

SetResourceVersion sets ResourceVersion field to given value.

### HasResourceVersion

`func (o *ArgoListMeta) HasResourceVersion() bool`

HasResourceVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


