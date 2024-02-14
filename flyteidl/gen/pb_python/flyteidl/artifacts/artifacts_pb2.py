# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: flyteidl/artifacts/artifacts.proto
"""Generated protocol buffer code."""
from google.protobuf.internal import builder as _builder
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from google.protobuf import struct_pb2 as google_dot_protobuf_dot_struct__pb2
from google.api import annotations_pb2 as google_dot_api_dot_annotations__pb2
from google.protobuf import timestamp_pb2 as google_dot_protobuf_dot_timestamp__pb2
from flyteidl.admin import launch_plan_pb2 as flyteidl_dot_admin_dot_launch__plan__pb2
from flyteidl.core import literals_pb2 as flyteidl_dot_core_dot_literals__pb2
from flyteidl.core import types_pb2 as flyteidl_dot_core_dot_types__pb2
from flyteidl.core import identifier_pb2 as flyteidl_dot_core_dot_identifier__pb2
from flyteidl.core import artifact_id_pb2 as flyteidl_dot_core_dot_artifact__id__pb2
from flyteidl.core import interface_pb2 as flyteidl_dot_core_dot_interface__pb2
from flyteidl.event import cloudevents_pb2 as flyteidl_dot_event_dot_cloudevents__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\"flyteidl/artifacts/artifacts.proto\x12\x11\x66lyteidl.artifact\x1a\x1cgoogle/protobuf/struct.proto\x1a\x1cgoogle/api/annotations.proto\x1a\x1fgoogle/protobuf/timestamp.proto\x1a flyteidl/admin/launch_plan.proto\x1a\x1c\x66lyteidl/core/literals.proto\x1a\x19\x66lyteidl/core/types.proto\x1a\x1e\x66lyteidl/core/identifier.proto\x1a\x1f\x66lyteidl/core/artifact_id.proto\x1a\x1d\x66lyteidl/core/interface.proto\x1a flyteidl/event/cloudevents.proto\"\x8b\x02\n\x08\x41rtifact\x12:\n\x0b\x61rtifact_id\x18\x01 \x01(\x0b\x32\x19.flyteidl.core.ArtifactIDR\nartifactId\x12\x33\n\x04spec\x18\x02 \x01(\x0b\x32\x1f.flyteidl.artifact.ArtifactSpecR\x04spec\x12\x12\n\x04tags\x18\x03 \x03(\tR\x04tags\x12\x39\n\x06source\x18\x04 \x01(\x0b\x32!.flyteidl.artifact.ArtifactSourceR\x06source\x12?\n\x08metadata\x18\x05 \x01(\x0b\x32#.flyteidl.artifact.ArtifactMetadataR\x08metadata\"_\n\x10\x41rtifactMetadata\x12\x39\n\ncreated_at\x18\x01 \x01(\x0b\x32\x1a.google.protobuf.TimestampR\tcreatedAt\x12\x10\n\x03uri\x18\x02 \x01(\tR\x03uri\"\xc7\x03\n\x15\x43reateArtifactRequest\x12=\n\x0c\x61rtifact_key\x18\x01 \x01(\x0b\x32\x1a.flyteidl.core.ArtifactKeyR\x0b\x61rtifactKey\x12\x18\n\x07version\x18\x03 \x01(\tR\x07version\x12\x33\n\x04spec\x18\x02 \x01(\x0b\x32\x1f.flyteidl.artifact.ArtifactSpecR\x04spec\x12X\n\npartitions\x18\x04 \x03(\x0b\x32\x38.flyteidl.artifact.CreateArtifactRequest.PartitionsEntryR\npartitions\x12L\n\x14time_partition_value\x18\x05 \x01(\x0b\x32\x1a.google.protobuf.TimestampR\x12timePartitionValue\x12\x39\n\x06source\x18\x06 \x01(\x0b\x32!.flyteidl.artifact.ArtifactSourceR\x06source\x1a=\n\x0fPartitionsEntry\x12\x10\n\x03key\x18\x01 \x01(\tR\x03key\x12\x14\n\x05value\x18\x02 \x01(\tR\x05value:\x02\x38\x01\"\xfb\x01\n\x0e\x41rtifactSource\x12Y\n\x12workflow_execution\x18\x01 \x01(\x0b\x32*.flyteidl.core.WorkflowExecutionIdentifierR\x11workflowExecution\x12\x17\n\x07node_id\x18\x02 \x01(\tR\x06nodeId\x12\x32\n\x07task_id\x18\x03 \x01(\x0b\x32\x19.flyteidl.core.IdentifierR\x06taskId\x12#\n\rretry_attempt\x18\x04 \x01(\rR\x0cretryAttempt\x12\x1c\n\tprincipal\x18\x05 \x01(\tR\tprincipal\"\xd7\x01\n\x0c\x41rtifactSpec\x12,\n\x05value\x18\x01 \x01(\x0b\x32\x16.flyteidl.core.LiteralR\x05value\x12.\n\x04type\x18\x02 \x01(\x0b\x32\x1a.flyteidl.core.LiteralTypeR\x04type\x12+\n\x11short_description\x18\x03 \x01(\tR\x10shortDescription\x12<\n\ruser_metadata\x18\x04 \x01(\x0b\x32\x17.google.protobuf.StructR\x0cuserMetadata\"z\n\x07Trigger\x12\x38\n\ntrigger_id\x18\x01 \x01(\x0b\x32\x19.flyteidl.core.IdentifierR\ttriggerId\x12\x35\n\x08triggers\x18\x02 \x03(\x0b\x32\x19.flyteidl.core.ArtifactIDR\x08triggers\"Q\n\x16\x43reateArtifactResponse\x12\x37\n\x08\x61rtifact\x18\x01 \x01(\x0b\x32\x1b.flyteidl.artifact.ArtifactR\x08\x61rtifact\"b\n\x12GetArtifactRequest\x12\x32\n\x05query\x18\x01 \x01(\x0b\x32\x1c.flyteidl.core.ArtifactQueryR\x05query\x12\x18\n\x07\x64\x65tails\x18\x02 \x01(\x08R\x07\x64\x65tails\"N\n\x13GetArtifactResponse\x12\x37\n\x08\x61rtifact\x18\x01 \x01(\x0b\x32\x1b.flyteidl.artifact.ArtifactR\x08\x61rtifact\"`\n\rSearchOptions\x12+\n\x11strict_partitions\x18\x01 \x01(\x08R\x10strictPartitions\x12\"\n\rlatest_by_key\x18\x02 \x01(\x08R\x0blatestByKey\"\x80\x03\n\x16SearchArtifactsRequest\x12=\n\x0c\x61rtifact_key\x18\x01 \x01(\x0b\x32\x1a.flyteidl.core.ArtifactKeyR\x0b\x61rtifactKey\x12\x39\n\npartitions\x18\x02 \x01(\x0b\x32\x19.flyteidl.core.PartitionsR\npartitions\x12L\n\x14time_partition_value\x18\x03 \x01(\x0b\x32\x1a.google.protobuf.TimestampR\x12timePartitionValue\x12\x1c\n\tprincipal\x18\x04 \x01(\tR\tprincipal\x12\x18\n\x07version\x18\x05 \x01(\tR\x07version\x12:\n\x07options\x18\x06 \x01(\x0b\x32 .flyteidl.artifact.SearchOptionsR\x07options\x12\x14\n\x05token\x18\x07 \x01(\tR\x05token\x12\x14\n\x05limit\x18\x08 \x01(\x05R\x05limit\"j\n\x17SearchArtifactsResponse\x12\x39\n\tartifacts\x18\x01 \x03(\x0b\x32\x1b.flyteidl.artifact.ArtifactR\tartifacts\x12\x14\n\x05token\x18\x02 \x01(\tR\x05token\"\xdc\x01\n\x19\x46indByWorkflowExecRequest\x12\x43\n\x07\x65xec_id\x18\x01 \x01(\x0b\x32*.flyteidl.core.WorkflowExecutionIdentifierR\x06\x65xecId\x12T\n\tdirection\x18\x02 \x01(\x0e\x32\x36.flyteidl.artifact.FindByWorkflowExecRequest.DirectionR\tdirection\"$\n\tDirection\x12\n\n\x06INPUTS\x10\x00\x12\x0b\n\x07OUTPUTS\x10\x01\"\x7f\n\rAddTagRequest\x12:\n\x0b\x61rtifact_id\x18\x01 \x01(\x0b\x32\x19.flyteidl.core.ArtifactIDR\nartifactId\x12\x14\n\x05value\x18\x02 \x01(\tR\x05value\x12\x1c\n\toverwrite\x18\x03 \x01(\x08R\toverwrite\"\x10\n\x0e\x41\x64\x64TagResponse\"b\n\x14\x43reateTriggerRequest\x12J\n\x13trigger_launch_plan\x18\x01 \x01(\x0b\x32\x1a.flyteidl.admin.LaunchPlanR\x11triggerLaunchPlan\"\x17\n\x15\x43reateTriggerResponse\"T\n\x18\x44\x65\x61\x63tivateTriggerRequest\x12\x38\n\ntrigger_id\x18\x01 \x01(\x0b\x32\x19.flyteidl.core.IdentifierR\ttriggerId\"\x1b\n\x19\x44\x65\x61\x63tivateTriggerResponse\"\x1e\n\x1c\x44\x65\x61\x63tivateAllTriggersRequest\"H\n\x1d\x44\x65\x61\x63tivateAllTriggersResponse\x12\'\n\x0fnum_deactivated\x18\x01 \x01(\x04R\x0enumDeactivated\"\x80\x01\n\x10\x41rtifactProducer\x12\x36\n\tentity_id\x18\x01 \x01(\x0b\x32\x19.flyteidl.core.IdentifierR\x08\x65ntityId\x12\x34\n\x07outputs\x18\x02 \x01(\x0b\x32\x1a.flyteidl.core.VariableMapR\x07outputs\"\\\n\x17RegisterProducerRequest\x12\x41\n\tproducers\x18\x01 \x03(\x0b\x32#.flyteidl.artifact.ArtifactProducerR\tproducers\"\x7f\n\x10\x41rtifactConsumer\x12\x36\n\tentity_id\x18\x01 \x01(\x0b\x32\x19.flyteidl.core.IdentifierR\x08\x65ntityId\x12\x33\n\x06inputs\x18\x02 \x01(\x0b\x32\x1b.flyteidl.core.ParameterMapR\x06inputs\"\\\n\x17RegisterConsumerRequest\x12\x41\n\tconsumers\x18\x01 \x03(\x0b\x32#.flyteidl.artifact.ArtifactConsumerR\tconsumers\"\x12\n\x10RegisterResponse\"\x9a\x01\n\x16\x45xecutionInputsRequest\x12M\n\x0c\x65xecution_id\x18\x01 \x01(\x0b\x32*.flyteidl.core.WorkflowExecutionIdentifierR\x0b\x65xecutionId\x12\x31\n\x06inputs\x18\x02 \x03(\x0b\x32\x19.flyteidl.core.ArtifactIDR\x06inputs\"\x19\n\x17\x45xecutionInputsResponse\"N\n\x10ListUsageRequest\x12:\n\x0b\x61rtifact_id\x18\x01 \x01(\x0b\x32\x19.flyteidl.core.ArtifactIDR\nartifactId\"_\n\x11ListUsageResponse\x12J\n\nexecutions\x18\x01 \x03(\x0b\x32*.flyteidl.core.WorkflowExecutionIdentifierR\nexecutions2\xad\r\n\x10\x41rtifactRegistry\x12g\n\x0e\x43reateArtifact\x12(.flyteidl.artifact.CreateArtifactRequest\x1a).flyteidl.artifact.CreateArtifactResponse\"\x00\x12\x84\x01\n\x0bGetArtifact\x12%.flyteidl.artifact.GetArtifactRequest\x1a&.flyteidl.artifact.GetArtifactResponse\"&\x82\xd3\xe4\x93\x02 :\x01*\"\x1b/artifacts/api/v1/artifacts\x12\x8d\x01\n\x0fSearchArtifacts\x12).flyteidl.artifact.SearchArtifactsRequest\x1a*.flyteidl.artifact.SearchArtifactsResponse\"#\x82\xd3\xe4\x93\x02\x1d:\x01*\"\x18/artifacts/api/v1/search\x12\x64\n\rCreateTrigger\x12\'.flyteidl.artifact.CreateTriggerRequest\x1a(.flyteidl.artifact.CreateTriggerResponse\"\x00\x12\x9f\x01\n\x11\x44\x65\x61\x63tivateTrigger\x12+.flyteidl.artifact.DeactivateTriggerRequest\x1a,.flyteidl.artifact.DeactivateTriggerResponse\"/\x82\xd3\xe4\x93\x02):\x01*2$/artifacts/api/v1/trigger/deactivate\x12\xaf\x01\n\x15\x44\x65\x61\x63tivateAllTriggers\x12/.flyteidl.artifact.DeactivateAllTriggersRequest\x1a\x30.flyteidl.artifact.DeactivateAllTriggersResponse\"3\x82\xd3\xe4\x93\x02-:\x01*2(/artifacts/api/v1/trigger/deactivate/all\x12O\n\x06\x41\x64\x64Tag\x12 .flyteidl.artifact.AddTagRequest\x1a!.flyteidl.artifact.AddTagResponse\"\x00\x12\x65\n\x10RegisterProducer\x12*.flyteidl.artifact.RegisterProducerRequest\x1a#.flyteidl.artifact.RegisterResponse\"\x00\x12\x65\n\x10RegisterConsumer\x12*.flyteidl.artifact.RegisterConsumerRequest\x1a#.flyteidl.artifact.RegisterResponse\"\x00\x12m\n\x12SetExecutionInputs\x12).flyteidl.artifact.ExecutionInputsRequest\x1a*.flyteidl.artifact.ExecutionInputsResponse\"\x00\x12\xd8\x01\n\x12\x46indByWorkflowExec\x12,.flyteidl.artifact.FindByWorkflowExecRequest\x1a*.flyteidl.artifact.SearchArtifactsResponse\"h\x82\xd3\xe4\x93\x02\x62\x12`/artifacts/api/v1/search/execution/{exec_id.project}/{exec_id.domain}/{exec_id.name}/{direction}\x12\xf5\x01\n\tListUsage\x12#.flyteidl.artifact.ListUsageRequest\x1a$.flyteidl.artifact.ListUsageResponse\"\x9c\x01\x82\xd3\xe4\x93\x02\x95\x01\x12\x92\x01/artifacts/api/v1/usage/{artifact_id.artifact_key.project}/{artifact_id.artifact_key.domain}/{artifact_id.artifact_key.name}/{artifact_id.version}B\xcd\x01\n\x15\x63om.flyteidl.artifactB\x0e\x41rtifactsProtoP\x01Z?github.com/flyteorg/flyte/flyteidl/gen/pb-go/flyteidl/artifacts\xa2\x02\x03\x46\x41X\xaa\x02\x11\x46lyteidl.Artifact\xca\x02\x11\x46lyteidl\\Artifact\xe2\x02\x1d\x46lyteidl\\Artifact\\GPBMetadata\xea\x02\x12\x46lyteidl::Artifactb\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'flyteidl.artifacts.artifacts_pb2', _globals)
if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'\n\025com.flyteidl.artifactB\016ArtifactsProtoP\001Z?github.com/flyteorg/flyte/flyteidl/gen/pb-go/flyteidl/artifacts\242\002\003FAX\252\002\021Flyteidl.Artifact\312\002\021Flyteidl\\Artifact\342\002\035Flyteidl\\Artifact\\GPBMetadata\352\002\022Flyteidl::Artifact'
  _CREATEARTIFACTREQUEST_PARTITIONSENTRY._options = None
  _CREATEARTIFACTREQUEST_PARTITIONSENTRY._serialized_options = b'8\001'
  _ARTIFACTREGISTRY.methods_by_name['GetArtifact']._options = None
  _ARTIFACTREGISTRY.methods_by_name['GetArtifact']._serialized_options = b'\202\323\344\223\002 :\001*\"\033/artifacts/api/v1/artifacts'
  _ARTIFACTREGISTRY.methods_by_name['SearchArtifacts']._options = None
  _ARTIFACTREGISTRY.methods_by_name['SearchArtifacts']._serialized_options = b'\202\323\344\223\002\035:\001*\"\030/artifacts/api/v1/search'
  _ARTIFACTREGISTRY.methods_by_name['DeactivateTrigger']._options = None
  _ARTIFACTREGISTRY.methods_by_name['DeactivateTrigger']._serialized_options = b'\202\323\344\223\002):\001*2$/artifacts/api/v1/trigger/deactivate'
  _ARTIFACTREGISTRY.methods_by_name['DeactivateAllTriggers']._options = None
  _ARTIFACTREGISTRY.methods_by_name['DeactivateAllTriggers']._serialized_options = b'\202\323\344\223\002-:\001*2(/artifacts/api/v1/trigger/deactivate/all'
  _ARTIFACTREGISTRY.methods_by_name['FindByWorkflowExec']._options = None
  _ARTIFACTREGISTRY.methods_by_name['FindByWorkflowExec']._serialized_options = b'\202\323\344\223\002b\022`/artifacts/api/v1/search/execution/{exec_id.project}/{exec_id.domain}/{exec_id.name}/{direction}'
  _ARTIFACTREGISTRY.methods_by_name['ListUsage']._options = None
  _ARTIFACTREGISTRY.methods_by_name['ListUsage']._serialized_options = b'\202\323\344\223\002\225\001\022\222\001/artifacts/api/v1/usage/{artifact_id.artifact_key.project}/{artifact_id.artifact_key.domain}/{artifact_id.artifact_key.name}/{artifact_id.version}'
  _globals['_ARTIFACT']._serialized_start=372
  _globals['_ARTIFACT']._serialized_end=639
  _globals['_ARTIFACTMETADATA']._serialized_start=641
  _globals['_ARTIFACTMETADATA']._serialized_end=736
  _globals['_CREATEARTIFACTREQUEST']._serialized_start=739
  _globals['_CREATEARTIFACTREQUEST']._serialized_end=1194
  _globals['_CREATEARTIFACTREQUEST_PARTITIONSENTRY']._serialized_start=1133
  _globals['_CREATEARTIFACTREQUEST_PARTITIONSENTRY']._serialized_end=1194
  _globals['_ARTIFACTSOURCE']._serialized_start=1197
  _globals['_ARTIFACTSOURCE']._serialized_end=1448
  _globals['_ARTIFACTSPEC']._serialized_start=1451
  _globals['_ARTIFACTSPEC']._serialized_end=1666
  _globals['_TRIGGER']._serialized_start=1668
  _globals['_TRIGGER']._serialized_end=1790
  _globals['_CREATEARTIFACTRESPONSE']._serialized_start=1792
  _globals['_CREATEARTIFACTRESPONSE']._serialized_end=1873
  _globals['_GETARTIFACTREQUEST']._serialized_start=1875
  _globals['_GETARTIFACTREQUEST']._serialized_end=1973
  _globals['_GETARTIFACTRESPONSE']._serialized_start=1975
  _globals['_GETARTIFACTRESPONSE']._serialized_end=2053
  _globals['_SEARCHOPTIONS']._serialized_start=2055
  _globals['_SEARCHOPTIONS']._serialized_end=2151
  _globals['_SEARCHARTIFACTSREQUEST']._serialized_start=2154
  _globals['_SEARCHARTIFACTSREQUEST']._serialized_end=2538
  _globals['_SEARCHARTIFACTSRESPONSE']._serialized_start=2540
  _globals['_SEARCHARTIFACTSRESPONSE']._serialized_end=2646
  _globals['_FINDBYWORKFLOWEXECREQUEST']._serialized_start=2649
  _globals['_FINDBYWORKFLOWEXECREQUEST']._serialized_end=2869
  _globals['_FINDBYWORKFLOWEXECREQUEST_DIRECTION']._serialized_start=2833
  _globals['_FINDBYWORKFLOWEXECREQUEST_DIRECTION']._serialized_end=2869
  _globals['_ADDTAGREQUEST']._serialized_start=2871
  _globals['_ADDTAGREQUEST']._serialized_end=2998
  _globals['_ADDTAGRESPONSE']._serialized_start=3000
  _globals['_ADDTAGRESPONSE']._serialized_end=3016
  _globals['_CREATETRIGGERREQUEST']._serialized_start=3018
  _globals['_CREATETRIGGERREQUEST']._serialized_end=3116
  _globals['_CREATETRIGGERRESPONSE']._serialized_start=3118
  _globals['_CREATETRIGGERRESPONSE']._serialized_end=3141
  _globals['_DEACTIVATETRIGGERREQUEST']._serialized_start=3143
  _globals['_DEACTIVATETRIGGERREQUEST']._serialized_end=3227
  _globals['_DEACTIVATETRIGGERRESPONSE']._serialized_start=3229
  _globals['_DEACTIVATETRIGGERRESPONSE']._serialized_end=3256
  _globals['_DEACTIVATEALLTRIGGERSREQUEST']._serialized_start=3258
  _globals['_DEACTIVATEALLTRIGGERSREQUEST']._serialized_end=3288
  _globals['_DEACTIVATEALLTRIGGERSRESPONSE']._serialized_start=3290
  _globals['_DEACTIVATEALLTRIGGERSRESPONSE']._serialized_end=3362
  _globals['_ARTIFACTPRODUCER']._serialized_start=3365
  _globals['_ARTIFACTPRODUCER']._serialized_end=3493
  _globals['_REGISTERPRODUCERREQUEST']._serialized_start=3495
  _globals['_REGISTERPRODUCERREQUEST']._serialized_end=3587
  _globals['_ARTIFACTCONSUMER']._serialized_start=3589
  _globals['_ARTIFACTCONSUMER']._serialized_end=3716
  _globals['_REGISTERCONSUMERREQUEST']._serialized_start=3718
  _globals['_REGISTERCONSUMERREQUEST']._serialized_end=3810
  _globals['_REGISTERRESPONSE']._serialized_start=3812
  _globals['_REGISTERRESPONSE']._serialized_end=3830
  _globals['_EXECUTIONINPUTSREQUEST']._serialized_start=3833
  _globals['_EXECUTIONINPUTSREQUEST']._serialized_end=3987
  _globals['_EXECUTIONINPUTSRESPONSE']._serialized_start=3989
  _globals['_EXECUTIONINPUTSRESPONSE']._serialized_end=4014
  _globals['_LISTUSAGEREQUEST']._serialized_start=4016
  _globals['_LISTUSAGEREQUEST']._serialized_end=4094
  _globals['_LISTUSAGERESPONSE']._serialized_start=4096
  _globals['_LISTUSAGERESPONSE']._serialized_end=4191
  _globals['_ARTIFACTREGISTRY']._serialized_start=4194
  _globals['_ARTIFACTREGISTRY']._serialized_end=5903
# @@protoc_insertion_point(module_scope)