# LicensingRevokeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Reason** | Pointer to **string** | Reason is the operator&#39;s note, echoed back by verify so a support agent can explain the refusal. | [optional] 
**Scope** | **string** | Scope is what the revocation matches on: \&quot;nonce\&quot; kills one token, \&quot;holder\&quot; every token issued to one bearer, \&quot;fingerprint\&quot; every token bound to one device, and \&quot;release\&quot; every token scoped to one binary release. | 
**Value** | **string** | Value is the concrete nonce, holder, fingerprint or release id to revoke. | 

## Methods

### NewLicensingRevokeRequest

`func NewLicensingRevokeRequest(scope string, value string, ) *LicensingRevokeRequest`

NewLicensingRevokeRequest instantiates a new LicensingRevokeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLicensingRevokeRequestWithDefaults

`func NewLicensingRevokeRequestWithDefaults() *LicensingRevokeRequest`

NewLicensingRevokeRequestWithDefaults instantiates a new LicensingRevokeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReason

`func (o *LicensingRevokeRequest) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *LicensingRevokeRequest) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *LicensingRevokeRequest) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *LicensingRevokeRequest) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetScope

`func (o *LicensingRevokeRequest) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *LicensingRevokeRequest) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *LicensingRevokeRequest) SetScope(v string)`

SetScope sets Scope field to given value.


### GetValue

`func (o *LicensingRevokeRequest) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *LicensingRevokeRequest) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *LicensingRevokeRequest) SetValue(v string)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


