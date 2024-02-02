// @generated by protoc-gen-es v1.7.2 with parameter "target=ts"
// @generated from file flyteidl/admin/agent.proto (package flyteidl.admin, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Duration, Message, proto3, protoInt64, Timestamp } from "@bufbuild/protobuf";
import { TaskExecutionIdentifier } from "../core/identifier_pb.js";
import { LiteralMap } from "../core/literals_pb.js";
import { TaskTemplate } from "../core/tasks_pb.js";
import { TaskExecution_Phase, TaskLog } from "../core/execution_pb.js";
import { ExecutionMetricResult } from "../core/metrics_pb.js";

/**
 * The state of the execution is used to control its visibility in the UI/CLI.
 *
 * @generated from enum flyteidl.admin.State
 */
export enum State {
  /**
   * @generated from enum value: RETRYABLE_FAILURE = 0;
   */
  RETRYABLE_FAILURE = 0,

  /**
   * @generated from enum value: PERMANENT_FAILURE = 1;
   */
  PERMANENT_FAILURE = 1,

  /**
   * @generated from enum value: PENDING = 2;
   */
  PENDING = 2,

  /**
   * @generated from enum value: RUNNING = 3;
   */
  RUNNING = 3,

  /**
   * @generated from enum value: SUCCEEDED = 4;
   */
  SUCCEEDED = 4,
}
// Retrieve enum metadata with: proto3.getEnumType(State)
proto3.util.setEnumType(State, "flyteidl.admin.State", [
  { no: 0, name: "RETRYABLE_FAILURE" },
  { no: 1, name: "PERMANENT_FAILURE" },
  { no: 2, name: "PENDING" },
  { no: 3, name: "RUNNING" },
  { no: 4, name: "SUCCEEDED" },
]);

/**
 * Represents a subset of runtime task execution metadata that are relevant to external plugins.
 *
 * @generated from message flyteidl.admin.TaskExecutionMetadata
 */
export class TaskExecutionMetadata extends Message<TaskExecutionMetadata> {
  /**
   * ID of the task execution
   *
   * @generated from field: flyteidl.core.TaskExecutionIdentifier task_execution_id = 1;
   */
  taskExecutionId?: TaskExecutionIdentifier;

  /**
   * k8s namespace where the task is executed in
   *
   * @generated from field: string namespace = 2;
   */
  namespace = "";

  /**
   * Labels attached to the task execution
   *
   * @generated from field: map<string, string> labels = 3;
   */
  labels: { [key: string]: string } = {};

  /**
   * Annotations attached to the task execution
   *
   * @generated from field: map<string, string> annotations = 4;
   */
  annotations: { [key: string]: string } = {};

  /**
   * k8s service account associated with the task execution
   *
   * @generated from field: string k8s_service_account = 5;
   */
  k8sServiceAccount = "";

  /**
   * Environment variables attached to the task execution
   *
   * @generated from field: map<string, string> environment_variables = 6;
   */
  environmentVariables: { [key: string]: string } = {};

  constructor(data?: PartialMessage<TaskExecutionMetadata>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "flyteidl.admin.TaskExecutionMetadata";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "task_execution_id", kind: "message", T: TaskExecutionIdentifier },
    { no: 2, name: "namespace", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "labels", kind: "map", K: 9 /* ScalarType.STRING */, V: {kind: "scalar", T: 9 /* ScalarType.STRING */} },
    { no: 4, name: "annotations", kind: "map", K: 9 /* ScalarType.STRING */, V: {kind: "scalar", T: 9 /* ScalarType.STRING */} },
    { no: 5, name: "k8s_service_account", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 6, name: "environment_variables", kind: "map", K: 9 /* ScalarType.STRING */, V: {kind: "scalar", T: 9 /* ScalarType.STRING */} },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): TaskExecutionMetadata {
    return new TaskExecutionMetadata().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): TaskExecutionMetadata {
    return new TaskExecutionMetadata().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): TaskExecutionMetadata {
    return new TaskExecutionMetadata().fromJsonString(jsonString, options);
  }

  static equals(a: TaskExecutionMetadata | PlainMessage<TaskExecutionMetadata> | undefined, b: TaskExecutionMetadata | PlainMessage<TaskExecutionMetadata> | undefined): boolean {
    return proto3.util.equals(TaskExecutionMetadata, a, b);
  }
}

