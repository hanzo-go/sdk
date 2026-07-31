# CloudPairingApproved

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OwnerBootstrapped** | Pointer to **bool** | OwnerBootstrapped is true when this approval was the org&#39;s FIRST on the channel and therefore also made the sender its owner. | [optional] 
**Sender** | Pointer to **string** | Sender is the external chat identity that is now allowed to DM the org&#39;s bot. | [optional] 

## Methods

### NewCloudPairingApproved

`func NewCloudPairingApproved() *CloudPairingApproved`

NewCloudPairingApproved instantiates a new CloudPairingApproved object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudPairingApprovedWithDefaults

`func NewCloudPairingApprovedWithDefaults() *CloudPairingApproved`

NewCloudPairingApprovedWithDefaults instantiates a new CloudPairingApproved object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOwnerBootstrapped

`func (o *CloudPairingApproved) GetOwnerBootstrapped() bool`

GetOwnerBootstrapped returns the OwnerBootstrapped field if non-nil, zero value otherwise.

### GetOwnerBootstrappedOk

`func (o *CloudPairingApproved) GetOwnerBootstrappedOk() (*bool, bool)`

GetOwnerBootstrappedOk returns a tuple with the OwnerBootstrapped field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerBootstrapped

`func (o *CloudPairingApproved) SetOwnerBootstrapped(v bool)`

SetOwnerBootstrapped sets OwnerBootstrapped field to given value.

### HasOwnerBootstrapped

`func (o *CloudPairingApproved) HasOwnerBootstrapped() bool`

HasOwnerBootstrapped returns a boolean if a field has been set.

### GetSender

`func (o *CloudPairingApproved) GetSender() string`

GetSender returns the Sender field if non-nil, zero value otherwise.

### GetSenderOk

`func (o *CloudPairingApproved) GetSenderOk() (*string, bool)`

GetSenderOk returns a tuple with the Sender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSender

`func (o *CloudPairingApproved) SetSender(v string)`

SetSender sets Sender field to given value.

### HasSender

`func (o *CloudPairingApproved) HasSender() bool`

HasSender returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


