var createError = require('http-errors');
var express = require('express');
var path = require('path');
var cookieParser = require('cookie-parser');
var logger = require('morgan');

var indexRouter = require('./routes/index');
var usersRouter = require('./routes/users');
var dataClientRouter = require('./routes/dataclient');
var modelClientRouter = require('./routes/modelclient');
var computingClientRouter = require('./routes/computingclient');
var Web3 = require('web3');
var Tx = require('ethereumjs-tx');

var app = express();

// view engine setup
app.set('views', path.join(__dirname, 'views'));
app.set('view engine', 'pug');

app.use(logger('dev'));
app.use(express.json());
app.use(express.urlencoded({ extended: false }));
app.use(cookieParser());
app.use(express.static(path.join(__dirname, 'public')));

app.use('/', indexRouter);
app.use('/user', usersRouter);
app.use('/dataclient', dataClientRouter);
app.use('/modelclient', modelClientRouter);
app.use('/computingclient', computingClientRouter);


// catch 404 and forward to error handler
app.use(function(req, res, next) {
    next(createError(404));
});

// error handler
app.use(function(err, req, res, next) {
    res.sendFile(__dirname + '/routes/pages/404.html');
});

// web3 define
var web3 = undefined;
if (typeof web3 !== 'undefined') {
    console.log("web3 !== 'undefined'");
    web3 = new Web3(web3.currentProvider);
} else {
    // set the provider you want from Web3.providers
    web3 = new Web3(new Web3.providers.HttpProvider("http://localhost:8545"));
}


global.web3 = web3;
global.Tx = Tx;
global.DataTransactionTo = "0xb1012acadcdf1c462c0f6b665e84d892e1f4bcf9";
global.ModelTransactionTo = "0x7c5e1facd0af1c562c2d03eeade47e1c7e695ca6";
global.ComputingTransactionTo = "0x25127a6a0c6dcdd431425aa1929f93e339039ed1";
module.exports = app;