### Expected question in CKAD

1. Create Pod with specified name and image.
    * Set specified labels
    * Set environment variable
    
2. Two question related to multi container pod
    * Convert single container into multi container pod (Ambassador pattern)
    * Create a pod in which 
        * One pod spits out the log and other is processing the log and forward it oto fluentd
        * Shared volume should be not be persistent
        * It will ask you create a configmap from given file and mount it to one of container which is processing the 
          log.
        
3. Create a configmap and create a pod which uses the config map to set specified environment variable.

4. Configure a pod to use set of network policy. Network policies will be created already to control egress and ingress
   traffic

5. Create a secret and a pod. Mount the secret into the pod's container in specified path

6. Configure a pod to use readiness and liveness probe

7. Create a job with specified image name, completion, backoff limit and activeDeadlineSecond. 
   Specific command to run inside the container will be given
   
8. Give pods named app, db and cache. 4 network policy already created.
   Configure pod app in such way that it uses defined network policy to only allow traffic from db and cache 
   and should only able to send traffic to db and cache.
   
9. Create a cronJob with specified image, command and should run on mentioned schedule time.

10. find a pod which is using most CPU/Memory in a particular namespace

11. Fix a deployment with image pull error. 

12. Fix a deployment with readiness probe error.

13. Create a PV. Create a PVC to claim the PC and finally Create a pod which consume PVC

14. Create a deployment with specified image name
    * set new image name in deployment
    * rollout with previous version
    
15. Create a deployment and collect summary in json format.

16. Configure a pod to run process as root user and set specified capabilities.

17. Create a deployment and expose it inside cluster.

18. Create a deployment and expose it as node port.

19. Create a pod with specified resource request and limit

20. Create a namespace and create a pod in the namespace.