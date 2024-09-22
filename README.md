# golang-mongodb-http-server
This is an educational project that depicts the usage of Golang programming language and how it can be used to create a web server that interacts with MongoDB database.

## App config
This app requires environment variable with name of GMWS_ENV. The values can be eaither 'dev' or 'prod' and it defaults to 'dev'.
<br> Based on the value, app will read properties from specific config file.

There are Java-style config files are created at the root of the project with the name of ```application-{env}.yml```

This project uses Viper library to read the config files and Environment Variables.

## How to run
- Set up environmet variable GMWS_ENV as described above
- Update the DB connection URI in ```application-{env}.yml```
- Run ```go mod tidy``` command to get all required dependency 
- run the server using below command
```
go run api-server/main.go
```

## APIs
### `GET /record/movie/{id}`
This should return entire movie record for the given id from the Database
    
### `POST /record/movie`
This should add the movie data in the database
#### request body
```json
{
  "plot": "Young Pauline is left a lot of money when her wealthy uncle dies. However, her uncle's secretary has been named as her guardian until she marries, at which time she will officially take ...",
  "genres": [
    "Action"
  ],
  "runtime": 199,
  "cast": [
    "Pearl White",
    "Crane Wilbur",
    "Paul Panzer",
    "Edward Josè"
  ],
  "num_mflix_comments": 0,
  "poster": "https://m.media-amazon.com/images/M/MV5BMzgxODk1Mzk2Ml5BMl5BanBnXkFtZTgwMDg0NzkwMjE@._V1_SY1000_SX677_AL_.jpg",
  "title": "The Perils of Pauline",
  "fullplot": "Young Pauline is left a lot of money when her wealthy uncle dies. However, her uncle's secretary has been named as her guardian until she marries, at which time she will officially take possession of her inheritance. Meanwhile, her \"guardian\" and his confederates constantly come up with schemes to get rid of Pauline so that he can get his hands on the money himself.",
  "languages": [
    "English"
  ],
  "released": {
    "$date": {
      "$numberLong": "-1760227200000"
    }
  },
  "directors": [
    "Louis J. Gasnier",
    "Donald MacKenzie"
  ],
  "writers": [
    "Charles W. Goddard (screenplay)",
    "Basil Dickey (screenplay)",
    "Charles W. Goddard (novel)",
    "George B. Seitz",
    "Bertram Millhauser"
  ],
  "awards": {
    "wins": 1,
    "nominations": 0,
    "text": "1 win."
  },
  "lastupdated": "2015-09-12 00:01:18.647000000",
  "year": 1914,
  "imdb": {
    "rating": 7.6,
    "votes": 744,
    "id": 4465
  },
  "countries": [
    "USA"
  ],
  "type": "movie",
  "tomatoes": {
    "viewer": {
      "rating": 2.8,
      "numReviews": 9
    },
    "production": "Pathè Frères",
    "lastUpdated": {
      "$date": "2015-09-11T17:46:19.000Z"
    }
  },
  "plot_embedding": [
    0.004042261,
    -0.01163094
  ]
}
```

### `PUT /record/movie`
This should patch/update the movie data in the database
#### request body
```json
{
  "_id": {
    "$oid": "573a1390f29313caabcd5293"
  },
  "plot": "Young Pauline is left a lot of money when her wealthy uncle dies. However, her uncle's secretary has been named as her guardian until she marries, at which time she will officially take ...",
  "genres": [
    "Action"
  ],
  "runtime": 199,
  "cast": [
    "Pearl White",
    "Crane Wilbur",
    "Paul Panzer",
    "Edward Josè"
  ],
  "num_mflix_comments": 0,
  "poster": "https://m.media-amazon.com/images/M/MV5BMzgxODk1Mzk2Ml5BMl5BanBnXkFtZTgwMDg0NzkwMjE@._V1_SY1000_SX677_AL_.jpg",
  "title": "The Perils of Pauline",
  "fullplot": "Young Pauline is left a lot of money when her wealthy uncle dies. However, her uncle's secretary has been named as her guardian until she marries, at which time she will officially take possession of her inheritance. Meanwhile, her \"guardian\" and his confederates constantly come up with schemes to get rid of Pauline so that he can get his hands on the money himself.",
  "languages": [
    "English"
  ],
  "released": {
    "$date": {
      "$numberLong": "-1760227200000"
    }
  },
  "directors": [
    "Louis J. Gasnier",
    "Donald MacKenzie"
  ],
  "writers": [
    "Charles W. Goddard (screenplay)",
    "Basil Dickey (screenplay)",
    "Charles W. Goddard (novel)",
    "George B. Seitz",
    "Bertram Millhauser"
  ],
  "awards": {
    "wins": 1,
    "nominations": 0,
    "text": "1 win."
  },
  "lastupdated": "2015-09-12 00:01:18.647000000",
  "year": 1914,
  "imdb": {
    "rating": 7.6,
    "votes": 744,
    "id": 4465
  },
  "countries": [
    "USA"
  ],
  "type": "movie",
  "tomatoes": {
    "viewer": {
      "rating": 2.8,
      "numReviews": 9
    },
    "production": "Pathè Frères",
    "lastUpdated": {
      "$date": "2015-09-11T17:46:19.000Z"
    }
  },
  "plot_embedding": [
    0.004042261,
    -0.01163094
  ]
}
```

### `DELETE /record/movie/{id}`
This should delete the movie data in the database for the given id
