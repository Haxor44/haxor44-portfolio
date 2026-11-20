(function () {
  var root = document.documentElement;
  var reduceMotion = window.matchMedia('(prefers-reduced-motion: reduce)').matches;
  root.classList.add('js');

  var year = document.getElementById('year');
  if (year) year.textContent = new Date().getFullYear();

  /* ---------- Theme toggle ---------- */
  var toggle = document.getElementById('theme-toggle');
  if (toggle) {
    toggle.addEventListener('click', function () {
      var next = root.getAttribute('data-theme') === 'dark' ? 'light' : 'dark';
      root.setAttribute('data-theme', next);
      try { localStorage.setItem('theme', next); } catch (e) {}
    });
  }

  /* ---------- Header border on scroll ---------- */
  var header = document.querySelector('.site-header');
  function onScroll() { if (header) header.classList.toggle('is-scrolled', window.scrollY > 8); }
  window.addEventListener('scroll', onScroll, { passive: true });
  onScroll();

  /* ---------- Scroll reveal ---------- */
  // Only whole sections fade in — never individual cards — and only ones that
  // start below the fold. A fallback timer reveals everything regardless, so
  // content can never get stuck invisible (fast scrolling, anchors, printing).
  var targets = document.querySelectorAll('.section, .cta-band');
  if ('IntersectionObserver' in window && !reduceMotion) {
    var revealAll = function () {
      document.querySelectorAll('.reveal').forEach(function (el) { el.classList.add('is-visible'); });
    };
    var io = new IntersectionObserver(function (entries) {
      entries.forEach(function (entry) {
        if (entry.isIntersecting) {
          entry.target.classList.add('is-visible');
          io.unobserve(entry.target);
        }
      });
    }, { threshold: 0, rootMargin: '0px 0px -10% 0px' });
    targets.forEach(function (el) {
      if (el.getBoundingClientRect().top < window.innerHeight) return; // already on screen
      el.classList.add('reveal');
      io.observe(el);
    });
    window.addEventListener('beforeprint', revealAll);
    setTimeout(revealAll, 2000);
  }

  /* ---------- Starfield ---------- */
  var canvas = document.getElementById('starfield');
  if (!canvas || !canvas.getContext) return;
  var ctx = canvas.getContext('2d');
  var stars = [];
  var dpr = Math.min(window.devicePixelRatio || 1, 2);
  var w, h;

  function starColor() {
    return getComputedStyle(root).getPropertyValue('--star').trim() || '229, 231, 235';
  }

  function resize() {
    w = window.innerWidth;
    h = window.innerHeight;
    canvas.width = w * dpr;
    canvas.height = h * dpr;
    ctx.setTransform(dpr, 0, 0, dpr, 0, 0);
    var count = Math.round((w * h) / 9000);
    stars = [];
    for (var i = 0; i < count; i++) {
      stars.push({
        x: Math.random() * w,
        y: Math.random() * h,
        r: Math.random() * 1.1 + 0.2,
        a: Math.random() * 0.6 + 0.15,
        s: Math.random() * 0.02 + 0.004,
        p: Math.random() * Math.PI * 2,
        warm: Math.random() < 0.08
      });
    }
  }

  function draw(t) {
    var rgb = starColor();
    ctx.clearRect(0, 0, w, h);
    for (var i = 0; i < stars.length; i++) {
      var s = stars[i];
      var alpha = reduceMotion ? s.a : s.a * (0.55 + 0.45 * Math.sin(s.p + t * s.s * 0.06));
      ctx.beginPath();
      ctx.arc(s.x, s.y, s.r, 0, Math.PI * 2);
      ctx.fillStyle = s.warm ? 'rgba(249, 115, 22,' + alpha + ')' : 'rgba(' + rgb + ',' + alpha + ')';
      ctx.fill();
    }
    if (!reduceMotion) requestAnimationFrame(draw);
  }

  resize();
  window.addEventListener('resize', function () { resize(); if (reduceMotion) draw(0); });
  requestAnimationFrame(draw);
})();
