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
    public static class GetAstraDatabases
    {
        /// <summary>
        /// `astra.getAstraDatabases` provides a datasource for a list of Astra databases. This can be used to select databases within your Astra Organization.
        /// </summary>
        public static Task<GetAstraDatabasesResult> InvokeAsync(GetAstraDatabasesArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetAstraDatabasesResult>("astra:index/getAstraDatabases:getAstraDatabases", args ?? new GetAstraDatabasesArgs(), options.WithDefaults());

        /// <summary>
        /// `astra.getAstraDatabases` provides a datasource for a list of Astra databases. This can be used to select databases within your Astra Organization.
        /// </summary>
        public static Output<GetAstraDatabasesResult> Invoke(GetAstraDatabasesInvokeArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetAstraDatabasesResult>("astra:index/getAstraDatabases:getAstraDatabases", args ?? new GetAstraDatabasesInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetAstraDatabasesArgs : Pulumi.InvokeArgs
    {
        [Input("cloudProvider")]
        public string? CloudProvider { get; set; }

        [Input("status")]
        public string? Status { get; set; }

        public GetAstraDatabasesArgs()
        {
        }
    }

    public sealed class GetAstraDatabasesInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("cloudProvider")]
        public Input<string>? CloudProvider { get; set; }

        [Input("status")]
        public Input<string>? Status { get; set; }

        public GetAstraDatabasesInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetAstraDatabasesResult
    {
        public readonly string? CloudProvider;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly ImmutableArray<Outputs.GetAstraDatabasesResultResult> Results;
        public readonly string? Status;

        [OutputConstructor]
        private GetAstraDatabasesResult(
            string? cloudProvider,

            string id,

            ImmutableArray<Outputs.GetAstraDatabasesResultResult> results,

            string? status)
        {
            CloudProvider = cloudProvider;
            Id = id;
            Results = results;
            Status = status;
        }
    }
}
