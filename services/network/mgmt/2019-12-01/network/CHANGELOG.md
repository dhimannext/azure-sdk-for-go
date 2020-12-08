Generated from https://github.com/Azure/azure-rest-api-specs/tree/3c764635e7d442b3e74caf593029fcd440b3ef82

Code generator @microsoft.azure/autorest.go@~2.1.161

## Breaking Changes

- Function `NewRouteListResultPage` parameter(s) have been changed from `(func(context.Context, RouteListResult) (RouteListResult, error))` to `(RouteListResult, func(context.Context, RouteListResult) (RouteListResult, error))`
- Function `NewIPGroupListResultPage` parameter(s) have been changed from `(func(context.Context, IPGroupListResult) (IPGroupListResult, error))` to `(IPGroupListResult, func(context.Context, IPGroupListResult) (IPGroupListResult, error))`
- Function `NewPublicIPPrefixListResultPage` parameter(s) have been changed from `(func(context.Context, PublicIPPrefixListResult) (PublicIPPrefixListResult, error))` to `(PublicIPPrefixListResult, func(context.Context, PublicIPPrefixListResult) (PublicIPPrefixListResult, error))`
- Function `NewOperationListResultPage` parameter(s) have been changed from `(func(context.Context, OperationListResult) (OperationListResult, error))` to `(OperationListResult, func(context.Context, OperationListResult) (OperationListResult, error))`
- Function `NewVirtualNetworkGatewayListConnectionsResultPage` parameter(s) have been changed from `(func(context.Context, VirtualNetworkGatewayListConnectionsResult) (VirtualNetworkGatewayListConnectionsResult, error))` to `(VirtualNetworkGatewayListConnectionsResult, func(context.Context, VirtualNetworkGatewayListConnectionsResult) (VirtualNetworkGatewayListConnectionsResult, error))`
- Function `NewAvailablePrivateEndpointTypesResultPage` parameter(s) have been changed from `(func(context.Context, AvailablePrivateEndpointTypesResult) (AvailablePrivateEndpointTypesResult, error))` to `(AvailablePrivateEndpointTypesResult, func(context.Context, AvailablePrivateEndpointTypesResult) (AvailablePrivateEndpointTypesResult, error))`
- Function `NewSecurityGroupListResultPage` parameter(s) have been changed from `(func(context.Context, SecurityGroupListResult) (SecurityGroupListResult, error))` to `(SecurityGroupListResult, func(context.Context, SecurityGroupListResult) (SecurityGroupListResult, error))`
- Function `NewInterfaceLoadBalancerListResultPage` parameter(s) have been changed from `(func(context.Context, InterfaceLoadBalancerListResult) (InterfaceLoadBalancerListResult, error))` to `(InterfaceLoadBalancerListResult, func(context.Context, InterfaceLoadBalancerListResult) (InterfaceLoadBalancerListResult, error))`
- Function `NewApplicationSecurityGroupListResultPage` parameter(s) have been changed from `(func(context.Context, ApplicationSecurityGroupListResult) (ApplicationSecurityGroupListResult, error))` to `(ApplicationSecurityGroupListResult, func(context.Context, ApplicationSecurityGroupListResult) (ApplicationSecurityGroupListResult, error))`
- Function `NewExpressRouteLinkListResultPage` parameter(s) have been changed from `(func(context.Context, ExpressRouteLinkListResult) (ExpressRouteLinkListResult, error))` to `(ExpressRouteLinkListResult, func(context.Context, ExpressRouteLinkListResult) (ExpressRouteLinkListResult, error))`
- Function `NewListHubVirtualNetworkConnectionsResultPage` parameter(s) have been changed from `(func(context.Context, ListHubVirtualNetworkConnectionsResult) (ListHubVirtualNetworkConnectionsResult, error))` to `(ListHubVirtualNetworkConnectionsResult, func(context.Context, ListHubVirtualNetworkConnectionsResult) (ListHubVirtualNetworkConnectionsResult, error))`
- Function `NewExpressRouteServiceProviderListResultPage` parameter(s) have been changed from `(func(context.Context, ExpressRouteServiceProviderListResult) (ExpressRouteServiceProviderListResult, error))` to `(ExpressRouteServiceProviderListResult, func(context.Context, ExpressRouteServiceProviderListResult) (ExpressRouteServiceProviderListResult, error))`
- Function `NewListVpnSiteLinkConnectionsResultPage` parameter(s) have been changed from `(func(context.Context, ListVpnSiteLinkConnectionsResult) (ListVpnSiteLinkConnectionsResult, error))` to `(ListVpnSiteLinkConnectionsResult, func(context.Context, ListVpnSiteLinkConnectionsResult) (ListVpnSiteLinkConnectionsResult, error))`
- Function `NewVirtualNetworkGatewayListResultPage` parameter(s) have been changed from `(func(context.Context, VirtualNetworkGatewayListResult) (VirtualNetworkGatewayListResult, error))` to `(VirtualNetworkGatewayListResult, func(context.Context, VirtualNetworkGatewayListResult) (VirtualNetworkGatewayListResult, error))`
- Function `NewEndpointServicesListResultPage` parameter(s) have been changed from `(func(context.Context, EndpointServicesListResult) (EndpointServicesListResult, error))` to `(EndpointServicesListResult, func(context.Context, EndpointServicesListResult) (EndpointServicesListResult, error))`
- Function `NewRouteFilterListResultPage` parameter(s) have been changed from `(func(context.Context, RouteFilterListResult) (RouteFilterListResult, error))` to `(RouteFilterListResult, func(context.Context, RouteFilterListResult) (RouteFilterListResult, error))`
- Function `NewVirtualNetworkGatewayConnectionListResultPage` parameter(s) have been changed from `(func(context.Context, VirtualNetworkGatewayConnectionListResult) (VirtualNetworkGatewayConnectionListResult, error))` to `(VirtualNetworkGatewayConnectionListResult, func(context.Context, VirtualNetworkGatewayConnectionListResult) (VirtualNetworkGatewayConnectionListResult, error))`
- Function `NewAuthorizationListResultPage` parameter(s) have been changed from `(func(context.Context, AuthorizationListResult) (AuthorizationListResult, error))` to `(AuthorizationListResult, func(context.Context, AuthorizationListResult) (AuthorizationListResult, error))`
- Function `NewVirtualNetworkTapListResultPage` parameter(s) have been changed from `(func(context.Context, VirtualNetworkTapListResult) (VirtualNetworkTapListResult, error))` to `(VirtualNetworkTapListResult, func(context.Context, VirtualNetworkTapListResult) (VirtualNetworkTapListResult, error))`
- Function `NewExpressRouteCircuitPeeringListResultPage` parameter(s) have been changed from `(func(context.Context, ExpressRouteCircuitPeeringListResult) (ExpressRouteCircuitPeeringListResult, error))` to `(ExpressRouteCircuitPeeringListResult, func(context.Context, ExpressRouteCircuitPeeringListResult) (ExpressRouteCircuitPeeringListResult, error))`
- Function `NewExpressRoutePortsLocationListResultPage` parameter(s) have been changed from `(func(context.Context, ExpressRoutePortsLocationListResult) (ExpressRoutePortsLocationListResult, error))` to `(ExpressRoutePortsLocationListResult, func(context.Context, ExpressRoutePortsLocationListResult) (ExpressRoutePortsLocationListResult, error))`
- Function `NewLoadBalancerBackendAddressPoolListResultPage` parameter(s) have been changed from `(func(context.Context, LoadBalancerBackendAddressPoolListResult) (LoadBalancerBackendAddressPoolListResult, error))` to `(LoadBalancerBackendAddressPoolListResult, func(context.Context, LoadBalancerBackendAddressPoolListResult) (LoadBalancerBackendAddressPoolListResult, error))`
- Function `NewAvailableServiceAliasesResultPage` parameter(s) have been changed from `(func(context.Context, AvailableServiceAliasesResult) (AvailableServiceAliasesResult, error))` to `(AvailableServiceAliasesResult, func(context.Context, AvailableServiceAliasesResult) (AvailableServiceAliasesResult, error))`
- Function `NewLocalNetworkGatewayListResultPage` parameter(s) have been changed from `(func(context.Context, LocalNetworkGatewayListResult) (LocalNetworkGatewayListResult, error))` to `(LocalNetworkGatewayListResult, func(context.Context, LocalNetworkGatewayListResult) (LocalNetworkGatewayListResult, error))`
- Function `NewLoadBalancerFrontendIPConfigurationListResultPage` parameter(s) have been changed from `(func(context.Context, LoadBalancerFrontendIPConfigurationListResult) (LoadBalancerFrontendIPConfigurationListResult, error))` to `(LoadBalancerFrontendIPConfigurationListResult, func(context.Context, LoadBalancerFrontendIPConfigurationListResult) (LoadBalancerFrontendIPConfigurationListResult, error))`
- Function `NewWebApplicationFirewallPolicyListResultPage` parameter(s) have been changed from `(func(context.Context, WebApplicationFirewallPolicyListResult) (WebApplicationFirewallPolicyListResult, error))` to `(WebApplicationFirewallPolicyListResult, func(context.Context, WebApplicationFirewallPolicyListResult) (WebApplicationFirewallPolicyListResult, error))`
- Function `NewExpressRouteCircuitConnectionListResultPage` parameter(s) have been changed from `(func(context.Context, ExpressRouteCircuitConnectionListResult) (ExpressRouteCircuitConnectionListResult, error))` to `(ExpressRouteCircuitConnectionListResult, func(context.Context, ExpressRouteCircuitConnectionListResult) (ExpressRouteCircuitConnectionListResult, error))`
- Function `NewLoadBalancerLoadBalancingRuleListResultPage` parameter(s) have been changed from `(func(context.Context, LoadBalancerLoadBalancingRuleListResult) (LoadBalancerLoadBalancingRuleListResult, error))` to `(LoadBalancerLoadBalancingRuleListResult, func(context.Context, LoadBalancerLoadBalancingRuleListResult) (LoadBalancerLoadBalancingRuleListResult, error))`
- Function `NewPublicIPAddressListResultPage` parameter(s) have been changed from `(func(context.Context, PublicIPAddressListResult) (PublicIPAddressListResult, error))` to `(PublicIPAddressListResult, func(context.Context, PublicIPAddressListResult) (PublicIPAddressListResult, error))`
- Function `NewVirtualRouterPeeringListResultPage` parameter(s) have been changed from `(func(context.Context, VirtualRouterPeeringListResult) (VirtualRouterPeeringListResult, error))` to `(VirtualRouterPeeringListResult, func(context.Context, VirtualRouterPeeringListResult) (VirtualRouterPeeringListResult, error))`
- Function `NewInterfaceListResultPage` parameter(s) have been changed from `(func(context.Context, InterfaceListResult) (InterfaceListResult, error))` to `(InterfaceListResult, func(context.Context, InterfaceListResult) (InterfaceListResult, error))`
- Function `NewInterfaceTapConfigurationListResultPage` parameter(s) have been changed from `(func(context.Context, InterfaceTapConfigurationListResult) (InterfaceTapConfigurationListResult, error))` to `(InterfaceTapConfigurationListResult, func(context.Context, InterfaceTapConfigurationListResult) (InterfaceTapConfigurationListResult, error))`
- Function `NewDdosProtectionPlanListResultPage` parameter(s) have been changed from `(func(context.Context, DdosProtectionPlanListResult) (DdosProtectionPlanListResult, error))` to `(DdosProtectionPlanListResult, func(context.Context, DdosProtectionPlanListResult) (DdosProtectionPlanListResult, error))`
- Function `NewApplicationGatewayListResultPage` parameter(s) have been changed from `(func(context.Context, ApplicationGatewayListResult) (ApplicationGatewayListResult, error))` to `(ApplicationGatewayListResult, func(context.Context, ApplicationGatewayListResult) (ApplicationGatewayListResult, error))`
- Function `NewApplicationGatewayAvailableSslPredefinedPoliciesPage` parameter(s) have been changed from `(func(context.Context, ApplicationGatewayAvailableSslPredefinedPolicies) (ApplicationGatewayAvailableSslPredefinedPolicies, error))` to `(ApplicationGatewayAvailableSslPredefinedPolicies, func(context.Context, ApplicationGatewayAvailableSslPredefinedPolicies) (ApplicationGatewayAvailableSslPredefinedPolicies, error))`
- Function `NewRouteFilterRuleListResultPage` parameter(s) have been changed from `(func(context.Context, RouteFilterRuleListResult) (RouteFilterRuleListResult, error))` to `(RouteFilterRuleListResult, func(context.Context, RouteFilterRuleListResult) (RouteFilterRuleListResult, error))`
- Function `NewInboundNatRuleListResultPage` parameter(s) have been changed from `(func(context.Context, InboundNatRuleListResult) (InboundNatRuleListResult, error))` to `(InboundNatRuleListResult, func(context.Context, InboundNatRuleListResult) (InboundNatRuleListResult, error))`
- Function `NewExpressRoutePortListResultPage` parameter(s) have been changed from `(func(context.Context, ExpressRoutePortListResult) (ExpressRoutePortListResult, error))` to `(ExpressRoutePortListResult, func(context.Context, ExpressRoutePortListResult) (ExpressRoutePortListResult, error))`
- Function `NewPrivateLinkServiceListResultPage` parameter(s) have been changed from `(func(context.Context, PrivateLinkServiceListResult) (PrivateLinkServiceListResult, error))` to `(PrivateLinkServiceListResult, func(context.Context, PrivateLinkServiceListResult) (PrivateLinkServiceListResult, error))`
- Function `NewBastionHostListResultPage` parameter(s) have been changed from `(func(context.Context, BastionHostListResult) (BastionHostListResult, error))` to `(BastionHostListResult, func(context.Context, BastionHostListResult) (BastionHostListResult, error))`
- Function `NewBgpServiceCommunityListResultPage` parameter(s) have been changed from `(func(context.Context, BgpServiceCommunityListResult) (BgpServiceCommunityListResult, error))` to `(BgpServiceCommunityListResult, func(context.Context, BgpServiceCommunityListResult) (BgpServiceCommunityListResult, error))`
- Function `NewAvailableDelegationsResultPage` parameter(s) have been changed from `(func(context.Context, AvailableDelegationsResult) (AvailableDelegationsResult, error))` to `(AvailableDelegationsResult, func(context.Context, AvailableDelegationsResult) (AvailableDelegationsResult, error))`
- Function `NewNatGatewayListResultPage` parameter(s) have been changed from `(func(context.Context, NatGatewayListResult) (NatGatewayListResult, error))` to `(NatGatewayListResult, func(context.Context, NatGatewayListResult) (NatGatewayListResult, error))`
- Function `NewAzureFirewallListResultPage` parameter(s) have been changed from `(func(context.Context, AzureFirewallListResult) (AzureFirewallListResult, error))` to `(AzureFirewallListResult, func(context.Context, AzureFirewallListResult) (AzureFirewallListResult, error))`
- Function `NewExpressRouteCrossConnectionPeeringListPage` parameter(s) have been changed from `(func(context.Context, ExpressRouteCrossConnectionPeeringList) (ExpressRouteCrossConnectionPeeringList, error))` to `(ExpressRouteCrossConnectionPeeringList, func(context.Context, ExpressRouteCrossConnectionPeeringList) (ExpressRouteCrossConnectionPeeringList, error))`
- Function `NewAzureFirewallFqdnTagListResultPage` parameter(s) have been changed from `(func(context.Context, AzureFirewallFqdnTagListResult) (AzureFirewallFqdnTagListResult, error))` to `(AzureFirewallFqdnTagListResult, func(context.Context, AzureFirewallFqdnTagListResult) (AzureFirewallFqdnTagListResult, error))`
- Function `NewBastionShareableLinkListResultPage` parameter(s) have been changed from `(func(context.Context, BastionShareableLinkListResult) (BastionShareableLinkListResult, error))` to `(BastionShareableLinkListResult, func(context.Context, BastionShareableLinkListResult) (BastionShareableLinkListResult, error))`
- Function `NewAutoApprovedPrivateLinkServicesResultPage` parameter(s) have been changed from `(func(context.Context, AutoApprovedPrivateLinkServicesResult) (AutoApprovedPrivateLinkServicesResult, error))` to `(AutoApprovedPrivateLinkServicesResult, func(context.Context, AutoApprovedPrivateLinkServicesResult) (AutoApprovedPrivateLinkServicesResult, error))`
- Function `NewVirtualNetworkPeeringListResultPage` parameter(s) have been changed from `(func(context.Context, VirtualNetworkPeeringListResult) (VirtualNetworkPeeringListResult, error))` to `(VirtualNetworkPeeringListResult, func(context.Context, VirtualNetworkPeeringListResult) (VirtualNetworkPeeringListResult, error))`
- Function `NewListVirtualWANsResultPage` parameter(s) have been changed from `(func(context.Context, ListVirtualWANsResult) (ListVirtualWANsResult, error))` to `(ListVirtualWANsResult, func(context.Context, ListVirtualWANsResult) (ListVirtualWANsResult, error))`
- Function `NewListVpnSitesResultPage` parameter(s) have been changed from `(func(context.Context, ListVpnSitesResult) (ListVpnSitesResult, error))` to `(ListVpnSitesResult, func(context.Context, ListVpnSitesResult) (ListVpnSitesResult, error))`
- Function `NewLoadBalancerProbeListResultPage` parameter(s) have been changed from `(func(context.Context, LoadBalancerProbeListResult) (LoadBalancerProbeListResult, error))` to `(LoadBalancerProbeListResult, func(context.Context, LoadBalancerProbeListResult) (LoadBalancerProbeListResult, error))`
- Function `NewServiceEndpointPolicyDefinitionListResultPage` parameter(s) have been changed from `(func(context.Context, ServiceEndpointPolicyDefinitionListResult) (ServiceEndpointPolicyDefinitionListResult, error))` to `(ServiceEndpointPolicyDefinitionListResult, func(context.Context, ServiceEndpointPolicyDefinitionListResult) (ServiceEndpointPolicyDefinitionListResult, error))`
- Function `NewPeerExpressRouteCircuitConnectionListResultPage` parameter(s) have been changed from `(func(context.Context, PeerExpressRouteCircuitConnectionListResult) (PeerExpressRouteCircuitConnectionListResult, error))` to `(PeerExpressRouteCircuitConnectionListResult, func(context.Context, PeerExpressRouteCircuitConnectionListResult) (PeerExpressRouteCircuitConnectionListResult, error))`
- Function `NewVirtualNetworkListResultPage` parameter(s) have been changed from `(func(context.Context, VirtualNetworkListResult) (VirtualNetworkListResult, error))` to `(VirtualNetworkListResult, func(context.Context, VirtualNetworkListResult) (VirtualNetworkListResult, error))`
- Function `NewFirewallPolicyListResultPage` parameter(s) have been changed from `(func(context.Context, FirewallPolicyListResult) (FirewallPolicyListResult, error))` to `(FirewallPolicyListResult, func(context.Context, FirewallPolicyListResult) (FirewallPolicyListResult, error))`
- Function `NewExpressRouteCrossConnectionListResultPage` parameter(s) have been changed from `(func(context.Context, ExpressRouteCrossConnectionListResult) (ExpressRouteCrossConnectionListResult, error))` to `(ExpressRouteCrossConnectionListResult, func(context.Context, ExpressRouteCrossConnectionListResult) (ExpressRouteCrossConnectionListResult, error))`
- Function `NewSecurityRuleListResultPage` parameter(s) have been changed from `(func(context.Context, SecurityRuleListResult) (SecurityRuleListResult, error))` to `(SecurityRuleListResult, func(context.Context, SecurityRuleListResult) (SecurityRuleListResult, error))`
- Function `NewInterfaceIPConfigurationListResultPage` parameter(s) have been changed from `(func(context.Context, InterfaceIPConfigurationListResult) (InterfaceIPConfigurationListResult, error))` to `(InterfaceIPConfigurationListResult, func(context.Context, InterfaceIPConfigurationListResult) (InterfaceIPConfigurationListResult, error))`
- Function `NewPrivateEndpointConnectionListResultPage` parameter(s) have been changed from `(func(context.Context, PrivateEndpointConnectionListResult) (PrivateEndpointConnectionListResult, error))` to `(PrivateEndpointConnectionListResult, func(context.Context, PrivateEndpointConnectionListResult) (PrivateEndpointConnectionListResult, error))`
- Function `NewProfileListResultPage` parameter(s) have been changed from `(func(context.Context, ProfileListResult) (ProfileListResult, error))` to `(ProfileListResult, func(context.Context, ProfileListResult) (ProfileListResult, error))`
- Function `NewListVirtualHubsResultPage` parameter(s) have been changed from `(func(context.Context, ListVirtualHubsResult) (ListVirtualHubsResult, error))` to `(ListVirtualHubsResult, func(context.Context, ListVirtualHubsResult) (ListVirtualHubsResult, error))`
- Function `NewListVpnGatewaysResultPage` parameter(s) have been changed from `(func(context.Context, ListVpnGatewaysResult) (ListVpnGatewaysResult, error))` to `(ListVpnGatewaysResult, func(context.Context, ListVpnGatewaysResult) (ListVpnGatewaysResult, error))`
- Function `NewExpressRouteCircuitListResultPage` parameter(s) have been changed from `(func(context.Context, ExpressRouteCircuitListResult) (ExpressRouteCircuitListResult, error))` to `(ExpressRouteCircuitListResult, func(context.Context, ExpressRouteCircuitListResult) (ExpressRouteCircuitListResult, error))`
- Function `NewBastionActiveSessionListResultPage` parameter(s) have been changed from `(func(context.Context, BastionActiveSessionListResult) (BastionActiveSessionListResult, error))` to `(BastionActiveSessionListResult, func(context.Context, BastionActiveSessionListResult) (BastionActiveSessionListResult, error))`
- Function `NewServiceEndpointPolicyListResultPage` parameter(s) have been changed from `(func(context.Context, ServiceEndpointPolicyListResult) (ServiceEndpointPolicyListResult, error))` to `(ServiceEndpointPolicyListResult, func(context.Context, ServiceEndpointPolicyListResult) (ServiceEndpointPolicyListResult, error))`
- Function `NewListVpnServerConfigurationsResultPage` parameter(s) have been changed from `(func(context.Context, ListVpnServerConfigurationsResult) (ListVpnServerConfigurationsResult, error))` to `(ListVpnServerConfigurationsResult, func(context.Context, ListVpnServerConfigurationsResult) (ListVpnServerConfigurationsResult, error))`
- Function `NewUsagesListResultPage` parameter(s) have been changed from `(func(context.Context, UsagesListResult) (UsagesListResult, error))` to `(UsagesListResult, func(context.Context, UsagesListResult) (UsagesListResult, error))`
- Function `NewVirtualNetworkListUsageResultPage` parameter(s) have been changed from `(func(context.Context, VirtualNetworkListUsageResult) (VirtualNetworkListUsageResult, error))` to `(VirtualNetworkListUsageResult, func(context.Context, VirtualNetworkListUsageResult) (VirtualNetworkListUsageResult, error))`
- Function `NewListVirtualHubRouteTableV2sResultPage` parameter(s) have been changed from `(func(context.Context, ListVirtualHubRouteTableV2sResult) (ListVirtualHubRouteTableV2sResult, error))` to `(ListVirtualHubRouteTableV2sResult, func(context.Context, ListVirtualHubRouteTableV2sResult) (ListVirtualHubRouteTableV2sResult, error))`
- Function `NewLoadBalancerOutboundRuleListResultPage` parameter(s) have been changed from `(func(context.Context, LoadBalancerOutboundRuleListResult) (LoadBalancerOutboundRuleListResult, error))` to `(LoadBalancerOutboundRuleListResult, func(context.Context, LoadBalancerOutboundRuleListResult) (LoadBalancerOutboundRuleListResult, error))`
- Function `NewBastionSessionDeleteResultPage` parameter(s) have been changed from `(func(context.Context, BastionSessionDeleteResult) (BastionSessionDeleteResult, error))` to `(BastionSessionDeleteResult, func(context.Context, BastionSessionDeleteResult) (BastionSessionDeleteResult, error))`
- Function `NewListP2SVpnGatewaysResultPage` parameter(s) have been changed from `(func(context.Context, ListP2SVpnGatewaysResult) (ListP2SVpnGatewaysResult, error))` to `(ListP2SVpnGatewaysResult, func(context.Context, ListP2SVpnGatewaysResult) (ListP2SVpnGatewaysResult, error))`
- Function `NewListVpnConnectionsResultPage` parameter(s) have been changed from `(func(context.Context, ListVpnConnectionsResult) (ListVpnConnectionsResult, error))` to `(ListVpnConnectionsResult, func(context.Context, ListVpnConnectionsResult) (ListVpnConnectionsResult, error))`
- Function `NewFlowLogListResultPage` parameter(s) have been changed from `(func(context.Context, FlowLogListResult) (FlowLogListResult, error))` to `(FlowLogListResult, func(context.Context, FlowLogListResult) (FlowLogListResult, error))`
- Function `NewListVpnSiteLinksResultPage` parameter(s) have been changed from `(func(context.Context, ListVpnSiteLinksResult) (ListVpnSiteLinksResult, error))` to `(ListVpnSiteLinksResult, func(context.Context, ListVpnSiteLinksResult) (ListVpnSiteLinksResult, error))`
- Function `NewPrivateEndpointListResultPage` parameter(s) have been changed from `(func(context.Context, PrivateEndpointListResult) (PrivateEndpointListResult, error))` to `(PrivateEndpointListResult, func(context.Context, PrivateEndpointListResult) (PrivateEndpointListResult, error))`
- Function `NewLoadBalancerListResultPage` parameter(s) have been changed from `(func(context.Context, LoadBalancerListResult) (LoadBalancerListResult, error))` to `(LoadBalancerListResult, func(context.Context, LoadBalancerListResult) (LoadBalancerListResult, error))`
- Function `NewRouteTableListResultPage` parameter(s) have been changed from `(func(context.Context, RouteTableListResult) (RouteTableListResult, error))` to `(RouteTableListResult, func(context.Context, RouteTableListResult) (RouteTableListResult, error))`
- Function `NewVirtualApplianceListResultPage` parameter(s) have been changed from `(func(context.Context, VirtualApplianceListResult) (VirtualApplianceListResult, error))` to `(VirtualApplianceListResult, func(context.Context, VirtualApplianceListResult) (VirtualApplianceListResult, error))`
- Function `NewVirtualRouterListResultPage` parameter(s) have been changed from `(func(context.Context, VirtualRouterListResult) (VirtualRouterListResult, error))` to `(VirtualRouterListResult, func(context.Context, VirtualRouterListResult) (VirtualRouterListResult, error))`
- Function `NewFirewallPolicyRuleGroupListResultPage` parameter(s) have been changed from `(func(context.Context, FirewallPolicyRuleGroupListResult) (FirewallPolicyRuleGroupListResult, error))` to `(FirewallPolicyRuleGroupListResult, func(context.Context, FirewallPolicyRuleGroupListResult) (FirewallPolicyRuleGroupListResult, error))`
- Function `NewSubnetListResultPage` parameter(s) have been changed from `(func(context.Context, SubnetListResult) (SubnetListResult, error))` to `(SubnetListResult, func(context.Context, SubnetListResult) (SubnetListResult, error))`