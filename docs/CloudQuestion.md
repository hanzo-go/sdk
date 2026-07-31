# CloudQuestion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Account** | Pointer to **string** |  | [optional] 
**Amount** | Pointer to **string** | formatted figure ($…) | [optional] 
**Id** | Pointer to **string** | the source transaction id it concerns | [optional] 
**Kind** | Pointer to **string** | outlier|reversal|roundoff|uncosted|overdrawn | [optional] 
**PostedAt** | Pointer to **string** |  | [optional] 
**Text** | Pointer to **string** | the specific question to ask the founder | [optional] 

## Methods

### NewCloudQuestion

`func NewCloudQuestion() *CloudQuestion`

NewCloudQuestion instantiates a new CloudQuestion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudQuestionWithDefaults

`func NewCloudQuestionWithDefaults() *CloudQuestion`

NewCloudQuestionWithDefaults instantiates a new CloudQuestion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccount

`func (o *CloudQuestion) GetAccount() string`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *CloudQuestion) GetAccountOk() (*string, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *CloudQuestion) SetAccount(v string)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *CloudQuestion) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetAmount

`func (o *CloudQuestion) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *CloudQuestion) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *CloudQuestion) SetAmount(v string)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *CloudQuestion) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetId

`func (o *CloudQuestion) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CloudQuestion) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CloudQuestion) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CloudQuestion) HasId() bool`

HasId returns a boolean if a field has been set.

### GetKind

`func (o *CloudQuestion) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *CloudQuestion) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *CloudQuestion) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *CloudQuestion) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetPostedAt

`func (o *CloudQuestion) GetPostedAt() string`

GetPostedAt returns the PostedAt field if non-nil, zero value otherwise.

### GetPostedAtOk

`func (o *CloudQuestion) GetPostedAtOk() (*string, bool)`

GetPostedAtOk returns a tuple with the PostedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostedAt

`func (o *CloudQuestion) SetPostedAt(v string)`

SetPostedAt sets PostedAt field to given value.

### HasPostedAt

`func (o *CloudQuestion) HasPostedAt() bool`

HasPostedAt returns a boolean if a field has been set.

### GetText

`func (o *CloudQuestion) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *CloudQuestion) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *CloudQuestion) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *CloudQuestion) HasText() bool`

HasText returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


