package main
import (
	"bytes"
	"encoding/json"
	"fmt"
	"time"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
	"tcs.com/proofverify"
)

var logger = shim.NewLogger("known-traveller")

type TravellerChaincode struct {
}

// UserDetails structure to store general details of the user
type UserDetails struct {
	ObjectType        string    `json:"docType"`
	UserID            string    `json:"userID"`
	UserAddress       string    `json:"userAddress"`
	UserPhotoLocation string    `json:"userPhotoLocation"`
	CreatedBy         string    `json:"createdBy"`
	CreatedDate       time.Time `json:"createdDate"`
	UpdatedBy         string    `json:"updatedBy"`
	UpdatedDate       time.Time `json:"updatedDate"`
}

// UserPassportDetails structure to store passport details of the user
type UserPassportDetails struct {
	ObjectType       string    `json:"docType"`
	UserID           string    `json:"userID"`
	PassportNumber   string    `json:"passportNumber"`
	PassportScanCopy string    `json:"passportScanCopy"`
	PassportPages    string    `json:"passportPages"`
	CreatedBy        string    `json:"createdBy"`
	CreatedDate      time.Time `json:"createdDate"`
	UpdatedBy        string    `json:"updatedBy"`
	UpdatedDate      time.Time `json:"updatedDate"`
}

// UserVisaDetails structure to store passport details of the user
type UserVisaDetails struct {
	ObjectType  string `json:"docType"`
	UserID      string `json:"userID"`
	VisaID      string `json:"visaID"`
	CountryCode string `json:"countryCode"`
	VisaNumber  string `json:"visaNumber"`
	Country     string `json:"country"`
	ValidatedBy string `json:"validatedBy"`
	ExpiryDate  string `json:"expiryDate"`
}

// UserFlightDetails structure to store passport details of the user
type UserFlightDetails struct {
	ObjectType        string    `json:"docType"`
	FlightID          int       `json:"flightID"`
	FlightNo          string    `json:"flightNo"`
	FlightName        string    `json:"flightName"`
	FlightType        string    `json:"flightType"`
	Class             string    `json:"class"`
	SeatsLeft         int       `json:"seatsLeft"`
	StartLocation     string    `json:"startLocation"`
	EndLocation       string    `json:"endLocation"`
	StartTime         time.Time `json:"startTime"`
	FlightDuration    int       `json:"flightDuration"`
	CreatedBy         string    `json:"createdBy"`
	CreatedDate       time.Time `json:"createdDate"`
	UpdatedBy         string    `json:"updatedBy"`
	Status            string    `json:"status"`
	DateOfTravelStart time.Time `json:"dateOfTravelStart"`
}

type vote struct {
	ObjectType 	string 	`json:"docType"`
	PollID		string 	`json:"pollID"`
	VoterID		string 	`json:"voterID"`
	VoterSex 	string 	`json:"voterSex"`
	VoterAge	int 	`json:"voterAge"`
}

type votePrivateDetails struct {
	ObjectType 	string 	`json:"docType"`
	PollID		string 	`json:"pollID"`
	VoterID		string 	`json:"voterID"`
	Salt 		string 	`json:"salt"`
	VoteHash 	string 	`json:"voteHash"`
}

// ==== date and time layouts ====
const layoutDate = "1/2/2006"
const layoutTime = "3:4:5 PM"

// Init of the chaincode
// This function is called only one when the chaincode is instantiated.
// So the goal is to prepare the ledger to handle future requests.
func (t *TravellerChaincode) Init(stub shim.ChaincodeStubInterface) pb.Response {
	fmt.Println("########### TravellerChaincode Init ###########")

	// Get the function and arguments from the request
	function, _ := stub.GetFunctionAndParameters()

	// Check if the request is the init function
	if function != "init" {
		return shim.Error("Unknown function call")
	}

	// Put in the ledger the key/value hello/world
	err := stub.PutState("hello", []byte("world"))
	if err != nil {
		return shim.Error(err.Error())
	}

	// Return a successful message
	return shim.Success(nil)
}

