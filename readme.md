# jellyfin-notifications
This is really just a thin proxy to the jellyfin endpoint:
 - `http://<jellyfin>:<port>/Items/<library>/Refresh`

## What is does
This code just renames the header `X-Mediabrowser-Token` to `Mediabrowser Token` and calls Jellyfins refresh endpoint

## Why
- Isolate from more Emby to Jellyfin divergences.
- Support tools which only support Emby to call Jellyfin
- fun. (this all could of been done using a reverse proxy header rewrite)

## Config
*Flags*
```bash
server  = flag.String("jellyfin", "http://192.168.1.1", "Jellyfin Server URL. eg http://192.168.1.1:8096")
library = flag.String("library", "ajd4b4rb4htb4h5bt4bt45bt", "Jellyfin library eg 767bffe4f11c93ef34b805451a696a4e")
```

*Example Docker Compose*

[here](./docker-compose.yaml)