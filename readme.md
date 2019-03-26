# Prometheus + Kubernetes demo

## Step 1: Deploy
Helm is the easiest way to do this. Check out `https://github.com/kubernetes/charts/tree/master/stable/prometheus`.

We can run `helm install stable/prometheus` to get the stock prometheus server.

In this case, we will run: `helm install stable/prometheus --name prom-demo -f values.yaml` to use our custom yaml.

## Step 2: See it running
Running `minikube service prom-demo-prometheus-server` will bring up the browser with prometheus server running.

## Step 3: Check out Kubernetes stats
Check out `kube_pod_status_phase` to see our all our pods running.

## Step 4: Prometheus operators
More info here: `https://prometheus.io/docs/prometheus/latest/querying/operators/` but you can use the query language to clean up the data point we looked at above.

`sum(kube_pod_status_phase) by (phase)` will return the set of pods that are running grouped by their phases.


## Step 5: Check out node information

The node exporter gives you node relative information as well like CPU/disk usage etc.

Run `machine_cpu_cores` will return the CPU count which should match the number of CPU's in `kubectl describe nodes`

### Step 6: App metrics

1. Run the app: `kubectl apply -f twelve-clouds.yaml`
2. Visit the app after it's deployed: `minikube service twelve-clouds-service`
3. You'll see the  `/metrics` endpoint with go stats
4. Visit the `/hello` endpoint 2 times.
5. You'll see a new "hello_calls" metric in the dashboard
