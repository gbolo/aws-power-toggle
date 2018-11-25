# Get Version Information

Retrieves the current running version of the backend server.
This includes release number, commit hash, and date of build.


**URL** : `/api/v1/version`

**Method** : `GET`

## Success Response

**Code** : `200 OK`

**Example Response Body**

```json
{
  "version": "1.1",
  "git_hash": "746c10e",
  "build_date": "2018-11-24T23:21:02-0500"
}
```

## Notes

The version information is passed in at compile time when using the make target: `make all`

if the proper `ldflags` are not passed in during build,
then the following default content will be returned:

```json
{
  "version": "devel",
  "git_hash": "unknown",
  "build_date": "unknown"
}
```
