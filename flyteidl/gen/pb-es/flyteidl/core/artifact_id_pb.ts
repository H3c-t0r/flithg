// @generated by protoc-gen-es v1.7.2 with parameter "target=ts"
// @generated from file flyteidl/core/artifact_id.proto (package flyteidl.core, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3, Timestamp } from "@bufbuild/protobuf";

/**
 * @generated from message flyteidl.core.ArtifactKey
 */
export class ArtifactKey extends Message<ArtifactKey> {
  /**
   * Project and domain and suffix needs to be unique across a given artifact store.
   *
   * @generated from field: string project = 1;
   */
  project = "";

  /**
   * @generated from field: string domain = 2;
   */
  domain = "";

  /**
   * @generated from field: string name = 3;
   */
  name = "";

  /**
   * @generated from field: string org = 4;
   */
  org = "";

  constructor(data?: PartialMessage<ArtifactKey>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "flyteidl.core.ArtifactKey";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "project", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "domain", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 4, name: "org", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ArtifactKey {
    return new ArtifactKey().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ArtifactKey {
    return new ArtifactKey().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ArtifactKey {
    return new ArtifactKey().fromJsonString(jsonString, options);
  }

  static equals(a: ArtifactKey | PlainMessage<ArtifactKey> | undefined, b: ArtifactKey | PlainMessage<ArtifactKey> | undefined): boolean {
    return proto3.util.equals(ArtifactKey, a, b);
  }
}

/**
 * Only valid for triggers
 *
 * @generated from message flyteidl.core.ArtifactBindingData
 */
export class ArtifactBindingData extends Message<ArtifactBindingData> {
  /**
   * These two fields are only relevant in the partition value case
   *
   * @generated from oneof flyteidl.core.ArtifactBindingData.partition_data
   */
  partitionData: {
    /**
     * @generated from field: string partition_key = 2;
     */
    value: string;
    case: "partitionKey";
  } | {
    /**
     * @generated from field: bool bind_to_time_partition = 3;
     */
    value: boolean;
    case: "bindToTimePartition";
  } | { case: undefined; value?: undefined } = { case: undefined };

  /**
   * This is only relevant in the time partition case
   *
   * @generated from field: string transform = 4;
   */
  transform = "";

  constructor(data?: PartialMessage<ArtifactBindingData>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "flyteidl.core.ArtifactBindingData";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 2, name: "partition_key", kind: "scalar", T: 9 /* ScalarType.STRING */, oneof: "partition_data" },
    { no: 3, name: "bind_to_time_partition", kind: "scalar", T: 8 /* ScalarType.BOOL */, oneof: "partition_data" },
    { no: 4, name: "transform", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ArtifactBindingData {
    return new ArtifactBindingData().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ArtifactBindingData {
    return new ArtifactBindingData().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ArtifactBindingData {
    return new ArtifactBindingData().fromJsonString(jsonString, options);
  }

  static equals(a: ArtifactBindingData | PlainMessage<ArtifactBindingData> | undefined, b: ArtifactBindingData | PlainMessage<ArtifactBindingData> | undefined): boolean {
    return proto3.util.equals(ArtifactBindingData, a, b);
  }
}

/**
 * @generated from message flyteidl.core.InputBindingData
 */
export class InputBindingData extends Message<InputBindingData> {
  /**
   * @generated from field: string var = 1;
   */
  var = "";

  constructor(data?: PartialMessage<InputBindingData>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "flyteidl.core.InputBindingData";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "var", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): InputBindingData {
    return new InputBindingData().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): InputBindingData {
    return new InputBindingData().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): InputBindingData {
    return new InputBindingData().fromJsonString(jsonString, options);
  }

  static equals(a: InputBindingData | PlainMessage<InputBindingData> | undefined, b: InputBindingData | PlainMessage<InputBindingData> | undefined): boolean {
    return proto3.util.equals(InputBindingData, a, b);
  }
}

/**
 * @generated from message flyteidl.core.LabelValue
 */
export class LabelValue extends Message<LabelValue> {
  /**
   * @generated from oneof flyteidl.core.LabelValue.value
   */
  value: {
    /**
     * The string static value is for use in the Partitions object
     *
     * @generated from field: string static_value = 1;
     */
    value: string;
    case: "staticValue";
  } | {
    /**
     * The time value is for use in the TimePartition case
     *
     * @generated from field: google.protobuf.Timestamp time_value = 2;
     */
    value: Timestamp;
    case: "timeValue";
  } | {
    /**
     * @generated from field: flyteidl.core.ArtifactBindingData triggered_binding = 3;
     */
    value: ArtifactBindingData;
    case: "triggeredBinding";
  } | {
    /**
     * @generated from field: flyteidl.core.InputBindingData input_binding = 4;
     */
    value: InputBindingData;
    case: "inputBinding";
  } | { case: undefined; value?: undefined } = { case: undefined };

  constructor(data?: PartialMessage<LabelValue>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "flyteidl.core.LabelValue";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "static_value", kind: "scalar", T: 9 /* ScalarType.STRING */, oneof: "value" },
    { no: 2, name: "time_value", kind: "message", T: Timestamp, oneof: "value" },
    { no: 3, name: "triggered_binding", kind: "message", T: ArtifactBindingData, oneof: "value" },
    { no: 4, name: "input_binding", kind: "message", T: InputBindingData, oneof: "value" },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): LabelValue {
    return new LabelValue().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): LabelValue {
    return new LabelValue().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): LabelValue {
    return new LabelValue().fromJsonString(jsonString, options);
  }

  static equals(a: LabelValue | PlainMessage<LabelValue> | undefined, b: LabelValue | PlainMessage<LabelValue> | undefined): boolean {
    return proto3.util.equals(LabelValue, a, b);
  }
}

/**
 * @generated from message flyteidl.core.Partitions
 */
export class Partitions extends Message<Partitions> {
  /**
   * @generated from field: map<string, flyteidl.core.LabelValue> value = 1;
   */
  value: { [key: string]: LabelValue } = {};

  constructor(data?: PartialMessage<Partitions>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "flyteidl.core.Partitions";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "value", kind: "map", K: 9 /* ScalarType.STRING */, V: {kind: "message", T: LabelValue} },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Partitions {
    return new Partitions().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Partitions {
    return new Partitions().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Partitions {
    return new Partitions().fromJsonString(jsonString, options);
  }

  static equals(a: Partitions | PlainMessage<Partitions> | undefined, b: Partitions | PlainMessage<Partitions> | undefined): boolean {
    return proto3.util.equals(Partitions, a, b);
  }
}

/**
 * @generated from message flyteidl.core.TimePartition
 */
export class TimePartition extends Message<TimePartition> {
  /**
   * @generated from field: flyteidl.core.LabelValue value = 1;
   */
  value?: LabelValue;

  constructor(data?: PartialMessage<TimePartition>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "flyteidl.core.TimePartition";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "value", kind: "message", T: LabelValue },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): TimePartition {
    return new TimePartition().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): TimePartition {
    return new TimePartition().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): TimePartition {
    return new TimePartition().fromJsonString(jsonString, options);
  }

  static equals(a: TimePartition | PlainMessage<TimePartition> | undefined, b: TimePartition | PlainMessage<TimePartition> | undefined): boolean {
    return proto3.util.equals(TimePartition, a, b);
  }
}

/**
 * @generated from message flyteidl.core.ArtifactID
 */
export class ArtifactID extends Message<ArtifactID> {
  /**
   * @generated from field: flyteidl.core.ArtifactKey artifact_key = 1;
   */
  artifactKey?: ArtifactKey;

  /**
   * @generated from field: string version = 2;
   */
  version = "";

  /**
   * Think of a partition as a tag on an Artifact, except it's a key-value pair.
   * Different partitions naturally have different versions (execution ids).
   *
   * @generated from field: flyteidl.core.Partitions partitions = 3;
   */
  partitions?: Partitions;

  /**
   * There is no such thing as an empty time partition - if it's not set, then there is no time partition.
   *
   * @generated from field: flyteidl.core.TimePartition time_partition = 4;
   */
  timePartition?: TimePartition;

  constructor(data?: PartialMessage<ArtifactID>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "flyteidl.core.ArtifactID";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "artifact_key", kind: "message", T: ArtifactKey },
    { no: 2, name: "version", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "partitions", kind: "message", T: Partitions },
    { no: 4, name: "time_partition", kind: "message", T: TimePartition },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ArtifactID {
    return new ArtifactID().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ArtifactID {
    return new ArtifactID().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ArtifactID {
    return new ArtifactID().fromJsonString(jsonString, options);
  }

  static equals(a: ArtifactID | PlainMessage<ArtifactID> | undefined, b: ArtifactID | PlainMessage<ArtifactID> | undefined): boolean {
    return proto3.util.equals(ArtifactID, a, b);
  }
}

/**
 * @generated from message flyteidl.core.ArtifactTag
 */
export class ArtifactTag extends Message<ArtifactTag> {
  /**
   * @generated from field: flyteidl.core.ArtifactKey artifact_key = 1;
   */
  artifactKey?: ArtifactKey;

  /**
   * @generated from field: flyteidl.core.LabelValue value = 2;
   */
  value?: LabelValue;

  constructor(data?: PartialMessage<ArtifactTag>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "flyteidl.core.ArtifactTag";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "artifact_key", kind: "message", T: ArtifactKey },
    { no: 2, name: "value", kind: "message", T: LabelValue },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ArtifactTag {
    return new ArtifactTag().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ArtifactTag {
    return new ArtifactTag().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ArtifactTag {
    return new ArtifactTag().fromJsonString(jsonString, options);
  }

  static equals(a: ArtifactTag | PlainMessage<ArtifactTag> | undefined, b: ArtifactTag | PlainMessage<ArtifactTag> | undefined): boolean {
    return proto3.util.equals(ArtifactTag, a, b);
  }
}

/**
 * Uniqueness constraints for Artifacts
 *  - project, domain, name, version, partitions
 * Option 2 (tags are standalone, point to an individual artifact id):
 *  - project, domain, name, alias (points to one partition if partitioned)
 *  - project, domain, name, partition key, partition value
 *
 * @generated from message flyteidl.core.ArtifactQuery
 */
export class ArtifactQuery extends Message<ArtifactQuery> {
  /**
   * @generated from oneof flyteidl.core.ArtifactQuery.identifier
   */
  identifier: {
    /**
     * @generated from field: flyteidl.core.ArtifactID artifact_id = 1;
     */
    value: ArtifactID;
    case: "artifactId";
  } | {
    /**
     * @generated from field: flyteidl.core.ArtifactTag artifact_tag = 2;
     */
    value: ArtifactTag;
    case: "artifactTag";
  } | {
    /**
     * @generated from field: string uri = 3;
     */
    value: string;
    case: "uri";
  } | {
    /**
     * This is used in the trigger case, where a user specifies a value for an input that is one of the triggering
     * artifacts, or a partition value derived from a triggering artifact.
     *
     * @generated from field: flyteidl.core.ArtifactBindingData binding = 4;
     */
    value: ArtifactBindingData;
    case: "binding";
  } | { case: undefined; value?: undefined } = { case: undefined };

  constructor(data?: PartialMessage<ArtifactQuery>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "flyteidl.core.ArtifactQuery";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "artifact_id", kind: "message", T: ArtifactID, oneof: "identifier" },
    { no: 2, name: "artifact_tag", kind: "message", T: ArtifactTag, oneof: "identifier" },
    { no: 3, name: "uri", kind: "scalar", T: 9 /* ScalarType.STRING */, oneof: "identifier" },
    { no: 4, name: "binding", kind: "message", T: ArtifactBindingData, oneof: "identifier" },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ArtifactQuery {
    return new ArtifactQuery().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ArtifactQuery {
    return new ArtifactQuery().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ArtifactQuery {
    return new ArtifactQuery().fromJsonString(jsonString, options);
  }

  static equals(a: ArtifactQuery | PlainMessage<ArtifactQuery> | undefined, b: ArtifactQuery | PlainMessage<ArtifactQuery> | undefined): boolean {
    return proto3.util.equals(ArtifactQuery, a, b);
  }
}