// Invoke
// All future requests named invoke will arrive here.
func (t *TravellerChaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	fmt.Println("########### TravellerChaincode Invoke ###########")

	fmt.Println("########### TravellerChaincode Init1 ###########")
	//////////////////////////

	indycreatorbytes, _ := stub.GetCreator()
	type IndyCreator struct {
		Did string
	}
	indycreator := &IndyCreator{}
	if err := json.Unmarshal(indycreatorbytes, &indycreator); err != nil {
		panic(err)
	}
	fmt.Println(indycreator)
	status, err := proofverify.VerifyVoterProof(indycreator.Did)
	if status == false || err != nil {
		return shim.Error("Proof verification failed : " + err.Error())
	}
	fmt.Println("proof verification success")

	// Get the function and arguments from the request
	function, args := stub.GetFunctionAndParameters()

	// Check whether it is an invoke request
	if function == "AddUserDetails" {
		return t.AddUserDetails(stub, args)
	}

	if function == "UpdateUserDetails" {
		return t.UpdateUserDetails(stub, args)
	}

	if function == "QueryUserDetails" {
		return t.QueryUserDetails(stub, args)
	}

	if function == "initVote" {
		return t.initVote(stub, args)
	}
	if function == "getVote" {
		return t.getVote(stub, args)
	}

	if function != "invoke" {
		return shim.Error("Unknown function call")
	}

	// Check whether the number of arguments is sufficient
	if len(args) < 1 {
		return shim.Error("The number of arguments is insufficient.")
	}

	// In order to manage multiple type of request, we will check the first argument.
	// Here we have one possible argument: query (every query request will read in the ledger without modification)
	if args[0] == "query" {
		return t.query(stub, args)
	}

	// The update argument will manage all update in the ledger
	if args[0] == "change" {
		return t.change(stub, args)
	}

	// If the arguments given don’t match any function, we return an error
	return shim.Error("Unknown action, check the first argument")
}

func (t *TravellerChaincode) AddUserDetails(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var err error

	if len(args) != 7 {
		return shim.Error("AddUserDetails: Incorrect number of arguments. Expecting 7")
	}

	userID := args[0]
	userAddress := args[1]
	userPhotoLocation := args[2]
	createdBy := args[3]
	createdDate, err1 := time.Parse(layoutDate, args[4])
	updatedBy := args[5]
	updatedDate, err1 := time.Parse(layoutDate, args[6])

	fmt.Println("Error", err1)
	// ==== Check if User details already exists ====
	userDetailsAsBytes, err := stub.GetState(userID)
	if err != nil {
		return shim.Error("AddUserDetails: Failed to get User details: " + err.Error())
	} else if userDetailsAsBytes != nil {
		logger.Errorf("AddUserDetails : User details already exists")
		return shim.Error("AddUserDetails: User details already exists")
	}

	newUserDetails := &UserDetails{"UserDetails", userID, userAddress, userPhotoLocation, createdBy, createdDate, updatedBy, updatedDate}
	newUserDetailsJSONasBytes, err := json.Marshal(newUserDetails)

	if err != nil {
		return shim.Error(err.Error())
	}

	// === Save user details to state ===
	err = stub.PutState(userID, newUserDetailsJSONasBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	err = stub.SetEvent("eventAddUser", []byte{})
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(nil)
}

// UpdateUserDetails updates general details of the user
func (t *TravellerChaincode) UpdateUserDetails(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var err error

	if len(args) != 7 {
		return shim.Error("UpdateUserDetails: Incorrect number of arguments. Expecting 7")
	}

	userID := args[0]
	userAddress := args[1]
	userPhotoLocation := args[2]
	createdBy := args[3]
	createdDate, err1 := time.Parse(layoutDate, args[4])
	updatedBy := args[5]
	updatedDate, err1 := time.Parse(layoutDate, args[6])
	fmt.Println("Error", err1)

	// ==== Check if User details already exists ====
	userDetailsAsBytes, err := stub.GetState(userID)
	if err != nil {
		return shim.Error("AddUserDetails: Failed to get User details: " + err.Error())
	} else if userDetailsAsBytes == nil {
		logger.Errorf("AddUserDetails : User details doesn't exist ")
		return shim.Error("AddUserDetails: User details doesn't exist")
	}

	newUserDetails := &UserDetails{"UserDetails", userID, userAddress, userPhotoLocation, createdBy, createdDate, updatedBy, updatedDate}
	newUserDetailsJSONasBytes, err := json.Marshal(newUserDetails)

	if err != nil {
		return shim.Error(err.Error())
	}

	// === Save user details to state ===
	err = stub.PutState(userID, newUserDetailsJSONasBytes)
	if err != nil {
		return shim.Error(err.Error())
	}
	err = stub.SetEvent("eventUpdateUser", []byte{})
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(nil)
}

// QueryUserDetails updates general details of the user
func (t *TravellerChaincode) QueryUserDetails(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	userID := args[0]
	//queryString := fmt.Sprintf("{\"selector\":{\"docType\":\"UserDetails\"}}")
	state, err := stub.GetState(userID)
	//queryResults, err := getQueryResultForQueryString(stub, queryString)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(state)
}

/* ======================================================================================================
* Function 		: GetQueryResultForQueryString(Internal Function)
* Description 	: Internal function to query public datasets
* Parameters	: Parm1 - stub; Parm2 - querystring
* Return 		: Dataset for the given query
* ====================================================================================================== */
func getQueryResultForQueryString(stub shim.ChaincodeStubInterface, queryString string) ([]byte, error) {

	logger.Info("- getQueryResultForQueryString queryString:\n%s\n", queryString)

	resultsIterator, err := stub.GetQueryResult(queryString)
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	// buffer is a JSON array containing QueryRecords
	var buffer bytes.Buffer
	buffer.WriteString("[")

	bArrayMemberAlreadyWritten := false
	for resultsIterator.HasNext() {
		//queryResultKey, queryResultRecord, err := resultsIterator.Next()
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}
		// Add a comma before array members, suppress it for the first array member
		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
		}

		buffer.WriteString(string(queryResponse.Value))

		bArrayMemberAlreadyWritten = true
	}
	buffer.WriteString("]")

	logger.Info("- getQueryResultForQueryString queryResult:\n%s\n", buffer.String())

	return buffer.Bytes(), nil
}

