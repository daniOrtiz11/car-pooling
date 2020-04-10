# Table booking API using Go


The purpose of this API is to provide table-occupancy service as people arrive at the restaurant.


In this service we assume that groups of diners arrive at the restaurant and have to be assigned a table. On the other hand, the tables have a number of seats available. So if a table has enough seats available it could accommodate a group of customers.

People requests tables in groups of 1 to 6. People in the same group want to eat on the same table. You can take any group at any table that has enough empty seats for them. If it's not possible to accommodate them, they're willing to wait until 
there's a table available for them. Once a table is available for a group that is waiting, they should eat. 

Once they get a table assigned, they will eat until they leave the restaurant, you cannot ask them to sit at another table (i.e. you cannot swap them to another table to make space for another group). In terms of fairness: groups are served in the order they arrive, but they eat opportunistically.

For example: a group of 6 is waiting for a table and there are 4 empty seats at a table for 6; if a group of 2 requests a table you may take them in the table for 6 but only if you have nowhere else to make them eat. This may mean that the group of 6 waits a long time, possibly until they become frustrated and leave the restaurant.

## API

This API must comply with the following contract:

### GET /status

Indicate the service has started up correctly and is ready to accept requests.

Responses:

* **200 OK** When the service is ready to receive requests.

### PUT /tables

Load the list of available tables in the service and remove all previous data (existing bookings and tables). This method may be called more than once during the life cycle of the service.

**Body** _required_ The list of tables to load.

**Content Type** `application/json`

Sample:

```json
[
  {
    "id": 1,
    "seats": 4
  },
  {
    "id": 2,
    "seats": 6
  }
]
```

Responses:

* **200 OK** When the list is registered correctly.
* **400 Bad Request** When there is a failure in the request format, expected headers, or the payload can't be unmarshalled.

### POST /booking

A group of people requests to perform a booking.

**Body** _required_ The group of people that wants to perform the booking

**Content Type** `application/json`

Sample:

```json
{
  "id": 1,
  "people": 4
}
```

Responses:

* **200 OK** or **202 Accepted** When the group is registered correctly
* **400 Bad Request** When there is a failure in the request format or the payload can't be unmarshalled.

### POST /bill

A group of people ask for the restaurant bill to leave.

**Body** _required_ A form with the group ID, such that `ID=X`

**Content Type** `application/x-www-form-urlencoded`

Responses:

* **200 OK** or **204 No Content** When the group is unregistered correctly.
* **404 Not Found** When the group is not to be found.
* **400 Bad Request** When there is a failure in the request format or the payload can't be unmarshalled.

### POST /locate

Given a group ID such that `ID=X`, return the table the group is eating with, or no table if they are still waiting to be seated.

**Body** _required_ A url encoded form with the group ID such that `ID=X`

**Content Type** `application/x-www-form-urlencoded`

**Accept** `application/json`

Responses:

* **200 OK** With the table as the payload when the group is assigned to a table.
* **204 No Content** When the group is waiting to be assigned to a table.
* **404 Not Found** When the group is not to be found.
* **400 Bad Request** When there is a failure in the request format or the payload can't be unmarshalled.

## Dependencies

Work in progress

## Deployment

Work in progress

### Local environment

Work in progress

### Production environment
Work in progress