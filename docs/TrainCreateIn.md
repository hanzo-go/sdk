# TrainCreateIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Corrections** | Pointer to **[]string** | Corrections are the corrections it is taught from: captured corrections of one question, which becomes its route. | [optional] 
**Name** | Pointer to **string** | Name is its name: lowercase words joined by hyphens, at most 64. | [optional] 

## Methods

### NewTrainCreateIn

`func NewTrainCreateIn() *TrainCreateIn`

NewTrainCreateIn instantiates a new TrainCreateIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrainCreateInWithDefaults

`func NewTrainCreateInWithDefaults() *TrainCreateIn`

NewTrainCreateInWithDefaults instantiates a new TrainCreateIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCorrections

`func (o *TrainCreateIn) GetCorrections() []string`

GetCorrections returns the Corrections field if non-nil, zero value otherwise.

### GetCorrectionsOk

`func (o *TrainCreateIn) GetCorrectionsOk() (*[]string, bool)`

GetCorrectionsOk returns a tuple with the Corrections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorrections

`func (o *TrainCreateIn) SetCorrections(v []string)`

SetCorrections sets Corrections field to given value.

### HasCorrections

`func (o *TrainCreateIn) HasCorrections() bool`

HasCorrections returns a boolean if a field has been set.

### GetName

`func (o *TrainCreateIn) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TrainCreateIn) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TrainCreateIn) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TrainCreateIn) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


