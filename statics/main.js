
// starts a greyhound websocket, 
// @log is a function pointer to a method that takes a string, for
// logging purposes.  
//
// the query method is the prefered way to query
// greyhound for data. One sends an action, data, and a callback
// function which takes an array of strings as a response.
function startGHWebSocket(log) {
    log("Starting websockets...");
    ws = new WebSocket("ws://localhost:8081/socket");
    ws.log = log;
    ws.binaryType = "blob";
    ws.onopen = function () {
        this.log("socket opened");
    };
    ws.onclose = function (e) {
        this.log("socket closed");
    };
    ws.onmessage = function (e) {
        if (this.callback) {
            this.callback($.parseJSON(e.data));
        }
    };
    ws.query = function(action, data, callback) {
        m = JSON.stringify({action: action, queryData: data});
        this.callback = callback;
        this.send(m);
    };
    return ws;
}
