//*** (IMPORTS STARTS ) ****
var update = require('./invoke.js');
var query = require('./query.js');
//***  IMPORTS ENDS ****
exports.invoke = function (request,reply) {
	console.log('the function='+request.body.header.event_const);
	var header = request.body.header.event_const;
	var jsonblob = request.body.jsonblob;
	console.log("Header value:"+header);
	var fnName;
	if(header === "CNST_CREATEPROSPECT"){
		console.log("In route js entered and stored CreateProspect function");
		fnName = "CreateProspect";
	}
	if(header === "CNST_CREATEAPPLICATION"){
        console.log("In route js entered and stored CreateApplication function");
        fnName = "CreateApplication";
	}
	if(header === "CNST_UPDATEVALUERBYPROPERTYID"){
		console.log("In route js entered and stored UpdateValuerByPropertyId function");
		fnName = "UpdateValuerByPropertyId";
	}
	if(header == "CNST_UPDATESOLICITORBYPROPERTYID"){
        console.log("In route js entered UpdateSolicitorByPropertyId function");
        var fnName = "UpdateSolicitorByPropertyId";
    }
	if(header == "CNST_UPDATEUNDERWRITERBYAPPLICANTID"){
        console.log("In route js entered UpdateUnderwriterByApplicantId function");
        var fnName = "UpdateUnderwriterByApplicantId";
    }
	if(header == "CNST_QUALITYCHECK_EVENT"){
        console.log("In route js entered Invoice_QualityCheck function");
        var fnName = "Invoice_QualityCheck";
    }
	if(header == "CNST_ADJUSTMENT_EVENT"){
        console.log("In route js entered Invoice_Adjustment function");
        var fnName = "Invoice_Adjustment";
    }
	if(header == "CNST_PAYMENTFROMBANKERTOSUPPLIER_EVENT"){
        console.log("In route js entered ProcessPayment_BankerToSupplier function");
        var fnName = "ProcessPayment_BankerToSupplier";
    }
	if(header == "CNST_PAYMENTFROMOEMTOSUPPLIER_EVENT"){
        console.log("In route js entered CollectPayment_OEMToSupplier function");
        var fnName = "CollectPayment_OEMToSupplier";
    }
	if(header == "CNST_PAYMENTFROMSUPPLIERTOBANKER_EVENT"){
        console.log("In route js entered CollectPayment_SupplierToBanker function");
        var fnName = "CollectPayment_SupplierToBanker";
    }
	console.log("Function name at Routes layer:"+fnName);
	update.invokeSDK(fnName,jsonblob,reply);
}

// ***  ( QUERY SDK METHODS STARTS ) ***//      
exports.queryGetDetailsByProspectId = function (request,reply) {
	var fnName = "GetDetailsByProspectId";
	query.querySDK(fnName, request, reply);
}
exports.queryViewProspect = function (request,reply) {
  var fnName = "ViewProspect";
  query.querySDK(fnName, request, reply);
}
exports.queryGetDetailsByApplicantId = function (request,reply) {
	var fnName = "GetDetailsByApplicantId";
	query.querySDK(fnName, request, reply);
}
exports.queryViewApplication = function (request,reply) {
      var fnName = "ViewApplication";
      query.querySDK(fnName, request, reply);
}
exports.queryViewProperties = function (request,reply) {
    var fnName = "ViewProperties";
    query.querySDK(fnName, request, reply);
    }
exports.queryGetDetailsByPropertyId = function (request,reply) {
    var fnName = "GetDetailsByPropertyId";
    query.querySDK(fnName, request, reply);
    }
exports.querygetReceiverDetails = function (request,reply) {
    var fnName = "getReceiverDetails";
    query.querySDK(fnName, request, reply);
    }