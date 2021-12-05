window.addEventListener('load', function () {
  if (window.location.pathname.indexOf('/post') === 0) {
    let ctimeStr = document.querySelector('.ctime').innerText;
    let ctime = new Date(ctimeStr);
    let diff = parseInt((new Date() - ctime) / 1000 / 60 / 60 / 24);
    if (diff > 60) {
      let tipNode = document.createElement("div");
      tipNode.style.border = "1px dashed #CCC";
      tipNode.style.padding = "4px 10px";
      if (window.navigator.language === 'zh-CN') {
        tipNode.innerText = "文章已发表 " + diff + " 天，内容可能已经过期。";
      } else {
        tipNode.innerText = "It has been " + diff + " days since the article was published, the article may have expired.";
      }
      document.querySelector('.entry-content').parentNode.insertBefore(tipNode, document.querySelector('.entry-content'))
    }
  }
  hljs.init();
});

var hljs = {
  $code: document.querySelectorAll('pre code'),
  addClass: function (ele, cls) {
    ele.className += ' ' + cls;
  },
};

hljs.init = function () {
  [].slice.call(hljs.$code).forEach(function (elem, i) {
    var lines = elem.innerHTML.trim().split(/[\r\n]+/);
    var html = lines.map(function (item, index) {
      return '<li><span class="line-num" data-line="' + (index + 1) + '"></span>' + item + '</li>';
    }).join('');
    html = '<ul>' + html + '</ul>';

    if (lines.length > 3 && elem.className.match(/lang-(\w+)/) && RegExp.$1 !== 'undefined') {
      html += '<b class="name">' + RegExp.$1 + '</b>';
    }

    elem.innerHTML = html;

    hljs.addClass(elem, 'firekylin-code');
  });
};