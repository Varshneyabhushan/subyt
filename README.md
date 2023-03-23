# SubYT
This project is choose the subset of youtube videos and display them on a web page. 
We can browse over them in the reverse chronological order of creation date of the videos,
or search by the title of the videos

# Setup

## step 1
Copy all the contents of sample.env file and fill them up accordingly. 


`MONGO_PORT` : port at which mongo server should be exposed
`MONGO_DATA_PATH` : folder that should be used by mongo database, to store it's data.
When the folder mentioned doesn't exist, it will be created.

```
##These are the credentials which can be used to login to mongodb
MONGO_INITDB_ROOT_USERNAME=randomUser
MONGO_INITDB_ROOT_PASSWORD=randomPassword
MONGO_INITDB_DATABASE=videosService
```

`VIDEOS_SERVICE_PORT` : port at which apis are exposed. (core of the project)

`YOUTUBE_SEARCH_DEV_KEY` -> Developer key from youtube api
Get you own key at https://console.cloud.google.com/apis/library/youtube.googleapis.com

`YOUTUBE_MAX_RESULTS` -> maximum results to get from youtube in a single request


`TOPIC` -> youtube videos are synced from youtube, based on the topic we choose here

`SCHEDULER_REQUEST_COOL_DOWN` -> cool down between requests to youtube api. 
If this is too frequent, we may run out of quota per hour/ quota per day. 

`SCHEDULER_SYNC_COOL_DOWN` -> total time to rest, when all the videos are fetched initially.
Initial fetch brings youtube videos those have been published from _120 days_

`YTWORKER_CHECKPOINT` -> file path where the clue about what videos to fetch next, 
or if at all to stop fetching the videos. This has to point to a valid file, even when it's empty. 
When its pointed to an existing directory, <span style="color:red"> be aware that </span>, it may delete it's contents. 

`SEARCH_DB_PORT` -> port at which elastic search apis are exposed.

`SEARCH_DB_PATH` -> folder used by elastic search, to store it's data. 
When the folder mentioned doesn't exist, a new folder is created and used


## step 2
Once the .env file is made, run docker compose command with the docker-compose.yml file

`docker compose up`

Worker starts its job and fetches videos to the database. Required APIS are also hosted,
which can be tested at the port that is mentioned in variable `VIDEOS_SERVICE_PORT` in .env file

## step 3

Move to `web` folder and make the .env file guided by sample.env file that is present in the same folder.

`REACT_APP_SERVICE_URL` : should point at the URL where, we have hosted the videosService (core of the project). 

`REACT_APP_TOPIC` : topic on which youtube videos are searched. This will be displayed in the web page
Leave it empty if the tag is not required on the webpage.

## step 4

start the web page by running

`npm start`

or build the code using `npm run build`, but it should be hosted by a server

`serve -s build`

