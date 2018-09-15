//Modifiedd
var express = require('express');
var path = require('path');
var favicon = require('serve-favicon');
var logger = require('morgan');

var cookieParser = require('cookie-parser');
var bodyParser = require('body-parser'); //for fileupload
var fs = require('fs'); //for fileupload
var formidable = require('formidable');
var routes = require('./routes/index');
var users = require('./routes/user');
var route = require('./route.js');

var items =[];
var app = express();
app.use(bodyParser.json());       // to support JSON-encoded bodies
app.use(bodyParser.urlencoded({     // to support URL-encoded bodies
  extended: true
}));
var bodyParser = require('body-parser').json();
var env = process.env.NODE_ENV || 'development';

var OEM_Dashboard =  require('./routes/OEM_Dashboard');
var OEM_inwardGoods =  require('./routes/OEM_inwardGoods');
var OEM_viewDetailedInvoice =  require('./routes/OEM_viewDetailedInvoice');
var Supplier_viewDetailedInvoice =  require('./routes/Supplier_viewDetailedInvoice');
var Financer_viewDetailedInvoice =  require('./routes/Financer_viewDetailedInvoice');
var newPurchaseOrder =  require('./routes/newPurchaseOrder');
var viewPurchaseOrder =  require('./routes/viewPurchaseOrder');
var viewInvoice =  require('./routes/viewInvoice');
var viewInvoicesupplier =  require('./routes/viewInvoicesupplier');
var verifyGoods =  require('./routes/verifyGoods');
var qaAcceptReject =  require('./routes/qaAcceptReject');
var facilitator =  require('./routes/facilitator');
var supplierDashboard =  require('./routes/supplierDashboard');
var supplierViewPO =  require('./routes/supplierViewPO');
var oemDispatch =  require('./routes/supplierDispatch');
var oem_poview1 =  require('./routes/oem_poview');
var colPayFromOEM_Supplier =  require('./routes/colPayFromOEM_Supplier');
var viewInvoices_Financer =  require('./routes/viewInvoices_Financer');
var processPayToSupplier =  require('./routes/processPayToSupplier');
var collectPayFromSupplier =  require('./routes/collectPayFromSupplier');
var acceptorrejectPO =  require('./routes/acceptorrejectPO');

app.locals.ENV = env;
app.locals.ENV_DEVELOPMENT = env === 'development';

// view engine setup

app.set('views', path.join(__dirname, 'views'));
app.set('view engine', 'jade');

// app.use(favicon(__dirname + '/public/img/favicon.ico'));
app.use(logger('dev'));

app.use(cookieParser());
app.use(express.static(path.join(__dirname, 'public')));

app.use('/', routes);

app.use('/OEM_Dashboard.html',OEM_Dashboard);
app.use('/OEM_inwardGoods.html',OEM_inwardGoods);
app.use('/OEM_viewDetailedInvoice.html',OEM_viewDetailedInvoice);
app.use('/Supplier_viewDetailedInvoice.html',Supplier_viewDetailedInvoice);
app.use('/Financer_viewDetailedInvoice.html',Financer_viewDetailedInvoice);
app.use('/newPurchaseOrder.html',newPurchaseOrder);
app.use('/viewPurchaseOrder.html',viewPurchaseOrder);
app.use('/viewInvoice.html',viewInvoice);
app.use('/viewInvoicesupplier.html',viewInvoicesupplier);
app.use('/qaAcceptReject.html',qaAcceptReject);
app.use('/verifyGoods.html',verifyGoods);
app.use('/facilitator.html',facilitator);
app.use('/supplierDispatch.html',oemDispatch);
app.use('/oem_poview.html',oem_poview1);
app.use('/colPayFromOEM_Supplier.html',colPayFromOEM_Supplier);
app.use('/supplierDashboard.html',supplierDashboard);
app.use('/supplierViewPO.html',supplierViewPO);
app.use('/viewInvoices_Financer.html',viewInvoices_Financer);
app.use('/processPayToSupplier.html',processPayToSupplier);
app.use('/collectPayFromSupplier.html',collectPayFromSupplier);
app.use('/acceptorrejectPO.html',acceptorrejectPO);


app.get('/query/GetDetailsByProspectId/:arg1',function(req,res){
  //console.log("req: jax: "+req.params.arg1);
                 route.queryGetDetailsByProspectId(req,res);

})

app.get('/query/ViewProspect',function(req,res){
  //console.log("req: jax: "+req.params.arg1);
    route.queryViewProspect(req,res);

})

app.get('/query/GetDetailsByApplicantId/:arg1',function(req,res){
  //console.log("req: jax: "+req.params.arg1);
    route.queryGetDetailsByApplicantId(req,res);

})

