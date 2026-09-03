# RiskLabelsOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int64** | Count is how many this page holds. It is not a total: a total over an unbounded append-only log is a full scan of a single-writer file. | [optional] 
**Labels** | Pointer to [**[]RiskLabelRecord**](RiskLabelRecord.md) | Labels is the page, newest event first. | [optional] 

## Methods

### NewRiskLabelsOut

`func NewRiskLabelsOut() *RiskLabelsOut`

NewRiskLabelsOut instantiates a new RiskLabelsOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRiskLabelsOutWithDefaults

`func NewRiskLabelsOutWithDefaults() *RiskLabelsOut`

NewRiskLabelsOutWithDefaults instantiates a new RiskLabelsOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *RiskLabelsOut) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *RiskLabelsOut) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *RiskLabelsOut) SetCount(v int64)`

SetCount sets Count field to given value.

### HasCount

`func (o *RiskLabelsOut) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetLabels

`func (o *RiskLabelsOut) GetLabels() []RiskLabelRecord`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *RiskLabelsOut) GetLabelsOk() (*[]RiskLabelRecord, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *RiskLabelsOut) SetLabels(v []RiskLabelRecord)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *RiskLabelsOut) HasLabels() bool`

HasLabels returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