// query
// Every readonly functions in the ledger will be here
func (t *TravellerChaincode) query(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	fmt.Println("########### TravellerChaincode query ###########")

	// Check whether the number of arguments is sufficient
	if len(args) < 2 {
		return shim.Error("The number of arguments is insufficient.")
	}

	// Like the Invoke function, we manage multiple type of query requests with the second argument.
	// We also have only one possible argument: hello
	if args[1] == "hello" {

		// Get the state of the value matching the key hello in the ledger
		state, err := stub.GetState("hello")
		if err != nil {
			return shim.Error("Failed to get state of hello")
		}

		// Return this value in response
		return shim.Success(state)
	}

	// If the arguments given don’t match any function, we return an error
	return shim.Error("Unknown query action, check the second argument.")
}

// invoke
// Every functions that read and write in the ledger will be here
func (t *TravellerChaincode) change(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	fmt.Println("########### TravellerChaincode invoke ###########")

	if len(args) < 2 {
		return shim.Error("The number of arguments is insufficient.")
	}
	transMap, err := stub.GetTransient()
	if err != nil {
		return shim.Error("Error getting transient: " + err.Error())
	}
	fmt.Println("Transient", transMap)

	// Check if the ledger key is "hello" and process if it is the case. Otherwise it returns an error.
	if args[1] == "hello" && len(args) == 3 {

		fmt.Println("Transient inside ", transMap)
		

		// Write the new value in the ledger
		err := stub.PutState("hello", []byte(args[2]))
		if err != nil {
			return shim.Error("Failed to update state of hello")
		}

		// Notify listeners that an event "eventInvoke" have been executed (check line 19 in the file invoke.go)
		err = stub.SetEvent("eventInvoke", []byte{})
		if err != nil {
			return shim.Error(err.Error())
		}

		// Return this value in response
		return shim.Success(nil)
	}

	// If the arguments given don’t match any function, we return an error
	return shim.Error("Unknown invoke action, check the second argument.")
}



