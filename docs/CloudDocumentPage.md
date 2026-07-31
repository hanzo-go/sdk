# CloudDocumentPage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]CloudDocumentSummary**](CloudDocumentSummary.md) | Data are the documents, WITHOUT their rendered content — fetch one to read it. | [optional] 
**Disclaimer** | Pointer to **string** | Disclaimer is the boundary made visible on the wire. | [optional] 

## Methods

### NewCloudDocumentPage

`func NewCloudDocumentPage() *CloudDocumentPage`

NewCloudDocumentPage instantiates a new CloudDocumentPage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudDocumentPageWithDefaults

`func NewCloudDocumentPageWithDefaults() *CloudDocumentPage`

NewCloudDocumentPageWithDefaults instantiates a new CloudDocumentPage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *CloudDocumentPage) GetData() []CloudDocumentSummary`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CloudDocumentPage) GetDataOk() (*[]CloudDocumentSummary, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CloudDocumentPage) SetData(v []CloudDocumentSummary)`

SetData sets Data field to given value.

### HasData

`func (o *CloudDocumentPage) HasData() bool`

HasData returns a boolean if a field has been set.

### GetDisclaimer

`func (o *CloudDocumentPage) GetDisclaimer() string`

GetDisclaimer returns the Disclaimer field if non-nil, zero value otherwise.

### GetDisclaimerOk

`func (o *CloudDocumentPage) GetDisclaimerOk() (*string, bool)`

GetDisclaimerOk returns a tuple with the Disclaimer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisclaimer

`func (o *CloudDocumentPage) SetDisclaimer(v string)`

SetDisclaimer sets Disclaimer field to given value.

### HasDisclaimer

`func (o *CloudDocumentPage) HasDisclaimer() bool`

HasDisclaimer returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


