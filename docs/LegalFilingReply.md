# LegalFilingReply

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Disclaimer** | Pointer to **string** | Disclaimer is the boundary made visible on the wire. | [optional] 
**Filing** | Pointer to [**LegalLegalFiling**](LegalLegalFiling.md) | Filing is the tracking record. | [optional] 

## Methods

### NewLegalFilingReply

`func NewLegalFilingReply() *LegalFilingReply`

NewLegalFilingReply instantiates a new LegalFilingReply object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLegalFilingReplyWithDefaults

`func NewLegalFilingReplyWithDefaults() *LegalFilingReply`

NewLegalFilingReplyWithDefaults instantiates a new LegalFilingReply object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisclaimer

`func (o *LegalFilingReply) GetDisclaimer() string`

GetDisclaimer returns the Disclaimer field if non-nil, zero value otherwise.

### GetDisclaimerOk

`func (o *LegalFilingReply) GetDisclaimerOk() (*string, bool)`

GetDisclaimerOk returns a tuple with the Disclaimer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisclaimer

`func (o *LegalFilingReply) SetDisclaimer(v string)`

SetDisclaimer sets Disclaimer field to given value.

### HasDisclaimer

`func (o *LegalFilingReply) HasDisclaimer() bool`

HasDisclaimer returns a boolean if a field has been set.

### GetFiling

`func (o *LegalFilingReply) GetFiling() LegalLegalFiling`

GetFiling returns the Filing field if non-nil, zero value otherwise.

### GetFilingOk

`func (o *LegalFilingReply) GetFilingOk() (*LegalLegalFiling, bool)`

GetFilingOk returns a tuple with the Filing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiling

`func (o *LegalFilingReply) SetFiling(v LegalLegalFiling)`

SetFiling sets Filing field to given value.

### HasFiling

`func (o *LegalFilingReply) HasFiling() bool`

HasFiling returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


