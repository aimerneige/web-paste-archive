const get_url = 'http://127.0.0.1:13033/record';

var access_path = window.location.pathname;
$.ajax({
  url: get_url + access_path,
  type: 'GET',
  success: function (data) {
    let title = data['data']['title'];
    $('#title').text(title);
    let content = data['data']['content'];
    $('#content').text(content);
  }
});