app.get('/query/ViewApplication',function(req,res){
  //console.log("req: jax: "+req.params.arg1);
	route.queryViewApplication(req,res);

})
app.get('/query/ViewProperties',function(req,res){
  //console.log("req: jax: "+req.params.arg1);
	route.queryViewProperties(req,res);

})

app.get('/query/GetDetailsByPropertyId/:arg1',function(req,res){
  //console.log("req: jax: "+req.params.arg1);
                 route.queryGetDetailsByPropertyId(req,res);

})


app.post('/invoke',function(request,reply){
	
	
   route.invoke(request,reply);
   // console.log("response from ajax:"+response);
           

})




//**** ( FILE UPLOAD STARTS ) ******
app.post('/upload', function(req, res){

	  // create an incoming form object
	  var form = new formidable.IncomingForm();

	  // specify that we want to allow the user to upload multiple files in a single request
	  form.multiples = true;

	  // store all uploads in the /uploads directory
	  form.uploadDir = path.join(__dirname, '/public/uploads');
	  console.log("jax pan type:"+ form.jax);
	//var id = file.name;
	  // every time a file has been uploaded successfully,
	  // rename it to it's orignal name
	  form.on('file', function(field, file) {
		//var id = file.name; 
		console.log(file.name);
		uid = Date.now()
		file.name = uid+"_"+file.name;
		filename = file.name;
	    fs.rename(file.path, path.join(form.uploadDir, file.name));
	  });

	  // log any errors that occur
	  form.on('error', function(err) {
	    console.log('An error has occured: \n' + err);
	  });

	  // once all the files have been uploaded, send a response to the client
	  form.on('end', function() {
	    res.end(filename);
	  });

	  // parse the incoming request containing the form data
	  form.parse(req);

	});
//**** ( FILE UPLOAD ENDS ) ******

// Array storage for multiple items
app.post('/addtocart', function (req, response){
	const itemId = req.body.par_itemId;
	const itemName = req.body.par_itemName;
	const po_id = req.body.par_poid;
    const quantity = req.body.par_quantity;
    const unitPrice = req.body.par_unitPrice;
    const totalPrice = req.body.par_totalPrice;
	var itemObj;
	//console.log(itemId+" "+quantity+" "+unitPrice+" "+totalPrice);
	
	itemObj = {
	    "item_id":itemId,
	    "item_name":itemName,
	    "po_id":po_id,
	    "item_orderedQuantity":quantity,
	    "unitPrice":unitPrice
	    
	}
	console.log(itemObj+"check");
	items.push(itemObj);
	console.log("xxXXXxx");
	console.log(items);
	response.send(items);
	console.log(JSON.stringify(items));
	//response.send(JSON.stringify(item2));
	console.log("done");
	})
//Ends
	//query starts
	app.get('/getItems',function(req,res){
		console.log("requested get items"+items);
		res.send(items)
	});
	//query ends
app.post('/clearcart',function(req,res){
	console.log("requested get items"+items);
	items = [];
	res.send(items)
});
//delete	
app.post('/deletefromcart',function(req,res){
	const itemIdDel = req.body.par_itemId;
	console.log(itemIdDel+"'ll be deleted");
	console.log(items);
	for (i=0;i<items.length;i++){
		console.log(items[i].itemId);
		if(items[i].itemId == itemIdDel){
			console.log("found similarity zzzssshhh");
			var delIndex = items[i];
			console.log(delIndex+'indexValue');
			console.log(items[i]);
			items.splice(i,1);
			
			console.log(itemIdDel+"  deleted successfully");
			break;
			
		}
		else {
			console.log("found NO similarity zzzssshhh")
		}
	}
	console.log('out of loop');
	console.log(item2);
	res.send(item2)
	/*var index = item2.indexOf(itemId);
	console.log(index+"index val!!!");*/
	
})
	
//pretextZZZZZZZZZZZZ
if (app.get('env') === 'development') {
    app.use(function(err, req, res, next) {
        res.status(err.status || 500);
        res.render('error', {
            message: err.message,
            error: err,
            title: 'error'
        });
    });
}





/// catch 404 and forward to error handler
app.use(function(req, res, next) {
    var err = new Error('Not Found');
    err.status = 404;
    next(err);
});
// production error handler
// no stacktraces leaked to user
app.use(function(err, req, res, next) {
    res.status(err.status || 500);
    res.render('error', {
        message: err.message,
        error: {},
        title: 'error'
    });
});

app.set('port', process.env.PORT || 3000);

var server = app.listen(app.get('port'), function() {

  console.log('Express server listening on port ' + server.address().port);

});


module.exports = app;
