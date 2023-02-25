# clay
A Go reverse proxy

## how it works
- The proxy listens on port 3000
- When a get request is received on the path `/forward`, it will look for a URL in the `target` header and try to proxy the request

idk why i made this i was bored