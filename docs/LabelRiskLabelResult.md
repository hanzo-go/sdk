# LabelRiskLabelResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | ID is the content digest of the assertion — the id a redelivery of the same fact resolves to. | [optional] 
**Refusal** | Pointer to **string** | Refusal states what was wrong, for the refused. | [optional] 
**Status** | Pointer to **string** | Status is recorded, duplicate or refused. | [optional] 

## Methods

### NewLabelRiskLabelResult

`func NewLabelRiskLabelResult() *LabelRiskLabelResult`

NewLabelRiskLabelResult instantiates a new LabelRiskLabelResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLabelRiskLabelResultWithDefaults

`func NewLabelRiskLabelResultWithDefaults() *LabelRiskLabelResult`

NewLabelRiskLabelResultWithDefaults instantiates a new LabelRiskLabelResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *LabelRiskLabelResult) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LabelRiskLabelResult) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LabelRiskLabelResult) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *LabelRiskLabelResult) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRefusal

`func (o *LabelRiskLabelResult) GetRefusal() string`

GetRefusal returns the Refusal field if non-nil, zero value otherwise.

### GetRefusalOk

`func (o *LabelRiskLabelResult) GetRefusalOk() (*string, bool)`

GetRefusalOk returns a tuple with the Refusal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefusal

`func (o *LabelRiskLabelResult) SetRefusal(v string)`

SetRefusal sets Refusal field to given value.

### HasRefusal

`func (o *LabelRiskLabelResult) HasRefusal() bool`

HasRefusal returns a boolean if a field has been set.

### GetStatus

`func (o *LabelRiskLabelResult) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *LabelRiskLabelResult) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *LabelRiskLabelResult) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *LabelRiskLabelResult) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


