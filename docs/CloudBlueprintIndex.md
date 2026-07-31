# CloudBlueprintIndex

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]CloudBlueprintRow**](CloudBlueprintRow.md) | Data is one row per embedded blueprint, sorted by template id. | [optional] 

## Methods

### NewCloudBlueprintIndex

`func NewCloudBlueprintIndex() *CloudBlueprintIndex`

NewCloudBlueprintIndex instantiates a new CloudBlueprintIndex object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudBlueprintIndexWithDefaults

`func NewCloudBlueprintIndexWithDefaults() *CloudBlueprintIndex`

NewCloudBlueprintIndexWithDefaults instantiates a new CloudBlueprintIndex object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *CloudBlueprintIndex) GetData() []CloudBlueprintRow`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CloudBlueprintIndex) GetDataOk() (*[]CloudBlueprintRow, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CloudBlueprintIndex) SetData(v []CloudBlueprintRow)`

SetData sets Data field to given value.

### HasData

`func (o *CloudBlueprintIndex) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


