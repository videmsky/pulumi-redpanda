// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Redpanda
{
    public static class GetNetwork
    {
        /// <summary>
        /// Data source for a Redpanda Cloud network
        /// 
        /// 
        /// ## Usage
        /// 
        /// &lt;!--Start PulumiCodeChooser --&gt;
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Redpanda = Pulumi.Redpanda;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var example = Redpanda.GetNetwork.Invoke(new()
        ///     {
        ///         Id = "network_id",
        ///     });
        /// 
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Task<GetNetworkResult> InvokeAsync(GetNetworkArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetNetworkResult>("redpanda:index/getNetwork:getNetwork", args ?? new GetNetworkArgs(), options.WithDefaults());

        /// <summary>
        /// Data source for a Redpanda Cloud network
        /// 
        /// 
        /// ## Usage
        /// 
        /// &lt;!--Start PulumiCodeChooser --&gt;
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Redpanda = Pulumi.Redpanda;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var example = Redpanda.GetNetwork.Invoke(new()
        ///     {
        ///         Id = "network_id",
        ///     });
        /// 
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Output<GetNetworkResult> Invoke(GetNetworkInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetNetworkResult>("redpanda:index/getNetwork:getNetwork", args ?? new GetNetworkInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetNetworkArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// UUID of the network
        /// </summary>
        [Input("id", required: true)]
        public string Id { get; set; } = null!;

        public GetNetworkArgs()
        {
        }
        public static new GetNetworkArgs Empty => new GetNetworkArgs();
    }

    public sealed class GetNetworkInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// UUID of the network
        /// </summary>
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        public GetNetworkInvokeArgs()
        {
        }
        public static new GetNetworkInvokeArgs Empty => new GetNetworkInvokeArgs();
    }


    [OutputType]
    public sealed class GetNetworkResult
    {
        /// <summary>
        /// The cidr_block to create the network in
        /// </summary>
        public readonly string CidrBlock;
        /// <summary>
        /// The cloud provider to create the network in. Can also be set at the provider level
        /// </summary>
        public readonly string CloudProvider;
        /// <summary>
        /// The type of cluster this network is associated with, can be one of dedicated or cloud
        /// </summary>
        public readonly string ClusterType;
        /// <summary>
        /// UUID of the network
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Name of the network
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The id of the namespace in which to create the network
        /// </summary>
        public readonly string NamespaceId;
        /// <summary>
        /// The region to create the network in. Can also be set at the provider level
        /// </summary>
        public readonly string Region;

        [OutputConstructor]
        private GetNetworkResult(
            string cidrBlock,

            string cloudProvider,

            string clusterType,

            string id,

            string name,

            string namespaceId,

            string region)
        {
            CidrBlock = cidrBlock;
            CloudProvider = cloudProvider;
            ClusterType = clusterType;
            Id = id;
            Name = name;
            NamespaceId = namespaceId;
            Region = region;
        }
    }
}
