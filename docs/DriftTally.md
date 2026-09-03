# DriftTally

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ok** | Pointer to **int64** | OK is how many rows run what they declare. | [optional] 
**Red** | Pointer to **int64** | Red is how many have drifted badly. | [optional] 
**Yellow** | Pointer to **int64** | Yellow is how many have drifted within tolerance. | [optional] 

## Methods

### NewDriftTally

`func NewDriftTally() *DriftTally`

NewDriftTally instantiates a new DriftTally object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDriftTallyWithDefaults

`func NewDriftTallyWithDefaults() *DriftTally`

NewDriftTallyWithDefaults instantiates a new DriftTally object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOk

`func (o *DriftTally) GetOk() int64`

GetOk returns the Ok field if non-nil, zero value otherwise.

### GetOkOk

`func (o *DriftTally) GetOkOk() (*int64, bool)`

GetOkOk returns a tuple with the Ok field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOk

`func (o *DriftTally) SetOk(v int64)`

SetOk sets Ok field to given value.

### HasOk

`func (o *DriftTally) HasOk() bool`

HasOk returns a boolean if a field has been set.

### GetRed

`func (o *DriftTally) GetRed() int64`

GetRed returns the Red field if non-nil, zero value otherwise.

### GetRedOk

`func (o *DriftTally) GetRedOk() (*int64, bool)`

GetRedOk returns a tuple with the Red field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRed

`func (o *DriftTally) SetRed(v int64)`

SetRed sets Red field to given value.

### HasRed

`func (o *DriftTally) HasRed() bool`

HasRed returns a boolean if a field has been set.

### GetYellow

`func (o *DriftTally) GetYellow() int64`

GetYellow returns the Yellow field if non-nil, zero value otherwise.

### GetYellowOk

`func (o *DriftTally) GetYellowOk() (*int64, bool)`

GetYellowOk returns a tuple with the Yellow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYellow

`func (o *DriftTally) SetYellow(v int64)`

SetYellow sets Yellow field to given value.

### HasYellow

`func (o *DriftTally) HasYellow() bool`

HasYellow returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


