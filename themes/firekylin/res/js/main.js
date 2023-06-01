window.addEventListener('load', function () {
  if (window.location.pathname.indexOf('/post') === 0) {
    for (let ele of document.querySelectorAll('.entry-content a:not(.toc a), .post-info a')) {
      ele.setAttribute('target', '_blank');
    }
  }
});

document.addEventListener('DOMContentLoaded', function () {
  let io = new IntersectionObserver(entries => {
    entries.forEach(entry => {
      if (entry.intersectionRatio > 0) {
        let img = entry.target;
        io.unobserve(img);

        const imgObj = new Image()
        imgObj.onload = function() {
          img.src = img.dataset.src;
          img.style.filter = '';
        }
        imgObj.src = img.dataset.src;
      }
    });
  }, { rootMargin: '300px' });

  document.querySelectorAll('img[data-src]').forEach(function (img) {
    img.style.filter = 'opacity(0.2)';
    io.observe(img);
  });
});