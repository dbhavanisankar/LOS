//Modifiedd
var express = require('express');
var path = require('path');
var favicon = require('serve-favicon');
var logger = require('morgan');

var cookieParser = require('cookie-parser');
var bodyParser = require('body-parser'); //for fileupload
var fs = require('fs'); //for fileupload
var formidable = require('formidable');
var routes = require('./routes/user-login');


var items =[];
var app = express();
app.use(bodyParser.json());       // to support JSON-encoded bodies
app.use(bodyParser.urlencoded({     // to support URL-encoded bodies
  extended: true
}));
var bodyParser = require('body-parser').json();
var env = process.env.NODE_ENV || 'development';

var userlogin =  require('./routes/user-login');
var prequalificationuserviewprospect =  require('./routes/prequalification-user-view-prospect');
var bankerviewprospect =  require('./routes/banker-view-prospect');
var solictorviewprospect =  require('./routes/solictor-view-prospect');


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

app.use('/user-login.html',userlogin);
app.use('/prequalification-user-view-prospect.html',prequalificationuserviewprospect);
app.use('/prequalification-user-new-prospect.html',prequalificationusernewprospect);

app.use('/banker-view-prospect.html',bankerviewprospect);
app.use('/solictor-view-prospect.html',solictorviewprospect);


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
	  console.log("path :"+path.join(__dirname, '/public/uploads'));
	 // console.log("jax pan type:"+ form.jax);
	//var id = file.name;
	  // every time a file has been uploaded successfully,
	  // rename it to it's orignal name
	  form.on('file', function(field, file) {
		//var id = file.name; 
		console.log("file.name :" +file.name);
		uid = Date.now();
		file.name = uid+"_"+file.name;
		filename = file.name;
	    fs.rename(file.path, path.join(form.uploadDir, file.name));
	 

// priyanka code starts here

 console.log( "filename :"+filename);
 
var crypto = require('crypto');
var hash = crypto.createHash('sha256');
var md5sum = crypto.createHash('md5'); 	

var algo = 'md5';
var shasum = crypto.createHash(algo);
//var fileN = path.join(__dirname, '/public/uploads',filename);

//generating hash for the filename
filehash = crypto.createHash('md5').update(filename).digest('hex');
console.log('Generated Hash for filename :' + filehash);


//priyanka code ends here



	  });

	  // log any errors that occur
	  form.on('error', function(err) {
	    console.log('An error has occured: \n' + err);
	  });

	  // once all the files have been uploaded, send a response to the client
	  form.on('end', function() {
	   // res.end(filename);
		  res.end(filehash +","+filename);
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
			
			console.log(itemIdDel+" deleted successfully");
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
