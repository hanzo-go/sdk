# TrustPublish

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attester** | Pointer to **string** | Attester is who vouched for it: \&quot;self\&quot; for anything the org states itself, or \&quot;auditor\&quot; for anything an independent auditor put their name to. REQUIRED, and anything other than \&quot;self\&quot; is read as \&quot;auditor\&quot; — the safe direction, since an auditor-signed item can only ever be released on request. | [optional] 
**Body** | Pointer to **string** | Body is the item&#39;s content for the kinds that are text rather than a file: an article, a subprocessor entry, a dated note. | [optional] 
**Document** | Pointer to **string** | Document is a data-room document holding the item&#39;s bytes, uploaded first through POST /v1/dataroom/documents. Optional: an item can be content with no file. The document must already exist in the caller org&#39;s own store. | [optional] 
**Framework** | Pointer to **string** | Framework is the standard it speaks to. Optional and free text — the value is the org&#39;s own, not a list this API keeps. | [optional] 
**Kind** | Pointer to **string** | Kind is what the item is: report, letter, policy, questionnaire, subprocessor, article or update. Required. | [optional] 
**Name** | Pointer to **string** | Name is the item&#39;s title. Required. | [optional] 
**Summary** | Pointer to **string** | Summary is a line about it. Optional. | [optional] 
**Tier** | Pointer to **string** | Tier is who may read it: \&quot;public\&quot; or \&quot;gated\&quot;. It DEFAULTS TO GATED and anything that is not exactly \&quot;public\&quot; is gated, so an item published by a caller that says nothing is private and someone has to release it on purpose. \&quot;public\&quot; is refused for an auditor-signed item. | [optional] 

## Methods

### NewTrustPublish

`func NewTrustPublish() *TrustPublish`

NewTrustPublish instantiates a new TrustPublish object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrustPublishWithDefaults

`func NewTrustPublishWithDefaults() *TrustPublish`

NewTrustPublishWithDefaults instantiates a new TrustPublish object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttester

`func (o *TrustPublish) GetAttester() string`

GetAttester returns the Attester field if non-nil, zero value otherwise.

### GetAttesterOk

`func (o *TrustPublish) GetAttesterOk() (*string, bool)`

GetAttesterOk returns a tuple with the Attester field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttester

`func (o *TrustPublish) SetAttester(v string)`

SetAttester sets Attester field to given value.

### HasAttester

`func (o *TrustPublish) HasAttester() bool`

HasAttester returns a boolean if a field has been set.

### GetBody

`func (o *TrustPublish) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *TrustPublish) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *TrustPublish) SetBody(v string)`

SetBody sets Body field to given value.

### HasBody

`func (o *TrustPublish) HasBody() bool`

HasBody returns a boolean if a field has been set.

### GetDocument

`func (o *TrustPublish) GetDocument() string`

GetDocument returns the Document field if non-nil, zero value otherwise.

### GetDocumentOk

`func (o *TrustPublish) GetDocumentOk() (*string, bool)`

GetDocumentOk returns a tuple with the Document field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocument

`func (o *TrustPublish) SetDocument(v string)`

SetDocument sets Document field to given value.

### HasDocument

`func (o *TrustPublish) HasDocument() bool`

HasDocument returns a boolean if a field has been set.

### GetFramework

`func (o *TrustPublish) GetFramework() string`

GetFramework returns the Framework field if non-nil, zero value otherwise.

### GetFrameworkOk

`func (o *TrustPublish) GetFrameworkOk() (*string, bool)`

GetFrameworkOk returns a tuple with the Framework field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFramework

`func (o *TrustPublish) SetFramework(v string)`

SetFramework sets Framework field to given value.

### HasFramework

`func (o *TrustPublish) HasFramework() bool`

HasFramework returns a boolean if a field has been set.

### GetKind

`func (o *TrustPublish) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *TrustPublish) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *TrustPublish) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *TrustPublish) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetName

`func (o *TrustPublish) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TrustPublish) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TrustPublish) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TrustPublish) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSummary

`func (o *TrustPublish) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *TrustPublish) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *TrustPublish) SetSummary(v string)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *TrustPublish) HasSummary() bool`

HasSummary returns a boolean if a field has been set.

### GetTier

`func (o *TrustPublish) GetTier() string`

GetTier returns the Tier field if non-nil, zero value otherwise.

### GetTierOk

`func (o *TrustPublish) GetTierOk() (*string, bool)`

GetTierOk returns a tuple with the Tier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTier

`func (o *TrustPublish) SetTier(v string)`

SetTier sets Tier field to given value.

### HasTier

`func (o *TrustPublish) HasTier() bool`

HasTier returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


