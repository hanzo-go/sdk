# DestinationField

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Example** | Pointer to **string** | a sample value of the right shape (\&quot;G-XXXXXXX\&quot;), when one helps | [optional] 
**Key** | Pointer to **string** | the camelCase key on both the connect body and the stored config | [optional] 
**Label** | Pointer to **string** | human label for the console card&#39;s input | [optional] 
**Required** | Pointer to **bool** | when true, a connect that leaves it empty is refused 400 | [optional] 

## Methods

### NewDestinationField

`func NewDestinationField() *DestinationField`

NewDestinationField instantiates a new DestinationField object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDestinationFieldWithDefaults

`func NewDestinationFieldWithDefaults() *DestinationField`

NewDestinationFieldWithDefaults instantiates a new DestinationField object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExample

`func (o *DestinationField) GetExample() string`

GetExample returns the Example field if non-nil, zero value otherwise.

### GetExampleOk

`func (o *DestinationField) GetExampleOk() (*string, bool)`

GetExampleOk returns a tuple with the Example field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExample

`func (o *DestinationField) SetExample(v string)`

SetExample sets Example field to given value.

### HasExample

`func (o *DestinationField) HasExample() bool`

HasExample returns a boolean if a field has been set.

### GetKey

`func (o *DestinationField) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *DestinationField) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *DestinationField) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *DestinationField) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetLabel

`func (o *DestinationField) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *DestinationField) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *DestinationField) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *DestinationField) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetRequired

`func (o *DestinationField) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *DestinationField) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *DestinationField) SetRequired(v bool)`

SetRequired sets Required field to given value.

### HasRequired

`func (o *DestinationField) HasRequired() bool`

HasRequired returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


