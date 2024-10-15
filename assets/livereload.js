conn = new WebSocket("ws://" + document.location.host + "/ws");
let time = 2000
conn.onmessage = function(event)  {
  if (parseInt(event.data) == event.data) {
    time = parseInt(event.data)
  }
  console.log("Time for refresh: " + time)
}
conn.onclose = function (evt) {  
  console.log("Connection Closed")  
  setTimeout(function () {  
    location.reload();  
  }, time);  
};
