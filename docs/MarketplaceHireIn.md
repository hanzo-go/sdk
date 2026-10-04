# MarketplaceHireIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | Pointer to **string** | Amount is what the job pays, U.S. dollars to the cent: \&quot;250.00\&quot;. | [optional] 
**Attempt** | Pointer to **string** | Attempt is the buyer&#39;s own key for this hire: chosen once, when the hire starts — a random id, 1 to 128 letters, digits, dots, colons, dashes or underscores — and sent with every request it makes, the quote and the payment alike, with the rest of the request unchanged. The same attempt answers the same job: the quote&#39;s terms to sign until it is paid, the job once it opened, whatever it has come to since. Another request under an attempt already used is refused (422): a new hire is a new attempt. | [optional] 
**Brief** | Pointer to **string** | Brief is what the buyer asks for, and what delivered looks like. Required, at most 4096 characters. | [optional] 
**Category** | Pointer to **string** | Category is what the payment is for: service (the default), goods, transfer or royalty. | [optional] 
**Deadline** | Pointer to **int64** | Deadline is when delivery must land by, unix seconds. Without one, it is the latest the payment leaves room for: its expiry, less the review window, fourteen days for a ruling on a dispute, and a day for the clock. | [optional] 
**Listing** | Pointer to **string** | Listing is the public listing the job is hired through. Empty for a direct offer, which names Seller instead. | [optional] 
**Payment** | Pointer to **string** | Payment is the buyer&#39;s signed x402 authorization for the job — the value PAYMENT-SIGNATURE carries, for a caller that sends no headers of its own, over MCP or the command line. The header wins when both are sent. | [optional] 
**Performed** | Pointer to **string** | Performed is where the work is performed, ISO 3166-1 alpha-2. It decides whether a foreign seller&#39;s pay is U.S.-source. | [optional] 
**Rail** | Pointer to **string** | Rail is x402 (the default) or chain. | [optional] 
**Review** | Pointer to **int64** | Review is how many seconds the buyer has after delivery to release or dispute: three days by default, thirty at most. | [optional] 
**Seller** | Pointer to **string** | Seller is the org offered the job directly; only without a listing. | [optional] 
**Title** | Pointer to **string** | Title names a direct offer; a listing&#39;s job takes the listing&#39;s title. | [optional] 
**Wallet** | Pointer to **string** | Wallet is the buyer&#39;s wallet the payment is signed from. When named, the authorization must be signed by its address. | [optional] 

## Methods

### NewMarketplaceHireIn

`func NewMarketplaceHireIn() *MarketplaceHireIn`

NewMarketplaceHireIn instantiates a new MarketplaceHireIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarketplaceHireInWithDefaults

`func NewMarketplaceHireInWithDefaults() *MarketplaceHireIn`

NewMarketplaceHireInWithDefaults instantiates a new MarketplaceHireIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *MarketplaceHireIn) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *MarketplaceHireIn) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *MarketplaceHireIn) SetAmount(v string)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *MarketplaceHireIn) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetAttempt

`func (o *MarketplaceHireIn) GetAttempt() string`

GetAttempt returns the Attempt field if non-nil, zero value otherwise.

### GetAttemptOk

`func (o *MarketplaceHireIn) GetAttemptOk() (*string, bool)`

GetAttemptOk returns a tuple with the Attempt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttempt

`func (o *MarketplaceHireIn) SetAttempt(v string)`

SetAttempt sets Attempt field to given value.

### HasAttempt

`func (o *MarketplaceHireIn) HasAttempt() bool`

HasAttempt returns a boolean if a field has been set.

### GetBrief

`func (o *MarketplaceHireIn) GetBrief() string`

GetBrief returns the Brief field if non-nil, zero value otherwise.

### GetBriefOk

`func (o *MarketplaceHireIn) GetBriefOk() (*string, bool)`

GetBriefOk returns a tuple with the Brief field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrief

`func (o *MarketplaceHireIn) SetBrief(v string)`

SetBrief sets Brief field to given value.

### HasBrief

`func (o *MarketplaceHireIn) HasBrief() bool`

HasBrief returns a boolean if a field has been set.

### GetCategory

