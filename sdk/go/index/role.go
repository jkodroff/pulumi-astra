// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package index

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// `Role` resource represents custom roles for a particular Astra Org. Custom roles can be assigned to an Astra user is to grant them granular permissions when the default roles in the UI are not specific enough. Roles are composed of policies which are granted to resources.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/mapped/pulumi-index/sdk/go/index"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := astra.NewRole(ctx, "example", &astra.RoleArgs{
// 			Description: pulumi.String("test role"),
// 			Effect:      pulumi.String("allow"),
// 			Policies: pulumi.StringArray{
// 				pulumi.String("db-all-keyspace-create"),
// 			},
// 			Resources: pulumi.StringArray{
// 				pulumi.String("drn:astra:org:f9f4b1e0-4c05-451e-9bba-d631295a7f73"),
// 			},
// 			RoleName: pulumi.String("puppies"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = astra.NewRole(ctx, "example2", &astra.RoleArgs{
// 			Description: pulumi.String("complex role"),
// 			Effect:      pulumi.String("allow"),
// 			Policies: pulumi.StringArray{
// 				pulumi.String("accesslist-read"),
// 				pulumi.String("db-all-keyspace-describe"),
// 				pulumi.String("db-keyspace-describe"),
// 				pulumi.String("db-table-select"),
// 				pulumi.String("db-table-describe"),
// 				pulumi.String("db-graphql"),
// 				pulumi.String("db-rest"),
// 				pulumi.String("db-cql"),
// 			},
// 			Resources: pulumi.StringArray{
// 				pulumi.String("drn:astra:org:f9f4b1e0-4c05-451e-9bba-d631295a7f73"),
// 				pulumi.String("drn:astra:org:f9f4b1e0-4c05-451e-9bba-d631295a7f73:db:5b70892f-e01a-4595-98e6-19ecc9985d50"),
// 				pulumi.String("drn:astra:org:f9f4b1e0-4c05-451e-9bba-d631295a7f73:db:5b70892f-e01a-4595-98e6-19ecc9985d50:keyspace:system_schema:table:*"),
// 				pulumi.String("drn:astra:org:f9f4b1e0-4c05-451e-9bba-d631295a7f73:db:5b70892f-e01a-4595-98e6-19ecc9985d50:keyspace:system:table:*"),
// 				pulumi.String("drn:astra:org:f9f4b1e0-4c05-451e-9bba-d631295a7f73:db:5b70892f-e01a-4595-98e6-19ecc9985d50:keyspace:system_virtual_schema:table:*"),
// 				pulumi.String("drn:astra:org:f9f4b1e0-4c05-451e-9bba-d631295a7f73:db:5b70892f-e01a-4595-98e6-19ecc9985d50:keyspace:*"),
// 				pulumi.String("drn:astra:org:f9f4b1e0-4c05-451e-9bba-d631295a7f73:db:5b70892f-e01a-4595-98e6-19ecc9985d50:keyspace:*:table:*"),
// 			},
// 			RoleName: pulumi.String("puppies"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## Import
//
// ```sh
//  $ pulumi import astra:index/role:Role example role-id
// ```
type Role struct {
	pulumi.CustomResourceState

	// Role description
	Description pulumi.StringOutput `pulumi:"description"`
	// Role effect
	Effect pulumi.StringOutput `pulumi:"effect"`
	// List of policies for the role. See
	// https://docs.datastax.com/en/astra/docs/user-permissions.html#_operational_roles_detail for supported policies.
	Policies pulumi.StringArrayOutput `pulumi:"policies"`
	// Resources for which role is applicable (format is "drn:astra:org:<org UUID>", followed by optional resource criteria.
	// See example usage above).
	Resources pulumi.StringArrayOutput `pulumi:"resources"`
	// Role ID, system generated
	RoleId pulumi.StringOutput `pulumi:"roleId"`
	// Role name
	RoleName pulumi.StringOutput `pulumi:"roleName"`
}

