# LabelRiskHoldIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hold** | Pointer to **bool** | Hold is the state to put them in: true places the hold, false releases it. One op both ways, because a hold that can be placed and not released pins a compliance record past every retention boundary with nothing able to let it go — and an operator who cannot release a hold stops placing them. | [optional] 
**Ids** | Pointer to **[]string** | IDs are the content digests of the records, as returned by the write and by the read. They name records in THIS tenant&#39;s plane; an id belonging to anybody else names nothing here, because the statement runs against this tenant&#39;s own file and there is no other file it could reach. | [optional] 

## Methods

### NewLabelRiskHoldIn

`func NewLabelRiskHoldIn() *LabelRiskHoldIn`

NewLabelRiskHoldIn instantiates a new LabelRiskHoldIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLabelRiskHoldInWithDefaults

`func NewLabelRiskHoldInWithDefaults() *LabelRiskHoldIn`

NewLabelRiskHoldInWithDefaults instantiates a new LabelRiskHoldIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHold

`func (o *LabelRiskHoldIn) GetHold() bool`

GetHold returns the Hold field if non-nil, zero value otherwise.

### GetHoldOk

`func (o *LabelRiskHoldIn) GetHoldOk() (*bool, bool)`

GetHoldOk returns a tuple with the Hold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHold

`func (o *LabelRiskHoldIn) SetHold(v bool)`

SetHold sets Hold field to given value.

### HasHold

`func (o *LabelRiskHoldIn) HasHold() bool`

HasHold returns a boolean if a field has been set.

### GetIds

`func (o *LabelRiskHoldIn) GetIds() []string`

GetIds returns the Ids field if non-nil, zero value otherwise.

### GetIdsOk

`func (o *LabelRiskHoldIn) GetIdsOk() (*[]string, bool)`

GetIdsOk returns a tuple with the Ids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIds

`func (o *LabelRiskHoldIn) SetIds(v []string)`

SetIds sets Ids field to given value.

### HasIds

`func (o *LabelRiskHoldIn) HasIds() bool`

HasIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


