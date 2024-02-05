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
      ".proto2\277z\n\014AdminService\022\216\001\n\nCreateTask\022!" +
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
      "{domain}Z:\0228/api/v1/active_launch_plans/" +
      "org/{org}/{project}/{domain}\022\334\001\n\021ListLau" +
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
      "/org/{id.org}/executions/{id.project}/{i" +
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
      "ution_id.name}/{id.node_id}\022\236\003\n\026GetDynam" +
      "icNodeWorkflow\022-.flyteidl.admin.GetDynam" +
      "icNodeWorkflowRequest\032+.flyteidl.admin.D" +
      "ynamicNodeWorkflowResponse\"\247\002\202\323\344\223\002\240\002\022\177/a" +
      "pi/v1/node_executions/{id.execution_id.p" +
      "roject}/{id.execution_id.domain}/{id.exe" +
      "cution_id.name}/{id.node_id}/dynamic_wor" +
      "kflowZ\234\001\022\231\001/api/v1/node_executions/org/{" +
      "id.execution_id.org}/{id.execution_id.pr" +
      "oject}/{id.execution_id.domain}/{id.exec" +
      "ution_id.name}/{id.node_id}/dynamic_work" +
      "flow\022\371\002\n\022ListNodeExecutions\022(.flyteidl.a" +
      "dmin.NodeExecutionListRequest\032!.flyteidl" +
      ".admin.NodeExecutionList\"\225\002\202\323\344\223\002\216\002\022s/api" +
      "/v1/node_executions/{workflow_execution_" +
      "id.project}/{workflow_execution_id.domai" +
      "n}/{workflow_execution_id.name}Z\226\001\022\223\001/ap" +
      "i/v1/node_executions/org/{workflow_execu" +
      "tion_id.org}/{workflow_execution_id.proj" +
      "ect}/{workflow_execution_id.domain}/{wor" +
      "kflow_execution_id.name}\022\217\010\n\031ListNodeExe" +
      "cutionsForTask\022/.flyteidl.admin.NodeExec" +
      "utionForTaskListRequest\032!.flyteidl.admin" +
      ".NodeExecutionList\"\235\007\202\323\344\223\002\226\007\022\251\003/api/v1/c" +
      "hildren/task_executions/{task_execution_" +
      "id.node_execution_id.execution_id.projec" +
      "t}/{task_execution_id.node_execution_id." +
      "execution_id.domain}/{task_execution_id." +
      "node_execution_id.execution_id.name}/{ta" +
      "sk_execution_id.node_execution_id.node_i" +
      "d}/{task_execution_id.task_id.project}/{" +
      "task_execution_id.task_id.domain}/{task_" +
      "execution_id.task_id.name}/{task_executi" +
      "on_id.task_id.version}/{task_execution_i" +
      "d.retry_attempt}Z\347\003\022\344\003/api/v1/children/o" +
      "rg/{task_execution_id.node_execution_id." +
      "execution_id.org}/task_executions/{task_" +
      "execution_id.node_execution_id.execution" +
      "_id.project}/{task_execution_id.node_exe" +
      "cution_id.execution_id.domain}/{task_exe" +
      "cution_id.node_execution_id.execution_id" +
      ".name}/{task_execution_id.node_execution" +
      "_id.node_id}/{task_execution_id.task_id." +
      "project}/{task_execution_id.task_id.doma" +
      "in}/{task_execution_id.task_id.name}/{ta" +
      "sk_execution_id.task_id.version}/{task_e" +
      "xecution_id.retry_attempt}\022\203\003\n\024GetNodeEx" +
      "ecutionData\022+.flyteidl.admin.NodeExecuti" +
      "onGetDataRequest\032,.flyteidl.admin.NodeEx" +
      "ecutionGetDataResponse\"\217\002\202\323\344\223\002\210\002\022s/api/v" +
      "1/data/node_executions/{id.execution_id." +
      "project}/{id.execution_id.domain}/{id.ex" +
      "ecution_id.name}/{id.node_id}Z\220\001\022\215\001/api/" +
      "v1/data/org/{id.execution_id.org}/node_e" +
      "xecutions/{id.execution_id.project}/{id." +
      "execution_id.domain}/{id.execution_id.na" +
      "me}/{id.node_id}\022\250\001\n\017RegisterProject\022&.f" +
      "lyteidl.admin.ProjectRegisterRequest\032\'.f" +
      "lyteidl.admin.ProjectRegisterResponse\"D\202" +
      "\323\344\223\002>\"\020/api/v1/projects:\001*Z\'\"\"/api/v1/pr" +
      "ojects/org/{project.org}:\001*\022\227\001\n\rUpdatePr" +
      "oject\022\027.flyteidl.admin.Project\032%.flyteid" +
      "l.admin.ProjectUpdateResponse\"F\202\323\344\223\002@\032\025/" +
      "api/v1/projects/{id}:\001*Z$\032\037/api/v1/proje" +
      "cts/org/{org}/{id}:\001*\022\204\001\n\nGetProject\022\036.f" +
      "lyteidl.admin.ProjectRequest\032\027.flyteidl." +
      "admin.Project\"=\202\323\344\223\0027\022\020/api/v1/projectsZ" +
      "#\022!/api/v1/projects/org/{project.id}\022\204\001\n" +
      "\014ListProjects\022\".flyteidl.admin.ProjectLi" +
      "stRequest\032\030.flyteidl.admin.Projects\"6\202\323\344" +
      "\223\0020\022\020/api/v1/projectsZ\034\022\032/api/v1/project" +
      "s/org/{org}\022\325\001\n\023CreateWorkflowEvent\022-.fl" +
      "yteidl.admin.WorkflowExecutionEventReque" +
      "st\032..flyteidl.admin.WorkflowExecutionEve" +
      "ntResponse\"_\202\323\344\223\002Y\"\030/api/v1/events/workf" +
      "lows:\001*Z:\"5/api/v1/events/org/{event.exe" +
      "cution_id.org}/workflows:\001*\022\304\001\n\017CreateNo" +
      "deEvent\022).flyteidl.admin.NodeExecutionEv" +
      "entRequest\032*.flyteidl.admin.NodeExecutio" +
      "nEventResponse\"Z\202\323\344\223\002T\"\024/api/v1/events/n" +
      "odes:\001*Z9\"4/api/v1/events/org/{event.id." +
      "execution_id.org}/nodes:\001*\022\332\001\n\017CreateTas" +
      "kEvent\022).flyteidl.admin.TaskExecutionEve" +
      "ntRequest\032*.flyteidl.admin.TaskExecution" +
      "EventResponse\"p\202\323\344\223\002j\"\024/api/v1/events/ta" +
      "sks:\001*ZO\"J/api/v1/events/org/{event.pare" +
      "nt_node_execution_id.execution_id.org}/t" +
      "asks:\001*\022\313\005\n\020GetTaskExecution\022\'.flyteidl." +
      "admin.TaskExecutionGetRequest\032\035.flyteidl" +
      ".admin.TaskExecution\"\356\004\202\323\344\223\002\347\004\022\231\002/api/v1" +
      "/task_executions/{id.node_execution_id.e" +
      "xecution_id.project}/{id.node_execution_" +
      "id.execution_id.domain}/{id.node_executi" +
      "on_id.execution_id.name}/{id.node_execut" +
      "ion_id.node_id}/{id.task_id.project}/{id" +
      ".task_id.domain}/{id.task_id.name}/{id.t" +
      "ask_id.version}/{id.retry_attempt}Z\310\002\022\305\002" +
      "/api/v1/task_executions/org/{id.node_exe" +
      "cution_id.execution_id.org}/{id.node_exe" +
      "cution_id.execution_id.project}/{id.node" +
      "_execution_id.execution_id.domain}/{id.n" +
      "ode_execution_id.execution_id.name}/{id." +
      "node_execution_id.node_id}/{id.task_id.p" +
      "roject}/{id.task_id.domain}/{id.task_id." +
      "name}/{id.task_id.version}/{id.retry_att" +
      "empt}\022\361\003\n\022ListTaskExecutions\022(.flyteidl." +
      "admin.TaskExecutionListRequest\032!.flyteid" +
      "l.admin.TaskExecutionList\"\215\003\202\323\344\223\002\206\003\022\252\001/a" +
      "pi/v1/task_executions/{node_execution_id" +
      ".execution_id.project}/{node_execution_i" +
      "d.execution_id.domain}/{node_execution_i" +
      "d.execution_id.name}/{node_execution_id." +
      "node_id}Z\326\001\022\323\001/api/v1/task_executions/or" +
      "g/{node_execution_id.execution_id.org}/{" +
      "node_execution_id.execution_id.project}/" +
      "{node_execution_id.execution_id.domain}/" +
      "{node_execution_id.execution_id.name}/{n" +
      "ode_execution_id.node_id}\022\354\005\n\024GetTaskExe" +
      "cutionData\022+.flyteidl.admin.TaskExecutio" +
      "nGetDataRequest\032,.flyteidl.admin.TaskExe" +
      "cutionGetDataResponse\"\370\004\202\323\344\223\002\361\004\022\236\002/api/v" +
      "1/data/task_executions/{id.node_executio" +
      "n_id.execution_id.project}/{id.node_exec" +
      "ution_id.execution_id.domain}/{id.node_e" +
      "xecution_id.execution_id.name}/{id.node_" +
      "execution_id.node_id}/{id.task_id.projec" +
      "t}/{id.task_id.domain}/{id.task_id.name}" +
      "/{id.task_id.version}/{id.retry_attempt}" +
      "Z\315\002\022\312\002/api/v1/data/org/{id.node_executio" +
      "n_id.execution_id.org}/task_executions/{" +
      "id.node_execution_id.execution_id.projec" +
      "t}/{id.node_execution_id.execution_id.do" +
      "main}/{id.node_execution_id.execution_id" +
      ".name}/{id.node_execution_id.node_id}/{i" +
      "d.task_id.project}/{id.task_id.domain}/{" +
      "id.task_id.name}/{id.task_id.version}/{i" +
      "d.retry_attempt}\022\313\002\n\035UpdateProjectDomain" +
      "Attributes\0224.flyteidl.admin.ProjectDomai" +
      "nAttributesUpdateRequest\0325.flyteidl.admi" +
      "n.ProjectDomainAttributesUpdateResponse\"" +
      "\274\001\202\323\344\223\002\265\001\032J/api/v1/project_domain_attrib" +
      "utes/{attributes.project}/{attributes.do" +
      "main}:\001*Zd\032_/api/v1/project_domain_attri" +
      "butes/org/{attributes.org}/{attributes.p" +
      "roject}/{attributes.domain}:\001*\022\203\002\n\032GetPr" +
      "ojectDomainAttributes\0221.flyteidl.admin.P" +
      "rojectDomainAttributesGetRequest\0322.flyte" +
      "idl.admin.ProjectDomainAttributesGetResp" +
      "onse\"~\202\323\344\223\002x\0224/api/v1/project_domain_att" +
      "ributes/{project}/{domain}Z@\022>/api/v1/pr" +
      "oject_domain_attributes/org/{org}/{proje" +
      "ct}/{domain}\022\223\002\n\035DeleteProjectDomainAttr" +
      "ibutes\0224.flyteidl.admin.ProjectDomainAtt" +
      "ributesDeleteRequest\0325.flyteidl.admin.Pr" +
      "ojectDomainAttributesDeleteResponse\"\204\001\202\323" +
      "\344\223\002~*4/api/v1/project_domain_attributes/" +
      "{project}/{domain}:\001*ZC*>/api/v1/project" +
      "_domain_attributes/org/{org}/{project}/{" +
      "domain}:\001*\022\212\002\n\027UpdateProjectAttributes\022." +
      ".flyteidl.admin.ProjectAttributesUpdateR" +
      "equest\032/.flyteidl.admin.ProjectAttribute" +
      "sUpdateResponse\"\215\001\202\323\344\223\002\206\001\032//api/v1/proje" +
      "ct_attributes/{attributes.project}:\001*ZP\032" +
      "K/api/v1/project_domain_attributes/org/{" +
      "attributes.org}/{attributes.project}:\001*\022" +
      "\330\001\n\024GetProjectAttributes\022+.flyteidl.admi" +
      "n.ProjectAttributesGetRequest\032,.flyteidl" +
      ".admin.ProjectAttributesGetResponse\"e\202\323\344" +
      "\223\002_\022$/api/v1/project_attributes/{project" +
      "}Z7\0225/api/v1/project_domain_attributes/o" +
      "rg/{org}/{project}\022\347\001\n\027DeleteProjectAttr" +
      "ibutes\022..flyteidl.admin.ProjectAttribute" +
      "sDeleteRequest\032/.flyteidl.admin.ProjectA" +
      "ttributesDeleteResponse\"k\202\323\344\223\002e*$/api/v1" +
      "/project_attributes/{project}:\001*Z:*5/api" +
      "/v1/project_domain_attributes/org/{org}/" +
      "{project}:\001*\022\334\002\n\030UpdateWorkflowAttribute" +
      "s\022/.flyteidl.admin.WorkflowAttributesUpd" +
      "ateRequest\0320.flyteidl.admin.WorkflowAttr" +
      "ibutesUpdateResponse\"\334\001\202\323\344\223\002\325\001\032Z/api/v1/" +
      "workflow_attributes/{attributes.project}" +
      "/{attributes.domain}/{attributes.workflo" +
      "w}:\001*Zt\032o/api/v1/workflow_attributes/org" +
      "/{attributes.org}/{attributes.project}/{" +
      "attributes.domain}/{attributes.workflow}" +
      ":\001*\022\200\002\n\025GetWorkflowAttributes\022,.flyteidl" +
      ".admin.WorkflowAttributesGetRequest\032-.fl" +
      "yteidl.admin.WorkflowAttributesGetRespon" +
      "se\"\211\001\202\323\344\223\002\202\001\0229/api/v1/workflow_attribute" +
      "s/{project}/{domain}/{workflow}ZE\022C/api/" +
      "v1/workflow_attributes/org/{org}/{projec" +
      "t}/{domain}/{workflow}\022\217\002\n\030DeleteWorkflo" +
      "wAttributes\022/.flyteidl.admin.WorkflowAtt" +
      "ributesDeleteRequest\0320.flyteidl.admin.Wo" +
      "rkflowAttributesDeleteResponse\"\217\001\202\323\344\223\002\210\001" +
      "*9/api/v1/workflow_attributes/{project}/" +
      "{domain}/{workflow}:\001*ZH*C/api/v1/workfl" +
      "ow_attributes/org/{org}/{project}/{domai" +
      "n}/{workflow}:\001*\022\312\001\n\027ListMatchableAttrib" +
      "utes\022..flyteidl.admin.ListMatchableAttri" +
      "butesRequest\032/.flyteidl.admin.ListMatcha" +
      "bleAttributesResponse\"N\202\323\344\223\002H\022\034/api/v1/m" +
      "atchable_attributesZ(\022&/api/v1/matchable" +
      "_attributes/org/{org}\022\350\001\n\021ListNamedEntit" +
      "ies\022&.flyteidl.admin.NamedEntityListRequ" +
      "est\032\037.flyteidl.admin.NamedEntityList\"\211\001\202" +
      "\323\344\223\002\202\001\0229/api/v1/named_entities/{resource" +
      "_type}/{project}/{domain}ZE\022C/api/v1/nam" +
      "ed_entities/org/{org}/{resource_type}/{p" +
      "roject}/{domain}\022\203\002\n\016GetNamedEntity\022%.fl" +
      "yteidl.admin.NamedEntityGetRequest\032\033.fly" +
      "teidl.admin.NamedEntity\"\254\001\202\323\344\223\002\245\001\022I/api/" +
      "v1/named_entities/{resource_type}/{id.pr" +
      "oject}/{id.domain}/{id.name}ZX\022V/api/v1/" +
      "named_entities/org/{id.org}/{resource_ty" +
      "pe}/{id.project}/{id.domain}/{id.name}\022\235" +
      "\002\n\021UpdateNamedEntity\022(.flyteidl.admin.Na" +
      "medEntityUpdateRequest\032).flyteidl.admin." +
      "NamedEntityUpdateResponse\"\262\001\202\323\344\223\002\253\001\032I/ap" +
      "i/v1/named_entities/{resource_type}/{id." +
      "project}/{id.domain}/{id.name}:\001*Z[\032V/ap" +
      "i/v1/named_entities/org/{id.org}/{resour" +
      "ce_type}/{id.project}/{id.domain}/{id.na" +
      "me}:\001*\022l\n\nGetVersion\022!.flyteidl.admin.Ge" +
      "tVersionRequest\032\".flyteidl.admin.GetVers" +
      "ionResponse\"\027\202\323\344\223\002\021\022\017/api/v1/version\022\266\002\n" +
      "\024GetDescriptionEntity\022 .flyteidl.admin.O" +
      "bjectGetRequest\032!.flyteidl.admin.Descrip" +
      "tionEntity\"\330\001\202\323\344\223\002\321\001\022_/api/v1/descriptio" +
      "n_entities/{id.resource_type}/{id.projec" +
      "t}/{id.domain}/{id.name}/{id.version}Zn\022" +
      "l/api/v1/description_entities/org/{id.or" +
      "g}/{id.resource_type}/{id.project}/{id.d" +
      "omain}/{id.name}/{id.version}\022\310\003\n\027ListDe" +
      "scriptionEntities\022,.flyteidl.admin.Descr" +
      "iptionEntityListRequest\032%.flyteidl.admin" +
      ".DescriptionEntityList\"\327\002\202\323\344\223\002\320\002\022O/api/v" +
      "1/description_entities/{resource_type}/{" +
      "id.project}/{id.domain}/{id.name}Z^\022\\/ap" +
      "i/v1/description_entities/org/{id.org}/{" +
      "resource_type}/{id.project}/{id.domain}/" +
      "{id.name}ZG\022E/api/v1/description_entitie" +
      "s/{resource_type}/{id.project}/{id.domai" +
      "n}ZT\022R/api/v1/description_entities/org/{" +
      "id.org}/{resource_type}/{id.project}/{id",
      ".domain}\022\225\002\n\023GetExecutionMetrics\0222.flyte" +
      "idl.admin.WorkflowExecutionGetMetricsReq" +
      "uest\0323.flyteidl.admin.WorkflowExecutionG" +
      "etMetricsResponse\"\224\001\202\323\344\223\002\215\001\022=/api/v1/met" +
      "rics/executions/{id.project}/{id.domain}" +
      "/{id.name}ZL\022J/api/v1/metrics/executions" +
      "/org/{id.org}/{id.project}/{id.domain}/{" +
      "id.name}B?Z=github.com/flyteorg/flyte/fl" +
      "yteidl/gen/pb-go/flyteidl/serviceb\006proto" +
      "3"
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
