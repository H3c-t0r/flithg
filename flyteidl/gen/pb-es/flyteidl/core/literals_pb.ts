// @generated by protoc-gen-es v1.7.2 with parameter "target=ts"
// @generated from file flyteidl/core/literals.proto (package flyteidl.core, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Duration, Message, proto3, Struct, Timestamp } from "@bufbuild/protobuf";
import { BlobType, Error, LiteralType, OutputReference, SchemaType, StructuredDatasetType } from "./types_pb.js";

/**
 * Primitive Types
 *
 * @generated from message flyteidl.core.Primitive
 */
export class Primitive extends Message<Primitive> {
  /**
   * Defines one of simple primitive types. These types will get translated into different programming languages as
   * described in https://developers.google.com/protocol-buffers/docs/proto#scalar.
   *
   * @generated from oneof flyteidl.core.Primitive.value
   */
  value: {
    /**
     * @generated from field: int64 integer = 1;
     */
    value: bigint;
    case: "integer";
  } | {
    /**
     * @generated from field: double float_value = 2;
     */
    value: number;
    case: "floatValue";
  } | {
    /**
     * @generated from field: string string_value = 3;
     */
    value: string;
    case: "stringValue";
  } | {
    /**
     * @generated from field: bool boolean = 4;
     */
    value: boolean;
    case: "boolean";
  } | {
    /**
     * @generated from field: google.protobuf.Timestamp datetime = 5;
     */
    value: Timestamp;
    case: "datetime";
  } | {
    /**
     * @generated from field: google.protobuf.Duration duration = 6;
     */
    value: Duration;
    case: "duration";
  } | { case: undefined; value?: undefined } = { case: undefined };

