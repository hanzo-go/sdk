# O11yO11yAuthNSupport

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Callback** | Pointer to [**[]O11yO11yCallbackAuthN**](O11yO11yCallbackAuthN.md) | Callback are the SSO routes; each is begun by visiting its URL. | [optional] 
**Password** | Pointer to [**[]O11yO11yPasswordAuthN**](O11yO11yPasswordAuthN.md) | Password are the password routes. | [optional] 

## Methods

### NewO11yO11yAuthNSupport

`func NewO11yO11yAuthNSupport() *O11yO11yAuthNSupport`

NewO11yO11yAuthNSupport instantiates a new O11yO11yAuthNSupport object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yO11yAuthNSupportWithDefaults

`func NewO11yO11yAuthNSupportWithDefaults() *O11yO11yAuthNSupport`

NewO11yO11yAuthNSupportWithDefaults instantiates a new O11yO11yAuthNSupport object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCallback

`func (o *O11yO11yAuthNSupport) GetCallback() []O11yO11yCallbackAuthN`

GetCallback returns the Callback field if non-nil, zero value otherwise.

### GetCallbackOk

`func (o *O11yO11yAuthNSupport) GetCallbackOk() (*[]O11yO11yCallbackAuthN, bool)`

GetCallbackOk returns a tuple with the Callback field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallback

`func (o *O11yO11yAuthNSupport) SetCallback(v []O11yO11yCallbackAuthN)`

SetCallback sets Callback field to given value.

### HasCallback

`func (o *O11yO11yAuthNSupport) HasCallback() bool`

HasCallback returns a boolean if a field has been set.

### GetPassword

`func (o *O11yO11yAuthNSupport) GetPassword() []O11yO11yPasswordAuthN`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *O11yO11yAuthNSupport) GetPasswordOk() (*[]O11yO11yPasswordAuthN, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *O11yO11yAuthNSupport) SetPassword(v []O11yO11yPasswordAuthN)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *O11yO11yAuthNSupport) HasPassword() bool`

HasPassword returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


