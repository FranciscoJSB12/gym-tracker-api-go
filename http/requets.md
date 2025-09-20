# API endpoints

### Sign up

curl --location 'localhost:3000/users/login' \
--header 'Content-Type: application/json' \
--data-raw '{
    "email": "",
    "password": "",
    "name": "Francisco",
    "height": 175.00
}'

### Login 

curl --location 'localhost:3000/users/login' \
--header 'Content-Type: application/json' \
--data-raw '{
    "email": "",
    "password": ""
}'

### Create routine by user

curl --location 'localhost:3000/routines' \
--header 'Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImphdmlmOEBnbWFpbC5jb20iLCJleHAiOjE3NTgzOTAxNzIsInVzZXJJZCI6MX0.ym4pSOrYO6MTdA9YomGbDIX2fJs1HLmcP64JQvWwwxM' \
--header 'Content-Type: application/json' \
--data '{
    "name": "Rutina de femoral"
}'

### Get routines by user

curl --location 'localhost:3000/routines' \
--header 'Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImphdmlmOEBnbWFpbC5jb20iLCJleHAiOjE3NTgzMzczMTQsInVzZXJJZCI6MX0.2fce5g4pVGeSwdb82__W22iIoWjpmWqpAIpR0ummAjI'

### Create exercise by user

curl --location 'localhost:3000/exercises' \
--header 'Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImphdmlmOEBnbWFpbC5jb20iLCJleHAiOjE3NTgzMzczMTQsInVzZXJJZCI6MX0.2fce5g4pVGeSwdb82__W22iIoWjpmWqpAIpR0ummAjI' \
--header 'Content-Type: application/json' \
--data '{
    "name": "Dominadas",
    "routineID": 2,
    "proposedRounds": 4,
    "repetitions": 15 
}'

### Get execises by user

curl --location 'localhost:3000/exercises/routine/2' \
--header 'Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImphdmlmOEBnbWFpbC5jb20iLCJleHAiOjE3NTgzMzczMTQsInVzZXJJZCI6MX0.2fce5g4pVGeSwdb82__W22iIoWjpmWqpAIpR0ummAjI'

### Create progress history by exercise

curl --location 'localhost:3000/progress-histories' \
--header 'Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImphdmlmOEBnbWFpbC5jb20iLCJleHAiOjE3NTgzMzczMTQsInVzZXJJZCI6MX0.2fce5g4pVGeSwdb82__W22iIoWjpmWqpAIpR0ummAjI' \
--header 'Content-Type: application/json' \
--data '{
    "exerciseID": 1,
    "weightUsedKG": 65.2,
    "date": "2025-09-20T01:15:03.177Z"
}'

## Get progress history by exercise

curl --location 'localhost:3000/progress-histories/exercise/1' \
--header 'Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImphdmlmOEBnbWFpbC5jb20iLCJleHAiOjE3NTgzOTAxNzIsInVzZXJJZCI6MX0.ym4pSOrYO6MTdA9YomGbDIX2fJs1HLmcP64JQvWwwxM'

### Create weight history by user

curl --location 'localhost:3000/weight-histories' \
--header 'Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImphdmlmOEBnbWFpbC5jb20iLCJleHAiOjE3NTgzNDQ3MTYsInVzZXJJZCI6MX0.BLcDDs2g9j-mvBwDVBF1jIQ8U9y8-e6bDWTRmGKXa6U' \
--header 'Content-Type: application/json' \
--data '{
    "weightKG": 65.2,
    "date": "2025-09-20T01:15:03.177Z"
}'

## Get weight history by user

curl --location 'localhost:3000/weight-histories' \
--header 'Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImphdmlmOEBnbWFpbC5jb20iLCJleHAiOjE3NTgzNDQ3MTYsInVzZXJJZCI6MX0.BLcDDs2g9j-mvBwDVBF1jIQ8U9y8-e6bDWTRmGKXa6U'