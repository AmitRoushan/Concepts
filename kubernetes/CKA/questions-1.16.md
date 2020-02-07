1. Nginx deployments with pod name nginx-app,with volumemounts
2. How to get max CPU consuming per pod and redirecting output to specified file
3. List all PersistentVolumes sorted by their name and redirecting output to specified file
4. In master its showing not ready , how to troubleshoot (Check list service file, binaries is their or not /usr/local , conf file is their or not)
5. Take the etcd backup and copy to specified path and etcd bootstrapping (ca certificate set in Etcd service file, and restart) make sure We need to know All parameters in ETCD service file
6. Setup a cluster using kubeadm init using the config /etc/kubeadm.conf 
7. init container (weather container path one specific file is their or not, if its not their pod should restart )
8. Create redis Pod and it should be deployed Specified NameSpace (with Volume)
9. Create Nginx Pod with specified Labels and Pod should schedule on Specific node.
10. Create a Deployment of Pod with Specified Label and selectors with VolumeMounts, Volumes
11. Deploy a pod with Three Containers(nginx+Redis+Memcached)
12. Deploy a pod with replica 3,(Deployment)
13. Deploy pod nginx:v1, then rollupdate to v2, then rollback previous Version
14. Create pod and Svc, Use the utility nslookup to look up the DNS records of the service and pod and Outputs redirecting to Specified Path.
15. Start a pod automatically by keeping manifest in /etc/kubernetes/manifests
16. Create Pv with 1Gi, host path
17. Write deployment yaml  keep it in given path
18. Deploy redis pod with Specified Volumemounts,volumes, but data should Persistent , volume should not persistent
19. New CA generation and Attaching to kubernetes cluster component ETCD.
20. Create Random-nginx-svc service should routed to Front-end pod
21. Increase Presentation Pods count 4
22. Create Daemon set with specified Labels, volumeMounts, volumes, Secrets,
23. Set the label "name=wk8s-node-0" to wk8s-node-0 and redirect the output to Specified File path
24. Deploy a pod with Specified secret should mapped to Configmap, configmap mapped to volumeMount with Volume