`func (o *MarketplaceHireIn) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *MarketplaceHireIn) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *MarketplaceHireIn) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *MarketplaceHireIn) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetDeadline

`func (o *MarketplaceHireIn) GetDeadline() int64`

GetDeadline returns the Deadline field if non-nil, zero value otherwise.

### GetDeadlineOk

`func (o *MarketplaceHireIn) GetDeadlineOk() (*int64, bool)`

GetDeadlineOk returns a tuple with the Deadline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeadline

`func (o *MarketplaceHireIn) SetDeadline(v int64)`

SetDeadline sets Deadline field to given value.

### HasDeadline

`func (o *MarketplaceHireIn) HasDeadline() bool`

HasDeadline returns a boolean if a field has been set.

### GetListing

`func (o *MarketplaceHireIn) GetListing() string`

GetListing returns the Listing field if non-nil, zero value otherwise.

### GetListingOk

`func (o *MarketplaceHireIn) GetListingOk() (*string, bool)`

GetListingOk returns a tuple with the Listing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListing

`func (o *MarketplaceHireIn) SetListing(v string)`

SetListing sets Listing field to given value.

### HasListing

`func (o *MarketplaceHireIn) HasListing() bool`

HasListing returns a boolean if a field has been set.

### GetPayment

`func (o *MarketplaceHireIn) GetPayment() string`

GetPayment returns the Payment field if non-nil, zero value otherwise.

### GetPaymentOk

`func (o *MarketplaceHireIn) GetPaymentOk() (*string, bool)`

GetPaymentOk returns a tuple with the Payment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayment

`func (o *MarketplaceHireIn) SetPayment(v string)`

SetPayment sets Payment field to given value.

### HasPayment

`func (o *MarketplaceHireIn) HasPayment() bool`

HasPayment returns a boolean if a field has been set.

### GetPerformed

`func (o *MarketplaceHireIn) GetPerformed() string`

GetPerformed returns the Performed field if non-nil, zero value otherwise.

### GetPerformedOk

`func (o *MarketplaceHireIn) GetPerformedOk() (*string, bool)`

GetPerformedOk returns a tuple with the Performed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerformed

`func (o *MarketplaceHireIn) SetPerformed(v string)`

SetPerformed sets Performed field to given value.

### HasPerformed

`func (o *MarketplaceHireIn) HasPerformed() bool`

HasPerformed returns a boolean if a field has been set.

### GetRail

`func (o *MarketplaceHireIn) GetRail() string`

GetRail returns the Rail field if non-nil, zero value otherwise.

### GetRailOk

`func (o *MarketplaceHireIn) GetRailOk() (*string, bool)`

GetRailOk returns a tuple with the Rail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRail

`func (o *MarketplaceHireIn) SetRail(v string)`

SetRail sets Rail field to given value.

### HasRail

`func (o *MarketplaceHireIn) HasRail() bool`

HasRail returns a boolean if a field has been set.

### GetReview

`func (o *MarketplaceHireIn) GetReview() int64`

GetReview returns the Review field if non-nil, zero value otherwise.

### GetReviewOk

`func (o *MarketplaceHireIn) GetReviewOk() (*int64, bool)`

GetReviewOk returns a tuple with the Review field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReview

`func (o *MarketplaceHireIn) SetReview(v int64)`

SetReview sets Review field to given value.

### HasReview

`func (o *MarketplaceHireIn) HasReview() bool`

HasReview returns a boolean if a field has been set.

### GetSeller

`func (o *MarketplaceHireIn) GetSeller() string`

GetSeller returns the Seller field if non-nil, zero value otherwise.

### GetSellerOk

`func (o *MarketplaceHireIn) GetSellerOk() (*string, bool)`

GetSellerOk returns a tuple with the Seller field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeller

`func (o *MarketplaceHireIn) SetSeller(v string)`

SetSeller sets Seller field to given value.

### HasSeller

`func (o *MarketplaceHireIn) HasSeller() bool`

HasSeller returns a boolean if a field has been set.

### GetTitle

`func (o *MarketplaceHireIn) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *MarketplaceHireIn) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *MarketplaceHireIn) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *MarketplaceHireIn) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetWallet

`func (o *MarketplaceHireIn) GetWallet() string`

GetWallet returns the Wallet field if non-nil, zero value otherwise.

### GetWalletOk

`func (o *MarketplaceHireIn) GetWalletOk() (*string, bool)`

GetWalletOk returns a tuple with the Wallet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWallet

`func (o *MarketplaceHireIn) SetWallet(v string)`

SetWallet sets Wallet field to given value.

### HasWallet

`func (o *MarketplaceHireIn) HasWallet() bool`

HasWallet returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


