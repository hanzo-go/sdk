# TrainOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Labels** | Pointer to **[]string** | Labels are its option labels, in order: a choice&#39;s criteria, a score&#39;s levels as text, a noul&#39;s two sides. | [optional] 
**Type** | Pointer to **string** | Type is choice, score or noul. | [optional] 

## Methods

### NewTrainOptions

`func NewTrainOptions() *TrainOptions`

NewTrainOptions instantiates a new TrainOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrainOptionsWithDefaults

`func NewTrainOptionsWithDefaults() *TrainOptions`

NewTrainOptionsWithDefaults instantiates a new TrainOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLabels

`func (o *TrainOptions) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *TrainOptions) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *TrainOptions) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *TrainOptions) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetType

`func (o *TrainOptions) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TrainOptions) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TrainOptions) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *TrainOptions) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


