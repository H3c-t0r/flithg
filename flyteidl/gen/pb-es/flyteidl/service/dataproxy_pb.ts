// @generated by protoc-gen-es v1.7.2 with parameter "target=ts"
// @generated from file flyteidl/service/dataproxy.proto (package flyteidl.service, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Duration, Message, proto3, Timestamp } from "@bufbuild/protobuf";
import { NodeExecutionIdentifier } from "../core/identifier_pb.js";
import { Literal, LiteralMap } from "../core/literals_pb.js";

/**
 * ArtifactType
 *
 * @generated from enum flyteidl.service.ArtifactType
 */
export enum ArtifactType {
  /**
   * ARTIFACT_TYPE_UNDEFINED is the default, often invalid, value for the enum.
   *
   * @generated from enum value: ARTIFACT_TYPE_UNDEFINED = 0;
   */
  UNDEFINED = 0,

  /**
   * ARTIFACT_TYPE_DECK refers to the deck html file optionally generated after a task, a workflow or a launch plan
   * finishes executing.
   *
   * @generated from enum value: ARTIFACT_TYPE_DECK = 1;
   */
  DECK = 1,
}
// Retrieve enum metadata with: proto3.getEnumType(ArtifactType)
proto3.util.setEnumType(ArtifactType, "flyteidl.service.ArtifactType", [
  { no: 0, name: "ARTIFACT_TYPE_UNDEFINED" },
  { no: 1, name: "ARTIFACT_TYPE_DECK" },
]);

/**
 * @generated from message flyteidl.service.CreateUploadLocationResponse
 */
export class CreateUploadLocationResponse extends Message<CreateUploadLocationResponse> {
  /**
   * SignedUrl specifies the url to use to upload content to (e.g. https://my-bucket.s3.amazonaws.com/randomstring/suffix.tar?X-...)
   *
   * @generated from field: string signed_url = 1;
   */
  signedUrl = "";

  /**
   * NativeUrl specifies the url in the format of the configured storage provider (e.g. s3://my-bucket/randomstring/suffix.tar)
   *
   * @generated from field: string native_url = 2;
   */
  nativeUrl = "";

  /**
   * ExpiresAt defines when will the signed URL expires.
   *
   * @generated from field: google.protobuf.Timestamp expires_at = 3;
   */
  expiresAt?: Timestamp;

  /**
   * Data proxy generates these headers for client, and they have to add these headers to the request when uploading the file.
   *
   * @generated from field: map<string, string> headers = 4;
   */
  headers: { [key: string]: string } = {};

  constructor(data?: PartialMessage<CreateUploadLocationResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "flyteidl.service.CreateUploadLocationResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "signed_url", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "native_url", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "expires_at", kind: "message", T: Timestamp },
    { no: 4, name: "headers", kind: "map", K: 9 /* ScalarType.STRING */, V: {kind: "scalar", T: 9 /* ScalarType.STRING */} },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): CreateUploadLocationResponse {
    return new CreateUploadLocationResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): CreateUploadLocationResponse {
    return new CreateUploadLocationResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): CreateUploadLocationResponse {
    return new CreateUploadLocationResponse().fromJsonString(jsonString, options);
  }

  static equals(a: CreateUploadLocationResponse | PlainMessage<CreateUploadLocationResponse> | undefined, b: CreateUploadLocationResponse | PlainMessage<CreateUploadLocationResponse> | undefined): boolean {
    return proto3.util.equals(CreateUploadLocationResponse, a, b);
  }
}

/**
 * CreateUploadLocationRequest specified request for the CreateUploadLocation API.
 * The implementation in data proxy service will create the s3 location with some server side configured prefixes,
 * and then:
 *   - project/domain/(a deterministic str representation of the content_md5)/filename (if present); OR
 *   - project/domain/filename_root (if present)/filename (if present).
 *
 * @generated from message flyteidl.service.CreateUploadLocationRequest
 */
export class CreateUploadLocationRequest extends Message<CreateUploadLocationRequest> {
  /**
   * Project to create the upload location for
   * +required
   *
   * @generated from field: string project = 1;
   */
  project = "";

  /**
   * Domain to create the upload location for.
   * +required
   *
   * @generated from field: string domain = 2;
   */
  domain = "";

  /**
   * Filename specifies a desired suffix for the generated location. E.g. `file.py` or `pre/fix/file.zip`.
   * +optional. By default, the service will generate a consistent name based on the provided parameters.
   *
   * @generated from field: string filename = 3;
   */
  filename = "";

