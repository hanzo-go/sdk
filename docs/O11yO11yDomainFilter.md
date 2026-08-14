# O11yO11yDomainFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Expression** | Pointer to **string** | Expression is the predicate, e.g. &#x60;http.status_code &gt;&#x3D; 500&#x60;. | [optional] 

## Methods

### NewO11yO11yDomainFilter

`func NewO11yO11yDomainFilter() *O11yO11yDomainFilter`

NewO11yO11yDomainFilter instantiates a new O11yO11yDomainFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yO11yDomainFilterWithDefaults

`func NewO11yO11yDomainFilterWithDefaults() *O11yO11yDomainFilter`

NewO11yO11yDomainFilterWithDefaults instantiates a new O11yO11yDomainFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpression

`func (o *O11yO11yDomainFilter) GetExpression() string`

GetExpression returns the Expression field if non-nil, zero value otherwise.

### GetExpressionOk

`func (o *O11yO11yDomainFilter) GetExpressionOk() (*string, bool)`

GetExpressionOk returns a tuple with the Expression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpression

`func (o *O11yO11yDomainFilter) SetExpression(v string)`

SetExpression sets Expression field to given value.

### HasExpression

`func (o *O11yO11yDomainFilter) HasExpression() bool`

HasExpression returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


