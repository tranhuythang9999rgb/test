//login

curl --location '127.0.0.1:8080/user/login' \
--form 'user_name="thangth1"' \
--form 'password="1234"'

//get list session in user name

curl --location '127.0.0.1:8080/user/list' \
--header 'Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MjE5NjcyNDUsInVzZXJfbmFtZSI6InRoYW5ndGgxIn0.NzfvQmW6Umw1ITMvaMyvy3t4kORTiFk_gv8mswzgvsk'

// logout

curl --location --request POST '127.0.0.1:8080/user/logout' \
--header 'Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MjE5NjcyNDUsInVzZXJfbmFtZSI6InRoYW5ndGgxIn0.NzfvQmW6Umw1ITMvaMyvy3t4kORTiFk_gv8mswzgvsk'

//register

curl --location '127.0.0.1:8080/user/register' \
--header 'Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MjE5NjcyNDUsInVzZXJfbmFtZSI6InRoYW5ndGgxIn0.NzfvQmW6Umw1ITMvaMyvy3t4kORTiFk_gv8mswzgvsk' \
--form 'user_name=""' \
--form 'display_name=""' \
--form 'password=""' \
--form 'avatar=""'