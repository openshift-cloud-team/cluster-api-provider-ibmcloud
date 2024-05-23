// Code generated by smithy-go-codegen DO NOT EDIT.

package codebuild

import (
	"context"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	internalauth "github.com/aws/aws-sdk-go-v2/internal/auth"
	"github.com/aws/aws-sdk-go-v2/service/codebuild/types"
	smithyendpoints "github.com/aws/smithy-go/endpoints"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a build project.
func (c *Client) CreateProject(ctx context.Context, params *CreateProjectInput, optFns ...func(*Options)) (*CreateProjectOutput, error) {
	if params == nil {
		params = &CreateProjectInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateProject", params, optFns, c.addOperationCreateProjectMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateProjectOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateProjectInput struct {

	// Information about the build output artifacts for the build project.
	//
	// This member is required.
	Artifacts *types.ProjectArtifacts

	// Information about the build environment for the build project.
	//
	// This member is required.
	Environment *types.ProjectEnvironment

	// The name of the build project.
	//
	// This member is required.
	Name *string

	// The ARN of the IAM role that enables CodeBuild to interact with dependent
	// Amazon Web Services services on behalf of the Amazon Web Services account.
	//
	// This member is required.
	ServiceRole *string

	// Information about the build input source code for the build project.
	//
	// This member is required.
	Source *types.ProjectSource

	// Set this to true to generate a publicly accessible URL for your project's build
	// badge.
	BadgeEnabled *bool

	// A ProjectBuildBatchConfig object that defines the batch build options for the
	// project.
	BuildBatchConfig *types.ProjectBuildBatchConfig

	// Stores recently used information so that it can be quickly accessed at a later
	// time.
	Cache *types.ProjectCache

	// The maximum number of concurrent builds that are allowed for this project. New
	// builds are only started if the current number of builds is less than or equal to
	// this limit. If the current build count meets this limit, new builds are
	// throttled and are not run.
	ConcurrentBuildLimit *int32

	// A description that makes the build project easy to identify.
	Description *string

	// The Key Management Service customer master key (CMK) to be used for encrypting
	// the build output artifacts. You can use a cross-account KMS key to encrypt the
	// build output artifacts if your service role has permission to that key. You can
	// specify either the Amazon Resource Name (ARN) of the CMK or, if available, the
	// CMK's alias (using the format alias/ ).
	EncryptionKey *string

	// An array of ProjectFileSystemLocation objects for a CodeBuild build project. A
	// ProjectFileSystemLocation object specifies the identifier , location ,
	// mountOptions , mountPoint , and type of a file system created using Amazon
	// Elastic File System.
	FileSystemLocations []types.ProjectFileSystemLocation

	// Information about logs for the build project. These can be logs in CloudWatch
	// Logs, logs uploaded to a specified S3 bucket, or both.
	LogsConfig *types.LogsConfig

	// The number of minutes a build is allowed to be queued before it times out.
	QueuedTimeoutInMinutes *int32

	// An array of ProjectArtifacts objects.
	SecondaryArtifacts []types.ProjectArtifacts

	// An array of ProjectSourceVersion objects. If secondarySourceVersions is
	// specified at the build level, then they take precedence over these
	// secondarySourceVersions (at the project level).
	SecondarySourceVersions []types.ProjectSourceVersion

	// An array of ProjectSource objects.
	SecondarySources []types.ProjectSource

	// A version of the build input to be built for this project. If not specified,
	// the latest version is used. If specified, it must be one of:
	//   - For CodeCommit: the commit ID, branch, or Git tag to use.
	//   - For GitHub: the commit ID, pull request ID, branch name, or tag name that
	//   corresponds to the version of the source code you want to build. If a pull
	//   request ID is specified, it must use the format pr/pull-request-ID (for
	//   example pr/25 ). If a branch name is specified, the branch's HEAD commit ID is
	//   used. If not specified, the default branch's HEAD commit ID is used.
	//   - For Bitbucket: the commit ID, branch name, or tag name that corresponds to
	//   the version of the source code you want to build. If a branch name is specified,
	//   the branch's HEAD commit ID is used. If not specified, the default branch's HEAD
	//   commit ID is used.
	//   - For Amazon S3: the version ID of the object that represents the build input
	//   ZIP file to use.
	// If sourceVersion is specified at the build level, then that version takes
	// precedence over this sourceVersion (at the project level). For more
	// information, see Source Version Sample with CodeBuild (https://docs.aws.amazon.com/codebuild/latest/userguide/sample-source-version.html)
	// in the CodeBuild User Guide.
	SourceVersion *string

	// A list of tag key and value pairs associated with this build project. These
	// tags are available for use by Amazon Web Services services that support
	// CodeBuild build project tags.
	Tags []types.Tag

	// How long, in minutes, from 5 to 480 (8 hours), for CodeBuild to wait before it
	// times out any build that has not been marked as completed. The default is 60
	// minutes.
	TimeoutInMinutes *int32

	// VpcConfig enables CodeBuild to access resources in an Amazon VPC.
	VpcConfig *types.VpcConfig

	noSmithyDocumentSerde
}

type CreateProjectOutput struct {

	// Information about the build project that was created.
	Project *types.Project

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateProjectMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateProject{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateProject{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addlegacyEndpointContextSetter(stack, options); err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack, options); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addCreateProjectResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreateProjectValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateProject(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecursionDetection(stack); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	if err = addendpointDisableHTTPSMiddleware(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opCreateProject(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "codebuild",
		OperationName: "CreateProject",
	}
}

type opCreateProjectResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver  builtInParameterResolver
}

func (*opCreateProjectResolveEndpointMiddleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *opCreateProjectResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	if awsmiddleware.GetRequiresLegacyEndpoints(ctx) {
		return next.HandleSerialize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	if m.EndpointResolver == nil {
		return out, metadata, fmt.Errorf("expected endpoint resolver to not be nil")
	}

	params := EndpointParameters{}

	m.BuiltInResolver.ResolveBuiltIns(&params)

	var resolvedEndpoint smithyendpoints.Endpoint
	resolvedEndpoint, err = m.EndpointResolver.ResolveEndpoint(ctx, params)
	if err != nil {
		return out, metadata, fmt.Errorf("failed to resolve service endpoint, %w", err)
	}

	req.URL = &resolvedEndpoint.URI

	for k := range resolvedEndpoint.Headers {
		req.Header.Set(
			k,
			resolvedEndpoint.Headers.Get(k),
		)
	}

	authSchemes, err := internalauth.GetAuthenticationSchemes(&resolvedEndpoint.Properties)
	if err != nil {
		var nfe *internalauth.NoAuthenticationSchemesFoundError
		if errors.As(err, &nfe) {
			// if no auth scheme is found, default to sigv4
			signingName := "codebuild"
			signingRegion := m.BuiltInResolver.(*builtInResolver).Region
			ctx = awsmiddleware.SetSigningName(ctx, signingName)
			ctx = awsmiddleware.SetSigningRegion(ctx, signingRegion)

		}
		var ue *internalauth.UnSupportedAuthenticationSchemeSpecifiedError
		if errors.As(err, &ue) {
			return out, metadata, fmt.Errorf(
				"This operation requests signer version(s) %v but the client only supports %v",
				ue.UnsupportedSchemes,
				internalauth.SupportedSchemes,
			)
		}
	}

	for _, authScheme := range authSchemes {
		switch authScheme.(type) {
		case *internalauth.AuthenticationSchemeV4:
			v4Scheme, _ := authScheme.(*internalauth.AuthenticationSchemeV4)
			var signingName, signingRegion string
			if v4Scheme.SigningName == nil {
				signingName = "codebuild"
			} else {
				signingName = *v4Scheme.SigningName
			}
			if v4Scheme.SigningRegion == nil {
				signingRegion = m.BuiltInResolver.(*builtInResolver).Region
			} else {
				signingRegion = *v4Scheme.SigningRegion
			}
			if v4Scheme.DisableDoubleEncoding != nil {
				// The signer sets an equivalent value at client initialization time.
				// Setting this context value will cause the signer to extract it
				// and override the value set at client initialization time.
				ctx = internalauth.SetDisableDoubleEncoding(ctx, *v4Scheme.DisableDoubleEncoding)
			}
			ctx = awsmiddleware.SetSigningName(ctx, signingName)
			ctx = awsmiddleware.SetSigningRegion(ctx, signingRegion)
			break
		case *internalauth.AuthenticationSchemeV4A:
			v4aScheme, _ := authScheme.(*internalauth.AuthenticationSchemeV4A)
			if v4aScheme.SigningName == nil {
				v4aScheme.SigningName = aws.String("codebuild")
			}
			if v4aScheme.DisableDoubleEncoding != nil {
				// The signer sets an equivalent value at client initialization time.
				// Setting this context value will cause the signer to extract it
				// and override the value set at client initialization time.
				ctx = internalauth.SetDisableDoubleEncoding(ctx, *v4aScheme.DisableDoubleEncoding)
			}
			ctx = awsmiddleware.SetSigningName(ctx, *v4aScheme.SigningName)
			ctx = awsmiddleware.SetSigningRegion(ctx, v4aScheme.SigningRegionSet[0])
			break
		case *internalauth.AuthenticationSchemeNone:
			break
		}
	}

	return next.HandleSerialize(ctx, in)
}

func addCreateProjectResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
	return stack.Serialize.Insert(&opCreateProjectResolveEndpointMiddleware{
		EndpointResolver: options.EndpointResolverV2,
		BuiltInResolver: &builtInResolver{
			Region:       options.Region,
			UseDualStack: options.EndpointOptions.UseDualStackEndpoint,
			UseFIPS:      options.EndpointOptions.UseFIPSEndpoint,
			Endpoint:     options.BaseEndpoint,
		},
	}, "ResolveEndpoint", middleware.After)
}
