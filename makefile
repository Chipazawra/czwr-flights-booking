CLIENT_ID=68ipF36GBniX4THOJBYovQ564gG86cjj
CLIENT_SECRET=JBVIKTlKjcWCBAPy
TOKEN=dFg0trGA5xwVr58wq1uOEfBeR6VF

gettoken: 
	curl "https://test.api.amadeus.com/v1/security/oauth2/token" \
		-H "Content-Type: application/x-www-form-urlencoded" \
   		-d "grant_type=client_credentials&client_id=$(CLIENT_ID)&client_secret=$(CLIENT_SECRET)"

request:
	curl -X GET "https://test.api.amadeus.com/v2/reference-data/urls/checkin-links?airlineCode=BA" \
		-H "Authorization: Bearer $(TOKEN)"