# IngressBackend

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | Pointer to **string** | URL is the upstream server, http(s)://host[:port]. | [optional] 
**Weight** | Pointer to **int64** | Weight is this member&#39;s share of the round-robin; must be &gt;&#x3D; 0. | [optional] 

## Methods

### NewIngressBackend

`func NewIngressBackend() *IngressBackend`

NewIngressBackend instantiates a new IngressBackend object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIngressBackendWithDefaults

`func NewIngressBackendWithDefaults() *IngressBackend`

NewIngressBackendWithDefaults instantiates a new IngressBackend object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *IngressBackend) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *IngressBackend) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *IngressBackend) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *IngressBackend) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetWeight

`func (o *IngressBackend) GetWeight() int64`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *IngressBackend) GetWeightOk() (*int64, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *IngressBackend) SetWeight(v int64)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *IngressBackend) HasWeight() bool`

HasWeight returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