  constructor(data?: PartialMessage<Primitive>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "flyteidl.core.Primitive";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "integer", kind: "scalar", T: 3 /* ScalarType.INT64 */, oneof: "value" },
    { no: 2, name: "float_value", kind: "scalar", T: 1 /* ScalarType.DOUBLE */, oneof: "value" },
    { no: 3, name: "string_value", kind: "scalar", T: 9 /* ScalarType.STRING */, oneof: "value" },
    { no: 4, name: "boolean", kind: "scalar", T: 8 /* ScalarType.BOOL */, oneof: "value" },
    { no: 5, name: "datetime", kind: "message", T: Timestamp, oneof: "value" },
    { no: 6, name: "duration", kind: "message", T: Duration, oneof: "value" },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Primitive {
    return new Primitive().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Primitive {
    return new Primitive().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Primitive {
    return new Primitive().fromJsonString(jsonString, options);
  }

  static equals(a: Primitive | PlainMessage<Primitive> | undefined, b: Primitive | PlainMessage<Primitive> | undefined): boolean {
    return proto3.util.equals(Primitive, a, b);
  }
}

/**
 * Used to denote a nil/null/None assignment to a scalar value. The underlying LiteralType for Void is intentionally
 * undefined since it can be assigned to a scalar of any LiteralType.
 *
 * @generated from message flyteidl.core.Void
 */
export class Void extends Message<Void> {
  constructor(data?: PartialMessage<Void>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "flyteidl.core.Void";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Void {
    return new Void().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Void {
    return new Void().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Void {
    return new Void().fromJsonString(jsonString, options);
  }

  static equals(a: Void | PlainMessage<Void> | undefined, b: Void | PlainMessage<Void> | undefined): boolean {
    return proto3.util.equals(Void, a, b);
  }
}

/**
 * Refers to an offloaded set of files. It encapsulates the type of the store and a unique uri for where the data is.
 * There are no restrictions on how the uri is formatted since it will depend on how to interact with the store.
 *
 * @generated from message flyteidl.core.Blob
 */
export class Blob extends Message<Blob> {
  /**
   * @generated from field: flyteidl.core.BlobMetadata metadata = 1;
   */
  metadata?: BlobMetadata;

  /**
   * @generated from field: string uri = 3;
   */
  uri = "";

  constructor(data?: PartialMessage<Blob>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "flyteidl.core.Blob";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "metadata", kind: "message", T: BlobMetadata },
    { no: 3, name: "uri", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Blob {
    return new Blob().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Blob {
    return new Blob().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Blob {
    return new Blob().fromJsonString(jsonString, options);
  }

  static equals(a: Blob | PlainMessage<Blob> | undefined, b: Blob | PlainMessage<Blob> | undefined): boolean {
    return proto3.util.equals(Blob, a, b);
  }
}

/**
 * @generated from message flyteidl.core.BlobMetadata
 */
export class BlobMetadata extends Message<BlobMetadata> {
  /**
   * @generated from field: flyteidl.core.BlobType type = 1;
   */
  type?: BlobType;

  constructor(data?: PartialMessage<BlobMetadata>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "flyteidl.core.BlobMetadata";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "type", kind: "message", T: BlobType },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): BlobMetadata {
    return new BlobMetadata().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): BlobMetadata {
    return new BlobMetadata().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): BlobMetadata {
    return new BlobMetadata().fromJsonString(jsonString, options);
  }

  static equals(a: BlobMetadata | PlainMessage<BlobMetadata> | undefined, b: BlobMetadata | PlainMessage<BlobMetadata> | undefined): boolean {
    return proto3.util.equals(BlobMetadata, a, b);
  }
}

/**
 * A simple byte array with a tag to help different parts of the system communicate about what is in the byte array.
 * It's strongly advisable that consumers of this type define a unique tag and validate the tag before parsing the data.
 *
 * @generated from message flyteidl.core.Binary
 */
export class Binary extends Message<Binary> {
  /**
   * @generated from field: bytes value = 1;
   */
  value = new Uint8Array(0);

  /**
   * @generated from field: string tag = 2;
   */
  tag = "";

  constructor(data?: PartialMessage<Binary>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "flyteidl.core.Binary";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "value", kind: "scalar", T: 12 /* ScalarType.BYTES */ },
    { no: 2, name: "tag", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Binary {
    return new Binary().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Binary {
    return new Binary().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Binary {
    return new Binary().fromJsonString(jsonString, options);
  }

  static equals(a: Binary | PlainMessage<Binary> | undefined, b: Binary | PlainMessage<Binary> | undefined): boolean {
    return proto3.util.equals(Binary, a, b);
  }
}

/**
 * A strongly typed schema that defines the interface of data retrieved from the underlying storage medium.
 *
 * @generated from message flyteidl.core.Schema
 */
export class Schema extends Message<Schema> {
  /**
   * @generated from field: string uri = 1;
   */
  uri = "";

  /**
   * @generated from field: flyteidl.core.SchemaType type = 3;
   */
  type?: SchemaType;

  constructor(data?: PartialMessage<Schema>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "flyteidl.core.Schema";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "uri", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "type", kind: "message", T: SchemaType },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Schema {
    return new Schema().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Schema {
    return new Schema().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Schema {
    return new Schema().fromJsonString(jsonString, options);
  }

  static equals(a: Schema | PlainMessage<Schema> | undefined, b: Schema | PlainMessage<Schema> | undefined): boolean {
    return proto3.util.equals(Schema, a, b);
  }
}

/**
 * The runtime representation of a tagged union value. See `UnionType` for more details.
 *
 * @generated from message flyteidl.core.Union
 */
export class Union extends Message<Union> {
  /**
   * @generated from field: flyteidl.core.Literal value = 1;
   */
  value?: Literal;

  /**
   * @generated from field: flyteidl.core.LiteralType type = 2;
   */
  type?: LiteralType;

  constructor(data?: PartialMessage<Union>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "flyteidl.core.Union";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "value", kind: "message", T: Literal },
    { no: 2, name: "type", kind: "message", T: LiteralType },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Union {
    return new Union().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Union {
    return new Union().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Union {
    return new Union().fromJsonString(jsonString, options);
  }

  static equals(a: Union | PlainMessage<Union> | undefined, b: Union | PlainMessage<Union> | undefined): boolean {
    return proto3.util.equals(Union, a, b);
  }
}

/**
 * @generated from message flyteidl.core.StructuredDatasetMetadata
 */
export class StructuredDatasetMetadata extends Message<StructuredDatasetMetadata> {
  /**
   * Bundle the type information along with the literal.
   * This is here because StructuredDatasets can often be more defined at run time than at compile time.
   * That is, at compile time you might only declare a task to return a pandas dataframe or a StructuredDataset,
   * without any column information, but at run time, you might have that column information.
   * flytekit python will copy this type information into the literal, from the type information, if not provided by
   * the various plugins (encoders).
   * Since this field is run time generated, it's not used for any type checking.
   *
   * @generated from field: flyteidl.core.StructuredDatasetType structured_dataset_type = 1;
   */
  structuredDatasetType?: StructuredDatasetType;

  constructor(data?: PartialMessage<StructuredDatasetMetadata>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "flyteidl.core.StructuredDatasetMetadata";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "structured_dataset_type", kind: "message", T: StructuredDatasetType },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): StructuredDatasetMetadata {
    return new StructuredDatasetMetadata().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): StructuredDatasetMetadata {
    return new StructuredDatasetMetadata().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): StructuredDatasetMetadata {
    return new StructuredDatasetMetadata().fromJsonString(jsonString, options);
  }

  static equals(a: StructuredDatasetMetadata | PlainMessage<StructuredDatasetMetadata> | undefined, b: StructuredDatasetMetadata | PlainMessage<StructuredDatasetMetadata> | undefined): boolean {
    return proto3.util.equals(StructuredDatasetMetadata, a, b);
  }
}

/**
 * @generated from message flyteidl.core.StructuredDataset
 */
export class StructuredDataset extends Message<StructuredDataset> {
  /**
   * String location uniquely identifying where the data is.
   * Should start with the storage location (e.g. s3://, gs://, bq://, etc.)
   *
   * @generated from field: string uri = 1;
   */
  uri = "";

  /**
   * @generated from field: flyteidl.core.StructuredDatasetMetadata metadata = 2;
   */
  metadata?: StructuredDatasetMetadata;

  constructor(data?: PartialMessage<StructuredDataset>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "flyteidl.core.StructuredDataset";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "uri", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "metadata", kind: "message", T: StructuredDatasetMetadata },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): StructuredDataset {
    return new StructuredDataset().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): StructuredDataset {
    return new StructuredDataset().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): StructuredDataset {
    return new StructuredDataset().fromJsonString(jsonString, options);
  }

  static equals(a: StructuredDataset | PlainMessage<StructuredDataset> | undefined, b: StructuredDataset | PlainMessage<StructuredDataset> | undefined): boolean {
    return proto3.util.equals(StructuredDataset, a, b);
  }
}

/**
 * @generated from message flyteidl.core.Scalar
 */
export class Scalar extends Message<Scalar> {
  /**
   * @generated from oneof flyteidl.core.Scalar.value
   */
  value: {
    /**
     * @generated from field: flyteidl.core.Primitive primitive = 1;
     */
    value: Primitive;
    case: "primitive";
  } | {
    /**
     * @generated from field: flyteidl.core.Blob blob = 2;
     */
    value: Blob;
    case: "blob";
  } | {
    /**
     * @generated from field: flyteidl.core.Binary binary = 3;
     */
    value: Binary;
    case: "binary";
  } | {
    /**
     * @generated from field: flyteidl.core.Schema schema = 4;
     */
    value: Schema;
    case: "schema";
  } | {
    /**
     * @generated from field: flyteidl.core.Void none_type = 5;
     */
    value: Void;
    case: "noneType";
  } | {
    /**
     * @generated from field: flyteidl.core.Error error = 6;
     */
    value: Error;
    case: "error";
  } | {
    /**
     * @generated from field: google.protobuf.Struct generic = 7;
     */
    value: Struct;
    case: "generic";
  } | {
    /**
     * @generated from field: flyteidl.core.StructuredDataset structured_dataset = 8;
     */
    value: StructuredDataset;
    case: "structuredDataset";
  } | {
    /**
     * @generated from field: flyteidl.core.Union union = 9;
     */
    value: Union;
    case: "union";
  } | { case: undefined; value?: undefined } = { case: undefined };

  constructor(data?: PartialMessage<Scalar>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "flyteidl.core.Scalar";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "primitive", kind: "message", T: Primitive, oneof: "value" },
    { no: 2, name: "blob", kind: "message", T: Blob, oneof: "value" },
    { no: 3, name: "binary", kind: "message", T: Binary, oneof: "value" },
    { no: 4, name: "schema", kind: "message", T: Schema, oneof: "value" },
    { no: 5, name: "none_type", kind: "message", T: Void, oneof: "value" },
    { no: 6, name: "error", kind: "message", T: Error, oneof: "value" },
    { no: 7, name: "generic", kind: "message", T: Struct, oneof: "value" },
    { no: 8, name: "structured_dataset", kind: "message", T: StructuredDataset, oneof: "value" },
    { no: 9, name: "union", kind: "message", T: Union, oneof: "value" },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Scalar {
    return new Scalar().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Scalar {
    return new Scalar().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Scalar {
    return new Scalar().fromJsonString(jsonString, options);
  }

  static equals(a: Scalar | PlainMessage<Scalar> | undefined, b: Scalar | PlainMessage<Scalar> | undefined): boolean {
    return proto3.util.equals(Scalar, a, b);
  }
}

/**
 * A simple value. This supports any level of nesting (e.g. array of array of array of Blobs) as well as simple primitives.
 *
 * @generated from message flyteidl.core.Literal
 */
export class Literal extends Message<Literal> {
  /**
   * @generated from oneof flyteidl.core.Literal.value
   */
  value: {
    /**
     * A simple value.
     *
     * @generated from field: flyteidl.core.Scalar scalar = 1;
     */
    value: Scalar;
    case: "scalar";
  } | {
    /**
     * A collection of literals to allow nesting.
     *
     * @generated from field: flyteidl.core.LiteralCollection collection = 2;
     */
    value: LiteralCollection;
    case: "collection";
  } | {
    /**
     * A map of strings to literals.
     *
     * @generated from field: flyteidl.core.LiteralMap map = 3;
     */
    value: LiteralMap;
    case: "map";
  } | { case: undefined; value?: undefined } = { case: undefined };

  /**
   * A hash representing this literal.
   * This is used for caching purposes. For more details refer to RFC 1893
   * (https://github.com/flyteorg/flyte/blob/master/rfc/system/1893-caching-of-offloaded-objects.md)
   *
   * @generated from field: string hash = 4;
   */
  hash = "";

  /**
   * Additional metadata for literals.
   *
   * @generated from field: map<string, string> metadata = 5;
   */
  metadata: { [key: string]: string } = {};

  /**
   * Testing out gitattributes
   *
   * @generated from field: string test = 6;
   */
  test = "";

  constructor(data?: PartialMessage<Literal>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "flyteidl.core.Literal";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "scalar", kind: "message", T: Scalar, oneof: "value" },
    { no: 2, name: "collection", kind: "message", T: LiteralCollection, oneof: "value" },
    { no: 3, name: "map", kind: "message", T: LiteralMap, oneof: "value" },
    { no: 4, name: "hash", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 5, name: "metadata", kind: "map", K: 9 /* ScalarType.STRING */, V: {kind: "scalar", T: 9 /* ScalarType.STRING */} },
    { no: 6, name: "test", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Literal {
    return new Literal().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Literal {
    return new Literal().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Literal {
    return new Literal().fromJsonString(jsonString, options);
  }

  static equals(a: Literal | PlainMessage<Literal> | undefined, b: Literal | PlainMessage<Literal> | undefined): boolean {
    return proto3.util.equals(Literal, a, b);
  }
}

/**
 * A collection of literals. This is a workaround since oneofs in proto messages cannot contain a repeated field.
 *
 * @generated from message flyteidl.core.LiteralCollection
 */
export class LiteralCollection extends Message<LiteralCollection> {
  /**
   * @generated from field: repeated flyteidl.core.Literal literals = 1;
   */
  literals: Literal[] = [];

  constructor(data?: PartialMessage<LiteralCollection>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "flyteidl.core.LiteralCollection";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "literals", kind: "message", T: Literal, repeated: true },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): LiteralCollection {
    return new LiteralCollection().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): LiteralCollection {
    return new LiteralCollection().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): LiteralCollection {
    return new LiteralCollection().fromJsonString(jsonString, options);
  }

  static equals(a: LiteralCollection | PlainMessage<LiteralCollection> | undefined, b: LiteralCollection | PlainMessage<LiteralCollection> | undefined): boolean {
    return proto3.util.equals(LiteralCollection, a, b);
  }
}

/**
 * A map of literals. This is a workaround since oneofs in proto messages cannot contain a repeated field.
 *
 * @generated from message flyteidl.core.LiteralMap
 */
export class LiteralMap extends Message<LiteralMap> {
  /**
   * @generated from field: map<string, flyteidl.core.Literal> literals = 1;
   */
  literals: { [key: string]: Literal } = {};

  constructor(data?: PartialMessage<LiteralMap>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "flyteidl.core.LiteralMap";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "literals", kind: "map", K: 9 /* ScalarType.STRING */, V: {kind: "message", T: Literal} },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): LiteralMap {
    return new LiteralMap().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): LiteralMap {
    return new LiteralMap().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): LiteralMap {
    return new LiteralMap().fromJsonString(jsonString, options);
  }

  static equals(a: LiteralMap | PlainMessage<LiteralMap> | undefined, b: LiteralMap | PlainMessage<LiteralMap> | undefined): boolean {
    return proto3.util.equals(LiteralMap, a, b);
  }
}

/**
 * A collection of BindingData items.
 *
 * @generated from message flyteidl.core.BindingDataCollection
 */
export class BindingDataCollection extends Message<BindingDataCollection> {
  /**
   * @generated from field: repeated flyteidl.core.BindingData bindings = 1;
   */
  bindings: BindingData[] = [];

  constructor(data?: PartialMessage<BindingDataCollection>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "flyteidl.core.BindingDataCollection";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "bindings", kind: "message", T: BindingData, repeated: true },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): BindingDataCollection {
    return new BindingDataCollection().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): BindingDataCollection {
    return new BindingDataCollection().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): BindingDataCollection {
    return new BindingDataCollection().fromJsonString(jsonString, options);
  }

  static equals(a: BindingDataCollection | PlainMessage<BindingDataCollection> | undefined, b: BindingDataCollection | PlainMessage<BindingDataCollection> | undefined): boolean {
    return proto3.util.equals(BindingDataCollection, a, b);
  }
}

/**
 * A map of BindingData items.
 *
 * @generated from message flyteidl.core.BindingDataMap
 */
export class BindingDataMap extends Message<BindingDataMap> {
  /**
   * @generated from field: map<string, flyteidl.core.BindingData> bindings = 1;
   */
  bindings: { [key: string]: BindingData } = {};

  constructor(data?: PartialMessage<BindingDataMap>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "flyteidl.core.BindingDataMap";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "bindings", kind: "map", K: 9 /* ScalarType.STRING */, V: {kind: "message", T: BindingData} },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): BindingDataMap {
    return new BindingDataMap().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): BindingDataMap {
    return new BindingDataMap().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): BindingDataMap {
    return new BindingDataMap().fromJsonString(jsonString, options);
  }

  static equals(a: BindingDataMap | PlainMessage<BindingDataMap> | undefined, b: BindingDataMap | PlainMessage<BindingDataMap> | undefined): boolean {
    return proto3.util.equals(BindingDataMap, a, b);
  }
}

/**
 * @generated from message flyteidl.core.UnionInfo
 */
export class UnionInfo extends Message<UnionInfo> {
  /**
   * @generated from field: flyteidl.core.LiteralType targetType = 1;
   */
  targetType?: LiteralType;

  constructor(data?: PartialMessage<UnionInfo>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "flyteidl.core.UnionInfo";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "targetType", kind: "message", T: LiteralType },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): UnionInfo {
    return new UnionInfo().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): UnionInfo {
    return new UnionInfo().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): UnionInfo {
    return new UnionInfo().fromJsonString(jsonString, options);
  }

  static equals(a: UnionInfo | PlainMessage<UnionInfo> | undefined, b: UnionInfo | PlainMessage<UnionInfo> | undefined): boolean {
    return proto3.util.equals(UnionInfo, a, b);
  }
}

/**
 * Specifies either a simple value or a reference to another output.
 *
 * @generated from message flyteidl.core.BindingData
 */
export class BindingData extends Message<BindingData> {
  /**
   * @generated from oneof flyteidl.core.BindingData.value
   */
  value: {
    /**
     * A simple scalar value.
     *
     * @generated from field: flyteidl.core.Scalar scalar = 1;
     */
    value: Scalar;
    case: "scalar";
  } | {
    /**
     * A collection of binding data. This allows nesting of binding data to any number
     * of levels.
     *
     * @generated from field: flyteidl.core.BindingDataCollection collection = 2;
     */
    value: BindingDataCollection;
    case: "collection";
  } | {
    /**
     * References an output promised by another node.
     *
     * @generated from field: flyteidl.core.OutputReference promise = 3;
     */
    value: OutputReference;
    case: "promise";
  } | {
    /**
     * A map of bindings. The key is always a string.
     *
     * @generated from field: flyteidl.core.BindingDataMap map = 4;
     */
    value: BindingDataMap;
    case: "map";
  } | { case: undefined; value?: undefined } = { case: undefined };

  /**
   * @generated from field: flyteidl.core.UnionInfo union = 5;
   */
  union?: UnionInfo;

  constructor(data?: PartialMessage<BindingData>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "flyteidl.core.BindingData";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "scalar", kind: "message", T: Scalar, oneof: "value" },
    { no: 2, name: "collection", kind: "message", T: BindingDataCollection, oneof: "value" },
    { no: 3, name: "promise", kind: "message", T: OutputReference, oneof: "value" },
    { no: 4, name: "map", kind: "message", T: BindingDataMap, oneof: "value" },
    { no: 5, name: "union", kind: "message", T: UnionInfo },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): BindingData {
    return new BindingData().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): BindingData {
    return new BindingData().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): BindingData {
    return new BindingData().fromJsonString(jsonString, options);
  }

  static equals(a: BindingData | PlainMessage<BindingData> | undefined, b: BindingData | PlainMessage<BindingData> | undefined): boolean {
    return proto3.util.equals(BindingData, a, b);
  }
}

/**
 * An input/output binding of a variable to either static value or a node output.
 *
 * @generated from message flyteidl.core.Binding
 */
export class Binding extends Message<Binding> {
  /**
   * Variable name must match an input/output variable of the node.
   *
   * @generated from field: string var = 1;
   */
  var = "";

  /**
   * Data to use to bind this variable.
   *
   * @generated from field: flyteidl.core.BindingData binding = 2;
   */
  binding?: BindingData;

  constructor(data?: PartialMessage<Binding>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "flyteidl.core.Binding";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "var", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "binding", kind: "message", T: BindingData },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Binding {
    return new Binding().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Binding {
    return new Binding().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Binding {
    return new Binding().fromJsonString(jsonString, options);
  }

  static equals(a: Binding | PlainMessage<Binding> | undefined, b: Binding | PlainMessage<Binding> | undefined): boolean {
    return proto3.util.equals(Binding, a, b);
  }
}

/**
 * A generic key value pair.
 *
 * @generated from message flyteidl.core.KeyValuePair
 */
export class KeyValuePair extends Message<KeyValuePair> {
  /**
   * required.
   *
   * @generated from field: string key = 1;
   */
  key = "";

  /**
   * +optional.
   *
   * @generated from field: string value = 2;
   */
  value = "";

  constructor(data?: PartialMessage<KeyValuePair>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "flyteidl.core.KeyValuePair";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "key", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "value", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): KeyValuePair {
    return new KeyValuePair().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): KeyValuePair {
    return new KeyValuePair().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): KeyValuePair {
    return new KeyValuePair().fromJsonString(jsonString, options);
  }

  static equals(a: KeyValuePair | PlainMessage<KeyValuePair> | undefined, b: KeyValuePair | PlainMessage<KeyValuePair> | undefined): boolean {
    return proto3.util.equals(KeyValuePair, a, b);
  }
}

/**
 * Retry strategy associated with an executable unit.
 *
 * @generated from message flyteidl.core.RetryStrategy
 */
export class RetryStrategy extends Message<RetryStrategy> {
  /**
   * Number of retries. Retries will be consumed when the job fails with a recoverable error.
   * The number of retries must be less than or equals to 10.
   *
   * @generated from field: uint32 retries = 5;
   */
  retries = 0;

  constructor(data?: PartialMessage<RetryStrategy>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "flyteidl.core.RetryStrategy";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 5, name: "retries", kind: "scalar", T: 13 /* ScalarType.UINT32 */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): RetryStrategy {
    return new RetryStrategy().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): RetryStrategy {
    return new RetryStrategy().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): RetryStrategy {
    return new RetryStrategy().fromJsonString(jsonString, options);
  }

  static equals(a: RetryStrategy | PlainMessage<RetryStrategy> | undefined, b: RetryStrategy | PlainMessage<RetryStrategy> | undefined): boolean {
    return proto3.util.equals(RetryStrategy, a, b);
  }
}

