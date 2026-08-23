# \CommerceAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteCommerceCollectionByCollectionid**](CommerceAPI.md#DeleteCommerceCollectionByCollectionid) | **Delete** /v1/commerce/collection/{collectionid} | Delete a collection, keeping a recoverable copy
[**DeleteCommerceDisclosureByDisclosureid**](CommerceAPI.md#DeleteCommerceDisclosureByDisclosureid) | **Delete** /v1/commerce/disclosure/{disclosureid} | Delete a disclosure, keeping a recoverable copy
[**DeleteCommerceDiscountByDiscountid**](CommerceAPI.md#DeleteCommerceDiscountByDiscountid) | **Delete** /v1/commerce/discount/{discountid} | Delete a discount, keeping a recoverable copy
[**DeleteCommerceMovieByMovieid**](CommerceAPI.md#DeleteCommerceMovieByMovieid) | **Delete** /v1/commerce/movie/{movieid} | Delete a movie, keeping a recoverable copy
[**DeleteCommerceNoteByNoteid**](CommerceAPI.md#DeleteCommerceNoteByNoteid) | **Delete** /v1/commerce/note/{noteid} | Delete a note, keeping a recoverable copy
[**DeleteCommercePlansEntriesBySlug**](CommerceAPI.md#DeleteCommercePlansEntriesBySlug) | **Delete** /v1/commerce/plans/entries/{slug} | Remove a plan from the authority
[**DeleteCommerceProductByProductid**](CommerceAPI.md#DeleteCommerceProductByProductid) | **Delete** /v1/commerce/product/{productid} | Delete a product, keeping a recoverable copy
[**DeleteCommerceRatesEntriesByProductByMeter**](CommerceAPI.md#DeleteCommerceRatesEntriesByProductByMeter) | **Delete** /v1/commerce/rates/entries/{product}/{meter} | Remove a rate outright
[**DeleteCommerceReturnByReturnid**](CommerceAPI.md#DeleteCommerceReturnByReturnid) | **Delete** /v1/commerce/return/{returnid} | Delete a return, keeping a recoverable copy
[**DeleteCommerceSaleschannelBySaleschannelid**](CommerceAPI.md#DeleteCommerceSaleschannelBySaleschannelid) | **Delete** /v1/commerce/saleschannel/{saleschannelid} | Delete a sales channel, keeping a recoverable copy
[**DeleteCommerceStocklocationByStocklocationid**](CommerceAPI.md#DeleteCommerceStocklocationByStocklocationid) | **Delete** /v1/commerce/stocklocation/{stocklocationid} | Delete a stock location, keeping a recoverable copy
[**DeleteCommerceStoreByStoreid**](CommerceAPI.md#DeleteCommerceStoreByStoreid) | **Delete** /v1/commerce/store/{storeid} | Delete a storefront, keeping a recoverable copy
[**DeleteCommerceStoreByStoreidListingByKey**](CommerceAPI.md#DeleteCommerceStoreByStoreidListingByKey) | **Delete** /v1/commerce/store/{storeid}/listing/{key} | Remove a listing override
[**DeleteCommerceSubmissionBySubmissionid**](CommerceAPI.md#DeleteCommerceSubmissionBySubmissionid) | **Delete** /v1/commerce/submission/{submissionid} | Delete a submission, keeping a recoverable copy
[**DeleteCommerceSubscriberBySubscriberid**](CommerceAPI.md#DeleteCommerceSubscriberBySubscriberid) | **Delete** /v1/commerce/subscriber/{subscriberid} | Delete a subscriber, keeping a recoverable copy
[**DeleteCommerceTokentransactionByTokentransactionid**](CommerceAPI.md#DeleteCommerceTokentransactionByTokentransactionid) | **Delete** /v1/commerce/tokentransaction/{tokentransactionid} | Delete a token transaction, keeping a recoverable copy
[**DeleteCommerceTransferByTransferid**](CommerceAPI.md#DeleteCommerceTransferByTransferid) | **Delete** /v1/commerce/transfer/{transferid} | Delete a transfer, keeping a recoverable copy
[**DeleteCommerceVariantByVariantid**](CommerceAPI.md#DeleteCommerceVariantByVariantid) | **Delete** /v1/commerce/variant/{variantid} | Delete a variant, keeping a recoverable copy
[**DeleteCommerceWalletByWalletid**](CommerceAPI.md#DeleteCommerceWalletByWalletid) | **Delete** /v1/commerce/wallet/{walletid} | Delete a wallet, keeping a recoverable copy
[**DeleteCommerceWatchlistByWatchlistid**](CommerceAPI.md#DeleteCommerceWatchlistByWatchlistid) | **Delete** /v1/commerce/watchlist/{watchlistid} | Delete a watchlist, keeping a recoverable copy
[**DeleteCommerceWebhookByWebhookid**](CommerceAPI.md#DeleteCommerceWebhookByWebhookid) | **Delete** /v1/commerce/webhook/{webhookid} | Delete a webhook, keeping a recoverable copy
[**DiscardCart**](CommerceAPI.md#DiscardCart) | **Post** /v1/commerce/cart/{id}/discard | Discard a cart the shopper abandoned
[**GetCart**](CommerceAPI.md#GetCart) | **Get** /v1/commerce/cart/{id} | Read one cart with its lines and totals
[**GetCommerceAdminCatalog**](CommerceAPI.md#GetCommerceAdminCatalog) | **Get** /v1/commerce/admin/catalog | The catalog projection with cost and margin included
[**GetCommerceCatalog**](CommerceAPI.md#GetCommerceCatalog) | **Get** /v1/commerce/catalog | The public product catalog projection for a brand
[**GetCommerceCatalogEntries**](CommerceAPI.md#GetCommerceCatalogEntries) | **Get** /v1/commerce/catalog/entries | The raw catalog entries, including the unpublished ones
[**GetCommerceCollection**](CommerceAPI.md#GetCommerceCollection) | **Get** /v1/commerce/collection/ | List your org&#39;s collections, as a page
[**GetCommerceCollectionByCollectionid**](CommerceAPI.md#GetCommerceCollectionByCollectionid) | **Get** /v1/commerce/collection/{collectionid} | Fetch one collection
[**GetCommerceCurrencies**](CommerceAPI.md#GetCommerceCurrencies) | **Get** /v1/commerce/currencies | The reference currency list the price and settings pickers render
[**GetCommerceDeposits**](CommerceAPI.md#GetCommerceDeposits) | **Get** /v1/commerce/deposits | Read the crypto deposit watcher&#39;s runtime state, asset by asset
[**GetCommerceDisclosure**](CommerceAPI.md#GetCommerceDisclosure) | **Get** /v1/commerce/disclosure/ | List your org&#39;s disclosures, as a page
[**GetCommerceDisclosureByDisclosureid**](CommerceAPI.md#GetCommerceDisclosureByDisclosureid) | **Get** /v1/commerce/disclosure/{disclosureid} | Fetch one disclosure
[**GetCommerceDiscount**](CommerceAPI.md#GetCommerceDiscount) | **Get** /v1/commerce/discount/ | List your org&#39;s discounts, as a page
[**GetCommerceDiscountByDiscountid**](CommerceAPI.md#GetCommerceDiscountByDiscountid) | **Get** /v1/commerce/discount/{discountid} | Fetch one discount
[**GetCommerceHealth**](CommerceAPI.md#GetCommerceHealth) | **Get** /v1/commerce/health | Answers ok whenever the commerce subsystem is mounted.
[**GetCommerceMovie**](CommerceAPI.md#GetCommerceMovie) | **Get** /v1/commerce/movie/ | List your org&#39;s movies, as a page
[**GetCommerceMovieByMovieid**](CommerceAPI.md#GetCommerceMovieByMovieid) | **Get** /v1/commerce/movie/{movieid} | Fetch one movie
[**GetCommerceNote**](CommerceAPI.md#GetCommerceNote) | **Get** /v1/commerce/note/ | List your org&#39;s notes, as a page
[**GetCommerceNoteByNoteid**](CommerceAPI.md#GetCommerceNoteByNoteid) | **Get** /v1/commerce/note/{noteid} | Fetch one note
[**GetCommerceOrg**](CommerceAPI.md#GetCommerceOrg) | **Get** /v1/commerce/org | The public org configuration a checkout page boots from
[**GetCommercePlansEntries**](CommerceAPI.md#GetCommercePlansEntries) | **Get** /v1/commerce/plans/entries | The raw plan authority rows
[**GetCommerceProduct**](CommerceAPI.md#GetCommerceProduct) | **Get** /v1/commerce/product/ | List your org&#39;s products, as a page
[**GetCommerceProductByProductid**](CommerceAPI.md#GetCommerceProductByProductid) | **Get** /v1/commerce/product/{productid} | Fetch one product
[**GetCommerceRatesEntries**](CommerceAPI.md#GetCommerceRatesEntries) | **Get** /v1/commerce/rates/entries | List what one unit of each metered thing costs
[**GetCommerceReturn**](CommerceAPI.md#GetCommerceReturn) | **Get** /v1/commerce/return/ | List your org&#39;s returns, as a page
[**GetCommerceReturnByReturnid**](CommerceAPI.md#GetCommerceReturnByReturnid) | **Get** /v1/commerce/return/{returnid} | Fetch one return
[**GetCommerceSaleschannel**](CommerceAPI.md#GetCommerceSaleschannel) | **Get** /v1/commerce/saleschannel/ | List your org&#39;s sales channels, as a page
[**GetCommerceSaleschannelBySaleschannelid**](CommerceAPI.md#GetCommerceSaleschannelBySaleschannelid) | **Get** /v1/commerce/saleschannel/{saleschannelid} | Fetch one sales channel
[**GetCommerceStocklocation**](CommerceAPI.md#GetCommerceStocklocation) | **Get** /v1/commerce/stocklocation/ | List your org&#39;s stock locations, as a page
[**GetCommerceStocklocationByStocklocationid**](CommerceAPI.md#GetCommerceStocklocationByStocklocationid) | **Get** /v1/commerce/stocklocation/{stocklocationid} | Fetch one stock location
[**GetCommerceStore**](CommerceAPI.md#GetCommerceStore) | **Get** /v1/commerce/store/ | List your org&#39;s storefronts as a page
[**GetCommerceStoreAccess**](CommerceAPI.md#GetCommerceStoreAccess) | **Get** /v1/commerce/store/access | Whether a store is entitled to trade, and why
[**GetCommerceStoreByStoreid**](CommerceAPI.md#GetCommerceStoreByStoreid) | **Get** /v1/commerce/store/{storeid} | Fetch one storefront
[**GetCommerceStoreByStoreidBundleByKey**](CommerceAPI.md#GetCommerceStoreByStoreidBundleByKey) | **Get** /v1/commerce/store/{storeid}/bundle/{key} | Fetch a bundle as this storefront sells it
[**GetCommerceStoreByStoreidListing**](CommerceAPI.md#GetCommerceStoreByStoreidListing) | **Get** /v1/commerce/store/{storeid}/listing | The storefront&#39;s whole listing override map
[**GetCommerceStoreByStoreidListingByKey**](CommerceAPI.md#GetCommerceStoreByStoreidListingByKey) | **Get** /v1/commerce/store/{storeid}/listing/{key} | Fetch one listing override, by item id or by its slug or SKU
[**GetCommerceStoreByStoreidProductByKey**](CommerceAPI.md#GetCommerceStoreByStoreidProductByKey) | **Get** /v1/commerce/store/{storeid}/product/{key} | Fetch a product as this storefront sells it
[**GetCommerceStoreByStoreidVariantByKey**](CommerceAPI.md#GetCommerceStoreByStoreidVariantByKey) | **Get** /v1/commerce/store/{storeid}/variant/{key} | Fetch a variant as this storefront sells it
[**GetCommerceStoreCurrent**](CommerceAPI.md#GetCommerceStoreCurrent) | **Get** /v1/commerce/store/current | Resolve your org&#39;s active storefront without naming an id
[**GetCommerceSubmission**](CommerceAPI.md#GetCommerceSubmission) | **Get** /v1/commerce/submission/ | List your org&#39;s submissions, as a page
[**GetCommerceSubmissionBySubmissionid**](CommerceAPI.md#GetCommerceSubmissionBySubmissionid) | **Get** /v1/commerce/submission/{submissionid} | Fetch one submission
[**GetCommerceSubscriber**](CommerceAPI.md#GetCommerceSubscriber) | **Get** /v1/commerce/subscriber/ | List your org&#39;s subscribers, as a page
[**GetCommerceSubscriberBySubscriberid**](CommerceAPI.md#GetCommerceSubscriberBySubscriberid) | **Get** /v1/commerce/subscriber/{subscriberid} | Fetch one subscriber
[**GetCommerceTokentransaction**](CommerceAPI.md#GetCommerceTokentransaction) | **Get** /v1/commerce/tokentransaction/ | List your org&#39;s token transactions, as a page
[**GetCommerceTokentransactionByTokentransactionid**](CommerceAPI.md#GetCommerceTokentransactionByTokentransactionid) | **Get** /v1/commerce/tokentransaction/{tokentransactionid} | Fetch one token transaction
[**GetCommerceTransfer**](CommerceAPI.md#GetCommerceTransfer) | **Get** /v1/commerce/transfer/ | List your org&#39;s transfers, as a page
[**GetCommerceTransferByTransferid**](CommerceAPI.md#GetCommerceTransferByTransferid) | **Get** /v1/commerce/transfer/{transferid} | Fetch one transfer
[**GetCommerceVariant**](CommerceAPI.md#GetCommerceVariant) | **Get** /v1/commerce/variant/ | List your org&#39;s variants, as a page
[**GetCommerceVariantByVariantid**](CommerceAPI.md#GetCommerceVariantByVariantid) | **Get** /v1/commerce/variant/{variantid} | Fetch one variant
[**GetCommerceWallet**](CommerceAPI.md#GetCommerceWallet) | **Get** /v1/commerce/wallet/ | List your org&#39;s wallets, as a page
[**GetCommerceWalletByWalletid**](CommerceAPI.md#GetCommerceWalletByWalletid) | **Get** /v1/commerce/wallet/{walletid} | Fetch one wallet
[**GetCommerceWatchlist**](CommerceAPI.md#GetCommerceWatchlist) | **Get** /v1/commerce/watchlist/ | List your org&#39;s watchlists, as a page
[**GetCommerceWatchlistByWatchlistid**](CommerceAPI.md#GetCommerceWatchlistByWatchlistid) | **Get** /v1/commerce/watchlist/{watchlistid} | Fetch one watchlist
[**GetCommerceWebhook**](CommerceAPI.md#GetCommerceWebhook) | **Get** /v1/commerce/webhook/ | List your org&#39;s webhooks, as a page
[**GetCommerceWebhookByWebhookid**](CommerceAPI.md#GetCommerceWebhookByWebhookid) | **Get** /v1/commerce/webhook/{webhookid} | Fetch one webhook
[**GetPayment**](CommerceAPI.md#GetPayment) | **Get** /v1/commerce/payments/{id} | Read one settled payment by its id
[**OpenCart**](CommerceAPI.md#OpenCart) | **Post** /v1/commerce/cart | Open a cart for a shopper to fill
[**PatchCommerceCollectionByCollectionid**](CommerceAPI.md#PatchCommerceCollectionByCollectionid) | **Patch** /v1/commerce/collection/{collectionid} | Change part of a collection
[**PatchCommerceDisclosureByDisclosureid**](CommerceAPI.md#PatchCommerceDisclosureByDisclosureid) | **Patch** /v1/commerce/disclosure/{disclosureid} | Change part of a disclosure
[**PatchCommerceDiscountByDiscountid**](CommerceAPI.md#PatchCommerceDiscountByDiscountid) | **Patch** /v1/commerce/discount/{discountid} | Change part of a discount
[**PatchCommerceMovieByMovieid**](CommerceAPI.md#PatchCommerceMovieByMovieid) | **Patch** /v1/commerce/movie/{movieid} | Change part of a movie
[**PatchCommerceNoteByNoteid**](CommerceAPI.md#PatchCommerceNoteByNoteid) | **Patch** /v1/commerce/note/{noteid} | Change part of a note
[**PatchCommerceProductByProductid**](CommerceAPI.md#PatchCommerceProductByProductid) | **Patch** /v1/commerce/product/{productid} | Change part of a product
[**PatchCommerceReturnByReturnid**](CommerceAPI.md#PatchCommerceReturnByReturnid) | **Patch** /v1/commerce/return/{returnid} | Change part of a return
[**PatchCommerceSaleschannelBySaleschannelid**](CommerceAPI.md#PatchCommerceSaleschannelBySaleschannelid) | **Patch** /v1/commerce/saleschannel/{saleschannelid} | Change part of a sales channel
[**PatchCommerceStocklocationByStocklocationid**](CommerceAPI.md#PatchCommerceStocklocationByStocklocationid) | **Patch** /v1/commerce/stocklocation/{stocklocationid} | Change part of a stock location
[**PatchCommerceStoreByStoreid**](CommerceAPI.md#PatchCommerceStoreByStoreid) | **Patch** /v1/commerce/store/{storeid} | Change part of a storefront
[**PatchCommerceStoreByStoreidListingByKey**](CommerceAPI.md#PatchCommerceStoreByStoreidListingByKey) | **Patch** /v1/commerce/store/{storeid}/listing/{key} | Confirm a listing override exists and re-save the store
[**PatchCommerceSubmissionBySubmissionid**](CommerceAPI.md#PatchCommerceSubmissionBySubmissionid) | **Patch** /v1/commerce/submission/{submissionid} | Change part of a submission
[**PatchCommerceSubscriberBySubscriberid**](CommerceAPI.md#PatchCommerceSubscriberBySubscriberid) | **Patch** /v1/commerce/subscriber/{subscriberid} | Change part of a subscriber
[**PatchCommerceTokentransactionByTokentransactionid**](CommerceAPI.md#PatchCommerceTokentransactionByTokentransactionid) | **Patch** /v1/commerce/tokentransaction/{tokentransactionid} | Change part of a token transaction
[**PatchCommerceTransferByTransferid**](CommerceAPI.md#PatchCommerceTransferByTransferid) | **Patch** /v1/commerce/transfer/{transferid} | Change part of a transfer
[**PatchCommerceVariantByVariantid**](CommerceAPI.md#PatchCommerceVariantByVariantid) | **Patch** /v1/commerce/variant/{variantid} | Change part of a variant
[**PatchCommerceWalletByWalletid**](CommerceAPI.md#PatchCommerceWalletByWalletid) | **Patch** /v1/commerce/wallet/{walletid} | Change part of a wallet
[**PatchCommerceWatchlistByWatchlistid**](CommerceAPI.md#PatchCommerceWatchlistByWatchlistid) | **Patch** /v1/commerce/watchlist/{watchlistid} | Change part of a watchlist
[**PatchCommerceWebhookByWebhookid**](CommerceAPI.md#PatchCommerceWebhookByWebhookid) | **Patch** /v1/commerce/webhook/{webhookid} | Change part of a webhook
[**PostCommerceCatalogEntries**](CommerceAPI.md#PostCommerceCatalogEntries) | **Post** /v1/commerce/catalog/entries | Add a catalog entry
[**PostCommerceCatalogModels**](CommerceAPI.md#PostCommerceCatalogModels) | **Post** /v1/commerce/catalog/models | Land a syncer&#39;s view of the model catalog: upstream costs and machine facts
[**PostCommerceCatalogModelsRefresh**](CommerceAPI.md#PostCommerceCatalogModelsRefresh) | **Post** /v1/commerce/catalog/models/refresh | Refresh the model catalog by reading the upstream provider
[**PostCommerceCatalogSeed**](CommerceAPI.md#PostCommerceCatalogSeed) | **Post** /v1/commerce/catalog/seed | Seed the embedded catalog, without disturbing edits already made
[**PostCommerceCollection**](CommerceAPI.md#PostCommerceCollection) | **Post** /v1/commerce/collection/ | Create a collection
[**PostCommerceCollectionByCollectionid**](CommerceAPI.md#PostCommerceCollectionByCollectionid) | **Post** /v1/commerce/collection/{collectionid} | Method-override tunnel for a collection — for clients that cannot send PUT, PATCH or DELETE
[**PostCommerceDisclosure**](CommerceAPI.md#PostCommerceDisclosure) | **Post** /v1/commerce/disclosure/ | Create a disclosure
[**PostCommerceDisclosureByDisclosureid**](CommerceAPI.md#PostCommerceDisclosureByDisclosureid) | **Post** /v1/commerce/disclosure/{disclosureid} | Method-override tunnel for a disclosure — for clients that cannot send PUT, PATCH or DELETE
[**PostCommerceDiscount**](CommerceAPI.md#PostCommerceDiscount) | **Post** /v1/commerce/discount/ | Create a discount
[**PostCommerceDiscountByDiscountid**](CommerceAPI.md#PostCommerceDiscountByDiscountid) | **Post** /v1/commerce/discount/{discountid} | Method-override tunnel for a discount — for clients that cannot send PUT, PATCH or DELETE
[**PostCommerceMovie**](CommerceAPI.md#PostCommerceMovie) | **Post** /v1/commerce/movie/ | Create a movie
[**PostCommerceMovieByMovieid**](CommerceAPI.md#PostCommerceMovieByMovieid) | **Post** /v1/commerce/movie/{movieid} | Method-override tunnel for a movie — for clients that cannot send PUT, PATCH or DELETE
[**PostCommerceNote**](CommerceAPI.md#PostCommerceNote) | **Post** /v1/commerce/note/ | Create a note
[**PostCommerceNoteByNoteid**](CommerceAPI.md#PostCommerceNoteByNoteid) | **Post** /v1/commerce/note/{noteid} | Method-override tunnel for a note — for clients that cannot send PUT, PATCH or DELETE
[**PostCommercePlansEntries**](CommerceAPI.md#PostCommercePlansEntries) | **Post** /v1/commerce/plans/entries | Add a subscription plan
[**PostCommercePlansSeed**](CommerceAPI.md#PostCommercePlansSeed) | **Post** /v1/commerce/plans/seed | Seed the embedded plan catalog, without overwriting administrative edits
[**PostCommerceProduct**](CommerceAPI.md#PostCommerceProduct) | **Post** /v1/commerce/product/ | Create a product
[**PostCommerceProductByProductid**](CommerceAPI.md#PostCommerceProductByProductid) | **Post** /v1/commerce/product/{productid} | Method-override tunnel for a product — for clients that cannot send PUT, PATCH or DELETE
[**PostCommerceRatesEntries**](CommerceAPI.md#PostCommerceRatesEntries) | **Post** /v1/commerce/rates/entries | Add a rate
[**PostCommerceRatesImport**](CommerceAPI.md#PostCommerceRatesImport) | **Post** /v1/commerce/rates/import | Load the published price document, reconciling rather than replacing
[**PostCommerceReturn**](CommerceAPI.md#PostCommerceReturn) | **Post** /v1/commerce/return/ | Create a return
[**PostCommerceReturnByReturnid**](CommerceAPI.md#PostCommerceReturnByReturnid) | **Post** /v1/commerce/return/{returnid} | Method-override tunnel for a return — for clients that cannot send PUT, PATCH or DELETE
[**PostCommerceSaleschannel**](CommerceAPI.md#PostCommerceSaleschannel) | **Post** /v1/commerce/saleschannel/ | Create a sales channel
[**PostCommerceSaleschannelBySaleschannelid**](CommerceAPI.md#PostCommerceSaleschannelBySaleschannelid) | **Post** /v1/commerce/saleschannel/{saleschannelid} | Method-override tunnel for a sales channel — for clients that cannot send PUT, PATCH or DELETE
[**PostCommerceStocklocation**](CommerceAPI.md#PostCommerceStocklocation) | **Post** /v1/commerce/stocklocation/ | Create a stock location
[**PostCommerceStocklocationByStocklocationid**](CommerceAPI.md#PostCommerceStocklocationByStocklocationid) | **Post** /v1/commerce/stocklocation/{stocklocationid} | Method-override tunnel for a stock location — for clients that cannot send PUT, PATCH or DELETE
[**PostCommerceStore**](CommerceAPI.md#PostCommerceStore) | **Post** /v1/commerce/store/ | Create a storefront
[**PostCommerceStoreByStoreid**](CommerceAPI.md#PostCommerceStoreByStoreid) | **Post** /v1/commerce/store/{storeid} | Method-override tunnel for clients that cannot send PUT, PATCH or DELETE
[**PostCommerceStoreByStoreidAuthorize**](CommerceAPI.md#PostCommerceStoreByStoreidAuthorize) | **Post** /v1/commerce/store/{storeid}/authorize | Authorize a new order against a storefront, holding the funds without settling them
[**PostCommerceStoreByStoreidAuthorizeByOrderid**](CommerceAPI.md#PostCommerceStoreByStoreidAuthorizeByOrderid) | **Post** /v1/commerce/store/{storeid}/authorize/{orderid} | Authorize an order that already exists, holding the funds without settling them
[**PostCommerceStoreByStoreidCaptureByOrderid**](CommerceAPI.md#PostCommerceStoreByStoreidCaptureByOrderid) | **Post** /v1/commerce/store/{storeid}/capture/{orderid} | Capture a previously authorized order and settle the payment
[**PostCommerceStoreByStoreidCharge**](CommerceAPI.md#PostCommerceStoreByStoreidCharge) | **Post** /v1/commerce/store/{storeid}/charge | Authorize and capture a new order in one call
[**PostCommerceStoreByStoreidCheckoutAuthorize**](CommerceAPI.md#PostCommerceStoreByStoreidCheckoutAuthorize) | **Post** /v1/commerce/store/{storeid}/checkout/authorize | Authorize a new order against a storefront, holding the funds — the checkout spelling
[**PostCommerceStoreByStoreidCheckoutAuthorizeByOrderid**](CommerceAPI.md#PostCommerceStoreByStoreidCheckoutAuthorizeByOrderid) | **Post** /v1/commerce/store/{storeid}/checkout/authorize/{orderid} | Authorize an existing order, holding the funds — the checkout spelling
[**PostCommerceStoreByStoreidCheckoutCaptureByOrderid**](CommerceAPI.md#PostCommerceStoreByStoreidCheckoutCaptureByOrderid) | **Post** /v1/commerce/store/{storeid}/checkout/capture/{orderid} | Capture a previously authorized order and settle it — the checkout spelling
[**PostCommerceStoreByStoreidCheckoutCharge**](CommerceAPI.md#PostCommerceStoreByStoreidCheckoutCharge) | **Post** /v1/commerce/store/{storeid}/checkout/charge | Authorize and capture a new order in one call — the checkout spelling
[**PostCommerceStoreByStoreidCheckoutPaypalCancelByPaykey**](CommerceAPI.md#PostCommerceStoreByStoreidCheckoutPaypalCancelByPaykey) | **Post** /v1/commerce/store/{storeid}/checkout/paypal/cancel/{payKey} | PayPal cancel by pay key — refuses, exactly as the unprefixed address does
[**PostCommerceStoreByStoreidCheckoutPaypalConfirmByPaykey**](CommerceAPI.md#PostCommerceStoreByStoreidCheckoutPaypalConfirmByPaykey) | **Post** /v1/commerce/store/{storeid}/checkout/paypal/confirm/{payKey} | PayPal confirm by pay key — refuses, exactly as the unprefixed address does
[**PostCommerceStoreByStoreidCheckoutPaypalPay**](CommerceAPI.md#PostCommerceStoreByStoreidCheckoutPaypalPay) | **Post** /v1/commerce/store/{storeid}/checkout/paypal/pay | Start a PayPal authorization for a new order — the checkout spelling
[**PostCommerceStoreByStoreidListingByKey**](CommerceAPI.md#PostCommerceStoreByStoreidListingByKey) | **Post** /v1/commerce/store/{storeid}/listing/{key} | Add a listing override under a new key
[**PostCommerceStoreByStoreidPaypalCancelByPaykey**](CommerceAPI.md#PostCommerceStoreByStoreidPaypalCancelByPaykey) | **Post** /v1/commerce/store/{storeid}/paypal/cancel/{payKey} | PayPal cancel by pay key — refuses, because a pay key alone does not identify the order
[**PostCommerceStoreByStoreidPaypalConfirmByPaykey**](CommerceAPI.md#PostCommerceStoreByStoreidPaypalConfirmByPaykey) | **Post** /v1/commerce/store/{storeid}/paypal/confirm/{payKey} | PayPal confirm by pay key — refuses, because a pay key alone does not identify the order
[**PostCommerceStoreByStoreidPaypalPay**](CommerceAPI.md#PostCommerceStoreByStoreidPaypalPay) | **Post** /v1/commerce/store/{storeid}/paypal/pay | Start a PayPal authorization for a new order
[**PostCommerceStoreByStoreidTrial**](CommerceAPI.md#PostCommerceStoreByStoreidTrial) | **Post** /v1/commerce/store/{storeid}/trial | Start this store&#39;s no-card trial on the entry plan
[**PostCommerceStoreToken**](CommerceAPI.md#PostCommerceStoreToken) | **Post** /v1/commerce/store/token | Mint your org&#39;s least-privilege storefront read key
[**PostCommerceSubmission**](CommerceAPI.md#PostCommerceSubmission) | **Post** /v1/commerce/submission/ | Create a submission
[**PostCommerceSubmissionBySubmissionid**](CommerceAPI.md#PostCommerceSubmissionBySubmissionid) | **Post** /v1/commerce/submission/{submissionid} | Method-override tunnel for a submission — for clients that cannot send PUT, PATCH or DELETE
[**PostCommerceSubscriber**](CommerceAPI.md#PostCommerceSubscriber) | **Post** /v1/commerce/subscriber/ | Create a subscriber
[**PostCommerceSubscriberBySubscriberid**](CommerceAPI.md#PostCommerceSubscriberBySubscriberid) | **Post** /v1/commerce/subscriber/{subscriberid} | Method-override tunnel for a subscriber — for clients that cannot send PUT, PATCH or DELETE
[**PostCommerceTokentransaction**](CommerceAPI.md#PostCommerceTokentransaction) | **Post** /v1/commerce/tokentransaction/ | Create a token transaction
[**PostCommerceTokentransactionByTokentransactionid**](CommerceAPI.md#PostCommerceTokentransactionByTokentransactionid) | **Post** /v1/commerce/tokentransaction/{tokentransactionid} | Method-override tunnel for a token transaction — for clients that cannot send PUT, PATCH or DELETE
[**PostCommerceTransfer**](CommerceAPI.md#PostCommerceTransfer) | **Post** /v1/commerce/transfer/ | Create a transfer
[**PostCommerceTransferByTransferid**](CommerceAPI.md#PostCommerceTransferByTransferid) | **Post** /v1/commerce/transfer/{transferid} | Method-override tunnel for a transfer — for clients that cannot send PUT, PATCH or DELETE
[**PostCommerceVariant**](CommerceAPI.md#PostCommerceVariant) | **Post** /v1/commerce/variant/ | Create a variant
[**PostCommerceVariantByVariantid**](CommerceAPI.md#PostCommerceVariantByVariantid) | **Post** /v1/commerce/variant/{variantid} | Method-override tunnel for a variant — for clients that cannot send PUT, PATCH or DELETE
[**PostCommerceWallet**](CommerceAPI.md#PostCommerceWallet) | **Post** /v1/commerce/wallet/ | Create a wallet
[**PostCommerceWalletByWalletid**](CommerceAPI.md#PostCommerceWalletByWalletid) | **Post** /v1/commerce/wallet/{walletid} | Method-override tunnel for a wallet — for clients that cannot send PUT, PATCH or DELETE
[**PostCommerceWatchlist**](CommerceAPI.md#PostCommerceWatchlist) | **Post** /v1/commerce/watchlist/ | Create a watchlist
[**PostCommerceWatchlistByWatchlistid**](CommerceAPI.md#PostCommerceWatchlistByWatchlistid) | **Post** /v1/commerce/watchlist/{watchlistid} | Method-override tunnel for a watchlist — for clients that cannot send PUT, PATCH or DELETE
[**PostCommerceWebhook**](CommerceAPI.md#PostCommerceWebhook) | **Post** /v1/commerce/webhook/ | Create a webhook
[**PostCommerceWebhookByWebhookid**](CommerceAPI.md#PostCommerceWebhookByWebhookid) | **Post** /v1/commerce/webhook/{webhookid} | Method-override tunnel for a webhook — for clients that cannot send PUT, PATCH or DELETE
[**PostCommerceWebhooksByProvider**](CommerceAPI.md#PostCommerceWebhooksByProvider) | **Post** /v1/commerce/webhooks/{provider} | Payment-provider webhook intake for settlement and subscription lifecycle events
[**PutCommerceCollectionByCollectionid**](CommerceAPI.md#PutCommerceCollectionByCollectionid) | **Put** /v1/commerce/collection/{collectionid} | Replace a collection outright
[**PutCommerceDisclosureByDisclosureid**](CommerceAPI.md#PutCommerceDisclosureByDisclosureid) | **Put** /v1/commerce/disclosure/{disclosureid} | Replace a disclosure outright
[**PutCommerceDiscountByDiscountid**](CommerceAPI.md#PutCommerceDiscountByDiscountid) | **Put** /v1/commerce/discount/{discountid} | Replace a discount outright
[**PutCommerceMovieByMovieid**](CommerceAPI.md#PutCommerceMovieByMovieid) | **Put** /v1/commerce/movie/{movieid} | Replace a movie outright
[**PutCommerceNoteByNoteid**](CommerceAPI.md#PutCommerceNoteByNoteid) | **Put** /v1/commerce/note/{noteid} | Replace a note outright
[**PutCommercePlansEntriesBySlug**](CommerceAPI.md#PutCommercePlansEntriesBySlug) | **Put** /v1/commerce/plans/entries/{slug} | Edit a plan, leaving the fields you omit alone
[**PutCommerceProductByProductid**](CommerceAPI.md#PutCommerceProductByProductid) | **Put** /v1/commerce/product/{productid} | Replace a product outright
[**PutCommerceRatesEntriesByProductByMeter**](CommerceAPI.md#PutCommerceRatesEntriesByProductByMeter) | **Put** /v1/commerce/rates/entries/{product}/{meter} | Edit a rate, and mark it as operator-set
[**PutCommerceReturnByReturnid**](CommerceAPI.md#PutCommerceReturnByReturnid) | **Put** /v1/commerce/return/{returnid} | Replace a return outright
[**PutCommerceSaleschannelBySaleschannelid**](CommerceAPI.md#PutCommerceSaleschannelBySaleschannelid) | **Put** /v1/commerce/saleschannel/{saleschannelid} | Replace a sales channel outright
[**PutCommerceStocklocationByStocklocationid**](CommerceAPI.md#PutCommerceStocklocationByStocklocationid) | **Put** /v1/commerce/stocklocation/{stocklocationid} | Replace a stock location outright
[**PutCommerceStoreByStoreid**](CommerceAPI.md#PutCommerceStoreByStoreid) | **Put** /v1/commerce/store/{storeid} | Replace a storefront outright
[**PutCommerceStoreByStoreidListingByKey**](CommerceAPI.md#PutCommerceStoreByStoreidListingByKey) | **Put** /v1/commerce/store/{storeid}/listing/{key} | Upsert a listing override
[**PutCommerceSubmissionBySubmissionid**](CommerceAPI.md#PutCommerceSubmissionBySubmissionid) | **Put** /v1/commerce/submission/{submissionid} | Replace a submission outright
[**PutCommerceSubscriberBySubscriberid**](CommerceAPI.md#PutCommerceSubscriberBySubscriberid) | **Put** /v1/commerce/subscriber/{subscriberid} | Replace a subscriber outright
[**PutCommerceTokentransactionByTokentransactionid**](CommerceAPI.md#PutCommerceTokentransactionByTokentransactionid) | **Put** /v1/commerce/tokentransaction/{tokentransactionid} | Replace a token transaction outright
[**PutCommerceTransferByTransferid**](CommerceAPI.md#PutCommerceTransferByTransferid) | **Put** /v1/commerce/transfer/{transferid} | Replace a transfer outright
[**PutCommerceVariantByVariantid**](CommerceAPI.md#PutCommerceVariantByVariantid) | **Put** /v1/commerce/variant/{variantid} | Replace a variant outright
[**PutCommerceWalletByWalletid**](CommerceAPI.md#PutCommerceWalletByWalletid) | **Put** /v1/commerce/wallet/{walletid} | Replace a wallet outright
[**PutCommerceWatchlistByWatchlistid**](CommerceAPI.md#PutCommerceWatchlistByWatchlistid) | **Put** /v1/commerce/watchlist/{watchlistid} | Replace a watchlist outright
[**PutCommerceWebhookByWebhookid**](CommerceAPI.md#PutCommerceWebhookByWebhookid) | **Put** /v1/commerce/webhook/{webhookid} | Replace a webhook outright
[**SetCartItem**](CommerceAPI.md#SetCartItem) | **Post** /v1/commerce/cart/{id}/item | Set one item&#39;s quantity in a cart; zero removes it
[**TakePayment**](CommerceAPI.md#TakePayment) | **Post** /v1/commerce/payments | Take a card payment and credit the org&#39;s balance



## DeleteCommerceCollectionByCollectionid

> DeleteCommerceCollectionByCollectionid(ctx, collectionid).Execute()

Delete a collection, keeping a recoverable copy



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	collectionid := "collectionid_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CommerceAPI.DeleteCommerceCollectionByCollectionid(context.Background(), collectionid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommerceAPI.DeleteCommerceCollectionByCollectionid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCommerceCollectionByCollectionidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCommerceDisclosureByDisclosureid

> DeleteCommerceDisclosureByDisclosureid(ctx, disclosureid).Execute()

Delete a disclosure, keeping a recoverable copy



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	disclosureid := "disclosureid_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CommerceAPI.DeleteCommerceDisclosureByDisclosureid(context.Background(), disclosureid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommerceAPI.DeleteCommerceDisclosureByDisclosureid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**disclosureid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCommerceDisclosureByDisclosureidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCommerceDiscountByDiscountid

> DeleteCommerceDiscountByDiscountid(ctx, discountid).Execute()

Delete a discount, keeping a recoverable copy



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	discountid := "discountid_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CommerceAPI.DeleteCommerceDiscountByDiscountid(context.Background(), discountid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommerceAPI.DeleteCommerceDiscountByDiscountid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**discountid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCommerceDiscountByDiscountidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCommerceMovieByMovieid

> DeleteCommerceMovieByMovieid(ctx, movieid).Execute()

Delete a movie, keeping a recoverable copy



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	movieid := "movieid_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CommerceAPI.DeleteCommerceMovieByMovieid(context.Background(), movieid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommerceAPI.DeleteCommerceMovieByMovieid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**movieid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCommerceMovieByMovieidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCommerceNoteByNoteid

> DeleteCommerceNoteByNoteid(ctx, noteid).Execute()

Delete a note, keeping a recoverable copy



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	noteid := "noteid_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CommerceAPI.DeleteCommerceNoteByNoteid(context.Background(), noteid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommerceAPI.DeleteCommerceNoteByNoteid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**noteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCommerceNoteByNoteidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCommercePlansEntriesBySlug

> DeleteCommercePlansEntriesBySlug(ctx, slug).Execute()

Remove a plan from the authority



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	slug := "slug_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CommerceAPI.DeleteCommercePlansEntriesBySlug(context.Background(), slug).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommerceAPI.DeleteCommercePlansEntriesBySlug``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slug** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCommercePlansEntriesBySlugRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCommerceProductByProductid

> DeleteCommerceProductByProductid(ctx, productid).Execute()

Delete a product, keeping a recoverable copy



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	productid := "productid_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CommerceAPI.DeleteCommerceProductByProductid(context.Background(), productid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommerceAPI.DeleteCommerceProductByProductid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**productid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCommerceProductByProductidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCommerceRatesEntriesByProductByMeter

> DeleteCommerceRatesEntriesByProductByMeter(ctx, product, meter).Execute()

Remove a rate outright



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	product := "product_example" // string | 
	meter := "meter_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CommerceAPI.DeleteCommerceRatesEntriesByProductByMeter(context.Background(), product, meter).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommerceAPI.DeleteCommerceRatesEntriesByProductByMeter``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**product** | **string** |  | 
**meter** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCommerceRatesEntriesByProductByMeterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCommerceReturnByReturnid

> DeleteCommerceReturnByReturnid(ctx, returnid).Execute()

Delete a return, keeping a recoverable copy



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	returnid := "returnid_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CommerceAPI.DeleteCommerceReturnByReturnid(context.Background(), returnid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommerceAPI.DeleteCommerceReturnByReturnid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**returnid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCommerceReturnByReturnidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCommerceSaleschannelBySaleschannelid

> DeleteCommerceSaleschannelBySaleschannelid(ctx, saleschannelid).Execute()

Delete a sales channel, keeping a recoverable copy



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	saleschannelid := "saleschannelid_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CommerceAPI.DeleteCommerceSaleschannelBySaleschannelid(context.Background(), saleschannelid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommerceAPI.DeleteCommerceSaleschannelBySaleschannelid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**saleschannelid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCommerceSaleschannelBySaleschannelidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCommerceStocklocationByStocklocationid

> DeleteCommerceStocklocationByStocklocationid(ctx, stocklocationid).Execute()

Delete a stock location, keeping a recoverable copy



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	stocklocationid := "stocklocationid_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CommerceAPI.DeleteCommerceStocklocationByStocklocationid(context.Background(), stocklocationid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommerceAPI.DeleteCommerceStocklocationByStocklocationid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stocklocationid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCommerceStocklocationByStocklocationidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCommerceStoreByStoreid

> DeleteCommerceStoreByStoreid(ctx, storeid).Execute()

Delete a storefront, keeping a recoverable copy



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	storeid := "storeid_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CommerceAPI.DeleteCommerceStoreByStoreid(context.Background(), storeid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommerceAPI.DeleteCommerceStoreByStoreid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**storeid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCommerceStoreByStoreidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCommerceStoreByStoreidListingByKey

> DeleteCommerceStoreByStoreidListingByKey(ctx, storeid, key).Execute()

Remove a listing override



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	storeid := "storeid_example" // string | 
	key := "key_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CommerceAPI.DeleteCommerceStoreByStoreidListingByKey(context.Background(), storeid, key).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommerceAPI.DeleteCommerceStoreByStoreidListingByKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**storeid** | **string** |  | 
**key** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCommerceStoreByStoreidListingByKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCommerceSubmissionBySubmissionid

> DeleteCommerceSubmissionBySubmissionid(ctx, submissionid).Execute()

Delete a submission, keeping a recoverable copy



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	submissionid := "submissionid_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CommerceAPI.DeleteCommerceSubmissionBySubmissionid(context.Background(), submissionid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommerceAPI.DeleteCommerceSubmissionBySubmissionid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**submissionid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCommerceSubmissionBySubmissionidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCommerceSubscriberBySubscriberid

> DeleteCommerceSubscriberBySubscriberid(ctx, subscriberid).Execute()

Delete a subscriber, keeping a recoverable copy



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	subscriberid := "subscriberid_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CommerceAPI.DeleteCommerceSubscriberBySubscriberid(context.Background(), subscriberid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommerceAPI.DeleteCommerceSubscriberBySubscriberid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriberid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCommerceSubscriberBySubscriberidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCommerceTokentransactionByTokentransactionid

> DeleteCommerceTokentransactionByTokentransactionid(ctx, tokentransactionid).Execute()

Delete a token transaction, keeping a recoverable copy



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	tokentransactionid := "tokentransactionid_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CommerceAPI.DeleteCommerceTokentransactionByTokentransactionid(context.Background(), tokentransactionid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommerceAPI.DeleteCommerceTokentransactionByTokentransactionid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tokentransactionid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCommerceTokentransactionByTokentransactionidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCommerceTransferByTransferid

> DeleteCommerceTransferByTransferid(ctx, transferid).Execute()

Delete a transfer, keeping a recoverable copy



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	transferid := "transferid_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CommerceAPI.DeleteCommerceTransferByTransferid(context.Background(), transferid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommerceAPI.DeleteCommerceTransferByTransferid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**transferid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCommerceTransferByTransferidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCommerceVariantByVariantid

> DeleteCommerceVariantByVariantid(ctx, variantid).Execute()

Delete a variant, keeping a recoverable copy



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	variantid := "variantid_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CommerceAPI.DeleteCommerceVariantByVariantid(context.Background(), variantid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommerceAPI.DeleteCommerceVariantByVariantid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**variantid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCommerceVariantByVariantidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCommerceWalletByWalletid

> DeleteCommerceWalletByWalletid(ctx, walletid).Execute()

Delete a wallet, keeping a recoverable copy



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	walletid := "walletid_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CommerceAPI.DeleteCommerceWalletByWalletid(context.Background(), walletid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommerceAPI.DeleteCommerceWalletByWalletid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**walletid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCommerceWalletByWalletidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCommerceWatchlistByWatchlistid

> DeleteCommerceWatchlistByWatchlistid(ctx, watchlistid).Execute()

Delete a watchlist, keeping a recoverable copy



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	watchlistid := "watchlistid_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CommerceAPI.DeleteCommerceWatchlistByWatchlistid(context.Background(), watchlistid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommerceAPI.DeleteCommerceWatchlistByWatchlistid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**watchlistid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCommerceWatchlistByWatchlistidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCommerceWebhookByWebhookid

> DeleteCommerceWebhookByWebhookid(ctx, webhookid).Execute()

Delete a webhook, keeping a recoverable copy



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	webhookid := "webhookid_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CommerceAPI.DeleteCommerceWebhookByWebhookid(context.Background(), webhookid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommerceAPI.DeleteCommerceWebhookByWebhookid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**webhookid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCommerceWebhookByWebhookidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DiscardCart

> Cart DiscardCart(ctx, id).Execute()

Discard a cart the shopper abandoned



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	id := "id_example" // string | ID is the cart's id, as the open call answered it.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CommerceAPI.DiscardCart(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommerceAPI.DiscardCart``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DiscardCart`: Cart
	fmt.Fprintf(os.Stdout, "Response from `CommerceAPI.DiscardCart`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the cart&#39;s id, as the open call answered it. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDiscardCartRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Cart**](Cart.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCart

> Cart GetCart(ctx, id).Execute()

Read one cart with its lines and totals



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	id := "id_example" // string | ID is the cart's id, as the open call answered it.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CommerceAPI.GetCart(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommerceAPI.GetCart``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCart`: Cart
	fmt.Fprintf(os.Stdout, "Response from `CommerceAPI.GetCart`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the cart&#39;s id, as the open call answered it. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCartRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Cart**](Cart.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCommerceAdminCatalog

> GetCommerceAdminCatalog(ctx).Execute()

The catalog projection with cost and margin included



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CommerceAPI.GetCommerceAdminCatalog(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommerceAPI.GetCommerceAdminCatalog``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetCommerceAdminCatalogRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCommerceCatalog

> GetCommerceCatalog(ctx).Execute()

The public product catalog projection for a brand



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CommerceAPI.GetCommerceCatalog(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommerceAPI.GetCommerceCatalog``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetCommerceCatalogRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCommerceCatalogEntries

> GetCommerceCatalogEntries(ctx).Execute()

The raw catalog entries, including the unpublished ones



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CommerceAPI.GetCommerceCatalogEntries(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommerceAPI.GetCommerceCatalogEntries``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetCommerceCatalogEntriesRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCommerceCollection

> GetCommerceCollection(ctx).Execute()

List your org's collections, as a page



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CommerceAPI.GetCommerceCollection(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommerceAPI.GetCommerceCollection``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetCommerceCollectionRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCommerceCollectionByCollectionid

> GetCommerceCollectionByCollectionid(ctx, collectionid).Execute()

Fetch one collection



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	collectionid := "collectionid_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CommerceAPI.GetCommerceCollectionByCollectionid(context.Background(), collectionid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommerceAPI.GetCommerceCollectionByCollectionid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCommerceCollectionByCollectionidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCommerceCurrencies

> GetCommerceCurrencies(ctx).Execute()

The reference currency list the price and settings pickers render



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CommerceAPI.GetCommerceCurrencies(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommerceAPI.GetCommerceCurrencies``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetCommerceCurrenciesRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCommerceDeposits

> GetCommerceDeposits(ctx).Execute()

Read the crypto deposit watcher's runtime state, asset by asset



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CommerceAPI.GetCommerceDeposits(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommerceAPI.GetCommerceDeposits``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetCommerceDepositsRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCommerceDisclosure

> GetCommerceDisclosure(ctx).Execute()

List your org's disclosures, as a page



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CommerceAPI.GetCommerceDisclosure(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommerceAPI.GetCommerceDisclosure``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetCommerceDisclosureRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCommerceDisclosureByDisclosureid

> GetCommerceDisclosureByDisclosureid(ctx, disclosureid).Execute()

Fetch one disclosure



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	disclosureid := "disclosureid_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CommerceAPI.GetCommerceDisclosureByDisclosureid(context.Background(), disclosureid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommerceAPI.GetCommerceDisclosureByDisclosureid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**disclosureid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCommerceDisclosureByDisclosureidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCommerceDiscount

> GetCommerceDiscount(ctx).Execute()

List your org's discounts, as a page



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CommerceAPI.GetCommerceDiscount(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommerceAPI.GetCommerceDiscount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetCommerceDiscountRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCommerceDiscountByDiscountid

> GetCommerceDiscountByDiscountid(ctx, discountid).Execute()

Fetch one discount



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	discountid := "discountid_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CommerceAPI.GetCommerceDiscountByDiscountid(context.Background(), discountid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommerceAPI.GetCommerceDiscountByDiscountid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**discountid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCommerceDiscountByDiscountidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCommerceHealth

> Liveness GetCommerceHealth(ctx).Execute()

Answers ok whenever the commerce subsystem is mounted.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CommerceAPI.GetCommerceHealth(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommerceAPI.GetCommerceHealth``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCommerceHealth`: Liveness
	fmt.Fprintf(os.Stdout, "Response from `CommerceAPI.GetCommerceHealth`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetCommerceHealthRequest struct via the builder pattern


### Return type

[**Liveness**](Liveness.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCommerceMovie

> GetCommerceMovie(ctx).Execute()

List your org's movies, as a page



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CommerceAPI.GetCommerceMovie(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommerceAPI.GetCommerceMovie``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetCommerceMovieRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCommerceMovieByMovieid

> GetCommerceMovieByMovieid(ctx, movieid).Execute()

Fetch one movie



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	movieid := "movieid_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CommerceAPI.GetCommerceMovieByMovieid(context.Background(), movieid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommerceAPI.GetCommerceMovieByMovieid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**movieid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCommerceMovieByMovieidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCommerceNote

> GetCommerceNote(ctx).Execute()

List your org's notes, as a page



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CommerceAPI.GetCommerceNote(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommerceAPI.GetCommerceNote``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetCommerceNoteRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCommerceNoteByNoteid

> GetCommerceNoteByNoteid(ctx, noteid).Execute()

Fetch one note



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	noteid := "noteid_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CommerceAPI.GetCommerceNoteByNoteid(context.Background(), noteid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommerceAPI.GetCommerceNoteByNoteid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**noteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCommerceNoteByNoteidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCommerceOrg

> GetCommerceOrg(ctx).Execute()

The public org configuration a checkout page boots from



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CommerceAPI.GetCommerceOrg(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommerceAPI.GetCommerceOrg``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetCommerceOrgRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCommercePlansEntries

> GetCommercePlansEntries(ctx).Execute()

The raw plan authority rows



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CommerceAPI.GetCommercePlansEntries(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommerceAPI.GetCommercePlansEntries``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetCommercePlansEntriesRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCommerceProduct

> GetCommerceProduct(ctx).Execute()

List your org's products, as a page



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CommerceAPI.GetCommerceProduct(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommerceAPI.GetCommerceProduct``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetCommerceProductRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCommerceProductByProductid

> GetCommerceProductByProductid(ctx, productid).Execute()

Fetch one product



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	productid := "productid_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CommerceAPI.GetCommerceProductByProductid(context.Background(), productid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommerceAPI.GetCommerceProductByProductid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**productid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCommerceProductByProductidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCommerceRatesEntries

> GetCommerceRatesEntries(ctx).Execute()

List what one unit of each metered thing costs



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CommerceAPI.GetCommerceRatesEntries(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommerceAPI.GetCommerceRatesEntries``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetCommerceRatesEntriesRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCommerceReturn

> GetCommerceReturn(ctx).Execute()

List your org's returns, as a page



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CommerceAPI.GetCommerceReturn(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommerceAPI.GetCommerceReturn``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetCommerceReturnRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCommerceReturnByReturnid

> GetCommerceReturnByReturnid(ctx, returnid).Execute()

Fetch one return



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	returnid := "returnid_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CommerceAPI.GetCommerceReturnByReturnid(context.Background(), returnid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommerceAPI.GetCommerceReturnByReturnid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**returnid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCommerceReturnByReturnidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCommerceSaleschannel

> GetCommerceSaleschannel(ctx).Execute()

List your org's sales channels, as a page



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CommerceAPI.GetCommerceSaleschannel(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommerceAPI.GetCommerceSaleschannel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetCommerceSaleschannelRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCommerceSaleschannelBySaleschannelid

> GetCommerceSaleschannelBySaleschannelid(ctx, saleschannelid).Execute()

Fetch one sales channel



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	saleschannelid := "saleschannelid_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CommerceAPI.GetCommerceSaleschannelBySaleschannelid(context.Background(), saleschannelid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommerceAPI.GetCommerceSaleschannelBySaleschannelid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**saleschannelid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCommerceSaleschannelBySaleschannelidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCommerceStocklocation

> GetCommerceStocklocation(ctx).Execute()

List your org's stock locations, as a page



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CommerceAPI.GetCommerceStocklocation(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommerceAPI.GetCommerceStocklocation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetCommerceStocklocationRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCommerceStocklocationByStocklocationid

> GetCommerceStocklocationByStocklocationid(ctx, stocklocationid).Execute()

Fetch one stock location



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	stocklocationid := "stocklocationid_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CommerceAPI.GetCommerceStocklocationByStocklocationid(context.Background(), stocklocationid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommerceAPI.GetCommerceStocklocationByStocklocationid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stocklocationid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCommerceStocklocationByStocklocationidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCommerceStore

> GetCommerceStore(ctx).Execute()

List your org's storefronts as a page



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CommerceAPI.GetCommerceStore(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommerceAPI.GetCommerceStore``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetCommerceStoreRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCommerceStoreAccess

> GetCommerceStoreAccess(ctx).Execute()

Whether a store is entitled to trade, and why



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CommerceAPI.GetCommerceStoreAccess(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommerceAPI.GetCommerceStoreAccess``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetCommerceStoreAccessRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCommerceStoreByStoreid

> GetCommerceStoreByStoreid(ctx, storeid).Execute()

Fetch one storefront



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	storeid := "storeid_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CommerceAPI.GetCommerceStoreByStoreid(context.Background(), storeid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommerceAPI.GetCommerceStoreByStoreid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**storeid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCommerceStoreByStoreidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCommerceStoreByStoreidBundleByKey

> GetCommerceStoreByStoreidBundleByKey(ctx, storeid, key).Execute()

Fetch a bundle as this storefront sells it



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	storeid := "storeid_example" // string | 
	key := "key_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CommerceAPI.GetCommerceStoreByStoreidBundleByKey(context.Background(), storeid, key).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommerceAPI.GetCommerceStoreByStoreidBundleByKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**storeid** | **string** |  | 
**key** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCommerceStoreByStoreidBundleByKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCommerceStoreByStoreidListing

> GetCommerceStoreByStoreidListing(ctx, storeid).Execute()

The storefront's whole listing override map



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	storeid := "storeid_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CommerceAPI.GetCommerceStoreByStoreidListing(context.Background(), storeid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommerceAPI.GetCommerceStoreByStoreidListing``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**storeid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCommerceStoreByStoreidListingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCommerceStoreByStoreidListingByKey

> GetCommerceStoreByStoreidListingByKey(ctx, storeid, key).Execute()

Fetch one listing override, by item id or by its slug or SKU



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	storeid := "storeid_example" // string | 
	key := "key_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CommerceAPI.GetCommerceStoreByStoreidListingByKey(context.Background(), storeid, key).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommerceAPI.GetCommerceStoreByStoreidListingByKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**storeid** | **string** |  | 
**key** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCommerceStoreByStoreidListingByKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCommerceStoreByStoreidProductByKey

> GetCommerceStoreByStoreidProductByKey(ctx, storeid, key).Execute()

Fetch a product as this storefront sells it



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	storeid := "storeid_example" // string | 
	key := "key_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CommerceAPI.GetCommerceStoreByStoreidProductByKey(context.Background(), storeid, key).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommerceAPI.GetCommerceStoreByStoreidProductByKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**storeid** | **string** |  | 
**key** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCommerceStoreByStoreidProductByKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCommerceStoreByStoreidVariantByKey

> GetCommerceStoreByStoreidVariantByKey(ctx, storeid, key).Execute()

Fetch a variant as this storefront sells it



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	storeid := "storeid_example" // string | 
	key := "key_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CommerceAPI.GetCommerceStoreByStoreidVariantByKey(context.Background(), storeid, key).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommerceAPI.GetCommerceStoreByStoreidVariantByKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**storeid** | **string** |  | 
**key** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCommerceStoreByStoreidVariantByKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCommerceStoreCurrent

> GetCommerceStoreCurrent(ctx).Execute()

Resolve your org's active storefront without naming an id



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CommerceAPI.GetCommerceStoreCurrent(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommerceAPI.GetCommerceStoreCurrent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetCommerceStoreCurrentRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCommerceSubmission

> GetCommerceSubmission(ctx).Execute()

List your org's submissions, as a page



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CommerceAPI.GetCommerceSubmission(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommerceAPI.GetCommerceSubmission``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetCommerceSubmissionRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCommerceSubmissionBySubmissionid

> GetCommerceSubmissionBySubmissionid(ctx, submissionid).Execute()

Fetch one submission



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	submissionid := "submissionid_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CommerceAPI.GetCommerceSubmissionBySubmissionid(context.Background(), submissionid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommerceAPI.GetCommerceSubmissionBySubmissionid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**submissionid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCommerceSubmissionBySubmissionidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCommerceSubscriber

> GetCommerceSubscriber(ctx).Execute()

List your org's subscribers, as a page



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CommerceAPI.GetCommerceSubscriber(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommerceAPI.GetCommerceSubscriber``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetCommerceSubscriberRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCommerceSubscriberBySubscriberid

> GetCommerceSubscriberBySubscriberid(ctx, subscriberid).Execute()

Fetch one subscriber



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	subscriberid := "subscriberid_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CommerceAPI.GetCommerceSubscriberBySubscriberid(context.Background(), subscriberid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommerceAPI.GetCommerceSubscriberBySubscriberid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriberid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCommerceSubscriberBySubscriberidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCommerceTokentransaction

> GetCommerceTokentransaction(ctx).Execute()

List your org's token transactions, as a page



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CommerceAPI.GetCommerceTokentransaction(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommerceAPI.GetCommerceTokentransaction``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetCommerceTokentransactionRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCommerceTokentransactionByTokentransactionid

> GetCommerceTokentransactionByTokentransactionid(ctx, tokentransactionid).Execute()

Fetch one token transaction



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	tokentransactionid := "tokentransactionid_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CommerceAPI.GetCommerceTokentransactionByTokentransactionid(context.Background(), tokentransactionid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommerceAPI.GetCommerceTokentransactionByTokentransactionid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tokentransactionid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCommerceTokentransactionByTokentransactionidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCommerceTransfer

> GetCommerceTransfer(ctx).Execute()

List your org's transfers, as a page



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CommerceAPI.GetCommerceTransfer(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommerceAPI.GetCommerceTransfer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetCommerceTransferRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCommerceTransferByTransferid

> GetCommerceTransferByTransferid(ctx, transferid).Execute()

Fetch one transfer



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	transferid := "transferid_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CommerceAPI.GetCommerceTransferByTransferid(context.Background(), transferid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommerceAPI.GetCommerceTransferByTransferid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**transferid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCommerceTransferByTransferidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCommerceVariant

> GetCommerceVariant(ctx).Execute()

List your org's variants, as a page



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CommerceAPI.GetCommerceVariant(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommerceAPI.GetCommerceVariant``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetCommerceVariantRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCommerceVariantByVariantid

> GetCommerceVariantByVariantid(ctx, variantid).Execute()

Fetch one variant



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	variantid := "variantid_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CommerceAPI.GetCommerceVariantByVariantid(context.Background(), variantid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommerceAPI.GetCommerceVariantByVariantid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**variantid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCommerceVariantByVariantidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCommerceWallet

> GetCommerceWallet(ctx).Execute()

List your org's wallets, as a page



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CommerceAPI.GetCommerceWallet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommerceAPI.GetCommerceWallet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetCommerceWalletRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCommerceWalletByWalletid

> GetCommerceWalletByWalletid(ctx, walletid).Execute()

Fetch one wallet



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	walletid := "walletid_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CommerceAPI.GetCommerceWalletByWalletid(context.Background(), walletid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommerceAPI.GetCommerceWalletByWalletid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**walletid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCommerceWalletByWalletidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCommerceWatchlist

> GetCommerceWatchlist(ctx).Execute()

List your org's watchlists, as a page



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CommerceAPI.GetCommerceWatchlist(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommerceAPI.GetCommerceWatchlist``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetCommerceWatchlistRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCommerceWatchlistByWatchlistid

> GetCommerceWatchlistByWatchlistid(ctx, watchlistid).Execute()

Fetch one watchlist



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	watchlistid := "watchlistid_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CommerceAPI.GetCommerceWatchlistByWatchlistid(context.Background(), watchlistid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommerceAPI.GetCommerceWatchlistByWatchlistid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**watchlistid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCommerceWatchlistByWatchlistidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCommerceWebhook

> GetCommerceWebhook(ctx).Execute()

List your org's webhooks, as a page



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CommerceAPI.GetCommerceWebhook(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommerceAPI.GetCommerceWebhook``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetCommerceWebhookRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCommerceWebhookByWebhookid

> GetCommerceWebhookByWebhookid(ctx, webhookid).Execute()

Fetch one webhook



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	webhookid := "webhookid_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CommerceAPI.GetCommerceWebhookByWebhookid(context.Background(), webhookid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommerceAPI.GetCommerceWebhookByWebhookid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**webhookid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCommerceWebhookByWebhookidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPayment

> PaymentRecord GetPayment(ctx, id).Execute()

Read one settled payment by its id



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	id := "id_example" // string | ID is the ledger transaction id a payment returned.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CommerceAPI.GetPayment(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommerceAPI.GetPayment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPayment`: PaymentRecord
	fmt.Fprintf(os.Stdout, "Response from `CommerceAPI.GetPayment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the ledger transaction id a payment returned. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPaymentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PaymentRecord**](PaymentRecord.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OpenCart

> Cart OpenCart(ctx).CartOpen(cartOpen).Execute()

Open a cart for a shopper to fill



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	cartOpen := *openapiclient.NewCartOpen() // CartOpen | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CommerceAPI.OpenCart(context.Background()).CartOpen(cartOpen).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommerceAPI.OpenCart``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OpenCart`: Cart
	fmt.Fprintf(os.Stdout, "Response from `CommerceAPI.OpenCart`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOpenCartRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cartOpen** | [**CartOpen**](CartOpen.md) |  | 

### Return type

[**Cart**](Cart.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchCommerceCollectionByCollectionid

> PatchCommerceCollectionByCollectionid(ctx, collectionid).Execute()

Change part of a collection



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	collectionid := "collectionid_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CommerceAPI.PatchCommerceCollectionByCollectionid(context.Background(), collectionid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommerceAPI.PatchCommerceCollectionByCollectionid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchCommerceCollectionByCollectionidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchCommerceDisclosureByDisclosureid

> PatchCommerceDisclosureByDisclosureid(ctx, disclosureid).Execute()

Change part of a disclosure



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	disclosureid := "disclosureid_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CommerceAPI.PatchCommerceDisclosureByDisclosureid(context.Background(), disclosureid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommerceAPI.PatchCommerceDisclosureByDisclosureid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**disclosureid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchCommerceDisclosureByDisclosureidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchCommerceDiscountByDiscountid

> PatchCommerceDiscountByDiscountid(ctx, discountid).Execute()

Change part of a discount



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	discountid := "discountid_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CommerceAPI.PatchCommerceDiscountByDiscountid(context.Background(), discountid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommerceAPI.PatchCommerceDiscountByDiscountid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**discountid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchCommerceDiscountByDiscountidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchCommerceMovieByMovieid

> PatchCommerceMovieByMovieid(ctx, movieid).Execute()

Change part of a movie



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	movieid := "movieid_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CommerceAPI.PatchCommerceMovieByMovieid(context.Background(), movieid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommerceAPI.PatchCommerceMovieByMovieid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**movieid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchCommerceMovieByMovieidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchCommerceNoteByNoteid

> PatchCommerceNoteByNoteid(ctx, noteid).Execute()

Change part of a note



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	noteid := "noteid_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CommerceAPI.PatchCommerceNoteByNoteid(context.Background(), noteid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommerceAPI.PatchCommerceNoteByNoteid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**noteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchCommerceNoteByNoteidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchCommerceProductByProductid

> PatchCommerceProductByProductid(ctx, productid).Execute()

Change part of a product



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	productid := "productid_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CommerceAPI.PatchCommerceProductByProductid(context.Background(), productid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommerceAPI.PatchCommerceProductByProductid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**productid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchCommerceProductByProductidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchCommerceReturnByReturnid

> PatchCommerceReturnByReturnid(ctx, returnid).Execute()

Change part of a return



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	returnid := "returnid_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CommerceAPI.PatchCommerceReturnByReturnid(context.Background(), returnid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommerceAPI.PatchCommerceReturnByReturnid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**returnid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchCommerceReturnByReturnidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchCommerceSaleschannelBySaleschannelid

> PatchCommerceSaleschannelBySaleschannelid(ctx, saleschannelid).Execute()

Change part of a sales channel



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	saleschannelid := "saleschannelid_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CommerceAPI.PatchCommerceSaleschannelBySaleschannelid(context.Background(), saleschannelid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommerceAPI.PatchCommerceSaleschannelBySaleschannelid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**saleschannelid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchCommerceSaleschannelBySaleschannelidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchCommerceStocklocationByStocklocationid

> PatchCommerceStocklocationByStocklocationid(ctx, stocklocationid).Execute()

Change part of a stock location



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	stocklocationid := "stocklocationid_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CommerceAPI.PatchCommerceStocklocationByStocklocationid(context.Background(), stocklocationid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommerceAPI.PatchCommerceStocklocationByStocklocationid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stocklocationid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchCommerceStocklocationByStocklocationidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchCommerceStoreByStoreid

> PatchCommerceStoreByStoreid(ctx, storeid).Execute()

Change part of a storefront



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	storeid := "storeid_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CommerceAPI.PatchCommerceStoreByStoreid(context.Background(), storeid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommerceAPI.PatchCommerceStoreByStoreid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**storeid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchCommerceStoreByStoreidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchCommerceStoreByStoreidListingByKey

> PatchCommerceStoreByStoreidListingByKey(ctx, storeid, key).Execute()

Confirm a listing override exists and re-save the store



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	storeid := "storeid_example" // string | 
	key := "key_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CommerceAPI.PatchCommerceStoreByStoreidListingByKey(context.Background(), storeid, key).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommerceAPI.PatchCommerceStoreByStoreidListingByKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**storeid** | **string** |  | 
**key** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchCommerceStoreByStoreidListingByKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchCommerceSubmissionBySubmissionid

> PatchCommerceSubmissionBySubmissionid(ctx, submissionid).Execute()

Change part of a submission



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	submissionid := "submissionid_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CommerceAPI.PatchCommerceSubmissionBySubmissionid(context.Background(), submissionid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommerceAPI.PatchCommerceSubmissionBySubmissionid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**submissionid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchCommerceSubmissionBySubmissionidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchCommerceSubscriberBySubscriberid

> PatchCommerceSubscriberBySubscriberid(ctx, subscriberid).Execute()

Change part of a subscriber



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	subscriberid := "subscriberid_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CommerceAPI.PatchCommerceSubscriberBySubscriberid(context.Background(), subscriberid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommerceAPI.PatchCommerceSubscriberBySubscriberid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriberid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchCommerceSubscriberBySubscriberidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchCommerceTokentransactionByTokentransactionid

> PatchCommerceTokentransactionByTokentransactionid(ctx, tokentransactionid).Execute()

Change part of a token transaction



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	tokentransactionid := "tokentransactionid_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CommerceAPI.PatchCommerceTokentransactionByTokentransactionid(context.Background(), tokentransactionid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommerceAPI.PatchCommerceTokentransactionByTokentransactionid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tokentransactionid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchCommerceTokentransactionByTokentransactionidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchCommerceTransferByTransferid

> PatchCommerceTransferByTransferid(ctx, transferid).Execute()

Change part of a transfer



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	transferid := "transferid_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CommerceAPI.PatchCommerceTransferByTransferid(context.Background(), transferid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommerceAPI.PatchCommerceTransferByTransferid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**transferid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchCommerceTransferByTransferidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchCommerceVariantByVariantid

> PatchCommerceVariantByVariantid(ctx, variantid).Execute()

Change part of a variant



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	variantid := "variantid_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CommerceAPI.PatchCommerceVariantByVariantid(context.Background(), variantid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommerceAPI.PatchCommerceVariantByVariantid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**variantid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchCommerceVariantByVariantidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchCommerceWalletByWalletid

> PatchCommerceWalletByWalletid(ctx, walletid).Execute()

Change part of a wallet



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	walletid := "walletid_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CommerceAPI.PatchCommerceWalletByWalletid(context.Background(), walletid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommerceAPI.PatchCommerceWalletByWalletid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**walletid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchCommerceWalletByWalletidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchCommerceWatchlistByWatchlistid

> PatchCommerceWatchlistByWatchlistid(ctx, watchlistid).Execute()

Change part of a watchlist



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	watchlistid := "watchlistid_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CommerceAPI.PatchCommerceWatchlistByWatchlistid(context.Background(), watchlistid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommerceAPI.PatchCommerceWatchlistByWatchlistid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**watchlistid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchCommerceWatchlistByWatchlistidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchCommerceWebhookByWebhookid

> PatchCommerceWebhookByWebhookid(ctx, webhookid).Execute()

Change part of a webhook



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	webhookid := "webhookid_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CommerceAPI.PatchCommerceWebhookByWebhookid(context.Background(), webhookid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommerceAPI.PatchCommerceWebhookByWebhookid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**webhookid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchCommerceWebhookByWebhookidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCommerceCatalogEntries

> PostCommerceCatalogEntries(ctx).Execute()

Add a catalog entry



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CommerceAPI.PostCommerceCatalogEntries(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommerceAPI.PostCommerceCatalogEntries``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPostCommerceCatalogEntriesRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCommerceCatalogModels

> PostCommerceCatalogModels(ctx).Execute()

Land a syncer's view of the model catalog: upstream costs and machine facts



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CommerceAPI.PostCommerceCatalogModels(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommerceAPI.PostCommerceCatalogModels``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPostCommerceCatalogModelsRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCommerceCatalogModelsRefresh

> PostCommerceCatalogModelsRefresh(ctx).Execute()

Refresh the model catalog by reading the upstream provider



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CommerceAPI.PostCommerceCatalogModelsRefresh(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommerceAPI.PostCommerceCatalogModelsRefresh``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPostCommerceCatalogModelsRefreshRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCommerceCatalogSeed

> PostCommerceCatalogSeed(ctx).Execute()

Seed the embedded catalog, without disturbing edits already made



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CommerceAPI.PostCommerceCatalogSeed(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommerceAPI.PostCommerceCatalogSeed``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPostCommerceCatalogSeedRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCommerceCollection

> PostCommerceCollection(ctx).Execute()

Create a collection



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CommerceAPI.PostCommerceCollection(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommerceAPI.PostCommerceCollection``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPostCommerceCollectionRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCommerceCollectionByCollectionid

> PostCommerceCollectionByCollectionid(ctx, collectionid).Execute()

Method-override tunnel for a collection — for clients that cannot send PUT, PATCH or DELETE



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	collectionid := "collectionid_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CommerceAPI.PostCommerceCollectionByCollectionid(context.Background(), collectionid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommerceAPI.PostCommerceCollectionByCollectionid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostCommerceCollectionByCollectionidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCommerceDisclosure

> PostCommerceDisclosure(ctx).Execute()

Create a disclosure



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CommerceAPI.PostCommerceDisclosure(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommerceAPI.PostCommerceDisclosure``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPostCommerceDisclosureRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCommerceDisclosureByDisclosureid

> PostCommerceDisclosureByDisclosureid(ctx, disclosureid).Execute()

Method-override tunnel for a disclosure — for clients that cannot send PUT, PATCH or DELETE



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	disclosureid := "disclosureid_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CommerceAPI.PostCommerceDisclosureByDisclosureid(context.Background(), disclosureid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommerceAPI.PostCommerceDisclosureByDisclosureid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**disclosureid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostCommerceDisclosureByDisclosureidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCommerceDiscount

> PostCommerceDiscount(ctx).Execute()

Create a discount



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CommerceAPI.PostCommerceDiscount(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommerceAPI.PostCommerceDiscount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPostCommerceDiscountRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCommerceDiscountByDiscountid

> PostCommerceDiscountByDiscountid(ctx, discountid).Execute()

Method-override tunnel for a discount — for clients that cannot send PUT, PATCH or DELETE



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	discountid := "discountid_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CommerceAPI.PostCommerceDiscountByDiscountid(context.Background(), discountid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommerceAPI.PostCommerceDiscountByDiscountid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**discountid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostCommerceDiscountByDiscountidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCommerceMovie

> PostCommerceMovie(ctx).Execute()

Create a movie



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CommerceAPI.PostCommerceMovie(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommerceAPI.PostCommerceMovie``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPostCommerceMovieRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCommerceMovieByMovieid

> PostCommerceMovieByMovieid(ctx, movieid).Execute()

Method-override tunnel for a movie — for clients that cannot send PUT, PATCH or DELETE



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	movieid := "movieid_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CommerceAPI.PostCommerceMovieByMovieid(context.Background(), movieid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommerceAPI.PostCommerceMovieByMovieid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**movieid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostCommerceMovieByMovieidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCommerceNote

> PostCommerceNote(ctx).Execute()

Create a note



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CommerceAPI.PostCommerceNote(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommerceAPI.PostCommerceNote``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPostCommerceNoteRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCommerceNoteByNoteid

> PostCommerceNoteByNoteid(ctx, noteid).Execute()

Method-override tunnel for a note — for clients that cannot send PUT, PATCH or DELETE



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	noteid := "noteid_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CommerceAPI.PostCommerceNoteByNoteid(context.Background(), noteid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommerceAPI.PostCommerceNoteByNoteid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**noteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostCommerceNoteByNoteidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCommercePlansEntries

> PostCommercePlansEntries(ctx).Execute()

Add a subscription plan



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CommerceAPI.PostCommercePlansEntries(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommerceAPI.PostCommercePlansEntries``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPostCommercePlansEntriesRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCommercePlansSeed

> PostCommercePlansSeed(ctx).Execute()

Seed the embedded plan catalog, without overwriting administrative edits



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CommerceAPI.PostCommercePlansSeed(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommerceAPI.PostCommercePlansSeed``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPostCommercePlansSeedRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCommerceProduct

> PostCommerceProduct(ctx).Execute()

Create a product



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CommerceAPI.PostCommerceProduct(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommerceAPI.PostCommerceProduct``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPostCommerceProductRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCommerceProductByProductid

> PostCommerceProductByProductid(ctx, productid).Execute()

Method-override tunnel for a product — for clients that cannot send PUT, PATCH or DELETE



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	productid := "productid_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CommerceAPI.PostCommerceProductByProductid(context.Background(), productid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommerceAPI.PostCommerceProductByProductid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**productid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostCommerceProductByProductidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCommerceRatesEntries

> PostCommerceRatesEntries(ctx).Execute()

Add a rate



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CommerceAPI.PostCommerceRatesEntries(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommerceAPI.PostCommerceRatesEntries``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPostCommerceRatesEntriesRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCommerceRatesImport

> PostCommerceRatesImport(ctx).Execute()

Load the published price document, reconciling rather than replacing



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CommerceAPI.PostCommerceRatesImport(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommerceAPI.PostCommerceRatesImport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPostCommerceRatesImportRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCommerceReturn

> PostCommerceReturn(ctx).Execute()

Create a return



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CommerceAPI.PostCommerceReturn(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommerceAPI.PostCommerceReturn``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPostCommerceReturnRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCommerceReturnByReturnid

> PostCommerceReturnByReturnid(ctx, returnid).Execute()

Method-override tunnel for a return — for clients that cannot send PUT, PATCH or DELETE



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	returnid := "returnid_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CommerceAPI.PostCommerceReturnByReturnid(context.Background(), returnid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommerceAPI.PostCommerceReturnByReturnid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**returnid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostCommerceReturnByReturnidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCommerceSaleschannel

> PostCommerceSaleschannel(ctx).Execute()

Create a sales channel



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CommerceAPI.PostCommerceSaleschannel(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommerceAPI.PostCommerceSaleschannel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPostCommerceSaleschannelRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCommerceSaleschannelBySaleschannelid

> PostCommerceSaleschannelBySaleschannelid(ctx, saleschannelid).Execute()

Method-override tunnel for a sales channel — for clients that cannot send PUT, PATCH or DELETE



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	saleschannelid := "saleschannelid_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CommerceAPI.PostCommerceSaleschannelBySaleschannelid(context.Background(), saleschannelid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommerceAPI.PostCommerceSaleschannelBySaleschannelid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**saleschannelid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostCommerceSaleschannelBySaleschannelidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCommerceStocklocation

> PostCommerceStocklocation(ctx).Execute()

Create a stock location



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CommerceAPI.PostCommerceStocklocation(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommerceAPI.PostCommerceStocklocation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPostCommerceStocklocationRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCommerceStocklocationByStocklocationid

> PostCommerceStocklocationByStocklocationid(ctx, stocklocationid).Execute()

Method-override tunnel for a stock location — for clients that cannot send PUT, PATCH or DELETE



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	stocklocationid := "stocklocationid_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CommerceAPI.PostCommerceStocklocationByStocklocationid(context.Background(), stocklocationid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommerceAPI.PostCommerceStocklocationByStocklocationid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stocklocationid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostCommerceStocklocationByStocklocationidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCommerceStore

> PostCommerceStore(ctx).Execute()

Create a storefront



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CommerceAPI.PostCommerceStore(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommerceAPI.PostCommerceStore``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPostCommerceStoreRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCommerceStoreByStoreid

> PostCommerceStoreByStoreid(ctx, storeid).Execute()

Method-override tunnel for clients that cannot send PUT, PATCH or DELETE



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	storeid := "storeid_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CommerceAPI.PostCommerceStoreByStoreid(context.Background(), storeid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommerceAPI.PostCommerceStoreByStoreid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**storeid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostCommerceStoreByStoreidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCommerceStoreByStoreidAuthorize

> PostCommerceStoreByStoreidAuthorize(ctx, storeid).Execute()

Authorize a new order against a storefront, holding the funds without settling them



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	storeid := "storeid_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CommerceAPI.PostCommerceStoreByStoreidAuthorize(context.Background(), storeid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommerceAPI.PostCommerceStoreByStoreidAuthorize``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**storeid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostCommerceStoreByStoreidAuthorizeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCommerceStoreByStoreidAuthorizeByOrderid

> PostCommerceStoreByStoreidAuthorizeByOrderid(ctx, storeid, orderid).Execute()

Authorize an order that already exists, holding the funds without settling them



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	storeid := "storeid_example" // string | 
	orderid := "orderid_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CommerceAPI.PostCommerceStoreByStoreidAuthorizeByOrderid(context.Background(), storeid, orderid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommerceAPI.PostCommerceStoreByStoreidAuthorizeByOrderid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**storeid** | **string** |  | 
**orderid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostCommerceStoreByStoreidAuthorizeByOrderidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCommerceStoreByStoreidCaptureByOrderid

> PostCommerceStoreByStoreidCaptureByOrderid(ctx, storeid, orderid).Execute()

Capture a previously authorized order and settle the payment



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	storeid := "storeid_example" // string | 
	orderid := "orderid_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CommerceAPI.PostCommerceStoreByStoreidCaptureByOrderid(context.Background(), storeid, orderid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommerceAPI.PostCommerceStoreByStoreidCaptureByOrderid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**storeid** | **string** |  | 
**orderid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostCommerceStoreByStoreidCaptureByOrderidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCommerceStoreByStoreidCharge

> PostCommerceStoreByStoreidCharge(ctx, storeid).Execute()

Authorize and capture a new order in one call



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	storeid := "storeid_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CommerceAPI.PostCommerceStoreByStoreidCharge(context.Background(), storeid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommerceAPI.PostCommerceStoreByStoreidCharge``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**storeid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostCommerceStoreByStoreidChargeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCommerceStoreByStoreidCheckoutAuthorize

> PostCommerceStoreByStoreidCheckoutAuthorize(ctx, storeid).Execute()

Authorize a new order against a storefront, holding the funds — the checkout spelling



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	storeid := "storeid_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CommerceAPI.PostCommerceStoreByStoreidCheckoutAuthorize(context.Background(), storeid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommerceAPI.PostCommerceStoreByStoreidCheckoutAuthorize``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**storeid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostCommerceStoreByStoreidCheckoutAuthorizeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCommerceStoreByStoreidCheckoutAuthorizeByOrderid

> PostCommerceStoreByStoreidCheckoutAuthorizeByOrderid(ctx, storeid, orderid).Execute()

Authorize an existing order, holding the funds — the checkout spelling



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	storeid := "storeid_example" // string | 
	orderid := "orderid_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CommerceAPI.PostCommerceStoreByStoreidCheckoutAuthorizeByOrderid(context.Background(), storeid, orderid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommerceAPI.PostCommerceStoreByStoreidCheckoutAuthorizeByOrderid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**storeid** | **string** |  | 
**orderid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostCommerceStoreByStoreidCheckoutAuthorizeByOrderidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCommerceStoreByStoreidCheckoutCaptureByOrderid

> PostCommerceStoreByStoreidCheckoutCaptureByOrderid(ctx, storeid, orderid).Execute()

Capture a previously authorized order and settle it — the checkout spelling



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	storeid := "storeid_example" // string | 
	orderid := "orderid_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CommerceAPI.PostCommerceStoreByStoreidCheckoutCaptureByOrderid(context.Background(), storeid, orderid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommerceAPI.PostCommerceStoreByStoreidCheckoutCaptureByOrderid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**storeid** | **string** |  | 
**orderid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostCommerceStoreByStoreidCheckoutCaptureByOrderidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCommerceStoreByStoreidCheckoutCharge

> PostCommerceStoreByStoreidCheckoutCharge(ctx, storeid).Execute()

Authorize and capture a new order in one call — the checkout spelling



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	storeid := "storeid_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CommerceAPI.PostCommerceStoreByStoreidCheckoutCharge(context.Background(), storeid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommerceAPI.PostCommerceStoreByStoreidCheckoutCharge``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**storeid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostCommerceStoreByStoreidCheckoutChargeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCommerceStoreByStoreidCheckoutPaypalCancelByPaykey

> PostCommerceStoreByStoreidCheckoutPaypalCancelByPaykey(ctx, storeid, payKey).Execute()

PayPal cancel by pay key — refuses, exactly as the unprefixed address does



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	storeid := "storeid_example" // string | 
	payKey := "payKey_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CommerceAPI.PostCommerceStoreByStoreidCheckoutPaypalCancelByPaykey(context.Background(), storeid, payKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommerceAPI.PostCommerceStoreByStoreidCheckoutPaypalCancelByPaykey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**storeid** | **string** |  | 
**payKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostCommerceStoreByStoreidCheckoutPaypalCancelByPaykeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCommerceStoreByStoreidCheckoutPaypalConfirmByPaykey

> PostCommerceStoreByStoreidCheckoutPaypalConfirmByPaykey(ctx, storeid, payKey).Execute()

PayPal confirm by pay key — refuses, exactly as the unprefixed address does



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	storeid := "storeid_example" // string | 
	payKey := "payKey_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CommerceAPI.PostCommerceStoreByStoreidCheckoutPaypalConfirmByPaykey(context.Background(), storeid, payKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommerceAPI.PostCommerceStoreByStoreidCheckoutPaypalConfirmByPaykey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**storeid** | **string** |  | 
**payKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostCommerceStoreByStoreidCheckoutPaypalConfirmByPaykeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCommerceStoreByStoreidCheckoutPaypalPay

> PostCommerceStoreByStoreidCheckoutPaypalPay(ctx, storeid).Execute()

Start a PayPal authorization for a new order — the checkout spelling



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	storeid := "storeid_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CommerceAPI.PostCommerceStoreByStoreidCheckoutPaypalPay(context.Background(), storeid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommerceAPI.PostCommerceStoreByStoreidCheckoutPaypalPay``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**storeid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostCommerceStoreByStoreidCheckoutPaypalPayRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCommerceStoreByStoreidListingByKey

> PostCommerceStoreByStoreidListingByKey(ctx, storeid, key).Execute()

Add a listing override under a new key



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	storeid := "storeid_example" // string | 
	key := "key_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CommerceAPI.PostCommerceStoreByStoreidListingByKey(context.Background(), storeid, key).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommerceAPI.PostCommerceStoreByStoreidListingByKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**storeid** | **string** |  | 
**key** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostCommerceStoreByStoreidListingByKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCommerceStoreByStoreidPaypalCancelByPaykey

> PostCommerceStoreByStoreidPaypalCancelByPaykey(ctx, storeid, payKey).Execute()

PayPal cancel by pay key — refuses, because a pay key alone does not identify the order



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	storeid := "storeid_example" // string | 
	payKey := "payKey_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CommerceAPI.PostCommerceStoreByStoreidPaypalCancelByPaykey(context.Background(), storeid, payKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommerceAPI.PostCommerceStoreByStoreidPaypalCancelByPaykey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**storeid** | **string** |  | 
**payKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostCommerceStoreByStoreidPaypalCancelByPaykeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCommerceStoreByStoreidPaypalConfirmByPaykey

> PostCommerceStoreByStoreidPaypalConfirmByPaykey(ctx, storeid, payKey).Execute()

PayPal confirm by pay key — refuses, because a pay key alone does not identify the order



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	storeid := "storeid_example" // string | 
	payKey := "payKey_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CommerceAPI.PostCommerceStoreByStoreidPaypalConfirmByPaykey(context.Background(), storeid, payKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommerceAPI.PostCommerceStoreByStoreidPaypalConfirmByPaykey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**storeid** | **string** |  | 
**payKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostCommerceStoreByStoreidPaypalConfirmByPaykeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCommerceStoreByStoreidPaypalPay

> PostCommerceStoreByStoreidPaypalPay(ctx, storeid).Execute()

Start a PayPal authorization for a new order



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	storeid := "storeid_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CommerceAPI.PostCommerceStoreByStoreidPaypalPay(context.Background(), storeid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommerceAPI.PostCommerceStoreByStoreidPaypalPay``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**storeid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostCommerceStoreByStoreidPaypalPayRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCommerceStoreByStoreidTrial

> PostCommerceStoreByStoreidTrial(ctx, storeid).Execute()

Start this store's no-card trial on the entry plan



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	storeid := "storeid_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CommerceAPI.PostCommerceStoreByStoreidTrial(context.Background(), storeid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommerceAPI.PostCommerceStoreByStoreidTrial``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**storeid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostCommerceStoreByStoreidTrialRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCommerceStoreToken

> PostCommerceStoreToken(ctx).Execute()

Mint your org's least-privilege storefront read key



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CommerceAPI.PostCommerceStoreToken(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommerceAPI.PostCommerceStoreToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPostCommerceStoreTokenRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCommerceSubmission

> PostCommerceSubmission(ctx).Execute()

Create a submission



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CommerceAPI.PostCommerceSubmission(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommerceAPI.PostCommerceSubmission``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPostCommerceSubmissionRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCommerceSubmissionBySubmissionid

> PostCommerceSubmissionBySubmissionid(ctx, submissionid).Execute()

Method-override tunnel for a submission — for clients that cannot send PUT, PATCH or DELETE



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	submissionid := "submissionid_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CommerceAPI.PostCommerceSubmissionBySubmissionid(context.Background(), submissionid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommerceAPI.PostCommerceSubmissionBySubmissionid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**submissionid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostCommerceSubmissionBySubmissionidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCommerceSubscriber

> PostCommerceSubscriber(ctx).Execute()

Create a subscriber



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CommerceAPI.PostCommerceSubscriber(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommerceAPI.PostCommerceSubscriber``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPostCommerceSubscriberRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCommerceSubscriberBySubscriberid

> PostCommerceSubscriberBySubscriberid(ctx, subscriberid).Execute()

Method-override tunnel for a subscriber — for clients that cannot send PUT, PATCH or DELETE



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	subscriberid := "subscriberid_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CommerceAPI.PostCommerceSubscriberBySubscriberid(context.Background(), subscriberid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommerceAPI.PostCommerceSubscriberBySubscriberid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriberid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostCommerceSubscriberBySubscriberidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCommerceTokentransaction

> PostCommerceTokentransaction(ctx).Execute()

Create a token transaction



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CommerceAPI.PostCommerceTokentransaction(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommerceAPI.PostCommerceTokentransaction``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPostCommerceTokentransactionRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCommerceTokentransactionByTokentransactionid

> PostCommerceTokentransactionByTokentransactionid(ctx, tokentransactionid).Execute()

Method-override tunnel for a token transaction — for clients that cannot send PUT, PATCH or DELETE



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	tokentransactionid := "tokentransactionid_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CommerceAPI.PostCommerceTokentransactionByTokentransactionid(context.Background(), tokentransactionid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommerceAPI.PostCommerceTokentransactionByTokentransactionid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tokentransactionid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostCommerceTokentransactionByTokentransactionidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCommerceTransfer

> PostCommerceTransfer(ctx).Execute()

Create a transfer



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CommerceAPI.PostCommerceTransfer(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommerceAPI.PostCommerceTransfer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPostCommerceTransferRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCommerceTransferByTransferid

> PostCommerceTransferByTransferid(ctx, transferid).Execute()

Method-override tunnel for a transfer — for clients that cannot send PUT, PATCH or DELETE



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	transferid := "transferid_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CommerceAPI.PostCommerceTransferByTransferid(context.Background(), transferid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommerceAPI.PostCommerceTransferByTransferid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**transferid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostCommerceTransferByTransferidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCommerceVariant

> PostCommerceVariant(ctx).Execute()

Create a variant



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CommerceAPI.PostCommerceVariant(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommerceAPI.PostCommerceVariant``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPostCommerceVariantRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCommerceVariantByVariantid

> PostCommerceVariantByVariantid(ctx, variantid).Execute()

Method-override tunnel for a variant — for clients that cannot send PUT, PATCH or DELETE



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	variantid := "variantid_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CommerceAPI.PostCommerceVariantByVariantid(context.Background(), variantid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommerceAPI.PostCommerceVariantByVariantid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**variantid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostCommerceVariantByVariantidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCommerceWallet

> PostCommerceWallet(ctx).Execute()

Create a wallet



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CommerceAPI.PostCommerceWallet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommerceAPI.PostCommerceWallet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPostCommerceWalletRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCommerceWalletByWalletid

> PostCommerceWalletByWalletid(ctx, walletid).Execute()

Method-override tunnel for a wallet — for clients that cannot send PUT, PATCH or DELETE



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	walletid := "walletid_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CommerceAPI.PostCommerceWalletByWalletid(context.Background(), walletid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommerceAPI.PostCommerceWalletByWalletid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**walletid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostCommerceWalletByWalletidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCommerceWatchlist

> PostCommerceWatchlist(ctx).Execute()

Create a watchlist



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CommerceAPI.PostCommerceWatchlist(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommerceAPI.PostCommerceWatchlist``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPostCommerceWatchlistRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCommerceWatchlistByWatchlistid

> PostCommerceWatchlistByWatchlistid(ctx, watchlistid).Execute()

Method-override tunnel for a watchlist — for clients that cannot send PUT, PATCH or DELETE



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	watchlistid := "watchlistid_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CommerceAPI.PostCommerceWatchlistByWatchlistid(context.Background(), watchlistid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommerceAPI.PostCommerceWatchlistByWatchlistid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**watchlistid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostCommerceWatchlistByWatchlistidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCommerceWebhook

> PostCommerceWebhook(ctx).Execute()

Create a webhook



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CommerceAPI.PostCommerceWebhook(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommerceAPI.PostCommerceWebhook``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPostCommerceWebhookRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCommerceWebhookByWebhookid

> PostCommerceWebhookByWebhookid(ctx, webhookid).Execute()

Method-override tunnel for a webhook — for clients that cannot send PUT, PATCH or DELETE



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	webhookid := "webhookid_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CommerceAPI.PostCommerceWebhookByWebhookid(context.Background(), webhookid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommerceAPI.PostCommerceWebhookByWebhookid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**webhookid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostCommerceWebhookByWebhookidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCommerceWebhooksByProvider

> PostCommerceWebhooksByProvider(ctx, provider).Execute()

Payment-provider webhook intake for settlement and subscription lifecycle events



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	provider := "provider_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CommerceAPI.PostCommerceWebhooksByProvider(context.Background(), provider).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommerceAPI.PostCommerceWebhooksByProvider``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**provider** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostCommerceWebhooksByProviderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutCommerceCollectionByCollectionid

> PutCommerceCollectionByCollectionid(ctx, collectionid).Execute()

Replace a collection outright



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	collectionid := "collectionid_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CommerceAPI.PutCommerceCollectionByCollectionid(context.Background(), collectionid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommerceAPI.PutCommerceCollectionByCollectionid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutCommerceCollectionByCollectionidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutCommerceDisclosureByDisclosureid

> PutCommerceDisclosureByDisclosureid(ctx, disclosureid).Execute()

Replace a disclosure outright



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	disclosureid := "disclosureid_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CommerceAPI.PutCommerceDisclosureByDisclosureid(context.Background(), disclosureid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommerceAPI.PutCommerceDisclosureByDisclosureid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**disclosureid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutCommerceDisclosureByDisclosureidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutCommerceDiscountByDiscountid

> PutCommerceDiscountByDiscountid(ctx, discountid).Execute()

Replace a discount outright



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	discountid := "discountid_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CommerceAPI.PutCommerceDiscountByDiscountid(context.Background(), discountid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommerceAPI.PutCommerceDiscountByDiscountid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**discountid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutCommerceDiscountByDiscountidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutCommerceMovieByMovieid

> PutCommerceMovieByMovieid(ctx, movieid).Execute()

Replace a movie outright



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	movieid := "movieid_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CommerceAPI.PutCommerceMovieByMovieid(context.Background(), movieid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommerceAPI.PutCommerceMovieByMovieid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**movieid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutCommerceMovieByMovieidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutCommerceNoteByNoteid

> PutCommerceNoteByNoteid(ctx, noteid).Execute()

Replace a note outright



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	noteid := "noteid_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CommerceAPI.PutCommerceNoteByNoteid(context.Background(), noteid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommerceAPI.PutCommerceNoteByNoteid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**noteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutCommerceNoteByNoteidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutCommercePlansEntriesBySlug

> PutCommercePlansEntriesBySlug(ctx, slug).Execute()

Edit a plan, leaving the fields you omit alone



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	slug := "slug_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CommerceAPI.PutCommercePlansEntriesBySlug(context.Background(), slug).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommerceAPI.PutCommercePlansEntriesBySlug``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slug** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutCommercePlansEntriesBySlugRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutCommerceProductByProductid

> PutCommerceProductByProductid(ctx, productid).Execute()

Replace a product outright



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	productid := "productid_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CommerceAPI.PutCommerceProductByProductid(context.Background(), productid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommerceAPI.PutCommerceProductByProductid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**productid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutCommerceProductByProductidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutCommerceRatesEntriesByProductByMeter

> PutCommerceRatesEntriesByProductByMeter(ctx, product, meter).Execute()

Edit a rate, and mark it as operator-set



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	product := "product_example" // string | 
	meter := "meter_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CommerceAPI.PutCommerceRatesEntriesByProductByMeter(context.Background(), product, meter).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommerceAPI.PutCommerceRatesEntriesByProductByMeter``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**product** | **string** |  | 
**meter** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutCommerceRatesEntriesByProductByMeterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutCommerceReturnByReturnid

> PutCommerceReturnByReturnid(ctx, returnid).Execute()

Replace a return outright



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	returnid := "returnid_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CommerceAPI.PutCommerceReturnByReturnid(context.Background(), returnid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommerceAPI.PutCommerceReturnByReturnid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**returnid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutCommerceReturnByReturnidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutCommerceSaleschannelBySaleschannelid

> PutCommerceSaleschannelBySaleschannelid(ctx, saleschannelid).Execute()

Replace a sales channel outright



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	saleschannelid := "saleschannelid_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CommerceAPI.PutCommerceSaleschannelBySaleschannelid(context.Background(), saleschannelid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommerceAPI.PutCommerceSaleschannelBySaleschannelid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**saleschannelid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutCommerceSaleschannelBySaleschannelidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutCommerceStocklocationByStocklocationid

> PutCommerceStocklocationByStocklocationid(ctx, stocklocationid).Execute()

Replace a stock location outright



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	stocklocationid := "stocklocationid_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CommerceAPI.PutCommerceStocklocationByStocklocationid(context.Background(), stocklocationid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommerceAPI.PutCommerceStocklocationByStocklocationid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stocklocationid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutCommerceStocklocationByStocklocationidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutCommerceStoreByStoreid

> PutCommerceStoreByStoreid(ctx, storeid).Execute()

Replace a storefront outright



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	storeid := "storeid_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CommerceAPI.PutCommerceStoreByStoreid(context.Background(), storeid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommerceAPI.PutCommerceStoreByStoreid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**storeid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutCommerceStoreByStoreidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutCommerceStoreByStoreidListingByKey

> PutCommerceStoreByStoreidListingByKey(ctx, storeid, key).Execute()

Upsert a listing override



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	storeid := "storeid_example" // string | 
	key := "key_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CommerceAPI.PutCommerceStoreByStoreidListingByKey(context.Background(), storeid, key).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommerceAPI.PutCommerceStoreByStoreidListingByKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**storeid** | **string** |  | 
**key** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutCommerceStoreByStoreidListingByKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutCommerceSubmissionBySubmissionid

> PutCommerceSubmissionBySubmissionid(ctx, submissionid).Execute()

Replace a submission outright



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	submissionid := "submissionid_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CommerceAPI.PutCommerceSubmissionBySubmissionid(context.Background(), submissionid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommerceAPI.PutCommerceSubmissionBySubmissionid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**submissionid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutCommerceSubmissionBySubmissionidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutCommerceSubscriberBySubscriberid

> PutCommerceSubscriberBySubscriberid(ctx, subscriberid).Execute()

Replace a subscriber outright



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	subscriberid := "subscriberid_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CommerceAPI.PutCommerceSubscriberBySubscriberid(context.Background(), subscriberid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommerceAPI.PutCommerceSubscriberBySubscriberid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriberid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutCommerceSubscriberBySubscriberidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutCommerceTokentransactionByTokentransactionid

> PutCommerceTokentransactionByTokentransactionid(ctx, tokentransactionid).Execute()

Replace a token transaction outright



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	tokentransactionid := "tokentransactionid_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CommerceAPI.PutCommerceTokentransactionByTokentransactionid(context.Background(), tokentransactionid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommerceAPI.PutCommerceTokentransactionByTokentransactionid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tokentransactionid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutCommerceTokentransactionByTokentransactionidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutCommerceTransferByTransferid

> PutCommerceTransferByTransferid(ctx, transferid).Execute()

Replace a transfer outright



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	transferid := "transferid_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CommerceAPI.PutCommerceTransferByTransferid(context.Background(), transferid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommerceAPI.PutCommerceTransferByTransferid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**transferid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutCommerceTransferByTransferidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutCommerceVariantByVariantid

> PutCommerceVariantByVariantid(ctx, variantid).Execute()

Replace a variant outright



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	variantid := "variantid_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CommerceAPI.PutCommerceVariantByVariantid(context.Background(), variantid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommerceAPI.PutCommerceVariantByVariantid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**variantid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutCommerceVariantByVariantidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutCommerceWalletByWalletid

> PutCommerceWalletByWalletid(ctx, walletid).Execute()

Replace a wallet outright



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	walletid := "walletid_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CommerceAPI.PutCommerceWalletByWalletid(context.Background(), walletid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommerceAPI.PutCommerceWalletByWalletid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**walletid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutCommerceWalletByWalletidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutCommerceWatchlistByWatchlistid

> PutCommerceWatchlistByWatchlistid(ctx, watchlistid).Execute()

Replace a watchlist outright



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	watchlistid := "watchlistid_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CommerceAPI.PutCommerceWatchlistByWatchlistid(context.Background(), watchlistid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommerceAPI.PutCommerceWatchlistByWatchlistid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**watchlistid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutCommerceWatchlistByWatchlistidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutCommerceWebhookByWebhookid

> PutCommerceWebhookByWebhookid(ctx, webhookid).Execute()

Replace a webhook outright



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	webhookid := "webhookid_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CommerceAPI.PutCommerceWebhookByWebhookid(context.Background(), webhookid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommerceAPI.PutCommerceWebhookByWebhookid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**webhookid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutCommerceWebhookByWebhookidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetCartItem

> Cart SetCartItem(ctx, id).CartItemSet(cartItemSet).Execute()

Set one item's quantity in a cart; zero removes it



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	id := "id_example" // string | ID is the cart to amend, from the path.
	cartItemSet := *openapiclient.NewCartItemSet() // CartItemSet | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CommerceAPI.SetCartItem(context.Background(), id).CartItemSet(cartItemSet).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommerceAPI.SetCartItem``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetCartItem`: Cart
	fmt.Fprintf(os.Stdout, "Response from `CommerceAPI.SetCartItem`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the cart to amend, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetCartItemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cartItemSet** | [**CartItemSet**](CartItemSet.md) |  | 

### Return type

[**Cart**](Cart.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TakePayment

> PaymentOut TakePayment(ctx).PaymentIn(paymentIn).Execute()

Take a card payment and credit the org's balance



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	paymentIn := *openapiclient.NewPaymentIn() // PaymentIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CommerceAPI.TakePayment(context.Background()).PaymentIn(paymentIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommerceAPI.TakePayment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TakePayment`: PaymentOut
	fmt.Fprintf(os.Stdout, "Response from `CommerceAPI.TakePayment`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTakePaymentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **paymentIn** | [**PaymentIn**](PaymentIn.md) |  | 

### Return type

[**PaymentOut**](PaymentOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

