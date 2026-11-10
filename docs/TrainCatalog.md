# TrainCatalog

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]TrainFamily**](TrainFamily.md) | Data is every base the caller can train, each with its methods, where it runs and what a device-second costs there. | [optional] 

## Methods

### NewTrainCatalog

`func NewTrainCatalog() *TrainCatalog`

NewTrainCatalog instantiates a new TrainCatalog object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrainCatalogWithDefaults

`func NewTrainCatalogWithDefaults() *TrainCatalog`

NewTrainCatalogWithDefaults instantiates a new TrainCatalog object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *TrainCatalog) GetData() []TrainFamily`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *TrainCatalog) GetDataOk() (*[]TrainFamily, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *TrainCatalog) SetData(v []TrainFamily)`

SetData sets Data field to given value.

### HasData

`func (o *TrainCatalog) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


