# kubernetes-a-b-deployment-example

For personal testing only.

Steps:

0) Install kubernetes somewhere. Docker Desktop for example can run a single node cluster with one simple click.
1) Build the image: `docker build -t color:0.0.1 .`
2) Deploy: `kubectl apply -f ab.yml`
3) Create an endless loop curling the service in another shell: `while true; do curl localhost:8080; sleep 1; done`
4) Remove the comments in ab.yml from the second deployment
5) Apply the config again. Observe the curl output

Then play around. Make changes to the yaml file and apply it. Observe the curl outputs. Suggestions:

- Change the replica numbers an see how the distibutions change
- Change the "color"
- Scale a deployment to 0

Have fun.