docker day 7
------------

Topics

* kubernetes operators
* crds
* container metrics
* google sre
* deploy prometheus stack
* prometheus
* alertmanager
* prometheus-operator
* deploy hello app with metrics

kubernetes operators
--------------------

The Operator pattern aims to capture the key aim of a human operator who is managing a service or set of services. Human operators who look after specific applications and services have deep knowledge of how the system ought to behave, how to deploy it, and how to react if there are problems.

People who run workloads on Kubernetes often like to use automation to take care of repeatable tasks. The Operator pattern captures how you can write code to automate a task beyond what Kubernetes itself provides.

* [Overview](https://kubernetes.io/docs/concepts/extend-kubernetes/operator/)
* [Pros](https://thenewstack.io/why-kubernetes-operators-will-unleash-your-developers-by-reducing-complexity/)
* [Cons](https://thenewstack.io/kubernetes-when-to-use-and-when-to-avoid-the-operator-pattern/)

crds
----

Custom resources are extensions of the Kubernetes API.

* [Overview](https://kubernetes.io/docs/concepts/extend-kubernetes/api-extension/custom-resources/)
* [Defining a schema](https://kubernetes.io/blog/2019/06/20/crd-structural-schema/)

container metrics
-----------------

When running containers in a distributed environment like kubernetes a proper metrics and monitoring system becomes an absolute necessity. With pods constantly coming and going you need some unified method of determining whether they are functioning correctly.

google sre
----------

It’s impossible to manage a service correctly, let alone well, without understanding which behaviors really matter for that service and how to measure and evaluate those behaviors. To this end, we would like to define and deliver a given level of service to our users, whether they use an internal API or a public product.

Books

* [Books](https://landing.google.com/sre/books/)
* [SLOs](https://landing.google.com/sre/sre-book/chapters/service-level-objectives/)
* [Monitoring Distributed Systems](https://landing.google.com/sre/sre-book/chapters/monitoring-distributed-systems/) Defines 4 Golden signals here
* [Practical Alerting](https://landing.google.com/sre/sre-book/chapters/practical-alerting/)
* [Alerting on SLOs](https://landing.google.com/sre/workbook/chapters/alerting-on-slos/)

deploy prometheus stack
-----------------------

    kubectl apply -f prom-manifests/setup
    kubectl apply -f prom-manifests/

* View Prometheus: [http://localhost:9090](http://localhost:9090)
* View Alertmanager: [http://localhost:9093](http://localhost:9093)

prometheus
----------

Prometheus is a metrics and monitoring system that fits particularly well into the kubernetes environment

* [Overview](https://prometheus.io/docs/introduction/overview/)
* [Data Model](https://prometheus.io/docs/concepts/data_model/)
* [Metric Types](https://prometheus.io/docs/concepts/metric_types/)

After reading a basic overview log into Prometheus and go through all of the tabs

You can find the Openshift Prometheus link on the Administrator->Monitoring->Metrics page

alertmanager
------------

Alertmanager manages and forwards alerts that come from prometheus. It is not responsible for generating alerts.

* [Overview](https://prometheus.io/docs/alerting/latest/overview/)

After reading a basic overview log into Alertmanager and go through all of the tabs and alerts

You can find the Openshift Alertmanager link on the Administrator->Monitoring->Alerting page

prometheus-operator
-------------------

Prometheus Operator allows easy automation of Prometheus inside of the kubernetes environment using CRDs. Important CRDs include:

* Prometheus
* Alertmanager
* ServiceMonitor
* PodMonitor
* PrometheusRule

Links:

* [Intro with good diagram](https://devops.college/prometheus-operator-how-to-monitor-an-external-service-3cb6ac8d5acb)
* [Overview](https://github.com/prometheus-operator/prometheus-operator/blob/master/README.md)
* [API Docs](https://github.com/prometheus-operator/prometheus-operator/blob/master/Documentation/api.md)

deploy hello app with metrics
-----------------------------

Deploy hello app and query the metrics in prometheus

* Build and deploy app:

        docker build app/ -t hello-api:v0.2.0
        kubectl apply -f app/deployment.yml
        kubectl apply -f app/svc.yml
        kubectl apply -f app/servicemonitor.yml

* Query the apps metric endpoint [http://localhost:8090/metrics](http://localhost:8090/metrics)
* Choose a metric from the `/metrics` view and query it in prometheus
    * If the metrics haven’t shown up in prometheus view the Targets tab and look for http-api. It can take up to 5 minutes for it to show up
* Call [http://localhost:8090/](http://localhost:8090/) several times to increase metric counters
* Check the data in prometheus or the `/metrics` endpoint to see those requests reflected
* Check alertmanager to see if the example alert has show up yet