  /**
   * ExpiresIn defines a requested expiration duration for the generated url. The request will be rejected if this
   * exceeds the platform allowed max.
   * +optional. The default value comes from a global config.
   *
   * @generated from field: google.protobuf.Duration expires_in = 4;
   */
  expiresIn?: Duration;

  /**
   * ContentMD5 restricts the upload location to the specific MD5 provided. The ContentMD5 will also appear in the
   * generated path.
   * +required
   *
   * @generated from field: bytes content_md5 = 5;
   */
  contentMd5 = new Uint8Array(0);

  /**
   * If present, data proxy will use this string in lieu of the md5 hash in the path. When the filename is also included
   * this makes the upload location deterministic. The native url will still be prefixed by the upload location prefix
   * in data proxy config. This option is useful when uploading multiple files.
   * +optional
   *
   * @generated from field: string filename_root = 6;
   */
  filenameRoot = "";

  /**
   * If true, the data proxy will add extra header (x-ams-meta-, x-goog-meta, etc) to the signed URL and
   * it will force clients to add metadata to the object.
   *
   * @generated from field: bool add_metadata = 7;
   */
  addMetadata = false;

  constructor(data?: PartialMessage<CreateUploadLocationRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "flyteidl.service.CreateUploadLocationRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "project", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "domain", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "filename", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 4, name: "expires_in", kind: "message", T: Duration },
    { no: 5, name: "content_md5", kind: "scalar", T: 12 /* ScalarType.BYTES */ },
    { no: 6, name: "filename_root", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 7, name: "add_metadata", kind: "scalar", T: 8 /* ScalarType.BOOL */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): CreateUploadLocationRequest {
    return new CreateUploadLocationRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): CreateUploadLocationRequest {
    return new CreateUploadLocationRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): CreateUploadLocationRequest {
    return new CreateUploadLocationRequest().fromJsonString(jsonString, options);
  }

  static equals(a: CreateUploadLocationRequest | PlainMessage<CreateUploadLocationRequest> | undefined, b: CreateUploadLocationRequest | PlainMessage<CreateUploadLocationRequest> | undefined): boolean {
    return proto3.util.equals(CreateUploadLocationRequest, a, b);
  }
}

/**
 * CreateDownloadLocationRequest specified request for the CreateDownloadLocation API.
 *
 * @generated from message flyteidl.service.CreateDownloadLocationRequest
 * @deprecated
 */
export class CreateDownloadLocationRequest extends Message<CreateDownloadLocationRequest> {
  /**
   * NativeUrl specifies the url in the format of the configured storage provider (e.g. s3://my-bucket/randomstring/suffix.tar)
   *
   * @generated from field: string native_url = 1;
   */
  nativeUrl = "";

  /**
   * ExpiresIn defines a requested expiration duration for the generated url. The request will be rejected if this
   * exceeds the platform allowed max.
   * +optional. The default value comes from a global config.
   *
   * @generated from field: google.protobuf.Duration expires_in = 2;
   */
  expiresIn?: Duration;

  constructor(data?: PartialMessage<CreateDownloadLocationRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "flyteidl.service.CreateDownloadLocationRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "native_url", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "expires_in", kind: "message", T: Duration },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): CreateDownloadLocationRequest {
    return new CreateDownloadLocationRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): CreateDownloadLocationRequest {
    return new CreateDownloadLocationRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): CreateDownloadLocationRequest {
    return new CreateDownloadLocationRequest().fromJsonString(jsonString, options);
  }

  static equals(a: CreateDownloadLocationRequest | PlainMessage<CreateDownloadLocationRequest> | undefined, b: CreateDownloadLocationRequest | PlainMessage<CreateDownloadLocationRequest> | undefined): boolean {
    return proto3.util.equals(CreateDownloadLocationRequest, a, b);
  }
}

/**
 * @generated from message flyteidl.service.CreateDownloadLocationResponse
 * @deprecated
 */
export class CreateDownloadLocationResponse extends Message<CreateDownloadLocationResponse> {
  /**
   * SignedUrl specifies the url to use to download content from (e.g. https://my-bucket.s3.amazonaws.com/randomstring/suffix.tar?X-...)
   *
   * @generated from field: string signed_url = 1;
   */
  signedUrl = "";

  /**
   * ExpiresAt defines when will the signed URL expires.
   *
   * @generated from field: google.protobuf.Timestamp expires_at = 2;
   */
  expiresAt?: Timestamp;

  constructor(data?: PartialMessage<CreateDownloadLocationResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "flyteidl.service.CreateDownloadLocationResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "signed_url", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "expires_at", kind: "message", T: Timestamp },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): CreateDownloadLocationResponse {
    return new CreateDownloadLocationResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): CreateDownloadLocationResponse {
    return new CreateDownloadLocationResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): CreateDownloadLocationResponse {
    return new CreateDownloadLocationResponse().fromJsonString(jsonString, options);
  }

