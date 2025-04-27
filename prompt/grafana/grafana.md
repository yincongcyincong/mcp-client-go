1. **List all available label names in a Loki datasource for the given time range:**

   - "Please list all available label names in the Loki datasource with UID `datasourceUid` for the time range starting from `startRfc3339` and ending at `endRfc3339`."

2. **Retrieve all possible values for a specific label in Loki within the given time range:**

   - "Retrieve all possible values for the label `labelName` in the Loki datasource with UID `datasourceUid` for the time range starting from `startRfc3339` and ending at `endRfc3339`."

3. **List OnCall schedules for a specific team:**

   - "Please list the OnCall schedules for the team with ID `teamId`, returning page `page`. If a specific schedule ID `scheduleId` is provided, return the details for that schedule."

4. **List teams from Grafana OnCall:**

   - "Please list all teams from Grafana OnCall, returning page `page`."

5. **List users from Grafana OnCall:**

   - "Please list all users from Grafana OnCall, returning page `page`. If a user ID `userId` or username `username` is provided, return the details for that specific user."

6. **List label names in a Prometheus datasource:**

   - "List all label names in the Prometheus datasource with UID `datasourceUid` for the time range starting from `startRfc3339` and ending at `endRfc3339`. Optionally, filter the results with `matches` and limit the results to `limit`."

7. **Get the values of a label in Prometheus:**

   - "Retrieve the values of the label `labelName` in the Prometheus datasource with UID `datasourceUid` for the time range starting from `startRfc3339` and ending at `endRfc3339`. Optionally, filter the results with `matches` and limit the results to `limit`."

8. **List Prometheus metric metadata:**

   - "Please list the metric metadata in the Prometheus datasource with UID `datasourceUid`, with an optional limit of `limit` metrics and `limitPerMetric` metrics per metric. If a specific metric `metric` is provided, return details for that metric."

9. **List metric names in a Prometheus datasource that match a given regex:**

   - "List all metric names in the Prometheus datasource with UID `datasourceUid` that match the regex `regex`. Return up to `limit` results on page `page`."

10. **Query and retrieve log entries or metric values from a Loki datasource:**

    - "Please query the Loki datasource with UID `datasourceUid` using the LogQL query `logql` for the time range starting from `startRfc3339` and ending at `endRfc3339`. Optionally, return results in `direction` order and limit the number of log lines to `limit`."

11. **Query statistics about log streams in a Loki datasource:**

    - "Please query statistics about log streams in the Loki datasource with UID `datasourceUid` using the LogQL matcher `logql` for the time range starting from `startRfc3339` and ending at `endRfc3339`."

12. **Query Prometheus using a range or instant request:**

    - "Please query the Prometheus datasource with UID `datasourceUid` using the PromQL expression `expr`. The query type is `queryType` with a time range from `startRfc3339` to `endRfc3339`. If `queryType` is 'range', use a step size of `stepSeconds`."

13. **Search for dashboards:**

    - "Please search for dashboards using the query `query`."

14. **Create or update a dashboard:**

    - "Please create or update the dashboard with the provided JSON `dashboard`. If the dashboard is to be saved in a folder, use the folder UID `folderUid`. Set the commit message to `message` and specify whether to overwrite the dashboard with the `overwrite` flag."

15. **Prompt for `add_activity_to_incident`:**
    - "Add a note to incident `incident-12345` with the content 'The body of the activity content' and event time `2025-04-27T14:00:00Z`."

16. **Prompt for `create_incident`:**
    - "Create a high-priority incident titled 'Critical Incident' with the severity set to 'High', status set to 'Active', and a label 'Critical'. Attach a file with URL 'https://example.com/attachment' and caption 'Attachment Caption'."

17. **Prompt for `get_alert_rule_by_uid`:**
    - "Retrieve details for the alert rule with UID `alert-rule-123`."

18. **Prompt for `get_current_oncall_users`:**
    - "Get the current on-call users for the schedule `schedule-abc123`."

19. **Prompt for `get_dashboard_by_uid`:**
    - "Retrieve the dashboard with UID `dashboard-xyz789`."

20. **Prompt for `get_datasource_by_name`:**
    - "Get the data source details for `prometheus-datasource`."

21. **Prompt for `get_datasource_by_uid`:**
    - "Retrieve data source information for UID `datasource-uid-456`."

22. **Prompt for `get_incident`:**
    - "Get the details of the incident with ID `incident-12345`."

23. **Prompt for `get_oncall_shift`:**
    - "Get the details of the on-call shift with ID `shift-001`."

24. **Prompt for `list_alert_rules`:**
    - "List all alert rules with severity 'critical', returning the first 50 records."

25. **Prompt for `list_contact_points`:**
    - "List all notification contact points filtered by name 'Alert Contact', returning up to 100 records."

26. **Prompt for `list_datasources`:**
    - "List all data sources of type 'Prometheus'."

27. **Prompt for `list_incidents`:**
    - "List all incidents with status 'active', excluding drill incidents, and return up to 100 records."

