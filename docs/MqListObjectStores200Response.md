# MqListObjectStores200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Stores** | Pointer to [**[]MqObjectStoreInfo**](MqObjectStoreInfo.md) |  | [optional] 
**Total** | Pointer to **int32** |  | [optional] 

## Methods

### NewMqListObjectStores200Response

`func NewMqListObjectStores200Response() *MqListObjectStores200Response`

NewMqListObjectStores200Response instantiates a new MqListObjectStores200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMqListObjectStores200ResponseWithDefaults

`func NewMqListObjectStores200ResponseWithDefaults() *MqListObjectStores200Response`

NewMqListObjectStores200ResponseWithDefaults instantiates a new MqListObjectStores200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStores

`func (o *MqListObjectStores200Response) GetStores() []MqObjectStoreInfo`

GetStores returns the Stores field if non-nil, zero value otherwise.

### GetStoresOk

`func (o *MqListObjectStores200Response) GetStoresOk() (*[]MqObjectStoreInfo, bool)`

GetStoresOk returns a tuple with the Stores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStores

`func (o *MqListObjectStores200Response) SetStores(v []MqObjectStoreInfo)`

SetStores sets Stores field to given value.

### HasStores

`func (o *MqListObjectStores200Response) HasStores() bool`

HasStores returns a boolean if a field has been set.

### GetTotal

`func (o *MqListObjectStores200Response) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *MqListObjectStores200Response) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *MqListObjectStores200Response) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *MqListObjectStores200Response) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


