# MarketplaceJob

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | Pointer to **interface{}** |  | [optional] 
**Attempt** | Pointer to **string** | Attempt is the buyer&#39;s own key for the hire that made the job, when it sent one. | [optional] 
**Brief** | Pointer to **string** | Brief is what the buyer asked for. | [optional] 
**BuyerOrg** | Pointer to **string** | BuyerOrg is the org that pays. | [optional] 
**Category** | Pointer to **string** | Category is what the payment is for — service, goods, transfer or royalty — as the buyer declared it and the seller accepted it. It is what the economic event says. | [optional] 
**Clearance** | Pointer to **string** | Clearance is principal&#39;s latest decision that the buyer may pay the seller this amount — asked when the job was quoted, again when it was funded, and again when it was released: GET /v1/principal/clearance/{id}. | [optional] 
**CreatedAt** | Pointer to **int64** | CreatedAt is when the job was quoted, unix seconds. | [optional] 
**Currency** | Pointer to **string** | Currency labels Amount: USD. | [optional] 
**Deadline** | Pointer to **int64** | Deadline is when delivery must land by, unix seconds. | [optional] 
**Delivery** | Pointer to [**MarketplaceDelivery**](MarketplaceDelivery.md) | Delivery is what the seller delivered, once it has. | [optional] 
**Dispute** | Pointer to [**MarketplaceDispute**](MarketplaceDispute.md) | Dispute is why the job was stopped, once it was. | [optional] 
**Ending** | Pointer to **string** | Ending is the unpaid ending the job was given — declined, cancelled or refunded — while the rail returns the amount it set aside. It takes no other step meanwhile, and Status becomes Ending once the money is back. | [optional] 
**Escrow** | Pointer to [**MarketplaceEscrow**](MarketplaceEscrow.md) | Escrow is how the job is paid. | [optional] 
**History** | Pointer to [**[]MarketplaceStep**](MarketplaceStep.md) | History is every step the job took, oldest first. | [optional] 
**Id** | Pointer to **string** | ID is the job, \&quot;job_\&quot;-prefixed. | [optional] 
**Listing** | Pointer to **string** | Listing is the listing it was hired through; empty for a direct offer. | [optional] 
**Performed** | Pointer to **string** | Performed is where the work is performed, ISO 3166-1 alpha-2, when said. | [optional] 
**Review** | Pointer to **int64** | Review is how many seconds after delivery the buyer has to release or dispute before the job releases itself. | [optional] 
**SellerOrg** | Pointer to **string** | SellerOrg is the org that does the work and is paid. | [optional] 
**Status** | Pointer to **string** | Status is where the job stands: quoted or funding, which only its buyer sees, then open, accepted, delivered, released, disputed, declined, cancelled or refunded. | [optional] 
**Title** | Pointer to **string** | Title names the work: the listing&#39;s title, or the offer&#39;s own. | [optional] 
**UpdatedAt** | Pointer to **int64** | UpdatedAt is when it last moved, unix seconds. | [optional] 

## Methods

### NewMarketplaceJob

`func NewMarketplaceJob() *MarketplaceJob`

NewMarketplaceJob instantiates a new MarketplaceJob object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarketplaceJobWithDefaults

`func NewMarketplaceJobWithDefaults() *MarketplaceJob`

NewMarketplaceJobWithDefaults instantiates a new MarketplaceJob object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *MarketplaceJob) GetAmount() interface{}`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *MarketplaceJob) GetAmountOk() (*interface{}, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *MarketplaceJob) SetAmount(v interface{})`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *MarketplaceJob) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### SetAmountNil

`func (o *MarketplaceJob) SetAmountNil(b bool)`

 SetAmountNil sets the value for Amount to be an explicit nil

### UnsetAmount
`func (o *MarketplaceJob) UnsetAmount()`

UnsetAmount ensures that no value is present for Amount, not even an explicit nil
### GetAttempt

`func (o *MarketplaceJob) GetAttempt() string`

GetAttempt returns the Attempt field if non-nil, zero value otherwise.

### GetAttemptOk

`func (o *MarketplaceJob) GetAttemptOk() (*string, bool)`

GetAttemptOk returns a tuple with the Attempt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttempt

`func (o *MarketplaceJob) SetAttempt(v string)`

SetAttempt sets Attempt field to given value.

### HasAttempt

`func (o *MarketplaceJob) HasAttempt() bool`

