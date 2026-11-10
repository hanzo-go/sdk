# BillingGatewayTerms

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Accepts** | Pointer to [**[]BillingGatewayAccept**](BillingGatewayAccept.md) |  | [optional] 
**MaxCents** | Pointer to **int64** |  | [optional] 
**MaxValiditySeconds** | Pointer to **int64** |  | [optional] 
**MinCents** | Pointer to **int64** |  | [optional] 
**Nonce** | Pointer to **string** |  | [optional] 
**PayTo** | Pointer to **string** |  | [optional] 
**PrimaryType** | Pointer to **string** |  | [optional] 
**Salt** | Pointer to **string** |  | [optional] 
**Scheme** | Pointer to **string** |  | [optional] 
**Types** | Pointer to **interface{}** |  | [optional] 
**X402Version** | Pointer to **int64** |  | [optional] 

## Methods

### NewBillingGatewayTerms

`func NewBillingGatewayTerms() *BillingGatewayTerms`

NewBillingGatewayTerms instantiates a new BillingGatewayTerms object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBillingGatewayTermsWithDefaults

`func NewBillingGatewayTermsWithDefaults() *BillingGatewayTerms`

NewBillingGatewayTermsWithDefaults instantiates a new BillingGatewayTerms object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccepts

`func (o *BillingGatewayTerms) GetAccepts() []BillingGatewayAccept`

GetAccepts returns the Accepts field if non-nil, zero value otherwise.

### GetAcceptsOk

`func (o *BillingGatewayTerms) GetAcceptsOk() (*[]BillingGatewayAccept, bool)`

GetAcceptsOk returns a tuple with the Accepts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccepts

`func (o *BillingGatewayTerms) SetAccepts(v []BillingGatewayAccept)`

SetAccepts sets Accepts field to given value.

### HasAccepts

`func (o *BillingGatewayTerms) HasAccepts() bool`

HasAccepts returns a boolean if a field has been set.

### GetMaxCents

`func (o *BillingGatewayTerms) GetMaxCents() int64`

GetMaxCents returns the MaxCents field if non-nil, zero value otherwise.

### GetMaxCentsOk

`func (o *BillingGatewayTerms) GetMaxCentsOk() (*int64, bool)`

GetMaxCentsOk returns a tuple with the MaxCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCents

`func (o *BillingGatewayTerms) SetMaxCents(v int64)`

SetMaxCents sets MaxCents field to given value.

### HasMaxCents

`func (o *BillingGatewayTerms) HasMaxCents() bool`

HasMaxCents returns a boolean if a field has been set.

### GetMaxValiditySeconds

`func (o *BillingGatewayTerms) GetMaxValiditySeconds() int64`

GetMaxValiditySeconds returns the MaxValiditySeconds field if non-nil, zero value otherwise.

### GetMaxValiditySecondsOk

`func (o *BillingGatewayTerms) GetMaxValiditySecondsOk() (*int64, bool)`

GetMaxValiditySecondsOk returns a tuple with the MaxValiditySeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxValiditySeconds

`func (o *BillingGatewayTerms) SetMaxValiditySeconds(v int64)`

SetMaxValiditySeconds sets MaxValiditySeconds field to given value.

### HasMaxValiditySeconds

`func (o *BillingGatewayTerms) HasMaxValiditySeconds() bool`

HasMaxValiditySeconds returns a boolean if a field has been set.

### GetMinCents

`func (o *BillingGatewayTerms) GetMinCents() int64`

GetMinCents returns the MinCents field if non-nil, zero value otherwise.

### GetMinCentsOk

`func (o *BillingGatewayTerms) GetMinCentsOk() (*int64, bool)`

GetMinCentsOk returns a tuple with the MinCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinCents

`func (o *BillingGatewayTerms) SetMinCents(v int64)`

SetMinCents sets MinCents field to given value.

### HasMinCents

`func (o *BillingGatewayTerms) HasMinCents() bool`

HasMinCents returns a boolean if a field has been set.

### GetNonce

`func (o *BillingGatewayTerms) GetNonce() string`

GetNonce returns the Nonce field if non-nil, zero value otherwise.

### GetNonceOk

