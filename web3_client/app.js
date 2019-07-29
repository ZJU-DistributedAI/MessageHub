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
    web3 = new Web3(new Web3.providers.HttpProvider("http://47.52.163.119:8545"));
}


global.web3 = web3;
global.Tx = Tx;
global.DataTransactionTo = "0x7c2387f88aca12aca97a938b88f6d0512b9195e2";
global.ModelTransactionTo = "0xf6eb003cf5fcdfdaedcc4540028f9f305f1011aa";
global.ComputingTransactionTo = "0x25127a6a0c6dcdd431425aa1929f93e339039ed1";
module.exports = app;