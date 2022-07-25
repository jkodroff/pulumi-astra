// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package astra

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// `AccessList` provides a datasource that lists the access lists for an Astra database.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-astra/sdk/go/astra"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// 	"github.com/pulumiverse/pulumi-index/sdk/go/index"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := astra.LookupAccessList(ctx, &GetAccessListArgs{
// 			DatabaseId: "8d356587-73b3-430a-9c0e-d780332e2afb",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func LookupAccessList(ctx *pulumi.Context, args *LookupAccessListArgs, opts ...pulumi.InvokeOption) (*LookupAccessListResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv LookupAccessListResult
	err := ctx.Invoke("astra:index/getAccessList:getAccessList", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getAccessList.
type LookupAccessListArgs struct {
	DatabaseId string `pulumi:"databaseId"`
}

// A collection of values returned by getAccessList.
type LookupAccessListResult struct {
	Addresses  []GetAccessListAddress `pulumi:"addresses"`
	DatabaseId string                 `pulumi:"databaseId"`
	Enabled    bool                   `pulumi:"enabled"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
}

func LookupAccessListOutput(ctx *pulumi.Context, args LookupAccessListOutputArgs, opts ...pulumi.InvokeOption) LookupAccessListResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupAccessListResult, error) {
			args := v.(LookupAccessListArgs)
			r, err := LookupAccessList(ctx, &args, opts...)
			var s LookupAccessListResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupAccessListResultOutput)
}

// A collection of arguments for invoking getAccessList.
type LookupAccessListOutputArgs struct {
	DatabaseId pulumi.StringInput `pulumi:"databaseId"`
}

func (LookupAccessListOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAccessListArgs)(nil)).Elem()
}

// A collection of values returned by getAccessList.
type LookupAccessListResultOutput struct{ *pulumi.OutputState }

func (LookupAccessListResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAccessListResult)(nil)).Elem()
}

func (o LookupAccessListResultOutput) ToLookupAccessListResultOutput() LookupAccessListResultOutput {
	return o
}

func (o LookupAccessListResultOutput) ToLookupAccessListResultOutputWithContext(ctx context.Context) LookupAccessListResultOutput {
	return o
}

func (o LookupAccessListResultOutput) Addresses() GetAccessListAddressArrayOutput {
	return o.ApplyT(func(v LookupAccessListResult) []GetAccessListAddress { return v.Addresses }).(GetAccessListAddressArrayOutput)
}

func (o LookupAccessListResultOutput) DatabaseId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAccessListResult) string { return v.DatabaseId }).(pulumi.StringOutput)
}

func (o LookupAccessListResultOutput) Enabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupAccessListResult) bool { return v.Enabled }).(pulumi.BoolOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupAccessListResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAccessListResult) string { return v.Id }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupAccessListResultOutput{})
}
