# CloudScaleIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int32** | Count is the node count to set. | [optional] 
**Id** | Pointer to **string** | ID is the DOKS cluster id, from the path. | [optional] 
**Pool** | Pointer to **string** | Pool is the node pool, from the path. Its DO id or its name — both are unique within a cluster, and an operator reads the name off the board. | [optional] 

## Methods

### NewCloudScaleIn

`func NewCloudScaleIn() *CloudScaleIn`

NewCloudScaleIn instantiates a new CloudScaleIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudScaleInWithDefaults

`func NewCloudScaleInWithDefaults() *CloudScaleIn`

NewCloudScaleInWithDefaults instantiates a new CloudScaleIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *CloudScaleIn) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *CloudScaleIn) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *CloudScaleIn) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *CloudScaleIn) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetId

`func (o *CloudScaleIn) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CloudScaleIn) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CloudScaleIn) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CloudScaleIn) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPool

`func (o *CloudScaleIn) GetPool() string`

GetPool returns the Pool field if non-nil, zero value otherwise.

### GetPoolOk

`func (o *CloudScaleIn) GetPoolOk() (*string, bool)`

GetPoolOk returns a tuple with the Pool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPool

`func (o *CloudScaleIn) SetPool(v string)`

SetPool sets Pool field to given value.

### HasPool

`func (o *CloudScaleIn) HasPool() bool`

HasPool returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


