# Weather API
Provide Forecast for a given Location

## **Tools used**
* Web Framework: Echo
* Config file: Viper


# API : Provide Weather Forecast for a given Lat, Lons

# API Endpoints

# GET
_____________________________________________
### /healthcheck  -- Provides the Health information of the service


------------------------------------------------------------
### api/v1/weather/forecast  --  provides Weather forecast information 

### Example 
 #### Query Params
 * lat : 44.1
 * lon : -82.929611

#### Example Response
````

{
    "weather_condition": "overcast clouds",
    "temperature": "42.44 F",
    "weather_type": "cold",
    "visibility": "10 Miles",
    "wind_speed": "21.56 miles/sec",
    "cloud_coverage": "99 %",
    "sunrise": "2024-02-26T06:34:47-06:00",
    "sunset": "2024-02-26T17:34:56-06:00"
}
````


### Assumptions:
 ````
   1. Metrics : imperial (farenhiet, miles )
   2. temperature < 60 --> Cold
      temperature 60 to 85 --> moderate
      temperature > 85 --> hot
 ````

 ### TODO
  #### Add Swagger , Unit Test cases , More Logs for dashboard
   
 