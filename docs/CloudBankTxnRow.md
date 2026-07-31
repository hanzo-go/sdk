# CloudBankTxnRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AmountCents** | Pointer to **int32** |  | [optional] 
**Connector** | Pointer to **string** |  | [optional] 
**Currency** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Direction** | Pointer to **string** |  | [optional] 
**ExternalId** | Pointer to **string** |  | [optional] 
**MatchedVoucher** | Pointer to **string** |  | [optional] 
**Merchant** | Pointer to **string** |  | [optional] 
**PostedAt** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 

## Methods

### NewCloudBankTxnRow

`func NewCloudBankTxnRow() *CloudBankTxnRow`

NewCloudBankTxnRow instantiates a new CloudBankTxnRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudBankTxnRowWithDefaults

`func NewCloudBankTxnRowWithDefaults() *CloudBankTxnRow`

NewCloudBankTxnRowWithDefaults instantiates a new CloudBankTxnRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmountCents

`func (o *CloudBankTxnRow) GetAmountCents() int32`

GetAmountCents returns the AmountCents field if non-nil, zero value otherwise.

### GetAmountCentsOk

`func (o *CloudBankTxnRow) GetAmountCentsOk() (*int32, bool)`

GetAmountCentsOk returns a tuple with the AmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountCents

`func (o *CloudBankTxnRow) SetAmountCents(v int32)`

SetAmountCents sets AmountCents field to given value.

### HasAmountCents

`func (o *CloudBankTxnRow) HasAmountCents() bool`

HasAmountCents returns a boolean if a field has been set.

### GetConnector

`func (o *CloudBankTxnRow) GetConnector() string`

GetConnector returns the Connector field if non-nil, zero value otherwise.

### GetConnectorOk

`func (o *CloudBankTxnRow) GetConnectorOk() (*string, bool)`

GetConnectorOk returns a tuple with the Connector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnector

`func (o *CloudBankTxnRow) SetConnector(v string)`

SetConnector sets Connector field to given value.

### HasConnector

`func (o *CloudBankTxnRow) HasConnector() bool`

HasConnector returns a boolean if a field has been set.

### GetCurrency

`func (o *CloudBankTxnRow) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *CloudBankTxnRow) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *CloudBankTxnRow) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *CloudBankTxnRow) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetDescription

`func (o *CloudBankTxnRow) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CloudBankTxnRow) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CloudBankTxnRow) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CloudBankTxnRow) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDirection

`func (o *CloudBankTxnRow) GetDirection() string`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *CloudBankTxnRow) GetDirectionOk() (*string, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *CloudBankTxnRow) SetDirection(v string)`

SetDirection sets Direction field to given value.

### HasDirection

`func (o *CloudBankTxnRow) HasDirection() bool`

HasDirection returns a boolean if a field has been set.

### GetExternalId

`func (o *CloudBankTxnRow) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *CloudBankTxnRow) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *CloudBankTxnRow) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *CloudBankTxnRow) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetMatchedVoucher

`func (o *CloudBankTxnRow) GetMatchedVoucher() string`

GetMatchedVoucher returns the MatchedVoucher field if non-nil, zero value otherwise.

### GetMatchedVoucherOk

`func (o *CloudBankTxnRow) GetMatchedVoucherOk() (*string, bool)`

GetMatchedVoucherOk returns a tuple with the MatchedVoucher field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchedVoucher

`func (o *CloudBankTxnRow) SetMatchedVoucher(v string)`

SetMatchedVoucher sets MatchedVoucher field to given value.

### HasMatchedVoucher

`func (o *CloudBankTxnRow) HasMatchedVoucher() bool`

HasMatchedVoucher returns a boolean if a field has been set.

### GetMerchant

`func (o *CloudBankTxnRow) GetMerchant() string`

GetMerchant returns the Merchant field if non-nil, zero value otherwise.

### GetMerchantOk

`func (o *CloudBankTxnRow) GetMerchantOk() (*string, bool)`

GetMerchantOk returns a tuple with the Merchant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchant

`func (o *CloudBankTxnRow) SetMerchant(v string)`

SetMerchant sets Merchant field to given value.

### HasMerchant

`func (o *CloudBankTxnRow) HasMerchant() bool`

HasMerchant returns a boolean if a field has been set.

### GetPostedAt

`func (o *CloudBankTxnRow) GetPostedAt() string`

GetPostedAt returns the PostedAt field if non-nil, zero value otherwise.

### GetPostedAtOk

`func (o *CloudBankTxnRow) GetPostedAtOk() (*string, bool)`

GetPostedAtOk returns a tuple with the PostedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostedAt

`func (o *CloudBankTxnRow) SetPostedAt(v string)`

SetPostedAt sets PostedAt field to given value.

### HasPostedAt

`func (o *CloudBankTxnRow) HasPostedAt() bool`

HasPostedAt returns a boolean if a field has been set.

### GetStatus

`func (o *CloudBankTxnRow) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CloudBankTxnRow) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CloudBankTxnRow) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CloudBankTxnRow) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


