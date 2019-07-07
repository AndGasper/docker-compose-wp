## Docker Notes
```
ERROR: for   Cannot start service {servicename}: driver failed {projectname}_{servicename}  (f9e15ccc238b8519214e4999e057bb44db3b09ca4f5175ba55b2f2336f0f6b0a): ready allocated
```
- Error message:
```
ERROR: for {projectname}_{servicename}  
Cannot start service {servicename}: driver failed programming external connectivity on endpoint {projectname}_{servicename} ({HASH}): 
Bind for {ipv4}:{port} failed: port is already allocated
```
- Resolution: 
```
# Get a list of the running containers: 
$ docker ps
```
```
# "kill" the container allocated to the needed port 
> docker kill {container-id}
```


More advanced form: 
```
docker ps -q | xargs docker inspect --format '{{ .Id }} - {{ .NetworkSettings.Ports }}'
```
```
docker inspect --format='{{(index (index .NetworkSettings.Ports "8787/tcp") 0).HostPort}}' $INSTANCE_ID
```

## Find a specific port mapping 

```
docker inspect --format='{{(index (index .NetworkSettings.Ports "8787/tcp") 0).HostPort}}' $INSTANCE_ID
```
```
docker inspect --format='{{(index (index .NetworkSettings.Ports "80/udp") 0).HostPort}}' $INSTANCE_ID
```


```
docker ps -q | xargs docker inspect --format '{{range $p, $conf := .NetworkSettings.Ports}} {{$p}} -> {{(index $conf 0).HostPort}} {{end}}' $INSTANCE_ID
```
```
'{{range $p, $conf := .NetworkSettings.Ports}} {{$p}} -> {{(index $conf 0).HostPort}} {{end}}' $INSTANCE_ID
```

## Multiple Docker Compose Files: 
- It's something like:
``` 
docker-compose -f docker-compose.yml -f {environment}.yml
```



## Miscellaneous 
- [Docker Compose: Volumes + Networks Example](https://www.linux.com/learn/docker-volumes-and-networks-compose)
- [Docker Swarm Configs - Example Configurations](https://docs.docker.com/engine/swarm/configs/)
- [Docker Compose: Gutenberg Example](https://github.com/WordPress/gutenberg/blob/master/docker-compose.yml)s
- [Docker Library: WordPress](https://github.com/docker-library/wordpress)
