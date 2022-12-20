# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: flyteidl/admin/task_execution.proto
"""Generated protocol buffer code."""
from google.protobuf.internal import builder as _builder
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from flyteidl.admin import common_pb2 as flyteidl_dot_admin_dot_common__pb2
from flyteidl.core import execution_pb2 as flyteidl_dot_core_dot_execution__pb2
from flyteidl.core import identifier_pb2 as flyteidl_dot_core_dot_identifier__pb2
from flyteidl.core import literals_pb2 as flyteidl_dot_core_dot_literals__pb2
from flyteidl.event import event_pb2 as flyteidl_dot_event_dot_event__pb2
from google.protobuf import timestamp_pb2 as google_dot_protobuf_dot_timestamp__pb2
from google.protobuf import duration_pb2 as google_dot_protobuf_dot_duration__pb2
from google.protobuf import struct_pb2 as google_dot_protobuf_dot_struct__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n#flyteidl/admin/task_execution.proto\x12\x0e\x66lyteidl.admin\x1a\x1b\x66lyteidl/admin/common.proto\x1a\x1d\x66lyteidl/core/execution.proto\x1a\x1e\x66lyteidl/core/identifier.proto\x1a\x1c\x66lyteidl/core/literals.proto\x1a\x1a\x66lyteidl/event/event.proto\x1a\x1fgoogle/protobuf/timestamp.proto\x1a\x1egoogle/protobuf/duration.proto\x1a\x1cgoogle/protobuf/struct.proto\"Q\n\x17TaskExecutionGetRequest\x12\x36\n\x02id\x18\x01 \x01(\x0b\x32&.flyteidl.core.TaskExecutionIdentifierR\x02id\"\xe3\x01\n\x18TaskExecutionListRequest\x12R\n\x11node_execution_id\x18\x01 \x01(\x0b\x32&.flyteidl.core.NodeExecutionIdentifierR\x0fnodeExecutionId\x12\x14\n\x05limit\x18\x02 \x01(\rR\x05limit\x12\x14\n\x05token\x18\x03 \x01(\tR\x05token\x12\x18\n\x07\x66ilters\x18\x04 \x01(\tR\x07\x66ilters\x12-\n\x07sort_by\x18\x05 \x01(\x0b\x32\x14.flyteidl.admin.SortR\x06sortBy\"\xc1\x01\n\rTaskExecution\x12\x36\n\x02id\x18\x01 \x01(\x0b\x32&.flyteidl.core.TaskExecutionIdentifierR\x02id\x12\x1b\n\tinput_uri\x18\x02 \x01(\tR\x08inputUri\x12>\n\x07\x63losure\x18\x03 \x01(\x0b\x32$.flyteidl.admin.TaskExecutionClosureR\x07\x63losure\x12\x1b\n\tis_parent\x18\x04 \x01(\x08R\x08isParent\"q\n\x11TaskExecutionList\x12\x46\n\x0ftask_executions\x18\x01 \x03(\x0b\x32\x1d.flyteidl.admin.TaskExecutionR\x0etaskExecutions\x12\x14\n\x05token\x18\x02 \x01(\tR\x05token\"\xea\x05\n\x14TaskExecutionClosure\x12#\n\noutput_uri\x18\x01 \x01(\tB\x02\x18\x01H\x00R\toutputUri\x12\x35\n\x05\x65rror\x18\x02 \x01(\x0b\x32\x1d.flyteidl.core.ExecutionErrorH\x00R\x05\x65rror\x12@\n\x0boutput_data\x18\x0c \x01(\x0b\x32\x19.flyteidl.core.LiteralMapB\x02\x18\x01H\x00R\noutputData\x12\x38\n\x05phase\x18\x03 \x01(\x0e\x32\".flyteidl.core.TaskExecution.PhaseR\x05phase\x12*\n\x04logs\x18\x04 \x03(\x0b\x32\x16.flyteidl.core.TaskLogR\x04logs\x12\x39\n\nstarted_at\x18\x05 \x01(\x0b\x32\x1a.google.protobuf.TimestampR\tstartedAt\x12\x35\n\x08\x64uration\x18\x06 \x01(\x0b\x32\x19.google.protobuf.DurationR\x08\x64uration\x12\x39\n\ncreated_at\x18\x07 \x01(\x0b\x32\x1a.google.protobuf.TimestampR\tcreatedAt\x12\x39\n\nupdated_at\x18\x08 \x01(\x0b\x32\x1a.google.protobuf.TimestampR\tupdatedAt\x12\x38\n\x0b\x63ustom_info\x18\t \x01(\x0b\x32\x17.google.protobuf.StructR\ncustomInfo\x12\x16\n\x06reason\x18\n \x01(\tR\x06reason\x12\x1b\n\ttask_type\x18\x0b \x01(\tR\x08taskType\x12\x41\n\x08metadata\x18\x10 \x01(\x0b\x32%.flyteidl.event.TaskExecutionMetadataR\x08metadata\x12#\n\revent_version\x18\x11 \x01(\x05R\x0c\x65ventVersionB\x0f\n\routput_result\"U\n\x1bTaskExecutionGetDataRequest\x12\x36\n\x02id\x18\x01 \x01(\x0b\x32&.flyteidl.core.TaskExecutionIdentifierR\x02id\"\x84\x02\n\x1cTaskExecutionGetDataResponse\x12\x33\n\x06inputs\x18\x01 \x01(\x0b\x32\x17.flyteidl.admin.UrlBlobB\x02\x18\x01R\x06inputs\x12\x35\n\x07outputs\x18\x02 \x01(\x0b\x32\x17.flyteidl.admin.UrlBlobB\x02\x18\x01R\x07outputs\x12:\n\x0b\x66ull_inputs\x18\x03 \x01(\x0b\x32\x19.flyteidl.core.LiteralMapR\nfullInputs\x12<\n\x0c\x66ull_outputs\x18\x04 \x01(\x0b\x32\x19.flyteidl.core.LiteralMapR\x0b\x66ullOutputsB\xb8\x01\n\x12\x63om.flyteidl.adminB\x12TaskExecutionProtoP\x01Z5github.com/flyteorg/flyteidl/gen/pb-go/flyteidl/admin\xa2\x02\x03\x46\x41X\xaa\x02\x0e\x46lyteidl.Admin\xca\x02\x0e\x46lyteidl\\Admin\xe2\x02\x1a\x46lyteidl\\Admin\\GPBMetadata\xea\x02\x0f\x46lyteidl::Adminb\x06proto3')