/**
 * Represents a request structure to create task.
 *
 * @generated from message flyteidl.admin.CreateTaskRequest
 */
export class CreateTaskRequest extends Message<CreateTaskRequest> {
  /**
   * The inputs required to start the execution. All required inputs must be
   * included in this map. If not required and not provided, defaults apply.
   * +optional
   *
   * @generated from field: flyteidl.core.LiteralMap inputs = 1;
   */
  inputs?: LiteralMap;

  /**
   * Template of the task that encapsulates all the metadata of the task.
   *
   * @generated from field: flyteidl.core.TaskTemplate template = 2;
   */
  template?: TaskTemplate;

  /**
   * Prefix for where task output data will be written. (e.g. s3://my-bucket/randomstring)
   *
   * @generated from field: string output_prefix = 3;
   */
  outputPrefix = "";

  /**
   * subset of runtime task execution metadata.
   *
   * @generated from field: flyteidl.admin.TaskExecutionMetadata task_execution_metadata = 4;
   */
  taskExecutionMetadata?: TaskExecutionMetadata;

  constructor(data?: PartialMessage<CreateTaskRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "flyteidl.admin.CreateTaskRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "inputs", kind: "message", T: LiteralMap },
    { no: 2, name: "template", kind: "message", T: TaskTemplate },
    { no: 3, name: "output_prefix", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 4, name: "task_execution_metadata", kind: "message", T: TaskExecutionMetadata },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): CreateTaskRequest {
    return new CreateTaskRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): CreateTaskRequest {
    return new CreateTaskRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): CreateTaskRequest {
    return new CreateTaskRequest().fromJsonString(jsonString, options);
  }

  static equals(a: CreateTaskRequest | PlainMessage<CreateTaskRequest> | undefined, b: CreateTaskRequest | PlainMessage<CreateTaskRequest> | undefined): boolean {
    return proto3.util.equals(CreateTaskRequest, a, b);
  }
}

/**
 * Represents a create response structure.
 *
 * @generated from message flyteidl.admin.CreateTaskResponse
 */
export class CreateTaskResponse extends Message<CreateTaskResponse> {
  /**
   * Metadata is created by the agent. It could be a string (jobId) or a dict (more complex metadata).
   * Resource is for synchronous task execution.
   *
   * @generated from oneof flyteidl.admin.CreateTaskResponse.res
   */
  res: {
    /**
     * @generated from field: bytes resource_meta = 1;
     */
    value: Uint8Array;
    case: "resourceMeta";
  } | {
    /**
     * @generated from field: flyteidl.admin.Resource resource = 2;
     */
    value: Resource;
    case: "resource";
  } | { case: undefined; value?: undefined } = { case: undefined };

  constructor(data?: PartialMessage<CreateTaskResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "flyteidl.admin.CreateTaskResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "resource_meta", kind: "scalar", T: 12 /* ScalarType.BYTES */, oneof: "res" },
    { no: 2, name: "resource", kind: "message", T: Resource, oneof: "res" },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): CreateTaskResponse {
    return new CreateTaskResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): CreateTaskResponse {
    return new CreateTaskResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): CreateTaskResponse {
    return new CreateTaskResponse().fromJsonString(jsonString, options);
  }

  static equals(a: CreateTaskResponse | PlainMessage<CreateTaskResponse> | undefined, b: CreateTaskResponse | PlainMessage<CreateTaskResponse> | undefined): boolean {
    return proto3.util.equals(CreateTaskResponse, a, b);
  }
}

/**
 * A message used to fetch a job resource from flyte agent server.
 *
 * @generated from message flyteidl.admin.GetTaskRequest
 */
export class GetTaskRequest extends Message<GetTaskRequest> {
  /**
   * A predefined yet extensible Task type identifier.
   *
   * @generated from field: string task_type = 1;
   */
  taskType = "";

  /**
   * Metadata about the resource to be pass to the agent.
   *
   * @generated from field: bytes resource_meta = 2;
   */
  resourceMeta = new Uint8Array(0);

