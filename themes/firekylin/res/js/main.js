window.addEventListener('load', function () {
  if (window.location.pathname.indexOf('/post') === 0) {
    for (let ele of document.querySelectorAll('.entry-content a:not(.toc a), .post-info a')) { 
      ele.setAttribute('target', '_blank');
    }
  }
});