// ============================================================
// initVote - create a new vote and store into chaincode state
// ============================================================
func (t *TravellerChaincode) initVote(stub shim.ChaincodeStubInterface, args []string) pb.Response {


	type voteTransientInput struct {
		PollID		string 	`json:"pollID"`
		VoterID		string 	`json:"voterID"`
		VoterSex 	string 	`json:"voterSex"`
		VoterAge	int 	`json:"voterAge"`
		Salt 		string 	`json:"salt"`
		VoteHash 	string 	`json:"voteHash"`
	}

	fmt.Println("- start init vote")


	if len(args) != 0 {
		return shim.Error("Private data should be passed in transient map.")
	}

	transMap, err := stub.GetTransient()
	if err != nil {
		return shim.Error("Error getting transient: " + err.Error())
	}

	voteJsonBytes, success := transMap["vote"]
	if !success {
		return shim.Error("vote must be a key in the transient map")
	}

	if len(voteJsonBytes) == 0 {
		return shim.Error("vote value in transient map cannot be empty JSON string")
	}

	var voteInput voteTransientInput
	err = json.Unmarshal(voteJsonBytes, &voteInput)
	if err != nil {
		return shim.Error("failed to decode JSON of: " + string(voteJsonBytes))
	}

	// input sanitation

	if len(voteInput.PollID) == 0 {
		return shim.Error("poll ID field must be a non-empty string")
	} 

	if len(voteInput.VoterID) == 0 {
		return shim.Error("voter ID field must be a non-empty string")
	} 

	if voteInput.VoterAge <= 0 {
		return shim.Error("age field must be > 0")
	}

	if len(voteInput.VoterSex) == 0 {
		return shim.Error("sex field must be a non-empty string")
	} 

	if len(voteInput.Salt) == 0 {
		return shim.Error("salt must be > 0")
	}

	if len(voteInput.VoteHash) == 0 {
		return shim.Error("vote hash field must be a non-empty string")
	} 

	fmt.Println("Pollid vote ip",voteInput.PollID , voteInput.VoterID)
	existingVoteAsBytes, err := stub.GetPrivateData("collectionVote", voteInput.PollID + voteInput.VoterID)
	if err != nil {
		return shim.Error("Failed to get vote: " + err.Error())
	} else if existingVoteAsBytes != nil {
		fmt.Println("This vote already exists: " + voteInput.PollID + voteInput.VoterID)
		return shim.Error("This vote already exists: " + voteInput.PollID + voteInput.VoterID)
	}

	vote := &vote{
		ObjectType: "vote",
		PollID: voteInput.PollID,
		VoterID: voteInput.VoterID,
		VoterAge: voteInput.VoterAge,
		VoterSex: voteInput.VoterSex,
	}
	voteJSONasBytes, err := json.Marshal(vote)
	if err != nil {
		return shim.Error(err.Error())
	}

	err = stub.PutPrivateData("collectionVote", voteInput.PollID + voteInput.VoterID, voteJSONasBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	votePrivateDetails := &votePrivateDetails {
		ObjectType: "votePrivateDetails",
		PollID: voteInput.PollID,
		VoterID: voteInput.VoterID,
		Salt: voteInput.Salt,
		VoteHash: voteInput.VoteHash,
	}
	votePrivateDetailsBytes, err := json.Marshal(votePrivateDetails)
	if err != nil {
		return shim.Error(err.Error())
	}

	err = stub.PutPrivateData(
		"collectionVotePrivateDetails", 
		voteInput.PollID + voteInput.VoterID + voteInput.Salt, 
		votePrivateDetailsBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	//register event
	err = stub.SetEvent("eventInitVote", []byte{})
	if err != nil {
		return shim.Error(err.Error())
	}

	fmt.Println("- end init vote (success)")
	return shim.Success([]byte(voteInput.Salt))
}

// =====================================================
// getVote - retrieve vote metadata from chaincode state
// =====================================================

func (t *TravellerChaincode) getVote(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting vote key to query")
	}

	voteKey := args[0]
	// ==== retrieve the vote ====
	voteAsBytes, err := stub.GetPrivateData("collectionVote", voteKey)
	if err != nil {
		return shim.Error("{\"Error\":\"Failed to get state for " + voteKey + "\"}")
	} else if voteAsBytes == nil {
		return shim.Error("{\"Error\":\"Vote does not exist: " + voteKey + "\"}")
	}

	return shim.Success(voteAsBytes)
}



func main() {
	// Start the chaincode and make it ready for futures requests
	err := shim.Start(new(TravellerChaincode))
	if err != nil {
		fmt.Printf("Error starting Heroes Service chaincode: %s", err)
	}
}
