strategy:
  type: RollingUpdate  # Specify the update strategy
  rollingUpdate:
    maxSurge: 1        
# Allow one additional pod over the desired replica count
# during update
    maxUnavailable: 1   
# Allow one pod to be unavailable during update


#We added this section to the .yaml file
