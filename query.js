var hfc = require('fabric-client');
var path = require('path');
var util = require('util');
exports.querySDK = function (fnName,request,reply) {
    var func_name = fnName;
    console.log ("Function Name Rcvd:"+func_name);
	
    if(func_name == "GetDetailsByProspectId"){
        var arg1_f2 = request.params.arg1;
    }
    if(func_name == "ViewProspect"){
          //var arg1_f3 = request.params.arg1;
    }
    if(func_name == "GetDetailsByApplicantId"){
        var arg1_f3 = request.params.arg1;
    }
    if(func_name == "ViewApplication"){
        //var arg1_f4 = request.params.arg1;
    }
	if(func_name == "ViewProperties"){
        //var arg1_f4 = request.params.arg1;
    }
    if(func_name == "GetDetailsByPropertyId"){
        var arg1_f5 = request.params.arg1;
    }
    if(func_name == "getReceiverDetails"){
        var arg1_f6 = request.params.arg1;
    }
    
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
    console.log("Assigning transaction_id: ", transaction_id._transaction_id);

    if(func_name == "GetDetailsByProspectId"){
         const request = {
                chaincodeId: options.chaincode_id,
                txId: transaction_id,
                fcn: func_name,
                args: [arg1_f2]
                //args: []
            };
             console.log("Calling One GetDetailsByProspectId Function..");
            return channel.queryByChaincode(request);
    }

    if(func_name == "ViewProspect"){
    	const request = {
    			chaincodeId: options.chaincode_id,
    			txId: transaction_id,
    			fcn: func_name,
    			//  args: [arg1_f1]
    			args: []
        };
        console.log("Calling One ViewProspect Function..");
        return channel.queryByChaincode(request);
    }
    if(func_name == "GetDetailsByApplicantId"){
       const request = {
        chaincodeId: options.chaincode_id,
        txId: transaction_id,
        fcn: func_name,
        //  args: [arg1_f1]
        args: [arg1_f3]
       };
      console.log("Calling One GetDetailsByApplicantId Function..");
      return channel.queryByChaincode(request);
    }
    if(func_name == "ViewApplication"){
    	const request = {
    			chaincodeId: options.chaincode_id,
    			txId: transaction_id,
    			fcn: func_name,
    			args: []
    			//args: [arg1_f4]
        };
		console.log("Calling One ViewApplication Function..");
        return channel.queryByChaincode(request);
    }
    
    if(func_name == "ViewProperties"){
        const request = {
			chaincodeId: options.chaincode_id,
			txId: transaction_id,
			fcn: func_name,
			args: []
			//args: [arg1_f5]
        
         };
         console.log("Calling One ViewProperties Function..");
         return channel.queryByChaincode(request);
	}
    if(func_name == "GetDetailsByPropertyId"){
        const request = {
			chaincodeId: options.chaincode_id,
			txId: transaction_id,
			fcn: func_name,
			//  args: []
			 args: [arg1_f5]
        
         };
         console.log("Calling One GetDetailsByPropertyId Function..");
         return channel.queryByChaincode(request);
	}
	if(func_name == "getReceiverDetails"){
        const request = {
			chaincodeId: options.chaincode_id,
			txId: transaction_id,
			fcn: func_name,
			//  args: [arg1_f1]
			 args: [arg1_f6]
        
         };
         console.log("Calling One getReceiverDetails Function..");
         return channel.queryByChaincode(request);
	}
    }).then((query_responses) => {
	console.log("returned from query:"+query_responses);
	if (!query_responses.length) {
		console.log("No payloads were returned from query");
    } else {
    	console.log("Query result count = ", query_responses.length)
    }
    if (query_responses[0] instanceof Error) {
    	console.error("error from query = ", query_responses[0]);
    }
    console.log("Response is ", query_responses[0].toString());
    reply.send(util.format(query_responses[0].toString('utf8')));
    }).catch((err) => {
    	console.error("Caught Error", err);
    	reply.send ("Caught Error", err);
    });
}