  static equals(a: CreateDownloadLocationResponse | PlainMessage<CreateDownloadLocationResponse> | undefined, b: CreateDownloadLocationResponse | PlainMessage<CreateDownloadLocationResponse> | undefined): boolean {
    return proto3.util.equals(CreateDownloadLocationResponse, a, b);
  }
}

/**
 * CreateDownloadLinkRequest defines the request parameters to create a download link (signed url)
 *
 * @generated from message flyteidl.service.CreateDownloadLinkRequest
 */
export class CreateDownloadLinkRequest extends Message<CreateDownloadLinkRequest> {
  /**
   * ArtifactType of the artifact requested.
   *
   * @generated from field: flyteidl.service.ArtifactType artifact_type = 1;
   */
  artifactType = ArtifactType.UNDEFINED;

  /**
   * ExpiresIn defines a requested expiration duration for the generated url. The request will be rejected if this
   * exceeds the platform allowed max.
   * +optional. The default value comes from a global config.
   *
   * @generated from field: google.protobuf.Duration expires_in = 2;
   */
  expiresIn?: Duration;

  /**
   * @generated from oneof flyteidl.service.CreateDownloadLinkRequest.source
   */
  source: {
    /**
     * NodeId is the unique identifier for the node execution. For a task node, this will retrieve the output of the
     * most recent attempt of the task.
     *
     * @generated from field: flyteidl.core.NodeExecutionIdentifier node_execution_id = 3;
     */
    value: NodeExecutionIdentifier;
    case: "nodeExecutionId";
  } | { case: undefined; value?: undefined } = { case: undefined };

  constructor(data?: PartialMessage<CreateDownloadLinkRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "flyteidl.service.CreateDownloadLinkRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "artifact_type", kind: "enum", T: proto3.getEnumType(ArtifactType) },
    { no: 2, name: "expires_in", kind: "message", T: Duration },
    { no: 3, name: "node_execution_id", kind: "message", T: NodeExecutionIdentifier, oneof: "source" },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): CreateDownloadLinkRequest {
    return new CreateDownloadLinkRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): CreateDownloadLinkRequest {
    return new CreateDownloadLinkRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): CreateDownloadLinkRequest {
    return new CreateDownloadLinkRequest().fromJsonString(jsonString, options);
  }

  static equals(a: CreateDownloadLinkRequest | PlainMessage<CreateDownloadLinkRequest> | undefined, b: CreateDownloadLinkRequest | PlainMessage<CreateDownloadLinkRequest> | undefined): boolean {
    return proto3.util.equals(CreateDownloadLinkRequest, a, b);
  }
}

/**
 * CreateDownloadLinkResponse defines the response for the generated links
 *
 * @generated from message flyteidl.service.CreateDownloadLinkResponse
 */
export class CreateDownloadLinkResponse extends Message<CreateDownloadLinkResponse> {
  /**
   * SignedUrl specifies the url to use to download content from (e.g. https://my-bucket.s3.amazonaws.com/randomstring/suffix.tar?X-...)
   *
   * @generated from field: repeated string signed_url = 1 [deprecated = true];
   * @deprecated
   */
  signedUrl: string[] = [];

  /**
   * ExpiresAt defines when will the signed URL expire.
   *
   * @generated from field: google.protobuf.Timestamp expires_at = 2 [deprecated = true];
   * @deprecated
   */
  expiresAt?: Timestamp;

  /**
   * New wrapper object containing the signed urls and expiration time
   *
   * @generated from field: flyteidl.service.PreSignedURLs pre_signed_urls = 3;
   */
  preSignedUrls?: PreSignedURLs;

  constructor(data?: PartialMessage<CreateDownloadLinkResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "flyteidl.service.CreateDownloadLinkResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "signed_url", kind: "scalar", T: 9 /* ScalarType.STRING */, repeated: true },
    { no: 2, name: "expires_at", kind: "message", T: Timestamp },
    { no: 3, name: "pre_signed_urls", kind: "message", T: PreSignedURLs },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): CreateDownloadLinkResponse {
    return new CreateDownloadLinkResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): CreateDownloadLinkResponse {
    return new CreateDownloadLinkResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): CreateDownloadLinkResponse {
    return new CreateDownloadLinkResponse().fromJsonString(jsonString, options);
  }

