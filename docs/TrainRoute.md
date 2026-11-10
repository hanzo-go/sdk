# TrainRoute

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Keys** | Pointer to **[]string** | Keys are the question keys it answers. | [optional] 
**Options** | Pointer to [**map[string]TrainOptions**](TrainOptions.md) | Options narrows a key to one question: a question under that key routes here only when its type and option labels are exactly these. A key not named here routes whatever its question. | [optional] 

## Methods

### NewTrainRoute

`func NewTrainRoute() *TrainRoute`

NewTrainRoute instantiates a new TrainRoute object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrainRouteWithDefaults

`func NewTrainRouteWithDefaults() *TrainRoute`

NewTrainRouteWithDefaults instantiates a new TrainRoute object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKeys

`func (o *TrainRoute) GetKeys() []string`

GetKeys returns the Keys field if non-nil, zero value otherwise.

### GetKeysOk

`func (o *TrainRoute) GetKeysOk() (*[]string, bool)`

GetKeysOk returns a tuple with the Keys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeys

`func (o *TrainRoute) SetKeys(v []string)`

SetKeys sets Keys field to given value.

### HasKeys

`func (o *TrainRoute) HasKeys() bool`

HasKeys returns a boolean if a field has been set.

### GetOptions

`func (o *TrainRoute) GetOptions() map[string]TrainOptions`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *TrainRoute) GetOptionsOk() (*map[string]TrainOptions, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *TrainRoute) SetOptions(v map[string]TrainOptions)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *TrainRoute) HasOptions() bool`

HasOptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


