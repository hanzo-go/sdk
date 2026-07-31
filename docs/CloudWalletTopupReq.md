# CloudWalletTopupReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FromAddress** | Pointer to **string** | FromAddress is the wallet the transfer was sent from. Optional; when given it must match the transfer&#39;s on-chain sender. | [optional] 
**Rail** | Pointer to **string** | Which accepted rail the transfer was sent on, e.g. \&quot;base-usdc\&quot;. The client names it rather than the server guessing from the tx: the same address can exist on several chains, so inferring would risk crediting against the wrong treasury. It may be omitted only while exactly one rail is enabled. | [optional] 
**TxHash** | Pointer to **string** | TxHash is the hash of the ERC-20 transfer that was already sent to the rail&#39;s treasury. The receipt is read from that chain; nothing is credited that the chain did not confirm. | [optional] 

## Methods

### NewCloudWalletTopupReq

`func NewCloudWalletTopupReq() *CloudWalletTopupReq`

NewCloudWalletTopupReq instantiates a new CloudWalletTopupReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudWalletTopupReqWithDefaults

`func NewCloudWalletTopupReqWithDefaults() *CloudWalletTopupReq`

NewCloudWalletTopupReqWithDefaults instantiates a new CloudWalletTopupReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFromAddress

`func (o *CloudWalletTopupReq) GetFromAddress() string`

GetFromAddress returns the FromAddress field if non-nil, zero value otherwise.

### GetFromAddressOk

`func (o *CloudWalletTopupReq) GetFromAddressOk() (*string, bool)`

GetFromAddressOk returns a tuple with the FromAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromAddress

`func (o *CloudWalletTopupReq) SetFromAddress(v string)`

SetFromAddress sets FromAddress field to given value.

### HasFromAddress

`func (o *CloudWalletTopupReq) HasFromAddress() bool`

HasFromAddress returns a boolean if a field has been set.

### GetRail

`func (o *CloudWalletTopupReq) GetRail() string`

GetRail returns the Rail field if non-nil, zero value otherwise.

### GetRailOk

`func (o *CloudWalletTopupReq) GetRailOk() (*string, bool)`

GetRailOk returns a tuple with the Rail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRail

`func (o *CloudWalletTopupReq) SetRail(v string)`

SetRail sets Rail field to given value.

### HasRail

`func (o *CloudWalletTopupReq) HasRail() bool`

HasRail returns a boolean if a field has been set.

### GetTxHash

`func (o *CloudWalletTopupReq) GetTxHash() string`

GetTxHash returns the TxHash field if non-nil, zero value otherwise.

### GetTxHashOk

`func (o *CloudWalletTopupReq) GetTxHashOk() (*string, bool)`

GetTxHashOk returns a tuple with the TxHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxHash

`func (o *CloudWalletTopupReq) SetTxHash(v string)`

SetTxHash sets TxHash field to given value.

### HasTxHash

`func (o *CloudWalletTopupReq) HasTxHash() bool`

HasTxHash returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


