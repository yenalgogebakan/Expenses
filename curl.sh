
curl -X GET "localhost:8085/testing?name=appleboy&address=xyz"


curl -X GET localhost:8085/testing --data '{"name":"JJ", "address":"xyz"}' -H "Content-Type:application/json"

curl -d '{"option": "value", "something": "anothervalue"}' -H "Content-Type: application/json" -X POST https://flaviocopes.com/
curl -d "@my-file.json" -X POST https://flaviocopes.com/ ---Send JSON file
curl -o file.html https://flaviocopes.com/ ---Save the response to a file

