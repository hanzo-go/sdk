# KvListClusters200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Clusters** | Pointer to [**[]KvCluster**](KvCluster.md) |  | [optional] 
**Total** | Pointer to **int32** |  | [optional] 

## Methods

### NewKvListClusters200Response

`func NewKvListClusters200Response() *KvListClusters200Response`

NewKvListClusters200Response instantiates a new KvListClusters200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKvListClusters200ResponseWithDefaults

`func NewKvListClusters200ResponseWithDefaults() *KvListClusters200Response`

NewKvListClusters200ResponseWithDefaults instantiates a new KvListClusters200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusters

`func (o *KvListClusters200Response) GetClusters() []KvCluster`

GetClusters returns the Clusters field if non-nil, zero value otherwise.

### GetClustersOk

`func (o *KvListClusters200Response) GetClustersOk() (*[]KvCluster, bool)`

GetClustersOk returns a tuple with the Clusters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusters

`func (o *KvListClusters200Response) SetClusters(v []KvCluster)`

SetClusters sets Clusters field to given value.

### HasClusters

`func (o *KvListClusters200Response) HasClusters() bool`

HasClusters returns a boolean if a field has been set.

### GetTotal

`func (o *KvListClusters200Response) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *KvListClusters200Response) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *KvListClusters200Response) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *KvListClusters200Response) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


