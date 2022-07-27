# Test service
---
	1.docker compose up -d				//to start DB container
	2.add to ENV "KV_VIPER_FILE=./config.yaml"    	//path to config
	3.run service with "migrate up" args            //to create nessecary tables
	4.run service with "run service" args          	//to start service

Create blob
---
	POST    /integrations/testService/blob            

Get blob by id
---
	GET     /integrations/testService/blob/{id}      

Get all blobs
---
	GET     /integrations/testService/blob?filter?[author_id]={author_id}

Delete blob by id
---
	DELETE  /integrations/testService/blob/{id}        
