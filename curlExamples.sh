curl --header "Content-Type: application/json" \
  --request POST \
  --data '{"username":"xyz","password":"xyz", "email":"xyz"}' \
  http://goserver.pl:8083/register

curl --header "Content-Type: application/json" \
  --request POST \
  -u xyz123:xyz123 \
  http://goserver.pl:8083/login
  
  

  curl --request POST -H 'Accept: application/json' -H "Authorization: Bearer d20a874d-6a8e-4bf3-691c-bbc29dac3ef6"  http://goserver.pl:8083/logout