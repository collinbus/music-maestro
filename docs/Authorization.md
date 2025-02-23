# Authorization

### Requirements

You need a MongoDB server thats running on `localhost:27017`<br>
Inside you'll need a database `musicMaestro` with collection `applicationData`. That collection must contain one document with `applicationName: musicMaestro`

### Spotify Authorization

If your application wants to use the Spotify api you can follow the [Spotify Authorization Code Flow](https://developer.spotify.com/documentation/general/guides/authorization-guide/#authorization-code-flow)<br>
The very first time, you'll need to register your app in order to request an access token. Once that's successful, copy the code
from the redirect url and initialize it in the application with following command:
> musicMaestro -accessCode CODE -clientId CLIENT_ID -clientSecret CLIENT_SECRET


#### Requesting an access token

Every hour, you will need to request a new access token. That token can be used in the authorization header for every service call that you make. Once it's expired, request a new one by following step two of the Authorization Code Flow.
