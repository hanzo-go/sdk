# PairingApproved

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OwnerBootstrapped** | Pointer to **bool** | OwnerBootstrapped is true when this approval was the org&#39;s FIRST on the channel and therefore also made the sender its owner. | [optional] 
**Sender** | Pointer to **string** | Sender is the external chat identity that is now allowed to DM the org&#39;s bot. | [optional] 

## Methods

### NewPairingApproved

`func NewPairingApproved() *PairingApproved`

NewPairingApproved instantiates a new PairingApproved object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPairingApprovedWithDefaults

`func NewPairingApprovedWithDefaults() *PairingApproved`

NewPairingApprovedWithDefaults instantiates a new PairingApproved object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOwnerBootstrapped

`func (o *PairingApproved) GetOwnerBootstrapped() bool`

GetOwnerBootstrapped returns the OwnerBootstrapped field if non-nil, zero value otherwise.

### GetOwnerBootstrappedOk

`func (o *PairingApproved) GetOwnerBootstrappedOk() (*bool, bool)`

GetOwnerBootstrappedOk returns a tuple with the OwnerBootstrapped field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerBootstrapped

`func (o *PairingApproved) SetOwnerBootstrapped(v bool)`

SetOwnerBootstrapped sets OwnerBootstrapped field to given value.

### HasOwnerBootstrapped

`func (o *PairingApproved) HasOwnerBootstrapped() bool`

HasOwnerBootstrapped returns a boolean if a field has been set.

### GetSender

`func (o *PairingApproved) GetSender() string`

GetSender returns the Sender field if non-nil, zero value otherwise.

### GetSenderOk

`func (o *PairingApproved) GetSenderOk() (*string, bool)`

GetSenderOk returns a tuple with the Sender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSender

`func (o *PairingApproved) SetSender(v string)`

SetSender sets Sender field to given value.

### HasSender

`func (o *PairingApproved) HasSender() bool`

HasSender returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


