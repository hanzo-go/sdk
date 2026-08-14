# O11yFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Expression** | Pointer to **string** | expression to filter by following the filter syntax | [optional] 

## Methods

### NewO11yFilter

`func NewO11yFilter() *O11yFilter`

NewO11yFilter instantiates a new O11yFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yFilterWithDefaults

`func NewO11yFilterWithDefaults() *O11yFilter`

NewO11yFilterWithDefaults instantiates a new O11yFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpression

`func (o *O11yFilter) GetExpression() string`

GetExpression returns the Expression field if non-nil, zero value otherwise.

### GetExpressionOk

`func (o *O11yFilter) GetExpressionOk() (*string, bool)`

GetExpressionOk returns a tuple with the Expression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpression

`func (o *O11yFilter) SetExpression(v string)`

SetExpression sets Expression field to given value.

### HasExpression

`func (o *O11yFilter) HasExpression() bool`

HasExpression returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


