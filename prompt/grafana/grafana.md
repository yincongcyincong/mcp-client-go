1. `add_activity_to_incident`
    - "Add a note to incident ID `abc123`: “We’ve contacted customer support.” The event time is `2025-05-06T14:00:00Z`."

2. `create_incident`
    - "Create a new incident titled “Service Outage” with severity “critical” and status “active”. It’s a drill. Add labels `region:us-east-1` and `service:payment`. The room prefix is `incident-`. Attach this screenshot: `https://example.com/screenshot.png` with the caption “Error page screenshot”".

3. `get_alert_rule_by_uid`
    - "Show me the details of the alert rule with UID `alert-456`".

4. `get_current_oncall_users`
    - "Who is currently on-call for the schedule ID `schedule-789`?"

5. `get_dashboard_by_uid`
    - "Retrieve the dashboard with UID `dash-001`".

6. `get_datasource_by_name`
    - "Get the datasource named `Loki-Logs`".

7. `get_datasource_by_uid`
    - "I want the datasource with UID `ds-987`".

8. `get_incident`
    - "Get the full details of incident ID `incident-abc456`".

9. `get_oncall_shift`
    - "Show me the shift details for shift ID `shift-101`".

10. `list_alert_rules`
    - "List all alert rules that are currently firing and have the label `service=api`. Limit the results to 50".

11. `list_contact_points`
    - "List all notification contact points".

12. `list_datasources`
    - "List all datasources of type `prometheus`".

13. `list_incidents`
    - "List all active incidents, excluding drills".

14. `list_loki_label_names`
    - "Show me all unique label keys from the Loki datasource with UID `loki-123` for the past 2 hours".

15. `add_activity_to_incident`:
    - "Add a note to incident `incident-12345` with the content 'The body of the activity content' and event time `2025-04-27T14:00:00Z`."

16. `create_incident`:
    - "Create a high-priority incident titled 'Critical Incident' with the severity set to 'High', status set to 'Active', and a label 'Critical'. Attach a file with URL 'https://example.com/attachment' and caption 'Attachment Caption'."

17. `get_alert_rule_by_uid`:
    - "Retrieve details for the alert rule with UID `alert-rule-123`."

18. `get_current_oncall_users`:
    - "Get the current on-call users for the schedule `schedule-abc123`."

19. `get_dashboard_by_uid`:
    - "Retrieve the dashboard with UID `dashboard-xyz789`."

20. `get_datasource_by_name`:
    - "Get the data source details for `prometheus-datasource`."

21. `get_datasource_by_uid`:
    - "Retrieve data source information for UID `datasource-uid-456`."

22. `get_incident`:
    - "Get the details of the incident with ID `incident-12345`."

23. `get_oncall_shift`:
    - "Get the details of the on-call shift with ID `shift-001`."

24. `list_alert_rules`:
    - "List all alert rules with severity 'critical', returning the first 50 records."

25. `list_contact_points`:
    - "List all notification contact points filtered by name 'Alert Contact', returning up to 100 records."

26. `list_datasources`:
    - "List all data sources of type 'Prometheus'."

27. `list_incidents`:
    - "List all incidents with status 'active', excluding drill incidents, and return up to 100 records."

