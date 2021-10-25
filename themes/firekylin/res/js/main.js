window.addEventListener('load', function(){
  if (window.location.pathname.indexOf('/post') === 0) {
    let ctimeStr = document.querySelector('.ctime').innerText;
    let ctime = new Date(ctimeStr);
    let diff = parseInt((new Date() - ctime) / 1000 / 60 / 60 / 24);
    if (diff > 60) {
      let tipNode = document.createElement("div");
      tipNode.style.border = "1px dashed #CCC";
      tipNode.style.padding = "4px 10px";
      tipNode.innerText = "It has been " + diff +  " days since the article was published, the article may have expired.";
      document.querySelector('.entry-content').parentNode.insertBefore(tipNode, document.querySelector('.entry-content'))
    }
  }
});

// 行号和高亮行处理 @xuexb
 var hljs = {
    $code: document.querySelectorAll('pre code'),
    hasClass: function (ele, cls) {
      return ele.className.match(new RegExp('(\\s|^)' + cls + '(\\s|$)'));
    },
    addClass: function (ele, cls) {
      if (!hljs.hasClass(ele, cls)) {
        ele.className += ' ' + cls;
      }
    },
    removeClass: function (ele, cls) {
      if (hljs.hasClass(ele, cls)) {
        ele.className = ele.className.replace(new RegExp('(\\s|^)' + cls + '(\\s|$)'), ' ');
      }
    }
  };

  /**
   * 使用数据生成hash
   *
   * @param  {Object} data 数据
   * @param {number} data.index 代码块位置, 以1开始
   * @param {number} data.start 行号开始
   * @param {number} data.end 行号结束
   *
   * @return {string}
   */
  hljs.stringHash = function (data) {
    var hash = '';
    if (data.index > 1) {
      hash += data.index + '-';
    }
    hash += 'L' + data.start;
    if (data.end && data.end > data.start) {
      hash += '-' + 'L' + data.end;
    }
    return hash;
  };

  /**
   * 解析hash为数据
   *
   * @return {Object} {index: 当前代码块位置, 以1开始,  start: 行号开始,  end: 结束位置}
   */
  hljs.parseHash = function () {
    var parse = location.hash.substr(1).match(/((\d+)-)?L(\d+)(-L(\d+))?/);

    if (!parse) {
      return null;
    }

    return {
      index: parseInt(parse[2], 10) || 1,
      start: parseInt(parse[3], 10) || 1,
      end: parseInt(parse[5], 10) || parseInt(parse[3], 10) || 1
    }
  };

  /**
   * 标记行颜色并跳转
   */
  hljs.mark = function (go) {
    var hash = hljs.parseHash();
    if (!hash || !hljs.$code || !hljs.$code[hash.index - 1]) {
      return;
    }

    var $li = hljs.$code[hash.index - 1].querySelectorAll('li');
    for (var i = hash.start - 1; i < hash.end; i++) {
      if ($li[i]) {
        hljs.addClass($li[i], 'mark');
      }
    }

    // 如果需要定位且元素存在
    if (go && $li && $li[0]) {
      setTimeout(function () {
        window.scrollTo(0, getRect($li[0]).top - 50);
      });
    }
  };

  /**
   * 移除所有高亮行号
   */
  hljs.removeMark = function () {
    [].slice.call(document.querySelectorAll('pre code li.mark')).forEach(function (elem) {
      hljs.removeClass(elem, 'mark');
    });
  };

  /**
   * 初始化
   */
  hljs.init = function () {
    [].slice.call(hljs.$code).forEach(function (elem, i) {
      // 输出行号, -1是为了让最后一个换行忽略
      var lines = elem.innerHTML.split(/\n/).slice(0, -1);
      var html = lines.map(function (item, index) {
        return '<li><span class="line-num" data-line="' + (index + 1) + '"></span>' + item + '</li>';
      }).join('');
      html = '<ul>' + html + '</ul>';

      // 输出语言
      if (lines.length > 3 && elem.className.match(/lang-(\w+)/) && RegExp.$1 !== 'undefined') {
        html += '<b class="name">' + RegExp.$1 + '</b>';
      }

      elem.innerHTML = html;

      hljs.addClass(elem, 'firekylin-code');

      // 绑定点击高亮行事件
      elem.addEventListener('click', function (event) {
        // 小小的委托
        if (!event.target || !hljs.hasClass(event.target, 'line-num')) {
          return;
        }

        // 如果是区间
        if (event.shiftKey) {
          var hash = hljs.parseHash();
          hash.newIndex = i + 1;
          hash.current = event.target.getAttribute('data-line');
          if (hash.index !== hash.newIndex - 0) {
            hash.index = hash.newIndex;
            hash.start = hash.current;
            hash.end = 0;
          } else {
            if (hash.current > hash.start) {
              hash.end = hash.current;
            } else {
              hash.end = hash.start;
              hash.start = hash.current;
            }
          }
          location.hash = hljs.stringHash(hash);
        } else {
          location.hash = hljs.stringHash({
            index: i + 1,
            start: event.target.getAttribute('data-line')
          });
        }
      });
    });
  };

  hljs.init();
  window.addEventListener('load', function () {
    hljs.mark(true);
  });
  window.addEventListener('hashchange', function () {
    hljs.removeMark();
    hljs.mark();
  });