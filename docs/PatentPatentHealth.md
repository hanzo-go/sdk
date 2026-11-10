# PatentPatentHealth

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Anchored** | Pointer to **bool** | Anchored is whether record roots are committed outside this platform. False is the common case and the important one: without it the chain proves order WITHIN the platform&#39;s own honesty and nothing beyond it. | [optional] 
**Durable** | Pointer to **bool** | Durable is whether acknowledged writes are fenced to a durable object. False means a store is local to this pod: the records stand, failover does not carry them. | [optional] 
**Now** | Pointer to **string** | RecordedAt is the server&#39;s time now, RFC3339 — so a reader can tell how stale its own view is. | [optional] 
**SignatureVerification** | Pointer to **bool** | SignatureVerification is whether the server can verify a submitted detached signature. Attribution does not depend on it — the author of every record is the validated principal, never a body field — but a record carrying a signature it cannot check says so. | [optional] 
**StampedBy** | Pointer to **string** | ServerClockSays is what stamps a record: always \&quot;server\&quot;, because no client-supplied timestamp is read anywhere in this package. | [optional] 

## Methods

### NewPatentPatentHealth

`func NewPatentPatentHealth() *PatentPatentHealth`

NewPatentPatentHealth instantiates a new PatentPatentHealth object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatentPatentHealthWithDefaults

`func NewPatentPatentHealthWithDefaults() *PatentPatentHealth`

NewPatentPatentHealthWithDefaults instantiates a new PatentPatentHealth object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnchored

`func (o *PatentPatentHealth) GetAnchored() bool`

GetAnchored returns the Anchored field if non-nil, zero value otherwise.

### GetAnchoredOk

`func (o *PatentPatentHealth) GetAnchoredOk() (*bool, bool)`

GetAnchoredOk returns a tuple with the Anchored field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnchored

`func (o *PatentPatentHealth) SetAnchored(v bool)`

SetAnchored sets Anchored field to given value.

### HasAnchored

`func (o *PatentPatentHealth) HasAnchored() bool`

HasAnchored returns a boolean if a field has been set.

### GetDurable

`func (o *PatentPatentHealth) GetDurable() bool`

GetDurable returns the Durable field if non-nil, zero value otherwise.

### GetDurableOk

`func (o *PatentPatentHealth) GetDurableOk() (*bool, bool)`

GetDurableOk returns a tuple with the Durable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurable

`func (o *PatentPatentHealth) SetDurable(v bool)`

SetDurable sets Durable field to given value.

### HasDurable

`func (o *PatentPatentHealth) HasDurable() bool`

HasDurable returns a boolean if a field has been set.

### GetNow

`func (o *PatentPatentHealth) GetNow() string`

GetNow returns the Now field if non-nil, zero value otherwise.

### GetNowOk

`func (o *PatentPatentHealth) GetNowOk() (*string, bool)`

GetNowOk returns a tuple with the Now field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNow

`func (o *PatentPatentHealth) SetNow(v string)`

SetNow sets Now field to given value.

### HasNow

`func (o *PatentPatentHealth) HasNow() bool`

HasNow returns a boolean if a field has been set.

### GetSignatureVerification

`func (o *PatentPatentHealth) GetSignatureVerification() bool`

GetSignatureVerification returns the SignatureVerification field if non-nil, zero value otherwise.

### GetSignatureVerificationOk

`func (o *PatentPatentHealth) GetSignatureVerificationOk() (*bool, bool)`

GetSignatureVerificationOk returns a tuple with the SignatureVerification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignatureVerification

`func (o *PatentPatentHealth) SetSignatureVerification(v bool)`

SetSignatureVerification sets SignatureVerification field to given value.

### HasSignatureVerification

`func (o *PatentPatentHealth) HasSignatureVerification() bool`

HasSignatureVerification returns a boolean if a field has been set.

### GetStampedBy

`func (o *PatentPatentHealth) GetStampedBy() string`

GetStampedBy returns the StampedBy field if non-nil, zero value otherwise.

### GetStampedByOk

`func (o *PatentPatentHealth) GetStampedByOk() (*string, bool)`

GetStampedByOk returns a tuple with the StampedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStampedBy

`func (o *PatentPatentHealth) SetStampedBy(v string)`

SetStampedBy sets StampedBy field to given value.

### HasStampedBy

`func (o *PatentPatentHealth) HasStampedBy() bool`

HasStampedBy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


