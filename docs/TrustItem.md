# TrustItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Available** | Pointer to **string** | Available is \&quot;now\&quot; when the item can be read immediately, or \&quot;on request\&quot; when it is released only to a party who asks and is answered. It is the one field a page renders the difference from. | [optional] 
**Body** | Pointer to **string** | Body is the item&#39;s content, for the kinds that are text rather than a file — an article, a subprocessor entry, a note. Empty for anything released on request: a summary of a document is still the document. | [optional] 
**Framework** | Pointer to **string** | Framework is the standard the item speaks to, when it speaks to one. | [optional] 
**Id** | Pointer to **string** | ID addresses the item — for reading it if it is available now, or for naming it in a request if it is not. | [optional] 
**Kind** | Pointer to **string** | Kind is what the item is: report, letter, policy, questionnaire, subprocessor, article or update. | [optional] 
**Name** | Pointer to **string** | Name is the item&#39;s title. | [optional] 
**Signed** | Pointer to **string** | Signed is \&quot;self\&quot; when the org states it itself and \&quot;auditor\&quot; when an independent auditor put their name to it. It is the reason an item is available now or on request, so a reader can see the rule rather than infer it. | [optional] 
**Summary** | Pointer to **string** | Summary is a line about the item. | [optional] 
**UpdatedAt** | Pointer to **int32** | UpdatedAt is when the item last changed, in unix milliseconds. | [optional] 

## Methods

### NewTrustItem

`func NewTrustItem() *TrustItem`

NewTrustItem instantiates a new TrustItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrustItemWithDefaults

`func NewTrustItemWithDefaults() *TrustItem`

NewTrustItemWithDefaults instantiates a new TrustItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvailable

`func (o *TrustItem) GetAvailable() string`

GetAvailable returns the Available field if non-nil, zero value otherwise.

### GetAvailableOk

`func (o *TrustItem) GetAvailableOk() (*string, bool)`

GetAvailableOk returns a tuple with the Available field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailable

`func (o *TrustItem) SetAvailable(v string)`

SetAvailable sets Available field to given value.

### HasAvailable

`func (o *TrustItem) HasAvailable() bool`

HasAvailable returns a boolean if a field has been set.

### GetBody

`func (o *TrustItem) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *TrustItem) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *TrustItem) SetBody(v string)`

SetBody sets Body field to given value.

### HasBody

`func (o *TrustItem) HasBody() bool`

HasBody returns a boolean if a field has been set.

### GetFramework

`func (o *TrustItem) GetFramework() string`

GetFramework returns the Framework field if non-nil, zero value otherwise.

### GetFrameworkOk

`func (o *TrustItem) GetFrameworkOk() (*string, bool)`

GetFrameworkOk returns a tuple with the Framework field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFramework

`func (o *TrustItem) SetFramework(v string)`

SetFramework sets Framework field to given value.

### HasFramework

`func (o *TrustItem) HasFramework() bool`

HasFramework returns a boolean if a field has been set.

### GetId

`func (o *TrustItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TrustItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TrustItem) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TrustItem) HasId() bool`

HasId returns a boolean if a field has been set.

### GetKind

`func (o *TrustItem) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *TrustItem) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *TrustItem) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *TrustItem) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetName

`func (o *TrustItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TrustItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TrustItem) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TrustItem) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSigned

`func (o *TrustItem) GetSigned() string`

GetSigned returns the Signed field if non-nil, zero value otherwise.

### GetSignedOk

`func (o *TrustItem) GetSignedOk() (*string, bool)`

GetSignedOk returns a tuple with the Signed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigned

`func (o *TrustItem) SetSigned(v string)`

SetSigned sets Signed field to given value.

### HasSigned

`func (o *TrustItem) HasSigned() bool`

HasSigned returns a boolean if a field has been set.

### GetSummary

`func (o *TrustItem) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *TrustItem) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *TrustItem) SetSummary(v string)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *TrustItem) HasSummary() bool`

HasSummary returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *TrustItem) GetUpdatedAt() int32`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *TrustItem) GetUpdatedAtOk() (*int32, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *TrustItem) SetUpdatedAt(v int32)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *TrustItem) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


