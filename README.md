## Commands to build
```
docker-compose build && docker-compose up -d
```


## Connecting to an already running image
```
# Windows Specific Command: winpty docker exec -it 54263fde073e bash
docker exec -it {containerID} bash
```


## Configuring the IDE
- [Custom Dockerfile names](https://code.visualstudio.com/docs/languages/identifiers)


