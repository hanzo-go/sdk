# SecuritySecurityGetScan200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Scan** | Pointer to [**SecurityScan**](SecurityScan.md) |  | [optional] 
**Findings** | Pointer to [**[]SecurityFinding**](SecurityFinding.md) |  | [optional] 

## Methods

### NewSecuritySecurityGetScan200Response

`func NewSecuritySecurityGetScan200Response() *SecuritySecurityGetScan200Response`

NewSecuritySecurityGetScan200Response instantiates a new SecuritySecurityGetScan200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecuritySecurityGetScan200ResponseWithDefaults

`func NewSecuritySecurityGetScan200ResponseWithDefaults() *SecuritySecurityGetScan200Response`

NewSecuritySecurityGetScan200ResponseWithDefaults instantiates a new SecuritySecurityGetScan200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetScan

`func (o *SecuritySecurityGetScan200Response) GetScan() SecurityScan`

GetScan returns the Scan field if non-nil, zero value otherwise.

### GetScanOk

`func (o *SecuritySecurityGetScan200Response) GetScanOk() (*SecurityScan, bool)`

GetScanOk returns a tuple with the Scan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScan

`func (o *SecuritySecurityGetScan200Response) SetScan(v SecurityScan)`

SetScan sets Scan field to given value.

### HasScan

`func (o *SecuritySecurityGetScan200Response) HasScan() bool`

HasScan returns a boolean if a field has been set.

### GetFindings

`func (o *SecuritySecurityGetScan200Response) GetFindings() []SecurityFinding`

GetFindings returns the Findings field if non-nil, zero value otherwise.

### GetFindingsOk

`func (o *SecuritySecurityGetScan200Response) GetFindingsOk() (*[]SecurityFinding, bool)`

GetFindingsOk returns a tuple with the Findings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFindings

`func (o *SecuritySecurityGetScan200Response) SetFindings(v []SecurityFinding)`

SetFindings sets Findings field to given value.

### HasFindings

`func (o *SecuritySecurityGetScan200Response) HasFindings() bool`

HasFindings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


