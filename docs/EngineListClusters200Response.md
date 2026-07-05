# EngineListClusters200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Clusters** | Pointer to [**[]EngineCluster**](EngineCluster.md) |  | [optional] 
**Total** | Pointer to **int32** |  | [optional] 

## Methods

### NewEngineListClusters200Response

`func NewEngineListClusters200Response() *EngineListClusters200Response`

NewEngineListClusters200Response instantiates a new EngineListClusters200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEngineListClusters200ResponseWithDefaults

`func NewEngineListClusters200ResponseWithDefaults() *EngineListClusters200Response`

NewEngineListClusters200ResponseWithDefaults instantiates a new EngineListClusters200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusters

`func (o *EngineListClusters200Response) GetClusters() []EngineCluster`

GetClusters returns the Clusters field if non-nil, zero value otherwise.

### GetClustersOk

`func (o *EngineListClusters200Response) GetClustersOk() (*[]EngineCluster, bool)`

GetClustersOk returns a tuple with the Clusters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusters

`func (o *EngineListClusters200Response) SetClusters(v []EngineCluster)`

SetClusters sets Clusters field to given value.

### HasClusters

`func (o *EngineListClusters200Response) HasClusters() bool`

HasClusters returns a boolean if a field has been set.

### GetTotal

`func (o *EngineListClusters200Response) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *EngineListClusters200Response) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *EngineListClusters200Response) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *EngineListClusters200Response) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


