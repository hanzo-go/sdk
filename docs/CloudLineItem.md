# CloudLineItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AmountCents** | Pointer to **int32** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 

## Methods

### NewCloudLineItem

`func NewCloudLineItem() *CloudLineItem`

NewCloudLineItem instantiates a new CloudLineItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudLineItemWithDefaults

`func NewCloudLineItemWithDefaults() *CloudLineItem`

NewCloudLineItemWithDefaults instantiates a new CloudLineItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmountCents

`func (o *CloudLineItem) GetAmountCents() int32`

GetAmountCents returns the AmountCents field if non-nil, zero value otherwise.

### GetAmountCentsOk

`func (o *CloudLineItem) GetAmountCentsOk() (*int32, bool)`

GetAmountCentsOk returns a tuple with the AmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountCents

`func (o *CloudLineItem) SetAmountCents(v int32)`

SetAmountCents sets AmountCents field to given value.

### HasAmountCents

`func (o *CloudLineItem) HasAmountCents() bool`

HasAmountCents returns a boolean if a field has been set.

### GetDescription

`func (o *CloudLineItem) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CloudLineItem) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CloudLineItem) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CloudLineItem) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


