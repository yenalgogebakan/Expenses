
#curl -X GET "localhost:8085/testing?name=appleboy&address=xyz"
#curl -X GET localhost:8085/testing --data '{"name":"JJ", "address":"xyz"}' -H "Content-Type:application/json"
#curl -d '{"Name": "Ali", "Surname": "Deneme"}' -H "Content-Type: application/json" -X PUT localhost:8085/v1/users
#curl -d "@my-file.json" -X POST https://flaviocopes.com/ ---Send JSON file
#curl -o file.html https://flaviocopes.com/ ---Save the response to a file

#curl -d '{"name":"Ali","surname":"Deneme","email":"yenal.gogebakan@cybersoft.com.tr", "info":"cyberpark kat 2", "gender":"E"}' -H "Content-Type: application/json" -X PUT localhost:8085/v1/users
#curl -d '{"name":"YENAL","surname":"GOGEBAKAN","email":"yenal.gogebakan@cybersoft.com.tr", "info":"cyberpark Teper", "gender":"E"}' -H "Content-Type: application/json" -X PUT localhost:8085/v1/users
#curl -d '{"Name": "Ali", "Surname": "Deneme"}' -H "Content-Type: application/json" -X PUT localhost:8085/v1/users
#curl -d '{"XXX": "Ali", "Surname": "Deneme"}' -H "Content-Type: application/json" -X PUT localhost:8085/v1/users
#curl  -X GET localhost:8085/v1/users/Yenal
#curl  -X GET localhost:8085/v1/users/YENAL
curl -d '{"servicename": "ListUsers"}' -H "Content-Type: application/json" -X POST localhost:8085/v1/services