  constructor(data?: PartialMessage<GetTaskRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "flyteidl.admin.GetTaskRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "task_type", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "resource_meta", kind: "scalar", T: 12 /* ScalarType.BYTES */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetTaskRequest {
    return new GetTaskRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetTaskRequest {
    return new GetTaskRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetTaskRequest {
    return new GetTaskRequest().fromJsonString(jsonString, options);
  }

  static equals(a: GetTaskRequest | PlainMessage<GetTaskRequest> | undefined, b: GetTaskRequest | PlainMessage<GetTaskRequest> | undefined): boolean {
    return proto3.util.equals(GetTaskRequest, a, b);
  }
}

/**
 * Response to get an individual task resource.
 *
 * @generated from message flyteidl.admin.GetTaskResponse
 */
export class GetTaskResponse extends Message<GetTaskResponse> {
  /**
   * @generated from field: flyteidl.admin.Resource resource = 1;
   */
  resource?: Resource;

  /**
   * log information for the task execution
   *
   * @generated from field: repeated flyteidl.core.TaskLog log_links = 2;
   */
  logLinks: TaskLog[] = [];

  constructor(data?: PartialMessage<GetTaskResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "flyteidl.admin.GetTaskResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "resource", kind: "message", T: Resource },
    { no: 2, name: "log_links", kind: "message", T: TaskLog, repeated: true },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetTaskResponse {
    return new GetTaskResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetTaskResponse {
    return new GetTaskResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetTaskResponse {
    return new GetTaskResponse().fromJsonString(jsonString, options);
  }

  static equals(a: GetTaskResponse | PlainMessage<GetTaskResponse> | undefined, b: GetTaskResponse | PlainMessage<GetTaskResponse> | undefined): boolean {
    return proto3.util.equals(GetTaskResponse, a, b);
  }
}

/**
 * @generated from message flyteidl.admin.Resource
 */
export class Resource extends Message<Resource> {
  /**
   * DEPRECATED. The state of the execution is used to control its visibility in the UI/CLI.
   *
   * @generated from field: flyteidl.admin.State state = 1 [deprecated = true];
   * @deprecated
   */
  state = State.RETRYABLE_FAILURE;

  /**
   * The outputs of the execution. It's typically used by sql task. Agent service will create a
   * Structured dataset pointing to the query result table.
   * +optional
   *
   * @generated from field: flyteidl.core.LiteralMap outputs = 2;
   */
  outputs?: LiteralMap;

  /**
   * A descriptive message for the current state. e.g. waiting for cluster.
   *
   * @generated from field: string message = 3;
   */
  message = "";

  /**
   * log information for the task execution.
   *
   * @generated from field: repeated flyteidl.core.TaskLog log_links = 4;
   */
  logLinks: TaskLog[] = [];

  /**
   * The phase of the execution is used to determine the phase of the plugin's execution.
   *
   * @generated from field: flyteidl.core.TaskExecution.Phase phase = 5;
   */
  phase = TaskExecution_Phase.UNDEFINED;

  constructor(data?: PartialMessage<Resource>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "flyteidl.admin.Resource";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "state", kind: "enum", T: proto3.getEnumType(State) },
    { no: 2, name: "outputs", kind: "message", T: LiteralMap },
    { no: 3, name: "message", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 4, name: "log_links", kind: "message", T: TaskLog, repeated: true },
    { no: 5, name: "phase", kind: "enum", T: proto3.getEnumType(TaskExecution_Phase) },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Resource {
    return new Resource().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Resource {
    return new Resource().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Resource {
    return new Resource().fromJsonString(jsonString, options);
  }

  static equals(a: Resource | PlainMessage<Resource> | undefined, b: Resource | PlainMessage<Resource> | undefined): boolean {
    return proto3.util.equals(Resource, a, b);
  }
}

/**
 * A message used to delete a task.
 *
 * @generated from message flyteidl.admin.DeleteTaskRequest
 */
export class DeleteTaskRequest extends Message<DeleteTaskRequest> {
  /**
   * A predefined yet extensible Task type identifier.
   *
   * @generated from field: string task_type = 1;
   */
  taskType = "";

  /**
   * Metadata about the resource to be pass to the agent.
   *
   * @generated from field: bytes resource_meta = 2;
   */
  resourceMeta = new Uint8Array(0);

  constructor(data?: PartialMessage<DeleteTaskRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "flyteidl.admin.DeleteTaskRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "task_type", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "resource_meta", kind: "scalar", T: 12 /* ScalarType.BYTES */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): DeleteTaskRequest {
    return new DeleteTaskRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): DeleteTaskRequest {
    return new DeleteTaskRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): DeleteTaskRequest {
    return new DeleteTaskRequest().fromJsonString(jsonString, options);
  }

  static equals(a: DeleteTaskRequest | PlainMessage<DeleteTaskRequest> | undefined, b: DeleteTaskRequest | PlainMessage<DeleteTaskRequest> | undefined): boolean {
    return proto3.util.equals(DeleteTaskRequest, a, b);
  }
}

/**
 * Response to delete a task.
 *
 * @generated from message flyteidl.admin.DeleteTaskResponse
 */
export class DeleteTaskResponse extends Message<DeleteTaskResponse> {
  constructor(data?: PartialMessage<DeleteTaskResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "flyteidl.admin.DeleteTaskResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): DeleteTaskResponse {
    return new DeleteTaskResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): DeleteTaskResponse {
    return new DeleteTaskResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): DeleteTaskResponse {
    return new DeleteTaskResponse().fromJsonString(jsonString, options);
  }

  static equals(a: DeleteTaskResponse | PlainMessage<DeleteTaskResponse> | undefined, b: DeleteTaskResponse | PlainMessage<DeleteTaskResponse> | undefined): boolean {
    return proto3.util.equals(DeleteTaskResponse, a, b);
  }
}

/**
 * A message containing the agent metadata.
 *
 * @generated from message flyteidl.admin.Agent
 */
export class Agent extends Message<Agent> {
  /**
   * Name is the developer-assigned name of the agent.
   *
   * @generated from field: string name = 1;
   */
  name = "";

  /**
   * SupportedTaskTypes are the types of the tasks that the agent can handle.
   *
   * @generated from field: repeated string supported_task_types = 2;
   */
  supportedTaskTypes: string[] = [];

  constructor(data?: PartialMessage<Agent>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "flyteidl.admin.Agent";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "supported_task_types", kind: "scalar", T: 9 /* ScalarType.STRING */, repeated: true },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Agent {
    return new Agent().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Agent {
    return new Agent().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Agent {
    return new Agent().fromJsonString(jsonString, options);
  }

  static equals(a: Agent | PlainMessage<Agent> | undefined, b: Agent | PlainMessage<Agent> | undefined): boolean {
    return proto3.util.equals(Agent, a, b);
  }
}

/**
 * A request to get an agent.
 *
 * @generated from message flyteidl.admin.GetAgentRequest
 */
export class GetAgentRequest extends Message<GetAgentRequest> {
  /**
   * The name of the agent.
   *
   * @generated from field: string name = 1;
   */
  name = "";

  constructor(data?: PartialMessage<GetAgentRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "flyteidl.admin.GetAgentRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetAgentRequest {
    return new GetAgentRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetAgentRequest {
    return new GetAgentRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetAgentRequest {
    return new GetAgentRequest().fromJsonString(jsonString, options);
  }

  static equals(a: GetAgentRequest | PlainMessage<GetAgentRequest> | undefined, b: GetAgentRequest | PlainMessage<GetAgentRequest> | undefined): boolean {
    return proto3.util.equals(GetAgentRequest, a, b);
  }
}

/**
 * A response containing an agent.
 *
 * @generated from message flyteidl.admin.GetAgentResponse
 */
export class GetAgentResponse extends Message<GetAgentResponse> {
  /**
   * @generated from field: flyteidl.admin.Agent agent = 1;
   */
  agent?: Agent;

  constructor(data?: PartialMessage<GetAgentResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "flyteidl.admin.GetAgentResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "agent", kind: "message", T: Agent },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetAgentResponse {
    return new GetAgentResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetAgentResponse {
    return new GetAgentResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetAgentResponse {
    return new GetAgentResponse().fromJsonString(jsonString, options);
  }

  static equals(a: GetAgentResponse | PlainMessage<GetAgentResponse> | undefined, b: GetAgentResponse | PlainMessage<GetAgentResponse> | undefined): boolean {
    return proto3.util.equals(GetAgentResponse, a, b);
  }
}

/**
 * A request to list all agents.
 *
 * @generated from message flyteidl.admin.ListAgentsRequest
 */
export class ListAgentsRequest extends Message<ListAgentsRequest> {
  constructor(data?: PartialMessage<ListAgentsRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "flyteidl.admin.ListAgentsRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ListAgentsRequest {
    return new ListAgentsRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ListAgentsRequest {
    return new ListAgentsRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ListAgentsRequest {
    return new ListAgentsRequest().fromJsonString(jsonString, options);
  }

  static equals(a: ListAgentsRequest | PlainMessage<ListAgentsRequest> | undefined, b: ListAgentsRequest | PlainMessage<ListAgentsRequest> | undefined): boolean {
    return proto3.util.equals(ListAgentsRequest, a, b);
  }
}

/**
 * A response containing a list of agents.
 *
 * @generated from message flyteidl.admin.ListAgentsResponse
 */
export class ListAgentsResponse extends Message<ListAgentsResponse> {
  /**
   * @generated from field: repeated flyteidl.admin.Agent agents = 1;
   */
  agents: Agent[] = [];

  constructor(data?: PartialMessage<ListAgentsResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "flyteidl.admin.ListAgentsResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "agents", kind: "message", T: Agent, repeated: true },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ListAgentsResponse {
    return new ListAgentsResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ListAgentsResponse {
    return new ListAgentsResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ListAgentsResponse {
    return new ListAgentsResponse().fromJsonString(jsonString, options);
  }

  static equals(a: ListAgentsResponse | PlainMessage<ListAgentsResponse> | undefined, b: ListAgentsResponse | PlainMessage<ListAgentsResponse> | undefined): boolean {
    return proto3.util.equals(ListAgentsResponse, a, b);
  }
}

/**
 * A request to get the metrics from a task execution.
 *
 * @generated from message flyteidl.admin.GetTaskMetricsRequest
 */
export class GetTaskMetricsRequest extends Message<GetTaskMetricsRequest> {
  /**
   * A predefined yet extensible Task type identifier.
   *
   * @generated from field: string task_type = 1;
   */
  taskType = "";

  /**
   * Metadata is created by the agent. It could be a string (jobId) or a dict (more complex metadata).
   *
   * @generated from field: bytes resource_meta = 2;
   */
  resourceMeta = new Uint8Array(0);

  /**
   * The metrics to query. If empty, will return a default set of metrics.
   * e.g. EXECUTION_METRIC_USED_CPU_AVG or EXECUTION_METRIC_USED_MEMORY_BYTES_AVG
   *
   * @generated from field: repeated string queries = 3;
   */
  queries: string[] = [];

  /**
   * Start timestamp, inclusive.
   *
   * @generated from field: google.protobuf.Timestamp start_time = 4;
   */
  startTime?: Timestamp;

  /**
   * End timestamp, inclusive..
   *
   * @generated from field: google.protobuf.Timestamp end_time = 5;
   */
  endTime?: Timestamp;

  /**
   * Query resolution step width in duration format or float number of seconds.
   *
   * @generated from field: google.protobuf.Duration step = 6;
   */
  step?: Duration;

  constructor(data?: PartialMessage<GetTaskMetricsRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "flyteidl.admin.GetTaskMetricsRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "task_type", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "resource_meta", kind: "scalar", T: 12 /* ScalarType.BYTES */ },
    { no: 3, name: "queries", kind: "scalar", T: 9 /* ScalarType.STRING */, repeated: true },
    { no: 4, name: "start_time", kind: "message", T: Timestamp },
    { no: 5, name: "end_time", kind: "message", T: Timestamp },
    { no: 6, name: "step", kind: "message", T: Duration },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetTaskMetricsRequest {
    return new GetTaskMetricsRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetTaskMetricsRequest {
    return new GetTaskMetricsRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetTaskMetricsRequest {
    return new GetTaskMetricsRequest().fromJsonString(jsonString, options);
  }

  static equals(a: GetTaskMetricsRequest | PlainMessage<GetTaskMetricsRequest> | undefined, b: GetTaskMetricsRequest | PlainMessage<GetTaskMetricsRequest> | undefined): boolean {
    return proto3.util.equals(GetTaskMetricsRequest, a, b);
  }
}

/**
 * A response containing a list of metrics for a task execution.
 *
 * @generated from message flyteidl.admin.GetTaskMetricsResponse
 */
export class GetTaskMetricsResponse extends Message<GetTaskMetricsResponse> {
  /**
   * The execution metric results.
   *
   * @generated from field: repeated flyteidl.core.ExecutionMetricResult results = 1;
   */
  results: ExecutionMetricResult[] = [];

  constructor(data?: PartialMessage<GetTaskMetricsResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "flyteidl.admin.GetTaskMetricsResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "results", kind: "message", T: ExecutionMetricResult, repeated: true },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetTaskMetricsResponse {
    return new GetTaskMetricsResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetTaskMetricsResponse {
    return new GetTaskMetricsResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetTaskMetricsResponse {
    return new GetTaskMetricsResponse().fromJsonString(jsonString, options);
  }

  static equals(a: GetTaskMetricsResponse | PlainMessage<GetTaskMetricsResponse> | undefined, b: GetTaskMetricsResponse | PlainMessage<GetTaskMetricsResponse> | undefined): boolean {
    return proto3.util.equals(GetTaskMetricsResponse, a, b);
  }
}

/**
 * A request to get the log from a task execution.
 *
 * @generated from message flyteidl.admin.GetTaskLogsRequest
 */
export class GetTaskLogsRequest extends Message<GetTaskLogsRequest> {
  /**
   * A predefined yet extensible Task type identifier.
   *
   * @generated from field: string task_type = 1;
   */
  taskType = "";

  /**
   * Metadata is created by the agent. It could be a string (jobId) or a dict (more complex metadata).
   *
   * @generated from field: bytes resource_meta = 2;
   */
  resourceMeta = new Uint8Array(0);

  /**
   * Number of lines to return.
   *
   * @generated from field: uint64 lines = 3;
   */
  lines = protoInt64.zero;

  /**
   * In the case of multiple pages of results, the server-provided token can be used to fetch the next page
   * in a query. If there are no more results, this value will be empty.
   *
   * @generated from field: string token = 4;
   */
  token = "";

  constructor(data?: PartialMessage<GetTaskLogsRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "flyteidl.admin.GetTaskLogsRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "task_type", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "resource_meta", kind: "scalar", T: 12 /* ScalarType.BYTES */ },
    { no: 3, name: "lines", kind: "scalar", T: 4 /* ScalarType.UINT64 */ },
    { no: 4, name: "token", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetTaskLogsRequest {
    return new GetTaskLogsRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetTaskLogsRequest {
    return new GetTaskLogsRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetTaskLogsRequest {
    return new GetTaskLogsRequest().fromJsonString(jsonString, options);
  }

  static equals(a: GetTaskLogsRequest | PlainMessage<GetTaskLogsRequest> | undefined, b: GetTaskLogsRequest | PlainMessage<GetTaskLogsRequest> | undefined): boolean {
    return proto3.util.equals(GetTaskLogsRequest, a, b);
  }
}

/**
 * A response containing the logs for a task execution.
 *
 * @generated from message flyteidl.admin.GetTaskLogsResponse
 */
export class GetTaskLogsResponse extends Message<GetTaskLogsResponse> {
  /**
   * The execution log results.
   *
   * @generated from field: repeated string results = 1;
   */
  results: string[] = [];

  /**
   * In the case of multiple pages of results, the server-provided token can be used to fetch the next page
   * in a query. If there are no more results, this value will be empty.
   *
   * @generated from field: string token = 2;
   */
  token = "";

  constructor(data?: PartialMessage<GetTaskLogsResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "flyteidl.admin.GetTaskLogsResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "results", kind: "scalar", T: 9 /* ScalarType.STRING */, repeated: true },
    { no: 2, name: "token", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetTaskLogsResponse {
    return new GetTaskLogsResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetTaskLogsResponse {
    return new GetTaskLogsResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetTaskLogsResponse {
    return new GetTaskLogsResponse().fromJsonString(jsonString, options);
  }

  static equals(a: GetTaskLogsResponse | PlainMessage<GetTaskLogsResponse> | undefined, b: GetTaskLogsResponse | PlainMessage<GetTaskLogsResponse> | undefined): boolean {
    return proto3.util.equals(GetTaskLogsResponse, a, b);
  }
}

