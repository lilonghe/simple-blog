window.addEventListener('load', function () {
  if (window.location.pathname.indexOf('/post') === 0) {
    for (let ele of document.querySelectorAll('.entry-content a:not(.toc a), .post-info a')) {
      ele.setAttribute('target', '_blank');
    }
  }
});

document.addEventListener('DOMContentLoaded', function () {
  const cdnHost = 'https://cdn.lilonghe.net';
  let io = new IntersectionObserver(entries => {
    entries.forEach(entry => {
      if (entry.intersectionRatio > 0) {
        let img = entry.target;
        io.unobserve(img);

        const imgObj = new Image()

        let src = img.dataset.src;
        if (src.startsWith('/static/upload')) {
          src = cdnHost + src
        } else if (img.dataset.src.startsWith('https://note.lilonghe.net/static/upload')) {
          src = src.replace('https://note.lilonghe.net', cdnHost)
        }
        
        imgObj.onload = function() {
          img.src = src;
          img.style.filter = '';
        }
        
        imgObj.src = src;
      }
    });
  }, { rootMargin: '300px' });

  document.querySelectorAll('img[data-src]').forEach(function (img) {
    img.style.filter = 'opacity(0.2)';
    io.observe(img);
  });
});
