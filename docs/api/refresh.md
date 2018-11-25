# Refresh Environment Data

Forces a refresh of the cached environment data.
Requests to this endpoint triggers an API call to the AWS API.


**URL** : `/api/v1/refresh`

**Method** : `POST`

## Success Response

**Code** : `200 OK`

**Example Response Body**

```json
{
  "status": "OK"
}
```

## Notes

The refresh is **NOT** environment specific, and affects the **entire cache.**


**The backend has an internal configurable timer that
also refreshes the cache periodically.**