  static equals(a: CreateDownloadLinkResponse | PlainMessage<CreateDownloadLinkResponse> | undefined, b: CreateDownloadLinkResponse | PlainMessage<CreateDownloadLinkResponse> | undefined): boolean {
    return proto3.util.equals(CreateDownloadLinkResponse, a, b);
  }
}

/**
 * Wrapper object since the message is shared across this and the GetDataResponse
 *
 * @generated from message flyteidl.service.PreSignedURLs
 */
export class PreSignedURLs extends Message<PreSignedURLs> {
  /**
   * SignedUrl specifies the url to use to download content from (e.g. https://my-bucket.s3.amazonaws.com/randomstring/suffix.tar?X-...)
   *
   * @generated from field: repeated string signed_url = 1;
   */
  signedUrl: string[] = [];

  /**
   * ExpiresAt defines when will the signed URL expire.
   *
   * @generated from field: google.protobuf.Timestamp expires_at = 2;
   */
  expiresAt?: Timestamp;

  constructor(data?: PartialMessage<PreSignedURLs>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "flyteidl.service.PreSignedURLs";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "signed_url", kind: "scalar", T: 9 /* ScalarType.STRING */, repeated: true },
    { no: 2, name: "expires_at", kind: "message", T: Timestamp },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): PreSignedURLs {
    return new PreSignedURLs().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): PreSignedURLs {
    return new PreSignedURLs().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): PreSignedURLs {
    return new PreSignedURLs().fromJsonString(jsonString, options);
  }

  static equals(a: PreSignedURLs | PlainMessage<PreSignedURLs> | undefined, b: PreSignedURLs | PlainMessage<PreSignedURLs> | undefined): boolean {
    return proto3.util.equals(PreSignedURLs, a, b);
  }
}

/**
 * General request artifact to retrieve data from a Flyte artifact url.
 *
 * @generated from message flyteidl.service.GetDataRequest
 */
export class GetDataRequest extends Message<GetDataRequest> {
  /**
   * A unique identifier in the form of flyte://<something> that uniquely, for a given Flyte
   * backend, identifies a Flyte artifact ([i]nput, [o]output, flyte [d]eck, etc.).
   * e.g. flyte://v1/proj/development/execid/n2/0/i (for 0th task execution attempt input)
   *      flyte://v1/proj/development/execid/n2/i (for node execution input)
   *      flyte://v1/proj/development/execid/n2/o/o3 (the o3 output of the second node)
   *
   * @generated from field: string flyte_url = 1;
   */
  flyteUrl = "";

  constructor(data?: PartialMessage<GetDataRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "flyteidl.service.GetDataRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "flyte_url", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetDataRequest {
    return new GetDataRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetDataRequest {
    return new GetDataRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetDataRequest {
    return new GetDataRequest().fromJsonString(jsonString, options);
  }

  static equals(a: GetDataRequest | PlainMessage<GetDataRequest> | undefined, b: GetDataRequest | PlainMessage<GetDataRequest> | undefined): boolean {
    return proto3.util.equals(GetDataRequest, a, b);
  }
}

/**
 * @generated from message flyteidl.service.GetDataResponse
 */
export class GetDataResponse extends Message<GetDataResponse> {
  /**
   * @generated from oneof flyteidl.service.GetDataResponse.data
   */
  data: {
    /**
     * literal map data will be returned
     *
     * @generated from field: flyteidl.core.LiteralMap literal_map = 1;
     */
    value: LiteralMap;
    case: "literalMap";
  } | {
    /**
     * Flyte deck html will be returned as a signed url users can download
     *
     * @generated from field: flyteidl.service.PreSignedURLs pre_signed_urls = 2;
     */
    value: PreSignedURLs;
    case: "preSignedUrls";
  } | {
    /**
     * Single literal will be returned. This is returned when the user/url requests a specific output or input
     * by name. See the o3 example above.
     *
     * @generated from field: flyteidl.core.Literal literal = 3;
     */
    value: Literal;
    case: "literal";
  } | { case: undefined; value?: undefined } = { case: undefined };

  constructor(data?: PartialMessage<GetDataResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "flyteidl.service.GetDataResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "literal_map", kind: "message", T: LiteralMap, oneof: "data" },
    { no: 2, name: "pre_signed_urls", kind: "message", T: PreSignedURLs, oneof: "data" },
    { no: 3, name: "literal", kind: "message", T: Literal, oneof: "data" },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetDataResponse {
    return new GetDataResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetDataResponse {
    return new GetDataResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetDataResponse {
    return new GetDataResponse().fromJsonString(jsonString, options);
  }

  static equals(a: GetDataResponse | PlainMessage<GetDataResponse> | undefined, b: GetDataResponse | PlainMessage<GetDataResponse> | undefined): boolean {
    return proto3.util.equals(GetDataResponse, a, b);
  }
}

