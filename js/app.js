$(".send-message-button").click(function(e) {
    var msg = $(".send-message").val();
    e.preventDefault();
    $.post("/send", {
        message: msg
    }).success(function() {
        $(".send-message").val('');
    });
});

var m = $("#message");
var s = new EventSource("/sse");
s.onmessage = function(e) {
  m.text(e.data);
};