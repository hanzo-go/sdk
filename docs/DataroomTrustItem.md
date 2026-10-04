# DataroomTrustItem

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
**UpdatedAt** | Pointer to **int64** | UpdatedAt is when the item last changed, in unix milliseconds. | [optional] 

## Methods

### NewDataroomTrustItem

`func NewDataroomTrustItem() *DataroomTrustItem`

NewDataroomTrustItem instantiates a new DataroomTrustItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataroomTrustItemWithDefaults

`func NewDataroomTrustItemWithDefaults() *DataroomTrustItem`

NewDataroomTrustItemWithDefaults instantiates a new DataroomTrustItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvailable

`func (o *DataroomTrustItem) GetAvailable() string`

GetAvailable returns the Available field if non-nil, zero value otherwise.

### GetAvailableOk

`func (o *DataroomTrustItem) GetAvailableOk() (*string, bool)`

GetAvailableOk returns a tuple with the Available field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailable

`func (o *DataroomTrustItem) SetAvailable(v string)`

SetAvailable sets Available field to given value.

### HasAvailable

`func (o *DataroomTrustItem) HasAvailable() bool`

HasAvailable returns a boolean if a field has been set.

### GetBody

`func (o *DataroomTrustItem) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *DataroomTrustItem) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *DataroomTrustItem) SetBody(v string)`

SetBody sets Body field to given value.

### HasBody

`func (o *DataroomTrustItem) HasBody() bool`

HasBody returns a boolean if a field has been set.

### GetFramework

`func (o *DataroomTrustItem) GetFramework() string`

GetFramework returns the Framework field if non-nil, zero value otherwise.

### GetFrameworkOk

`func (o *DataroomTrustItem) GetFrameworkOk() (*string, bool)`

GetFrameworkOk returns a tuple with the Framework field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFramework

`func (o *DataroomTrustItem) SetFramework(v string)`

SetFramework sets Framework field to given value.

### HasFramework

`func (o *DataroomTrustItem) HasFramework() bool`

HasFramework returns a boolean if a field has been set.

### GetId

`func (o *DataroomTrustItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DataroomTrustItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DataroomTrustItem) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DataroomTrustItem) HasId() bool`

HasId returns a boolean if a field has been set.

### GetKind

`func (o *DataroomTrustItem) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *DataroomTrustItem) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *DataroomTrustItem) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *DataroomTrustItem) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetName

`func (o *DataroomTrustItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DataroomTrustItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DataroomTrustItem) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DataroomTrustItem) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSigned

`func (o *DataroomTrustItem) GetSigned() string`

GetSigned returns the Signed field if non-nil, zero value otherwise.

### GetSignedOk

`func (o *DataroomTrustItem) GetSignedOk() (*string, bool)`

GetSignedOk returns a tuple with the Signed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigned

`func (o *DataroomTrustItem) SetSigned(v string)`

SetSigned sets Signed field to given value.

### HasSigned

`func (o *DataroomTrustItem) HasSigned() bool`

HasSigned returns a boolean if a field has been set.

### GetSummary

`func (o *DataroomTrustItem) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *DataroomTrustItem) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *DataroomTrustItem) SetSummary(v string)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *DataroomTrustItem) HasSummary() bool`

HasSummary returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *DataroomTrustItem) GetUpdatedAt() int64`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *DataroomTrustItem) GetUpdatedAtOk() (*int64, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *DataroomTrustItem) SetUpdatedAt(v int64)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *DataroomTrustItem) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