// NewRole registers a new resource with the given unique name, arguments, and options.
func NewRole(ctx *pulumi.Context,
	name string, args *RoleArgs, opts ...pulumi.ResourceOption) (*Role, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Description == nil {
		return nil, errors.New("invalid value for required argument 'Description'")
	}
	if args.Effect == nil {
		return nil, errors.New("invalid value for required argument 'Effect'")
	}
	if args.Policies == nil {
		return nil, errors.New("invalid value for required argument 'Policies'")
	}
	if args.Resources == nil {
		return nil, errors.New("invalid value for required argument 'Resources'")
	}
	if args.RoleName == nil {
		return nil, errors.New("invalid value for required argument 'RoleName'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource Role
	err := ctx.RegisterResource("astra:index/role:Role", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRole gets an existing Role resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRole(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RoleState, opts ...pulumi.ResourceOption) (*Role, error) {
	var resource Role
	err := ctx.ReadResource("astra:index/role:Role", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Role resources.
type roleState struct {
	// Role description
	Description *string `pulumi:"description"`
	// Role effect
	Effect *string `pulumi:"effect"`
	// List of policies for the role. See
	// https://docs.datastax.com/en/astra/docs/user-permissions.html#_operational_roles_detail for supported policies.
	Policies []string `pulumi:"policies"`
	// Resources for which role is applicable (format is "drn:astra:org:<org UUID>", followed by optional resource criteria.
	// See example usage above).
	Resources []string `pulumi:"resources"`
	// Role ID, system generated
	RoleId *string `pulumi:"roleId"`
	// Role name
	RoleName *string `pulumi:"roleName"`
}

type RoleState struct {
	// Role description
	Description pulumi.StringPtrInput
	// Role effect
	Effect pulumi.StringPtrInput
	// List of policies for the role. See
	// https://docs.datastax.com/en/astra/docs/user-permissions.html#_operational_roles_detail for supported policies.
	Policies pulumi.StringArrayInput
	// Resources for which role is applicable (format is "drn:astra:org:<org UUID>", followed by optional resource criteria.
	// See example usage above).
	Resources pulumi.StringArrayInput
	// Role ID, system generated
	RoleId pulumi.StringPtrInput
	// Role name
	RoleName pulumi.StringPtrInput
}

func (RoleState) ElementType() reflect.Type {
	return reflect.TypeOf((*roleState)(nil)).Elem()
}

type roleArgs struct {
	// Role description
	Description string `pulumi:"description"`
	// Role effect
	Effect string `pulumi:"effect"`
	// List of policies for the role. See
	// https://docs.datastax.com/en/astra/docs/user-permissions.html#_operational_roles_detail for supported policies.
	Policies []string `pulumi:"policies"`
	// Resources for which role is applicable (format is "drn:astra:org:<org UUID>", followed by optional resource criteria.
	// See example usage above).
	Resources []string `pulumi:"resources"`
	// Role name
	RoleName string `pulumi:"roleName"`
}

// The set of arguments for constructing a Role resource.
type RoleArgs struct {
	// Role description
	Description pulumi.StringInput
	// Role effect
	Effect pulumi.StringInput
	// List of policies for the role. See
	// https://docs.datastax.com/en/astra/docs/user-permissions.html#_operational_roles_detail for supported policies.
	Policies pulumi.StringArrayInput
	// Resources for which role is applicable (format is "drn:astra:org:<org UUID>", followed by optional resource criteria.
	// See example usage above).
	Resources pulumi.StringArrayInput
	// Role name
	RoleName pulumi.StringInput
}

func (RoleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*roleArgs)(nil)).Elem()
}

type RoleInput interface {
	pulumi.Input

	ToRoleOutput() RoleOutput
	ToRoleOutputWithContext(ctx context.Context) RoleOutput
}

func (*Role) ElementType() reflect.Type {
	return reflect.TypeOf((**Role)(nil)).Elem()
}

func (i *Role) ToRoleOutput() RoleOutput {
	return i.ToRoleOutputWithContext(context.Background())
}

func (i *Role) ToRoleOutputWithContext(ctx context.Context) RoleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RoleOutput)
}

// RoleArrayInput is an input type that accepts RoleArray and RoleArrayOutput values.
// You can construct a concrete instance of `RoleArrayInput` via:
//
//          RoleArray{ RoleArgs{...} }
type RoleArrayInput interface {
	pulumi.Input

	ToRoleArrayOutput() RoleArrayOutput
	ToRoleArrayOutputWithContext(context.Context) RoleArrayOutput
}

type RoleArray []RoleInput

func (RoleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Role)(nil)).Elem()
}

