# Authorization

### Spotify Authorization

If your application wants to use the Spotify api you can follow the [Spotify Authorization Code Flow](https://developer.spotify.com/documentation/general/guides/authorization-guide/#authorization-code-flow)<br>
The very first time, you'll need to register your app in order to request an access token. Once that's successful, copy the code
from the redirect url and initialize it in the application with following command:
> musicMaestro -code CODE


#### Requesting an access token

Every hour, you will need to request a new access token. That token can be used in the authorization header for every service call that you make. Once it's expired, request a new one by following step two of the Authorization Code Flow.
