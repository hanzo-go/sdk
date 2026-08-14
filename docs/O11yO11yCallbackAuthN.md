# O11yO11yCallbackAuthN

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Provider** | Pointer to **string** | Provider is the route&#39;s provider — google_auth, saml or oidc. | [optional] 
**Url** | Pointer to **string** | URL is where the browser goes to begin the flow. | [optional] 

## Methods

### NewO11yO11yCallbackAuthN

`func NewO11yO11yCallbackAuthN() *O11yO11yCallbackAuthN`

NewO11yO11yCallbackAuthN instantiates a new O11yO11yCallbackAuthN object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yO11yCallbackAuthNWithDefaults

`func NewO11yO11yCallbackAuthNWithDefaults() *O11yO11yCallbackAuthN`

NewO11yO11yCallbackAuthNWithDefaults instantiates a new O11yO11yCallbackAuthN object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProvider

`func (o *O11yO11yCallbackAuthN) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *O11yO11yCallbackAuthN) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *O11yO11yCallbackAuthN) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *O11yO11yCallbackAuthN) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetUrl

`func (o *O11yO11yCallbackAuthN) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *O11yO11yCallbackAuthN) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *O11yO11yCallbackAuthN) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *O11yO11yCallbackAuthN) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


