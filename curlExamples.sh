curl -v --header "Content-Type: application/json" \
  --request POST \
  --data '{"username":"xyz","password":"xyz", "email":"xyz"}' \
  http://goserver.pl:8083/register

curl -v --header "Content-Type: application/json" \
  --request POST \
  -u xyz123:xyz123 \
  http://goserver.pl:8083/login
  
  

  curl -v --request POST -H 'Accept: application/json' -H "Authorization: Bearer a2e3663d-48a5-4564-645c-cc734b00899f"  http://goserver.pl:8083/logout