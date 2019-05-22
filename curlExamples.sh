curl -v --header "Content-Type: application/json" \
  --request POST \
  --data '{"username":"xyz","password":"xyz", "email":"xyz"}' \
  http://goserver.pl:8083/register

curl -v --header "Content-Type: application/json" \
  --request POST \
  -u xyz123:xyz123 \
  http://goserver.pl:8083/login
  
  

  curl -v --request POST -H 'Accept: application/json' -H "Authorization: Bearer dc55e2b5-c520-4640-5101-38e250beba5d"  http://goserver.pl:8083/logout