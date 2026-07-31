# CloudListingPage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Listings** | Pointer to [**[]CloudListing**](CloudListing.md) | Listings is every listing this org has published, private ones included (Public says which are discoverable by others). | [optional] 

## Methods

### NewCloudListingPage

`func NewCloudListingPage() *CloudListingPage`

NewCloudListingPage instantiates a new CloudListingPage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudListingPageWithDefaults

`func NewCloudListingPageWithDefaults() *CloudListingPage`

NewCloudListingPageWithDefaults instantiates a new CloudListingPage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetListings

`func (o *CloudListingPage) GetListings() []CloudListing`

GetListings returns the Listings field if non-nil, zero value otherwise.

### GetListingsOk

`func (o *CloudListingPage) GetListingsOk() (*[]CloudListing, bool)`

GetListingsOk returns a tuple with the Listings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListings

`func (o *CloudListingPage) SetListings(v []CloudListing)`

SetListings sets Listings field to given value.

### HasListings

`func (o *CloudListingPage) HasListings() bool`

HasListings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