_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, globals())
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'flyteidl.admin.task_execution_pb2', globals())
if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'\n\022com.flyteidl.adminB\022TaskExecutionProtoP\001Z5github.com/flyteorg/flyteidl/gen/pb-go/flyteidl/admin\242\002\003FAX\252\002\016Flyteidl.Admin\312\002\016Flyteidl\\Admin\342\002\032Flyteidl\\Admin\\GPBMetadata\352\002\017Flyteidl::Admin'
  _TASKEXECUTIONCLOSURE.fields_by_name['output_uri']._options = None
  _TASKEXECUTIONCLOSURE.fields_by_name['output_uri']._serialized_options = b'\030\001'
  _TASKEXECUTIONCLOSURE.fields_by_name['output_data']._options = None
  _TASKEXECUTIONCLOSURE.fields_by_name['output_data']._serialized_options = b'\030\001'
  _TASKEXECUTIONGETDATARESPONSE.fields_by_name['inputs']._options = None
  _TASKEXECUTIONGETDATARESPONSE.fields_by_name['inputs']._serialized_options = b'\030\001'
  _TASKEXECUTIONGETDATARESPONSE.fields_by_name['outputs']._options = None
  _TASKEXECUTIONGETDATARESPONSE.fields_by_name['outputs']._serialized_options = b'\030\001'
  _TASKEXECUTIONGETREQUEST._serialized_start=300
  _TASKEXECUTIONGETREQUEST._serialized_end=381
  _TASKEXECUTIONLISTREQUEST._serialized_start=384
  _TASKEXECUTIONLISTREQUEST._serialized_end=611
  _TASKEXECUTION._serialized_start=614
  _TASKEXECUTION._serialized_end=807
  _TASKEXECUTIONLIST._serialized_start=809
  _TASKEXECUTIONLIST._serialized_end=922
  _TASKEXECUTIONCLOSURE._serialized_start=925
  _TASKEXECUTIONCLOSURE._serialized_end=1671
  _TASKEXECUTIONGETDATAREQUEST._serialized_start=1673
  _TASKEXECUTIONGETDATAREQUEST._serialized_end=1758
  _TASKEXECUTIONGETDATARESPONSE._serialized_start=1761
  _TASKEXECUTIONGETDATARESPONSE._serialized_end=2021
# @@protoc_insertion_point(module_scope)
