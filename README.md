## fetchapi

This code has two API calls, a POST request and a GET Request. 

- Submit the receipt details through the POST request to receive a unique ID. <br>
- This ID can be sent to the GET request to get the number of points the user received. 


## Code structure. 

### Models has all the relevant structs for the application     <br>   

- item is the structure of an item in the receipt.   
- Receipt is the structure of the receipt. <br>
- points.go has a map to save the ID with the corresponding points. <br>
- PostResponse is the structure of the response of POST Request, similarly GetResponse.  <br>

### Handlers. 

- receipthandler takes the receipt data to calculate points, generate ID and save.  <br>
- pointshandler takes the ID to respond with number of points received      
