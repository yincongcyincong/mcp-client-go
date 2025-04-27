1. **RunCommand**:
   "Please execute the following command on the specified ECS instances: `ls -l /home/user`. The command type is `RunShellScript`."

2. **StartInstances**:
   "I need to start the following ECS instances with IDs: `i-12345`, `i-67890` in the `cn-hangzhou` region."

3. **StopInstances**:
   "Please stop ECS instances with IDs: `i-12345`, `i-67890` in `cn-hangzhou` region. Force stop is not required."

4. **RebootInstances**:
   "Reboot ECS instances with IDs: `i-12345`, `i-67890` in the `cn-hangzhou` region. Force stop is not required."

5. **RunInstances**:
   "Create 5 new ECS instances using the specified image `img-12345`, type `ecs.g5.large`, with a security group ID `sg-12345` and VSwitch ID `vsw-12345` in `cn-hangzhou`."

6. **ResetPassword**:
   "Reset the password for ECS instances with IDs: `i-12345`, `i-67890` to `NewPassword123!` in the `cn-hangzhou` region."

7. **ReplaceSystemDisk**:
   "Replace the system disk of ECS instances with IDs: `i-12345`, `i-67890` with a new image `img-12345` in the `cn-hangzhou` region."

8. **StartRDSInstances**:
   "Please start the RDS instances with IDs: `rds-12345`, `rds-67890` in the `cn-hangzhou` region."

9. **StopRDSInstances**:
   "Stop RDS instances with IDs: `rds-12345`, `rds-67890` in the `cn-hangzhou` region."

10. **RebootRDSInstances**:
    "Reboot RDS instances with IDs: `rds-12345`, `rds-67890` in the `cn-hangzhou` region."

11. **GetCpuUsageData**:
    "Retrieve the CPU usage data for ECS instances with IDs: `i-12345`, `i-67890` in the `cn-hangzhou` region."

12. **GetCpuLoadavgData**:
    "Fetch the 1-minute average CPU load for ECS instances with IDs: `i-12345`, `i-67890` in the `cn-hangzhou` region."

13. **Get CPU 5-minute average load data:**
    - "Retrieve the 5-minute average CPU load for multiple ECS instances."
    - "Can you fetch the CPU load average for my ECS instances over the past 5 minutes?"

14. **Get CPU 15-minute average load data:**
    - "Fetch the 15-minute average CPU load data for selected ECS instances."
    - "Please show the CPU load average for my ECS instances over the past 15 minutes."

15. **Get memory usage data:**
    - "Can you get the memory usage statistics for my ECS instances?"
    - "Retrieve memory utilization data for the following ECS instances."

16. **Get memory utilization data:**
    - "Get the memory usage rate for my ECS instances."
    - "Fetch memory utilization data across my ECS instances."

17. **Get disk usage data:**
    - "Please fetch disk usage data for the ECS instances."
    - "Retrieve disk usage information for multiple ECS instances."

18. **Get total disk capacity data:**
    - "Fetch the total disk capacity for the ECS instances."
    - "Show me the disk total capacity for my ECS instances."

19. **Get disk used data:**
    - "Get the used disk space data for my ECS instances."
    - "Retrieve the disk usage information for ECS instances."

20. **List all OSS storage spaces in a region:**
    - "List all OSS storage buckets in the region cn-hangzhou."
    - "Can you show the available OSS buckets in the region?"

21. **Get all files in a specific OSS storage space:**
    - "List all objects in the specified OSS bucket in cn-hangzhou."
    - "Show me the contents of the OSS bucket named 'my-bucket' in the cn-hangzhou region."

22. **Create a new OSS storage space:**
    - "Create a new OSS bucket named 'new-bucket' with LRS redundancy in the cn-hangzhou region."
    - "Can you set up a new OSS bucket called 'my-new-bucket' with standard storage in cn-hangzhou?"

