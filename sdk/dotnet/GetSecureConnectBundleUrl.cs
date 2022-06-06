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
    public static class GetSecureConnectBundleUrl
    {
        /// <summary>
        /// `astra.getSecureConnectBundleUrl` provides a datasource that generates a temporary secure connect bundle URL. This URL lasts five minutes. Secure connect bundles are used to connect to Astra using cql cassandra drivers. See the [docs](https://docs.datastax.com/en/astra/docs/connecting-to-database.html) for more information on how to connect.
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
        ///         var dev = Output.Create(Astra.GetSecureConnectBundleUrl.InvokeAsync(new Astra.GetSecureConnectBundleUrlArgs
        ///         {
        ///             DatabaseId = "f9f4b1e0-4c05-451e-9bba-d631295a7f73",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetSecureConnectBundleUrlResult> InvokeAsync(GetSecureConnectBundleUrlArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetSecureConnectBundleUrlResult>("astra:index/getSecureConnectBundleUrl:getSecureConnectBundleUrl", args ?? new GetSecureConnectBundleUrlArgs(), options.WithDefaults());

        /// <summary>
        /// `astra.getSecureConnectBundleUrl` provides a datasource that generates a temporary secure connect bundle URL. This URL lasts five minutes. Secure connect bundles are used to connect to Astra using cql cassandra drivers. See the [docs](https://docs.datastax.com/en/astra/docs/connecting-to-database.html) for more information on how to connect.
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
        ///         var dev = Output.Create(Astra.GetSecureConnectBundleUrl.InvokeAsync(new Astra.GetSecureConnectBundleUrlArgs
        ///         {
        ///             DatabaseId = "f9f4b1e0-4c05-451e-9bba-d631295a7f73",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetSecureConnectBundleUrlResult> Invoke(GetSecureConnectBundleUrlInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetSecureConnectBundleUrlResult>("astra:index/getSecureConnectBundleUrl:getSecureConnectBundleUrl", args ?? new GetSecureConnectBundleUrlInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetSecureConnectBundleUrlArgs : Pulumi.InvokeArgs
    {
        [Input("databaseId", required: true)]
        public string DatabaseId { get; set; } = null!;

        public GetSecureConnectBundleUrlArgs()
        {
        }
    }

    public sealed class GetSecureConnectBundleUrlInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("databaseId", required: true)]
        public Input<string> DatabaseId { get; set; } = null!;

        public GetSecureConnectBundleUrlInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetSecureConnectBundleUrlResult
    {
        public readonly string DatabaseId;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string Url;

        [OutputConstructor]
        private GetSecureConnectBundleUrlResult(
            string databaseId,

            string id,

            string url)
        {
            DatabaseId = databaseId;
            Id = id;
            Url = url;
        }
    }
}
