.. _deployment-cluster-config:

##############
Cluster Config
##############


.. panels::
    :header: text-center
    :column: col-lg-12 p-2

    .. link-button:: deployment-cluster-config-auth-setup
       :type: ref
       :text: Authenticating in Flyte
       :classes: btn-block stretched-link
    ^^^^^^^^^^^^
    Basic OIDC and Authentication Setup

    ---

    .. link-button:: deployment-cluster-config-auth-migration
       :type: ref
       :text: Migrating Your Authentication Config
       :classes: btn-block stretched-link
    ^^^^^^^^^^^^
    Migration guide to move to Admin's own authorization server.

    ---

    .. link-button:: deployment-cluster-config-auth-appendix
       :type: ref
       :text: Understanding Authentication in Detail
       :classes: btn-block stretched-link
    ^^^^^^^^^^^^
    Migration guide to move to Admin's own authorization server.

    ---

    .. link-button:: deployment-cluster-config-general
       :type: ref
       :text: Configuring Custom K8s Resources
       :classes: btn-block stretched-link
    ^^^^^^^^^^^^
    How to use Flyte's cluster-resource-controller to control specific Kubernetes resources and administer project/domain-specific resource quotas (say, to limit the number of CPUs/GPUs/mem per tenant).

    ---

    .. link-button:: deployment-customizable-resources
       :type: ref
       :text: Adding New Customizable Resources
       :classes: btn-block stretched-link
    ^^^^^^^^^^^^
    Creating new default configurations or overriding certain values for specific combinations of user projects, domains and workflows through Flyte APIs.

    ---

    .. link-button:: flyteadmin-config-specification
       :type: ref
       :text: FlyteAdmin Configuration
       :classes: btn-block stretched-link
    ^^^^^^^^^^^^
    All about FlyteAdmin configuration specification.

    ---

    .. link-button:: flytepropeller-config-specification
       :type: ref
       :text: FlytePropeller Configuration
       :classes: btn-block stretched-link
    ^^^^^^^^^^^^
    Detailed list of all Flyte Propeller configuration options and their default values.

    ---

    .. link-button:: flytedatacatalog-config-specification
       :type: ref
       :text: Flyte Datacatalog Configuration
       :classes: btn-block stretched-link
    ^^^^^^^^^^^^
    Flyte Datacatalog configuration specification are covered here.

    ---

    .. link-button:: flytescheduler-config-specification
       :type: ref
       :text: Flyte Scheduler Configuration
       :classes: btn-block stretched-link
    ^^^^^^^^^^^^
    Flyte Scheduler configuration specification are covered here.

    ---

    .. link-button:: deployment-cluster-config-notifications
       :type: ref
       :text: Notifications
       :classes: btn-block stretched-link
    ^^^^^^^^^^^^
    Guide to setting up and configuring notifications.

    ---

    .. link-button:: deployment-cluster-config-eventing
       :type: ref
       :text: External Events
       :classes: btn-block stretched-link
    ^^^^^^^^^^^^
    How to set up Flyte to emit events to third-parties.

    ---

    .. link-button:: deployment-cluster-config-cloud-event
       :type: ref
       :text: Cloud Events
       :classes: btn-block stretched-link
    ^^^^^^^^^^^^
    How to set up Flyte to emit events to third-parties in CloudEvents format.

    ---

    .. link-button:: deployment-cluster-config-monitoring
       :type: ref
       :text: Monitoring
       :classes: btn-block stretched-link
    ^^^^^^^^^^^^
    Guide to setting up and configuring observability.

    ---

    .. link-button:: deployment-cluster-config-performance
       :type: ref
       :text: Optimizing Performance
       :classes: btn-block stretched-link
    ^^^^^^^^^^^^
    Improve the performance of the core Flyte engine.

.. toctree::
    :maxdepth: 1
    :name: Cluster Config
    :hidden:

    auth_setup
    auth_migration
    auth_appendix
    general
    customizable_resources
    flyteadmin_config
    flytepropeller_config
    datacatalog_config
    scheduler_config
    monitoring
    notifications
    performance
    eventing
    cloud_event
