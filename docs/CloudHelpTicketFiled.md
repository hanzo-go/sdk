# CloudHelpTicketFiled

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** | Status is the lifecycle state the ticket was filed in — always \&quot;Open\&quot;. | [optional] 
**Ticket** | Pointer to **string** | Ticket is the opaque, random customer-facing reference (\&quot;tkt_\&quot; + 24 hex characters). It is NOT the ticket&#39;s internal name: that name is sequential, and handing it out would disclose the center&#39;s ticket volume. | [optional] 

## Methods

### NewCloudHelpTicketFiled

`func NewCloudHelpTicketFiled() *CloudHelpTicketFiled`

NewCloudHelpTicketFiled instantiates a new CloudHelpTicketFiled object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudHelpTicketFiledWithDefaults

`func NewCloudHelpTicketFiledWithDefaults() *CloudHelpTicketFiled`

NewCloudHelpTicketFiledWithDefaults instantiates a new CloudHelpTicketFiled object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *CloudHelpTicketFiled) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CloudHelpTicketFiled) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CloudHelpTicketFiled) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CloudHelpTicketFiled) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTicket

`func (o *CloudHelpTicketFiled) GetTicket() string`

GetTicket returns the Ticket field if non-nil, zero value otherwise.

### GetTicketOk

`func (o *CloudHelpTicketFiled) GetTicketOk() (*string, bool)`

GetTicketOk returns a tuple with the Ticket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTicket

`func (o *CloudHelpTicketFiled) SetTicket(v string)`

SetTicket sets Ticket field to given value.

### HasTicket

`func (o *CloudHelpTicketFiled) HasTicket() bool`

HasTicket returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


