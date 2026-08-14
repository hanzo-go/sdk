# IngressMiddlewares

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Middlewares** | Pointer to [**[]Middleware**](Middleware.md) | Middlewares is the org&#39;s middlewares, ordered by id. | [optional] 

## Methods

### NewIngressMiddlewares

`func NewIngressMiddlewares() *IngressMiddlewares`

NewIngressMiddlewares instantiates a new IngressMiddlewares object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIngressMiddlewaresWithDefaults

`func NewIngressMiddlewaresWithDefaults() *IngressMiddlewares`

NewIngressMiddlewaresWithDefaults instantiates a new IngressMiddlewares object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMiddlewares

`func (o *IngressMiddlewares) GetMiddlewares() []Middleware`

GetMiddlewares returns the Middlewares field if non-nil, zero value otherwise.

### GetMiddlewaresOk

`func (o *IngressMiddlewares) GetMiddlewaresOk() (*[]Middleware, bool)`

GetMiddlewaresOk returns a tuple with the Middlewares field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMiddlewares

`func (o *IngressMiddlewares) SetMiddlewares(v []Middleware)`

SetMiddlewares sets Middlewares field to given value.

### HasMiddlewares

`func (o *IngressMiddlewares) HasMiddlewares() bool`

HasMiddlewares returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


