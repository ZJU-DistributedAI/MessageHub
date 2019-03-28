var express = require('express');
var router = express.Router();

function completeRes(msg, code){
    var response = {
        msg: msg,
        code: code,
    };
    return JSON.stringify(response);
}

/* GET users listing. */
router.get('/', function(req, res, next) {
    res.send('respond with a resource');
});

module.exports = router;

router.get('/register', function (req, res) {
    res.sendFile( __dirname + "/pages/" + "register.html" );
});

router.get('/login', function (req, res) {
    res.sendFile( __dirname + "/pages/" + "login.html" );
});

router.post('/checklogin', function (req, res) {
    var username = req.body.username;
    var password = req.body.password;
    var userType = req.body.userType;

    if(username === 'dcd' && password === '123456') {
        if(userType == 0) {
            res.end("/dataclient/index");
        }
        if(userType == 1) {
            res.end("/modelclient/index");
        }
        if(userType == 2) {
            res.end("/computingclient/index");
        }
    }
    else {
        res.end("/user/login");
    }
});

router.post('/createwallet', function (req, res) {
    var response;
    var password = req.body.password;
    if(password == undefined || password == '') {
        response = completeRes("未找到password参数", 201);
    }
    else {
        var account = web3.eth.personal.newAccount(password);
        // account 是一个promise对象
        account.then(function(data){
            response = completeRes(data,200);
            res.end(response);
        });
    }
});

router.get('/createwalletpage', function (req, res) {
    res.sendFile( __dirname + "/pages/" + "createwallet.html" );
});

router.get('/downloadtool', function (req, res) {
    res.sendFile( __dirname + "/pages/" + "download_tool.html" );
});