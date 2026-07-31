# CloudPoolScale

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClusterId** | Pointer to **string** | ClusterID and PoolID address the pool, from the URL path. | [optional] 
**Count** | Pointer to **int32** | Count is the node count to scale TO — an absolute target, not a delta, and never negative. | [optional] 
**PoolId** | Pointer to **string** |  | [optional] 
**Provider** | Pointer to **string** | Provider is the cloud the cluster lives on. Required; body or ?provider&#x3D;. | [optional] 

## Methods

### NewCloudPoolScale

`func NewCloudPoolScale() *CloudPoolScale`

NewCloudPoolScale instantiates a new CloudPoolScale object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudPoolScaleWithDefaults

`func NewCloudPoolScaleWithDefaults() *CloudPoolScale`

NewCloudPoolScaleWithDefaults instantiates a new CloudPoolScale object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusterId

`func (o *CloudPoolScale) GetClusterId() string`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *CloudPoolScale) GetClusterIdOk() (*string, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *CloudPoolScale) SetClusterId(v string)`

SetClusterId sets ClusterId field to given value.

### HasClusterId

`func (o *CloudPoolScale) HasClusterId() bool`

HasClusterId returns a boolean if a field has been set.

### GetCount

`func (o *CloudPoolScale) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *CloudPoolScale) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *CloudPoolScale) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *CloudPoolScale) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetPoolId

`func (o *CloudPoolScale) GetPoolId() string`

GetPoolId returns the PoolId field if non-nil, zero value otherwise.

### GetPoolIdOk

`func (o *CloudPoolScale) GetPoolIdOk() (*string, bool)`

GetPoolIdOk returns a tuple with the PoolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolId

`func (o *CloudPoolScale) SetPoolId(v string)`

SetPoolId sets PoolId field to given value.

### HasPoolId

`func (o *CloudPoolScale) HasPoolId() bool`

HasPoolId returns a boolean if a field has been set.

### GetProvider

`func (o *CloudPoolScale) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *CloudPoolScale) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *CloudPoolScale) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *CloudPoolScale) HasProvider() bool`

HasProvider returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


