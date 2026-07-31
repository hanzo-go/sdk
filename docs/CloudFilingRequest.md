# CloudFilingRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DocumentIds** | Pointer to **[]string** | DocumentIDs are the documents to file. At least one is required, and every one must belong to the caller&#39;s org — a filing can never reach across orgs. | [optional] 
**Jurisdiction** | Pointer to **string** | Jurisdiction is the state or agency the filing is for, e.g. \&quot;DE\&quot;. | [optional] 

## Methods

### NewCloudFilingRequest

`func NewCloudFilingRequest() *CloudFilingRequest`

NewCloudFilingRequest instantiates a new CloudFilingRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudFilingRequestWithDefaults

`func NewCloudFilingRequestWithDefaults() *CloudFilingRequest`

NewCloudFilingRequestWithDefaults instantiates a new CloudFilingRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDocumentIds

`func (o *CloudFilingRequest) GetDocumentIds() []string`

GetDocumentIds returns the DocumentIds field if non-nil, zero value otherwise.

### GetDocumentIdsOk

`func (o *CloudFilingRequest) GetDocumentIdsOk() (*[]string, bool)`

GetDocumentIdsOk returns a tuple with the DocumentIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentIds

`func (o *CloudFilingRequest) SetDocumentIds(v []string)`

SetDocumentIds sets DocumentIds field to given value.

### HasDocumentIds

`func (o *CloudFilingRequest) HasDocumentIds() bool`

HasDocumentIds returns a boolean if a field has been set.

### GetJurisdiction

`func (o *CloudFilingRequest) GetJurisdiction() string`

GetJurisdiction returns the Jurisdiction field if non-nil, zero value otherwise.

### GetJurisdictionOk

`func (o *CloudFilingRequest) GetJurisdictionOk() (*string, bool)`

GetJurisdictionOk returns a tuple with the Jurisdiction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJurisdiction

`func (o *CloudFilingRequest) SetJurisdiction(v string)`

SetJurisdiction sets Jurisdiction field to given value.

### HasJurisdiction

`func (o *CloudFilingRequest) HasJurisdiction() bool`

HasJurisdiction returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


