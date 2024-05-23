// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudtrail

import (
	"context"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	internalauth "github.com/aws/aws-sdk-go-v2/internal/auth"
	"github.com/aws/aws-sdk-go-v2/service/cloudtrail/types"
	smithyendpoints "github.com/aws/smithy-go/endpoints"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Creates a new event data store.
func (c *Client) CreateEventDataStore(ctx context.Context, params *CreateEventDataStoreInput, optFns ...func(*Options)) (*CreateEventDataStoreOutput, error) {
	if params == nil {
		params = &CreateEventDataStoreInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateEventDataStore", params, optFns, c.addOperationCreateEventDataStoreMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateEventDataStoreOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateEventDataStoreInput struct {

	// The name of the event data store.
	//
	// This member is required.
	Name *string

	// The advanced event selectors to use to select the events for the data store.
	// You can configure up to five advanced event selectors for each event data store.
	// For more information about how to use advanced event selectors to log CloudTrail
	// events, see Log events by using advanced event selectors (https://docs.aws.amazon.com/awscloudtrail/latest/userguide/logging-data-events-with-cloudtrail.html#creating-data-event-selectors-advanced)
	// in the CloudTrail User Guide. For more information about how to use advanced
	// event selectors to include Config configuration items in your event data store,
	// see Create an event data store for Config configuration items (https://docs.aws.amazon.com/awscloudtrail/latest/userguide/query-lake-cli.html#lake-cli-create-eds-config)
	// in the CloudTrail User Guide. For more information about how to use advanced
	// event selectors to include non-Amazon Web Services events in your event data
	// store, see Create an integration to log events from outside Amazon Web Services (https://docs.aws.amazon.com/awscloudtrail/latest/userguide/query-lake-cli.html#lake-cli-create-integration)
	// in the CloudTrail User Guide.
	AdvancedEventSelectors []types.AdvancedEventSelector

	// Specifies the KMS key ID to use to encrypt the events delivered by CloudTrail.
	// The value can be an alias name prefixed by alias/ , a fully specified ARN to an
	// alias, a fully specified ARN to a key, or a globally unique identifier.
	// Disabling or deleting the KMS key, or removing CloudTrail permissions on the
	// key, prevents CloudTrail from logging events to the event data store, and
	// prevents users from querying the data in the event data store that was encrypted
	// with the key. After you associate an event data store with a KMS key, the KMS
	// key cannot be removed or changed. Before you disable or delete a KMS key that
	// you are using with an event data store, delete or back up your event data store.
	// CloudTrail also supports KMS multi-Region keys. For more information about
	// multi-Region keys, see Using multi-Region keys (https://docs.aws.amazon.com/kms/latest/developerguide/multi-region-keys-overview.html)
	// in the Key Management Service Developer Guide. Examples:
	//   - alias/MyAliasName
	//   - arn:aws:kms:us-east-2:123456789012:alias/MyAliasName
	//   - arn:aws:kms:us-east-2:123456789012:key/12345678-1234-1234-1234-123456789012
	//   - 12345678-1234-1234-1234-123456789012
	KmsKeyId *string

	// Specifies whether the event data store includes events from all Regions, or
	// only from the Region in which the event data store is created.
	MultiRegionEnabled *bool

	// Specifies whether an event data store collects events logged for an
	// organization in Organizations.
	OrganizationEnabled *bool

	// The retention period of the event data store, in days. You can set a retention
	// period of up to 2557 days, the equivalent of seven years. CloudTrail Lake
	// determines whether to retain an event by checking if the eventTime of the event
	// is within the specified retention period. For example, if you set a retention
	// period of 90 days, CloudTrail will remove events when the eventTime is older
	// than 90 days. If you plan to copy trail events to this event data store, we
	// recommend that you consider both the age of the events that you want to copy as
	// well as how long you want to keep the copied events in your event data store.
	// For example, if you copy trail events that are 5 years old and specify a
	// retention period of 7 years, the event data store will retain those events for
	// two years.
	RetentionPeriod *int32

	// Specifies whether the event data store should start ingesting live events. The
	// default is true.
	StartIngestion *bool

	// A list of tags.
	TagsList []types.Tag

	// Specifies whether termination protection is enabled for the event data store.
	// If termination protection is enabled, you cannot delete the event data store
	// until termination protection is disabled.
	TerminationProtectionEnabled *bool

	noSmithyDocumentSerde
}

type CreateEventDataStoreOutput struct {

	// The advanced event selectors that were used to select the events for the data
	// store.
	AdvancedEventSelectors []types.AdvancedEventSelector

	// The timestamp that shows when the event data store was created.
	CreatedTimestamp *time.Time

	// The ARN of the event data store.
	EventDataStoreArn *string

	// Specifies the KMS key ID that encrypts the events delivered by CloudTrail. The
	// value is a fully specified ARN to a KMS key in the following format.
	// arn:aws:kms:us-east-2:123456789012:key/12345678-1234-1234-1234-123456789012
	KmsKeyId *string

	// Indicates whether the event data store collects events from all Regions, or
	// only from the Region in which it was created.
	MultiRegionEnabled *bool

	// The name of the event data store.
	Name *string

	// Indicates whether an event data store is collecting logged events for an
	// organization in Organizations.
	OrganizationEnabled *bool

	// The retention period of an event data store, in days.
	RetentionPeriod *int32

	// The status of event data store creation.
	Status types.EventDataStoreStatus

	// A list of tags.
	TagsList []types.Tag

	// Indicates whether termination protection is enabled for the event data store.
	TerminationProtectionEnabled *bool

	// The timestamp that shows when an event data store was updated, if applicable.
	// UpdatedTimestamp is always either the same or newer than the time shown in
	// CreatedTimestamp .
	UpdatedTimestamp *time.Time

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateEventDataStoreMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateEventDataStore{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateEventDataStore{}, middleware.After)
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
	if err = addCreateEventDataStoreResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreateEventDataStoreValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateEventDataStore(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateEventDataStore(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cloudtrail",
		OperationName: "CreateEventDataStore",
	}
}

type opCreateEventDataStoreResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver  builtInParameterResolver
}

func (*opCreateEventDataStoreResolveEndpointMiddleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *opCreateEventDataStoreResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
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
			signingName := "cloudtrail"
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
				signingName = "cloudtrail"
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
				v4aScheme.SigningName = aws.String("cloudtrail")
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

func addCreateEventDataStoreResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
	return stack.Serialize.Insert(&opCreateEventDataStoreResolveEndpointMiddleware{
		EndpointResolver: options.EndpointResolverV2,
		BuiltInResolver: &builtInResolver{
			Region:       options.Region,
			UseDualStack: options.EndpointOptions.UseDualStackEndpoint,
			UseFIPS:      options.EndpointOptions.UseFIPSEndpoint,
			Endpoint:     options.BaseEndpoint,
		},
	}, "ResolveEndpoint", middleware.After)
}
