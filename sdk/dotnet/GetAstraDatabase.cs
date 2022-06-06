// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Astra
{
    public static class GetAstraDatabase
    {
        /// <summary>
        /// `astra.Database` provides a datasource for Astra an Astra database. This can be used to select an existing database within your Astra Organization.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using Astra = Pulumi.Astra;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var db = Output.Create(Astra.GetAstraDatabase.InvokeAsync(new Astra.GetAstraDatabaseArgs
        ///         {
        ///             DatabaseId = "8d356587-73b3-430a-9c0e-d780332e2afb",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetAstraDatabaseResult> InvokeAsync(GetAstraDatabaseArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetAstraDatabaseResult>("astra:index/getAstraDatabase:getAstraDatabase", args ?? new GetAstraDatabaseArgs(), options.WithDefaults());

        /// <summary>
        /// `astra.Database` provides a datasource for Astra an Astra database. This can be used to select an existing database within your Astra Organization.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using Astra = Pulumi.Astra;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var db = Output.Create(Astra.GetAstraDatabase.InvokeAsync(new Astra.GetAstraDatabaseArgs
        ///         {
        ///             DatabaseId = "8d356587-73b3-430a-9c0e-d780332e2afb",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetAstraDatabaseResult> Invoke(GetAstraDatabaseInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetAstraDatabaseResult>("astra:index/getAstraDatabase:getAstraDatabase", args ?? new GetAstraDatabaseInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetAstraDatabaseArgs : Pulumi.InvokeArgs
    {
        [Input("databaseId", required: true)]
        public string DatabaseId { get; set; } = null!;

        public GetAstraDatabaseArgs()
        {
        }
    }

    public sealed class GetAstraDatabaseInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("databaseId", required: true)]
        public Input<string> DatabaseId { get; set; } = null!;

        public GetAstraDatabaseInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetAstraDatabaseResult
    {
        public readonly ImmutableArray<string> AdditionalKeyspaces;
        public readonly string CloudProvider;
        public readonly string CqlshUrl;
        public readonly string DataEndpointUrl;
        public readonly string DatabaseId;
        public readonly string GrafanaUrl;
        public readonly string GraphqlUrl;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string Keyspace;
        public readonly string Name;
        public readonly int NodeCount;
        public readonly string OrganizationId;
        public readonly string OwnerId;
        public readonly ImmutableArray<string> Regions;
        public readonly int ReplicationFactor;
        public readonly string Status;
        public readonly int TotalStorage;

        [OutputConstructor]
        private GetAstraDatabaseResult(
            ImmutableArray<string> additionalKeyspaces,

            string cloudProvider,

            string cqlshUrl,

            string dataEndpointUrl,

            string databaseId,

            string grafanaUrl,

            string graphqlUrl,

            string id,

            string keyspace,

            string name,

            int nodeCount,

            string organizationId,

            string ownerId,

            ImmutableArray<string> regions,

            int replicationFactor,

            string status,

            int totalStorage)
        {
            AdditionalKeyspaces = additionalKeyspaces;
            CloudProvider = cloudProvider;
            CqlshUrl = cqlshUrl;
            DataEndpointUrl = dataEndpointUrl;
            DatabaseId = databaseId;
            GrafanaUrl = grafanaUrl;
            GraphqlUrl = graphqlUrl;
            Id = id;
            Keyspace = keyspace;
            Name = name;
            NodeCount = nodeCount;
            OrganizationId = organizationId;
            OwnerId = ownerId;
            Regions = regions;
            ReplicationFactor = replicationFactor;
            Status = status;
            TotalStorage = totalStorage;
        }
    }
}