func (i RoleArray) ToRoleArrayOutput() RoleArrayOutput {
	return i.ToRoleArrayOutputWithContext(context.Background())
}

func (i RoleArray) ToRoleArrayOutputWithContext(ctx context.Context) RoleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RoleArrayOutput)
}

// RoleMapInput is an input type that accepts RoleMap and RoleMapOutput values.
// You can construct a concrete instance of `RoleMapInput` via:
//
//          RoleMap{ "key": RoleArgs{...} }
type RoleMapInput interface {
	pulumi.Input

	ToRoleMapOutput() RoleMapOutput
	ToRoleMapOutputWithContext(context.Context) RoleMapOutput
}

type RoleMap map[string]RoleInput

func (RoleMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Role)(nil)).Elem()
}

func (i RoleMap) ToRoleMapOutput() RoleMapOutput {
	return i.ToRoleMapOutputWithContext(context.Background())
}

func (i RoleMap) ToRoleMapOutputWithContext(ctx context.Context) RoleMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RoleMapOutput)
}

type RoleOutput struct{ *pulumi.OutputState }

func (RoleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Role)(nil)).Elem()
}

func (o RoleOutput) ToRoleOutput() RoleOutput {
	return o
}

func (o RoleOutput) ToRoleOutputWithContext(ctx context.Context) RoleOutput {
	return o
}

// Role description
func (o RoleOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v *Role) pulumi.StringOutput { return v.Description }).(pulumi.StringOutput)
}

// Role effect
func (o RoleOutput) Effect() pulumi.StringOutput {
	return o.ApplyT(func(v *Role) pulumi.StringOutput { return v.Effect }).(pulumi.StringOutput)
}

// List of policies for the role. See
// https://docs.datastax.com/en/astra/docs/user-permissions.html#_operational_roles_detail for supported policies.
func (o RoleOutput) Policies() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Role) pulumi.StringArrayOutput { return v.Policies }).(pulumi.StringArrayOutput)
}

// Resources for which role is applicable (format is "drn:astra:org:<org UUID>", followed by optional resource criteria.
// See example usage above).
func (o RoleOutput) Resources() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Role) pulumi.StringArrayOutput { return v.Resources }).(pulumi.StringArrayOutput)
}

// Role ID, system generated
func (o RoleOutput) RoleId() pulumi.StringOutput {
	return o.ApplyT(func(v *Role) pulumi.StringOutput { return v.RoleId }).(pulumi.StringOutput)
}

// Role name
func (o RoleOutput) RoleName() pulumi.StringOutput {
	return o.ApplyT(func(v *Role) pulumi.StringOutput { return v.RoleName }).(pulumi.StringOutput)
}

type RoleArrayOutput struct{ *pulumi.OutputState }

func (RoleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Role)(nil)).Elem()
}

func (o RoleArrayOutput) ToRoleArrayOutput() RoleArrayOutput {
	return o
}

func (o RoleArrayOutput) ToRoleArrayOutputWithContext(ctx context.Context) RoleArrayOutput {
	return o
}

func (o RoleArrayOutput) Index(i pulumi.IntInput) RoleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Role {
		return vs[0].([]*Role)[vs[1].(int)]
	}).(RoleOutput)
}

type RoleMapOutput struct{ *pulumi.OutputState }

func (RoleMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Role)(nil)).Elem()
}

func (o RoleMapOutput) ToRoleMapOutput() RoleMapOutput {
	return o
}

func (o RoleMapOutput) ToRoleMapOutputWithContext(ctx context.Context) RoleMapOutput {
	return o
}

func (o RoleMapOutput) MapIndex(k pulumi.StringInput) RoleOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Role {
		return vs[0].(map[string]*Role)[vs[1].(string)]
	}).(RoleOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RoleInput)(nil)).Elem(), &Role{})
	pulumi.RegisterInputType(reflect.TypeOf((*RoleArrayInput)(nil)).Elem(), RoleArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*RoleMapInput)(nil)).Elem(), RoleMap{})
	pulumi.RegisterOutputType(RoleOutput{})
	pulumi.RegisterOutputType(RoleArrayOutput{})
	pulumi.RegisterOutputType(RoleMapOutput{})
}
