// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: flyteidl/service/admin.proto

package flyteidl.service;

public final class Admin {
  private Admin() {}
  public static void registerAllExtensions(
      com.google.protobuf.ExtensionRegistryLite registry) {
  }

  public static void registerAllExtensions(
      com.google.protobuf.ExtensionRegistry registry) {
    registerAllExtensions(
        (com.google.protobuf.ExtensionRegistryLite) registry);
  }

  public static com.google.protobuf.Descriptors.FileDescriptor
      getDescriptor() {
    return descriptor;
  }
  private static  com.google.protobuf.Descriptors.FileDescriptor
      descriptor;
  static {
    java.lang.String[] descriptorData = {
      "\n\034flyteidl/service/admin.proto\022\020flyteidl" +
      ".service\032\034google/api/annotations.proto\032\034" +
      "flyteidl/admin/project.proto\032.flyteidl/a" +
      "dmin/project_domain_attributes.proto\032\'fl" +
      "yteidl/admin/project_attributes.proto\032\031f" +
      "lyteidl/admin/task.proto\032\035flyteidl/admin" +
      "/workflow.proto\032(flyteidl/admin/workflow" +
      "_attributes.proto\032 flyteidl/admin/launch" +
      "_plan.proto\032\032flyteidl/admin/event.proto\032" +
      "\036flyteidl/admin/execution.proto\032\'flyteid" +
      "l/admin/matchable_resource.proto\032#flytei" +
      "dl/admin/node_execution.proto\032#flyteidl/" +
      "admin/task_execution.proto\032\034flyteidl/adm" +
      "in/version.proto\032\033flyteidl/admin/common." +
      "proto\032\'flyteidl/admin/description_entity" +
      ".proto2\227v\n\014AdminService\022\216\001\n\nCreateTask\022!" +
      ".flyteidl.admin.TaskCreateRequest\032\".flyt" +
      "eidl.admin.TaskCreateResponse\"9\202\323\344\223\0023\"\r/" +
      "api/v1/tasks:\001*Z\037\"\032/api/v1/tasks/org/{id" +
      ".org}:\001*\022\330\001\n\007GetTask\022 .flyteidl.admin.Ob" +
      "jectGetRequest\032\024.flyteidl.admin.Task\"\224\001\202" +
      "\323\344\223\002\215\001\022=/api/v1/tasks/{id.project}/{id.d" +
      "omain}/{id.name}/{id.version}ZL\022J/api/v1" +
      "/tasks/org/{id.org}/{id.project}/{id.dom" +
      "ain}/{id.name}/{id.version}\022\305\001\n\013ListTask" +
      "Ids\0220.flyteidl.admin.NamedEntityIdentifi" +
      "erListRequest\032).flyteidl.admin.NamedEnti" +
      "tyIdentifierList\"Y\202\323\344\223\002S\022#/api/v1/task_i" +
      "ds/{project}/{domain}Z,\022*/api/v1/tasks/o" +
      "rg/{org}/{project}/{domain}\022\250\002\n\tListTask" +
      "s\022#.flyteidl.admin.ResourceListRequest\032\030" +
      ".flyteidl.admin.TaskList\"\333\001\202\323\344\223\002\324\001\0220/api" +
      "/v1/tasks/{id.project}/{id.domain}/{id.n" +
      "ame}Z?\022=/api/v1/tasks/org/{id.org}/{id.p" +
      "roject}/{id.domain}/{id.name}Z(\022&/api/v1" +
      "/tasks/{id.project}/{id.domain}Z5\0223/api/" +
      "v1/tasks/org/{id.org}/{id.project}/{id.d" +
      "omain}\022\242\001\n\016CreateWorkflow\022%.flyteidl.adm" +
      "in.WorkflowCreateRequest\032&.flyteidl.admi" +
      "n.WorkflowCreateResponse\"A\202\323\344\223\002;\"\021/api/v" +
      "1/workflows:\001*Z#\"\036/api/v1/workflows/org/" +
      "{id.org}:\001*\022\350\001\n\013GetWorkflow\022 .flyteidl.a" +
      "dmin.ObjectGetRequest\032\030.flyteidl.admin.W" +
      "orkflow\"\234\001\202\323\344\223\002\225\001\022A/api/v1/workflows/{id" +
      ".project}/{id.domain}/{id.name}/{id.vers" +
      "ion}ZP\022N/api/v1/workflows/org/{id.org}/{" +
      "id.project}/{id.domain}/{id.name}/{id.ve" +
      "rsion}\022\321\001\n\017ListWorkflowIds\0220.flyteidl.ad" +
      "min.NamedEntityIdentifierListRequest\032).f" +
      "lyteidl.admin.NamedEntityIdentifierList\"" +
      "a\202\323\344\223\002[\022\'/api/v1/workflow_ids/{project}/" +
      "{domain}Z0\022./api/v1/workflows/org/{org}/" +
      "{project}/{domain}\022\300\002\n\rListWorkflows\022#.f" +
      "lyteidl.admin.ResourceListRequest\032\034.flyt" +
      "eidl.admin.WorkflowList\"\353\001\202\323\344\223\002\344\001\0224/api/" +
      "v1/workflows/{id.project}/{id.domain}/{i" +
      "d.name}ZC\022A/api/v1/workflows/org/{id.org" +
      "}/{id.project}/{id.domain}/{id.name}Z,\022*" +
      "/api/v1/workflows/{id.project}/{id.domai" +
      "n}Z9\0227/api/v1/workflows/org/{id.org}/{id" +
      ".project}/{id.domain}\022\256\001\n\020CreateLaunchPl" +
      "an\022\'.flyteidl.admin.LaunchPlanCreateRequ" +
      "est\032(.flyteidl.admin.LaunchPlanCreateRes" +
      "ponse\"G\202\323\344\223\002A\"\024/api/v1/launch_plans:\001*Z&" +
      "\"!/api/v1/launch_plans/org/{id.org}:\001*\022\362" +
      "\001\n\rGetLaunchPlan\022 .flyteidl.admin.Object" +
      "GetRequest\032\032.flyteidl.admin.LaunchPlan\"\242" +
      "\001\202\323\344\223\002\233\001\022D/api/v1/launch_plans/{id.proje" +
      "ct}/{id.domain}/{id.name}/{id.version}ZS" +
      "\022Q/api/v1/launch_plans/org/{id.org}/{id." +
      "project}/{id.domain}/{id.name}/{id.versi" +
      "on}\022\363\001\n\023GetActiveLaunchPlan\022\'.flyteidl.a" +
      "dmin.ActiveLaunchPlanRequest\032\032.flyteidl." +
      "admin.LaunchPlan\"\226\001\202\323\344\223\002\217\001\022>/api/v1/acti" +
      "ve_launch_plans/{id.project}/{id.domain}" +
      "/{id.name}ZM\022K/api/v1/active_launch_plan" +
      "s/org/{id.org}/{id.project}/{id.domain}/" +
      "{id.name}\022\330\001\n\025ListActiveLaunchPlans\022+.fl" +
      "yteidl.admin.ActiveLaunchPlanListRequest" +
      "\032\036.flyteidl.admin.LaunchPlanList\"r\202\323\344\223\002l" +
      "\022./api/v1/active_launch_plans/{project}/" +
      "{domain}Z:\0228/api/v1/org/{org}/active_lau" +
      "nch_plans/{project}/{domain}\022\334\001\n\021ListLau" +
      "nchPlanIds\0220.flyteidl.admin.NamedEntityI" +
      "dentifierListRequest\032).flyteidl.admin.Na" +
      "medEntityIdentifierList\"j\202\323\344\223\002d\022*/api/v1" +
      "/launch_plan_ids/{project}/{domain}Z6\0224/" +
      "api/v1/launch_plan_ids/org/{org}/{projec" +
      "t}/{domain}\022\320\002\n\017ListLaunchPlans\022#.flytei" +
      "dl.admin.ResourceListRequest\032\036.flyteidl." +
      "admin.LaunchPlanList\"\367\001\202\323\344\223\002\360\001\0227/api/v1/" +
      "launch_plans/{id.project}/{id.domain}/{i" +
      "d.name}ZF\022D/api/v1/launch_plans/org/{id." +
      "org}/{id.project}/{id.domain}/{id.name}Z" +
      "/\022-/api/v1/launch_plans/{id.project}/{id" +
      ".domain}Z<\022:/api/v1/launch_plans/org/{id" +
      ".org}/{id.project}/{id.domain}\022\215\002\n\020Updat" +
      "eLaunchPlan\022\'.flyteidl.admin.LaunchPlanU" +
      "pdateRequest\032(.flyteidl.admin.LaunchPlan" +
      "UpdateResponse\"\245\001\202\323\344\223\002\236\001\032D/api/v1/launch" +
      "_plans/{id.project}/{id.domain}/{id.name" +
      "}/{id.version}:\001*ZS\032Q/api/v1/launch_plan" +
      "s/org/{id.org}/{id.project}/{id.domain}/" +
      "{id.name}/{id.version}\022\244\001\n\017CreateExecuti" +
      "on\022&.flyteidl.admin.ExecutionCreateReque" +
      "st\032\'.flyteidl.admin.ExecutionCreateRespo" +
      "nse\"@\202\323\344\223\002:\"\022/api/v1/executions:\001*Z!\032\034/a" +
      "pi/v1/executions/org/{org}:\001*\022\275\001\n\021Relaun" +
      "chExecution\022(.flyteidl.admin.ExecutionRe" +
      "launchRequest\032\'.flyteidl.admin.Execution" +
      "CreateResponse\"U\202\323\344\223\002O\"\033/api/v1/executio" +
      "ns/relaunch:\001*Z-\"(/api/v1/executions/org" +
      "/{id.org}/relaunch:\001*\022\271\001\n\020RecoverExecuti" +
      "on\022\'.flyteidl.admin.ExecutionRecoverRequ" +
      "est\032\'.flyteidl.admin.ExecutionCreateResp" +
      "onse\"S\202\323\344\223\002M\"\032/api/v1/executions/recover" +
      ":\001*Z,\"\'/api/v1/executions/org/{id.org}/r" +
      "ecover:\001*\022\334\001\n\014GetExecution\022+.flyteidl.ad" +
      "min.WorkflowExecutionGetRequest\032\031.flytei" +
      "dl.admin.Execution\"\203\001\202\323\344\223\002}\0225/api/v1/exe" +
      "cutions/{id.project}/{id.domain}/{id.nam" +
      "e}ZD\022B/api/v1/executions/org/{id.org}/{i" +
      "d.project}/{id.domain}/{id.name}\022\357\001\n\017Upd" +
      "ateExecution\022&.flyteidl.admin.ExecutionU" +
      "pdateRequest\032\'.flyteidl.admin.ExecutionU" +
      "pdateResponse\"\212\001\202\323\344\223\002\203\001\0325/api/v1/executi" +
      "ons/{id.project}/{id.domain}/{id.name}:\001" +
      "*ZG\032B/api/v1/executions/org/{id.org}/{id" +
      ".project}/{id.domain}/{id.name}:\001*\022\206\002\n\020G" +
      "etExecutionData\022/.flyteidl.admin.Workflo" +
      "wExecutionGetDataRequest\0320.flyteidl.admi" +
      "n.WorkflowExecutionGetDataResponse\"\216\001\202\323\344" +
      "\223\002\207\001\022:/api/v1/data/executions/{id.projec" +
      "t}/{id.domain}/{id.name}ZI\022G/api/v1/data" +
      "/executions/org/{id.org}/{id.project}/{i" +
      "d.domain}/{id.name}\022\305\001\n\016ListExecutions\022#" +
      ".flyteidl.admin.ResourceListRequest\032\035.fl" +
      "yteidl.admin.ExecutionList\"o\202\323\344\223\002i\022+/api" +
      "/v1/executions/{id.project}/{id.domain}Z" +
      ":\0228/api/v1/executions/org/{id.org}/{id.p" +
      "roject}/{id.domain}\022\370\001\n\022TerminateExecuti" +
      "on\022).flyteidl.admin.ExecutionTerminateRe" +
      "quest\032*.flyteidl.admin.ExecutionTerminat" +
      "eResponse\"\212\001\202\323\344\223\002\203\001*5/api/v1/executions/" +
      "{id.project}/{id.domain}/{id.name}:\001*ZG*" +
      "B/api/v1/executions/org/{id.org}/{id.pro" +
      "ject}/{id.domain}/{id.name}:\001*\022\342\002\n\020GetNo" +
      "deExecution\022\'.flyteidl.admin.NodeExecuti" +
      "onGetRequest\032\035.flyteidl.admin.NodeExecut" +
      "ion\"\205\002\202\323\344\223\002\376\001\022n/api/v1/node_executions/{" +
      "id.execution_id.project}/{id.execution_i" +
      "d.domain}/{id.execution_id.name}/{id.nod" +
      "e_id}Z\213\001\022\210\001/api/v1/node_executions/org/{" +
      "id.execution_id.org}/{id.execution_id.pr" +
      "oject}/{id.execution_id.domain}/{id.exec" +
      "ution_id.name}/{id.node_id}\022\371\002\n\022ListNode" +
      "Executions\022(.flyteidl.admin.NodeExecutio" +
      "nListRequest\032!.flyteidl.admin.NodeExecut" +
      "ionList\"\225\002\202\323\344\223\002\216\002\022s/api/v1/node_executio" +
      "ns/{workflow_execution_id.project}/{work" +
      "flow_execution_id.domain}/{workflow_exec" +
      "ution_id.name}Z\226\001\022\223\001/api/v1/node_executi" +
      "ons/org/{workflow_execution_id.org}/{wor" +
      "kflow_execution_id.project}/{workflow_ex" +
      "ecution_id.domain}/{workflow_execution_i" +
      "d.name}\022\217\010\n\031ListNodeExecutionsForTask\022/." +
      "flyteidl.admin.NodeExecutionForTaskListR" +
      "equest\032!.flyteidl.admin.NodeExecutionLis" +
      "t\"\235\007\202\323\344\223\002\226\007\022\251\003/api/v1/children/task_exec" +
      "utions/{task_execution_id.node_execution" +
      "_id.execution_id.project}/{task_executio" +
      "n_id.node_execution_id.execution_id.doma" +
      "in}/{task_execution_id.node_execution_id" +
      ".execution_id.name}/{task_execution_id.n" +
      "ode_execution_id.node_id}/{task_executio" +
      "n_id.task_id.project}/{task_execution_id" +
      ".task_id.domain}/{task_execution_id.task" +
      "_id.name}/{task_execution_id.task_id.ver" +
      "sion}/{task_execution_id.retry_attempt}Z" +
      "\347\003\022\344\003/api/v1/children/task_executions/or" +
      "g/{task_execution_id.node_execution_id.e" +
      "xecution_id.org}/{task_execution_id.node" +
      "_execution_id.execution_id.project}/{tas" +
      "k_execution_id.node_execution_id.executi" +
      "on_id.domain}/{task_execution_id.node_ex" +
      "ecution_id.execution_id.name}/{task_exec" +
      "ution_id.node_execution_id.node_id}/{tas" +
      "k_execution_id.task_id.project}/{task_ex" +
      "ecution_id.task_id.domain}/{task_executi" +
      "on_id.task_id.name}/{task_execution_id.t" +
      "ask_id.version}/{task_execution_id.retry" +
      "_attempt}\022\203\003\n\024GetNodeExecutionData\022+.fly" +
      "teidl.admin.NodeExecutionGetDataRequest\032" +
      ",.flyteidl.admin.NodeExecutionGetDataRes" +
      "ponse\"\217\002\202\323\344\223\002\210\002\022s/api/v1/data/node_execu" +
      "tions/{id.execution_id.project}/{id.exec" +
      "ution_id.domain}/{id.execution_id.name}/" +
      "{id.node_id}Z\220\001\022\215\001/api/v1/data/node_exec" +
      "utions/org/{id.execution_id.org}/{id.exe" +
      "cution_id.project}/{id.execution_id.doma" +
      "in}/{id.execution_id.name}/{id.node_id}\022" +
      "\250\001\n\017RegisterProject\022&.flyteidl.admin.Pro" +
      "jectRegisterRequest\032\'.flyteidl.admin.Pro" +
      "jectRegisterResponse\"D\202\323\344\223\002>\"\020/api/v1/pr" +
      "ojects:\001*Z\'\"\"/api/v1/projects/org/{proje" +
      "ct.org}:\001*\022\227\001\n\rUpdateProject\022\027.flyteidl." +
      "admin.Project\032%.flyteidl.admin.ProjectUp" +
      "dateResponse\"F\202\323\344\223\002@\032\025/api/v1/projects/{" +
      "id}:\001*Z$\032\037/api/v1/projects/org/{org}/{id" +
      "}:\001*\022\204\001\n\014ListProjects\022\".flyteidl.admin.P" +
      "rojectListRequest\032\030.flyteidl.admin.Proje" +
      "cts\"6\202\323\344\223\0020\022\020/api/v1/projectsZ\034\022\032/api/v1" +
      "/projects/org/{org}\022\325\001\n\023CreateWorkflowEv" +
      "ent\022-.flyteidl.admin.WorkflowExecutionEv" +
      "entRequest\032..flyteidl.admin.WorkflowExec" +
      "utionEventResponse\"_\202\323\344\223\002Y\"\030/api/v1/even" +
      "ts/workflows:\001*Z:\"5/api/v1/events/workfl" +
      "ows/org/{event.execution_id.org}:\001*\022\304\001\n\017" +
      "CreateNodeEvent\022).flyteidl.admin.NodeExe" +
      "cutionEventRequest\032*.flyteidl.admin.Node" +
      "ExecutionEventResponse\"Z\202\323\344\223\002T\"\024/api/v1/" +
      "events/nodes:\001*Z9\"4/api/v1/events/nodes/" +
      "org/{event.id.execution_id.org}:\001*\022\332\001\n\017C" +
      "reateTaskEvent\022).flyteidl.admin.TaskExec" +
      "utionEventRequest\032*.flyteidl.admin.TaskE" +
      "xecutionEventResponse\"p\202\323\344\223\002j\"\024/api/v1/e" +
      "vents/tasks:\001*ZO\"J/api/v1/events/tasks/o" +
      "rg/{event.parent_node_execution_id.execu" +
      "tion_id.org}:\001*\022\313\005\n\020GetTaskExecution\022\'.f" +
      "lyteidl.admin.TaskExecutionGetRequest\032\035." +
      "flyteidl.admin.TaskExecution\"\356\004\202\323\344\223\002\347\004\022\231" +
      "\002/api/v1/task_executions/{id.node_execut" +
      "ion_id.execution_id.project}/{id.node_ex" +
      "ecution_id.execution_id.domain}/{id.node" +
      "_execution_id.execution_id.name}/{id.nod" +
      "e_execution_id.node_id}/{id.task_id.proj" +
      "ect}/{id.task_id.domain}/{id.task_id.nam" +
      "e}/{id.task_id.version}/{id.retry_attemp" +
      "t}Z\310\002\022\305\002/api/v1/task_executions/org/{id." +
      "node_execution_id.execution_id.org}/{id." +
      "node_execution_id.execution_id.project}/" +
      "{id.node_execution_id.execution_id.domai" +
      "n}/{id.node_execution_id.execution_id.na" +
      "me}/{id.node_execution_id.node_id}/{id.t" +
      "ask_id.project}/{id.task_id.domain}/{id." +
      "task_id.name}/{id.task_id.version}/{id.r" +
      "etry_attempt}\022\361\003\n\022ListTaskExecutions\022(.f" +
      "lyteidl.admin.TaskExecutionListRequest\032!" +
      ".flyteidl.admin.TaskExecutionList\"\215\003\202\323\344\223" +
      "\002\206\003\022\252\001/api/v1/task_executions/{node_exec" +
      "ution_id.execution_id.project}/{node_exe" +
      "cution_id.execution_id.domain}/{node_exe" +
      "cution_id.execution_id.name}/{node_execu" +
      "tion_id.node_id}Z\326\001\022\323\001/api/v1/task_execu" +
      "tions/org/{node_execution_id.execution_i" +
      "d.org}/{node_execution_id.execution_id.p" +
      "roject}/{node_execution_id.execution_id." +
      "domain}/{node_execution_id.execution_id." +
      "name}/{node_execution_id.node_id}\022\354\005\n\024Ge" +
      "tTaskExecutionData\022+.flyteidl.admin.Task" +
      "ExecutionGetDataRequest\032,.flyteidl.admin" +
      ".TaskExecutionGetDataResponse\"\370\004\202\323\344\223\002\361\004\022" +
      "\236\002/api/v1/data/task_executions/{id.node_" +
      "execution_id.execution_id.project}/{id.n" +
      "ode_execution_id.execution_id.domain}/{i" +
      "d.node_execution_id.execution_id.name}/{" +
      "id.node_execution_id.node_id}/{id.task_i" +
      "d.project}/{id.task_id.domain}/{id.task_" +
      "id.name}/{id.task_id.version}/{id.retry_" +
      "attempt}Z\315\002\022\312\002/api/v1/data/task_executio" +
      "ns/org/{id.node_execution_id.execution_i" +
      "d.org}/{id.node_execution_id.execution_i" +
      "d.project}/{id.node_execution_id.executi" +
      "on_id.domain}/{id.node_execution_id.exec" +
      "ution_id.name}/{id.node_execution_id.nod" +
      "e_id}/{id.task_id.project}/{id.task_id.d" +
      "omain}/{id.task_id.name}/{id.task_id.ver" +
      "sion}/{id.retry_attempt}\022\313\002\n\035UpdateProje" +
      "ctDomainAttributes\0224.flyteidl.admin.Proj" +
      "ectDomainAttributesUpdateRequest\0325.flyte" +
      "idl.admin.ProjectDomainAttributesUpdateR" +
      "esponse\"\274\001\202\323\344\223\002\265\001\032J/api/v1/project_domai" +
      "n_attributes/{attributes.project}/{attri" +
      "butes.domain}:\001*Zd\032_/api/v1/project_doma" +
      "in_attributes/org/{attributes.org}/{attr" +
      "ibutes.project}/{attributes.domain}:\001*\022\203" +
      "\002\n\032GetProjectDomainAttributes\0221.flyteidl" +
      ".admin.ProjectDomainAttributesGetRequest" +
      "\0322.flyteidl.admin.ProjectDomainAttribute" +
      "sGetResponse\"~\202\323\344\223\002x\0224/api/v1/project_do" +
      "main_attributes/{project}/{domain}Z@\022>/a" +
      "pi/v1/project_domain_attributes/org/{org" +
      "}/{project}/{domain}\022\223\002\n\035DeleteProjectDo" +
      "mainAttributes\0224.flyteidl.admin.ProjectD" +
      "omainAttributesDeleteRequest\0325.flyteidl." +
      "admin.ProjectDomainAttributesDeleteRespo" +
      "nse\"\204\001\202\323\344\223\002~*4/api/v1/project_domain_att" +
      "ributes/{project}/{domain}:\001*ZC*>/api/v1" +
      "/project_domain_attributes/org/{org}/{pr" +
      "oject}/{domain}:\001*\022\212\002\n\027UpdateProjectAttr" +
      "ibutes\022..flyteidl.admin.ProjectAttribute" +
      "sUpdateRequest\032/.flyteidl.admin.ProjectA" +
      "ttributesUpdateResponse\"\215\001\202\323\344\223\002\206\001\032//api/" +
      "v1/project_attributes/{attributes.projec" +
      "t}:\001*ZP\032K/api/v1/project_domain_attribut" +
      "es/org/{attributes.org}/{attributes.proj" +
      "ect}:\001*\022\330\001\n\024GetProjectAttributes\022+.flyte" +
      "idl.admin.ProjectAttributesGetRequest\032,." +
      "flyteidl.admin.ProjectAttributesGetRespo" +
      "nse\"e\202\323\344\223\002_\022$/api/v1/project_attributes/" +
      "{project}Z7\0225/api/v1/project_domain_attr" +
      "ibutes/org/{org}/{project}\022\347\001\n\027DeletePro" +
      "jectAttributes\022..flyteidl.admin.ProjectA" +
      "ttributesDeleteRequest\032/.flyteidl.admin." +
      "ProjectAttributesDeleteResponse\"k\202\323\344\223\002e*" +
      "$/api/v1/project_attributes/{project}:\001*" +
      "Z:*5/api/v1/project_domain_attributes/or" +
      "g/{org}/{project}:\001*\022\334\002\n\030UpdateWorkflowA" +
      "ttributes\022/.flyteidl.admin.WorkflowAttri" +
      "butesUpdateRequest\0320.flyteidl.admin.Work" +
      "flowAttributesUpdateResponse\"\334\001\202\323\344\223\002\325\001\032Z" +
      "/api/v1/workflow_attributes/{attributes." +
      "project}/{attributes.domain}/{attributes" +
      ".workflow}:\001*Zt\032o/api/v1/workflow_attrib" +
      "utes/org/{attributes.org}/{attributes.pr" +
      "oject}/{attributes.domain}/{attributes.w" +
      "orkflow}:\001*\022\200\002\n\025GetWorkflowAttributes\022,." +
      "flyteidl.admin.WorkflowAttributesGetRequ" +
      "est\032-.flyteidl.admin.WorkflowAttributesG" +
      "etResponse\"\211\001\202\323\344\223\002\202\001\0229/api/v1/workflow_a" +
      "ttributes/{project}/{domain}/{workflow}Z" +
      "E\022C/api/v1/workflow_attributes/org/{org}" +
      "/{project}/{domain}/{workflow}\022\217\002\n\030Delet" +
      "eWorkflowAttributes\022/.flyteidl.admin.Wor" +
      "kflowAttributesDeleteRequest\0320.flyteidl." +
      "admin.WorkflowAttributesDeleteResponse\"\217" +
      "\001\202\323\344\223\002\210\001*9/api/v1/workflow_attributes/{p" +
      "roject}/{domain}/{workflow}:\001*ZH*C/api/v" +
      "1/workflow_attributes/org/{org}/{project" +
      "}/{domain}/{workflow}:\001*\022\312\001\n\027ListMatchab" +
      "leAttributes\022..flyteidl.admin.ListMatcha" +
      "bleAttributesRequest\032/.flyteidl.admin.Li" +
      "stMatchableAttributesResponse\"N\202\323\344\223\002H\022\034/" +
      "api/v1/matchable_attributesZ(\022&/api/v1/m" +
      "atchable_attributes/org/{org}\022\350\001\n\021ListNa" +
      "medEntities\022&.flyteidl.admin.NamedEntity" +
      "ListRequest\032\037.flyteidl.admin.NamedEntity" +
      "List\"\211\001\202\323\344\223\002\202\001\0229/api/v1/named_entities/{" +
      "resource_type}/{project}/{domain}ZE\022C/ap" +
      "i/v1/named_entities/{resource_type}/org/" +
      "{org}/{project}/{domain}\022\203\002\n\016GetNamedEnt" +
      "ity\022%.flyteidl.admin.NamedEntityGetReque" +
      "st\032\033.flyteidl.admin.NamedEntity\"\254\001\202\323\344\223\002\245" +
      "\001\022I/api/v1/named_entities/{resource_type" +
      "}/{id.project}/{id.domain}/{id.name}ZX\022V" +
      "/api/v1/named_entities/{resource_type}/o" +
      "rg/{id.org}/{id.project}/{id.domain}/{id" +
      ".name}\022\235\002\n\021UpdateNamedEntity\022(.flyteidl." +
      "admin.NamedEntityUpdateRequest\032).flyteid" +
      "l.admin.NamedEntityUpdateResponse\"\262\001\202\323\344\223" +
      "\002\253\001\032I/api/v1/named_entities/{resource_ty" +
      "pe}/{id.project}/{id.domain}/{id.name}:\001" +
      "*Z[\032V/api/v1/named_entities/{resource_ty" +
      "pe}/org/{id.org}/{id.project}/{id.domain" +
      "}/{id.name}:\001*\022l\n\nGetVersion\022!.flyteidl." +
      "admin.GetVersionRequest\032\".flyteidl.admin" +
      ".GetVersionResponse\"\027\202\323\344\223\002\021\022\017/api/v1/ver" +
      "sion\022\266\002\n\024GetDescriptionEntity\022 .flyteidl" +
      ".admin.ObjectGetRequest\032!.flyteidl.admin" +
      ".DescriptionEntity\"\330\001\202\323\344\223\002\321\001\022_/api/v1/de" +
      "scription_entities/{id.resource_type}/{i" +
      "d.project}/{id.domain}/{id.name}/{id.ver" +
      "sion}Zn\022l/api/v1/description_entities/or" +
      "g/{id.org}/{id.resource_type}/{id.projec" +
      "t}/{id.domain}/{id.name}/{id.version}\022\310\003" +
      "\n\027ListDescriptionEntities\022,.flyteidl.adm" +
      "in.DescriptionEntityListRequest\032%.flytei" +
      "dl.admin.DescriptionEntityList\"\327\002\202\323\344\223\002\320\002" +
      "\022O/api/v1/description_entities/{resource" +
      "_type}/{id.project}/{id.domain}/{id.name" +
      "}Z^\022\\/api/v1/description_entities/{resou" +
      "rce_type}/org/{id.org}/{id.project}/{id." +
      "domain}/{id.name}ZG\022E/api/v1/description" +
      "_entities/{resource_type}/{id.project}/{" +
      "id.domain}ZT\022R/api/v1/description_entiti" +
      "es/{resource_type}/org/{id.org}/{id.proj" +
      "ect}/{id.domain}\022\225\002\n\023GetExecutionMetrics" +
      "\0222.flyteidl.admin.WorkflowExecutionGetMe" +
      "tricsRequest\0323.flyteidl.admin.WorkflowEx" +
      "ecutionGetMetricsResponse\"\224\001\202\323\344\223\002\215\001\022=/ap" +
      "i/v1/metrics/executions/{id.project}/{id" +
      ".domain}/{id.name}ZL\022J/api/v1/metrics/ex" +
      "ecutions/org/{id.org}/{id.project}/{id.d" +
      "omain}/{id.name}B?Z=github.com/flyteorg/" +
      "flyte/flyteidl/gen/pb-go/flyteidl/servic" +
      "eb\006proto3"
    };
    com.google.protobuf.Descriptors.FileDescriptor.InternalDescriptorAssigner assigner =
        new com.google.protobuf.Descriptors.FileDescriptor.    InternalDescriptorAssigner() {
          public com.google.protobuf.ExtensionRegistry assignDescriptors(
              com.google.protobuf.Descriptors.FileDescriptor root) {
            descriptor = root;
            return null;
          }
        };
    com.google.protobuf.Descriptors.FileDescriptor
      .internalBuildGeneratedFileFrom(descriptorData,
        new com.google.protobuf.Descriptors.FileDescriptor[] {
          com.google.api.AnnotationsProto.getDescriptor(),
          flyteidl.admin.ProjectOuterClass.getDescriptor(),
          flyteidl.admin.ProjectDomainAttributesOuterClass.getDescriptor(),
          flyteidl.admin.ProjectAttributesOuterClass.getDescriptor(),
          flyteidl.admin.TaskOuterClass.getDescriptor(),
          flyteidl.admin.WorkflowOuterClass.getDescriptor(),
          flyteidl.admin.WorkflowAttributesOuterClass.getDescriptor(),
          flyteidl.admin.LaunchPlanOuterClass.getDescriptor(),
          flyteidl.admin.Event.getDescriptor(),
          flyteidl.admin.ExecutionOuterClass.getDescriptor(),
          flyteidl.admin.MatchableResourceOuterClass.getDescriptor(),
          flyteidl.admin.NodeExecutionOuterClass.getDescriptor(),
          flyteidl.admin.TaskExecutionOuterClass.getDescriptor(),
          flyteidl.admin.VersionOuterClass.getDescriptor(),
          flyteidl.admin.Common.getDescriptor(),
          flyteidl.admin.DescriptionEntityOuterClass.getDescriptor(),
        }, assigner);
    com.google.protobuf.ExtensionRegistry registry =
        com.google.protobuf.ExtensionRegistry.newInstance();
    registry.add(com.google.api.AnnotationsProto.http);
    com.google.protobuf.Descriptors.FileDescriptor
        .internalUpdateFileDescriptor(descriptor, registry);
    com.google.api.AnnotationsProto.getDescriptor();
    flyteidl.admin.ProjectOuterClass.getDescriptor();
    flyteidl.admin.ProjectDomainAttributesOuterClass.getDescriptor();
    flyteidl.admin.ProjectAttributesOuterClass.getDescriptor();
    flyteidl.admin.TaskOuterClass.getDescriptor();
    flyteidl.admin.WorkflowOuterClass.getDescriptor();
    flyteidl.admin.WorkflowAttributesOuterClass.getDescriptor();
    flyteidl.admin.LaunchPlanOuterClass.getDescriptor();
    flyteidl.admin.Event.getDescriptor();
    flyteidl.admin.ExecutionOuterClass.getDescriptor();
    flyteidl.admin.MatchableResourceOuterClass.getDescriptor();
    flyteidl.admin.NodeExecutionOuterClass.getDescriptor();
    flyteidl.admin.TaskExecutionOuterClass.getDescriptor();
    flyteidl.admin.VersionOuterClass.getDescriptor();
    flyteidl.admin.Common.getDescriptor();
    flyteidl.admin.DescriptionEntityOuterClass.getDescriptor();
  }

  // @@protoc_insertion_point(outer_class_scope)
}
