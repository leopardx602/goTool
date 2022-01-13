$(document).ready(() => {
	var ws = new WebSocket("ws://127.0.0.1:5000/connect");
	ws.onopen = function(evt) {
		console.log("OPEN")
	}
	ws.onclose = function(evt) {
		console.log("CLOSE")
		ws = null;
	}
	ws.onmessage = function(evt) {
		console.log("RESPONSE: " + evt.data)
	}
	ws.onerror = function(evt) {
		console.log("ERROR: " + evt.data)
	}

	$("#send").click(()=>{
		var msg = $("#msg").val()
		ws.send(msg);
	})
});