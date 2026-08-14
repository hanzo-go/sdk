# O11yO11yMetricOrder

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Direction** | Pointer to **string** | Direction is asc or desc. | [optional] 
**Key** | Pointer to [**O11yO11yMetricField**](O11yO11yMetricField.md) | Key is the field to order by. | [optional] 

## Methods

### NewO11yO11yMetricOrder

`func NewO11yO11yMetricOrder() *O11yO11yMetricOrder`

NewO11yO11yMetricOrder instantiates a new O11yO11yMetricOrder object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yO11yMetricOrderWithDefaults

`func NewO11yO11yMetricOrderWithDefaults() *O11yO11yMetricOrder`

NewO11yO11yMetricOrderWithDefaults instantiates a new O11yO11yMetricOrder object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDirection

`func (o *O11yO11yMetricOrder) GetDirection() string`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *O11yO11yMetricOrder) GetDirectionOk() (*string, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *O11yO11yMetricOrder) SetDirection(v string)`

SetDirection sets Direction field to given value.

### HasDirection

`func (o *O11yO11yMetricOrder) HasDirection() bool`

HasDirection returns a boolean if a field has been set.

### GetKey

`func (o *O11yO11yMetricOrder) GetKey() O11yO11yMetricField`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *O11yO11yMetricOrder) GetKeyOk() (*O11yO11yMetricField, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *O11yO11yMetricOrder) SetKey(v O11yO11yMetricField)`

SetKey sets Key field to given value.

### HasKey

`func (o *O11yO11yMetricOrder) HasKey() bool`

HasKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


