#New Voting Enrollment Voting
echo "POST request Enroll on Org1 Voting  ..."
curl -s -X POST \
  http://localhost:3000/application \
  -H "content-type: application/x-www-form-urlencoded" \
  -d "{\"Name\": \"Voting\",\"Secret\": \"kerapwd\",\"Type\": \"user\"}"

echo "POST completed"

# New commerce Enrollment 
echo "POST request Enroll on Org1 commerce  ..."
curl -s -X POST \
  http://localhost:3000/application \
  -H "content-type: application/x-www-form-urlencoded" \
  -d "{\"Name\": \"commerce\",\"Secret\": \"yalepwd\",\"Type\": \"user\"}"

echo "POST completed"


echo "Query value of key hello"
#query function in Fabric SDK Go
curl -s -X GET \
  "http://localhost:3000/query?key=hello"

echo
echo "Invoke change function in chaincode"
#Invoke a new transaction
curl -d "fcn=change&key=hello&value=newval"  http://localhost:3000/invoke


#Add User Details
#curl -d "username=Voting&userID=1001&userAddress=Kochi&userPhotoLocation=Ernakulam&createdBy=abc&createdDate=12/11/2020&updatedBy=xyz&updatedDate=23/11/2020"  http://localhost:3000/adduser


#Update User Details
#curl -d "username=Voting&userID=1001&userAddress=NKochi&userPhotoLocation=NErnakulam&createdBy=Nabc&createdDate=13/11/2020&updatedBy=xyz&updatedDate=23/11/2020"  http://localhost:3000/updateuser


#query function in Fabric SDK Go
#curl -s -X GET  "http://localhost:3000/userquery?username=Voting&userID=1001"


#Initialize vote private data
#curl -d "username=Voting&PollID=P101&VoterID=V201&VoterSex=Male&VoterAge=38&Salt=Salt&VoteHash=VoteHash" http://localhost:3000/initVote

#Get private data vote
#curl -s -X GET  "http://localhost:3000/getVote?username=Voting&pollID=P101&voterID=V201"

# Get history by Transaction Id
#curl --location --request GET 'http://localhost:3000/history?username=Voting&query=txn&txnid=


