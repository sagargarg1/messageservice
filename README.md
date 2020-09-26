# messageservice
a small service to add/delete/update/messages

API:
The small message service exposes 6 API's
1.To add a message
  POST with api /messageservice/v1/message
  Json Body
  {
	  "message": "Happy Birthday" 
  }
  The response to this api will be an id which can be used to perform further operations
  
2.To get a message based on ID
  GET with api /messageservice/v1/message/1
  The response will give the message for this id=1
  
3.To delete a message based on ID
  DELETE with api /messageservice/v1/message/1
  This will delete the message with id=1
  
4.To update a message
  PUT with api /messageservice/v1/message
  Json Body
  {
    "id": 1,
	  "message": "Happy Birthday" 
  }
  This will update the message with id=1
  
5.To get all the messages
  GET with api /messageservice/v1/message/all
  The response will give all message in the DB
  
6.To get metrics
  GET with api /messageservice/v1/message/metrics
  The response will give metrics like
  "Number": How many messages in the system
	"BirthdayMessages": How many messages are birthday messages
	"SorryMessages": How many are sorry messages
	"GoodMorningMessages": How many are good morning messages
	"PalindromeMessages": How many messages are palindromes
  
  More metrics can be added as per the requriement . More categories can be added as well
  
  
Test:
  Unit test has been added to test the functionality of the handlers and DB . More testing capability can be added based on whether a real DB is being used. 
  The code is structures in a wasy which makes the testing quite simple . go_mock can be used to enhance the capability if more handlers and data layers are there.
  Testing can be done using "go test"
  
Logging:
  Logging capability has been added to log every request .From a design perspectve all the logging is being done from handlers and not from DB interface.
  "github.com/hashicorp/go-hclog" is used as a logger
 
Metrics:
  To collect the metrics from single point a middleware has been added which keeps the application clean . More middlewares can be added based on oAuth and RBAC requirements.
  
Code Structure:
  Code structuring is quite granular with
  handlers :- Containing the main handlers which is further segregated to message and metrics
  data:- For storing and performing DB operations
  model:- To keep data structures
  utils:- To perform utility operations
  middleware:- Specially for mtrics
  
 Deployement:
  The application is containarized and can be deployed either on cloud or on prem
  
  On Premise :-
    1.Take a centOS VM
    2.Login as root
    3.Install docker,git
    4.Clone the git repository 
    5.docker build -t main .
    
    OR
    
    1.Take a centOS VM
    2.Loging as root
    3.Install docker,git,golang
    4.Create a dir src at GOPATH
    5.Copy the code
    6.go build main.go
    7."./main"
    
    The functionality can be tested using curl or postman
    
  On Cloud(Azure):-
    1.Create a reource group 
    2.Create a centOS VM
    3.Repeat above steps
    
    Using AKS:
      1. Create a POD definition
      2. Cretae a replica set for this POD
      3. Crete a node port service
