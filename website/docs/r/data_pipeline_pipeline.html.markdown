---
# ----------------------------------------------------------------------------
#
#     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
#
# ----------------------------------------------------------------------------
#
#     This code is generated by Magic Modules using the following:
#
#     Configuration: https:#github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/datapipeline/Pipeline.yaml
#     Template:      https:#github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/templates/terraform/resource.html.markdown.tmpl
#
#     DO NOT EDIT this file directly. Any changes made to this file will be
#     overwritten during the next generation cycle.
#
# ----------------------------------------------------------------------------
subcategory: "DataPipeline"
description: |-
  The main pipeline entity and all the necessary metadata for launching and managing linked jobs.
---

# google_data_pipeline_pipeline

The main pipeline entity and all the necessary metadata for launching and managing linked jobs.


To get more information about Pipeline, see:

* [API documentation](https://cloud.google.com/dataflow/docs/reference/data-pipelines/rest/v1/projects.locations.pipelines)
* How-to Guides
    * [Official Documentation](https://cloud.google.com/dataflow)

<div class = "oics-button" style="float: right; margin: 0 0 -15px">
  <a href="https://console.cloud.google.com/cloudshell/open?cloudshell_git_repo=https%3A%2F%2Fgithub.com%2Fterraform-google-modules%2Fdocs-examples.git&cloudshell_image=gcr.io%2Fcloudshell-images%2Fcloudshell%3Alatest&cloudshell_print=.%2Fmotd&cloudshell_tutorial=.%2Ftutorial.md&cloudshell_working_dir=data_pipeline_pipeline&open_in_editor=main.tf" target="_blank">
    <img alt="Open in Cloud Shell" src="//gstatic.com/cloudssh/images/open-btn.svg" style="max-height: 44px; margin: 32px auto; max-width: 100%;">
  </a>
</div>
## Example Usage - Data Pipeline Pipeline


```hcl
resource "google_service_account" "service_account" {
  account_id   = "my-account"
  display_name = "Service Account"
}

resource "google_data_pipeline_pipeline" "primary" {
  name         = "my-pipeline"
  display_name = "my-pipeline"
  type         = "PIPELINE_TYPE_BATCH"
  state        = "STATE_ACTIVE"
  region       = "us-central1"

  workload {
    dataflow_launch_template_request {
      project_id = "my-project"
      gcs_path   = "gs://my-bucket/path"
      launch_parameters {
        job_name = "my-job"
        parameters = {
          "name" : "wrench"
        }
        environment {
          num_workers                = 5
          max_workers                = 5
          zone                       = "us-centra1-a"
          service_account_email      = google_service_account.service_account.email
          network                    = "default"
          temp_location              = "gs://my-bucket/tmp_dir"
          bypass_temp_dir_validation = false
          machine_type               = "E2"
          additional_user_labels = {
            "context" : "test"
          }
          worker_region    = "us-central1"
          worker_zone      = "us-central1-a"

          enable_streaming_engine = "false"
        }
        update                 = false
        transform_name_mapping = { "name" : "wrench" }
      }
      location = "us-central1"
    }
  }
  schedule_info {
    schedule = "* */2 * * *"
  }
}
```

## Argument Reference

The following arguments are supported:


* `name` -
  (Required)
  "The pipeline name. For example': 'projects/PROJECT_ID/locations/LOCATION_ID/pipelines/PIPELINE_ID."
  "- PROJECT_ID can contain letters ([A-Za-z]), numbers ([0-9]), hyphens (-), colons (:), and periods (.). For more information, see Identifying projects."
  "LOCATION_ID is the canonical ID for the pipeline's location. The list of available locations can be obtained by calling google.cloud.location.Locations.ListLocations. Note that the Data Pipelines service is not available in all regions. It depends on Cloud Scheduler, an App Engine application, so it's only available in App Engine regions."
  "PIPELINE_ID is the ID of the pipeline. Must be unique for the selected project and location."

* `type` -
  (Required)
  The type of the pipeline. This field affects the scheduling of the pipeline and the type of metrics to show for the pipeline.
  https://cloud.google.com/dataflow/docs/reference/data-pipelines/rest/v1/projects.locations.pipelines#pipelinetype
  Possible values are: `PIPELINE_TYPE_UNSPECIFIED`, `PIPELINE_TYPE_BATCH`, `PIPELINE_TYPE_STREAMING`.

* `state` -
  (Required)
  The state of the pipeline. When the pipeline is created, the state is set to 'PIPELINE_STATE_ACTIVE' by default. State changes can be requested by setting the state to stopping, paused, or resuming. State cannot be changed through pipelines.patch requests.
  https://cloud.google.com/dataflow/docs/reference/data-pipelines/rest/v1/projects.locations.pipelines#state
  Possible values are: `STATE_UNSPECIFIED`, `STATE_RESUMING`, `STATE_ACTIVE`, `STATE_STOPPING`, `STATE_ARCHIVED`, `STATE_PAUSED`.


* `display_name` -
  (Optional)
  The display name of the pipeline. It can contain only letters ([A-Za-z]), numbers ([0-9]), hyphens (-), and underscores (_).

* `workload` -
  (Optional)
  Workload information for creating new jobs.
  https://cloud.google.com/dataflow/docs/reference/data-pipelines/rest/v1/projects.locations.pipelines#workload
  Structure is [documented below](#nested_workload).

* `schedule_info` -
  (Optional)
  Internal scheduling information for a pipeline. If this information is provided, periodic jobs will be created per the schedule. If not, users are responsible for creating jobs externally.
  https://cloud.google.com/dataflow/docs/reference/data-pipelines/rest/v1/projects.locations.pipelines#schedulespec
  Structure is [documented below](#nested_schedule_info).

* `scheduler_service_account_email` -
  (Optional)
  Optional. A service account email to be used with the Cloud Scheduler job. If not specified, the default compute engine service account will be used.

* `pipeline_sources` -
  (Optional)
  The sources of the pipeline (for example, Dataplex). The keys and values are set by the corresponding sources during pipeline creation.
  An object containing a list of "key": value pairs. Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.

* `region` -
  (Optional)
  A reference to the region

* `project` - (Optional) The ID of the project in which the resource belongs.
    If it is not provided, the provider project is used.



<a name="nested_workload"></a>The `workload` block supports:

* `dataflow_launch_template_request` -
  (Optional)
  Template information and additional parameters needed to launch a Dataflow job using the standard launch API.
  https://cloud.google.com/dataflow/docs/reference/data-pipelines/rest/v1/projects.locations.pipelines#launchtemplaterequest
  Structure is [documented below](#nested_workload_dataflow_launch_template_request).

* `dataflow_flex_template_request` -
  (Optional)
  Template information and additional parameters needed to launch a Dataflow job using the flex launch API.
  https://cloud.google.com/dataflow/docs/reference/data-pipelines/rest/v1/projects.locations.pipelines#launchflextemplaterequest
  Structure is [documented below](#nested_workload_dataflow_flex_template_request).


<a name="nested_workload_dataflow_launch_template_request"></a>The `dataflow_launch_template_request` block supports:

* `project_id` -
  (Required)
  The ID of the Cloud Platform project that the job belongs to.

* `validate_only` -
  (Optional)

* `launch_parameters` -
  (Optional)
  The parameters of the template to launch. This should be part of the body of the POST request.
  https://cloud.google.com/dataflow/docs/reference/data-pipelines/rest/v1/projects.locations.pipelines#launchtemplateparameters
  Structure is [documented below](#nested_workload_dataflow_launch_template_request_launch_parameters).

* `location` -
  (Optional)
  The regional endpoint to which to direct the request.

* `gcs_path` -
  (Optional)
  A Cloud Storage path to the template from which to create the job. Must be a valid Cloud Storage URL, beginning with 'gs://'.


<a name="nested_workload_dataflow_launch_template_request_launch_parameters"></a>The `launch_parameters` block supports:

* `job_name` -
  (Required)
  The job name to use for the created job.

* `parameters` -
  (Optional)
  The runtime parameters to pass to the job.
  'An object containing a list of "key": value pairs. Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.'

* `environment` -
  (Optional)
  The runtime environment for the job.
  https://cloud.google.com/dataflow/docs/reference/data-pipelines/rest/v1/projects.locations.pipelines#RuntimeEnvironment
  Structure is [documented below](#nested_workload_dataflow_launch_template_request_launch_parameters_environment).

* `update` -
  (Optional)
  If set, replace the existing pipeline with the name specified by jobName with this pipeline, preserving state.

* `transform_name_mapping` -
  (Optional)
  Map of transform name prefixes of the job to be replaced to the corresponding name prefixes of the new job. Only applicable when updating a pipeline.
  'An object containing a list of "key": value pairs. Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.'


<a name="nested_workload_dataflow_launch_template_request_launch_parameters_environment"></a>The `environment` block supports:

* `num_workers` -
  (Optional)
  The initial number of Compute Engine instances for the job.

* `max_workers` -
  (Optional)
  The maximum number of Compute Engine instances to be made available to your pipeline during execution, from 1 to 1000.

* `zone` -
  (Optional)
  The Compute Engine availability zone for launching worker instances to run your pipeline. In the future, workerZone will take precedence.

* `service_account_email` -
  (Optional)
  The email address of the service account to run the job as.

* `temp_location` -
  (Optional)
  The Cloud Storage path to use for temporary files. Must be a valid Cloud Storage URL, beginning with gs://.

* `bypass_temp_dir_validation` -
  (Optional)
  Whether to bypass the safety checks for the job's temporary directory. Use with caution.

* `machine_type` -
  (Optional)
  The machine type to use for the job. Defaults to the value from the template if not specified.

* `additional_experiments` -
  (Optional)
  Additional experiment flags for the job.

* `network` -
  (Optional)
  Network to which VMs will be assigned. If empty or unspecified, the service will use the network "default".

* `subnetwork` -
  (Optional)
  Subnetwork to which VMs will be assigned, if desired. You can specify a subnetwork using either a complete URL or an abbreviated path. Expected to be of the form "https://www.googleapis.com/compute/v1/projects/HOST_PROJECT_ID/regions/REGION/subnetworks/SUBNETWORK" or "regions/REGION/subnetworks/SUBNETWORK". If the subnetwork is located in a Shared VPC network, you must use the complete URL.

* `additional_user_labels` -
  (Optional)
  Additional user labels to be specified for the job. Keys and values should follow the restrictions specified in the labeling restrictions page. An object containing a list of key/value pairs.
  'Example: { "name": "wrench", "mass": "1kg", "count": "3" }.'
  'An object containing a list of "key": value pairs. Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.'

* `kms_key_name` -
  (Optional)
  'Name for the Cloud KMS key for the job. The key format is: projects//locations//keyRings//cryptoKeys/'

* `ip_configuration` -
  (Optional)
  Configuration for VM IPs.
  https://cloud.google.com/dataflow/docs/reference/data-pipelines/rest/v1/projects.locations.pipelines#WorkerIPAddressConfiguration
  Possible values are: `WORKER_IP_UNSPECIFIED`, `WORKER_IP_PUBLIC`, `WORKER_IP_PRIVATE`.

* `worker_region` -
  (Optional)
  The Compute Engine region (https://cloud.google.com/compute/docs/regions-zones/regions-zones) in which worker processing should occur, e.g. "us-west1". Mutually exclusive with workerZone. If neither workerRegion nor workerZone is specified, default to the control plane's region.

* `worker_zone` -
  (Optional)
  The Compute Engine zone (https://cloud.google.com/compute/docs/regions-zones/regions-zones) in which worker processing should occur, e.g. "us-west1-a". Mutually exclusive with workerRegion. If neither workerRegion nor workerZone is specified, a zone in the control plane's region is chosen based on available capacity. If both workerZone and zone are set, workerZone takes precedence.

* `enable_streaming_engine` -
  (Optional)
  Whether to enable Streaming Engine for the job.

<a name="nested_workload_dataflow_flex_template_request"></a>The `dataflow_flex_template_request` block supports:

* `project_id` -
  (Required)
  The ID of the Cloud Platform project that the job belongs to.

* `launch_parameter` -
  (Required)
  Parameter to launch a job from a Flex Template.
  https://cloud.google.com/dataflow/docs/reference/data-pipelines/rest/v1/projects.locations.pipelines#launchflextemplateparameter
  Structure is [documented below](#nested_workload_dataflow_flex_template_request_launch_parameter).

* `location` -
  (Required)
  The regional endpoint to which to direct the request. For example, us-central1, us-west1.

* `validate_only` -
  (Optional)
  If true, the request is validated but not actually executed. Defaults to false.


<a name="nested_workload_dataflow_flex_template_request_launch_parameter"></a>The `launch_parameter` block supports:

* `job_name` -
  (Required)
  The job name to use for the created job. For an update job request, the job name should be the same as the existing running job.

* `parameters` -
  (Optional)
  'The parameters for the Flex Template. Example: {"numWorkers":"5"}'
  'An object containing a list of "key": value pairs. Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.'

* `launch_options` -
  (Optional)
  Launch options for this Flex Template job. This is a common set of options across languages and templates. This should not be used to pass job parameters.
  'An object containing a list of "key": value pairs. Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.'

* `environment` -
  (Optional)
  The runtime environment for the Flex Template job.
  https://cloud.google.com/dataflow/docs/reference/data-pipelines/rest/v1/projects.locations.pipelines#FlexTemplateRuntimeEnvironment
  Structure is [documented below](#nested_workload_dataflow_flex_template_request_launch_parameter_environment).

* `update` -
  (Optional)
  Set this to true if you are sending a request to update a running streaming job. When set, the job name should be the same as the running job.

* `transform_name_mappings` -
  (Optional)
  'Use this to pass transform name mappings for streaming update jobs. Example: {"oldTransformName":"newTransformName",...}'
  'An object containing a list of "key": value pairs. Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.'

* `container_spec_gcs_path` -
  (Optional)
  Cloud Storage path to a file with a JSON-serialized ContainerSpec as content.


<a name="nested_workload_dataflow_flex_template_request_launch_parameter_environment"></a>The `environment` block supports:

* `num_workers` -
  (Optional)
  The initial number of Compute Engine instances for the job.

* `max_workers` -
  (Optional)
  The maximum number of Compute Engine instances to be made available to your pipeline during execution, from 1 to 1000.

* `zone` -
  (Optional)
  The Compute Engine availability zone for launching worker instances to run your pipeline. In the future, workerZone will take precedence.

* `service_account_email` -
  (Optional)
  The email address of the service account to run the job as.

* `temp_location` -
  (Optional)
  The Cloud Storage path to use for temporary files. Must be a valid Cloud Storage URL, beginning with gs://.

* `machine_type` -
  (Optional)
  The machine type to use for the job. Defaults to the value from the template if not specified.

* `additional_experiments` -
  (Optional)
  Additional experiment flags for the job.

* `network` -
  (Optional)
  Network to which VMs will be assigned. If empty or unspecified, the service will use the network "default".

* `subnetwork` -
  (Optional)
  Subnetwork to which VMs will be assigned, if desired. You can specify a subnetwork using either a complete URL or an abbreviated path. Expected to be of the form "https://www.googleapis.com/compute/v1/projects/HOST_PROJECT_ID/regions/REGION/subnetworks/SUBNETWORK" or "regions/REGION/subnetworks/SUBNETWORK". If the subnetwork is located in a Shared VPC network, you must use the complete URL.

* `additional_user_labels` -
  (Optional)
  Additional user labels to be specified for the job. Keys and values should follow the restrictions specified in the labeling restrictions page. An object containing a list of key/value pairs.
  'Example: { "name": "wrench", "mass": "1kg", "count": "3" }.'
  'An object containing a list of "key": value pairs. Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.'

* `kms_key_name` -
  (Optional)
  'Name for the Cloud KMS key for the job. The key format is: projects//locations//keyRings//cryptoKeys/'

* `ip_configuration` -
  (Optional)
  Configuration for VM IPs.
  https://cloud.google.com/dataflow/docs/reference/data-pipelines/rest/v1/projects.locations.pipelines#WorkerIPAddressConfiguration
  Possible values are: `WORKER_IP_UNSPECIFIED`, `WORKER_IP_PUBLIC`, `WORKER_IP_PRIVATE`.

* `worker_region` -
  (Optional)
  The Compute Engine region (https://cloud.google.com/compute/docs/regions-zones/regions-zones) in which worker processing should occur, e.g. "us-west1". Mutually exclusive with workerZone. If neither workerRegion nor workerZone is specified, default to the control plane's region.

* `worker_zone` -
  (Optional)
  The Compute Engine zone (https://cloud.google.com/compute/docs/regions-zones/regions-zones) in which worker processing should occur, e.g. "us-west1-a". Mutually exclusive with workerRegion. If neither workerRegion nor workerZone is specified, a zone in the control plane's region is chosen based on available capacity. If both workerZone and zone are set, workerZone takes precedence.

* `enable_streaming_engine` -
  (Optional)
  Whether to enable Streaming Engine for the job.

* `flexrs_goal` -
  (Optional)
  Set FlexRS goal for the job. https://cloud.google.com/dataflow/docs/guides/flexrs
  https://cloud.google.com/dataflow/docs/reference/data-pipelines/rest/v1/projects.locations.pipelines#FlexResourceSchedulingGoal
  Possible values are: `FLEXRS_UNSPECIFIED`, `FLEXRS_SPEED_OPTIMIZED`, `FLEXRS_COST_OPTIMIZED`.

<a name="nested_schedule_info"></a>The `schedule_info` block supports:

* `schedule` -
  (Optional)
  Unix-cron format of the schedule. This information is retrieved from the linked Cloud Scheduler.

* `time_zone` -
  (Optional)
  Timezone ID. This matches the timezone IDs used by the Cloud Scheduler API. If empty, UTC time is assumed.

* `next_job_time` -
  (Output)
  When the next Scheduler job is going to run.
  A timestamp in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine fractional digits. Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z".

## Attributes Reference

In addition to the arguments listed above, the following computed attributes are exported:

* `id` - an identifier for the resource with format `projects/{{project}}/locations/{{region}}/pipelines/{{name}}`

* `create_time` -
  The timestamp when the pipeline was initially created. Set by the Data Pipelines service.
  A timestamp in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine fractional digits. Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z".

* `last_update_time` -
  The timestamp when the pipeline was last modified. Set by the Data Pipelines service.
  A timestamp in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine fractional digits. Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z".

* `job_count` -
  Number of jobs.


## Timeouts

This resource provides the following
[Timeouts](https://developer.hashicorp.com/terraform/plugin/sdkv2/resources/retries-and-customizable-timeouts) configuration options:

- `create` - Default is 20 minutes.
- `update` - Default is 20 minutes.
- `delete` - Default is 20 minutes.

## Import


Pipeline can be imported using any of these accepted formats:

* `projects/{{project}}/locations/{{region}}/pipelines/{{name}}`
* `{{project}}/{{region}}/{{name}}`
* `{{region}}/{{name}}`
* `{{name}}`


In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import Pipeline using one of the formats above. For example:

```tf
import {
  id = "projects/{{project}}/locations/{{region}}/pipelines/{{name}}"
  to = google_data_pipeline_pipeline.default
}
```

When using the [`terraform import` command](https://developer.hashicorp.com/terraform/cli/commands/import), Pipeline can be imported using one of the formats above. For example:

```
$ terraform import google_data_pipeline_pipeline.default projects/{{project}}/locations/{{region}}/pipelines/{{name}}
$ terraform import google_data_pipeline_pipeline.default {{project}}/{{region}}/{{name}}
$ terraform import google_data_pipeline_pipeline.default {{region}}/{{name}}
$ terraform import google_data_pipeline_pipeline.default {{name}}
```

## User Project Overrides

This resource supports [User Project Overrides](https://registry.terraform.io/providers/hashicorp/google/latest/docs/guides/provider_reference#user_project_override).
