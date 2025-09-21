# API endpoints

### Sign up

curl --location 'localhost:3000/users/sign-up' \
--header 'Content-Type: application/json' \
--data-raw '{
    "email": "",
    "password": "",
    "name": "Francisco",
    "heightCM": 175.00
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
--header 'Authorization: ' \
--header 'Content-Type: application/json' \
--data '{
    "name": "Rutina de femoral"
}'

### Get routines by user

curl --location 'localhost:3000/routines' \
--header 'Authorization: '

### Create exercise by user

curl --location 'localhost:3000/exercises' \
--header 'Authorization: ' \
--header 'Content-Type: application/json' \
--data '{
    "name": "Dominadas",
    "routineID": 2,
    "proposedRounds": 4,
    "repetitions": 15 
}'

### Get execises by user

curl --location 'localhost:3000/exercises/routine/2' \
--header 'Authorization: '

### Create progress history by exercise

curl --location 'localhost:3000/progress-histories' \
--header 'Authorization: ' \
--header 'Content-Type: application/json' \
--data '{
    "exerciseID": 1,
    "weightUsedKG": 65.2,
    "date": "2025-09-20T01:15:03.177Z",
    "repsDone": 12,
    "roundsDone": 4
}'

## Get progress history by exercise

curl --location 'localhost:3000/progress-histories/exercise/1' \
--header 'Authorization: '

### Create weight history by user

curl --location 'localhost:3000/weight-histories/user' \
--header 'Authorization: ' \
--header 'Content-Type: application/json' \
--data '{
    "weightKG": 65.2,
    "date": "2025-09-20T01:15:03.177Z"
}'

## Get weight history by user

curl --location 'localhost:3000/weight-histories/user' \
--header 'Authorization: '