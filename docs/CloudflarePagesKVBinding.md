# CloudflarePagesKVBinding

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NamespaceId** | Pointer to **string** | NamespaceID is the KV namespace this binding points at, by Cloudflare&#39;s id rather than its title. The BINDING NAME — what the Worker code reads it as — is the map key this value sits under, not a field here. | [optional] 

## Methods

### NewCloudflarePagesKVBinding

`func NewCloudflarePagesKVBinding() *CloudflarePagesKVBinding`

NewCloudflarePagesKVBinding instantiates a new CloudflarePagesKVBinding object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudflarePagesKVBindingWithDefaults

`func NewCloudflarePagesKVBindingWithDefaults() *CloudflarePagesKVBinding`

NewCloudflarePagesKVBindingWithDefaults instantiates a new CloudflarePagesKVBinding object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNamespaceId

`func (o *CloudflarePagesKVBinding) GetNamespaceId() string`

GetNamespaceId returns the NamespaceId field if non-nil, zero value otherwise.

### GetNamespaceIdOk

`func (o *CloudflarePagesKVBinding) GetNamespaceIdOk() (*string, bool)`

GetNamespaceIdOk returns a tuple with the NamespaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespaceId

`func (o *CloudflarePagesKVBinding) SetNamespaceId(v string)`

SetNamespaceId sets NamespaceId field to given value.

### HasNamespaceId

`func (o *CloudflarePagesKVBinding) HasNamespaceId() bool`

HasNamespaceId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


