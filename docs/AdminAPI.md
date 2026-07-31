# \AdminAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AnalyticsAdminListUsers**](AdminAPI.md#AnalyticsAdminListUsers) | **Get** /v1/analytics/admin/users | List all users (admin only)
[**AnalyticsAdminListWebsites**](AdminAPI.md#AnalyticsAdminListWebsites) | **Get** /v1/analytics/admin/websites | List all websites for a user (admin only)
[**CloudAdminAIMetrics**](AdminAPI.md#CloudAdminAIMetrics) | **Get** /v1/admin/aimetrics | Is the fleet AI board: O11yAI generations (count, cost, avg/p95 latency, per-model), per-model usage from the live cloud_usage ledger, and the eval plane (traces, scores, score names, runs, and the average-score trend).
[**CloudAdminAnalytics**](AdminAPI.md#CloudAdminAnalytics) | **Get** /v1/admin/analytics | Is the SaaS product-analytics board over the caller&#39;s tenant window: active customers, new and churned, retention, MRR, ARPU, the usage trend and the top customers by spend — every number folded from the commerce ledger, not sampled.
[**CloudAdminApplications**](AdminAPI.md#CloudAdminApplications) | **Get** /v1/admin/applications | Lists IAM applications for one owner org, forwarded VERBATIM from IAM&#39;s get-applications.
[**CloudAdminAudit**](AdminAPI.md#CloudAdminAudit) | **Get** /v1/admin/audit | Reads cloud&#39;s tamper-evident audit trail, newest first, with the chain&#39;s live integrity attached so a listing can be badged as verified.
[**CloudAdminAuditVerify**](AdminAPI.md#CloudAdminAuditVerify) | **Get** /v1/admin/audit/verify | Walks the WHOLE hash chain and reports whether it is intact: how many records were checked, the head hash to pin externally against tail-truncation, and — when the chain is broken — the seq of the first bad record and why.
[**CloudAdminBases**](AdminAPI.md#CloudAdminBases) | **Get** /v1/admin/bases | Lists the tenant Base instances in the caller&#39;s window — a SuperAdmin sees every tenant&#39;s, anyone else only their own subtree&#39;s.
[**CloudAdminBlockStorage**](AdminAPI.md#CloudAdminBlockStorage) | **Get** /v1/admin/block-storage | Is the realtime block-storage board: the DigitalOcean volume fleet (count, capacity, monthly list cost, per-volume region and attachment) plus the analytics datastore&#39;s OWN fill, read from its system.disks.
[**CloudAdminCompute**](AdminAPI.md#CloudAdminCompute) | **Get** /v1/admin/compute | Rolls the fleet&#39;s compute usage up to one row per (org, app, project, kind): how many distinct machines ran in the window, how many are still active, what they billed, and when each group last emitted an event.
[**CloudAdminCordonNode**](AdminAPI.md#CloudAdminCordonNode) | **Post** /v1/admin/infra/nodes/{id}/cordon | Marks one cluster node unschedulable — or schedulable again — and can drain the pods already on it.
[**CloudAdminCreateCreditGrant**](AdminAPI.md#CloudAdminCreateCreditGrant) | **Post** /v1/admin/credit-grants | Mints credit for one org.
[**CloudAdminCreateSpendCap**](AdminAPI.md#CloudAdminCreateSpendCap) | **Post** /v1/admin/spend-caps | Sets a usage cap on one org — a platform override of a customer budget, written to the customer&#39;s own spend-alert rows.
[**CloudAdminCustomer**](AdminAPI.md#CloudAdminCustomer) | **Get** /v1/admin/customers/{org} | Answers GET /v1/admin/customers/:org.
[**CloudAdminCustomers**](AdminAPI.md#CloudAdminCustomers) | **Get** /v1/admin/customers | Lists every customer org at a glance, sorted by slug: owner email, plan, suspend status, member count, balance, month-to-date spend and MRR.
[**CloudAdminDeleteDroplet**](AdminAPI.md#CloudAdminDeleteDroplet) | **Delete** /v1/admin/infra/droplets/{id} | Destroys a droplet the board has just proven is NOT a DOKS node.
[**CloudAdminDeleteLoadBalancer**](AdminAPI.md#CloudAdminDeleteLoadBalancer) | **Delete** /v1/admin/infra/loadbalancers/{id} | Destroys a load balancer the board has just proven no live type&#x3D;LoadBalancer Service in any cluster targets.
[**CloudAdminDeleteSpendCap**](AdminAPI.md#CloudAdminDeleteSpendCap) | **Delete** /v1/admin/spend-caps/{id} | Removes one cap by id, lifting the ceiling entirely.
[**CloudAdminDeleteVolume**](AdminAPI.md#CloudAdminDeleteVolume) | **Delete** /v1/admin/infra/volumes/{id} | Destroys a volume the board has just proven no PersistentVolume in any cluster references.
[**CloudAdminDisablePlugin**](AdminAPI.md#CloudAdminDisablePlugin) | **Post** /v1/admin/plugins/{name}/disable | Stops the plugin.
[**CloudAdminEnablePlugin**](AdminAPI.md#CloudAdminEnablePlugin) | **Post** /v1/admin/plugins/{name}/enable | Brings a stopped or disabled plugin back on the artifact it already has: the zero Plugin names no new artifact, so Reload reuses the loaded spec and clears the disabled flag.
[**CloudAdminFinance**](AdminAPI.md#CloudAdminFinance) | **Get** /v1/admin/finance | Answers GET /v1/admin/finance.
[**CloudAdminFinanceBackfill**](AdminAPI.md#CloudAdminFinanceBackfill) | **Post** /v1/admin/finance/backfill | Carries ONE org&#39;s current commerce prepaid balance into the native finance wallet — the one-time cutover between the two ledgers.
[**CloudAdminFlags**](AdminAPI.md#CloudAdminFlags) | **Get** /v1/admin/flags | Reads the platform control-plane board: every runtime launch/release switch (waitlist, public signup, subsystem activation, gateway limits, network ids) with its LIVE value and where that value came from — a stored definition or the compiled-in default.
[**CloudAdminGrantCredit**](AdminAPI.md#CloudAdminGrantCredit) | **Post** /v1/admin/customers/{org}/credit | Issues a staff credit grant to the org named in the path — a comp, refund or promo — through the ONE credit-write path core.ApplyGrant, which validates the amount against the per-grant cap, checks the org exists, moves the money and records the tamper-evident audit row.
[**CloudAdminGrants**](AdminAPI.md#CloudAdminGrants) | **Get** /v1/admin/grants | Reads the credit-grant ledger across ALL orgs, newest first — who granted what to whom, when, and from which money bucket.
[**CloudAdminInfra**](AdminAPI.md#CloudAdminInfra) | **Get** /v1/admin/infra | Serves the whole DigitalOcean infrastructure board: droplets, volumes, DOKS clusters and load balancers, each cross-referenced against every cluster&#39;s live Kubernetes state so the board can say what is safe to destroy and what is not.
[**CloudAdminInvoices**](AdminAPI.md#CloudAdminInvoices) | **Get** /v1/admin/invoices | Answers GET /v1/admin/invoices.
[**CloudAdminIssueGrant**](AdminAPI.md#CloudAdminIssueGrant) | **Post** /v1/admin/grants | Issues a credit grant to any org from the operator Grants view, with the target named in the body.
[**CloudAdminMe**](AdminAPI.md#CloudAdminMe) | **Get** /v1/admin/me | Answers with the validated operator identity — who the console is signed in as, which tier they are, and how wide their tenant window is.
[**CloudAdminMetrics**](AdminAPI.md#CloudAdminMetrics) | **Get** /v1/admin/metrics | Answers GET /v1/admin/metrics by aggregating commerce.events directly (fleet-wide, no per-org fan-out).
[**CloudAdminMoney**](AdminAPI.md#CloudAdminMoney) | **Get** /v1/admin/money | moneyBoardHandler answers GET /v1/admin/money.
[**CloudAdminO11y**](AdminAPI.md#CloudAdminO11y) | **Get** /v1/admin/o11y | Is the fleet-wide observability board: LLM usage (requests, tokens, cost, errors, top orgs, top models), trace RED metrics (count, p50/p95/p99 latency in ms, error rate, top services), fleet log volume, and the O11yAI generation rollup — all aggregated across EVERY tenant, with no org filter applied.
[**CloudAdminOrgs**](AdminAPI.md#CloudAdminOrgs) | **Get** /v1/admin/orgs | Lists the tenant directory one row per org, sorted by slug: member count and the org&#39;s month-to-date spend and credit balance, read live from IAM and commerce.
[**CloudAdminOverview**](AdminAPI.md#CloudAdminOverview) | **Get** /v1/admin/overview | Is the Platform Overview tiles: how many orgs and users are in the caller&#39;s tenant window, the fleet workload counts, and month-to-date spend and credits.
[**CloudAdminPlugins**](AdminAPI.md#CloudAdminPlugins) | **Get** /v1/admin/plugins | Reports what each host is actually running: every loaded plugin with its version, pid, uptime, reload and restart counts, and its measured CPU, RSS, thread and fd cost — read from the kernel, which is only answerable at all because a plugin is a process.
[**CloudAdminProducts**](AdminAPI.md#CloudAdminProducts) | **Get** /v1/admin/products | Lists the fleet workload registry: every operator App CR across the platform namespaces with its declared vs running image tag, reconciled health/phase and drift verdict.
[**CloudAdminPromo**](AdminAPI.md#CloudAdminPromo) | **Get** /v1/admin/promos | Reads the current platform plan promo — the singleton discount offer, e.g.
[**CloudAdminProvidersCredit**](AdminAPI.md#CloudAdminProvidersCredit) | **Get** /v1/admin/providers/credit | Serves GET /v1/admin/providers/credit — the per-provider upstream credit ledger.
[**CloudAdminReactivateCustomer**](AdminAPI.md#CloudAdminReactivateCustomer) | **Post** /v1/admin/customers/{org}/reactivate | Restores access for every member of the org, undoing a suspend.
[**CloudAdminReloadPlugin**](AdminAPI.md#CloudAdminReloadPlugin) | **Post** /v1/admin/plugins/{name}/reload | Swaps a plugin for another build without dropping a request.
[**CloudAdminResizeDroplet**](AdminAPI.md#CloudAdminResizeDroplet) | **Post** /v1/admin/infra/droplets/{id}/resize | Changes a droplet&#39;s plan.
[**CloudAdminResizeVolume**](AdminAPI.md#CloudAdminResizeVolume) | **Post** /v1/admin/infra/volumes/{id}/resize | Grows a volume.
[**CloudAdminRevenue**](AdminAPI.md#CloudAdminRevenue) | **Get** /v1/admin/revenue | Is the fleet money board: total prepaid balances held, total realized spend, MRR, ARPU, a per-customer table sorted highest-revenue first, and a real 30-day spend trend from the usage ledger.
[**CloudAdminRoles**](AdminAPI.md#CloudAdminRoles) | **Get** /v1/admin/roles | Lists IAM roles for one owner org, forwarded VERBATIM from IAM&#39;s get-roles.
[**CloudAdminScaleNodePool**](AdminAPI.md#CloudAdminScaleNodePool) | **Post** /v1/admin/infra/clusters/{id}/nodepools/{pool}/scale | Sets a node pool&#39;s node count — the ONE correct way to change how many nodes a DOKS cluster has.
[**CloudAdminServices**](AdminAPI.md#CloudAdminServices) | **Get** /v1/admin/services | Reads the launch board: every hosted service in the registry with its LIVE waitlist mode, evaluated through the flag engine.
[**CloudAdminSetFlag**](AdminAPI.md#CloudAdminSetFlag) | **Put** /v1/admin/flags/{key} | Stores or overwrites ONE platform switch&#39;s definition and answers with the whole board as it now stands.
[**CloudAdminSetPromo**](AdminAPI.md#CloudAdminSetPromo) | **Put** /v1/admin/promos | Upserts the platform plan promo — the ONE place the offer is configured.
[**CloudAdminSetServiceMode**](AdminAPI.md#CloudAdminSetServiceMode) | **Post** /v1/admin/services/{service}/mode | Flips ONE service&#39;s waitlist switch — the launch lever.
[**CloudAdminSnapshotVolume**](AdminAPI.md#CloudAdminSnapshotVolume) | **Post** /v1/admin/infra/volumes/{id}/snapshot | Takes a point-in-time snapshot of one volume — the undo a delete relies on, available on its own so an operator can take one before any risky change.
[**CloudAdminSpendCaps**](AdminAPI.md#CloudAdminSpendCaps) | **Get** /v1/admin/spend-caps | Reads one org&#39;s usage caps: its spend alerts plus the derived period spend, over/warn state and reset time.
[**CloudAdminSubscriptions**](AdminAPI.md#CloudAdminSubscriptions) | **Get** /v1/admin/subscriptions | Answers GET /v1/admin/subscriptions.
[**CloudAdminSubsystems**](AdminAPI.md#CloudAdminSubsystems) | **Get** /v1/admin/subsystems | subsystems answers GET /v1/admin/subsystems.
[**CloudAdminSuspendCustomer**](AdminAPI.md#CloudAdminSuspendCustomer) | **Post** /v1/admin/customers/{org}/suspend | Cuts off every member of the org: IAM refuses a forbidden user at login AND at token issuance, so a suspended customer can neither sign in nor mint a fresh token.
[**CloudAdminSync**](AdminAPI.md#CloudAdminSync) | **Post** /v1/admin/sync | Answers the operator&#39;s \&quot;Sync now\&quot; button.
[**CloudAdminUpdateSpendCap**](AdminAPI.md#CloudAdminUpdateSpendCap) | **Patch** /v1/admin/spend-caps/{id} | Edits one cap by id — raise or lower the ceiling, flip enforcement.
[**CloudAdminUpsertService**](AdminAPI.md#CloudAdminUpsertService) | **Post** /v1/admin/services | Onboards a hosted service, or edits one, so a new host comes under the launch gate WITHOUT a redeploy.
[**CloudAdminUsage**](AdminAPI.md#CloudAdminUsage) | **Get** /v1/admin/usage | Returns the month-to-date money totals: one org&#39;s when org names one, else the fleet sum across every org a SuperAdmin can see.
[**CloudAdminUsageFunding**](AdminAPI.md#CloudAdminUsageFunding) | **Get** /v1/admin/usage/funding | Splits our upstream AI usage by how it was FUNDED: one row per (provider, model) over the window, tagged credit (provider grant still remaining), paid (grant exhausted) or paid_only (no grant at all).
[**CloudAdminUsers**](AdminAPI.md#CloudAdminUsers) | **Get** /v1/admin/users | Lists the user directory across the caller&#39;s tenant window, one page at a time.
[**CloudAdminWaitlist**](AdminAPI.md#CloudAdminWaitlist) | **Get** /v1/admin/waitlist | Reads one waitlist&#39;s leaderboard from the Hanzo waitlist engine — position, points and referral standing per entry — proxied server-authed with the engine secret, never a client credential.
[**CloudAdminWaitlistBoost**](AdminAPI.md#CloudAdminWaitlistBoost) | **Post** /v1/admin/waitlist/boost | Grants a user waitlist points, moving them up toward the access cutoff.
[**CloudGetV1AdminAffiliates**](AdminAPI.md#CloudGetV1AdminAffiliates) | **Get** /v1/admin/affiliates | 
[**CloudGetV1AdminAuthors**](AdminAPI.md#CloudGetV1AdminAuthors) | **Get** /v1/admin/authors | ListAuthors returns the platform&#39;s whole author program — every org&#39;s author record, not the caller&#39;s — with each one&#39;s repository and deploy counts and a fleet roll-up of the money accrued, pending and paid.
[**CloudGetV1AdminAuthorsIdBasis**](AdminAPI.md#CloudGetV1AdminAuthorsIdBasis) | **Get** /v1/admin/authors/{id}/basis | AuthorRoyaltyBasis returns the audit trail behind ONE author&#39;s royalty — the same payload the author reads at /v1/authors/basis, from the same builder, so support sees exactly what the author sees rather than a parallel view free to drift.
[**CloudGetV1AdminCatalog**](AdminAPI.md#CloudGetV1AdminCatalog) | **Get** /v1/admin/catalog | GetAdminCatalog returns the full model and provider catalog annotated with each entry&#39;s enablement state, for the operator console.
[**CloudGetV1AdminEnablement**](AdminAPI.md#CloudGetV1AdminEnablement) | **Get** /v1/admin/enablement | ListEnablement returns every item an operator has set an enablement state on — its global state (off, beta or ga) and the orgs granted its beta.
[**CloudGetV1AdminReferrals**](AdminAPI.md#CloudGetV1AdminReferrals) | **Get** /v1/admin/referrals | 
[**CloudGetV1AdminReferralsBonuses**](AdminAPI.md#CloudGetV1AdminReferralsBonuses) | **Get** /v1/admin/referrals/bonuses | Returns every one-time referral bonus in the ledger with a fleet summary.
[**CloudGetV1AdminTreasury**](AdminAPI.md#CloudGetV1AdminTreasury) | **Get** /v1/admin/treasury | GetAdminTreasury returns the whole treasury board for a SuperAdmin: the reserve fund report, the recent double-entry journal, and the Hanzo L1 anchor status of the ledger root.
[**CloudPatchV1AdminCatalogModelsByWildcard1**](AdminAPI.md#CloudPatchV1AdminCatalogModelsByWildcard1) | **Patch** /v1/admin/catalog/models/{wildcard1} | 
[**CloudPatchV1AdminCatalogProvidersName**](AdminAPI.md#CloudPatchV1AdminCatalogProvidersName) | **Patch** /v1/admin/catalog/providers/{name} | PatchProvider sets one provider&#39;s availability overlay.
[**CloudPostV1AdminAffiliatesByIdApprove**](AdminAPI.md#CloudPostV1AdminAffiliatesByIdApprove) | **Post** /v1/admin/affiliates/{id}/approve | 
[**CloudPostV1AdminAffiliatesByIdPayout**](AdminAPI.md#CloudPostV1AdminAffiliatesByIdPayout) | **Post** /v1/admin/affiliates/{id}/payout | 
[**CloudPostV1AdminAffiliatesByIdRate**](AdminAPI.md#CloudPostV1AdminAffiliatesByIdRate) | **Post** /v1/admin/affiliates/{id}/rate | 
[**CloudPostV1AdminAffiliatesByIdSuspend**](AdminAPI.md#CloudPostV1AdminAffiliatesByIdSuspend) | **Post** /v1/admin/affiliates/{id}/suspend | 
[**CloudPostV1AdminAffiliatesSweep**](AdminAPI.md#CloudPostV1AdminAffiliatesSweep) | **Post** /v1/admin/affiliates/sweep | 
[**CloudPostV1AdminAuthorsIdApprove**](AdminAPI.md#CloudPostV1AdminAuthorsIdApprove) | **Post** /v1/admin/authors/{id}/approve | ApproveAuthor admits one author to EARNING, optionally on a negotiated royalty share.
[**CloudPostV1AdminAuthorsIdPayout**](AdminAPI.md#CloudPostV1AdminAuthorsIdPayout) | **Post** /v1/admin/authors/{id}/payout | PayAuthor records a payout of accrued royalty and settles it.
[**CloudPostV1AdminAuthorsIdSuspend**](AdminAPI.md#CloudPostV1AdminAuthorsIdSuspend) | **Post** /v1/admin/authors/{id}/suspend | SuspendAuthor stops one author earning.
[**CloudPostV1AdminAuthorsSweep**](AdminAPI.md#CloudPostV1AdminAuthorsSweep) | **Post** /v1/admin/authors/sweep | SweepAuthorRoyalty runs the accrual sweep across every approved author: for each of their deploying orgs it computes this period&#39;s royalty from that org&#39;s metered spend and latches it at most once per period.
[**CloudPostV1AdminReferralsSweep**](AdminAPI.md#CloudPostV1AdminReferralsSweep) | **Post** /v1/admin/referrals/sweep | Qualify-checks every pending referral and grants the ones that now qualify.
[**CloudPostV1AdminTreasuryAnchor**](AdminAPI.md#CloudPostV1AdminTreasuryAnchor) | **Post** /v1/admin/treasury/anchor | AnchorTreasury commits the current ledger root to Hanzo L1, making the books tamper-evident on chain, and returns the anchoring status.
[**CloudPostV1AdminTreasuryBindAnchor**](AdminAPI.md#CloudPostV1AdminTreasuryBindAnchor) | **Post** /v1/admin/treasury/bind-anchor | BindTreasuryAnchorSigner makes the reserve&#39;s threshold MPC wallet the signer for on-chain anchors, and returns its EVM address so an operator can fund it for gas.
[**CloudPostV1AdminTreasuryPolicy**](AdminAPI.md#CloudPostV1AdminTreasuryPolicy) | **Post** /v1/admin/treasury/policy | SetTreasuryPolicy sets the revenue-share basis points a sweep accrues into the reserve fund and returns the stored policy.
[**CloudPostV1AdminTreasurySeed**](AdminAPI.md#CloudPostV1AdminTreasurySeed) | **Post** /v1/admin/treasury/seed | SeedTreasury injects bootstrap capital into the reserve fund so backed payouts can begin before the first revenue-share sweep, and returns the journal entry it wrote.
[**CloudPostV1AdminTreasurySweep**](AdminAPI.md#CloudPostV1AdminTreasurySweep) | **Post** /v1/admin/treasury/sweep | SweepTreasury posts the revenue-share accrual for one period — revenue into the reserve fund, at the current policy&#39;s basis points — and returns what it moved.
[**CloudPutV1AdminEnablement**](AdminAPI.md#CloudPutV1AdminEnablement) | **Put** /v1/admin/enablement | SetEnablement sets one item&#39;s global enablement state — off, beta or ga — and optionally replaces the list of orgs granted its beta.
[**S3AdminInfo**](AdminAPI.md#S3AdminInfo) | **Get** /v1/s3/admin/info | Server information
[**S3AdminUsage**](AdminAPI.md#S3AdminUsage) | **Get** /v1/s3/admin/usage | Storage usage
[**S3CreateServiceAccount**](AdminAPI.md#S3CreateServiceAccount) | **Post** /v1/s3/admin/service-accounts | Create a service account
[**S3ListServiceAccounts**](AdminAPI.md#S3ListServiceAccounts) | **Get** /v1/s3/admin/service-accounts | List service accounts



## AnalyticsAdminListUsers

> []AnalyticsAdminListUsers200ResponseInner AnalyticsAdminListUsers(ctx).Page(page).PageSize(pageSize).OrderBy(orderBy).Search(search).Execute()

List all users (admin only)

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
	page := int32(56) // int32 |  (optional)
	pageSize := int32(56) // int32 |  (optional)
	orderBy := "orderBy_example" // string |  (optional)
	search := "search_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminAPI.AnalyticsAdminListUsers(context.Background()).Page(page).PageSize(pageSize).OrderBy(orderBy).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.AnalyticsAdminListUsers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AnalyticsAdminListUsers`: []AnalyticsAdminListUsers200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.AnalyticsAdminListUsers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAnalyticsAdminListUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** |  | 
 **pageSize** | **int32** |  | 
 **orderBy** | **string** |  | 
 **search** | **string** |  | 

### Return type

[**[]AnalyticsAdminListUsers200ResponseInner**](AnalyticsAdminListUsers200ResponseInner.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AnalyticsAdminListWebsites

> []AnalyticsWebsite AnalyticsAdminListWebsites(ctx).UserId(userId).IncludeOwnedTeams(includeOwnedTeams).IncludeAllTeams(includeAllTeams).Page(page).PageSize(pageSize).OrderBy(orderBy).Search(search).Execute()

List all websites for a user (admin only)

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
	userId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	includeOwnedTeams := "includeOwnedTeams_example" // string |  (optional)
	includeAllTeams := "includeAllTeams_example" // string |  (optional)
	page := int32(56) // int32 |  (optional)
	pageSize := int32(56) // int32 |  (optional)
	orderBy := "orderBy_example" // string |  (optional)
	search := "search_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminAPI.AnalyticsAdminListWebsites(context.Background()).UserId(userId).IncludeOwnedTeams(includeOwnedTeams).IncludeAllTeams(includeAllTeams).Page(page).PageSize(pageSize).OrderBy(orderBy).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.AnalyticsAdminListWebsites``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AnalyticsAdminListWebsites`: []AnalyticsWebsite
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.AnalyticsAdminListWebsites`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAnalyticsAdminListWebsitesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userId** | **string** |  | 
 **includeOwnedTeams** | **string** |  | 
 **includeAllTeams** | **string** |  | 
 **page** | **int32** |  | 
 **pageSize** | **int32** |  | 
 **orderBy** | **string** |  | 
 **search** | **string** |  | 

### Return type

[**[]AnalyticsWebsite**](AnalyticsWebsite.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudAdminAIMetrics

> CloudAimetricsOut CloudAdminAIMetrics(ctx).Range_(range_).Execute()

Is the fleet AI board: O11yAI generations (count, cost, avg/p95 latency, per-model), per-model usage from the live cloud_usage ledger, and the eval plane (traces, scores, score names, runs, and the average-score trend).



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
	range_ := "7d" // string | Range is the lower time bound: 24h, 7d or 30d. Anything else reads as the board's own default. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminAPI.CloudAdminAIMetrics(context.Background()).Range_(range_).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.CloudAdminAIMetrics``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudAdminAIMetrics`: CloudAimetricsOut
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.CloudAdminAIMetrics`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudAdminAIMetricsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **range_** | **string** | Range is the lower time bound: 24h, 7d or 30d. Anything else reads as the board&#39;s own default. | 

### Return type

[**CloudAimetricsOut**](CloudAimetricsOut.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudAdminAnalytics

> CloudAnalyticsOut CloudAdminAnalytics(ctx).Range_(range_).Execute()

Is the SaaS product-analytics board over the caller's tenant window: active customers, new and churned, retention, MRR, ARPU, the usage trend and the top customers by spend — every number folded from the commerce ledger, not sampled.



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
	range_ := "30d" // string | Range is the lower time bound: 24h, 7d or 30d. Anything else reads as the board's own default. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminAPI.CloudAdminAnalytics(context.Background()).Range_(range_).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.CloudAdminAnalytics``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudAdminAnalytics`: CloudAnalyticsOut
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.CloudAdminAnalytics`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudAdminAnalyticsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **range_** | **string** | Range is the lower time bound: 24h, 7d or 30d. Anything else reads as the board&#39;s own default. | 

### Return type

[**CloudAnalyticsOut**](CloudAnalyticsOut.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudAdminApplications

> CloudIamRowsOut CloudAdminApplications(ctx).Owner(owner).P(p).PageSize(pageSize).Execute()

Lists IAM applications for one owner org, forwarded VERBATIM from IAM's get-applications.



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
	owner := "admin" // string | Owner is the org whose rows to read. Defaults to the admin org, which owns the platform's roles and applications. (optional)
	p := "1" // string | Page is the 1-based page number. Forwarded only when set — IAM applies its own default otherwise. (optional)
	pageSize := "50" // string | PageSize is rows per page. Forwarded only when set. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminAPI.CloudAdminApplications(context.Background()).Owner(owner).P(p).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.CloudAdminApplications``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudAdminApplications`: CloudIamRowsOut
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.CloudAdminApplications`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudAdminApplicationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **owner** | **string** | Owner is the org whose rows to read. Defaults to the admin org, which owns the platform&#39;s roles and applications. | 
 **p** | **string** | Page is the 1-based page number. Forwarded only when set — IAM applies its own default otherwise. | 
 **pageSize** | **string** | PageSize is rows per page. Forwarded only when set. | 

### Return type

[**CloudIamRowsOut**](CloudIamRowsOut.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudAdminAudit

> CloudRecordsOut CloudAdminAudit(ctx).Org(org).Sub(sub).Action(action).Resource(resource).ResourceId(resourceId).Result(result).Since(since).Until(until).PageSize(pageSize).P(p).Execute()

Reads cloud's tamper-evident audit trail, newest first, with the chain's live integrity attached so a listing can be badged as verified.



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
	org := "acme" // string | Org restricts the trail to one tenant. (optional)
	sub := "sub_example" // string | Sub restricts it to one actor (the validated subject that made the request). (optional)
	action := "admin.waitlist.grant" // string | Action restricts it to one action name, e.g. \"admin.waitlist.grant\". (optional)
	resource := "resource_example" // string | Resource restricts it to one resource kind, e.g. \"credit-grant\". (optional)
	resourceId := "resourceId_example" // string | ResourceID restricts it to one resource instance. (optional)
	result := "result_example" // string | Result restricts it to \"success\" or \"error\". (optional)
	since := "2026-07-01T00:00:00Z" // string | Since is the inclusive lower time bound, RFC3339. An unparseable value is ignored rather than refused — one malformed filter must not hide the trail. (optional)
	until := "until_example" // string | Until is the upper time bound, RFC3339, with the same tolerance. (optional)
	pageSize := "50" // string | PageSize is rows per page, default 100. (optional)
	p := "p_example" // string | Page is the 1-based page number, driving the offset. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminAPI.CloudAdminAudit(context.Background()).Org(org).Sub(sub).Action(action).Resource(resource).ResourceId(resourceId).Result(result).Since(since).Until(until).PageSize(pageSize).P(p).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.CloudAdminAudit``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudAdminAudit`: CloudRecordsOut
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.CloudAdminAudit`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudAdminAuditRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **org** | **string** | Org restricts the trail to one tenant. | 
 **sub** | **string** | Sub restricts it to one actor (the validated subject that made the request). | 
 **action** | **string** | Action restricts it to one action name, e.g. \&quot;admin.waitlist.grant\&quot;. | 
 **resource** | **string** | Resource restricts it to one resource kind, e.g. \&quot;credit-grant\&quot;. | 
 **resourceId** | **string** | ResourceID restricts it to one resource instance. | 
 **result** | **string** | Result restricts it to \&quot;success\&quot; or \&quot;error\&quot;. | 
 **since** | **string** | Since is the inclusive lower time bound, RFC3339. An unparseable value is ignored rather than refused — one malformed filter must not hide the trail. | 
 **until** | **string** | Until is the upper time bound, RFC3339, with the same tolerance. | 
 **pageSize** | **string** | PageSize is rows per page, default 100. | 
 **p** | **string** | Page is the 1-based page number, driving the offset. | 

### Return type

[**CloudRecordsOut**](CloudRecordsOut.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudAdminAuditVerify

> CloudVerifyOut CloudAdminAuditVerify(ctx).Execute()

Walks the WHOLE hash chain and reports whether it is intact: how many records were checked, the head hash to pin externally against tail-truncation, and — when the chain is broken — the seq of the first bad record and why.



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
	resp, r, err := apiClient.AdminAPI.CloudAdminAuditVerify(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.CloudAdminAuditVerify``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudAdminAuditVerify`: CloudVerifyOut
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.CloudAdminAuditVerify`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudAdminAuditVerifyRequest struct via the builder pattern


### Return type

[**CloudVerifyOut**](CloudVerifyOut.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudAdminBases

> CloudBasesOut CloudAdminBases(ctx).Execute()

Lists the tenant Base instances in the caller's window — a SuperAdmin sees every tenant's, anyone else only their own subtree's.



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
	resp, r, err := apiClient.AdminAPI.CloudAdminBases(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.CloudAdminBases``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudAdminBases`: CloudBasesOut
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.CloudAdminBases`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudAdminBasesRequest struct via the builder pattern


### Return type

[**CloudBasesOut**](CloudBasesOut.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudAdminBlockStorage

> CloudBlockStorageOut CloudAdminBlockStorage(ctx).Execute()

Is the realtime block-storage board: the DigitalOcean volume fleet (count, capacity, monthly list cost, per-volume region and attachment) plus the analytics datastore's OWN fill, read from its system.disks.



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
	resp, r, err := apiClient.AdminAPI.CloudAdminBlockStorage(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.CloudAdminBlockStorage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudAdminBlockStorage`: CloudBlockStorageOut
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.CloudAdminBlockStorage`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudAdminBlockStorageRequest struct via the builder pattern


### Return type

[**CloudBlockStorageOut**](CloudBlockStorageOut.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudAdminCompute

> CloudComputeOut CloudAdminCompute(ctx).Kind(kind).Org(org).Range_(range_).Execute()

Rolls the fleet's compute usage up to one row per (org, app, project, kind): how many distinct machines ran in the window, how many are still active, what they billed, and when each group last emitted an event.



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
	kind := "bot" // string | Kind narrows to one workload class (bot | machine | cluster | nodepool | container | function | …). An OPEN spectrum matched as a plain string, lowercased to the warehouse's convention; empty means every kind. (optional)
	org := "acme" // string | Org narrows to one tenant. Empty means every tenant — this board is cross-tenant by nature. (optional)
	range_ := "7d" // string | Range is the lower time bound: 24h, 7d or 30d. Anything else reads as 30d. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminAPI.CloudAdminCompute(context.Background()).Kind(kind).Org(org).Range_(range_).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.CloudAdminCompute``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudAdminCompute`: CloudComputeOut
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.CloudAdminCompute`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudAdminComputeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **kind** | **string** | Kind narrows to one workload class (bot | machine | cluster | nodepool | container | function | …). An OPEN spectrum matched as a plain string, lowercased to the warehouse&#39;s convention; empty means every kind. | 
 **org** | **string** | Org narrows to one tenant. Empty means every tenant — this board is cross-tenant by nature. | 
 **range_** | **string** | Range is the lower time bound: 24h, 7d or 30d. Anything else reads as 30d. | 

### Return type

[**CloudComputeOut**](CloudComputeOut.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudAdminCordonNode

> CloudMutationOut CloudAdminCordonNode(ctx, id).CloudCordonIn(cloudCordonIn).Execute()

Marks one cluster node unschedulable — or schedulable again — and can drain the pods already on it.



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
	id := "id_example" // string | ID is the node's droplet id, from the path.
	cloudCordonIn := *openapiclient.NewCloudCordonIn() // CloudCordonIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminAPI.CloudAdminCordonNode(context.Background(), id).CloudCordonIn(cloudCordonIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.CloudAdminCordonNode``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudAdminCordonNode`: CloudMutationOut
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.CloudAdminCordonNode`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the node&#39;s droplet id, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudAdminCordonNodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cloudCordonIn** | [**CloudCordonIn**](CloudCordonIn.md) |  | 

### Return type

[**CloudMutationOut**](CloudMutationOut.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudAdminCreateCreditGrant

> CloudRawOut CloudAdminCreateCreditGrant(ctx).RequestBody(requestBody).Execute()

Mints credit for one org.



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
	requestBody := map[string]map[string]interface{}{"key": map[string]interface{}(123)} // map[string]map[string]interface{} | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminAPI.CloudAdminCreateCreditGrant(context.Background()).RequestBody(requestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.CloudAdminCreateCreditGrant``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudAdminCreateCreditGrant`: CloudRawOut
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.CloudAdminCreateCreditGrant`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudAdminCreateCreditGrantRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestBody** | **map[string]map[string]interface{}** |  | 

### Return type

[**CloudRawOut**](CloudRawOut.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudAdminCreateSpendCap

> CloudRawOut CloudAdminCreateSpendCap(ctx).CloudCapIn(cloudCapIn).Execute()

Sets a usage cap on one org — a platform override of a customer budget, written to the customer's own spend-alert rows.



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
	cloudCapIn := *openapiclient.NewCloudCapIn() // CloudCapIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminAPI.CloudAdminCreateSpendCap(context.Background()).CloudCapIn(cloudCapIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.CloudAdminCreateSpendCap``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudAdminCreateSpendCap`: CloudRawOut
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.CloudAdminCreateSpendCap`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudAdminCreateSpendCapRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cloudCapIn** | [**CloudCapIn**](CloudCapIn.md) |  | 

### Return type

[**CloudRawOut**](CloudRawOut.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudAdminCustomer

> CloudCustomerDetailOut CloudAdminCustomer(ctx, org).Execute()

Answers GET /v1/admin/customers/:org.



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
	org := "org_example" // string | Org is the tenant slug from the path.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminAPI.CloudAdminCustomer(context.Background(), org).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.CloudAdminCustomer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudAdminCustomer`: CloudCustomerDetailOut
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.CloudAdminCustomer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**org** | **string** | Org is the tenant slug from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudAdminCustomerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CloudCustomerDetailOut**](CloudCustomerDetailOut.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudAdminCustomers

> CloudCustomersOut CloudAdminCustomers(ctx).Execute()

Lists every customer org at a glance, sorted by slug: owner email, plan, suspend status, member count, balance, month-to-date spend and MRR.



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
	resp, r, err := apiClient.AdminAPI.CloudAdminCustomers(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.CloudAdminCustomers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudAdminCustomers`: CloudCustomersOut
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.CloudAdminCustomers`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudAdminCustomersRequest struct via the builder pattern


### Return type

[**CloudCustomersOut**](CloudCustomersOut.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudAdminDeleteDroplet

> CloudMutationOut CloudAdminDeleteDroplet(ctx, id).Size(size).Disk(disk).Execute()

Destroys a droplet the board has just proven is NOT a DOKS node.



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
	id := "id_example" // string | ID is the DO droplet id, from the path. Numeric.
	size := "size_example" // string | Size is the target DigitalOcean size slug on resize, e.g. \"s-4vcpu-8gb\". (optional)
	disk := true // bool | Disk requests a PERMANENT resize that grows the disk. DO can never resize such a droplet down again, so it defaults false — a CPU/RAM-only change, reversible. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminAPI.CloudAdminDeleteDroplet(context.Background(), id).Size(size).Disk(disk).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.CloudAdminDeleteDroplet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudAdminDeleteDroplet`: CloudMutationOut
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.CloudAdminDeleteDroplet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the DO droplet id, from the path. Numeric. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudAdminDeleteDropletRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **size** | **string** | Size is the target DigitalOcean size slug on resize, e.g. \&quot;s-4vcpu-8gb\&quot;. | 
 **disk** | **bool** | Disk requests a PERMANENT resize that grows the disk. DO can never resize such a droplet down again, so it defaults false — a CPU/RAM-only change, reversible. | 

### Return type

[**CloudMutationOut**](CloudMutationOut.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudAdminDeleteLoadBalancer

> CloudMutationOut CloudAdminDeleteLoadBalancer(ctx, id).Execute()

Destroys a load balancer the board has just proven no live type=LoadBalancer Service in any cluster targets.



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
	id := "id_example" // string | ID is the DO load balancer id, from the path.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminAPI.CloudAdminDeleteLoadBalancer(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.CloudAdminDeleteLoadBalancer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudAdminDeleteLoadBalancer`: CloudMutationOut
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.CloudAdminDeleteLoadBalancer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the DO load balancer id, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudAdminDeleteLoadBalancerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CloudMutationOut**](CloudMutationOut.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudAdminDeleteSpendCap

> CloudRawOut CloudAdminDeleteSpendCap(ctx, id).Org(org).Execute()

Removes one cap by id, lifting the ceiling entirely.



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
	id := "cap_1" // string | ID is the cap to edit or remove, from the path. Unused by the list and create ops.
	org := "acme" // string | Org is the tenant to act on. Required for a SuperAdmin — they must name their target; ignored for a white-label admin, who always acts on their own org. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminAPI.CloudAdminDeleteSpendCap(context.Background(), id).Org(org).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.CloudAdminDeleteSpendCap``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudAdminDeleteSpendCap`: CloudRawOut
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.CloudAdminDeleteSpendCap`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the cap to edit or remove, from the path. Unused by the list and create ops. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudAdminDeleteSpendCapRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **org** | **string** | Org is the tenant to act on. Required for a SuperAdmin — they must name their target; ignored for a white-label admin, who always acts on their own org. | 

### Return type

[**CloudRawOut**](CloudRawOut.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudAdminDeleteVolume

> CloudMutationOut CloudAdminDeleteVolume(ctx, id).Snapshot(snapshot).Name(name).SizeGiB(sizeGiB).Execute()

Destroys a volume the board has just proven no PersistentVolume in any cluster references.



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
	id := "id_example" // string | ID is the DO volume id, from the path.
	snapshot := "snapshot_example" // string | Snapshot is the snapshot-first switch on DELETE. Anything other than the literal \"false\" snapshots before destroying — the snapshot IS the undo, so waiving it is deliberate and explicit. (optional)
	name := "name_example" // string | Name is the snapshot name on the snapshot action. Blank gets a deterministic \"<volume>-predelete-<unix>\" so the undo is findable in the DO console. (optional)
	sizeGiB := int32(56) // int32 | SizeGiB is the target size on the resize action. A volume only ever grows — ExpandTo is the verdict that refuses a shrink, so this is not validated here. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminAPI.CloudAdminDeleteVolume(context.Background(), id).Snapshot(snapshot).Name(name).SizeGiB(sizeGiB).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.CloudAdminDeleteVolume``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudAdminDeleteVolume`: CloudMutationOut
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.CloudAdminDeleteVolume`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the DO volume id, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudAdminDeleteVolumeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **snapshot** | **string** | Snapshot is the snapshot-first switch on DELETE. Anything other than the literal \&quot;false\&quot; snapshots before destroying — the snapshot IS the undo, so waiving it is deliberate and explicit. | 
 **name** | **string** | Name is the snapshot name on the snapshot action. Blank gets a deterministic \&quot;&lt;volume&gt;-predelete-&lt;unix&gt;\&quot; so the undo is findable in the DO console. | 
 **sizeGiB** | **int32** | SizeGiB is the target size on the resize action. A volume only ever grows — ExpandTo is the verdict that refuses a shrink, so this is not validated here. | 

### Return type

[**CloudMutationOut**](CloudMutationOut.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudAdminDisablePlugin

> CloudActionOut CloudAdminDisablePlugin(ctx, name).CloudNameIn(cloudNameIn).Execute()

Stops the plugin.



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
	name := "billing" // string | Name is the app, from the path.
	cloudNameIn := *openapiclient.NewCloudNameIn() // CloudNameIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminAPI.CloudAdminDisablePlugin(context.Background(), name).CloudNameIn(cloudNameIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.CloudAdminDisablePlugin``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudAdminDisablePlugin`: CloudActionOut
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.CloudAdminDisablePlugin`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name is the app, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudAdminDisablePluginRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cloudNameIn** | [**CloudNameIn**](CloudNameIn.md) |  | 

### Return type

[**CloudActionOut**](CloudActionOut.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudAdminEnablePlugin

> CloudActionOut CloudAdminEnablePlugin(ctx, name).CloudNameIn(cloudNameIn).Execute()

Brings a stopped or disabled plugin back on the artifact it already has: the zero Plugin names no new artifact, so Reload reuses the loaded spec and clears the disabled flag.



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
	name := "billing" // string | Name is the app, from the path.
	cloudNameIn := *openapiclient.NewCloudNameIn() // CloudNameIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminAPI.CloudAdminEnablePlugin(context.Background(), name).CloudNameIn(cloudNameIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.CloudAdminEnablePlugin``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudAdminEnablePlugin`: CloudActionOut
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.CloudAdminEnablePlugin`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name is the app, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudAdminEnablePluginRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cloudNameIn** | [**CloudNameIn**](CloudNameIn.md) |  | 

### Return type

[**CloudActionOut**](CloudActionOut.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudAdminFinance

> CloudFinanceOut CloudAdminFinance(ctx).Execute()

Answers GET /v1/admin/finance.



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
	resp, r, err := apiClient.AdminAPI.CloudAdminFinance(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.CloudAdminFinance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudAdminFinance`: CloudFinanceOut
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.CloudAdminFinance`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudAdminFinanceRequest struct via the builder pattern


### Return type

[**CloudFinanceOut**](CloudFinanceOut.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudAdminFinanceBackfill

> CloudBackfillOut CloudAdminFinanceBackfill(ctx).CloudBackfillIn(cloudBackfillIn).Execute()

Carries ONE org's current commerce prepaid balance into the native finance wallet — the one-time cutover between the two ledgers.



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
	cloudBackfillIn := *openapiclient.NewCloudBackfillIn() // CloudBackfillIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminAPI.CloudAdminFinanceBackfill(context.Background()).CloudBackfillIn(cloudBackfillIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.CloudAdminFinanceBackfill``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudAdminFinanceBackfill`: CloudBackfillOut
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.CloudAdminFinanceBackfill`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudAdminFinanceBackfillRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cloudBackfillIn** | [**CloudBackfillIn**](CloudBackfillIn.md) |  | 

### Return type

[**CloudBackfillOut**](CloudBackfillOut.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudAdminFlags

> CloudFlagsOut CloudAdminFlags(ctx).Execute()

Reads the platform control-plane board: every runtime launch/release switch (waitlist, public signup, subsystem activation, gateway limits, network ids) with its LIVE value and where that value came from — a stored definition or the compiled-in default.



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
	resp, r, err := apiClient.AdminAPI.CloudAdminFlags(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.CloudAdminFlags``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudAdminFlags`: CloudFlagsOut
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.CloudAdminFlags`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudAdminFlagsRequest struct via the builder pattern


### Return type

[**CloudFlagsOut**](CloudFlagsOut.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudAdminGrantCredit

> CloudGrantOut CloudAdminGrantCredit(ctx, org).CloudGrantIn(cloudGrantIn).Execute()

Issues a staff credit grant to the org named in the path — a comp, refund or promo — through the ONE credit-write path core.ApplyGrant, which validates the amount against the per-grant cap, checks the org exists, moves the money and records the tamper-evident audit row.



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
	org := "org_example" // string | Org is the tenant to credit. Required.
	cloudGrantIn := *openapiclient.NewCloudGrantIn() // CloudGrantIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminAPI.CloudAdminGrantCredit(context.Background(), org).CloudGrantIn(cloudGrantIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.CloudAdminGrantCredit``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudAdminGrantCredit`: CloudGrantOut
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.CloudAdminGrantCredit`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**org** | **string** | Org is the tenant to credit. Required. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudAdminGrantCreditRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cloudGrantIn** | [**CloudGrantIn**](CloudGrantIn.md) |  | 

### Return type

[**CloudGrantOut**](CloudGrantOut.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudAdminGrants

> CloudGrantsOut CloudAdminGrants(ctx).Org(org).Result(result).Limit(limit).Execute()

Reads the credit-grant ledger across ALL orgs, newest first — who granted what to whom, when, and from which money bucket.



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
	org := "org_example" // string | Org filters by the ACTOR's org (the staff org that issued the grant), which is rarely what a reader wants — the target org is a row field, not a filter. (optional)
	result := "success" // string | Result filters by outcome: \"success\" or \"error\". Empty returns both, which is the point of this view — a refused grant is as interesting as a granted one. (optional)
	limit := "50" // string | Limit caps the rows returned. Default 200. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminAPI.CloudAdminGrants(context.Background()).Org(org).Result(result).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.CloudAdminGrants``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudAdminGrants`: CloudGrantsOut
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.CloudAdminGrants`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudAdminGrantsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **org** | **string** | Org filters by the ACTOR&#39;s org (the staff org that issued the grant), which is rarely what a reader wants — the target org is a row field, not a filter. | 
 **result** | **string** | Result filters by outcome: \&quot;success\&quot; or \&quot;error\&quot;. Empty returns both, which is the point of this view — a refused grant is as interesting as a granted one. | 
 **limit** | **string** | Limit caps the rows returned. Default 200. | 

### Return type

[**CloudGrantsOut**](CloudGrantsOut.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudAdminInfra

> CloudReadOut CloudAdminInfra(ctx).Refresh(refresh).Execute()

Serves the whole DigitalOcean infrastructure board: droplets, volumes, DOKS clusters and load balancers, each cross-referenced against every cluster's live Kubernetes state so the board can say what is safe to destroy and what is not.



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
	refresh := "1" // string | Refresh, when present, forces a full re-scan instead of serving the cached snapshot. Every MUTATION re-scans regardless — this is only for the reader. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminAPI.CloudAdminInfra(context.Background()).Refresh(refresh).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.CloudAdminInfra``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudAdminInfra`: CloudReadOut
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.CloudAdminInfra`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudAdminInfraRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **refresh** | **string** | Refresh, when present, forces a full re-scan instead of serving the cached snapshot. Every MUTATION re-scans regardless — this is only for the reader. | 

### Return type

[**CloudReadOut**](CloudReadOut.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudAdminInvoices

> CloudInvoicesOut CloudAdminInvoices(ctx).Status(status).Org(org).Limit(limit).Execute()

Answers GET /v1/admin/invoices.



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
	status := "status_example" // string | Status filters on the invoice's LATEST lifecycle status (paid, open, void, …), matched case-insensitively. (optional)
	org := "org_example" // string | Org filters to one tenant, matched exactly. (optional)
	limit := "limit_example" // string | Limit caps the rows returned. total still reports the full match count. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminAPI.CloudAdminInvoices(context.Background()).Status(status).Org(org).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.CloudAdminInvoices``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudAdminInvoices`: CloudInvoicesOut
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.CloudAdminInvoices`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudAdminInvoicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **status** | **string** | Status filters on the invoice&#39;s LATEST lifecycle status (paid, open, void, …), matched case-insensitively. | 
 **org** | **string** | Org filters to one tenant, matched exactly. | 
 **limit** | **string** | Limit caps the rows returned. total still reports the full match count. | 

### Return type

[**CloudInvoicesOut**](CloudInvoicesOut.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudAdminIssueGrant

> CloudGrantOut CloudAdminIssueGrant(ctx).CloudGrantIn(cloudGrantIn).Execute()

Issues a credit grant to any org from the operator Grants view, with the target named in the body.



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
	cloudGrantIn := *openapiclient.NewCloudGrantIn() // CloudGrantIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminAPI.CloudAdminIssueGrant(context.Background()).CloudGrantIn(cloudGrantIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.CloudAdminIssueGrant``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudAdminIssueGrant`: CloudGrantOut
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.CloudAdminIssueGrant`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudAdminIssueGrantRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cloudGrantIn** | [**CloudGrantIn**](CloudGrantIn.md) |  | 

### Return type

[**CloudGrantOut**](CloudGrantOut.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudAdminMe

> CloudMeOut CloudAdminMe(ctx).Execute()

Answers with the validated operator identity — who the console is signed in as, which tier they are, and how wide their tenant window is.



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
	resp, r, err := apiClient.AdminAPI.CloudAdminMe(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.CloudAdminMe``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudAdminMe`: CloudMeOut
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.CloudAdminMe`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudAdminMeRequest struct via the builder pattern


### Return type

[**CloudMeOut**](CloudMeOut.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudAdminMetrics

> CloudMetricsOut CloudAdminMetrics(ctx).Window(window).Limit(limit).Execute()

Answers GET /v1/admin/metrics by aggregating commerce.events directly (fleet-wide, no per-org fan-out).



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
	window := "window_example" // string | Window is the movement window the new/churned MRR and the recent feed are measured over. Anything unrecognised falls back to the board default. (optional)
	limit := "limit_example" // string | Limit caps the top-customers table. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminAPI.CloudAdminMetrics(context.Background()).Window(window).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.CloudAdminMetrics``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudAdminMetrics`: CloudMetricsOut
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.CloudAdminMetrics`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudAdminMetricsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **window** | **string** | Window is the movement window the new/churned MRR and the recent feed are measured over. Anything unrecognised falls back to the board default. | 
 **limit** | **string** | Limit caps the top-customers table. | 

### Return type

[**CloudMetricsOut**](CloudMetricsOut.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudAdminMoney

> CloudMoneyOut CloudAdminMoney(ctx).Execute()

moneyBoardHandler answers GET /v1/admin/money.



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
	resp, r, err := apiClient.AdminAPI.CloudAdminMoney(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.CloudAdminMoney``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudAdminMoney`: CloudMoneyOut
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.CloudAdminMoney`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudAdminMoneyRequest struct via the builder pattern


### Return type

[**CloudMoneyOut**](CloudMoneyOut.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudAdminO11y

> CloudO11yOut CloudAdminO11y(ctx).Range_(range_).Execute()

Is the fleet-wide observability board: LLM usage (requests, tokens, cost, errors, top orgs, top models), trace RED metrics (count, p50/p95/p99 latency in ms, error rate, top services), fleet log volume, and the O11yAI generation rollup — all aggregated across EVERY tenant, with no org filter applied.



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
	range_ := "7d" // string | Range is the lower time bound: 24h, 7d or 30d. Anything else reads as the board's own default. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminAPI.CloudAdminO11y(context.Background()).Range_(range_).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.CloudAdminO11y``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudAdminO11y`: CloudO11yOut
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.CloudAdminO11y`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudAdminO11yRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **range_** | **string** | Range is the lower time bound: 24h, 7d or 30d. Anything else reads as the board&#39;s own default. | 

### Return type

[**CloudO11yOut**](CloudO11yOut.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudAdminOrgs

> CloudOrgsOut CloudAdminOrgs(ctx).Execute()

Lists the tenant directory one row per org, sorted by slug: member count and the org's month-to-date spend and credit balance, read live from IAM and commerce.



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
	resp, r, err := apiClient.AdminAPI.CloudAdminOrgs(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.CloudAdminOrgs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudAdminOrgs`: CloudOrgsOut
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.CloudAdminOrgs`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudAdminOrgsRequest struct via the builder pattern


### Return type

[**CloudOrgsOut**](CloudOrgsOut.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudAdminOverview

> CloudOverviewOut CloudAdminOverview(ctx).Execute()

Is the Platform Overview tiles: how many orgs and users are in the caller's tenant window, the fleet workload counts, and month-to-date spend and credits.



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
	resp, r, err := apiClient.AdminAPI.CloudAdminOverview(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.CloudAdminOverview``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudAdminOverview`: CloudOverviewOut
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.CloudAdminOverview`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudAdminOverviewRequest struct via the builder pattern


### Return type

[**CloudOverviewOut**](CloudOverviewOut.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudAdminPlugins

> CloudListOut CloudAdminPlugins(ctx).Scope(scope).Execute()

Reports what each host is actually running: every loaded plugin with its version, pid, uptime, reload and restart counts, and its measured CPU, RSS, thread and fd cost — read from the kernel, which is only answerable at all because a plugin is a process.



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
	scope := "fleet" // string | Scope \"host\" answers for THIS host only. Default \"fleet\" fans out to every live peer. A peer answers a host-scoped read, which is what stops the fan-out recursing. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminAPI.CloudAdminPlugins(context.Background()).Scope(scope).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.CloudAdminPlugins``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudAdminPlugins`: CloudListOut
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.CloudAdminPlugins`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudAdminPluginsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **scope** | **string** | Scope \&quot;host\&quot; answers for THIS host only. Default \&quot;fleet\&quot; fans out to every live peer. A peer answers a host-scoped read, which is what stops the fan-out recursing. | 

### Return type

[**CloudListOut**](CloudListOut.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudAdminProducts

> CloudProductsOut CloudAdminProducts(ctx).Kind(kind).Tier(tier).Env(env).Execute()

Lists the fleet workload registry: every operator App CR across the platform namespaces with its declared vs running image tag, reconciled health/phase and drift verdict.



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
	kind := "kind_example" // string | Kind matches the operator App CR's declared spec.role (sql|kv|generic|ingress). (optional)
	tier := "data" // string | Tier matches the derived infra grouping (cloud|data|edge|daemon|paas|app). (optional)
	env := "main" // string | Env matches the lifecycle namespace (main|test|dev). (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminAPI.CloudAdminProducts(context.Background()).Kind(kind).Tier(tier).Env(env).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.CloudAdminProducts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudAdminProducts`: CloudProductsOut
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.CloudAdminProducts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudAdminProductsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **kind** | **string** | Kind matches the operator App CR&#39;s declared spec.role (sql|kv|generic|ingress). | 
 **tier** | **string** | Tier matches the derived infra grouping (cloud|data|edge|daemon|paas|app). | 
 **env** | **string** | Env matches the lifecycle namespace (main|test|dev). | 

### Return type

[**CloudProductsOut**](CloudProductsOut.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudAdminPromo

> CloudRawOut CloudAdminPromo(ctx).Execute()

Reads the current platform plan promo — the singleton discount offer, e.g.



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
	resp, r, err := apiClient.AdminAPI.CloudAdminPromo(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.CloudAdminPromo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudAdminPromo`: CloudRawOut
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.CloudAdminPromo`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudAdminPromoRequest struct via the builder pattern


### Return type

[**CloudRawOut**](CloudRawOut.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudAdminProvidersCredit

> CloudProvidersCreditOut CloudAdminProvidersCredit(ctx).Execute()

Serves GET /v1/admin/providers/credit — the per-provider upstream credit ledger.



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
	resp, r, err := apiClient.AdminAPI.CloudAdminProvidersCredit(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.CloudAdminProvidersCredit``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudAdminProvidersCredit`: CloudProvidersCreditOut
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.CloudAdminProvidersCredit`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudAdminProvidersCreditRequest struct via the builder pattern


### Return type

[**CloudProvidersCreditOut**](CloudProvidersCreditOut.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudAdminReactivateCustomer

> CloudAccessOut CloudAdminReactivateCustomer(ctx, org).Execute()

Restores access for every member of the org, undoing a suspend.



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
	org := "org_example" // string | Org is the tenant slug from the path.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminAPI.CloudAdminReactivateCustomer(context.Background(), org).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.CloudAdminReactivateCustomer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudAdminReactivateCustomer`: CloudAccessOut
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.CloudAdminReactivateCustomer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**org** | **string** | Org is the tenant slug from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudAdminReactivateCustomerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CloudAccessOut**](CloudAccessOut.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudAdminReloadPlugin

> CloudActionOut CloudAdminReloadPlugin(ctx, name).CloudReloadIn(cloudReloadIn).Execute()

Swaps a plugin for another build without dropping a request.



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
	name := "billing" // string | Name is the app, from the path. It must be one the manifest declares.
	cloudReloadIn := *openapiclient.NewCloudReloadIn() // CloudReloadIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminAPI.CloudAdminReloadPlugin(context.Background(), name).CloudReloadIn(cloudReloadIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.CloudAdminReloadPlugin``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudAdminReloadPlugin`: CloudActionOut
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.CloudAdminReloadPlugin`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name is the app, from the path. It must be one the manifest declares. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudAdminReloadPluginRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cloudReloadIn** | [**CloudReloadIn**](CloudReloadIn.md) |  | 

### Return type

[**CloudActionOut**](CloudActionOut.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudAdminResizeDroplet

> CloudMutationOut CloudAdminResizeDroplet(ctx, id).CloudDropletIn(cloudDropletIn).Execute()

Changes a droplet's plan.



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
	id := "id_example" // string | ID is the DO droplet id, from the path. Numeric.
	cloudDropletIn := *openapiclient.NewCloudDropletIn() // CloudDropletIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminAPI.CloudAdminResizeDroplet(context.Background(), id).CloudDropletIn(cloudDropletIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.CloudAdminResizeDroplet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudAdminResizeDroplet`: CloudMutationOut
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.CloudAdminResizeDroplet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the DO droplet id, from the path. Numeric. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudAdminResizeDropletRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cloudDropletIn** | [**CloudDropletIn**](CloudDropletIn.md) |  | 

### Return type

[**CloudMutationOut**](CloudMutationOut.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudAdminResizeVolume

> CloudMutationOut CloudAdminResizeVolume(ctx, id).CloudVolumeIn(cloudVolumeIn).Execute()

Grows a volume.



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
	id := "id_example" // string | ID is the DO volume id, from the path.
	cloudVolumeIn := *openapiclient.NewCloudVolumeIn() // CloudVolumeIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminAPI.CloudAdminResizeVolume(context.Background(), id).CloudVolumeIn(cloudVolumeIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.CloudAdminResizeVolume``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudAdminResizeVolume`: CloudMutationOut
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.CloudAdminResizeVolume`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the DO volume id, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudAdminResizeVolumeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cloudVolumeIn** | [**CloudVolumeIn**](CloudVolumeIn.md) |  | 

### Return type

[**CloudMutationOut**](CloudMutationOut.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudAdminRevenue

> CloudRevenueOut CloudAdminRevenue(ctx).Execute()

Is the fleet money board: total prepaid balances held, total realized spend, MRR, ARPU, a per-customer table sorted highest-revenue first, and a real 30-day spend trend from the usage ledger.



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
	resp, r, err := apiClient.AdminAPI.CloudAdminRevenue(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.CloudAdminRevenue``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudAdminRevenue`: CloudRevenueOut
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.CloudAdminRevenue`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudAdminRevenueRequest struct via the builder pattern


### Return type

[**CloudRevenueOut**](CloudRevenueOut.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudAdminRoles

> CloudIamRowsOut CloudAdminRoles(ctx).Owner(owner).P(p).PageSize(pageSize).Execute()

Lists IAM roles for one owner org, forwarded VERBATIM from IAM's get-roles.



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
	owner := "admin" // string | Owner is the org whose rows to read. Defaults to the admin org, which owns the platform's roles and applications. (optional)
	p := "1" // string | Page is the 1-based page number. Forwarded only when set — IAM applies its own default otherwise. (optional)
	pageSize := "50" // string | PageSize is rows per page. Forwarded only when set. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminAPI.CloudAdminRoles(context.Background()).Owner(owner).P(p).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.CloudAdminRoles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudAdminRoles`: CloudIamRowsOut
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.CloudAdminRoles`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudAdminRolesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **owner** | **string** | Owner is the org whose rows to read. Defaults to the admin org, which owns the platform&#39;s roles and applications. | 
 **p** | **string** | Page is the 1-based page number. Forwarded only when set — IAM applies its own default otherwise. | 
 **pageSize** | **string** | PageSize is rows per page. Forwarded only when set. | 

### Return type

[**CloudIamRowsOut**](CloudIamRowsOut.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudAdminScaleNodePool

> CloudMutationOut CloudAdminScaleNodePool(ctx, id, pool).CloudScaleIn(cloudScaleIn).Execute()

Sets a node pool's node count — the ONE correct way to change how many nodes a DOKS cluster has.



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
	id := "id_example" // string | ID is the DOKS cluster id, from the path.
	pool := "pool_example" // string | Pool is the node pool, from the path. Its DO id or its name — both are unique within a cluster, and an operator reads the name off the board.
	cloudScaleIn := *openapiclient.NewCloudScaleIn() // CloudScaleIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminAPI.CloudAdminScaleNodePool(context.Background(), id, pool).CloudScaleIn(cloudScaleIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.CloudAdminScaleNodePool``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudAdminScaleNodePool`: CloudMutationOut
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.CloudAdminScaleNodePool`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the DOKS cluster id, from the path. | 
**pool** | **string** | Pool is the node pool, from the path. Its DO id or its name — both are unique within a cluster, and an operator reads the name off the board. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudAdminScaleNodePoolRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **cloudScaleIn** | [**CloudScaleIn**](CloudScaleIn.md) |  | 

### Return type

[**CloudMutationOut**](CloudMutationOut.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudAdminServices

> CloudServicesOut CloudAdminServices(ctx).Execute()

Reads the launch board: every hosted service in the registry with its LIVE waitlist mode, evaluated through the flag engine.



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
	resp, r, err := apiClient.AdminAPI.CloudAdminServices(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.CloudAdminServices``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudAdminServices`: CloudServicesOut
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.CloudAdminServices`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudAdminServicesRequest struct via the builder pattern


### Return type

[**CloudServicesOut**](CloudServicesOut.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudAdminSetFlag

> CloudFlagsOut CloudAdminSetFlag(ctx, key).CloudSetFlagIn(cloudSetFlagIn).Execute()

Stores or overwrites ONE platform switch's definition and answers with the whole board as it now stands.



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
	key := "key_example" // string | Key is the switch to write, taken from the path (e.g. \"waitlist.chat\").
	cloudSetFlagIn := *openapiclient.NewCloudSetFlagIn() // CloudSetFlagIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminAPI.CloudAdminSetFlag(context.Background(), key).CloudSetFlagIn(cloudSetFlagIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.CloudAdminSetFlag``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudAdminSetFlag`: CloudFlagsOut
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.CloudAdminSetFlag`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string** | Key is the switch to write, taken from the path (e.g. \&quot;waitlist.chat\&quot;). | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudAdminSetFlagRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cloudSetFlagIn** | [**CloudSetFlagIn**](CloudSetFlagIn.md) |  | 

### Return type

[**CloudFlagsOut**](CloudFlagsOut.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudAdminSetPromo

> CloudRawOut CloudAdminSetPromo(ctx).CloudPromoIn(cloudPromoIn).Execute()

Upserts the platform plan promo — the ONE place the offer is configured.



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
	cloudPromoIn := *openapiclient.NewCloudPromoIn() // CloudPromoIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminAPI.CloudAdminSetPromo(context.Background()).CloudPromoIn(cloudPromoIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.CloudAdminSetPromo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudAdminSetPromo`: CloudRawOut
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.CloudAdminSetPromo`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudAdminSetPromoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cloudPromoIn** | [**CloudPromoIn**](CloudPromoIn.md) |  | 

### Return type

[**CloudRawOut**](CloudRawOut.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudAdminSetServiceMode

> CloudServiceOut CloudAdminSetServiceMode(ctx, service).CloudServiceModeIn(cloudServiceModeIn).Execute()

Flips ONE service's waitlist switch — the launch lever.



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
	service := "service_example" // string | Service is the slug to flip, taken from the path.
	cloudServiceModeIn := *openapiclient.NewCloudServiceModeIn() // CloudServiceModeIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminAPI.CloudAdminSetServiceMode(context.Background(), service).CloudServiceModeIn(cloudServiceModeIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.CloudAdminSetServiceMode``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudAdminSetServiceMode`: CloudServiceOut
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.CloudAdminSetServiceMode`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**service** | **string** | Service is the slug to flip, taken from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudAdminSetServiceModeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cloudServiceModeIn** | [**CloudServiceModeIn**](CloudServiceModeIn.md) |  | 

### Return type

[**CloudServiceOut**](CloudServiceOut.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudAdminSnapshotVolume

> CloudVolumeSnapshotOut CloudAdminSnapshotVolume(ctx, id).CloudVolumeIn(cloudVolumeIn).Execute()

Takes a point-in-time snapshot of one volume — the undo a delete relies on, available on its own so an operator can take one before any risky change.



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
	id := "id_example" // string | ID is the DO volume id, from the path.
	cloudVolumeIn := *openapiclient.NewCloudVolumeIn() // CloudVolumeIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminAPI.CloudAdminSnapshotVolume(context.Background(), id).CloudVolumeIn(cloudVolumeIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.CloudAdminSnapshotVolume``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudAdminSnapshotVolume`: CloudVolumeSnapshotOut
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.CloudAdminSnapshotVolume`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the DO volume id, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudAdminSnapshotVolumeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cloudVolumeIn** | [**CloudVolumeIn**](CloudVolumeIn.md) |  | 

### Return type

[**CloudVolumeSnapshotOut**](CloudVolumeSnapshotOut.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudAdminSpendCaps

> CloudRawOut CloudAdminSpendCaps(ctx).Org(org).Id(id).Execute()

Reads one org's usage caps: its spend alerts plus the derived period spend, over/warn state and reset time.



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
	org := "acme" // string | Org is the tenant to act on. Required for a SuperAdmin — they must name their target; ignored for a white-label admin, who always acts on their own org. (optional)
	id := "id_example" // string | ID is the cap to edit or remove, from the path. Unused by the list and create ops. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminAPI.CloudAdminSpendCaps(context.Background()).Org(org).Id(id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.CloudAdminSpendCaps``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudAdminSpendCaps`: CloudRawOut
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.CloudAdminSpendCaps`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudAdminSpendCapsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **org** | **string** | Org is the tenant to act on. Required for a SuperAdmin — they must name their target; ignored for a white-label admin, who always acts on their own org. | 
 **id** | **string** | ID is the cap to edit or remove, from the path. Unused by the list and create ops. | 

### Return type

[**CloudRawOut**](CloudRawOut.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudAdminSubscriptions

> CloudSubscriptionsOut CloudAdminSubscriptions(ctx).Status(status).Org(org).Limit(limit).Execute()

Answers GET /v1/admin/subscriptions.



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
	status := "status_example" // string | Status filters on the subscription's LATEST lifecycle status (active, trialing, canceled, …), matched case-insensitively. (optional)
	org := "org_example" // string | Org filters to one tenant, matched exactly. (optional)
	limit := "limit_example" // string | Limit caps the rows returned. total still reports the full match count. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminAPI.CloudAdminSubscriptions(context.Background()).Status(status).Org(org).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.CloudAdminSubscriptions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudAdminSubscriptions`: CloudSubscriptionsOut
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.CloudAdminSubscriptions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudAdminSubscriptionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **status** | **string** | Status filters on the subscription&#39;s LATEST lifecycle status (active, trialing, canceled, …), matched case-insensitively. | 
 **org** | **string** | Org filters to one tenant, matched exactly. | 
 **limit** | **string** | Limit caps the rows returned. total still reports the full match count. | 

### Return type

[**CloudSubscriptionsOut**](CloudSubscriptionsOut.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudAdminSubsystems

> CloudSubsystemsOut CloudAdminSubsystems(ctx).Range_(range_).Execute()

subsystems answers GET /v1/admin/subsystems.



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
	range_ := "range__example" // string | Range bounds the telemetry window: 24h, 7d or 30d. Anything else, including empty, resolves to the default through the same o11yRange the o11y board uses. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminAPI.CloudAdminSubsystems(context.Background()).Range_(range_).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.CloudAdminSubsystems``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudAdminSubsystems`: CloudSubsystemsOut
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.CloudAdminSubsystems`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudAdminSubsystemsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **range_** | **string** | Range bounds the telemetry window: 24h, 7d or 30d. Anything else, including empty, resolves to the default through the same o11yRange the o11y board uses. | 

### Return type

[**CloudSubsystemsOut**](CloudSubsystemsOut.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudAdminSuspendCustomer

> CloudAccessOut CloudAdminSuspendCustomer(ctx, org).Execute()

Cuts off every member of the org: IAM refuses a forbidden user at login AND at token issuance, so a suspended customer can neither sign in nor mint a fresh token.



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
	org := "org_example" // string | Org is the tenant slug from the path.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminAPI.CloudAdminSuspendCustomer(context.Background(), org).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.CloudAdminSuspendCustomer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudAdminSuspendCustomer`: CloudAccessOut
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.CloudAdminSuspendCustomer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**org** | **string** | Org is the tenant slug from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudAdminSuspendCustomerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CloudAccessOut**](CloudAccessOut.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudAdminSync

> CloudSyncOut CloudAdminSync(ctx).Execute()

Answers the operator's \"Sync now\" button.



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
	resp, r, err := apiClient.AdminAPI.CloudAdminSync(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.CloudAdminSync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudAdminSync`: CloudSyncOut
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.CloudAdminSync`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudAdminSyncRequest struct via the builder pattern


### Return type

[**CloudSyncOut**](CloudSyncOut.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudAdminUpdateSpendCap

> CloudRawOut CloudAdminUpdateSpendCap(ctx, id).CloudCapIn(cloudCapIn).Execute()

Edits one cap by id — raise or lower the ceiling, flip enforcement.



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
	id := "id_example" // string | ID is the cap to edit or remove, from the path. Unused by the list and create ops.
	cloudCapIn := *openapiclient.NewCloudCapIn() // CloudCapIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminAPI.CloudAdminUpdateSpendCap(context.Background(), id).CloudCapIn(cloudCapIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.CloudAdminUpdateSpendCap``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudAdminUpdateSpendCap`: CloudRawOut
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.CloudAdminUpdateSpendCap`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the cap to edit or remove, from the path. Unused by the list and create ops. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudAdminUpdateSpendCapRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cloudCapIn** | [**CloudCapIn**](CloudCapIn.md) |  | 

### Return type

[**CloudRawOut**](CloudRawOut.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudAdminUpsertService

> CloudServiceOut CloudAdminUpsertService(ctx).CloudServiceInput(cloudServiceInput).Execute()

Onboards a hosted service, or edits one, so a new host comes under the launch gate WITHOUT a redeploy.



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
	cloudServiceInput := *openapiclient.NewCloudServiceInput() // CloudServiceInput | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminAPI.CloudAdminUpsertService(context.Background()).CloudServiceInput(cloudServiceInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.CloudAdminUpsertService``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudAdminUpsertService`: CloudServiceOut
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.CloudAdminUpsertService`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudAdminUpsertServiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cloudServiceInput** | [**CloudServiceInput**](CloudServiceInput.md) |  | 

### Return type

[**CloudServiceOut**](CloudServiceOut.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudAdminUsage

> CloudUsageOut CloudAdminUsage(ctx).Org(org).Execute()

Returns the month-to-date money totals: one org's when org names one, else the fleet sum across every org a SuperAdmin can see.



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
	org := "acme" // string | Org reads ONE tenant's month-to-date total instead of the fleet sum. Honoured for a SuperAdmin only — a white-label admin always reads their own org. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminAPI.CloudAdminUsage(context.Background()).Org(org).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.CloudAdminUsage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudAdminUsage`: CloudUsageOut
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.CloudAdminUsage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudAdminUsageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **org** | **string** | Org reads ONE tenant&#39;s month-to-date total instead of the fleet sum. Honoured for a SuperAdmin only — a white-label admin always reads their own org. | 

### Return type

[**CloudUsageOut**](CloudUsageOut.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudAdminUsageFunding

> CloudUsageFundingOut CloudAdminUsageFunding(ctx).From(from).To(to).Execute()

Splits our upstream AI usage by how it was FUNDED: one row per (provider, model) over the window, tagged credit (provider grant still remaining), paid (grant exhausted) or paid_only (no grant at all).



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
	from := "2026-07-01T00:00:00Z" // string | From is the inclusive start of the window. Unparseable or absent, together with To, falls back to the last 30 days. (optional)
	to := "2026-07-27T00:00:00Z" // string | To is the exclusive end of the window. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminAPI.CloudAdminUsageFunding(context.Background()).From(from).To(to).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.CloudAdminUsageFunding``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudAdminUsageFunding`: CloudUsageFundingOut
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.CloudAdminUsageFunding`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudAdminUsageFundingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **from** | **string** | From is the inclusive start of the window. Unparseable or absent, together with To, falls back to the last 30 days. | 
 **to** | **string** | To is the exclusive end of the window. | 

### Return type

[**CloudUsageFundingOut**](CloudUsageFundingOut.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudAdminUsers

> CloudUsersOut CloudAdminUsers(ctx).Org(org).Q(q).P(p).PageSize(pageSize).Execute()

Lists the user directory across the caller's tenant window, one page at a time.



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
	org := "acme" // string | Org narrows the directory to ONE tenant. Honoured for a SuperAdmin only — a white-label admin is pinned to their own org and this is ignored. (optional)
	q := "ada" // string | Query is a free-text filter, matched by IAM as a \"contains\" over the user name. (optional)
	p := "1" // string | Page is the 1-based page number. Defaults to \"1\"; IAM returns zero rows AND a zero total when it is unset, so this layer never leaves it empty. (optional)
	pageSize := "50" // string | PageSize is rows per page. Defaults to \"200\", the shared admin page size. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminAPI.CloudAdminUsers(context.Background()).Org(org).Q(q).P(p).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.CloudAdminUsers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudAdminUsers`: CloudUsersOut
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.CloudAdminUsers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudAdminUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **org** | **string** | Org narrows the directory to ONE tenant. Honoured for a SuperAdmin only — a white-label admin is pinned to their own org and this is ignored. | 
 **q** | **string** | Query is a free-text filter, matched by IAM as a \&quot;contains\&quot; over the user name. | 
 **p** | **string** | Page is the 1-based page number. Defaults to \&quot;1\&quot;; IAM returns zero rows AND a zero total when it is unset, so this layer never leaves it empty. | 
 **pageSize** | **string** | PageSize is rows per page. Defaults to \&quot;200\&quot;, the shared admin page size. | 

### Return type

[**CloudUsersOut**](CloudUsersOut.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudAdminWaitlist

> CloudRawOut CloudAdminWaitlist(ctx).Waitlist(waitlist).Page(page).PageSize(pageSize).Execute()

Reads one waitlist's leaderboard from the Hanzo waitlist engine — position, points and referral standing per entry — proxied server-authed with the engine secret, never a client credential.



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
	waitlist := "chat" // string | Waitlist is the waitlist slug to read (e.g. \"chat\"). The engine decides what an empty slug means. (optional)
	page := "1" // string | Page is the 1-based page number. (optional)
	pageSize := "50" // string | PageSize is entries per page. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminAPI.CloudAdminWaitlist(context.Background()).Waitlist(waitlist).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.CloudAdminWaitlist``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudAdminWaitlist`: CloudRawOut
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.CloudAdminWaitlist`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudAdminWaitlistRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **waitlist** | **string** | Waitlist is the waitlist slug to read (e.g. \&quot;chat\&quot;). The engine decides what an empty slug means. | 
 **page** | **string** | Page is the 1-based page number. | 
 **pageSize** | **string** | PageSize is entries per page. | 

### Return type

[**CloudRawOut**](CloudRawOut.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudAdminWaitlistBoost

> CloudRawOut CloudAdminWaitlistBoost(ctx).CloudWaitlistBoostRequest(cloudWaitlistBoostRequest).Execute()

Grants a user waitlist points, moving them up toward the access cutoff.



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
	cloudWaitlistBoostRequest := *openapiclient.NewCloudWaitlistBoostRequest() // CloudWaitlistBoostRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminAPI.CloudAdminWaitlistBoost(context.Background()).CloudWaitlistBoostRequest(cloudWaitlistBoostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.CloudAdminWaitlistBoost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudAdminWaitlistBoost`: CloudRawOut
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.CloudAdminWaitlistBoost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudAdminWaitlistBoostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cloudWaitlistBoostRequest** | [**CloudWaitlistBoostRequest**](CloudWaitlistBoostRequest.md) |  | 

### Return type

[**CloudRawOut**](CloudRawOut.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1AdminAffiliates

> CloudGetV1AdminAffiliates(ctx).Execute()



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
	r, err := apiClient.AdminAPI.CloudGetV1AdminAffiliates(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.CloudGetV1AdminAffiliates``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1AdminAffiliatesRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1AdminAuthors

> CloudAdminBook CloudGetV1AdminAuthors(ctx).Limit(limit).Execute()

ListAuthors returns the platform's whole author program — every org's author record, not the caller's — with each one's repository and deploy counts and a fleet roll-up of the money accrued, pending and paid.



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
	limit := int32(56) // int32 | Limit bounds the page. 0 or less means the default of 500; anything above 1000 is clamped to 1000. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminAPI.CloudGetV1AdminAuthors(context.Background()).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.CloudGetV1AdminAuthors``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1AdminAuthors`: CloudAdminBook
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.CloudGetV1AdminAuthors`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1AdminAuthorsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Limit bounds the page. 0 or less means the default of 500; anything above 1000 is clamped to 1000. | 

### Return type

[**CloudAdminBook**](CloudAdminBook.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1AdminAuthorsIdBasis

> CloudBasisResult CloudGetV1AdminAuthorsIdBasis(ctx, id).Period(period).Execute()

AuthorRoyaltyBasis returns the audit trail behind ONE author's royalty — the same payload the author reads at /v1/authors/basis, from the same builder, so support sees exactly what the author sees rather than a parallel view free to drift.



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
	id := "aut_1f…" // string | ID is the author record's handle, from the path.
	period := "2026-07" // string | Period is the UTC accrual month, YYYY-MM. Empty means every period; any other shape is refused with 400. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminAPI.CloudGetV1AdminAuthorsIdBasis(context.Background(), id).Period(period).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.CloudGetV1AdminAuthorsIdBasis``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1AdminAuthorsIdBasis`: CloudBasisResult
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.CloudGetV1AdminAuthorsIdBasis`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the author record&#39;s handle, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1AdminAuthorsIdBasisRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **period** | **string** | Period is the UTC accrual month, YYYY-MM. Empty means every period; any other shape is refused with 400. | 

### Return type

[**CloudBasisResult**](CloudBasisResult.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1AdminCatalog

> CloudAdminCatalogOut CloudGetV1AdminCatalog(ctx).Execute()

GetAdminCatalog returns the full model and provider catalog annotated with each entry's enablement state, for the operator console.



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
	resp, r, err := apiClient.AdminAPI.CloudGetV1AdminCatalog(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.CloudGetV1AdminCatalog``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1AdminCatalog`: CloudAdminCatalogOut
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.CloudGetV1AdminCatalog`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1AdminCatalogRequest struct via the builder pattern


### Return type

[**CloudAdminCatalogOut**](CloudAdminCatalogOut.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1AdminEnablement

> CloudAdminEnablementBoard CloudGetV1AdminEnablement(ctx).Execute()

ListEnablement returns every item an operator has set an enablement state on — its global state (off, beta or ga) and the orgs granted its beta.



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
	resp, r, err := apiClient.AdminAPI.CloudGetV1AdminEnablement(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.CloudGetV1AdminEnablement``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1AdminEnablement`: CloudAdminEnablementBoard
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.CloudGetV1AdminEnablement`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1AdminEnablementRequest struct via the builder pattern


### Return type

[**CloudAdminEnablementBoard**](CloudAdminEnablementBoard.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1AdminReferrals

> CloudGetV1AdminReferrals(ctx).Execute()



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
	r, err := apiClient.AdminAPI.CloudGetV1AdminReferrals(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.CloudGetV1AdminReferrals``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1AdminReferralsRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1AdminReferralsBonuses

> CloudAdminBonusesEnvelope CloudGetV1AdminReferralsBonuses(ctx).Limit(limit).Execute()

Returns every one-time referral bonus in the ledger with a fleet summary.



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
	limit := "limit_example" // string | Limit is how many referrals to return, as a decimal string in the `?limit=` query. Absent, unparseable or non-positive means 500; over 1000 is clamped to 1000. It is a string rather than a number because the parse that has always served this route trims surrounding whitespace, and one parse rule is better than two. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminAPI.CloudGetV1AdminReferralsBonuses(context.Background()).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.CloudGetV1AdminReferralsBonuses``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1AdminReferralsBonuses`: CloudAdminBonusesEnvelope
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.CloudGetV1AdminReferralsBonuses`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1AdminReferralsBonusesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **string** | Limit is how many referrals to return, as a decimal string in the &#x60;?limit&#x3D;&#x60; query. Absent, unparseable or non-positive means 500; over 1000 is clamped to 1000. It is a string rather than a number because the parse that has always served this route trims surrounding whitespace, and one parse rule is better than two. | 

### Return type

[**CloudAdminBonusesEnvelope**](CloudAdminBonusesEnvelope.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1AdminTreasury

> CloudAdminReportOut CloudGetV1AdminTreasury(ctx).Limit(limit).Execute()

GetAdminTreasury returns the whole treasury board for a SuperAdmin: the reserve fund report, the recent double-entry journal, and the Hanzo L1 anchor status of the ledger root.



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
	limit := int32(56) // int32 | Limit caps the journal entries returned. Out of range or unparseable takes the default. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminAPI.CloudGetV1AdminTreasury(context.Background()).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.CloudGetV1AdminTreasury``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1AdminTreasury`: CloudAdminReportOut
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.CloudGetV1AdminTreasury`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1AdminTreasuryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Limit caps the journal entries returned. Out of range or unparseable takes the default. | 

### Return type

[**CloudAdminReportOut**](CloudAdminReportOut.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPatchV1AdminCatalogModelsByWildcard1

> CloudPatchV1AdminCatalogModelsByWildcard1(ctx, wildcard1).Execute()



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
	wildcard1 := "wildcard1_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AdminAPI.CloudPatchV1AdminCatalogModelsByWildcard1(context.Background(), wildcard1).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.CloudPatchV1AdminCatalogModelsByWildcard1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**wildcard1** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudPatchV1AdminCatalogModelsByWildcard1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPatchV1AdminCatalogProvidersName

> CloudOverlay CloudPatchV1AdminCatalogProvidersName(ctx, name).CloudProviderPatchIn(cloudProviderPatchIn).Execute()

PatchProvider sets one provider's availability overlay.



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
	name := "name_example" // string | Name is the provider the overlay belongs to, from the URL.
	cloudProviderPatchIn := *openapiclient.NewCloudProviderPatchIn() // CloudProviderPatchIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminAPI.CloudPatchV1AdminCatalogProvidersName(context.Background(), name).CloudProviderPatchIn(cloudProviderPatchIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.CloudPatchV1AdminCatalogProvidersName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudPatchV1AdminCatalogProvidersName`: CloudOverlay
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.CloudPatchV1AdminCatalogProvidersName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name is the provider the overlay belongs to, from the URL. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudPatchV1AdminCatalogProvidersNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cloudProviderPatchIn** | [**CloudProviderPatchIn**](CloudProviderPatchIn.md) |  | 

### Return type

[**CloudOverlay**](CloudOverlay.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPostV1AdminAffiliatesByIdApprove

> CloudPostV1AdminAffiliatesByIdApprove(ctx, id).Execute()



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
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AdminAPI.CloudPostV1AdminAffiliatesByIdApprove(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.CloudPostV1AdminAffiliatesByIdApprove``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudPostV1AdminAffiliatesByIdApproveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPostV1AdminAffiliatesByIdPayout

> CloudPostV1AdminAffiliatesByIdPayout(ctx, id).Execute()



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
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AdminAPI.CloudPostV1AdminAffiliatesByIdPayout(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.CloudPostV1AdminAffiliatesByIdPayout``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudPostV1AdminAffiliatesByIdPayoutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPostV1AdminAffiliatesByIdRate

> CloudPostV1AdminAffiliatesByIdRate(ctx, id).Execute()



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
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AdminAPI.CloudPostV1AdminAffiliatesByIdRate(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.CloudPostV1AdminAffiliatesByIdRate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudPostV1AdminAffiliatesByIdRateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPostV1AdminAffiliatesByIdSuspend

> CloudPostV1AdminAffiliatesByIdSuspend(ctx, id).Execute()



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
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AdminAPI.CloudPostV1AdminAffiliatesByIdSuspend(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.CloudPostV1AdminAffiliatesByIdSuspend``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudPostV1AdminAffiliatesByIdSuspendRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPostV1AdminAffiliatesSweep

> CloudPostV1AdminAffiliatesSweep(ctx).Execute()



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
	r, err := apiClient.AdminAPI.CloudPostV1AdminAffiliatesSweep(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.CloudPostV1AdminAffiliatesSweep``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudPostV1AdminAffiliatesSweepRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPostV1AdminAuthorsIdApprove

> CloudAuthorResult CloudPostV1AdminAuthorsIdApprove(ctx, id).CloudApproveRequest(cloudApproveRequest).Execute()

ApproveAuthor admits one author to EARNING, optionally on a negotiated royalty share.



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
	id := "aut_1f…" // string | ID is the author to approve, from the path.
	cloudApproveRequest := *openapiclient.NewCloudApproveRequest() // CloudApproveRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminAPI.CloudPostV1AdminAuthorsIdApprove(context.Background(), id).CloudApproveRequest(cloudApproveRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.CloudPostV1AdminAuthorsIdApprove``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudPostV1AdminAuthorsIdApprove`: CloudAuthorResult
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.CloudPostV1AdminAuthorsIdApprove`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the author to approve, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudPostV1AdminAuthorsIdApproveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cloudApproveRequest** | [**CloudApproveRequest**](CloudApproveRequest.md) |  | 

### Return type

[**CloudAuthorResult**](CloudAuthorResult.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPostV1AdminAuthorsIdPayout

> CloudPayoutResult CloudPostV1AdminAuthorsIdPayout(ctx, id).CloudPayoutRequest(cloudPayoutRequest).Execute()

PayAuthor records a payout of accrued royalty and settles it.



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
	id := "aut_1f…" // string | ID is the author to pay, from the path.
	cloudPayoutRequest := *openapiclient.NewCloudPayoutRequest() // CloudPayoutRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminAPI.CloudPostV1AdminAuthorsIdPayout(context.Background(), id).CloudPayoutRequest(cloudPayoutRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.CloudPostV1AdminAuthorsIdPayout``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudPostV1AdminAuthorsIdPayout`: CloudPayoutResult
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.CloudPostV1AdminAuthorsIdPayout`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the author to pay, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudPostV1AdminAuthorsIdPayoutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cloudPayoutRequest** | [**CloudPayoutRequest**](CloudPayoutRequest.md) |  | 

### Return type

[**CloudPayoutResult**](CloudPayoutResult.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPostV1AdminAuthorsIdSuspend

> CloudAuthorResult CloudPostV1AdminAuthorsIdSuspend(ctx, id).Execute()

SuspendAuthor stops one author earning.



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
	id := "id_example" // string | ID is the author record's handle, \"aut_\"-prefixed.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminAPI.CloudPostV1AdminAuthorsIdSuspend(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.CloudPostV1AdminAuthorsIdSuspend``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudPostV1AdminAuthorsIdSuspend`: CloudAuthorResult
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.CloudPostV1AdminAuthorsIdSuspend`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the author record&#39;s handle, \&quot;aut_\&quot;-prefixed. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudPostV1AdminAuthorsIdSuspendRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CloudAuthorResult**](CloudAuthorResult.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPostV1AdminAuthorsSweep

> CloudAuthorSweepResult CloudPostV1AdminAuthorsSweep(ctx).Execute()

SweepAuthorRoyalty runs the accrual sweep across every approved author: for each of their deploying orgs it computes this period's royalty from that org's metered spend and latches it at most once per period.



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
	resp, r, err := apiClient.AdminAPI.CloudPostV1AdminAuthorsSweep(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.CloudPostV1AdminAuthorsSweep``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudPostV1AdminAuthorsSweep`: CloudAuthorSweepResult
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.CloudPostV1AdminAuthorsSweep`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudPostV1AdminAuthorsSweepRequest struct via the builder pattern


### Return type

[**CloudAuthorSweepResult**](CloudAuthorSweepResult.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPostV1AdminReferralsSweep

> CloudSweepEnvelope CloudPostV1AdminReferralsSweep(ctx).Execute()

Qualify-checks every pending referral and grants the ones that now qualify.



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
	resp, r, err := apiClient.AdminAPI.CloudPostV1AdminReferralsSweep(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.CloudPostV1AdminReferralsSweep``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudPostV1AdminReferralsSweep`: CloudSweepEnvelope
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.CloudPostV1AdminReferralsSweep`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudPostV1AdminReferralsSweepRequest struct via the builder pattern


### Return type

[**CloudSweepEnvelope**](CloudSweepEnvelope.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPostV1AdminTreasuryAnchor

> CloudAnchorOut CloudPostV1AdminTreasuryAnchor(ctx).Execute()

AnchorTreasury commits the current ledger root to Hanzo L1, making the books tamper-evident on chain, and returns the anchoring status.



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
	resp, r, err := apiClient.AdminAPI.CloudPostV1AdminTreasuryAnchor(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.CloudPostV1AdminTreasuryAnchor``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudPostV1AdminTreasuryAnchor`: CloudAnchorOut
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.CloudPostV1AdminTreasuryAnchor`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudPostV1AdminTreasuryAnchorRequest struct via the builder pattern


### Return type

[**CloudAnchorOut**](CloudAnchorOut.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPostV1AdminTreasuryBindAnchor

> CloudBindOut CloudPostV1AdminTreasuryBindAnchor(ctx).Execute()

BindTreasuryAnchorSigner makes the reserve's threshold MPC wallet the signer for on-chain anchors, and returns its EVM address so an operator can fund it for gas.



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
	resp, r, err := apiClient.AdminAPI.CloudPostV1AdminTreasuryBindAnchor(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.CloudPostV1AdminTreasuryBindAnchor``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudPostV1AdminTreasuryBindAnchor`: CloudBindOut
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.CloudPostV1AdminTreasuryBindAnchor`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudPostV1AdminTreasuryBindAnchorRequest struct via the builder pattern


### Return type

[**CloudBindOut**](CloudBindOut.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPostV1AdminTreasuryPolicy

> CloudPolicyOut CloudPostV1AdminTreasuryPolicy(ctx).CloudPolicyRequest(cloudPolicyRequest).Execute()

SetTreasuryPolicy sets the revenue-share basis points a sweep accrues into the reserve fund and returns the stored policy.



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
	cloudPolicyRequest := *openapiclient.NewCloudPolicyRequest() // CloudPolicyRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminAPI.CloudPostV1AdminTreasuryPolicy(context.Background()).CloudPolicyRequest(cloudPolicyRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.CloudPostV1AdminTreasuryPolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudPostV1AdminTreasuryPolicy`: CloudPolicyOut
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.CloudPostV1AdminTreasuryPolicy`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudPostV1AdminTreasuryPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cloudPolicyRequest** | [**CloudPolicyRequest**](CloudPolicyRequest.md) |  | 

### Return type

[**CloudPolicyOut**](CloudPolicyOut.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPostV1AdminTreasurySeed

> CloudSeedOut CloudPostV1AdminTreasurySeed(ctx).CloudSeedRequest(cloudSeedRequest).Execute()

SeedTreasury injects bootstrap capital into the reserve fund so backed payouts can begin before the first revenue-share sweep, and returns the journal entry it wrote.



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
	cloudSeedRequest := *openapiclient.NewCloudSeedRequest() // CloudSeedRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminAPI.CloudPostV1AdminTreasurySeed(context.Background()).CloudSeedRequest(cloudSeedRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.CloudPostV1AdminTreasurySeed``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudPostV1AdminTreasurySeed`: CloudSeedOut
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.CloudPostV1AdminTreasurySeed`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudPostV1AdminTreasurySeedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cloudSeedRequest** | [**CloudSeedRequest**](CloudSeedRequest.md) |  | 

### Return type

[**CloudSeedOut**](CloudSeedOut.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPostV1AdminTreasurySweep

> CloudSweepOut CloudPostV1AdminTreasurySweep(ctx).CloudSweepRequest(cloudSweepRequest).Execute()

SweepTreasury posts the revenue-share accrual for one period — revenue into the reserve fund, at the current policy's basis points — and returns what it moved.



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
	cloudSweepRequest := *openapiclient.NewCloudSweepRequest() // CloudSweepRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminAPI.CloudPostV1AdminTreasurySweep(context.Background()).CloudSweepRequest(cloudSweepRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.CloudPostV1AdminTreasurySweep``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudPostV1AdminTreasurySweep`: CloudSweepOut
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.CloudPostV1AdminTreasurySweep`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudPostV1AdminTreasurySweepRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cloudSweepRequest** | [**CloudSweepRequest**](CloudSweepRequest.md) |  | 

### Return type

[**CloudSweepOut**](CloudSweepOut.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPutV1AdminEnablement

> CloudAdminEnablementItem CloudPutV1AdminEnablement(ctx).CloudSetEnablementBody(cloudSetEnablementBody).Execute()

SetEnablement sets one item's global enablement state — off, beta or ga — and optionally replaces the list of orgs granted its beta.



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
	cloudSetEnablementBody := *openapiclient.NewCloudSetEnablementBody() // CloudSetEnablementBody | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminAPI.CloudPutV1AdminEnablement(context.Background()).CloudSetEnablementBody(cloudSetEnablementBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.CloudPutV1AdminEnablement``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudPutV1AdminEnablement`: CloudAdminEnablementItem
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.CloudPutV1AdminEnablement`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudPutV1AdminEnablementRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cloudSetEnablementBody** | [**CloudSetEnablementBody**](CloudSetEnablementBody.md) |  | 

### Return type

[**CloudAdminEnablementItem**](CloudAdminEnablementItem.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## S3AdminInfo

> S3AdminInfo200Response S3AdminInfo(ctx).Execute()

Server information



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
	resp, r, err := apiClient.AdminAPI.S3AdminInfo(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.S3AdminInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `S3AdminInfo`: S3AdminInfo200Response
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.S3AdminInfo`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiS3AdminInfoRequest struct via the builder pattern


### Return type

[**S3AdminInfo200Response**](S3AdminInfo200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## S3AdminUsage

> S3UsageInfo S3AdminUsage(ctx).Execute()

Storage usage



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
	resp, r, err := apiClient.AdminAPI.S3AdminUsage(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.S3AdminUsage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `S3AdminUsage`: S3UsageInfo
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.S3AdminUsage`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiS3AdminUsageRequest struct via the builder pattern


### Return type

[**S3UsageInfo**](S3UsageInfo.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## S3CreateServiceAccount

> S3ServiceAccount S3CreateServiceAccount(ctx).S3CreateServiceAccountRequest(s3CreateServiceAccountRequest).Execute()

Create a service account



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
	s3CreateServiceAccountRequest := *openapiclient.NewS3CreateServiceAccountRequest() // S3CreateServiceAccountRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminAPI.S3CreateServiceAccount(context.Background()).S3CreateServiceAccountRequest(s3CreateServiceAccountRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.S3CreateServiceAccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `S3CreateServiceAccount`: S3ServiceAccount
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.S3CreateServiceAccount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiS3CreateServiceAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **s3CreateServiceAccountRequest** | [**S3CreateServiceAccountRequest**](S3CreateServiceAccountRequest.md) |  | 

### Return type

[**S3ServiceAccount**](S3ServiceAccount.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## S3ListServiceAccounts

> S3ListServiceAccounts200Response S3ListServiceAccounts(ctx).Execute()

List service accounts

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
	resp, r, err := apiClient.AdminAPI.S3ListServiceAccounts(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.S3ListServiceAccounts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `S3ListServiceAccounts`: S3ListServiceAccounts200Response
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.S3ListServiceAccounts`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiS3ListServiceAccountsRequest struct via the builder pattern


### Return type

[**S3ListServiceAccounts200Response**](S3ListServiceAccounts200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

