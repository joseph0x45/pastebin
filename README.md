# pastebin
Personal self hosted pastebin

Create pastes by just providing a title and pasting the content

<img width="1865" height="1016" alt="image" src="https://github.com/user-attachments/assets/44dee493-7c72-4b76-a05a-69a24cf2ba8b" />

Pastes are persisted in the database.

<img width="1865" height="1016" alt="image" src="https://github.com/user-attachments/assets/bfec712d-b27f-4829-a469-0e5cf2778a7c" />

## Customizing port
By default Pastebin starts on port 8080. You can change this by providing the `-port` flag like so
```sh
pastebin -port=6969
```

## Customizing SQLite database used
By default Pastebin creates a file named `pastebin.db` in the directory where it is launched. You can change this by providing the `-db` flag like so
```sh
pastebin -db=/path/to/my/db
```

## Pastebin also exposes a HTTP interface

### Creating a paste
Curl example 
```sh
curl -X POST -d '{"title":"super cool article on hackaday", "content":"https://hackaday.com/2025/09/29/mini-laptop-needs-custom-kernel/#more-834825"}' http://localhost:8080/api/pastes
```
Returns:
- `HTTP 201` if the paste was created.
- `HTTP 400` if one of the fields is missing or empty
- `HTTP 500` if there was a server error. Consult logs to see what went wrong

### Getting all pastes
Curl example
```sh
curl http://localhost:8080/api/pastes
```
Returns the list of pastes under this structure
```json
{
  "pastes": [
    {
      "id": "PH6lC8g2b",
      "title": "Some link I can't send to my other device directly",
      "content": "https://electricfiredesign.com/2021/02/05/leds-for-light-art-part-2-optics/"
    }
  ]
}
```

### Getting a paste with it's id
Curl example
```sh
curl http://localhost:8080/api/pastes/<paste_id>
```
Returns the paste under this structure
```json
{
  "paste": {
    "id": "PH6lC8g2b",
    "title": "Some link I can't send to my other device directly",
    "content": "https://electricfiredesign.com/2021/02/05/leds-for-light-art-part-2-optics/"
  }
}
```
Returns:
- `HTTP 200` if the paste was found.
- `HTTP 404` if the paste was not found
- `HTTP 500` if there was a server error. Consult logs to see what went wrong

### Deleting a paste
Curl example
```sh
curl -X DELETE http://localhost:8080/api/pastes/<paste_id>
```
Returns:
- `HTTP 200` if the paste was deleted (or wasn't found)
- `HTTP 500` if there was a server error. Consult logs to see what went wrong