HasAttempt returns a boolean if a field has been set.

### GetBrief

`func (o *MarketplaceJob) GetBrief() string`

GetBrief returns the Brief field if non-nil, zero value otherwise.

### GetBriefOk

`func (o *MarketplaceJob) GetBriefOk() (*string, bool)`

GetBriefOk returns a tuple with the Brief field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrief

`func (o *MarketplaceJob) SetBrief(v string)`

SetBrief sets Brief field to given value.

### HasBrief

`func (o *MarketplaceJob) HasBrief() bool`

HasBrief returns a boolean if a field has been set.

### GetBuyerOrg

`func (o *MarketplaceJob) GetBuyerOrg() string`

GetBuyerOrg returns the BuyerOrg field if non-nil, zero value otherwise.

### GetBuyerOrgOk

`func (o *MarketplaceJob) GetBuyerOrgOk() (*string, bool)`

GetBuyerOrgOk returns a tuple with the BuyerOrg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuyerOrg

`func (o *MarketplaceJob) SetBuyerOrg(v string)`

SetBuyerOrg sets BuyerOrg field to given value.

### HasBuyerOrg

`func (o *MarketplaceJob) HasBuyerOrg() bool`

HasBuyerOrg returns a boolean if a field has been set.

### GetCategory

`func (o *MarketplaceJob) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *MarketplaceJob) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *MarketplaceJob) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *MarketplaceJob) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetClearance

`func (o *MarketplaceJob) GetClearance() string`

GetClearance returns the Clearance field if non-nil, zero value otherwise.

### GetClearanceOk

`func (o *MarketplaceJob) GetClearanceOk() (*string, bool)`

GetClearanceOk returns a tuple with the Clearance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClearance

`func (o *MarketplaceJob) SetClearance(v string)`

SetClearance sets Clearance field to given value.

### HasClearance

`func (o *MarketplaceJob) HasClearance() bool`

HasClearance returns a boolean if a field has been set.

### GetCreatedAt

`func (o *MarketplaceJob) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *MarketplaceJob) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *MarketplaceJob) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *MarketplaceJob) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCurrency

