import typing
from grafanalib.core import (
    Alert, AlertCondition, Dashboard, Graph,
    GreaterThan, OP_AND, OPS_FORMAT, Row, RTYPE_SUM, SECONDS_FORMAT,
    SHORT_FORMAT, single_y_axis, Target, TimeRange, YAxes, YAxis, MILLISECONDS_FORMAT, DataSourceInput
)

# ------------------------------
# For Gostats we recommend using
# Grafana dashboard ID 10826 - https://grafana.com/grafana/dashboards/10826
#

DATASOURCE_NAME = "DS_PROM"
DATASOURCE = "${%s}" % DATASOURCE_NAME


class FlytePropeller(object):

    @staticmethod
    def create_free_workers() -> Graph:
        return Graph(
            title="Free workers count",
            dataSource=DATASOURCE,
            targets=[
                Target(
                    expr='sum(flyte:propeller:all:free_workers_count) by (kubernetes_pod_name)',
                    refId='A',
                ),
            ],
            yAxes=YAxes(
                YAxis(format=OPS_FORMAT),
                YAxis(format=SHORT_FORMAT),
            ),
        )

    @staticmethod
    def round_latency_per_wf(interval: int = 1) -> Graph:
        return Graph(
            title=f"round Latency per workflow",
            dataSource=DATASOURCE,
            targets=[
                Target(
                    expr=f'sum(rate(flyte:propeller:all:round:raw_ms[{interval}m])) by (wf)',
                    refId='A',
                ),
            ],
            yAxes=single_y_axis(format=MILLISECONDS_FORMAT),
        )

    @staticmethod
    def round_latency(interval: int = 1) -> Graph:
        return Graph(
            title=f"round Latency by quantile",
            dataSource=DATASOURCE,
            targets=[
                Target(
                    expr=f'sum(rate(flyte:propeller:all:round:raw_unlabeled_ms[{interval}m])) by (quantile)',
                    refId='A',
                ),
            ],
            yAxes=single_y_axis(format=MILLISECONDS_FORMAT),
        )

    @staticmethod
    def round_panic() -> Graph:
        return Graph(
            title="Round panic",
            dataSource=DATASOURCE,
            targets=[
                Target(
                    expr='sum(rate(flyte:propeller:all:round:panic_unlabeled[5m]))',
                    refId='A',
                ),
            ],
            yAxes=YAxes(
                YAxis(format=OPS_FORMAT),
                YAxis(format=SHORT_FORMAT),
            ),
        )

    @staticmethod
    def streak_length() -> Graph:
        return Graph(
            title="Avg streak length",
            dataSource=DATASOURCE,
            targets=[
                Target(
                    expr='avg(flyte:propeller:all:round:streak_length_unlabeled)',
                    refId='A',
                ),
            ],
            yAxes=YAxes(
                YAxis(format=OPS_FORMAT),
                YAxis(format=SHORT_FORMAT),
            ),
        )

    @staticmethod
    def system_errors() -> Graph:
        return Graph(
            title="System errors",
            dataSource=DATASOURCE,
            targets=[
                Target(
                    expr='sum(deriv(flyte:propeller:all:round:system_error_unlabeled[5m]))*300',
                    refId='A',
                ),
            ],
            yAxes=YAxes(
                YAxis(format=OPS_FORMAT),
                YAxis(format=SHORT_FORMAT),
            ),
        )

    @staticmethod
    def abort_errors() -> Graph:
        return Graph(
            title="System errors",
            dataSource=DATASOURCE,
            targets=[
                Target(
                    expr='sum(rate(flyte:propeller:all:round:abort_error[5m]))*300',
                    refId='A',
                ),
            ],
            yAxes=YAxes(
                YAxis(format=OPS_FORMAT),
                YAxis(format=SHORT_FORMAT),
            ),
        )

    @staticmethod
    def workflows_per_project() -> Graph:
        return Graph(
            title=f"Running Workflows per project",
            dataSource=DATASOURCE,
            targets=[
                Target(
                    expr=f'sum(flyte:propeller:all:collector:flyteworkflow) by (project)',
                    refId='A',
                ),
            ],
            yAxes=YAxes(
                YAxis(format=OPS_FORMAT),
                YAxis(format=SHORT_FORMAT),
            ),
        )

    @staticmethod
    def plugin_success_vs_failures() -> Graph:
        """
        TODO We need to convert the plugin names to be labels, so that prometheus can perform queries correctly
        """
        return Graph(
            title=f"Plugin Failures",
            dataSource=DATASOURCE,
            targets=[
                Target(
                    expr='{__name__=~"flyte:propeller:all:node:plugin:.*_failure_unlabeled"}',
                    refId='A',
                ),
                Target(
                    expr='{__name__=~"flyte:propeller:all:node:plugin:.*_success_unlabeled"}',
                    refId='A',
                ),
            ],
            yAxes=YAxes(
                YAxis(format=OPS_FORMAT),
                YAxis(format=SHORT_FORMAT),
            ),
        )

    @staticmethod
    def node_exec_latency() -> Graph:
        return Graph(
            title=f"Node Exec latency quantile and workflow",
            dataSource=DATASOURCE,
            targets=[
                Target(
                    expr=f'sum(flyte:propeller:all:node:node_exec_latency_us) by (quantile, wf) / 1000',
                    refId='A',
                ),
            ],
            yAxes=single_y_axis(format=MILLISECONDS_FORMAT),
        )

    @staticmethod
    def node_event_recording_latency() -> Graph:
        return Graph(
            title=f"Node Event event recording latency quantile and workflow",
            dataSource=DATASOURCE,
            targets=[
                Target(
                    expr=f'sum(flyte:propeller:all:node:event_recording:success_duration_ms) by (quantile, wf)',
                    refId='A',
                ),
                Target(
                    expr=f'sum(flyte:propeller:all:node:event_recording:failure_duration_ms) by (quantile, wf)',
                    refId='B',
                ),
            ],
            yAxes=single_y_axis(format=MILLISECONDS_FORMAT),
        )

    @staticmethod
    def node_input_latency() -> Graph:
        return Graph(
            title=f"Node Input latency quantile and workflow",
            dataSource=DATASOURCE,
            targets=[
                Target(
                    expr=f'sum(flyte:propeller:all:node:node_input_latency_ms) by (quantile, wf)',
                    refId='A',
                ),
            ],
            yAxes=single_y_axis(format=MILLISECONDS_FORMAT),
        )

    @staticmethod
    def metastore_failures():
        # Copy counts sum(rate(flyte:propeller:all:metastore:copy:overall_unlabeled_ms_count[5m]))
        return Graph(
            title=f"Failures from metastore",
            dataSource=DATASOURCE,
            targets=[
                Target(
                    expr=f'sum(rate(flyte:propeller:all:metastore:head_failure_unlabeled[5m]))',
                    legendFormat="head-failure",
                    refId='A',
                ),
                Target(
                    expr=f'sum(rate(flyte:propeller:all:metastore:bad_container_unlabeled[5m]))',
                    legendFormat="bad-container",
                    refId='A',
                ),
                Target(
                    expr=f'sum(rate(flyte:propeller:all:metastore:bad_key_unlabeled[5m]))',
                    legendFormat="bad-key",
                    refId='A',
                ),
                Target(
                    expr=f'sum(rate(flyte:propeller:all:metastore:read_failure_unlabeled[5m]))',
                    legendFormat="read-failure",
                    refId='A',
                ),
                Target(
                    expr=f'sum(rate(flyte:propeller:all:metastore:write_failure_unlabeled[5m]))',
                    legendFormat="write-failure",
                    refId='A',
                ),
            ],
            yAxes=YAxes(
                YAxis(format=OPS_FORMAT),
                YAxis(format=SHORT_FORMAT),
            ),
        )

    @staticmethod
    def metastore_cache_hit_percentage(interval: int) -> Graph:
        """
        TODO replace with metric math maybe?
        """
        return Graph(
            title="cache hit percentage",
            dataSource=DATASOURCE,
            targets=[
                Target(
                    expr=f'(sum(rate(flyte:propeller:all:metastore:cache_hit[{interval}m])) * 100) / (sum(rate(flyte:propeller:all:metastore:cache_miss[{interval}m])) + sum(rate(flyte:propeller:all:metastore:cache_hit[{interval}m])))',
                    refId='A',
                ),
            ],
            yAxes=YAxes(
                YAxis(format=OPS_FORMAT),
                YAxis(format=SHORT_FORMAT),
            ),
        )

    @staticmethod
    def metastore_latencies(collapse: bool) -> Row:
        return Row(
            title=f"Metastore latencies",
            collapse=collapse,
            panels=[
                Graph(
                    title=f"Metastore copy latency",
                    dataSource=DATASOURCE,
                    targets=[
                        Target(
                            expr=f'sum(flyte:propeller:all:metastore:copy:overall_unlabeled_ms) by (quantile)',
                            refId='A',
                        ),
                    ],
                    yAxes=single_y_axis(format=MILLISECONDS_FORMAT),
                ),
                Graph(
                    title=f"Metastore write latency by workflow",
                    dataSource=DATASOURCE,
                    targets=[
                        Target(
                            expr='sum(flyte:propeller:all:metastore:write_ms) by (quantile, wf)',
                            refId='A',
                        ),
                    ],
                    yAxes=single_y_axis(format=MILLISECONDS_FORMAT),
                ),
                Graph(
                    title=f"Metastore read open latency by workflow",
                    dataSource=DATASOURCE,
                    targets=[
                        Target(
                            expr='sum(flyte:propeller:all:metastore:read_open_ms) by (quantile, wf)',
                            refId='A',
                        ),
                    ],
                    yAxes=single_y_axis(format=MILLISECONDS_FORMAT),
                ),
                Graph(
                    title=f"Metastore head latency by workflow",
                    dataSource=DATASOURCE,
                    targets=[
                        Target(
                            expr='sum(flyte:propeller:all:metastore:head_ms) by (quantile, wf)',
                            refId='A',
                        ),
                    ],
                    yAxes=single_y_axis(format=MILLISECONDS_FORMAT),
                ),
                Graph(
                    title=f"Metastore fetch latency by workflow",
                    dataSource=DATASOURCE,
                    targets=[
                        Target(
                            expr='sum(flyte:propeller:all:metastore:proto_fetch_ms) by (quantile, wf)',
                            legendFormat="proto-fetch",
                            refId='A',
                        ),

                        Target(
                            expr='sum(flyte:propeller:all:metastore:remote_fetch_ms) by (quantile, wf)',
                            legendFormat="remote-fetch",
                            refId='B',
                        ),
                    ],
                    yAxes=single_y_axis(format=MILLISECONDS_FORMAT),
                ),
            ]
        )

    @staticmethod
    def metastore_metrics(interval: int, collapse: bool) -> Row:
        return Row(
            title="Metastore failures and cache",
            collapse=collapse,
            panels=[
                FlytePropeller.metastore_cache_hit_percentage(interval),
                FlytePropeller.metastore_failures(),
            ],
        )

    @staticmethod
    def node_metrics(collapse: bool) -> Row:
        return Row(
            title="Node Metrics",
            collapse=collapse,
            panels=[
                FlytePropeller.node_exec_latency(),
                FlytePropeller.node_input_latency(),
                FlytePropeller.node_event_recording_latency(),
            ],
        )

    @staticmethod
    def core_metrics(interval: int, collapse: bool) -> Row:
        return Row(
            title="Core metrics",
            collapse=collapse,
            panels=[
                FlytePropeller.create_free_workers(),
                FlytePropeller.abort_errors(),
                FlytePropeller.system_errors(),
                FlytePropeller.plugin_success_vs_failures(),
                FlytePropeller.round_latency(interval),
                FlytePropeller.round_latency_per_wf(interval),
                FlytePropeller.round_panic(),
                FlytePropeller.workflows_per_project(),
            ],
        )

    @staticmethod
    def create_all_rows(interval: int = 5) -> typing.List[Row]:
        return [
            FlytePropeller.core_metrics(interval, False),
            FlytePropeller.metastore_metrics(interval, True),
            FlytePropeller.metastore_latencies(True),
            FlytePropeller.node_metrics(True),
        ]


dashboard = Dashboard(
    tags=["flyte", "prometheus", "flytepropeller", "flyte-dataplane"],
    inputs=[
        DataSourceInput(
            name=DATASOURCE_NAME,
            label="Prometheus",
            description="Prometheus server that connects to Flyte",
            pluginId="prometheus",
            pluginName="Prometheus",
        ),
    ],
    editable=False,
    title="Flyte Propeller Dashboard (via Prometheus)",
    rows=FlytePropeller.create_all_rows(interval=5),
    description="Flyte Propeller Dashboard. This is great for monitoring FlytePropeller / Flyte data plane deployments. This is mostly useful for the Flyte deployment maintainer",
).auto_panel_ids()

if __name__ == "__main__":
    print(dashboard.to_json_data())
