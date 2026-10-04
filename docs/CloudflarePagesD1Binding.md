# CloudflarePagesD1Binding

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | ID is the D1 database this binding points at, by Cloudflare&#39;s uuid. The binding name the Worker code reads it as is the map key, not a field here. | [optional] 

## Methods

### NewCloudflarePagesD1Binding

`func NewCloudflarePagesD1Binding() *CloudflarePagesD1Binding`

NewCloudflarePagesD1Binding instantiates a new CloudflarePagesD1Binding object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudflarePagesD1BindingWithDefaults

`func NewCloudflarePagesD1BindingWithDefaults() *CloudflarePagesD1Binding`

NewCloudflarePagesD1BindingWithDefaults instantiates a new CloudflarePagesD1Binding object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CloudflarePagesD1Binding) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CloudflarePagesD1Binding) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CloudflarePagesD1Binding) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CloudflarePagesD1Binding) HasId() bool`

HasId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


