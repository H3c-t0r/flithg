# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: flyteidl/service/cache.proto
"""Generated protocol buffer code."""
from google.protobuf.internal import builder as _builder
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from google.api import annotations_pb2 as google_dot_api_dot_annotations__pb2
from flyteidl.core import errors_pb2 as flyteidl_dot_core_dot_errors__pb2
from flyteidl.core import identifier_pb2 as flyteidl_dot_core_dot_identifier__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x1c\x66lyteidl/service/cache.proto\x12\x10\x66lyteidl.service\x1a\x1cgoogle/api/annotations.proto\x1a\x1a\x66lyteidl/core/errors.proto\x1a\x1e\x66lyteidl/core/identifier.proto\"t\n\x1e\x45victTaskExecutionCacheRequest\x12R\n\x11task_execution_id\x18\x01 \x01(\x0b\x32&.flyteidl.core.TaskExecutionIdentifierR\x0ftaskExecutionId\"S\n\x12\x45victCacheResponse\x12=\n\x06\x65rrors\x18\x01 \x01(\x0b\x32%.flyteidl.core.CacheEvictionErrorListR\x06\x65rrors2\xb5\x04\n\x0c\x43\x61\x63heService\x12\xa4\x04\n\x17\x45victTaskExecutionCache\x12\x30.flyteidl.service.EvictTaskExecutionCacheRequest\x1a$.flyteidl.service.EvictCacheResponse\"\xb0\x03\x82\xd3\xe4\x93\x02\xa9\x03*\xa6\x03/api/v1/cache/task_executions/{task_execution_id.node_execution_id.execution_id.project}/{task_execution_id.node_execution_id.execution_id.domain}/{task_execution_id.node_execution_id.execution_id.name}/{task_execution_id.node_execution_id.node_id}/{task_execution_id.task_id.project}/{task_execution_id.task_id.domain}/{task_execution_id.task_id.name}/{task_execution_id.task_id.version}/{task_execution_id.retry_attempt}B\xc2\x01\n\x14\x63om.flyteidl.serviceB\nCacheProtoP\x01Z=github.com/flyteorg/flyte/flyteidl/gen/pb-go/flyteidl/service\xa2\x02\x03\x46SX\xaa\x02\x10\x46lyteidl.Service\xca\x02\x10\x46lyteidl\\Service\xe2\x02\x1c\x46lyteidl\\Service\\GPBMetadata\xea\x02\x11\x46lyteidl::Serviceb\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'flyteidl.service.cache_pb2', _globals)
if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'\n\024com.flyteidl.serviceB\nCacheProtoP\001Z=github.com/flyteorg/flyte/flyteidl/gen/pb-go/flyteidl/service\242\002\003FSX\252\002\020Flyteidl.Service\312\002\020Flyteidl\\Service\342\002\034Flyteidl\\Service\\GPBMetadata\352\002\021Flyteidl::Service'
  _CACHESERVICE.methods_by_name['EvictTaskExecutionCache']._options = None
  _CACHESERVICE.methods_by_name['EvictTaskExecutionCache']._serialized_options = b'\202\323\344\223\002\251\003*\246\003/api/v1/cache/task_executions/{task_execution_id.node_execution_id.execution_id.project}/{task_execution_id.node_execution_id.execution_id.domain}/{task_execution_id.node_execution_id.execution_id.name}/{task_execution_id.node_execution_id.node_id}/{task_execution_id.task_id.project}/{task_execution_id.task_id.domain}/{task_execution_id.task_id.name}/{task_execution_id.task_id.version}/{task_execution_id.retry_attempt}'
  _globals['_EVICTTASKEXECUTIONCACHEREQUEST']._serialized_start=140
  _globals['_EVICTTASKEXECUTIONCACHEREQUEST']._serialized_end=256
  _globals['_EVICTCACHERESPONSE']._serialized_start=258
  _globals['_EVICTCACHERESPONSE']._serialized_end=341
  _globals['_CACHESERVICE']._serialized_start=344
  _globals['_CACHESERVICE']._serialized_end=909
# @@protoc_insertion_point(module_scope)