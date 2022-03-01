const post_url = 'http://127.0.0.1:13033/record';

function publish() {
  var inst = new mdui.Dialog('#dialog');
  $.ajax({
    type: 'post',
    url: post_url,
    data: JSON.stringify({
      title: $('#title').val(),
      content: $('#content').val()
    }),
    success: function (data) {
      inst.open();
      let short_link = data['data']['short_link'];
      let msg = "Your share link: " + short_link;
      $('#dialog-content').text(msg);
    }
  });
}

