# CloudDomainAddIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name is the custom domain to attach, e.g. \&quot;www.acme.com\&quot;. | [optional] 
**Project** | Pointer to **string** | Project is the Pages project name, from the path. | [optional] 

## Methods

### NewCloudDomainAddIn

`func NewCloudDomainAddIn() *CloudDomainAddIn`

NewCloudDomainAddIn instantiates a new CloudDomainAddIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudDomainAddInWithDefaults

`func NewCloudDomainAddInWithDefaults() *CloudDomainAddIn`

NewCloudDomainAddInWithDefaults instantiates a new CloudDomainAddIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CloudDomainAddIn) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CloudDomainAddIn) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CloudDomainAddIn) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CloudDomainAddIn) HasName() bool`

HasName returns a boolean if a field has been set.

### GetProject

`func (o *CloudDomainAddIn) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *CloudDomainAddIn) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *CloudDomainAddIn) SetProject(v string)`

SetProject sets Project field to given value.

### HasProject

`func (o *CloudDomainAddIn) HasProject() bool`

HasProject returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