`func (o *BillingGatewayTerms) GetNonceOk() (*string, bool)`

GetNonceOk returns a tuple with the Nonce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonce

`func (o *BillingGatewayTerms) SetNonce(v string)`

SetNonce sets Nonce field to given value.

### HasNonce

`func (o *BillingGatewayTerms) HasNonce() bool`

HasNonce returns a boolean if a field has been set.

### GetPayTo

`func (o *BillingGatewayTerms) GetPayTo() string`

GetPayTo returns the PayTo field if non-nil, zero value otherwise.

### GetPayToOk

`func (o *BillingGatewayTerms) GetPayToOk() (*string, bool)`

GetPayToOk returns a tuple with the PayTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayTo

`func (o *BillingGatewayTerms) SetPayTo(v string)`

SetPayTo sets PayTo field to given value.

### HasPayTo

`func (o *BillingGatewayTerms) HasPayTo() bool`

HasPayTo returns a boolean if a field has been set.

### GetPrimaryType

`func (o *BillingGatewayTerms) GetPrimaryType() string`

GetPrimaryType returns the PrimaryType field if non-nil, zero value otherwise.

### GetPrimaryTypeOk

`func (o *BillingGatewayTerms) GetPrimaryTypeOk() (*string, bool)`

GetPrimaryTypeOk returns a tuple with the PrimaryType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryType

`func (o *BillingGatewayTerms) SetPrimaryType(v string)`

SetPrimaryType sets PrimaryType field to given value.

### HasPrimaryType

`func (o *BillingGatewayTerms) HasPrimaryType() bool`

HasPrimaryType returns a boolean if a field has been set.

### GetSalt

`func (o *BillingGatewayTerms) GetSalt() string`

GetSalt returns the Salt field if non-nil, zero value otherwise.

### GetSaltOk

`func (o *BillingGatewayTerms) GetSaltOk() (*string, bool)`

GetSaltOk returns a tuple with the Salt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalt

`func (o *BillingGatewayTerms) SetSalt(v string)`

SetSalt sets Salt field to given value.

### HasSalt

`func (o *BillingGatewayTerms) HasSalt() bool`

HasSalt returns a boolean if a field has been set.

### GetScheme

`func (o *BillingGatewayTerms) GetScheme() string`

GetScheme returns the Scheme field if non-nil, zero value otherwise.

### GetSchemeOk

`func (o *BillingGatewayTerms) GetSchemeOk() (*string, bool)`

GetSchemeOk returns a tuple with the Scheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheme

`func (o *BillingGatewayTerms) SetScheme(v string)`

SetScheme sets Scheme field to given value.

### HasScheme

`func (o *BillingGatewayTerms) HasScheme() bool`

HasScheme returns a boolean if a field has been set.

### GetTypes

`func (o *BillingGatewayTerms) GetTypes() interface{}`

GetTypes returns the Types field if non-nil, zero value otherwise.

### GetTypesOk

`func (o *BillingGatewayTerms) GetTypesOk() (*interface{}, bool)`

GetTypesOk returns a tuple with the Types field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypes

`func (o *BillingGatewayTerms) SetTypes(v interface{})`

SetTypes sets Types field to given value.

### HasTypes

`func (o *BillingGatewayTerms) HasTypes() bool`

HasTypes returns a boolean if a field has been set.

### SetTypesNil

`func (o *BillingGatewayTerms) SetTypesNil(b bool)`

 SetTypesNil sets the value for Types to be an explicit nil

### UnsetTypes
`func (o *BillingGatewayTerms) UnsetTypes()`

UnsetTypes ensures that no value is present for Types, not even an explicit nil
### GetX402Version

`func (o *BillingGatewayTerms) GetX402Version() int64`

GetX402Version returns the X402Version field if non-nil, zero value otherwise.

### GetX402VersionOk

`func (o *BillingGatewayTerms) GetX402VersionOk() (*int64, bool)`

GetX402VersionOk returns a tuple with the X402Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetX402Version

`func (o *BillingGatewayTerms) SetX402Version(v int64)`

SetX402Version sets X402Version field to given value.

### HasX402Version

`func (o *BillingGatewayTerms) HasX402Version() bool`

HasX402Version returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


