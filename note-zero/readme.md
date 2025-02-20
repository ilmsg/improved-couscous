
**get note**

    $ curl --request GET \
    --url http://localhost:3060/api/v1/notes

**create note**

    $ curl --request POST \
    --url http://localhost:3060/api/v1/notes \
    --header 'Content-Type: application/json' \
    --data '{
        "note": "read book"
    }'

**get note by id**

    $ curl --request GET \
    --url http://localhost:3060/api/v1/notes/92509906-455c-42ad-9df8-da779555fc86

**update note**

    $ curl --request PATCH \
    --url http://localhost:3060/api/v1/notes/92509906-455c-42ad-9df8-da779555fc86 \
    --header 'Content-Type: application/json' \
    --header 'User-Agent: insomnia/10.3.1' \
    --data '{
        "note": "play games"
    }'

**delete note**

    $ curl --request DELETE \
    --url http://localhost:3060/api/v1/notes/a2fe0dd5-38a8-4b41-a62c-7d7a8642b90f \