`func (o *MarketplaceJob) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *MarketplaceJob) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *MarketplaceJob) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *MarketplaceJob) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetDeadline

`func (o *MarketplaceJob) GetDeadline() int64`

GetDeadline returns the Deadline field if non-nil, zero value otherwise.

### GetDeadlineOk

`func (o *MarketplaceJob) GetDeadlineOk() (*int64, bool)`

GetDeadlineOk returns a tuple with the Deadline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeadline

`func (o *MarketplaceJob) SetDeadline(v int64)`

SetDeadline sets Deadline field to given value.

### HasDeadline

`func (o *MarketplaceJob) HasDeadline() bool`

HasDeadline returns a boolean if a field has been set.

### GetDelivery

`func (o *MarketplaceJob) GetDelivery() MarketplaceDelivery`

GetDelivery returns the Delivery field if non-nil, zero value otherwise.

### GetDeliveryOk

`func (o *MarketplaceJob) GetDeliveryOk() (*MarketplaceDelivery, bool)`

GetDeliveryOk returns a tuple with the Delivery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelivery

`func (o *MarketplaceJob) SetDelivery(v MarketplaceDelivery)`

SetDelivery sets Delivery field to given value.

### HasDelivery

`func (o *MarketplaceJob) HasDelivery() bool`

HasDelivery returns a boolean if a field has been set.

### GetDispute

`func (o *MarketplaceJob) GetDispute() MarketplaceDispute`

GetDispute returns the Dispute field if non-nil, zero value otherwise.

### GetDisputeOk

`func (o *MarketplaceJob) GetDisputeOk() (*MarketplaceDispute, bool)`

GetDisputeOk returns a tuple with the Dispute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDispute

`func (o *MarketplaceJob) SetDispute(v MarketplaceDispute)`

SetDispute sets Dispute field to given value.

### HasDispute

`func (o *MarketplaceJob) HasDispute() bool`

HasDispute returns a boolean if a field has been set.

### GetEnding

`func (o *MarketplaceJob) GetEnding() string`

GetEnding returns the Ending field if non-nil, zero value otherwise.

### GetEndingOk

`func (o *MarketplaceJob) GetEndingOk() (*string, bool)`

GetEndingOk returns a tuple with the Ending field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnding

`func (o *MarketplaceJob) SetEnding(v string)`

SetEnding sets Ending field to given value.

### HasEnding

`func (o *MarketplaceJob) HasEnding() bool`

HasEnding returns a boolean if a field has been set.

### GetEscrow

`func (o *MarketplaceJob) GetEscrow() MarketplaceEscrow`

GetEscrow returns the Escrow field if non-nil, zero value otherwise.

### GetEscrowOk

`func (o *MarketplaceJob) GetEscrowOk() (*MarketplaceEscrow, bool)`

GetEscrowOk returns a tuple with the Escrow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEscrow

`func (o *MarketplaceJob) SetEscrow(v MarketplaceEscrow)`

SetEscrow sets Escrow field to given value.

### HasEscrow

`func (o *MarketplaceJob) HasEscrow() bool`

HasEscrow returns a boolean if a field has been set.

### GetHistory

`func (o *MarketplaceJob) GetHistory() []MarketplaceStep`

GetHistory returns the History field if non-nil, zero value otherwise.

### GetHistoryOk

`func (o *MarketplaceJob) GetHistoryOk() (*[]MarketplaceStep, bool)`

GetHistoryOk returns a tuple with the History field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHistory

`func (o *MarketplaceJob) SetHistory(v []MarketplaceStep)`

SetHistory sets History field to given value.

### HasHistory

`func (o *MarketplaceJob) HasHistory() bool`

HasHistory returns a boolean if a field has been set.

### GetId

`func (o *MarketplaceJob) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MarketplaceJob) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MarketplaceJob) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MarketplaceJob) HasId() bool`

HasId returns a boolean if a field has been set.

### GetListing

`func (o *MarketplaceJob) GetListing() string`

GetListing returns the Listing field if non-nil, zero value otherwise.

### GetListingOk

`func (o *MarketplaceJob) GetListingOk() (*string, bool)`

GetListingOk returns a tuple with the Listing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListing

`func (o *MarketplaceJob) SetListing(v string)`

SetListing sets Listing field to given value.

### HasListing

`func (o *MarketplaceJob) HasListing() bool`

HasListing returns a boolean if a field has been set.

### GetPerformed

`func (o *MarketplaceJob) GetPerformed() string`

GetPerformed returns the Performed field if non-nil, zero value otherwise.

### GetPerformedOk

`func (o *MarketplaceJob) GetPerformedOk() (*string, bool)`

GetPerformedOk returns a tuple with the Performed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerformed

`func (o *MarketplaceJob) SetPerformed(v string)`

SetPerformed sets Performed field to given value.

### HasPerformed

`func (o *MarketplaceJob) HasPerformed() bool`

HasPerformed returns a boolean if a field has been set.

### GetReview

`func (o *MarketplaceJob) GetReview() int64`

GetReview returns the Review field if non-nil, zero value otherwise.

### GetReviewOk

`func (o *MarketplaceJob) GetReviewOk() (*int64, bool)`

GetReviewOk returns a tuple with the Review field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReview

`func (o *MarketplaceJob) SetReview(v int64)`

SetReview sets Review field to given value.

### HasReview

`func (o *MarketplaceJob) HasReview() bool`

HasReview returns a boolean if a field has been set.

### GetSellerOrg

`func (o *MarketplaceJob) GetSellerOrg() string`

GetSellerOrg returns the SellerOrg field if non-nil, zero value otherwise.

### GetSellerOrgOk

`func (o *MarketplaceJob) GetSellerOrgOk() (*string, bool)`

GetSellerOrgOk returns a tuple with the SellerOrg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSellerOrg

`func (o *MarketplaceJob) SetSellerOrg(v string)`

SetSellerOrg sets SellerOrg field to given value.

### HasSellerOrg

`func (o *MarketplaceJob) HasSellerOrg() bool`

HasSellerOrg returns a boolean if a field has been set.

### GetStatus

`func (o *MarketplaceJob) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *MarketplaceJob) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *MarketplaceJob) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *MarketplaceJob) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTitle

`func (o *MarketplaceJob) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *MarketplaceJob) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *MarketplaceJob) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *MarketplaceJob) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *MarketplaceJob) GetUpdatedAt() int64`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *MarketplaceJob) GetUpdatedAtOk() (*int64, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *MarketplaceJob) SetUpdatedAt(v int64)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *MarketplaceJob) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


