# DataroomTrustPublish

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

### NewDataroomTrustPublish

`func NewDataroomTrustPublish() *DataroomTrustPublish`

NewDataroomTrustPublish instantiates a new DataroomTrustPublish object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataroomTrustPublishWithDefaults

`func NewDataroomTrustPublishWithDefaults() *DataroomTrustPublish`

NewDataroomTrustPublishWithDefaults instantiates a new DataroomTrustPublish object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttester

`func (o *DataroomTrustPublish) GetAttester() string`

GetAttester returns the Attester field if non-nil, zero value otherwise.

### GetAttesterOk

`func (o *DataroomTrustPublish) GetAttesterOk() (*string, bool)`

GetAttesterOk returns a tuple with the Attester field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttester

`func (o *DataroomTrustPublish) SetAttester(v string)`

SetAttester sets Attester field to given value.

### HasAttester

`func (o *DataroomTrustPublish) HasAttester() bool`

HasAttester returns a boolean if a field has been set.

### GetBody

`func (o *DataroomTrustPublish) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *DataroomTrustPublish) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *DataroomTrustPublish) SetBody(v string)`

SetBody sets Body field to given value.

### HasBody

`func (o *DataroomTrustPublish) HasBody() bool`

HasBody returns a boolean if a field has been set.

### GetDocument

`func (o *DataroomTrustPublish) GetDocument() string`

GetDocument returns the Document field if non-nil, zero value otherwise.

### GetDocumentOk

`func (o *DataroomTrustPublish) GetDocumentOk() (*string, bool)`

GetDocumentOk returns a tuple with the Document field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocument

`func (o *DataroomTrustPublish) SetDocument(v string)`

SetDocument sets Document field to given value.

### HasDocument

`func (o *DataroomTrustPublish) HasDocument() bool`

HasDocument returns a boolean if a field has been set.

### GetFramework

`func (o *DataroomTrustPublish) GetFramework() string`

GetFramework returns the Framework field if non-nil, zero value otherwise.

### GetFrameworkOk

`func (o *DataroomTrustPublish) GetFrameworkOk() (*string, bool)`

GetFrameworkOk returns a tuple with the Framework field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFramework

`func (o *DataroomTrustPublish) SetFramework(v string)`

SetFramework sets Framework field to given value.

### HasFramework

`func (o *DataroomTrustPublish) HasFramework() bool`

HasFramework returns a boolean if a field has been set.

### GetKind

`func (o *DataroomTrustPublish) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *DataroomTrustPublish) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *DataroomTrustPublish) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *DataroomTrustPublish) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetName

`func (o *DataroomTrustPublish) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DataroomTrustPublish) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DataroomTrustPublish) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DataroomTrustPublish) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSummary

`func (o *DataroomTrustPublish) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *DataroomTrustPublish) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *DataroomTrustPublish) SetSummary(v string)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *DataroomTrustPublish) HasSummary() bool`

HasSummary returns a boolean if a field has been set.

### GetTier

`func (o *DataroomTrustPublish) GetTier() string`

GetTier returns the Tier field if non-nil, zero value otherwise.

### GetTierOk

`func (o *DataroomTrustPublish) GetTierOk() (*string, bool)`

GetTierOk returns a tuple with the Tier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTier

`func (o *DataroomTrustPublish) SetTier(v string)`

SetTier sets Tier field to given value.

### HasTier

`func (o *DataroomTrustPublish) HasTier() bool`

HasTier returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


