$(".send-message-button").click(function(e) {
    var msg = $(".send-message").val();
    e.preventDefault();
    $.post("/send", {
        message: msg
    }).success(function() {
        $(".send-message").val('');
    });
});

var $d = $("#datetime");
var $l = $("#event-log");
var s = new EventSource("/sse");

// handle datetimes
s.addEventListener("datetime", function(e) {
  $d.text(e.data)
});

// handle pull_request
s.addEventListener("pullrequest", function(e) {
  $l.append(
    $('<li>')
      .append("Pull request from " + e.data))
});

// handle plain messages
s.onmessage = function(e) {
  $l.append(
    $('<li>')
      .append("Message: " + e.data))
};

