#### Modify value of ```ATTACKED_HOST``` of locust-cm.yaml file resides inside locust-descriptors directory with due api endpoint
#### ```kubectl apply -f locust-cm.yaml ```
#### ```kubectl apply -f scripts-cm.yaml ```
#### ```kubectl apply -f master-deployment.yaml ```
#### ```kubectl apply -f slave-deployment.yaml ```
#### ```kubectl apply -f service.yaml ```

Apply ingress if needed, 
#### ```kubectl apply -f service.yaml ```

####  To port-forward:

```kubectl port-forward svc/locust-master 8081:8089```

kubectl port-forward svc/locust-master 8081:8089
https://auto-lb-locust-yqgyiafw-qrpactdj-prod.bd-1.wpc.waltonelectronics.com/api/v1/blocks