
var express = require ('express');
var app= express();
var bodyParser = require('body-parser');
var mysql  = require('mysql');
var dbconn = mysql.createConnection({
  host     : 'localhost',
  user     : 'root',
  password : '',
  database : 'blockactivity'
});


var hfc = require('fabric-client');
var path = require('path');
var util = require('util');
exports.queryblockSDK = function (fnName,par_txnId,reply,blockjson) {
    var func_name = fnName;
    
    console.log ("Function Name Rcvd:"+func_name + " transaction id received to block activity:"+par_txnId);
   /** 
    if(func_name == "getVehicle_CustomerDetails"){
          var arg1_f3 = request.params.arg1;
    }
    if(func_name == "getAll_VehicleStatus_Pending_PR"){
       // var arg1_f3 = request.params.arg1;
    }
    if(func_name == "getVehicleDetails_EngineNo"){
        var arg1_f4 = request.params.arg1;
    }
    if(func_name == "getOwnershipTransferDetails"){
        var arg1_f5 = request.params.arg1;
    }
    if(func_name == "getReceiverDetails"){
        var arg1_f6 = request.params.arg1;
    }
   */ 
    var options = {
        wallet_path: path.join(__dirname, './creds'),
        user_id: 'PeerAdmin',
        channel_id: 'mychannel',
        chaincode_id: 'loanOrginatSystem',
        network_url: 'grpc://192.168.99.100:7051',
    };
    var channel = {};
    var client = null;
    Promise.resolve().then(() => {
       console.log("Create a client and set the wallet location");
    client = new hfc();
    return hfc.newDefaultKeyValueStore({ path: options.wallet_path });
    }).then((wallet) => {
        console.log("Set wallet path, and associate user ", options.user_id, " with application");
    client.setStateStore(wallet);
    return client.getUserContext(options.user_id, true);
    }).then((user) => {
        console.log("Check user is enrolled, and set a query URL in the network");
    if (user === undefined || user.isEnrolled() === false) {
        console.error("User not defined, or not enrolled - error");
    }
    channel = client.newChannel(options.channel_id);
    channel.addPeer(client.newPeer(options.network_url));
    return;
    }).then(() => {
    console.log("Query Function to be called....");
    var transaction_id = client.newTransactionID();
    console.log("Assigning transaction_idzzzz: ", transaction_id._transaction_id);
    	
/*    if(func_name == "getAll_VehicleStatus_Pending_PR"){
         const request = {
                chaincodeId: options.chaincode_id,
                txId: transaction_id,
                fcn: func_name,
              //  args: [arg1_f1]
                args: []
            };
             console.log("Calling One getAll_VehicleStatus_Pending_PR Function..");
            return channel.queryByChaincode(request);
    } */
/*
    if(func_name == "getAllSalesOrderDetails"){
    	const request = {
    			chaincodeId: options.chaincode_id,
    			txId: tra
    			nsaction_id,
    			fcn: func_name,
    			//  args: [arg1_f1]
    			args: []
        };
        console.log("Calling One getAllSalesOrderDetails Function..");
        return channel.queryByChaincode(request);
    }
    
    */
    
    if(func_name === "trackblockactivity"){
        console.log('inside trackblockactivity');
    	//add this function to query blockchain transaction jax
        return channel.queryTransaction(par_txnId)

        // return channel.queryTransaction(tx_id.getTransactionID());
        //return channel.BlockchainInfo;
   	
      // var arg1_f2 = request.params.arg1;
   }
  
    
    
 
	
    }).then((query_responses) => {
//	console.log("TYPE:"+query_responses.transactionEnvelope.payload.header.channel_header.type);
//	console.log("VERSION:"+query_responses.transactionEnvelope.payload.header.channel_header.version);
	//console.log("TIMESTAMP:"+query_responses.transactionEnvelope.payload.header.channel_header.timestamp);
	//console.log("CHANNEL ID:"+query_responses.transactionEnvelope.payload.header.channel_header.channel_id);
	//console.log("TXN ID:"+query_responses.transactionEnvelope.payload.header.channel_header.tx_id);
	//console.log("EPOCH :"+query_responses.transactionEnvelope.payload.data.actions[0].payload.action.proposal_response_payload.extension.results.ns_rwset[1].rwset.writes[0].value );
	
	//console.log("BlockNUMsdsf"+query_responses.transactionEnvelope.payload.data.actions[0].payload.chaincode_proposal_payload.input);
	
	
	
	//var arr_result = JSON.parse(query_responses);
	//console.log("arr values:"+arr_result);
	
	var txnid = query_responses.transactionEnvelope.payload.header.channel_header.tx_id;
	var version = query_responses.transactionEnvelope.payload.header.channel_header.version;
	var timestamp = query_responses.transactionEnvelope.payload.header.channel_header.timestamp;
	var txn_type  = query_responses.transactionEnvelope.payload.header.channel_header.type;
	var payload =  blockjson;
	//var payload =  query_responses.transactionEnvelope.payload.data.actions[0].payload.action.proposal_response_payload.extension.results.ns_rwset[1].rwset.writes[0].value;
	//var db_payload = payload.salesID;
	//console.log("Sales id :jax : "+payload);
	//var stringify = JSON.stringify(query_responses)
	//content = JSON.parse(stringify);  
	//console.log(content);
	//console.log("payload: "+payload)

	// alternative shortcut
	console.log(util.inspect(query_responses, false, null))
	//console.log("BlockNUMsdsf"+query_responses.transactionEnvelope.payload.data.actions[0].payload.chaincode_proposal_payload.input);
	  //chcking starts
	 // console.log("BlockNUM"+query_responses.transactionEnvelope.payload.data.actions.payload.action[0].proposal_response_payload.proposal_hash);

	// console.log("payload1"+query_responses.transactionEnvelope.payload.data.actions.payload.action.proposal_response_payload.extension.events.payload);

	// console.log("payload2"+query_responses.transactionEnvelope.payload.data.actions.payload.action.proposal_response_payload.extension.response.payload);
	
	
	
	
	
	//console.log("signature_header creator Mspid :"+query_responses.transactionEnvelope.payload.header.creator.Mspid);
//	console.log("signature_header creator IdBytes  :"+query_responses.transactionEnvelope.payload.header.channel_header.creator.IdBytes);
	//console.log("signature_header nonce   :"+query_responses.transactionEnvelope.payload.header.channel_header.nonce );
	
	//var obj = JSON.parse(query_responses[0],validationCode);
	//console.log(typeof obj);
	//reply.send(typeof obj);
	//db connection starts
//	dbconn.connect(function(err){
	 // if(err){
	 //   console.log('Database connection error'+err);
	 // }
	  //else{
	 //   console.log('Database connection successful');
		
		var urlencodedParser = bodyParser.urlencoded({ extended: false })

	app.use(express.static('public'));

	
	//var userList = {name:response.userName,mailId:response.mailId,password:response.password,contactNo:response.mbno,city:response.city,country:response.nation}
		var db_list = {type:"INVOKE",txn_id: txnid,chaincode_id : "loanOrginatSystem",payload: payload,version : version,bc_timestamp : timestamp,txn_type : txn_type, deployments: "0" ,invocations:"1" }
	    dbconn.query('insert into blocks set?',db_list);

	   console.log("done Insertion");


		
		
	 // }
	//});
	//db con ends
	
	if (!query_responses.length) {
		console.log("No payloads were returned from query");
    } else {
    	console.log("Query result count = ", query_responses.length)
    }
    if (query_responses[0] instanceof Error) {
    	console.error("error from query = ", query_responses[0]);
    }
  //  console.log("Response is ", query_responses[0].toString());
   // reply.send(util.format(query_responses[0].toString('utf8')));
    }).catch((err) => {
    	console.error("Caught Error", err);
    	reply.send ("Caught Error", err);
    });
}
