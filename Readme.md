## Run container
### Prerequisites: Docker Daemon should be running (install and run Docker Desktop or Orbstack.dev)

### Basic commands:
- `make build` (build and up container)

- `make up` (just for up container)
- `make down` (shutdown container)
- `show_logs` (show logs from the container)

Ensure the Container has started and the default network was created before testing the endpoints:
```[+] Running 2/2
 ✔ Network template-webserver_default  Created                                    
 ✔ Container go-app                    Started    
 ```

### Test the Endpoints: 
You can use tools like Postman or curl (make sure one or another is installed) to test the endpoints:
```
GET: curl http://localhost:8000/items
POST: curl -X POST -d '{"name":"NewItem","price":300}' http://localhost:8000/items
PUT: curl -X PUT -d '{"name":"UpdatedItem","price":500}' http://localhost:8000/items/1
DELETE: curl -X DELETE http://localhost:8000/items/1
```
You can use Webbrowser to send GET requests: e.g. `http://localhost:8000/items`