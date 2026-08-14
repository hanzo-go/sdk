# FilingReply

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Disclaimer** | Pointer to **string** | Disclaimer is the boundary made visible on the wire. | [optional] 
**Filing** | Pointer to [**LegalFiling**](LegalFiling.md) | Filing is the tracking record. | [optional] 

## Methods

### NewFilingReply

`func NewFilingReply() *FilingReply`

NewFilingReply instantiates a new FilingReply object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFilingReplyWithDefaults

`func NewFilingReplyWithDefaults() *FilingReply`

NewFilingReplyWithDefaults instantiates a new FilingReply object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisclaimer

`func (o *FilingReply) GetDisclaimer() string`

GetDisclaimer returns the Disclaimer field if non-nil, zero value otherwise.

### GetDisclaimerOk

`func (o *FilingReply) GetDisclaimerOk() (*string, bool)`

GetDisclaimerOk returns a tuple with the Disclaimer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisclaimer

`func (o *FilingReply) SetDisclaimer(v string)`

SetDisclaimer sets Disclaimer field to given value.

### HasDisclaimer

`func (o *FilingReply) HasDisclaimer() bool`

HasDisclaimer returns a boolean if a field has been set.

### GetFiling

`func (o *FilingReply) GetFiling() LegalFiling`

GetFiling returns the Filing field if non-nil, zero value otherwise.

### GetFilingOk

`func (o *FilingReply) GetFilingOk() (*LegalFiling, bool)`

GetFilingOk returns a tuple with the Filing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiling

`func (o *FilingReply) SetFiling(v LegalFiling)`

SetFiling sets Filing field to given value.

### HasFiling

`func (o *FilingReply) HasFiling() bool`

HasFiling returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


