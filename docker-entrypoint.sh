#!/bin/bash
set -eo pipefail

./dockerhub-webhook start \
   --hostname "${DW_HOSTNAME}" \
   --port ${DW_PORT} \
   --profile-port ${DW_PROF_PORT} \
   --max-conn ${DW_MAX_CONN} \
   --max-procs ${DW_MAX_PROCS} \
   --debug "${DW_DEBUG}" \
   --valid-tokens "${DW_VALID_TOKENS}" \
   --namespace "${DW_NAMESPACE}" \
   --alive-path "${DW_ALIVE_PATH}" \
   --notify-path "${DW_NOTIFY_PATH}" \
   --status-path "${DW_STATUS_PATH}" \
   --target-host "${DW_TARGET_HOST}" \
   --target-port ${DW_TARGET_PORT} \
   --target-path "${DW_TARGET_PATH}" \
   --target-token "${DW_TARGET_TOKEN}"
