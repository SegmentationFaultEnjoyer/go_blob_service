# testService

1. docker compose up -d                       -->   to start DB container
2. add to ENV "KV_VIPER_FILE=./config.yaml"   -->   path to config
3. run service with "migrate up" args         -->   to create nessecary tables
4. run service with "run service" args        -->   to start service

POST    /integrations/testService/user            -->   create user
POST    /integrations/testService/blob            -->   create blob
GET     /integrations/testService/blobs/:id       -->   get blob by id
GET     /integrations/testService/blobs/:user_id  -->   get all blobs by user id
DELETE  /integrations/testService/blob/:id        -->   delete blob by id
