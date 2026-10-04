# ContentTransitionResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Distribution** | Pointer to [**ContentPublishResult**](ContentPublishResult.md) | Distribution is the channel fan-out this move triggered. Present ONLY on the move to published, the single edge that distributes — so its absence means no fan-out was attempted, never that one failed quietly. A fan-out that DID fail is present carrying its own honest status, because distribution never rolls the status change back. | [optional] 
**Doctype** | Pointer to **string** | DocType is the content type that moved — Campaign, SocialPost or Asset — echoed from the path. | [optional] 
**From** | Pointer to **string** | From is the state the item held when it was read. A document carrying no status yet reads as \&quot;draft\&quot;. | [optional] 
**Name** | Pointer to **string** | Name is the document that moved, echoed from the path. | [optional] 
**Storefront** | Pointer to [**ContentStorefrontResult**](ContentStorefrontResult.md) | Storefront is the catalog side effect, present only when a published Asset was product imagery — it carries a design and a kind of ecom, product or lifestyle. Absent for everything else, so absence reads as \&quot;not catalog imagery\&quot; rather than \&quot;the catalog failed\&quot;. | [optional] 
**To** | Pointer to **string** | To is the state it holds now. From &#x3D;&#x3D; To on an idempotent re-transition, which is legal and is where a caller that lost a publish race lands. | [optional] 

## Methods

### NewContentTransitionResult

`func NewContentTransitionResult() *ContentTransitionResult`

NewContentTransitionResult instantiates a new ContentTransitionResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContentTransitionResultWithDefaults

`func NewContentTransitionResultWithDefaults() *ContentTransitionResult`

NewContentTransitionResultWithDefaults instantiates a new ContentTransitionResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDistribution

`func (o *ContentTransitionResult) GetDistribution() ContentPublishResult`

GetDistribution returns the Distribution field if non-nil, zero value otherwise.

### GetDistributionOk

`func (o *ContentTransitionResult) GetDistributionOk() (*ContentPublishResult, bool)`

GetDistributionOk returns a tuple with the Distribution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistribution

`func (o *ContentTransitionResult) SetDistribution(v ContentPublishResult)`

SetDistribution sets Distribution field to given value.

### HasDistribution

`func (o *ContentTransitionResult) HasDistribution() bool`

HasDistribution returns a boolean if a field has been set.

### GetDoctype

`func (o *ContentTransitionResult) GetDoctype() string`

GetDoctype returns the Doctype field if non-nil, zero value otherwise.

### GetDoctypeOk

`func (o *ContentTransitionResult) GetDoctypeOk() (*string, bool)`

GetDoctypeOk returns a tuple with the Doctype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDoctype

`func (o *ContentTransitionResult) SetDoctype(v string)`

SetDoctype sets Doctype field to given value.

### HasDoctype

`func (o *ContentTransitionResult) HasDoctype() bool`

HasDoctype returns a boolean if a field has been set.

### GetFrom

`func (o *ContentTransitionResult) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *ContentTransitionResult) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *ContentTransitionResult) SetFrom(v string)`

SetFrom sets From field to given value.

### HasFrom

`func (o *ContentTransitionResult) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetName

`func (o *ContentTransitionResult) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ContentTransitionResult) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ContentTransitionResult) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ContentTransitionResult) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStorefront

`func (o *ContentTransitionResult) GetStorefront() ContentStorefrontResult`

GetStorefront returns the Storefront field if non-nil, zero value otherwise.

### GetStorefrontOk

`func (o *ContentTransitionResult) GetStorefrontOk() (*ContentStorefrontResult, bool)`

GetStorefrontOk returns a tuple with the Storefront field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorefront

`func (o *ContentTransitionResult) SetStorefront(v ContentStorefrontResult)`

SetStorefront sets Storefront field to given value.

### HasStorefront

`func (o *ContentTransitionResult) HasStorefront() bool`

HasStorefront returns a boolean if a field has been set.

### GetTo

`func (o *ContentTransitionResult) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *ContentTransitionResult) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *ContentTransitionResult) SetTo(v string)`

SetTo sets To field to given value.

### HasTo

`func (o *ContentTransitionResult) HasTo() bool`

HasTo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


