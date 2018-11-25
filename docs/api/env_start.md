# Initialize a Start-Up of an Environment

Sends a startup request to the AWS API for all associated instances
in the specified environment

**URL** : `/api/v1/env/{env-id}/start`

**Method** : `POST`

## Success Response

**Code** : `200 OK`

## Notes

Responses vary depending on upstream AWS API
since we pass the exact aws response.
**This will be addressed in an upcomming release for consistency.**

**ONLY** instances which are in a `stopped` state get included for the AWS API call.